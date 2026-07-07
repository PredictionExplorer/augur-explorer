//go:build integration

package randomwalk

// Store read suite (§4.2 of docs/MODERNIZATION.md) for the RandomWalk domain:
// every public query function runs against the shared fixture dataset plus a
// suite-local extension seed (rw_uranks, which production fills from the
// rwctl top-rated cron rather than triggers), with results pinned as goldens.
//
// Error paths on the legacy os.Exit methods stay untested until Phase 1
// converts them to returned errors.

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	dbs "github.com/PredictionExplorer/augur-explorer/internal/store"
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
	sharedWrapper *SQLStorageWrapper
	errSetupSkip  error // non-nil: integration environment unavailable, skip
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

	storage := dbs.NewSQLStorageFromDB(db.SQL, log.New(os.Stderr, "store: ", 0))
	storage.Db_set_schema_name("public")
	sharedWrapper = &SQLStorageWrapper{S: storage}

	return m.Run()
}

// store returns the shared wrapper, skipping the test when the integration
// environment (Docker) is unavailable.
func store(t *testing.T) *SQLStorageWrapper {
	t.Helper()
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	if sharedWrapper == nil {
		t.Fatal("store harness not initialized (TestMain did not run?)")
	}
	return sharedWrapper
}

// golden pins fetch() as testdata/golden/<name>.json, fetching twice to
// prove the query is deterministic.
func golden(t *testing.T, name string, fetch func() any) {
	t.Helper()
	testutil.GoldenJSON(t, filepath.Join("testdata", "golden", name+".json"), fetch)
}
