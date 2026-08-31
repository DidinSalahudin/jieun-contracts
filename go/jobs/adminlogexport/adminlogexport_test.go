package adminlogexport

import "testing"

func TestAdminLogExportTask_MarshalUnmarshalRoundTrip(t *testing.T) {
	task := &AdminLogExportTask{JobID: "logexp_1", UserID: "u_1", Kind: "engine", IdempotencyKey: "logexp_1"}
	data, err := task.Marshal()
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	got, err := UnmarshalAdminLogExportTask(data)
	if err != nil {
		t.Fatalf("UnmarshalAdminLogExportTask: %v", err)
	}
	if got.JobID != "logexp_1" || got.UserID != "u_1" || got.Kind != "engine" || got.IdempotencyKey != "logexp_1" {
		t.Errorf("got = %+v, want JobID/UserID/Kind/IdempotencyKey = logexp_1/u_1/engine/logexp_1", got)
	}
}
