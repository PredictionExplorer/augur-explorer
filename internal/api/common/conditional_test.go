package common

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// serveConditional runs one exchange through ConditionalETag wrapping handler.
func serveConditional(t *testing.T, method string, handler http.HandlerFunc, mutate func(*http.Request)) *httptest.ResponseRecorder {
	t.Helper()
	h := ConditionalETag()(handler)
	req := httptest.NewRequest(method, "/x", nil)
	if mutate != nil {
		mutate(req)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(httpx.WrapResponseWriter(w), req)
	return w
}

func TestConditionalETagAddsValidatorAndDefaultCachePolicy(t *testing.T) {
	t.Parallel()
	body := []byte(`{"status":1,"rounds":[1,2,3]}`)
	w := serveConditional(t, http.MethodGet, jsonHandler(body), nil)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	etag := w.Header().Get("ETag")
	if !strings.HasPrefix(etag, `W/"`) || !strings.HasSuffix(etag, `"`) {
		t.Fatalf("ETag = %q, want a weak validator", etag)
	}
	if got := w.Header().Get("Cache-Control"); got != "no-cache" {
		t.Fatalf("Cache-Control = %q, want the no-cache default", got)
	}
	if !bytes.Equal(w.Body.Bytes(), body) {
		t.Fatal("body must be released unchanged")
	}
}

func TestConditionalETagDeterministicPerBody(t *testing.T) {
	t.Parallel()
	a1 := serveConditional(t, http.MethodGet, jsonHandler([]byte(`{"a":1}`)), nil).Header().Get("ETag")
	a2 := serveConditional(t, http.MethodGet, jsonHandler([]byte(`{"a":1}`)), nil).Header().Get("ETag")
	b := serveConditional(t, http.MethodGet, jsonHandler([]byte(`{"a":2}`)), nil).Header().Get("ETag")
	if a1 != a2 {
		t.Fatalf("same body produced different tags: %q vs %q", a1, a2)
	}
	if a1 == b {
		t.Fatalf("different bodies share tag %q", a1)
	}
}

func TestConditionalETag304Flow(t *testing.T) {
	t.Parallel()
	body := jsonPayload(2048)
	first := serveConditional(t, http.MethodGet, jsonHandler(body), nil)
	etag := first.Header().Get("ETag")
	if etag == "" {
		t.Fatal("no ETag on the first response")
	}

	tests := []struct {
		name        string
		ifNoneMatch string
		want        int
	}{
		{"exact weak tag", etag, http.StatusNotModified},
		{"strong form of the same tag", strings.TrimPrefix(etag, "W/"), http.StatusNotModified},
		{"wildcard", "*", http.StatusNotModified},
		{"match in a list", `"other", ` + etag + `, "third"`, http.StatusNotModified},
		{"list without a match", `"aaa", W/"bbb"`, http.StatusOK},
		{"different tag", `W/"0000000000000000000000000000dead"`, http.StatusOK},
		{"malformed member skipped", `not-quoted, ` + etag, http.StatusNotModified},
		{"only malformed members", `not-quoted, also"broken`, http.StatusOK},
		{"empty header", "", http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := serveConditional(t, http.MethodGet, jsonHandler(body), func(r *http.Request) {
				if tt.ifNoneMatch != "" {
					r.Header.Set("If-None-Match", tt.ifNoneMatch)
				}
			})
			if w.Code != tt.want {
				t.Fatalf("status = %d, want %d", w.Code, tt.want)
			}
			if tt.want == http.StatusNotModified {
				if w.Body.Len() != 0 {
					t.Fatalf("304 carried %d body bytes", w.Body.Len())
				}
				if got := w.Header().Get("ETag"); got != etag {
					t.Fatalf("304 ETag = %q, want %q", got, etag)
				}
				if got := w.Header().Get("Cache-Control"); got != "no-cache" {
					t.Fatalf("304 Cache-Control = %q", got)
				}
				if got := w.Header().Get("Content-Type"); got != "" {
					t.Fatalf("304 must not describe an omitted body, Content-Type = %q", got)
				}
			} else if !bytes.Equal(w.Body.Bytes(), body) {
				t.Fatal("200 must carry the full body")
			}
		})
	}
}

func TestConditionalETagHEAD(t *testing.T) {
	t.Parallel()
	body := jsonPayload(2048)
	first := serveConditional(t, http.MethodHead, jsonHandler(body), nil)
	etag := first.Header().Get("ETag")
	if etag == "" {
		t.Fatal("HEAD response missing ETag")
	}
	w := serveConditional(t, http.MethodHead, jsonHandler(body), func(r *http.Request) {
		r.Header.Set("If-None-Match", etag)
	})
	if w.Code != http.StatusNotModified {
		t.Fatalf("conditional HEAD status = %d, want 304", w.Code)
	}
}

