//go:build integration

package randomwalk

import (
	"context"
	"testing"

	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
)

// TestProcessingStatusRoundTrip covers ProcessingStatus (which lazily
// inserts the default row) and the update path, restoring the original
// watermark afterwards.
func TestProcessingStatusRoundTrip(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	initial, err := r.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("ProcessingStatus: %v", err)
	}
	t.Cleanup(func() {
		if err := r.UpdateProcessingStatus(context.Background(), &initial); err != nil {
			t.Errorf("restoring processing status: %v", err)
		}
	})

	want := rwp.ProcStatus{LastIdProcessed: 5095, LastBlockNum: 141}
	if err := r.UpdateProcessingStatus(ctx, &want); err != nil {
		t.Fatalf("UpdateProcessingStatus: %v", err)
	}

	got, err := r.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("ProcessingStatus after update: %v", err)
	}
	if got != want {
		t.Fatalf("processing status round trip: got %+v, want %+v", got, want)
	}
}

func TestContractAddrs(t *testing.T) {
	r := repo(t)
	golden(t, "randomwalk_contract_addresses", func() any {
		addrs, err := r.ContractAddrs(context.Background())
		if err != nil {
			t.Fatalf("ContractAddrs: %v", err)
		}
		return addrs
	})
}

func TestRawContractAddrs(t *testing.T) {
	r := repo(t)
	marketplace, randomwalk, err := r.RawContractAddrs(context.Background())
	if err != nil {
		t.Fatalf("RawContractAddrs: %v", err)
	}
	if marketplace != addrMarketplace || randomwalk != addrRandomWalk {
		t.Errorf("raw contract addrs: got (%v, %v), want (%v, %v)",
			marketplace, randomwalk, addrMarketplace, addrRandomWalk)
	}
}

func TestRankingDataForAllUsers(t *testing.T) {
	r := repo(t)
	golden(t, "randomwalk_ranking_data_for_all_users", func() any {
		recs, err := r.RankingDataForAllUsers(context.Background())
		if err != nil {
			t.Fatalf("RankingDataForAllUsers: %v", err)
		}
		return recs
	})
}

func TestTopProfitMakers(t *testing.T) {
	r := repo(t)
	golden(t, "randomwalk_top_profit_makers", func() any {
		recs, err := r.TopProfitMakers(context.Background())
		if err != nil {
			t.Fatalf("TopProfitMakers: %v", err)
		}
		return recs
	})
}

func TestTopTradeMakers(t *testing.T) {
	r := repo(t)
	golden(t, "randomwalk_top_trade_makers", func() any {
		recs, err := r.TopTradeMakers(context.Background())
		if err != nil {
			t.Fatalf("TopTradeMakers: %v", err)
		}
		return recs
	})
}

func TestTopVolumeMakers(t *testing.T) {
	r := repo(t)
	golden(t, "randomwalk_top_volume_makers", func() any {
		recs, err := r.TopVolumeMakers(context.Background())
		if err != nil {
			t.Fatalf("TopVolumeMakers: %v", err)
		}
		return recs
	})
}

