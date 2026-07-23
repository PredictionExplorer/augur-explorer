# Operations

How to run, deploy, and maintain the RWCG backend.

## Local development

Prerequisites: Go 1.26+, Docker (for Postgres and integration tests).

```bash
# Start a migrated database
docker compose up db migrate

# Environment
cp .env.example .env   # fill in RPC_URL at minimum
export $(grep -v '^#' .env | xargs)
# The database can be one DATABASE_URL instead of the PGSQL_* variables.
# Services log structured records to stdout: text by default (dev), JSON
# with LOG_FORMAT=json (production); LOG_LEVEL selects the minimum level.

# Build everything into ./bin
make build

# Run the API server
./bin/apiserver

# Run the indexers (need RPC_URL for Arbitrum)
./bin/cg-etl
./bin/rw-etl
```

Or run the full stack in containers: `docker compose --profile etl up`.
The apiserver Compose service is healthy only after `/readyz` confirms its
database. Run the D6-safe, no-credentials whole-stack probe with:

```bash
docker compose --profile smoke up --build --abort-on-container-exit \
  --exit-code-from smoketest smoketest
```

## Testing

```bash
make test              # unit tests, race detector, shuffled
make test-integration  # + testcontainers Postgres (needs Docker)
make lint              # golangci-lint
make vuln              # govulncheck dependency scan
```

Integration tests get a fresh PostgreSQL container with all migrations
applied via `internal/testdb`; they skip automatically when Docker is not
available.

## Database migrations

The schema is owned by goose migrations in `db/migrations`.

```bash
# Apply pending migrations
GOOSE_DBSTRING="postgres://user:pass@host:5432/db" make migrate-up

# Create a new migration
goose -dir db/migrations create add_my_table sql
```

### Adopting migrations on the existing production database

The baseline migrations (00001–00003) describe the schema production already
has. On first rollout, record them as applied without executing them:

```sql
-- run once against the production database
CREATE TABLE IF NOT EXISTS goose_db_version (
    id SERIAL PRIMARY KEY,
    version_id BIGINT NOT NULL,
    is_applied BOOLEAN NOT NULL,
    tstamp TIMESTAMP DEFAULT now()
);
INSERT INTO goose_db_version (version_id, is_applied)
VALUES (0, true), (1, true), (2, true), (3, true);
```

All subsequent migrations are applied normally with `goose up`.

## Production deployment (systemd)

Unit files live in `deploy/systemd/`. Layout on the host:

