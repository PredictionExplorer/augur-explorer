package common

// Benchmarks for the rate-limit middleware (§4.5 of docs/MODERNIZATION.md).
// Baselines live in docs/benchmarks.md; re-run with:
//
//	go test ./internal/api/common/ -bench BenchmarkRateLimiter -benchmem -count=6

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func rateLimitedRouter() *httpx.Router {
	r := httpx.NewRouter()
	// Effectively unlimited so the benchmark measures limiter bookkeeping,
	// not 429 short-circuits.
	r.Use(RateLimit(1e9, 1<<30))
	r.GET("/ping", func(c *httpx.Context) { c.Status(http.StatusNoContent) })
	return r
}

// BenchmarkRateLimiter/distinct_ips exercises the per-IP map path (every
// goroutine is its own client); /shared_ip hammers a single token bucket to
// expose mutex contention on the hottest key.
func BenchmarkRateLimiter(b *testing.B) {
	b.Run("distinct_ips", func(b *testing.B) {
		router := rateLimitedRouter()
		var ipCounter atomic.Uint64
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			n := ipCounter.Add(1)
			remote := fmt.Sprintf("10.%d.%d.%d:4242", (n>>16)&0xff, (n>>8)&0xff, n&0xff)
			for pb.Next() {
				req := httptest.NewRequest(http.MethodGet, "/ping", nil)
				req.RemoteAddr = remote
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				if w.Code != http.StatusNoContent {
					b.Fatalf("unexpected status %d", w.Code)
				}
			}
		})
	})

	b.Run("shared_ip", func(b *testing.B) {
		router := rateLimitedRouter()
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				req := httptest.NewRequest(http.MethodGet, "/ping", nil)
				req.RemoteAddr = "10.0.0.1:4242"
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
				if w.Code != http.StatusNoContent {
					b.Fatalf("unexpected status %d", w.Code)
				}
			}
		})
	})
}
