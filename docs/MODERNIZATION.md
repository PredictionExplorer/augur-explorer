# Modernization Roadmap — Go-Native Rewrite

This is the working document for turning this codebase into world-class,
idiomatic Go. It tracks everything: what has been done, what remains, which
tests protect each change, and how to verify every step. The end state is a
backend that behaves the same (or better) with a dramatically cleaner
implementation and a redesigned v2 API that the frontend migrates onto.

**Status legend:** `[ ]` todo · `[x]` done · `[~]` in progress (note the PR/date)

---

## 1. How to use this document

1. Work top-to-bottom within a phase; phases may overlap only where noted.
2. Check items off as they land. Add the date and commit hash for anything
   non-trivial in the [Progress log](#12-progress-log).
3. **The golden rule: no rewrite lands without characterization tests.**
   Phase 0 builds a safety net that pins current behavior; every subsequent
   rewrite PR must keep that suite green. If a behavior change is
   *intentional*, update the golden files in the same PR and say so in the
   PR description.
4. **Definition of done** for every item, unless stated otherwise:
   - `go build ./...`, `go vet ./...`, `go test ./...` pass
   - `golangci-lint run` introduces no new issues (target: zero, see §10)
   - `make test-integration` passes (Docker required)
   - Parity suite (§4.1) passes once it exists
   - New/changed exported symbols have doc comments
5. Re-measure the [metrics dashboard](#2-metrics-dashboard) at the end of
   each phase and update the table — the numbers are the honest progress bar.

---

## 2. Metrics dashboard

Measured 2026-07-06 (commit `85941dba`). Update after each phase.

| Metric | Baseline (start of project) | Current | Target | How to measure |
|---|---|---|---|---|
| Hand-written Go LOC (`cmd/` + `internal/`) | ~60,000 (plus 124k frozen legacy) | 57,660 | n/a (informational) | `rg --files internal cmd -g '*.go' -0 \| xargs -0 wc -l \| tail -1` |
| snake_case functions | ~700 | **656** (store 387, api 135, cg-etl 89, notibot 24, rw-etl 11) | **0** | `rg "^func (\([^)]+\) )?[A-Za-z]+_[A-Za-z0-9_]*\(" --type go -c internal cmd` |
| `os.Exit` in library code (`internal/`) | ~560 | **~485** (store/cosmicgame ~410, store/randomwalk 75, api 3) | **0** (allowed only in `cmd/*/main.go` startup) | `rg -c "os\.Exit" internal` |
| Dot-import files | ~70 | **21** | **0** | `rg -l '^\s*\. "github' --type go` |
| Package-level mutable globals (api + etl) | ~120 | ~80 (state.go ~50, cg-etl ~25, rw-etl ~8) | ~0 (DI everywhere) | manual review per package |
| golangci-lint issues | 433 (first run) | **180** | **0** | `golangci-lint run` |
| Test files | 17 | **19** | 100+ | `rg --files -g '*_test.go' \| wc -l` |
| Fuzz targets | 0 | **0** | **25+** (see §4.4) | `rg -l "func Fuzz" --type go` |
| Coverage on `internal/` | <5% | ~8% | **≥70%** enforced in CI | `go test -coverprofile=c.out ./internal/... && go tool cover -func=c.out \| tail -1` |
| Queries on sqlc | 0 | 8 (layer1 pattern) | all static queries | count in `internal/store/queries/*.sql` |
| Routes on stdlib router | 0 | 0 (gin) | all (v1 compat + v2) | n/a |
| `context.Context` on store methods | 0% | ~5% (base pkg only) | 100% | manual |

---

## 3. Completed foundation (for context)

Everything below landed in commit `85941dba` — the platform the rewrite
builds on:

- [x] Legacy `previous-code/` tree deleted (~450k LOC); live daemons rescued
- [x] Standard `cmd/` + `internal/` layout; single root Makefile; `bin/` output
- [x] Go 1.26.4; go-ethereum 1.17.4; gin 1.12; pgx/v5 driver (keepalive + retry preserved)
- [x] ~90 script mains → `cgctl` / `rwctl` / `opsctl` cobra CLIs; orphans deleted
- [x] `/black/*` HTML explorer removed (~4,200 LOC + 150 templates); JSON-only
- [x] ETL monoliths split into topical files; graceful shutdown; error-log bug fixed
- [x] Base store (`internal/store/*.go`): errors returned, zero `os.Exit`, no lib/pq
- [x] goose migrations (`db/migrations/`) + testcontainers harness (`internal/testdb`)
- [x] sqlc scaffolding (`sqlc.yaml`, `internal/store/queries/`, generated `sqlcgen`)
- [x] Fail-closed admin auth + per-IP rate limiting (+ unit tests)
- [x] `/healthz`, `/readyz`, Prometheus metrics, private pprof listener
- [x] CI: build, tidy-check, race tests, integration tests, lint, govulncheck; Renovate
- [x] Docker multi-stage build, `compose.yaml` dev env, systemd units
- [x] Docs: README, architecture, operations, CONTRIBUTING, 4 ADRs, `.env.example`
- [x] OpenAPI 3.1 spec covering all 187 v1 routes (`docs/openapi.yaml`)
- [x] 210 lint issues fixed in leaf packages; 3 real bugs found and fixed

---

## 4. Phase 0 — Test safety net (build this FIRST)

Purpose: pin current behavior so the rewrite can be verified mechanically.
These tests must pass against the code as it is today, then keep passing
(or be deliberately updated) as each piece is rewritten.

### 4.1 API parity suite (characterization tests)

Golden-file HTTP tests: seed a testcontainers database with a fixed fixture
set, start the real router via `httptest`, hit every route, snapshot the
JSON to `testdata/golden/`. The suite is the contract for the v1 API freeze.

- [ ] `internal/api/apitest/` package: fixture loader + golden-file helpers
      (`-update` flag to regenerate; byte-stable via canonical JSON marshal)
- [ ] Fixture dataset: SQL seed files under `internal/api/apitest/testdata/seed/`
      covering ≥3 rounds, ≥5 bidders, every prize type (main, raffle ETH/NFT,
      endurance, chrono warrior, lastcst), donations (ETH/ERC20/NFT + claims),
      staking (CST + RWalk, stake/unstake/rewards), banned bids, RandomWalk
      mints/offers/sales/withdrawals/name-changes, ranking votes with Elo state
- [ ] Route enumeration test: parse `docs/openapi.yaml` and assert every
      documented path is registered (and vice versa) — spec can never drift
- [ ] Golden tests for all CosmicGame GET routes (~120 paths)
- [ ] Golden tests for all RandomWalk GET routes (~45 paths)
- [ ] Golden tests for `/metadata/:token_id` host dispatch (both hosts) and
      `/cg/metadata/:token_id`
- [ ] Mutation-route tests: `ban_bid`/`unban_bid`/`token-ranking/match`
      (auth 503/401/success matrix), `add_game` signature verification paths
- [ ] Error-shape tests: invalid params on representative routes pin the
      `{"status":0,"error":...}` envelope
- [ ] Wire into `make test-integration` and CI

### 4.2 Store integration suite

One test file per store file; every public query function called against the
seeded database at least once, asserting golden rows. This is what lets the
sqlc/pgx rewrite (§5) proceed file-by-file with confidence.

CosmicGame (`internal/store/cosmicgame/`):

- [ ] `inserts.go` (73 funcs — cover via ETL fixture replay in §4.3 plus direct calls for edge cases)
- [ ] `deletes.go` (72 funcs — insert-then-delete round-trips; verify trigger side effects)
- [ ] `statistics.go` (20 funcs)
- [ ] `user-specific.go` (20 funcs)
- [ ] `staking.go` (18 funcs)
- [ ] `admin_events_resolve.go` (14 — extend existing test)
- [ ] `eth-donations.go` (13)
- [ ] `bidding.go` (13 — extend existing `bidding_v2_test.go`)
- [ ] `contract_params.go` (12 — extend existing test)
- [ ] `tokens-erc721.go` (11)
- [ ] `erc20-donations.go` (9)
- [ ] `bidding_analytics.go` (8)
- [ ] `raffle-eth.go` (8)
- [ ] `nft-donations.go` (8)
- [ ] `tokens-erc20.go` (5)
- [ ] `cosmicgame.go` (4)
- [ ] `raffle-nft.go` (3), `main-prize.go` (3), `banned_bids.go` (3)
- [ ] `prize-history.go` (2), `marketing.go` (2), `admin_events.go` (2)

RandomWalk (`internal/store/randomwalk/`):

- [ ] `randomwalk.go` (29 funcs)
- [ ] `randomwalk_api.go` (24 funcs)
- [ ] `ranking.go` (9 funcs — include Elo transaction semantics)

Base (`internal/store/`):

- [ ] `lookups.go`, `blockchain.go`, `blockchain_insert.go`, `archive.go`
      (round-trip block/tx/evt_log insertion incl. reorg path)
- [ ] Trigger behavior tests: `cg_bidder`/`cg_glob_stats`/`cg_winner`
      aggregates update correctly on insert/delete (the plpgsql functions in
      migration 00002 are load-bearing and currently untested)

### 4.3 ETL decode fixtures (golden events)

- [ ] `cmd/cg-etl/fixtures_test.go`: for each of the ~80 event types in
      `events_registry.go`, build a synthetic `types.Log` by ABI-packing known
      values (topics + data), run the decode path, assert the decoded struct
- [ ] Golden DB-state test: process a scripted sequence of ~30 events through
      `process_single_event` against testdb; snapshot the resulting rows
- [ ] Same for `cmd/rw-etl` (9 event types)
- [ ] RLP replay test: feed archived `evt_log.log_rlp` samples (checked-in
      testdata, ~50 real events exported once via `opsctl archive export`)
      through decode and compare against their known DB rows
- [ ] Reorg simulation test: two blocks with same height/different hash
      exercise the chain-split handling in `internal/etl`

### 4.4 Fuzz test inventory (the panic-hunting fleet)

Native Go fuzzing (`go test -fuzz`). Every target gets a seed corpus in
`testdata/fuzz/` (committed) and runs in the nightly CI fuzz job (§4.6).
Invariant unless stated otherwise: *never panics, never hangs*.

Decoders (highest value — they consume chain data):

- [ ] `FuzzReceiptsDecode` — `internal/freezer/decode`: arbitrary bytes → RLP receipt decode
- [ ] `FuzzArbitrumLegacyDecode` — `internal/freezer/decode`: the Arbitrum-specific format
- [ ] `FuzzFreezerIndexRead` — `internal/freezer`: corrupt index/data-file bytes
- [ ] `FuzzEventDecodeCG` — `cmd/cg-etl`: fuzz `types.Log{Topics,Data}` through
      every registered decode branch (drive via the signature registry)
- [ ] `FuzzEventDecodeRW` — `cmd/rw-etl`: same for RandomWalk events
- [ ] `FuzzEvtlogRLP` — `internal/etl`: arbitrary bytes as stored `log_rlp`

HTTP/API input handling (security-relevant):

- [ ] `FuzzResolveAssetFile` — `cmd/apiserver`: invariant: resolved path is
      always under the asset root (path traversal cannot escape)
- [ ] `FuzzSafeFileUnderRoot` — same invariant, lower-level helper
- [ ] `FuzzNormalizeSeedSegment` + `FuzzIsHex` — `cmd/apiserver`
- [ ] `FuzzMetadataHostDispatch` — host/X-Forwarded-Host strings never panic,
      always route to exactly one handler
- [ ] `FuzzParseOptionalIntQuery` — `internal/store` (or its new home)
- [ ] `FuzzIsAddressValid` — `internal/api/common`
- [ ] `FuzzNFTAssetsPublicBase` — `internal/api/common`: normalization is
      idempotent and always yields `/images`-suffixed or empty result
- [ ] `FuzzRecoverPersonalSignSigner` — `internal/api/randomwalk`: arbitrary
      signature bytes/messages never panic; only 65-byte sigs can succeed

Domain logic (property-based invariants):

- [ ] `FuzzEloUpdate` — `internal/api/randomwalk`: ratings stay finite;
      winner's rating never decreases; loser's never increases
- [ ] `FuzzOrderByWhitelists` — every dynamic ORDER BY builder in
      `internal/store` returns a string from its whitelist for ANY input
- [ ] `FuzzShortAddress` / `FuzzShortHash` / `FuzzThousandsFormat` —
      `internal/primitives`: no panics, output shape invariants
- [ ] `FuzzDateUtils` — `internal/primitives`: round-trip/bounds invariants

Parsers and builders:

- [ ] `FuzzLogAnomalyScan` — `cmd/loganomaly`: arbitrary log lines
- [ ] `FuzzTwitterRequestBuild` — `internal/notify/tweets`: URL/params encoding
- [ ] `FuzzWhatsappPayload` — `internal/notify/wanotif`: JSON payload builder
- [ ] `FuzzTxCollectorConfig` — `cmd/opsctl`: arbitrary JSON config bytes
- [ ] `FuzzSrvmonitorConfig` — `cmd/srvmonitor`: config parser
- [ ] `FuzzConnStringEscape` — `internal/store`: `escapeConnParam` round-trip
      safety (quotes/backslashes cannot break out)

### 4.5 Benchmarks (guard the hot paths)

- [ ] `BenchmarkEventDecode` (cg-etl, most common event: BidPlaced)
- [ ] `BenchmarkReceiptsDecode` (freezer)
- [ ] `BenchmarkRateLimiter` (middleware, parallel)
- [ ] `BenchmarkStatisticsQueries` (top-3 heaviest read queries vs testdb)
- [ ] Record baselines in `docs/benchmarks.md`; re-run after each rewrite phase
      (`go test -bench=. -benchmem -count=6 | benchstat`)

### 4.6 CI additions for the safety net

- [ ] Nightly fuzz workflow (`.github/workflows/fuzz.yml`, cron): run every
      `Fuzz*` target 5–10 min each; upload crashers as artifacts; open issue on failure
- [ ] Coverage job: fail if `internal/` coverage drops below the ratchet value
      (start at current %, raise the floor after each phase — ratchet, never lower)
- [ ] Parity suite runs in the integration job (already tagged `integration`)

---

## 5. Phase 1 — Store layer made idiomatic

Goal: `internal/store` becomes a modern, context-first, error-returning,
pool-based data layer with type-safe queries. Rewrite file-by-file; each file
lands only with its §4.2 tests green.

### 5.1 Structural groundwork (do once, first)

- [ ] `Store` type owning a `*pgxpool.Pool` (replaces `SQLStorage.db *sql.DB`);
      constructor `New(ctx, Config) (*Store, error)`; `Close()`
- [ ] All methods take `ctx context.Context` as the first parameter
- [ ] Typed sentinel errors: `ErrNotFound`, `ErrConflict`; helpers wrap pgx errors
- [ ] Delete the `SQLStorageWrapper` shim; subpackages become methods on
      domain-scoped types (`store.CosmicGame()`, `store.RandomWalk()`) or
      separate injected repo structs — decide and record in §11
- [ ] Remove the `SchemaName()` string concatenation (schema is always `public`;
      set `search_path` at pool level; queries reference bare table names)
- [ ] Structured query logging via `slog` + pgx tracer (replaces `Log_msg` file logger)

### 5.2 Per-file conversion checklist

For each file: convert static SQL to sqlc (`internal/store/queries/*.sql`),
keep dynamic queries as hand-written pgx, remove every `os.Exit`, add
context, rename functions to idiomatic Go (drop `Get_`/`Insert_` snake_case —
e.g. `Get_bids_by_round_num` → `BidsByRound`, `Insert_prize_claim_event` →
`InsertPrizeClaim`), and update all callers.

`internal/store/cosmicgame/` (order: leaf files first, inserts/deletes last —
they need the §4.3 ETL fixtures in place):

- [ ] `marketing.go`, `admin_events.go`, `prize-history.go` (small warm-ups)
- [ ] `main-prize.go`, `raffle-nft.go`, `banned_bids.go`, `cosmicgame.go`
- [ ] `tokens-erc20.go`
- [ ] `bidding_analytics.go`
- [ ] `raffle-eth.go`, `nft-donations.go`
- [ ] `erc20-donations.go`
- [ ] `tokens-erc721.go`
- [ ] `contract_params.go`
- [ ] `bidding.go` (dynamic query builder → safe pgx builder w/ whitelist test)
- [ ] `eth-donations.go`
- [ ] `admin_events_resolve.go`
- [ ] `staking.go`
- [ ] `user-specific.go`
- [ ] `statistics.go` (heaviest CTEs — verify against benchmarks §4.5)
- [ ] `deletes.go` (72 funcs; consider generic `deleteByEvtlogID(table)` helper — the bodies are near-identical)
- [ ] `inserts.go` (73 funcs; last, with full ETL fixture coverage)

`internal/store/randomwalk/`:

- [ ] `ranking.go` (transactional Elo update via `pgx.Tx`)
- [ ] `randomwalk_api.go`
- [ ] `randomwalk.go`

Base:

- [ ] Migrate base files (`lookups.go`, `blockchain.go`, `blockchain_insert.go`,
      `archive.go`) from `database/sql` handles to the pgxpool-native `Store`
- [ ] Address cache in `lookups.go`: keep, but as a field on `Store` (not package
      state), with an LRU bound

### 5.3 Callers updated as each file lands

- [ ] `internal/api/*` handlers propagate store errors (500 + problem detail, no exits)
- [ ] `cmd/cg-etl` / `cmd/rw-etl` processors return errors to the loop (batch
      retry w/ backoff; crash only from `main`)
- [ ] `cmd/notibot`, `cmd/imggen-monitor`, CLIs (`cgctl`, `rwctl`, `opsctl`)
- [ ] Delete `db/{layer1,cosmicgame,randomwalk}/` raw DDL dirs once nothing
      references them (update the `opsctl archive node-fill` error message);
      goose migrations become the only schema source

---

## 6. Phase 2 — API v2 + stdlib router

Goal: a properly designed v2 API on `net/http` (Go 1.22+ pattern routing),
with v1 kept byte-identical behind the parity suite until the frontend has
migrated.

### 6.1 Design (write `docs/adr/0005-api-v2.md` first)

- [ ] Resource-oriented paths: `/api/v2/cosmicgame/rounds/{round}/bids`
      instead of `/api/cosmicgame/bid/list/by_round/:round_num/:sort/:offset/:limit`
- [ ] Pagination via query params (`?limit=&offset=` or cursor — decide in §11),
      never path segments; consistent `meta` block in list responses
- [ ] Typed top-level responses (no `{"status":1,"error":""}` envelope);
      errors as RFC 9457 `application/problem+json`
- [ ] Consistent field naming (camelCase JSON, ISO-8601 timestamps, amounts as
      decimal strings with explicit `*Wei`/`*Eth` fields)
- [ ] OpenAPI-first: author `docs/openapi-v2.yaml`, generate server stubs +
      typed models (oapi-codegen); handlers implement the generated interface
- [ ] Versioning and deprecation policy section in the ADR (v1 sunset criteria)

### 6.2 Implementation

- [ ] `internal/api/v2/` package: `Server` struct with injected `*store.Store`,
      contract-state cache, logger; no package-level state
- [ ] stdlib `http.ServeMux` with method+pattern routes; middleware as
      `func(http.Handler) http.Handler` chain (port CORS, rate limit, auth,
      metrics, recovery from gin)
- [ ] Contract-state cache: extract the ~50 globals + 3 refresh goroutines from
      [internal/api/cosmicgame/state.go](../internal/api/cosmicgame/state.go)
      into an injected `ContractState` component with lifecycle (`Run(ctx)`),
      RWMutex-guarded snapshot reads, and unit tests with a mocked eth client
- [ ] v1 compatibility layer: existing 187 routes re-registered on the new
      router calling the same service code; parity suite green
- [ ] Graceful shutdown: `http.Server.Shutdown` on SIGTERM (replaces
      `gin.Run` + `select {}`); readiness flips false during drain
- [ ] Remove gin (and the vestigial autocert manager) from go.mod once v1
      compat runs on stdlib
- [ ] Response compression + ETag/Cache-Control on hot read routes
- [ ] httptest suite for v2 (same fixtures as §4.1, new goldens)
- [ ] OpenAPI contract validation in tests (kin-openapi response validator)

### 6.3 Frontend migration

- [ ] Publish v2 spec + changelog mapping every v1 path to its v2 replacement
- [ ] Frontend switches endpoint-group by group (tracked as external checklist)
- [ ] v1 marked deprecated in spec; add `Deprecation`/`Sunset` headers
- [ ] Remove v1 layer + its goldens when traffic hits zero (final step, gated)

---

## 7. Phase 3 — ETL engine rewrite

Goal: one shared, tested indexing engine; the two binaries become thin
configuration of it. §4.3 fixtures must be green before starting.

- [ ] `internal/indexer` package: `Engine` struct (rpc client, store, registry,
      slog logger, batch policy) with `Run(ctx) error`
- [ ] `EventHandler` interface: `Topic() common.Hash; Decode(types.Log) (Event, error);
      Store(ctx, Store, Event) error` — decode separated from persistence
      (enables the decode-only fuzz/golden tests to bypass the DB)
- [ ] Port all ~80 CosmicGame handlers from `proc_*` functions to the registry;
      delete `cmd/cg-etl/events_*.go` bodies as they move
- [ ] Port the 9 RandomWalk handlers
- [ ] Replace package-level globals (eclient, ABIs, `evt_*` vars, storagew,
      Info/Error) with injected dependencies
- [ ] Batch loop: context-aware retry with exponential backoff + jitter on RPC
      and DB errors (no more crash-per-blip); circuit-break to exit only after
      N consecutive failures
- [ ] Status/progress persisted transactionally with the batch's inserts
- [ ] `log/slog` structured logging (block ranges, event counts, timings);
      keep file output via slog handler during transition
- [ ] Prometheus metrics: `rwcg_etl_last_block`, `rwcg_etl_events_total{type}`,
      `rwcg_etl_batch_duration_seconds`, `rwcg_etl_reorgs_total`
- [ ] `contract_sync.go` startup sync ported and unit-tested with mocked reads
- [ ] Fixture replay + golden DB tests pass unchanged against the new engine

---

## 8. Phase 4 — Cross-cutting Go polish

### 8.1 Naming (snake_case → idiomatic) — per package

Mechanical but wide; use gopls rename / IDE refactor per identifier, one
package per PR, parity suite green after each.

- [ ] `internal/store` + subpackages (387 funcs — mostly renamed during Phase 1; this item is the sweep for leftovers)
- [ ] `internal/api` (135 funcs; handler names like `api_cosmic_game_bid_list` → `handleBidList`)
- [ ] `cmd/cg-etl` (89 — superseded by Phase 3 port; sweep leftovers)
- [ ] `cmd/notibot` (24), `cmd/rw-etl` (11), `cmd/rwalk-alarm` (4)
- [ ] `internal/primitives` (3) + rename package to `internal/model` (or fold
      types into their owning packages — decide in §11)
- [ ] Local variables and struct fields in touched files follow along
      (err_str → errStr, bid_position → bidPosition, ...)

### 8.2 Imports and files

Eliminate all dot-imports (21 files): `cmd/cg-etl` (9), `internal/api/cosmicgame` (5),
`cmd/rw-etl` (2), `cmd/notibot` (1), `cmd/apiserver` (1), `internal/api/common` (1),
`internal/store/cosmicgame` (2 incl. test).

- [ ] Replace with named imports/aliases (`cgstore`, `rwcontracts`, ...)
- [ ] Enable `gci` + `gofumpt` in `.golangci.yml` formatters once the big moves
      are done (avoids noisy interim diffs); run `golangci-lint fmt` repo-wide

### 8.3 Configuration and logging

- [ ] `internal/config`: typed structs per service (`APIServerConfig`,
      `ETLConfig`, ...) loaded+validated at startup (fail fast, print effective
      config sans secrets); replaces ~45 scattered `os.Getenv` sites;
      `.env.example` generated or verified from the struct tags by a test
- [ ] Accept `DATABASE_URL` alongside `PGSQL_*` (12-factor friendliness)
- [ ] `log/slog` across all services (JSON in prod, text in dev); delete the
      `Info`/`Error` file-logger pairs; systemd/journald handles persistence;
      request logging middleware emits structured fields (route, status, ms, ip)

### 8.4 Lint to zero and keep it there

- [ ] Fix the remaining 180 issues (store/api/etl/CLI packages) — mostly
      falls out of Phases 1–3; sweep the rest
- [ ] Re-enable `staticcheck ST1003` (naming) once §8.1 is done; remove the
      exclusion from `.golangci.yml`
- [ ] Add `godot`, `noctx`, `contextcheck`, `paralleltest`, `testifylint` to
      the linter set once the codebase is ready for each
- [ ] `errcheck` exclusions reviewed and minimized

### 8.5 Documentation of code

- [ ] Package comment (`doc.go` where useful) for every package
- [ ] Doc comment on every exported symbol (`golangci-lint` + `revive`
      exported rule to enforce)
- [ ] Testable examples (`Example*`) for: store usage, indexer engine setup,
      middleware composition, config loading
- [ ] Regenerate abigen bindings via `go:generate` directives with a pinned
      abigen version; document in `contracts/README.md`

---

## 9. Phase 5 — Finish line

- [ ] `/version` endpoint + `--version` flags: git SHA/tag/build date via
      `-ldflags` (wire into Makefile and Dockerfile)
- [ ] Coverage ratchet reaches ≥70% on `internal/`; gate is enforced in CI
- [ ] `docs/openapi.yaml` (v1) frozen and marked deprecated; v2 spec canonical
- [ ] Delete `cmd/*/run-loop.sh` scripts after systemd cutover confirmed in prod
- [ ] Update `docs/architecture.md`, `docs/BACKEND.md`, `docs/operations.md`
      to the post-rewrite reality; refresh ADR statuses
- [ ] Final metrics dashboard update — all targets met
- [ ] Retire this document to `docs/adr/` as a historical record (or keep as
      the living maintenance checklist)

---

## 10. Test taxonomy summary (what "plenty of tests" means here)

| Kind | Where | Phase | Purpose |
|---|---|---|---|
| Unit | next to code | 0–4 | pure logic, middleware, parsers, helpers |
| Golden / characterization | `internal/api/apitest`, ETL fixtures | 0 | pin behavior across the rewrite |
| Integration (testcontainers) | store suite, migrations, triggers | 0–1 | real Postgres, every query + trigger |
| Fuzz (25+ targets) | §4.4 inventory | 0, nightly forever | decoders, parsers, security invariants |
| Property-based | Elo, whitelists, formatters | 0 | invariants under random input |
| Contract (OpenAPI) | v1 route-drift test, v2 validator | 0, 2 | spec ⇄ code can never diverge |
| HTTP handler (httptest) | v1 parity + v2 suites | 0, 2 | full request→response through real router |
| Benchmarks | §4.5 + benchstat | 0, re-run 1–3 | no performance regressions |
| E2E smoke | `opsctl smoketest` in compose | 2 | whole stack boots and answers |
| Race + shuffle | CI, all tests | always | concurrency safety |

---

## 11. Decision log (record choices here as they're made)

| # | Decision | Options | Status |
|---|---|---|---|
| D1 | Module/repo rename (`augur-explorer` → `rwcg-backend`?) | rename now / at v2 / never | open |
| D2 | v2 pagination | offset+limit / opaque cursor | open |
| D3 | Store shape | one `Store` with domain methods / per-domain repo structs | open |
| D4 | `internal/primitives` future | rename to `internal/model` / dissolve into owners | open |
| D5 | Property-testing lib | stdlib fuzz only / add `pgregory.net/rapid` | open |
| D6 | v1 sunset criteria | zero traffic for 30d / hard date | open |

---

## 12. Progress log

| Date | Commit | What landed |
|---|---|---|
| 2026-07-06 | `85941dba` | Foundation: layout, Go 1.26, CI, hardening, docs (see §3) |
| | | |
