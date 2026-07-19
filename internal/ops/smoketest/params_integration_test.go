//go:build integration

package smoketest

import (
	"context"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

func TestSQLParameterSourceLoadsSeededDatabase(t *testing.T) {
	db := testdb.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("apply fixtures: %v", err)
	}

	params, err := (SQLParameterSource{DB: db.Pool}).Parameters(ctx)
	if err != nil {
		t.Fatalf("Parameters(): %v", err)
	}
	if params.TokenName != "Genesis" {
		t.Fatalf("token name = %q, want fixture value Genesis", params.TokenName)
	}
	if params.RoundNumber != "2" {
		t.Fatalf("round = %q, want latest claimed round 2", params.RoundNumber)
	}
	if params.TokenID != "9" {
		t.Fatalf("token id = %q, want latest fixture token 9", params.TokenID)
	}
	if params.UserAddress == zeroAddress || params.CSTStakerAddress == zeroAddress {
		t.Fatalf("fixture addresses were not loaded: %#v", params)
	}
}
