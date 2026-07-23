// Bounded request bodies: the process edge accepts no unbounded client
// input. The largest legitimate bodies on this API are small JSON documents
// (ranking votes, moderation commands, FAQ questions), so one fixed global
// cap protects every handler — including future ones — from memory
// exhaustion without per-route knobs (the D9 no-knob precedent).

package common

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// ProblemTypeBase is the URI prefix of every RFC 9457 problem type the API
// emits. internal/api/v2 derives its problem types from this constant, so
// middleware-level and handler-level problems can never drift apart.
const ProblemTypeBase = "https://api.cosmicsignature.com/problems/"

// MaxRequestBodyBytes is the fixed production request-body cap (1 MiB).
const MaxRequestBodyBytes int64 = 1 << 20

// MaxRequestBody bounds every request body at limit bytes. A request that
// declares a larger Content-Length is rejected immediately with 413 —
// RFC 9457 problem+json under /api/v2/, the legacy {"status":0,"error":...}
// envelope everywhere else — without reading a byte. Bodies without a
// declared length (chunked, or lying clients) are wrapped in
// http.MaxBytesReader, so the first read past the limit fails with
// *http.MaxBytesError and the consuming layer answers: the generated v2
// decoders map it to the same 413 problem, the FAQ proxy to the same legacy
// 413 envelope, and the frozen v1 handlers keep their 400 bind-error shape.
func MaxRequestBody(limit int64) httpx.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength > limit {
				respondBodyTooLarge(w, r, limit)
				return
			}
			r.Body = http.MaxBytesReader(w, r.Body, limit)
			next.ServeHTTP(w, r)
		})
	}
}

// RequestBodyTooLargeDetail is the problem-detail sentence for an oversized
// request body; internal/api/v2 uses it for read-time overflows so the
// middleware's pre-check and the decoder path render identical problems.
func RequestBodyTooLargeDetail(limit int64) string {
	return "The request body exceeds the " + strconv.FormatInt(limit, 10) + "-byte limit."
}

// RespondRequestBodyTooLarge answers 413 in the legacy error envelope. The
// FAQ proxy calls it when its body read trips the MaxRequestBody wrapper, so
// declared and undeclared oversized bodies get the same response.
func RespondRequestBodyTooLarge(c *httpx.Context, limit int64) {
	c.JSON(http.StatusRequestEntityTooLarge, httpx.H{
		"status": 0,
		"error":  "request body exceeds the " + strconv.FormatInt(limit, 10) + "-byte limit",
	})
}

// problemDocument mirrors the generated v2 Problem model's field order, so
// the middleware's 413 body is byte-identical to a handler-level problem.
type problemDocument struct {
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
	Status   int    `json:"status"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

func respondBodyTooLarge(w http.ResponseWriter, r *http.Request, limit int64) {
	if !strings.HasPrefix(r.URL.Path, "/api/v2/") {
		RespondRequestBodyTooLarge(httpx.NewContext(w, r), limit)
		return
	}
	body, err := json.Marshal(problemDocument{
		Detail:   RequestBodyTooLargeDetail(limit),
		Instance: r.URL.Path,
		Status:   http.StatusRequestEntityTooLarge,
		Title:    "Request body too large",
		Type:     ProblemTypeBase + "request-too-large",
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusRequestEntityTooLarge)
	_, _ = w.Write(append(body, '\n'))
}
