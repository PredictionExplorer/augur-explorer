package common

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// bodyLimitRouter builds a router with the cap under test and one POST route
// per API family that reports what the handler could read.
func bodyLimitRouter(limit int64, handler httpx.HandlerFunc) *httpx.Router {
	r := httpx.NewRouter()
	r.Use(MaxRequestBody(limit))
	r.POST("/api/randomwalk/echo", handler)
	r.POST("/api/v2/randomwalk/echo", handler)
	return r
}

// readAllHandler consumes the whole body and answers 200 with the byte count
// or classifies the read failure (413 for the cap, 500 otherwise).
func readAllHandler(c *httpx.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		var maxBytes *http.MaxBytesError
		if errors.As(err, &maxBytes) {
			RespondRequestBodyTooLarge(c, maxBytes.Limit)
			return
		}
		c.JSON(http.StatusInternalServerError, httpx.H{"error": "unexpected read failure"})
		return
	}
	c.JSON(http.StatusOK, httpx.H{"read": len(body)})
}

// undeclaredLengthReader hides its size from httptest.NewRequest so the
// request carries no Content-Length (the chunked/lying-client case).
type undeclaredLengthReader struct{ r io.Reader }

func (u undeclaredLengthReader) Read(p []byte) (int, error) { return u.r.Read(p) }

func TestMaxRequestBodyDeclaredOversizeAnswers413(t *testing.T) {
	t.Parallel()
	const limit = 64
	r := bodyLimitRouter(limit, func(c *httpx.Context) {
		t.Error("handler must not run for a declared oversized body")
	})

	t.Run("legacy envelope outside v2", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/api/randomwalk/echo",
			strings.NewReader(strings.Repeat("x", limit+1))))

		if w.Code != http.StatusRequestEntityTooLarge {
			t.Fatalf("status = %d, want 413\n%s", w.Code, w.Body.String())
		}
		if ct := w.Header().Get("Content-Type"); !strings.HasPrefix(ct, "application/json") {
			t.Fatalf("Content-Type = %q", ct)
		}
		var envelope struct {
			Status int    `json:"status"`
			Error  string `json:"error"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &envelope); err != nil {
			t.Fatalf("decode envelope: %v\n%s", err, w.Body.String())
		}
		if envelope.Status != 0 || envelope.Error != "request body exceeds the 64-byte limit" {
			t.Fatalf("envelope = %+v", envelope)
		}
	})

	t.Run("problem under v2", func(t *testing.T) {
		t.Parallel()
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/api/v2/randomwalk/echo",
			strings.NewReader(strings.Repeat("x", limit+1))))

		if w.Code != http.StatusRequestEntityTooLarge {
			t.Fatalf("status = %d, want 413\n%s", w.Code, w.Body.String())
		}
		if ct := w.Header().Get("Content-Type"); ct != "application/problem+json" {
			t.Fatalf("Content-Type = %q", ct)
		}
		var problem struct {
			Type     string `json:"type"`
			Title    string `json:"title"`
			Status   int    `json:"status"`
			Detail   string `json:"detail"`
			Instance string `json:"instance"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &problem); err != nil {
			t.Fatalf("decode problem: %v\n%s", err, w.Body.String())
		}
		if problem.Type != ProblemTypeBase+"request-too-large" ||
			problem.Title != "Request body too large" ||
			problem.Status != http.StatusRequestEntityTooLarge ||
			problem.Detail != "The request body exceeds the 64-byte limit." ||
			problem.Instance != "/api/v2/randomwalk/echo" {
			t.Fatalf("problem = %+v", problem)
		}
		if !strings.HasSuffix(w.Body.String(), "\n") {
			t.Fatal("problem body must end with a newline (matches the v2 writeProblem rendering)")
		}
	})
}

