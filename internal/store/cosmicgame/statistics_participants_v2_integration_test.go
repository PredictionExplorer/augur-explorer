//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"math/big"
	"reflect"
	"strconv"
	"testing"
)

func TestParticipantPageGoldens(t *testing.T) {
	r := repo(t)
	golden(t, "bidder_participants_page_v2", func() any {
		rows, more, err := r.BidderParticipantsPage(context.Background(), nil, 2)
		if err != nil {
			t.Fatalf("BidderParticipantsPage: %v", err)
		}
		return struct {
			Rows    []BidderParticipantRecord
			HasMore bool
		}{rows, more}
	})
	golden(t, "winner_participants_page_v2", func() any {
		rows, more, err := r.WinnerParticipantsPage(context.Background(), nil, 2)
		if err != nil {
			t.Fatalf("WinnerParticipantsPage: %v", err)
		}
		return struct {
			Rows    []WinnerParticipantRecord
			HasMore bool
		}{rows, more}
	})
	golden(t, "donor_participants_page_v2", func() any {
		rows, more, err := r.DonorParticipantsPage(context.Background(), nil, 1)
		if err != nil {
			t.Fatalf("DonorParticipantsPage: %v", err)
		}
		return struct {
			Rows    []DonorParticipantRecord
			HasMore bool
		}{rows, more}
	})
	golden(t, "cst_staker_participants_page_v2", func() any {
		rows, more, err := r.CSTStakerParticipantsPage(context.Background(), nil, 1)
		if err != nil {
			t.Fatalf("CSTStakerParticipantsPage: %v", err)
		}
		return struct {
			Rows    []CSTStakerParticipantRecord
			HasMore bool
		}{rows, more}
	})
	golden(t, "randomwalk_staker_participants_page_v2", func() any {
		rows, more, err := r.RandomWalkStakerParticipantsPage(context.Background(), nil, 2)
		if err != nil {
			t.Fatalf("RandomWalkStakerParticipantsPage: %v", err)
		}
		return struct {
			Rows    []RandomWalkStakerParticipantRecord
			HasMore bool
		}{rows, more}
	})
	golden(t, "dual_staker_participants_page_v2", func() any {
		rows, more, err := r.DualStakerParticipantsPage(context.Background(), nil, 1)
		if err != nil {
			t.Fatalf("DualStakerParticipantsPage: %v", err)
		}
		return struct {
			Rows    []DualStakerParticipantRecord
			HasMore bool
		}{rows, more}
	})
}

