//go:build integration

package cosmicgame

import (
	"context"
	"testing"
)

func TestGetCosmicGameStatistics(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_game_statistics", func() any {
		stats, err := r.CosmicGameStatistics(context.Background())
		if err != nil {
			t.Fatalf("CosmicGameStatistics: %v", err)
		}
		return stats
	})
}

func TestGetStakeStatisticsCst(t *testing.T) {
	r := repo(t)
	golden(t, "stake_statistics_cst", func() any {
		stats, err := r.StakeStatisticsCst(context.Background())
		if err != nil {
			t.Fatalf("StakeStatisticsCst: %v", err)
		}
		return stats
	})
}

func TestGetStakeStatisticsRwalk(t *testing.T) {
	r := repo(t)
	golden(t, "stake_statistics_rwalk", func() any {
		stats, err := r.StakeStatisticsRwalk(context.Background())
		if err != nil {
			t.Fatalf("StakeStatisticsRwalk: %v", err)
		}
		return stats
	})
}

func TestGetCosmicGameRoundStatistics(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_game_round_statistics_0", func() any {
		stats, err := r.CosmicGameRoundStatistics(context.Background(), 0)
		if err != nil {
			t.Fatalf("CosmicGameRoundStatistics(0): %v", err)
		}
		return stats
	})
	golden(t, "cosmic_game_round_statistics_3_open", func() any {
		stats, err := r.CosmicGameRoundStatistics(context.Background(), 3)
		if err != nil {
			t.Fatalf("CosmicGameRoundStatistics(3): %v", err)
		}
		return stats
	})
}

func TestGetUniqueBidders(t *testing.T) {
	r := repo(t)
	golden(t, "unique_bidders", func() any {
		recs, err := r.UniqueBidders(context.Background())
		if err != nil {
			t.Fatalf("UniqueBidders: %v", err)
		}
		return recs
	})
}

func TestGetUniqueWinners(t *testing.T) {
	r := repo(t)
	golden(t, "unique_winners", func() any {
		recs, err := r.UniqueWinners(context.Background())
		if err != nil {
			t.Fatalf("UniqueWinners: %v", err)
		}
		return recs
	})
}

func TestGetRoiLeaderboard(t *testing.T) {
	r := repo(t)
	// Every whitelisted sort column plus the default branch.
	for _, sortBy := range []string{"roi", "winrate", "spent", "nfts", "bids", "default"} {
		sortArg := sortBy
		if sortArg == "default" {
			sortArg = ""
		}
		golden(t, "roi_leaderboard_"+sortBy, func() any {
			recs, err := r.RoiLeaderboard(context.Background(), 0, sortArg, 0, 100)
			if err != nil {
				t.Fatalf("RoiLeaderboard(%q): %v", sortArg, err)
			}
			return recs
		})
	}
	golden(t, "roi_leaderboard_min_bids_3", func() any {
		recs, err := r.RoiLeaderboard(context.Background(), 3, "roi", 0, 100)
		if err != nil {
			t.Fatalf("RoiLeaderboard(min_bids=3): %v", err)
		}
		return recs
	})
}

func TestGetClaimsByRound(t *testing.T) {
	r := repo(t)
	golden(t, "claims_by_round", func() any {
		recs, err := r.ClaimsByRound(context.Background())
		if err != nil {
			t.Fatalf("ClaimsByRound: %v", err)
		}
		return recs
	})
}

func TestGetClaimDetailByRound(t *testing.T) {
	r := repo(t)
	golden(t, "claim_detail_by_round_0", func() any {
		detail, err := r.ClaimDetailByRound(context.Background(), 0)
		if err != nil {
			t.Fatalf("ClaimDetailByRound(0): %v", err)
		}
		return detail
	})
	golden(t, "claim_detail_by_round_1", func() any {
		detail, err := r.ClaimDetailByRound(context.Background(), 1)
		if err != nil {
			t.Fatalf("ClaimDetailByRound(1): %v", err)
		}
		return detail
	})
}

func TestGetUniqueStakersCst(t *testing.T) {
	r := repo(t)
	golden(t, "unique_stakers_cst", func() any {
		recs, err := r.UniqueStakersCst(context.Background())
		if err != nil {
			t.Fatalf("UniqueStakersCst: %v", err)
		}
		return recs
	})
}

func TestGetUniqueStakersRwalk(t *testing.T) {
	r := repo(t)
	golden(t, "unique_stakers_rwalk", func() any {
		recs, err := r.UniqueStakersRwalk(context.Background())
		if err != nil {
			t.Fatalf("UniqueStakersRwalk: %v", err)
		}
		return recs
	})
}

func TestGetUniqueStakersBoth(t *testing.T) {
	r := repo(t)
	golden(t, "unique_stakers_both", func() any {
		recs, err := r.UniqueStakersBoth(context.Background())
		if err != nil {
			t.Fatalf("UniqueStakersBoth: %v", err)
		}
		return recs
	})
}

func TestGetUniqueDonors(t *testing.T) {
	r := repo(t)
	golden(t, "unique_donors", func() any {
		recs, err := r.UniqueDonors(context.Background())
		if err != nil {
			t.Fatalf("UniqueDonors: %v", err)
		}
		return recs
	})
}

func TestGetNFTDonationStats(t *testing.T) {
	r := repo(t)
	golden(t, "nft_donation_stats", func() any {
		recs, err := r.NFTDonationStats(context.Background())
		if err != nil {
			t.Fatalf("NFTDonationStats: %v", err)
		}
		return recs
	})
}

func TestGetRecordCounters(t *testing.T) {
	r := repo(t)
	golden(t, "record_counters", func() any {
		counters, err := r.RecordCounters(context.Background())
		if err != nil {
			t.Fatalf("RecordCounters: %v", err)
		}
		return counters
	})
}
