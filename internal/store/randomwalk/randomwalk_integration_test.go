//go:build integration

package randomwalk

import (
	"testing"

	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
)

// TestProcessingStatusRoundTrip covers Get_randomwalk_processing_status
// (which lazily inserts the default row) and the update path, restoring the
// original watermark afterwards.
func TestProcessingStatusRoundTrip(t *testing.T) {
	sw := wrapper(t)

	initial := sw.Get_randomwalk_processing_status()
	t.Cleanup(func() { sw.Update_randomwalk_process_status(&initial) })

	want := rwp.ProcStatus{LastIdProcessed: 5095, LastBlockNum: 141}
	sw.Update_randomwalk_process_status(&want)

	if got := sw.Get_randomwalk_processing_status(); got != want {
		t.Fatalf("processing status round trip: got %+v, want %+v", got, want)
	}
}

func TestGetRandomwalkContractAddresses(t *testing.T) {
	sw := wrapper(t)
	golden(t, "randomwalk_contract_addresses", func() any {
		return sw.Get_randomwalk_contract_addresses()
	})
}

func TestGetRandomwalkRankingDataForAllUsers(t *testing.T) {
	sw := wrapper(t)
	golden(t, "randomwalk_ranking_data_for_all_users", func() any {
		return sw.Get_randomwalk_ranking_data_for_all_users()
	})
}

func TestGetRandomwalkTopProfitMakers(t *testing.T) {
	sw := wrapper(t)
	golden(t, "randomwalk_top_profit_makers", func() any {
		return sw.Get_randomwalk_top_profit_makers()
	})
}

func TestGetRandomwalkTopTradeMakers(t *testing.T) {
	sw := wrapper(t)
	golden(t, "randomwalk_top_trade_makers", func() any {
		return sw.Get_randomwalk_top_trade_makers()
	})
}

func TestGetRandomwalkTopVolumeMakers(t *testing.T) {
	sw := wrapper(t)
	golden(t, "randomwalk_top_volume_makers", func() any {
		return sw.Get_randomwalk_top_volume_makers()
	})
}

// TestUpdateRankRoundTrip exercises the three UPDATE-then-INSERT rank
// writers: the update path against carol's extension-seed row, and the
// insert path for alice (no rw_uranks row), restoring both afterwards.
func TestUpdateRankRoundTrip(t *testing.T) {
	sw := wrapper(t)
	db := sw.S.Db()

	t.Cleanup(func() {
		// Restore carol's extension-seed row and drop alice's inserted rows.
		if _, err := db.Exec(
			"UPDATE rw_uranks SET total_trades=1, top_profit=1.0, top_trades=1.0, top_volume=1.0, profit=950000000000000000, volume=1000000000000000000 WHERE aid=$1",
			int64(aidCarol),
		); err != nil {
			t.Errorf("restoring carol's rank row: %v", err)
		}
		if _, err := db.Exec("DELETE FROM rw_uranks WHERE aid=$1", int64(aidAlice)); err != nil {
			t.Errorf("removing alice's rank row: %v", err)
		}
	})

	// Update path: carol has a row, one row must be affected.
	if affected := sw.Update_randomwalk_top_profit_rank(aidCarol, 2.5, 900000000000000000); affected != 1 {
		t.Errorf("top_profit update for carol: got %d affected rows, want 1", affected)
	}
	if affected := sw.Update_randomwalk_top_total_trades_rank(aidCarol, 2.5, 2); affected != 1 {
		t.Errorf("top_trades update for carol: got %d affected rows, want 1", affected)
	}
	if affected := sw.Update_randomwalk_top_volume_rank(aidCarol, 2.5, 2000000000000000000); affected != 1 {
		t.Errorf("top_volume update for carol: got %d affected rows, want 1", affected)
	}

	// Insert path: alice has no row; affected=0 and an INSERT happens.
	if affected := sw.Update_randomwalk_top_profit_rank(aidAlice, 75.0, 0); affected != 0 {
		t.Errorf("top_profit insert for alice: got %d affected rows, want 0", affected)
	}
	var profit float64
	if err := db.QueryRow("SELECT top_profit FROM rw_uranks WHERE aid=$1", int64(aidAlice)).Scan(&profit); err != nil {
		t.Fatalf("reading alice's inserted rank row: %v", err)
	}
	if profit != 75.0 {
		t.Errorf("alice's inserted top_profit: got %v, want 75.0", profit)
	}
}

