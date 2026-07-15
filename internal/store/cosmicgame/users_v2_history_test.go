package cosmicgame

import "testing"

func TestUserHistoryPagesRejectInvalidAddressIDOrLimit(t *testing.T) {
	t.Parallel()

	var repo Repo
	calls := map[string]func(userAid int64, limit int) error{
		"UserPrizesPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserPrizesPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserRaffleEthDepositsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserRaffleEthDepositsPage(t.Context(), userAid, nil, nil, limit)
			return err
		},
		"UserRaffleNftWinsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserRaffleNftWinsPage(t.Context(), userAid, nil, limit)
			return err
		},
		"EthDonationsByUserPage": func(userAid int64, limit int) error {
			_, _, err := repo.EthDonationsByUserPage(t.Context(), userAid, nil, limit)
			return err
		},
		"ERC20DonationsByUserPage": func(userAid int64, limit int) error {
			_, _, err := repo.ERC20DonationsByUserPage(t.Context(), userAid, nil, limit)
			return err
		},
		"NFTDonationsByUserPage": func(userAid int64, limit int) error {
			_, _, err := repo.NFTDonationsByUserPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserDonatedNftsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserDonatedNftsPage(t.Context(), userAid, nil, nil, limit)
			return err
		},
		"UserDonatedErc20Page": func(userAid int64, limit int) error {
			_, _, err := repo.UserDonatedErc20Page(t.Context(), userAid, nil, limit)
			return err
		},
	}
	inputs := []struct {
		userAid int64
		limit   int
	}{
		{userAid: 0, limit: 1},
		{userAid: -1, limit: 1},
		{userAid: 1, limit: 0},
		{userAid: 1, limit: -1},
	}
	for name, call := range calls {
		for _, input := range inputs {
			if err := call(input.userAid, input.limit); err == nil {
				t.Errorf("%s(aid=%d, limit=%d) succeeded", name, input.userAid, input.limit)
			}
		}
	}
}

func TestUserHistoryPagesRejectInvalidEventCursor(t *testing.T) {
	t.Parallel()

	var repo Repo
	badCursors := []*UserEventPageCursor{{}, {EventLogID: 0}, {EventLogID: -5}}
	calls := map[string]func(after *UserEventPageCursor) error{
		"UserRaffleEthDepositsPage": func(after *UserEventPageCursor) error {
			_, _, err := repo.UserRaffleEthDepositsPage(t.Context(), 1, nil, after, 1)
			return err
		},
		"UserRaffleNftWinsPage": func(after *UserEventPageCursor) error {
			_, _, err := repo.UserRaffleNftWinsPage(t.Context(), 1, after, 1)
			return err
		},
		"EthDonationsByUserPage": func(after *UserEventPageCursor) error {
			_, _, err := repo.EthDonationsByUserPage(t.Context(), 1, after, 1)
			return err
		},
		"ERC20DonationsByUserPage": func(after *UserEventPageCursor) error {
			_, _, err := repo.ERC20DonationsByUserPage(t.Context(), 1, after, 1)
			return err
		},
		"NFTDonationsByUserPage": func(after *UserEventPageCursor) error {
			_, _, err := repo.NFTDonationsByUserPage(t.Context(), 1, after, 1)
			return err
		},
		"UserDonatedNftsPage": func(after *UserEventPageCursor) error {
			_, _, err := repo.UserDonatedNftsPage(t.Context(), 1, nil, after, 1)
			return err
		},
	}
	for name, call := range calls {
		for _, cursor := range badCursors {
			if err := call(cursor); err == nil {
				t.Errorf("%s(cursor=%+v) succeeded", name, cursor)
			}
		}
	}
}

func TestUserPrizesPageRejectsInvalidTupleCursor(t *testing.T) {
	t.Parallel()

	var repo Repo
	badCursors := []*UserPrizePageCursor{
		{Round: -1, PrizeType: 0, WinnerIndex: 0},
		{Round: 0, PrizeType: -1, WinnerIndex: 0},
		{Round: 0, PrizeType: 16, WinnerIndex: 0},
		{Round: 0, PrizeType: 0, WinnerIndex: -1},
	}
	for _, cursor := range badCursors {
		if _, _, err := repo.UserPrizesPage(t.Context(), 1, cursor, 1); err == nil {
			t.Errorf("UserPrizesPage(cursor=%+v) succeeded", cursor)
		}
	}
}

func TestUserDonatedErc20PageRejectsInvalidCursor(t *testing.T) {
	t.Parallel()

	var repo Repo
	badCursors := []*UserDonatedErc20PageCursor{
		{Round: -1, TokenAid: 1},
		{Round: 0, TokenAid: 0},
		{Round: 0, TokenAid: -1},
	}
	for _, cursor := range badCursors {
		if _, _, err := repo.UserDonatedErc20Page(t.Context(), 1, cursor, 1); err == nil {
			t.Errorf("UserDonatedErc20Page(cursor=%+v) succeeded", cursor)
		}
	}
}
