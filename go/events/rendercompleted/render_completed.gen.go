// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    renderCompletedEvent, err := UnmarshalRenderCompletedEvent(bytes)
//    bytes, err = renderCompletedEvent.Marshal()

package rendercompleted

import "encoding/json"

func UnmarshalRenderCompletedEvent(data []byte) (RenderCompletedEvent, error) {
	var r RenderCompletedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RenderCompletedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// data payload of jieun.render.completed. See api/API-CLIPPER.md §5.
type RenderCompletedEvent struct {
	Degraded      bool    `json:"degraded"`
	DurationMS    int64   `json:"duration_ms"`
	FfmpegVersion string  `json:"ffmpeg_version"`
	OutputKey     string  `json:"output_key"`
	Ratio         string  `json:"ratio"`
	RenderMS      int64   `json:"render_ms"`
	Resolution    []int64 `json:"resolution"`
	SizeBytes     int64   `json:"size_bytes"`
	TargetVideoID string  `json:"target_video_id"`
	ThumbnailKey  string  `json:"thumbnail_key"`
}
