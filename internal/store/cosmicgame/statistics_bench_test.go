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

import (
	"context"
	"testing"
)

func benchRepo(b *testing.B) *Repo {
	b.Helper()
	if errSetupSkip != nil {
		b.Skipf("skipping: %v", errSetupSkip)
	}
	if sharedRepo == nil {
		b.Fatal("store harness not initialized (TestMain did not run?)")
	}
	return sharedRepo
}

func BenchmarkStatisticsQueries(b *testing.B) {
	r := benchRepo(b)
	ctx := context.Background()

	// The three heaviest read paths: the multi-query dashboard aggregate,
	// the per-round claim summary CTE, and the ROI leaderboard join.
	b.Run("cosmic_game_statistics", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			stats, err := r.CosmicGameStatistics(ctx)
			if err != nil {
				b.Fatalf("statistics query: %v", err)
			}
			if stats.TotalBids == 0 {
				b.Fatal("statistics query returned no bids")
			}
		}
	})

	b.Run("claims_by_round", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, err := r.ClaimsByRound(ctx)
			if err != nil {
				b.Fatalf("claims_by_round: %v", err)
			}
			if len(rows) == 0 {
				b.Fatal("claims_by_round returned no rows")
			}
		}
	})

	b.Run("roi_leaderboard", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, err := r.RoiLeaderboard(ctx, 0, "roi", 0, 100)
			if err != nil {
				b.Fatalf("roi_leaderboard: %v", err)
			}
			if len(rows) == 0 {
				b.Fatal("roi_leaderboard returned no rows")
			}
		}
	})
}
