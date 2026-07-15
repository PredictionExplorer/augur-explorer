# ADR-0007: HTTP edge policy — compression, conditional requests, version reporting

- Status: Accepted
- Date: 2026-07-13

## Context

The API serves large, highly redundant JSON documents (round histories,
statistics, user profiles) directly to browsers with no reverse proxy in
front: the Go binary terminates TLS and owns every response header. Until
this decision the server never compressed a response, never emitted a cache
validator, and dashboards re-downloaded identical multi-hundred-kilobyte
bodies on every poll. Deployments were also unidentifiable at runtime — no
endpoint or flag reported what build was running.

Frontends poll endpoints whose content changes only when a new chain event
arrives, so most polls return bytes the client already holds. The fix
belongs in the shared middleware chain, not in handlers: the v1 surface is
frozen byte-for-byte behind the parity suite, and v2 handlers should stay
transport-agnostic.

## Decision

### Compression (`common.Compress`)

Every response passes through a gzip middleware (stdlib `compress/gzip`,
level 6, pooled writers). A response is compressed only when all hold:

- the request negotiates gzip (`Accept-Encoding` with RFC 9110 q-values;
  `x-gzip` alias and `*` wildcard honored; malformed elements refuse
  conservatively),
- the status is 200 — error envelopes are small, and 206/304 must keep
  their exact representation,
- no `Content-Encoding`/`Content-Range` is already set,
- the `Content-Type` is text-like (JSON, text, XML, SVG — never raster
  images or video, which are pre-compressed formats),
- the body reaches 1 KiB (below that the gzip framing and CPU buy nothing).

Every response gains `Vary: Accept-Encoding` (deduplicated) so shared
caches key on the negotiation dimension. Compressed responses drop any
handler-set `Content-Length`.

### Conditional requests (`common.ConditionalETag`)

Successful GET/HEAD responses with hashable (text-like) bodies gain a weak
validator, `W/"<truncated SHA-256>"`, computed over the identity body.
A request presenting a matching `If-None-Match` (weak comparison, list and
`*` forms per RFC 9110 §8.8.3.2) is answered `304 Not Modified` with the
body dropped. `Cache-Control: no-cache` (store, but revalidate) is applied
only when no other layer chose a policy — the static-asset middleware's
`no-store`/`max-age=3600` win by construction.

The tag is weak because the same entity travels gzip- or identity-encoded
depending on negotiation; the validator is computed inside the compression
layer so it is stable across encodings. File responses served by
`http.ServeFile` (images, `/static` ABI files) keep their native
`Last-Modified` conditional handling and are never double-validated: the
middleware skips any response already carrying a validator.

The middleware buffers hashable bodies to hash before the header section is
sent. `Context.JSON` already materializes every body in memory, so the
added cost is one copy; the 32 KiB benchmark round-trip is ~17 µs against
170 µs+ for the cheapest database read underneath it.

### Chain order

`CORS → Recovery → AccessLog → Compress → RateLimit → metrics →
ConditionalETag → routes`. Compression sits inside the access log so logged
`bytes` are wire bytes; the validator sits innermost so it hashes identity
bodies and its 304s flow out through the compression layer untouched. Both
writers keep the `httpx.ResponseWriter` contract (`Status`/`Written`/
`Size`/`Unwrap`), and neither flushes anything when a handler panics, so
`Recovery` still answers a clean 500.

### Transport headers stay out of the OpenAPI contracts

`Vary`, `Content-Encoding`, `ETag` and `Cache-Control` are transport-level
concerns like the CORS headers: they are documented here and in
[operations.md](../operations.md), not declared per-operation in
`openapi.yaml`/`openapi-v2.yaml`. kin-openapi response validation ignores
undeclared headers, so the v2 golden suite continues to validate every
response. The parity goldens pin status/content-type/body and never
negotiate gzip or send validators, so the v1 freeze is unaffected.

### Version reporting

`internal/version` resolves the build identity — version/tag, commit, build
date, Go toolchain — from `-ldflags -X` stamping (Makefile, Dockerfile)
with `debug.ReadBuildInfo` VCS fallback for plain `go build`. Every binary
answers `--version` (cobra CLIs via `root.Version`, plain mains via
`version.HandleFlag` before flag parsing); every service logs one
structured `build info` record at startup; the API server serves
`GET /version` (declared in `openapi.yaml`, guarded by the route-drift
test; excluded from the parity sweep because the values are
build-dependent, with a dedicated shape test instead).

## Consequences

- Large JSON responses shrink ~10× for gzip clients; unchanged resources
  revalidate as empty 304s. No handler changed and no golden regenerated.
- Access-log `bytes` now reports wire (compressed) bytes.
- 304 responses appear in metrics as `3xx` for routes that previously only
  produced `2xx`.
- Every hashable 200 pays one body copy plus a SHA-256; compressed
  responses pay the gzip CPU. Both are far below the database time of the
  routes they ride on (see [benchmarks.md](../benchmarks.md)).
- Operators can identify any deployment via `/version`, `--version` or the
  startup log record; production builds must go through `make build` or
  the Dockerfile to be stamped (unstamped binaries degrade to VCS metadata
  or `unknown`, never fail).
- If a CDN or reverse proxy is ever introduced, `Vary: Accept-Encoding`
  and the weak validators are already cache-correct; response compression
  can then be delegated outward by deleting one middleware line.
