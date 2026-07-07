//go:build integration

package cosmicgame

import "testing"

func TestGetCosmicGameStatistics(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_game_statistics", func() any {
		return sw.Get_cosmic_game_statistics()
	})
}

func TestGetStakeStatisticsCst(t *testing.T) {
	sw := store(t)
	golden(t, "stake_statistics_cst", func() any {
		return sw.Get_stake_statistics_cst()
	})
}

func TestGetStakeStatisticsRwalk(t *testing.T) {
	sw := store(t)
	golden(t, "stake_statistics_rwalk", func() any {
		return sw.Get_stake_statistics_rwalk()
	})
}

func TestGetCosmicGameRoundStatistics(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_game_round_statistics_0", func() any {
		return sw.Get_cosmic_game_round_statistics(0)
	})
	golden(t, "cosmic_game_round_statistics_3_open", func() any {
		return sw.Get_cosmic_game_round_statistics(3)
	})
}

func TestGetUniqueBidders(t *testing.T) {
	sw := store(t)
	golden(t, "unique_bidders", func() any {
		return sw.Get_unique_bidders()
	})
}

func TestGetUniqueWinners(t *testing.T) {
	sw := store(t)
	golden(t, "unique_winners", func() any {
		return sw.Get_unique_winners()
	})
}

func TestGetRoiLeaderboard(t *testing.T) {
	sw := store(t)
	// Every whitelisted sort column plus the default branch.
	for _, sortBy := range []string{"roi", "winrate", "spent", "nfts", "bids", "default"} {
		sortArg := sortBy
		if sortArg == "default" {
			sortArg = ""
		}
		golden(t, "roi_leaderboard_"+sortBy, func() any {
			return sw.Get_roi_leaderboard(0, sortArg, 0, 100)
		})
	}
	golden(t, "roi_leaderboard_min_bids_3", func() any {
		return sw.Get_roi_leaderboard(3, "roi", 0, 100)
	})
}

func TestGetClaimsByRound(t *testing.T) {
	sw := store(t)
	golden(t, "claims_by_round", func() any {
		return sw.Get_claims_by_round()
	})
}

func TestGetClaimDetailByRound(t *testing.T) {
	sw := store(t)
	golden(t, "claim_detail_by_round_0", func() any {
		return sw.Get_claim_detail_by_round(0)
	})
	golden(t, "claim_detail_by_round_1", func() any {
		return sw.Get_claim_detail_by_round(1)
	})
}

func TestGetUniqueStakersCst(t *testing.T) {
	sw := store(t)
	golden(t, "unique_stakers_cst", func() any {
		return sw.Get_unique_stakers_cst()
	})
}

func TestGetUniqueStakersRwalk(t *testing.T) {
	sw := store(t)
	golden(t, "unique_stakers_rwalk", func() any {
		return sw.Get_unique_stakers_rwalk()
	})
}

func TestGetUniqueStakersBoth(t *testing.T) {
	sw := store(t)
	golden(t, "unique_stakers_both", func() any {
		return sw.Get_unique_stakers_both()
	})
}

func TestGetUniqueDonors(t *testing.T) {
	sw := store(t)
	golden(t, "unique_donors", func() any {
		return sw.Get_unique_donors()
	})
}

func TestGetNFTDonationStats(t *testing.T) {
	sw := store(t)
	golden(t, "nft_donation_stats", func() any {
		return sw.Get_NFT_donation_stats()
	})
}

func TestGetRecordCounters(t *testing.T) {
	sw := store(t)
	golden(t, "record_counters", func() any {
		return sw.Get_record_counters()
	})
}

func TestGetNumPrizeClaims(t *testing.T) {
	sw := store(t)
	if got := sw.Get_num_prize_claims(); got != 3 {
		t.Errorf("num prize claims: got %d, want 3", got)
	}
}