func TestConditionalETagSkips(t *testing.T) {
	t.Parallel()
	big := jsonPayload(2048)
	tests := []struct {
		name    string
		method  string
		handler http.HandlerFunc
	}{
		{"POST responses", http.MethodPost, jsonHandler(big)},
		{"non-200 status", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"status":0,"error":"nope"}`))
		}},
		{"handler-owned ETag", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("ETag", `"handler-owned"`)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(big)
		}},
		{"file responses with Last-Modified", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Last-Modified", "Mon, 13 Jul 2026 12:00:00 GMT")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(big)
		}},
		{"pre-encoded bodies", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Encoding", "gzip")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(big)
		}},
		{"partial content", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Range", "bytes 0-99/2048")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(big[:100])
		}},
		{"non-text content", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/png")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(big)
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := serveConditional(t, tt.method, tt.handler, func(r *http.Request) {
				r.Header.Set("If-None-Match", "*")
			})
			if w.Code == http.StatusNotModified {
				t.Fatal("response must not be downgraded to 304")
			}
			if tt.name == "handler-owned ETag" {
				if got := w.Header().Get("ETag"); got != `"handler-owned"` {
					t.Fatalf("handler ETag overwritten: %q", got)
				}
			} else if got := w.Header().Get("ETag"); got != "" {
				t.Fatalf("unexpected ETag %q", got)
			}
			if got := w.Header().Get("Cache-Control"); got != "" {
				t.Fatalf("Cache-Control %q added to a skipped response", got)
			}
		})
	}
}

func TestConditionalETagPreservesHandlerCachePolicy(t *testing.T) {
	t.Parallel()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"secret":1}`))
	}
	w := serveConditional(t, http.MethodGet, handler, nil)
	if got := w.Header().Get("Cache-Control"); got != "no-store" {
		t.Fatalf("Cache-Control = %q, want the handler's no-store preserved", got)
	}
	if w.Header().Get("ETag") == "" {
		t.Fatal("validator should still be added; only the cache policy is owned elsewhere")
	}
}

func TestConditionalETagImplicitWriteHeader(t *testing.T) {
	t.Parallel()
	// The handler writes without calling WriteHeader: the implicit 200 must
	// still be hashed and validated.
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"implicit":true}`))
	}
	first := serveConditional(t, http.MethodGet, handler, nil)
	etag := first.Header().Get("ETag")
	if first.Code != http.StatusOK || etag == "" {
		t.Fatalf("status=%d etag=%q", first.Code, etag)
	}
	w := serveConditional(t, http.MethodGet, handler, func(r *http.Request) {
		r.Header.Set("If-None-Match", etag)
	})
	if w.Code != http.StatusNotModified {
		t.Fatalf("revalidation = %d, want 304", w.Code)
	}
}

func TestConditionalETagDuplicateWriteHeader(t *testing.T) {
	t.Parallel()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.WriteHeader(http.StatusTeapot) // ignored: first write wins
		_, _ = w.Write([]byte(`{"ok":1}`))
	}
	w := serveConditional(t, http.MethodGet, handler, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want the first WriteHeader to win", w.Code)
	}
}

func TestConditionalETagSilentHandler(t *testing.T) {
	t.Parallel()
	// A handler that writes nothing leaves the implicit-200 behavior to
	// net/http: no validator, no cache policy, nothing forwarded early.
	handler := func(w http.ResponseWriter, r *http.Request) {}
	w := serveConditional(t, http.MethodGet, handler, nil)
	if w.Code != http.StatusOK || w.Body.Len() != 0 {
		t.Fatalf("status=%d bodyBytes=%d", w.Code, w.Body.Len())
	}
	if w.Header().Get("ETag") != "" || w.Header().Get("Cache-Control") != "" {
		t.Fatal("silent responses must not gain cache metadata")
	}
}

func TestConditionalETagPanicLeavesResponseUnwritten(t *testing.T) {
	t.Parallel()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"partial":`))
		panic("boom")
	})
	h := Recovery(logger)(ConditionalETag()(handler))

	w := httptest.NewRecorder()
	h.ServeHTTP(httpx.WrapResponseWriter(w), httptest.NewRequest(http.MethodGet, "/x", nil))

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500 from Recovery", w.Code)
	}
	if w.Body.Len() != 0 {
		t.Fatalf("panic leaked %d buffered bytes", w.Body.Len())
	}
	if w.Header().Get("ETag") != "" {
		t.Fatal("no validator may be emitted for a failed response")
	}
}

