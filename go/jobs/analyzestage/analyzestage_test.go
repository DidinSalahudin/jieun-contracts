package analyzestage

import "testing"

func TestAnalyzeStageTask_MarshalUnmarshalRoundTrip(t *testing.T) {
	original := AnalyzeStageTask{
		RunID: "run_5c9", Stage: StageTranscribe, UserID: "u_4821", VideoID: "v_4821",
		TraceID: "trace_1", IdempotencyKey: "stage:run_5c9:transcribe",
		EnginePolicy: EnginePolicy{Llm: []string{"gemini-flash"}, Transcribe: []string{"local-whisper"}},
		Constraints:  Constraints{MaxCandidates: 8, MaxLenS: 60, MinLenS: 15},
	}

	data, err := original.Marshal()
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	got, err := UnmarshalAnalyzeStageTask(data)
	if err != nil {
		t.Fatalf("UnmarshalAnalyzeStageTask: %v", err)
	}
	if got.RunID != original.RunID || got.Stage != original.Stage {
		t.Errorf("got RunID=%q Stage=%q, want RunID=%q Stage=%q", got.RunID, got.Stage, original.RunID, original.Stage)
	}
	if len(got.EnginePolicy.Llm) != 1 || got.EnginePolicy.Llm[0] != "gemini-flash" {
		t.Errorf("EnginePolicy.Llm = %v, want [gemini-flash]", got.EnginePolicy.Llm)
	}
}

func TestStageConstants_AreDistinct(t *testing.T) {
	stages := []Stage{StageTranscribe, StageSignals, StageCandidates, StageJudge, StageDispatch}
	seen := map[Stage]bool{}
	for _, s := range stages {
		if seen[s] {
			t.Errorf("duplicate stage value %q", s)
		}
		seen[s] = true
	}
	if len(seen) != 5 {
		t.Fatalf("got %d distinct stages, want 5", len(seen))
	}
}
