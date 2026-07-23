//go:build integration

package apitest

// End-to-end coverage for the v1 deprecation announcement through the real
// production chain (routes.New): every deprecated v1 response carries the
// RFC 9745 Deprecation header and the migration-guide Link, the exempt
// surfaces stay clean, and the headers survive negotiated representations
// (gzip, 304 revalidations, HEAD). The parity goldens stay byte-identical
// because they pin status/content-type/body, never headers; the Sunset
// header stays absent because the harness — like production today — sets
// no V1_SUNSET_AT.

import (
	"net/http"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/policy"
)

// wantV1Deprecation asserts the exact deprecation header triple of a
// deprecated v1 response: the pinned Deprecation date, the migration Link
// and no Sunset (not configured).
func wantV1Deprecation(t *testing.T, header http.Header, context string) {
	t.Helper()
	if got, want := header.Get("Deprecation"), "@1784160000"; got != want {
		t.Errorf("%s: Deprecation = %q, want %q", context, got, want)
	}
	wantLink := `<` + policy.V1MigrationGuideURL + `>; rel="deprecation"; type="text/markdown"`
	if got := header.Get("Link"); got != wantLink {
		t.Errorf("%s: Link = %q, want %q", context, got, wantLink)
	}
	if got := header.Get("Sunset"); got != "" {
		t.Errorf("%s: Sunset = %q, want absent (no V1_SUNSET_AT configured)", context, got)
	}
}

// TestV1DeprecationHeaders sweeps representative routes on both sides of
// the policy through the full production chain.
func TestV1DeprecationHeaders(t *testing.T) {
	h := server(t)

	t.Run("deprecated v1 routes announce", func(t *testing.T) {
		for _, route := range []string{
			"/api/cosmicgame/statistics/dashboard",
			"/api/cosmicgame/rounds/list/0/10",
			"/api/randomwalk/statistics/by_market",
		} {
			w := h.get(t, route)
			if w.Code != http.StatusOK {
				t.Fatalf("%s: status = %d", route, w.Code)
			}
			wantV1Deprecation(t, w.Header(), route)
		}
	})

	t.Run("exempt surfaces stay clean", func(t *testing.T) {
		for _, route := range []string{
			"/api/v2/cosmicgame/rounds", // the replacement surface
			"/healthz",
			"/readyz",
			"/version",
			"/cg/metadata/1", // contract-pinned tokenURI target (D12)
		} {
			w := h.get(t, route)
			if got := w.Header().Get("Deprecation"); got != "" {
				t.Errorf("%s: unexpected Deprecation header %q", route, got)
			}
			if got := w.Header().Get("Sunset"); got != "" {
				t.Errorf("%s: unexpected Sunset header %q", route, got)
			}
		}
	})

	t.Run("v1 error responses announce too", func(t *testing.T) {
		w := h.get(t, "/api/cosmicgame/bid/info/not-a-number")
		if w.Code == http.StatusOK {
			t.Fatalf("expected an error status, got 200")
		}
		wantV1Deprecation(t, w.Header(), "malformed-param error response")
	})

	t.Run("HEAD carries the headers", func(t *testing.T) {
		w := h.do(t, request{method: http.MethodHead, path: "/api/cosmicgame/statistics/dashboard"})
		if w.Code != http.StatusOK {
			t.Fatalf("HEAD status = %d", w.Code)
		}
		wantV1Deprecation(t, w.Header(), "HEAD response")
	})

	t.Run("304 revalidation carries the headers", func(t *testing.T) {
		const route = "/api/cosmicgame/statistics/dashboard"
		first := h.get(t, route)
		etag := first.Header().Get("ETag")
		if etag == "" {
			t.Fatal("no ETag on first response")
		}
		w := h.do(t, request{path: route, headers: map[string]string{"If-None-Match": etag}})
		if w.Code != http.StatusNotModified {
			t.Fatalf("revalidation status = %d, want 304", w.Code)
		}
		wantV1Deprecation(t, w.Header(), "304 response")
	})

	t.Run("compressed representation carries the headers", func(t *testing.T) {
		w := h.do(t, request{path: fatV1Route, headers: map[string]string{"Accept-Encoding": "gzip"}})
		if w.Code != http.StatusOK || w.Header().Get("Content-Encoding") != "gzip" {
			t.Fatalf("status=%d encoding=%q, want compressed 200", w.Code, w.Header().Get("Content-Encoding"))
		}
		wantV1Deprecation(t, w.Header(), "gzip response")
	})

	t.Run("FAQ proxy is exempt", func(t *testing.T) {
		w := h.get(t, "/api/cosmicgame/faq/health")
		if got := w.Header().Get("Deprecation"); got != "" {
			t.Errorf("faq health: unexpected Deprecation header %q", got)
		}
	})
}
