package jobqueue

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

const maxRetryCount = 3

// Client wraps asynq.Client with this project's task-type registry and
// idempotency convention (INTEGRATION.md §6.3, §5.2).
type Client struct {
	ac *asynq.Client
}

func NewClient(redisAddr string) *Client {
	return &Client{ac: asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})}
}

func (c *Client) Close() error {
	return c.ac.Close()
}

// Enqueue submits payload as a task of the given registered type, using
// idempotencyKey as asynq's native task ID — asynq's Redis backend rejects
// a duplicate active/pending task sharing an ID, which is this project's
// idempotency_key enforcement (INTEGRATION.md §5.2). Retries max out at 3
// with asynq's built-in exponential backoff; an exhausted task moves to
// asynq's Archive automatically (this project's dead-letter queue — see
// this plan's Global Constraints for why no custom DLQ was built).
func (c *Client) Enqueue(ctx context.Context, taskType TaskType, idempotencyKey string, payload []byte) (*asynq.TaskInfo, error) {
	queue, ok := QueueForTask[taskType]
	if !ok {
		return nil, fmt.Errorf("jobqueue: task type %q is not registered", taskType)
	}
	task := asynq.NewTask(string(taskType), payload)
	info, err := c.ac.EnqueueContext(ctx, task,
		asynq.Queue(queue),
		asynq.TaskID(idempotencyKey),
		asynq.MaxRetry(maxRetryCount),
	)
	if err != nil {
		return nil, fmt.Errorf("jobqueue: enqueue %q: %w", taskType, err)
	}
	return info, nil
}
