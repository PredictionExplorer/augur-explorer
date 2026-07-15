package common

// Benchmarks for the response-edge middleware: gzip compression and the
// conditional-request validator (§4.5 of docs/MODERNIZATION.md). Baselines
// live in docs/benchmarks.md; re-run with:
//
//	go test ./internal/api/common/ -bench 'BenchmarkCompress|BenchmarkConditionalETag' -benchmem -count=6 -run '^$'

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// benchPayload reuses the deterministic JSON generator from the unit tests.
func benchSizes() []int { return []int{1 << 10, 32 << 10, 512 << 10} }

// BenchmarkCompress measures one full middleware exchange per iteration:
// negotiation, threshold buffering, pooled gzip encoding and the write-out.
func BenchmarkCompress(b *testing.B) {
	for _, size := range benchSizes() {
		body := jsonPayload(size)
		h := Compress()(jsonHandler(body))
		b.Run(fmt.Sprintf("gzip_%dKiB", size>>10), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(body)))
			for b.Loop() {
				req := httptest.NewRequest(http.MethodGet, "/x", nil)
				req.Header.Set("Accept-Encoding", "gzip")
				w := httptest.NewRecorder()
				h.ServeHTTP(httpx.WrapResponseWriter(w), req)
			}
		})
		b.Run(fmt.Sprintf("identity_%dKiB", size>>10), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(body)))
			for b.Loop() {
				req := httptest.NewRequest(http.MethodGet, "/x", nil)
				w := httptest.NewRecorder()
				h.ServeHTTP(httpx.WrapResponseWriter(w), req)
			}
		})
	}
}

// BenchmarkConditionalETag measures the buffering + SHA-256 validator cost
// for a full response and the short-circuit revalidation (304) path.
func BenchmarkConditionalETag(b *testing.B) {
	for _, size := range benchSizes() {
		body := jsonPayload(size)
		h := ConditionalETag()(jsonHandler(body))
		etag := weakETag(body)
		b.Run(fmt.Sprintf("tag_%dKiB", size>>10), func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(int64(len(body)))
			for b.Loop() {
				req := httptest.NewRequest(http.MethodGet, "/x", nil)
				w := httptest.NewRecorder()
				h.ServeHTTP(httpx.WrapResponseWriter(w), req)
			}
		})
		b.Run(fmt.Sprintf("revalidate_304_%dKiB", size>>10), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				req := httptest.NewRequest(http.MethodGet, "/x", nil)
				req.Header.Set("If-None-Match", etag)
				w := httptest.NewRecorder()
				h.ServeHTTP(httpx.WrapResponseWriter(w), req)
			}
		})
	}
}
