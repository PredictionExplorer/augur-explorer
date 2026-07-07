package testutil

import (
	"bytes"
	"encoding/json"
	"testing"
)

// CompareGoldenJSON pins an arbitrary Go value as indented JSON against the
// golden file at path (same -update semantics as CompareGolden). Struct field
// order and sorted map keys make the encoding stable.
func CompareGoldenJSON(t *testing.T, path string, v any) {
	t.Helper()
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("marshaling value for golden file %s: %v", path, err)
	}
	CompareGolden(t, path, data)
}

// GoldenJSON evaluates fetch twice, fails the test if the two results encode
// differently (nondeterministic query), and pins the result against the
// golden file at path. Integration suites use it so every pinned query also
// proves its own determinism.
func GoldenJSON(t *testing.T, path string, fetch func() any) {
	t.Helper()
	first, err := json.MarshalIndent(fetch(), "", "  ")
	if err != nil {
		t.Fatalf("marshaling value for golden file %s: %v", path, err)
	}
	second, err := json.MarshalIndent(fetch(), "", "  ")
	if err != nil {
		t.Fatalf("marshaling second fetch for golden file %s: %v", path, err)
	}
	if !bytes.Equal(first, second) {
		t.Fatalf("nondeterministic result for golden file %s:\n%s", path, DiffSummary(first, second))
	}
	CompareGolden(t, path, first)
}
