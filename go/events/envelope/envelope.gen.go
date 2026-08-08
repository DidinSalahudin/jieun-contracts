// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    eventEnvelope, err := UnmarshalEventEnvelope(bytes)
//    bytes, err = eventEnvelope.Marshal()

package envelope

import "time"

import "encoding/json"

func UnmarshalEventEnvelope(data []byte) (EventEnvelope, error) {
	var r EventEnvelope
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EventEnvelope) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Common envelope every jieun.* event uses. See INTEGRATION.md §6.4.
type EventEnvelope struct {
	Data                            map[string]interface{} `json:"data"`
	// e.g. jieun.analysis.completed                       
	Event                           string                 `json:"event"`
	EventID                         string                 `json:"event_id"`
	JobID                           *string                `json:"job_id,omitempty"`
	OccurredAt                      time.Time              `json:"occurred_at"`
	TraceID                         string                 `json:"trace_id"`
	UserID                          string                 `json:"user_id"`
	Version                         int64                  `json:"version"`
}
