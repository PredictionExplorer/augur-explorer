//go:build integration

package verify

import (
	"context"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

const (
	contractA = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	contractB = "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
)

// seedEvents installs two contracts and a handful of evt_log rows:
//
//	block 100, topic ddf252ad, contract A: 2 events
//	block 100, topic 8c5be1e5, contract A: 1 event
//	block 101, topic ddf252ad, contract B: 1 event
func seedEvents(t *testing.T, db *testdb.DB) {
	t.Helper()
	ctx := context.Background()

	_, err := db.Pool.Exec(ctx, `
		INSERT INTO block(block_hash, parent_hash, block_num, ts, num_tx)
		VALUES ('0xb100','0xb099',100,NOW(),1), ('0xb101','0xb100',101,NOW(),1)`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Pool.Exec(ctx, `
		INSERT INTO address(block_num, tx_id, addr)
		VALUES (100, 0, $1), (100, 0, $2)`, contractA, contractB)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Pool.Exec(ctx, `
		INSERT INTO transaction(block_num, tx_hash) VALUES (100, '0xt1'), (101, '0xt2')`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Pool.Exec(ctx, `
		INSERT INTO evt_log(block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		SELECT v.block_num, tx.id, a.address_id, v.sig, v.log_index, ''::bytea
		FROM (VALUES
			(100, '0xt1', $1::text, 'ddf252ad', 0),
			(100, '0xt1', $1::text, 'ddf252ad', 1),
			(100, '0xt1', $1::text, '8c5be1e5', 2),
			(101, '0xt2', $2::text, 'ddf252ad', 0)
		) AS v(block_num, tx_hash, addr, sig, log_index)
		JOIN transaction tx ON tx.tx_hash = v.tx_hash
		JOIN address a ON a.addr = v.addr`, contractA, contractB)
	if err != nil {
		t.Fatal(err)
	}
}

func TestContractAidsAndDBEventsIntegration(t *testing.T) {
	db := testdb.New(t)
	seedEvents(t, db)
	ctx := context.Background()

	aids, err := ContractAids(ctx, db.Pool, []string{contractA, contractB, "0xdeadbeef00000000000000000000000000000000"})
	if err != nil {
		t.Fatal(err)
	}
	if len(aids) != 2 {
		t.Fatalf("aids = %v", aids)
	}

	events, err := DBEvents(ctx, db.Pool, "evt_log", 100, 101, aids)
	if err != nil {
		t.Fatal(err)
	}
	want := map[Key]int{
		{BlockNum: 100, Topic0: "ddf252ad", Contract: contractA}: 2,
		{BlockNum: 100, Topic0: "8c5be1e5", Contract: contractA}: 1,
		{BlockNum: 101, Topic0: "ddf252ad", Contract: contractB}: 1,
	}
	if len(events) != len(want) {
		t.Fatalf("events = %v", events)
	}
	for key, count := range want {
		if events[key] != count {
			t.Fatalf("events[%+v] = %d, want %d (all: %v)", key, events[key], count, events)
		}
	}

	// Block range filtering.
	events, err = DBEvents(ctx, db.Pool, "evt_log", 100, 100, aids)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 2 {
		t.Fatalf("filtered events = %v", events)
	}

	// Contract filtering.
	events, err = DBEvents(ctx, db.Pool, "evt_log", 100, 101, aids[:1])
	if err != nil {
		t.Fatal(err)
	}
	for key := range events {
		if key.Contract == contractB && aids[0] != 2 {
			// aids order follows the address table; ensure only one
			// contract's rows are returned.
			t.Fatalf("events = %v, want a single contract", events)
		}
	}

	// Invalid table name is rejected before touching the database.
	if _, err := DBEvents(ctx, db.Pool, "evt;drop", 0, 1, aids); err == nil ||
		!strings.Contains(err.Error(), "invalid table name") {
		t.Fatalf("err = %v", err)
	}

	// A valid but missing table surfaces the database error.
	if _, err := DBEvents(ctx, db.Pool, "no_such_table", 0, 1, aids); err == nil {
		t.Fatal("missing table must error")
	}

	// Cancellation propagates.
	cancelled, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := ContractAids(cancelled, db.Pool, []string{contractA}); err == nil {
		t.Fatal("cancelled ContractAids must error")
	}
	if _, err := DBEvents(cancelled, db.Pool, "evt_log", 0, 1, aids); err == nil {
		t.Fatal("cancelled DBEvents must error")
	}
}
