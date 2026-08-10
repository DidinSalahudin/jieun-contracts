package jobqueue

import "github.com/hibiken/asynq"

// NewServer builds an asynq worker server configured with this project's
// queue priorities (INTEGRATION.md §6.2). Retry backoff is asynq's own
// default RetryDelayFunc, which is already exponential — no override
// needed to satisfy "backoff eksponensial" (this plan's Global
// Constraints explain why no custom DLQ is built either: an exhausted
// task moves to asynq's Archive automatically).
func NewServer(redisAddr string) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{Queues: QueuePriorities},
	)
}

// NewServeMux re-exports asynq.NewServeMux so callers only need to import
// internal/jobqueue, not asynq directly, to register handlers.
func NewServeMux() *asynq.ServeMux {
	return asynq.NewServeMux()
}
