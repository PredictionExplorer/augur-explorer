package common

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func adminProtectedRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/admin", RequireAdminKey("X-Admin-Key", "TEST_ADMIN_KEY"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})
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

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/admin", RequireAdminKey("X-Admin-Key", "TEST_PRIMARY_KEY", "TEST_FALLBACK_KEY"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/admin", nil)
	req.Header.Set("X-Admin-Key", "fallback-secret")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 via fallback env var, got %d", w.Code)
	}
}

func TestRateLimitEnforcesBurst(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/limited", RateLimit(1, 2), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

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
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/limited", RateLimit(1, 1), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

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
