package testutil

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCompareGoldenJSONRoundTrip(t *testing.T) {
	type rec struct {
		Name  string
		Count int64
		Ratio float64
	}
	v := []rec{{Name: "alice", Count: 3, Ratio: 0.5}, {Name: "bob", Count: 1, Ratio: 1.25}}

	path := filepath.Join(t.TempDir(), "rec.json")
	want := "[\n  {\n    \"Name\": \"alice\",\n    \"Count\": 3,\n    \"Ratio\": 0.5\n  },\n  {\n    \"Name\": \"bob\",\n    \"Count\": 1,\n    \"Ratio\": 1.25\n  }\n]\n"
	if err := os.WriteFile(path, []byte(want), 0o600); err != nil {
		t.Fatal(err)
	}

	CompareGoldenJSON(t, path, v) // must match the pre-written golden exactly
}

func TestGoldenJSONDetectsNondeterminism(t *testing.T) {
	path := filepath.Join(t.TempDir(), "flaky.json")
	n := 0
	mock := &testing.T{}
	done := make(chan struct{})
	go func() {
		defer close(done)
		// Runs in a goroutine because Fatalf calls runtime.Goexit.
		GoldenJSON(mock, path, func() any {
			n++
			return n
		})
	}()
	<-done
	if !mock.Failed() {
		t.Fatal("GoldenJSON accepted a fetch that returned different values on each call")
	}
}
