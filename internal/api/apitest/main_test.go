//go:build integration

package apitest

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

var (
	sharedHarness *harness
	errSetupSkip  error // non-nil: integration environment unavailable, skip
)

// TestMain owns the one database container and the one process-wide harness:
// the API packages keep their dependencies in package state, so Init can run
// only once, and the container must outlive every test in the package.
func TestMain(m *testing.M) {
	os.Exit(runMain(m))
}

func runMain(m *testing.M) int {
	flag.Parse()

	// Golden files must not depend on the machine's timezone: handlers that
	// scan timestamptz columns into time.Time render them in time.Local.
	time.Local = time.UTC

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	db, stop, err := testdb.Start(ctx)
	if err != nil {
		if errors.Is(err, testdb.ErrContainerUnavailable) {
			errSetupSkip = err
			return m.Run() // every test skips with the reason
		}
		fmt.Fprintf(os.Stderr, "apitest: starting test database: %v\n", err)
		return 1
	}
	defer stop()

	if err := seedDatabase(ctx, db.SQL); err != nil {
		fmt.Fprintf(os.Stderr, "apitest: seeding fixture data: %v\n", err)
		return 1
	}

	h, err := newHarness(ctx, db.SQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "apitest: building harness: %v\n", err)
		return 1
	}
	sharedHarness = h

	return m.Run()
}

// server returns the shared harness, skipping the test when the integration
// environment (Docker) is unavailable.
func server(t *testing.T) *harness {
	t.Helper()
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	if sharedHarness == nil {
		t.Fatal("apitest harness not initialized (TestMain did not run?)")
	}
	return sharedHarness
}
