//go:build integration

package cosmicgame

import "testing"

func TestGetCosmicTokenHolders(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_token_holders", func() any {
		return sw.Get_cosmic_token_holders()
	})
}

func TestGetCosmicTokenStatistics(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_token_statistics", func() any {
		return sw.Get_cosmic_token_statistics()
	})
}

func TestGetUserCosmicTokenSummary(t *testing.T) {
	sw := store(t)
	golden(t, "user_cosmic_token_summary_alice", func() any {
		return sw.Get_user_cosmic_token_summary(aidAlice)
	})
	golden(t, "user_cosmic_token_summary_zero", func() any {
		// No transfer activity: pins the all-defaults shape.
		return sw.Get_user_cosmic_token_summary(aidZero)
	})
}

func TestGetCosmicTokenTotalSupplyHistoryByBid(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_token_total_supply_history_by_bid", func() any {
		return sw.Get_cosmic_token_total_supply_history_by_bid()
	})
}

func TestGetCosmicTokenTotalSupplyHistoryByDate(t *testing.T) {
	sw := store(t)
	// Fixture activity happens on 2026-01-01 (ts 1767225600 onward).
	golden(t, "cosmic_token_total_supply_history_by_date", func() any {
		return sw.Get_cosmic_token_total_supply_history_by_date("20251231", "20260301")
	})
	if got := sw.Get_cosmic_token_total_supply_history_by_date("19990101", "19990102"); len(got) != 0 {
		t.Errorf("expected no supply history rows in 1999, got %d", len(got))
	}
}
