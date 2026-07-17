package common

import (
	"net/http"
	"strconv"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// DeprecationPolicy configures the DeprecationHeaders middleware.
type DeprecationPolicy struct {
	// Match selects the requests that receive the headers (typically a
	// path-prefix check). Required.
	Match func(*http.Request) bool

	// DeprecatedAt is the moment the surface became deprecated, emitted as
	// the RFC 9745 Deprecation header ("@" + Unix seconds). Required.
	DeprecatedAt time.Time

	// LinkURL, when set, is emitted as a Link header with the RFC 9745
	// "deprecation" relation, pointing clients at the migration guide.
	LinkURL string

	// SunsetAt, when non-zero, is emitted as the RFC 8594 Sunset header
	// (HTTP-date): the earliest moment the surface may be removed. Zero
	// omits the header — the deprecation is announced without a removal
	// date yet.
	SunsetAt time.Time
}

// DeprecationHeaders announces a deprecated API surface (RFC 9745): every
// matched response — success, error, 304 revalidation or router 404 under
// the matched prefix — carries a Deprecation header, an optional
// migration-guide Link and, once a removal date is decided, a Sunset
// header. Values are precomputed; the per-request cost is the match and
// two or three header writes. Panics at construction on an invalid policy
// so a bad route table fails at startup.
func DeprecationHeaders(p DeprecationPolicy) httpx.Middleware {
	if p.Match == nil {
		panic("common.DeprecationHeaders: Match is required")
	}
	if p.DeprecatedAt.IsZero() {
		panic("common.DeprecationHeaders: DeprecatedAt is required")
	}
	deprecation := "@" + strconv.FormatInt(p.DeprecatedAt.Unix(), 10)
	link := ""
	if p.LinkURL != "" {
		link = "<" + p.LinkURL + `>; rel="deprecation"; type="text/markdown"`
	}
	sunset := ""
	if !p.SunsetAt.IsZero() {
		sunset = p.SunsetAt.UTC().Format(http.TimeFormat)
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if p.Match(req) {
				h := w.Header()
				h.Set("Deprecation", deprecation)
				if link != "" {
					h.Set("Link", link)
				}
				if sunset != "" {
					h.Set("Sunset", sunset)
				}
			}
			next.ServeHTTP(w, req)
		})
	}
}
