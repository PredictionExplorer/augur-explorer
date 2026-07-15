package cosmicgame

import (
	"strings"
	"testing"
)

func TestUserStakingPagesRejectInvalidAddressIDOrLimit(t *testing.T) {
	t.Parallel()

	var repo Repo
	calls := map[string]func(userAid int64, limit int) error{
		"UserCstStakingActionsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserCstStakingActionsPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserRwalkStakingActionsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserRwalkStakingActionsPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserStakedCstTokensPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserStakedCstTokensPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserStakedRwalkTokensPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserStakedRwalkTokensPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserStakingDepositsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserStakingDepositsPage(t.Context(), userAid, nil, nil, limit)
			return err
		},
		"UserStakingDepositRewardsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserStakingDepositRewardsPage(t.Context(), userAid, 1, nil, limit)
			return err
		},
		"UserStakingTokenRewardsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserStakingTokenRewardsPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserStakingTokenRewardDepositsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserStakingTokenRewardDepositsPage(t.Context(), userAid, 1, nil, limit)
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

func TestUserStakingPagesRejectInvalidScopeIdentifiers(t *testing.T) {
	t.Parallel()

	var repo Repo
	if _, _, err := repo.UserStakingDepositRewardsPage(t.Context(), 1, -1, nil, 1); err == nil {
		t.Error("UserStakingDepositRewardsPage accepted a negative deposit id")
	}
	if _, _, err := repo.UserStakingTokenRewardDepositsPage(t.Context(), 1, -1, nil, 1); err == nil {
		t.Error("UserStakingTokenRewardDepositsPage accepted a negative token id")
	}
}

func TestUserStakingPagesRejectInvalidCursors(t *testing.T) {
	t.Parallel()

	var repo Repo
	badEventCursors := []*UserEventPageCursor{{}, {EventLogID: 0}, {EventLogID: -5}}
	for _, cursor := range badEventCursors {
		if _, _, err := repo.UserCstStakingActionsPage(t.Context(), 1, cursor, 1); err == nil {
			t.Errorf("UserCstStakingActionsPage(cursor=%+v) succeeded", cursor)
		}
		if _, _, err := repo.UserRwalkStakingActionsPage(t.Context(), 1, cursor, 1); err == nil {
			t.Errorf("UserRwalkStakingActionsPage(cursor=%+v) succeeded", cursor)
		}
	}

	badTokenCursor := &UserStakingTokenPageCursor{TokenID: -1}
	if _, _, err := repo.UserStakedCstTokensPage(t.Context(), 1, badTokenCursor, 1); err == nil {
		t.Error("UserStakedCstTokensPage accepted a negative token cursor")
	}
	if _, _, err := repo.UserStakedRwalkTokensPage(t.Context(), 1, badTokenCursor, 1); err == nil {
		t.Error("UserStakedRwalkTokensPage accepted a negative token cursor")
	}
	if _, _, err := repo.UserStakingTokenRewardsPage(t.Context(), 1, badTokenCursor, 1); err == nil {
		t.Error("UserStakingTokenRewardsPage accepted a negative token cursor")
	}
	if _, _, err := repo.UserStakingDepositsPage(t.Context(), 1, nil,
		&UserStakingDepositPageCursor{DepositID: -1}, 1); err == nil {
		t.Error("UserStakingDepositsPage accepted a negative deposit cursor")
	}
	if _, _, err := repo.UserStakingDepositRewardsPage(t.Context(), 1, 1,
		&UserStakingRewardPageCursor{ActionID: -1}, 1); err == nil {
		t.Error("UserStakingDepositRewardsPage accepted a negative action cursor")
	}
	if _, _, err := repo.UserStakingTokenRewardDepositsPage(t.Context(), 1, 1,
		&UserStakingTokenDepositPageCursor{DepositID: -1}, 1); err == nil {
		t.Error("UserStakingTokenRewardDepositsPage accepted a negative deposit cursor")
	}
}

