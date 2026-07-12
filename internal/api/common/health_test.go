package common

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func newHealthRouter() *httpx.Router {
	r := httpx.NewRouter()
	RegisterHealthRoutes(r, nil)
	return r
}

func newHealthRouterWithPing(ping func(context.Context) error) *httpx.Router {
	r := httpx.NewRouter()
	registerHealthRoutes(r, ping)
	return r
}

func doGet(r *httpx.Router, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
	return w
}

func TestHealthzAlwaysOK(t *testing.T) {
	draining.Store(false)
	t.Cleanup(func() { draining.Store(false) })
	r := newHealthRouter()

	if w := doGet(r, "/healthz"); w.Code != http.StatusOK || w.Body.String() != "ok" {
		t.Fatalf("healthz = %d %q, want 200 ok", w.Code, w.Body.String())
	}

	// Liveness must hold during a drain: the process is still alive.
	SetDraining()
	if w := doGet(r, "/healthz"); w.Code != http.StatusOK {
		t.Fatalf("healthz while draining = %d, want 200", w.Code)
	}
}

func TestReadyzWithoutStore(t *testing.T) {
	draining.Store(false)
	t.Cleanup(func() { draining.Store(false) })
	r := newHealthRouter()

	w := doGet(r, "/readyz")
	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("readyz without store = %d, want 503", w.Code)
	}
	if body := w.Body.String(); body != `{"reason":"database not configured","status":"unready"}` {
		t.Fatalf("unexpected readyz body: %s", body)
	}
}

func TestReadyzDrainingWinsOverEverything(t *testing.T) {
	draining.Store(false)
	t.Cleanup(func() { draining.Store(false) })
	pingCalls := 0
	r := newHealthRouterWithPing(func(context.Context) error {
		pingCalls++
		return nil
	})

	SetDraining()
	w := doGet(r, "/readyz")
	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("readyz while draining = %d, want 503", w.Code)
	}
	if body := w.Body.String(); body != `{"status":"draining"}` {
		t.Fatalf("unexpected draining body: %s", body)
	}
	if pingCalls != 0 {
		t.Fatalf("database ping called %d times while draining", pingCalls)
	}
}

func TestReadyzDatabasePingResults(t *testing.T) {
	draining.Store(false)
	t.Cleanup(func() { draining.Store(false) })

	tests := []struct {
		name       string
		pingErr    error
		wantStatus int
		wantBody   string
	}{
		{
			name:       "ready",
			wantStatus: http.StatusOK,
			wantBody:   `{"status":"ready"}`,
		},
		{
			name:       "database unavailable",
			pingErr:    errors.New("database unavailable"),
			wantStatus: http.StatusServiceUnavailable,
			wantBody:   `{"reason":"database unavailable","status":"unready"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pingCalls := 0
			r := newHealthRouterWithPing(func(ctx context.Context) error {
				pingCalls++
				if ctx == nil {
					t.Fatal("ping received a nil request context")
				}
				return tt.pingErr
			})

			w := doGet(r, "/readyz")
			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}
			if got := w.Body.String(); got != tt.wantBody {
				t.Errorf("body = %q, want %q", got, tt.wantBody)
			}
			if pingCalls != 1 {
				t.Errorf("database ping calls = %d, want 1", pingCalls)
			}
		})
	}
}
