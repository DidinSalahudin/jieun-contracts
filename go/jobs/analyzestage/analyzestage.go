package analyzestage

import "encoding/json"

func UnmarshalAnalyzeStageTask(data []byte) (AnalyzeStageTask, error) {
	var r AnalyzeStageTask
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AnalyzeStageTask) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Stage identifies which of the 5 self-chained stages (everything after
// extract, which runs inside analyze:start) this task carries out.
type Stage string

const (
	StageTranscribe Stage = "transcribe"
	StageSignals    Stage = "signals"
	StageCandidates Stage = "candidates"
	StageJudge      Stage = "judge"
	StageDispatch   Stage = "dispatch"
)

// AnalyzeStageTask is the payload of the analyze:stage asynq task —
// INTEGRATION.md §6.3's self-chained follow-up to analyze:start. Payload
// stays minimal: each stage re-derives its own working data (transcript,
// signal windows, candidates) from Postgres rather than carrying it here.
type AnalyzeStageTask struct {
	RunID          string       `json:"run_id"`
	Stage          Stage        `json:"stage"`
	UserID         string       `json:"user_id"`
	VideoID        string       `json:"video_id"`
	TraceID        string       `json:"trace_id"`
	IdempotencyKey string       `json:"idempotency_key"`
	EnginePolicy   EnginePolicy `json:"engine_policy"`
	Budget         *Budget      `json:"budget,omitempty"`
	Constraints    Constraints  `json:"constraints"`
}

// EnginePolicy, Budget, and Constraints mirror analyzestart's identically
// named types (jieun-contracts/go/jobs/analyzestart) — duplicated here
// rather than imported so this package has no dependency on analyzestart
// (they're conceptually independent contracts that happen to share
// shape), matching how neither package imports the other today.
type EnginePolicy struct {
	Llm        []string `json:"llm"`
	Transcribe []string `json:"transcribe"`
}

type Budget struct {
	MaxIdr *int64 `json:"max_idr,omitempty"`
}

type Constraints struct {
	MaxCandidates int64 `json:"max_candidates"`
	MaxLenS       int64 `json:"max_len_s"`
	MinLenS       int64 `json:"min_len_s"`
}
