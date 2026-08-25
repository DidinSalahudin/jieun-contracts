// Package storyboardrender defines the storyboard:render asynq task
// payload — hand-written (no JSON-schema source exists for this task
// yet, same as storyboardgenerate/analyticsexport), api/API-STORYBOARD.md
// §5. Scenes are embedded in full at enqueue time — unlike render:clip's
// subtitle words (fetched live via analyzer's internal API because clip
// candidates reference a dynamic transcript), a storyboard is already
// static once status=ready, so there's nothing to re-fetch at render time.
package storyboardrender

import "encoding/json"

func UnmarshalStoryboardRenderTask(data []byte) (StoryboardRenderTask, error) {
	var r StoryboardRenderTask
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StoryboardRenderTask) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StoryboardRenderTask struct {
	JobID          string  `json:"job_id"`
	TraceID        string  `json:"trace_id"`
	UserID         string  `json:"user_id"`
	IdempotencyKey string  `json:"idempotency_key"`
	Payload        Payload `json:"payload"`
}

type Payload struct {
	StoryboardID  string       `json:"storyboard_id"`
	TargetVideoID string       `json:"target_video_id"`
	OutputKey     string       `json:"output_key"`
	Scenes        []Scene      `json:"scenes"`
	Profile       Profile      `json:"profile"`
	EnginePolicy  EnginePolicy `json:"engine_policy"`
}

type Scene struct {
	N            int64        `json:"n"`
	DurationS    int64        `json:"duration_s"`
	Narration    string       `json:"narration"`
	VisualSource VisualSource `json:"visual_source"`
}

type VisualSource struct {
	Type  string `json:"type"`
	Query string `json:"query"`
}

type Profile struct {
	Ratio     string    `json:"ratio"`
	Voice     Voice     `json:"voice"`
	Subtitle  Subtitle  `json:"subtitle"`
	Music     Music     `json:"music"`
	Watermark Watermark `json:"watermark"`
}

type Voice struct {
	Engine  string  `json:"engine"`
	VoiceID string  `json:"voice_id"`
	Speed   float64 `json:"speed"`
}

type Subtitle struct {
	Enabled bool   `json:"enabled"`
	Style   string `json:"style"`
}

type Music struct {
	Enabled bool    `json:"enabled"`
	Mood    string  `json:"mood"`
	Volume  float64 `json:"volume"`
}

type Watermark struct {
	Enabled  bool   `json:"enabled"`
	Text     string `json:"text"`
	Position string `json:"position"`
}

type EnginePolicy struct {
	TTS   []string `json:"tts"`
	Image []string `json:"image"`
	Music []string `json:"music"`
}