func TestParticipantPagesMatchLegacyLists(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	legacyBidders, err := r.UniqueBidders(ctx)
	if err != nil {
		t.Fatal(err)
	}
	bidders := collectParticipantPages(t, ParticipantBidders, r.BidderParticipantsPage,
		func(row BidderParticipantRecord) ParticipantPageCursor {
			return ParticipantPageCursor{
				Kind: ParticipantBidders, SortValue: strconv.FormatInt(row.BidCount, 10), AddressID: row.BidderAid,
			}
		})
	if len(bidders) != len(legacyBidders) {
		t.Fatalf("bidders = %d, legacy = %d", len(bidders), len(legacyBidders))
	}
	legacyBidderByID := make(map[int64]struct {
		address, maxBid string
		count           int64
	}, len(legacyBidders))
	for _, row := range legacyBidders {
		legacyBidderByID[row.BidderAid] = struct {
			address, maxBid string
			count           int64
		}{row.BidderAddr, row.MaxBidAmount, row.NumBids}
	}
	for _, row := range bidders {
		want, ok := legacyBidderByID[row.BidderAid]
		if !ok || row.Address != want.address || row.BidCount != want.count || row.MaxBidWei != want.maxBid {
			t.Errorf("bidder row %+v does not match legacy %+v", row, want)
		}
	}

	legacyWinners, err := r.UniqueWinners(ctx)
	if err != nil {
		t.Fatal(err)
	}
	winners := collectParticipantPages(t, ParticipantWinners, r.WinnerParticipantsPage,
		func(row WinnerParticipantRecord) ParticipantPageCursor {
			return ParticipantPageCursor{
				Kind: ParticipantWinners, SortValue: strconv.FormatInt(row.PrizeCount, 10), AddressID: row.WinnerAid,
			}
		})
	if len(winners) != len(legacyWinners) {
		t.Fatalf("winners = %d, legacy = %d", len(winners), len(legacyWinners))
	}
	legacyWinnerByID := make(map[int64]WinnerParticipantRecord, len(legacyWinners))
	for _, row := range legacyWinners {
		legacyWinnerByID[row.WinnerAid] = WinnerParticipantRecord{
			WinnerAid: row.WinnerAid, Address: row.WinnerAddr, PrizeCount: row.PrizesCount,
			MaxMainPrizeETHWei: row.WinnerStats.MaxWinAmount,
			TotalETHWonWei:     row.WinnerStats.PrizesSum,
			CSTPrizeCount:      row.WinnerStats.ERC20Count, NFTPrizeCount: row.WinnerStats.ERC721Count,
			UnclaimedNFTCount: row.WinnerStats.UnclaimedNfts,
			TotalETHSpentWei:  row.WinnerStats.TotalSpent,
		}
	}
	for _, row := range winners {
		if want, ok := legacyWinnerByID[row.WinnerAid]; !ok || !reflect.DeepEqual(row, want) {
			t.Errorf("winner row %+v does not match legacy %+v", row, want)
		}
	}

	legacyDonors, err := r.UniqueDonors(ctx)
	if err != nil {
		t.Fatal(err)
	}
	donors := collectParticipantPages(t, ParticipantDonors, r.DonorParticipantsPage,
		func(row DonorParticipantRecord) ParticipantPageCursor {
			return ParticipantPageCursor{
				Kind: ParticipantDonors, SortValue: row.TotalDonatedWei, AddressID: row.DonorAid,
			}
		})
	if len(donors) != len(legacyDonors) {
		t.Fatalf("donors = %d, legacy = %d", len(donors), len(legacyDonors))
	}
	legacyDonorByID := make(map[int64]struct {
		address, total string
		count          int64
	}, len(legacyDonors))
	for _, row := range legacyDonors {
		legacyDonorByID[row.DonorAid] = struct {
			address, total string
			count          int64
		}{row.DonorAddr, row.TotalDonated, row.CountDonations}
	}
	for _, row := range donors {
		want, ok := legacyDonorByID[row.DonorAid]
		if !ok || row.Address != want.address ||
			row.DonationCount != want.count || row.TotalDonatedWei != want.total {
			t.Errorf("donor row %+v does not match legacy %+v", row, want)
		}
	}

	legacyCST, err := r.UniqueStakersCst(ctx)
	if err != nil {
		t.Fatal(err)
	}
	cst := collectParticipantPages(t, ParticipantCSTStakers, r.CSTStakerParticipantsPage,
		func(row CSTStakerParticipantRecord) ParticipantPageCursor {
			return ParticipantPageCursor{
				Kind: ParticipantCSTStakers, SortValue: row.TotalRewardWei, AddressID: row.StakerAid,
			}
		})
	if len(cst) != len(legacyCST) {
		t.Fatalf("CST stakers = %d, legacy = %d", len(cst), len(legacyCST))
	}
	legacyCSTByID := make(map[int64]CSTStakerParticipantRecord, len(legacyCST))
	for _, row := range legacyCST {
		legacyCSTByID[row.StakerAid] = CSTStakerParticipantRecord{
			StakerAid: row.StakerAid, Address: row.StakerAddr,
			StakedTokenCount: row.TotalTokensStaked, StakeActionCount: row.NumStakeActions,
			UnstakeActionCount: row.NumUnstakeActions, TotalRewardWei: row.TotalReward,
			UnclaimedRewardWei: row.UnclaimedReward,
		}
	}
	for _, row := range cst {
		if want, ok := legacyCSTByID[row.StakerAid]; !ok || !reflect.DeepEqual(row, want) {
			t.Errorf("CST row %+v does not match legacy %+v", row, want)
		}
	}

	legacyRandomWalk, err := r.UniqueStakersRwalk(ctx)
	if err != nil {
		t.Fatal(err)
	}
	randomWalk := collectParticipantPages(t, ParticipantRandomWalkStakers, r.RandomWalkStakerParticipantsPage,
		func(row RandomWalkStakerParticipantRecord) ParticipantPageCursor {
			return ParticipantPageCursor{
				Kind:      ParticipantRandomWalkStakers,
				SortValue: strconv.FormatInt(row.StakedTokenCount, 10), AddressID: row.StakerAid,
			}
		})
	if len(randomWalk) != len(legacyRandomWalk) {
		t.Fatalf("RandomWalk stakers = %d, legacy = %d", len(randomWalk), len(legacyRandomWalk))
	}
	legacyRandomWalkByID := make(map[int64]RandomWalkStakerParticipantRecord, len(legacyRandomWalk))
	for _, row := range legacyRandomWalk {
		legacyRandomWalkByID[row.StakerAid] = RandomWalkStakerParticipantRecord{
			StakerAid: row.StakerAid, Address: row.StakerAddr,
			StakedTokenCount: row.TotalTokensStaked, StakeActionCount: row.NumStakeActions,
			UnstakeActionCount: row.NumUnstakeActions, MintedTokenCount: row.TotalTokensMinted,
		}
	}
	for _, row := range randomWalk {
		if want, ok := legacyRandomWalkByID[row.StakerAid]; !ok || !reflect.DeepEqual(row, want) {
			t.Errorf("RandomWalk row %+v does not match legacy %+v", row, want)
		}
	}

	legacyDual, err := r.UniqueStakersBoth(ctx)
	if err != nil {
		t.Fatal(err)
	}
	dual := collectParticipantPages(t, ParticipantDualStakers, r.DualStakerParticipantsPage,
		func(row DualStakerParticipantRecord) ParticipantPageCursor {
			return ParticipantPageCursor{
				Kind:      ParticipantDualStakers,
				SortValue: strconv.FormatInt(row.TotalStakedTokenCount, 10), AddressID: row.StakerAid,
			}
		})
	if len(dual) != len(legacyDual) {
		t.Fatalf("dual stakers = %d, legacy = %d", len(dual), len(legacyDual))
	}
	legacyDualByID := make(map[int64]DualStakerParticipantRecord, len(legacyDual))
	for _, row := range legacyDual {
		legacyDualByID[row.StakerAid] = DualStakerParticipantRecord{
			StakerAid: row.StakerAid, Address: row.StakerAddr,
			TotalStakedTokenCount:        row.TotalStakedTokensBoth,
			CSTStakedTokenCount:          row.CSTStats.TotalTokensStaked,
			CSTStakeActionCount:          row.CSTStats.NumStakeActions,
			CSTUnstakeActionCount:        row.CSTStats.NumUnstakeActions,
			CSTTotalRewardWei:            row.CSTStats.TotalReward,
			CSTUnclaimedRewardWei:        row.CSTStats.UnclaimedReward,
			RandomWalkStakedTokenCount:   row.RWalkStats.TotalTokensStaked,
			RandomWalkStakeActionCount:   row.RWalkStats.NumStakeActions,
			RandomWalkUnstakeActionCount: row.RWalkStats.NumUnstakeActions,
			RandomWalkMintedTokenCount:   row.RWalkStats.TotalTokensMinted,
		}
	}
	for _, row := range dual {
		if want, ok := legacyDualByID[row.StakerAid]; !ok || !reflect.DeepEqual(row, want) {
			t.Errorf("dual row %+v does not match legacy %+v", row, want)
		}
	}
}

