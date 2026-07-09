package common

import (
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func adminProtectedRouter() *httpx.Router {
	r := httpx.NewRouter()
	r.POST("/admin", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"ok": true})
	}, RequireAdminKey("X-Admin-Key", "TEST_ADMIN_KEY"))
	return r
}

func TestRequireAdminKeyFailsClosedWhenUnset(t *testing.T) {
	t.Setenv("TEST_ADMIN_KEY", "")
	r := adminProtectedRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 when no admin key is configured, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), `"status":0`) {
		t.Fatalf("expected the legacy error envelope, got %q", w.Body.String())
	}
}

func TestRequireAdminKeyRejectsWrongKey(t *testing.T) {
	t.Setenv("TEST_ADMIN_KEY", "correct-key")
	r := adminProtectedRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	req.Header.Set("X-Admin-Key", "wrong-key")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for wrong key, got %d", w.Code)
	}
}

func TestRequireAdminKeyAcceptsCorrectKey(t *testing.T) {
	t.Setenv("TEST_ADMIN_KEY", "correct-key")
	r := adminProtectedRouter()

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	req.Header.Set("X-Admin-Key", "correct-key")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 for correct key, got %d", w.Code)
	}
}

func TestRequireAdminKeyUsesFallbackEnvVar(t *testing.T) {
	t.Setenv("TEST_PRIMARY_KEY", "")
	t.Setenv("TEST_FALLBACK_KEY", "fallback-secret")

	r := httpx.NewRouter()
	r.POST("/admin", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"ok": true})
	}, RequireAdminKey("X-Admin-Key", "TEST_PRIMARY_KEY", "TEST_FALLBACK_KEY"))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	req.Header.Set("X-Admin-Key", "fallback-secret")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 via fallback env var, got %d", w.Code)
	}
}

func TestRateLimitEnforcesBurst(t *testing.T) {
	r := httpx.NewRouter()
	r.GET("/limited", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"ok": true})
	}, RateLimit(1, 2))

	codes := make([]int, 0, 4)
	for range 4 {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/limited", nil)
		req.RemoteAddr = "203.0.113.7:1234"
		r.ServeHTTP(w, req)
		codes = append(codes, w.Code)
	}

	if codes[0] != http.StatusOK || codes[1] != http.StatusOK {
		t.Fatalf("first two requests should pass the burst, got %v", codes)
	}
	if codes[2] != http.StatusTooManyRequests && codes[3] != http.StatusTooManyRequests {
		t.Fatalf("expected 429 once burst exhausted, got %v", codes)
	}
}

func TestRateLimitIsolatesClients(t *testing.T) {
	r := httpx.NewRouter()
	r.GET("/limited", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"ok": true})
	}, RateLimit(1, 1))

	// Exhaust the bucket for client A.
	for range 2 {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/limited", nil)
		req.RemoteAddr = "203.0.113.1:1111"
		r.ServeHTTP(w, req)
	}

	// Client B must be unaffected.
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/limited", nil)
	req.RemoteAddr = "203.0.113.2:2222"
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected different client to have its own bucket, got %d", w.Code)
	}
}

func TestCORSSetsHeadersAndShortCircuitsOptions(t *testing.T) {
	r := httpx.NewRouter()
	r.Use(CORS())
	r.GET("/x", func(c *httpx.Context) { c.Status(http.StatusOK) })

	t.Run("headers_on_normal_response", func(t *testing.T) {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/x", nil))
		if w.Header().Get("Access-Control-Allow-Origin") != "*" {
			t.Errorf("missing CORS origin header")
		}
	})

	t.Run("options_answers_204_for_any_path", func(t *testing.T) {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodOptions, "/anything/at/all", nil))
		if w.Code != http.StatusNoContent {
			t.Errorf("OPTIONS = %d, want 204", w.Code)
		}
		if w.Body.Len() != 0 {
			t.Errorf("OPTIONS body = %q, want empty", w.Body.String())
		}
		if w.Header().Get("Access-Control-Allow-Methods") == "" {
			t.Errorf("preflight response is missing the methods header")
		}
	})

	t.Run("headers_on_404", func(t *testing.T) {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/missing", nil))
		if w.Code != http.StatusNotFound || w.Header().Get("Access-Control-Allow-Origin") != "*" {
			t.Errorf("404 must still carry CORS headers: %d %q",
				w.Code, w.Header().Get("Access-Control-Allow-Origin"))
		}
	})
}

func TestRecoveryTurnsPanicsInto500(t *testing.T) {
	logs := &strings.Builder{}
	logger := slog.New(slog.NewTextHandler(logs, nil))

	r := httpx.NewRouter()
	r.Use(CORS(), Recovery(logger))
	r.GET("/boom", func(c *httpx.Context) { panic("kaboom") })

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/boom", nil))

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", w.Code)
	}
	// Headers set before the panic (CORS runs first) survive.
	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("headers set before the panic must survive recovery")
	}
	if !strings.Contains(logs.String(), "kaboom") {
		t.Error("panic value must be logged")
	}
}

func TestRecoveryDoesNotOverrideStartedResponse(t *testing.T) {
	r := httpx.NewRouter()
	r.Use(Recovery(slog.New(slog.NewTextHandler(io.Discard, nil))))
	r.GET("/late-boom", func(c *httpx.Context) {
		c.Status(http.StatusOK)
		panic("after header")
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/late-boom", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d; recovery must not rewrite an already-sent header", w.Code)
	}
}

func TestRecoveryStaysSilentOnClientDisconnect(t *testing.T) {
	logs := &strings.Builder{}
	r := httpx.NewRouter()
	r.Use(Recovery(slog.New(slog.NewTextHandler(logs, nil))))
	r.GET("/gone", func(c *httpx.Context) {
		panic(&net.OpError{Op: "write", Err: errors.New("write: broken pipe")})
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/gone", nil))

	// No 500 is written: the client is gone. (The recorder keeps its 200
	// default because nothing was written at all.)
	if w.Body.Len() != 0 {
		t.Errorf("disconnect recovery must not write a body, got %q", w.Body.String())
	}
	if !strings.Contains(logs.String(), "client disconnected") {
		t.Errorf("disconnect must be logged as a warning, logs: %s", logs.String())
	}
}

func TestAccessLogEmitsStructuredFields(t *testing.T) {
	logs := &strings.Builder{}
	logger := slog.New(slog.NewTextHandler(logs, nil))

	r := httpx.NewRouter()
	r.Use(AccessLog(logger))
	r.GET("/api/things/{id}", func(c *httpx.Context) { c.JSON(http.StatusOK, httpx.H{"ok": true}) })

	req := httptest.NewRequest(http.MethodGet, "/api/things/7?limit=2", nil)
	req.RemoteAddr = "203.0.113.5:9999"
	r.ServeHTTP(httptest.NewRecorder(), req)

	line := logs.String()
	for _, want := range []string{
		"method=GET",
		`path="/api/things/7?limit=2"`,
		"route=/api/things/{id}",
		"status=200",
		"ip=203.0.113.5",
	} {
		if !strings.Contains(line, want) {
			t.Errorf("access log line missing %q: %s", want, line)
		}
	}
}
