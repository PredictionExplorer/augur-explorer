//go:build integration

package cosmicgame

import (
	"context"
	"testing"
)

func TestCosmicTokenHolders(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_token_holders", func() any {
		recs, err := r.CosmicTokenHolders(context.Background())
		if err != nil {
			t.Fatalf("CosmicTokenHolders: %v", err)
		}
		return recs
	})
}

func TestCosmicTokenStatistics(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_token_statistics", func() any {
		stats, err := r.CosmicTokenStatistics(context.Background())
		if err != nil {
			t.Fatalf("CosmicTokenStatistics: %v", err)
		}
		return stats
	})
}

func TestUserCosmicTokenSummary(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "user_cosmic_token_summary_alice", func() any {
		summary, err := r.UserCosmicTokenSummary(ctx, aidAlice)
		if err != nil {
			t.Fatalf("UserCosmicTokenSummary(alice): %v", err)
		}
		return summary
	})
	golden(t, "user_cosmic_token_summary_zero", func() any {
		// No transfer activity: pins the all-defaults shape.
		summary, err := r.UserCosmicTokenSummary(ctx, aidZero)
		if err != nil {
			t.Fatalf("UserCosmicTokenSummary(zero): %v", err)
		}
		return summary
	})
}

func TestCosmicTokenSupplyHistoryByBid(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_token_total_supply_history_by_bid", func() any {
		recs, err := r.CosmicTokenSupplyHistoryByBid(context.Background())
		if err != nil {
			t.Fatalf("CosmicTokenSupplyHistoryByBid: %v", err)
		}
		return recs
	})
}

func TestCosmicTokenSupplyHistoryByDate(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// Fixture activity happens on 2026-01-01 (ts 1767225600 onward).
	golden(t, "cosmic_token_total_supply_history_by_date", func() any {
		recs, err := r.CosmicTokenSupplyHistoryByDate(ctx, "20251231", "20260301")
		if err != nil {
			t.Fatalf("CosmicTokenSupplyHistoryByDate: %v", err)
		}
		return recs
	})
	got, err := r.CosmicTokenSupplyHistoryByDate(ctx, "19990101", "19990102")
	if err != nil {
		t.Fatalf("CosmicTokenSupplyHistoryByDate(1999): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no supply history rows in 1999, got %d", len(got))
	}
}
