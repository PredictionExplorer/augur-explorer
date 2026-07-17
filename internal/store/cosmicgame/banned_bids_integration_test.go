//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestBannedBids(t *testing.T) {
	r := repo(t)
	golden(t, "banned_bids", func() any {
		recs, err := r.BannedBids(context.Background())
		if err != nil {
			t.Fatalf("BannedBids: %v", err)
		}
		return recs
	})
}

func TestBannedBidsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	const (
		olderBid = int64(999_997)
		newerBid = int64(999_998)
		laterBid = int64(999_999)
	)
	for _, bidID := range []int64{olderBid, newerBid, laterBid} {
		if err := r.DeleteBannedBid(ctx, bidID); err != nil {
			t.Fatalf("pre-clean bid %d: %v", bidID, err)
		}
		t.Cleanup(func() {
			if err := r.DeleteBannedBid(ctx, bidID); err != nil {
				t.Errorf("cleanup bid %d: %v", bidID, err)
			}
		})
	}
	older, err := r.CreateBannedBid(ctx, olderBid,
		"0x2100000000000000000000000000000000000021", time.Unix(100, 0))
	if err != nil {
		t.Fatalf("create older ban: %v", err)
	}
	newer, err := r.CreateBannedBid(ctx, newerBid,
		"0x2200000000000000000000000000000000000022", time.Unix(200, 0))
	if err != nil {
		t.Fatalf("create newer ban: %v", err)
	}

	first, hasMore, err := r.BannedBidsPage(ctx, nil, 1)
	if err != nil || !hasMore || len(first) != 1 || first[0].Id != newer.Id {
		t.Fatalf("first page = %+v, hasMore=%v, err=%v; newer=%+v", first, hasMore, err, newer)
	}

	// A row inserted ahead of the page boundary must not leak into this
	// traversal; clients poll the head page to see newly active bans.
	later, err := r.CreateBannedBid(ctx, laterBid,
		"0x2300000000000000000000000000000000000023", time.Unix(300, 0))
	if err != nil {
		t.Fatalf("create later ban: %v", err)
	}
	second, hasMore, err := r.BannedBidsPage(ctx, &BannedBidPageCursor{ID: newer.Id}, 1)
	if err != nil || !hasMore || len(second) != 1 || second[0].Id != older.Id {
		t.Fatalf("second page = %+v, hasMore=%v, err=%v; older=%+v", second, hasMore, err, older)
	}
	if second[0].Id == later.Id {
		t.Fatal("new head row leaked into an existing traversal")
	}

	var collected []int64
	after := (*BannedBidPageCursor)(nil)
	for {
		page, more, pageErr := r.BannedBidsPage(ctx, after, 1)
		if pageErr != nil {
			t.Fatalf("page walk: %v", pageErr)
		}
		for _, rec := range page {
			if len(collected) > 0 && rec.Id >= collected[len(collected)-1] {
				t.Fatalf("page walk is not strictly descending: %v then %d", collected, rec.Id)
			}
			collected = append(collected, rec.Id)
		}
		if !more {
			break
		}
		if len(page) != 1 {
			t.Fatalf("hasMore page returned %d rows", len(page))
		}
		after = &BannedBidPageCursor{ID: page[0].Id}
	}
	if len(collected) < 4 {
		t.Fatalf("page walk returned %d rows, want fixture plus three temporary bans", len(collected))
	}

	if _, _, err := r.BannedBidsPage(ctx, &BannedBidPageCursor{}, 1); err == nil {
		t.Error("BannedBidsPage accepted an invalid cursor")
	}
	if _, _, err := r.BannedBidsPage(ctx, nil, 0); err == nil {
		t.Error("BannedBidsPage accepted a non-positive limit")
	}
}

