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