func TestWinnerParticipantsIgnoreReplaySensitiveAggregate(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	var (
		maxWin, totalWon                          string
		prizeCount, cstCount, nftCount, unclaimed int64
	)
	if err := r.pool().QueryRow(ctx, `SELECT max_win_amount::TEXT,prizes_sum::TEXT,
		prizes_count,erc20_count,erc721_count,unclaimed_nfts
		FROM cg_winner WHERE winner_aid=$1`, aidAlice).Scan(
		&maxWin, &totalWon, &prizeCount, &cstCount, &nftCount, &unclaimed,
	); err != nil {
		t.Fatal(err)
	}
	defer func() {
		_, err := r.pool().Exec(ctx, `UPDATE cg_winner SET
			max_win_amount=$1,prizes_sum=$2,prizes_count=$3,erc20_count=$4,
			erc721_count=$5,unclaimed_nfts=$6 WHERE winner_aid=$7`,
			maxWin, totalWon, prizeCount, cstCount, nftCount, unclaimed, aidAlice)
		if err != nil {
			t.Errorf("restoring winner aggregate: %v", err)
		}
	}()
	if _, err := r.pool().Exec(ctx, `UPDATE cg_winner SET
		max_win_amount=NULL,prizes_sum=NULL,prizes_count=NULL,erc20_count=NULL,
		erc721_count=NULL,unclaimed_nfts=NULL WHERE winner_aid=$1`, aidAlice); err != nil {
		t.Fatal(err)
	}
	rows, _, err := r.WinnerParticipantsPage(ctx, nil, 10)
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range rows {
		if row.WinnerAid != aidAlice {
			continue
		}
		if row.PrizeCount != 7 ||
			row.MaxMainPrizeETHWei != "500000000000000000" ||
			row.TotalETHWonWei != "640000000000000000" ||
			row.CSTPrizeCount != 2 || row.NFTPrizeCount != 2 {
			t.Fatalf("canonical winner row = %+v", row)
		}
		return
	}
	t.Fatal("canonical winner directory omitted Alice")
}

