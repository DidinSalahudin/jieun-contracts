// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    analysisCompletedEvent, err := UnmarshalAnalysisCompletedEvent(bytes)
//    bytes, err = analysisCompletedEvent.Marshal()

package analysiscompleted

import "encoding/json"

func UnmarshalAnalysisCompletedEvent(data []byte) (AnalysisCompletedEvent, error) {
	var r AnalysisCompletedEvent
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AnalysisCompletedEvent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// data payload of jieun.analysis.completed. See api/API-ANALYZER.md §6.
type AnalysisCompletedEvent struct {
	Candidates    []Candidate `json:"candidates"`
	Cost          Cost        `json:"cost"`
	DurationMS    int64       `json:"duration_ms"`
	EnginesUsed   EnginesUsed `json:"engines_used"`
	RunID         string      `json:"run_id"`
	TranscriptRef string      `json:"transcript_ref"`
	VideoID       string      `json:"video_id"`
}

type Candidate struct {
	CandidateID string   `json:"candidate_id"`
	EndMS       int64    `json:"end_ms"`
	Hook        string   `json:"hook"`
	Reason      string   `json:"reason"`
	Score       int64    `json:"score"`
	Signals     []string `json:"signals"`
	StartMS     int64    `json:"start_ms"`
}

type Cost struct {
	Idr         int64 `json:"idr"`
	LlmTokens   int64 `json:"llm_tokens"`
	TranscribeS int64 `json:"transcribe_s"`
}

type EnginesUsed struct {
	Llm        string `json:"llm"`
	Transcribe string `json:"transcribe"`
}
