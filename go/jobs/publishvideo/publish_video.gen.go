// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    publishVideoTask, err := UnmarshalPublishVideoTask(bytes)
//    bytes, err = publishVideoTask.Marshal()

package publishvideo

import "encoding/json"

func UnmarshalPublishVideoTask(data []byte) (PublishVideoTask, error) {
	var r PublishVideoTask
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PublishVideoTask) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Payload of the publish:video asynq task. See api/API-PUBLISHER.md §5.
type PublishVideoTask struct {
	IdempotencyKey string  `json:"idempotency_key"`
	JobID          string  `json:"job_id"`
	Payload        Payload `json:"payload"`
	TraceID        string  `json:"trace_id"`
	UserID         string  `json:"user_id"`
}

type Payload struct {
	AccountID      string    `json:"account_id"`
	Caption        Caption   `json:"caption"`
	DistributionID string    `json:"distribution_id"`
	ObjectKey      string    `json:"object_key"`
	Platform       Platform  `json:"platform"`
	VideoMeta      VideoMeta `json:"video_meta"`
}

type Caption struct {
	Description string     `json:"description"`
	Tags        []string   `json:"tags"`
	Title       string     `json:"title"`
	Visibility  Visibility `json:"visibility"`
}

type VideoMeta struct {
	DurationMS int64  `json:"duration_ms"`
	Ratio      string `json:"ratio"`
	SizeBytes  int64  `json:"size_bytes"`
}

type Visibility string

const (
	Private  Visibility = "private"
	Public   Visibility = "public"
	Unlisted Visibility = "unlisted"
)

type Platform string

const (
	Facebook  Platform = "facebook"
	Instagram Platform = "instagram"
	Tiktok    Platform = "tiktok"
	Youtube   Platform = "youtube"
)
