//go:build integration

package cosmicgame

import (
	"cmp"
	"context"
	"errors"
	"reflect"
	"slices"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

func walkGlobalStakingActions(
	t *testing.T,
	fetch func(context.Context, *GlobalStakingActionPageCursor, int) ([]GlobalStakingActionRecord, bool, error),
	pageSize int,
) []GlobalStakingActionRecord {
	t.Helper()
	ctx := context.Background()
	var all []GlobalStakingActionRecord
	var after *GlobalStakingActionPageCursor
	for {
		page, hasMore, err := fetch(ctx, after, pageSize)
		if err != nil {
			t.Fatal(err)
		}
		if len(page) > pageSize {
			t.Fatalf("page length=%d limit=%d", len(page), pageSize)
		}
		for i := range page {
			if len(all) > 0 && page[i].Tx.EvtLogId >= all[len(all)-1].Tx.EvtLogId {
				t.Fatalf("unordered event ids: %d after %d",
					page[i].Tx.EvtLogId, all[len(all)-1].Tx.EvtLogId)
			}
			all = append(all, page[i])
		}
		if !hasMore {
			return all
		}
		if len(page) == 0 {
			t.Fatal("hasMore without cursor row")
		}
		after = &GlobalStakingActionPageCursor{
			EventLogID: page[len(page)-1].Tx.EvtLogId,
		}
	}
}

func TestGlobalStakingActionPagesMatchLegacy(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	for _, tc := range []struct {
		name      string
		fetch     func(context.Context, *GlobalStakingActionPageCursor, int) ([]GlobalStakingActionRecord, bool, error)
		legacyIDs func() ([]int64, error)
	}{
		{
			name:  "CST",
			fetch: r.GlobalCstStakingActionsPage,
			legacyIDs: func() ([]int64, error) {
				rows, err := r.GlobalStakingCstHistory(ctx, 0, 1000)
				ids := make([]int64, len(rows))
				for i := range rows {
					ids[i] = rows[i].Tx.EvtLogId
				}
				return ids, err
			},
		},
		{
			name:  "RandomWalk",
			fetch: r.GlobalRwalkStakingActionsPage,
			legacyIDs: func() ([]int64, error) {
				rows, err := r.GlobalStakingRwalkHistory(ctx, 0, 1000)
				ids := make([]int64, len(rows))
				for i := range rows {
					ids[i] = rows[i].Tx.EvtLogId
				}
				return ids, err
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			legacy, err := tc.legacyIDs()
			if err != nil {
				t.Fatal(err)
			}
			slices.SortFunc(legacy, func(a, b int64) int { return cmp.Compare(b, a) })
			paged := walkGlobalStakingActions(t, tc.fetch, 1)
			got := make([]int64, len(paged))
			for i := range paged {
				got[i] = paged[i].Tx.EvtLogId
				if paged[i].StakerAddress == "" || paged[i].RoundNum < 0 {
					t.Fatalf("incomplete action %+v", paged[i])
				}
			}
			if !reflect.DeepEqual(got, legacy) {
				t.Fatalf("paged=%v legacy=%v", got, legacy)
			}
		})
	}

	cst := walkGlobalStakingActions(t, r.GlobalCstStakingActionsPage, 1)
	var sawStake, sawUnstake bool
	for i := range cst {
		switch cst[i].Kind {
		case UserStakingActionStake:
			sawStake = true
			if cst[i].RewardWei != "" || cst[i].RewardPerTokenWei != "" {
				t.Fatalf("stake has reward: %+v", cst[i])
			}
		case UserStakingActionUnstake:
			sawUnstake = true
			if cst[i].RewardWei == "" || cst[i].RewardPerTokenWei == "" {
				t.Fatalf("unstake misses reward: %+v", cst[i])
			}
		}
	}
	if !sawStake || !sawUnstake {
		t.Fatalf("fixture did not exercise both CST action kinds: %+v", cst)
	}
}

func walkGlobalStakedCstTokens(t *testing.T, r *Repo, pageSize int) []GlobalStakedCstTokenRecord {
	t.Helper()
	var all []GlobalStakedCstTokenRecord
	var after *GlobalStakedTokenPageCursor
	for {
		page, more, err := r.GlobalStakedCstTokensPage(context.Background(), after, pageSize)
		if err != nil {
			t.Fatal(err)
		}
		all = append(all, page...)
		if !more {
			return all
		}
		after = &GlobalStakedTokenPageCursor{TokenID: page[len(page)-1].TokenID}
	}
}

func walkGlobalStakedRwalkTokens(t *testing.T, r *Repo, pageSize int) []GlobalStakedRwalkTokenRecord {
	t.Helper()
	var all []GlobalStakedRwalkTokenRecord
	var after *GlobalStakedTokenPageCursor
	for {
		page, more, err := r.GlobalStakedRwalkTokensPage(context.Background(), after, pageSize)
		if err != nil {
			t.Fatal(err)
		}
		all = append(all, page...)
		if !more {
			return all
		}
		after = &GlobalStakedTokenPageCursor{TokenID: page[len(page)-1].TokenID}
	}
}

func TestGlobalStakedTokenPagesMatchLegacy(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	legacyCst, err := r.StakedTokensCstGlobal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	pagedCst := walkGlobalStakedCstTokens(t, r, 1)
	wantCst := make([]int64, len(legacyCst))
	gotCst := make([]int64, len(pagedCst))
	for i := range legacyCst {
		wantCst[i] = legacyCst[i].TokenInfo.TokenId
	}
	for i := range pagedCst {
		gotCst[i] = pagedCst[i].TokenID
		if pagedCst[i].Seed == "" || pagedCst[i].StakerAddress == "" {
			t.Fatalf("incomplete CST token %+v", pagedCst[i])
		}
	}
	if !reflect.DeepEqual(gotCst, wantCst) {
		t.Fatalf("CST paged=%v legacy=%v", gotCst, wantCst)
	}

	legacyRwalk, err := r.StakedTokensRwalkGlobal(ctx)
	if err != nil {
		t.Fatal(err)
	}
	pagedRwalk := walkGlobalStakedRwalkTokens(t, r, 1)
	wantRwalk := make([]int64, len(legacyRwalk))
	gotRwalk := make([]int64, len(pagedRwalk))
	for i := range legacyRwalk {
		wantRwalk[i] = legacyRwalk[i].StakedTokenId
	}
	for i := range pagedRwalk {
		gotRwalk[i] = pagedRwalk[i].TokenID
		if pagedRwalk[i].StakerAddress == "" {
			t.Fatalf("incomplete RandomWalk token %+v", pagedRwalk[i])
		}
	}
	if !reflect.DeepEqual(gotRwalk, wantRwalk) {
		t.Fatalf("RandomWalk paged=%v legacy=%v", gotRwalk, wantRwalk)
	}
}

func walkGlobalStakingDeposits(t *testing.T, r *Repo, pageSize int) []GlobalStakingDepositRecord {
	t.Helper()
	var all []GlobalStakingDepositRecord
	var after *GlobalStakingDepositPageCursor
	for {
		page, more, err := r.GlobalStakingDepositsPage(context.Background(), after, pageSize)
		if err != nil {
			t.Fatal(err)
		}
		all = append(all, page...)
		if !more {
			return all
		}
		after = &GlobalStakingDepositPageCursor{DepositID: page[len(page)-1].DepositID}
	}
}

func TestGlobalStakingDepositsMatchLegacyAndCloseExactly(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	legacy, err := r.GlobalStakingRewards(ctx)
	if err != nil {
		t.Fatal(err)
	}
	paged := walkGlobalStakingDeposits(t, r, 1)
	if len(paged) != len(legacy) || len(paged) == 0 {
		t.Fatalf("paged=%d legacy=%d", len(paged), len(legacy))
	}
	legacyByRound := make(map[int64]string, len(legacy))
	for i := range legacy {
		legacyByRound[legacy[i].RoundNum] = legacy[i].TotalDepositAmount
	}
	for i := range paged {
		if legacyByRound[paged[i].RoundNum] != paged[i].TotalDepositWei {
			t.Fatalf("deposit %+v not in legacy %+v", paged[i], legacyByRound)
		}
		if paged[i].RewardCount < paged[i].PendingRewardCount {
			t.Fatalf("invalid counts %+v", paged[i])
		}
	}
	exhausted, more, err := r.GlobalStakingDepositsPage(
		ctx,
		&GlobalStakingDepositPageCursor{DepositID: 0},
		50,
	)
	if err != nil || more || len(exhausted) != 0 {
		t.Fatalf("exhausted=%+v more=%v err=%v", exhausted, more, err)
	}
}

func walkRoundStakingRewards(
	t *testing.T,
	r *Repo,
	round int64,
	pageSize int,
) []RoundStakingRewardRecord {
	t.Helper()
	var all []RoundStakingRewardRecord
	var after *RoundStakingRewardPageCursor
	for {
		page, more, err := r.RoundStakingRewardsPage(
			context.Background(),
			round,
			after,
			pageSize,
		)
		if err != nil {
			t.Fatal(err)
		}
		all = append(all, page...)
		if !more {
			return all
		}
		after = &RoundStakingRewardPageCursor{
			DepositID: page[len(page)-1].DepositID,
			StakerAid: page[len(page)-1].StakerAid,
		}
	}
}

func TestRoundStakingRewardsPageMatchesLegacy(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	legacy, err := r.StakingCstRewardsByRound(ctx, 0)
	if err != nil {
		t.Fatal(err)
	}
	paged := walkRoundStakingRewards(t, r, 0, 1)
	if len(paged) != len(legacy) || len(paged) == 0 {
		t.Fatalf("paged=%d legacy=%d", len(paged), len(legacy))
	}
	legacyByStaker := make(map[int64][3]string, len(legacy))
	for i := range legacy {
		legacyByStaker[legacy[i].StakerAid] = [3]string{
			legacy[i].StakerAmount,
			legacy[i].AmountCollected,
			legacy[i].AmountPendingToClaim,
		}
	}
	for i := range paged {
		want, ok := legacyByStaker[paged[i].StakerAid]
		if !ok || want != [3]string{
			paged[i].RewardWei,
			paged[i].CollectedWei,
			paged[i].PendingWei,
		} {
			t.Fatalf("reward %+v, legacy=%+v", paged[i], want)
		}
	}
	empty, more, err := r.RoundStakingRewardsPage(ctx, 999, nil, 50)
	if err != nil || more || len(empty) != 0 {
		t.Fatalf("empty round=%+v more=%v err=%v", empty, more, err)
	}
}

func walkGlobalStakerRaffle(
	t *testing.T,
	r *Repo,
	isRwalk bool,
	pageSize int,
) []cgmodel.CGRaffleNFTWinnerRec {
	t.Helper()
	var all []cgmodel.CGRaffleNFTWinnerRec
	var after *GlobalStakerRafflePageCursor
	for {
		page, more, err := r.GlobalStakerRaffleNftWinsPage(
			context.Background(),
			isRwalk,
			after,
			pageSize,
		)
		if err != nil {
			t.Fatal(err)
		}
		all = append(all, page...)
		if !more {
			return all
		}
		after = &GlobalStakerRafflePageCursor{
			EventLogID: page[len(page)-1].Tx.EvtLogId,
		}
	}
}

func TestGlobalStakerRafflePagesMatchLegacy(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	for _, tc := range []struct {
		name    string
		isRwalk bool
		legacy  func(context.Context, int, int) ([]cgmodel.CGRaffleNFTWinnerRec, error)
	}{
		{name: "CST", legacy: r.StakingCstMintsGlobal},
		{name: "RandomWalk", isRwalk: true, legacy: r.StakingRwalkMintsGlobal},
	} {
		t.Run(tc.name, func(t *testing.T) {
			legacy, err := tc.legacy(ctx, 0, 1000)
			if err != nil {
				t.Fatal(err)
			}
			paged := walkGlobalStakerRaffle(t, r, tc.isRwalk, 1)
			want := make([]int64, len(legacy))
			got := make([]int64, len(paged))
			for i := range legacy {
				want[i] = legacy[i].Tx.EvtLogId
			}
			for i := range paged {
				got[i] = paged[i].Tx.EvtLogId
				if !paged[i].IsStaker || paged[i].IsRWalk != tc.isRwalk {
					t.Fatalf("wrong pool row %+v", paged[i])
				}
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("paged=%v legacy=%v", got, want)
			}
		})
	}
}

func TestGlobalStakingMultiDepositPaginationAndRestoration(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	before := takeStakingAccumulatorSnapshot(t, r, ctx)
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8020, 117, 1018, 9, 'sdglob02', 97, '\x00')`); err != nil {
		t.Fatal(err)
	}
	removeExtension := func() {
		if _, err := r.pool().Exec(context.Background(), "DELETE FROM evt_log WHERE id=8020"); err != nil {
			t.Errorf("cleaning synthetic global staking deposit: %v", err)
		}
	}
	t.Cleanup(removeExtension)
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_staking_eth_deposit(evtlog_id, block_num, tx_id, time_stamp, contract_aid,
			round_num, deposit_time, deposit_id, num_staked_nfts, deposit_amount, amount_per_token, modulo)
		VALUES (8020, 117, 1018, TO_TIMESTAMP(1767227300), 9, 1, TO_TIMESTAMP(1767227300), 502, 1,
			3000000000000000001, 3000000000000000001, 0)`); err != nil {
		t.Fatal(err)
	}

	deposits := walkGlobalStakingDeposits(t, r, 1)
	if len(deposits) < 2 || deposits[0].DepositID != 502 ||
		deposits[0].TotalDepositWei != "3000000000000000001" ||
		deposits[0].PendingWei != "3000000000000000001" ||
		deposits[0].RemainderWei != "0" {
		t.Fatalf("global deposits=%+v", deposits)
	}
	rewards := walkRoundStakingRewards(t, r, 1, 1)
	if len(rewards) != 1 || rewards[0].DepositID != 502 ||
		rewards[0].PendingWei != "3000000000000000001" {
		t.Fatalf("round-1 rewards=%+v", rewards)
	}

	removeExtension()
	after := takeStakingAccumulatorSnapshot(t, r, ctx)
	if !reflect.DeepEqual(before, after) {
		t.Fatalf("staking accumulators not restored:\nbefore %+v\nafter  %+v", before, after)
	}
	restored := walkGlobalStakingDeposits(t, r, 50)
	if len(restored) != 1 || restored[0].DepositID != 501 {
		t.Fatalf("restored deposits=%+v", restored)
	}
}

