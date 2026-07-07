//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"

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
