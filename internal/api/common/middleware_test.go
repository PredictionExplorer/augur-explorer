package common

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func adminProtectedRouter(secret string) *httpx.Router {
	r := httpx.NewRouter()
	r.POST("/admin", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"ok": true})
	}, RequireAdminKey("X-Admin-Key", AdminKey{Name: "TEST_ADMIN_KEY", Value: secret}))
	return r
}

func TestRequireAdminKeyFailsClosedWhenUnset(t *testing.T) {
	t.Parallel()
	r := adminProtectedRouter("")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 when no admin key is configured, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), `"status":0`) {
		t.Fatalf("expected the legacy error envelope, got %q", w.Body.String())
	}
	// The disabled message names the variables the operator must set.
	if !strings.Contains(w.Body.String(), "no admin key configured (TEST_ADMIN_KEY)") {
		t.Fatalf("disabled message must name the env var, got %q", w.Body.String())
	}
}

func TestRequireAdminKeyTreatsWhitespaceValueAsUnset(t *testing.T) {
	t.Parallel()
	r := adminProtectedRouter("   ")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	req.Header.Set("X-Admin-Key", "   ")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 for a whitespace-only configured key, got %d", w.Code)
	}
}

func TestRequireAdminKeyRejectsWrongKey(t *testing.T) {
	t.Parallel()
	r := adminProtectedRouter("correct-key")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	req.Header.Set("X-Admin-Key", "wrong-key")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 for wrong key, got %d", w.Code)
	}
}

func TestRequireAdminKeyAcceptsCorrectKey(t *testing.T) {
	t.Parallel()
	r := adminProtectedRouter("correct-key")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	req.Header.Set("X-Admin-Key", "correct-key")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 for correct key, got %d", w.Code)
	}
}

func TestRequireAdminKeyUsesFallbackKey(t *testing.T) {
	t.Parallel()
	r := httpx.NewRouter()
	r.POST("/admin", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"ok": true})
	}, RequireAdminKey("X-Admin-Key",
		AdminKey{Name: "TEST_PRIMARY_KEY", Value: ""},
		AdminKey{Name: "TEST_FALLBACK_KEY", Value: "fallback-secret"}))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	req.Header.Set("X-Admin-Key", "fallback-secret")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 via fallback key, got %d", w.Code)
	}
}

func TestRequireAdminKeyDisabledMessageNamesAllVars(t *testing.T) {
	t.Parallel()
	r := httpx.NewRouter()
	r.POST("/admin", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"ok": true})
	}, RequireAdminKey("X-Admin-Key",
		AdminKey{Name: "RANKING_ADMIN_KEY"},
		AdminKey{Name: "ADMIN_API_KEY"}))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "(RANKING_ADMIN_KEY or ADMIN_API_KEY)") {
		t.Fatalf("disabled message must join the variable names, got %q", w.Body.String())
	}
}

func TestRateLimitEnforcesBurst(t *testing.T) {
	r := httpx.NewRouter()
	r.GET("/limited", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"ok": true})
	}, RateLimit(0, 2))

	codes := make([]int, 0, 4)
	for range 4 {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/limited", nil)
		req.RemoteAddr = "203.0.113.7:1234"
		r.ServeHTTP(w, req)
		codes = append(codes, w.Code)
	}

	want := []int{http.StatusOK, http.StatusOK, http.StatusTooManyRequests, http.StatusTooManyRequests}
	for i := range want {
		if codes[i] != want[i] {
			t.Fatalf("request statuses = %v, want %v", codes, want)
		}
	}
}

func TestRateLimitConfiguration(t *testing.T) {
	tests := []struct {
		name      string
		rps       float64
		burst     int
		requests  int
		wantCodes []int
		wantCalls int
	}{
		{
			name:      "zero burst rejects every request",
			rps:       10,
			burst:     0,
			requests:  2,
			wantCodes: []int{http.StatusTooManyRequests, http.StatusTooManyRequests},
		},
		{
			name:      "zero rate spends initial burst without refill",
			rps:       0,
			burst:     1,
			requests:  3,
			wantCodes: []int{http.StatusOK, http.StatusTooManyRequests, http.StatusTooManyRequests},
			wantCalls: 1,
		},
		{
			name:      "infinite rate still requires positive burst",
			rps:       math.Inf(1),
			burst:     0,
			requests:  3,
			wantCodes: []int{http.StatusTooManyRequests, http.StatusTooManyRequests, http.StatusTooManyRequests},
		},
		{
			name:      "infinite rate refills immediately",
			rps:       math.Inf(1),
			burst:     1,
			requests:  3,
			wantCodes: []int{http.StatusOK, http.StatusOK, http.StatusOK},
			wantCalls: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlerCalls := 0
			r := httpx.NewRouter()
			r.GET("/limited", func(c *httpx.Context) {
				handlerCalls++
				c.JSON(http.StatusOK, httpx.H{"ok": true})
			}, RateLimit(tt.rps, tt.burst))

			for i := 0; i < tt.requests; i++ {
				w := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/limited", nil)
				req.RemoteAddr = "203.0.113.10:1234"
				r.ServeHTTP(w, req)

				if w.Code != tt.wantCodes[i] {
					t.Errorf("request %d status = %d, want %d", i+1, w.Code, tt.wantCodes[i])
				}
				if tt.wantCodes[i] == http.StatusTooManyRequests {
					if got := w.Body.String(); got != `{"error":"rate limit exceeded, slow down","status":0}` {
						t.Errorf("request %d body = %q", i+1, got)
					}
					if got := w.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
						t.Errorf("request %d Content-Type = %q, want JSON", i+1, got)
					}
				}
			}

			if handlerCalls != tt.wantCalls {
				t.Errorf("handler calls = %d, want %d", handlerCalls, tt.wantCalls)
			}
		})
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
		wantHeaders := map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET, POST, PUT, PATCH, DELETE, OPTIONS",
			"Access-Control-Allow-Headers": "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		}
		for name, want := range wantHeaders {
			if got := w.Header().Get(name); got != want {
				t.Errorf("%s = %q, want %q", name, got, want)
			}
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

func TestRecoveryPassesNormalResponsesWithDefaultLogger(t *testing.T) {
	h := Recovery(nil)(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}))

	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/ok", nil))
	if w.Code != http.StatusAccepted {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusAccepted)
	}
}

