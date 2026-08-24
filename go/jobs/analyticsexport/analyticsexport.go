// Package analyticsexport defines the analytics:export asynq task
// payload — hand-written (no JSON-schema source for this task yet,
// unlike its quicktype-generated siblings), api/API-ANALYTICS.md §7.
package analyticsexport

import "encoding/json"

func UnmarshalAnalyticsExportTask(data []byte) (AnalyticsExportTask, error) {
	var r AnalyticsExportTask
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AnalyticsExportTask) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// AnalyticsExportTask is the analytics:export asynq task payload.
type AnalyticsExportTask struct {
	JobID          string  `json:"job_id"`
	UserID         string  `json:"user_id"`
	IdempotencyKey string  `json:"idempotency_key"`
	Payload        Payload `json:"payload"`
}

type Payload struct {
	RangeDays int      `json:"range_days"`
	Format    string   `json:"format"`
	Include   []string `json:"include"`
}
