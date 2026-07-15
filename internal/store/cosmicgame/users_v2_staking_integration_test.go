//go:build integration

package cosmicgame

import (
	"context"
	"reflect"
	"strings"
	"testing"
)

func TestUserCstStakingActionsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.StakingActionsCstByUser(ctx, aidAlice, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	wantIDs := make([]int64, len(legacy))
	for i := range legacy {
		wantIDs[i] = legacy[i].Tx.EvtLogId
	}
	gotIDs := walkUserEventPages(t, 1,
		func(after *UserEventPageCursor, limit int) ([]UserStakingActionRecord, bool, error) {
			return r.UserCstStakingActionsPage(ctx, aidAlice, after, limit)
		},
		func(record UserStakingActionRecord) int64 { return record.Tx.EvtLogId },
	)
	if !reflect.DeepEqual(gotIDs, wantIDs) {
		t.Fatalf("paged actions = %v, legacy = %v", gotIDs, wantIDs)
	}

	alice, hasMore, err := r.UserCstStakingActionsPage(ctx, aidAlice, nil, 50)
	if err != nil || hasMore || len(alice) != 2 {
		t.Fatalf("alice actions = %d, more=%v, err=%v", len(alice), hasMore, err)
	}
	unstake, stake := alice[0], alice[1]
	if unstake.Tx.EvtLogId != 5056 || unstake.Kind != UserStakingActionUnstake ||
		unstake.ActionID != 1 || unstake.TokenID != 1 ||
		unstake.RewardWei != "1000000000000000000" || unstake.StakerAid != aidAlice {
		t.Fatalf("alice unstake = %+v", unstake)
	}
	if stake.Tx.EvtLogId != 5051 || stake.Kind != UserStakingActionStake ||
		stake.ActionID != 1 || stake.TokenID != 1 ||
		stake.RewardWei != "" || stake.TotalStakedNfts != 1 {
		t.Fatalf("alice stake = %+v", stake)
	}

	bob, hasMore, err := r.UserCstStakingActionsPage(ctx, aidBob, nil, 50)
	if err != nil || hasMore || len(bob) != 1 || bob[0].Kind != UserStakingActionStake ||
		bob[0].Tx.EvtLogId != 5052 || bob[0].TokenID != 5 {
		t.Fatalf("bob actions = %+v, more=%v, err=%v", bob, hasMore, err)
	}

	exhausted, hasMore, err := r.UserCstStakingActionsPage(ctx, aidAlice,
		&UserEventPageCursor{EventLogID: 5051}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}

	inactive, hasMore, err := r.UserCstStakingActionsPage(ctx, aidEmma, nil, 50)
	if err != nil || hasMore || len(inactive) != 0 || inactive == nil {
		t.Fatalf("inactive-wallet page = len %d nil=%v more=%v err=%v",
			len(inactive), inactive == nil, hasMore, err)
	}
}

func TestUserRwalkStakingActionsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.StakingActionsRwalkByUser(ctx, aidCarol, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	wantIDs := make([]int64, len(legacy))
	for i := range legacy {
		wantIDs[i] = legacy[i].Tx.EvtLogId
	}
	gotIDs := walkUserEventPages(t, 1,
		func(after *UserEventPageCursor, limit int) ([]UserStakingActionRecord, bool, error) {
			return r.UserRwalkStakingActionsPage(ctx, aidCarol, after, limit)
		},
		func(record UserStakingActionRecord) int64 { return record.Tx.EvtLogId },
	)
	if !reflect.DeepEqual(gotIDs, wantIDs) {
		t.Fatalf("paged actions = %v, legacy = %v", gotIDs, wantIDs)
	}

	carol, hasMore, err := r.UserRwalkStakingActionsPage(ctx, aidCarol, nil, 50)
	if err != nil || hasMore || len(carol) != 2 {
		t.Fatalf("carol actions = %d, more=%v, err=%v", len(carol), hasMore, err)
	}
	if carol[0].Tx.EvtLogId != 5057 || carol[0].Kind != UserStakingActionUnstake ||
		carol[0].ActionID != 101 || carol[0].TokenID != 10 || carol[0].RewardWei != "" {
		t.Fatalf("carol unstake = %+v", carol[0])
	}
	if carol[1].Tx.EvtLogId != 5053 || carol[1].Kind != UserStakingActionStake ||
		carol[1].RewardWei != "" {
		t.Fatalf("carol stake = %+v", carol[1])
	}

	bob, hasMore, err := r.UserRwalkStakingActionsPage(ctx, aidBob, nil, 50)
	if err != nil || hasMore || len(bob) != 1 ||
		bob[0].Tx.EvtLogId != 5098 || bob[0].Kind != UserStakingActionStake ||
		bob[0].ActionID != 103 || bob[0].TokenID != 13 {
		t.Fatalf("bob actions = %+v, more=%v, err=%v", bob, hasMore, err)
	}
}

func TestUserStakedCstTokensPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.StakedTokensCstByUser(ctx, aidBob)
	if err != nil {
		t.Fatal(err)
	}
	bob, hasMore, err := r.UserStakedCstTokensPage(ctx, aidBob, nil, 50)
	if err != nil || hasMore || len(bob) != len(legacy) || len(bob) != 1 {
		t.Fatalf("bob staked tokens = %d, legacy %d, more=%v, err=%v",
			len(bob), len(legacy), hasMore, err)
	}
	token := bob[0]
	if token.TokenID != 5 || token.ActionID != 2 || token.StakerAid != aidBob ||
		token.StakeTx.EvtLogId != 5052 || token.MintRound != 0 ||
		!strings.HasPrefix(token.Seed, "seed") || token.TokenName != "" {
		t.Fatalf("bob staked token = %+v", token)
	}

	// Alice unstaked her only token: live membership is empty.
	alice, hasMore, err := r.UserStakedCstTokensPage(ctx, aidAlice, nil, 50)
	if err != nil || hasMore || len(alice) != 0 || alice == nil {
		t.Fatalf("alice staked tokens = len %d nil=%v more=%v err=%v",
			len(alice), alice == nil, hasMore, err)
	}

	exhausted, hasMore, err := r.UserStakedCstTokensPage(ctx, aidBob,
		&UserStakingTokenPageCursor{TokenID: 5}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}
}

func TestUserStakedRwalkTokensPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	dave, hasMore, err := r.UserStakedRwalkTokensPage(ctx, aidDave, nil, 50)
	if err != nil || hasMore || len(dave) != 1 ||
		dave[0].TokenID != 11 || dave[0].ActionID != 102 || dave[0].StakeTx.EvtLogId != 5054 {
		t.Fatalf("dave staked tokens = %+v, more=%v, err=%v", dave, hasMore, err)
	}
	bob, hasMore, err := r.UserStakedRwalkTokensPage(ctx, aidBob, nil, 50)
	if err != nil || hasMore || len(bob) != 1 ||
		bob[0].TokenID != 13 || bob[0].ActionID != 103 || bob[0].StakerAid != aidBob {
		t.Fatalf("bob staked tokens = %+v, more=%v, err=%v", bob, hasMore, err)
	}

	// Carol unstaked hers.
	carol, hasMore, err := r.UserStakedRwalkTokensPage(ctx, aidCarol, nil, 50)
	if err != nil || hasMore || len(carol) != 0 || carol == nil {
		t.Fatalf("carol staked tokens = len %d nil=%v more=%v err=%v",
			len(carol), carol == nil, hasMore, err)
	}

	exhausted, hasMore, err := r.UserStakedRwalkTokensPage(ctx, aidBob,
		&UserStakingTokenPageCursor{TokenID: 13}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}
}

func TestUserStakingDepositsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Alice fully collected her share of deposit 501 by unstaking.
	alice, hasMore, err := r.UserStakingDepositsPage(ctx, aidAlice, nil, nil, 50)
	if err != nil || hasMore || len(alice) != 1 {
		t.Fatalf("alice deposits = %d, more=%v, err=%v", len(alice), hasMore, err)
	}
	deposit := alice[0]
	if deposit.DepositID != 501 || deposit.RoundNum != 0 || deposit.Tx.EvtLogId != 5055 ||
		deposit.StakerAid != aidAlice {
		t.Fatalf("alice deposit identity = %+v", deposit)
	}
	if deposit.TotalDepositWei != "2000000000000000000" || deposit.TotalStakedNfts != 2 ||
		deposit.AmountPerTokenWei != "1000000000000000000" {
		t.Fatalf("alice deposit pool amounts = %+v", deposit)
	}
	if deposit.StakedNftCount != 1 || deposit.AmountDepositedWei != "1000000000000000000" ||
		deposit.AmountToClaimWei != "0" ||
		deposit.ClaimedRewardWei != "1000000000000000000" || deposit.PendingRewardWei != "0" ||
		deposit.ClaimedNftCount != 1 || deposit.PendingNftCount != 0 {
		t.Fatalf("alice deposit share = %+v", deposit)
	}

	// Bob is still staked: his share is fully pending.
	bob, hasMore, err := r.UserStakingDepositsPage(ctx, aidBob, nil, nil, 50)
	if err != nil || hasMore || len(bob) != 1 {
		t.Fatalf("bob deposits = %d, more=%v, err=%v", len(bob), hasMore, err)
	}
	if bob[0].PendingRewardWei != "1000000000000000000" || bob[0].ClaimedRewardWei != "0" ||
		bob[0].AmountToClaimWei != "1000000000000000000" ||
		bob[0].ClaimedNftCount != 0 || bob[0].PendingNftCount != 1 {
		t.Fatalf("bob deposit share = %+v", bob[0])
	}

	// The legacy views agree with the unified ledger.
	legacyPending, err := r.StakingRewardsToBeClaimed(ctx, aidBob)
	if err != nil || len(legacyPending) != 1 || legacyPending[0].PendingToClaim != "1000000000000000000" {
		t.Fatalf("legacy pending = %+v, err=%v", legacyPending, err)
	}
	legacyCollected, err := r.StakingRewardsCollected(ctx, aidAlice, 0, 100)
	if err != nil || len(legacyCollected) != 1 ||
		legacyCollected[0].YourCollectedAmount != "1000000000000000000" ||
		legacyCollected[0].YourAmountToClaim != "0" {
		t.Fatalf("legacy collected = %+v, err=%v", legacyCollected, err)
	}

	// The claimed filter splits the two populations.
	claimedOnly := true
	pendingOnly := false
	if page, _, err := r.UserStakingDepositsPage(ctx, aidAlice, &claimedOnly, nil, 50); err != nil || len(page) != 1 {
		t.Fatalf("alice claimed filter = %d rows, err=%v", len(page), err)
	}
	if page, _, err := r.UserStakingDepositsPage(ctx, aidAlice, &pendingOnly, nil, 50); err != nil || len(page) != 0 {
		t.Fatalf("alice pending filter = %d rows, err=%v", len(page), err)
	}
	if page, _, err := r.UserStakingDepositsPage(ctx, aidBob, &claimedOnly, nil, 50); err != nil || len(page) != 0 {
		t.Fatalf("bob claimed filter = %d rows, err=%v", len(page), err)
	}
	if page, _, err := r.UserStakingDepositsPage(ctx, aidBob, &pendingOnly, nil, 50); err != nil || len(page) != 1 {
		t.Fatalf("bob pending filter = %d rows, err=%v", len(page), err)
	}

	// Carol never staked CST: empty ledger.
	carol, hasMore, err := r.UserStakingDepositsPage(ctx, aidCarol, nil, nil, 50)
	if err != nil || hasMore || len(carol) != 0 || carol == nil {
		t.Fatalf("carol deposits = len %d nil=%v more=%v err=%v",
			len(carol), carol == nil, hasMore, err)
	}
}

