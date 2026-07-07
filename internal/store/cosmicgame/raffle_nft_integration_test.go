//go:build integration

package cosmicgame

import (
	"context"
	"testing"
)

func TestRaffleNFTWinnersByRound(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	for _, tc := range []struct {
		name     string
		round    int64
		isStaker bool
	}{
		{"raffle_nft_winners_by_round_0_bidders", 0, false},
		{"raffle_nft_winners_by_round_0_stakers", 0, true},
		{"raffle_nft_winners_by_round_2_bidders", 2, false},
	} {
		golden(t, tc.name, func() any {
			recs, err := r.RaffleNFTWinnersByRound(ctx, tc.round, tc.isStaker)
			if err != nil {
				t.Fatalf("RaffleNFTWinnersByRound(%d, %v): %v", tc.round, tc.isStaker, err)
			}
			return recs
		})
	}
}

func TestRaffleNFTWinners(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "raffle_nft_winners", func() any {
		recs, err := r.RaffleNFTWinners(ctx, 0, 100)
		if err != nil {
			t.Fatalf("RaffleNFTWinners: %v", err)
		}
		return recs
	})
	golden(t, "raffle_nft_winners_paged", func() any {
		recs, err := r.RaffleNFTWinners(ctx, 1, 1)
		if err != nil {
			t.Fatalf("RaffleNFTWinners paged: %v", err)
		}
		return recs
	})
}
