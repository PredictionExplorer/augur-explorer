package common

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

var testDeprecatedAt = time.Date(2026, time.July, 16, 0, 0, 0, 0, time.UTC)

// depPolicy returns a policy matching /api/old/* with the standard test
// deprecation moment; sunset is optional.
func depPolicy(sunset time.Time) DeprecationPolicy {
	return DeprecationPolicy{
		Match:        func(r *http.Request) bool { return strings.HasPrefix(r.URL.Path, "/api/old/") },
		DeprecatedAt: testDeprecatedAt,
		LinkURL:      "https://example.test/migration.md",
		SunsetAt:     sunset,
	}
}

// serveThroughDeprecation runs one request through DeprecationHeaders
// wrapping a trivial 200 handler.
func serveThroughDeprecation(t *testing.T, p DeprecationPolicy, method, path string) *httptest.ResponseRecorder {
	t.Helper()
	handler := DeprecationHeaders(p)(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":1}`))
	}))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(method, path, nil))
	return rec
}

func TestDeprecationHeadersMatchedRequest(t *testing.T) {
	rec := serveThroughDeprecation(t, depPolicy(time.Time{}), http.MethodGet, "/api/old/thing")

	// RFC 9745 format: "@" + Unix seconds of the deprecation moment.
	if got, want := rec.Header().Get("Deprecation"), "@1784160000"; got != want {
		t.Errorf("Deprecation = %q, want %q", got, want)
	}
	if got, want := rec.Header().Get("Link"), `<https://example.test/migration.md>; rel="deprecation"; type="text/markdown"`; got != want {
		t.Errorf("Link = %q, want %q", got, want)
	}
	if got := rec.Header().Get("Sunset"); got != "" {
		t.Errorf("Sunset emitted without a configured date: %q", got)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want 200 (middleware must not affect the response itself)", rec.Code)
	}
}

// TestDeprecationHeadersRoundTrip proves the emitted value parses back to
// the configured moment under the RFC 9745 "@" + Unix-seconds grammar.
func TestDeprecationHeadersRoundTrip(t *testing.T) {
	rec := serveThroughDeprecation(t, depPolicy(time.Time{}), http.MethodGet, "/api/old/thing")
	got := rec.Header().Get("Deprecation")
	unix, err := strconv.ParseInt(strings.TrimPrefix(got, "@"), 10, 64)
	if err != nil {
		t.Fatalf("Deprecation %q is not '@' + integer seconds: %v", got, err)
	}
	if !time.Unix(unix, 0).UTC().Equal(testDeprecatedAt) {
		t.Errorf("Deprecation %q decodes to %v, want %v", got, time.Unix(unix, 0).UTC(), testDeprecatedAt)
	}
}

func TestDeprecationHeadersUnmatchedRequest(t *testing.T) {
	rec := serveThroughDeprecation(t, depPolicy(time.Time{}), http.MethodGet, "/api/v2/new/thing")
	for _, header := range []string{"Deprecation", "Link", "Sunset"} {
		if got := rec.Header().Get(header); got != "" {
			t.Errorf("%s emitted on unmatched path: %q", header, got)
		}
	}
}

func TestDeprecationHeadersSunsetConfigured(t *testing.T) {
	sunset := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC)
	rec := serveThroughDeprecation(t, depPolicy(sunset), http.MethodGet, "/api/old/thing")
	if got, want := rec.Header().Get("Sunset"), "Fri, 01 Jan 2027 00:00:00 GMT"; got != want {
		t.Errorf("Sunset = %q, want %q", got, want)
	}
}

func TestDeprecationHeadersSunsetConvertsToUTC(t *testing.T) {
	est := time.FixedZone("EST", -5*3600)
	sunset := time.Date(2026, time.December, 31, 19, 0, 0, 0, est) // = 2027-01-01T00:00:00Z
	rec := serveThroughDeprecation(t, depPolicy(sunset), http.MethodGet, "/api/old/thing")
	if got, want := rec.Header().Get("Sunset"), "Fri, 01 Jan 2027 00:00:00 GMT"; got != want {
		t.Errorf("Sunset = %q, want %q", got, want)
	}
}

func TestDeprecationHeadersHEAD(t *testing.T) {
	rec := serveThroughDeprecation(t, depPolicy(time.Time{}), http.MethodHead, "/api/old/thing")
	if got := rec.Header().Get("Deprecation"); got == "" {
		t.Error("HEAD response missing Deprecation header")
	}
}

func TestDeprecationHeadersNoLinkURL(t *testing.T) {
	p := depPolicy(time.Time{})
	p.LinkURL = ""
	rec := serveThroughDeprecation(t, p, http.MethodGet, "/api/old/thing")
	if got := rec.Header().Get("Link"); got != "" {
		t.Errorf("Link emitted without a URL: %q", got)
	}
	if got := rec.Header().Get("Deprecation"); got == "" {
		t.Error("Deprecation missing")
	}
}

// TestDeprecationHeadersSurvive304 composes the middleware with
// ConditionalETag the way routes.New does (deprecation outermost) and
// proves a revalidated 304 still announces the deprecation.
func TestDeprecationHeadersSurvive304(t *testing.T) {
	r := httpx.NewRouter()
	r.Use(DeprecationHeaders(depPolicy(time.Time{})))
	r.Use(ConditionalETag())
	r.GET("/api/old/data", func(c *httpx.Context) {
		c.JSON(http.StatusOK, httpx.H{"status": 1})
	})

	first := httptest.NewRecorder()
	r.ServeHTTP(first, httptest.NewRequest(http.MethodGet, "/api/old/data", nil))
	etag := first.Header().Get("ETag")
	if etag == "" {
		t.Fatal("no ETag on first response")
	}

	req := httptest.NewRequest(http.MethodGet, "/api/old/data", nil)
	req.Header.Set("If-None-Match", etag)
	second := httptest.NewRecorder()
	r.ServeHTTP(second, req)

	if second.Code != http.StatusNotModified {
		t.Fatalf("revalidation status = %d, want 304", second.Code)
	}
	if got := second.Header().Get("Deprecation"); got == "" {
		t.Error("304 response missing Deprecation header")
	}
	if got := second.Header().Get("Link"); got == "" {
		t.Error("304 response missing Link header")
	}
}

// TestDeprecationHeadersOn404 pins that unrouted paths under a deprecated
// prefix also announce the deprecation: the middleware is a global concern
// wrapping the whole route table, mirroring the routes.New chain position.
func TestDeprecationHeadersOn404(t *testing.T) {
	r := httpx.NewRouter()
	r.Use(DeprecationHeaders(depPolicy(time.Time{})))
	r.GET("/api/old/data", func(c *httpx.Context) { c.JSON(http.StatusOK, httpx.H{"status": 1}) })

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/old/missing", nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
	if got := rec.Header().Get("Deprecation"); got == "" {
		t.Error("404 under deprecated prefix missing Deprecation header")
	}
}

func TestDeprecationHeadersPanicsOnInvalidPolicy(t *testing.T) {
	assertPanics := func(name string, p DeprecationPolicy) {
		t.Helper()
		defer func() {
			if recover() == nil {
				t.Errorf("%s: expected panic", name)
			}
		}()
		DeprecationHeaders(p)
	}
	assertPanics("missing Match", DeprecationPolicy{DeprecatedAt: testDeprecatedAt})
	assertPanics("missing DeprecatedAt", DeprecationPolicy{Match: func(*http.Request) bool { return true }})
}
