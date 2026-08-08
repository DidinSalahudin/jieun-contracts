// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    analyzeStartTask, err := UnmarshalAnalyzeStartTask(bytes)
//    bytes, err = analyzeStartTask.Marshal()

package analyzestart

import "time"

import "encoding/json"

func UnmarshalAnalyzeStartTask(data []byte) (AnalyzeStartTask, error) {
	var r AnalyzeStartTask
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AnalyzeStartTask) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Payload of the analyze:start asynq task. See INTEGRATION.md §7.1.
type AnalyzeStartTask struct {
	ActorType      ActorType `json:"actor_type"`
	EnqueuedAt     time.Time `json:"enqueued_at"`
	IdempotencyKey string    `json:"idempotency_key"`
	JobID          string    `json:"job_id"`
	Payload        Payload   `json:"payload"`
	TraceID        string    `json:"trace_id"`
	UserID         string    `json:"user_id"`
}

type Payload struct {
	Budget        *Budget      `json:"budget,omitempty"`
	Constraints   Constraints  `json:"constraints"`
	EnginePolicy  EnginePolicy `json:"engine_policy"`
	LanguageHint  *string      `json:"language_hint,omitempty"`
	PromptVersion string       `json:"prompt_version"`
	Source        Source       `json:"source"`
	VideoID       string       `json:"video_id"`
}

type Budget struct {
	MaxIdr *int64 `json:"max_idr,omitempty"`
}

type Constraints struct {
	MaxCandidates int64 `json:"max_candidates"`
	MaxLenS       int64 `json:"max_len_s"`
	MinLenS       int64 `json:"min_len_s"`
}

type EnginePolicy struct {
	Llm        []string `json:"llm"`
	Transcribe []string `json:"transcribe"`
}

type Source struct {
	ObjectKey *string `json:"object_key"`
	Type      Type    `json:"type"`
	URL       *string `json:"url"`
}

type ActorType string

const (
	Admin  ActorType = "admin"
	System ActorType = "system"
	User   ActorType = "user"
)

type Type string

const (
	Upload     Type = "upload"
	YoutubeURL Type = "youtube_url"
)