func TestStakingDepositAndTokenExistence(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	if exists, err := r.StakingDepositExists(ctx, 501); err != nil || !exists {
		t.Fatalf("deposit 501 exists = %v, err=%v", exists, err)
	}
	if exists, err := r.StakingDepositExists(ctx, 999); err != nil || exists {
		t.Fatalf("deposit 999 exists = %v, err=%v", exists, err)
	}
	if exists, err := r.CosmicSignatureTokenExists(ctx, 1); err != nil || !exists {
		t.Fatalf("token 1 exists = %v, err=%v", exists, err)
	}
	if exists, err := r.CosmicSignatureTokenExists(ctx, 999); err != nil || exists {
		t.Fatalf("token 999 exists = %v, err=%v", exists, err)
	}
}

func TestUserStakingDepositRewardsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// No legacy cross-check here: v1's action_ids_by_deposit eligibility
	// window compares stake-action IDs against the deposit's action counter,
	// which the synthetic seed does not reproduce. cg_st_reward itself is
	// the canonical source these pages read.
	alice, hasMore, err := r.UserStakingDepositRewardsPage(ctx, aidAlice, 501, nil, 50)
	if err != nil || hasMore || len(alice) != 1 {
		t.Fatalf("alice rewards = %d, more=%v, err=%v", len(alice), hasMore, err)
	}
	if alice[0].ActionID != 1 || alice[0].TokenID != 1 ||
		alice[0].RewardWei != "1000000000000000000" || !alice[0].Claimed ||
		alice[0].StakerAid != aidAlice {
		t.Fatalf("alice reward = %+v", alice[0])
	}

	bob, hasMore, err := r.UserStakingDepositRewardsPage(ctx, aidBob, 501, nil, 50)
	if err != nil || hasMore || len(bob) != 1 ||
		bob[0].ActionID != 2 || bob[0].TokenID != 5 || bob[0].Claimed {
		t.Fatalf("bob rewards = %+v, more=%v, err=%v", bob, hasMore, err)
	}

	exhausted, hasMore, err := r.UserStakingDepositRewardsPage(ctx, aidBob, 501,
		&UserStakingRewardPageCursor{ActionID: 2}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}

	// Carol has no share in an existing deposit.
	carol, hasMore, err := r.UserStakingDepositRewardsPage(ctx, aidCarol, 501, nil, 50)
	if err != nil || hasMore || len(carol) != 0 || carol == nil {
		t.Fatalf("carol rewards = len %d nil=%v more=%v err=%v",
			len(carol), carol == nil, hasMore, err)
	}
}

func TestUserStakingTokenRewardsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.StakingCstUserTokenRewards(ctx, aidAlice)
	if err != nil || len(legacy) != 1 {
		t.Fatalf("legacy token rewards = %d, err=%v", len(legacy), err)
	}

	alice, hasMore, err := r.UserStakingTokenRewardsPage(ctx, aidAlice, nil, 50)
	if err != nil || hasMore || len(alice) != 1 {
		t.Fatalf("alice token rewards = %d, more=%v, err=%v", len(alice), hasMore, err)
	}
	if alice[0].TokenID != legacy[0].TokenId || alice[0].TokenID != 1 ||
		alice[0].TotalWei != "1000000000000000000" ||
		alice[0].CollectedWei != "1000000000000000000" || alice[0].PendingWei != "0" {
		t.Fatalf("alice token reward = %+v", alice[0])
	}

	bob, hasMore, err := r.UserStakingTokenRewardsPage(ctx, aidBob, nil, 50)
	if err != nil || hasMore || len(bob) != 1 ||
		bob[0].TokenID != 5 || bob[0].CollectedWei != "0" ||
		bob[0].PendingWei != "1000000000000000000" {
		t.Fatalf("bob token reward = %+v, more=%v, err=%v", bob, hasMore, err)
	}

	exhausted, hasMore, err := r.UserStakingTokenRewardsPage(ctx, aidBob,
		&UserStakingTokenPageCursor{TokenID: 5}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}
}

func TestUserStakingTokenRewardDepositsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.StakingCstUserTokenRewardDetails(ctx, aidAlice, 1)
	if err != nil || len(legacy) != 1 {
		t.Fatalf("legacy token details = %d, err=%v", len(legacy), err)
	}

	alice, hasMore, err := r.UserStakingTokenRewardDepositsPage(ctx, aidAlice, 1, nil, 50)
	if err != nil || hasMore || len(alice) != 1 {
		t.Fatalf("alice token deposits = %d, more=%v, err=%v", len(alice), hasMore, err)
	}
	row := alice[0]
	if row.DepositID != 501 || row.RoundNum != 0 || row.Tx.EvtLogId != 5055 ||
		row.RewardWei != "1000000000000000000" || !row.Claimed || row.StakerAid != aidAlice {
		t.Fatalf("alice token deposit = %+v", row)
	}

	// Bob never earned rewards with token 1: empty page on an existing token.
	bob, hasMore, err := r.UserStakingTokenRewardDepositsPage(ctx, aidBob, 1, nil, 50)
	if err != nil || hasMore || len(bob) != 0 || bob == nil {
		t.Fatalf("bob token-1 deposits = len %d nil=%v more=%v err=%v",
			len(bob), bob == nil, hasMore, err)
	}

	exhausted, hasMore, err := r.UserStakingTokenRewardDepositsPage(ctx, aidAlice, 1,
		&UserStakingTokenDepositPageCursor{DepositID: 501}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}
}

// stakingAccumulatorSnapshot captures every accumulator the staking-deposit
// triggers touch, so extension-seed tests can prove exact restoration.
type stakingAccumulatorSnapshot struct {
	StakerRows        []string
	StatsRow          string
	RewardCount       int64
	StakerDepositRows int64
	AllocationPrizes  int64
}

