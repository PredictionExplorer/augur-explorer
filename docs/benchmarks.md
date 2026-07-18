# Benchmark baselines

Performance guardrails for the hot paths (§4.5 of
[MODERNIZATION.md](MODERNIZATION.md)). Re-run after each rewrite phase and
compare with [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat);
a regression outside the noise band needs a justification in the PR that
introduces it.

## How to run

```sh
# Unit benchmarks (no Docker)
go test ./internal/indexer/cosmicgame/ -bench BenchmarkEventDecode -benchmem -count=6 -run '^$' | tee old.txt
go test ./internal/freezer/decode/ -bench BenchmarkReceiptsDecode -benchmem -count=6 -run '^$'
go test ./internal/api/common/   -bench BenchmarkRateLimiter    -benchmem -count=6 -run '^$'
go test ./internal/api/common/   -bench 'BenchmarkCompress|BenchmarkConditionalETag' -benchmem -count=6 -run '^$'

# DB benchmarks (Docker required; runs against the seeded test container)
go test -tags=integration ./internal/store/cosmicgame/ -bench BenchmarkStatisticsQueries -benchmem -count=6 -run '^$' -timeout 15m
go test -tags=integration ./internal/store/randomwalk/ -bench BenchmarkRandomWalkV2Queries -benchmem -count=6 -run '^$' -timeout 15m
go test -tags=integration ./internal/indexer/ -bench BenchmarkIngestBlock -benchmem -count=6 -run '^$' -timeout 15m

# Compare against a previous run
go install golang.org/x/perf/cmd/benchstat@latest
benchstat old.txt new.txt
```

## Baselines — 2026-07-07

The initial rows were recorded on an Apple M4 Max (arm64, 16 threads), Go
1.26.4, macOS 15. The 2026-07-10 bidding-analytics rows use the same machine
with Go 1.26.5 and macOS 26.5.1. Numbers are medians of `-count=6`. The
statistics queries include the testcontainers-Postgres round trip on the same
machine — compare them only against runs captured the same way.

