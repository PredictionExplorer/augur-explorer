//go:build integration

package txcollector

import (
	"context"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

func TestLoadEventRowsPostgresSuccessFilterCancelAndError(t *testing.T) {
	db := testdb.New(t)
	if err := testfixtures.Apply(t.Context(), db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}

	const contract = "0x2000000000000000000000000000000000000002"
	rows, err := LoadEventRows(t.Context(), db.Pool, []string{contract}, 127)
	if err != nil {
		t.Fatalf("LoadEventRows: %v", err)
	}
	if len(rows) == 0 {
		t.Fatal("filtered PostgreSQL query returned no rows")
	}
	for _, row := range rows {
		if row.BlockNum < 127 {
			t.Fatalf("row below fromBlock: %+v", row)
		}
		if row.ContractAddr != common.HexToAddress(contract).Hex() {
			t.Fatalf("row from another contract: %+v", row)
		}
	}

	empty, err := LoadEventRows(t.Context(), db.Pool, []string{contract}, 1_000_000)
	if err != nil {
		t.Fatalf("empty filtered query: %v", err)
	}
	if len(empty) != 0 {
		t.Fatalf("rows above chain fixture range = %+v", empty)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := LoadEventRows(ctx, db.Pool, []string{contract}, 0); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled query error = %v, want context canceled", err)
	}

	db.Pool.Close()
	if err := db.SQL.Close(); err != nil {
		t.Fatalf("closing database: %v", err)
	}
	if _, err := LoadEventRows(context.Background(), db.Pool, []string{contract}, 0); err == nil {
		t.Fatal("query on closed PostgreSQL database succeeded")
	}
}
