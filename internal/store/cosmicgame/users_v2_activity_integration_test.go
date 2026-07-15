//go:build integration

package cosmicgame

import (
	"context"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// walkUserTokenPages exhausts an ascending token-keyed page method with the
// given page size and returns every collected token ID, verifying page
// bounds and strict ascending order on the way.
func walkUserTokenPages(
	t *testing.T,
	pageSize int,
	fetch func(after *UserTokenPageCursor, limit int) ([]UserOwnedTokenRecord, bool, error),
) []int64 {
	t.Helper()
	var collected []int64
	var after *UserTokenPageCursor
	for {
		page, hasMore, err := fetch(after, pageSize)
		if err != nil {
			t.Fatalf("page walk: %v", err)
		}
		if len(page) > pageSize {
			t.Fatalf("page length = %d, limit %d", len(page), pageSize)
		}
		for i := range page {
			id := page[i].TokenID
			if len(collected) > 0 && id <= collected[len(collected)-1] {
				t.Fatalf("unordered token IDs: %v then %d", collected, id)
			}
			collected = append(collected, id)
		}
		if !hasMore {
			return collected
		}
		if len(page) == 0 {
			t.Fatal("hasMore without a cursor row")
		}
		after = &UserTokenPageCursor{TokenID: page[len(page)-1].TokenID}
	}
}

func TestUserCosmicSignatureTokensPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// The directory must agree with the legacy owned-token list.
	legacy, err := r.CosmicSignatureTokensByUser(ctx, aidBob, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	legacyIDs := make([]int64, len(legacy))
	for i := range legacy {
		legacyIDs[i] = legacy[i].TokenId
	}
	sort.Slice(legacyIDs, func(i, j int) bool { return legacyIDs[i] < legacyIDs[j] })

	gotIDs := walkUserTokenPages(t, 1,
		func(after *UserTokenPageCursor, limit int) ([]UserOwnedTokenRecord, bool, error) {
			return r.UserCosmicSignatureTokensPage(ctx, aidBob, after, limit)
		})
	if !reflect.DeepEqual(gotIDs, legacyIDs) {
		t.Fatalf("paged tokens = %v, legacy = %v", gotIDs, legacyIDs)
	}

	bob, hasMore, err := r.UserCosmicSignatureTokensPage(ctx, aidBob, nil, 50)
	if err != nil || hasMore || len(bob) != 3 {
		t.Fatalf("bob tokens = %d, more=%v, err=%v", len(bob), hasMore, err)
	}
	// Token 2 was minted to dave (bidder raffle) and later transferred to
	// bob: ownership follows the transfer, provenance stays dave's.
	transferred := bob[0]
	if transferred.TokenID != 2 || transferred.MintSource != MintSourceBidderRaffle ||
		transferred.OwnerAid != aidBob || transferred.Staked ||
		!strings.EqualFold(transferred.WinnerAddr, "0x2400000000000000000000000000000000000024") ||
		transferred.MintTx.EvtLogId != 5026 || transferred.MintRound != 0 {
		t.Fatalf("bob token 2 = %+v", transferred)
	}
	// Token 5 (endurance champion prize) is still inside the staking wallet.
	staked := bob[1]
	if staked.TokenID != 5 || staked.MintSource != MintSourceEnduranceChampion || !staked.Staked {
		t.Fatalf("bob token 5 = %+v", staked)
	}
	// Token 9 is a Cosmic Signature staker raffle mint (staker, not rwalk).
	stakerMint := bob[2]
	if stakerMint.TokenID != 9 || stakerMint.MintSource != MintSourceCosmicSigStaker ||
		stakerMint.MintRound != 2 || stakerMint.TokenName != "" {
		t.Fatalf("bob token 9 = %+v", stakerMint)
	}

	alice, hasMore, err := r.UserCosmicSignatureTokensPage(ctx, aidAlice, nil, 50)
	if err != nil || hasMore || len(alice) != 2 {
		t.Fatalf("alice tokens = %d, more=%v, err=%v", len(alice), hasMore, err)
	}
	// Token 1: the named round-0 main prize; unstaked again after the
	// stake/unstake cycle so membership is gone.
	if alice[0].TokenID != 1 || alice[0].MintSource != MintSourceMainPrize ||
		alice[0].TokenName != "Genesis" || alice[0].Staked ||
		alice[0].Seed == "" {
		t.Fatalf("alice token 1 = %+v", alice[0])
	}
	// Token 6: the chrono-warrior NFT v1 misclassified as a main prize.
	if alice[1].TokenID != 6 || alice[1].MintSource != MintSourceChronoWarriorPrize {
		t.Fatalf("alice token 6 = %+v", alice[1])
	}

	carol, _, err := r.UserCosmicSignatureTokensPage(ctx, aidCarol, nil, 50)
	if err != nil || len(carol) != 2 ||
		carol[0].MintSource != MintSourceRandomWalkStaker ||
		carol[1].MintSource != MintSourceLastCstBidder {
		t.Fatalf("carol tokens = %+v, err=%v", carol, err)
	}

	exhausted, hasMore, err := r.UserCosmicSignatureTokensPage(ctx, aidBob,
		&UserTokenPageCursor{TokenID: 9}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}

	// The external charity receiver is indexed but owns nothing.
	inactive, hasMore, err := r.UserCosmicSignatureTokensPage(ctx, 28, nil, 50)
	if err != nil || hasMore || len(inactive) != 0 || inactive == nil {
		t.Fatalf("inactive-wallet page = len %d nil=%v more=%v err=%v",
			len(inactive), inactive == nil, hasMore, err)
	}
}

func TestUserCosmicSignatureTransfersPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.CosmicSignatureTransfersByUser(ctx, aidDave, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	wantIDs := make([]int64, len(legacy))
	for i := range legacy {
		wantIDs[i] = legacy[i].Tx.EvtLogId
	}
	gotIDs := walkUserEventPages(t, 1,
		func(after *UserEventPageCursor, limit int) ([]UserCosmicSignatureTransferRecord, bool, error) {
			return r.UserCosmicSignatureTransfersPage(ctx, aidDave, after, limit)
		},
		func(record UserCosmicSignatureTransferRecord) int64 { return record.Tx.EvtLogId },
	)
	if !reflect.DeepEqual(gotIDs, wantIDs) {
		t.Fatalf("paged transfers = %v, legacy = %v", gotIDs, wantIDs)
	}

	dave, hasMore, err := r.UserCosmicSignatureTransfersPage(ctx, aidDave, nil, 50)
	if err != nil || hasMore || len(dave) != 3 {
		t.Fatalf("dave transfers = %d, more=%v, err=%v", len(dave), hasMore, err)
	}
	// Newest first: round-1 main-prize mint in, token 2 out to bob, mint in.
	if dave[0].Tx.EvtLogId != 5064 || dave[0].TransferType != 1 ||
		dave[0].Direction != UserTransferIn || dave[0].TokenID != 7 {
		t.Fatalf("dave mint = %+v", dave[0])
	}
	if dave[1].Tx.EvtLogId != 5048 || dave[1].TransferType != 0 ||
		dave[1].Direction != UserTransferOut || dave[1].TokenID != 2 ||
		!strings.EqualFold(dave[1].ToAddr, "0x2200000000000000000000000000000000000022") {
		t.Fatalf("dave transfer out = %+v", dave[1])
	}
	if dave[2].Tx.EvtLogId != 5027 || dave[2].Direction != UserTransferIn {
		t.Fatalf("dave first mint = %+v", dave[2])
	}

	// The same 5048 row appears as an incoming transfer on bob's ledger.
	bob, _, err := r.UserCosmicSignatureTransfersPage(ctx, aidBob, nil, 50)
	if err != nil || len(bob) != 3 {
		t.Fatalf("bob transfers = %d, err=%v", len(bob), err)
	}
	if bob[1].Tx.EvtLogId != 5048 || bob[1].Direction != UserTransferIn ||
		!strings.EqualFold(bob[1].FromAddr, "0x2400000000000000000000000000000000000024") {
		t.Fatalf("bob incoming transfer = %+v", bob[1])
	}

	exhausted, hasMore, err := r.UserCosmicSignatureTransfersPage(ctx, aidDave,
		&UserEventPageCursor{EventLogID: 5027}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}
}

func TestUserCosmicTokenTransfersPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.CosmicTokenTransfersByUser(ctx, aidAlice, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	wantIDs := make([]int64, len(legacy))
	for i := range legacy {
		wantIDs[i] = legacy[i].Tx.EvtLogId
	}
	gotIDs := walkUserEventPages(t, 2,
		func(after *UserEventPageCursor, limit int) ([]UserCosmicTokenTransferRecord, bool, error) {
			return r.UserCosmicTokenTransfersPage(ctx, aidAlice, after, limit)
		},
		func(record UserCosmicTokenTransferRecord) int64 { return record.Tx.EvtLogId },
	)
	if !reflect.DeepEqual(gotIDs, wantIDs) {
		t.Fatalf("paged transfers = %v, legacy = %v", gotIDs, wantIDs)
	}

	alice, hasMore, err := r.UserCosmicTokenTransfersPage(ctx, aidAlice, nil, 50)
	if err != nil || hasMore || len(alice) != 4 {
		t.Fatalf("alice transfers = %d, more=%v, err=%v", len(alice), hasMore, err)
	}
	// Newest first: the 10-CST payment to bob, then three reward mints.
	if alice[0].Tx.EvtLogId != 5049 || alice[0].TransferType != 0 ||
		alice[0].Direction != UserTransferOut ||
		alice[0].AmountWei != "10000000000000000000" {
		t.Fatalf("alice payment = %+v", alice[0])
	}
	for i, wantEvt := range []int64{5043, 5011, 5005} {
		mint := alice[i+1]
		if mint.Tx.EvtLogId != wantEvt || mint.TransferType != 1 ||
			mint.Direction != UserTransferIn ||
			mint.AmountWei != "100000000000000000000" {
			t.Fatalf("alice mint %d = %+v", wantEvt, mint)
		}
	}

	// Carol's ledger ends with her burn.
	carol, _, err := r.UserCosmicTokenTransfersPage(ctx, aidCarol, nil, 50)
	if err != nil || len(carol) != 2 {
		t.Fatalf("carol transfers = %d, err=%v", len(carol), err)
	}
	if carol[0].Tx.EvtLogId != 5050 || carol[0].TransferType != 2 ||
		carol[0].Direction != UserTransferOut ||
		carol[0].AmountWei != "5000000000000000000" {
		t.Fatalf("carol burn = %+v", carol[0])
	}
}

func TestUserCosmicTokenTransfersSelfTransferDeduplicates(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	var balanceBefore, countBefore string
	if err := r.pool().QueryRow(ctx, `SELECT o.cur_balance::TEXT || ':' || s.erc20_num_transfers
		FROM cg_costok_owner o JOIN cg_transfer_stats s ON s.user_aid=o.owner_aid
		WHERE o.owner_aid=$1`, aidBob).Scan(&balanceBefore); err != nil {
		t.Fatal(err)
	}
	countBefore = balanceBefore

	// A synthetic self-transfer qualifies on both UNION branches and must
	// come back exactly once, flagged self.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8020, 129, 1043, 4, 'st000001', 97, '\x00')`); err != nil {
		t.Fatal(err)
	}
	removeExtension := func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM evt_log WHERE id=8020"); err != nil {
			t.Errorf("cleaning synthetic self-transfer: %v", err)
		}
	}
	t.Cleanup(removeExtension)
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_erc20_transfer(evtlog_id, block_num, tx_id, time_stamp, contract_aid, value, from_aid, to_aid, otype)
		VALUES (8020, 129, 1043, TO_TIMESTAMP(1767228500), 4, 7000000000000000000, $1, $1, 0)`, aidBob); err != nil {
		t.Fatal(err)
	}

	bob, _, err := r.UserCosmicTokenTransfersPage(ctx, aidBob, nil, 50)
	if err != nil {
		t.Fatal(err)
	}
	var selfRows int
	for _, transfer := range bob {
		if transfer.Tx.EvtLogId == 8020 {
			selfRows++
			if transfer.Direction != UserTransferSelf ||
				transfer.AmountWei != "7000000000000000000" ||
				transfer.FromAid != aidBob || transfer.ToAid != aidBob {
				t.Fatalf("self transfer = %+v", transfer)
			}
		}
	}
	if selfRows != 1 {
		t.Fatalf("self transfer rows = %d, want exactly 1", selfRows)
	}

	// A one-row page walk must also see it exactly once.
	ids := walkUserEventPages(t, 1,
		func(after *UserEventPageCursor, limit int) ([]UserCosmicTokenTransferRecord, bool, error) {
			return r.UserCosmicTokenTransfersPage(ctx, aidBob, after, limit)
		},
		func(record UserCosmicTokenTransferRecord) int64 { return record.Tx.EvtLogId },
	)
	var walkSelfRows int
	for _, id := range ids {
		if id == 8020 {
			walkSelfRows++
		}
	}
	if walkSelfRows != 1 {
		t.Fatalf("page walk saw the self transfer %d times", walkSelfRows)
	}

	// Removing the extension must restore the trigger-maintained balance
	// and counters exactly (a self transfer is balance-neutral but counted).
	removeExtension()
	var balanceAfter string
	if err := r.pool().QueryRow(ctx, `SELECT o.cur_balance::TEXT || ':' || s.erc20_num_transfers
		FROM cg_costok_owner o JOIN cg_transfer_stats s ON s.user_aid=o.owner_aid
		WHERE o.owner_aid=$1`, aidBob).Scan(&balanceAfter); err != nil {
		t.Fatal(err)
	}
	if balanceAfter != countBefore {
		t.Fatalf("trigger state not restored: before %q after %q", countBefore, balanceAfter)
	}
}

func TestUserMarketingRewardsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.MarketingRewardHistoryByUser(ctx, aidEmma, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	emma, hasMore, err := r.UserMarketingRewardsPage(ctx, aidEmma, nil, 50)
	if err != nil || hasMore || len(emma) != len(legacy) || len(emma) != 1 {
		t.Fatalf("emma rewards = %d, legacy %d, more=%v, err=%v",
			len(emma), len(legacy), hasMore, err)
	}
	reward := emma[0]
	if reward.Tx.EvtLogId != 5017 || reward.MarketerAid != aidEmma ||
		reward.AmountWei != "50000000000000000000" {
		t.Fatalf("emma reward = %+v", reward)
	}

	exhausted, hasMore, err := r.UserMarketingRewardsPage(ctx, aidEmma,
		&UserEventPageCursor{EventLogID: 5017}, 50)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}

	inactive, hasMore, err := r.UserMarketingRewardsPage(ctx, aidAlice, nil, 50)
	if err != nil || hasMore || len(inactive) != 0 || inactive == nil {
		t.Fatalf("inactive-wallet page = len %d nil=%v more=%v err=%v",
			len(inactive), inactive == nil, hasMore, err)
	}
}

func TestUserCosmicTokenSummaryV2(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	alice, err := r.UserCosmicTokenSummaryV2(ctx, aidAlice)
	if err != nil {
		t.Fatal(err)
	}
	want := UserCosmicTokenSummaryRecord{
		BalanceWei:                 "290000000000000000000",
		BiddingRewardsWei:          "500000000000000000000",
		MainPrizesWei:              "100000000000000000000",
		RafflePrizesWei:            "0",
		ChronoWarriorPrizesWei:     "35000000000000000000",
		EnduranceChampionPrizesWei: "0",
		LastCstBidderPrizesWei:     "0",
		MarketingRewardsWei:        "0",
		TotalEarnedWei:             "635000000000000000000",
		ConsumedInBidsWei:          "0",
		NetWei:                     "635000000000000000000",
		TransferCount:              4,
		MintCount:                  3,
		BurnCount:                  0,
	}
	if alice != want {
		t.Fatalf("alice summary = %+v, want %+v", alice, want)
	}

	// The summary balance must agree with the trigger-maintained ledger.
	var triggerBalance string
	if err := r.pool().QueryRow(ctx,
		"SELECT cur_balance::TEXT FROM cg_costok_owner WHERE owner_aid=$1",
		aidAlice).Scan(&triggerBalance); err != nil {
		t.Fatal(err)
	}
	if alice.BalanceWei != triggerBalance {
		t.Fatalf("summary balance %s != trigger balance %s", alice.BalanceWei, triggerBalance)
	}

	// Carol consumed more CST in bids than she earned: net is negative and
	// her last-CST prize is counted (v1's summary missed that source).
	carol, err := r.UserCosmicTokenSummaryV2(ctx, aidCarol)
	if err != nil {
		t.Fatal(err)
	}
	if carol.BalanceWei != "95000000000000000000" ||
		carol.BiddingRewardsWei != "200000000000000000000" ||
		carol.RafflePrizesWei != "30000000000000000000" ||
		carol.LastCstBidderPrizesWei != "40000000000000000000" ||
		carol.TotalEarnedWei != "270000000000000000000" ||
		carol.ConsumedInBidsWei != "410000000000000000000" ||
		carol.NetWei != "-140000000000000000000" ||
		carol.TransferCount != 2 || carol.MintCount != 1 || carol.BurnCount != 1 {
		t.Fatalf("carol summary = %+v", carol)
	}

	// Bob's endurance prize CST is counted.
	bob, err := r.UserCosmicTokenSummaryV2(ctx, aidBob)
	if err != nil || bob.EnduranceChampionPrizesWei != "45000000000000000000" {
		t.Fatalf("bob summary = %+v, err=%v", bob, err)
	}

	// An indexed wallet without any CST activity gets the exact zero shape.
	inactive, err := r.UserCosmicTokenSummaryV2(ctx, 28)
	if err != nil {
		t.Fatal(err)
	}
	zero := UserCosmicTokenSummaryRecord{
		BalanceWei: "0", BiddingRewardsWei: "0", MainPrizesWei: "0",
		RafflePrizesWei: "0", ChronoWarriorPrizesWei: "0",
		EnduranceChampionPrizesWei: "0", LastCstBidderPrizesWei: "0",
		MarketingRewardsWei: "0", TotalEarnedWei: "0",
		ConsumedInBidsWei: "0", NetWei: "0",
	}
	if inactive != zero {
		t.Fatalf("inactive summary = %+v", inactive)
	}
}

func TestUserPendingWinnings(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Alice: both her deposits are unclaimed; the chrono join must split
	// them by winner_index, not by v1's hardcoded index threshold.
	alice, err := r.UserPendingWinnings(ctx, aidAlice)
	if err != nil {
		t.Fatal(err)
	}
	want := UserPendingWinningsRecord{
		RaffleEthWei:           "60000000000000000",
		ChronoWarriorEthWei:    "80000000000000000",
		DonatedNftCount:        0,
		StakingRewardWei:       "0",
		DonatedErc20TokenCount: 0,
	}
	if alice != want {
		t.Fatalf("alice pending = %+v, want %+v", alice, want)
	}

	// Bob withdrew his raffle deposit but still has a pending staking
	// reward on the open deposit.
	bob, err := r.UserPendingWinnings(ctx, aidBob)
	if err != nil {
		t.Fatal(err)
	}
	if bob.RaffleEthWei != "0" || bob.ChronoWarriorEthWei != "0" ||
		bob.StakingRewardWei != "1000000000000000000" ||
		bob.DonatedNftCount != 0 || bob.DonatedErc20TokenCount != 0 {
		t.Fatalf("bob pending = %+v", bob)
	}

	// Emma won round 2 whose donated NFT nobody claimed.
	emma, err := r.UserPendingWinnings(ctx, aidEmma)
	if err != nil {
		t.Fatal(err)
	}
	if emma.DonatedNftCount != 1 || emma.RaffleEthWei != "0" ||
		emma.StakingRewardWei != "0" {
		t.Fatalf("emma pending = %+v", emma)
	}

	// The pending amounts must agree with the deposit ledger's claim flags.
	deposits, _, err := r.UserRaffleEthDepositsPage(ctx, aidAlice, nil, nil, 50)
	if err != nil {
		t.Fatal(err)
	}
	var raffleSum, chronoSum int64
	for _, deposit := range deposits {
		if deposit.Claimed {
			continue
		}
		amount, err := strconv.ParseInt(deposit.EthAmountWei, 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		if deposit.IsChronoWarrior {
			chronoSum += amount
		} else {
			raffleSum += amount
		}
	}
	if raffleSum != 60000000000000000 || chronoSum != 80000000000000000 {
		t.Fatalf("ledger cross-check = raffle %d, chrono %d", raffleSum, chronoSum)
	}

	// An indexed wallet with nothing to claim gets the exact zero shape.
	inactive, err := r.UserPendingWinnings(ctx, 28)
	if err != nil {
		t.Fatal(err)
	}
	zero := UserPendingWinningsRecord{
		RaffleEthWei:        "0",
		ChronoWarriorEthWei: "0",
		StakingRewardWei:    "0",
	}
	if inactive != zero {
		t.Fatalf("inactive pending = %+v", inactive)
	}
}

func TestUserCosmicSignatureTokensRejectAmbiguousMintSource(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	var statsBefore string
	if err := r.pool().QueryRow(ctx,
		"SELECT total_cst_given_in_prizes::TEXT FROM cg_glob_stats").Scan(&statsBefore); err != nil {
		t.Fatal(err)
	}

	// A synthetic round-1 endurance prize pointing at dave's round-1
	// main-prize token makes token 7 match two mint sources; the directory
	// must fail loudly instead of picking one silently.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8030, 120, 1021, 2, 'ep000001', 98, '\x00')`); err != nil {
		t.Fatal(err)
	}
	removeExtension := func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM evt_log WHERE id=8030"); err != nil {
			t.Errorf("cleaning synthetic endurance prize: %v", err)
		}
	}
	t.Cleanup(removeExtension)
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_endurance_prize(evtlog_id, block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num, erc721_token_id, erc20_amount)
		VALUES (8030, 120, 1021, TO_TIMESTAMP(1767227600), 2, 23, 1, 7, 1)`); err != nil {
		t.Fatal(err)
	}

	if _, _, err := r.UserCosmicSignatureTokensPage(ctx, aidDave, nil, 50); err == nil {
		t.Fatal("owned-token page accepted a token with two mint sources")
	} else if !strings.Contains(err.Error(), "mint sources") {
		t.Fatalf("unexpected ambiguity error: %v", err)
	}

	// Removing the extension restores the directory and the accumulators.
	removeExtension()
	dave, _, err := r.UserCosmicSignatureTokensPage(ctx, aidDave, nil, 50)
	if err != nil || len(dave) != 1 || dave[0].MintSource != MintSourceMainPrize {
		t.Fatalf("dave tokens after restore = %+v, err=%v", dave, err)
	}
	var statsAfter string
	if err := r.pool().QueryRow(ctx,
		"SELECT total_cst_given_in_prizes::TEXT FROM cg_glob_stats").Scan(&statsAfter); err != nil {
		t.Fatal(err)
	}
	if statsAfter != statsBefore {
		t.Fatalf("global stats not restored: before %s after %s", statsBefore, statsAfter)
	}
}

func TestUserActivityErrorPaths(t *testing.T) {
	r := repo(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	calls := map[string]func(ctx context.Context) error{
		"UserCosmicSignatureTokensPage": func(ctx context.Context) error {
			_, _, err := r.UserCosmicSignatureTokensPage(ctx, aidAlice, nil, 1)
			return err
		},
		"UserCosmicSignatureTransfersPage": func(ctx context.Context) error {
			_, _, err := r.UserCosmicSignatureTransfersPage(ctx, aidAlice, nil, 1)
			return err
		},
		"UserCosmicTokenTransfersPage": func(ctx context.Context) error {
			_, _, err := r.UserCosmicTokenTransfersPage(ctx, aidAlice, nil, 1)
			return err
		},
		"UserMarketingRewardsPage": func(ctx context.Context) error {
			_, _, err := r.UserMarketingRewardsPage(ctx, aidEmma, nil, 1)
			return err
		},
		"UserCosmicTokenSummaryV2": func(ctx context.Context) error {
			_, err := r.UserCosmicTokenSummaryV2(ctx, aidAlice)
			return err
		},
		"UserPendingWinnings": func(ctx context.Context) error {
			_, err := r.UserPendingWinnings(ctx, aidAlice)
			return err
		},
	}
	for name, call := range calls {
		if err := call(cancelled); err == nil {
			t.Errorf("%s succeeded on a cancelled context", name)
		}
	}
}

func TestUserActivityReadIndexesExist(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// #nosec G101 -- index names and column lists, not credentials.
	wantIndexes := map[string]string{
		"cg_erc721_transfer_from_evt_idx": "(from_aid, evtlog_id desc)",
		"cg_erc721_transfer_to_evt_idx":   "(to_aid, evtlog_id desc)",
		"cg_erc20_transfer_from_evt_idx":  "(from_aid, evtlog_id desc)",
		"cg_erc20_transfer_to_evt_idx":    "(to_aid, evtlog_id desc)",
		"cg_mkt_reward_marketer_evt_idx":  "(marketer_aid, evtlog_id desc)",
		"cg_mint_event_owner_token_idx":   "(cur_owner_aid, token_id)",
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