func TestUserStakingActionsPageSQLShapes(t *testing.T) {
	t.Parallel()

	queries := map[string]string{
		"cst first":        userStakingActionsPageSQL("cg_nft_staked_cst", "cg_nft_unstaked_cst", "e.reward::TEXT", false),
		"cst continuation": userStakingActionsPageSQL("cg_nft_staked_cst", "cg_nft_unstaked_cst", "e.reward::TEXT", true),
		"rwalk first":      userStakingActionsPageSQL("cg_nft_staked_rwalk", "cg_nft_unstaked_rwalk", "NULL::TEXT", false),
	}
	for name, query := range queries {
		if strings.Count(query, "UNION ALL") != 1 {
			t.Errorf("%s: expected one UNION ALL branch merge:\n%s", name, query)
		}
		if strings.Count(query, "WHERE e.staker_aid = $1") != 2 {
			t.Errorf("%s: both branches must filter on the staker:\n%s", name, query)
		}
		if strings.Count(query, "ORDER BY e.evtlog_id DESC") != 2 ||
			strings.Count(query, "ORDER BY evtlog_id DESC") != 1 {
			t.Errorf("%s: both branches and the merge must sort newest first:\n%s", name, query)
		}
		// Each branch must be bounded before the merge: three LIMITs total.
		if strings.Count(query, "LIMIT") != 3 {
			t.Errorf("%s: expected three LIMIT clauses, got %d", name, strings.Count(query, "LIMIT"))
		}
	}

	first := queries["cst first"]
	if strings.Contains(first, "e.evtlog_id < $2") || !strings.Contains(first, "LIMIT $2") {
		t.Errorf("first page must bind only staker and limit:\n%s", first)
	}
	continuation := queries["cst continuation"]
	if strings.Count(continuation, "AND e.evtlog_id < $2") != 2 ||
		!strings.Contains(continuation, "LIMIT $3") {
		t.Errorf("continuation must bound both branches by the cursor:\n%s", continuation)
	}
	if !strings.Contains(first, "e.reward::TEXT AS reward_wei") {
		t.Errorf("cst unstake branch must select the exact reward:\n%s", first)
	}
	if !strings.Contains(queries["rwalk first"], "NULL::TEXT AS reward_wei") {
		t.Errorf("rwalk unstake branch must select no reward:\n%s", queries["rwalk first"])
	}
}

func TestUserStakingCursorValidity(t *testing.T) {
	t.Parallel()

	var nilToken *UserStakingTokenPageCursor
	var nilDeposit *UserStakingDepositPageCursor
	var nilReward *UserStakingRewardPageCursor
	var nilTokenDeposit *UserStakingTokenDepositPageCursor
	if !nilToken.valid() || !nilDeposit.valid() || !nilReward.valid() || !nilTokenDeposit.valid() {
		t.Error("nil cursors must be valid (first page)")
	}
	if !(&UserStakingTokenPageCursor{TokenID: 0}).valid() {
		t.Error("token id 0 is a legal cursor position")
	}
	if (&UserStakingTokenPageCursor{TokenID: -1}).valid() {
		t.Error("negative token cursor accepted")
	}
	if !(&UserStakingDepositPageCursor{DepositID: 0}).valid() ||
		(&UserStakingDepositPageCursor{DepositID: -1}).valid() {
		t.Error("deposit cursor validity is wrong")
	}
	if !(&UserStakingRewardPageCursor{ActionID: 0}).valid() ||
		(&UserStakingRewardPageCursor{ActionID: -1}).valid() {
		t.Error("reward cursor validity is wrong")
	}
	if !(&UserStakingTokenDepositPageCursor{DepositID: 0}).valid() ||
		(&UserStakingTokenDepositPageCursor{DepositID: -1}).valid() {
		t.Error("token deposit cursor validity is wrong")
	}
}
