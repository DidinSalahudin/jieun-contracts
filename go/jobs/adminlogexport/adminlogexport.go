// Package adminlogexport defines the admin:log-export asynq task
// payload — mirrors jobs/dataexport's exact shape, plus a Kind field
// (T-104 design: one task type parameterized by log kind, not five).
package adminlogexport

import "encoding/json"

func UnmarshalAdminLogExportTask(data []byte) (AdminLogExportTask, error) {
	var r AdminLogExportTask
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AdminLogExportTask) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// AdminLogExportTask is the admin:log-export asynq task payload. Kind
// is one of "system"/"server"/"engine"/"user"/"audit".
type AdminLogExportTask struct {
	JobID          string `json:"job_id"`
	UserID         string `json:"user_id"`
	Kind           string `json:"kind"`
	IdempotencyKey string `json:"idempotency_key"`
}
