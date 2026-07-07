//go:build integration

package cosmicgame

import "testing"

func TestGetRaffleNFTWinnersByRound(t *testing.T) {
	sw := store(t)
	golden(t, "raffle_nft_winners_by_round_0_bidders", func() any {
		return sw.Get_raffle_nft_winners_by_round(0, false)
	})
	golden(t, "raffle_nft_winners_by_round_0_stakers", func() any {
		return sw.Get_raffle_nft_winners_by_round(0, true)
	})
	golden(t, "raffle_nft_winners_by_round_2_bidders", func() any {
		return sw.Get_raffle_nft_winners_by_round(2, false)
	})
}

func TestGetRaffleNFTWinners(t *testing.T) {
	sw := store(t)
	golden(t, "raffle_nft_winners", func() any {
		return sw.Get_raffle_nft_winners(0, 100)
	})
	golden(t, "raffle_nft_winners_paged", func() any {
		return sw.Get_raffle_nft_winners(1, 1)
	})
}
