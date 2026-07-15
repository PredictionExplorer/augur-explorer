package common

import (
	"compress/gzip"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// compressThreshold is the smallest body (in bytes) worth compressing.
// Below it the gzip header/trailer overhead and the CPU spent buy nothing:
// the response is written identity-encoded.
const compressThreshold = 1024

// gzipWriterPool recycles gzip writers (they allocate ~256 KiB of
// compression state each; see gzip.NewWriterLevel).
var gzipWriterPool = sync.Pool{
	New: func() any {
		// DefaultCompression (level 6): JSON API bodies are highly
		// redundant, so the extra ratio over BestSpeed is worth the CPU on
		// database-bound request latencies.
		w, _ := gzip.NewWriterLevel(nil, gzip.DefaultCompression)
		return w
	},
}

// Compress returns middleware that gzip-encodes responses for clients that
// negotiate it via Accept-Encoding. A response is compressed only when all
// of the following hold:
//
//   - the request negotiates gzip (q-values honored, "*" supported),
//   - the response status is 200 (error envelopes are small; 206/304 must
//     keep their exact representation),
//   - no Content-Encoding or Content-Range is already set,
//   - the Content-Type is compressible (JSON, text, XML — never images),
//   - the body reaches compressThreshold (smaller bodies pass through).
//
// Every response gains "Vary: Accept-Encoding" so shared caches never serve
// a cached representation with the wrong encoding. Compressed responses
// carry "Content-Encoding: gzip" and drop any handler-set Content-Length.
// The response writer keeps the httpx.ResponseWriter contract, so status
// and byte accounting in outer middleware (access log) see wire bytes.
func Compress() httpx.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cw := &compressWriter{
				rw:         httpx.WrapResponseWriter(w),
				negotiated: AcceptsGzip(r.Header.Get("Accept-Encoding")),
			}
			next.ServeHTTP(cw, r)
			// Deliberately not deferred: when the handler panics nothing is
			// flushed, so the Recovery middleware (outside) still sees an
			// unwritten response and can answer 500.
			cw.finish()
		})
	}
}

// addVaryAcceptEncoding appends "Accept-Encoding" to Vary unless already
// listed.
func addVaryAcceptEncoding(h http.Header) {
	for _, v := range h.Values("Vary") {
		for token := range strings.SplitSeq(v, ",") {
			if strings.EqualFold(strings.TrimSpace(token), "Accept-Encoding") {
				return
			}
		}
	}
	h.Add("Vary", "Accept-Encoding")
}

// AcceptsGzip reports whether an Accept-Encoding request header value
// negotiates gzip: an explicit "gzip" (or its "x-gzip" alias) with q > 0,
// or a "*" wildcard with q > 0 while gzip is not explicitly refused.
// Malformed elements are ignored, so garbage input conservatively answers
// false rather than sending an encoding the client may not understand.
func AcceptsGzip(header string) bool {
	wildcard := false
	for element := range strings.SplitSeq(header, ",") {
		coding, q, ok := parseAcceptEncodingElement(element)
		if !ok {
			continue
		}
		switch coding {
		case "gzip", "x-gzip":
			// Explicit mention is definitive either way.
			return q > 0
		case "*":
			wildcard = q > 0
		}
	}
	return wildcard
}

// parseAcceptEncodingElement splits one Accept-Encoding list element into
// its lowercase coding name and quality value. ok is false for empty or
// malformed elements.
func parseAcceptEncodingElement(element string) (coding string, q float64, ok bool) {
	coding, params, _ := strings.Cut(element, ";")
	coding = strings.ToLower(strings.TrimSpace(coding))
	if coding == "" || strings.ContainsAny(coding, " \t") {
		return "", 0, false
	}
	q = 1
	for param := range strings.SplitSeq(params, ";") {
		key, value, found := strings.Cut(param, "=")
		if !found || !strings.EqualFold(strings.TrimSpace(key), "q") {
			continue
		}
		parsed, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil || parsed < 0 || parsed > 1 {
			// A malformed quality refuses the element (conservative).
			return "", 0, false
		}
		q = parsed
	}
	return coding, q, true
}

// isCompressibleContentType reports whether a Content-Type header value
// names a text-like media type that benefits from gzip. Already-compressed
// formats (images, video, archives) and unknown/empty types answer false.
func isCompressibleContentType(contentType string) bool {
	mediaType, _, _ := strings.Cut(contentType, ";")
	mediaType = strings.ToLower(strings.TrimSpace(mediaType))
	if mediaType == "" {
		return false
	}
	if strings.HasPrefix(mediaType, "text/") {
		return true
	}
	switch mediaType {
	case "application/json", "application/javascript", "application/xml",
		"image/svg+xml":
		return true
	}
	return strings.HasSuffix(mediaType, "+json") || strings.HasSuffix(mediaType, "+xml")
}

