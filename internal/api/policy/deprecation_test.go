package policy

import "testing"

func TestV1DeprecatedBoundaries(t *testing.T) {
	t.Parallel()
	tests := map[string]bool{
		"/api/cosmicgame/statistics/dashboard": true,
		"/api/randomwalk/statistics":           true,
		"/api/cosmicgame/faq/query":            false,
		"/api/v2/cosmicgame/rounds":            false,
		"/healthz":                             false,
		"/readyz":                              false,
		"/version":                             false,
		"/metadata/42":                         false,
		"/api/cosmicgame":                      false,
		"/api/randomwalk":                      false,
	}
	for path, want := range tests {
		if got := V1Deprecated(path); got != want {
			t.Errorf("V1Deprecated(%q) = %v, want %v", path, got, want)
		}
	}
}

func FuzzV1Deprecated(f *testing.F) {
	f.Add("/api/cosmicgame/statistics/dashboard")
	f.Add("/api/v2/randomwalk/tokens")
	f.Fuzz(func(t *testing.T, path string) {
		first := V1Deprecated(path)
		if second := V1Deprecated(path); second != first {
			t.Fatalf("non-deterministic result for %q: %v then %v", path, first, second)
		}
	})
}