| Benchmark | ns/op | B/op | allocs/op | Notes |
|---|---|---|---|---|
| `BenchmarkEventDecode` (cg-etl) | 2,130 | 2,920 | 43 | v1 `BidPlaced` ABI unpack + topic extraction |
| `BenchmarkReceiptsDecode/raw_rlp` (freezer) | 24,600 | 31,617 | 287 | 10 receipts x 3 logs, ~300 MB/s |
| `BenchmarkReceiptsDecode/snappy` (freezer) | 26,800 | 39,822 | 288 | same payload, snappy-compressed |
| `BenchmarkRateLimiter/distinct_ips` (api/common) | 1,144 | 5,373 | 15 | parallel, per-IP limiter map path (stdlib router) |
| `BenchmarkRateLimiter/shared_ip` (api/common) | 1,298 | 5,374 | 15 | parallel, single-bucket contention (stdlib router) |
| `BenchmarkCompress/gzip_32KiB` (api/common) | 114,300 | 59,500 | 30 | full middleware exchange, pooled gzip level 6, ~286 MB/s |
| `BenchmarkCompress/identity_32KiB` (api/common) | 4,410 | 47,184 | 21 | negotiation + passthrough for a non-gzip client |
| `BenchmarkConditionalETag/tag_32KiB` (api/common) | 16,800 | 88,252 | 24 | buffer + SHA-256 weak validator + release, ~1.95 GB/s |
| `BenchmarkConditionalETag/revalidate_304_32KiB` (api/common) | 13,700 | 47,627 | 26 | matching If-None-Match short-circuits to an empty 304 |
| `BenchmarkStatisticsQueries/cosmic_game_statistics` | 2,530,000 | 14,390 | 298 | multi-query dashboard aggregate (pgx-native Repo) |
| `BenchmarkStatisticsQueries/claims_by_round` | 936,000 | 9,625 | 82 | per-round claim summary CTE (pgx-native Repo) |
| `BenchmarkStatisticsQueries/roi_leaderboard` | 315,000 | 23,870 | 323 | ROI leaderboard join, sort=roi (pgx-native Repo) |
| `BenchmarkStatisticsQueries/bidding_frequency_15m` | 196,000 | 8,604 | 37 | indexed, zero-filled 15-minute frequency series with round-open exclusion |
| `BenchmarkStatisticsQueries/bidding_frequency_1h` | 192,000 | 8,131 | 20 | indexed UTC epoch-aligned hourly frequency branch |
| `BenchmarkStatisticsQueries/bidding_type_ratio_15m` | 188,000 | 19,308 | 32 | indexed, zero-filled 15-minute bid-type composition series |
| `BenchmarkStatisticsQueries/top_bidder_active_periods` | 400,000 | 13,831 | 54 | bounded lifetime top-20 lookup plus windowed session segmentation |
| `BenchmarkStatisticsQueries/bid_time_bounds` | 167,000 | 468 | 6 | first/last indexed bid timestamps |
| `BenchmarkStatisticsQueries/participant_bidders` | 172,000 | 4,092 | 26 | first 50 bidders on the indexed bid-count keyset |
| `BenchmarkStatisticsQueries/participant_winners` | 465,000 | 7,844 | 33 | canonical prize/event reconstruction; independent of replay-sensitive aggregates |
| `BenchmarkStatisticsQueries/participant_donors` | 172,000 | 3,555 | 14 | first 50 donors on the indexed exact-wei keyset |
| `BenchmarkStatisticsQueries/participant_cst_stakers` | 173,000 | 5,243 | 15 | first 50 CST stakers on the indexed reward keyset |
| `BenchmarkStatisticsQueries/participant_randomwalk_stakers` | 171,000 | 4,259 | 15 | first 50 RandomWalk stakers on the indexed token-count keyset |
| `BenchmarkStatisticsQueries/participant_dual_stakers` | 194,000 | 7,316 | 11 | first 50 dual stakers; computed cross-table token-count order |
| `BenchmarkStatisticsQueries/user_profile` | 499,000 | 3,453 | 18 | exact bounded profile with canonical prize reconstruction |
| `BenchmarkStatisticsQueries/user_bids_page` | 280,000 | 41,961 | 488 | first 50 full bid resources on the indexed user/event keyset |
| `BenchmarkStatisticsQueries/global_token_page` | 360,000 | 20,303 | 124 | first 50 global tokens with scalar-subquery mint provenance |
| `BenchmarkStatisticsQueries/cosmic_token_statistics` | 433,000 | 2,648 | 33 | one-snapshot ERC-20 aggregate with jsonb top-holder list |
| `BenchmarkStatisticsQueries/supply_by_bid_page` | 322,000 | 17,893 | 142 | first 50 supply-ledger rows with streamed running totals |
| `BenchmarkStatisticsQueries/global_staking_actions_page` | 207,000 | 14,283 | 55 | bounded two-branch global CST stake/unstake event merge |
| `BenchmarkStatisticsQueries/global_staked_tokens_page` | 196,000 | 10,273 | 18 | live globally staked CST membership with mint provenance |
| `BenchmarkStatisticsQueries/global_staking_deposits_page` | 433,000 | 12,397 | 19 | page-first exact reward-deposit aggregates and claim progress |
| `BenchmarkStatisticsQueries/round_staking_rewards_page` | 181,000 | 7,192 | 19 | one round's per-staker exact reward allocations |
| `BenchmarkStatisticsQueries/global_staker_raffle_page` | 186,000 | 10,576 | 35 | one filtered global staker-raffle NFT page |
| `BenchmarkRandomWalkV2Queries/token_events` | 546,000 | 43,640 | 224 | six-branch bounded per-token provenance merge |
| `BenchmarkRandomWalkV2Queries/offer_history` | 410,000 | 20,444 | 176 | offer-creation ledger with purchase/cancel outcome joins |
| `BenchmarkRandomWalkV2Queries/offers_price_asc` | 345,000 | 10,110 | 37 | live order book on the partial (price, evtlog) index |
| `BenchmarkRandomWalkV2Queries/statistics` | 399,000 | 4,888 | 90 | one-snapshot collection/marketplace/withdrawal aggregate |
| `BenchmarkRankingQueries/ratings_page` | 180,000 | 2,645 | 17 | full rating directory page with per-token match counts |
| `BenchmarkRankingQueries/statistics` | 166,000 | 580 | 6 | one-snapshot beauty-contest vote/voter/rated counters |
| `BenchmarkIngestBlock/tx_block_3_logs` (indexer) | 3,940,000 | 7,326 | 127 | steady-state replay of a 3-log block through the per-block ingestion transaction (ADR-0010): BEGIN + pipeline + watermark + COMMIT |
| `BenchmarkIngestBlock/autocommit_3_logs` (indexer) | 3,620,000 | 7,665 | 120 | the identical pipeline without the transaction wrapper — the delta (~330µs, ~9%) is the atomicity cost |

