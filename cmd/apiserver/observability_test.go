package main

import (
	"crypto/tls"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
)

// logSink is a mutex-guarded writer for records emitted by server goroutines.
type logSink struct {
	mu  sync.Mutex
	buf strings.Builder
}

func (s *logSink) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.Write(p)
}

func (s *logSink) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.buf.String()
}

func TestStartInternalServerDisabledWhenUnset(t *testing.T) {
	t.Parallel()
	if srv := startInternalServer("  ", slog.New(slog.DiscardHandler)); srv != nil {
		t.Fatal("empty METRICS_ADDR must disable the internal server")
	}
}

func TestStartInternalServerServesMetrics(t *testing.T) {
	t.Parallel()
	sink := &logSink{}
	logger := slog.New(slog.NewTextHandler(sink, nil))
	srv := startInternalServer("127.0.0.1:0", logger)
	if srv == nil {
		t.Fatal("internal server not started")
	}
	t.Cleanup(func() { _ = srv.Close() })

	// ListenAndServe binds asynchronously on the configured address; with
	// port 0 the bound port is not observable, so this test only proves the
	// startup record appears and the server participates in Close.
	deadline := time.Now().Add(5 * time.Second)
	for !strings.Contains(sink.String(), "internal metrics/pprof server listening") {
		if time.Now().After(deadline) {
			t.Fatalf("startup record missing: %q", sink.String())
		}
		time.Sleep(2 * time.Millisecond)
	}
}

func TestStartInternalServerLogsListenFailure(t *testing.T) {
	t.Parallel()
	// Squat the port first so ListenAndServe fails.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = ln.Close() })

	sink := &logSink{}
	logger := slog.New(slog.NewTextHandler(sink, nil))
	srv := startInternalServer(ln.Addr().String(), logger)
	if srv == nil {
		t.Fatal("internal server not constructed")
	}
	t.Cleanup(func() { _ = srv.Close() })

	deadline := time.Now().Add(5 * time.Second)
	for !strings.Contains(sink.String(), "internal metrics server:") {
		if time.Now().After(deadline) {
			t.Fatalf("listen failure was not logged: %q", sink.String())
		}
		time.Sleep(2 * time.Millisecond)
	}
}

// TestNewPublicServerTimeoutPolicy pins the listener hardening: header,
// body-read and keep-alive-idle bounds are set, WriteTimeout is deliberately
// zero while frozen v1 endpoints stream unbounded arrays to slow clients,
// and each TLS server clones the shared config (ServeTLS mutates it).
func TestNewPublicServerTimeoutPolicy(t *testing.T) {
	t.Parallel()
	handler := http.NewServeMux()
	tlsConfig := &tls.Config{ServerName: "api.example"}

	srv := newPublicServer(handler, tlsConfig)
	if srv.ReadHeaderTimeout != readHeaderTimeout || srv.ReadTimeout != readTimeout || srv.IdleTimeout != idleTimeout {
		t.Errorf("timeouts = (header %v, read %v, idle %v), want (%v, %v, %v)",
			srv.ReadHeaderTimeout, srv.ReadTimeout, srv.IdleTimeout,
			readHeaderTimeout, readTimeout, idleTimeout)
	}
	if srv.WriteTimeout != 0 {
		t.Errorf("WriteTimeout = %v, want 0 (v1 streams unbounded arrays; see the constants' doc)", srv.WriteTimeout)
	}
	if srv.TLSConfig == tlsConfig {
		t.Error("TLS config must be cloned per server, not shared")
	}
	if srv.TLSConfig.ServerName != tlsConfig.ServerName {
		t.Error("TLS clone lost the original settings")
	}

	if plain := newPublicServer(handler, nil); plain.TLSConfig != nil {
		t.Error("plain listener must not carry a TLS config")
	}
}

// TestStatusClass pins the metric label mapping.
func TestStatusClass(t *testing.T) {
	t.Parallel()
	for code, want := range map[int]string{
		200: "2xx", 301: "3xx", 404: "4xx", 500: "5xx", 503: "5xx", 101: "2xx",
	} {
		if got := statusClass(code); got != want {
			t.Errorf("statusClass(%d) = %q, want %q", code, got, want)
		}
	}
}

// TestMetricsMiddlewareDeprecatedLabel drives requests through the real
// production router (routes.New with the metrics middleware in its
// documented chain position) and pins the deprecated label to the RFC 9745
// header policy: v1 paths count deprecated="true" — matched or 404 — while
// operational, FAQ-proxy and v2-style paths count deprecated="false". The
// D6 sunset gate reads exactly these series (docs/operations.md).
func TestMetricsMiddlewareDeprecatedLabel(t *testing.T) {
	t.Parallel()
	r := routes.New(nil, routes.Options{
		Extra: []httpx.Middleware{metricsMiddleware()},
		RegisterExtra: func(rt *httpx.Router) {
			// Stands in for a registered v1 route so the matched-route
			// arm is observable without booting the full v1 modules (the
			// boot integration test covers that wiring).
			rt.GET("/api/cosmicgame/observability-probe", func(c *httpx.Context) {
				c.String(http.StatusOK, "ok")
			})
		},
	})

	requests := []struct {
		path       string
		route      string // expected route label
		status     string // expected status-class label
		deprecated string // expected deprecated label
	}{
		{"/api/cosmicgame/observability-probe", "/api/cosmicgame/observability-probe", "2xx", "true"},
		{"/api/randomwalk/observability-missing", "unmatched", "4xx", "true"},
		{"/api/cosmicgame/faq/observability-missing", "unmatched", "4xx", "false"},
		{"/healthz", "/healthz", "2xx", "false"},
	}
	for _, req := range requests {
		before := testutil.ToFloat64(httpRequestsTotal.WithLabelValues(
			http.MethodGet, req.route, req.status, req.deprecated))

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, req.path, nil))

		got := testutil.ToFloat64(httpRequestsTotal.WithLabelValues(
			http.MethodGet, req.route, req.status, req.deprecated))
		if got != before+1 {
			t.Errorf("GET %s: series (route=%q status=%q deprecated=%q) = %v, want %v",
				req.path, req.route, req.status, req.deprecated, got, before+1)
		}

		// The label must agree with the header policy on every response.
		wantHeader := req.deprecated == "true"
		if gotHeader := rec.Header().Get("Deprecation") != ""; gotHeader != wantHeader {
			t.Errorf("GET %s: Deprecation header present = %v, want %v (metric and header policy must agree)",
				req.path, gotHeader, wantHeader)
		}
	}
}
