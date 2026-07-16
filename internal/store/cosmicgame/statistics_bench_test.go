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

	b.Run("bidding_frequency_15m", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, err := r.BidFrequencyByPeriodBounded(ctx, fixtureStartTs, fixtureEndTs, 900)
			if err != nil || len(rows) == 0 {
				b.Fatalf("bidding frequency: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("bidding_frequency_1h", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, err := r.BidFrequencyByPeriodBounded(ctx, fixtureStartTs, fixtureEndTs, 3600)
			if err != nil || len(rows) == 0 {
				b.Fatalf("hourly bidding frequency: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("bidding_type_ratio_15m", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, err := r.BidTypeRatioByPeriodBounded(ctx, fixtureStartTs, fixtureEndTs, 900)
			if err != nil || len(rows) == 0 {
				b.Fatalf("bidding type ratio: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("top_bidder_active_periods", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			bidders, periods, hasMore, err := r.TopBidderActivePeriodsBounded(
				ctx, 20, fixtureStartTs, fixtureEndTs, 6, 2,
			)
			if err != nil || hasMore || len(bidders) == 0 || len(periods) == 0 {
				b.Fatalf("top bidder active periods: bidders=%d periods=%d more=%v err=%v",
					len(bidders), len(periods), hasMore, err)
			}
		}
	})

	b.Run("bid_time_bounds", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			minTimestamp, maxTimestamp, err := r.BidTimeBounds(ctx)
			if err != nil || minTimestamp == 0 || maxTimestamp < minTimestamp {
				b.Fatalf("bid time bounds: min=%d max=%d err=%v",
					minTimestamp, maxTimestamp, err)
			}
		}
	})

	b.Run("participant_bidders", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.BidderParticipantsPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("bidder participants: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("participant_winners", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.WinnerParticipantsPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("winner participants: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("participant_donors", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.DonorParticipantsPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("donor participants: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("participant_cst_stakers", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.CSTStakerParticipantsPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("CST-staker participants: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("participant_randomwalk_stakers", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.RandomWalkStakerParticipantsPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("RandomWalk-staker participants: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("participant_dual_stakers", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.DualStakerParticipantsPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("dual-staker participants: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("user_profile", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			profile, err := r.UserProfile(ctx, aidAlice)
			if err != nil || profile.BidCount == 0 {
				b.Fatalf("user profile: bids=%d err=%v", profile.BidCount, err)
			}
		}
	})

	b.Run("user_bids_page", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.BidsByUserPage(ctx, aidAlice, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("user bids page: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("global_token_page", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.CosmicSignatureTokensGlobalPage(ctx, GlobalTokenFilter{}, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("global token page: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("cosmic_token_statistics", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			record, err := r.CosmicTokenStatisticsV2(ctx)
			if err != nil || record.HolderCount == 0 {
				b.Fatalf("cosmic token statistics: holders=%d err=%v", record.HolderCount, err)
			}
		}
	})

	b.Run("supply_by_bid_page", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.CosmicTokenSupplyByBidPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("supply-by-bid page: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("global_staking_actions_page", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.GlobalCstStakingActionsPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("global staking actions page: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("global_staked_tokens_page", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.GlobalStakedCstTokensPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("global staked tokens page: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("global_staking_deposits_page", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.GlobalStakingDepositsPage(ctx, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("global staking deposits page: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("round_staking_rewards_page", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.RoundStakingRewardsPage(ctx, 0, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("round staking rewards page: rows=%d err=%v", len(rows), err)
			}
		}
	})

	b.Run("global_staker_raffle_page", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			rows, _, err := r.GlobalStakerRaffleNftWinsPage(ctx, false, nil, 50)
			if err != nil || len(rows) == 0 {
				b.Fatalf("global staker raffle page: rows=%d err=%v", len(rows), err)
			}
		}
	})
}
