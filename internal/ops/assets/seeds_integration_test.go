//go:build integration

package assets

import (
	"context"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

func TestSQLTokenSourceLoadsSeededDatabase(t *testing.T) {
	db := testdb.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("apply fixtures: %v", err)
	}

	tokens, err := (SQLTokenSource{DB: db.Pool}).TokenSeeds(ctx, "public")
	if err != nil {
		t.Fatalf("TokenSeeds(): %v", err)
	}
	if len(tokens) < 9 {
		t.Fatalf("token count = %d, want at least 9", len(tokens))
	}
	if tokens[0].TokenID != 1 ||
		tokens[0].Seed != "seed0000000000000000000000000000000000000000000000000000000001" {
		t.Fatalf("first token = %#v", tokens[0])
	}
	for i := 1; i < len(tokens); i++ {
		if tokens[i-1].TokenID >= tokens[i].TokenID {
			t.Fatalf("tokens are not ordered: %#v then %#v", tokens[i-1], tokens[i])
		}
	}
}
