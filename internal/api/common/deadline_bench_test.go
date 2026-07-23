package common

// Benchmarks for the request-deadline middleware (§4.5 of
// docs/MODERNIZATION.md). Baselines live in docs/benchmarks.md; re-run with:
//
//	go test ./internal/api/common/ -bench BenchmarkRequestDeadline -benchmem -count=6
//
// The middleware sits on every non-exempt request, so the interesting number
// is the pass-through cost: one context.WithTimeout (a timer), one request
// clone and the writer wrapper.

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func deadlineBenchRouter(policy DeadlinePolicy) *httpx.Router {
	r := httpx.NewRouter()
	r.Use(RequestDeadline(policy, nil))
	r.GET("/ping", func(c *httpx.Context) { c.Status(http.StatusNoContent) })
	return r
}

func BenchmarkRequestDeadline(b *testing.B) {
	b.Run("bounded_passthrough", func(b *testing.B) {
		router := deadlineBenchRouter(func(*http.Request) time.Duration { return DefaultRequestDeadline })
		b.ReportAllocs()
		for b.Loop() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/ping", nil))
			if w.Code != http.StatusNoContent {
				b.Fatalf("unexpected status %d", w.Code)
			}
		}
	})

	b.Run("exempt_passthrough", func(b *testing.B) {
		router := deadlineBenchRouter(func(*http.Request) time.Duration { return 0 })
		b.ReportAllocs()
		for b.Loop() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/ping", nil))
			if w.Code != http.StatusNoContent {
				b.Fatalf("unexpected status %d", w.Code)
			}
		}
	})
}
