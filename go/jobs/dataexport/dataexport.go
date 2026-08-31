// Package dataexport defines the data:export asynq task payload —
// hand-written (no JSON-schema source yet), mirrors
// jobs/analyticsexport's exact shape. api/API-CONTENT.md §10a.
package dataexport

import "encoding/json"

func UnmarshalDataExportTask(data []byte) (DataExportTask, error) {
	var r DataExportTask
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DataExportTask) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// DataExportTask is the data:export asynq task payload — no
// range/format fields like analytics:export's Payload, an export is
// always "every category of the user's data, right now".
type DataExportTask struct {
	JobID          string `json:"job_id"`
	UserID         string `json:"user_id"`
	IdempotencyKey string `json:"idempotency_key"`
}