func TestGlobalStakingPagesPropagateCancellation(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	calls := map[string]func() error{
		"CST actions": func() error {
			_, _, err := r.GlobalCstStakingActionsPage(ctx, nil, 1)
			return err
		},
		"RandomWalk actions": func() error {
			_, _, err := r.GlobalRwalkStakingActionsPage(ctx, nil, 1)
			return err
		},
		"CST staked": func() error {
			_, _, err := r.GlobalStakedCstTokensPage(ctx, nil, 1)
			return err
		},
		"RandomWalk staked": func() error {
			_, _, err := r.GlobalStakedRwalkTokensPage(ctx, nil, 1)
			return err
		},
		"deposits": func() error {
			_, _, err := r.GlobalStakingDepositsPage(ctx, nil, 1)
			return err
		},
		"round rewards": func() error {
			_, _, err := r.RoundStakingRewardsPage(ctx, 0, nil, 1)
			return err
		},
		"raffle": func() error {
			_, _, err := r.GlobalStakerRaffleNftWinsPage(ctx, false, nil, 1)
			return err
		},
	}
	for name, call := range calls {
		t.Run(name, func(t *testing.T) {
			if err := call(); !errors.Is(err, context.Canceled) {
				t.Fatalf("error=%v, want context.Canceled", err)
			}
		})
	}
}

