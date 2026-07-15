//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
)

// walkUserEventPages exhausts a newest-first event-keyed page method with
// the given page size and returns every collected event-log ID, verifying
// page bounds and strict descending order on the way.
func walkUserEventPages[T any](
	t *testing.T,
	pageSize int,
	fetch func(after *UserEventPageCursor, limit int) ([]T, bool, error),
	eventLogID func(T) int64,
) []int64 {
	t.Helper()
	var collected []int64
	var after *UserEventPageCursor
	for {
		page, hasMore, err := fetch(after, pageSize)
		if err != nil {
			t.Fatalf("page walk: %v", err)
		}
		if len(page) > pageSize {
			t.Fatalf("page length = %d, limit %d", len(page), pageSize)
		}
		for i := range page {
			id := eventLogID(page[i])
			if len(collected) > 0 && id >= collected[len(collected)-1] {
				t.Fatalf("unordered event IDs: %v then %d", collected, id)
			}
			collected = append(collected, id)
		}
		if !hasMore {
			return collected
		}
		if len(page) == 0 {
			t.Fatal("hasMore without a cursor row")
		}
		after = &UserEventPageCursor{EventLogID: collected[len(collected)-1]}
	}
}

func TestUserPrizesPageMatchesLegacyHistory(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	for name, aid := range map[string]int64{
		"alice": aidAlice,
		"bob":   aidBob,
		"carol": aidCarol,
	} {
		t.Run(name, func(t *testing.T) {
			legacy, err := r.PrizeHistoryByUser(ctx, aid, 0, 1000)
			if err != nil {
				t.Fatal(err)
			}
			type prizeKey struct {
				round, ptype, winnerIndex int64
			}
			want := map[prizeKey]string{}
			for _, row := range legacy {
				if row.IsTimeoutClaim {
					continue
				}
				want[prizeKey{row.RoundNum, row.RecordType, row.WinnerIndex}] = row.Amount
			}

			var got []prizeKey
			var after *UserPrizePageCursor
			for {
				page, hasMore, err := r.UserPrizesPage(ctx, aid, after, 2)
				if err != nil {
					t.Fatalf("UserPrizesPage: %v", err)
				}
				if len(page) > 2 {
					t.Fatalf("page length = %d", len(page))
				}
				for i := range page {
					row := page[i]
					if row.WinnerAid != aid {
						t.Fatalf("out-of-scope winner aid %d", row.WinnerAid)
					}
					key := prizeKey{row.RoundNum, row.RecordType, row.WinnerIndex}
					if amount, ok := want[key]; !ok || amount != row.Amount {
						t.Fatalf("prize %+v amount %q not in legacy set %v", key, row.Amount, want)
					}
					if len(got) > 0 {
						previous := got[len(got)-1]
						ordered := key.round < previous.round ||
							(key.round == previous.round &&
								(key.ptype > previous.ptype ||
									(key.ptype == previous.ptype && key.winnerIndex > previous.winnerIndex)))
						if !ordered {
							t.Fatalf("unordered prizes: %+v then %+v", previous, key)
						}
					}
					got = append(got, key)
				}
				if !hasMore {
					break
				}
				if len(page) == 0 {
					t.Fatal("hasMore without a cursor row")
				}
				last := page[len(page)-1]
				after = &UserPrizePageCursor{
					Round:       last.RoundNum,
					PrizeType:   last.RecordType,
					WinnerIndex: last.WinnerIndex,
				}
			}
			if len(got) != len(want) {
				t.Fatalf("paged %d prizes, legacy has %d", len(got), len(want))
			}
		})
	}
}

func TestUserPrizesPageBoundaries(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	full, hasMore, err := r.UserPrizesPage(ctx, aidAlice, nil, 200)
	if err != nil || hasMore {
		t.Fatalf("full page: hasMore=%v err=%v", hasMore, err)
	}
	if len(full) < 5 {
		t.Fatalf("alice prize count = %d, fixture expects at least main(3)+chrono(3)+raffle(1)", len(full))
	}
	first, hasMore, err := r.UserPrizesPage(ctx, aidAlice, nil, 1)
	if err != nil || !hasMore || len(first) != 1 {
		t.Fatalf("first page = %d rows, more=%v, err=%v", len(first), hasMore, err)
	}
	if !reflect.DeepEqual(first[0], full[0]) {
		t.Fatalf("first page row differs from full list head")
	}
	last := full[len(full)-1]
	exhausted, hasMore, err := r.UserPrizesPage(ctx, aidAlice, &UserPrizePageCursor{
		Round:       last.RoundNum,
		PrizeType:   last.RecordType,
		WinnerIndex: last.WinnerIndex,
	}, 5)
	if err != nil || hasMore || len(exhausted) != 0 {
		t.Fatalf("exhausted page = %d rows, more=%v, err=%v", len(exhausted), hasMore, err)
	}

	empty, hasMore, err := r.UserPrizesPage(ctx, aidPrizesWallet, nil, 5)
	if err != nil || hasMore || len(empty) != 0 || empty == nil {
		t.Fatalf("inactive-wallet page = len %d nil=%v more=%v err=%v",
			len(empty), empty == nil, hasMore, err)
	}
}

func TestUserRaffleEthDepositsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacy, err := r.EthDepositsByUser(ctx, aidAlice)
	if err != nil {
		t.Fatal(err)
	}
	wantIDs := make([]int64, len(legacy))
	for i := range legacy {
		wantIDs[i] = legacy[i].Tx.EvtLogId
	}
	gotIDs := walkUserEventPages(t, 1,
		func(after *UserEventPageCursor, limit int) ([]UserRaffleEthDepositRecord, bool, error) {
			return r.UserRaffleEthDepositsPage(ctx, aidAlice, nil, after, limit)
		},
		func(record UserRaffleEthDepositRecord) int64 { return record.Tx.EvtLogId },
	)
	if !reflect.DeepEqual(gotIDs, wantIDs) {
		t.Fatalf("paged deposits = %v, legacy = %v", gotIDs, wantIDs)
	}

	all, hasMore, err := r.UserRaffleEthDepositsPage(ctx, aidAlice, nil, nil, 50)
	if err != nil || hasMore {
		t.Fatalf("alice deposits: more=%v err=%v", hasMore, err)
	}
	if len(all) != 2 {
		t.Fatalf("alice deposit count = %d, want 2", len(all))
	}
	byEvent := map[int64]UserRaffleEthDepositRecord{}
	for _, record := range all {
		if record.WinnerAid != aidAlice || record.Claimed || record.Withdrawal != nil {
			t.Fatalf("alice deposit %+v: want unclaimed without withdrawal", record)
		}
		byEvent[record.Tx.EvtLogId] = record
	}
	if !byEvent[5040].IsChronoWarrior || byEvent[5040].EthAmountWei != "80000000000000000" {
		t.Fatalf("chrono deposit = %+v", byEvent[5040])
	}
	if byEvent[5066].IsChronoWarrior || byEvent[5066].EthAmountWei != "60000000000000000" {
		t.Fatalf("raffle deposit = %+v", byEvent[5066])
	}

	claimedOnly := true
	unclaimedOnly := false
	bobClaimed, hasMore, err := r.UserRaffleEthDepositsPage(ctx, aidBob, &claimedOnly, nil, 50)
	if err != nil || hasMore || len(bobClaimed) != 1 {
		t.Fatalf("bob claimed deposits = %d, more=%v, err=%v", len(bobClaimed), hasMore, err)
	}
	deposit := bobClaimed[0]
	if !deposit.Claimed || deposit.IsChronoWarrior || deposit.Withdrawal == nil {
		t.Fatalf("bob claimed deposit = %+v", deposit)
	}
	if deposit.Withdrawal.EventLogID != 5044 ||
		!strings.EqualFold(deposit.Withdrawal.BeneficiaryAddr, "0x2200000000000000000000000000000000000022") {
		t.Fatalf("bob withdrawal = %+v", deposit.Withdrawal)
	}
	bobUnclaimed, hasMore, err := r.UserRaffleEthDepositsPage(ctx, aidBob, &unclaimedOnly, nil, 50)
	if err != nil || hasMore || len(bobUnclaimed) != 0 {
		t.Fatalf("bob unclaimed deposits = %d, more=%v, err=%v", len(bobUnclaimed), hasMore, err)
	}
	aliceUnclaimed, _, err := r.UserRaffleEthDepositsPage(ctx, aidAlice, &unclaimedOnly, nil, 50)
	if err != nil || len(aliceUnclaimed) != 2 {
		t.Fatalf("alice unclaimed deposits = %d, err=%v", len(aliceUnclaimed), err)
	}
}

