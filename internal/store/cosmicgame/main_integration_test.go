//go:build integration

package cosmicgame

// Store read suite (§4.2 of docs/MODERNIZATION.md): every public query
// function is called against the shared fixture dataset at least once and
// its result pinned as a golden file. This is the safety net that let the
// Phase 1 store rewrite proceed file-by-file; the whole CosmicGame read
// layer now runs on the Repo (context-first, error-returning), so the suite
// no longer carries a legacy-wrapper handle. The remaining legacy files
// (inserts.go / deletes.go) are exercised by the cmd/cg-etl fixture replay.

import (
	"context"
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

// Fixture handles referenced by the read-suite cases
// (see internal/testfixtures/seed/*.sql).
const (
	aidZero            = 1 // 0x0000...0000
	aidCosmicGame      = 2
	aidCosmicSignature = 3
	aidCosmicToken     = 4
	aidCharityWallet   = 6
	aidPrizesWallet    = 7
	aidRandomWalk      = 8
	aidStakingCST      = 9
	aidStakingRWalk    = 10
	aidMarketingWallet = 11
	aidMarketplace     = 12
	aidAlice           = 21
	aidBob             = 22
	aidCarol           = 23
	aidDave            = 24
	aidEmma            = 25
	aidDonatedERC20    = 26
	aidDonatedNFT      = 27
)

var (
	sharedRepo       *Repo
	sharedConnString string
	errSetupSkip     error // non-nil: integration environment unavailable, skip
)

// TestMain owns one database container seeded with the shared fixture set;
// every test in the package reads from it through the same wrapper.
func TestMain(m *testing.M) {
	os.Exit(runMain(m))
}

func runMain(m *testing.M) int {
	flag.Parse()

	// Golden files must not depend on the machine's timezone: timestamptz
	// columns scanned into time.Time render in time.Local.
	time.Local = time.UTC

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	db, stop, err := testdb.Start(ctx)
	if err != nil {
		if errors.Is(err, testdb.ErrContainerUnavailable) {
			errSetupSkip = err
			return m.Run() // integration tests skip with the reason
		}
		fmt.Fprintf(os.Stderr, "store/cosmicgame: starting test database: %v\n", err)
		return 1
	}
	defer stop()

	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		fmt.Fprintf(os.Stderr, "store/cosmicgame: seeding fixture data: %v\n", err)
		return 1
	}

	// One Store over the container's pool, mirroring the production wiring:
	// every query in this package runs on the Repo.
	st := store.NewFromPool(db.Pool)
	sharedRepo = NewRepo(st)
	sharedConnString = db.ConnString

	return m.Run()
}

// spareStore opens a second, independent Store on the test database for
// cases that need to close or break a pool without affecting the shared one.
func spareStore(ctx context.Context) (*store.Store, error) {
	pool, err := pgxpool.New(ctx, sharedConnString)
	if err != nil {
		return nil, err
	}
	return store.NewFromPool(pool), nil
}

// repo returns the shared Repo (converted query methods), skipping the test
// when the integration environment (Docker) is unavailable.
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

// golden pins fetch() as testdata/golden/<name>.json, fetching twice to
// prove the query is deterministic.
func golden(t *testing.T, name string, fetch func() any) {
	t.Helper()
	testutil.GoldenJSON(t, filepath.Join("testdata", "golden", name+".json"), fetch)
}
