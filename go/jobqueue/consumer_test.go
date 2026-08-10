package jobqueue

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/DidinSalahudin/jieun-contracts/go/events/envelope"
)

func TestEventConsumer_ReadGroup_DeliversPublishedEvent(t *testing.T) {
	// Unique per run — a fixed event_id would collide with the "processed"
	// dedup key (24h TTL) left behind by a previous run of this same test.
	suffix := newTestSuffix()
	streamName := "test-stream-" + t.Name() + "-" + suffix
	groupName := "test-group-" + t.Name() + "-" + suffix
	eventID := "evt_" + t.Name() + "-" + suffix

	pub := NewEventPublisher(testRedisAddr)
	defer pub.Close()
	cons := NewEventConsumer(testRedisAddr)
	defer cons.Close()
	ctx := context.Background()

	if err := cons.EnsureGroup(ctx, streamName, groupName); err != nil {
		t.Fatalf("EnsureGroup: %v", err)
	}

	env := envelope.EventEnvelope{
		Event: "jieun.render.completed", Version: 1, EventID: eventID,
		OccurredAt: time.Now().UTC(), TraceID: "trace_1", UserID: "u_1",
		Data: map[string]interface{}{},
	}
	if _, err := pub.Publish(ctx, streamName, env); err != nil {
		t.Fatalf("Publish: %v", err)
	}

	msgs, err := cons.ReadGroup(ctx, streamName, groupName, "consumer-1", 10)
	if err != nil {
		t.Fatalf("ReadGroup: %v", err)
	}
	if len(msgs) != 1 {
		t.Fatalf("len(msgs) = %d, want 1", len(msgs))
	}
	if msgs[0].Envelope.EventID != env.EventID {
		t.Errorf("msgs[0].Envelope.EventID = %q, want %q", msgs[0].Envelope.EventID, env.EventID)
	}

	if err := cons.Ack(ctx, streamName, groupName, msgs[0].ID); err != nil {
		t.Fatalf("Ack: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer rdb.Close()
	pending, err := rdb.XPending(ctx, streamName, groupName).Result()
	if err != nil {
		t.Fatalf("XPending: %v", err)
	}
	if pending.Count != 0 {
		t.Errorf("pending.Count = %d, want 0 (message was acked)", pending.Count)
	}

	rdb.Del(ctx, streamName)
	rdb.Del(ctx, "processed:"+eventID)
}

func TestEventConsumer_ReadGroup_SkipsAlreadyProcessedEventID(t *testing.T) {
	// Simulates at-least-once redelivery: the same event_id arrives twice
	// (e.g. two Streams entries carrying the same logical event, or a
	// redelivery after a crash before ack). The second delivery must not
	// appear in ReadGroup's results — TASKS.md T-040's "Konsumen event
	// idempoten terhadap event_id".
	suffix := newTestSuffix()
	streamName := "test-stream-" + t.Name() + "-" + suffix
	groupName := "test-group-" + t.Name() + "-" + suffix

	pub := NewEventPublisher(testRedisAddr)
	defer pub.Close()
	cons := NewEventConsumer(testRedisAddr)
	defer cons.Close()
	ctx := context.Background()

	if err := cons.EnsureGroup(ctx, streamName, groupName); err != nil {
		t.Fatalf("EnsureGroup: %v", err)
	}

	env := envelope.EventEnvelope{
		Event: "jieun.render.completed", Version: 1, EventID: "evt_dup_" + t.Name() + "-" + suffix,
		OccurredAt: time.Now().UTC(), TraceID: "trace_1", UserID: "u_1",
		Data: map[string]interface{}{},
	}
	// Same event_id published twice — two distinct Streams entries.
	if _, err := pub.Publish(ctx, streamName, env); err != nil {
		t.Fatalf("Publish 1: %v", err)
	}
	if _, err := pub.Publish(ctx, streamName, env); err != nil {
		t.Fatalf("Publish 2: %v", err)
	}

	msgs, err := cons.ReadGroup(ctx, streamName, groupName, "consumer-1", 10)
	if err != nil {
		t.Fatalf("ReadGroup: %v", err)
	}
	if len(msgs) != 1 {
		t.Fatalf("len(msgs) = %d, want 1 (second delivery of the same event_id must be filtered)", len(msgs))
	}

	rdb := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer rdb.Close()
	rdb.Del(ctx, streamName)
	rdb.Del(ctx, "processed:"+env.EventID)
}
