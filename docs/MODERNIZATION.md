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

Measured 2026-07-08 (Phase 1 completion sprint). Update after each phase.

| Metric | Baseline (start of project) | Current | Target | How to measure |
|---|---|---|---|---|
| Hand-written Go LOC (`cmd/` + `internal/`) | ~60,000 (plus 124k frozen legacy) | 66,735 (legacy store layer + sqlc scaffolding deleted) | n/a (informational) | `rg --files internal cmd -g '*.go' \| tr '\n' '\0' \| xargs -0 wc -l \| tail -1` |
| snake_case functions | ~700 | **267** (api 135, cg-etl 88 `proc_*` renamed with the Phase-3 port, notibot 24, rw-etl 10, rwalk-alarm 4, misc 6; **the entire store layer is 0**) | **0** | `rg "^func (\([^)]+\) )?[A-Za-z]+_[A-Za-z0-9_]*\(" --type go -c internal cmd` |
| `os.Exit` in library code (`internal/`) | ~560 | **12 matches = 5 real calls** (3 test-harness `TestMain`s, the `cosmic_game_init` startup fatal, `primitives.Fatalf`; the other 7 are doc comments). **The whole store layer is exit-free.** | **0** (allowed only in `cmd/*/main.go` startup) | `rg -c "os\.Exit" internal` |
| Dot-import files | ~70 | **17** | **0** | `rg -l '^\s*\. "github' --type go` |
| Package-level mutable globals (api + etl) | ~120 | ~80 (state.go ~50, cg-etl ~25, rw-etl ~8) | ~0 (DI everywhere) | manual review per package |
| golangci-lint issues | 433 (first run) | **145** capped output (uncapped 492, down from 674 — deleting the legacy base/RW store layers removed their issue load) | **0** | `golangci-lint run` |
| Test files | 17 | **95** | 100+ | `rg --files -g '*_test.go' \| wc -l` |
| Fuzz targets | 0 | **28** | **25+** (see §4.4) — met; grows with Phase 3 | `rg "func Fuzz" internal cmd contracts -c` |
| Benchmarks | 0 | **4** (8 sub-benchmarks; baselines in `docs/benchmarks.md`, re-checked after Phase 1: no regression) | keep green vs baselines | `rg "func Benchmark" cmd internal -c` |
| Coverage on `internal/` (unit) | 2.4% | **8.2%** | superseded by the integration ratchet below | `go test -coverprofile=c.out ./internal/... && go tool cover -func=c.out \| tail -1` |
| Coverage on `internal/` (integration, enforced) | n/a | **67.2%** (ratchet floor raised 64% → 66% in CI) | **≥70%**, floor only moves up | `go test -race -tags=integration -coverprofile=c.out -coverpkg=./internal/... ./... && go tool cover -func=c.out \| tail -1` |
| Queries on sqlc | 0 | 0 — scaffolding retired with Phase 1 (D7 amended: hand-written pgx everywhere) | n/a | n/a |
| Routes on stdlib router | 0 | 0 (gin) | all (v1 compat + v2) | n/a |
| `context.Context` on store methods | 0% | **100% — 366 Repo methods (CosmicGame 304 + RandomWalk 62) + 22 base `Store` methods; `SQLStorage` and both wrappers are deleted** | 100% | `rg -c "func \(r \*Repo\)" internal/store/cosmicgame internal/store/randomwalk` |
| Store queries on pgx-native pool | 0 | all (one shared `pgxpool` per process; the `database/sql` view is gone) | all | manual |

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

**Status: complete as of 2026-07-07** — every section landed; the only open
item is the production-RLP replay in §4.3, which needs prod access and does
not gate Phase 1. The store rewrite (§5) is unblocked.

### 4.1 API parity suite (characterization tests)

Golden-file HTTP tests: seed a testcontainers database with a fixed fixture
set, start the real router via `httptest`, hit every route, snapshot the
JSON to `testdata/golden/`. The suite is the contract for the v1 API freeze.

**Sprint 2026-07-07: landed in full (183 goldens, 3 real bugs found — see
progress log).** Scoping decision: routes whose happy path needs live contract
reads (`bid/cst_price`, `bid/eth_price`, `bid/current_special_winners`,
`ct/statistics`, `marketing/config/current`, `time/*`, parts of
`user/balances` and the dashboard) are pinned in their deterministic
"RPC unavailable" shape against a stub Ethereum JSON-RPC server; their happy
paths become testable in Phase 2 when contract state is injectable
(`ContractState` + mocked eth client) and the goldens get regenerated then.

- [x] `internal/api/apitest/` package: fixture loader + golden-file helpers
      (`-update` flag to regenerate; byte-stable via canonical JSON with
      `json.Number`; per-case redaction for genuinely random fields; every
      case fetched twice to prove determinism). One container + one process
      -wide harness via `TestMain` (package Init can only run once until
      Phase 2 removes the package-level state).
- [x] Fixture dataset: SQL seed files under `internal/testfixtures/seed/`
      (embedded, shared with the §4.2 store suites via `testfixtures.Apply`;
      originally authored under `internal/api/apitest/testdata/seed/`)
      covering 3 complete rounds + 1 open round, 5 bidders, every prize type
      (main, raffle ETH/NFT incl. staker raffles, endurance, chrono warrior,
      lastcst), donations (ETH simple/with-info JSON, ERC20 + claim, NFT +
      claimed & unclaimed), staking (CST + RWalk stake/unstake, reward
      deposit with per-token accounting), banned bids, marketing rewards,
      admin events, charity flows, and RandomWalk mints/offers/sale/
      cancellation/withdrawals/name-changes/ranking votes with Elo state.
      Direct inserts fire the migration-00002/00003 triggers, so every
      aggregate (`cg_bidder`, `cg_winner`, `cg_glob_stats`, `cg_round_stats`,
      staking accumulators, `rw_stats`, ...) is computed by the real plpgsql
      — first real coverage of the trigger layer.
- [x] Route enumeration test: parse `docs/openapi.yaml` and assert every
      documented path is registered (and vice versa) — spec can never drift.
      Runs as a unit test (no Docker); found zero drift: all 187 operations
      match.
- [x] Golden tests for all CosmicGame GET routes (100 registered paths incl.
      legacy aliases)
- [x] Golden tests for all RandomWalk GET routes (30 registered paths)
- [x] Golden tests for `/metadata/:token_id` host dispatch (both hosts) and
      `/cg/metadata/:token_id`; `/healthz`, `/readyz`, FAQ proxy against a
      stub upstream
- [x] Mutation-route tests: `ban_bid`/`unban_bid`/`token-ranking/match`
      (auth 503/401/success matrix), `add_game` signature verification paths
      (challenge nonce, EIP-191 recovery, duplicate-pair 409, nonce replay,
      chain-id rejection); all mutations restore fixture state so test order
      and `-shuffle` cannot affect goldens
- [x] Error-shape tests: invalid params on representative routes pin the
      `{"status":0,"error":...}` envelope (including the HTTP-200-on-invalid
      -address legacy quirk)
- [x] Wire into `make test-integration` and CI (integration build tag; suite
      auto-skips when Docker is unavailable)

### 4.2 Store integration suite

One test file per store file; every public query function called against the
seeded database at least once, asserting golden rows. This is what lets the
sqlc/pgx rewrite (§5) proceed file-by-file with confidence.

**Sprint 2026-07-07: read suite landed in full (196 goldens, 3 production
bugs found — see progress log).** Harnesses: `main_integration_test.go` in
each store package boots `testdb.Start` + the shared embedded seed dataset
(`internal/testfixtures`, extracted from apitest so both suites pin the same
data); every golden is fetched twice to prove determinism
(`testutil.GoldenJSON`). Error paths of the legacy `os.Exit` methods stay
untestable by design until Phase 1 converts them to returned errors — the
goldens are the safety net for exactly that conversion.

CosmicGame (`internal/store/cosmicgame/`):

- [x] `inserts.go` (73 funcs — covered via §4.3 ETL fixture replay: every
      dispatched event type inserts through its store function against the
      real triggers; goldens pin the rows *(2026-07-07)*)
- [x] `deletes.go` (72 funcs — covered via §4.3 replay idempotence: every CG
      fixture re-processes, exercising Delete_* then Insert_* per event type;
      trigger delete paths additionally covered by the reorg rollback golden
      *(2026-07-07; also fixed two wrong table names in deletes.go)*)
