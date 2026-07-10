//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"reflect"
	"testing"

	cgprimitives "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestPrizeClaims(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "prize_claims", func() any {
		recs, err := r.PrizeClaims(ctx, 0, 100)
		if err != nil {
			t.Fatalf("PrizeClaims: %v", err)
		}
		return recs
	})
	golden(t, "prize_claims_paged", func() any {
		recs, err := r.PrizeClaims(ctx, 1, 1)
		if err != nil {
			t.Fatalf("PrizeClaims paged: %v", err)
		}
		return recs
	})
}

func TestPrizeInfo(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	for _, round := range []int64{0, 1, 2} {
		golden(t, "prize_info_round_"+itoa(round), func() any {
			rec, err := r.PrizeInfo(ctx, round)
			if err != nil {
				t.Fatalf("PrizeInfo(round %d): %v", round, err)
			}
			return rec
		})
	}
	if _, err := r.PrizeInfo(ctx, 999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("PrizeInfo(999) = %v, want store.ErrNotFound", err)
	}
}

func TestPrizeClaimsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	first, hasMore, err := r.PrizeClaimsPage(ctx, nil, 2)
	if err != nil {
		t.Fatalf("first page: %v", err)
	}
	if !hasMore || len(first) != 2 {
		t.Fatalf("first page = %d records, hasMore=%v; want 2,true", len(first), hasMore)
	}
	if first[0].RoundNum != 2 || first[0].ClaimPrizeTx.Tx.EvtLogId != 5072 ||
		first[1].RoundNum != 1 || first[1].ClaimPrizeTx.Tx.EvtLogId != 5062 {
		t.Fatalf("first page order = (%d,%d),(%d,%d)",
			first[0].RoundNum, first[0].ClaimPrizeTx.Tx.EvtLogId,
			first[1].RoundNum, first[1].ClaimPrizeTx.Tx.EvtLogId)
	}

	second, hasMore, err := r.PrizeClaimsPage(ctx, &RoundPageCursor{
		RoundNum:   int64(first[1].RoundNum),
		EventLogID: first[1].ClaimPrizeTx.Tx.EvtLogId,
	}, 2)
	if err != nil {
		t.Fatalf("second page: %v", err)
	}
	if hasMore || len(second) != 1 {
		t.Fatalf("second page = %d records, hasMore=%v; want 1,false", len(second), hasMore)
	}
	if second[0].RoundNum != 0 || second[0].ClaimPrizeTx.Tx.EvtLogId != 5018 {
		t.Fatalf("second page record = (%d,%d), want (0,5018)",
			second[0].RoundNum, second[0].ClaimPrizeTx.Tx.EvtLogId)
	}

	exhausted, hasMore, err := r.PrizeClaimsPage(ctx, &RoundPageCursor{
		RoundNum:   0,
		EventLogID: 5018,
	}, 2)
	if err != nil {
		t.Fatalf("exhausted page: %v", err)
	}
	if hasMore || exhausted == nil || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %#v, hasMore=%v; want non-nil empty,false", exhausted, hasMore)
	}
}

func TestPrizeClaimsPagePropagatesCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, _, err := r.PrizeClaimsPage(ctx, nil, 2); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled page error = %v, want context.Canceled", err)
	}
}

func TestRoundInfoIsLeanPrizeInfoBase(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	lean, err := r.RoundInfo(ctx, 0)
	if err != nil {
		t.Fatalf("RoundInfo: %v", err)
	}
	if lean.RaffleNFTWinners != nil || lean.StakingNFTWinners != nil ||
		lean.RaffleETHDeposits != nil || lean.AllPrizes != nil {
		t.Fatalf("RoundInfo loaded nested collections: %+v", lean)
	}

	full, err := r.PrizeInfo(ctx, 0)
	if err != nil {
		t.Fatalf("PrizeInfo: %v", err)
	}
	full.RaffleNFTWinners = nil
	full.StakingNFTWinners = nil
	full.RaffleETHDeposits = nil
	full.AllPrizes = nil
	if !reflect.DeepEqual(lean, full) {
		t.Fatal("RoundInfo base differs from PrizeInfo base")
	}
}

func TestAllPrizesForRound(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "all_prizes_for_round_0", func() any {
		recs, err := r.AllPrizesForRound(ctx, 0)
		if err != nil {
			t.Fatalf("AllPrizesForRound(0): %v", err)
		}
		return recs
	})
	golden(t, "all_prizes_for_round_2", func() any {
		recs, err := r.AllPrizesForRound(ctx, 2)
		if err != nil {
			t.Fatalf("AllPrizesForRound(2): %v", err)
		}
		return recs
	})
	got, err := r.AllPrizesForRound(ctx, 999)
	if err != nil {
		t.Fatalf("AllPrizesForRound(999): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no prizes for round 999, got %d", len(got))
	}
}

func TestAllPrizesForRoundPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	want, err := r.AllPrizesForRound(ctx, 0)
	if err != nil {
		t.Fatalf("AllPrizesForRound: %v", err)
	}
	var (
		after *PrizePageCursor
		got   []cgprimitives.CGPrizeHistory
	)
	for {
		page, hasMore, err := r.AllPrizesForRoundPage(ctx, 0, after, 4)
		if err != nil {
			t.Fatalf("AllPrizesForRoundPage(after=%+v): %v", after, err)
		}
		if page == nil {
			t.Fatal("page encoded as nil")
		}
		got = append(got, page...)
		if !hasMore {
			break
		}
		if len(page) != 4 {
			t.Fatalf("short page reported hasMore: %d records", len(page))
		}
		last := page[len(page)-1]
		after = &PrizePageCursor{
			PrizeType:   last.RecordType,
			WinnerIndex: last.WinnerIndex,
		}
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("paged prizes differ from full list\ngot:  %#v\nwant: %#v", got, want)
	}

	last := want[len(want)-1]
	exhausted, hasMore, err := r.AllPrizesForRoundPage(ctx, 0, &PrizePageCursor{
		PrizeType:   last.RecordType,
		WinnerIndex: last.WinnerIndex,
	}, 4)
	if err != nil {
		t.Fatalf("exhausted page: %v", err)
	}
	if hasMore || exhausted == nil || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %#v, hasMore=%v; want non-nil empty,false", exhausted, hasMore)
	}
}

func TestRoundPrizeQueriesPropagateCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, _, err := r.AllPrizesForRoundPage(ctx, 0, nil, 2); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled prize page error = %v, want context.Canceled", err)
	}
	if _, err := r.CompletedRoundExists(ctx, 0); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled existence error = %v, want context.Canceled", err)
	}
}

func TestCompletedRoundExists(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	for _, round := range []int64{0, 1, 2} {
		exists, err := r.CompletedRoundExists(ctx, round)
		if err != nil {
			t.Fatalf("CompletedRoundExists(%d): %v", round, err)
		}
		if !exists {
			t.Errorf("completed round %d was not found", round)
		}
	}
	for _, round := range []int64{3, 999} {
		exists, err := r.CompletedRoundExists(ctx, round)
		if err != nil {
			t.Fatalf("CompletedRoundExists(%d): %v", round, err)
		}
		if exists {
			t.Errorf("uncompleted round %d reported complete", round)
		}
	}
}
