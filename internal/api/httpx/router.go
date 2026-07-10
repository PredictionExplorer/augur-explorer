package httpx

import (
	"net/http"
	"net/url"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
)

// Middleware wraps an http.Handler with cross-cutting behavior. Global
// middleware (Use) surrounds the whole route table — including 404/405
// responses — while per-route middleware surrounds one handler.
type Middleware func(http.Handler) http.Handler

// Route is one registered (method, pattern) pair. Pattern uses ServeMux
// syntax: "/a/{param}/b", with "{name...}" for trailing wildcards.
type Route struct {
	Method  string
	Pattern string
}

// Router is a small routing layer over http.ServeMux. It adds: a route
// registry (Routes), global and per-route middleware, HandlerFunc adaption,
// and the legacy trailing-slash redirect (a request that only differs from
// a registered route by one trailing slash is redirected, 301 for GET and
// 307 otherwise, preserving the query string).
//
// Registration (Use, GET, POST, ...) must finish before the first request
// is served; the middleware chain is frozen at that point.
type Router struct {
	mux        *http.ServeMux
	middleware []Middleware
	routes     []Route

	buildOnce sync.Once
	handler   http.Handler
	frozen    atomic.Bool
}

// NewRouter returns an empty Router.
func NewRouter() *Router {
	return &Router{mux: http.NewServeMux()}
}

// Use appends global middleware; the first registered runs outermost. It
// panics when called after the router has started serving.
func (r *Router) Use(mw ...Middleware) {
	if r.frozen.Load() {
		panic("httpx: Use called after the router started serving")
	}
	r.middleware = append(r.middleware, mw...)
}

// GET registers a GET handler (per ServeMux semantics it also answers HEAD
// requests, with the body discarded by net/http).
func (r *Router) GET(pattern string, h HandlerFunc, mw ...Middleware) {
	r.Handle(http.MethodGet, pattern, h, mw...)
}

// POST registers a POST handler.
func (r *Router) POST(pattern string, h HandlerFunc, mw ...Middleware) {
	r.Handle(http.MethodPost, pattern, h, mw...)
}

// HEAD registers a HEAD handler.
func (r *Router) HEAD(pattern string, h HandlerFunc, mw ...Middleware) {
	r.Handle(http.MethodHead, pattern, h, mw...)
}

// Handle registers a handler for the method and pattern, wrapped in the
// given per-route middleware (first runs outermost). Conflicting or
// malformed patterns panic at registration, so a bad route table fails at
// startup, never at request time.
func (r *Router) Handle(method, pattern string, h HandlerFunc, mw ...Middleware) {
	if r.frozen.Load() {
		panic("httpx: Handle called after the router started serving")
	}
	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		h(NewContext(w, req))
	})
	for i := len(mw) - 1; i >= 0; i-- {
		handler = mw[i](handler)
	}
	r.mux.Handle(method+" "+pattern, handler)
	r.routes = append(r.routes, Route{Method: method, Pattern: pattern})
}

// HandleFunc implements the narrow mux interface used by generated stdlib
// OpenAPI servers. pattern must use the ServeMux "METHOD /path" form. The
// route is retained in the registry so generated and handwritten operations
// participate in the same drift tests and conflict checks.
func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if r.frozen.Load() {
		panic("httpx: HandleFunc called after the router started serving")
	}
	method, routePattern, ok := strings.Cut(pattern, " ")
	if !ok || method == "" || routePattern == "" {
		panic("httpx: HandleFunc requires a METHOD /path pattern")
	}
	r.mux.HandleFunc(pattern, handler)
	r.routes = append(r.routes, Route{Method: method, Pattern: routePattern})
}

// Routes returns the registered route table (a copy, in registration order).
func (r *Router) Routes() []Route {
	return slices.Clone(r.routes)
}

// ServeHTTP freezes the middleware chain on first use and dispatches the
// request through it.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.buildOnce.Do(r.build)
	r.handler.ServeHTTP(WrapResponseWriter(w), req)
}

func (r *Router) build() {
	r.frozen.Store(true)
	var h http.Handler = http.HandlerFunc(r.dispatch)
	for i := len(r.middleware) - 1; i >= 0; i-- {
		h = r.middleware[i](h)
	}
	r.handler = h
}

// dispatch routes one request, applying the trailing-slash redirect for
// slash-suffixed paths that match no pattern directly. Only that direction
// needs handling here: for the reverse (subtree pattern registered, request
// missing the trailing slash) ServeMux redirects natively. Non-slash paths —
// the hot path — go straight to the mux with a single route match.
func (r *Router) dispatch(w http.ResponseWriter, req *http.Request) {
	if strings.HasSuffix(req.URL.Path, "/") {
		if target, code, ok := r.trailingSlashRedirect(req); ok {
			http.Redirect(w, req, target, code) // #nosec G710 -- target is the request's own path minus one trailing slash; "//"-prefixed paths are refused, so it cannot leave the site
			return
		}
	}
	r.mux.ServeHTTP(w, req)
}

// trailingSlashRedirect reports whether the slash-suffixed request path is
// itself unrouted while the path without the trailing slash matches a route
// for the same method; if so it returns the redirect target (query string
// preserved), 301 for GET and 307 otherwise.
//
// The target is always a same-site, path-only URL: it is the request's own
// path minus one slash and never carries a scheme or host. Paths starting
// with "//" are refused because a Location of that shape would be read as
// scheme-relative (an off-site redirect) by browsers.
func (r *Router) trailingSlashRedirect(req *http.Request) (target string, code int, ok bool) {
	path := req.URL.Path
	if len(path) <= 1 || strings.HasPrefix(path, "//") {
		return "", 0, false
	}
	if _, pattern := r.mux.Handler(req); pattern != "" {
		return "", 0, false // routed as-is (e.g. a subtree wildcard)
	}

	altURL := &url.URL{Path: path[:len(path)-1], RawQuery: req.URL.RawQuery}
	probe := req.Clone(req.Context())
	probe.URL = altURL
	probe.RequestURI = altURL.RequestURI()
	if _, pattern := r.mux.Handler(probe); pattern == "" {
		return "", 0, false
	}

	code = http.StatusMovedPermanently
	if req.Method != http.MethodGet {
		code = http.StatusTemporaryRedirect
	}
	return altURL.String(), code, true
}
