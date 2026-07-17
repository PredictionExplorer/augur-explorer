package common

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// ConditionalETag returns middleware that gives cacheable read responses a
// validator: successful GET/HEAD responses with hashable bodies gain a weak
// entity tag derived from the identity body, requests presenting a matching
// If-None-Match are answered 304 Not Modified with the body dropped, and
// Cache-Control defaults to "no-cache" (store, but revalidate) when no
// other layer chose a policy.
//
// The tag is weak (W/"…") because the same entity may travel gzip- or
// identity-encoded depending on negotiation; weak comparison treats those
// representations as equivalent, and RFC 9110 requires If-None-Match to use
// weak comparison anyway. It runs innermost so it hashes the body before
// compression and its 304 responses pass out through the compression layer
// untouched.
//
// Responses that already carry a validator (ETag or Last-Modified — file
// responses served by http.ServeFile own their conditional semantics),
// partial content, non-200 statuses and non-text types pass through
// unbuffered.
func ConditionalETag() httpx.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet && r.Method != http.MethodHead {
				next.ServeHTTP(w, r)
				return
			}
			ew := &etagWriter{rw: httpx.WrapResponseWriter(w), ifNoneMatch: r.Header.Get("If-None-Match")}
			next.ServeHTTP(ew, r)
			// Deliberately not deferred: a handler panic must leave the
			// response unwritten for the Recovery middleware outside.
			ew.finish()
		})
	}
}

// etagWriter buffers hashable 200 GET/HEAD bodies so the entity tag can be
// computed and compared before the header section is sent.
type etagWriter struct {
	rw          httpx.ResponseWriter
	ifNoneMatch string

	status      int
	handlerSent bool
	buffering   bool
	buf         []byte
	done        bool
}

// Header returns the response header map.
func (ew *etagWriter) Header() http.Header { return ew.rw.Header() }

// WriteHeader decides at the header boundary whether the response is
// hashable; everything else forwards immediately.
func (ew *etagWriter) WriteHeader(status int) {
	if ew.handlerSent {
		return
	}
	ew.handlerSent = true
	ew.status = status

	h := ew.rw.Header()
	if status == http.StatusOK &&
		h.Get("ETag") == "" &&
		h.Get("Last-Modified") == "" &&
		h.Get("Content-Range") == "" &&
		h.Get("Content-Encoding") == "" &&
		isCompressibleContentType(h.Get("Content-Type")) {
		ew.buffering = true
		return
	}
	ew.done = true
	ew.rw.WriteHeader(status)
}

// Write buffers hashable bodies and forwards everything else.
func (ew *etagWriter) Write(b []byte) (int, error) {
	if !ew.handlerSent {
		ew.WriteHeader(http.StatusOK)
	}
	if ew.buffering {
		ew.buf = append(ew.buf, b...)
		return len(b), nil
	}
	return ew.rw.Write(b)
}

// finish computes the validator for a buffered response, answers 304 when
// the client already holds the current entity, and otherwise releases the
// buffered body. Called by the middleware on the success path only. A
// writer that reaches here unfinished is always buffering: WriteHeader
// marked every other response done and forwarded it immediately.
func (ew *etagWriter) finish() {
	if ew.done {
		return
	}
	ew.done = true
	if !ew.handlerSent {
		return
	}

	h := ew.rw.Header()
	etag := weakETag(ew.buf)
	h.Set("ETag", etag)
	if h.Get("Cache-Control") == "" {
		// Store, but revalidate before reuse: correct-by-default for API
		// reads whose freshness lifetime is unknowable here. Layers that
		// know better (static assets) set their policy first and win.
		h.Set("Cache-Control", "no-cache")
	}

	if ifNoneMatchSatisfied(ew.ifNoneMatch, etag) {
		// 304 carries the validator and caching metadata but no
		// representation: Content-Type/Length describe the omitted body.
		h.Del("Content-Type")
		h.Del("Content-Length")
		ew.rw.WriteHeader(http.StatusNotModified)
		ew.buf = nil
		return
	}

	ew.rw.WriteHeader(http.StatusOK)
	if len(ew.buf) > 0 {
		_, _ = ew.rw.Write(ew.buf)
		ew.buf = nil
	}
}

// Status reports the status chosen by the handler (or 304 once finish
// downgraded the response; outer middleware reads wire status from the
// shared underlying writer after finish).
func (ew *etagWriter) Status() int {
	if s := ew.rw.Status(); s != 0 {
		return s
	}
	return ew.status
}

// Written reports whether the handler started the response, even while the
// body is still buffered.
func (ew *etagWriter) Written() bool { return ew.handlerSent }

// Size reports the body bytes that reached the underlying writer.
func (ew *etagWriter) Size() int { return ew.rw.Size() }

// Unwrap exposes the underlying writer for http.ResponseController.
func (ew *etagWriter) Unwrap() http.ResponseWriter { return ew.rw }

// weakETag derives the weak validator for an identity body: a truncated
// SHA-256 in the W/"…" form. Truncation to 128 bits keeps headers short
// while collisions stay out of reach for a per-URL validator.
func weakETag(body []byte) string {
	sum := sha256.Sum256(body)
	return `W/"` + hex.EncodeToString(sum[:16]) + `"`
}

// ifNoneMatchSatisfied reports whether an If-None-Match header value lists
// an entity tag weakly matching current, or is the "*" wildcard (which
// matches any existing representation). Malformed members are skipped, so
// garbage conservatively answers false and the full response is sent.
func ifNoneMatchSatisfied(header, current string) bool {
	header = strings.TrimSpace(header)
	if header == "" {
		return false
	}
	if header == "*" {
		return true
	}
	currentOpaque, ok := opaqueTag(current)
	if !ok {
		return false
	}
	for _, member := range splitETagList(header) {
		memberOpaque, ok := opaqueTag(member)
		if !ok {
			continue
		}
		// Weak comparison (RFC 9110 §8.8.3.2): weakness flags are ignored,
		// opaque tags compare byte-wise.
		if memberOpaque == currentOpaque {
			return true
		}
	}
	return false
}

// opaqueTag strips an optional weakness prefix and the surrounding quotes,
// returning the opaque tag characters. ok is false when the input is not a
// well-formed entity tag.
func opaqueTag(tag string) (string, bool) {
	tag = strings.TrimSpace(tag)
	tag = strings.TrimPrefix(tag, "W/")
	if len(tag) < 2 || tag[0] != '"' || tag[len(tag)-1] != '"' {
		return "", false
	}
	inner := tag[1 : len(tag)-1]
	if strings.Contains(inner, `"`) {
		return "", false
	}
	return inner, true
}

// splitETagList splits an If-None-Match value into entity-tag members.
// Commas inside quoted opaque tags do not split (etagc permits ","), so the
// scan tracks quoting rather than using strings.Split.
func splitETagList(header string) []string {
	var members []string
	start := 0
	inQuotes := false
	for i := range len(header) {
		switch header[i] {
		case '"':
			inQuotes = !inQuotes
		case ',':
			if !inQuotes {
				if m := strings.TrimSpace(header[start:i]); m != "" {
					members = append(members, m)
				}
				start = i + 1
			}
		}
	}
	if m := strings.TrimSpace(header[start:]); m != "" {
		members = append(members, m)
	}
	return members
}
