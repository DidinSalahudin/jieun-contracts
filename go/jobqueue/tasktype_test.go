package jobqueue

import "testing"

func TestQueueForTask_HasAllThreeRegisteredTypes(t *testing.T) {
	want := map[TaskType]string{
		TaskAnalyzeStart: QueueAnalyze,
		TaskAnalyzeStage: QueueAnalyze,
		TaskRenderClip:   QueueRender,
		TaskPublishVideo: QueuePublish,
	}
	if len(QueueForTask) != len(want) {
		t.Fatalf("QueueForTask has %d entries, want %d", len(QueueForTask), len(want))
	}
	for taskType, wantQueue := range want {
		gotQueue, ok := QueueForTask[taskType]
		if !ok {
			t.Errorf("QueueForTask missing %q", taskType)
			continue
		}
		if gotQueue != wantQueue {
			t.Errorf("QueueForTask[%q] = %q, want %q", taskType, gotQueue, wantQueue)
		}
	}
}

func TestQueuePriorities_MatchIntegrationDoc(t *testing.T) {
	want := map[string]int{
		QueuePublish: 6,
		QueueRender:  4,
		QueueAnalyze: 3,
		QueueDefault: 1,
	}
	if len(QueuePriorities) != len(want) {
		t.Fatalf("QueuePriorities has %d entries, want %d", len(QueuePriorities), len(want))
	}
	for queue, wantPriority := range want {
		gotPriority, ok := QueuePriorities[queue]
		if !ok || gotPriority != wantPriority {
			t.Errorf("QueuePriorities[%q] = %d, want %d", queue, gotPriority, wantPriority)
		}
	}
}

func TestTaskTypeConstants_MatchCatalogNames(t *testing.T) {
	if TaskAnalyzeStart != "analyze:start" {
		t.Errorf("TaskAnalyzeStart = %q, want \"analyze:start\"", TaskAnalyzeStart)
	}
	if TaskRenderClip != "render:clip" {
		t.Errorf("TaskRenderClip = %q, want \"render:clip\"", TaskRenderClip)
	}
	if TaskPublishVideo != "publish:video" {
		t.Errorf("TaskPublishVideo = %q, want \"publish:video\"", TaskPublishVideo)
	}
	if TaskAnalyzeStage != "analyze:stage" {
		t.Errorf("TaskAnalyzeStage = %q, want \"analyze:stage\"", TaskAnalyzeStage)
	}
}
