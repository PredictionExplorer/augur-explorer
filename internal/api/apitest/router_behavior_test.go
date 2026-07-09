//go:build integration

package apitest

import (
	"net/http"
	"strings"
	"testing"
)

// TestRouterBehavior pins router-level semantics of the stdlib mux port that
// no golden covers: unmatched paths, method mismatches, preflight, HEAD and
// the trailing-slash redirect. Three behaviors deliberately differ from the
// legacy gin router (documented in docs/MODERNIZATION.md §6.2):
//
//   - a wrong method on a known path answers 405 with an Allow header
//     (legacy: plain 404),
//   - the router-level 404 body carries stdlib's trailing newline,
//   - HEAD requests are served by GET routes (legacy: 404).
func TestRouterBehavior(t *testing.T) {
	h := server(t)

	t.Run("unknown_path_404", func(t *testing.T) {
		w := h.get(t, "/api/does/not/exist")
		if w.Code != http.StatusNotFound {
			t.Fatalf("status = %d, want 404", w.Code)
		}
		if body := w.Body.String(); body != "404 page not found\n" {
			t.Errorf("404 body = %q", body)
		}
		// Global middleware wraps unmatched requests: CORS headers present.
		if w.Header().Get("Access-Control-Allow-Origin") != "*" {
			t.Error("404 response must carry the CORS headers")
		}
	})

	t.Run("wrong_method_405_with_allow", func(t *testing.T) {
		w := h.do(t, request{method: http.MethodPost, path: "/api/cosmicgame/statistics/dashboard"})
		if w.Code != http.StatusMethodNotAllowed {
			t.Fatalf("status = %d, want 405", w.Code)
		}
		if allow := w.Header().Get("Allow"); !strings.Contains(allow, http.MethodGet) {
			t.Errorf("Allow = %q, want it to list GET", allow)
		}
	})

	t.Run("head_served_by_get_route", func(t *testing.T) {
		w := h.do(t, request{method: http.MethodHead, path: "/healthz"})
		if w.Code != http.StatusOK {
			t.Fatalf("HEAD /healthz = %d, want 200", w.Code)
		}
	})

	t.Run("options_preflight_204_any_path", func(t *testing.T) {
		for _, path := range []string{"/api/cosmicgame/statistics/dashboard", "/api/never/registered"} {
			w := h.do(t, request{method: http.MethodOptions, path: path})
			if w.Code != http.StatusNoContent {
				t.Errorf("OPTIONS %s = %d, want 204", path, w.Code)
			}
			if w.Body.Len() != 0 {
				t.Errorf("OPTIONS %s body = %q, want empty", path, w.Body.String())
			}
			if w.Header().Get("Access-Control-Allow-Methods") == "" {
				t.Errorf("OPTIONS %s missing preflight headers", path)
			}
		}
	})

	t.Run("trailing_slash_redirects_to_route", func(t *testing.T) {
		w := h.get(t, "/api/cosmicgame/statistics/dashboard/")
		if w.Code != http.StatusMovedPermanently {
			t.Fatalf("status = %d, want 301", w.Code)
		}
		if loc := w.Header().Get("Location"); loc != "/api/cosmicgame/statistics/dashboard" {
			t.Errorf("Location = %q", loc)
		}
	})

	t.Run("rate_limit_envelope_429", func(t *testing.T) {
		// One fixed IP hammers a cheap route past the global burst (100).
		const ip = "10.255.0.1:4242"
		var last int
		for range 120 {
			w := h.do(t, request{path: "/healthz", remoteAddr: ip})
			last = w.Code
			if last == http.StatusTooManyRequests {
				break
			}
		}
		if last != http.StatusTooManyRequests {
			t.Fatalf("burst was never limited, last status %d", last)
		}
	})
}
