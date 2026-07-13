// Package faq proxies requests to the Python FAQ bot service (Haystack + Codex MCP).
// The proxy is an injected value (no package-level state); route registration
// is a method the shared router constructor calls.
package faq

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// Proxy forwards FAQ requests to the configured upstream.
type Proxy struct {
	upstreamURL string
	client      *http.Client
	info        *log.Logger
	errlog      *log.Logger
}

// Options configures New. An empty UpstreamURL falls back to
// AI_BOT_BACKEND_URL, then the legacy FAQ_BOT_UPSTREAM_URL alias, then
// http://127.0.0.1:8000. Client defaults to a 180s-timeout HTTP client.
type Options struct {
	UpstreamURL string
	Client      *http.Client
	Info        *log.Logger
	Error       *log.Logger
}

// New builds the FAQ proxy and logs the effective upstream.
func New(opts Options) *Proxy {
	discard := log.New(io.Discard, "", 0)
	if opts.Info == nil {
		opts.Info = discard
	}
	if opts.Error == nil {
		opts.Error = discard
	}
	if opts.Client == nil {
		opts.Client = &http.Client{Timeout: 180 * time.Second}
	}
	upstream := strings.TrimRight(strings.TrimSpace(opts.UpstreamURL), "/")
	if upstream == "" {
		upstream = strings.TrimRight(strings.TrimSpace(os.Getenv("AI_BOT_BACKEND_URL")), "/")
	}
	if upstream == "" {
		upstream = strings.TrimRight(strings.TrimSpace(os.Getenv("FAQ_BOT_UPSTREAM_URL")), "/")
	}
	if upstream == "" {
		upstream = "http://127.0.0.1:8000"
	}
	p := &Proxy{
		upstreamURL: upstream,
		client:      opts.Client,
		info:        opts.Info,
		errlog:      opts.Error,
	}
	p.info.Printf("FAQ proxy enabled; upstream=%s", p.upstreamURL)
	return p
}

// RegisterRoutes registers /api/cosmicgame/faq/* proxy routes.
func (p *Proxy) RegisterRoutes(r *httpx.Router) {
	r.GET("/api/cosmicgame/faq/health", p.proxyFAQ)
	r.POST("/api/cosmicgame/faq/query", p.proxyFAQ)
	r.POST("/api/cosmicgame/faq/reindex", p.proxyFAQ)
}

func (p *Proxy) proxyFAQ(c *httpx.Context) {
	target := p.upstreamURL + mapFAQPath(c.Request.URL.Path)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		p.errlog.Printf("faq proxy read body: %v", err)
		c.JSON(http.StatusBadGateway, httpx.H{"status": 0, "error": "faq proxy read failed"})
		return
	}

	req, err := http.NewRequestWithContext(c.Request.Context(), c.Request.Method, target, bytes.NewReader(body))
	if err != nil {
		c.JSON(http.StatusBadGateway, httpx.H{"status": 0, "error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if accept := c.GetHeader("Accept"); accept != "" {
		req.Header.Set("Accept", accept)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		p.errlog.Printf("faq proxy upstream %s: %v", target, err)
		c.JSON(http.StatusBadGateway, httpx.H{
			"status":    0,
			"error":     "FAQ service unavailable. Is the Python faq-bot backend running?",
			"component": "faq_bot",
			"upstream":  p.upstreamURL,
		})
		return
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, httpx.H{"status": 0, "error": "faq proxy response read failed"})
		return
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/json"
	}
	c.Data(resp.StatusCode, contentType, respBody)
}

func mapFAQPath(path string) string {
	switch path {
	case "/api/cosmicgame/faq/health":
		return "/health"
	case "/api/cosmicgame/faq/query":
		return "/api/query"
	case "/api/cosmicgame/faq/reindex":
		return "/api/reindex"
	default:
		return "/"
	}
}
