// Bounded request processing: no API request may wait on its dependencies
// forever. The middleware puts one deadline on the request context, so every
// context-aware call a handler makes (PostgreSQL through pgx, contract reads
// through go-ethereum) fails with context.DeadlineExceeded once the budget is
// spent — and the failure renders as an explicit 503 timeout instead of the
// generic 500, in whichever envelope the route family speaks (the D21
// per-family precedent). Client disconnects (context.Canceled) keep their
// existing behavior: there is nobody left to answer.

package common

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// DefaultRequestDeadline is the fixed processing budget of one API request
// (D22). It bounds context-aware handler work — database queries and
// contract reads — not the response write to a slow client (that is the
// listener's concern). 30s is 60x the store's slow-query warning threshold;
// a request that needs more is a fault, not a workload.
const DefaultRequestDeadline = 30 * time.Second

// DeadlinePolicy selects the processing deadline for one request. A zero or
// negative duration exempts the request: its context gets no deadline (the
// FAQ proxy is the one exempt family — its LLM upstream legitimately exceeds
// the API budget and is bounded by the proxy's own HTTP client timeout).
type DeadlinePolicy func(*http.Request) time.Duration

// RequestDeadline bounds every request's processing time with a context
// deadline chosen by policy. When the deadline expires, the first
// context-aware call inside the handler fails and the handler renders its
// usual internal-error response; the middleware intercepts that 5xx and
// answers 503 instead — an RFC 9457 request-timeout problem under /api/v2/,
// the legacy {"status":0,"error":...} envelope everywhere else. onTimeout
// (nil-safe) observes every such interception; production counts it in
// rwcg_http_request_timeouts_total.
func RequestDeadline(policy DeadlinePolicy, onTimeout func(*http.Request)) httpx.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			d := policy(r)
			if d <= 0 {
				next.ServeHTTP(w, r)
				return
			}
			ctx, cancel := context.WithTimeout(r.Context(), d)
			defer cancel()
			bounded := r.WithContext(ctx)
			tw := &timeoutWriter{
				rw:        httpx.WrapResponseWriter(w),
				r:         bounded,
				ctx:       ctx,
				deadline:  d,
				onTimeout: onTimeout,
			}
			next.ServeHTTP(tw, bounded)
			// WithContext cloned the request, so the route pattern the mux
			// recorded during dispatch lives on the clone. Outer middleware
			// (access log, metrics) reads it from the original — copy it
			// back so their route labels keep working.
			if r.Pattern == "" {
				r.Pattern = bounded.Pattern
			}
		})
	}
}

// timeoutWriter converts an internal-error response produced after the
// request deadline expired into the explicit 503 timeout rendering. Handlers
// keep rendering their normal 5xx (they cannot tell a timed-out query from a
// broken one at every call site); the writer, which can, swaps the response
// at the header boundary. A response that already started (Written) is never
// touched, and non-5xx statuses pass through: a handler that finished its
// work just past the deadline still delivers its result.
type timeoutWriter struct {
	rw        httpx.ResponseWriter
	r         *http.Request
	ctx       context.Context
	deadline  time.Duration
	onTimeout func(*http.Request)
	hijacked  bool
}

// Header returns the response header map.
func (tw *timeoutWriter) Header() http.Header { return tw.rw.Header() }

// WriteHeader forwards the handler's status, replacing a post-deadline 5xx
// with the 503 timeout rendering.
func (tw *timeoutWriter) WriteHeader(status int) {
	if tw.hijacked {
		return
	}
	if status >= http.StatusInternalServerError && !tw.rw.Written() &&
		errors.Is(tw.ctx.Err(), context.DeadlineExceeded) {
		tw.hijacked = true
		if tw.onTimeout != nil {
			tw.onTimeout(tw.r)
		}
		respondTimeout(tw.rw, tw.r, tw.deadline)
		return
	}
	tw.rw.WriteHeader(status)
}

// Write forwards body bytes, swallowing the replaced response's body after a
// hijack (the handler keeps writing its 5xx body; nobody must receive it).
func (tw *timeoutWriter) Write(b []byte) (int, error) {
	if tw.hijacked {
		return len(b), nil
	}
	return tw.rw.Write(b)
}

// Status reports the status on the wire (503 after a hijack).
func (tw *timeoutWriter) Status() int { return tw.rw.Status() }

// Written reports whether the header section has been sent.
func (tw *timeoutWriter) Written() bool { return tw.rw.Written() }

// Size reports the body bytes that reached the underlying writer.
func (tw *timeoutWriter) Size() int { return tw.rw.Size() }

// Unwrap exposes the underlying writer for http.ResponseController.
func (tw *timeoutWriter) Unwrap() http.ResponseWriter { return tw.rw }

// RequestTimeoutDetail is the problem-detail sentence for a request that
// exceeded its processing deadline; kept here so a future handler-level
// rendering can never drift from the middleware's.
func RequestTimeoutDetail(d time.Duration) string {
	return "The request did not complete within the " + d.String() + " processing deadline."
}

// respondTimeout renders the 503 timeout in the request's family shape:
// problem+json under /api/v2/, the legacy envelope everywhere else. Headers
// a handler set before its replaced WriteHeader may remain; Content-Type is
// overwritten so the body and its declared type always agree.
func respondTimeout(w httpx.ResponseWriter, r *http.Request, d time.Duration) {
	if strings.HasPrefix(r.URL.Path, "/api/v2/") {
		body, err := json.Marshal(problemDocument{
			Detail:   RequestTimeoutDetail(d),
			Instance: r.URL.Path,
			Status:   http.StatusServiceUnavailable,
			Title:    "Request timeout",
			Type:     ProblemTypeBase + "request-timeout",
		})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write(append(body, '\n'))
		return
	}
	body, err := json.Marshal(httpx.H{
		"status": 0,
		"error":  "request timed out after " + d.String(),
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusServiceUnavailable)
	_, _ = w.Write(body)
}
