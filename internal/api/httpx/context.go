package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// H is a shorthand for JSON object literals in handler responses.
// It is a type alias, so H values are plain maps: encoding/json renders them
// with sorted keys, which keeps responses deterministic.
type H = map[string]any

// HandlerFunc is a request handler operating on a Context.
type HandlerFunc func(*Context)

// Context carries one HTTP exchange through a handler. It is created by the
// Router per request and exposes the small helper surface the v1 handlers
// use; anything not covered is reachable through Writer and Request.
type Context struct {
	Writer  ResponseWriter
	Request *http.Request
}

// NewContext builds a Context for w and r outside the Router (tests,
// adapters). The writer is wrapped idempotently.
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{Writer: WrapResponseWriter(w), Request: r}
}

// Param returns the value of the named path parameter ("" when absent).
// For end-of-pattern wildcards ({name...}) the value spans the remaining
// path segments without a leading slash.
func (c *Context) Param(name string) string {
	return c.Request.PathValue(name)
}

// Query returns the first value of the named URL query parameter.
func (c *Context) Query(name string) string {
	return c.Request.URL.Query().Get(name)
}

// GetHeader returns the named request header.
func (c *Context) GetHeader(name string) string {
	return c.Request.Header.Get(name)
}

// FullPath returns the route pattern that matched this request without the
// method prefix (e.g. "/api/cosmicgame/bid/info/{evtlog_id}"), or "" when
// the request was not dispatched through a pattern. Used for log and metric
// labels, never for responses.
func (c *Context) FullPath() string {
	return PatternPath(c.Request)
}

// PatternPath extracts the path part of the ServeMux pattern that matched r
// ("" when unmatched). Patterns look like "GET /a/{b}"; the method prefix
// and optional host are stripped.
func PatternPath(r *http.Request) string {
	pattern := r.Pattern
	if i := strings.IndexByte(pattern, ' '); i >= 0 {
		pattern = pattern[i+1:]
	}
	if i := strings.IndexByte(pattern, '/'); i > 0 {
		// Host-qualified pattern ("host/path"): keep the path only.
		pattern = pattern[i:]
	}
	return pattern
}

// ClientIP returns the IP portion of the request's RemoteAddr, or "" when it
// cannot be parsed. Proxy headers are deliberately not consulted: the
// servers run without trusted proxies, matching the legacy configuration.
func ClientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return ""
	}
	if parsed := net.ParseIP(ip); parsed == nil {
		return ""
	}
	return ip
}

// ShouldBindJSON decodes the request body as JSON into obj. It mirrors the
// legacy binding: a plain decode with no strictness options, so error text
// (including io.EOF for an empty body) is the encoding/json one.
func (c *Context) ShouldBindJSON(obj any) error {
	if c.Request == nil || c.Request.Body == nil {
		return errors.New("invalid request")
	}
	return json.NewDecoder(c.Request.Body).Decode(obj)
}

// JSON renders obj as an application/json response with the given status.
// Marshaling happens before anything is written: a marshal failure (e.g. a
// NaN that escaped sanitization) panics with nothing on the wire, so the
// Recovery middleware can still answer 500 — the exact legacy behavior.
func (c *Context) JSON(status int, obj any) {
	data, err := json.Marshal(obj)
	if err != nil {
		panic(fmt.Errorf("httpx: rendering JSON response: %w", err))
	}
	c.setContentType("application/json; charset=utf-8")
	c.Writer.WriteHeader(status)
	_, _ = c.Writer.Write(data)
}

// String renders a formatted plain-text response with the given status.
func (c *Context) String(status int, format string, values ...any) {
	c.setContentType("text/plain; charset=utf-8")
	c.Writer.WriteHeader(status)
	_, _ = fmt.Fprintf(c.Writer, format, values...)
}

// Data writes raw bytes with the given status and content type (used by the
// reverse proxy to relay upstream responses unchanged).
func (c *Context) Data(status int, contentType string, data []byte) {
	if contentType != "" {
		c.Writer.Header().Set("Content-Type", contentType)
	}
	c.Writer.WriteHeader(status)
	_, _ = c.Writer.Write(data)
}

// Status sends a header-only response with the given status code.
func (c *Context) Status(status int) {
	c.Writer.WriteHeader(status)
}

// File serves the named file, honoring conditional and range requests.
func (c *Context) File(path string) {
	http.ServeFile(c.Writer, c.Request, path)
}

// setContentType sets the Content-Type header unless the handler already
// chose one.
func (c *Context) setContentType(value string) {
	h := c.Writer.Header()
	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", value)
	}
}
