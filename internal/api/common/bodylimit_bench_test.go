package common

// Benchmarks for the request-body cap middleware (§4.5 of
// docs/MODERNIZATION.md). Baselines live in docs/benchmarks.md; re-run with:
//
//	go test ./internal/api/common/ -bench BenchmarkMaxRequestBody -benchmem -count=6
//
// The middleware sits on every request, so the interesting numbers are the
// pass-through costs (bodyless GET, small POST), not the 413 rejection.

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func bodyCappedRouter() *httpx.Router {
	r := httpx.NewRouter()
	r.Use(MaxRequestBody(MaxRequestBodyBytes))
	r.GET("/ping", func(c *httpx.Context) { c.Status(http.StatusNoContent) })
	r.POST("/echo", func(c *httpx.Context) {
		_, _ = io.Copy(io.Discard, c.Request.Body)
		c.Status(http.StatusNoContent)
	})
	return r
}

func BenchmarkMaxRequestBody(b *testing.B) {
	b.Run("get_passthrough", func(b *testing.B) {
		router := bodyCappedRouter()
		b.ReportAllocs()
		for b.Loop() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/ping", nil))
			if w.Code != http.StatusNoContent {
				b.Fatalf("unexpected status %d", w.Code)
			}
		}
	})

	b.Run("post_1KiB_within_cap", func(b *testing.B) {
		router := bodyCappedRouter()
		payload := strings.Repeat("x", 1<<10)
		b.ReportAllocs()
		for b.Loop() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/echo", strings.NewReader(payload)))
			if w.Code != http.StatusNoContent {
				b.Fatalf("unexpected status %d", w.Code)
			}
		}
	})

	b.Run("declared_oversize_413", func(b *testing.B) {
		router := bodyCappedRouter()
		payload := strings.Repeat("x", int(MaxRequestBodyBytes)+1)
		b.ReportAllocs()
		for b.Loop() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/echo", strings.NewReader(payload)))
			if w.Code != http.StatusRequestEntityTooLarge {
				b.Fatalf("unexpected status %d", w.Code)
			}
		}
	})
}