- [x] `statistics.go` (20 funcs — all 6 ROI-leaderboard sort branches pinned)
- [x] `user-specific.go` (20 funcs)
- [x] `staking.go` (18 funcs — open + closed stake actions, per-deposit and
      per-token reward views)
- [x] `admin_events_resolve.go` (14 — fixture events resolved end-to-end plus
      synthetic events driving every SQL-backed resolver branch: real
      activation-span/ETH-price hits and the documented fallbacks)
- [x] `eth-donations.go` (13)
- [x] `bidding.go` (13 — v1/v2 bid shapes, pagination, sort directions,
      not-found soft paths)
- [x] `contract_params.go` (12 — reads incl. the error path on a bad table;
      the ETL-facing Sync* write surface is exercised by `contract_sync.go`
      startup sync and moves with Phase 3)
- [x] `tokens-erc721.go` (11)
- [x] `erc20-donations.go` (9)
- [x] `bidding_analytics.go` (8 — incl. zero-filled bucket semantics)
- [x] `raffle-eth.go` (8)
- [x] `nft-donations.go` (8 — claimed and unclaimed donation states)
- [x] `tokens-erc20.go` (5)
- [x] `cosmicgame.go` (4 — incl. processing-status lazy-insert round trip)
- [x] `raffle-nft.go` (3), `main-prize.go` (3), `banned_bids.go` (3 —
      insert/delete round trip restores fixture state)
- [x] `prize-history.go` (2), `marketing.go` (2), `admin_events.go` (2)

RandomWalk (`internal/store/randomwalk/`):

- [x] `randomwalk.go` (29 funcs — ETL-facing inserts covered by the §4.3
      rw-etl fixtures; the notification/top-stats read surface (notibot's
      `Get_all_events_for_notification*`, `Get_messaging_status` watermark
      round trip, `Get_last_mint_timestamp`, rw_uranks top-makers with a
      suite-local extension seed) pinned by the store suite *(2026-07-07;
      found the unseeded rw_messaging_status watermark bug — migration
      `00008`)*)
- [x] `randomwalk_api.go` (24 funcs — all `Get_active_offers` order branches,
      floor price, full token history; *fixed `Check_rwalk_token_exists`,
      which was broken on every call*)
- [x] `ranking.go` (9 funcs — Elo transaction semantics: rollback leaves
      count/ratings/votes untouched, commit applies match + both rating
      upserts; nonce lifecycle incl. replay rejection and expiry purge)

Base (`internal/store/`):

- [x] `lookups.go`, `blockchain.go`, `blockchain_insert.go`, `archive.go`
      (round-trip block/tx/evt_log insertion incl. reorg path) — covered by
      `internal/etl/blockops_integration_test.go` (§4.3 sprint): block insert
      + hash verification + watermark, chain-split cascade, tx three-level
      fallback (RPC / archive / minimal), evt_log dedup-replace, FilterLogs
      range+address filtering *(2026-07-07)*
- [x] Trigger behavior tests: `cg_bidder`/`cg_glob_stats`/`cg_winner`
      aggregates update correctly on insert/delete (the plpgsql functions in
      migration 00002 are load-bearing). *Insert paths pinned end-to-end by
      the §4.1 fixtures; found+fixed the token-name trigger collision
      (migration 00004). Delete/reorg paths now pinned by the §4.3 replay
      idempotence checks and reorg rollback goldens; found+fixed the
      unstake-delete restore gap (00006) and the item-bought delete NEW/OLD
      + profit-reversal bugs (00007). (2026-07-07)*

### 4.3 ETL decode fixtures (golden events)

**Sprint 2026-07-07: landed in full — 97 golden files, 7 production bugs
found and fixed (see progress log).** Infrastructure: `internal/testchain`
(deterministic in-memory Ethereum JSON-RPC node: real header hashes, signed
txs, receipts, `eth_getLogs`, registrable `eth_call` handlers, `Reorg()`)
and `internal/testutil` (golden compare + canonical DB snapshot/diff with
FK resolution to natural keys — addresses, tx hashes, `evt:block/logindex`).

- [x] `cmd/cg-etl/fixtures_test.go`: one fixture per dispatched event type
      (all `select_event_and_process` branches — 74 at the time, 75 since the
      write-layer sprint added the missing ERC20TransferFailed dispatch —
      incl. both address-guarded
      handlers of the two duplicate-topic events), built by ABI-packing known
      values and pushed through the real pipeline (`EnsureBlockExists` →
      `EnsureTransactionExists` → `InsertEventLog` → `process_single_event`);
      the golden pins the full DB diff including every plpgsql trigger side
      effect. Each fixture is re-processed afterwards and must be
      state-neutral (pins the delete-then-insert reorg recovery of every CG
      handler). Negative cases: wrong-address skip, unknown topic0, zero
      topics. `topics_test.go` additionally pins every hand-maintained
      registry constant against the ABI-derived event ID as a unit test (no
      Docker), so ABI regeneration can never silently retire a dispatch.
- [x] Golden DB-state test: `TestScriptedRoundGolden` processes a complete
      round (admin bootstrap, ETH/RandomWalk/CST bids, donations of every
      kind, the 13-log claim transaction, withdrawal, donated-prize claims,
      transfers, staking with reward deposit) — one cumulative golden
- [x] Same for `cmd/rw-etl` (all 7 dispatched event types + skip paths for
      unknown token/offer, marketplace story golden). RW handlers insert
      without delete-first, so single-event replay is impossible by design
      (UNIQUE(evtlog_id) aborts); re-processing is pinned via the reorg test.
- [~] RLP replay: the synthetic fixtures round-trip real RLP through
      `evt_log.log_rlp` (encode on insert, decode in `process_single_event`),
      which pins the storage format end to end. Replaying production-exported
      samples (via `opsctl archive export`) remains open — needs prod access.
- [x] Reorg simulation: `TestReorgRollbackAndReplay` (both ETLs) reorgs the
      fake chain mid-story, drives `EnsureBlockExists` → `HandleChainSplit`,
      pins the rolled-back state as a golden (trigger delete paths reverse
      the aggregates) and asserts that re-processing the replacement fork
      reproduces the pre-reorg state exactly. Found the unstake/item-bought
      trigger reversal bugs (migrations `00006`/`00007`).

### 4.4 Fuzz test inventory (the panic-hunting fleet)

Native Go fuzzing (`go test -fuzz`). Every target carries its seed corpus
inline via `f.Add(...)` (reviewed in source; crashers found later get
committed under `testdata/fuzz/` as regression inputs) and runs in the
nightly CI fuzz job (§4.6). Run locally: `make fuzz-smoke` (10s/target) or
`scripts/fuzz-all.sh 5m`.
Invariant unless stated otherwise: *never panics, never hangs*.

**Sprint 2026-07-06 (`a7971755`):** 28 fuzz targets landed (all checked items below).
One real bug found and fixed: a corrupt freezer index could make
`ReadItem` allocate up to 2^48 bytes and OOM-kill the process
(`internal/freezer`, both reader paths; regression test
`TestReadItemCorruptIndexHugeOffset`).

Decoders (highest value — they consume chain data):

- [x] `FuzzReceiptsDecode` — `internal/freezer/decode`: arbitrary bytes → RLP receipt decode
- [x] `FuzzArbitrumLegacyDecode` — `internal/freezer/decode`: the Arbitrum-specific format
      (+ `FuzzArbitrumLogDecode` for the single-log decoder)
- [x] `FuzzFreezerIndexRead` — `internal/freezer`: corrupt index/data-file bytes
      (+ `FuzzUint48RoundTrip`; found the OOM bug above)
- [~] `FuzzEventDecodeCG` — decode-only half landed as `FuzzABIEventUnpack` in
      `contracts/cosmicgame` (fuzzes every event decoder in all 10 CosmicGame
      ABIs). The full registry-driven fuzz over the `cmd/cg-etl` `proc_*`
      handlers is blocked until Phase 3 separates decode from persistence.
- [~] `FuzzEventDecodeRW` — same split: `FuzzABIEventUnpack` in
      `contracts/randomwalk` landed; handler-level fuzz waits for Phase 3
- [x] `FuzzEvtlogRLP` — lives in `cmd/cg-etl` (the real decode site of stored
      `log_rlp`, not `internal/etl` which only fetches); includes
      decode→encode→decode round-trip property

HTTP/API input handling (security-relevant):

- [x] `FuzzResolveAssetFile` — `cmd/apiserver`: invariant: resolved path is
      always under the asset root (path traversal cannot escape); also pins
      that the package-layout fallback stays scoped to `new/cosmicsignature/`