// compressWriter defers the decision to compress until the response shape
// is known. Eligible 200 responses buffer until compressThreshold: crossing
// it switches to streaming gzip, finishing below it flushes the buffer
// identity-encoded. Ineligible responses pass straight through.
type compressWriter struct {
	rw         httpx.ResponseWriter
	negotiated bool // the request accepts gzip

	status      int  // status the handler chose (0 until WriteHeader)
	handlerSent bool // handler called WriteHeader or wrote body bytes

	deciding bool   // buffering an eligible response below the threshold
	buf      []byte // pending identity bytes while deciding
	gz       *gzip.Writer
	done     bool // finish ran (or the response passed through)
}

// Header returns the response header map.
func (cw *compressWriter) Header() http.Header { return cw.rw.Header() }

// WriteHeader records the handler's status and decides eligibility. The
// underlying WriteHeader is deferred while the compression decision is
// pending, because Content-Encoding must be set before the header goes out.
// Vary is added here — after the handler chose its headers — so it dedupes
// against a handler-set value.
func (cw *compressWriter) WriteHeader(status int) {
	if cw.handlerSent {
		return
	}
	cw.handlerSent = true
	cw.status = status

	h := cw.rw.Header()
	addVaryAcceptEncoding(h)
	if cw.negotiated &&
		status == http.StatusOK &&
		h.Get("Content-Encoding") == "" &&
		h.Get("Content-Range") == "" &&
		isCompressibleContentType(h.Get("Content-Type")) {
		cw.deciding = true
		return
	}
	cw.done = true
	cw.rw.WriteHeader(status)
}

// Write buffers, streams through gzip, or passes through, depending on the
// decision state.
func (cw *compressWriter) Write(b []byte) (int, error) {
	if !cw.handlerSent {
		cw.WriteHeader(http.StatusOK)
	}
	switch {
	case cw.gz != nil:
		return cw.gz.Write(b)
	case cw.deciding:
		cw.buf = append(cw.buf, b...)
		if len(cw.buf) >= compressThreshold {
			if err := cw.engage(); err != nil {
				return 0, err
			}
		}
		return len(b), nil
	default:
		return cw.rw.Write(b)
	}
}

// engage commits to gzip: headers are finalized, the deferred WriteHeader
// goes out and the buffered bytes stream through a pooled gzip writer.
func (cw *compressWriter) engage() error {
	h := cw.rw.Header()
	h.Set("Content-Encoding", "gzip")
	// The handler's length (if any) describes the identity body.
	h.Del("Content-Length")
	cw.deciding = false
	cw.rw.WriteHeader(cw.status)

	gz := gzipWriterPool.Get().(*gzip.Writer)
	gz.Reset(cw.rw)
	cw.gz = gz

	buffered := cw.buf
	cw.buf = nil
	_, err := gz.Write(buffered)
	return err
}

// finish completes the response on the handler's success path: a still-
// undecided buffer flushes identity-encoded, an engaged gzip stream closes
// (writing the trailer). Called by the middleware, never deferred.
func (cw *compressWriter) finish() {
	if cw.done {
		return
	}
	cw.done = true
	if cw.gz != nil {
		_ = cw.gz.Close()
		gzipWriterPool.Put(cw.gz)
		cw.gz = nil
		return
	}
	if !cw.handlerSent {
		// The handler wrote nothing at all; leave the implicit-200 behavior
		// to net/http exactly as it would be without this middleware.
		return
	}
	cw.rw.WriteHeader(cw.status)
	if len(cw.buf) > 0 {
		_, _ = cw.rw.Write(cw.buf)
		cw.buf = nil
	}
}

// Status reports the status the handler chose, before it reaches the wire.
func (cw *compressWriter) Status() int { return cw.status }

// Written reports whether the handler has started the response, even while
// the compression decision (and the underlying write) is still pending.
func (cw *compressWriter) Written() bool { return cw.handlerSent }

// Size reports the wire bytes forwarded to the underlying writer so far.
func (cw *compressWriter) Size() int { return cw.rw.Size() }

// Unwrap exposes the underlying writer for http.ResponseController.
func (cw *compressWriter) Unwrap() http.ResponseWriter { return cw.rw }
