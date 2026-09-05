# Graph Report - jieun-contracts  (2026-08-23)

## Corpus Check
- Corpus is ~18,163 words - fits in a single context window. You may not need a graph.

## Summary
- 1109 nodes · 1346 edges · 63 communities (60 shown, 3 thin omitted)
- Extraction: 97% EXTRACTED · 3% INFERRED · 0% AMBIGUOUS · INFERRED: 38 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- Gatewayapi — Contenttype Getbody
- Events — Type Key
- Jobs — Type Properties
- Jobs — Type Key
- Jobqueue — Client Enqueue
- Jobqueue — Actortype Analyzestarttask
- Jobs — Type Properties
- Events — Envelope Gen
- Jobs — Render Clip
- Events — Stage Type
- Events — Type Description
- Jobs — Type User
- Jobs — Analyze Start
- Renderclip — Attribution Audio
- Events — Type Minimum
- Jobs — Publish Video
- Jobs — Type User
- Events — Analysis Completed
- Events — Render Failed
- Events — Render Completed
- Events — Render Progress
- Jobs — Type Description
- Events — Type Properties
- Jobs — Type Max
- Publishvideo — Caption Payload
- Package.Json — Openapi Typescript
- Events — Type Minimum
- Events — Type Minimum
- Events — Schema
- Jobs — Type Llm
- Analysiscompleted — Analysiscompletedevent Candidate
- Events — Step Render
- Events — Type Step
- Jobs — Type Properties
- Events — Type Llm
- Jobs — Minimum Type
- Events — Render Failed
- Jobs — Constraints Engine
- Events — Candidates Duration
- Jobs — Publish Video
- Jobs — Type Items
- Jobs — Constraints Engine
- Renderfailed — Error Renderfailedevent
- Jobs — Candidates Dispatch
- Jobs — Type Properties
- Jobs — Type Object
- Ts — Package Json
- Events — Analysis Completed
- Jobs — Schema
- Jobs — Analyze Stage
- Jobs — Type Properties
- Jobs — Type Properties
- Jobs — Max Len
- Jobs — Idempotency Key
- Rendercompleted — Rendercompletedevent Gen
- Renderprogress — Renderprogressevent Gen
- Jobs — Required Type
- Jobs — Type Enum
- Events — Required Idr
- Jobs — Type Source
- Codegen — Generate Pair
- Githooks — Pre Commit
- Go.Mod — Github Com

## God Nodes (most connected - your core abstractions)
1. `required` - 11 edges
2. `GetHealthResponse` - 10 edges
3. `transform()` - 10 edges
4. `transform()` - 10 edges
5. `transform()` - 10 edges
6. `transform()` - 10 edges
7. `transform()` - 10 edges
8. `transform()` - 10 edges
9. `transform()` - 10 edges
10. `transform()` - 10 edges

## Surprising Connections (you probably didn't know these)
- `$schema` --shares_data_with--> `$schema`  [INFERRED]
  jobs/analyze-start.v1.json → events/analysis-completed.v1.json
- `$schema` --shares_data_with--> `$schema`  [INFERRED]
  events/render-completed.v1.json → jobs/publish-video.v1.json
- `$schema` --shares_data_with--> `$schema`  [INFERRED]
  events/analysis-completed.v1.json → jobs/render-clip.v1.json
- `AnalysisCompletedEvent generated Go client` --references--> `$schema`  [EXTRACTED]
  go/events/analysiscompleted/analysis_completed.gen.go → events/analysis-completed.v1.json
- `$schema` --shares_data_with--> `$schema`  [INFERRED]
  jobs/analyze-stage.v1.json → events/analysis-progress.v1.json

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Analysis async job chain: analyze:start seeds run, analyze:stage self-chains remaining stages, progress/completion events report status** — jobs_analyze_start_v1_schema, jobs_analyze_stage_v1_schema, events_analysis_progress_v1_schema, events_analysis_completed_v1_schema [INFERRED 0.85]
- **Render async job: render:clip task produces progress, completion, or failure events keyed by target_video_id** — jobs_render_clip_v1_schema, events_render_progress_v1_schema, events_render_completed_v1_schema, events_render_failed_v1_schema [INFERRED 0.85]
- **Every jieun.* event payload schema is carried inside the common EventEnvelope's data field** — events_envelope_v1_schema, events_analysis_completed_v1_schema, events_analysis_progress_v1_schema, events_render_completed_v1_schema, events_render_failed_v1_schema, events_render_progress_v1_schema [EXTRACTED 1.00]

