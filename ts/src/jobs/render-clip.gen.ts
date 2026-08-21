// To parse this data:
//
//   import { Convert, RenderClipTask } from "./render-clip.gen";
//
//   const renderClipTask = Convert.toRenderClipTask(json);
//
// These functions will throw an error if the JSON doesn't
// match the expected interface, even if the JSON is valid.

/**
 * Payload of the render:clip asynq task. See api/API-CLIPPER.md §2.
 */
export interface RenderClipTask {
    actor_type:      ActorType;
    idempotency_key: string;
    job_id:          string;
    payload:         Payload;
    trace_id:        string;
    user_id:         string;
}

export type ActorType = "user" | "admin" | "system";

export interface Payload {
    candidate_id:    string;
    cut:             Cut;
    output_key:      string;
    profile:         Profile;
    source_key:      string;
    source_video_id: string;
    target_video_id: string;
    thumbnail_key:   string;
    [property: string]: unknown;
}

export interface Cut {
    end_ms:   number;
    start_ms: number;
    [property: string]: unknown;
}

export interface Profile {
    attribution?: Attribution;
    audio:        Audio;
    encode:       Encode;
    fit:          Fit;
    ratio:        Ratio;
    resolution:   [number, number, ...number[]];
    subtitle:     Subtitle;
    watermark:    Watermark;
    [property: string]: unknown;
}

export interface Attribution {
    enabled: boolean;
    text?:   string;
    [property: string]: unknown;
}

export interface Audio {
    loudnorm_lufs?: number;
    true_peak_db?:  number;
    [property: string]: unknown;
}

export interface Encode {
    audio_bitrate_kbps?: number;
    audio_codec?:        string;
    codec?:              string;
    crf?:                number;
    faststart?:          boolean;
    preset?:             string;
    [property: string]: unknown;
}

export type Fit = "center_crop";

export type Ratio = "9:16" | "4:5" | "1:1" | "4:3" | "16:9" | "21:9";

export interface Subtitle {
    dual_line?:       boolean;
    enabled:          boolean;
    font?:            string;
    highlight_color?: string;
    position_ratio?:  number;
    size_ratio?:      number;
    source_lang?:     string;
    style?:           Style;
    target_lang?:     string;
    words_ref?:       string;
    [property: string]: unknown;
}

export type Style = "karaoke" | "minimal";

export interface Watermark {
    enabled:     boolean;
    opacity?:    number;
    position?:   Position;
    size_ratio?: number;
    text?:       string;
    [property: string]: unknown;
}

export type Position = "top_left" | "top_right" | "bottom_left" | "bottom_right";

// Converts JSON strings to/from your types
// and asserts the results of JSON.parse at runtime
export class Convert {
    public static toRenderClipTask(json: string): RenderClipTask {
        return cast(JSON.parse(json), r("RenderClipTask"));
    }

    public static renderClipTaskToJson(value: RenderClipTask): string {
        return JSON.stringify(uncast(value, r("RenderClipTask")), null, 2);
    }
}

function invalidValue(typ: any, val: any, key: any, parent: any = ''): never {
    const prettyTyp = prettyTypeName(typ);
    const parentText = parent ? ` on ${parent}` : '';
    const keyText = key ? ` for key "${key}"` : '';
    throw Error(`Invalid value${keyText}${parentText}. Expected ${prettyTyp} but got ${JSON.stringify(val)}`);
}

function prettyTypeName(typ: any): string {
    if (Array.isArray(typ)) {
        if (typ.length === 2 && typ[0] === undefined) {
            return `an optional ${prettyTypeName(typ[1])}`;
        } else {
            return `one of [${typ.map(a => { return prettyTypeName(a); }).join(", ")}]`;
        }
    } else if (typeof typ === "object" && typ.literal !== undefined) {
        return typ.literal;
    } else {
        return typeof typ;
    }
}

function jsonToJSProps(typ: any): any {
    if (typ.jsonToJS === undefined) {
        const map: any = {};
        typ.props.forEach((p: any) => map[p.json] = { key: p.js, typ: p.typ });
        typ.jsonToJS = map;
    }
    return typ.jsonToJS;
}

