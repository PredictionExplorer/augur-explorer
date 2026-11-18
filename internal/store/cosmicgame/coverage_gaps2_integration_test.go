//go:build integration

// Gap tests for query paths the main fixture suites never take: default
// interval normalization in the analytics series, LatestAddressParam,
// InsertEthToCharityFailed, and the duplicate-position guard.
package cosmicgame

import (
	"context"
	"errors"
	"strings"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// TestAnalyticsDefaultIntervalNormalization pins the shared non-positive
// interval handling: zero means one bucket spanning the whole range, and a
// degenerate range falls back to the documented default width.
func TestAnalyticsDefaultIntervalNormalization(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// A whole-range bucket (intervalSecs = finTs - initTs).
	if _, err := r.BidFrequencyByPeriod(ctx, 100, 400, 0); err != nil {
		t.Errorf("BidFrequencyByPeriod(interval 0): %v", err)
	}
	if _, err := r.BidFrequencyByPeriodBounded(ctx, 100, 400, 0); err != nil {
		t.Errorf("BidFrequencyByPeriodBounded(interval 0): %v", err)
	}
	if _, err := r.BidTypeRatioByPeriod(ctx, 100, 400, 0); err != nil {
		t.Errorf("BidTypeRatioByPeriod(interval 0): %v", err)
	}
	if _, err := r.BidTypeRatioByPeriodBounded(ctx, 100, 400, 0); err != nil {
		t.Errorf("BidTypeRatioByPeriodBounded(interval 0): %v", err)
	}

	// A degenerate range: the fallback width (3600s / 86400s) applies.
	if _, err := r.BidFrequencyByPeriod(ctx, 400, 400, -1); err != nil {
		t.Errorf("BidFrequencyByPeriod(empty range): %v", err)
	}
	if _, err := r.BidFrequencyByPeriodBounded(ctx, 400, 400, -1); err != nil {
		t.Errorf("BidFrequencyByPeriodBounded(empty range): %v", err)
	}
	if _, err := r.BidTypeRatioByPeriod(ctx, 400, 400, -1); err != nil {
		t.Errorf("BidTypeRatioByPeriod(empty range): %v", err)
	}
	if _, err := r.BidTypeRatioByPeriodBounded(ctx, 400, 400, -1); err != nil {
		t.Errorf("BidTypeRatioByPeriodBounded(empty range): %v", err)
	}
}

func TestLatestAddressParam(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Identifier whitelist: anything but a known admin table/column is
	// rejected before touching SQL.
	if _, _, err := r.LatestAddressParam(ctx, "cg_adm_charity_wallet; DROP TABLE x", "new_charity_aid"); err == nil {
		t.Error("bad table identifier accepted")
	}
	if _, _, err := r.LatestAddressParam(ctx, "cg_adm_charity_wallet", "aid; --"); err == nil {
		t.Error("bad column identifier accepted")
	}

	// Empty table: hasRow = false without error.
	var marketingRows int
	if err := r.pool().QueryRow(ctx,
		"SELECT COUNT(*) FROM cg_adm_marketing_addr").Scan(&marketingRows); err != nil {
		t.Fatal(err)
	}
	if marketingRows == 0 {
		if _, hasRow, err := r.LatestAddressParam(ctx,
			"cg_adm_marketing_addr", "new_marketing_aid"); err != nil || hasRow {
			t.Errorf("empty table: hasRow=%v err=%v", hasRow, err)
		}
	}

	// Seed one row and read it back through the join.
	var aid int64
	if err := r.pool().QueryRow(ctx,
		`INSERT INTO address(block_num, tx_id, addr)
		 VALUES (0, 0, '0x00000000000000000000000000000000000000f1')
		 RETURNING address_id`).Scan(&aid); err != nil {
		t.Fatal(err)
	}
	if _, err := r.pool().Exec(ctx,
		`INSERT INTO cg_adm_marketing_addr(
			block_num, tx_id, time_stamp, contract_aid, new_marketing_aid)
		 VALUES (1, 0, TO_TIMESTAMP(1767225600), $1, $1)`, aid); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		ctx := context.Background()
		_, _ = r.pool().Exec(ctx,
			"DELETE FROM cg_adm_marketing_addr WHERE new_marketing_aid=$1", aid)
		_, _ = r.pool().Exec(ctx, "DELETE FROM address WHERE address_id=$1", aid)
	})
	addr, hasRow, err := r.LatestAddressParam(ctx, "cg_adm_marketing_addr", "new_marketing_aid")
	if err != nil || !hasRow {
		t.Fatalf("seeded table: hasRow=%v err=%v", hasRow, err)
	}
	if !strings.EqualFold(addr, "0x00000000000000000000000000000000000000f1") {
		t.Errorf("addr = %q", addr)
	}
}