func TestConditionalETagWriterContract(t *testing.T) {
	t.Parallel()
	under := httpx.WrapResponseWriter(httptest.NewRecorder())
	ew := &etagWriter{rw: under}

	if ew.Written() || ew.Status() != 0 {
		t.Fatal("fresh writer must be unwritten")
	}
	ew.Header().Set("Content-Type", "application/json")
	ew.WriteHeader(http.StatusOK)
	if !ew.Written() || ew.Status() != http.StatusOK {
		t.Fatalf("after WriteHeader: written=%v status=%d", ew.Written(), ew.Status())
	}
	if under.Written() {
		t.Fatal("underlying write must be deferred while buffering")
	}
	if _, err := ew.Write([]byte("body")); err != nil {
		t.Fatal(err)
	}
	ew.finish()
	if under.Status() != http.StatusOK || under.Size() != 4 {
		t.Fatalf("underlying status=%d size=%d", under.Status(), under.Size())
	}
	if ew.Size() != 4 {
		t.Fatalf("Size() = %d, want released body bytes", ew.Size())
	}
	if ew.Unwrap() != under {
		t.Fatal("Unwrap must return the underlying writer")
	}
}

// TestCompressedConditionalStack proves the production composition: the
// validator hashes the identity body (stable across encodings), 304s pass
// out through the compression layer, and full responses still compress.
func TestCompressedConditionalStack(t *testing.T) {
	t.Parallel()
	body := jsonPayload(8 * 1024)
	h := Compress()(ConditionalETag()(jsonHandler(body)))

	do := func(acceptEncoding, ifNoneMatch string) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, "/x", nil)
		if acceptEncoding != "" {
			req.Header.Set("Accept-Encoding", acceptEncoding)
		}
		if ifNoneMatch != "" {
			req.Header.Set("If-None-Match", ifNoneMatch)
		}
		w := httptest.NewRecorder()
		h.ServeHTTP(httpx.WrapResponseWriter(w), req)
		return w
	}

	gzipResp := do("gzip", "")
	identityResp := do("", "")

	if gzipResp.Header().Get("Content-Encoding") != "gzip" {
		t.Fatal("negotiated response must be compressed")
	}
	if !bytes.Equal(gunzip(t, gzipResp.Body.Bytes()), body) {
		t.Fatal("compressed body must decompress to the identity body")
	}
	if identityResp.Header().Get("Content-Encoding") != "" {
		t.Fatal("identity response must stay identity")
	}

	gzTag := gzipResp.Header().Get("ETag")
	idTag := identityResp.Header().Get("ETag")
	if gzTag == "" || gzTag != idTag {
		t.Fatalf("validator must not depend on the encoding: %q vs %q", gzTag, idTag)
	}

	// A client that cached the gzip representation revalidates: 304, no
	// body, no Content-Encoding (nothing was encoded).
	notModified := do("gzip", gzTag)
	if notModified.Code != http.StatusNotModified {
		t.Fatalf("revalidation status = %d, want 304", notModified.Code)
	}
	if notModified.Body.Len() != 0 {
		t.Fatal("304 must not carry a body")
	}
	if notModified.Header().Get("Content-Encoding") != "" {
		t.Fatal("304 must not claim an encoding")
	}
	if got := notModified.Header().Get("Vary"); got != "Accept-Encoding" {
		t.Fatalf("304 Vary = %q; caches need the negotiation dimension", got)
	}
}

func TestIfNoneMatchSatisfied(t *testing.T) {
	t.Parallel()
	const current = `W/"abc123"`
	tests := []struct {
		header string
		want   bool
	}{
		{`W/"abc123"`, true},
		{`"abc123"`, true},
		{`*`, true},
		{`W/"abc123", "zzz"`, true},
		{`"zzz", W/"abc123"`, true},
		{`"zzz"`, false},
		{``, false},
		{`   `, false},
		{`W/`, false},
		{`W/""`, false},
		{`"tag,with,commas", W/"abc123"`, true},
		{`W/"unterminated`, false},
		{`abc123`, false},
	}
	for _, tt := range tests {
		t.Run(tt.header, func(t *testing.T) {
			t.Parallel()
			if got := ifNoneMatchSatisfied(tt.header, current); got != tt.want {
				t.Fatalf("ifNoneMatchSatisfied(%q) = %v, want %v", tt.header, got, tt.want)
			}
		})
	}

	t.Run("malformed current tag never matches a list", func(t *testing.T) {
		t.Parallel()
		// Defensive: current is always our own weakETag output, but a
		// malformed value must fail closed (full response) rather than 304.
		if ifNoneMatchSatisfied(`"abc123"`, "not-a-tag") {
			t.Fatal("malformed current tag matched")
		}
		// The wildcard still matches: it asserts existence, not equality.
		if !ifNoneMatchSatisfied(`*`, "not-a-tag") {
			t.Fatal("wildcard must match any current representation")
		}
	})
}

