//go:build integration

package cosmicgame

import (
	"context"
	"testing"
)

func TestPrizeHistoryByUser(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// alice won round 0's main prize; dave and emma won later rounds.
	golden(t, "prize_history_detailed_by_user_alice", func() any {
		recs, err := r.PrizeHistoryByUser(ctx, aidAlice, 0, 100)
		if err != nil {
			t.Fatalf("PrizeHistoryByUser(alice): %v", err)
		}
		return recs
	})
	golden(t, "prize_history_detailed_by_user_dave", func() any {
		recs, err := r.PrizeHistoryByUser(ctx, aidDave, 0, 100)
		if err != nil {
			t.Fatalf("PrizeHistoryByUser(dave): %v", err)
		}
		return recs
	})
	got, err := r.PrizeHistoryByUser(ctx, aidZero, 0, 100)
	if err != nil {
		t.Fatalf("PrizeHistoryByUser(zero addr): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no prize history for the zero address, got %d", len(got))
	}
}

func TestClaimHistoryGlobal(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "claim_history_detailed_global", func() any {
		recs, err := r.ClaimHistoryGlobal(ctx, 0, 100)
		if err != nil {
			t.Fatalf("ClaimHistoryGlobal: %v", err)
		}
		return recs
	})
	golden(t, "claim_history_detailed_global_paged", func() any {
		recs, err := r.ClaimHistoryGlobal(ctx, 2, 3)
		if err != nil {
			t.Fatalf("ClaimHistoryGlobal paged: %v", err)
		}
		return recs
	})
}
