package cosmicgame

import "testing"

func TestPrizeClaimsPageRejectsInvalidArguments(t *testing.T) {
	t.Parallel()

	var repo Repo
	for _, limit := range []int{0, -1} {
		if _, _, err := repo.PrizeClaimsPage(t.Context(), nil, limit); err == nil {
			t.Errorf("PrizeClaimsPage(limit=%d) succeeded", limit)
		}
	}
	for name, cursor := range map[string]RoundPageCursor{
		"negative round": {RoundNum: -1, EventLogID: 1},
		"zero event":     {RoundNum: 1, EventLogID: 0},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, _, err := repo.PrizeClaimsPage(t.Context(), &cursor, 1); err == nil {
				t.Fatal("PrizeClaimsPage accepted invalid cursor")
			}
		})
	}
}

func TestAllPrizesForRoundPageRejectsInvalidArguments(t *testing.T) {
	t.Parallel()

	var repo Repo
	if _, _, err := repo.AllPrizesForRoundPage(t.Context(), -1, nil, 1); err == nil {
		t.Fatal("AllPrizesForRoundPage accepted a negative round")
	}
	for _, limit := range []int{0, -1} {
		if _, _, err := repo.AllPrizesForRoundPage(t.Context(), 1, nil, limit); err == nil {
			t.Errorf("AllPrizesForRoundPage(limit=%d) succeeded", limit)
		}
	}
	for name, cursor := range map[string]PrizePageCursor{
		"negative type":   {PrizeType: -1, WinnerIndex: 0},
		"unknown type":    {PrizeType: 16, WinnerIndex: 0},
		"negative winner": {PrizeType: 1, WinnerIndex: -1},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, _, err := repo.AllPrizesForRoundPage(t.Context(), 1, &cursor, 1); err == nil {
				t.Fatal("AllPrizesForRoundPage accepted invalid cursor")
			}
		})
	}
	if _, err := repo.CompletedRoundExists(t.Context(), -1); err == nil {
		t.Fatal("CompletedRoundExists accepted a negative round")
	}
}