History:

- **2026-07-07** — initial baselines (this sprint). The receipts decoder
  numbers include the format-detection fix from the same sprint (snappy blobs
  whose length-uvarint starts with an RLP-like byte were previously
  undecodable; see `rlpListCoversExactly`).
- **2026-07-07 (store groundwork sprint)** — statistics queries re-measured
  after the connection layer moved to a shared `pgxpool` (the queries
  themselves are still legacy `database/sql` code running through the pool's
  `Store.DB()` view): `cosmic_game_statistics` 2,390,000 ns/op / 40,830 B/op /
  512 allocs, `claims_by_round` 845,000 / 19,728 / 186, `roi_leaderboard`
  267,000 / 20,282 / 161. Latency improved across the board; B/op roughly
  doubled through the stdlib-over-pool bridge — acceptable while
  transitional, re-measure when `statistics.go` converts to pgx-native.
- **2026-07-07 (read-layer conversion sprint 3)** — `statistics.go` converted
  to pgx-native Repo methods; the benchmark now runs `CosmicGameStatistics`,
  `ClaimsByRound` and `RoiLeaderboard` (table above updated). The
  stdlib-over-pool B/op inflation flagged in the previous entry is gone:
  `cosmic_game_statistics` 40,830 → 14,390 B/op (298 allocs, down from 512)
  and `claims_by_round` 19,728 → 9,625 B/op (82 allocs, down from 186).
  `roi_leaderboard` allocates more (20,282 → 23,870 B/op, 161 → 323 allocs)
  because pgx's binary protocol decodes its four NUMERIC columns through
  big-number types before formatting the strings the record shape pins —
  latency is unchanged inside the container noise band (ns/op medians moved
  −4%/+11%/+18% vs the bridge run and −5%/−2%/+1% vs the original
  database/sql baselines).
- **2026-07-09 (stdlib router sprint)** — the rate-limiter benchmark now
  drives the `httpx.Router` (net/http ServeMux) instead of the gin engine;
  table re-based. The full stack got faster: `distinct_ips` 1,510 → 1,144
  ns/op (−24%), `shared_ip` 1,600 → 1,298 ns/op (−19%), B/op within 20 bytes
  and allocs identical at 15. The statistics queries (untouched by the
  router) re-ran clean: 2,219,000 / 780,000 / 259,000 ns/op medians vs the
  2,530,000 / 936,000 / 315,000 baselines with byte-identical B/op and
  allocs — no regression.
- **2026-07-10 (API-v2 participant-directory sprint)** — added first-page
  baselines for all six directories after their deterministic keyset queries and
  concurrent read indexes landed. Five queries run in 171–194 µs against the
  seeded container. Winners take 465 µs because counts and ETH totals are
  reconstructed from canonical prize/event rows instead of the
  replay-sensitive `cg_winner` aggregate. The computed dual-staker order
  remains in the indexed-query latency band, so no ineffective cross-table
  index or materialized cache was added.
- **2026-07-10 (API-v2 bidding-analytics sprint)** — added six-run medians for
  the four backing queries, including both anchored 15-minute and UTC
  epoch-aligned hourly frequency branches. Migration 00016 adds the timestamp
  range index; v2 queries filter once and use `DATE_BIN` before joining the
  zero-fill series, fixing partial-tail leakage and avoiding bucket-by-row
  range joins. Frequency, type composition and time bounds run in 167–196 µs;
  the bounded two-query top-bidder/session path takes 400 µs. All remain below
  the existing 465 µs participant ceiling on the seeded container.
- **2026-07-11 (API-v2 user-foundation sprint)** — added six-run medians for
  the collection-free profile and first 50 newest user bids. Canonical prize
  reconstruction keeps the profile at 499 µs; the full bid projection is
  280 µs on migration 00017's `(bidder_aid, evtlog_id DESC)` index.
