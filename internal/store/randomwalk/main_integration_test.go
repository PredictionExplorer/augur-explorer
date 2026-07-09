//go:build integration

package randomwalk

// Store read suite (§4.2 of docs/MODERNIZATION.md) for the RandomWalk domain:
// every public query function runs against the shared fixture dataset plus a
// suite-local extension seed (rw_uranks, which production fills from the
// rwctl top-rated cron rather than triggers), with results pinned as goldens.
// The suite runs on the production wiring: one pgx pool wrapped by
// store.NewFromPool and queried through the Repo.

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// Fixture handles (see internal/testfixtures/seed/*.sql).
const (
	aidRandomWalk  = 8  // RandomWalk NFT contract
	aidMarketplace = 12 // RW marketplace contract
	aidAlice       = 21
	aidBob         = 22
	aidCarol       = 23
	aidDave        = 24

	addrRandomWalk  = "0x8000000000000000000000000000000000000008"
	addrMarketplace = "0x1200000000000000000000000000000000000012"
)

//go:embed testdata/extension/*.sql
var extensionSeedFS embed.FS

var (
	sharedStore      *store.Store
	sharedRepo       *Repo
	sharedConnString string
	errSetupSkip     error // non-nil: integration environment unavailable, skip
)

func TestMain(m *testing.M) {
	os.Exit(runMain(m))
}

func runMain(m *testing.M) int {
	flag.Parse()

	// Golden files must not depend on the machine's timezone.
	time.Local = time.UTC

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	db, stop, err := testdb.Start(ctx)
	if err != nil {
		if errors.Is(err, testdb.ErrContainerUnavailable) {
			errSetupSkip = err
			return m.Run() // integration tests skip with the reason
		}
		fmt.Fprintf(os.Stderr, "store/randomwalk: starting test database: %v\n", err)
		return 1
	}
	defer stop()

	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		fmt.Fprintf(os.Stderr, "store/randomwalk: seeding fixture data: %v\n", err)
		return 1
	}
	if err := testfixtures.ApplyFS(ctx, db.SQL, extensionSeedFS, "testdata/extension"); err != nil {
		fmt.Fprintf(os.Stderr, "store/randomwalk: applying extension seed: %v\n", err)
		return 1
	}

	sharedStore = store.NewFromPool(db.Pool)
	sharedRepo = NewRepo(sharedStore)
	sharedConnString = db.ConnString

	return m.Run()
}

// spareStore opens a second Store on the same container so tests can
// exercise closed-pool behavior without harming the shared pool.
func spareStore(ctx context.Context) (*store.Store, error) {
	pool, err := pgxpool.New(ctx, sharedConnString)
	if err != nil {
		return nil, err
	}
	return store.NewFromPool(pool), nil
}

// repo returns the shared Repo, skipping the test when the integration
// environment (Docker) is unavailable.
func repo(t *testing.T) *Repo {
	t.Helper()
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	if sharedRepo == nil {
		t.Fatal("store harness not initialized (TestMain did not run?)")
	}
	return sharedRepo
}

// pool exposes the shared pgx pool for the fixture-restoring cleanup
// statements of the write tests.
func pool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	repo(t)
	return sharedStore.Pool()
}

// golden pins fetch() as testdata/golden/<name>.json, fetching twice to
// prove the query is deterministic.
func golden(t *testing.T, name string, fetch func() any) {
	t.Helper()
	testutil.GoldenJSON(t, filepath.Join("testdata", "golden", name+".json"), fetch)
}
