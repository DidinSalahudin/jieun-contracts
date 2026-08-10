package jobqueue

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/DidinSalahudin/jieun-contracts/go/events/envelope"
)

// EventPublisher writes result-notification events to Redis Streams —
// INTEGRATION.md §6.1's "Streams untuk melapor" half of the split (the
// other half, commands, is Client/Server in this package). One stream per
// event family is the intended usage (e.g. a "jieun.analysis" stream);
// this package doesn't prescribe stream naming, callers choose.
type EventPublisher struct {
	rdb *redis.Client
}

func NewEventPublisher(redisAddr string) *EventPublisher {
	return &EventPublisher{rdb: redis.NewClient(&redis.Options{Addr: redisAddr})}
}

func (p *EventPublisher) Close() error {
	return p.rdb.Close()
}

// Publish marshals env (INTEGRATION.md §6.4's envelope, already generated
// as envelope.EventEnvelope) into a single "payload" field on a new
// Streams entry, returning the Streams-assigned message ID.
func (p *EventPublisher) Publish(ctx context.Context, stream string, env envelope.EventEnvelope) (string, error) {
	data, err := env.Marshal()
	if err != nil {
		return "", fmt.Errorf("jobqueue: marshal event envelope: %w", err)
	}
	id, err := p.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{"payload": string(data)},
	}).Result()
	if err != nil {
		return "", fmt.Errorf("jobqueue: publish event to stream %q: %w", stream, err)
	}
	return id, nil
}
