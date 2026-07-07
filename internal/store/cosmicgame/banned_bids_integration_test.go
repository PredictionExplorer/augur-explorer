//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"
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
	if _, err := spareRepo.ContractAddrs(ctx); err == nil {
		t.Error("ContractAddrs on closed pool succeeded, want error")
	}
}