func TestMaxRequestBodyUndeclaredOversizeFailsAtRead(t *testing.T) {
	t.Parallel()
	const limit = 64
	r := bodyLimitRouter(limit, readAllHandler)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/randomwalk/echo",
		undeclaredLengthReader{strings.NewReader(strings.Repeat("x", limit+1))})
	if req.ContentLength != -1 {
		t.Fatalf("ContentLength = %d, want -1 (undeclared)", req.ContentLength)
	}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("status = %d, want 413 from the handler's MaxBytesError arm\n%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "request body exceeds the 64-byte limit") {
		t.Fatalf("body = %s", w.Body.String())
	}
}

func TestMaxRequestBodyWithinLimitPasses(t *testing.T) {
	t.Parallel()
	const limit = 64
	r := bodyLimitRouter(limit, readAllHandler)

	for name, size := range map[string]int{"under": limit - 1, "exactly at": limit} {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/api/randomwalk/echo",
			strings.NewReader(strings.Repeat("x", size))))
		if w.Code != http.StatusOK {
			t.Fatalf("%s the limit: status = %d, want 200\n%s", name, w.Code, w.Body.String())
		}
		var reply struct {
			Read int `json:"read"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &reply); err != nil || reply.Read != size {
			t.Fatalf("%s the limit: read %d of %d bytes (err=%v)", name, reply.Read, size, err)
		}
	}
}

func TestMaxRequestBodyLeavesBodylessRequestsAlone(t *testing.T) {
	t.Parallel()
	r := httpx.NewRouter()
	r.Use(MaxRequestBody(8))
	r.GET("/api/randomwalk/echo", func(c *httpx.Context) { c.Status(http.StatusOK) })

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/echo", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
}

// TestShouldBindJSONSurfacesMaxBytesError pins the v1 handler contract: a
// bind on a capped body returns *http.MaxBytesError, and the frozen
// {"error": err.Error()} rendering stays a stable Go message.
func TestShouldBindJSONSurfacesMaxBytesError(t *testing.T) {
	t.Parallel()
	const limit = 16
	r := httpx.NewRouter()
	r.Use(MaxRequestBody(limit))
	r.POST("/api/randomwalk/bind", func(c *httpx.Context) {
		var payload struct {
			Field string `json:"field"`
		}
		err := c.ShouldBindJSON(&payload)
		var maxBytes *http.MaxBytesError
		if !errors.As(err, &maxBytes) {
			t.Errorf("ShouldBindJSON error = %v, want *http.MaxBytesError", err)
		}
		if maxBytes != nil && maxBytes.Limit != limit {
			t.Errorf("MaxBytesError.Limit = %d, want %d", maxBytes.Limit, limit)
		}
		c.JSON(http.StatusBadRequest, httpx.H{"error": err.Error()})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/randomwalk/bind",
		undeclaredLengthReader{strings.NewReader(`{"field":"` + strings.Repeat("x", limit*2) + `"}`)})
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want the frozen v1 400 bind-error shape", w.Code)
	}
	if !strings.Contains(w.Body.String(), "http: request body too large") {
		t.Fatalf("body = %s", w.Body.String())
	}
}

func TestProductionBodyCapAdmitsLegitimatePayloads(t *testing.T) {
	t.Parallel()
	if MaxRequestBodyBytes != 1<<20 {
		t.Fatalf("MaxRequestBodyBytes = %d, want 1 MiB (documented in openapi-v2.yaml and operations.md)", MaxRequestBodyBytes)
	}
	// The largest legitimate body on the API is a ranking vote (a few
	// hundred bytes); one MiB leaves two orders of magnitude of headroom.
	r := bodyLimitRouter(MaxRequestBodyBytes, readAllHandler)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/api/randomwalk/echo",
		strings.NewReader(`{"nft1":1,"nft2":2,"nft1_win":1,"sign_nonce":"`+strings.Repeat("a", 64)+`","signature":"0x`+strings.Repeat("b", 130)+`","chain_id":"42161"}`)))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
}
