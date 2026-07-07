//go:build integration

package apitest

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// update regenerates golden files instead of comparing against them:
//
//	go test -tags=integration ./internal/api/apitest/ -run TestAPIParity -update
var update = flag.Bool("update", false, "rewrite golden files with current responses")

// response captures everything a golden file pins about one HTTP exchange.
type response struct {
	Status      int    `json:"status"`
	ContentType string `json:"contentType"`
	Body        any    `json:"body"`
}

// redactor mutates a decoded JSON body to remove legitimately volatile fields
// before comparison (e.g. randomly generated nonces). It returns the value to
// store; most redactors mutate in place and return their argument.
type redactor func(t *testing.T, body any) any

// canonicalJSON re-encodes a JSON document with sorted keys and stable
// indentation. Numbers pass through as json.Number so 256-bit wei amounts are
// not damaged by float64 round-tripping.
func canonicalJSON(t *testing.T, data []byte) any {
	t.Helper()
	if len(bytes.TrimSpace(data)) == 0 {
		return nil
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		// Not JSON (e.g. /healthz plain text): pin the raw string.
		return string(data)
	}
	return v
}

// goldenPath maps a snapshot name to its file under testdata/golden.
func goldenPath(name string) string {
	return filepath.Join("testdata", "golden", name+".json")
}

// goldenName converts a route path (with concrete parameters substituted)
// into a filesystem-safe golden file name.
func goldenName(method, path string, suffix string) string {
	name := strings.Trim(path, "/")
	name = strings.NewReplacer("/", "_", ":", "", "*", "", "?", "_", "&", "_", "=", "_").Replace(name)
	if name == "" {
		name = "root"
	}
	if suffix != "" {
		name += "__" + suffix
	}
	if method != "GET" {
		name = strings.ToLower(method) + "__" + name
	}
	return name
}

// compareGolden asserts that got matches the golden file, or rewrites the
// file when -update is set.
func compareGolden(t *testing.T, name string, got response) {
	t.Helper()

	encoded, err := json.MarshalIndent(got, "", "  ")
	if err != nil {
		t.Fatalf("%s: marshaling response for golden comparison: %v", name, err)
	}
	encoded = append(encoded, '\n')

	path := goldenPath(name)
	if *update {
		if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
			t.Fatalf("%s: creating golden dir: %v", name, err)
		}
		if err := os.WriteFile(path, encoded, 0o600); err != nil {
			t.Fatalf("%s: writing golden file: %v", name, err)
		}
		return
	}

	want, err := os.ReadFile(path) //nolint:gosec // path derives from sanitized route names under testdata/
	if err != nil {
		t.Fatalf("%s: missing golden file %s (run with -update to create): %v", name, path, err)
	}
	if !bytes.Equal(want, encoded) {
		t.Errorf("%s: response differs from golden file %s\n%s",
			name, path, diffSummary(want, encoded))
	}
}

// diffSummary renders a compact first-difference report so a failing parity
// test points at the drift without dumping two full JSON documents.
func diffSummary(want, got []byte) string {
	wantLines := strings.Split(string(want), "\n")
	gotLines := strings.Split(string(got), "\n")
	limit := min(len(wantLines), len(gotLines))
	for i := 0; i < limit; i++ {
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