- **2026-07-13 (HTTP edge sprint)** — added the compression and
  conditional-request middleware baselines (32 KiB rows in the table; the
  1 KiB and 512 KiB size points recorded in the same run scale linearly:
  gzip 22.5 µs/1.75 ms, validator 1.8 µs/227 µs). Compression costs ~110 µs
  per 32 KiB response — well under the 170 µs+ database time of the
  cheapest indexed query it would ride on — and reduces JSON bodies ~10×.
  The rate-limiter benchmark re-ran clean against its baselines
  (`distinct_ips` 1,140 ns/op median vs 1,144, `shared_ip` 1,270 vs 1,298,
  B/op and allocs identical).
- **2026-07-15 (API-v2 global-directories sprint)** — added the directory
  query class: the global token page, the one-snapshot Cosmic Token
  statistics and the supply-by-bid page (table rows above). The token page
  originally reused the v1 ten-relation join shape and benchmarked at
  7.25 ms; EXPLAIN showed PostgreSQL spending 7.5 ms *planning* a query
  that executes in 0.2 ms, so the five prize-family joins became scalar
  subqueries — 360 µs steady-state, a 20× improvement, with the added
  property that a duplicated prize row now fails loudly instead of
  duplicating directory rows. Statistics/claims/ROI re-ran within their
  envelopes with byte-identical allocation counts. `user_profile` medians
  moved to ~780 µs against the 499 µs baseline with identical B/op and
  allocs — three sprints of extension seeds have grown the joined tables
  since that baseline, so this is data volume, not a plan change;
  re-baseline when the fixture set stabilizes.
- **2026-07-16 (API-v2 global-staking sprint)** — added five global
  staking query baselines (table rows above). The bounded action union,
  live membership, exact deposit aggregation, round allocation and
  pool-filtered raffle queries run in 181–207 µs against the seeded
  container; the page-first reward-deposit aggregate is 433 µs and bounds
  its reward scan to the selected deposits. The eight concurrent indexes
  in migration 00023 support action detail, deposit aggregation,
  round-allocation and global raffle keysets; no cache or denormalized read
  model was justified.
- **2026-07-16 (API-v2 RandomWalk sprint)** — added the RandomWalk query
  class in its own `BenchmarkRandomWalkV2Queries` suite (table rows above).
  The structurally heaviest read — the six-branch bounded per-token event
  merge, where every branch applies its own keyset filter and limit before
  the outer merge — runs in 546 µs; the outcome-joined offer ledger in
  410 µs, the price-ranked live book in 345 µs on migration 00024's
  partial `(contract, price, evtlog)` index, and the one-snapshot
  statistics aggregate in 399 µs. All sit inside the established
  170–550 µs container-round-trip band, so no denormalized read model or
  cache was added.
- **2026-07-16 (API-v2 ranking-write sprint)** — added the two
  beauty-contest read baselines (`BenchmarkRankingQueries`, table rows
  above). Both queries scan tables bounded by the frozen ~4k-token
  collection (`rw_token`, `rw_token_ranking`, `rw_ranking_match`), landing
  at the floor of the container-round-trip band — which is why the sprint
  deliberately added **no migration**: a keyset index over
  `COALESCE(rating, 1200)` would optimize a sub-200 µs sequential scan of
  a size-capped table.
- **2026-07-18 (transactional-ingestion sprint)** — added the
  `BenchmarkIngestBlock` pair quantifying ADR-0010's per-block transaction:
  a three-log block replayed in steady state costs ~3.94 ms through the
  transaction versus ~3.62 ms through the raw autocommit pipeline (~10
  container round trips either way), so the BEGIN/COMMIT/watermark wrapper
  adds ~330 µs (~9%) per block — far below the RPC latency that dominates
  live ingestion, and in production the single fsync per block *replaces*
  one per autocommitted statement. The guarded query benchmarks re-ran
  after every repo method moved onto the per-call querier resolution: the
  CosmicGame suite ran at-or-faster than its baselines
  (`cosmic_token_statistics` 233 µs vs 433, `supply_by_bid_page` 206 µs vs
  322, the staking pages 170–212 µs vs 181–433) and a serial A/B against
  the pre-change commit (same machine, back-to-back worktree runs) put the
  ranking queries dead even — `ratings_page` 178 µs new vs 177 µs old,
  `statistics` 164 µs new vs 166 µs old — with byte-identical B/op and
  allocation counts everywhere: the querier resolution is one context
  lookup per query and costs nothing measurable. (A first parallel run had
  shown the RandomWalk suite +14–18%; that was contention from running two
  container benchmark suites simultaneously, disproven by the serial A/B.)
