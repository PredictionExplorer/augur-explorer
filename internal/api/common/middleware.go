package common

import (
	"crypto/subtle"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RequireAdminKey guards a mutating endpoint with a shared-secret header.
//
// The expected secret is taken from the first non-empty environment variable
// in envVars. The endpoint FAILS CLOSED: if none of the variables is set the
// route responds 503 instead of allowing anonymous access, so a missing
// deployment variable can never silently expose an admin operation.
func RequireAdminKey(header string, envVars ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var secret string
		for _, name := range envVars {
			if v := strings.TrimSpace(os.Getenv(name)); v != "" {
				secret = v
				break
			}
		}
		if secret == "" {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"status": 0,
				"error":  "endpoint disabled: no admin key configured (" + strings.Join(envVars, " or ") + ")",
			})
			return
		}
		provided := c.GetHeader(header)
		if subtle.ConstantTimeCompare([]byte(provided), []byte(secret)) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": 0,
				"error":  "invalid or missing " + header,
			})
			return
		}
		c.Next()
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
func RateLimit(rps float64, burst int) gin.HandlerFunc {
	l := newIPLimiter(rps, burst)
	return func(c *gin.Context) {
		if !l.allow(c.ClientIP()) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"status": 0,
				"error":  "rate limit exceeded, slow down",
			})
			return
		}
		c.Next()
	}
}