- [x] `FuzzSafeFileUnderRoot` — same invariant, lower-level helper
- [x] `FuzzNormalizeSeedSegment` + `FuzzIsHex` — `cmd/apiserver`
- [x] `FuzzMetadataHostDispatch` — host/X-Forwarded-Host strings never panic,
      always route to exactly one handler (logic extracted from the `main()`
      closure; now `common.MetadataHostServesCosmicSignature` in
      `internal/api/common` since the parity harness registers the same
      dispatch, fuzz + unit tests moved along)
- [x] `FuzzParseOptionalIntQuery` — `internal/store/cosmicgame`
- [x] `FuzzIsAddressValid` — `internal/api/common` (accepted ⇒ EIP-55
      checksummed; rejected ⇒ JSON error envelope written)
- [x] `FuzzNFTAssetsPublicBase` — `internal/api/common`: normalization is
      idempotent and always yields `/images`-suffixed or empty result
- [x] `FuzzRecoverPersonalSignSigner` — `internal/api/randomwalk`: arbitrary
      signature bytes/messages never panic; only 65-byte sigs can succeed
      (+ sign→recover round-trip unit test with a generated key)

Domain logic (property-based invariants):

- [x] `FuzzEloUpdate` — `internal/api/randomwalk`: no NaN; winner's rating
      never decreases; loser's never increases; pair total conserved
- [x] `FuzzOrderByWhitelists` — landed as `FuzzRoiLeaderboardOrderClause`
      (`internal/store/cosmicgame`) and `FuzzActiveOffersOrderClause`
      (`internal/store/randomwalk`) after extracting the sort switches into
      pure whitelist functions; these are the only two sites where request
      input selects an ORDER BY (the bid/NFT query builders take literals)
- [x] `FuzzShortAddress` / `FuzzShortHash` / `FuzzThousandsFormat` —
      `internal/primitives`: no panics, output shape invariants,
      strip-commas-and-parse round trip
- [x] `FuzzDateUtils` — `internal/primitives`: component bounds, symmetry,
      zero-diff identity

Parsers and builders:

- [x] `FuzzLogAnomalyScan` — `cmd/loganomaly`: arbitrary log lines
- [x] `FuzzTwitterRequestBuild` — `internal/notify/tweets`: percent-encoding
      round-trips; OAuth base string keeps exactly 2 top-level `&` separators
- [x] `FuzzWhatsappPayload` — `internal/notify/wanotif`: error-body parser
      total; template payload marshals to valid round-tripping JSON
- [x] ~~`FuzzTxCollectorConfig`~~ — dropped: `opsctl txcollector` is
      flag-driven; no JSON config parser exists to fuzz
- [x] ~~`FuzzSrvmonitorConfig`~~ — deferred to §8.3: config is `LoadFromEnv`
      (env-var driven); fuzz the typed config loader once it exists
- [x] `FuzzConnStringEscape` — `internal/store`: `escapeConnParam` cannot
      break out of quotes (scanner proof) and round-trips through
      `pgx.ParseConfig` without parameter injection

### 4.5 Benchmarks (guard the hot paths)

**Sprint 2026-07-07: landed in full; baselines recorded.**

- [x] `BenchmarkEventDecode` (cg-etl, most common event: BidPlaced — ABI
      unpack + topic extraction, persistence excluded)
- [x] `BenchmarkReceiptsDecode` (freezer; raw-RLP and snappy sub-benchmarks —
      writing it found the snappy misdetection bug, see progress log)
- [x] `BenchmarkRateLimiter` (middleware, parallel; distinct-IP map path and
      shared-IP bucket contention)
- [x] `BenchmarkStatisticsQueries` (top-3 heaviest read queries vs the seeded
      testdb: dashboard aggregate, claims-by-round CTE, ROI leaderboard;
      integration tag)
- [x] Record baselines in `docs/benchmarks.md`; re-run after each rewrite phase
      (`go test -bench=. -benchmem -count=6 | benchstat`)

### 4.6 CI additions for the safety net

- [x] Nightly fuzz workflow (`.github/workflows/fuzz.yml`, cron 03:17 UTC +
      `workflow_dispatch`): runs every `Fuzz*` target 5 min each via
      `scripts/fuzz-all.sh`; uploads crashers as artifacts; opens/updates a
      `fuzz-failure` issue on failure. Local equivalent: `make fuzz-smoke`.
      *(2026-07-06)*
- [x] Coverage job: fail if `internal/` coverage drops below the ratchet value
      (start at current %, raise the floor after each phase — ratchet, never
      lower). *(2026-07-07: integration job computes `-coverpkg=./internal/...`
      coverage and fails below 50%; measured 53.0%. ETL fixture sprint raised
      the floor to 60%; measured 62.7%. Store read suite raised the floor to
      64%; measured 66.7%.)*
- [x] Parity suite runs in the integration job (already tagged `integration`)
      *(2026-07-07)*

---

## 5. Phase 1 — Store layer made idiomatic

Goal: `internal/store` becomes a modern, context-first, error-returning,
pool-based data layer with type-safe queries. Rewrite file-by-file; each file
lands only with its §4.2 tests green.

**Status: complete as of 2026-07-08.** The store layer (base + CosmicGame +
RandomWalk) is fully pgx-native and context-first; `SQLStorage`, both
`SQLStorageWrapper`s, the `database/sql` pool view and the package-level
address cache are deleted. The only remaining §5.3 caller work (renaming the
`proc_*` ETL handlers) is the Phase 3 port.

### 5.1 Structural groundwork (do once, first)

**Sprint 2026-07-07: landed in full (see progress log).** Package renamed
`dbs` → `store` on the way (all ~23 importers updated).

- [x] `Store` type owning a `*pgxpool.Pool` (`internal/store/store.go`):
      `New(ctx, Config) (*Store, error)` with keepalive dialer, bounded
      startup ping-retry (replaces `retryConnector`), pool-wide `timezone=UTC`
      (fixes the legacy one-connection-only `SET timezone TO 0`), explicit
      `DefaultMaxConns`; `NewFromPool` for tests; `Close()`. Transitional
      `Store.DB()` exposes a `database/sql` view of the same pool so every
      binary runs ONE pool while unconverted `SQLStorage` methods remain;
      `Connect_to_storage`/`openDB` deleted, all 9 call sites converted
      (apiserver, cg-etl, rw-etl, notibot, imggen-monitor, cgctl ×2, rwctl,
      opsctl). `ConnectHint` preserves the operator diagnostics of the old
      `show_connect_error`. *2026-07-08: the transitional `Store.DB()` view
      and the `SQLStorage` type are deleted — the pool is pgx-only.*
- [x] All methods take `ctx context.Context` as the first parameter
      (established by the Repo pattern; applies per file as each converts)
- [x] Typed sentinel errors: `ErrNotFound`, `ErrConflict`; `WrapError` maps
      pgx/sql no-rows and unique-violations, preserving chains (unit-tested)
- [x] `SQLStorageWrapper` shrinks per converted file (deleted at the end of
      Phase 1); **D3 decided: separate injected repo structs** — the
      CosmicGame wrapper deleted 2026-07-07 (write-layer sprint), the
      RandomWalk wrapper and the `SQLStorage` type itself deleted 2026-07-08
      (Phase 1 completion sprint; `store.QueryList` extracted so both repos
      share the scan loop)
- [x] `SchemaName()` concatenation removed from converted queries (bare table
      names; pool pins `search_path=public`); unconverted files keep it until
      their conversion
- [x] Structured query logging via `slog` + pgx `QueryTracer` (failed + slow
      queries, SQL truncated, cancellations skipped; unit-tested). Converted
      code no longer uses `Log_msg`.
- [x] `internal/testdb` exposes the container's `pgxpool.Pool`; store suite,
      apitest and cg-etl harnesses build the same one-pool wiring as
      production (`store.NewFromPool` → Repo + legacy view)
- [x] `store.TimeText` scan adapter: timestamptz → RFC3339Nano string,
      byte-identical to the `database/sql` conversion the goldens pin
      (+ `store.NullTimeText` for nullable timestamp columns: NULL leaves
      the destination unchanged, matching the legacy NullString pattern)

### 5.2 Per-file conversion checklist

