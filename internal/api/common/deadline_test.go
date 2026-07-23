package common

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// deadlineRouter builds a router with the deadline under test and one GET
// route per API family driven by the given handler.
func deadlineRouter(policy DeadlinePolicy, onTimeout func(*http.Request), handler httpx.HandlerFunc) *httpx.Router {
	r := httpx.NewRouter()
	r.Use(RequestDeadline(policy, onTimeout))
	r.GET("/api/randomwalk/slow", handler)
	r.GET("/api/v2/randomwalk/slow", handler)
	return r
}

// fixedDeadline returns d for every request.
func fixedDeadline(d time.Duration) DeadlinePolicy {
	return func(*http.Request) time.Duration { return d }
}

// expireAndRespond500 waits for the request deadline to expire (the way a
// context-aware store query would) and then renders the ordinary v1
// internal-error envelope — the exact sequence a handler runs through when
// its query dies of the deadline.
func expireAndRespond500(c *httpx.Context) {
	<-c.Request.Context().Done()
	c.JSON(http.StatusInternalServerError, httpx.H{"status": 0, "error": "Internal server error"})
}

func TestRequestDeadlineSetsContextDeadline(t *testing.T) {
	t.Parallel()
	const budget = time.Minute
	start := time.Now()
	var deadline time.Time
	var ok bool
	r := deadlineRouter(fixedDeadline(budget), nil, func(c *httpx.Context) {
		deadline, ok = c.Request.Context().Deadline()
		c.Status(http.StatusNoContent)
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/slow", nil))
	if w.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", w.Code)
	}
	if !ok {
		t.Fatal("handler context carries no deadline")
	}
	if remaining := time.Until(deadline); remaining > budget || deadline.Before(start) {
		t.Fatalf("deadline %v outside (start, start+%v]", deadline, budget)
	}
}

func TestRequestDeadlineExemptionLeavesContextUnbounded(t *testing.T) {
	t.Parallel()
	r := deadlineRouter(fixedDeadline(0), nil, func(c *httpx.Context) {
		if _, ok := c.Request.Context().Deadline(); ok {
			t.Error("exempt request must carry no deadline")
		}
		c.Status(http.StatusNoContent)
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/slow", nil))
	if w.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", w.Code)
	}
}

// TestRequestDeadlineTimeoutRendersLegacy503 pins the exact legacy-family
// timeout rendering: the handler's post-deadline 500 becomes a 503 envelope
// and the handler's own body never reaches the client.
func TestRequestDeadlineTimeoutRendersLegacy503(t *testing.T) {
	t.Parallel()
	var timedOut []*http.Request
	r := deadlineRouter(fixedDeadline(5*time.Millisecond), func(r *http.Request) {
		timedOut = append(timedOut, r)
	}, expireAndRespond500)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/slow", nil))

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503\n%s", w.Code, w.Body.String())
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Fatalf("Content-Type = %q", ct)
	}
	if got, want := w.Body.String(), `{"error":"request timed out after 5ms","status":0}`; got != want {
		t.Fatalf("body = %s, want %s", got, want)
	}
	if len(timedOut) != 1 {
		t.Fatalf("onTimeout observed %d requests, want exactly 1", len(timedOut))
	}
	if path := timedOut[0].URL.Path; path != "/api/randomwalk/slow" {
		t.Fatalf("onTimeout request path = %q", path)
	}
	if route := httpx.PatternPath(timedOut[0]); route != "/api/randomwalk/slow" {
		t.Fatalf("onTimeout request pattern = %q (metrics need the matched route)", route)
	}
}

// TestRequestDeadlineTimeoutRendersProblem503 pins the RFC 9457 rendering
// under /api/v2/, byte for byte including the trailing newline the generated
// writeProblem emits.
func TestRequestDeadlineTimeoutRendersProblem503(t *testing.T) {
	t.Parallel()
	r := deadlineRouter(fixedDeadline(5*time.Millisecond), nil, func(c *httpx.Context) {
		<-c.Request.Context().Done()
		// The generated v2 layer renders internal problems with 500.
		c.Writer.Header().Set("Content-Type", "application/problem+json")
		c.Writer.WriteHeader(http.StatusInternalServerError)
		_, _ = c.Writer.Write([]byte(`{"detail":"The request could not be completed."}` + "\n"))
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/v2/randomwalk/slow", nil))

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503\n%s", w.Code, w.Body.String())
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
	if problem.Type != ProblemTypeBase+"request-timeout" ||
		problem.Title != "Request timeout" ||
		problem.Status != http.StatusServiceUnavailable ||
		problem.Detail != "The request did not complete within the 5ms processing deadline." ||
		problem.Instance != "/api/v2/randomwalk/slow" {
		t.Fatalf("problem = %+v", problem)
	}
	if !strings.HasSuffix(w.Body.String(), "\n") {
		t.Fatal("problem body must end with a newline (matches the v2 writeProblem rendering)")
	}
}

// TestRequestDeadlineSuccessPastDeadlinePassesThrough pins the deliberate
// non-hijack: a handler that finished its work just past the deadline still
// delivers its 200 result untouched.
func TestRequestDeadlineSuccessPastDeadlinePassesThrough(t *testing.T) {
	t.Parallel()
	r := deadlineRouter(fixedDeadline(time.Millisecond), func(*http.Request) {
		t.Error("onTimeout must not fire for a successful response")
	}, func(c *httpx.Context) {
		<-c.Request.Context().Done()
		c.JSON(http.StatusOK, httpx.H{"late": true})
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/slow", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want the handler's own 200", w.Code)
	}
	if got := w.Body.String(); got != `{"late":true}` {
		t.Fatalf("body = %s", got)
	}
}

// TestRequestDeadlinePreDeadline500PassesThrough pins that an ordinary
// internal error keeps its 500: the middleware only reinterprets failures
// after the deadline expired.
func TestRequestDeadlinePreDeadline500PassesThrough(t *testing.T) {
	t.Parallel()
	r := deadlineRouter(fixedDeadline(time.Minute), func(*http.Request) {
		t.Error("onTimeout must not fire before the deadline")
	}, func(c *httpx.Context) {
		RespondInternalErrorJSON(c)
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/slow", nil))
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want the genuine 500", w.Code)
	}
	if got := w.Body.String(); got != `{"error":"Internal server error","status":0}` {
		t.Fatalf("body = %s", got)
	}
}

// TestRequestDeadlineStartedResponseIsNeverHijacked pins that a response
// whose header section already went out stays untouched even if the handler
// later signals a 5xx (WriteHeader after Write is a no-op status-wise, and
// the body keeps flowing).
func TestRequestDeadlineStartedResponseIsNeverHijacked(t *testing.T) {
	t.Parallel()
	r := deadlineRouter(fixedDeadline(time.Millisecond), func(*http.Request) {
		t.Error("onTimeout must not fire for a started response")
	}, func(c *httpx.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		_, _ = c.Writer.Write([]byte("partial,"))
		<-c.Request.Context().Done()
		c.Writer.WriteHeader(http.StatusInternalServerError) // ignored: already written
		_, _ = c.Writer.Write([]byte("rest"))
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/slow", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want the started 200", w.Code)
	}
	if got := w.Body.String(); got != "partial,rest" {
		t.Fatalf("body = %q", got)
	}
}

// TestRequestDeadlineClientCancelKeepsCurrentBehavior pins that a client
// disconnect (context.Canceled, not DeadlineExceeded) does not trigger the
// timeout rendering: the handler's 500 write proceeds as before (into a dead
// connection in production).
func TestRequestDeadlineClientCancelKeepsCurrentBehavior(t *testing.T) {
	t.Parallel()
	clientGone, cancel := context.WithCancel(context.Background())
	r := deadlineRouter(fixedDeadline(time.Minute), func(*http.Request) {
		t.Error("onTimeout must not fire for a client cancel")
	}, func(c *httpx.Context) {
		cancel()
		<-c.Request.Context().Done()
		RespondInternalErrorJSON(c)
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/slow", nil).WithContext(clientGone))
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want the unhijacked 500", w.Code)
	}
}

// TestRequestDeadlinePreservesRoutePattern pins the pattern copy-back: the
// deadline middleware clones the request (WithContext), and outer middleware
// (access log, metrics) must still see the ServeMux route pattern on the
// original request after dispatch.
func TestRequestDeadlinePreservesRoutePattern(t *testing.T) {
	t.Parallel()
	var outerPattern string
	r := httpx.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(w, req)
			outerPattern = httpx.PatternPath(req)
		})
	})
	r.Use(RequestDeadline(fixedDeadline(time.Minute), nil))
	r.GET("/api/randomwalk/token/{id}", func(c *httpx.Context) { c.Status(http.StatusNoContent) })

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/token/7", nil))
	if w.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", w.Code)
	}
	if outerPattern != "/api/randomwalk/token/{id}" {
		t.Fatalf("outer middleware saw pattern %q, want the matched route (metrics/access-log labels depend on it)", outerPattern)
	}
}