func TestRecoveryHandlesOrdinaryErrorPanic(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	h := Recovery(logger)(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		panic(errors.New("ordinary failure"))
	}))

	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/boom", nil))
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", w.Code)
	}
	if w.Body.Len() != 0 {
		t.Errorf("recovery body = %q, want empty", w.Body.String())
	}
}

func TestRecoveryRepanicsAbortHandler(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{name: "direct sentinel", err: http.ErrAbortHandler},
		{name: "wrapped sentinel", err: fmt.Errorf("wrapped: %w", http.ErrAbortHandler)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Recovery(slog.New(slog.NewTextHandler(io.Discard, nil)))(
				http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
					panic(tt.err)
				}),
			)

			var recovered any
			func() {
				defer func() { recovered = recover() }()
				h.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/abort", nil))
			}()

			err, ok := recovered.(error)
			if !ok || !errors.Is(err, http.ErrAbortHandler) {
				t.Fatalf("recovered panic = %#v, want http.ErrAbortHandler", recovered)
			}
		})
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

func TestIsClientDisconnect(t *testing.T) {
	brokenPipe := &net.OpError{Op: "write", Err: errors.New("write: broken pipe")}
	tests := []struct {
		name string
		rec  any
		want bool
	}{
		{name: "non-error panic", rec: "boom", want: false},
		{name: "ordinary error", rec: errors.New("broken pipe"), want: false},
		{name: "broken pipe operation", rec: brokenPipe, want: true},
		{
			name: "connection reset operation",
			rec:  &net.OpError{Op: "write", Err: errors.New("connection reset by peer")},
			want: true,
		},
		{
			name: "matching is case insensitive",
			rec:  &net.OpError{Op: "write", Err: errors.New("BROKEN PIPE")},
			want: true,
		},
		{name: "wrapped operation", rec: fmt.Errorf("render response: %w", brokenPipe), want: true},
		{
			name: "other network operation error",
			rec:  &net.OpError{Op: "write", Err: errors.New("i/o timeout")},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isClientDisconnect(tt.rec); got != tt.want {
				t.Errorf("isClientDisconnect(%#v) = %v, want %v", tt.rec, got, tt.want)
			}
		})
	}
}

func TestRecoveryStaysSilentOnClientDisconnect(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{name: "broken pipe", err: &net.OpError{Op: "write", Err: errors.New("write: broken pipe")}},
		{name: "connection reset", err: &net.OpError{Op: "write", Err: errors.New("connection reset by peer")}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logs := &strings.Builder{}
			r := httpx.NewRouter()
			r.Use(Recovery(slog.New(slog.NewTextHandler(logs, nil))))
			r.GET("/gone", func(c *httpx.Context) {
				panic(tt.err)
			})

			w := httptest.NewRecorder()
			r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/gone", nil))

			// No 500 is written: the client is gone. (The recorder keeps its
			// 200 default because nothing was written at all.)
			if w.Body.Len() != 0 {
				t.Errorf("disconnect recovery must not write a body, got %q", w.Body.String())
			}
			if !strings.Contains(logs.String(), "level=WARN") ||
				!strings.Contains(logs.String(), "client disconnected") {
				t.Errorf("disconnect must be logged as a warning, logs: %s", logs.String())
			}
			if strings.Contains(logs.String(), "panic recovered") {
				t.Errorf("disconnect must not be logged as a server panic, logs: %s", logs.String())
			}
		})
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

func TestAccessLogUsesDefaultLoggerAndCountsBytes(t *testing.T) {
	logs := &strings.Builder{}
	previous := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(logs, nil)))
	t.Cleanup(func() { slog.SetDefault(previous) })

	r := httpx.NewRouter()
	r.Use(AccessLog(nil))
	r.GET("/plain", func(c *httpx.Context) {
		c.String(http.StatusCreated, "abc")
	})

	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/plain", nil))

	line := logs.String()
	for _, want := range []string{
		"method=GET",
		"path=/plain",
		"route=/plain",
		"status=201",
		"bytes=3",
	} {
		if !strings.Contains(line, want) {
			t.Errorf("access log line missing %q: %s", want, line)
		}
	}
}