For each file: move the queries onto `Repo` (pgx-native, hand-written SQL —
see D7 for the narrowed sqlc scope), remove every `os.Exit`, add context,
rename functions to idiomatic Go (drop `Get_`/`Insert_` snake_case — e.g.
`Get_bids_by_round_num` → `BidsByRound`, `Insert_prize_claim_event` →
`InsertPrizeClaim`), and update all callers. Byte-identical goldens are the
gate: pin timestamptz-into-string columns with `store.TimeText`, keep
`make([]T, 0, n)` list semantics (empty JSON `[]`, never `null`), and swap
text literals feeding bool fields (`'T'`) for real booleans — native pgx does
not repeat `database/sql`'s permissive string conversions.

`internal/store/cosmicgame/` (order: leaf files first, inserts/deletes last —
they need the §4.3 ETL fixtures in place):

- [x] `marketing.go`, `admin_events.go`, `prize-history.go` (small warm-ups)
      *(2026-07-07 — `MarketingRewardHistoryGlobal`, `MarketingRewardsByUser`,
      `SystemModeChanges`, `AdminEventsInRange`; the 39-branch admin UNION is
      now generated from a registry table with a completeness unit test and a
      valid-SQL integration test covering fixture-empty tables)*
- [x] `main-prize.go`, `raffle-nft.go`, `banned_bids.go`, `cosmicgame.go`
      *(2026-07-07 — `PrizeClaims`, `PrizeInfo` (`(bool, rec)` →
      `(rec, error)` with `ErrNotFound`), `AllPrizesForRound`,
      `RaffleNFTWinnersByRound` (the `'T'/'F'` string-built staker flag is a
      bound parameter now), `RaffleNFTWinners`, `BannedBids`/`InsertBannedBid`/
      `DeleteBannedBid`, `ContractAddrs`, `ProcessingStatus`,
      `UpdateProcessingStatus`)*
- [x] `tokens-erc20.go` *(2026-07-07 — `CosmicTokenHolders`,
      `CosmicTokenStatistics`, `UserCosmicTokenSummary`,
      `CosmicTokenSupplyHistoryByBid`, `CosmicTokenSupplyHistoryByDate`)*
- [x] `bidding_analytics.go` *(2026-07-07 — `BidFrequencyByPeriod`,
      `BidTypeRatioByPeriod`, `TopBidders`, `TopBidderActivePeriods`,
      `BidTimeBounds`; the epoch-aligned vs anchored bucket SQL extracted to
      `bidFrequencySQL` with unit tests pinning branch selection and that
      only the integer interval is ever interpolated)*
- [x] `raffle-eth.go` *(2026-07-07 — `UnclaimedPrizeEthDeposits`,
      `PrizeEthDeposits`, `RaffleEthDeposits`, `ChronoWarriorEthDeposits`,
      `EthDepositsByUser`, `RaffleEthDepositsByUser`,
      `ChronoWarriorEthDepositsByUser` + the earlier `PrizeDepositsByRound`;
      nullable claim timestamp/date via `store.NullTimeText`)*
- [x] `nft-donations.go` *(2026-07-07 — `NFTDonations`, `NFTDonationInfo`
      (`(bool, rec)` → `(rec, error)`), `DonatedNFTClaims`,
      `NFTDonationsByRound`, `NFTDonationsByToken`,
      `UnclaimedDonatedNFTsByRound`, `DonatedTokenDistribution`,
      `NFTDonationsByUser`)*
- [x] `erc20-donations.go` *(2026-07-07 — `ERC20DonationsByRoundDetailed`/
      `All`/`Summarized`, `ERC20Donations`, `ERC20DonationInfo`,
      `ERC20DonationsByUser`, `ERC20DonationClaims`/`ByUser`/`ByRound`)*
- [x] `tokens-erc721.go` *(2026-07-07 — `CosmicSignatureTokens`,
      `CosmicSignatureTokenInfo`, `TokenNameHistory`,
      `TokenOwnershipTransfers`, `CosmicSignatureTokenDistribution`,
      `SearchTokensByName` (first store-suite coverage + golden),
      `NamedTokens`, `CosmicSignatureTokenCount`, `CosmicSignatureTokenSeed`;
      legacy `buildNFTSelectQuery`/`scanNFTRecord` twins stay for
      `user-specific.go` and die with its conversion)*
- [x] `contract_params.go` *(2026-07-07 — `GlobStatsCstRewardForBidding`,
      `LatestDecimalParam`, `InsertAdminCorrectionDecimal`,
      `InsertAdminCorrectionERC20Reward`, all with a lowercase-identifier
      guard on the interpolated table/column names; the check-then-correct
      sync policy (`SyncAdmin*IfMismatch`, `SyncCstRewardIfMismatch`) moved
      to `cmd/cg-etl/contract_sync.go` as `paramSyncer` — storage keeps the
      primitives, the ETL owns the policy, and address resolution stays lazy
      so a clean sync run leaves the address table untouched)*
- [x] `bidding.go` *(2026-07-07 — `Bids`, `BidInfo`, `BidsByRound`,
      `BidsWithMessageByRound`, `BidIDByEvtlog`, `BidRowIDByEvtlogID`,
      `EvtlogIDByRoundAndBidPosition`, `BidCountForRound`,
      `LastCstBidEvtlogForBidder`, `RoundStartTimestamp`,
      `RandomWalkTokensUsedInBids`; the §5.2 builder item landed as
      `bidSelectQuery` over WHERE/ORDER BY/paging whitelists —
      `TestBidSelectQueryWhitelists` exercises every combination and the
      rejection path, so request input can never reach ORDER BY)*
- [x] `eth-donations.go` *(2026-07-07 — `CharityDonations`(+`FromCosmicGame`,
      `Voluntary`), `CharityWalletWithdrawals`, `SimpleEthDonations`(+`ByRound`),
      `EthDonationsWithInfo`(+`ByRound`), `EthDonationWithInfoRecord`
      (zero-record-on-miss → `ErrNotFound`; handler keeps the legacy
      zero-record response), `EthDonationsByUser`/`ByRound`/`EthDonations`
      (the UNION builder shared), `DonationReceivedEvtIDByTx`)*
- [x] `admin_events_resolve.go` *(2026-07-07 — `ResolveAdminEventValues(ctx,
      events) error`; the three SQL lookups are ctx-aware Repo helpers whose
      interpolated table/column names pass the `checkAdminIdent` guard; real
      DB errors now propagate while no-rows keeps the documented ""/fallback
      renderings, so the goldens hold; pure format helpers unchanged)*
- [x] `staking.go` *(2026-07-07 — 16 methods incl.
      `StakeActionCstInfo`/`StakeActionRwalkInfo` (`(bool, rec)` →
      `(rec, error)` + `ErrNotFound`), `StakingRewardsToBeClaimed`/
      `Collected`, `StakedTokensCst/RwalkGlobal`, `GlobalStakingRewards`,
      `GlobalStakingCst/RwalkHistory` (+ Repo-local `lastBlockTimestamp`),
      staking mints, `StakingCstUserDepositRewards`/`UserTokenRewards`/
      `UserTokenRewardDetails`; the stake-action queries are pure functions
      (`stakeActionQueryCST/RWalk`) pinned by a no-Docker unit test.
      **Bug found & fixed:** `Get_staking_cst_mints_global` hardcoded
      `IsRWalk=true` on rows its own WHERE clause filters to
      `is_rwalk=FALSE` (copy-paste from the RWalk variant), so
      `staking/cst/mints/global` mislabeled every CST-staker mint; store +
      parity goldens updated for the corrected flag)*
- [x] `user-specific.go` *(2026-07-07 — 20 methods: `UserInfo` (`(bool,rec)`
      → `(rec, error)` + `ErrNotFound`; the rwalk-staking sub-query keeps its
      no-rows-keeps-zeros semantics), `PrizeClaimsByUser`, `BidsByUser` (on
      the whitelisted `bidList` builder), `UnclaimedDonatedNFTsByUser`,
      `RaffleNFTWinningsByUser`, `PrizeDepositsChronoWarrior/RaffleEthByUser`,
      `DonatedNFTClaimsByUser`, `CosmicSignatureTokensByUser` (on
      `nftListSelectSQL`), transfers, `MarketingRewardHistoryByUser` (shares
      the marketing.go column list), staked tokens/actions/mints by user,
      `UserNotifRedBoxRewards` (preserves the legacy early return that leaves
      `DonatedERC20Tokens` nil for non-CST-stakers),
      `ERC20DonatedPrizesByWinner`; the legacy `buildBidSelectQuery`/
      `scanBidRecord` and `buildNFTSelectQuery`/`scanNFTRecord` twins are
      deleted)*
