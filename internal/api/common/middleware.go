package common

import (
	"crypto/subtle"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// CORS allows cross-origin browser requests (the frontends live on other
// origins). It adds the CORS headers to every response and answers OPTIONS
// preflight with 204 before routing, for any path.
func CORS() httpx.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := w.Header()
			h.Set("Access-Control-Allow-Origin", "*")
			h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			h.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Recovery converts handler panics into empty 500 responses instead of
// killing the connection (and, before Go recovers it, the process). Panics
// caused by the client going away (broken pipe / connection reset) produce
// no response at all. http.ErrAbortHandler is re-panicked so the server's
// own suppression applies. A nil logger falls back to slog.Default().
func Recovery(logger *slog.Logger) httpx.Middleware {
	if logger == nil {
		logger = slog.Default()
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				rec := recover()
				if rec == nil {
					return
				}
				if err, ok := rec.(error); ok && errors.Is(err, http.ErrAbortHandler) {
					panic(rec)
				}
				if isClientDisconnect(rec) {
					// The peer is gone; there is nobody to answer.
					logger.Warn("client disconnected mid-response",
						"method", r.Method, "path", r.URL.Path, "panic", rec)
					return
				}
				logger.Error("panic recovered",
					"method", r.Method, "path", r.URL.Path,
					"panic", rec, "stack", string(debug.Stack()))
				if rw, ok := w.(httpx.ResponseWriter); !ok || !rw.Written() {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// isClientDisconnect reports whether a recovered panic is a write failure
// caused by the client closing the connection.
func isClientDisconnect(rec any) bool {
	err, ok := rec.(error)
	if !ok {
		return false
	}
	var opErr *net.OpError
	if !errors.As(err, &opErr) {
		return false
	}
	msg := strings.ToLower(opErr.Error())
	return strings.Contains(msg, "broken pipe") || strings.Contains(msg, "connection reset by peer")
}

// AccessLog emits one structured log line per completed request (replaces
// the legacy framework logger). The route field is the matched pattern, so
// log cardinality stays bounded; path carries the concrete URL.
func AccessLog(logger *slog.Logger) httpx.Middleware {
	if logger == nil {
		logger = slog.Default()
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := httpx.WrapResponseWriter(w)
			next.ServeHTTP(rw, r)

			path := r.URL.Path
			if raw := r.URL.RawQuery; raw != "" {
				path += "?" + raw
			}
			logger.Info("request",
				"method", r.Method,
				"path", path,
				"route", httpx.PatternPath(r),
				"status", rw.Status(),
				"bytes", rw.Size(),
				"duration_ms", float64(time.Since(start).Microseconds())/1000.0,
				"ip", httpx.ClientIP(r),
			)
		})
	}
}

// AdminKey is one configured admin secret: Name is the environment variable
// the deployment sets it through (named in the fail-closed disabled message
// so operators know what to configure) and Value is the loaded secret, which
// may be empty.
type AdminKey struct {
	Name  string
	Value string
}

// RequireAdminKey guards a mutating endpoint with a shared-secret header.
//
// The expected secret is the first key with a non-empty Value (values are
// injected from the typed service configuration at construction). The
// endpoint FAILS CLOSED: if every key is empty the route responds 503
// instead of allowing anonymous access, so a missing deployment variable can
// never silently expose an admin operation.
func RequireAdminKey(header string, keys ...AdminKey) httpx.Middleware {
	var secret string
	names := make([]string, 0, len(keys))
	for _, k := range keys {
		names = append(names, k.Name)
		if secret == "" {
			secret = strings.TrimSpace(k.Value)
		}
	}
	disabledMsg := "endpoint disabled: no admin key configured (" + strings.Join(names, " or ") + ")"
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := httpx.NewContext(w, r)
			if secret == "" {
				c.JSON(http.StatusServiceUnavailable, httpx.H{
					"status": 0,
					"error":  disabledMsg,
				})
				return
			}
			provided := r.Header.Get(header)
			if subtle.ConstantTimeCompare([]byte(provided), []byte(secret)) != 1 {
				c.JSON(http.StatusUnauthorized, httpx.H{
					"status": 0,
					"error":  "invalid or missing " + header,
				})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// ipLimiter tracks one token bucket per client IP with lazy eviction.
type ipLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	rps      rate.Limit
	burst    int
}

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

const visitorTTL = 10 * time.Minute

func newIPLimiter(rps float64, burst int) *ipLimiter {
	l := &ipLimiter{
		visitors: make(map[string]*visitor),
		rps:      rate.Limit(rps),
		burst:    burst,
	}
	// Evict idle entries so the map cannot grow without bound.
	go func() {
		for range time.Tick(visitorTTL) {
			l.mu.Lock()
			for ip, v := range l.visitors {
				if time.Since(v.lastSeen) > visitorTTL {
					delete(l.visitors, ip)
				}
			}
			l.mu.Unlock()
		}
	}()
	return l
}

func (l *ipLimiter) allow(ip string) bool {
	l.mu.Lock()
	v, ok := l.visitors[ip]
	if !ok {
		v = &visitor{limiter: rate.NewLimiter(l.rps, l.burst)}
		l.visitors[ip] = v
	}
	v.lastSeen = time.Now()
	l.mu.Unlock()
	return v.limiter.Allow()
}

// RateLimit returns a per-client-IP token-bucket limiter middleware.
// Requests over the limit receive 429 Too Many Requests.
//
// Use a generous limit globally (this is a public read API) and a strict one
// on mutating or expensive endpoints.
func RateLimit(rps float64, burst int) httpx.Middleware {
	l := newIPLimiter(rps, burst)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !l.allow(httpx.ClientIP(r)) {
				httpx.NewContext(w, r).JSON(http.StatusTooManyRequests, httpx.H{
					"status": 0,
					"error":  "rate limit exceeded, slow down",
				})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