- Binaries in `/opt/rwcg/bin` (`make build`, then copy `bin/*`).
- Config in `/etc/rwcg/*.env` (verified against the code by
  `internal/config`'s `.env.example` test; set `LOG_FORMAT=json` here).
- Logs go to stdout and journald owns persistence — the legacy
  `$HOME/ae_logs` file logs (`webserver_*.log`, `*_info/_error/_db.log`) are
  gone since the §8.3 slog migration.

```bash
sudo cp deploy/systemd/*.service deploy/systemd/*.timer /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable --now rwcg-apiserver@cosmic rwcg-apiserver@rwalk
sudo systemctl enable --now rwcg-etl@cg rwcg-etl@rw
sudo systemctl enable --now rwcg-notibot rwcg-imggen-monitor.timer \
  rwcg-loganomaly.timer
```

These replace the legacy `run-loop-*.sh` nohup scripts (still present next to
each command for the transition; delete them once systemd is live).
`rwcg-loganomaly.timer` is a persistent five-minute timer over a bounded
one-shot service. The service first exports the configured apiserver journal,
then atomically generates the anomaly file; a failed export or scan exits
nonzero and is retried by systemd. Configure its non-secret inputs in
`/etc/rwcg/loganomaly.env`:

```bash
LOGANOMALY_JOURNAL_UNIT=rwcg-apiserver@cosmic
LOGANOMALY_SINCE=-2d
LOGANOMALY_IN=/var/lib/rwcg/loganomaly/webserver_cosmic_journal.log
LOGANOMALY_OUT=/var/lib/rwcg/loganomaly/webserver_anomalies.log
LOGANOMALY_MIN_STATUS=500
LOGANOMALY_KEEP=50
```

The `rwcg` account must be able to read its apiserver journal. Keep both paths
under the unit-managed `/var/lib/rwcg/loganomaly` state directory so
`ProtectSystem=strict` and `ProtectHome=read-only` remain effective.

### Reading logs

```bash
journalctl -u rwcg-apiserver@cosmic -f          # follow the API server
journalctl -u rwcg-etl@cg --since "1 hour ago"  # ETL records
journalctl -u rwcg-apiserver@cosmic -o cat | jq 'select(.msg=="request")'
```

Every service emits one JSON record per line (`LOG_FORMAT=json`): access
lines are `msg="request"` with `route`, `status`, `duration_ms` and `ip`
fields (`bytes` counts wire bytes — compressed size when the client
negotiated gzip); database traces carry `component=db`; startup prints a
`build info` record (version, commit, build date) followed by an
`effective configuration` record with secrets redacted.

## HTTP edge behavior (compression, caching, build identity)

The API server terminates TLS itself — there is no reverse proxy — so the
response-edge policy lives in the shared middleware chain
([ADR-0007](adr/0007-http-edge.md)):

- **Listener timeouts**: every listener bounds header reads (10s), full
  request reads (30s — bodies are small JSON) and idle keep-alive
  connections (120s). There is deliberately **no WriteTimeout**: responses
  are fully buffered, but several frozen v1 routes deliver multi-megabyte
  unpaginated arrays and a write cap would sever them to legitimately slow
  clients. Handler *execution* time is bounded separately by the request
  processing deadline (next section); revisit at v1 removal.
- **TLS floor**: the public HTTPS listeners require TLS 1.2 or newer
  (Go's server default still admits 1.0/1.1). Handshake failures and other
  `http.Server` records land on the process logger at `WARN` instead of
  stderr, so scanner noise is filterable with `LOG_LEVEL=error`.
- **Request-body cap**: every request body is bounded at **1 MiB** by the
  shared middleware (the largest legitimate bodies — ranking votes — are a
  few hundred bytes). A request declaring a larger `Content-Length` is
  answered `413` immediately; one that streams past the cap fails at the
  first read inside the consuming layer. The 413 shape follows the route
  family: RFC 9457 `application/problem+json`
  (`…/problems/request-too-large`, declared on the three body-accepting v2
  operations) under `/api/v2/`, the legacy `{"status":0,"error":…}`
  envelope everywhere else, and the frozen v1 handlers keep their
  `400 {"error":"http: request body too large"}` bind shape for undeclared
  overflows. The FAQ proxy additionally caps buffered upstream responses at
  4 MiB (`502` beyond it). There are no environment switches: the limits
  are fixed, like the rest of the edge policy.
- **Compression**: JSON/text responses ≥ 1 KiB are gzip-encoded for clients
  sending `Accept-Encoding: gzip`. Images and video are never re-compressed.
  Every response carries `Vary: Accept-Encoding`.
- **Conditional requests**: successful JSON/text GETs carry a weak `ETag`
  and default `Cache-Control: no-cache` (store, but revalidate); clients
  re-sending the tag via `If-None-Match` get an empty `304 Not Modified`.
  `/images` responses keep their `max-age=3600` policy
  (`WEBSRV_IMAGE_NO_CACHE=1` switches to `no-store` for development), and
  files keep `http.ServeFile`'s `Last-Modified` handling. There are no new
  environment switches: the policy is fixed.
- Expect `3xx` samples in `rwcg_http_requests_total` for API routes now —
  they are revalidations, not redirects.

Verify from a shell:

```bash
curl -sI --compressed https://api.example/api/cosmicgame/rounds/list/0/10 \
  | grep -iE 'content-encoding|etag|cache-control|vary'
curl -sI -H 'If-None-Match: W/"<tag from the previous call>"' \
  https://api.example/api/cosmicgame/rounds/list/0/10   # HTTP/2 304
```

Every binary reports its build identity: `GET /version` on the API server,
`--version` on every command, and the `build info` startup log record.
Identity is stamped by `make build` and the Dockerfile (`git describe`,
commit, build date); ad-hoc `go build` binaries fall back to the
toolchain's embedded VCS metadata.

```bash
curl -s https://api.example/version | jq   # {"version":…,"commit":…,…}
/opt/rwcg/bin/apiserver --version          # same identity on the host
```

## Time bounds (no unbounded waiting)

Since the bounded-time change (D22 in `docs/MODERNIZATION.md`) every unit of
work is bounded in time; a black-holed dependency degrades to an explicit,
observable failure instead of a silent hang. Like the rest of the edge
policy the bounds are fixed constants — there are no environment switches.

- **API request deadline — 30s.** The shared middleware chain puts one
  deadline on every request context, so context-aware handler work
  (PostgreSQL queries, contract reads) fails with `DeadlineExceeded` once
  the budget is spent. The failure renders as **503** — an RFC 9457
  `…/problems/request-timeout` problem under `/api/v2/`, the legacy
  `{"status":0,"error":"request timed out after 30s"}` envelope everywhere
  else — and increments `rwcg_http_request_timeouts_total{method,route}`.
  A handler that finishes just past the deadline still delivers its result;
  only post-deadline internal errors are reinterpreted. The FAQ proxy is
  the one exempt family: its LLM upstream is bounded by the proxy's own
  180s client timeout and 4 MiB response cap. 30s is 60x the store's 500ms
  slow-query warning threshold — a request that needs more is a fault.
- **Server-side statement timeouts (defense in depth).** The services set
  PostgreSQL session GUCs on every pool connection: `statement_timeout`
  30s on the apiserver (matching the request deadline) and 60s on
  cg-etl/rw-etl/notibot/imggen-monitor; `idle_in_transaction_session_timeout`
  5m everywhere. These fire only when client-side cancellation never
  reaches the server (SQLSTATE `57014` / `25P03` in the logs — investigate,
  they mean the primary bound failed). Operator CLIs (`opsctl`, `cgctl`,
  `rwctl`, `freezer-verify`) deliberately keep the server defaults: their
  heavy statements are legitimate and Ctrl-C cancels them.
- **ETL chain RPC — 60s per call.** The indexer engine wraps its RPC client
  so every call (`eth_getLogs`, header/transaction/receipt fetches — the
  historical backfill included) carries a deadline, even mid-batch on the
  shutdown-immune context. A black-holed RPC endpoint now produces ordinary
  batch failures — backoff, adaptive batch shrinking, breaker after 10, exit,
  systemd restart — where it used to hang the ETL forever without a single
  failure recorded (and SIGTERM could not interrupt it). In-transaction
  contract reads (donated-NFT `tokenURI`, donation info) and the startup
  contract-parameter sync are bounded at 15s per call.
- **Everything else.** Twitter posts (2m — sized to video uploads),
  WhatsApp sends (30s), Discord REST calls (30s), autobid's whole JSON-RPC
  transport (30s per exchange), srvmonitor's RPC probes (10s, matching its
  other probes) and the opsctl chain tools (90s per exchange, above
  logscan's 60s per-fetch bound) are all bounded; `internal/ethtx` receipt
  waits were already capped at 2m.

Alert when timeouts appear at all — they are rare by construction:

```promql
sum by (route) (rate(rwcg_http_request_timeouts_total[15m])) > 0
```

## v1 deprecation headers and the sunset date

The frozen v1 API (everything under `/api/cosmicgame/` and
`/api/randomwalk/` except the FAQ proxy) is deprecated in favor of
`/api/v2`; the policy lives in `internal/api/policy.V1Deprecated` and a
drift test pins it to the `deprecated: true` flags in `docs/openapi.yaml`.
Every deprecated response — including errors, 304 revalidations and 404s
under those prefixes — carries:

- `Deprecation: @1784160000` (RFC 9745; the 2026-07-16 deprecation moment),
- `Link: <…/docs/api-v2-migration.md>; rel="deprecation"` (the v1→v2 map),
- `Sunset: <HTTP-date>` (RFC 8594) **only when** `V1_SUNSET_AT` is set.

Leave `V1_SUNSET_AT` unset until the ADR-0005 sunset gates are met (known
consumers migrated, 30 consecutive zero-traffic days, announced not-before
date). Once the date is announced, set it in the apiserver env file
(RFC 3339, e.g. `V1_SUNSET_AT=2027-01-01T00:00:00Z`) and restart; removal
of the v1 layer itself is a code change gated on the same criteria.

```bash
curl -sI https://api.example/api/cosmicgame/rounds/list/0/10 \
  | grep -iE 'deprecation|sunset|link'
```

### Measuring the sunset gate (v1 traffic)

`rwcg_http_requests_total` carries a `deprecated` label derived from the
same `policy.V1Deprecated` policy as the headers, so metric and header can
never disagree. `deprecated="true"` covers every request under the
deprecated prefixes — matched routes *and* 404s (`route="unmatched"`), so
probing bots don't hide in unrouted paths. The exempt surfaces (health,
`/version`, `/api/v2/*`, the FAQ proxy, the contract-pinned metadata
routes) count `deprecated="false"`.

Current v1 request rate, per route:

```promql
sum by (route) (rate(rwcg_http_requests_total{deprecated="true"}[1h]))
```

The D6 zero-traffic check — total v1 requests over the last 30 days,
excluding documented probes — must evaluate to zero:

```promql
sum(increase(rwcg_http_requests_total{deprecated="true"}[30d]))
```

**Documented probes:** monitoring that intentionally hits v1 (e.g. an
uptime check pinned to a legacy route) must either move to `/healthz` /
`/api/v2` before the measurement window starts, or be excluded explicitly
by its route label — record any such exclusion here with the route and the
reason. There are currently **no** documented probes: the expression above
must read exactly zero.

Repository-owned probes enforce that invariant:

- bare `opsctl smoketest` runs all 97 safe v2 GET operations; use explicit
  `--suite=v1` only for an operator-invoked compatibility regression, never
  as a scheduled uptime probe;
- `opsctl smoketest --suite=operational` needs no database credentials and
  checks `/healthz`, `/readyz`, `/version`, and stable DB-backed v2 reads;
- srvmonitor rejects configured internal or public URLs under either
  deprecated v1 prefix at startup. Use `/readyz` internally and a stable
  resource such as `/api/v2/cosmicgame/rounds?limit=1` publicly.

The full v2 suite intentionally fails on a contract-state 503 because that is
user-visible dependency degradation. The operational uptime suite excludes
`contracts/configuration`, `contracts/balances`, and `rounds/current*`: these
correctly return 503 while their RPC-backed cache is unavailable and must not
make the process/DB uptime probe flap.

Dashboards written before the label existed keep their shape with
`sum without (deprecated) (...)`; the label was added 2026-07-16 and old
series simply lack it.

## Security posture

- Read API is public by design; every route is rate limited per client IP.
- Mutating routes (v1 `ban_bid`/`unban_bid`, v2
  `/cosmicgame/moderation/banned-bids`, and ranking `match`) require
  `X-Admin-Key` / `X-Ranking-Admin-Key` and fail closed when the key env var
  is missing. V2 compares fixed-width key hashes in constant time and applies
  independent per-operation write buckets. Generate keys with
  `openssl rand -hex 32`.
- `/metrics` and `/debug/pprof` bind only to `METRICS_ADDR` — keep it on
  localhost or a private interface.
- Secrets come exclusively from environment files; nothing is committed.

## Monitoring

- `srvmonitor` — terminal dashboard checking RPC nodes, DB, D6-safe web API
  paths, disk, image server and anomaly-feed freshness; optional
  WhatsApp/Android alerts (see `cmd/srvmonitor/README.md`).
- `loganomaly` — scans the API access log for 5xx/panics, feeding srvmonitor.
  It understands slog JSON records (production), slog text lines and the
  legacy `[GIN]` lines still present in older capture files. Every successful
  generation begins with `#TS=<unix-seconds>`, even when no anomalies were
  found. `srvmonitor` strips that metadata, shows its age and emits a stable
  repeated alarm once it exceeds `ANOMALY_STALE_SECS` (default 1800);
  markerless/malformed-marker files remain legacy-compatible with no stale
  claim. The systemd one-shot above owns the journal export and generation;
  point `ANOMALY_REMOTE_FILE` at its `LOGANOMALY_OUT`.
- Prometheus can scrape `METRICS_ADDR/metrics`; alert on
  `rwcg_http_requests_total{status="5xx"}` and request-duration percentiles.
  Since the stdlib-router migration the `route` label uses ServeMux syntax
  (`/api/.../{param}` instead of `/api/.../:param`) — update dashboards that
  filter on route values. `rwcg_http_request_timeouts_total{method,route}`
  counts requests that hit the 30s processing deadline (rendered 503) —
  see the time-bounds section for the alert expression.
- The apiserver and both ETLs also export their shared pgx pool as
  `rwcg_db_pool_*`: saturation gauges (`acquired_conns`, `idle_conns`,
  `constructing_conns`, `total_conns` against the `max_conns` bound) and
  cumulative counters (`acquires_total`, `acquire_duration_seconds_total`,
  `empty_acquires_total` + `empty_acquire_wait_seconds_total` — the
  pool-exhaustion signals — `canceled_acquires_total`, `new_conns_total`,
  `max_lifetime_destroys_total`, `max_idle_destroys_total`). Useful alerts:

  ```promql
  # Queries queueing because the pool is exhausted (per second)
  rate(rwcg_db_pool_empty_acquires_total[5m]) > 0

  # Mean time spent waiting for a connection
  rate(rwcg_db_pool_acquire_duration_seconds_total[5m])
    / rate(rwcg_db_pool_acquires_total[5m])

  # Sustained saturation: every connection checked out
  rwcg_db_pool_acquired_conns >= rwcg_db_pool_max_conns
  ```
- The ETLs (`cg-etl`, `rw-etl`) honor `METRICS_ADDR` too (each process needs
  its own port on a shared host) and expose `rwcg_etl_last_block` (the
  processing watermark — alert when it stops advancing),
  `rwcg_etl_events_total{type}`, `rwcg_etl_batch_duration_seconds`,
  `rwcg_etl_reorgs_total` and `rwcg_etl_batch_failures_total{stage}`.
- ETL failure behavior (since the indexer-engine migration): a failed batch —
  RPC or database — is retried in-process with exponential backoff instead of
  crashing; the process exits non-zero only after 10 consecutive failures
  (systemd restarts it). Since the bounded-time change a *hung* RPC call is
  cut off at 60s and counts as an ordinary failure, so a black-holed
  endpoint trips the breaker instead of stalling the process silently. A
  restart never skips events: the watermark only advances past fully
  processed blocks, and the engine re-reads it from
  `cg_proc_status`/`rw_proc_status` at startup (rewind it with the ETL
  stopped if you need a manual replay).
- ETL atomicity (since the transactional-ingestion change, ADR-0010): each
  block's rows, domain writes and watermark update commit in one database
  transaction, so a crash or failure can no longer leave a partially applied
  block behind. `rwcg_etl_batch_failures_total{stage}` gained the `commit`
  stage (transaction begin/commit failures); the other stages — `fetch`,
  `chain_head`, `block`, `transaction`, `event_log`, `process`, `watermark`
  — keep their meanings.
- `cgctl backfill-dao-evtlog` uses the same one-transaction-per-block boundary
  for its layer-1 writes. A failed block contributes no durable rows or
  success stats; prior blocks stay committed and rerunning the range is
  idempotent.
- `GET /readyz` returns 503 whenever the database is unreachable — wire it
  into your load balancer health checks.

### Refreshing the production RLP replay corpus

The committed
`internal/indexer/{cosmicgame,randomwalk}/testdata/rlp_corpus.jsonl` files use
the same fields as `arch_evtlog`, with `log_rlp` rendered as 0x-prefixed hex.
`internal/rlpcorpus` owns the strict format: malformed JSON/hex/RLP, unknown
fields, inconsistent contract/topic data, duplicate `(tx_hash, log_index)`
identities, split transaction groups, descending sibling log indexes and
lines over 2 MiB are rejected before PostgreSQL is touched.

Export one project into a fresh scratch database. Select a small set of
representative transaction hashes during review, then let the native exporter
include every archived sibling log and render canonical JSONL:

```bash
opsctl archive export \
  --project cosmicgame \
  --src "$PRODUCTION_DATABASE_URL" \
  --dst "$SCRATCH_DATABASE_URL"

opsctl archive corpus-export \
  --db "$SCRATCH_DATABASE_URL" \
  --project cosmicgame \
  --tx-hash "$TX_HASH_1" \
  --tx-hash "$TX_HASH_2" \
  > /tmp/cosmicgame-rlp-corpus.jsonl
```

Use `--project randomwalk` and a fresh destination for the RandomWalk corpus.
The transaction hashes are repeatable and their order is preserved; a missing
transaction, malformed archive row or partial query is a hard error, and the
summary goes to stderr so stdout remains clean JSONL. Review the public-chain
payloads, then copy the result into the matching `testdata/rlp_corpus.jsonl`.
Validate the format, exact-byte direct replay and production `Engine.Run`
path before committing:

```bash
go test ./internal/rlpcorpus ./internal/testutil
go test -tags=integration \
  ./internal/indexer/cosmicgame ./internal/indexer/randomwalk \
  -run RLPCorpus
```

Fixture-derived rows keep this gate active without production access.
Replacing or extending them with production-exported rows completes the
remaining §4.3 production replay gate; never label synthetic rows as
production samples.

## Routine tasks

| Task | Command |
|------|---------|
| Verify DB contents against chain | `opsctl db verify`, `opsctl db evtlog-diff` |
| Export/verify RLP archives and corpora | `opsctl archive export` / `corpus-export` / `verify` / `node-fill` |
| Check NFT asset presence | `opsctl assets inventory`, `imggen-monitor` |
| Regenerate thumbnails | `opsctl assets gen-thumbnails` |
| Canonical v2 API smoke test | `opsctl smoketest` |
| D6-safe uptime/Compose probe | `opsctl smoketest --suite=operational` |
| Frozen v1 compatibility regression | `opsctl smoketest --suite=v1` (manual only; generates deprecated traffic) |
| Contract admin (owner ops) | `cgctl --help` (bid, claim-prize, set-* commands) |
| Social/notification tools | `rwctl --help` (tweet-mints, notify-bot, ...) |