func TestBidderParticipantsExcludeZeroCountTombstones(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	var tombstoneID int64
	if err := r.pool().QueryRow(ctx, `INSERT INTO address(block_num,tx_id,addr)
		VALUES (0,0,'0x9000030000000000000000000000000000000003')
		RETURNING address_id`).Scan(&tombstoneID); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if _, err := r.pool().Exec(ctx, "DELETE FROM cg_bidder WHERE bidder_aid=$1", tombstoneID); err != nil {
			t.Errorf("cleaning bidder tombstone: %v", err)
		}
		if _, err := r.pool().Exec(ctx, "DELETE FROM address WHERE address_id=$1", tombstoneID); err != nil {
			t.Errorf("cleaning bidder address: %v", err)
		}
	}()
	if _, err := r.pool().Exec(ctx, `INSERT INTO cg_bidder(bidder_aid,num_bids,max_bid)
		VALUES ($1,0,0)`, tombstoneID); err != nil {
		t.Fatal(err)
	}
	rows, _, err := r.BidderParticipantsPage(ctx, nil, 200)
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range rows {
		if row.BidderAid == tombstoneID {
			t.Fatalf("zero-count bidder leaked into directory: %+v", row)
		}
	}
}

func TestDualStakerParticipantPageTieBoundary(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	insertAddress := func(address string) int64 {
		t.Helper()
		var id int64
		if err := r.pool().QueryRow(ctx, `INSERT INTO address(block_num,tx_id,addr)
			VALUES (0,0,$1) RETURNING address_id`, address).Scan(&id); err != nil {
			t.Fatal(err)
		}
		return id
	}
	var firstID, secondID int64
	defer func() {
		for _, cleanup := range []struct {
			name, query string
		}{
			{"RandomWalk stakers", "DELETE FROM cg_staker_rwalk WHERE staker_aid IN ($1,$2)"},
			{"CST stakers", "DELETE FROM cg_staker_cst WHERE staker_aid IN ($1,$2)"},
			{"addresses", "DELETE FROM address WHERE address_id IN ($1,$2)"},
		} {
			if _, err := r.pool().Exec(ctx, cleanup.query, firstID, secondID); err != nil {
				t.Errorf("cleaning %s: %v", cleanup.name, err)
			}
		}
	}()
	firstID = insertAddress("0x9000010000000000000000000000000000000001")
	secondID = insertAddress("0x9000020000000000000000000000000000000002")
	for _, id := range []int64{firstID, secondID} {
		if _, err := r.pool().Exec(ctx, `INSERT INTO cg_staker_cst
			(staker_aid,total_tokens_staked,num_stake_actions,total_reward,unclaimed_reward)
			VALUES ($1,3,1,10,5)`, id); err != nil {
			t.Fatal(err)
		}
		if _, err := r.pool().Exec(ctx, `INSERT INTO cg_staker_rwalk
			(staker_aid,total_tokens_staked,num_stake_actions) VALUES ($1,3,1)`, id); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := r.pool().Exec(ctx,
		"UPDATE cg_staker_cst SET num_stake_actions=NULL WHERE staker_aid=$1", secondID); err != nil {
		t.Fatal(err)
	}
	if _, err := r.pool().Exec(ctx,
		"UPDATE cg_staker_rwalk SET num_stake_actions=NULL WHERE staker_aid=$1", secondID); err != nil {
		t.Fatal(err)
	}
	first, more, err := r.DualStakerParticipantsPage(ctx, nil, 1)
	if err != nil || !more || len(first) != 1 || first[0].StakerAid != firstID {
		t.Fatalf("first page = %+v,%v,%v", first, more, err)
	}
	after := &ParticipantPageCursor{
		Kind: ParticipantDualStakers, SortValue: "6", AddressID: firstID,
	}
	second, more, err := r.DualStakerParticipantsPage(ctx, after, 1)
	if err != nil || !more || len(second) != 1 || second[0].StakerAid != secondID ||
		second[0].CSTStakeActionCount != 0 || second[0].RandomWalkStakeActionCount != 0 {
		t.Fatalf("second page = %+v,%v,%v", second, more, err)
	}
}

func TestParticipantPagesPropagateFailures(t *testing.T) {
	r := repo(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	cancelledCalls := map[string]func() error{
		"bidders": func() error {
			_, _, err := r.BidderParticipantsPage(cancelled, nil, 1)
			return err
		},
		"winners": func() error {
			_, _, err := r.WinnerParticipantsPage(cancelled, nil, 1)
			return err
		},
		"donors": func() error {
			_, _, err := r.DonorParticipantsPage(cancelled, nil, 1)
			return err
		},
		"CST stakers": func() error {
			_, _, err := r.CSTStakerParticipantsPage(cancelled, nil, 1)
			return err
		},
		"RandomWalk stakers": func() error {
			_, _, err := r.RandomWalkStakerParticipantsPage(cancelled, nil, 1)
			return err
		},
		"dual stakers": func() error {
			_, _, err := r.DualStakerParticipantsPage(cancelled, nil, 1)
			return err
		},
	}
	for name, call := range cancelledCalls {
		if err := call(); !errors.Is(err, context.Canceled) {
			t.Errorf("%s cancellation = %v, want context.Canceled", name, err)
		}
	}

	st, err := spareStore(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	closedRepo := NewRepo(st)
	st.Close()
	closedCalls := []func() error{
		func() error {
			_, _, err := closedRepo.BidderParticipantsPage(context.Background(), nil, 1)
			return err
		},
		func() error {
			_, _, err := closedRepo.WinnerParticipantsPage(context.Background(), nil, 1)
			return err
		},
		func() error {
			_, _, err := closedRepo.DonorParticipantsPage(context.Background(), nil, 1)
			return err
		},
		func() error {
			_, _, err := closedRepo.CSTStakerParticipantsPage(context.Background(), nil, 1)
			return err
		},
		func() error {
			_, _, err := closedRepo.RandomWalkStakerParticipantsPage(context.Background(), nil, 1)
			return err
		},
		func() error {
			_, _, err := closedRepo.DualStakerParticipantsPage(context.Background(), nil, 1)
			return err
		},
	}
	for i, call := range closedCalls {
		if err := call(); err == nil {
			t.Errorf("closed-pool call %d returned nil", i)
		}
	}
}

func collectParticipantPages[T any](
	t *testing.T,
	kind ParticipantKind,
	page func(context.Context, *ParticipantPageCursor, int) ([]T, bool, error),
	cursorFor func(T) ParticipantPageCursor,
) []T {
	t.Helper()
	ctx := context.Background()
	var (
		all   []T
		after *ParticipantPageCursor
	)
	for pageNumber := 0; pageNumber < 100; pageNumber++ {
		rows, more, err := page(ctx, after, 1)
		if err != nil {
			t.Fatalf("%s page %d: %v", kind, pageNumber, err)
		}
		if len(rows) > 1 || (more && len(rows) != 1) {
			t.Fatalf("%s page %d rows=%d more=%v", kind, pageNumber, len(rows), more)
		}
		for _, row := range rows {
			current := cursorFor(row)
			if current.Kind != kind || !participantPageCursorFollows(current, after) {
				t.Fatalf("%s page %d cursor %+v does not follow %+v", kind, pageNumber, current, after)
			}
			all = append(all, row)
			currentCopy := current
			after = &currentCopy
		}
		if !more {
			rows, hasMore, err := page(ctx, after, 1)
			if err != nil || hasMore || len(rows) != 0 {
				t.Fatalf("%s terminal page = %v,%v,%v", kind, rows, hasMore, err)
			}
			return all
		}
	}
	t.Fatalf("%s pagination did not terminate", kind)
	return nil
}

func participantPageCursorFollows(current ParticipantPageCursor, previous *ParticipantPageCursor) bool {
	if previous == nil {
		return canonicalParticipantSortValue(current.SortValue) && current.AddressID > 0
	}
	if current.Kind != previous.Kind {
		return false
	}
	currentValue, currentOK := new(big.Int).SetString(current.SortValue, 10)
	previousValue, previousOK := new(big.Int).SetString(previous.SortValue, 10)
	if !currentOK || !previousOK {
		return false
	}
	comparison := currentValue.Cmp(previousValue)
	return comparison < 0 || (comparison == 0 && current.AddressID > previous.AddressID)
}
