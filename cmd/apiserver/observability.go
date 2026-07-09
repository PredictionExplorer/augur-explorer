package main

import (
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "rwcg_http_requests_total",
		Help: "HTTP requests processed, by method, route and status class.",
	}, []string{"method", "route", "status"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "rwcg_http_request_duration_seconds",
		Help:    "HTTP request latency by route.",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "route"})
)

// metricsMiddleware records Prometheus counters/latency for every request.
// The route label uses the matched route pattern (not the raw URL) to keep
// metric cardinality bounded.
func metricsMiddleware() httpx.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := httpx.WrapResponseWriter(w)
			next.ServeHTTP(rw, r)

			route := httpx.PatternPath(r)
			if route == "" {
				route = "unmatched" // 404s and unrouted paths share one label
			}
			status := statusClass(rw.Status())
			httpRequestsTotal.WithLabelValues(r.Method, route, status).Inc()
			httpRequestDuration.WithLabelValues(r.Method, route).Observe(time.Since(start).Seconds())
		})
	}
}

func statusClass(code int) string {
	switch {
	case code >= 500:
		return "5xx"
	case code >= 400:
		return "4xx"
	case code >= 300:
		return "3xx"
	default:
		return "2xx"
	}
}

// startInternalServer serves /metrics and /debug/pprof on a private listener.
// These must never be exposed publicly, so they live on their own port,
// controlled by METRICS_ADDR (e.g. "127.0.0.1:9090"). Unset means disabled.
// The returned server (nil when disabled) participates in graceful shutdown.
func startInternalServer() *http.Server {
	addr := strings.TrimSpace(os.Getenv("METRICS_ADDR"))
	if addr == "" {
		return nil
	}
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	go func() {
		Info.Printf("internal metrics/pprof server listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Error.Printf("internal metrics server: %v", err)
		}
	}()
	return srv
}
