package faq

import (
	"bytes"
	"io"
	"log"
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
	t.Run("explicit option wins and is normalized", func(t *testing.T) {
		p := New(Options{UpstreamURL: " http://upstream.example/ "})
		if p.upstreamURL != "http://upstream.example" {
			t.Fatalf("upstreamURL = %q", p.upstreamURL)
		}
	})

	t.Run("primary env fallback", func(t *testing.T) {
		t.Setenv("AI_BOT_BACKEND_URL", "http://primary.example/")
		t.Setenv("FAQ_BOT_UPSTREAM_URL", "http://legacy.example")
		if p := New(Options{}); p.upstreamURL != "http://primary.example" {
			t.Fatalf("upstreamURL = %q", p.upstreamURL)
		}
	})

	t.Run("legacy alias fallback", func(t *testing.T) {
		t.Setenv("AI_BOT_BACKEND_URL", "")
		t.Setenv("FAQ_BOT_UPSTREAM_URL", "http://legacy.example")
		if p := New(Options{}); p.upstreamURL != "http://legacy.example" {
			t.Fatalf("upstreamURL = %q", p.upstreamURL)
		}
	})

	t.Run("default upstream", func(t *testing.T) {
		t.Setenv("AI_BOT_BACKEND_URL", "")
		t.Setenv("FAQ_BOT_UPSTREAM_URL", "")
		if p := New(Options{}); p.upstreamURL != "http://127.0.0.1:8000" {
			t.Fatalf("upstreamURL = %q", p.upstreamURL)
		}
	})

	t.Run("construction logs the effective upstream", func(t *testing.T) {
		var buf bytes.Buffer
		New(Options{UpstreamURL: "http://logged.example", Info: log.New(&buf, "", 0)})
		if !strings.Contains(buf.String(), "upstream=http://logged.example") {
			t.Fatalf("log = %q", buf.String())
		}
	})
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
