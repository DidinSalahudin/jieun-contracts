#!/bin/sh
# Regenerates Go + TS types from openapi/, jobs/, events/.
# Run via `make generate`. Idempotent — safe to re-run.
#
# Each job/event schema gets its OWN Go subpackage (not a shared `jobs`
# package) because quicktype names nested types after JSON field names
# (e.g. every schema with a "payload" object generates a type literally
# called `Payload`) — sharing one package causes redeclaration errors.
set -eu
cd "$(dirname "$0")/.."

mkdir -p go/jobs/analyzestart go/jobs/renderclip go/jobs/publishvideo
mkdir -p go/events/envelope go/events/analysiscompleted go/events/rendercompleted go/events/renderprogress go/events/renderfailed
mkdir -p ts/src/jobs ts/src/events

gen_pair() {
  # $1 = schema file, $2 = top-level type name, $3 = go out, $4 = ts out, $5 = go package
  npx --yes quicktype --src "$1" --src-lang schema --lang go \
    --package "$5" --top-level "$2" --out "$3"
  npx --yes quicktype --src "$1" --src-lang schema --lang typescript \
    --top-level "$2" --out "$4"
}

gen_pair jobs/analyze-start.v1.json AnalyzeStartTask go/jobs/analyzestart/analyze_start.gen.go ts/src/jobs/analyze-start.gen.ts analyzestart
gen_pair jobs/render-clip.v1.json RenderClipTask go/jobs/renderclip/render_clip.gen.go ts/src/jobs/render-clip.gen.ts renderclip
gen_pair jobs/publish-video.v1.json PublishVideoTask go/jobs/publishvideo/publish_video.gen.go ts/src/jobs/publish-video.gen.ts publishvideo

gen_pair events/envelope.v1.json EventEnvelope go/events/envelope/envelope.gen.go ts/src/events/envelope.gen.ts envelope
gen_pair events/analysis-completed.v1.json AnalysisCompletedEvent go/events/analysiscompleted/analysis_completed.gen.go ts/src/events/analysis-completed.gen.ts analysiscompleted
gen_pair events/render-completed.v1.json RenderCompletedEvent go/events/rendercompleted/render_completed.gen.go ts/src/events/render-completed.gen.ts rendercompleted
gen_pair events/render-progress.v1.json RenderProgressEvent go/events/renderprogress/render_progress.gen.go ts/src/events/render-progress.gen.ts renderprogress
gen_pair events/render-failed.v1.json RenderFailedEvent go/events/renderfailed/render_failed.gen.go ts/src/events/render-failed.gen.ts renderfailed

mkdir -p go/gatewayapi
oapi-codegen -config codegen/oapi-codegen.gateway.yaml openapi/public-v1.yaml

npx --yes openapi-typescript openapi/public-v1.yaml -o ts/src/gateway.gen.ts

echo "generate.sh: done"
