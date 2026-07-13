//go:build integration

package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

const testContract = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

// seedDB installs one contract with three evt_log rows: two ddf252ad events
// in block 100 and one 8c5be1e5 event in block 100.
func seedDB(t *testing.T, db *testdb.DB) {
	t.Helper()
	ctx := context.Background()
	for _, stmt := range []string{
		`INSERT INTO block(block_hash, parent_hash, block_num, ts, num_tx) VALUES ('0xb100','0xb099',100,NOW(),1)`,
		`INSERT INTO address(block_num, tx_id, addr) VALUES (100, 0, '` + testContract + `')`,
		`INSERT INTO transaction(block_num, tx_hash) VALUES (100, '0xt1')`,
		`INSERT INTO evt_log(block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
		 SELECT 100, tx.id, a.address_id, v.sig, v.log_index, ''::bytea
		 FROM (VALUES ('ddf252ad', 0), ('ddf252ad', 1), ('8c5be1e5', 2)) AS v(sig, log_index)
		 JOIN transaction tx ON tx.tx_hash = '0xt1'
		 JOIN address a ON a.addr = '` + testContract + `'`,
	} {
		if _, err := db.Pool.Exec(ctx, stmt); err != nil {
			t.Fatalf("seeding: %v", err)
		}
	}
}

func writeJSONL(t *testing.T, lines ...string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "events.jsonl")
	if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestRunIntegrationPass(t *testing.T) {
	db := testdb.New(t)
	seedDB(t, db)

	// The freezer extract matches the database exactly.
	input := writeJSONL(t,
		`{"blockNumber":100,"contract":"`+testContract+`","topic0":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}`,
		`{"blockNumber":100,"contract":"`+testContract+`","topic0":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}`,
		`{"blockNumber":100,"contract":"`+testContract+`","topic0":"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"}`,
	)

	var out, errOut strings.Builder
	passed, err := run(context.Background(), []string{"--input", input, "--db", db.ConnString}, &out, &errOut)
	if err != nil {
		t.Fatal(err)
	}
	if !passed {
		t.Fatalf("verification failed:\n%s\n%s", out.String(), errOut.String())
	}
	if !strings.Contains(out.String(), "Matching:          2") {
		t.Fatalf("output = %q", out.String())
	}
	if !strings.Contains(out.String(), "Match rate: 100.00%") {
		t.Fatalf("output = %q", out.String())
	}
}

func TestRunIntegrationFailsBelowThreshold(t *testing.T) {
	db := testdb.New(t)
	seedDB(t, db)

	// The freezer extract misses the 8c5be1e5 event and carries an extra
	// unknown one: 1 of 2 DB keys match => 50% < 99%.
	input := writeJSONL(t,
		`{"blockNumber":100,"contract":"`+testContract+`","topic0":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}`,
		`{"blockNumber":100,"contract":"`+testContract+`","topic0":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}`,
		`{"blockNumber":105,"contract":"`+testContract+`","topic0":"0x1111111111111111111111111111111111111111111111111111111111111111"}`,
	)

	var out, errOut strings.Builder
	passed, err := run(context.Background(), []string{"--input", input, "--db", db.ConnString, "--verbose"}, &out, &errOut)
	if err != nil {
		t.Fatal(err)
	}
	if passed {
		t.Fatalf("verification passed unexpectedly:\n%s", out.String())
	}
	report := out.String()
	if !strings.Contains(report, "Missing (in DB, not in freezer): 1") {
		t.Fatalf("report = %q", report)
	}
	if !strings.Contains(report, "Extra (in freezer, not in DB):   1") {
		t.Fatalf("report = %q", report)
	}
	if !strings.Contains(report, "Missing from freezer (sample):") {
		t.Fatalf("report = %q", report)
	}
}

func TestRunIntegrationErrors(t *testing.T) {
	db := testdb.New(t)
	seedDB(t, db)
	ctx := context.Background()

	t.Run("empty input file", func(t *testing.T) {
		input := writeJSONL(t, "")
		var out, errOut strings.Builder
		_, err := run(ctx, []string{"--input", input, "--db", db.ConnString}, &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "no events found in input file") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("unknown contracts", func(t *testing.T) {
		input := writeJSONL(t,
			`{"blockNumber":100,"contract":"0xdeadbeef00000000000000000000000000000000","topic0":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}`,
		)
		var out, errOut strings.Builder
		_, err := run(ctx, []string{"--input", input, "--db", db.ConnString}, &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "none of the contracts found") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("bad connection string", func(t *testing.T) {
		input := writeJSONL(t,
			`{"blockNumber":100,"contract":"`+testContract+`","topic0":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}`,
		)
		var out, errOut strings.Builder
		_, err := run(ctx, []string{"--input", input, "--db", "postgres://nobody:wrong@127.0.0.1:1/nothing"}, &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "failed to connect to database") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("bad JSON input", func(t *testing.T) {
		input := writeJSONL(t, `{broken`)
		var out, errOut strings.Builder
		_, err := run(ctx, []string{"--input", input, "--db", db.ConnString}, &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "failed to load events") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("address table unreachable", func(t *testing.T) {
		// An empty search_path hides the address table: the contract-aid
		// query itself fails.
		input := writeJSONL(t,
			`{"blockNumber":100,"contract":"`+testContract+`","topic0":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}`,
		)
		var out, errOut strings.Builder
		_, err := run(ctx, []string{"--input", input, "--db", db.ConnString + "&options=-csearch_path%3Dpg_temp"}, &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "failed to query address table") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("event table missing", func(t *testing.T) {
		// A valid identifier that names no table passes validation and
		// fails at query time.
		input := writeJSONL(t,
			`{"blockNumber":100,"contract":"`+testContract+`","topic0":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}`,
		)
		var out, errOut strings.Builder
		_, err := run(ctx, []string{"--input", input, "--db", db.ConnString, "--table", "no_such_table"}, &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "failed to query events") {
			t.Fatalf("err = %v", err)
		}
	})
}