func TestBannedBidV2WritesAndLookup(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	const bidID = int64(999_996)
	const address = "0x2100000000000000000000000000000000000021"
	if err := r.DeleteBannedBid(ctx, bidID); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := r.DeleteBannedBid(ctx, bidID); err != nil {
			t.Errorf("cleanup: %v", err)
		}
	})

	bannedAt := time.Unix(1767231234, 0).UTC()
	rec, err := r.CreateBannedBid(ctx, bidID, address, bannedAt)
	if err != nil {
		t.Fatalf("CreateBannedBid: %v", err)
	}
	if rec.Id < 1 || rec.BidId != bidID || rec.UserAddr != address || rec.CreatedAt != bannedAt.Unix() {
		t.Fatalf("created record = %+v", rec)
	}
	if _, err := r.CreateBannedBid(ctx, bidID, address, bannedAt); !errors.Is(err, store.ErrConflict) {
		t.Fatalf("duplicate CreateBannedBid = %v, want store.ErrConflict", err)
	}
	removed, err := r.RemoveBannedBid(ctx, bidID)
	if err != nil || !removed {
		t.Fatalf("RemoveBannedBid = %v, %v; want true, nil", removed, err)
	}
	removed, err = r.RemoveBannedBid(ctx, bidID)
	if err != nil || removed {
		t.Fatalf("second RemoveBannedBid = %v, %v; want false, nil", removed, err)
	}

	got, err := r.BidderAddressForBid(ctx, 2002)
	if err != nil || got != "0x2200000000000000000000000000000000000022" {
		t.Fatalf("BidderAddressForBid(2002) = %q, %v", got, err)
	}
	if _, err := r.BidderAddressForBid(ctx, 999_999_999); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("BidderAddressForBid(missing) = %v, want store.ErrNotFound", err)
	}

	if _, err := r.CreateBannedBid(ctx, 0, address, bannedAt); err == nil {
		t.Error("CreateBannedBid accepted a zero bid id")
	}
	if _, err := r.CreateBannedBid(ctx, bidID, "", bannedAt); err == nil {
		t.Error("CreateBannedBid accepted an empty address")
	}
	if _, err := r.CreateBannedBid(ctx, bidID, address, time.Time{}); err == nil {
		t.Error("CreateBannedBid accepted a zero timestamp")
	}
	if _, err := r.CreateBannedBid(ctx, bidID, address, time.Unix(-1, 0)); err == nil {
		t.Error("CreateBannedBid accepted a pre-epoch timestamp")
	}
	if _, err := r.RemoveBannedBid(ctx, 0); err == nil {
		t.Error("RemoveBannedBid accepted a zero bid id")
	}
	if _, err := r.BidderAddressForBid(ctx, 0); err == nil {
		t.Error("BidderAddressForBid accepted a zero bid id")
	}
}

func TestBannedBidUniqueIndex(t *testing.T) {
	r := repo(t)
	var definition string
	err := r.pool().QueryRow(context.Background(), `SELECT indexdef
		FROM pg_indexes
		WHERE schemaname = 'public' AND indexname = 'idx_cg_banned_bids_bid_id'`).Scan(&definition)
	if err != nil {
		t.Fatalf("read bid-ban index: %v", err)
	}
	if want := "CREATE UNIQUE INDEX"; len(definition) < len(want) || definition[:len(want)] != want {
		t.Fatalf("bid-ban index is not unique: %s", definition)
	}
}

// TestBannedBidInsertDeleteRoundTrip exercises the write paths and restores
// the fixture state so golden-based tests are unaffected regardless of
// execution order.
func TestBannedBidInsertDeleteRoundTrip(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	const bidID = 999_999
	const userAddr = "0x2100000000000000000000000000000000000021"

	list, err := r.BannedBids(ctx)
	if err != nil {
		t.Fatalf("BannedBids: %v", err)
	}
	before := len(list)
	if err := r.InsertBannedBid(ctx, bidID, userAddr); err != nil {
		t.Fatalf("InsertBannedBid: %v", err)
	}
	// Ensure restoration even if the assertions below fail.
	t.Cleanup(func() {
		if err := r.DeleteBannedBid(ctx, bidID); err != nil {
			t.Errorf("cleanup DeleteBannedBid: %v", err)
		}
	})

	list, err = r.BannedBids(ctx)
	if err != nil {
		t.Fatalf("BannedBids after insert: %v", err)
	}
	if len(list) != before+1 {
		t.Fatalf("after insert: got %d banned bids, want %d", len(list), before+1)
	}
	inserted := list[len(list)-1]
	if inserted.BidId != bidID || inserted.UserAddr != userAddr {
		t.Fatalf("inserted row mismatch: got bid_id=%d addr=%s", inserted.BidId, inserted.UserAddr)
	}
	if inserted.CreatedAt <= 0 {
		t.Fatalf("inserted row has non-positive created_at: %d", inserted.CreatedAt)
	}

	if err := r.DeleteBannedBid(ctx, bidID); err != nil {
		t.Fatalf("DeleteBannedBid: %v", err)
	}
	list, err = r.BannedBids(ctx)
	if err != nil {
		t.Fatalf("BannedBids after delete: %v", err)
	}
	if got := len(list); got != before {
		t.Fatalf("after delete: got %d banned bids, want %d", got, before)
	}
}

