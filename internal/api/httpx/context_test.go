package httpx

import (
	"math"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// routedContext dispatches one request through a single-route router and
// returns the Context observed by the handler.
func routedContext(t *testing.T, pattern, path string) *Context {
	t.Helper()
	var captured *Context
	r := NewRouter()
	r.GET(pattern, func(c *Context) {
		captured = c
		c.Status(http.StatusNoContent)
	})
	r.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, path, nil))
	if captured == nil {
		t.Fatalf("handler for %q was not invoked by %q", pattern, path)
	}
	return captured
}

func TestContextParam(t *testing.T) {
	c := routedContext(t, "/users/{user_addr}/bids/{offset}", "/users/0xabc/bids/15")
	if got := c.Param("user_addr"); got != "0xabc" {
		t.Errorf("Param(user_addr) = %q, want %q", got, "0xabc")
	}
	if got := c.Param("offset"); got != "15" {
		t.Errorf("Param(offset) = %q, want %q", got, "15")
	}
	if got := c.Param("missing"); got != "" {
		t.Errorf("Param(missing) = %q, want empty", got)
	}
}

func TestContextParamWildcardHasNoLeadingSlash(t *testing.T) {
	c := routedContext(t, "/images/{filepath...}", "/images/new/cosmic/0xseed.png")
	if got := c.Param("filepath"); got != "new/cosmic/0xseed.png" {
		t.Errorf("wildcard Param = %q, want %q", got, "new/cosmic/0xseed.png")
	}
}

func TestContextQuery(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/x?limit=25&sort=desc&empty=", nil)
	c := NewContext(httptest.NewRecorder(), req)
	if got := c.Query("limit"); got != "25" {
		t.Errorf("Query(limit) = %q, want 25", got)
	}
	if got := c.Query("sort"); got != "desc" {
		t.Errorf("Query(sort) = %q, want desc", got)
	}
	if got := c.Query("empty"); got != "" {
		t.Errorf("Query(empty) = %q, want empty", got)
	}
	if got := c.Query("absent"); got != "" {
		t.Errorf("Query(absent) = %q, want empty", got)
	}
}

func TestContextGetHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("X-Admin-Key", "sekrit")
	c := NewContext(httptest.NewRecorder(), req)
	if got := c.GetHeader("X-Admin-Key"); got != "sekrit" {
		t.Errorf("GetHeader = %q, want sekrit", got)
	}
}

func TestContextFullPath(t *testing.T) {
	c := routedContext(t, "/api/bid/info/{evtlog_id}", "/api/bid/info/42")
	if got := c.FullPath(); got != "/api/bid/info/{evtlog_id}" {
		t.Errorf("FullPath = %q, want the pattern path", got)
	}

	// Unrouted request: no pattern was matched.
	c2 := NewContext(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/nowhere", nil))
	if got := c2.FullPath(); got != "" {
		t.Errorf("FullPath without a match = %q, want empty", got)
	}
}

func TestPatternPathStripsHost(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/p", nil)
	req.Pattern = "GET example.com/p/{id}"
	if got := PatternPath(req); got != "/p/{id}" {
		t.Errorf("PatternPath = %q, want /p/{id}", got)
	}
}

func TestClientIP(t *testing.T) {
	cases := []struct {
		remote string
		want   string
	}{
		{"203.0.113.9:1234", "203.0.113.9"},
		{"[2001:db8::1]:443", "2001:db8::1"},
		{"no-port", ""},
		{"not-an-ip:80", ""},
		{"", ""},
	}
	for _, tc := range cases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = tc.remote
		if got := ClientIP(req); got != tc.want {
			t.Errorf("ClientIP(%q) = %q, want %q", tc.remote, got, tc.want)
		}
	}
}

