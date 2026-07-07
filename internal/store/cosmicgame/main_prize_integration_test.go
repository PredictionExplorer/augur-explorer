//go:build integration

package cosmicgame

import "testing"

func TestGetPrizeClaims(t *testing.T) {
	sw := store(t)
	golden(t, "prize_claims", func() any {
		return sw.Get_prize_claims(0, 100)
	})
	golden(t, "prize_claims_paged", func() any {
		return sw.Get_prize_claims(1, 1)
	})
}

func TestGetPrizeInfo(t *testing.T) {
	sw := store(t)
	for _, round := range []int64{0, 1, 2} {
		golden(t, "prize_info_round_"+itoa(round), func() any {
			found, rec := sw.Get_prize_info(round)
			if !found {
				t.Fatalf("expected prize info for round %d", round)
			}
			return rec
		})
	}
	if found, _ := sw.Get_prize_info(999); found {
		t.Error("expected no prize info for unclaimed round 999")
	}
}

func TestGetAllPrizesForRound(t *testing.T) {
	sw := store(t)
	golden(t, "all_prizes_for_round_0", func() any {
		return sw.Get_all_prizes_for_round(0)
	})
	golden(t, "all_prizes_for_round_2", func() any {
		return sw.Get_all_prizes_for_round(2)
	})
	if got := sw.Get_all_prizes_for_round(999); len(got) != 0 {
		t.Errorf("expected no prizes for round 999, got %d", len(got))
	}
}