// TestErrorPaths pins the failure semantics the legacy layer could never
// express (it exited the process): a cancelled context aborts the query with
// context.Canceled in the chain, and a closed pool yields an error rather
// than a panic.
func TestErrorPaths(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cancelled, cancel := context.WithCancel(ctx)
	cancel()

	if _, err := r.BannedBids(cancelled); !errors.Is(err, context.Canceled) {
		t.Errorf("BannedBids(cancelled ctx) = %v, want context.Canceled in chain", err)
	}
	if _, _, err := r.BannedBidsPage(cancelled, nil, 1); !errors.Is(err, context.Canceled) {
		t.Errorf("BannedBidsPage(cancelled ctx) = %v, want context.Canceled in chain", err)
	}
	if _, err := r.BidderAddressForBid(cancelled, 2002); !errors.Is(err, context.Canceled) {
		t.Errorf("BidderAddressForBid(cancelled ctx) = %v, want context.Canceled in chain", err)
	}
	if _, err := r.CreateBannedBid(cancelled, 999_995,
		"0x2100000000000000000000000000000000000021", time.Now()); !errors.Is(err, context.Canceled) {
		t.Errorf("CreateBannedBid(cancelled ctx) = %v, want context.Canceled in chain", err)
	}
	if _, err := r.RemoveBannedBid(cancelled, 2002); !errors.Is(err, context.Canceled) {
		t.Errorf("RemoveBannedBid(cancelled ctx) = %v, want context.Canceled in chain", err)
	}
	if _, err := r.PrizeClaims(cancelled, 0, 1); !errors.Is(err, context.Canceled) {
		t.Errorf("PrizeClaims(cancelled ctx) = %v, want context.Canceled in chain", err)
	}
	if _, err := r.CosmicTokenStatistics(cancelled); !errors.Is(err, context.Canceled) {
		t.Errorf("CosmicTokenStatistics(cancelled ctx) = %v, want context.Canceled in chain", err)
	}
	// The status argument echoes the current watermark, so even a
	// cancellation bug could not corrupt shared fixture state.
	status, err := r.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("ProcessingStatus: %v", err)
	}
	if err := r.UpdateProcessingStatus(cancelled, &status); !errors.Is(err, context.Canceled) {
		t.Errorf("UpdateProcessingStatus(cancelled ctx) = %v, want context.Canceled in chain", err)
	}

	// A closed pool yields an error, not a panic or exit. Build a throwaway
	// store on the same database so the shared pool stays usable.
	spare, err := spareStore(ctx)
	if err != nil {
		t.Fatalf("connecting spare store: %v", err)
	}
	spareRepo := NewRepo(spare)
	spare.Close()
	if _, err := spareRepo.BannedBids(ctx); err == nil {
		t.Error("BannedBids on closed pool succeeded, want error")
	}
	if _, _, err := spareRepo.BannedBidsPage(ctx, nil, 1); err == nil {
		t.Error("BannedBidsPage on closed pool succeeded, want error")
	}
	if _, err := spareRepo.BidderAddressForBid(ctx, 2002); err == nil {
		t.Error("BidderAddressForBid on closed pool succeeded, want error")
	}
	if _, err := spareRepo.CreateBannedBid(ctx, 999_995,
		"0x2100000000000000000000000000000000000021", time.Now()); err == nil {
		t.Error("CreateBannedBid on closed pool succeeded, want error")
	}
	if _, err := spareRepo.RemoveBannedBid(ctx, 2002); err == nil {
		t.Error("RemoveBannedBid on closed pool succeeded, want error")
	}
	if _, err := spareRepo.ContractAddrs(ctx); err == nil {
		t.Error("ContractAddrs on closed pool succeeded, want error")
	}
}