func TestUserRaffleEthDepositsThirdPartyWithdrawal(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// A timeout withdrawal pays carol's round-0 deposit out to dave.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8003, 113, 1014, 7, 'tw000001', 90, '\x00')`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cleanupCtx := context.Background()
		if _, err := r.pool().Exec(cleanupCtx, "DELETE FROM evt_log WHERE id=8003"); err != nil {
			t.Errorf("cleaning synthetic withdrawal event: %v", err)
		}
		if _, err := r.pool().Exec(cleanupCtx, `UPDATE cg_prize_deposit
			SET claimed='F', withdrawal_id=0 WHERE evtlog_id=5024`); err != nil {
			t.Errorf("restoring carol's deposit: %v", err)
		}
	})
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_prize_withdrawal(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, winner_aid, beneficiary_aid, amount)
		VALUES (8003, 113, 1014, TO_TIMESTAMP(1767226900), 7, 0, $1, $2, 50000000000000000)`,
		aidCarol, aidDave); err != nil {
		t.Fatal(err)
	}

	deposits, _, err := r.UserRaffleEthDepositsPage(ctx, aidCarol, nil, nil, 50)
	if err != nil {
		t.Fatal(err)
	}
	var claimed *UserRaffleEthDepositRecord
	for i := range deposits {
		if deposits[i].Tx.EvtLogId == 5024 {
			claimed = &deposits[i]
		}
	}
	if claimed == nil || !claimed.Claimed || claimed.Withdrawal == nil {
		t.Fatalf("carol's deposit after timeout withdrawal = %+v", claimed)
	}
	if claimed.Withdrawal.EventLogID != 8003 ||
		!strings.EqualFold(claimed.Withdrawal.BeneficiaryAddr, "0x2400000000000000000000000000000000000024") {
		t.Fatalf("timeout withdrawal = %+v", claimed.Withdrawal)
	}
}

func TestUserRaffleNftWinsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	type wantWin struct {
		event   int64
		rwalk   bool
		staker  bool
		tokenID int64
		cstWei  string
	}
	for name, testCase := range map[string]struct {
		aid  int64
		want []wantWin
	}{
		"dave bidder pool": {aid: aidDave, want: []wantWin{
			{5025, false, false, 2, "30000000000000000000"},
		}},
		"carol randomwalk staker pool": {aid: aidCarol, want: []wantWin{
			{5028, true, true, 3, "30000000000000000000"},
		}},
		"bob cosmic signature staker pool": {aid: aidBob, want: []wantWin{
			{5099, false, true, 9, "32000000000000000000"},
		}},
	} {
		t.Run(name, func(t *testing.T) {
			// The legacy view never selected the CST amount; v2 adds it, so
			// only the event identities are comparable across generations.
			legacy, err := r.RaffleNFTWinningsByUser(ctx, testCase.aid)
			if err != nil {
				t.Fatal(err)
			}
			page, hasMore, err := r.UserRaffleNftWinsPage(ctx, testCase.aid, nil, 50)
			if err != nil || hasMore {
				t.Fatalf("page: more=%v err=%v", hasMore, err)
			}
			if len(page) != len(legacy) || len(page) != len(testCase.want) {
				t.Fatalf("win count = %d, legacy %d, want %d", len(page), len(legacy), len(testCase.want))
			}
			for i, want := range testCase.want {
				got := page[i]
				if got.Tx.EvtLogId != legacy[i].Tx.EvtLogId {
					t.Fatalf("win[%d] event = %d, legacy %d", i, got.Tx.EvtLogId, legacy[i].Tx.EvtLogId)
				}
				if got.Tx.EvtLogId != want.event || got.IsRWalk != want.rwalk ||
					got.IsStaker != want.staker || got.TokenID != want.tokenID ||
					got.CstAmountWei != want.cstWei || got.WinnerAid != testCase.aid {
					t.Fatalf("win[%d] = %+v, want %+v", i, got, want)
				}
			}
		})
	}
}

func TestUserEventPagesHonorContinuationCursors(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Each cursor sits just above the wallet's only event, so the
	// continuation query arm must return exactly that event.
	t.Run("raffle nft wins", func(t *testing.T) {
		page, hasMore, err := r.UserRaffleNftWinsPage(ctx, aidDave,
			&UserEventPageCursor{EventLogID: 5026}, 5)
		if err != nil || hasMore || len(page) != 1 || page[0].Tx.EvtLogId != 5025 {
			t.Fatalf("continuation page = %+v, more=%v, err=%v", page, hasMore, err)
		}
	})
	t.Run("erc20 donations", func(t *testing.T) {
		page, hasMore, err := r.ERC20DonationsByUserPage(ctx, aidAlice,
			&UserEventPageCursor{EventLogID: 5016}, 5)
		if err != nil || hasMore || len(page) != 1 || page[0].Tx.EvtLogId != 5015 {
			t.Fatalf("continuation page = %+v, more=%v, err=%v", page, hasMore, err)
		}
	})
	t.Run("nft donations", func(t *testing.T) {
		page, hasMore, err := r.NFTDonationsByUserPage(ctx, aidBob,
			&UserEventPageCursor{EventLogID: 5017}, 5)
		if err != nil || hasMore || len(page) != 1 || page[0].Tx.EvtLogId != 5016 {
			t.Fatalf("continuation page = %+v, more=%v, err=%v", page, hasMore, err)
		}
	})
	t.Run("donated nfts", func(t *testing.T) {
		page, hasMore, err := r.UserDonatedNftsPage(ctx, aidAlice, nil,
			&UserEventPageCursor{EventLogID: 5017}, 5)
		if err != nil || hasMore || len(page) != 1 || page[0].Tx.EvtLogId != 5016 {
			t.Fatalf("continuation page = %+v, more=%v, err=%v", page, hasMore, err)
		}
		exhausted, hasMore, err := r.UserDonatedNftsPage(ctx, aidAlice, nil,
			&UserEventPageCursor{EventLogID: 5016}, 5)
		if err != nil || hasMore || len(exhausted) != 0 {
			t.Fatalf("exhausted page = %+v, more=%v, err=%v", exhausted, hasMore, err)
		}
	})
}

func TestUserDonationPagesMatchLegacyLists(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	t.Run("eth donations", func(t *testing.T) {
		for name, aid := range map[string]int64{"dave plain": aidDave, "emma with info": aidEmma} {
			legacy, err := r.EthDonationsByUser(ctx, aid)
			if err != nil {
				t.Fatal(err)
			}
			wantIDs := make([]int64, len(legacy))
			for i := range legacy {
				wantIDs[i] = legacy[i].Tx.EvtLogId
			}
			gotIDs := walkUserEventPages(t, 1,
				func(after *UserEventPageCursor, limit int) ([]RoundEthDonationRecord, bool, error) {
					return r.EthDonationsByUserPage(ctx, aid, after, limit)
				},
				func(record RoundEthDonationRecord) int64 { return record.Tx.EvtLogId },
			)
			if !reflect.DeepEqual(gotIDs, wantIDs) {
				t.Fatalf("%s: paged = %v, legacy = %v", name, gotIDs, wantIDs)
			}
		}
		davePage, _, err := r.EthDonationsByUserPage(ctx, aidDave, nil, 50)
		if err != nil || len(davePage) != 1 || davePage[0].Kind != RoundEthDonationPlain {
			t.Fatalf("dave donations = %+v, err=%v", davePage, err)
		}
		emmaPage, _, err := r.EthDonationsByUserPage(ctx, aidEmma, nil, 50)
		if err != nil || len(emmaPage) != 1 || emmaPage[0].Kind != RoundEthDonationWithInfo ||
			emmaPage[0].ContractRecordID == nil || emmaPage[0].Data == nil {
			t.Fatalf("emma donations = %+v, err=%v", emmaPage, err)
		}
	})

	t.Run("erc20 donations", func(t *testing.T) {
		legacy, err := r.ERC20DonationsByUser(ctx, aidAlice)
		if err != nil {
			t.Fatal(err)
		}
		page, hasMore, err := r.ERC20DonationsByUserPage(ctx, aidAlice, nil, 50)
		if err != nil || hasMore || len(page) != len(legacy) || len(page) != 1 {
			t.Fatalf("alice erc20 donations = %d, legacy %d, more=%v, err=%v",
				len(page), len(legacy), hasMore, err)
		}
		if page[0].Tx.EvtLogId != 5015 || page[0].AmountBaseUnits != "500000000000000000000" {
			t.Fatalf("alice erc20 donation = %+v", page[0])
		}
	})

	t.Run("nft donations", func(t *testing.T) {
		for name, testCase := range map[string]struct {
			aid    int64
			events []int64
		}{
			"bob":  {aid: aidBob, events: []int64{5016}},
			"emma": {aid: aidEmma, events: []int64{5102}},
		} {
			legacy, err := r.NFTDonationsByUser(ctx, testCase.aid)
			if err != nil {
				t.Fatal(err)
			}
			gotIDs := walkUserEventPages(t, 1,
				func(after *UserEventPageCursor, limit int) ([]RoundNFTDonationRecord, bool, error) {
					return r.NFTDonationsByUserPage(ctx, testCase.aid, after, limit)
				},
				func(record RoundNFTDonationRecord) int64 { return record.Tx.EvtLogId },
			)
			if len(gotIDs) != len(legacy) || !reflect.DeepEqual(gotIDs, testCase.events) {
				t.Fatalf("%s nft donations = %v, want %v (legacy %d)",
					name, gotIDs, testCase.events, len(legacy))
			}
		}
	})
}

func TestEthDonationsByUserPageMergesBranchesInOrder(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Give dave one donation of each kind around his existing plain event
	// (5012) so the page must interleave both event tables.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp) VALUES
		(8010, 113, 1014, 2, 'ed000001', 91, '\x00'),
		(8011, 113, 1014, 2, 'ew000001', 92, '\x00')`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cleanupCtx := context.Background()
		if _, err := r.pool().Exec(cleanupCtx,
			"DELETE FROM cg_donation_json WHERE record_id=901"); err != nil {
			t.Errorf("cleaning synthetic donation json: %v", err)
		}
		if _, err := r.pool().Exec(cleanupCtx,
			"DELETE FROM evt_log WHERE id IN (8010, 8011)"); err != nil {
			t.Errorf("cleaning synthetic donation events: %v", err)
		}
	})
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_eth_donated(evtlog_id, block_num, tx_id, time_stamp, contract_aid, donor_aid, round_num, amount)
		VALUES (8010, 113, 1014, TO_TIMESTAMP(1767226900), 2, $1, 2, 111000000000000000)`,
		aidDave); err != nil {
		t.Fatal(err)
	}
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_eth_donated_wi(evtlog_id, block_num, tx_id, time_stamp, contract_aid, donor_aid, round_num, record_id, amount)
		VALUES (8011, 113, 1014, TO_TIMESTAMP(1767226900), 2, $1, 2, 901, 222000000000000000)`,
		aidDave); err != nil {
		t.Fatal(err)
	}
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_donation_json(record_id, data)
		VALUES (901, '{"title":"synthetic"}')`); err != nil {
		t.Fatal(err)
	}

	gotIDs := walkUserEventPages(t, 1,
		func(after *UserEventPageCursor, limit int) ([]RoundEthDonationRecord, bool, error) {
			return r.EthDonationsByUserPage(ctx, aidDave, after, limit)
		},
		func(record RoundEthDonationRecord) int64 { return record.Tx.EvtLogId },
	)
	if !reflect.DeepEqual(gotIDs, []int64{8011, 8010, 5012}) {
		t.Fatalf("merged donation order = %v, want [8011 8010 5012]", gotIDs)
	}

	middle, hasMore, err := r.EthDonationsByUserPage(ctx, aidDave,
		&UserEventPageCursor{EventLogID: 8011}, 1)
	if err != nil || !hasMore || len(middle) != 1 || middle[0].Tx.EvtLogId != 8010 ||
		middle[0].Kind != RoundEthDonationPlain {
		t.Fatalf("middle page = %+v, more=%v, err=%v", middle, hasMore, err)
	}
}

func TestUserDonatedNftsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Alice won round 0 and claimed its donated NFT.
	alice, hasMore, err := r.UserDonatedNftsPage(ctx, aidAlice, nil, nil, 50)
	if err != nil || hasMore || len(alice) != 1 {
		t.Fatalf("alice donated nfts = %d, more=%v, err=%v", len(alice), hasMore, err)
	}
	claimedRow := alice[0]
	if claimedRow.Tx.EvtLogId != 5016 || !claimedRow.Claimed || claimedRow.Claim == nil ||
		claimedRow.Claim.EventLogID != 5045 || claimedRow.Claim.ClaimerAid != aidAlice ||
		claimedRow.RoundWinnerAid != aidAlice || claimedRow.DonationIndex != 0 {
		t.Fatalf("alice claimed donated nft = %+v claim=%+v", claimedRow, claimedRow.Claim)
	}

	// Emma won round 2; its donation is still unclaimed.
	emma, _, err := r.UserDonatedNftsPage(ctx, aidEmma, nil, nil, 50)
	if err != nil || len(emma) != 1 {
		t.Fatalf("emma donated nfts = %d, err=%v", len(emma), err)
	}
	unclaimedRow := emma[0]
	if unclaimedRow.Tx.EvtLogId != 5102 || unclaimedRow.Claimed || unclaimedRow.Claim != nil ||
		unclaimedRow.RoundWinnerAid != aidEmma || unclaimedRow.DonationIndex != 1 {
		t.Fatalf("emma unclaimed donated nft = %+v", unclaimedRow)
	}

	// Status filters split the two populations.
	claimedOnly := true
	unclaimedOnly := false
	if page, _, err := r.UserDonatedNftsPage(ctx, aidEmma, &claimedOnly, nil, 50); err != nil || len(page) != 0 {
		t.Fatalf("emma claimed filter = %d rows, err=%v", len(page), err)
	}
	if page, _, err := r.UserDonatedNftsPage(ctx, aidEmma, &unclaimedOnly, nil, 50); err != nil || len(page) != 1 {
		t.Fatalf("emma unclaimed filter = %d rows, err=%v", len(page), err)
	}
	if page, _, err := r.UserDonatedNftsPage(ctx, aidAlice, &unclaimedOnly, nil, 50); err != nil || len(page) != 0 {
		t.Fatalf("alice unclaimed filter = %d rows, err=%v", len(page), err)
	}

	// The legacy views agree: claims by user and unclaimed-by-winner.
	legacyClaims, err := r.DonatedNFTClaimsByUser(ctx, aidAlice)
	if err != nil || len(legacyClaims) != 1 {
		t.Fatalf("legacy claims = %d, err=%v", len(legacyClaims), err)
	}
	legacyUnclaimed, err := r.UnclaimedDonatedNFTsByUser(ctx, aidEmma)
	if err != nil || len(legacyUnclaimed) != 1 || legacyUnclaimed[0].Tx.EvtLogId != 5102 {
		t.Fatalf("legacy unclaimed = %+v, err=%v", legacyUnclaimed, err)
	}

	// Wallets outside both branches see an empty page.
	if page, _, err := r.UserDonatedNftsPage(ctx, aidCarol, nil, nil, 50); err != nil || len(page) != 0 {
		t.Fatalf("carol donated nfts = %d rows, err=%v", len(page), err)
	}
}

func TestUserDonatedNftsPageIncludesTimeoutClaims(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Bob claims emma's round-2 donated NFT after the winner timeout: the
	// row must appear on bob's surface (claimer) and stay on emma's
	// (round winner), attributed to bob.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8001, 113, 1014, 7, 'nc000001', 93, '\x00')`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM evt_log WHERE id=8001"); err != nil {
			t.Errorf("cleaning synthetic nft claim: %v", err)
		}
	})
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_donated_nft_claimed(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, idx, token_aid, winner_aid, token_id)
		VALUES (8001, 113, 1014, TO_TIMESTAMP(1767226900), 7, 2, 1, 27, $1, 888)`,
		aidBob); err != nil {
		t.Fatal(err)
	}

	bob, _, err := r.UserDonatedNftsPage(ctx, aidBob, nil, nil, 50)
	if err != nil || len(bob) != 1 {
		t.Fatalf("bob donated nfts = %d, err=%v", len(bob), err)
	}
	if bob[0].Tx.EvtLogId != 5102 || !bob[0].Claimed || bob[0].Claim == nil ||
		bob[0].Claim.ClaimerAid != aidBob || bob[0].RoundWinnerAid != aidEmma {
		t.Fatalf("bob timeout claim row = %+v claim=%+v", bob[0], bob[0].Claim)
	}

	emma, _, err := r.UserDonatedNftsPage(ctx, aidEmma, nil, nil, 50)
	if err != nil || len(emma) != 1 || !emma[0].Claimed || emma[0].Claim == nil ||
		emma[0].Claim.ClaimerAid != aidBob {
		t.Fatalf("emma's row after bob's timeout claim = %+v", emma)
	}
}

func TestUserDonatedErc20Page(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Alice won round 0 and fully claimed the donated CST-token entitlement.
	alice, hasMore, err := r.UserDonatedErc20Page(ctx, aidAlice, nil, 50)
	if err != nil || hasMore || len(alice) != 1 {
		t.Fatalf("alice donated erc20 = %d, more=%v, err=%v", len(alice), hasMore, err)
	}
	summary := alice[0]
	if summary.RoundNum != 0 ||
		summary.DonatedBaseUnits != "500000000000000000000" ||
		summary.ClaimedBaseUnits != "500000000000000000000" ||
		summary.RemainingBaseUnits != "0" ||
		summary.LastClaim == nil || summary.LastClaim.EventLogID != 5046 ||
		summary.LastClaim.ClaimerAid != aidAlice {
		t.Fatalf("alice donated erc20 summary = %+v lastClaim=%+v", summary, summary.LastClaim)
	}

	// Deliberate correction over v1: cg_erc20_donation_stats.total_amount is
	// trigger-decremented on every claim, so the legacy winner view reports
	// the *remaining* amount as "donated" (zero after a full claim, with a
	// negative donate-claim diff). V2 reconstructs the true donated total.
	legacy, err := r.ERC20DonatedPrizesByWinner(ctx, aidAlice)
	if err != nil || len(legacy) != 1 {
		t.Fatalf("legacy winner summaries = %d, err=%v", len(legacy), err)
	}
	if legacy[0].AmountDonated != summary.RemainingBaseUnits ||
		legacy[0].AmountClaimed != summary.ClaimedBaseUnits {
		t.Fatalf("v2 summary %+v does not reconcile with legacy %+v", summary, legacy[0])
	}

	// Wallets without won rounds or claims see an empty page.
	if page, _, err := r.UserDonatedErc20Page(ctx, aidCarol, nil, 50); err != nil || len(page) != 0 {
		t.Fatalf("carol donated erc20 = %d rows, err=%v", len(page), err)
	}
}

func TestUserDonatedErc20PageTimeoutClaimAndPaging(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// A partially claimed round-1 entitlement: emma timeout-claims 40 of
	// 100 base units donated in dave's round.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8002, 113, 1014, 7, 'tc000001', 94, '\x00')`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cleanupCtx := context.Background()
		if _, err := r.pool().Exec(cleanupCtx,
			"DELETE FROM evt_log WHERE id=8002"); err != nil {
			t.Errorf("cleaning synthetic token claim: %v", err)
		}
		if _, err := r.pool().Exec(cleanupCtx,
			"DELETE FROM cg_erc20_donation_stats WHERE round_num=1 AND token_aid=26"); err != nil {
			t.Errorf("cleaning synthetic donation stats: %v", err)
		}
	})
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_erc20_donation_stats(round_num, token_aid, total_amount, claimed, winner_aid)
		VALUES (1, 26, 100, TRUE, $1)`, aidEmma); err != nil {
		t.Fatal(err)
	}
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_donated_tok_claimed(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, idx, token_aid, winner_aid, amount)
		VALUES (8002, 113, 1014, TO_TIMESTAMP(1767226900), 7, 1, 0, 26, $1, 40)`,
		aidEmma); err != nil {
		t.Fatal(err)
	}

	// Dave (round-1 winner) sees the partial entitlement.
	dave, _, err := r.UserDonatedErc20Page(ctx, aidDave, nil, 50)
	if err != nil || len(dave) != 1 {
		t.Fatalf("dave donated erc20 = %d, err=%v", len(dave), err)
	}
	partial := dave[0]
	if partial.RoundNum != 1 || partial.DonatedBaseUnits != "100" ||
		partial.ClaimedBaseUnits != "40" || partial.RemainingBaseUnits != "60" ||
		partial.LastClaim == nil || partial.LastClaim.ClaimerAid != aidEmma ||
		partial.LastClaim.AmountBaseUnits != "40" {
		t.Fatalf("dave partial entitlement = %+v lastClaim=%+v", partial, partial.LastClaim)
	}

	// Emma reaches the same row through the timeout-claim branch, and her
	// round-2 win contributes nothing (no donations there).
	emma, _, err := r.UserDonatedErc20Page(ctx, aidEmma, nil, 50)
	if err != nil || len(emma) != 1 || emma[0].RoundNum != 1 {
		t.Fatalf("emma donated erc20 = %+v, err=%v", emma, err)
	}

	// Alice now has two summaries across rounds 0 and 1? No: alice only
	// won round 0, and did not claim round 1 — her surface is unchanged.
	alice, _, err := r.UserDonatedErc20Page(ctx, aidAlice, nil, 50)
	if err != nil || len(alice) != 1 || alice[0].RoundNum != 0 {
		t.Fatalf("alice donated erc20 = %+v, err=%v", alice, err)
	}

	// Keyset paging across dave's two-round surface: give dave a claim on
	// round 0's token so he has two rows, then walk with limit 1.
	if _, err := r.pool().Exec(ctx, `INSERT INTO evt_log(id, block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		VALUES (8004, 113, 1014, 7, 'tc000002', 95, '\x00')`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := r.pool().Exec(context.Background(),
			"DELETE FROM evt_log WHERE id=8004"); err != nil {
			t.Errorf("cleaning second synthetic token claim: %v", err)
		}
	})
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_donated_tok_claimed(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, idx, token_aid, winner_aid, amount)
		VALUES (8004, 113, 1014, TO_TIMESTAMP(1767226910), 7, 0, 0, 26, $1, 0)`,
		aidDave); err != nil {
		t.Fatal(err)
	}

	first, hasMore, err := r.UserDonatedErc20Page(ctx, aidDave, nil, 1)
	if err != nil || !hasMore || len(first) != 1 || first[0].RoundNum != 1 {
		t.Fatalf("first page = %+v, more=%v, err=%v", first, hasMore, err)
	}
	second, hasMore, err := r.UserDonatedErc20Page(ctx, aidDave, &UserDonatedErc20PageCursor{
		Round:    first[0].RoundNum,
		TokenAid: first[0].TokenAid,
	}, 1)
	if err != nil || hasMore || len(second) != 1 || second[0].RoundNum != 0 {
		t.Fatalf("second page = %+v, more=%v, err=%v", second, hasMore, err)
	}
	// Round 0's last claim is now dave's later event, and totals include
	// both claim events.
	if second[0].LastClaim == nil || second[0].LastClaim.EventLogID != 8004 ||
		second[0].ClaimedBaseUnits != "500000000000000000000" {
		t.Fatalf("round-0 summary after second claim = %+v lastClaim=%+v",
			second[0], second[0].LastClaim)
	}
}

func TestUserHistoryPagesPropagateFailures(t *testing.T) {
	r := repo(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	cancelledCalls := map[string]func() error{
		"prizes": func() error {
			_, _, err := r.UserPrizesPage(cancelled, aidAlice, nil, 1)
			return err
		},
		"deposits": func() error {
			_, _, err := r.UserRaffleEthDepositsPage(cancelled, aidAlice, nil, nil, 1)
			return err
		},
		"nft wins": func() error {
			_, _, err := r.UserRaffleNftWinsPage(cancelled, aidAlice, nil, 1)
			return err
		},
		"eth donations": func() error {
			_, _, err := r.EthDonationsByUserPage(cancelled, aidAlice, nil, 1)
			return err
		},
		"erc20 donations": func() error {
			_, _, err := r.ERC20DonationsByUserPage(cancelled, aidAlice, nil, 1)
			return err
		},
		"nft donations": func() error {
			_, _, err := r.NFTDonationsByUserPage(cancelled, aidAlice, nil, 1)
			return err
		},
		"donated nfts": func() error {
			_, _, err := r.UserDonatedNftsPage(cancelled, aidAlice, nil, nil, 1)
			return err
		},
		"donated erc20": func() error {
			_, _, err := r.UserDonatedErc20Page(cancelled, aidAlice, nil, 1)
			return err
		},
	}
	for name, call := range cancelledCalls {
		if err := call(); !errors.Is(err, context.Canceled) {
			t.Errorf("%s cancellation = %v", name, err)
		}
	}

	st, err := spareStore(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	closedRepo := NewRepo(st)
	st.Close()
	closedCalls := map[string]func() error{
		"prizes": func() error {
			_, _, err := closedRepo.UserPrizesPage(context.Background(), aidAlice, nil, 1)
			return err
		},
		"deposits": func() error {
			_, _, err := closedRepo.UserRaffleEthDepositsPage(context.Background(), aidAlice, nil, nil, 1)
			return err
		},
		"donated erc20": func() error {
			_, _, err := closedRepo.UserDonatedErc20Page(context.Background(), aidAlice, nil, 1)
			return err
		},
	}
	for name, call := range closedCalls {
		if err := call(); err == nil {
			t.Errorf("%s succeeded on closed pool", name)
		}
	}
}

func TestUserHistoryReadIndexesExist(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// #nosec G101 -- index names and column lists, not credentials.
	wantIndexes := map[string]string{
		"cg_prize_deposit_winner_evt_idx":        "(winner_aid, evtlog_id desc)",
		"cg_raffle_nft_prize_winner_evt_idx":     "(winner_aid, evtlog_id desc)",
		"cg_eth_donated_donor_evt_idx":           "(donor_aid, evtlog_id desc)",
		"cg_eth_donated_wi_donor_evt_idx":        "(donor_aid, evtlog_id desc)",
		"cg_erc20_donation_donor_evt_idx":        "(donor_aid, evtlog_id desc)",
		"cg_nft_donation_donor_evt_idx":          "(donor_aid, evtlog_id desc)",
		"cg_donated_nft_claimed_winner_evt_idx":  "(winner_aid, evtlog_id desc)",
		"cg_donated_nft_claimed_donation_idx":    "(idx)",
		"cg_donated_tok_claimed_winner_idx":      "(winner_aid)",
		"cg_donated_tok_claimed_round_token_idx": "(round_num, token_aid)",
		"cg_raffle_eth_prize_winner_idx":         "(winner_aid)",
		"cg_lastcst_prize_winner_idx":            "(winner_aid)",
		"cg_endurance_prize_winner_idx":          "(winner_aid)",
		"cg_chrono_warrior_prize_winner_idx":     "(winner_aid)",
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
