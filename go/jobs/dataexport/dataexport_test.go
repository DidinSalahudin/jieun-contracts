package dataexport

import "testing"

func TestDataExportTask_MarshalUnmarshalRoundTrip(t *testing.T) {
	task := &DataExportTask{JobID: "exp_1", UserID: "u_1", IdempotencyKey: "exp_1"}
	data, err := task.Marshal()
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	got, err := UnmarshalDataExportTask(data)
	if err != nil {
		t.Fatalf("UnmarshalDataExportTask: %v", err)
	}
	if got.JobID != "exp_1" || got.UserID != "u_1" || got.IdempotencyKey != "exp_1" {
		t.Errorf("got = %+v, want JobID/UserID/IdempotencyKey = exp_1/u_1/exp_1", got)
	}
}