## Communities (63 total, 3 thin omitted)

### Community 0 - "Gatewayapi — Contenttype Getbody"
Cohesion: 0.06
Nodes (31): Client, ClientInterface, ClientOption, ClientWithResponses, ClientWithResponsesInterface, Error, ErrorErrorCode, GetHealth200JSONResponseBodyStatus (+23 more)

### Community 1 - "Events — Type Key"
Cohesion: 0.04
Nodes (46): additionalProperties, type, type, description, minimum, type, type, $id (+38 more)

### Community 2 - "Jobs — Type Properties"
Cohesion: 0.04
Nodes (47): type, type, minimum, type, type, type, duration_ms, ratio (+39 more)

### Community 3 - "Jobs — Type Key"
Cohesion: 0.04
Nodes (46): type, properties, required, type, minimum, type, candidate_id, end_ms (+38 more)

### Community 4 - "Jobqueue — Client Enqueue"
Cohesion: 0.07
Nodes (28): AnalyzeStageTask, Budget, Constraints, EnginePolicy, Stage, NewClient(), TestClient_Enqueue_DuplicateIdempotencyKeyRejected(), TestClient_Enqueue_RoutesToCorrectQueue() (+20 more)

### Community 5 - "Jobqueue — Actortype Analyzestarttask"
Cohesion: 0.06
Nodes (29): ActorType, AnalyzeStartTask, Budget, Constraints, EnginePolicy, Payload, Source, Type (+21 more)

### Community 6 - "Jobs — Type Properties"
Cohesion: 0.06
Nodes (38): properties, required, type, properties, type, properties, type, enum (+30 more)

### Community 7 - "Events — Envelope Gen"
Cohesion: 0.08
Nodes (29): cast(), Convert, EventEnvelope, invalidValue(), jsonToJSProps(), jsToJSONProps(), l(), prettyTypeName() (+21 more)

### Community 8 - "Jobs — Render Clip"
Cohesion: 0.08
Nodes (31): ActorType, Attribution, Audio, cast(), Convert, Cut, Encode, Fit (+23 more)

### Community 9 - "Events — Stage Type"
Cohesion: 0.06
Nodes (34): additionalProperties, description, $id, candidates, dispatch, judge, run_id, signals (+26 more)

### Community 10 - "Events — Type Description"
Cohesion: 0.06
Nodes (33): additionalProperties, type, description, description, type, type, $id, type (+25 more)

### Community 11 - "Jobs — Type User"
Cohesion: 0.06
Nodes (31): enum, type, additionalProperties, description, format, type, $id, type (+23 more)

### Community 12 - "Jobs — Analyze Start"
Cohesion: 0.10
Nodes (25): ActorType, AnalyzeStartTask, Budget, cast(), Constraints, Convert, EnginePolicy, invalidValue() (+17 more)

### Community 13 - "Renderclip — Attribution Audio"
Cohesion: 0.06
Nodes (29): Attribution, Audio, Cut, Encode, Fit, ActorType, Payload, UnmarshalRenderClipTask() (+21 more)

### Community 14 - "Events — Type Minimum"
Cohesion: 0.07
Nodes (30): type, items, minimum, type, type, properties, required, type (+22 more)

### Community 15 - "Jobs — Publish Video"
Cohesion: 0.11
Nodes (23): Caption, cast(), Convert, invalidValue(), jsonToJSProps(), jsToJSONProps(), l(), Payload (+15 more)

### Community 16 - "Jobs — Type User"
Cohesion: 0.07
Nodes (27): enum, type, additionalProperties, description, $id, type, type, actor_type (+19 more)

### Community 17 - "Events — Analysis Completed"
Cohesion: 0.12
Nodes (21): AnalysisCompletedEvent, Candidate, cast(), Convert, Cost, EnginesUsed, invalidValue(), jsonToJSProps() (+13 more)

### Community 18 - "Events — Render Failed"
Cohesion: 0.14
Nodes (19): cast(), Convert, Error, invalidValue(), jsonToJSProps(), jsToJSONProps(), l(), prettyTypeName() (+11 more)

