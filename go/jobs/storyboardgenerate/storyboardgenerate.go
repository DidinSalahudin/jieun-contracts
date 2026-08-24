// Package storyboardgenerate defines the storyboard:generate asynq task
// payload — hand-written (no JSON-schema source exists for this task
// yet, same as analyticsexport), api/API-STORYBOARD.md §5.
package storyboardgenerate

import "encoding/json"

func UnmarshalStoryboardGenerateTask(data []byte) (StoryboardGenerateTask, error) {
	var r StoryboardGenerateTask
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StoryboardGenerateTask) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StoryboardGenerateTask struct {
	JobID          string  `json:"job_id"`
	TraceID        string  `json:"trace_id"`
	UserID         string  `json:"user_id"`
	IdempotencyKey string  `json:"idempotency_key"`
	Payload        Payload `json:"payload"`
}

type Payload struct {
	StoryboardID  string       `json:"storyboard_id"`
	Idea          string       `json:"idea"`
	Options       Options      `json:"options"`
	PromptVersion string       `json:"prompt_version"`
	EnginePolicy  EnginePolicy `json:"engine_policy"`
	BudgetMaxIDR  int64        `json:"budget_max_idr"`
}

type Options struct {
	TargetDurationS int64  `json:"target_duration_s"`
	SceneCount      *int64 `json:"scene_count,omitempty"`
	Tone            string `json:"tone"`
	Language        string `json:"language"`
}

type EnginePolicy struct {
	LLM []string `json:"llm"`
}