func TestShouldBindJSON(t *testing.T) {
	type payload struct {
		BidID    int64  `json:"bid_id"`
		UserAddr string `json:"user_addr"`
	}

	t.Run("valid", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"bid_id":7,"user_addr":"0xabc"}`))
		var p payload
		if err := NewContext(httptest.NewRecorder(), req).ShouldBindJSON(&p); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if p.BidID != 7 || p.UserAddr != "0xabc" {
			t.Errorf("decoded %+v", p)
		}
	})

	t.Run("empty_body_is_EOF", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
		var p payload
		err := NewContext(httptest.NewRecorder(), req).ShouldBindJSON(&p)
		if err == nil || err.Error() != "EOF" {
			t.Fatalf("want EOF for empty body (legacy binding parity), got %v", err)
		}
	})

	t.Run("malformed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"bid_id":`))
		var p payload
		if err := NewContext(httptest.NewRecorder(), req).ShouldBindJSON(&p); err == nil {
			t.Fatal("want error for malformed JSON")
		}
	})

	t.Run("wrong_type", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"bid_id":"nope"}`))
		var p payload
		if err := NewContext(httptest.NewRecorder(), req).ShouldBindJSON(&p); err == nil {
			t.Fatal("want error for type mismatch")
		}
	})
}

func TestContextJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c := NewContext(w, httptest.NewRequest(http.MethodGet, "/", nil))
	c.JSON(http.StatusCreated, H{"status": 1, "error": ""})

	if w.Code != http.StatusCreated {
		t.Errorf("status = %d, want 201", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Errorf("content type = %q", ct)
	}
	// Map keys are rendered sorted: deterministic wire bytes.
	if body := w.Body.String(); body != `{"error":"","status":1}` {
		t.Errorf("body = %q", body)
	}
}

func TestContextJSONKeepsExistingContentType(t *testing.T) {
	w := httptest.NewRecorder()
	c := NewContext(w, httptest.NewRequest(http.MethodGet, "/", nil))
	c.Writer.Header().Set("Content-Type", "application/problem+json")
	c.JSON(http.StatusOK, H{})
	if ct := w.Header().Get("Content-Type"); ct != "application/problem+json" {
		t.Errorf("content type overridden to %q", ct)
	}
}

func TestContextJSONPanicsBeforeWritingOnMarshalFailure(t *testing.T) {
	w := httptest.NewRecorder()
	c := NewContext(w, httptest.NewRequest(http.MethodGet, "/", nil))

	defer func() {
		if recover() == nil {
			t.Fatal("want panic for unmarshalable value")
		}
		// Nothing may have reached the wire: Recovery still owns the response.
		if c.Writer.Written() || w.Body.Len() > 0 {
			t.Errorf("response started before marshal error: status=%d body=%q", w.Code, w.Body.String())
		}
	}()
	c.JSON(http.StatusOK, H{"bad": math.NaN()})
}

func TestContextString(t *testing.T) {
	w := httptest.NewRecorder()
	c := NewContext(w, httptest.NewRequest(http.MethodGet, "/", nil))
	c.String(http.StatusOK, "ok")

	if w.Code != http.StatusOK || w.Body.String() != "ok" {
		t.Errorf("got %d %q, want 200 ok", w.Code, w.Body.String())
	}
	if ct := w.Header().Get("Content-Type"); ct != "text/plain; charset=utf-8" {
		t.Errorf("content type = %q", ct)
	}

	w2 := httptest.NewRecorder()
	NewContext(w2, httptest.NewRequest(http.MethodGet, "/", nil)).String(http.StatusOK, "%d-%s", 7, "x")
	if w2.Body.String() != "7-x" {
		t.Errorf("formatted body = %q", w2.Body.String())
	}
}

func TestContextData(t *testing.T) {
	w := httptest.NewRecorder()
	c := NewContext(w, httptest.NewRequest(http.MethodGet, "/", nil))
	c.Data(http.StatusBadGateway, "application/json", []byte(`{"a":1}`))

	if w.Code != http.StatusBadGateway {
		t.Errorf("status = %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("content type = %q", ct)
	}
	if w.Body.String() != `{"a":1}` {
		t.Errorf("body = %q", w.Body.String())
	}
}

func TestContextStatus(t *testing.T) {
	w := httptest.NewRecorder()
	NewContext(w, httptest.NewRequest(http.MethodGet, "/", nil)).Status(http.StatusNotFound)
	if w.Code != http.StatusNotFound || w.Body.Len() != 0 {
		t.Errorf("got %d body %q, want bare 404", w.Code, w.Body.String())
	}
}

func TestContextFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "asset.txt")
	if err := os.WriteFile(path, []byte("payload"), 0o600); err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	NewContext(w, httptest.NewRequest(http.MethodGet, "/asset.txt", nil)).File(path)
	if w.Code != http.StatusOK || w.Body.String() != "payload" {
		t.Errorf("got %d %q", w.Code, w.Body.String())
	}

	w2 := httptest.NewRecorder()
	NewContext(w2, httptest.NewRequest(http.MethodGet, "/missing", nil)).File(filepath.Join(dir, "missing"))
	if w2.Code != http.StatusNotFound {
		t.Errorf("missing file: got %d, want 404", w2.Code)
	}
}
