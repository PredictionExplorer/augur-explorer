//go:build integration

package cosmicgame

import "testing"

func TestGetBannedBids(t *testing.T) {
	sw := store(t)
	golden(t, "banned_bids", func() any {
		return sw.Get_banned_bids()
	})
}

// TestBannedBidInsertDeleteRoundTrip exercises the two error-returning write
// functions and restores the fixture state so golden-based tests are
// unaffected regardless of execution order.
func TestBannedBidInsertDeleteRoundTrip(t *testing.T) {
	sw := store(t)
	const bidID = 999_999
	const userAddr = "0x2100000000000000000000000000000000000021"

	before := len(sw.Get_banned_bids())
	if err := sw.Insert_banned_bid(bidID, userAddr); err != nil {
		t.Fatalf("Insert_banned_bid: %v", err)
	}
	// Ensure restoration even if the assertions below fail.
	t.Cleanup(func() {
		if err := sw.Delete_banned_bid_by_bid_id(bidID); err != nil {
			t.Errorf("cleanup Delete_banned_bid_by_bid_id: %v", err)
		}
	})

	list := sw.Get_banned_bids()
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

	if err := sw.Delete_banned_bid_by_bid_id(bidID); err != nil {
		t.Fatalf("Delete_banned_bid_by_bid_id: %v", err)
	}
	if got := len(sw.Get_banned_bids()); got != before {
		t.Fatalf("after delete: got %d banned bids, want %d", got, before)
	}
}
