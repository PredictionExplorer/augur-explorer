//go:build integration

package cosmicgame

import "testing"

func TestGetBidIdByEvtlog(t *testing.T) {
	sw := wrapper(t)
	if got := sw.Get_bid_id_by_evtlog(5004); got != 2001 {
		t.Errorf("bid id for evtlog 5004: got %d, want 2001", got)
	}
	if got := sw.Get_bid_id_by_evtlog(999_999); got != -1 {
		t.Errorf("bid id for missing evtlog: got %d, want -1", got)
	}
}

func TestGetBids(t *testing.T) {
	sw := wrapper(t)
	golden(t, "bids", func() any {
		return sw.Get_bids(0, 100)
	})
	golden(t, "bids_paged", func() any {
		return sw.Get_bids(5, 3)
	})
	// limit=0 means "no limit".
	if all := sw.Get_bids(0, 0); len(all) != 12 {
		t.Errorf("expected 12 bids with limit=0, got %d", len(all))
	}
}

func TestGetBidInfo(t *testing.T) {
	sw := wrapper(t)
	golden(t, "bid_info_5004", func() any {
		found, rec := sw.Get_bid_info(5004)
		if !found {
			t.Fatal("expected bid at evtlog 5004")
		}
		return rec
	})
	// carol's CST bid carries the v2 reward/duration columns.
	golden(t, "bid_info_5008_cst", func() any {
		found, rec := sw.Get_bid_info(5008)
		if !found {
			t.Fatal("expected bid at evtlog 5008")
		}
		return rec
	})
	if found, _ := sw.Get_bid_info(999_999); found {
		t.Error("expected no bid at evtlog 999999")
	}
}

func TestGetEvtlogIdByRoundAndBidPosition(t *testing.T) {
	sw := wrapper(t)
	evtlogID, found := sw.Get_evtlog_id_by_round_and_bid_position(0, 1)
	if !found || evtlogID != 5004 {
		t.Errorf("round 0 position 1: got (%d,%v), want (5004,true)", evtlogID, found)
	}
	if _, found := sw.Get_evtlog_id_by_round_and_bid_position(0, 99); found {
		t.Error("expected no bid at round 0 position 99")
	}
}

func TestGetBidsWithMessageByRound(t *testing.T) {
	sw := wrapper(t)
	golden(t, "bids_with_message_by_round_0_asc", func() any {
		return sw.Get_bids_with_message_by_round(0, false, 0, 100)
	})
	golden(t, "bids_with_message_by_round_0_desc", func() any {
		return sw.Get_bids_with_message_by_round(0, true, 0, 100)
	})
}

func TestGetBidsByRoundNum(t *testing.T) {
	sw := wrapper(t)
	golden(t, "bids_by_round_num_0_asc", func() any {
		recs, total := sw.Get_bids_by_round_num(0, 0, 0, 100)
		return map[string]any{"records": recs, "total": total}
	})
	golden(t, "bids_by_round_num_0_desc_paged", func() any {
		recs, total := sw.Get_bids_by_round_num(0, 1, 1, 2)
		return map[string]any{"records": recs, "total": total}
	})
}

func TestGetBidCountForRound(t *testing.T) {
	sw := wrapper(t)
	if got := sw.Get_bid_count_for_round(0); got != 4 {
		t.Errorf("bid count round 0: got %d, want 4", got)
	}
	if got := sw.Get_bid_count_for_round(999); got != 0 {
		t.Errorf("bid count round 999: got %d, want 0", got)
	}
}

func TestGetLastCstBidEvtlogForBidder(t *testing.T) {
	sw := wrapper(t)
	const carol = "0x2300000000000000000000000000000000000023"
	evtlogID, found := sw.Get_last_cst_bid_evtlog_for_bidder(0, carol)
	if !found || evtlogID != 5008 {
		t.Errorf("carol's last CST bid in round 0: got (%d,%v), want (5008,true)", evtlogID, found)
	}
	// Case-insensitive address matching.
	if _, found := sw.Get_last_cst_bid_evtlog_for_bidder(0, "0X2300000000000000000000000000000000000023"); !found {
		t.Error("expected case-insensitive address match")
	}
	// alice never CST-bid in round 0.
	if _, found := sw.Get_last_cst_bid_evtlog_for_bidder(0, "0x2100000000000000000000000000000000000021"); found {
		t.Error("expected no CST bid for alice in round 0")
	}
}

func TestGetRoundStartTimestamp(t *testing.T) {
	sw := wrapper(t)
	if got := sw.Get_round_start_timestamp(0); got != 1767225700 {
		t.Errorf("round 0 start: got %d, want 1767225700", got)
	}
	if got := sw.Get_round_start_timestamp(999); got != 0 {
		t.Errorf("round 999 start: got %d, want 0", got)
	}
}

func TestGetRandomWalkTokensInBids(t *testing.T) {
	sw := wrapper(t)
	golden(t, "random_walk_tokens_in_bids", func() any {
		return sw.Get_random_walk_tokens_in_bids()
	})
}

func TestGetCosmicGameBidByEvtlogId(t *testing.T) {
	sw := wrapper(t)
	golden(t, "cosmic_game_bid_by_evtlog_id_5004", func() any {
		return sw.Get_cosmic_game_bid_by_evtlog_id(5004)
	})
}
