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

Measured 2026-07-12 (command-seam sprint slice 2a). Update after each phase.

| Metric | Baseline (start of project) | Current | Target | How to measure |
|---|---|---|---|---|
| Hand-written Go LOC (`cmd/` + `internal/`) | ~60,000 (plus 124k frozen legacy) | **113,213** (+6,634 this sprint: `internal/autobid` and the cgctl command/error/planner suites, net of the deleted 1,221-line bot and the duplicated cgctl ethtx); generated `internal/api/v2/api.gen.go` is separate | n/a (informational) | `rg --files internal cmd -g '*.go' -g '!internal/api/v2/api.gen.go' \| tr '\n' '\0' \| xargs -0 wc -l \| tail -1` |
| snake_case functions | ~700 | **135** (api 127, rwalk-alarm 4, primitives 3, apiserver 1; **the store layer, both ETLs, the notify stack and all three slice-1 binaries are 0**) | **0** | `rg "^func (\([^)]+\) )?[A-Za-z]+_[A-Za-z0-9_]*\(" --type go -c internal cmd` |
| `os.Exit` in library code (`internal/`) | ~560 | **15 matches = 6 real calls** (5 test-harness `TestMain`s + `primitives.Fatalf`; the rest are doc comments). | **0** (allowed only in `cmd/*/main.go` startup) | `rg -c "os\.Exit" internal` |
| Dot-import files | ~70 | **3** (apiserver 1, api/cosmicgame 2) | **0** | `rg -l '^\s*\. "github' --type go internal cmd contracts` |
| Package-level mutable globals (api + etl) | ~120 | ~12 (**legacy v1 API only**; the new `v2.Server` and both ETL binaries have zero package-level mutable state) | ~0 (DI everywhere) | manual review per package |
| golangci-lint issues | 433 (first run) | **74** capped / **335** uncapped (touched packages are clean; new diff is zero) | **0** | `golangci-lint run` |
| Test files | 17 | **218** | 100+ — met | `rg --files -g '*_test.go' \| wc -l` |
| Fuzz targets | 0 | **47** (including the wallet-scoped user-bid cursor plus all prior bounded opaque-cursor decoders, bidding and cached-contract-state invariants) | **25+** (see §4.4) — met | `rg "func Fuzz" internal cmd contracts -c` |
| Benchmarks | 0 | **4** (18 sub-benchmarks; baselines in `docs/benchmarks.md`) | keep green vs baselines | `rg "func Benchmark" cmd internal -c` |
| Coverage on `internal/` (unit) | 2.4% | **29.3%** | superseded by the integration ratchet below | `go test -coverprofile=c.out ./internal/... && go tool cover -func=c.out \| tail -1` |
| Legacy `internal/` integration coverage (enforced) | n/a | **89.19%** (floor 88.9%; original denominator retained for historical continuity) | ratchet only; never lower | `make coverage-check` |
| Handwritten production `internal/` coverage (enforced) | n/a | **92.34%** (floor 92.1%; generated API + test-only harnesses excluded by ADR-0006) | **≥90% — met**; commit gate active, next target 95% | `make coverage-check` |
| Handwritten all-production coverage (`cmd/` + `internal/`, enforced) | n/a | **76.31%** (floor 76.0%; notibot/rwctl/freezer-scan/cgctl covered by command-seam slices 1–2a; opsctl/srvmonitor remain) | ratchet toward **≥90%** | `make coverage-check` |
| Changed executable Go coverage (enforced) | n/a | per staged/PR diff | **≥95%** | `make coverage-check` |
| Queries on sqlc | 0 | 0 — scaffolding retired with Phase 1 (D7 amended: hand-written pgx everywhere) | n/a | n/a |
| Routes on stdlib router | 0 | **all** — frozen v1 (187 OpenAPI operations) plus 34 generated v2 operations (round resources, exact statistics/claims, participant/bidding analytics, five contract/live resources and the user profile/bid foundation), health, host-dispatched metadata and env-gated static assets on `net/http` via `internal/api/httpx`; **gin is out of the build graph** | all (v1 compat + v2) — active | route-drift tests + `go list -deps ./cmd/... ./internal/... \| rg -c gin-gonic` |
| `context.Context` on store methods | 0% | **100% — 396 Repo methods (CosmicGame 334 + RandomWalk 62) + 23 base `Store` methods; `SQLStorage` and both wrappers are deleted** | 100% | `rg -c "func \(r \*Repo\)" internal/store/cosmicgame internal/store/randomwalk` |
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
`user/balances` and the dashboard) were pinned in their deterministic
"RPC unavailable" shape against a stub Ethereum JSON-RPC server.
*Resolved 2026-07-09 (ContractState sprint): the harness now dials
`internal/testchain` serving fixture-coherent contract state through
ABI-driven stubs (`testchain.ContractStub`), and all 10 affected goldens are
regenerated in their happy-path shape. The degraded "RPC unavailable"
sentinels stay pinned by unit tests in
`internal/api/cosmicgame/contractstate`.*

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
      the pipeline integration suite (§4.3 sprint; now
      `internal/indexer/pipeline_integration_test.go` since the Phase-3
      engine extraction): block insert + hash verification + watermark,
      chain-split cascade, tx three-level fallback (RPC / archive / minimal),
      evt_log dedup-replace, FilterLogs range+address filtering *(2026-07-07)*
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

