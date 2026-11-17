//go:build integration

package etlapp_test

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/etlapp"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

// TestRunReportsEngineLoopFailure hands Run a module the engine loop's own
// validation rejects (no watermark adapter) and asserts the failure is
// surfaced as the fatal loop error, not swallowed: an ETL that cannot track
// its position must crash, silently idling would look healthy while
// indexing nothing.
func TestRunReportsEngineLoopFailure(t *testing.T) {
	db := testdb.New(t)
	chain := testchain.New(t)
	chain.EnsureBlock(1)

	env := map[string]string{
		"RPC_URL":      chain.URL(),
		"DATABASE_URL": db.ConnString,
	}
	getenv := func(key string) string { return env[key] }

	registry, err := indexer.NewRegistry()
	if err != nil {
		t.Fatalf("empty registry: %v", err)
	}
	wire := func(context.Context, etlapp.Deps) (etlapp.Module, error) {
		return etlapp.Module{Registry: registry}, nil
	}

	err = etlapp.Run(context.Background(), getenv, io.Discard,
		prometheus.NewRegistry(), prometheus.NewRegistry(), wire)
	if err == nil || !strings.Contains(err.Error(), "event processing loop terminated") {
		t.Fatalf("err = %v, want the engine loop failure", err)
	}
}
