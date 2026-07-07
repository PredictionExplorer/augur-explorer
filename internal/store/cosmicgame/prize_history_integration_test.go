//go:build integration

package cosmicgame

import "testing"

func TestGetPrizeHistoryDetailedByUser(t *testing.T) {
	sw := store(t)
	// alice won round 0's main prize; dave and emma won later rounds.
	golden(t, "prize_history_detailed_by_user_alice", func() any {
		return sw.Get_prize_history_detailed_by_user(aidAlice, 0, 100)
	})
	golden(t, "prize_history_detailed_by_user_dave", func() any {
		return sw.Get_prize_history_detailed_by_user(aidDave, 0, 100)
	})
	if got := sw.Get_prize_history_detailed_by_user(aidZero, 0, 100); len(got) != 0 {
		t.Errorf("expected no prize history for the zero address, got %d", len(got))
	}
}

func TestGetClaimHistoryDetailedGlobal(t *testing.T) {
	sw := store(t)
	golden(t, "claim_history_detailed_global", func() any {
		return sw.Get_claim_history_detailed_global(0, 100)
	})
	golden(t, "claim_history_detailed_global_paged", func() any {
		return sw.Get_claim_history_detailed_global(2, 3)
	})
}
