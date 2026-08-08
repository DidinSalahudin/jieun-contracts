// Namespaced re-exports — quicktype names nested/helper types after JSON
// field names (every schema with a "payload" object exports a type
// literally called `Payload`, plus a `Convert` helper per file), so a flat
// `export *` collides across files. Same root cause as the Go side
// (see codegen/generate.sh comment); fixed the same way, via isolation.
export * from "./gateway.gen";
export * as analyzeStart from "./jobs/analyze-start.gen";
export * as renderClip from "./jobs/render-clip.gen";
export * as publishVideo from "./jobs/publish-video.gen";
export * as envelope from "./events/envelope.gen";
export * as analysisCompleted from "./events/analysis-completed.gen";
export * as renderCompleted from "./events/render-completed.gen";
