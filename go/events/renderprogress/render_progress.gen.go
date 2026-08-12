// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    renderProgressEvent, err := UnmarshalRenderProgressEvent(bytes)
//    bytes, err = renderProgressEvent.Marshal()

package renderprogress

import "encoding/json"

func UnmarshalRenderProgressEvent(data []byte) (RenderProgressEvent, error) {
	var r RenderProgressEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RenderProgressEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// data payload of jieun.render.progress. See api/API-CLIPPER.md §5.
type RenderProgressEvent struct {
	Percent       int64  `json:"percent"`
	Step          string `json:"step"`
	StepLabel     string `json:"step_label"`
	TargetVideoID string `json:"target_video_id"`
}
