package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// TestRequestDeadlineForPolicy pins the route-family deadline policy: every
// surface carries the fixed production budget except the FAQ proxy, whose
// upstream is bounded by its own client timeout.
func TestRequestDeadlineForPolicy(t *testing.T) {
	t.Parallel()
	cases := map[string]time.Duration{
		"/api/cosmicgame/dashboard":       common.DefaultRequestDeadline,
		"/api/cosmicgame/faq/query":       0,
		"/api/cosmicgame/faq/health":      0,
		"/api/cosmicgame/faq/reindex":     0,
		"/api/cosmicgame/faqish":          common.DefaultRequestDeadline, // prefix must not over-match
		"/api/randomwalk/tokens":          common.DefaultRequestDeadline,
		"/api/v2/cosmicgame/rounds":       common.DefaultRequestDeadline,
		"/healthz":                        common.DefaultRequestDeadline,
		"/version":                        common.DefaultRequestDeadline,
		"/metadata/7":                     common.DefaultRequestDeadline,
		"/api/v2/randomwalk/ranking/hits": common.DefaultRequestDeadline,
	}
	for path, want := range cases {
		if got := RequestDeadlineFor(path); got != want {
			t.Errorf("RequestDeadlineFor(%q) = %v, want %v", path, got, want)
		}
	}
}

// TestNewInstallsRequestDeadline proves the production constructor puts the
// deadline on every non-exempt request context (probed through RegisterExtra
// so no module wiring is needed) and leaves FAQ-prefixed paths unbounded.
func TestNewInstallsRequestDeadline(t *testing.T) {
	t.Parallel()
	type probe struct {
		deadline time.Time
		bounded  bool
	}
	probes := map[string]*probe{}
	record := func(c *httpx.Context) {
		d, ok := c.Request.Context().Deadline()
		probes[c.Request.URL.Path] = &probe{deadline: d, bounded: ok}
		c.Status(http.StatusNoContent)
	}
	r := New(nil, Options{
		RegisterExtra: func(r *httpx.Router) {
			r.GET("/api/cosmicgame/probe", record)
			r.GET("/api/cosmicgame/faq/probe", record)
		},
	})

	start := time.Now()
	for _, path := range []string{"/api/cosmicgame/probe", "/api/cosmicgame/faq/probe"} {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, path, nil))
		if w.Code != http.StatusNoContent {
			t.Fatalf("%s: status = %d, want 204\n%s", path, w.Code, w.Body.String())
		}
	}

	bounded := probes["/api/cosmicgame/probe"]
	if !bounded.bounded {
		t.Fatal("non-exempt request carries no deadline through the production chain")
	}
	// The deadline is stamped when the request is served, a beat after
	// start — allow one second of scheduling slack above the budget.
	if remaining := bounded.deadline.Sub(start); remaining <= 0 || remaining > common.DefaultRequestDeadline+time.Second {
		t.Fatalf("deadline %v from request start, want within (0, %v+1s]", remaining, common.DefaultRequestDeadline)
	}
	if probes["/api/cosmicgame/faq/probe"].bounded {
		t.Fatal("FAQ-prefixed request must stay exempt from the request deadline")
	}
}
