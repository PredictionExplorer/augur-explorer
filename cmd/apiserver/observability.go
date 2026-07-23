package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "rwcg_http_requests_total",
		Help: "HTTP requests processed, by method, route, status class and " +
			"v1-deprecation membership (docs/operations.md documents the " +
			"sunset-gate queries over the deprecated label).",
	}, []string{"method", "route", "status", "deprecated"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "rwcg_http_request_duration_seconds",
		Help:    "HTTP request latency by route.",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "route"})

	httpRequestTimeoutsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "rwcg_http_request_timeouts_total",
		Help: "Requests whose processing deadline expired and were answered " +
			"with the 503 timeout rendering, by method and route.",
	}, []string{"method", "route"})
)

// countRequestTimeout records one deadline-expired request; it is the
// routes.Options.OnRequestTimeout hook. The route label uses the matched
// pattern like every other request metric.
func countRequestTimeout(r *http.Request) {
	route := httpx.PatternPath(r)
	if route == "" {
		route = "unmatched"
	}
	httpRequestTimeoutsTotal.WithLabelValues(r.Method, route).Inc()
}

// metricsMiddleware records Prometheus counters/latency for every request.
// The route label uses the matched route pattern (not the raw URL) to keep
// metric cardinality bounded. The deprecated label mirrors the RFC 9745
// header policy exactly — both derive from routes.V1Deprecated on the
// request path — so the D6 sunset gate ("30 consecutive zero-traffic days
// on v1") is measurable straight from rwcg_http_requests_total, including
// unmatched 404s under the deprecated prefixes.
func metricsMiddleware() httpx.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := httpx.WrapResponseWriter(w)
			deprecated := strconv.FormatBool(routes.V1Deprecated(r.URL.Path))
			next.ServeHTTP(rw, r)

			route := httpx.PatternPath(r)
			if route == "" {
				route = "unmatched" // 404s and unrouted paths share one label
			}
			status := statusClass(rw.Status())
			httpRequestsTotal.WithLabelValues(r.Method, route, status, deprecated).Inc()
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

// listenInternalServer binds /metrics and /debug/pprof on a private listener.
// These must never be exposed publicly, so they live on their own port,
// controlled by METRICS_ADDR (e.g. "127.0.0.1:9090"). Unset means disabled.
// Binding is synchronous so startup can roll back every open listener before
// any serving goroutine starts. Nil server/listener means disabled.
func listenInternalServer(ctx context.Context, metricsAddr string, errorLog *log.Logger) (*http.Server, net.Listener, error) {
	addr := strings.TrimSpace(metricsAddr)
	if addr == "" {
		return nil, nil, nil
	}
	listener, err := new(net.ListenConfig).Listen(ctx, "tcp", addr)
	if err != nil {
		return nil, nil, err
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
		ReadTimeout:       readTimeout,
		IdleTimeout:       idleTimeout,
		ErrorLog:          errorLog,
	}
	return srv, listener, nil
}
