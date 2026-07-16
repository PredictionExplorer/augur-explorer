package cosmicgame

import (
	"context"
	"strings"
	"testing"
)

func TestEscapeLikePattern(t *testing.T) {
	t.Parallel()
	cases := map[string]string{
		"plain":      "plain",
		"100%":       `100\%`,
		"a_b":        `a\_b`,
		`back\slash`: `back\\slash`,
		`%_\`:        `\%\_\\`,
		"":           "",
	}
	for input, want := range cases {
		if got := escapeLikePattern(input); got != want {
			t.Errorf("escapeLikePattern(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestGlobalTokenFilterValidity(t *testing.T) {
	t.Parallel()
	valid := []GlobalTokenFilter{
		{},
		{NamedOnly: true},
		{NameContains: "cosmic"},
	}
	for _, filter := range valid {
		if !filter.valid() {
			t.Errorf("filter %+v reported invalid", filter)
		}
	}
	if (GlobalTokenFilter{NamedOnly: true, NameContains: "x"}).valid() {
		t.Error("contradictory filter reported valid")
	}
}

func TestGlobalDirectoryCursorValidity(t *testing.T) {
	t.Parallel()
	var nilToken *GlobalTokenPageCursor
	if !nilToken.valid() || !(&GlobalTokenPageCursor{TokenID: 0}).valid() {
		t.Error("valid token cursors rejected")
	}
	if (&GlobalTokenPageCursor{TokenID: -1}).valid() {
		t.Error("negative token cursor accepted")
	}
	var nilEvent *TokenEventPageCursor
	if !nilEvent.valid() || !(&TokenEventPageCursor{EventLogID: 1}).valid() {
		t.Error("valid event cursors rejected")
	}
	if (&TokenEventPageCursor{EventLogID: 0}).valid() {
		t.Error("zero event cursor accepted")
	}
	var nilSupply *SupplyChangePageCursor
	if !nilSupply.valid() || !(&SupplyChangePageCursor{EventLogID: 1}).valid() {
		t.Error("valid supply cursors rejected")
	}
	if (&SupplyChangePageCursor{EventLogID: 0}).valid() {
		t.Error("zero supply cursor accepted")
	}
}

// TestGlobalTokenSelectContainsEveryProvenanceJoin pins the query shape the
// mint-source derivation depends on: all five prize-family subqueries plus
// the live staking membership must stay in the statement.
func TestGlobalTokenSelectContainsEveryProvenanceJoin(t *testing.T) {
	t.Parallel()
	for _, fragment := range []string{
		"cg_staked_token_cst",
		"cg_prize_claim",
		"cg_raffle_nft_prize",
		"cg_endurance_prize",
		"cg_lastcst_prize",
		"cg_chrono_warrior_prize",
	} {
		if !strings.Contains(globalTokenSelectSQL, fragment) {
			t.Errorf("global token select misses %s", fragment)
		}
	}
}

// TestGlobalDirectoryInputValidation proves every page method rejects
// invalid identifiers, limits and cursors before touching the pool — the
// zero-value Repo has no connection, so reaching the database would panic.
func TestGlobalDirectoryInputValidation(t *testing.T) {
	t.Parallel()
	r := &Repo{}
	ctx := context.Background()

	calls := map[string]func() error{
		"global page zero limit": func() error {
			_, _, err := r.CosmicSignatureTokensGlobalPage(ctx, GlobalTokenFilter{}, nil, 0)
			return err
		},
		"global page negative cursor": func() error {
			_, _, err := r.CosmicSignatureTokensGlobalPage(
				ctx, GlobalTokenFilter{}, &GlobalTokenPageCursor{TokenID: -1}, 1)
			return err
		},
		"global page contradictory filter": func() error {
			_, _, err := r.CosmicSignatureTokensGlobalPage(
				ctx, GlobalTokenFilter{NamedOnly: true, NameContains: "x"}, nil, 1)
			return err
		},
		"detail negative id": func() error {
			_, err := r.CosmicSignatureTokenDetailV2(ctx, -1)
			return err
		},
		"name history negative id": func() error {
			_, _, err := r.TokenNameHistoryPage(ctx, -1, nil, 1)
			return err
		},
		"name history zero cursor": func() error {
			_, _, err := r.TokenNameHistoryPage(ctx, 1, &TokenEventPageCursor{}, 1)
			return err
		},
		"transfers negative id": func() error {
			_, _, err := r.TokenTransfersPage(ctx, -1, nil, 1)
			return err
		},
		"transfers zero limit": func() error {
			_, _, err := r.TokenTransfersPage(ctx, 1, nil, 0)
			return err
		},
		"cs holders cross-kind cursor": func() error {
			_, _, err := r.CosmicSignatureHoldersPage(ctx, &ParticipantPageCursor{
				Kind: ParticipantBidders, SortValue: "1", AddressID: 1,
			}, 1)
			return err
		},
		"ct holders cross-kind cursor": func() error {
			_, _, err := r.CosmicTokenHoldersPage(ctx, &ParticipantPageCursor{
				Kind: ParticipantCsTokenHolders, SortValue: "1", AddressID: 1,
			}, 1)
			return err
		},
		"ct holders zero limit": func() error {
			_, _, err := r.CosmicTokenHoldersPage(ctx, nil, 0)
			return err
		},
		"supply zero limit": func() error {
			_, _, err := r.CosmicTokenSupplyByBidPage(ctx, nil, 0)
			return err
		},
		"supply zero cursor": func() error {
			_, _, err := r.CosmicTokenSupplyByBidPage(ctx, &SupplyChangePageCursor{}, 1)
			return err
		},
		"marketing zero limit": func() error {
			_, _, err := r.MarketingRewardsGlobalPage(ctx, nil, 0)
			return err
		},
		"marketing zero cursor": func() error {
			_, _, err := r.MarketingRewardsGlobalPage(ctx, &UserEventPageCursor{}, 1)
			return err
		},
	}
	for name, call := range calls {
		if err := call(); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
}
