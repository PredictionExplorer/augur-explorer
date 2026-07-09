// Private metrics/pprof listener for the ETL binaries, mirroring the API
// server's internal server: /metrics plus /debug/pprof on a port that must
// never be exposed publicly.

package indexer

import (
	"log/slog"
	"net"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartMetricsServer serves Prometheus metrics (from gatherer) and pprof on
// addr (e.g. "127.0.0.1:9091"). It returns the running server — callers Close
// it on shutdown — and the bound address (useful with a ":0" test listener).
func StartMetricsServer(addr string, gatherer prometheus.Gatherer, logger *slog.Logger) (*http.Server, net.Addr, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(gatherer, promhttp.HandlerOpts{}))
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	srv := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	go func() {
		logger.Info("internal metrics/pprof server listening", "addr", listener.Addr().String())
		if err := srv.Serve(listener); err != nil && err != http.ErrServerClosed {
			logger.Error("internal metrics server failed", "err", err)
		}
	}()
	return srv, listener.Addr(), nil
}