function jsToJSONProps(typ: any): any {
    if (typ.jsToJSON === undefined) {
        const map: any = {};
        typ.props.forEach((p: any) => map[p.js] = { key: p.json, typ: p.typ });
        typ.jsToJSON = map;
    }
    return typ.jsToJSON;
}

function transform(val: any, typ: any, getProps: any, key: any = '', parent: any = ''): any {
    function transformPrimitive(typ: string, val: any): any {
        if (typeof typ === typeof val) return val;
        return invalidValue(typ, val, key, parent);
    }

    function transformUnion(typs: any[], val: any): any {
        // val must validate against one typ in typs
        const l = typs.length;
        for (let i = 0; i < l; i++) {
            const typ = typs[i];
            try {
                return transform(val, typ, getProps);
            } catch (_) {}
        }
        return invalidValue(typs, val, key, parent);
    }

    function transformEnum(cases: string[], val: any): any {
        if (cases.indexOf(val) !== -1) return val;
        return invalidValue(cases.map(a => { return l(a); }), val, key, parent);
    }

    function transformArray(typ: any, val: any): any {
        // val must be an array with no invalid elements
        if (!Array.isArray(val)) return invalidValue(l("array"), val, key, parent);
        return val.map(el => transform(el, typ, getProps));
    }

    function transformDate(val: any): any {
        if (val === null) {
            return null;
        }
        const d = new Date(val);
        if (isNaN(d.valueOf())) {
            return invalidValue(l("Date"), val, key, parent);
        }
        return d;
    }

    function transformObject(props: { [k: string]: any }, additional: any, val: any): any {
        if (val === null || typeof val !== "object" || Array.isArray(val)) {
            return invalidValue(l(ref || "object"), val, key, parent);
        }
        const result: any = {};
        Object.getOwnPropertyNames(props).forEach(key => {
            const prop = props[key];
            const v = Object.prototype.hasOwnProperty.call(val, key) ? val[key] : undefined;
            result[prop.key] = transform(v, prop.typ, getProps, key, ref);
        });
        Object.getOwnPropertyNames(val).forEach(key => {
            if (!Object.prototype.hasOwnProperty.call(props, key)) {
                result[key] = transform(val[key], additional, getProps, key, ref);
            }
        });
        return result;
    }

    if (typ === "any") return val;
    if (typ === null) {
        if (val === null) return val;
        return invalidValue(typ, val, key, parent);
    }
    if (typ === false) return invalidValue(typ, val, key, parent);
    let ref: any = undefined;
    while (typeof typ === "object" && typ.ref !== undefined) {
        ref = typ.ref;
        typ = typeMap[typ.ref];
    }
    if (Array.isArray(typ)) return transformEnum(typ, val);
    if (typeof typ === "object") {
        return typ.hasOwnProperty("unionMembers") ? transformUnion(typ.unionMembers, val)
            : typ.hasOwnProperty("arrayItems")    ? transformArray(typ.arrayItems, val)
            : typ.hasOwnProperty("props")         ? transformObject(getProps(typ), typ.additional, val)
            : invalidValue(typ, val, key, parent);
    }
    // Numbers can be parsed by Date but shouldn't be.
    if (typ === Date && typeof val !== "number") return transformDate(val);
    return transformPrimitive(typ, val);
}

function cast<T>(val: any, typ: any): T {
    return transform(val, typ, jsonToJSProps);
}

function uncast<T>(val: T, typ: any): any {
    return transform(val, typ, jsToJSONProps);
}

function l(typ: any) {
    return { literal: typ };
}

function a(typ: any) {
    return { arrayItems: typ };
}

function u(...typs: any[]) {
    return { unionMembers: typs };
}

function o(props: any[], additional: any) {
    return { props, additional };
}

function m(additional: any) {
    return { props: [], additional };
}

function r(name: string) {
    return { ref: name };
}

