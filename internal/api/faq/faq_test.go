package faq

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// serveProxy routes one request through a router carrying only the proxy.
func serveProxy(p *Proxy, method, path string, body io.Reader) *httptest.ResponseRecorder {
	r := httpx.NewRouter()
	p.RegisterRoutes(r)
	req := httptest.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestNewUpstreamResolution(t *testing.T) {
	t.Parallel()
	t.Run("explicit option wins and is normalized", func(t *testing.T) {
		t.Parallel()
		p := New(Options{UpstreamURL: " http://upstream.example/ "})
		if p.upstreamURL != "http://upstream.example" {
			t.Fatalf("upstreamURL = %q", p.upstreamURL)
		}
	})

	// The AI_BOT_BACKEND_URL / FAQ_BOT_UPSTREAM_URL alias precedence lives
	// in the typed service configuration now (config.APIServer.FAQUpstream,
	// tested there); the module only applies the final default.
	t.Run("default upstream", func(t *testing.T) {
		t.Parallel()
		if p := New(Options{}); p.upstreamURL != "http://127.0.0.1:8000" {
			t.Fatalf("upstreamURL = %q", p.upstreamURL)
		}
	})

	t.Run("construction logs the effective upstream", func(t *testing.T) {
		t.Parallel()
		var buf bytes.Buffer
		New(Options{
			UpstreamURL: "http://logged.example",
			Logger:      slog.New(slog.NewTextHandler(&buf, nil)),
		})
		if !strings.Contains(buf.String(), "upstream=http://logged.example") {
			t.Fatalf("log = %q", buf.String())
		}
	})
}

// errReader fails the request-body read, driving the proxy's first error arm.
type errReader struct{}

func (errReader) Read([]byte) (int, error) { return 0, io.ErrUnexpectedEOF }

func TestProxyBodyReadFailureAnswers502(t *testing.T) {
	t.Parallel()
	var buf bytes.Buffer
	p := New(Options{
		UpstreamURL: "http://unused.example",
		Logger:      slog.New(slog.NewTextHandler(&buf, nil)),
	})
	w := serveProxy(p, http.MethodPost, "/api/cosmicgame/faq/query", errReader{})
	if w.Code != http.StatusBadGateway {
		t.Fatalf("status = %d\n%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "faq proxy read failed") {
		t.Fatalf("body = %s", w.Body.String())
	}
	if !strings.Contains(buf.String(), "faq proxy read body") {
		t.Fatalf("read failure was not logged: %q", buf.String())
	}
}

func TestProxyRequestBuildFailureAnswers502(t *testing.T) {
	// A control character in the upstream URL makes http.NewRequestWithContext
	// itself fail — the arm before any network I/O.
	p := New(Options{UpstreamURL: "http://bad.example/\x7f"})
	w := serveProxy(p, http.MethodPost, "/api/cosmicgame/faq/query", strings.NewReader(`{}`))
	if w.Code != http.StatusBadGateway {
		t.Fatalf("status = %d\n%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "invalid control character in URL") {
		t.Fatalf("body = %s", w.Body.String())
	}
}

func TestMapFAQPathDefaultsToRoot(t *testing.T) {
	if got := mapFAQPath("/api/cosmicgame/faq/unknown"); got != "/" {
		t.Fatalf("mapFAQPath = %q", got)
	}
}
