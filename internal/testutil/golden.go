// Package testutil provides shared helpers for integration test suites:
// golden-file comparison and canonical database snapshots/diffs.
//
// The package deliberately carries no build tag so its own unit tests run in
// the default suite; consumers under the `integration` tag import it as usual.
package testutil

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// updateGolden regenerates golden files instead of comparing against them:
//
//	go test -tags=integration ./cmd/cg-etl/ -run TestEventFixtures -update
var updateGolden = flag.Bool("update", false, "rewrite golden files with current values")

// UpdatingGolden reports whether the test run was started with -update.
func UpdatingGolden() bool { return *updateGolden }

// CompareGolden asserts that got matches the golden file at path, or rewrites
// the file when the -update flag is set. got should already be in a stable,
// canonical encoding; a trailing newline is added if missing.
func CompareGolden(t *testing.T, path string, got []byte) {
	t.Helper()

	if len(got) == 0 || got[len(got)-1] != '\n' {
		got = append(got, '\n')
	}

	if *updateGolden {
		if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
			t.Fatalf("creating golden dir for %s: %v", path, err)
		}
		if err := os.WriteFile(path, got, 0o600); err != nil {
			t.Fatalf("writing golden file %s: %v", path, err)
		}
		return
	}

	want, err := os.ReadFile(path) //nolint:gosec // test-authored fixture paths under testdata/
	if err != nil {
		t.Fatalf("missing golden file %s (run with -update to create): %v", path, err)
	}
	if !bytes.Equal(want, got) {
		t.Errorf("value differs from golden file %s\n%s", path, DiffSummary(want, got))
	}
}

// DiffSummary renders a compact first-difference report so a failing golden
// comparison points at the drift without dumping two full documents.
func DiffSummary(want, got []byte) string {
	wantLines := strings.Split(string(want), "\n")
	gotLines := strings.Split(string(got), "\n")
	limit := min(len(wantLines), len(gotLines))
	for i := range limit {
		if wantLines[i] != gotLines[i] {
			lo := max(0, i-2)
			hiW := min(len(wantLines), i+3)
			hiG := min(len(gotLines), i+3)
			return fmt.Sprintf("first difference at line %d:\n--- golden ---\n%s\n--- got ---\n%s",
				i+1,
				strings.Join(wantLines[lo:hiW], "\n"),
				strings.Join(gotLines[lo:hiG], "\n"))
		}
	}
	return fmt.Sprintf("documents diverge in length: golden %d lines, got %d lines",
		len(wantLines), len(gotLines))
}