- [x] `statistics.go` *(2026-07-07 — 16 methods; `CosmicGameStatistics`
      composes `StakeStatisticsCst/Rwalk` and the already-converted
      `DonatedTokenDistribution` (legacy private copy deleted);
      `CosmicGameRoundStatistics` + `activationTimeFromEvents` keep the
      open-round fallback shape; `ClaimsByRound`/`ClaimDetailByRound` with
      ctx-aware unclaimed-item scanners; `RoiLeaderboard` keeps the fuzzed
      ORDER BY whitelist; dead `Get_num_prize_claims` (no production caller)
      deleted; §4.5 benchmarks re-run on the Repo — see `docs/benchmarks.md`,
      the transitional stdlib-over-pool B/op inflation is gone)*
- [x] `deletes.go` *(2026-07-07 — one unexported `deleteByEvtlogID(ctx, table, id)`
      helper + 72 named methods paired with their inserts (`DeletePrizeClaim`,
      `DeleteBid`, ...); a reflection sweep (`TestDeleteMethodsValidSQL`)
      executes every `Delete*(ctx, int64) error` method against the real
      schema so a typo'd table name — the bug class fixed twice in §4.3 —
      fails loudly)*
- [x] `inserts.go` *(2026-07-07 — 73 context-first methods incl.
      `insertAdminValue` for the ~35 single-value admin tables and address
      FKs resolved through the new `Store.LookupOrCreateAddress` (bounded
      per-Store LRU replaces the package-level cache for converted callers);
      `InsertBid` keeps the bid-position/CST-reward pre-queries but **real DB
      errors now propagate instead of silently defaulting bid_position to 1**
      (the legacy layer mislabeled later bids of a round on any DB failure);
      `InsertDonationJSON` keeps the FK-cascade replay semantics. All 97 ETL
      fixture goldens byte-identical)*

`internal/store/randomwalk/`:

- [x] `ranking.go` (transactional Elo update via `pgx.Tx`) *(2026-07-08 —
      9 methods incl. `ApplyRankingMatch(ctx, tx, ...)` and
      `ConsumeRankingVoteNonce(ctx, tx, ...)` on `pgx.Tx`;
      `handlers_ranking.go` runs both vote transactions on `Pool().Begin`;
      nil-slice semantics of the ID-list queries preserved (handlers marshal
      them directly))*