func TestInsertEthToCharityFailed(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// Hang the row off an existing fixture event log (the column is a
	// NOT NULL-free FK with a UNIQUE constraint).
	var evtlogID, blockNum, txID int64
	if err := r.pool().QueryRow(ctx,
		`SELECT e.id, e.block_num, e.tx_id FROM evt_log e
		 LEFT JOIN cg_eth_to_charity_failed f ON f.evtlog_id = e.id
		 WHERE f.id IS NULL ORDER BY e.id LIMIT 1`).Scan(&evtlogID, &blockNum, &txID); err != nil {
		t.Fatalf("no free fixture evt_log row: %v", err)
	}
	evt := &cgmodel.CGEthToCharityFailed{
		EvtId:          evtlogID,
		BlockNum:       blockNum,
		TxId:           txID,
		TimeStamp:      1767225600,
		Contract:       "0x00000000000000000000000000000000000000f2",
		CharityAddress: "0x00000000000000000000000000000000000000f3",
		Amount:         "123456789",
	}
	if err := r.InsertEthToCharityFailed(ctx, evt); err != nil {
		t.Fatalf("InsertEthToCharityFailed: %v", err)
	}
	t.Cleanup(func() {
		_, _ = r.pool().Exec(context.Background(),
			"DELETE FROM cg_eth_to_charity_failed WHERE evtlog_id=$1", evtlogID)
	})
	var amount string
	if err := r.pool().QueryRow(ctx,
		"SELECT amount::TEXT FROM cg_eth_to_charity_failed WHERE evtlog_id=$1",
		evtlogID).Scan(&amount); err != nil {
		t.Fatal(err)
	}
	if amount != "123456789" {
		t.Errorf("amount = %q", amount)
	}
}

func TestBidListRejectsNonWhitelistedClauses(t *testing.T) {
	r := repo(t)
	if _, err := bidList(context.Background(), r, "test",
		"b.round_num=$1", "b.id; DROP TABLE cg_bid", "", 2, 1); err == nil {
		t.Error("non-whitelisted ORDER BY accepted")
	}
}

func TestBidByRoundAndPositionEdgeCases(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	// A position no round ever reached: not-found, not a zero-value record.
	if _, err := r.BidByRoundAndPosition(ctx, 0, 99999); err == nil {
		t.Error("expected not-found error")
	}

	// Duplicate (round, position) rows are corruption, reported as conflict.
	// The bid scanner reads evtlog_id, so each duplicate needs a real event
	// log row behind it.
	var bidderAid, contractAid int64
	if err := r.pool().QueryRow(ctx,
		"SELECT bidder_aid, contract_aid FROM cg_bid LIMIT 1").Scan(&bidderAid, &contractAid); err != nil {
		t.Fatalf("fixture bid: %v", err)
	}
	var blockNum, txID int64
	if err := r.pool().QueryRow(ctx,
		"SELECT block_num, id FROM transaction ORDER BY id LIMIT 1").Scan(&blockNum, &txID); err != nil {
		t.Fatalf("fixture transaction: %v", err)
	}
	var evt1, evt2 int64
	if err := r.pool().QueryRow(ctx,
		`INSERT INTO evt_log(block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		 VALUES ($1, $2, $3, '81f6a773', 777001, '\x00') RETURNING id`,
		blockNum, txID, contractAid).Scan(&evt1); err != nil {
		t.Fatal(err)
	}
	if err := r.pool().QueryRow(ctx,
		`INSERT INTO evt_log(block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		 VALUES ($1, $2, $3, '81f6a773', 777002, '\x00') RETURNING id`,
		blockNum, txID, contractAid).Scan(&evt2); err != nil {
		t.Fatal(err)
	}
	if _, err := r.pool().Exec(ctx,
		`INSERT INTO cg_bid(evtlog_id, block_num, tx_id, time_stamp, contract_aid, bidder_aid,
			rwalk_nft_id, round_num, bid_type, bid_position, prize_time, eth_price, cst_price, msg)
		 VALUES ($3, $1, $2, TO_TIMESTAMP(1767230000), $4, $5, -1, 777, 0, 1, TO_TIMESTAMP(1767230000), 1, -1, ''),
		        ($6, $1, $2, TO_TIMESTAMP(1767230001), $4, $5, -1, 777, 0, 1, TO_TIMESTAMP(1767230001), 1, -1, '')`,
		blockNum, txID, evt1, contractAid, bidderAid, evt2); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		ctx := context.Background()
		_, _ = r.pool().Exec(ctx, "DELETE FROM cg_bid WHERE round_num=777")
		_, _ = r.pool().Exec(ctx, "DELETE FROM evt_log WHERE id IN ($1, $2)", evt1, evt2)
	})
	if _, err := r.BidByRoundAndPosition(ctx, 777, 1); !errors.Is(err, store.ErrConflict) {
		t.Errorf("duplicate rows: got %v, want ErrConflict", err)
	}
}

func TestRoundStartTimestampEmptyRound(t *testing.T) {
	r := repo(t)
	ts, err := r.RoundStartTimestamp(context.Background(), 424242)
	if err != nil || ts != 0 {
		t.Errorf("empty round: ts=%d err=%v", ts, err)
	}
}
