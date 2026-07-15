package cosmicgame

import (
	"strings"
	"testing"
)

func TestUserActivityPagesRejectInvalidAddressIDOrLimit(t *testing.T) {
	t.Parallel()

	var repo Repo
	calls := map[string]func(userAid int64, limit int) error{
		"UserCosmicSignatureTokensPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserCosmicSignatureTokensPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserCosmicSignatureTransfersPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserCosmicSignatureTransfersPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserCosmicTokenTransfersPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserCosmicTokenTransfersPage(t.Context(), userAid, nil, limit)
			return err
		},
		"UserMarketingRewardsPage": func(userAid int64, limit int) error {
			_, _, err := repo.UserMarketingRewardsPage(t.Context(), userAid, nil, limit)
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

	if _, err := repo.UserCosmicTokenSummaryV2(t.Context(), 0); err == nil {
		t.Error("UserCosmicTokenSummaryV2 accepted address id 0")
	}
	if _, err := repo.UserPendingWinnings(t.Context(), -1); err == nil {
		t.Error("UserPendingWinnings accepted a negative address id")
	}
}

func TestUserActivityPagesRejectInvalidCursors(t *testing.T) {
	t.Parallel()

	var repo Repo
	badEventCursors := []*UserEventPageCursor{{}, {EventLogID: 0}, {EventLogID: -5}}
	for _, cursor := range badEventCursors {
		if _, _, err := repo.UserCosmicSignatureTransfersPage(t.Context(), 1, cursor, 1); err == nil {
			t.Errorf("UserCosmicSignatureTransfersPage(cursor=%+v) succeeded", cursor)
		}
		if _, _, err := repo.UserCosmicTokenTransfersPage(t.Context(), 1, cursor, 1); err == nil {
			t.Errorf("UserCosmicTokenTransfersPage(cursor=%+v) succeeded", cursor)
		}
		if _, _, err := repo.UserMarketingRewardsPage(t.Context(), 1, cursor, 1); err == nil {
			t.Errorf("UserMarketingRewardsPage(cursor=%+v) succeeded", cursor)
		}
	}
	if _, _, err := repo.UserCosmicSignatureTokensPage(t.Context(), 1,
		&UserTokenPageCursor{TokenID: -1}, 1); err == nil {
		t.Error("UserCosmicSignatureTokensPage accepted a negative token cursor")
	}
}

func TestUserTokenPageCursorValidity(t *testing.T) {
	t.Parallel()

	var nilCursor *UserTokenPageCursor
	if !nilCursor.valid() {
		t.Error("nil cursor must be valid (first page)")
	}
	if !(&UserTokenPageCursor{TokenID: 0}).valid() {
		t.Error("token id 0 is a legal cursor position")
	}
	if (&UserTokenPageCursor{TokenID: -1}).valid() {
		t.Error("negative token cursor accepted")
	}
}

func TestUserTransfersPageSQLShapes(t *testing.T) {
	t.Parallel()

	queries := map[string]string{
		"erc721 first":        userTransfersPageSQL("cg_erc721_transfer", "token_id", false),
		"erc721 continuation": userTransfersPageSQL("cg_erc721_transfer", "token_id", true),
		"erc20 first":         userTransfersPageSQL("cg_erc20_transfer", "value::TEXT", false),
	}
	for name, query := range queries {
		// UNION (not UNION ALL) so self-transfers qualifying on both sides
		// collapse into one row.
		if strings.Count(query, "UNION") != 1 || strings.Contains(query, "UNION ALL") {
			t.Errorf("%s: expected one deduplicating UNION:\n%s", name, query)
		}
		if strings.Count(query, "WHERE t.from_aid = $1") != 1 ||
			strings.Count(query, "WHERE t.to_aid = $1") != 1 {
			t.Errorf("%s: each branch must filter on one side:\n%s", name, query)
		}
		if strings.Count(query, "ORDER BY t.evtlog_id DESC") != 2 ||
			strings.Count(query, "ORDER BY evtlog_id DESC") != 1 {
			t.Errorf("%s: both branches and the merge must sort newest first:\n%s", name, query)
		}
		// Each branch must be bounded before the merge: three LIMITs total.
		if strings.Count(query, "LIMIT") != 3 {
			t.Errorf("%s: expected three LIMIT clauses, got %d", name, strings.Count(query, "LIMIT"))
		}
		// Direction is computed per branch so identical rows deduplicate.
		if strings.Count(query, "END AS direction") != 2 {
			t.Errorf("%s: both branches must compute the direction:\n%s", name, query)
		}
	}

	first := queries["erc721 first"]
	if strings.Contains(first, "t.evtlog_id < $2") || !strings.Contains(first, "LIMIT $2") {
		t.Errorf("first page must bind only wallet and limit:\n%s", first)
	}
	continuation := queries["erc721 continuation"]
	if strings.Count(continuation, "AND t.evtlog_id < $2") != 2 ||
		!strings.Contains(continuation, "LIMIT $3") {
		t.Errorf("continuation must bound both branches by the cursor:\n%s", continuation)
	}
	if !strings.Contains(first, "t.token_id AS value") {
		t.Errorf("erc721 branches must select the token id:\n%s", first)
	}
	if !strings.Contains(queries["erc20 first"], "t.value::TEXT AS value") {
		t.Errorf("erc20 branches must select the exact amount:\n%s", queries["erc20 first"])
	}
}
