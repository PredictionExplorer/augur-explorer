//go:build integration

package cosmicgame

import (
	"context"
	"log/slog"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestContractDriftAuditDependencyAndContextFailures(t *testing.T) {
	resetDB(t)
	logger := slog.New(slog.DiscardHandler)
	if _, err := CheckContractParamsDrift(
		context.Background(), nil, eclient,
		fxGameAddr, fxPrizesAddr, logger,
	); err == nil || !strings.Contains(err.Error(), "repo is nil") {
		t.Fatalf("nil repo error = %v", err)
	}

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := CheckContractParamsDrift(
		cancelled, cgRepo, eclient,
		fxGameAddr, fxPrizesAddr, logger,
	); err == nil || !strings.Contains(err.Error(), "latest header") {
		t.Fatalf("cancelled audit error = %v", err)
	}

	cfg, err := pgxpool.ParseConfig(testDB.ConnString)
	if err != nil {
		t.Fatal(err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	closedStore := store.NewFromPool(pool)
	closedRepo := cgstore.NewRepo(closedStore)
	closedStore.Close()
	testChain.EnsureBlock(5700)
	testChain.RegisterCall(addr(fxGameAddr), driftV1GameStub().Handler())
	t.Cleanup(registerCallHandlers)
	if _, err := CheckContractParamsDrift(
		context.Background(), closedRepo, eclient,
		fxGameAddr, "", logger,
	); err == nil {
		t.Fatal("closed repository drift audit succeeded")
	}

	restore := faultHarnessTable(t, "cg_glob_stats")
	_, err = CheckContractParamsDrift(
		context.Background(), cgRepo, eclient,
		fxGameAddr, "", logger,
	)
	restore()
	if err == nil || !strings.Contains(err.Error(), "cg_glob_stats") {
		t.Fatalf("missing glob stats error = %v", err)
	}
}