### Community 19 - "Events — Render Completed"
Cohesion: 0.15
Nodes (18): cast(), Convert, invalidValue(), jsonToJSProps(), jsToJSONProps(), l(), prettyTypeName(), r() (+10 more)

### Community 20 - "Events — Render Progress"
Cohesion: 0.15
Nodes (18): cast(), Convert, invalidValue(), jsonToJSProps(), jsToJSONProps(), l(), prettyTypeName(), r() (+10 more)

### Community 21 - "Jobs — Type Description"
Cohesion: 0.09
Nodes (22): properties, required, type, type, caption, description, tags, title (+14 more)

### Community 22 - "Events — Type Properties"
Cohesion: 0.11
Nodes (19): type, additionalProperties, properties, required, type, type, type, properties (+11 more)

### Community 23 - "Jobs — Type Max"
Cohesion: 0.12
Nodes (16): properties, required, type, max_candidates, max_len_s, min_len_s, minimum, type (+8 more)

### Community 24 - "Publishvideo — Caption Payload"
Cohesion: 0.13
Nodes (13): Caption, Payload, UnmarshalPublishVideoTask(), Platform, Caption, Payload, Platform, PublishVideoTask (+5 more)

### Community 25 - "Package.Json — Openapi Typescript"
Cohesion: 0.13
Nodes (14): openapi-typescript, author, description, devDependencies, openapi-typescript, quicktype, keywords, license (+6 more)

### Community 26 - "Events — Type Minimum"
Cohesion: 0.17
Nodes (12): type, minimum, type, properties, candidates, duration_ms, run_id, transcript_ref (+4 more)

### Community 27 - "Events — Type Minimum"
Cohesion: 0.17
Nodes (12): properties, type, minimum, type, minimum, type, cost, idr (+4 more)

### Community 28 - "Events — Schema"
Cohesion: 0.21
Nodes (12): $schema, $schema, $schema, $schema, $schema, AnalysisCompletedEvent generated Go client, EventEnvelope generated Go client, RenderCompletedEvent generated Go client (+4 more)

### Community 29 - "Jobs — Type Llm"
Cohesion: 0.17
Nodes (12): properties, required, type, llm, transcribe, items, type, engine_policy (+4 more)

### Community 30 - "Analysiscompleted — Analysiscompletedevent Candidate"
Cohesion: 0.18
Nodes (9): AnalysisCompletedEvent, Candidate, Cost, EnginesUsed, AnalysisCompletedEvent, Candidate, Cost, EnginesUsed (+1 more)

### Community 31 - "Events — Step Render"
Cohesion: 0.18
Nodes (10): additionalProperties, description, $id, target_video_id, required, title, type, percent (+2 more)

### Community 32 - "Events — Type Step"
Cohesion: 0.18
Nodes (11): maximum, minimum, type, properties, percent, step, step_label, target_video_id (+3 more)

### Community 33 - "Jobs — Type Properties"
Cohesion: 0.18
Nodes (11): type, properties, idempotency_key, run_id, trace_id, user_id, video_id, type (+3 more)

### Community 34 - "Events — Type Llm"
Cohesion: 0.20
Nodes (10): properties, required, type, llm, transcribe, type, engines_used, llm (+2 more)

### Community 35 - "Jobs — Minimum Type"
Cohesion: 0.20
Nodes (10): properties, minimum, type, minimum, type, minimum, type, max_candidates (+2 more)

### Community 36 - "Events — Render Failed"
Cohesion: 0.22
Nodes (8): additionalProperties, description, $id, target_video_id, required, title, type, error

### Community 37 - "Jobs — Constraints Engine"
Cohesion: 0.22
Nodes (9): constraints, engine_policy, idempotency_key, run_id, stage, trace_id, user_id, video_id (+1 more)

### Community 38 - "Events — Candidates Duration"
Cohesion: 0.25
Nodes (8): candidates, duration_ms, run_id, video_id, required, cost, engines_used, transcript_ref

### Community 39 - "Jobs — Publish Video"
Cohesion: 0.25
Nodes (7): PublishVideoTask generated Go client, additionalProperties, description, $id, $schema, title, type

### Community 40 - "Jobs — Type Items"
Cohesion: 0.29
Nodes (8): properties, type, items, type, llm, transcribe, items, type