// TestUpdateRankRoundTrip exercises the three UPDATE-then-INSERT rank
// writers: the update path against carol's extension-seed row, and the
// insert path for alice (no rw_uranks row), restoring both afterwards.
func TestUpdateRankRoundTrip(t *testing.T) {
	r := repo(t)
	p := pool(t)
	ctx := context.Background()

	t.Cleanup(func() {
		// Restore carol's extension-seed row and drop alice's inserted rows.
		if _, err := p.Exec(ctx,
			"UPDATE rw_uranks SET total_trades=1, top_profit=1.0, top_trades=1.0, top_volume=1.0, profit=950000000000000000, volume=1000000000000000000 WHERE aid=$1",
			int64(aidCarol),
		); err != nil {
			t.Errorf("restoring carol's rank row: %v", err)
		}
		if _, err := p.Exec(ctx, "DELETE FROM rw_uranks WHERE aid=$1", int64(aidAlice)); err != nil {
			t.Errorf("removing alice's rank row: %v", err)
		}
	})

	// Update path: carol has a row; her values must change in place.
	if err := r.UpdateTopProfitRank(ctx, aidCarol, 2.5, 900000000000000000); err != nil {
		t.Errorf("UpdateTopProfitRank(carol): %v", err)
	}
	if err := r.UpdateTopTotalTradesRank(ctx, aidCarol, 2.5, 2); err != nil {
		t.Errorf("UpdateTopTotalTradesRank(carol): %v", err)
	}
	if err := r.UpdateTopVolumeRank(ctx, aidCarol, 2.5, 2000000000000000000); err != nil {
		t.Errorf("UpdateTopVolumeRank(carol): %v", err)
	}
	var carolRank float64
	var carolTrades int64
	if err := p.QueryRow(ctx,
		"SELECT top_trades, total_trades FROM rw_uranks WHERE aid=$1", int64(aidCarol)).Scan(&carolRank, &carolTrades); err != nil {
		t.Fatalf("reading carol's updated rank row: %v", err)
	}
	if carolRank != 2.5 || carolTrades != 2 {
		t.Errorf("carol's updated ranks: got (%v, %v), want (2.5, 2)", carolRank, carolTrades)
	}
	var carolRows int
	if err := p.QueryRow(ctx,
		"SELECT COUNT(*) FROM rw_uranks WHERE aid=$1", int64(aidCarol)).Scan(&carolRows); err != nil {
		t.Fatalf("counting carol's rank rows: %v", err)
	}
	if carolRows != 1 {
		t.Errorf("carol's rank rows: got %d, want 1 (update must not insert)", carolRows)
	}

	// Insert path: alice has no row; the writer must create one.
	if err := r.UpdateTopProfitRank(ctx, aidAlice, 75.0, 0); err != nil {
		t.Errorf("UpdateTopProfitRank(alice): %v", err)
	}
	var profit float64
	if err := p.QueryRow(ctx, "SELECT top_profit FROM rw_uranks WHERE aid=$1", int64(aidAlice)).Scan(&profit); err != nil {
		t.Fatalf("reading alice's inserted rank row: %v", err)
	}
	if profit != 75.0 {
		t.Errorf("alice's inserted top_profit: got %v, want 75.0", profit)
	}
}

func TestMintEventsForNotification(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "mint_events_for_notification", func() any {
		recs, err := r.MintEventsForNotification(ctx, aidRandomWalk, 1767228600)
		if err != nil {
			t.Fatalf("MintEventsForNotification: %v", err)
		}
		return recs
	})
	got, err := r.MintEventsForNotification(ctx, aidRandomWalk, 1767229000)
	if err != nil {
		t.Fatalf("MintEventsForNotification(late): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no mints after the last fixture mint, got %d", len(got))
	}
}

// TestMessagingStatusRoundTrip pins the notibot watermark semantics:
// migration 00008 seeds the singleton row (previously the table was empty,
// the plain-UPDATE writer never persisted anything, and every notibot
// restart re-notified the full history).
func TestMessagingStatusRoundTrip(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	initial, err := r.MessagingStatus(ctx)
	if err != nil {
		t.Fatalf("MessagingStatus: %v", err)
	}
	t.Cleanup(func() {
		if err := r.UpdateMessagingStatus(context.Background(), &initial); err != nil {
			t.Errorf("restoring messaging status: %v", err)
		}
	})

	want := rwp.MsgStatus{TxId: 1036, EvtLogId: 5089, BlockNum: 136, TimeStamp: 1767229200}
	if err := r.UpdateMessagingStatus(ctx, &want); err != nil {
		t.Fatalf("UpdateMessagingStatus: %v", err)
	}

	got, err := r.MessagingStatus(ctx)
	if err != nil {
		t.Fatalf("MessagingStatus after update: %v", err)
	}
	if got != want {
		t.Fatalf("messaging status round trip: got %+v, want %+v (watermark not persisted?)", got, want)
	}
}

