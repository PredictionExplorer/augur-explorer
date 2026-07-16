package cosmicgame

import (
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestGlobalStakingActionsPageSQL(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name   string
		after  bool
		reward string
		per    string
	}{
		{name: "first page", reward: "e.reward::TEXT", per: "e.reward_per_tok::TEXT"},
		{name: "continued page", after: true, reward: "NULL::TEXT", per: "NULL::TEXT"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			query := globalStakingActionsPageSQL(
				"cg_nft_staked_cst",
				"cg_nft_unstaked_cst",
				tc.reward,
				tc.per,
				tc.after,
			)
			for _, fragment := range []string{
				"FROM cg_nft_staked_cst e",
				"FROM cg_nft_unstaked_cst e",
				"UNION ALL",
				"ORDER BY e.evtlog_id DESC",
				"LIMIT",
			} {
				if !strings.Contains(query, fragment) {
					t.Errorf("query misses %q", fragment)
				}
			}
			if got := strings.Count(query, "ORDER BY e.evtlog_id DESC"); got != 2 {
				t.Errorf("bounded branch sorts = %d, want 2", got)
			}
			hasCursor := strings.Contains(query, "WHERE e.evtlog_id < $1")
			if hasCursor != tc.after {
				t.Errorf("cursor filter present = %v, want %v", hasCursor, tc.after)
			}
		})
	}
}

func TestGlobalStakingPageValidation(t *testing.T) {
	t.Parallel()
	repo := NewRepo(&store.Store{})
	ctx := t.Context()

	for name, call := range map[string]func() error{
		"CST zero limit": func() error {
			_, _, err := repo.GlobalCstStakingActionsPage(ctx, nil, 0)
			return err
		},
		"CST bad cursor": func() error {
			_, _, err := repo.GlobalCstStakingActionsPage(
				ctx,
				&GlobalStakingActionPageCursor{},
				1,
			)
			return err
		},
		"RandomWalk zero limit": func() error {
			_, _, err := repo.GlobalRwalkStakingActionsPage(ctx, nil, 0)
			return err
		},
		"CST staked zero limit": func() error {
			_, _, err := repo.GlobalStakedCstTokensPage(ctx, nil, 0)
			return err
		},
		"CST staked bad cursor": func() error {
			_, _, err := repo.GlobalStakedCstTokensPage(
				ctx,
				&GlobalStakedTokenPageCursor{TokenID: -1},
				1,
			)
			return err
		},
		"RandomWalk staked zero limit": func() error {
			_, _, err := repo.GlobalStakedRwalkTokensPage(ctx, nil, 0)
			return err
		},
		"deposit zero limit": func() error {
			_, _, err := repo.GlobalStakingDepositsPage(ctx, nil, 0)
			return err
		},
		"deposit bad cursor": func() error {
			_, _, err := repo.GlobalStakingDepositsPage(
				ctx,
				&GlobalStakingDepositPageCursor{DepositID: -1},
				1,
			)
			return err
		},
		"round negative": func() error {
			_, _, err := repo.RoundStakingRewardsPage(ctx, -1, nil, 1)
			return err
		},
		"round zero limit": func() error {
			_, _, err := repo.RoundStakingRewardsPage(ctx, 0, nil, 0)
			return err
		},
		"round bad cursor deposit": func() error {
			_, _, err := repo.RoundStakingRewardsPage(
				ctx,
				0,
				&RoundStakingRewardPageCursor{DepositID: -1, StakerAid: 1},
				1,
			)
			return err
		},
		"round bad cursor staker": func() error {
			_, _, err := repo.RoundStakingRewardsPage(
				ctx,
				0,
				&RoundStakingRewardPageCursor{DepositID: 0},
				1,
			)
			return err
		},
		"raffle zero limit": func() error {
			_, _, err := repo.GlobalStakerRaffleNftWinsPage(ctx, false, nil, 0)
			return err
		},
		"raffle bad cursor": func() error {
			_, _, err := repo.GlobalStakerRaffleNftWinsPage(
				ctx,
				false,
				&GlobalStakerRafflePageCursor{},
				1,
			)
			return err
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := call(); err == nil {
				t.Fatal("invalid input accepted")
			}
		})
	}
}
