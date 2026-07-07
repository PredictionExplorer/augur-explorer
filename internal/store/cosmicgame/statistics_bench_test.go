//go:build integration

package cosmicgame

// Benchmarks for the heaviest read queries (§4.5 of docs/MODERNIZATION.md),
// run against the seeded test container. Baselines live in docs/benchmarks.md;
// re-run with:
//
//	go test -tags=integration ./internal/store/cosmicgame/ -bench BenchmarkStatisticsQueries -benchmem -count=6
//
// Numbers include the container round trip; compare only against baselines
// captured the same way. The §4.2 goldens keep results correct; these keep
// them fast.

import "testing"

func benchStore(b *testing.B) *SQLStorageWrapper {
	b.Helper()
	if errSetupSkip != nil {
		b.Skipf("skipping: %v", errSetupSkip)
	}
	if sharedWrapper == nil {
		b.Fatal("store harness not initialized (TestMain did not run?)")
	}
	return sharedWrapper
}

func BenchmarkStatisticsQueries(b *testing.B) {
	sw := benchStore(b)

	// The three heaviest read paths: the multi-query dashboard aggregate,
	// the per-round claim summary CTE, and the ROI leaderboard join.
	b.Run("cosmic_game_statistics", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stats := sw.Get_cosmic_game_statistics()
			if stats.TotalBids == 0 {
				b.Fatal("statistics query returned no bids")
			}
		}
	})

	b.Run("claims_by_round", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if rows := sw.Get_claims_by_round(); len(rows) == 0 {
				b.Fatal("claims_by_round returned no rows")
			}
		}
	})

	b.Run("roi_leaderboard", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			if rows := sw.Get_roi_leaderboard(0, "roi", 0, 100); len(rows) == 0 {
				b.Fatal("roi_leaderboard returned no rows")
			}
		}
	})
}