func TestAllEventsForNotification(t *testing.T) {
	r := repo(t)
	golden(t, "all_events_for_notification", func() any {
		recs, err := r.AllEventsForNotification(context.Background(), aidRandomWalk, 1767228000)
		if err != nil {
			t.Fatalf("AllEventsForNotification: %v", err)
		}
		return recs
	})
}

func TestAllEventsForNotificationSinceEvtlog(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "all_events_for_notification2", func() any {
		recs, err := r.AllEventsForNotificationSinceEvtlog(ctx, aidRandomWalk, 5080)
		if err != nil {
			t.Fatalf("AllEventsForNotificationSinceEvtlog: %v", err)
		}
		return recs
	})
	got, err := r.AllEventsForNotificationSinceEvtlog(ctx, aidRandomWalk, 999_999)
	if err != nil {
		t.Fatalf("AllEventsForNotificationSinceEvtlog(late): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no notification events past the last evtlog, got %d", len(got))
	}
}

func TestAllEventsForNotificationMintsOnly(t *testing.T) {
	r := repo(t)
	golden(t, "all_events_for_notification_test", func() any {
		recs, err := r.AllEventsForNotificationMintsOnly(context.Background(), aidRandomWalk, 1767228000)
		if err != nil {
			t.Fatalf("AllEventsForNotificationMintsOnly: %v", err)
		}
		return recs
	})
}

func TestServerTimestamp(t *testing.T) {
	r := repo(t)
	// now()-backed: assert sanity, not a golden.
	got, err := r.ServerTimestamp(context.Background())
	if err != nil {
		t.Fatalf("ServerTimestamp: %v", err)
	}
	if got <= 1767225600 {
		t.Errorf("server timestamp %d is before the fixture epoch", got)
	}
}

func TestLastMintTimestamp(t *testing.T) {
	r := repo(t)
	got, err := r.LastMintTimestamp(context.Background())
	if err != nil {
		t.Fatalf("LastMintTimestamp: %v", err)
	}
	if got != 1767228900 {
		t.Errorf("last mint timestamp: got %d, want 1767228900", got)
	}
}

func TestTokenTransfersByTxHash(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "rw_token_transfers_by_tx_hash", func() any {
		recs, err := r.TokenTransfersByTxHash(ctx, "0xf000000000000000000000000000000000000000000000000000000000001036")
		if err != nil {
			t.Fatalf("TokenTransfersByTxHash: %v", err)
		}
		return recs
	})
	got, err := r.TokenTransfersByTxHash(ctx, "0x"+"00"+"ff")
	if err != nil {
		t.Fatalf("TokenTransfersByTxHash(unknown): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no transfers for unknown tx hash, got %d", len(got))
	}
}

func TestOfferExists(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cases := []struct {
		contract string
		offerID  int64
		want     bool
		desc     string
	}{
		{addrMarketplace, 1, true, "offer 1 must exist"},
		{addrMarketplace, 999, false, "offer 999 must be missing"},
		{"0xdead000000000000000000000000000000000000", 1, false, "unknown contract has no offers"},
	}
	for _, c := range cases {
		got, err := r.OfferExists(ctx, c.contract, c.offerID)
		if err != nil {
			t.Fatalf("OfferExists(%v, %d): %v", c.contract, c.offerID, err)
		}
		if got != c.want {
			t.Errorf("%s: got %v", c.desc, got)
		}
	}
}

func TestTokenExists(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	got, err := r.TokenExists(ctx, addrRandomWalk, 10)
	if err != nil {
		t.Fatalf("TokenExists(10): %v", err)
	}
	if !got {
		t.Error("expected token 10 to exist")
	}
	got, err = r.TokenExists(ctx, addrRandomWalk, 999)
	if err != nil {
		t.Fatalf("TokenExists(999): %v", err)
	}
	if got {
		t.Error("expected token 999 to be missing")
	}
}
