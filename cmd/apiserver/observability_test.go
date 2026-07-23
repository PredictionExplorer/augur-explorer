package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
)

func TestListenInternalServerDisabledWhenUnset(t *testing.T) {
	t.Parallel()
	srv, listener, err := listenInternalServer(t.Context(), "  ", nil)
	if err != nil {
		t.Fatalf("listenInternalServer: %v", err)
	}
	if srv != nil || listener != nil {
		t.Fatal("empty METRICS_ADDR must disable the internal server")
	}
}

func TestListenInternalServerServesMetricsAndJoins(t *testing.T) {
	t.Parallel()
	errorLog := serverErrorLog(slog.New(slog.DiscardHandler))
	srv, listener, err := listenInternalServer(t.Context(), "127.0.0.1:0", errorLog)
	if err != nil {
		t.Fatalf("listenInternalServer: %v", err)
	}
	if srv.ErrorLog != errorLog {
		t.Error("internal server must route its ErrorLog through the process logger")
	}
	done := make(chan error, 1)
	go func() { done <- srv.Serve(listener) }()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://" + listener.Addr().String() + "/metrics")
	if err != nil {
		t.Fatalf("GET /metrics: %v", err)
	}
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET /metrics = %d, want 200", resp.StatusCode)
	}

	shutdownCtx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		t.Fatalf("Shutdown: %v", err)
	}
	select {
	case err := <-done:
		if !errors.Is(err, http.ErrServerClosed) {
			t.Fatalf("Serve returned %v, want %v", err, http.ErrServerClosed)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("internal server goroutine did not join after Shutdown")
	}
}

func TestListenInternalServerReportsBindFailureSynchronously(t *testing.T) {
	t.Parallel()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = ln.Close() })

	srv, listener, err := listenInternalServer(t.Context(), ln.Addr().String(), nil)
	if err == nil {
		t.Fatal("listenInternalServer accepted an occupied address")
	}
	if srv != nil || listener != nil {
		t.Fatal("failed bind returned a partially constructed server")
	}
}

// TestNewPublicServerTimeoutPolicy pins the listener hardening: header,
// body-read and keep-alive-idle bounds are set, WriteTimeout is deliberately
// zero while frozen v1 endpoints stream unbounded arrays to slow clients,
// each TLS server clones the shared config (ServeTLS mutates it), and the
// server's own error records route through the process logger.
func TestNewPublicServerTimeoutPolicy(t *testing.T) {
	t.Parallel()
	handler := http.NewServeMux()
	tlsConfig := &tls.Config{ServerName: "api.example"}
	errorLog := serverErrorLog(slog.New(slog.DiscardHandler))

	srv := newPublicServer(handler, tlsConfig, errorLog)
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
	if srv.ErrorLog != errorLog {
		t.Error("public server must route its ErrorLog through the process logger")
	}

	if plain := newPublicServer(handler, nil, nil); plain.TLSConfig != nil {
		t.Error("plain listener must not carry a TLS config")
	}
}

// TestServerErrorLogRoutesThroughSlog proves stdlib http.Server records land
// on the structured process logger at warning level instead of stderr.
func TestServerErrorLogRoutesThroughSlog(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelWarn}))

	serverErrorLog(logger).Printf("http: TLS handshake error from 198.51.100.7:4242: EOF")

	logged := buf.String()
	if !strings.Contains(logged, "level=WARN") || !strings.Contains(logged, "TLS handshake error") {
		t.Fatalf("server error record not routed through slog at WARN: %q", logged)
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
