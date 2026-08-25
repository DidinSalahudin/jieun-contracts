package jobqueue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/DidinSalahudin/jieun-contracts/go/events/envelope"
)

// dedupTTL bounds how long a processed event_id is remembered — long
// enough to cover any realistic at-least-once redelivery window, not
// forever (see this plan's Global Constraints for why this is a Redis
// key with a TTL, not a Postgres table).
const dedupTTL = 24 * time.Hour

type EventMessage struct {
	ID       string
	Envelope envelope.EventEnvelope
}

// EventConsumer reads result-notification events from Redis Streams via a
// consumer group — INTEGRATION.md §6.1/§6.4. ReadGroup filters out
// already-processed event_ids internally (auto-acking them without
// surfacing them to the caller), so every EventMessage it returns is
// guaranteed new from the caller's point of view.
type EventConsumer struct {
	rdb *redis.Client
}

func NewEventConsumer(redisAddr string) *EventConsumer {
	return &EventConsumer{rdb: redis.NewClient(&redis.Options{Addr: redisAddr})}
}

func (c *EventConsumer) Close() error {
	return c.rdb.Close()
}

// EnsureGroup creates the consumer group (and the stream itself, if it
// doesn't exist yet) — safe to call every time a consumer starts up.
func (c *EventConsumer) EnsureGroup(ctx context.Context, stream, group string) error {
	err := c.rdb.XGroupCreateMkStream(ctx, stream, group, "0").Err()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		return fmt.Errorf("jobqueue: ensure consumer group: %w", err)
	}
	return nil
}

// readGroupBlock bounds how long ReadGroup waits for new entries on an
// empty stream — go-redis's zero value for XReadGroupArgs.Block sends
// Redis's own "BLOCK 0" (block forever), not "don't block" as the Go zero
// value might suggest. A caller alternating between multiple streams (e.g.
// jieun.render.completed and jieun.render.failed) would hang forever the
// moment it reached whichever one had zero entries, never returning to
// poll the others again — a real bug found live, T-124b.
const readGroupBlock = 100 * time.Millisecond

// ReadGroup reads up to count new entries for this consumer, decodes each
// as an EventEnvelope, and applies event_id-level idempotency: an entry
// whose event_id was already processed is acked immediately and excluded
// from the returned slice — TASKS.md T-040's "Konsumen event idempoten
// terhadap event_id". Returns promptly (within readGroupBlock) even when
// the stream has nothing new.
func (c *EventConsumer) ReadGroup(ctx context.Context, stream, group, consumerName string, count int64) ([]EventMessage, error) {
	streams, err := c.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumerName,
		Streams:  []string{stream, ">"},
		Count:    count,
		Block:    readGroupBlock,
	}).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("jobqueue: read consumer group: %w", err)
	}

	var out []EventMessage
	for _, s := range streams {
		for _, entry := range s.Messages {
			raw, ok := entry.Values["payload"].(string)
			if !ok {
				continue
			}
			var env envelope.EventEnvelope
			if err := json.Unmarshal([]byte(raw), &env); err != nil {
				continue
			}

			alreadyProcessed, err := c.wasProcessed(ctx, env.EventID)
			if err != nil {
				return nil, err
			}
			if alreadyProcessed {
				_ = c.Ack(ctx, stream, group, entry.ID)
				continue
			}

			if err := c.markProcessed(ctx, env.EventID); err != nil {
				return nil, err
			}
			out = append(out, EventMessage{ID: entry.ID, Envelope: env})
		}
	}
	return out, nil
}

func (c *EventConsumer) wasProcessed(ctx context.Context, eventID string) (bool, error) {
	n, err := c.rdb.Exists(ctx, "processed:"+eventID).Result()
	if err != nil {
		return false, fmt.Errorf("jobqueue: check processed event: %w", err)
	}
	return n > 0, nil
}

func (c *EventConsumer) markProcessed(ctx context.Context, eventID string) error {
	if err := c.rdb.Set(ctx, "processed:"+eventID, "1", dedupTTL).Err(); err != nil {
		return fmt.Errorf("jobqueue: mark event processed: %w", err)
	}
	return nil
}

// Ack acknowledges a message, removing it from the consumer group's
// pending entries list.
func (c *EventConsumer) Ack(ctx context.Context, stream, group, messageID string) error {
	if err := c.rdb.XAck(ctx, stream, group, messageID).Err(); err != nil {
		return fmt.Errorf("jobqueue: ack message: %w", err)
	}
	return nil
}
