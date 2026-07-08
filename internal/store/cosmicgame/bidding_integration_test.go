//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestBidIDByEvtlog(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	got, err := r.BidIDByEvtlog(ctx, 5004)
	if err != nil {
		t.Fatalf("BidIDByEvtlog(5004): %v", err)
	}
	if got != 2001 {
		t.Errorf("bid id for evtlog 5004: got %d, want 2001", got)
	}
	if _, err := r.BidIDByEvtlog(ctx, 999_999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("BidIDByEvtlog(999999) = %v, want ErrNotFound", err)
	}
}

func TestBids(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "bids", func() any {
		recs, err := r.Bids(ctx, 0, 100)
		if err != nil {
			t.Fatalf("Bids: %v", err)
		}
		return recs
	})
	golden(t, "bids_paged", func() any {
		recs, err := r.Bids(ctx, 5, 3)
		if err != nil {
			t.Fatalf("Bids(paged): %v", err)
		}
		return recs
	})
	// limit=0 means "no limit".
	all, err := r.Bids(ctx, 0, 0)
	if err != nil {
		t.Fatalf("Bids(no limit): %v", err)
	}
	if len(all) != 12 {
		t.Errorf("expected 12 bids with limit=0, got %d", len(all))
	}
}

func TestBidInfo(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "bid_info_5004", func() any {
		rec, err := r.BidInfo(ctx, 5004)
		if err != nil {
			t.Fatalf("BidInfo(5004): %v", err)
		}
		return rec
	})
	// carol's CST bid carries the v2 reward/duration columns.
	golden(t, "bid_info_5008_cst", func() any {
		rec, err := r.BidInfo(ctx, 5008)
		if err != nil {
			t.Fatalf("BidInfo(5008): %v", err)
		}
		return rec
	})
	if _, err := r.BidInfo(ctx, 999_999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("BidInfo(999999) = %v, want ErrNotFound", err)
	}
}

func TestEvtlogIDByRoundAndBidPosition(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	evtlogID, err := r.EvtlogIDByRoundAndBidPosition(ctx, 0, 1)
	if err != nil || evtlogID != 5004 {
		t.Errorf("round 0 position 1: got (%d,%v), want (5004,nil)", evtlogID, err)
	}
	if _, err := r.EvtlogIDByRoundAndBidPosition(ctx, 0, 99); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("EvtlogIDByRoundAndBidPosition(0,99) = %v, want ErrNotFound", err)
	}
}

func TestBidsWithMessageByRound(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "bids_with_message_by_round_0_asc", func() any {
		recs, err := r.BidsWithMessageByRound(ctx, 0, false, 0, 100)
		if err != nil {
			t.Fatalf("BidsWithMessageByRound(asc): %v", err)
		}
		return recs
	})
	golden(t, "bids_with_message_by_round_0_desc", func() any {
		recs, err := r.BidsWithMessageByRound(ctx, 0, true, 0, 100)
		if err != nil {
			t.Fatalf("BidsWithMessageByRound(desc): %v", err)
		}
		return recs
	})
}

func TestBidsByRound(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "bids_by_round_num_0_asc", func() any {
		recs, total, err := r.BidsByRound(ctx, 0, 0, 0, 100)
		if err != nil {
			t.Fatalf("BidsByRound(asc): %v", err)
		}
		return map[string]any{"records": recs, "total": total}
	})
	golden(t, "bids_by_round_num_0_desc_paged", func() any {
		recs, total, err := r.BidsByRound(ctx, 0, 1, 1, 2)
		if err != nil {
			t.Fatalf("BidsByRound(desc paged): %v", err)
		}
		return map[string]any{"records": recs, "total": total}
	})
}

func TestBidCountForRound(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	got, err := r.BidCountForRound(ctx, 0)
	if err != nil {
		t.Fatalf("BidCountForRound(0): %v", err)
	}
	if got != 4 {
		t.Errorf("bid count round 0: got %d, want 4", got)
	}
	got, err = r.BidCountForRound(ctx, 999)
	if err != nil {
		t.Fatalf("BidCountForRound(999): %v", err)
	}
	if got != 0 {
		t.Errorf("bid count round 999: got %d, want 0", got)
	}
}

func TestLastCstBidEvtlogForBidder(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	const carol = "0x2300000000000000000000000000000000000023"
	evtlogID, err := r.LastCstBidEvtlogForBidder(ctx, 0, carol)
	if err != nil || evtlogID != 5008 {
		t.Errorf("carol's last CST bid in round 0: got (%d,%v), want (5008,nil)", evtlogID, err)
	}
	// Case-insensitive address matching.
	if _, err := r.LastCstBidEvtlogForBidder(ctx, 0, "0X2300000000000000000000000000000000000023"); err != nil {
		t.Errorf("expected case-insensitive address match, got %v", err)
	}
	// alice never CST-bid in round 0.
	if _, err := r.LastCstBidEvtlogForBidder(ctx, 0, "0x2100000000000000000000000000000000000021"); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("LastCstBidEvtlogForBidder(alice) = %v, want ErrNotFound", err)
	}
}

func TestRoundStartTimestamp(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	got, err := r.RoundStartTimestamp(ctx, 0)
	if err != nil {
		t.Fatalf("RoundStartTimestamp(0): %v", err)
	}
	if got != 1767225700 {
		t.Errorf("round 0 start: got %d, want 1767225700", got)
	}
	got, err = r.RoundStartTimestamp(ctx, 999)
	if err != nil {
		t.Fatalf("RoundStartTimestamp(999): %v", err)
	}
	if got != 0 {
		t.Errorf("round 999 start: got %d, want 0", got)
	}
}

func TestRandomWalkTokensUsedInBids(t *testing.T) {
	r := repo(t)
	golden(t, "random_walk_tokens_in_bids", func() any {
		recs, err := r.RandomWalkTokensUsedInBids(context.Background())
		if err != nil {
			t.Fatalf("RandomWalkTokensUsedInBids: %v", err)
		}
		return recs
	})
}

func TestBidRowIDByEvtlogID(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "cosmic_game_bid_by_evtlog_id_5004", func() any {
		id, err := r.BidRowIDByEvtlogID(ctx, 5004)
		if err != nil {
			t.Fatalf("BidRowIDByEvtlogID(5004): %v", err)
		}
		return id
	})
	// A missing bid means "pure Donate() transaction": 0, no error.
	id, err := r.BidRowIDByEvtlogID(ctx, 999_999)
	if err != nil {
		t.Fatalf("BidRowIDByEvtlogID(999999): %v", err)
	}
	if id != 0 {
		t.Errorf("bid row id for missing evtlog: got %d, want 0", id)
	}
}
