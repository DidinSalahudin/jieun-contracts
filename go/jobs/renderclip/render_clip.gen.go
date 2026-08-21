// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    renderClipTask, err := UnmarshalRenderClipTask(bytes)
//    bytes, err = renderClipTask.Marshal()

package renderclip

import "encoding/json"

func UnmarshalRenderClipTask(data []byte) (RenderClipTask, error) {
	var r RenderClipTask
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RenderClipTask) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Payload of the render:clip asynq task. See api/API-CLIPPER.md §2.
type RenderClipTask struct {
	ActorType      ActorType `json:"actor_type"`
	IdempotencyKey string    `json:"idempotency_key"`
	JobID          string    `json:"job_id"`
	Payload        Payload   `json:"payload"`
	TraceID        string    `json:"trace_id"`
	UserID         string    `json:"user_id"`
}

type Payload struct {
	CandidateID   string  `json:"candidate_id"`
	Cut           Cut     `json:"cut"`
	OutputKey     string  `json:"output_key"`
	Profile       Profile `json:"profile"`
	SourceKey     string  `json:"source_key"`
	SourceVideoID string  `json:"source_video_id"`
	TargetVideoID string  `json:"target_video_id"`
	ThumbnailKey  string  `json:"thumbnail_key"`
}

type Cut struct {
	EndMS   int64 `json:"end_ms"`
	StartMS int64 `json:"start_ms"`
}

type Profile struct {
	Attribution *Attribution `json:"attribution,omitempty"`
	Audio       Audio        `json:"audio"`
	Encode      Encode       `json:"encode"`
	Fit         Fit          `json:"fit"`
	Ratio       Ratio        `json:"ratio"`
	Resolution  []int64      `json:"resolution"`
	Subtitle    Subtitle     `json:"subtitle"`
	Watermark   Watermark    `json:"watermark"`
}

type Attribution struct {
	Enabled bool    `json:"enabled"`
	Text    *string `json:"text,omitempty"`
}

type Audio struct {
	LoudnormLufs *float64 `json:"loudnorm_lufs,omitempty"`
	TruePeakDB   *float64 `json:"true_peak_db,omitempty"`
}

type Encode struct {
	AudioBitrateKbps *int64  `json:"audio_bitrate_kbps,omitempty"`
	AudioCodec       *string `json:"audio_codec,omitempty"`
	Codec            *string `json:"codec,omitempty"`
	Crf              *int64  `json:"crf,omitempty"`
	Faststart        *bool   `json:"faststart,omitempty"`
	Preset           *string `json:"preset,omitempty"`
}

type Subtitle struct {
	DualLine       *bool    `json:"dual_line,omitempty"`
	Enabled        bool     `json:"enabled"`
	Font           *string  `json:"font,omitempty"`
	HighlightColor *string  `json:"highlight_color,omitempty"`
	PositionRatio  *float64 `json:"position_ratio,omitempty"`
	SizeRatio      *float64 `json:"size_ratio,omitempty"`
	SourceLang     *string  `json:"source_lang,omitempty"`
	Style          *Style   `json:"style,omitempty"`
	TargetLang     *string  `json:"target_lang,omitempty"`
	WordsRef       *string  `json:"words_ref,omitempty"`
}

type Watermark struct {
	Enabled   bool      `json:"enabled"`
	Opacity   *float64  `json:"opacity,omitempty"`
	Position  *Position `json:"position,omitempty"`
	SizeRatio *float64  `json:"size_ratio,omitempty"`
	Text      *string   `json:"text,omitempty"`
}

type ActorType string

const (
	Admin  ActorType = "admin"
	System ActorType = "system"
	User   ActorType = "user"
)

type Fit string

const (
	CenterCrop Fit = "center_crop"
)

type Ratio string

const (
	The11  Ratio = "1:1"
	The169 Ratio = "16:9"
	The219 Ratio = "21:9"
	The43  Ratio = "4:3"
	The45  Ratio = "4:5"
	The916 Ratio = "9:16"
)

type Style string

const (
	Karaoke Style = "karaoke"
	Minimal Style = "minimal"
)

type Position string

const (
	BottomLeft  Position = "bottom_left"
	BottomRight Position = "bottom_right"
	TopLeft     Position = "top_left"
	TopRight    Position = "top_right"
)
