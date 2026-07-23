package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/policy"
)

func TestV1Deprecated(t *testing.T) {
	cases := []struct {
		path string
		want bool
	}{
		// The frozen v1 surface.
		{"/api/cosmicgame/statistics/dashboard", true},
		{"/api/cosmicgame/bid/list/0/100", true},
		{"/api/randomwalk/statistics", true},
		{"/api/randomwalk/token-ranking/match", true},

		// The FAQ proxy fronts a separate service; no v2 replacement.
		{"/api/cosmicgame/faq/query", false},
		{"/api/cosmicgame/faq/health", false},

		// The replacement surface itself.
		{"/api/v2/cosmicgame/rounds", false},
		{"/api/v2/randomwalk/tokens", false},

		// Contract-pinned tokenURI targets (D12: can never move).
		{"/metadata/42", false},
		{"/cg/metadata/42", false},

		// Operational endpoints.
		{"/healthz", false},
		{"/readyz", false},
		{"/version", false},

		// Prefix edges: only true path segments below the module roots.
		{"/api/cosmicgame", false},
		{"/api/cosmicgames/thing", false},
		{"/api/randomwalk", false},
		{"/", false},
	}
	for _, tc := range cases {
		if got := policy.V1Deprecated(tc.path); got != tc.want {
			t.Errorf("policy.V1Deprecated(%q) = %v, want %v", tc.path, got, tc.want)
		}
	}
}

// TestNewWiresDeprecationHeaders proves the assembled router announces the
// v1 deprecation on every response under a deprecated prefix (the module
// route tables are empty here, so the 404s exercise the global chain) and
// keeps operational and v2 paths clean.
func TestNewWiresDeprecationHeaders(t *testing.T) {
	r := New(nil, Options{})

	get := func(path string) *httptest.ResponseRecorder {
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.RemoteAddr = "10.1.2.3:999"
		r.ServeHTTP(rec, req)
		return rec
	}

	deprecated := get("/api/cosmicgame/statistics/dashboard")
	if got := deprecated.Header().Get("Deprecation"); got != "@1784160000" {
		t.Errorf("Deprecation = %q, want @1784160000 (2026-07-16T00:00:00Z)", got)
	}
	wantLink := `<` + policy.V1MigrationGuideURL + `>; rel="deprecation"; type="text/markdown"`
	if got := deprecated.Header().Get("Link"); got != wantLink {
		t.Errorf("Link = %q, want %q", got, wantLink)
	}
	if got := deprecated.Header().Get("Sunset"); got != "" {
		t.Errorf("Sunset emitted without Options.V1SunsetAt: %q", got)
	}

	for _, clean := range []string{"/healthz", "/version", "/api/cosmicgame/faq/health"} {
		rec := get(clean)
		if got := rec.Header().Get("Deprecation"); got != "" {
			t.Errorf("%s: unexpected Deprecation header %q", clean, got)
		}
	}
}

func TestNewWiresSunsetWhenConfigured(t *testing.T) {
	r := New(nil, Options{
		V1SunsetAt: time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC),
	})
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/randomwalk/statistics", nil)
	req.RemoteAddr = "10.1.2.3:999"
	r.ServeHTTP(rec, req)
	if got, want := rec.Header().Get("Sunset"), "Fri, 01 Jan 2027 00:00:00 GMT"; got != want {
		t.Errorf("Sunset = %q, want %q", got, want)
	}
	if got := rec.Header().Get("Deprecation"); got == "" {
		t.Error("Deprecation missing alongside Sunset")
	}
}