func takeStakingAccumulatorSnapshot(t *testing.T, r *Repo, ctx context.Context) stakingAccumulatorSnapshot {
	t.Helper()
	var snapshot stakingAccumulatorSnapshot
	rows, err := r.pool().Query(ctx, `SELECT staker_aid || ':' || total_reward::TEXT || ':' || unclaimed_reward::TEXT
		FROM cg_staker_cst ORDER BY staker_aid`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var row string
		if err := rows.Scan(&row); err != nil {
			t.Fatal(err)
		}
		snapshot.StakerRows = append(snapshot.StakerRows, row)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
	if err := r.pool().QueryRow(ctx, `SELECT total_reward_amount::TEXT || ':' || total_unclaimed_reward::TEXT
			|| ':' || num_deposits || ':' || total_modulo::TEXT
		FROM cg_stake_stats_cst`).Scan(&snapshot.StatsRow); err != nil {
		t.Fatal(err)
	}
	if err := r.pool().QueryRow(ctx,
		"SELECT COUNT(*) FROM cg_st_reward").Scan(&snapshot.RewardCount); err != nil {
		t.Fatal(err)
	}
	if err := r.pool().QueryRow(ctx,
		"SELECT COUNT(*) FROM cg_staker_deposit").Scan(&snapshot.StakerDepositRows); err != nil {
		t.Fatal(err)
	}
	if err := r.pool().QueryRow(ctx,
		"SELECT COUNT(*) FROM cg_prize WHERE ptype=15").Scan(&snapshot.AllocationPrizes); err != nil {
		t.Fatal(err)
	}
	return snapshot
}

func TestUserStakingMultiDepositPagination(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	before := takeStakingAccumulatorSnapshot(t, r, ctx)

	// A second staking deposit (3 ETH over bob's single staked token) lands
	// in round 1. The insert trigger fans it out into cg_st_reward and
	// cg_staker_deposit; deleting the event must reverse everything.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8010, 117, 1018, 9, 'sd000002', 96, '\x00')`); err != nil {
		t.Fatal(err)
	}
	removeExtension := func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM evt_log WHERE id=8010"); err != nil {
			t.Errorf("cleaning synthetic staking deposit: %v", err)
		}
	}
	t.Cleanup(removeExtension)
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_staking_eth_deposit(evtlog_id, block_num, tx_id, time_stamp, contract_aid,
			round_num, deposit_time, deposit_id, num_staked_nfts, deposit_amount, amount_per_token, modulo)
		VALUES (8010, 117, 1018, TO_TIMESTAMP(1767227300), 9, 1, TO_TIMESTAMP(1767227300), 502, 1,
			3000000000000000000, 3000000000000000000, 0)`); err != nil {
		t.Fatal(err)
	}

	// Bob's ledger now pages newest deposit first across two pages.
	firstPage, hasMore, err := r.UserStakingDepositsPage(ctx, aidBob, nil, nil, 1)
	if err != nil || !hasMore || len(firstPage) != 1 || firstPage[0].DepositID != 502 {
		t.Fatalf("first deposit page = %+v, more=%v, err=%v", firstPage, hasMore, err)
	}
	if firstPage[0].PendingRewardWei != "3000000000000000000" ||
		firstPage[0].AmountPerTokenWei != "3000000000000000000" ||
		firstPage[0].RoundNum != 1 || firstPage[0].TotalStakedNfts != 1 {
		t.Fatalf("second deposit share = %+v", firstPage[0])
	}
	secondPage, hasMore, err := r.UserStakingDepositsPage(ctx, aidBob, nil,
		&UserStakingDepositPageCursor{DepositID: 502}, 1)
	if err != nil || hasMore || len(secondPage) != 1 || secondPage[0].DepositID != 501 {
		t.Fatalf("second deposit page = %+v, more=%v, err=%v", secondPage, hasMore, err)
	}

	// Token 5 accumulates rewards from both deposits...
	tokenRewards, _, err := r.UserStakingTokenRewardsPage(ctx, aidBob, nil, 50)
	if err != nil || len(tokenRewards) != 1 ||
		tokenRewards[0].TotalWei != "4000000000000000000" ||
		tokenRewards[0].PendingWei != "4000000000000000000" {
		t.Fatalf("bob token rewards = %+v, err=%v", tokenRewards, err)
	}
	// ...and its per-deposit breakdown pages in ascending deposit order.
	firstDetail, hasMore, err := r.UserStakingTokenRewardDepositsPage(ctx, aidBob, 5, nil, 1)
	if err != nil || !hasMore || len(firstDetail) != 1 || firstDetail[0].DepositID != 501 {
		t.Fatalf("first detail page = %+v, more=%v, err=%v", firstDetail, hasMore, err)
	}
	secondDetail, hasMore, err := r.UserStakingTokenRewardDepositsPage(ctx, aidBob, 5,
		&UserStakingTokenDepositPageCursor{DepositID: 501}, 1)
	if err != nil || hasMore || len(secondDetail) != 1 || secondDetail[0].DepositID != 502 ||
		secondDetail[0].RewardWei != "3000000000000000000" || secondDetail[0].Claimed {
		t.Fatalf("second detail page = %+v, more=%v, err=%v", secondDetail, hasMore, err)
	}

	// Alice took no part in the second deposit.
	alice, _, err := r.UserStakingDepositsPage(ctx, aidAlice, nil, nil, 50)
	if err != nil || len(alice) != 1 || alice[0].DepositID != 501 {
		t.Fatalf("alice deposits after extension = %+v, err=%v", alice, err)
	}

	// The legacy per-deposit reward tree sees the same two-deposit ledger.
	legacyTree, err := r.StakingCstUserDepositRewards(ctx, aidBob)
	if err != nil || len(legacyTree) != 2 {
		t.Fatalf("legacy deposit tree = %d deposits, err=%v", len(legacyTree), err)
	}
	if legacyTree[0].DepositId != 501 || legacyTree[1].DepositId != 502 ||
		legacyTree[1].YourClaimableAmountEth != 3 || legacyTree[1].FullyClaimed {
		t.Fatalf("legacy deposit tree = %+v", legacyTree)
	}

	// Removing the extension must restore every trigger-maintained
	// accumulator exactly.
	removeExtension()
	after := takeStakingAccumulatorSnapshot(t, r, ctx)
	if !reflect.DeepEqual(before, after) {
		t.Fatalf("staking accumulators not restored:\nbefore %+v\nafter  %+v", before, after)
	}
	restored, hasMore, err := r.UserStakingDepositsPage(ctx, aidBob, nil, nil, 50)
	if err != nil || hasMore || len(restored) != 1 || restored[0].DepositID != 501 {
		t.Fatalf("bob deposits after restore = %+v, more=%v, err=%v", restored, hasMore, err)
	}
}

func TestUserStakingErrorPaths(t *testing.T) {
	r := repo(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	calls := map[string]func(ctx context.Context) error{
		"UserCstStakingActionsPage": func(ctx context.Context) error {
			_, _, err := r.UserCstStakingActionsPage(ctx, aidAlice, nil, 1)
			return err
		},
		"UserRwalkStakingActionsPage": func(ctx context.Context) error {
			_, _, err := r.UserRwalkStakingActionsPage(ctx, aidCarol, nil, 1)
			return err
		},
		"UserStakedCstTokensPage": func(ctx context.Context) error {
			_, _, err := r.UserStakedCstTokensPage(ctx, aidBob, nil, 1)
			return err
		},
		"UserStakedRwalkTokensPage": func(ctx context.Context) error {
			_, _, err := r.UserStakedRwalkTokensPage(ctx, aidBob, nil, 1)
			return err
		},
		"UserStakingDepositsPage": func(ctx context.Context) error {
			_, _, err := r.UserStakingDepositsPage(ctx, aidBob, nil, nil, 1)
			return err
		},
		"StakingDepositExists": func(ctx context.Context) error {
			_, err := r.StakingDepositExists(ctx, 501)
			return err
		},
		"UserStakingDepositRewardsPage": func(ctx context.Context) error {
			_, _, err := r.UserStakingDepositRewardsPage(ctx, aidBob, 501, nil, 1)
			return err
		},
		"UserStakingTokenRewardsPage": func(ctx context.Context) error {
			_, _, err := r.UserStakingTokenRewardsPage(ctx, aidBob, nil, 1)
			return err
		},
		"CosmicSignatureTokenExists": func(ctx context.Context) error {
			_, err := r.CosmicSignatureTokenExists(ctx, 1)
			return err
		},
		"UserStakingTokenRewardDepositsPage": func(ctx context.Context) error {
			_, _, err := r.UserStakingTokenRewardDepositsPage(ctx, aidBob, 5, nil, 1)
			return err
		},
	}
	for name, call := range calls {
		if err := call(cancelled); err == nil {
			t.Errorf("%s succeeded on a cancelled context", name)
		}
	}
}

func TestUserStakingReadIndexesExist(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// #nosec G101 -- index names and column lists, not credentials.
	wantIndexes := map[string]string{
		"cg_nft_staked_cst_staker_evt_idx":       "(staker_aid, evtlog_id desc)",
		"cg_nft_unstaked_cst_staker_evt_idx":     "(staker_aid, evtlog_id desc)",
		"cg_nft_staked_rwalk_staker_evt_idx":     "(staker_aid, evtlog_id desc)",
		"cg_nft_unstaked_rwalk_staker_evt_idx":   "(staker_aid, evtlog_id desc)",
		"cg_staked_token_cst_staker_token_idx":   "(staker_aid, token_id)",
		"cg_staked_token_rwalk_staker_token_idx": "(staker_aid, token_id)",
		"cg_st_reward_staker_deposit_action_idx": "(staker_aid, deposit_id, action_id)",
		"cg_st_reward_staker_token_deposit_idx":  "(staker_aid, token_id, deposit_id)",
		"cg_staking_eth_deposit_deposit_idx":     "(deposit_id)",
	}
	for indexName, wantColumns := range wantIndexes {
		var definition string
		err := r.pool().QueryRow(ctx, `SELECT indexdef FROM pg_indexes
			WHERE schemaname='public' AND indexname=$1`, indexName).Scan(&definition)
		if err != nil {
			t.Errorf("read index %s: %v", indexName, err)
			continue
		}
		normalized := strings.ToLower(strings.Join(strings.Fields(definition), " "))
		if !strings.Contains(normalized, wantColumns) {
			t.Errorf("index %s definition = %s, want columns %s", indexName, definition, wantColumns)
		}
	}
}
