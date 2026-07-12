//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

const (
	aliceAddress         = "0x2100000000000000000000000000000000000021"
	unindexedUserAddress = "0x9900000000000000000000000000000000000099"
)

func TestUserProfileV2(t *testing.T) {
	r := repo(t)
	golden(t, "user_profile_v2_alice", func() any {
		record, err := r.UserProfile(context.Background(), aidAlice)
		if err != nil {
			t.Fatalf("UserProfile(alice): %v", err)
		}
		return record
	})

	empty, err := r.UserProfile(context.Background(), aidPrizesWallet)
	if err != nil {
		t.Fatalf("UserProfile(prizes wallet): %v", err)
	}
	if empty.BidCount != 0 || empty.PrizeCount != 0 ||
		empty.TotalETHSpentWei != "0" || empty.TotalETHWonWei != "0" ||
		empty.MaxETHBidWei != nil {
		t.Fatalf("inactive indexed profile = %+v", empty)
	}
	if _, err := r.UserProfile(context.Background(), 999999); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("UserProfile(missing) error = %v, want ErrNotFound", err)
	}
}

func TestUserProfileV2UsesCanonicalPrizeRows(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	before, err := r.UserProfile(ctx, aidAlice)
	if err != nil {
		t.Fatal(err)
	}

	var original struct {
		maxWin      string
		prizeCount  int64
		prizeSum    string
		erc20Count  int64
		erc721Count int64
		unclaimed   int64
	}
	err = r.pool().QueryRow(ctx, `SELECT max_win_amount::TEXT,prizes_count,
		prizes_sum::TEXT,erc20_count,erc721_count,unclaimed_nfts
		FROM cg_winner WHERE winner_aid=$1`, aidAlice).Scan(
		&original.maxWin,
		&original.prizeCount,
		&original.prizeSum,
		&original.erc20Count,
		&original.erc721Count,
		&original.unclaimed,
	)
	if err != nil {
		t.Fatalf("read original aggregate: %v", err)
	}
	_, err = r.pool().Exec(ctx, `UPDATE cg_winner
		SET max_win_amount=999999,prizes_count=999999,prizes_sum=999999,
			erc20_count=999999,erc721_count=999999,unclaimed_nfts=999999
		WHERE winner_aid=$1`, aidAlice)
	if err != nil {
		t.Fatalf("corrupt winner aggregate: %v", err)
	}
	t.Cleanup(func() {
		if _, restoreErr := r.pool().Exec(context.Background(), `UPDATE cg_winner
			SET max_win_amount=$2,prizes_count=$3,prizes_sum=$4,
				erc20_count=$5,erc721_count=$6,unclaimed_nfts=$7
			WHERE winner_aid=$1`,
			aidAlice,
			original.maxWin,
			original.prizeCount,
			original.prizeSum,
			original.erc20Count,
			original.erc721Count,
			original.unclaimed,
		); restoreErr != nil {
			t.Errorf("restore winner aggregate: %v", restoreErr)
		}
	})

	after, err := r.UserProfile(ctx, aidAlice)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(after, before) {
		t.Fatalf("profile depends on cg_winner aggregate\nafter:  %+v\nbefore: %+v", after, before)
	}
}

func TestBidsByUserPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	full, err := r.BidsByUser(ctx, aidAlice)
	if err != nil {
		t.Fatal(err)
	}

	var gotEventIDs []int64
	var after *UserBidPageCursor
	for {
		page, hasMore, err := r.BidsByUserPage(ctx, aidAlice, after, 2)
		if err != nil {
			t.Fatalf("BidsByUserPage: %v", err)
		}
		if len(page) > 2 {
			t.Fatalf("page length = %d", len(page))
		}
		for i := range page {
			if page[i].BidderAid != aidAlice {
				t.Fatalf("bidder aid = %d", page[i].BidderAid)
			}
			if len(gotEventIDs) > 0 &&
				page[i].Tx.EvtLogId >= gotEventIDs[len(gotEventIDs)-1] {
				t.Fatalf("unordered event IDs: %v then %d", gotEventIDs, page[i].Tx.EvtLogId)
			}
			gotEventIDs = append(gotEventIDs, page[i].Tx.EvtLogId)
		}
		if !hasMore {
			break
		}
		if len(page) == 0 {
			t.Fatal("hasMore without a cursor row")
		}
		after = &UserBidPageCursor{EventLogID: page[len(page)-1].Tx.EvtLogId}
	}

	wantEventIDs := make([]int64, len(full))
	for i := range full {
		wantEventIDs[i] = full[i].Tx.EvtLogId
	}
	if !reflect.DeepEqual(gotEventIDs, wantEventIDs) {
		t.Fatalf("paged event IDs = %v, full list = %v", gotEventIDs, wantEventIDs)
	}

	empty, hasMore, err := r.BidsByUserPage(ctx, aidPrizesWallet, nil, 2)
	if err != nil || hasMore || len(empty) != 0 || empty == nil {
		t.Fatalf("empty page = len %d nil=%v more=%v err=%v",
			len(empty), empty == nil, hasMore, err)
	}
}

func TestUserV2StoreErrorPaths(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := r.UserProfile(ctx, aidAlice); !errors.Is(err, context.Canceled) {
		t.Errorf("UserProfile cancellation = %v", err)
	}
	if _, _, err := r.BidsByUserPage(ctx, aidAlice, nil, 2); !errors.Is(err, context.Canceled) {
		t.Errorf("BidsByUserPage cancellation = %v", err)
	}

	st, err := spareStore(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	closedRepo := NewRepo(st)
	st.Close()
	if _, err := closedRepo.UserProfile(context.Background(), aidAlice); err == nil {
		t.Error("UserProfile succeeded on closed pool")
	}
	if _, _, err := closedRepo.BidsByUserPage(context.Background(), aidAlice, nil, 2); err == nil {
		t.Error("BidsByUserPage succeeded on closed pool")
	}
}

func TestUserAddressIDAndBidIndex(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	userAid, err := r.UserAddressID(ctx, aliceAddress)
	if err != nil || userAid != aidAlice {
		t.Fatalf("UserAddressID(alice) = %d, %v", userAid, err)
	}
	if _, err := r.UserAddressID(ctx, unindexedUserAddress); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("UserAddressID(unindexed) = %v, want ErrNotFound", err)
	}

	var indexDefinition string
	err = r.pool().QueryRow(ctx, `SELECT indexdef FROM pg_indexes
		WHERE schemaname='public' AND indexname='cg_bid_user_evtlog_idx'`).Scan(
		&indexDefinition,
	)
	if err != nil {
		t.Fatalf("read user bid index: %v", err)
	}
	normalized := strings.ToLower(strings.Join(strings.Fields(indexDefinition), " "))
	if !strings.Contains(normalized, "(bidder_aid, evtlog_id desc)") {
		t.Fatalf("unexpected index definition: %s", indexDefinition)
	}
}
