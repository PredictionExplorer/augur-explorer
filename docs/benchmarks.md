# Benchmark baselines

Performance guardrails for the hot paths (§4.5 of
[MODERNIZATION.md](MODERNIZATION.md)). Re-run after each rewrite phase and
compare with [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat);
a regression outside the noise band needs a justification in the PR that
introduces it.

## How to run

```sh
# Unit benchmarks (no Docker)
go test ./cmd/cg-etl/            -bench BenchmarkEventDecode    -benchmem -count=6 -run '^$' | tee old.txt
go test ./internal/freezer/decode/ -bench BenchmarkReceiptsDecode -benchmem -count=6 -run '^$'
go test ./internal/api/common/   -bench BenchmarkRateLimiter    -benchmem -count=6 -run '^$'

# DB benchmarks (Docker required; runs against the seeded test container)
go test -tags=integration ./internal/store/cosmicgame/ -bench BenchmarkStatisticsQueries -benchmem -count=6 -run '^$' -timeout 15m

# Compare against a previous run
go install golang.org/x/perf/cmd/benchstat@latest
benchstat old.txt new.txt
```

## Baselines — 2026-07-07

Recorded on an Apple M4 Max (arm64, 16 threads), Go 1.26.4, macOS 15.
Numbers are medians of `-count=6`. The statistics queries include the
testcontainers-Postgres round trip on the same machine — compare them only
against runs captured the same way.

| Benchmark | ns/op | B/op | allocs/op | Notes |
|---|---|---|---|---|
| `BenchmarkEventDecode` (cg-etl) | 2,130 | 2,920 | 43 | v1 `BidPlaced` ABI unpack + topic extraction |
| `BenchmarkReceiptsDecode/raw_rlp` (freezer) | 24,600 | 31,617 | 287 | 10 receipts x 3 logs, ~300 MB/s |
| `BenchmarkReceiptsDecode/snappy` (freezer) | 26,800 | 39,822 | 288 | same payload, snappy-compressed |
| `BenchmarkRateLimiter/distinct_ips` (api/common) | 1,510 | 5,357 | 15 | parallel, per-IP limiter map path |
| `BenchmarkRateLimiter/shared_ip` (api/common) | 1,600 | 5,356 | 15 | parallel, single-bucket contention |
| `BenchmarkStatisticsQueries/cosmic_game_statistics` | 2,660,000 | 18,064 | 344 | multi-query dashboard aggregate |
| `BenchmarkStatisticsQueries/claims_by_round` | 955,000 | 13,224 | 138 | per-round claim summary CTE |
| `BenchmarkStatisticsQueries/roi_leaderboard` | 313,000 | 18,656 | 149 | ROI leaderboard join, sort=roi |

History:

- **2026-07-07** — initial baselines (this sprint). The receipts decoder
  numbers include the format-detection fix from the same sprint (snappy blobs
  whose length-uvarint starts with an RLP-like byte were previously
  undecodable; see `rlpListCoversExactly`).