func TestGetMintEventsForNotification(t *testing.T) {
	sw := wrapper(t)
	golden(t, "mint_events_for_notification", func() any {
		return sw.Get_mint_events_for_notification(aidRandomWalk, 1767228600)
	})
	if got := sw.Get_mint_events_for_notification(aidRandomWalk, 1767229000); len(got) != 0 {
		t.Errorf("expected no mints after the last fixture mint, got %d", len(got))
	}
}

// TestMessagingStatusRoundTrip pins the notibot watermark semantics:
// migration 00008 seeds the singleton row (previously the table was empty,
// the plain-UPDATE writer never persisted anything, and every notibot
// restart re-notified the full history).
func TestMessagingStatusRoundTrip(t *testing.T) {
	sw := wrapper(t)

	initial := sw.Get_messaging_status()
	t.Cleanup(func() { sw.Update_messaging_status(&initial) })

	want := rwp.MsgStatus{TxId: 1036, EvtLogId: 5089, BlockNum: 136, TimeStamp: 1767229200}
	sw.Update_messaging_status(&want)

	got := sw.Get_messaging_status()
	if got != want {
		t.Fatalf("messaging status round trip: got %+v, want %+v (watermark not persisted?)", got, want)
	}
}

func TestGetAllEventsForNotification(t *testing.T) {
	sw := wrapper(t)
	golden(t, "all_events_for_notification", func() any {
		return sw.Get_all_events_for_notification(aidRandomWalk, 1767228000)
	})
}

func TestGetAllEventsForNotification2(t *testing.T) {
	sw := wrapper(t)
	golden(t, "all_events_for_notification2", func() any {
		return sw.Get_all_events_for_notification2(aidRandomWalk, 5080)
	})
	if got := sw.Get_all_events_for_notification2(aidRandomWalk, 999_999); len(got) != 0 {
		t.Errorf("expected no notification events past the last evtlog, got %d", len(got))
	}
}

func TestGetAllEventsForNotificationTest(t *testing.T) {
	sw := wrapper(t)
	golden(t, "all_events_for_notification_test", func() any {
		return sw.Get_all_events_for_notification_test(aidRandomWalk, 1767228000)
	})
}

func TestGetServerTimestamp(t *testing.T) {
	sw := wrapper(t)
	// now()-backed: assert sanity, not a golden.
	if got := sw.Get_server_timestamp(); got <= 1767225600 {
		t.Errorf("server timestamp %d is before the fixture epoch", got)
	}
}

func TestGetLastMintTimestamp(t *testing.T) {
	sw := wrapper(t)
	if got := sw.Get_last_mint_timestamp(); got != 1767228900 {
		t.Errorf("last mint timestamp: got %d, want 1767228900", got)
	}
}

func TestGetRwTokenTransfersByTxHash(t *testing.T) {
	sw := wrapper(t)
	golden(t, "rw_token_transfers_by_tx_hash", func() any {
		return sw.Get_rw_token_transfers_by_tx_hash("0xf000000000000000000000000000000000000000000000000000000000001036")
	})
	if got := sw.Get_rw_token_transfers_by_tx_hash("0x" + "00" + "ff"); len(got) != 0 {
		t.Errorf("expected no transfers for unknown tx hash, got %d", len(got))
	}
}

func TestOfferExists(t *testing.T) {
	sw := wrapper(t)
	if !sw.Offer_exists(addrMarketplace, 1) {
		t.Error("expected offer 1 to exist")
	}
	if sw.Offer_exists(addrMarketplace, 999) {
		t.Error("expected offer 999 to be missing")
	}
	if sw.Offer_exists("0xdead000000000000000000000000000000000000", 1) {
		t.Error("expected unknown contract to have no offers")
	}
}

func TestRWalkTokenExists(t *testing.T) {
	sw := wrapper(t)
	if !sw.RWalk_token_exists(addrRandomWalk, 10) {
		t.Error("expected token 10 to exist")
	}
	if sw.RWalk_token_exists(addrRandomWalk, 999) {
		t.Error("expected token 999 to be missing")
	}
}
