package jobqueue

// TaskType is an asynq task type name — matches the "Nama task" column of
// INTEGRATION.md §6.3's catalog exactly (e.g. "analyze:start").
type TaskType string

const (
	TaskAnalyzeStart       TaskType = "analyze:start"
	TaskAnalyzeStage       TaskType = "analyze:stage"
	TaskRenderClip         TaskType = "render:clip"
	TaskPublishVideo       TaskType = "publish:video"
	TaskAnalyticsExport    TaskType = "analytics:export"
	TaskStoryboardGenerate TaskType = "storyboard:generate"
)

// Queue names — INTEGRATION.md §6.2.
const (
	QueuePublish = "publish"
	QueueRender  = "render"
	QueueAnalyze = "analyze"
	QueueDefault = "default"
)

// QueuePriorities is asynq's weighted-queue config — INTEGRATION.md §6.2's
// priority column, used directly as asynq.Config.Queues.
var QueuePriorities = map[string]int{
	QueuePublish: 6,
	QueueRender:  4,
	QueueAnalyze: 3,
	QueueDefault: 1,
}

// QueueForTask maps each REGISTERED task type to its queue —
// INTEGRATION.md §6.3's catalog. Only task types with a real payload
// contract (jieun-contracts/jobs/*.v1.json) are registered here; the
// catalog's other 6 entries belong to services that don't exist yet and
// get registered by whichever future task builds them (see this plan's
// Global Constraints for the full catalog and why).
var QueueForTask = map[TaskType]string{
	TaskAnalyzeStart:       QueueAnalyze,
	TaskAnalyzeStage:       QueueAnalyze,
	TaskRenderClip:         QueueRender,
	TaskPublishVideo:       QueuePublish,
	TaskAnalyticsExport:    QueueDefault,
	TaskStoryboardGenerate: QueueAnalyze,
}