func TestOpaqueTagRejectsEmbeddedQuotes(t *testing.T) {
	t.Parallel()
	if _, ok := opaqueTag(`"a"b"`); ok {
		t.Fatal(`opaqueTag("a"b") must be malformed`)
	}
	if got, ok := opaqueTag(` W/"abc" `); !ok || got != "abc" {
		t.Fatalf("opaqueTag with padding = %q, %v", got, ok)
	}
}

func TestIfNoneMatchEmptyOpaqueTag(t *testing.T) {
	t.Parallel()
	// An empty opaque tag is well-formed (`""`); it must match only itself.
	if !ifNoneMatchSatisfied(`""`, `W/""`) {
		t.Fatal("empty tags should weak-compare equal")
	}
	if ifNoneMatchSatisfied(`""`, `W/"abc"`) {
		t.Fatal("empty tag must not match a non-empty one")
	}
}

func TestSplitETagListQuotedCommas(t *testing.T) {
	t.Parallel()
	got := splitETagList(`"a,b", W/"c", "d"`)
	want := []string{`"a,b"`, `W/"c"`, `"d"`}
	if len(got) != len(want) {
		t.Fatalf("got %d members %v, want %v", len(got), got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("member %d = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestWeakETagFormat(t *testing.T) {
	t.Parallel()
	tag := weakETag([]byte("x"))
	if !strings.HasPrefix(tag, `W/"`) || !strings.HasSuffix(tag, `"`) {
		t.Fatalf("weakETag = %q", tag)
	}
	inner := strings.TrimSuffix(strings.TrimPrefix(tag, `W/"`), `"`)
	if len(inner) != 32 {
		t.Fatalf("opaque tag %q length = %d, want 32 hex chars", inner, len(inner))
	}
	if tag == weakETag([]byte("y")) {
		t.Fatal("distinct bodies must hash differently")
	}
}

// FuzzIfNoneMatch drives arbitrary If-None-Match header values and bodies
// through the real middleware: whatever the client sends, the response must
// be either the complete 200 representation or an empty 304 carrying the
// current validator — and the validator must always revalidate against
// itself.
func FuzzIfNoneMatch(f *testing.F) {
	f.Add("", "body")
	f.Add("*", `{"a":1}`)
	f.Add(`W/"6b"`, "x")
	f.Add(`"a,b", W/"c"`, "commas")
	f.Add(`W/"unterminated`, "y")
	f.Add(strings.Repeat(`"t",`, 50), "list")
	f.Add("\x00\"\x01", "control")
	f.Fuzz(func(t *testing.T, ifNoneMatch, bodyStr string) {
		body := []byte(bodyStr)
		current := weakETag(body)

		// Pure-function invariants.
		if got, again := ifNoneMatchSatisfied(ifNoneMatch, current), ifNoneMatchSatisfied(ifNoneMatch, current); got != again {
			t.Fatalf("nondeterministic for %q", ifNoneMatch)
		}
		if !ifNoneMatchSatisfied(current, current) {
			t.Fatalf("validator %q does not match itself", current)
		}
		if !ifNoneMatchSatisfied(current+", "+ifNoneMatch, current) {
			t.Fatalf("leading exact member lost in %q", current+", "+ifNoneMatch)
		}

		// End-to-end: the middleware may only ever answer the full 200
		// representation or an empty 304 with the current validator.
		h := ConditionalETag()(jsonHandler(body))
		req := httptest.NewRequest(http.MethodGet, "/x", nil)
		req.Header.Set("If-None-Match", ifNoneMatch)
		w := httptest.NewRecorder()
		h.ServeHTTP(httpx.WrapResponseWriter(w), req)

		switch w.Code {
		case http.StatusOK:
			if !bytes.Equal(w.Body.Bytes(), body) {
				t.Fatalf("200 body corrupted for If-None-Match %q", ifNoneMatch)
			}
		case http.StatusNotModified:
			if w.Body.Len() != 0 {
				t.Fatalf("304 carried a body for If-None-Match %q", ifNoneMatch)
			}
		default:
			t.Fatalf("status %d for If-None-Match %q", w.Code, ifNoneMatch)
		}
		if got := w.Header().Get("ETag"); got != current {
			t.Fatalf("response ETag %q, want %q", got, current)
		}
	})
}
