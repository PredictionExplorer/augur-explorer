package faq

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// serveQueryProxy routes one POST to the query route through a router
// carrying only the proxy.
func serveQueryProxy(p *Proxy, body io.Reader) *httptest.ResponseRecorder {
	r := httpx.NewRouter()
	p.RegisterRoutes(r)
	req := httptest.NewRequest(http.MethodPost, "/api/cosmicgame/faq/query", body)
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
	w := serveQueryProxy(p, errReader{})
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
	// itself fail — the arm before any network I/O. The response carries a
	// fixed message; the underlying error (which can echo the URL) is only
	// logged.
	var buf bytes.Buffer
	p := New(Options{
		UpstreamURL: "http://bad.example/\x7f",
		Logger:      slog.New(slog.NewTextHandler(&buf, nil)),
	})
	w := serveQueryProxy(p, strings.NewReader(`{}`))
	if w.Code != http.StatusBadGateway {
		t.Fatalf("status = %d\n%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "faq proxy request build failed") {
		t.Fatalf("body = %s", w.Body.String())
	}
	if strings.Contains(w.Body.String(), "control character") {
		t.Fatalf("response echoes the internal error: %s", w.Body.String())
	}
	if !strings.Contains(buf.String(), "invalid control character in URL") {
		t.Fatalf("build failure was not logged: %q", buf.String())
	}
}

func TestProxyOversizedRequestBodyAnswers413(t *testing.T) {
	t.Parallel()
	const limit = 32
	p := New(Options{UpstreamURL: "http://unused.example"})
	r := httpx.NewRouter()
	r.Use(common.MaxRequestBody(limit))
	p.RegisterRoutes(r)

	// An undeclared length skips the middleware's Content-Length pre-check,
	// so the proxy's own io.ReadAll trips the MaxBytesReader wrapper.
	req := httptest.NewRequest(http.MethodPost, "/api/cosmicgame/faq/query",
		undeclaredLengthReader{strings.NewReader(strings.Repeat("q", limit+1))})
	if req.ContentLength != -1 {
		t.Fatalf("ContentLength = %d, want -1 (undeclared)", req.ContentLength)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("status = %d, want 413\n%s", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "request body exceeds the 32-byte limit") {
		t.Fatalf("body = %s", w.Body.String())
	}
}

// undeclaredLengthReader hides its size from httptest.NewRequest so the
// request carries no Content-Length.
type undeclaredLengthReader struct{ r io.Reader }

func (u undeclaredLengthReader) Read(p []byte) (int, error) { return u.r.Read(p) }

func TestProxyUpstreamResponseCap(t *testing.T) {
	t.Parallel()

	t.Run("over the cap answers 502 without relaying", func(t *testing.T) {
		t.Parallel()
		upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write(bytes.Repeat([]byte("a"), maxUpstreamResponseBytes+1))
		}))
		t.Cleanup(upstream.Close)

		var buf bytes.Buffer
		p := New(Options{UpstreamURL: upstream.URL, Logger: slog.New(slog.NewTextHandler(&buf, nil))})
		w := serveQueryProxy(p, strings.NewReader(`{}`))

		if w.Code != http.StatusBadGateway {
			t.Fatalf("status = %d\n%.128s", w.Code, w.Body.String())
		}
		if !strings.Contains(w.Body.String(), "faq upstream response too large") {
			t.Fatalf("body = %.128s", w.Body.String())
		}
		if !strings.Contains(buf.String(), "response exceeds") {
			t.Fatalf("oversized upstream response was not logged: %q", buf.String())
		}
	})

	t.Run("exactly at the cap relays unmodified", func(t *testing.T) {
		t.Parallel()
		payload := bytes.Repeat([]byte("b"), maxUpstreamResponseBytes)
		upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write(payload)
		}))
		t.Cleanup(upstream.Close)

		p := New(Options{UpstreamURL: upstream.URL})
		w := serveQueryProxy(p, strings.NewReader(`{}`))

		if w.Code != http.StatusOK {
			t.Fatalf("status = %d", w.Code)
		}
		if !bytes.Equal(w.Body.Bytes(), payload) {
			t.Fatalf("relayed body differs: got %d bytes, want %d", w.Body.Len(), len(payload))
		}
	})
}

func TestMapFAQPathDefaultsToRoot(t *testing.T) {
	if got := mapFAQPath("/api/cosmicgame/faq/unknown"); got != "/" {
		t.Fatalf("mapFAQPath = %q", got)
	}
}
