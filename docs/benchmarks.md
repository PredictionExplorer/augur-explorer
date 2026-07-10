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

# DB benchmarks (Docker required; runs against the seeded test container)
go test -tags=integration ./internal/store/cosmicgame/ -bench BenchmarkStatisticsQueries -benchmem -count=6 -run '^$' -timeout 15m

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