### Community 41 - "Jobs — Constraints Engine"
Cohesion: 0.25
Nodes (8): constraints, engine_policy, video_id, required, type, payload, prompt_version, source

### Community 42 - "Renderfailed — Error Renderfailedevent"
Cohesion: 0.29
Nodes (5): Error, UnmarshalRenderFailedEvent(), Error, RenderFailedEvent, RenderFailedEvent

### Community 43 - "Jobs — Candidates Dispatch"
Cohesion: 0.29
Nodes (7): candidates, dispatch, judge, signals, stage, enum, type

### Community 44 - "Jobs — Type Properties"
Cohesion: 0.29
Nodes (7): type, properties, type, language_hint, prompt_version, video_id, type

### Community 45 - "Jobs — Type Object"
Cohesion: 0.38
Nodes (7): type, object_key, url, properties, type, null, string

### Community 46 - "Ts — Package Json"
Cohesion: 0.29
Nodes (6): description, name, private, type, types, version

### Community 47 - "Events — Analysis Completed"
Cohesion: 0.33
Nodes (5): additionalProperties, description, $id, title, type

### Community 48 - "Jobs — Schema"
Cohesion: 0.40
Nodes (6): $schema, AnalyzeStageTask hand-written Go type, AnalyzeStageTask Go round-trip test, AnalyzeStartTask generated Go client, $schema, $schema

### Community 49 - "Jobs — Analyze Stage"
Cohesion: 0.33
Nodes (5): additionalProperties, description, $id, title, type

### Community 50 - "Jobs — Type Properties"
Cohesion: 0.33
Nodes (6): properties, type, minimum, type, budget, max_idr

### Community 51 - "Jobs — Type Properties"
Cohesion: 0.33
Nodes (6): properties, type, minimum, type, budget, max_idr

### Community 52 - "Jobs — Max Len"
Cohesion: 0.33
Nodes (6): required, type, max_candidates, max_len_s, min_len_s, constraints

### Community 53 - "Jobs — Idempotency Key"
Cohesion: 0.33
Nodes (6): idempotency_key, job_id, payload, trace_id, user_id, required

### Community 54 - "Rendercompleted — Rendercompletedevent Gen"
Cohesion: 0.40
Nodes (3): UnmarshalRenderCompletedEvent(), RenderCompletedEvent, RenderCompletedEvent

### Community 55 - "Renderprogress — Renderprogressevent Gen"
Cohesion: 0.40
Nodes (3): UnmarshalRenderProgressEvent(), RenderProgressEvent, RenderProgressEvent

### Community 56 - "Jobs — Required Type"
Cohesion: 0.40
Nodes (5): required, type, llm, transcribe, engine_policy

### Community 57 - "Jobs — Type Enum"
Cohesion: 0.40
Nodes (5): type, enum, type, upload, youtube_url

### Community 58 - "Events — Required Idr"
Cohesion: 0.50
Nodes (4): required, idr, llm_tokens, transcribe_s

### Community 59 - "Jobs — Type Source"
Cohesion: 0.50
Nodes (4): source, required, type, type

## Knowledge Gaps
- **490 isolated node(s):** `$id`, `title`, `description`, `type`, `video_id` (+485 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `$schema` connect `Events — Schema` to `Jobs — Schema`, `Events — Analysis Completed`?**
  _High betweenness centrality (0.155) - this node is a cross-community bridge._
- **Why does `$schema` connect `Events — Schema` to `Jobs — Type User`?**
  _High betweenness centrality (0.151) - this node is a cross-community bridge._
- **Why does `$schema` connect `Events — Schema` to `Events — Type Key`, `Jobs — Publish Video`?**
  _High betweenness centrality (0.129) - this node is a cross-community bridge._
- **What connects `$id`, `title`, `description` to the rest of the system?**
  _490 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Gatewayapi — Contenttype Getbody` be split into smaller, more focused modules?**
  _Cohesion score 0.06448087431693988 - nodes in this community are weakly interconnected._
- **Should `Events — Type Key` be split into smaller, more focused modules?**
  _Cohesion score 0.0425531914893617 - nodes in this community are weakly interconnected._
- **Should `Jobs — Type Properties` be split into smaller, more focused modules?**
  _Cohesion score 0.0425531914893617 - nodes in this community are weakly interconnected._