const typeMap: any = {
    "RenderClipTask": o([
        { json: "actor_type", js: "actor_type", typ: r("ActorType") },
        { json: "idempotency_key", js: "idempotency_key", typ: "" },
        { json: "job_id", js: "job_id", typ: "" },
        { json: "payload", js: "payload", typ: r("Payload") },
        { json: "trace_id", js: "trace_id", typ: "" },
        { json: "user_id", js: "user_id", typ: "" },
    ], false),
    "Payload": o([
        { json: "candidate_id", js: "candidate_id", typ: "" },
        { json: "cut", js: "cut", typ: r("Cut") },
        { json: "output_key", js: "output_key", typ: "" },
        { json: "profile", js: "profile", typ: r("Profile") },
        { json: "source_key", js: "source_key", typ: "" },
        { json: "source_video_id", js: "source_video_id", typ: "" },
        { json: "target_video_id", js: "target_video_id", typ: "" },
        { json: "thumbnail_key", js: "thumbnail_key", typ: "" },
    ], "any"),
    "Cut": o([
        { json: "end_ms", js: "end_ms", typ: 0 },
        { json: "start_ms", js: "start_ms", typ: 0 },
    ], "any"),
    "Profile": o([
        { json: "attribution", js: "attribution", typ: u(undefined, r("Attribution")) },
        { json: "audio", js: "audio", typ: r("Audio") },
        { json: "encode", js: "encode", typ: r("Encode") },
        { json: "fit", js: "fit", typ: r("Fit") },
        { json: "ratio", js: "ratio", typ: r("Ratio") },
        { json: "resolution", js: "resolution", typ: a(0) },
        { json: "subtitle", js: "subtitle", typ: r("Subtitle") },
        { json: "watermark", js: "watermark", typ: r("Watermark") },
    ], "any"),
    "Attribution": o([
        { json: "enabled", js: "enabled", typ: true },
        { json: "text", js: "text", typ: u(undefined, "") },
    ], "any"),
    "Audio": o([
        { json: "loudnorm_lufs", js: "loudnorm_lufs", typ: u(undefined, 3.14) },
        { json: "true_peak_db", js: "true_peak_db", typ: u(undefined, 3.14) },
    ], "any"),
    "Encode": o([
        { json: "audio_bitrate_kbps", js: "audio_bitrate_kbps", typ: u(undefined, 0) },
        { json: "audio_codec", js: "audio_codec", typ: u(undefined, "") },
        { json: "codec", js: "codec", typ: u(undefined, "") },
        { json: "crf", js: "crf", typ: u(undefined, 0) },
        { json: "faststart", js: "faststart", typ: u(undefined, true) },
        { json: "preset", js: "preset", typ: u(undefined, "") },
    ], "any"),
    "Subtitle": o([
        { json: "dual_line", js: "dual_line", typ: u(undefined, true) },
        { json: "enabled", js: "enabled", typ: true },
        { json: "font", js: "font", typ: u(undefined, "") },
        { json: "highlight_color", js: "highlight_color", typ: u(undefined, "") },
        { json: "position_ratio", js: "position_ratio", typ: u(undefined, 3.14) },
        { json: "size_ratio", js: "size_ratio", typ: u(undefined, 3.14) },
        { json: "source_lang", js: "source_lang", typ: u(undefined, "") },
        { json: "style", js: "style", typ: u(undefined, r("Style")) },
        { json: "target_lang", js: "target_lang", typ: u(undefined, "") },
        { json: "words_ref", js: "words_ref", typ: u(undefined, "") },
    ], "any"),
    "Watermark": o([
        { json: "enabled", js: "enabled", typ: true },
        { json: "opacity", js: "opacity", typ: u(undefined, 3.14) },
        { json: "position", js: "position", typ: u(undefined, r("Position")) },
        { json: "size_ratio", js: "size_ratio", typ: u(undefined, 3.14) },
        { json: "text", js: "text", typ: u(undefined, "") },
    ], "any"),
    "ActorType": [
        "admin",
        "system",
        "user",
    ],
    "Fit": [
        "center_crop",
    ],
    "Ratio": [
        "1:1",
        "16:9",
        "21:9",
        "4:3",
        "4:5",
        "9:16",
    ],
    "Style": [
        "karaoke",
        "minimal",
    ],
    "Position": [
        "bottom_left",
        "bottom_right",
        "top_left",
        "top_right",
    ],
};
