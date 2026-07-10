package cosmicgame

import "testing"

func TestRafflePageMethodsRejectInvalidArguments(t *testing.T) {
	t.Parallel()

	var repo Repo
	if _, _, err := repo.RaffleEthDepositsByRoundPage(t.Context(), -1, nil, 1); err == nil {
		t.Fatal("RaffleEthDepositsByRoundPage accepted a negative round")
	}
	if _, _, err := repo.RaffleNFTWinnersByRoundPage(t.Context(), -1, false, nil, 1); err == nil {
		t.Fatal("RaffleNFTWinnersByRoundPage accepted a negative round")
	}
	for _, limit := range []int{0, -1} {
		if _, _, err := repo.RaffleEthDepositsByRoundPage(t.Context(), 1, nil, limit); err == nil {
			t.Errorf("RaffleEthDepositsByRoundPage(limit=%d) succeeded", limit)
		}
		if _, _, err := repo.RaffleNFTWinnersByRoundPage(t.Context(), 1, false, nil, limit); err == nil {
			t.Errorf("RaffleNFTWinnersByRoundPage(limit=%d) succeeded", limit)
		}
	}

	for name, cursor := range map[string]RaffleEthDepositPageCursor{
		"negative winner": {WinnerIndex: -1, EventLogID: 1},
		"zero event":      {WinnerIndex: 0, EventLogID: 0},
	} {
		t.Run("eth/"+name, func(t *testing.T) {
			t.Parallel()
			if _, _, err := repo.RaffleEthDepositsByRoundPage(t.Context(), 1, &cursor, 1); err == nil {
				t.Fatal("RaffleEthDepositsByRoundPage accepted invalid cursor")
			}
		})
	}
	for name, cursor := range map[string]RaffleNFTWinnerPageCursor{
		"negative winner": {WinnerIndex: -1, EventLogID: 1},
		"zero event":      {WinnerIndex: 0, EventLogID: 0},
	} {
		t.Run("nft/"+name, func(t *testing.T) {
			t.Parallel()
			if _, _, err := repo.RaffleNFTWinnersByRoundPage(t.Context(), 1, false, &cursor, 1); err == nil {
				t.Fatal("RaffleNFTWinnersByRoundPage accepted invalid cursor")
			}
		})
	}
}
