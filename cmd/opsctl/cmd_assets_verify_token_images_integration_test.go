//go:build integration

package main

import (
	"context"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

func TestRandomWalkTokenCountSourceAgainstDB(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	st := store.NewFromPool(db.Pool)

	deps := defaultAssetsVerifyTokenImagesDeps()
	src := deps.newSource(st)

	// The RandomWalk contract address is not in the fixture database yet:
	// the lookup fails with a descriptive error.
	if _, err := src.MintedTokenCount(ctx); err == nil ||
		!strings.Contains(err.Error(), "RandomWalk contract address") {
		t.Fatalf("expected address-resolution error, got %v", err)
	}

	// With the address registered, the count comes from RandomWalkStats.
	if _, err := db.Pool.Exec(ctx,
		`INSERT INTO address(block_num, tx_id, addr) VALUES (0, 0, $1)`,
		rwalkNFTAddr); err != nil {
		t.Fatal(err)
	}
	count, err := src.MintedTokenCount(ctx)
	if err != nil {
		t.Fatalf("MintedTokenCount: %v", err)
	}
	if count != 0 {
		t.Fatalf("empty database: got %d minted tokens", count)
	}

	deps.closeStore(st)
}
