package common

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func newHealthRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	RegisterHealthRoutes(r, nil)
	return r
}

func doGet(r *gin.Engine, path string) *httptest.ResponseRecorder {
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
	r := newHealthRouter()

	SetDraining()
	w := doGet(r, "/readyz")
	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("readyz while draining = %d, want 503", w.Code)
	}
	if body := w.Body.String(); body != `{"status":"draining"}` {
		t.Fatalf("unexpected draining body: %s", body)
	}
}
