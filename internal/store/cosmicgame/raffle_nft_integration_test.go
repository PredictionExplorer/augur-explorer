//go:build integration

package cosmicgame

import (
	"cmp"
	"context"
	"errors"
	"reflect"
	"slices"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
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

func TestRaffleNFTWinnersByRoundPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	tests := []struct {
		round    int64
		isStaker bool
	}{
		{round: 0, isStaker: false},
		{round: 0, isStaker: true},
		{round: 2, isStaker: false},
	}
	for _, tc := range tests {
		legacy, err := r.RaffleNFTWinnersByRound(ctx, tc.round, tc.isStaker)
		if err != nil {
			t.Fatalf("legacy round %d pool %v: %v", tc.round, tc.isStaker, err)
		}
		var (
			after *RaffleNFTWinnerPageCursor
			paged = make([]cgmodel.CGRaffleNFTWinnerRec, 0)
		)
		for {
			page, hasMore, err := r.RaffleNFTWinnersByRoundPage(ctx, tc.round, tc.isStaker, after, 1)
			if err != nil {
				t.Fatalf("page round %d pool %v after %+v: %v", tc.round, tc.isStaker, after, err)
			}
			if page == nil {
				t.Fatal("page encoded as nil")
			}
			paged = append(paged, page...)
			if !hasMore {
				break
			}
			last := page[len(page)-1]
			after = &RaffleNFTWinnerPageCursor{
				WinnerIndex: last.WinnerIndex,
				EventLogID:  last.Tx.EvtLogId,
			}
		}
		sortWinners := func(records []cgmodel.CGRaffleNFTWinnerRec) {
			slices.SortFunc(records, func(a, b cgmodel.CGRaffleNFTWinnerRec) int {
				if byIndex := cmp.Compare(a.WinnerIndex, b.WinnerIndex); byIndex != 0 {
					return byIndex
				}
				return cmp.Compare(a.Tx.EvtLogId, b.Tx.EvtLogId)
			})
		}
		sortWinners(legacy)
		sortWinners(paged)
		if !reflect.DeepEqual(legacy, paged) {
			t.Fatalf("legacy/page records differ for round %d pool %v\nlegacy: %#v\npaged:  %#v",
				tc.round, tc.isStaker, legacy, paged)
		}

		if len(paged) > 0 {
			last := paged[len(paged)-1]
			exhausted, hasMore, err := r.RaffleNFTWinnersByRoundPage(ctx, tc.round, tc.isStaker, &RaffleNFTWinnerPageCursor{
				WinnerIndex: last.WinnerIndex,
				EventLogID:  last.Tx.EvtLogId,
			}, 1)
			if err != nil {
				t.Fatalf("exhausted page: %v", err)
			}
			if hasMore || exhausted == nil || len(exhausted) != 0 {
				t.Fatalf("exhausted page = %#v, hasMore=%v; want non-nil empty,false", exhausted, hasMore)
			}
		}
	}
}

func TestRaffleNFTWinnersByRoundPagePropagatesCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, _, err := r.RaffleNFTWinnersByRoundPage(ctx, 0, false, nil, 1); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled page error = %v, want context.Canceled", err)
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
