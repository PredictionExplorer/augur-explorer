// Package faq proxies requests to the Python FAQ bot service (Haystack + Codex MCP).
// The proxy is an injected value (no package-level state); route registration
// is a method the shared router constructor calls.
package faq

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// maxUpstreamResponseBytes caps how much of an upstream FAQ response the
// proxy buffers before relaying it (4 MiB). A misbehaving upstream must not
// be able to exhaust API server memory; legitimate FAQ answers are a few KB.
const maxUpstreamResponseBytes = 4 << 20

// Proxy forwards FAQ requests to the configured upstream.
type Proxy struct {
	upstreamURL string
	client      *http.Client
	logger      *slog.Logger
}

// Options configures New. UpstreamURL is the configured backend
// (AI_BOT_BACKEND_URL / the legacy FAQ_BOT_UPSTREAM_URL alias, resolved by
// the typed service configuration); empty applies the http://127.0.0.1:8000
// default. Client defaults to a 180s-timeout HTTP client; a nil Logger
// discards.
type Options struct {
	UpstreamURL string
	Client      *http.Client
	Logger      *slog.Logger
}

// New builds the FAQ proxy and logs the effective upstream.
func New(opts Options) *Proxy {
	if opts.Logger == nil {
		opts.Logger = slog.New(slog.DiscardHandler)
	}
	if opts.Client == nil {
		opts.Client = &http.Client{Timeout: 180 * time.Second}
	}
	upstream := strings.TrimRight(strings.TrimSpace(opts.UpstreamURL), "/")
	if upstream == "" {
		upstream = "http://127.0.0.1:8000"
	}
	p := &Proxy{
		upstreamURL: upstream,
		client:      opts.Client,
		logger:      opts.Logger,
	}
	p.logger.Info("FAQ proxy enabled", "upstream", p.upstreamURL)
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
		// The shared MaxRequestBody middleware bounds the body; a read past
		// the cap is the client's fault (413), any other failure is not.
		var maxBytes *http.MaxBytesError
		if errors.As(err, &maxBytes) {
			common.RespondRequestBodyTooLarge(c, maxBytes.Limit)
			return
		}
		p.logger.Error(fmt.Sprintf("faq proxy read body: %v", err))
		c.JSON(http.StatusBadGateway, httpx.H{"status": 0, "error": "faq proxy read failed"})
		return
	}

	req, err := http.NewRequestWithContext(c.Request.Context(), c.Request.Method, target, bytes.NewReader(body))
	if err != nil {
		p.logger.Error(fmt.Sprintf("faq proxy build request for %s: %v", target, err))
		c.JSON(http.StatusBadGateway, httpx.H{"status": 0, "error": "faq proxy request build failed"})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if accept := c.GetHeader("Accept"); accept != "" {
		req.Header.Set("Accept", accept)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		p.logger.Error(fmt.Sprintf("faq proxy upstream %s: %v", target, err))
		c.JSON(http.StatusBadGateway, httpx.H{
			"status":    0,
			"error":     "FAQ service unavailable. Is the Python faq-bot backend running?",
			"component": "faq_bot",
			"upstream":  p.upstreamURL,
		})
		return
	}
	defer func() { _ = resp.Body.Close() }()

	// Read one byte past the cap so a response of exactly the cap size
	// still relays while anything larger is rejected, not truncated.
	respBody, err := io.ReadAll(io.LimitReader(resp.Body, maxUpstreamResponseBytes+1))
	if err != nil {
		c.JSON(http.StatusBadGateway, httpx.H{"status": 0, "error": "faq proxy response read failed"})
		return
	}
	if len(respBody) > maxUpstreamResponseBytes {
		p.logger.Error(fmt.Sprintf("faq proxy upstream %s: response exceeds %d bytes", target, maxUpstreamResponseBytes))
		c.JSON(http.StatusBadGateway, httpx.H{"status": 0, "error": "faq upstream response too large"})
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