func TestGlobalStakingPagesFailOnClosedPool(t *testing.T) {
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	ctx := context.Background()
	st, err := spareStore(ctx)
	if err != nil {
		t.Fatal(err)
	}
	r := NewRepo(st)
	st.Close()
	calls := map[string]func() error{
		"CST actions": func() error {
			_, _, err := r.GlobalCstStakingActionsPage(ctx, nil, 1)
			return err
		},
		"RandomWalk actions": func() error {
			_, _, err := r.GlobalRwalkStakingActionsPage(ctx, nil, 1)
			return err
		},
		"CST staked": func() error {
			_, _, err := r.GlobalStakedCstTokensPage(ctx, nil, 1)
			return err
		},
		"RandomWalk staked": func() error {
			_, _, err := r.GlobalStakedRwalkTokensPage(ctx, nil, 1)
			return err
		},
		"deposits": func() error {
			_, _, err := r.GlobalStakingDepositsPage(ctx, nil, 1)
			return err
		},
		"round rewards": func() error {
			_, _, err := r.RoundStakingRewardsPage(ctx, 0, nil, 1)
			return err
		},
		"raffle": func() error {
			_, _, err := r.GlobalStakerRaffleNftWinsPage(ctx, false, nil, 1)
			return err
		},
	}
	for name, call := range calls {
		t.Run(name, func(t *testing.T) {
			if err := call(); err == nil {
				t.Fatal("closed pool query succeeded")
			}
		})
	}
}

func TestGlobalStakingReadIndexesExist(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	for _, name := range []string{
		"cg_nft_staked_cst_action_idx",
		"cg_nft_unstaked_cst_action_idx",
		"cg_nft_staked_rwalk_action_idx",
		"cg_nft_unstaked_rwalk_action_idx",
		"cg_staking_eth_deposit_round_deposit_idx",
		"cg_staker_deposit_deposit_staker_idx",
		"cg_st_reward_deposit_collected_idx",
		"cg_raffle_nft_prize_staker_pool_evt_idx",
	} {
		var exists bool
		if err := r.pool().QueryRow(ctx,
			`SELECT EXISTS(SELECT 1 FROM pg_indexes WHERE schemaname='public' AND indexname=$1)`,
			name,
		).Scan(&exists); err != nil {
			t.Fatal(err)
		}
		if !exists {
			t.Errorf("index %s does not exist", name)
		}
	}
}
