// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    renderFailedEvent, err := UnmarshalRenderFailedEvent(bytes)
//    bytes, err = renderFailedEvent.Marshal()

package renderfailed

import "encoding/json"

func UnmarshalRenderFailedEvent(data []byte) (RenderFailedEvent, error) {
	var r RenderFailedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RenderFailedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// data payload of jieun.render.failed. See api/API-CLIPPER.md §5.
type RenderFailedEvent struct {
	Error            Error   `json:"error"`
	FfmpegStderrTail *string `json:"ffmpeg_stderr_tail,omitempty"`
	TargetVideoID    string  `json:"target_video_id"`
}

type Error struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Retryable bool   `json:"retryable"`
}
