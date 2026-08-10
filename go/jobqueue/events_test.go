package jobqueue

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/DidinSalahudin/jieun-contracts/go/events/envelope"
)

func TestEventPublisher_Publish_WritesToStream(t *testing.T) {
	p := NewEventPublisher(testRedisAddr)
	defer p.Close()
	ctx := context.Background()
	streamName := "test-stream-" + t.Name()

	env := envelope.EventEnvelope{
		Event:      "jieun.analysis.completed",
		Version:    1,
		EventID:    "evt_test_1",
		OccurredAt: time.Now().UTC(),
		TraceID:    "trace_test_1",
		UserID:     "u_test_1",
		Data:       map[string]interface{}{"candidate_count": float64(3)},
	}

	msgID, err := p.Publish(ctx, streamName, env)
	if err != nil {
		t.Fatalf("Publish: %v", err)
	}
	if msgID == "" {
		t.Error("Publish returned empty message ID")
	}

	rdb := redis.NewClient(&redis.Options{Addr: testRedisAddr})
	defer rdb.Close()
	msgs, err := rdb.XRange(ctx, streamName, "-", "+").Result()
	if err != nil {
		t.Fatalf("XRange: %v", err)
	}
	if len(msgs) != 1 {
		t.Fatalf("len(msgs) = %d, want 1", len(msgs))
	}
	raw, ok := msgs[0].Values["payload"].(string)
	if !ok {
		t.Fatalf("msgs[0].Values[\"payload\"] is not a string: %#v", msgs[0].Values["payload"])
	}
	var got envelope.EventEnvelope
	if err := json.Unmarshal([]byte(raw), &got); err != nil {
		t.Fatalf("unmarshal published envelope: %v", err)
	}
	if got.EventID != "evt_test_1" {
		t.Errorf("got.EventID = %q, want evt_test_1", got.EventID)
	}

	rdb.Del(ctx, streamName)
}