- [x] `cmd/cg-etl/fixtures_test.go` (now
      `internal/indexer/cosmicgame/fixtures_test.go` since the Phase-3
      handler port): one fixture per dispatched event type
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
- [x] Same for `cmd/rw-etl` (now `internal/indexer/randomwalk`; all 7
      dispatched event types + skip paths for
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
- [x] `FuzzEventDecodeCG` — landed 2026-07-09 with the Phase-3 handler port:
      iterates every registered handler in `internal/indexer/cosmicgame` and
      fuzzes its `Decode` with arbitrary topics/data (never panics; indexed-
      topic bounds now checked instead of trusted). The ABI-level
      `FuzzABIEventUnpack` in `contracts/cosmicgame` stays as the lower layer.
- [x] `FuzzEventDecodeRW` — same, over `internal/indexer/randomwalk`
- [x] `FuzzEvtlogRLP` — lives in `internal/indexer` (moved with
      `LogProcessor`, the real decode site of stored `log_rlp`); includes
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
- [x] `FuzzDecodeBidCursor` — `internal/api/v2`: arbitrary and oversized
      opaque cursors never panic or bypass version/round/keyset validation
- [x] `FuzzDecodeRoundCursor` — `internal/api/v2`: completed-round cursors
      enforce length, encoding, version and descending keyset invariants

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
- [x] World-class coverage policy and shared CI/commit gate implementation
      *(2026-07-11 — [ADR-0006](adr/0006-coverage-policy.md):
      `internal/covergate` deduplicates multi-binary Go profiles and enforces
      the versioned `coverage/policy.json` in the required CI Coverage Gate.
      The tracked native/Cursor commit hooks are installed but deliberately
      remain deferred while coverage climbs toward 90%. Generated API and
      test-only harness code are excluded only from the new canonical metrics,
      never from the historical continuity metric. The authoritative
      race-enabled Sprint 1 baseline was 79.07% legacy / 80.43% handwritten
      internal / 52.80% all production.)*
- [x] Coverage-quality sprint 2: harden API boundaries and raise handwritten
      internal coverage past the 85% milestone *(2026-07-12 — exhaustive
      malformed-parameter, unindexed-wallet, missing-record, cancellation and
      partial-dependency failure matrices now exercise the real v1 router;
      shared HTTP, contract-state, RandomWalk ranking and v2 mapping/cursor/page
      invariants have direct tests. Sixty-eight path-parameter error branches
      made unreachable by Go 1.22 `ServeMux` matching were deleted rather than
      covered artificially. Canonical package coverage is now 98.21%
      `api/common`, 91.20% `api/cosmicgame`, 91.45% `contractstate`, 87.11%
      `api/randomwalk` and 92.09% `api/v2`. The authoritative profile closes
      at 83.73% legacy / 86.07% handwritten internal / 56.26% all production /
      97.52% changed code; floors ratchet to 83.5% / 85.8% / 56.0%, with
      changed executable Go still requiring 95%.)*
- [x] Coverage-quality sprint 3: raise handwritten internal from 86.07% to at
      least **88%** with store/indexer replay branches, notification adapters
      and the remaining high-risk API gaps; then ratchet every measured floor.
      *(2026-07-12 — landed at **91.82%**, past both the 88% sprint goal and
      the 90% activation milestone: the notification adapters got real
      conformance suites (OAuth1 known-answer vector, full chunked-video
      upload protocol, WhatsApp Graph API) after trimming the vendored dead
      surface; the freezer readers/decoders, store fallback/report branches,
      indexer chain-sync corrections and the v1 API's live-RPC failure
      matrices are covered behaviorally. See the progress log for the bug
      fixes and deletions.)*
- [x] Activate the fail-closed Git commit coverage gate **only after**
      handwritten production `internal/` coverage reaches at least **90.0%**
      under the authoritative race-enabled profile. In that same change, set
      `commitGateEnabled=true`, set both the commit and internal floors to at
      least 90.0%, verify a below-90 commit is rejected and a ≥90 commit
      succeeds, and never lower the floor afterward.
      *(2026-07-12 — activated in the sprint-3 change at 91.82% measured:
      `commitGateEnabled=true`, `commitFloor` 90.0, `internalFloor` 91.5
      (never lower). Verified `covergate -commit-status` reports
      `enabled 90.00`, the authoritative profile passes all floors, and the
      same profile against a higher-floor policy fails closed with
      `handwritten internal coverage 91.82% is below 95.00%`; the Cursor
      hook now denies `--no-verify` and the hook tests pin the enabled
      state. The next milestone is the 95% internal target.)*
- [~] Command-seam sprint: extract the operational logic of the zero-coverage
      binaries (`cmd/cgctl` 2,260 stmts, `cmd/opsctl` 1,973, `cmd/rwctl` 979,
      `cmd/srvmonitor` 1,128, `cmd/notibot` 438, `cmd/freezer-scan` 318)
      behind testable seams and raise handwritten all-production coverage
      from 59.65% toward 90%; ratchet `productionFloor` as each binary lands.
      Split into three slices:
  - [x] Slice 1 — notification stack + small binaries *(2026-07-12:
        `internal/notify/rwbot` unifies the duplicated notibot/rwctl bots
        behind injected seams; `internal/ethtx` extracts the shared signing
        session; `internal/freezer/scan` extracts the scanner pipeline;
        `internal/testchain` gains transaction submission. notibot, rwctl
        and freezer-scan are thin tested wiring; all-production coverage
        59.65% → 66.31%, floor 59.4% → 66.0%. See the progress log.)*
  - [x] Slice 2a — `cmd/cgctl` *(2026-07-12: `cgctl/internal/ethtx` merged
        into `internal/ethtx` (session-scoped GAS_PRICE_MULTIPLIER, CG gas
        limits, dev-chain time advance, `Session.Refresh`, format/output
        helpers); the 1,221-line autobid bot became `internal/autobid` — a
        pure `Decide` core plus an injected engine at 98.1% coverage; all
        ~30 subcommands are thin tested wiring over the shared session, and
        every transaction command now waits for its receipt (reverts were
        previously invisible). All-production coverage 66.31% → 76.31%,
        floor 66.0% → 76.0%. Slice 2 was split: cgctl signs real
        transactions and warranted its own sprint. See the progress log.)*
  - [ ] Slice 2b — `cmd/opsctl` (extract archive export/verify/node-fill,
        db verify/evtlog-diff, tx-collector, assets and smoketest engines
        behind seams; deduplicate the three FilterLogs retry loops)
  - [ ] Slice 3 — `cmd/srvmonitor` (relocate `cmd/srvmonitor/internal/*` to
        `internal/srvmonitor`, inject config/manager, add monitor tests)
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

- [x] Resource-oriented paths: `/api/v2/cosmicgame/rounds/{round}/bids`
      instead of `/api/cosmicgame/bid/list/by_round/:round_num/:sort/:offset/:limit`
- [x] Pagination via query params (`?limit=&cursor=`; D2 decided in §11),
      never path segments; consistent `meta` block in list responses
- [x] Typed top-level responses (no `{"status":1,"error":""}` envelope);
      errors as RFC 9457 `application/problem+json`
- [x] Consistent field naming (camelCase JSON, ISO-8601 timestamps, amounts as
      decimal strings with explicit `*Wei`/`*Eth` fields)
- [x] OpenAPI-first: author `docs/openapi-v2.yaml`, generate server stubs +
      typed models (oapi-codegen); handlers implement the generated interface
- [x] Versioning and deprecation policy section in the ADR (v1 sunset criteria)

      *(2026-07-10 — [ADR-0005](adr/0005-api-v2.md) accepted. D2 is a
      bounded, versioned opaque cursor with endpoint-scoped keyset state; D6
      requires known-consumer migration, 30 zero-traffic days and an announced
      not-before date. [openapi-v2.yaml](openapi-v2.yaml) is OpenAPI 3.0.3
      while stable oapi-codegen lacks 3.1 support; generated strict stdlib
      interfaces/models are pinned and checked for drift in CI. The first two
      slices implement completed-round and round-bid collection/item resources
      with camelCase models, UTC RFC3339 timestamps, exact decimal-string wei
      amounts and RFC 9457 errors.)*

### 6.2 Implementation

- [x] `internal/api/v2/` package: `Server` struct with injected `*store.Store`,
      contract-state cache, logger; no package-level state
      *(2026-07-10 — `v2.Server` injects the shared store/repo,
      `contractstate.State` and slog logger; generated routes register directly
      on `httpx.Router` through its stdlib mux seam. Round-bid pages use
      `(bid_position, evtlog_id)` keyset queries backed by migration 00009;
      malformed/cross-round cursors and invalid limits are 400 problems,
      missing items are 404, and internal details never escape 500s.)*
- [x] stdlib `http.ServeMux` with method+pattern routes; middleware as
      `func(http.Handler) http.Handler` chain (port CORS, rate limit, auth,
      metrics, recovery from gin) *(2026-07-09 —
      [internal/api/httpx](../internal/api/httpx): `Router` over
      `http.ServeMux` (route registry, global + per-route middleware,
      registration-time conflict panics, gin-parity trailing-slash redirect
      hardened against scheme-relative Locations) and `Context` reproducing
      the exact wire behavior the goldens pin (marshal-before-write JSON,
      legacy binding semantics incl. EOF error text); middleware in standard
      form in `internal/api/common`: `CORS`, `Recovery` (broken-pipe aware,
      re-panics `ErrAbortHandler`), `AccessLog` (slog; replaces gin.Logger),
      `RateLimit`, `RequireAdminKey`, plus the Prometheus middleware in
      `cmd/apiserver` reading `Request.Pattern`. All unit-tested.)*
- [x] Contract-state cache: extract the ~70 globals + 3 refresh goroutines from
      `internal/api/cosmicgame/state.go` into an injected `ContractState`
      component with lifecycle (`Run(ctx)`), RWMutex-guarded snapshot reads,
      and unit tests with a mocked eth client *(2026-07-09 —
      [internal/api/cosmicgame/contractstate](../internal/api/cosmicgame/contractstate):
      `State` with `New(Config)`, `LoadInitial(ctx)`, `Run(ctx)` on tickers
      (stops on cancel; replaces the unkillable `for{}` goroutines and the
      `DisableBackgroundRefresh` test hook), `Snapshot()` value copies,
      `SetBidPrice` write-back, `FetchLiveSpecialWinners(ctx)`, mechanics
      V1/V2 detection cache; the ~70 state globals, `contract_live_reads.go`
      package state and 7 dead never-read globals (with their 3 refresh-cycle
      contract calls) are deleted. `cosmicgame.Init(ctx, ...)` returns an
      error instead of calling `os.Exit`; handlers read one `Snapshot()` per
      request. Unit-tested against `testchain` + a new ABI-driven
      `testchain.ContractStub`: happy path pins every field, failure tests
      pin every "error"/-1/0 sentinel (reachable-node and dead-node cases),
      mechanics V1-only/V2-only/upgrade-flip, `Run` refresh + cancellation,
      concurrent snapshot/refresh under `-race`, special-winners happy path +
      ErrNotFound + RPC-failure. Full handler-level DI stays with the v2
      `Server` item.)*
- [x] v1 compatibility layer: existing 187 routes re-registered on the new
      router calling the same service code; parity suite green *(2026-07-09 —
      handlers kept their bodies and moved onto `*httpx.Context`; patterns
      converted `:param` → `{param}` (`*filepath` → `{filepath...}`); the
      duplicated main.go/apitest wiring replaced by one shared constructor,
      [internal/api/routes](../internal/api/routes)`.New(st, Options)` —
      production injects access log, metrics and static assets through
      Options, the harness runs the same chain minus those. All 196 parity +
      12 error-shape goldens byte-identical; route-drift test compares the
      OpenAPI spec directly against the router registry (one syntax now).
      Three deliberate router-level deltas, pinned by
      `apitest/router_behavior_test.go`: wrong method answers 405 + `Allow`
      (was 404), the router 404 body gained stdlib's trailing newline, HEAD
      is served by GET routes (was 404).)*
- [x] Graceful shutdown: `http.Server.Shutdown` on SIGTERM (replaces
      `gin.Run` + `select {}`); readiness flips false during drain
      *(2026-07-09 — `signal.NotifyContext` root ctx; all public listeners
      (TLS + plain) run as tracked `http.Server`s over explicit
      `net.Listen`ers with `ReadHeaderTimeout`; on SIGINT/SIGTERM `/readyz`
      answers `503 {"status":"draining"}` (`common.SetDraining`, unit-tested)
      while in-flight requests get 15s to finish; the internal metrics/pprof
      server shuts down too, refresh loops stop via ctx, the store pool
      closes last. The dead autocert manager is deleted and
      `golang.org/x/crypto` demoted to an indirect dependency.)*
- [x] Remove gin from go.mod once v1 compat runs on stdlib *(2026-07-09 —
      zero gin imports remain; `go mod tidy` dropped the direct requirement
      and the whole ecosystem (gin-contrib/sse, go-playground/validator,
      bytedance/sonic, ugorji). `go list -deps` confirms gin is linked into
      zero packages; one `// indirect` go.mod line survives because notibot's
      Discord library (disgord → nhooyr.io/websocket) lists gin among its
      test dependencies — never compiled. `cmd/loganomaly` now parses the
      slog access-log format alongside legacy `[GIN]` lines in old files.)*
- [ ] Response compression + ETag/Cache-Control on hot read routes
- [x] httptest suite for v2 (same fixtures as §4.1, new goldens)
      *(2026-07-11 — 213 deterministic v2 goldens cover current/completed
      rounds, round bids, prizes, raffles, donations, global statistics,
      counters, all ROI sorts, bounded claim summaries/details and all six
      participant directories, five bidding analytics resources and five
      contract/live resources plus the user profile/bid foundation: both
      mechanics generations, cache
      failure/recovery, live-v1/cached-v2 semantic equivalence, keyset pages,
      lean items, bounded time series, decimal-string percentages,
      bind/limit/cursor/sort/window/pool errors, open-round donation pages,
      indexed/unindexed zero-user shapes, cross-wallet cursors, open/missing
      404s for completed-only resources and cancelled-context 500s; v1's 196
      parity + 12 error goldens were not regenerated.)*
- [x] OpenAPI contract validation in tests (kin-openapi response validator)
      *(updated 2026-07-11 — the embedded spec is validated, spec and generated router
      are compared bidirectionally (34 operations), and every v2 golden
      status/header/body is response-validated with kin-openapi.)*

#### 6.2.1 V2 endpoint slices

- [x] Round bids: list + item under
      `/api/v2/cosmicgame/rounds/{round}/bids`
- [x] Completed rounds: cursor-paginated `/api/v2/cosmicgame/rounds` + lean
      `/api/v2/cosmicgame/rounds/{round}`; nested prize/raffle collections
      deliberately remain future sub-resources
- [x] Current/open round live state: `GET /api/v2/cosmicgame/rounds/current`
      combines one atomic contract-state snapshot with per-round aggregates
      and the authoritative bid count; exact wei/microsecond strings, UTC
      timing, corrected `secondsUntilMainPrize` semantics, no float ETH.
      Uninitialized/failed cache sentinels return an RFC 9457 503 with
      `Retry-After: 5`; corrupt data/store failures are opaque 500s. No
      request-time RPC fallback or global-dashboard fields in this slice.
- [x] Round prize/raffle/donation sub-resources: cursor-paginated
      `/api/v2/cosmicgame/rounds/{round}/prizes` landed with all 16 typed
      `cg_prize.ptype` variants, exact asset fields, completed-round 404 gate,
      shared strict cursor/page-limit primitives and matching
      `(round_num, ptype, winner_index)` index. Round-scoped
      `raffle-eth-deposits` and pool-selected `raffle-nft-winners` now expose
      exact payout records with claimed state and indexed round/pool keysets.
      `eth-donations`, `erc20-donations` and `nft-donations` expose immutable
      newest-first events with exact ETH wei/ERC-20 base-unit strings,
      info-record and NFT contract identities, cross-resource-safe cursors and
      matching `(round_num, evtlog_id DESC)` indexes. Donation pages remain
      available during open rounds and return empty 200 pages when no events
      exist; charity/user histories, summaries and claim events remain
      separate future resources.
- [x] Statistics/dashboard resources: exact DB-backed
      `/api/v2/cosmicgame/statistics` and `/statistics/counters` now complement
      `/rounds/current` without rebuilding the v1 dashboard mega-response or
      making request-time RPC calls. A six-mode, sort/filter-scoped keyset
      `/statistics/leaderboard/roi` exposes exact signed wei and decimal-string
      ratios. Cursor-paginated `/statistics/claims` summaries omit v1's
      unbounded inline assets; `/rounds/{round}/claims` provides completed-only
      bounded transaction, attached-token and unclaimed-item sections with
      exact ETH wei/ERC-20 base units. Six cursor-paginated participant
      directories now expose exact bidder, winner, donor, CST-staker,
      RandomWalk-staker and dual-staker projections with endpoint-scoped
      keysets; their mutable-rank traversal semantics are explicit. Five
      bounded `/statistics/bidding/*` resources expose frequency/activity,
      decimal-string bid-type percentages, lifetime-top-bidder periods and
      time bounds without unbounded `generate_series` or internal IDs.
      Migration 00016 backs half-open timestamp filtering; bids are assigned
      once with `DATE_BIN`, active periods cap at 2,000 and all analytics
      reads carry a five-second deadline. Five typed contract/live resources
      expose the complete address registry, mechanics-aware configuration,
      exact balances, block-pinned bid prices/reward/auction progress and
      block-coherent special winners. V2 reads only coherent caches (5m
      constants, 5s variables/balances, 30s champions), uses generation and
      address tags to prevent mixed snapshots, and never performs
      request-time RPC. [api-v2-migration.md](api-v2-migration.md) maps every
      v1 dashboard field to its replacement without recreating the mega-response.
- [~] User resources *(2026-07-11 — foundation landed:
      `/api/v2/cosmicgame/users/{address}` is an exact, collection-free
      profile over canonical bid/prize rows plus transfer/donation/staking
      statistics; valid unindexed wallets return the same checksummed zero
      shape. `/users/{address}/bids` reuses the typed Bid model and pages
      newest-first on a wallet-scoped event-log cursor backed by migration
      00017. Remaining prize/donation/claim/staking/transfer/token/marketing
      histories and live balances stay separate bounded slices.)*
- [ ] RandomWalk resources

### 6.3 Frontend migration

- [~] Publish v2 spec + changelog mapping every v1 path to its v2 replacement
      *(dashboard and user profile/bid mappings published in
      [api-v2-migration.md](api-v2-migration.md); remaining user histories,
      RandomWalk, CosmicToken and marketing groups remain)*
- [ ] Frontend switches endpoint-group by group (tracked as external checklist)
- [ ] v1 marked deprecated in spec; add `Deprecation`/`Sunset` headers
- [ ] Remove v1 layer + its goldens when traffic hits zero (final step, gated)

---

## 7. Phase 3 — ETL engine rewrite

Goal: one shared, tested indexing engine; the two binaries become thin
configuration of it. §4.3 fixtures must be green before starting.

**Status: complete as of 2026-07-09 (EventHandler-port sprint)** — the
engine core (loop, retry, metrics, slog) landed in the indexer-engine
sprint; the handler port finished the phase. Both ETL binaries are pure
wiring; the only open item is the deliberately deferred per-event
transactional status (see below).

- [x] `internal/indexer` package: `Engine` struct (rpc client, store, registry,
      slog logger, batch policy) with `Run(ctx) error` *(2026-07-09 — deps are
      injected via `indexer.Config`: `Store`, a narrow `Client` interface
      (satisfied by `*ethclient.Client`, faked in tests), a `Progress`
      watermark adapter per binary (preserves each domain's `last_evt_id`
      column), a `ProcessFunc` and the FilterLogs contract set. The handler
      registry stays in the binaries as `ProcessFunc` until the EventHandler
      port below; blockops/chainsplit/backfill moved in as Engine methods
      with their integration suite (`pipeline_integration_test.go`).)*
- [x] `EventHandler` interface: `Topic() common.Hash; Decode(types.Log) (Event, error);
      Store(ctx, Store, Event) error` — decode separated from persistence
      (enables the decode-only fuzz/golden tests to bypass the DB)
      *(2026-07-09 — landed as `indexer.EventHandler` with `Name()` (metric
      label) and `Sources()` (emitting contracts; the registry filters by
      source before Decode, replacing the in-handler address guards).
      `indexer.NewHandler[E]` keeps per-event type safety between the decode
      and store steps; `indexer.Registry` supports multiple handlers per
      topic0 (the two CharityAddressChanged meanings, the two
      FundsTransferredToCharity emitters, ERC721-vs-ERC20 Transfer) and
      validates at construction that same-topic handlers agree on one metric
      name. `indexer.LogProcessor(store, registry)` replaces the two
      byte-identical `process_single_event` copies and produces the engine's
      `ProcessFunc`. All unit-tested without Docker.)*
- [x] Port all ~80 CosmicGame handlers from `proc_*` functions to the registry;
      delete `cmd/cg-etl/events_*.go` bodies as they move *(2026-07-09 — the
      76 dispatched handlers (78 registrations incl. the Transfer split)
      live in `internal/indexer/cosmicgame` as decode/store method pairs on
      an injected `Handlers` set (`Config{Repo, Store, Caller, Contracts,
      Logger}`); ABIs parse once in `New`; `BootstrapContracts` replaces the
      main()/harness duplicated address bootstrap; the `decode.go` helpers
      became `Handlers` methods; `contract_sync.go` moved along as
      `SyncContractParams`. Decode steps are total: indexed-topic counts are
      bounds-checked (a malformed log now errors the batch instead of
      panicking the process).)*
- [x] Port the 9 RandomWalk handlers *(2026-07-09 —
      `internal/indexer/randomwalk`, same pattern; the existence-guard skips
      (unknown offer/token) live in the store steps where the DB read
      happens)*
- [x] Replace package-level globals (eclient, ABIs, `evt_*` vars, storagew,
      Info/Error) with injected dependencies *(2026-07-09 — both ETL binaries
      have zero package-level variables; every handler dependency is a
      `Handlers` field. The `evt_*` byte-slice vars are gone — registration
      uses the topic constants directly; handler logging is one structured
      slog record per event through the injected logger (the dual-file
      handler in production).)*
- [x] Batch loop: context-aware retry with exponential backoff + jitter on RPC
      and DB errors (no more crash-per-blip); circuit-break to exit only after
      N consecutive failures *(2026-07-09 — any failed batch retries from the
      last fully completed block with exponential backoff (±25% jitter, 1s→60s
      cap); `Run` returns only after `MaxConsecutiveFailures` (10) failures in
      a row, so systemd restarts resume from the watermark. Two data-loss bugs
      of the legacy loops fixed on the way — see the progress log.)*
- [ ] Status/progress persisted transactionally with the batch's inserts
      *(deliberately deferred, re-affirmed with the EventHandler port
      2026-07-09: per-event delete-then-insert replay is idempotent (pinned
      by the §4.3 replay tests), so batch-level transactionality buys no
      correctness today. The `Store(ctx, event)` seam is where a per-event
      `pgx.Tx` would slot in, but it still needs a querier abstraction
      across the 366 repo methods and a tx-aware address cache — not worth
      it for zero behavior change. Revisit only if a non-idempotent handler
      ever appears.)*
- [x] `log/slog` structured logging (block ranges, event counts, timings);
      keep file output via slog handler during transition *(2026-07-09 — the
      engine, startup sync and both mains log structured records through
      `indexer.NewDualLogHandler` into the legacy two-file layout (info +
      errors-duplicated-to-error-file). EventHandler-port sprint: the
      handlers emit one structured record per processed event through the
      injected logger — the `Info.Printf` multi-line dumps and the `Info`/
      `Error` `*log.Logger` globals are gone. §8.3 replaces file logging.)*
- [x] Prometheus metrics: `rwcg_etl_last_block`, `rwcg_etl_events_total{type}`,
      `rwcg_etl_batch_duration_seconds`, `rwcg_etl_reorgs_total`
      *(2026-07-09 — plus `rwcg_etl_batch_failures_total{stage}` for alerting
      on retry storms. The `type` label comes from the dispatch tables, which
      now carry event names (duplicate-topic entries assert one shared name).
      Both ETLs serve `/metrics` + pprof on `METRICS_ADDR` via
      `indexer.StartMetricsServer` — same private-listener rules as the API
      server (docs/operations.md).)*
- [x] `contract_sync.go` startup sync ported and unit-tested with mocked reads
      *(2026-07-09 — loggers ported to slog; V1/V2 mechanics probe and every
      versioned read fallback unit-tested against real abigen bindings over
      `testchain.ContractStub` (no Docker); the check-then-correct policy
      integration-tested end to end: fresh-DB corrections, clean-re-run
      no-op incl. the untouched address table, targeted correction on an
      on-chain change, and the skip-unreadable-params degraded mode.)*
- [x] Fixture replay + golden DB tests pass unchanged against the new engine
      *(2026-07-09 — both harnesses push fixtures through the Engine pipeline
      methods; all 492 goldens byte-identical, zero regenerated. The `Run`
      loop itself is covered by new unit tests (breaker, backoff, batch
      adaptation, caught-up/shutdown) and integration tests (end-to-end
      batches, transient-failure recovery, mid-block failure regression,
      shutdown-mid-batch, reorg-through-the-loop, backfill idempotence).)*

---

## 8. Phase 4 — Cross-cutting Go polish

### 8.1 Naming (snake_case → idiomatic) — per package

Mechanical but wide; use gopls rename / IDE refactor per identifier, one
package per PR, parity suite green after each.

- [ ] `internal/store` + subpackages (387 funcs — mostly renamed during Phase 1; this item is the sweep for leftovers)
- [ ] `internal/api` (135 funcs; handler names like `api_cosmic_game_bid_list` → `handleBidList`)
- [x] `cmd/cg-etl` — done by the Phase-3 handler port (0 snake_case funcs;
      the SCREAMING_SNAKE topic constants moved as-is to
      `internal/indexer/{cosmicgame,randomwalk}/topics.go` and are this
      item's remaining sweep)
- [x] `cmd/notibot` — done by the command-seam slice-1 rewrite (0 snake_case
      funcs; the bot logic lives in `internal/notify/rwbot`)
- [ ] ~~`cmd/rw-etl`~~ (0 — handler port), `cmd/rwalk-alarm` (4)
- [ ] `internal/primitives` (3) + rename package to `internal/model` (or fold
      types into their owning packages — decide in §11)
- [ ] Local variables and struct fields in touched files follow along
      (err_str → errStr, bid_position → bidPosition, ...)

### 8.2 Imports and files

Eliminate all dot-imports (3 files left after the command-seam slice-1
rewrite removed notibot's): `internal/api/cosmicgame` (2), `cmd/apiserver` (1).

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
- [x] Coverage ratchet reaches ≥70% on `internal/`; gate is enforced in CI
      *(73.0% measured after the first v2 slice; floor raised to 72%.)*
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
| D2 | v2 pagination | offset+limit / opaque cursor | **decided 2026-07-10: opaque cursor + limit** — versioned base64url payloads are bounded and endpoint-scoped; round bids keyset on `(bid_position, evtlog_id)`, with default/max limits 50/200 and no cursor at exhaustion. |
| D3 | Store shape | one `Store` with domain methods / per-domain repo structs | **decided 2026-07-07: per-domain repo structs** — `cosmicgame.Repo` wraps the shared `*store.Store`; `randomwalk.Repo` follows when its files convert. Keeps domain queries in their domain packages and the base package free of game knowledge. |
| D4 | `internal/primitives` future | rename to `internal/model` / dissolve into owners | open |
| D5 | Property-testing lib | stdlib fuzz only / add `pgregory.net/rapid` | **decided 2026-07-06: stdlib-only** — the §4.4 fleet needed no extra dependency; revisit only if a future property needs structured generators |
| D6 | v1 sunset criteria | zero traffic for 30d / hard date | **decided 2026-07-10: hybrid gate** — remove only after known consumers migrate, production metrics show 30 consecutive zero-traffic days (excluding documented probes), and an announced not-before date has passed. |
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
| 2026-07-08 | `25452558` | **Phase 1 completion sprint (§5 done: RandomWalk + base store on pgx, legacy bridge deleted):** the RandomWalk store's 62 legacy methods became context-first, error-returning, pgx-native `Repo` methods (`repo.go` mirrors CosmicGame's; the shared scan loop extracted to `store.QueryList`); the ranking transactions moved from `*sql.Tx` to `pgx.Tx` (`ApplyRankingMatch`/`ConsumeRankingVoteNonce`); the base files (`blockchain.go`, `blockchain_insert.go`, `archive.go`) became 17 ctx-first `Store` methods and 6 dead functions were deleted; `lookups.go` with the process-wide `amap` address cache is gone (per-Store LRU everywhere, new `AddressByID`). `internal/etl` runs on `ETLContext.Store *store.Store` with context-aware helpers. **The legacy bridge is deleted:** `SQLStorage`, `NewSQLStorageFromDB`, `Init_log`/`Log_msg` (replaced by the pgx slog tracer writing to the same db.log files), the transitional `Store.DB()` `database/sql` pool view, `common.Ctx.Db`, and the RandomWalk `SQLStorageWrapper`. Callers: ~49 RandomWalk API handler sites on `c.Request.Context()` + a package `respondStoreError` (not-found flows render the byte-identical legacy `DBError`/error strings via `store.ErrNotFound` mapping — pinned by the parity goldens incl. `errors__missing_rw_token`); ~39 CG API address-lookup sites on `Store.LookupAddressID`; three charity routes no longer `os.Exit` inside a request handler (a client disconnect could previously kill the whole API server once lookups became ctx-aware); rw-etl mirrors cg-etl (dispatch table checks every handler, RLP-decode panic → error, SIGTERM finishes the in-flight batch on `context.WithoutCancel`); notibot/rwctl/opsctl on Repo + Store with checked errors; `opsctl archive node-fill` resolves addresses through a pgx `Store`; `/readyz` pings `Store.Pool()`. Raw DDL dirs `db/{layer1,cosmicgame,randomwalk}` deleted (goose migrations are the only schema source; node-fill's help text updated) and the never-imported sqlc scaffolding retired (D7 amended). **Behavior fixes:** (1) `OfferExists`/`TokenExists` treated any DB failure during the existence check as "does not exist" and silently skipped the event — a data-loss bug; real errors now abort the batch for re-processing. (2) The rw-etl ABI-decode failures no longer kill the process mid-batch. **Tests:** all 492 goldens **byte-identical** (parity, CG+RW store suites, both ETL fixture suites incl. replay-idempotence and reorg; RW store suite rebuilt on `repo(t)` + production one-pool wiring); new rw-etl `TestWriteErrorPropagation` (read-only pool; all 7 event types must surface write errors, the 3 no-write negative fixtures must stay clean); new RW `TestErrorPathsConvertedFiles` (22 cancelled-ctx + 10 closed-pool cases across all three files); `TestStoreBaseHelpers` (AddressByID incl. ErrNotFound, case-insensitive `CountEvtLogsForContract`, `EvtLogExists`); blockops suite ported to the ctx-first API. Statistics benchmarks re-checked: no regression (2.24ms/787µs/259µs vs 2.53/936/315 baselines, B/op identical). Metrics: snake_case 359 → **267** (store layer 0), `os.Exit` in `internal/` 88 → **12 matches / 5 real calls** (3 test mains + startup fatal + `primitives.Fatalf`), Repo methods 304 → **366** (+ 22 base Store methods, ctx coverage 100%), lint uncapped 674 → **492** (capped 145), test files 93 → 95, LOC 68,252 → 66,735, integration coverage 66.6% → **67.2%** (CI floor raised 64% → 66%). |
| 2026-07-09 | `49c07f07` | **ContractState extraction sprint (Phase 2 kickoff: §6.2 ContractState + graceful shutdown; §4.1 happy-path goldens):** the ~70 package-level contract/database state globals in `internal/api/cosmicgame/state.go`, the three unkillable `for { refresh; sleep }` goroutines, the `DisableBackgroundRefresh` test hook and the `contract_live_reads.go` mutex state became one injected component, `internal/api/cosmicgame/contractstate.State`: `New(Config{EthClient, DataSource, Addrs, loggers, intervals})` → `LoadInitial(ctx)` (synchronous startup loads, legacy order) → `Run(ctx)` (ticker loops, stop on cancel) with `Snapshot()` value copies behind one RWMutex (refreshes publish per-group batches — no more torn reads), `SetBidPrice` (dashboard live-read write-back), `FetchLiveSpecialWinners(ctx)` and the V1/V2 mechanics detection cache. Handlers read one snapshot per request; 7 dead globals (`endurance_champ_addr`, `chrono_warrior_*`, `last_bidder_bid_time`, `lastcst_bidder_addr`, `round_activation_time_ts`) and their 3 refresh contract calls per 5s cycle are deleted rather than ported (written, never read). `cosmicgame.Init(ctx, ...)` returns an error — the `cosmic_game_init` startup `os.Exit(1)` is gone (`cmd/apiserver` keeps the fatal + HINT at main). **Graceful shutdown (§6.2):** `signal.NotifyContext` root ctx; public TLS/plain listeners are tracked `http.Server`s with `ReadHeaderTimeout`; SIGTERM flips `/readyz` to `503 draining` (`common.SetDraining`, unit-tested), drains in-flight requests 15s, stops refresh loops via ctx, shuts the metrics/pprof server, closes the pool; `select {}` and the dead autocert manager are deleted (`x/crypto` now indirect). **Parity harness (§4.1 deferred item resolved):** new ABI-driven `testchain.ContractStub` (selector-dispatched `Return`/`Handle` stubs packed via the real ABI codec, concurrency-safe re-stubbing); the apitest eth stub is replaced by `internal/testchain` serving fixture-coherent V1 game + CosmicToken/ERC-20 + MarketingWallet state (round 3 open, alice last bidder, tip block 142); the 10 RPC-dependent goldens regenerated in their happy-path shape (dashboard now pins live prices, percentages, `ContractMechanicsVersion=1`, real `ActivationTime`; `current_special_winners` drops its error redaction — response fully deterministic incl. the chrono-live derivation). Degraded-RPC shapes stay pinned by contractstate unit tests (reachable-node sentinels + dead-node balance failure + NaN game balance). **Tests:** +19 unit test funcs across `contractstate` (happy/sentinels/mechanics-flip/Run lifecycle/`-race` concurrency/special winners) and `common` health probes (draining wins, liveness holds); all 482 other goldens byte-identical; full `-race` integration suite green; fuzz smoke 28/28; statistics benchmarks unchanged (B/op & allocs identical to baselines). Metrics: snake_case 267 → **259**, `os.Exit` in `internal/` 12 matches/5 calls → **11 matches/4 calls**, dot-import files 17 → **15**, api+etl mutable globals ~80 → **~33**, lint uncapped 492 → **474** (new packages 0 issues), test files 95 → 96, integration coverage 67.2% → **69.1%** (CI floor 66% → 68%). |
| 2026-07-09 | `42100ea3` | **Indexer-engine sprint (Phase 3 kickoff: §7 engine core — batch loop, retry policy, metrics, slog, contract-sync tests):** the two near-identical `loop.go` polling loops and the `internal/etl` "common" package became one injected, tested component, [internal/indexer](../internal/indexer)`.Engine` — `New(Config{Store, Client, Progress, Process, Contracts, Logger, Metrics, TopicName, Batch, Retry})` and `Run(ctx) error`. The `Client` interface narrows `*ethclient.Client` to the five calls the engine makes (scriptable fakes in tests); each binary adapts its domain status row through a 20-line `Progress` implementation (preserving `last_evt_id`); the handlers stay behind a `ProcessFunc` until the EventHandler port. Blockops/chainsplit/backfill moved in as Engine methods with their integration suite. **Reliability semantics changed as planned (§7):** any failed batch — RPC, DB or handler — now retries in-process from the last fully completed block with exponential backoff (±25% jitter, 1s→60s cap) instead of crashing per blip; `Run` exits only after 10 consecutive failures (circuit breaker), and SIGTERM mid-batch still finishes the batch on `context.WithoutCancel` before exiting 0. **Two data-loss bugs of the legacy loops found & fixed:** (1) a pipeline failure (EnsureTransaction/InsertEventLog) on a later log of a block whose earlier logs had succeeded advanced the watermark *to that block*, so the failed log was never fetched again — silently lost; the engine only ever acknowledges completed block boundaries (regression test `TestRunMidBlockFailureDoesNotSkipRemainingLogs`). (2) On a fresh status row (`last_block_num=0`) the loops re-resolved the watermark *every iteration* through the store's block watermark, which a failed batch's own inserted blocks advance — the retry resumed past the events the batch still owed; the engine resolves the watermark once at startup and tracks it in memory. **Observability (§7):** Prometheus metrics `rwcg_etl_last_block`, `rwcg_etl_events_total{type}` (labels from the dispatch tables, which now carry event names), `rwcg_etl_batch_duration_seconds`, `rwcg_etl_reorgs_total`, `rwcg_etl_batch_failures_total{stage}`; both ETLs serve `/metrics`+pprof on `METRICS_ADDR` (`indexer.StartMetricsServer`); the engine, startup sync and mains log structured slog records into the legacy two-file layout via `indexer.NewDualLogHandler` (all records → info log, errors duplicated → error log). **contract_sync (§7 item done):** loggers ported to slog; mechanics probe + versioned read fallbacks unit-tested against abigen bindings over `testchain.ContractStub` (no Docker); check-then-correct policy integration-tested (fresh-DB corrections, clean-re-run no-op with untouched address table, targeted correction on a changed chain value, unreadable-params degraded mode). Tools ported: `cgctl backfill-dao-evtlog` runs on `Engine.BackfillContractEvtLogs` (first test coverage: insert + skip idempotence); `opsctl` uses `indexer.FetchLogs`/`client.BlockNumber`; dead code deleted (`inspected_events` registry ~310 LOC, `IMGGEN_PATH`, `rpcclient` globals, 2 orphaned vars). **Tests:** all 492 goldens byte-identical (parity, both store suites, both ETL fixture suites incl. replay-idempotence and reorg — the harnesses now push fixtures through the Engine pipeline); +30 test funcs: batch-policy/backoff/metrics/dual-handler/metrics-server unit tests, loop unit tests (breaker trip + reset, cancellation during caught-up/backoff, fetch-error batch shrink + empty-success growth + watermark acks), loop integration tests (end-to-end batches, transient processor-failure recovery, mid-block regression, shutdown-mid-batch completes the batch, reorg detected by the loop, backfill), contract-sync unit + integration, dispatch-name uniqueness per binary. Full `-race` integration suite green; fuzz smoke 28/28; `internal/indexer` lints clean. Metrics: snake_case 259 → **256**, LOC 69,474 → 71,010, test files 101 → **109**, lint capped 127 → **120** (uncapped 471 → 459), integration coverage 69.5% → **70.3%** (CI floor 68% → 69%; the ≥70% Phase-5 target is met). |
| 2026-07-09 | `d413ac62` | **Stdlib router sprint (§6.2: v1 API off gin onto net/http ServeMux; gin removed from the build):** new dependency-free `internal/api/httpx` package — `Router` over Go 1.22+ `http.ServeMux` (method+pattern routes, route registry replacing gin's `r.Routes()`, global + per-route middleware in standard `func(http.Handler) http.Handler` form, registration-time conflict panics, gin-parity trailing-slash redirect with the query preserved and scheme-relative `//` targets refused, freeze-after-first-request) and `Context` reproducing the wire behavior the goldens pin (`Param`/`Query`/`JSON` (marshal-before-write, panic → Recovery 500 like gin's render)/`String`/`Data`/`File`/`Status`/`ShouldBindJSON` with encoding/json error text incl. `EOF`); status-recording `ResponseWriter` with `Unwrap` for `http.ResponseController`. Middleware ported to standard form in `internal/api/common`: `CORS` (OPTIONS → 204 pre-routing), `Recovery` (broken-pipe silent, re-panics `ErrAbortHandler`, 500 only if unwritten), new slog `AccessLog` (route = matched pattern; replaces `gin.Logger()`), `RateLimit` + `RequireAdminKey` (same envelopes); Prometheus middleware reads `Request.Pattern` (labels now `{param}` syntax, noted in docs/operations.md). **The 160 v1 handlers kept their bodies** — mechanical `gofmt -r` port to `*httpx.Context`/`httpx.H` (a type alias, so map semantics and sorted-key JSON are unchanged); the four `binding:"required"` structs became explicit zero-value checks with identical 400 messages. **Shared router construction:** new `internal/api/routes.New(st, Options)` builds the middleware chain + full route table for both `cmd/apiserver` (Options inject access log, metrics, static assets) and the apitest harness — the "keep in sync with main.go" duplication is deleted. Static assets: `/images/{filepath...}` + `/static/{filepath...}` (files only, no directory listings) with the cache/log subtree middleware, first handler-level tests (200/404/HEAD/Cache-Control/traversal/no-cache env/registration gating). `cmd/loganomaly` parses the new slog access-log format alongside legacy `[GIN]` lines (logfmt parser fuzz-hardened). **gin removed:** zero imports; `go mod tidy` dropped gin, gin-contrib/sse, go-playground/validator, bytedance/sonic, ugorji et al.; `go list -deps` proves gin links into zero packages (one `// indirect` line remains — nhooyr.io/websocket (disgord test dep) lists it; never compiled). Go toolchain bumped 1.26.4 → 1.26.5 (fixes stdlib GO-2026-5856 crypto/tls finding; govulncheck now clean). **Deliberate router-level deltas, all pinned by new tests (`apitest/router_behavior_test.go`):** wrong method → 405 + `Allow` (gin: 404); router-level 404 body gains stdlib's trailing newline; HEAD served by GET routes (gin: 404); OPTIONS/trailing-slash/CORS-on-404/429-envelope re-pinned unchanged. **Tests:** all 196 parity + 12 error-shape goldens byte-identical with zero regenerations; route-drift test now compares OpenAPI `{param}` templates directly against the router registry (one syntax — the `openAPIPathToGin` translator is deleted); +40 unit test funcs (httpx Context/Router/writer, CORS/Recovery/AccessLog, static assets); full `-race` integration suite green; fuzz fleet 28/28 (one `FuzzConnStringEscape` CI-runner timeout reproduced 0/3 — infra flake, not a finding). Benchmarks: rate limiter re-based on the stdlib stack — `distinct_ips` 1,510 → 1,144 ns/op (−24%), `shared_ip` 1,600 → 1,298 (−19%), allocs identical; statistics queries unchanged (B/op byte-identical). Metrics: lint capped 130 → **127** (uncapped 474 → **471**, new packages 0 issues), test files 96 → **101**, LOC 67,875 → 69,474, integration coverage 69.1% → **69.5%** (floor stays 68%). |
| 2026-07-09 | `f54e4cfe` | **EventHandler-port sprint (Phase 3 complete: §7 handler port + the two blocked §4.4 fuzz targets):** the ~83 `proc_*` functions and their package-global wiring became typed, dependency-injected handler sets. New in [internal/indexer](../internal/indexer): `EventHandler` (`Topic`/`Name`/`Sources`/`Decode`/`Store`), the generic `NewHandler[E]` adapter (concrete event types survive the decode→store handoff), `Registry` (multi-handler-per-topic, source-address filtering before decode — the in-handler `bytes.Equal` guards became declarative registrations; construction validates that same-topic handlers share one metric label) and `LogProcessor` (the one copy of the twin `process_single_event`s, over a narrow `EventLogSource` seam; `FuzzEvtlogRLP` moved with it). New packages: [internal/indexer/cosmicgame](../internal/indexer/cosmicgame) — 76 handlers (78 registrations: the single legacy Transfer row split into ERC721/ERC20 per source) as `decode*`/`store*` method pairs on `Handlers` (`Config{Repo, Store, Caller (bind.ContractCaller), Contracts, Logger}`), ABIs parsed once in `New`, DB/RPC enrichment (CST-reward mint lookup, prize-round resolution, donation-info + tokenURI reads) in the store steps so every decode is pure; `BootstrapContracts` deduplicates the main()/harness address bootstrap; `contract_sync.go` moved in as `SyncContractParams`; [internal/indexer/randomwalk](../internal/indexer/randomwalk) — the 7 RW handlers, existence-guard skips in the store steps. **Both mains are pure wiring with zero package-level variables** (the ~30 ETL globals — 11 ABIs, 11 addresses, repos, eclient, `Info`/`Error`, ~67 `evt_*` topic slices — are deleted); handler logging is one structured slog record per event (dual-file layout preserved in prod). **Robustness:** decode steps bounds-check indexed-topic counts — a malformed log matching a known topic0 now fails the batch instead of panicking the process (pinned by the new fuzz targets). **Tests:** all 492 goldens byte-identical, zero regenerated (both fixture suites + story/reorg/replay/write-error suites moved into the handler packages as pure `git mv` renames; harnesses build `Handlers` + `Registry` per reset instead of mutating globals; the write-error suites re-process through a second read-only-pool handler set); +13 unit test funcs for the registry/LogProcessor (dispatch, source filtering, error propagation, RLP reconstruction, foreign-event-type rejection) plus registry-shape tests per package (metric-name consistency now enforced at construction, superseding `dispatch_names_test.go`); §4.4 unblocked: `FuzzEventDecodeCG`/`FuzzEventDecodeRW` iterate every registered handler's `Decode` (fleet 28 → 30, all green in the smoke run; one deadline flake on FuzzEventDecodeRW reproduced 0/3 — infra, not a finding). `BenchmarkEventDecode` moved with the bid handler and re-run: 2,105 ns/op / 2,920 B/op / 43 allocs — B/op and allocs byte-identical to the `docs/benchmarks.md` baseline. Full `-race` integration suite green. Docs: architecture/BACKEND/benchmarks/README updated to the new layout. Metrics: snake_case 256 → **161** (both ETLs 0), dot-import files 15 → **4**, api+etl mutable globals ~30 → **~12** (ETLs 0), lint capped 120 → **119** (uncapped 459 → **428**; all three indexer packages 0 issues), test files 109 → **112**, fuzz targets 28 → **30**, LOC 71,010 → 71,159, integration coverage 70.3% → **72.8%** (CI floor 69% → 71%). |
| 2026-07-10 | — | **API-v2 round-bids sprint (§6.1 complete + first §6.2 vertical slice):** accepted [ADR-0005](adr/0005-api-v2.md), deciding D2 (bounded/versioned opaque cursor + limit) and D6 (consumer migration + 30 zero-traffic days + announced not-before date); added the OpenAPI 3.0.3 [v2 contract](openapi-v2.yaml), pinned oapi-codegen v2.7.2 as a Go tool, committed generated strict stdlib interfaces/models and added a CI generation-drift gate. New zero-global `internal/api/v2.Server` injects the shared store/repo, existing `contractstate.State` and slog logger; `httpx.Router.HandleFunc` lets generated routes retain global middleware, conflict checks, metrics patterns and route enumeration. Shipped `GET /api/v2/cosmicgame/rounds/{round}/bids` + `/{position}` with camelCase typed models, exact decimal-string wei amounts, UTC timestamps, RFC 9457 errors and no internal-detail leakage. Pagination uses a strict/fuzzed base64url cursor over `(round,bid_position,evtlog_id)` and a `LIMIT n+1` keyset query; migration 00009 adds the matching index concurrently. **Tests:** 5 new test files (117 total), table-driven handler/model/cursor/router tests, store integration page-boundary/cancellation tests, bid-cursor fuzz target (31/31 smoke green), exact v2 spec↔router drift, and 10 deterministic real-Postgres v2 goldens whose statuses/headers/bodies are all kin-openapi validated. Full race+shuffle unit and race integration suites green; all existing 196 v1 parity + 12 error goldens remained unchanged; govulncheck clean; new/touched v2+httpx packages lint-clean (repository baseline remains 119). Integration coverage 72.8% → **73.0%**, CI floor 71% → **72%**; golden files 590 → **600**. |
| 2026-07-10 | — | **API-v2 completed-rounds sprint (second §6.2 vertical slice):** expanded [openapi-v2.yaml](openapi-v2.yaml) from 2 → 4 generated operations with cursor-paginated `GET /api/v2/cosmicgame/rounds` and lean `GET /api/v2/cosmicgame/rounds/{round}` resources. The contract exposes claim identity, exact main-prize wei amounts, aggregate/timing data, charity/staking allocations and special-prize summaries; legacy mega-response collections (`allPrizes`, raffle/staking winners and raffle deposits) are deliberately deferred to nested resources. `v2.Server` gained a narrow injected `roundReader`; strict mapping canonicalizes addresses/timestamps/decimal amounts, omits sentinels and rejects malformed repository data without leaking internals. Store work extracted the lean `RoundInfo` base while `PrizeInfo` still composes all four v1 collections byte-identically; `PrizeClaims` and new descending `(round_num, evtlog_id)` `PrizeClaimsPage` share one scanner, and concurrent migration 00010 adds the matching index. **Tests:** 4 new test files (121 total), unit coverage for mappings/handler failures/cursor order, real-Postgres page/lean-detail/cancellation tests, `FuzzDecodeRoundCursor` (32/32 smoke green), exact 4-route spec drift, and 11 new deterministic kin-openapi-validated round goldens (21 v2, 611 total). Full race+shuffle unit and race integration suites green; all 196 v1 parity + 12 error goldens and existing store goldens remained unchanged; go vet and govulncheck clean; v2 lint-clean (repository baseline remains 119). Integration coverage **73.0% → 73.4%**, CI floor **72% → 73%**. |
| 2026-07-10 | — | **API-v2 current-round sprint (third §6.2 vertical slice):** expanded [openapi-v2.yaml](openapi-v2.yaml) v0.2.0 → v0.3.0 and the generated strict router from 4 → 5 operations with `GET /api/v2/cosmicgame/rounds/current`. The resource reads exactly one injected `contractstate.Snapshot`, combines it with `CosmicGameRoundStatistics` and the authoritative `BidCountForRound`, and exposes only typed open-round identity/timing/aggregates plus exact decimal-string wei and microsecond values — no floating-point ETH, global dashboard payload or request-time RPC fallback. The legacy misnamed duration is corrected to `secondsUntilMainPrize`; uninitialized/failed cache sentinels (including a zero last bidder after bids exist) return an RFC 9457 503 with OpenAPI-required `Retry-After: 5`, while malformed data/store failures are opaque 500s. The mapper also normalizes the PostgreSQL timestamptz text retained by the frozen v1 repository into UTC RFC3339, fixing the compatibility seam found by the real fixture. **Tests:** 2 new test files (123 total) cover every live sentinel, malformed amounts/timestamps/counts/identities, zero-bid omission, single-snapshot/repository calls, error secrecy, both cancellation stages and deterministic HTTP output; store integration now pins open-round count + cancellation; 3 twice-fetched, kin-openapi-validated current-round goldens cover 200/503/500 (24 v2, 614 total). Full race+shuffle unit and race integration suites green; all 196 v1 parity + 12 error goldens unchanged; go vet and govulncheck clean; v2 + apitest lint-clean (repository baseline remains 119). The 32-target fuzz fleet had one unrelated `FuzzEloUpdate` stop-deadline flake with no crasher, reproduced 0/3 in isolation. Integration coverage **73.4% → 73.6%**, CI floor **73% → 73.5%**. |
| 2026-07-10 | — | **API-v2 round-prizes sprint (first round sub-resource slice):** expanded [openapi-v2.yaml](openapi-v2.yaml) v0.3.0 → v0.4.0 and the generated strict router from 5 → 6 operations with cursor-paginated `GET /api/v2/cosmicgame/rounds/{round}/prizes`. All 16 `cg_prize.ptype` values now have stable string enums; each response exposes only the applicable exact `ethAmountWei`, `cstAmountWei` or `nftTokenId`, canonical transaction/address/time fields, and deliberately omits legacy floats, address IDs, hardcoded claim flags and empty token metadata. The resource is completed-round-only (open/missing 404), validates repository identity/order before mapping, and pages on the unique `(ptype, winner_index)` suffix with migration 00011's matching concurrent `(round_num, ptype, winner_index)` index. Store work extracted the existing full-list SQL/scanner byte-identically for v1, added `AllPrizesForRoundPage` + `CompletedRoundExists`, and retained every old golden. **Platform quality:** bid, completed-round and prize cursors now share one bounded strict-JSON/base64url codec; all list handlers share the 50/200 limit resolver with existing payloads/errors unchanged. **Tests:** 5 new test files (128 total) exhaustively cover all 16 mappings, asset exclusivity, malformed rows, handler/error/order invariants, generic codec behavior, limit boundaries and `FuzzDecodePrizeCursor` (fleet 33); store integration proves full-list/page equivalence, page exhaustion, completion gates and cancellation. Twelve new twice-fetched, kin-openapi-validated prize goldens follow real continuation cursors across every prize type and cover empty/400/404/500 paths (36 v2, 626 total). Full race+shuffle unit and race integration suites green; all 196 v1 parity + 12 error goldens unchanged; build, vet and govulncheck clean; v2 + apitest lint-clean and repository baseline remains 119. The fuzz fleet had one unrelated `FuzzMetadataHostDispatch` stop-deadline flake with no crasher, reproduced 0/3 in isolation. Integration coverage **73.6% → 73.8%**, CI floor **73.5% → 73.7%**. |
| 2026-07-10 | — | **API-v2 round-raffles sprint (second round sub-resource slice):** expanded [openapi-v2.yaml](openapi-v2.yaml) v0.4.0 → v0.5.0 and the generated strict router from 6 → 8 operations with cursor-paginated `GET /api/v2/cosmicgame/rounds/{round}/raffle-eth-deposits` and pool-selected `.../raffle-nft-winners?pool=bidder\|randomWalkStaker`. The first resource exposes bidder-raffle PrizesWallet credits with exact `ethAmountWei` and claimed state; the second exposes each NFT+CST payout as one typed row with exact `cstAmountWei`. Both are completed-round-only, reject cross-round/cross-pool cursors and validate repository scope/order before mapping; legacy floats, row/address IDs and redundant pool flags stay out of v2. Store work added dedicated exact-wei ETH projections plus ascending `(winner_index, evtlog_id)` and pool-scoped `(winner_idx, evtlog_id)` page methods while keeping v1 full-list SQL/order byte-identical; concurrent migration 00012 adds matching `(round_num, winner_index, evtlog_id)` and `(round_num, is_staker, winner_idx, evtlog_id)` indexes. The shared `roundNotFoundProblem` now keeps all completed-round sub-resources wire-consistent. **Tests:** 4 new test files (132 total) cover claimed/unclaimed mappings, both NFT pools, strict scope/version/key bounds, handler paging/error/order invariants and two new fuzz decoders (fleet 35); store integration proves exact page boundaries, semantic equivalence with frozen full lists, exhaustion and cancellation. Twenty-four new twice-fetched, kin-openapi-validated raffle goldens cover both pools plus empty/400/404/500 matrices (60 v2, 650 total). Full race+shuffle unit and race integration suites green; all 196 v1 parity + 12 error goldens unchanged; build, vet and govulncheck clean; v2 + apitest lint-clean and repository baseline remains 119. The fuzz fleet had one unrelated `FuzzReceiptsDecode` stop-deadline flake with no crasher, reproduced 0/3 in isolation. Integration coverage **73.8% → 74.1%**, CI floor **73.7% → 74%**. |
| 2026-07-10 | — | **API-v2 round-donations sprint (round sub-resources complete):** expanded [openapi-v2.yaml](openapi-v2.yaml) v0.5.0 → v0.6.0 and the generated strict router from 8 → 11 operations with cursor-paginated `GET /api/v2/cosmicgame/rounds/{round}/{eth,erc20,nft}-donations`. Direct ETH records use a `plain`/`withInfo` discriminator with exact `ethAmountWei` and contract data; arbitrary token quantities are exact `amountBaseUnits` rather than incorrectly assuming 18-decimal wei; NFT records expose their public PrizesWallet index, token identity and recorded URI. Donations remain queryable during open rounds (empty 200 when no events), unlike completed-only prizes/raffles. All three newest-first cursors are round/resource-scoped, document their stable-boundary/live-poll semantics, and validate repository scope/order before mapping. Store work added v2-only exact event projections and strict `LIMIT n+1` page methods while preserving every v1 query; the combined ETH `UNION ALL` bounds both indexed branches before the merge. Concurrent migration 00013 adds four matching `(round_num, evtlog_id DESC)` indexes. **Tests:** 4 new test files (136 total) cover all ETH variants, exact large amounts, corrupt mappings, every handler/page invariant, open/empty rounds and three new cursor fuzzers (38/38 ten-second smoke green); store integration proves legacy/page equivalence, cross-table ETH ordering, exhaustion and cancellation. Twenty-eight new twice-fetched, kin-openapi-validated donation goldens cover success/next/empty/open-round and symmetric malformed/cross-round/cross-resource/limit/internal failures (88 v2, 678 total). Full race+shuffle unit and race integration suites, build and vet are green; generated output is reproducible; all 196 v1 parity + 12 error goldens remain unchanged; govulncheck reports zero reachable vulnerabilities. New-diff lint is zero and both `internal/api/v2` and `internal/store/cosmicgame` are clean (repository baseline 119 capped / 405 uncapped). Integration coverage **74.1% → 74.4%**, CI floor **74% → 74.3%**. |
| 2026-07-10 | — | **API-v2 statistics + claims sprint (first dashboard slice):** expanded [openapi-v2.yaml](openapi-v2.yaml) v0.6.0 → v0.7.0 and the generated strict router from 11 → 16 operations. DB-only `/api/v2/cosmicgame/statistics` and `/statistics/counters` expose exact global aggregates without request-time RPC or the v1 dashboard mega-response. Six-mode `/statistics/leaderboard/roi` replaces offsets with sort+`minBids`-scoped keysets, exact signed `netProfitWei`, decimal-string ROI/win-rate ratios and an internal bidder tie-breaker. Newest-round `/statistics/claims` summaries omit v1's unbounded inline assets; completed-only `/rounds/{round}/claims` returns independently bounded transaction, attached-token and unclaimed-item pages with exact ETH wei/ERC-20 base units. V2-only store projections preserve every v1 query; migration 00014 adds three `(round_num, evtlog_id)` claim-event indexes. **Correctness caught during implementation:** qualifying numeric ROI sort expressions was required because PostgreSQL otherwise resolved the text-cast output aliases and sorted lexically, which skipped rows across keyset pages; all six sorts now have full-list/page equivalence tests. **Tests:** 4 new test files (140 total), three cursor fuzzers (fleet 41), six exact store goldens and 39 twice-fetched, kin-openapi-validated HTTP goldens covering every sort, page boundary, section cursor and 400/404/500 path (127 v2, 723 total). Full race+shuffle unit and race integration suites, build, vet, generation reproducibility and govulncheck are green; all 196 v1 parity + 12 error goldens remain unchanged; new-diff lint is zero (repository baseline 119/405). Statistics benchmarks remain in the recorded noise envelope (global 2.61–2.78ms, claims 0.96–1.21ms, ROI 327–351µs; allocations unchanged). The 41-target smoke run had one unrelated `FuzzThousandsFormat` stop-deadline flake with no crasher, reproduced 0/3 in isolation; all three new fuzzers passed. Integration coverage **74.4% → 74.8%**, CI floor **74.3% → 74.7%**. |
| 2026-07-10 | — | **API-v2 participant-directories sprint (second dashboard slice):** expanded [openapi-v2.yaml](openapi-v2.yaml) v0.7.0 → v0.8.0 and the generated strict router from 16 → 22 operations with cursor-paginated bidder, winner, ETH-donor, CST-staker, RandomWalk-staker and dual-staker resources under `/api/v2/cosmicgame/statistics/participants/*`. V2-only projections expose canonical addresses, exact wei strings and deterministic descending aggregate keysets with an internal address-ID tie-breaker; legacy floats, IDs, duplicate winner shapes and zero-count bidder tombstones stay out of the contract. Winner counts and ETH totals are rebuilt from canonical prize/event rows rather than trusting the replay-sensitive `cg_winner` aggregate. Endpoint-scoped strict cursors reject cross-directory reuse and out-of-range count keys; ranked directories explicitly document their weak consistency under live aggregate changes, and handlers validate page cardinality/order/scope before mapping. Migration 00015 adds four concurrent aggregate read indexes. Five queries benchmark at 171–194µs; the canonical winner reconstruction is 465µs, still below the existing ROI/claims query class. **Tests:** 5 new test files (145 total), `FuzzDecodeParticipantCursor` (42/42 targets passed ten-second smoke), six deterministic store goldens and 32 twice-fetched, request/response OpenAPI-validated HTTP goldens covering first/next/empty pages, tie boundaries, malformed/cross-directory cursors, invalid limits and opaque 500s (159 v2 goldens total). Store integration proves full-list/page semantic equivalence for all six resources, independence from a corrupted winner aggregate, zero-count bidder exclusion, terminal exhaustion, cancellation/closed-pool behavior and a synthetic dual-staker tie. Full race+shuffle unit and race integration suites, build, vet, generated-code reproducibility and govulncheck are green; all 196 v1 parity goldens remained unchanged; touched packages are lint-clean. Integration coverage **74.8% → 75.1%**, CI floor **74.7% → 75.0%**. |
| 2026-07-10 | — | **API-v2 bidding-analytics sprint (third dashboard slice):** expanded [openapi-v2.yaml](openapi-v2.yaml) v0.8.0 → v0.9.0 and the generated strict router from 22 → 27 operations with DB-only `/api/v2/cosmicgame/statistics/bidding/{activity,frequency,type-ratio,top-active-periods,time-bounds}`. Windowed resources require `from`/`to`, cap scans at five years and time series at 2,000 buckets, preserve UTC/anchored bucket and first-hour exclusion semantics, and omit the legacy exact-boundary terminal bucket. Bid-type percentages are deterministic decimal strings derived from integer counts; top-period responses hide address IDs, use a v2-only stable tie-breaker and reject results above 2,000 periods. Every analytics query has a five-second deadline, and an injected clock makes the optional 30-day recent-spike marker deterministic. V2-only bounded store projections preserve frozen v1 query behavior: half-open timestamp filters run through migration 00016's concurrent `cg_bid(time_stamp)` index and aggregate each bid once with `DATE_BIN` before joining the zero-fill series. Six-run medians are 167–196µs for the single queries and 400µs for the bounded two-query top-period path. **Tests:** 2 new test files (147 total), `FuzzResolveAnalyticsWindow` + `FuzzDetectBidSpikes` (fleet 44), expanded cancellation, post-2038, partial-tail, timestamp-index, SQL period-cap and stable-tie coverage, plus 21 twice-fetched, request/response OpenAPI-validated HTTP goldens covering nonzero frequency, spike/recent-spike output, exact-boundary trimming, defaults, bind/window/timestamp/limit and opaque 500 paths (180 v2 total). Full race+shuffle unit and race integration suites, build, vet, generation reproducibility and govulncheck are green; all 196 v1 parity goldens remain unchanged; touched packages report zero lint issues (repository baseline 119). The one-second full fuzz smoke had a pre-existing claim-cursor stop-deadline timeout with no crasher, reproduced cleanly in isolation; both new targets passed dedicated five-second runs. Integration coverage **75.1% → 75.2%**, CI floor **75.0% → 75.1%**. |
| 2026-07-10 | — | **API-v2 contract-configuration sprint (dashboard decomposition complete):** expanded [openapi-v2.yaml](openapi-v2.yaml) v0.9.0 → v0.10.0 and the strict router from 27 → 32 operations with DB-backed `/contracts/addresses` plus cache-only `/contracts/configuration`, `/contracts/balances`, `/rounds/current/bid-prices` and `/rounds/current/special-winners`. The refresh engine now pins related RPC/balance reads to one block, serializes mechanics/address-dependent groups, bounds RPC/DB calls with deadlines, tags constant/variable mechanics generations and charity-balance generations, and exposes resource-specific readiness with accurate 5s/30s/300s retry guidance. V1 fixed and V2 dynamic CST rewards are modeled separately; V2 auction start timestamps normalize to clamped elapsed progress; special-winner round/contract reads and optional CST event lookup share one source block. The legacy dashboard write-back is isolated from v2's block-pinned price cache. [api-v2-migration.md](api-v2-migration.md) publishes the dashboard replacement map, so the §6.2.1 statistics/dashboard slice is complete without a v2 mega-response. **Tests:** 2 new test files (149 total), two cached-state fuzzers (fleet 46), 17 twice-fetched OpenAPI-validated HTTP goldens (197 v2 total), V1/V2 mechanics and bid-price goldens, symmetric cache 503/recovery cases, invalid-registry 500, complete address and v1/v2 semantic comparisons, block/generation/address coherence, overflow/auction normalization, optional DB failure, timeout and race coverage. Full build/vet/race unit gates are green; the first integration attempt hit an unrelated testcontainer connection reset and the immediate full race rerun passed. All 196 v1 parity goldens remain unchanged; touched packages are lint-clean and govulncheck reports no reachable vulnerability. Integration coverage **75.2% → 75.3%**, CI floor **75.1% → 75.2%**. |
| 2026-07-11 | — | **API-v2 user-foundation sprint (first §6.2.1 user slice):** expanded [openapi-v2.yaml](openapi-v2.yaml) v0.10.0 → v0.11.0 and the strict router from 32 → 34 operations with exact `GET /api/v2/cosmicgame/users/{address}` and cursor-paginated `.../users/{address}/bids`. The profile is deliberately collection-free: one checksummed identity with nested bidding, canonical prize/raffle, direct ETH-donation, transfer, CST-staking and RandomWalk-staking statistics; internal IDs, float ETH, magic sentinels, request-time RPC and the v1 mega-response's unbounded arrays stay out. Bid totals are rebuilt from canonical rows and prize totals share the participant directory's canonical event reconstruction, so corrupting `cg_winner` cannot alter the response. A valid unindexed wallet gets the same stable zero `200` shape and empty bid page. User bids reuse the existing typed `Bid`, order newest-first by immutable `evtlog_id`, reject cross-wallet cursors and use migration 00017's `(bidder_aid, evtlog_id DESC)` index. **Tests:** 4 new test files (153 total), `FuzzDecodeUserBidCursor` (fleet 47), one deterministic store golden and 16 twice-fetched kin-openapi-validated HTTP goldens (213 v2 total) cover active/indexed-zero/unindexed profiles, exact large amounts, every nested mapping, first/next/exhausted/empty pages, invalid addresses/limits/cursors, cross-wallet reuse, malformed repository rows, error secrecy and cancellation. Store integration proves canonical-aggregate independence, full-list/page equivalence with no gaps or duplicates, cancellation/closed-pool behavior and index presence. Six-run medians are 499µs for the canonical profile and 280µs for a 50-bid page. Full race+shuffle unit and race integration suites are green; all 47 fuzz targets pass the smoke fleet and the new cursor passes a dedicated 10-second run; generated output is reproducible; govulncheck is clean; all 196 v1 parity goldens remain unchanged; touched packages are lint-clean. Metrics: LOC **91,300 → 93,054**, integration coverage **75.3% → 75.5%**, CI floor **75.2% → 75.4%**. |
| 2026-07-11 | — | **Coverage-quality sprint 1 ([ADR-0006](adr/0006-coverage-policy.md), first 75.5→90 stage):** replaced the misleading single inline percentage with a tested policy engine, `internal/covergate` + `cmd/covergate`, which deduplicates repeated Go profile blocks, excludes generated/test-only code from canonical metrics, reports every package and intersects staged/PR diffs with executable blocks. `scripts/coverage-gate.sh` is shared by Make and CI; successful local profiles cache by staged-source hash. The tracked native/Cursor hooks are installed now but `commitGateEnabled=false` deliberately allows commits during the climb; once handwritten internal coverage reaches 90%, the checklist requires activating a permanent fail-closed 90% commit floor. CI meanwhile exposes one branch-protectable **Coverage Gate** and uploads its profile/diff/report. Coverage work targeted behavior rather than line filling: exhaustive v2 claims/ROI invariants; 100+ real-router cancellation/error-secrecy paths across CosmicGame and RandomWalk; every admin formatter and contract-backed resolver branch; canonical store resolver history/cancellation; JSONL/CSV output concurrency/error/append semantics; tool backup RLP/config/path/Postgres/contract lookups. **Two reliability bugs found and fixed:** (1) chain-sync synthetic log indexes used a 10,000-value time-based space; collisions triggered `InsertEventLog`'s delete-before-insert path and silently deleted an older correction. `Store.NextEventLogIndex` now allocates above the block's maximum, with the three-run correction test as regression coverage. (2) five RandomWalk ranking/user read paths leaked raw database/context errors or misclassified cancellations as client lookup failures; they now use the common opaque 500 envelope while retaining legacy not-found shapes. Full build/vet/race+shuffle unit and race integration gates pass; govulncheck is clean and new-diff lint is zero. The 47-target one-second fuzz smoke had one pre-existing `FuzzEventDecodeCG` stop-deadline timeout with no crasher, reproduced cleanly in a dedicated ten-second run. **Metrics (race-enabled CI profile):** 14 new test files (167 total), LOC **93,054 → 96,367**, legacy internal **75.5% → 79.07%** (floor 78.8%), canonical handwritten internal **80.43%** (floor 80.1%, target 90%), newly truthful all-production **52.80%** (floor 52.5%), and changed executable Go **96.16%** (floor 95%). |
| 2026-07-12 | — | **Coverage-quality sprint 2 (80→85 API-boundary stage):** completed a behavior-first HTTP boundary matrix across the real v1 router: every guarded CosmicGame GET fails cleanly before initialization; every numeric path position rejects malformed input; wallet-scoped routes handle malformed/unindexed addresses without 5xx or internal leakage; missing rows retain their public legacy shapes; and table-fault tests prove dashboard, user mega-response and 26 user-scoped resources fail opaquely after earlier dependencies succeed. Shared HTTP tests now pin readiness ping/draining precedence, parser limits, rate-limit extremes, recovery/disconnect semantics, access-log bytes and public NFT URL derivation. RandomWalk ranking uses typed client-error classification plus injected transaction seams, with exhaustive signed-vote/nonce/duplicate/rollback tests; v2 list handlers uniformly reject repository over-cardinality before emitting data or cursors; contract-state tests cover stale-mechanics fallback, overflow normalization, readiness coherence and refresh-failure isolation. **Reliability hardening:** (1) comma-separated `X-Forwarded-Proto` now honors the first proxy hop instead of producing malformed asset URLs; (2) signed-ranking DB errors can no longer be misclassified as client errors by substring matching; (3) eight v2 page builders now fail closed on repository cardinality violations. Sixty-eight path-parameter checks made impossible by `ServeMux` route matching were deleted rather than tested artificially. All 196 v1 parity + 12 established error goldens remain unchanged; 175 test files and all 47 fuzz targets are retained. Full build/vet/race integration passed; changed-code coverage is 97.52%. **Metrics:** LOC **96,367 → 99,920**, legacy internal **79.07% → 83.73%** (floor 83.5%), canonical handwritten internal **80.43% → 86.07%** (floor 85.8%, target 90%), all production **52.80% → 56.26%** (floor 56.0%). |
| 2026-07-12 | `4142524c` | **Command-seam sprint slice 1 (§4.6: notification stack + small binaries; §8.1 notibot, §8.2 notibot dot-import):** the three worst zero-coverage binaries became thin wiring over injected, heavily tested packages. **`internal/notify/rwbot`** unifies the two divergent RandomWalk notification bots (874-line `cmd/notibot/main.go` and `rwctl notify-bot`, ~60% duplicated) into one `Engine` with seams for every external system — `DataSource` (satisfied by `*randomwalk.Repo`), `Tweeter`, `Discord`, media `Fetcher`, `ResampleFunc` (ffmpeg) and `WithdrawalReader` — plus pure, pinned formatters for all five notification texts and the four Discord statistics-channel names. Both binaries now share the persisted `rw_messaging_status` watermark, context cancellation replaces the in-loop `os.Exit`s, and the engine is tested against scripted fakes (watermark persistence, 404-wait/403-skip media policy, retry-vs-skip semantics, floor-price dedup, mid-batch cancellation) plus a testdb integration run proving fixture events announce exactly once and restarts stay silent. **Deliberate fixes over legacy:** the 403 "skip" that actually stalled the bot forever now skips (watermark advances); the `last_mint_ts` data race is an atomic owned by the engine; ffmpeg failures no longer `os.Exit(1)` mid-loop; rwctl's bot no longer re-announces the floor price on every restart (it seeded `cur_floor_price=0`) and no longer re-notifies history after restarts (in-memory timestamp watermark → persisted watermark); a failed media fetch backs off for a poll interval instead of hammering the DB/image server; Discord embeds get a clean detail URL (the legacy embed URL carried a leading `\n\n`); `rwctl notify-bot` reads `RPC_URL` like every other subcommand (was `AUGUR_ETH_NODE_RPC_URL`); dead `DEV_MODE` block deleted. **`internal/ethtx`** extracts rwctl's transaction plumbing (connect, account prep, EIP-155 signing, legacy 2.0x gas policy, receipt wait, quiet/verbose output) behind an explicit `Options{RPCURL, PrivateKeyHex, Verbose, Out, ReceiptTimeout}` — no env reads or stdout writes inside the package; `cmd/cgctl/internal/ethtx` merges in with slice 2. **`internal/testchain` gained transaction support** (`eth_gasPrice`, `eth_getTransactionCount`, `eth_sendRawTransaction` mining into a fresh block, settable balances/nonces, one-shot revert/pending/reject knobs), so every rwctl transaction command now executes end-to-end in tests through cobra args → env config → abigen bindings → EIP-155 signing → receipt handling, including revert, receipt-timeout, insufficient-balance and verbose-output paths. **`internal/freezer/scan`** extracts the freezer-scan pipeline (chunk-parallel scan, resume-from-JSONL, chunk merge, error log, `--info`) as `Run(ctx, reader, Options) (Stats, error)` with synthetic cidx/cdat fixtures pinning filtering, resume/append, no-resume, best-effort vs fail-fast on corrupt index/payload, graceful cancellation (completed chunks still merge), directory outputs and progress logging; `cmd/freezer-scan` is a flag wrapper with its own tests, and the dead `--validate`/`--validateOnly` flags (advertised but never implemented) are gone. rwctl's scan/verify loops share one tested `scanLogsByRange`; the twitter-auth PIN flow runs against stub OAuth endpoints end to end; notibot's Discord sink is tested against a stub Discord REST API through the real disgord client (rate-limit retry parsing included, `ParseRetryAfterSeconds` unit-tested), and the ffmpeg adapter converts a real generated clip in tests. `.golangci.yml` now excludes best-effort `fmt.Fprint*` human-output writes from errcheck (data-file writers keep explicit bufio/Close error checks). **Tests:** 18 new test files (187 → 205), all 47 fuzz targets green on a full 5s re-run (one first-run stop-deadline infra flake, no crasher), full race+shuffle unit and race integration suites, build/vet/govulncheck clean; touched packages lint zero issues; repository lint 114 capped/393 uncapped → **85/346**. **Metrics:** LOC **102,366 → 106,579**, snake_case **159 → 135** (notibot 0), dot-import files **4 → 3**, `os.Exit` in `internal/` unchanged (bot exits now live in the binaries), package-level mutable state: notibot's ~30 globals deleted. Coverage (race profile): legacy internal **88.48% → 88.85%** (floor 88.2% → 88.6%), handwritten internal **91.82% → 92.08%** (floor 91.5% → 91.8%), all production **59.65% → 66.31%** (floor 59.4% → 66.0%), changed code 95.16%. |
| 2026-07-12 | `081fbea2` | **Command-seam sprint slice 2a (§4.6: `cmd/cgctl`; slice 2 split — cgctl signs real transactions and warranted its own sprint, opsctl becomes slice 2b):** the largest remaining zero-coverage binary (~4,900 LOC incl. the 1,221-line autobid bot) became thin, fully tested wiring. **ethtx merge:** `cmd/cgctl/internal/ethtx` (505 LOC, env-driven, stdout-only `Printer`) is deleted; [internal/ethtx](../internal/ethtx) gained a session-scoped `GasPriceMultiplier` option (`AdjustGasPriceBy`, `AdjustedGasPrice`), the CosmicGame gas-limit constants, `CallOpts`, `Network.Balance`, `Network.AdvanceDevChainTime` (Hardhat `evm_increaseTime`/`evm_mine`, gated on `IsDevChain`), `Session.Refresh` for multi-transaction commands, and the ported format/output helpers (`FormatTokenAmount`, `FmtDuration`, `ConvertToPercentage`, `MaxUint256`, `WeiToEthCompact`, `KeyValueDuration`, `ContractInfo`) — rwctl behavior unchanged, its suite untouched. **`internal/autobid`:** the bot's decision rules are a pure, exported `Decide(market, limits, myAddr)` core pinned by an exhaustive table (first-bid-must-be-ETH, CST-anyway ordering, RWalk mint+half-price economics, claim timing, boundary comparisons); the `Engine` owns the refresh loop, pending-transaction tracking with bounded receipt retries, single-flight background RWalk-token search (**the legacy goroutine raced on plain int64 fields — now atomics**), reconnect-with-chain-id-check, blockchain-reset abort (`ErrChainReset`) and session stats, with injected `Dial`/`Sleep`/`Out` seams and ctx cancellation replacing the in-loop signal channel. Scripted-round engine tests drive real abigen bindings over `testchain` + `ContractStub` (bid→win→claim, cheap-CST, RWalk pre-owned and mint-receipt paths, timeout-claim, initial bidding incl. safety cap and mid-loop failures, reconnect storms, chain-id mismatch, reset abort, cancellation, every submission/refresh/receipt error branch). **Deliberate fixes over legacy, all regression-tested:** (1) every transaction subcommand now waits for its receipt and fails on an on-chain revert — the legacy scripts printed `Success` at submission, so reverts were invisible; `claim-prize --delay` fired its second transaction after a blind 2s sleep and now waits properly; (2) the autobid bot honors `GAS_PRICE_MULTIPLIER` (it hardcoded 2.0×); (3) malformed numeric env config (`MAX_ETH_BID=lots`, bad `GAS_PRICE_MULTIPLIER`, ...) is a startup error instead of a silent fallback to defaults — a typo could previously bid with 50× the intended limit; (4) a pending transaction is resolved before the round-change exit, so a round-ending claim is still confirmed and counted. The `claim-and-set-time-increment` planner moved onto the shared session (all four documented paths + defer-exhaustion + Hardhat time-advance tested); `deploy-erc20`, `info` (V1+V2 variants), `erc20`/`nft` groups, `owner`, `donation-records`, `total-tokens`, `token-seed` and `backfill-dao-evtlog` are all covered (the DB commands against seeded testcontainers Postgres incl. backfill insert + idempotent re-run). **testchain grew** `evm_increaseTime`/`evm_mine` (offset shifts new block timestamps), `SetChainID`, receipt `contractAddress` for deploys, `SubmittedTxCount`, `SetMinedTxLogs` and the `FailNextRPC`/`FailRPCAfter` per-method failure injectors. **Tests:** 13 new test files (205 → 218), all 47 fuzz targets green on a 10s smoke, full `-race -shuffle` unit and `-race` integration suites, build/vet clean, touched packages lint-clean (repository 0 new issues). **Metrics (authoritative race profile):** LOC **106,579 → 113,213**, legacy internal **88.85% → 89.19%** (floor 88.9%), handwritten internal **92.08% → 92.34%** (floor 92.1%), all production **66.31% → 76.31%** (floor 66.0% → 76.0%), changed executable Go **98.37%**; package coverage `internal/autobid` 98.10%, `internal/ethtx` 99.16%, `cmd/cgctl` 92.10%. |
| 2026-07-12 | `8c7052a3` | **Coverage-quality sprint 3 (86→90 milestone reached; commit gate activated):** made the four weakest subsystems behaviorally tested instead of line-filled. **Notification adapters:** the vendored OAuth1 library (759 lines) is trimmed to the surface production uses — RSA/PLAINTEXT/HMAC-SHA256 signing, `SignForm`/`SetAuthorizationHeader`, `Put`/`Delete`, xAuth and session-renewal flows (zero callers) are deleted along with `SendTweetWithAttachment`/`SendTweetWithMedia`; the remaining HMAC-SHA1 signer is pinned against Twitter's documented signature vector plus a Python RFC 5849 cross-check, and httptest suites drive the full PIN authorization flow, wire-shape assertions and the complete INIT/APPEND/FINALIZE/STATUS chunked-video protocol (poll interval extracted so the stuck-processing abort is testable; the `ProcessingInfo.State` JSON tag pointed at a nonexistent field and now decodes `state`). `wanotif` gains an injectable base URL/HTTP client, loses two `fmt.Printf` debug lines that fired on every production send, and its send/template/error/transport paths are pinned against a stub Graph API. **Freezer:** the experimental `SequentialReader` (no production callers), `ParallelReader.ItemCount` and `FreezerReader.CdatFileInfo` are deleted; the production `ParallelReader`/`WorkerReader` path gets fixture tests including multi-cdat spanning reads, handle caching, `MaxAvailableBlock` partial-data binary search and the `--info` debug surface; synthetic RLP fixtures cover the whole decode fallback chain (`decodeReceiptAlternative`, `decodeReceiptLogsOnly`, pre-Byzantium roots, typed-prefix strip) and the Arbitrum branches (7-field extended format, fallback log-field scan, skip-bad-receipt, Nitro varint headers); `skipStreamValue`'s unreachable list recursion is reduced to the scalar skip its only caller can produce. **Store:** `MintReport` was 14% covered because every fixture mint falls outside its hardcoded 2021-11→2022-12 window — in-window seeds now pin month naming and cumulative redeem halving, and the query gains the `ORDER BY` its cumulative sum silently depended on (PostgreSQL never guaranteed `GROUP BY` output order — a latent wrong-report bug); reflection sweeps prove all 73 CosmicGame + 12 RandomWalk writes abort on address-resolution failure; the round-statistics activation fallback, unclaimed-ERC20 scans (v1 + v2 segment), `store.New` against a real container (UTC pinning, keepalive dialer, retry-loop cancellation) and the watermark's three row states are covered. **Indexer:** chain-sync gains a V2-mechanics end-to-end run (duration model, reward multiplier, V2-only change divisor) and failure-mode tests (nil client, cancelled allocation, read-only corrections must propagate); loop unit tests cover the startup-watermark breaker, persist failures and metric labels. **API:** the v1 live-read failure matrix drives per-method contract reverts through the real router (cst/eth price stages, token metadata, marketing config, balances, special winners, the prize-time 200-with-error shape) plus the dashboard's sentinel→live-fallback→cache-write-back recovery; ban/unban store failures, RandomWalk identifier/paging validation, floor-price defaults, beauty-pair voter filtering and the FAQ proxy (body/Accept forwarding, 502, legacy env alias, disabled registration) complete the boundary. **Commit gate activated** (§4.6): with the authoritative race profile at **91.82%** handwritten internal, `commitGateEnabled=true` with `commitFloor` 90.0 and `internalFloor` 91.5; verified `enabled 90.00` status, a passing ≥90 profile, a fail-closed rejection under a higher floor, and `--no-verify` denial; hook tests pin the enabled state. **Tests:** 12 new test files (187 total), all 47 fuzz targets green on a full 10s re-run (one stop-deadline infra flake on the first run, no crasher), full race+shuffle unit and race integration suites, build/vet/govulncheck clean, changed-code coverage 96.75%, touched packages lint-clean (capped total 116 → 114). **Metrics:** LOC **99,920 → 102,366**, legacy internal **83.73% → 88.48%** (floor 88.2%), canonical handwritten internal **86.07% → 91.82%** (floor 91.5%, next target 95%), all production **56.26% → 59.65%** (floor 59.4%). |