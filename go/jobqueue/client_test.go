package jobqueue

import (
	"context"
	"testing"

	"github.com/hibiken/asynq"
)

const testRedisAddr = "localhost:16379"

func TestClient_Enqueue_RoutesToCorrectQueue(t *testing.T) {
	c := NewClient(testRedisAddr)
	defer c.Close()
	ctx := context.Background()

	info, err := c.Enqueue(ctx, TaskRenderClip, "test-idem-"+t.Name(), []byte(`{"hello":"world"}`))
	if err != nil {
		t.Fatalf("Enqueue: %v", err)
	}
	if info.Queue != QueueRender {
		t.Errorf("info.Queue = %q, want %q", info.Queue, QueueRender)
	}
	if info.Type != string(TaskRenderClip) {
		t.Errorf("info.Type = %q, want %q", info.Type, TaskRenderClip)
	}

	inspector := asynq.NewInspector(asynq.RedisClientOpt{Addr: testRedisAddr})
	defer inspector.Close()
	if err := inspector.DeleteTask(QueueRender, info.ID); err != nil {
		t.Logf("cleanup: delete task: %v", err)
	}
}

func TestClient_Enqueue_DuplicateIdempotencyKeyRejected(t *testing.T) {
	c := NewClient(testRedisAddr)
	defer c.Close()
	ctx := context.Background()
	idemKey := "test-idem-dup-" + t.Name()

	info1, err := c.Enqueue(ctx, TaskRenderClip, idemKey, []byte(`{}`))
	if err != nil {
		t.Fatalf("first Enqueue: %v", err)
	}

	_, err = c.Enqueue(ctx, TaskRenderClip, idemKey, []byte(`{}`))
	if err == nil {
		t.Error("second Enqueue with same idempotency key: want an error (duplicate task ID), got nil")
	}

	inspector := asynq.NewInspector(asynq.RedisClientOpt{Addr: testRedisAddr})
	defer inspector.Close()
	if err := inspector.DeleteTask(QueueRender, info1.ID); err != nil {
		t.Logf("cleanup: delete task: %v", err)
	}
}

func TestClient_Enqueue_UnregisteredTaskTypeFails(t *testing.T) {
	c := NewClient(testRedisAddr)
	defer c.Close()
	ctx := context.Background()

	_, err := c.Enqueue(ctx, TaskType("nonexistent:task"), "test-idem-"+t.Name(), []byte(`{}`))
	if err == nil {
		t.Error("Enqueue with unregistered task type: want an error, got nil")
	}
}