// TestTimeoutWriterDelegatesAndStaysHijacked pins the wrapper's
// ResponseWriter contract: Status/Written/Size/Unwrap delegate to the
// underlying writer (503 on the wire after a hijack) and a handler that
// keeps calling WriteHeader after the hijack cannot disturb the sent
// response.
func TestTimeoutWriterDelegatesAndStaysHijacked(t *testing.T) {
	t.Parallel()
	var probed httpx.ResponseWriter
	r := deadlineRouter(fixedDeadline(5*time.Millisecond), nil, func(c *httpx.Context) {
		<-c.Request.Context().Done()
		probed = c.Writer
		c.Writer.WriteHeader(http.StatusInternalServerError) // hijacked here
		c.Writer.WriteHeader(http.StatusBadGateway)          // post-hijack: swallowed
		_, _ = c.Writer.Write([]byte("late body"))           // swallowed
	})

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/randomwalk/slow", nil))
	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want 503", w.Code)
	}
	if got, want := w.Body.String(), `{"error":"request timed out after 5ms","status":0}`; got != want {
		t.Fatalf("body = %s, want the timeout envelope only", got)
	}
	if probed.Status() != http.StatusServiceUnavailable {
		t.Errorf("Status() = %d, want the on-wire 503", probed.Status())
	}
	if !probed.Written() {
		t.Error("Written() = false after the hijack sent the response")
	}
	if probed.Size() == 0 {
		t.Error("Size() = 0, want the timeout body's bytes")
	}
	if probed.Unwrap() == nil {
		t.Error("Unwrap() = nil, want the underlying writer for http.ResponseController")
	}
}

func TestRequestTimeoutDetail(t *testing.T) {
	t.Parallel()
	if got, want := RequestTimeoutDetail(DefaultRequestDeadline),
		"The request did not complete within the 30s processing deadline."; got != want {
		t.Fatalf("RequestTimeoutDetail = %q, want %q", got, want)
	}
}

// TestDefaultRequestDeadlineValue pins the D22 production budget: 30 seconds,
// 60x the store's slow-query warning threshold, documented in operations.md.
func TestDefaultRequestDeadlineValue(t *testing.T) {
	t.Parallel()
	if DefaultRequestDeadline != 30*time.Second {
		t.Fatalf("DefaultRequestDeadline = %v, want 30s (documented in operations.md; change deliberately or not at all)", DefaultRequestDeadline)
	}
}
