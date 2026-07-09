// Package faq proxies requests to the Python FAQ bot service (Haystack + Codex MCP).
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

var (
	Info        *log.Logger
	Error       *log.Logger
	Enabled     bool
	upstreamURL string
	httpClient  = &http.Client{Timeout: 180 * time.Second}
)

// Init configures the FAQ proxy. Set AI_BOT_BACKEND_URL (default http://127.0.0.1:8000).
// FAQ_BOT_UPSTREAM_URL is accepted as a legacy alias. Set ENABLE_ROUTES_FAQ=false to disable registration.
func Init(info, errorLog *log.Logger, enabled bool) {
	Enabled = enabled
	Info = info
	Error = errorLog
	upstreamURL = strings.TrimRight(strings.TrimSpace(os.Getenv("AI_BOT_BACKEND_URL")), "/")
	if upstreamURL == "" {
		upstreamURL = strings.TrimRight(strings.TrimSpace(os.Getenv("FAQ_BOT_UPSTREAM_URL")), "/")
	}
	if upstreamURL == "" {
		upstreamURL = "http://127.0.0.1:8000"
	}
	if enabled {
		info.Printf("FAQ proxy enabled; upstream=%s", upstreamURL)
	} else {
		info.Printf("FAQ proxy disabled (ENABLE_ROUTES_FAQ=false)")
	}
}

// RegisterAPIRoutes registers /api/cosmicgame/faq/* proxy routes.
func RegisterAPIRoutes(r *httpx.Router) {
	if !Enabled {
		return
	}
	r.GET("/api/cosmicgame/faq/health", proxyFAQ)
	r.POST("/api/cosmicgame/faq/query", proxyFAQ)
	r.POST("/api/cosmicgame/faq/reindex", proxyFAQ)
}

func proxyFAQ(c *httpx.Context) {
	target := upstreamURL + mapFAQPath(c.Request.URL.Path)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		Error.Printf("faq proxy read body: %v", err)
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

	resp, err := httpClient.Do(req)
	if err != nil {
		Error.Printf("faq proxy upstream %s: %v", target, err)
		c.JSON(http.StatusBadGateway, httpx.H{
			"status":    0,
			"error":     "FAQ service unavailable. Is the Python faq-bot backend running?",
			"component": "faq_bot",
			"upstream":  upstreamURL,
		})
		return
	}
	defer resp.Body.Close()

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