- [x] `randomwalk_api.go` *(2026-07-08 — 23 methods: `ActiveOffers` (aids are
      bound parameters now, ORDER BY stays the fuzzed whitelist),
      `MintedTokensByPeriod`/`Sequentially`, `TradingHistory`,
      `RandomWalkStats`, `MarketStats`, `TokenFullHistory`,
      `MarketTradingVolumeByPeriod`, `TokenNameChanges`, `TokensByUser`,
      `FloorPrice` (empty order book keeps the explicit noOffers flag for
      notibot; the driver's ErrNoRows no longer leaks), `TradingHistoryByUser`,
      `UserInfo`/`TokenInfo` (`(rec, error)` + `ErrNotFound`; handlers render
      the byte-identical legacy not-found strings), `TokenMinted`,
      `Top5TradedTokens`, `MintIntervals`, `WithdrawalChart`, `SaleHistory`,
      `FloorPriceByPeriod`, `MintedTokensCSV`, `MintReport`)*
- [x] `randomwalk.go` *(2026-07-08 — 30 methods: processing/messaging status
      round trips (lazy singleton insert like the CG pattern),
      `ContractAddrs` + new `RawContractAddrs` (replaces rw-etl's raw
      `QueryRow` in main), all 7 ETL inserts on `Store.LookupOrCreateAddress`
      (`must_lookup_or_create_address` and its `os.Exit` die),
      `OfferExists`/`TokenExists` (**a DB failure during the existence check
      used to silently skip the event — a data-loss bug; real errors now
      abort the batch**), the rw_uranks rank upserts, notification reads,
      `ServerTimestamp`, `LastMintTimestamp`, `TokenTransfersByTxHash`)*

Base:

- [x] Migrate base files (`lookups.go`, `blockchain.go`, `blockchain_insert.go`,
      `archive.go`) from `database/sql` handles to the pgxpool-native `Store`
      *(2026-07-08 — 17 ctx-first `Store` methods: `EventLog`,
      `EventsBySigAndTx`, `EventLogRLPsBefore`, `BlockHash`, `LastBlockNum`,
      `SetLastBlockNum`, `DeleteBlock`, `EvtLogExists`,
      `CountEvtLogsForContract`, `TransactionIDByHash`, `InsertBlock`,
      `InsertMinimalTransaction`, `InsertTransaction`, `InsertEventLog`,
      `ArchivedTransactionByHash`, `InsertTransactionFromArchive`,
      `AddressByID`; the 6 dead functions (`Get_evtlogs_by_signature*`,
      `Get_last_evtlog_id`, `Get_last/first_block_timestamp`,
      `Get_archived_event_logs`) deleted; `lookups.go` with the process-wide
      `amap` cache deleted — every caller is on the per-Store LRU.
      `internal/etl` (`ETLContext.Store *store.Store`) and all its callers
      take a context; `FetchEvents`/`GetCurrentBlockNumber` too.)*
- [x] Address cache: field on `Store` with an LRU bound *(2026-07-07 —
      `internal/store/address.go`: `LookupOrCreateAddress`/`LookupAddressID`
      on the pgx pool with a bounded per-Store LRU (`DefaultAddressCacheSize`
      64k, race-safe, unit + integration tested; concurrent-create races
      resolve via the unique constraint + re-read). 2026-07-08: the
      package-level cache in `lookups.go` is deleted with the base-file
      migration; `AddressByID` added for the reverse lookups.)*

### 5.3 Callers updated as each file lands

- [x] `internal/api/*` handlers propagate store errors — every CosmicGame
      read handler now passes `c.Request.Context()` and answers failures with
      `respondStoreError` (logs + HTTP 500 in the legacy `{"status":0,...}`
      envelope via `common.RespondInternalErrorJSON`; these paths previously
      killed the process, so no parity golden constrains them; `ErrNotFound`
      keeps the exact legacy not-found responses — incl. the nine
      `Get_user_info` gate sites and both stake-action-info routes). The
      `state.go` background refreshers (`do_reload_database_variables` incl.
      the statistics aggregate) call converted methods with
      `context.Background()` and keep the previous value on failure
      (ContractState extraction is Phase 2). *2026-07-08: the last legacy
      surface is gone — the ~39 `Nonfatal_lookup_address_id` sites (CG API)
      and the RandomWalk handlers run on `Store.LookupAddressID`/`AddressByID`
      with the request context; the RandomWalk package got its own
      `respondStoreError` and its ~49 handler sites pass
      `c.Request.Context()`; the three charity routes that called `os.Exit`
      from inside a request handler on a failed contract-address lookup
      answer HTTP 500 instead (a cancelled request could previously kill the
      server); `/readyz` pings `Store.Pool()`.*
- [x] `cmd/cg-etl` fully error-propagating *(2026-07-07 — every `proc_*`
      handler takes `ctx` and returns errors (decode failures included);
      `select_event_and_process` became a dispatch table that checks every
      handler; `process_single_event`'s RLP-decode `panic` is a returned
      error; the loop leaves failed batches unacknowledged and crashes only
      from `main`. Graceful shutdown runs the in-flight batch's DB work on
      `context.WithoutCancel`, so SIGTERM mid-batch still gets the promised
      "finish batch, write status, exit 0". Full batch retry w/ backoff is
      Phase 3.)*
- [x] `cmd/notibot`, `cmd/imggen-monitor`, CLIs (`cgctl`, `rwctl`, `opsctl`)
      — all construct the shared `Store` (one pool per process) and handle
      connect errors; per-query conversion followed their files *(2026-07-08 —
      notibot polls through the Repo with ctx (DB failures keep the legacy
      crash-and-restart semantics under systemd, resuming from the persisted
      watermark); `rwctl` commands run on `(Repo, Store)` from
      `connectRWStorage` (top-rated rank writes now check errors); `cmd/rw-etl`
      mirrors cg-etl: dispatch-table handlers return errors, RLP-decode panic
      is a returned error, shutdown finishes the in-flight batch on
      `context.WithoutCancel`; `opsctl archive node-fill` resolves addresses
      through a pgx `Store` (its tool-local archive statements keep their own
      DSN handle); the `Init_log`/`Log_msg` file loggers are replaced by the
      pgx slog tracer writing to the same db.log files)*
- [x] Delete `db/{layer1,cosmicgame,randomwalk}/` raw DDL dirs once nothing
      references them (update the `opsctl archive node-fill` error message);
      goose migrations become the only schema source *(2026-07-08 — dirs
      deleted; node-fill help/error text points at `db/migrations`, which
      already carries the archive tables. The unused sqlc scaffolding
      (`sqlc.yaml`, `internal/store/queries/`, `internal/store/sqlcgen/`,
      `make generate`) is retired with it — D7 amended.)*

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
| D3 | Store shape | one `Store` with domain methods / per-domain repo structs | **decided 2026-07-07: per-domain repo structs** — `cosmicgame.Repo` wraps the shared `*store.Store`; `randomwalk.Repo` follows when its files convert. Keeps domain queries in their domain packages and the base package free of game knowledge. |
| D4 | `internal/primitives` future | rename to `internal/model` / dissolve into owners | open |
| D5 | Property-testing lib | stdlib fuzz only / add `pgregory.net/rapid` | **decided 2026-07-06: stdlib-only** — the §4.4 fleet needed no extra dependency; revisit only if a future property needs structured generators |
| D6 | v1 sunset criteria | zero traffic for 30d / hard date | open |
| D7 | sqlc scope (amends the §5.2 blanket "convert static SQL to sqlc") | all static queries / simple lookups only / none | **decided 2026-07-07: hand-written pgx for the read layer** — the store's heavy COALESCE/CASE/outer-join UNIONs defeat sqlc's nullability inference and would force pointer-mapped row types that fight the pinned JSON shapes. **Amended 2026-07-08: sqlc retired entirely** — the base-file conversion superseded its 8 layer1 queries with hand-written `Store` methods and the never-imported scaffolding was deleted. |

---

## 12. Progress log

| Date | Commit | What landed |
|---|---|---|
| 2026-07-06 | `85941dba` | Foundation: layout, Go 1.26, CI, hardening, docs (see §3) |
| 2026-07-06 | `a7971755` | **Fuzz fleet sprint (§4.4 + §4.6 nightly CI):** 28 fuzz targets + ~20 accompanying unit/property test funcs across 15 packages; `scripts/fuzz-all.sh`, `make fuzz-smoke`, nightly `.github/workflows/fuzz.yml`. Testability extractions (behavior-preserving, pinned by unit tests): metadata host dispatch → `metadataHostServesCosmicSignature` (`cmd/apiserver`), ORDER BY whitelists → `roiLeaderboardOrderClause` / `activeOffersOrderClause`. **Bug found & fixed:** corrupt freezer index entry could OOM-kill the process via an up-to-2^48-byte allocation in `FreezerReader.readBytes` / `WorkerReader.readBytes`; both now bounds-check against the data end (`TestReadItemCorruptIndexHugeOffset`). Test files 19 → 39. |
| 2026-07-07 | `aa9c686e` | **ETL fixture sprint (§4.3 complete + §4.2 write-path/trigger/base items):** `internal/testchain` (deterministic fake Ethereum node: real header hashes, signed txs so sender recovery works, receipts, `eth_getLogs`, registrable `eth_call` handlers, `Reorg()`) and `internal/testutil` (golden compare; canonical DB snapshot/diff with surrogate keys dropped and FKs resolved to natural identifiers). `cmd/cg-etl`: 74 per-event fixtures through the real pipeline with full trigger side effects pinned (84 goldens incl. scripted-round story + reorg rollback), every fixture re-processed to prove the delete-then-insert recovery path is state-neutral, plus a no-Docker unit test pinning all registry topic constants against ABI event IDs. `cmd/rw-etl`: all 7 event types, marketplace story, reorg test (13 goldens). `internal/etl`: blockops/chainsplit/tx-fallback/evt-log-dedup integration tests. **Seven production bugs found & fixed:** (1) `proc_admin_changed_event` unpacked ERC-1967 `AdminChanged` with the game ABI — the event is not in that ABI, so every occurrence killed the ETL (new `erc1967_abi`); (2) `proc_time_increase_changed_event` unpacked `TimeIncreaseChanged` by name from an ABI that doesn't define it — same crash mode (now decodes the raw uint256; `TestLegacyConstantsHaveNoABIEvent` guards the inverse direction); (3) `proc_token_generation_script_url_event` deleted from the *message-length* table before inserting, so re-processing a script-URL event aborted on the unique constraint; (4) `Delete_fund_transfer_failed_event`/`Delete_erc20_transfer_failed_event` referenced non-existent table names (`cg_fund_transfer_err`/`cg_erc20_transfer_err`) — any re-process/reorg of those events killed the ETL; (5) `last_block` was created rowless and every writer uses plain UPDATE, so on a fresh migrated DB the watermark never advanced and `HandleChainSplit` had nothing to roll back (migration `00005` seeds the row); (6) `on_nft_unstaked_{cst,rwalk}_delete()` never restored the staked-token row ("To Do" comment), so reorg rollback couldn't reverse reward deposits and replay double-counted `cg_staker_cst.total_reward` (migration `00006`); (7) `on_item_bought_delete()` referenced `NEW.*` in a DELETE trigger (always null → market volume/trade reversal silently skipped), restored the seller's `price_bought` to the sale price instead of the purchase price (profit became 0 on replay) and never reversed profit bookkeeping (migration `00007`). CI coverage ratchet floor raised 50% → 60% (measured 62.7%, up from 53.0%). Test files 44 → 53; golden files 183 → 280. |
| 2026-07-07 | `dd475c55` | **Store read suite + benchmarks sprint (§4.2 and §4.5 complete — Phase 0 done except prod-RLP replay):** shared seed dataset extracted to `internal/testfixtures` (embedded via `go:embed`, `Apply`/`ApplyFS`; apitest refactored onto it with parity goldens byte-identical, removing the CWD-relative glob). Store harnesses in `internal/store/{cosmicgame,randomwalk}` (TestMain + container + seed + wrapper); `testutil.CompareGoldenJSON`/`GoldenJSON` helpers (every golden fetched twice to prove determinism). 20 CosmicGame + 3 RandomWalk test files cover every public read function (~200 funcs, 196 goldens) incl. the notibot-only notification surface (rw_uranks extension seed), Elo transaction semantics (rollback/commit), nonce replay/expiry, processing-status and rank-writer round trips that restore fixture state. Benchmarks (§4.5): `BenchmarkEventDecode`, `BenchmarkReceiptsDecode`, `BenchmarkRateLimiter`, `BenchmarkStatisticsQueries`; baselines in `docs/benchmarks.md`. **Three production bugs found & fixed:** (1) `rw_messaging_status` was created rowless and `Update_messaging_status` uses a plain UPDATE, so on a freshly migrated DB the notibot watermark never persisted and every restart re-notified the entire event history to Twitter/Discord (migration `00008` seeds the row — same defect family as `last_block`/`00005`); (2) `Check_rwalk_token_exists` referenced placeholder `$2` while binding one argument, so PostgreSQL rejected every call — the error fell through to `return true, nil` ("token exists") and genuinely missing tokens returned `ErrNoRows`, which `rwctl scan-mints` treated as a transient DB error and retried forever; (3) the freezer receipts decoders detected "raw RLP" by first byte ≥ 0xc0, but snappy's decompressed-length uvarint starts with such a byte for half of all payload lengths > 127, making valid compressed blobs undecodable (`rlpListCoversExactly` now requires the list header to span the payload; applied to both `DecodeReceipts` and `DecodeArbitrumReceipts`, regression test added). CI coverage ratchet floor raised 60% → 64% (measured 66.7%, up from 62.7%). Test files 53 → 83; golden files 280 → 476. |
| 2026-07-07 | `dbf19cf1` | **Store groundwork + first conversion batch (Phase 1 kickoff: §5.1 complete, first three §5.2 rows, D3 + D7 decided):** base package renamed `dbs` → `store`; new `Store` on `*pgxpool.Pool` (`store.go`: `New`/`NewFromPool`/`Close`, keepalive dialer port, bounded startup ping-retry replacing `retryConnector`, pool-wide `timezone=UTC` + `search_path=public` runtime params, `DefaultMaxConns=16` — the legacy `*sql.DB` was unbounded), transitional `Store.DB()` `database/sql` view so every binary shares one pool, `ConnectHint` operator diagnostics; `errors.go` (`ErrNotFound`/`ErrConflict`/`WrapError`, multi-`%w` chains), `tracer.go` (slog `QueryTracer`: failed + slow queries, cancellations skipped), `scan.go` (`TimeText`: timestamptz → RFC3339Nano strings byte-identical to `database/sql`'s convertAssign, unit-tested against both formats). `Connect_to_storage`/`openDB` deleted; all 9 binaries converted (apiserver, cg-etl, rw-etl, notibot, imggen-monitor, cgctl ×2, rwctl, opsctl). `cosmicgame.Repo` (D3) + generic `queryList` helper preserving empty-slice JSON semantics; **24 functions across 8 files converted to context-first, error-returning, pgx-native methods** (`marketing.go`, `admin_events.go` — 39-branch UNION now registry-generated with completeness + valid-SQL tests —, `prize-history.go`, `main-prize.go`, `raffle-nft.go` — staker flag now a bound parameter —, `banned_bids.go`, `cosmicgame.go`, `tokens-erc20.go`, plus `PrizeDepositsByRound` from `raffle-eth.go` early because `PrizeInfo` composes it). ~20 API handlers pass `c.Request.Context()` and answer store failures with the new `respondStoreError` → `common.RespondInternalErrorJSON` (HTTP 500, legacy envelope, no internal detail; these paths previously killed the process). cg-etl loop reads/writes its watermark through the Repo and crashes only from `main`. `common.InitContext` carries the `*store.Store`; `testdb` exposes the container `pgxpool.Pool`; store-suite/apitest/cg-etl harnesses run the production one-pool wiring (`store(t)` helper renamed `wrapper(t)`, new `repo(t)`). **All 476 goldens byte-identical** (store suite, parity suite, ETL fixtures — pins the numeric→string, timestamp and bool scan semantics across the driver swap); new error-path tests land the first coverage the legacy `os.Exit` layer could never have: cancelled context, closed pool, `ErrNotFound` on missing round/status rows. Unit tests for Config/conn-string/ConnectHint (secret never echoed), error mapping, tracer output, TimeText. Metrics: snake_case 656 → 630, `os.Exit` in `internal/` ~490 → 469, dot-import files 21 → 19, lint 179 → 178, test files 83 → 87, integration coverage 66.8% (floor stays 64%). Statistics benchmarks re-run vs `docs/benchmarks.md` baselines: all three faster (2.66→2.39ms, 955→845µs, 313→267µs); B/op higher through the pool-backed `sql.DB` view (those queries are still on the legacy path — re-measure when `statistics.go` converts). |
| 2026-07-07 | `ca87801a` | **API parity suite sprint (§4.1 complete + §4.6 coverage ratchet):** `internal/api/apitest` boots the real gin router (production middleware chain, real Init sequence) against a seeded testcontainers Postgres and a deterministic Ethereum JSON-RPC stub; 183 golden files pin every registered GET route (each fetched twice to prove determinism), plus mutation-route tests (admin auth matrices, ban/unban round-trip, Elo match, EIP-191 signed `add_game` incl. replay/duplicate/chain rejections) and error-envelope goldens. Route-drift unit test proves `docs/openapi.yaml` ⇄ router equality (187/187 operations, both directions). Fixture dataset exercises the migration plpgsql triggers end-to-end. Supporting changes: `testdb.Start` for TestMain lifetimes, `DisableBackgroundRefresh` test hook in `state.go` (removed in Phase 2), metadata-host dispatch + health routes moved to `internal/api/common` for reuse. CI integration job now enforces the `internal/` coverage ratchet (floor 50%; measured 53.0%, up from 5.8%). **Three production bugs found & fixed:** (1) migrations 00002/00003 both defined `on_token_name_insert()`/`_delete()`, so the RandomWalk body silently overwrote the CosmicGame one and every `cg_token_name` insert failed — CS-NFT naming was broken and the ETL would crash on `NftNameChanged`; fixed by migration `00004` with per-project function names. (2) `Get_bid_frequency_by_period` / `Get_top_bidder_active_periods` passed Go ints into pgx text-concatenation placeholders (`$3 \|\| ' seconds'`), so `statistics/bidding/frequency` and `top_active_periods` failed on every call — and their `os.Exit(1)` error paths killed the whole API server when hit. (3) `Get_market_trading_volume_by_period` had a SQL typo (`TO_TIMESTAMP($1)i`), making `statistics/trading_volume` another process-killing route. Test files 39 → 44. |
| 2026-07-07 | `9018fcce` | **Read-layer conversion sprint 3 (§5.2: the four heavyweights — the CosmicGame read layer is now fully on the Repo):** `admin_events_resolve.go`, `staking.go`, `user-specific.go`, `statistics.go` converted to context-first, error-returning, pgx-native `Repo` methods (~52 public methods + ctx-aware helpers; `(bool, rec)` returns became `(rec, error)` + `ErrNotFound` on `UserInfo`/`StakeActionCstInfo`/`StakeActionRwalkInfo`); **every golden byte-identical except one deliberate fix** (see below). Safety/testability: the stake-action queries extracted to pure `stakeActionQueryCST/RWalk` functions with a no-Docker unit test pinning both production shapes; the admin-resolve lookups pass the `checkAdminIdent` guard; `RoiLeaderboard` keeps the fuzzed ORDER-BY whitelist; `BidsByUser`/`CosmicSignatureTokensByUser` reuse the whitelisted `bidList`/`nftListSelectSQL` builders and the legacy `buildBidSelectQuery`/`scanBidRecord` + `buildNFTSelectQuery`/`scanNFTRecord` twins are deleted (as are `donatedTokenDistributionLegacy` and the production-dead `Get_num_prize_claims`). Callers: all remaining CosmicGame read handlers (~35 sites incl. the big hybrid `api_cosmic_game_user_info`, the dashboard round-statistics call, 25 staking routes, admin-events resolve) now use `c.Request.Context()` + `respondStoreError`, with `ErrNotFound` mapped to the exact legacy not-found envelopes at all nine `Get_user_info` gates; `do_reload_database_variables` refreshes `bw_stats` via the Repo on `context.Background()` and keeps the previous value on failure. Store suite: the four integration test files moved to `repo(t)` (65 goldens unchanged), the legacy `wrapper(t)` harness deleted — the whole CG read suite runs the production one-pool wiring; `TestErrorPathsConvertedFiles` extended with 8 cancelled-context + 4 closed-pool cases; `BenchmarkStatisticsQueries` on the Repo (baselines re-recorded in `docs/benchmarks.md`: the stdlib-over-pool B/op inflation from the groundwork sprint is gone — statistics 40,830 → 14,390 B/op, claims 19,728 → 9,625 B/op; latency inside the container noise band). **Bug found & fixed:** `Get_staking_cst_mints_global` hardcoded `IsRWalk=true` on rows its own `WHERE is_rwalk=FALSE` filter selects (copy-paste from the RWalk variant), so `staking/cst/mints/global` mislabeled every CST-staker mint — store + parity goldens updated, regression assertion added. Metrics: snake_case 563 → 506, `os.Exit` in `internal/` 362 → 235 (the CG read layer is exit-free; 146 of the rest live in Phase-3 `inserts.go`/`deletes.go`), Repo methods 90 → 156, dot-import files 19 → 18, lint uncapped 1057 → 904 (capped 172), integration coverage 65.9% (−0.7pp: ~150 new error-only branches; floor stays 64%). |
| 2026-07-07 | `449dae2d` | **Read-layer conversion sprint 2 (§5.2: eight more files, 66 functions — the CosmicGame read layer is converted except the four heavyweights):** `bidding_analytics.go`, `raffle-eth.go`, `nft-donations.go`, `erc20-donations.go`, `tokens-erc721.go`, `contract_params.go`, `bidding.go`, `eth-donations.go` converted to context-first, error-returning, pgx-native `Repo` methods with idiomatic names; **all goldens byte-identical** (2 new store goldens: the epoch-aligned hourly bucket branch and first `SearchTokensByName` coverage). Safety/testability: `bidding.go`'s string-passthrough query builder replaced by `bidSelectQuery` over WHERE/ORDER BY/paging whitelists (`TestBidSelectQueryWhitelists` walks every combination and the rejection paths — request input can never reach ORDER BY); `bidding_analytics.go`'s bucket SQL extracted to `bidFrequencySQL` (unit tests pin epoch-aligned vs anchored branch selection and that only the integer interval is interpolated); `contract_params.go` admin table/column names pass a lowercase-identifier guard, and the `SyncAdmin*` check-then-correct policy moved out of storage into `cmd/cg-etl` (`paramSyncer`) with lazy address resolution preserved (a clean sync run leaves the address table untouched); `store.NullTimeText` added for nullable timestamps. Callers: ~60 more API handlers on `c.Request.Context()` + `respondStoreError`; `state.go` background refreshers keep the previous value on failure and log real errors; cg-etl donation handlers use adapters preserving the -1/0 sentinels (crash only on real DB errors until Phase 3); `cgctl total-tokens`/`token-seed` and `imggen-monitor` build the `Repo` directly. New tests: `TestErrorPathsConvertedFiles` (cancelled context + closed pool per file), malformed-identifier rejections, admin-correction insert round trip restoring fixture state. `Get_donated_token_distribution` stays as a private legacy copy inside unconverted `statistics.go` (dies with that file's conversion). Metrics: snake_case 630 → 563, `os.Exit` in `internal/` 469 → 362, Repo methods 24 → 90, lint uncapped 1210 → 1057 (capped display 179), test files 87 → 89, golden files 489 → 491, integration coverage 66.6% (floor stays 64%). |
| 2026-07-07 | `86b222ea` | **CG write-layer + ETL error-propagation sprint (§5.2 complete for CosmicGame: `deletes.go` + `inserts.go`, the §5.3 cg-etl item, and the §5.2-base address-cache item):** the 145 legacy write functions became context-first, error-returning, pgx-native `Repo` methods — `deletes.go` is one generic `deleteByEvtlogID` helper + 72 named methods, `inserts.go` 73 methods with `insertAdminValue` covering the ~35 single-value admin tables — and the CosmicGame `SQLStorageWrapper` (incl. `must_lookup_or_create_address`) is **deleted**. Address FKs resolve through the new `internal/store/address.go`: `Store.LookupOrCreateAddress`/`LookupAddressID` on the pgx pool with a bounded per-Store LRU (unit-tested incl. `-race`; insert races resolve via the unique constraint + re-read), pulled forward from the base-file batch. `cmd/cg-etl`: all 75 `proc_*` handlers take `ctx` and return errors (ABI-decode failures included — previously `os.Exit`), the if-chain dispatcher became a table that checks **every** handler's error (only bid v1/v2 were checked before), `process_single_event`'s RLP-decode `panic` is a returned error, and shutdown runs in-flight batch DB work on `context.WithoutCancel` so SIGTERM mid-batch still finishes the batch, writes status and exits 0 (previously the watermark write could fail with `context canceled` → exit 1). `internal/api/cosmicgame` dropped its wrapper handle (`arb_storage *store.SQLStorage` for the not-yet-converted base lookups; dead exported `ArbStoragew` deleted). **Behavior fixes:** (1) `InsertBid` no longer silently defaults `bid_position` to 1 when the position query fails — the legacy layer mislabeled every later bid of a round on any DB error; real errors now propagate. (2) The registry inspected `ERC20TransferFailed` (ICosmicSignatureErrors.sol) but never dispatched it, so fetched events were silently dropped with `cg_erc20_transf_err` forever empty; the event is now dispatched (`proc_erc20_transfer_failed_event` raw-decodes the body — no generated ABI carries the event; `TestERC20TransferFailedConstantMatchesSignature` pins the keccak signature and the no-ABI registry test guards the decode strategy) with fixture + golden. Dead `find_cosmic_token_721_transfer`/`find_cosmic_token_721_mint_event` (commented-out callers only) deleted. **Tests:** all 97 pre-existing ETL fixture goldens, replay-idempotence, reorg rollback, store read suite and the 183-golden parity suite **byte-identical** (1 new golden: `admin_erc20_transfer_failed`); new `TestWriteErrorPropagation` re-processes every fixture on a `default_transaction_read_only=on` pool and requires the error to surface from `process_single_event` for all 75 event types (and no error for the three no-write negative fixtures — proving their handlers write nothing); `TestDeleteMethodsValidSQL` reflection-sweeps all 73 `Delete*` methods against the real schema (the table-name-typo bug class found twice in §4.3); `TestLookupOrCreateAddress` integration round trip (create/cached/uncached/not-found/empty, first-seen block preserved); `TestErrorPathsConvertedFiles` extended with insert/delete cancelled-ctx + closed-pool cases. Metrics: snake_case 506 → 359, `os.Exit` in `internal/` 235 → **88** (store/cosmicgame production code exit-free; the rest is randomwalk 77 + api 7 + primitives 1 + test mains 3), Repo methods 156 → **304**, dot-import files 18 → 17, lint uncapped 904 → **674** (capped 170), test files 89 → 93, golden files 491 → 492, integration coverage 65.9% → **66.6%** (floor stays 64% until Phase 1 completes with the RandomWalk conversion). |
| 2026-07-08 | | **Phase 1 completion sprint (§5 done: RandomWalk + base store on pgx, legacy bridge deleted):** the RandomWalk store's 62 legacy methods became context-first, error-returning, pgx-native `Repo` methods (`repo.go` mirrors CosmicGame's; the shared scan loop extracted to `store.QueryList`); the ranking transactions moved from `*sql.Tx` to `pgx.Tx` (`ApplyRankingMatch`/`ConsumeRankingVoteNonce`); the base files (`blockchain.go`, `blockchain_insert.go`, `archive.go`) became 17 ctx-first `Store` methods and 6 dead functions were deleted; `lookups.go` with the process-wide `amap` address cache is gone (per-Store LRU everywhere, new `AddressByID`). `internal/etl` runs on `ETLContext.Store *store.Store` with context-aware helpers. **The legacy bridge is deleted:** `SQLStorage`, `NewSQLStorageFromDB`, `Init_log`/`Log_msg` (replaced by the pgx slog tracer writing to the same db.log files), the transitional `Store.DB()` `database/sql` pool view, `common.Ctx.Db`, and the RandomWalk `SQLStorageWrapper`. Callers: ~49 RandomWalk API handler sites on `c.Request.Context()` + a package `respondStoreError` (not-found flows render the byte-identical legacy `DBError`/error strings via `store.ErrNotFound` mapping — pinned by the parity goldens incl. `errors__missing_rw_token`); ~39 CG API address-lookup sites on `Store.LookupAddressID`; three charity routes no longer `os.Exit` inside a request handler (a client disconnect could previously kill the whole API server once lookups became ctx-aware); rw-etl mirrors cg-etl (dispatch table checks every handler, RLP-decode panic → error, SIGTERM finishes the in-flight batch on `context.WithoutCancel`); notibot/rwctl/opsctl on Repo + Store with checked errors; `opsctl archive node-fill` resolves addresses through a pgx `Store`; `/readyz` pings `Store.Pool()`. Raw DDL dirs `db/{layer1,cosmicgame,randomwalk}` deleted (goose migrations are the only schema source; node-fill's help text updated) and the never-imported sqlc scaffolding retired (D7 amended). **Behavior fixes:** (1) `OfferExists`/`TokenExists` treated any DB failure during the existence check as "does not exist" and silently skipped the event — a data-loss bug; real errors now abort the batch for re-processing. (2) The rw-etl ABI-decode failures no longer kill the process mid-batch. **Tests:** all 492 goldens **byte-identical** (parity, CG+RW store suites, both ETL fixture suites incl. replay-idempotence and reorg; RW store suite rebuilt on `repo(t)` + production one-pool wiring); new rw-etl `TestWriteErrorPropagation` (read-only pool; all 7 event types must surface write errors, the 3 no-write negative fixtures must stay clean); new RW `TestErrorPathsConvertedFiles` (22 cancelled-ctx + 10 closed-pool cases across all three files); `TestStoreBaseHelpers` (AddressByID incl. ErrNotFound, case-insensitive `CountEvtLogsForContract`, `EvtLogExists`); blockops suite ported to the ctx-first API. Statistics benchmarks re-checked: no regression (2.24ms/787µs/259µs vs 2.53/936/315 baselines, B/op identical). Metrics: snake_case 359 → **267** (store layer 0), `os.Exit` in `internal/` 88 → **12 matches / 5 real calls** (3 test mains + startup fatal + `primitives.Fatalf`), Repo methods 304 → **366** (+ 22 base Store methods, ctx coverage 100%), lint uncapped 674 → **492** (capped 145), test files 93 → 95, LOC 68,252 → 66,735, integration coverage 66.6% → **67.2%** (CI floor raised 64% → 66%). |
| | | |
