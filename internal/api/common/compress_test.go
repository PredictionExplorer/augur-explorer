package common

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// jsonPayload builds a deterministic JSON document of at least n bytes.
func jsonPayload(n int) []byte {
	var b bytes.Buffer
	b.WriteString(`{"rows":[`)
	for i := 0; b.Len() < n; i++ {
		if i > 0 {
			b.WriteByte(',')
		}
		fmt.Fprintf(&b, `{"id":%d,"addr":"0x00000000000000000000000000000000000000%02x","amountWei":"%d000000000000000000"}`, i, i%256, i)
	}
	b.WriteString(`]}`)
	return b.Bytes()
}

// serveCompressed runs one exchange through Compress wrapping handler.
func serveCompressed(t *testing.T, handler http.HandlerFunc, mutate func(*http.Request)) *httptest.ResponseRecorder {
	t.Helper()
	h := Compress()(handler)
	req := httptest.NewRequest(http.MethodGet, "/x", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	if mutate != nil {
		mutate(req)
	}
	w := httptest.NewRecorder()
	// The router wraps every writer before middleware; mirror that.
	h.ServeHTTP(httpx.WrapResponseWriter(w), req)
	return w
}

// gunzip decompresses a gzip body or fails the test.
func gunzip(t *testing.T, data []byte) []byte {
	t.Helper()
	zr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("gzip.NewReader: %v", err)
	}
	defer func() { _ = zr.Close() }() // checksum already verified by ReadAll
	out, err := io.ReadAll(zr)
	if err != nil {
		t.Fatalf("decompressing: %v", err)
	}
	return out
}

func jsonHandler(body []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	}
}

func TestCompressLargeJSONRoundTrip(t *testing.T) {
	t.Parallel()
	body := jsonPayload(64 * 1024)
	w := serveCompressed(t, jsonHandler(body), nil)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if got := w.Header().Get("Content-Encoding"); got != "gzip" {
		t.Fatalf("Content-Encoding = %q, want gzip", got)
	}
	if got := w.Header().Get("Vary"); got != "Accept-Encoding" {
		t.Fatalf("Vary = %q", got)
	}
	if cl := w.Header().Get("Content-Length"); cl != "" {
		t.Fatalf("Content-Length must be dropped on compressed responses, got %q", cl)
	}
	if w.Body.Len() >= len(body) {
		t.Fatalf("compressed body (%d) not smaller than identity (%d)", w.Body.Len(), len(body))
	}
	if got := gunzip(t, w.Body.Bytes()); !bytes.Equal(got, body) {
		t.Fatalf("decompressed body differs: %d vs %d bytes", len(got), len(body))
	}
}

func TestCompressNegotiation(t *testing.T) {
	t.Parallel()
	body := jsonPayload(8 * 1024)
	tests := []struct {
		header string
		want   bool
	}{
		{"gzip", true},
		{"GZIP", true},
		{"x-gzip", true},
		{"gzip, deflate, br", true},
		{"deflate, gzip;q=0.5", true},
		{"gzip;q=0", false},
		{"gzip;q=0.000", false},
		{"gzip; q=0", false},
		{"", false},
		{"identity", false},
		{"deflate, br", false},
		{"*", true},
		{"*;q=0", false},
		{"*;q=0.5, gzip;q=0", false},
		{"gzip;q=0, *;q=1", false},
		{"br;q=1, *;q=0.1", true},
		{"gzip;q=garbage", false},
		{"gzip;q=1.5", false},
		{";q=1", false},
		{"gz ip", false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("header=%q", tt.header), func(t *testing.T) {
			t.Parallel()
			w := serveCompressed(t, jsonHandler(body), func(r *http.Request) {
				r.Header.Del("Accept-Encoding")
				if tt.header != "" {
					r.Header.Set("Accept-Encoding", tt.header)
				}
			})
			gotGzip := w.Header().Get("Content-Encoding") == "gzip"
			if gotGzip != tt.want {
				t.Fatalf("compressed = %v, want %v", gotGzip, tt.want)
			}
			if !gotGzip && !bytes.Equal(w.Body.Bytes(), body) {
				t.Fatal("identity body must pass through unchanged")
			}
			// Vary advertises encoding-dependence to caches either way.
			if got := w.Header().Get("Vary"); got != "Accept-Encoding" {
				t.Fatalf("Vary = %q", got)
			}
		})
	}
}

func TestCompressThresholdBoundary(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		size int
		want bool
	}{
		{"below threshold stays identity", compressThreshold - 1, false},
		{"at threshold compresses", compressThreshold, true},
		{"empty body stays identity", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			body := bytes.Repeat([]byte("a"), tt.size)
			w := serveCompressed(t, jsonHandler(body), nil)
			if got := w.Header().Get("Content-Encoding") == "gzip"; got != tt.want {
				t.Fatalf("compressed = %v, want %v (size %d)", got, tt.want, tt.size)
			}
			if !tt.want && !bytes.Equal(w.Body.Bytes(), body) {
				t.Fatal("identity body must pass through unchanged")
			}
			if tt.want && !bytes.Equal(gunzip(t, w.Body.Bytes()), body) {
				t.Fatal("compressed body must decompress to the original")
			}
		})
	}
}

func TestCompressManySmallWritesCrossThreshold(t *testing.T) {
	t.Parallel()
	chunk := []byte(strings.Repeat("x", 100))
	const chunks = 64
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		for range chunks {
			_, _ = w.Write(chunk) // implicit 200 on the first write
		}
	}
	w := serveCompressed(t, handler, nil)
	if w.Header().Get("Content-Encoding") != "gzip" {
		t.Fatal("expected compression to engage mid-stream")
	}
	want := bytes.Repeat(chunk, chunks)
	if got := gunzip(t, w.Body.Bytes()); !bytes.Equal(got, want) {
		t.Fatalf("decompressed %d bytes, want %d", len(got), len(want))
	}
}

func TestCompressContentTypeGate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		contentType string
		want        bool
	}{
		{"application/json", true},
		{"application/json; charset=utf-8", true},
		{"application/problem+json", true},
		{"text/plain; charset=utf-8", true},
		{"text/html", true},
		{"application/xml", true},
		{"application/rss+xml", true},
		{"application/javascript", true},
		{"image/svg+xml", true},
		{"image/png", false},
		{"video/mp4", false},
		{"application/octet-stream", false},
		{"application/zip", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.contentType, func(t *testing.T) {
			t.Parallel()
			body := jsonPayload(8 * 1024)
			handler := func(w http.ResponseWriter, r *http.Request) {
				if tt.contentType != "" {
					w.Header().Set("Content-Type", tt.contentType)
				}
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(body)
			}
			w := serveCompressed(t, handler, nil)
			if got := w.Header().Get("Content-Encoding") == "gzip"; got != tt.want {
				t.Fatalf("compressed = %v, want %v for %q", got, tt.want, tt.contentType)
			}
		})
	}
}

func TestCompressSkipsNon200(t *testing.T) {
	t.Parallel()
	for _, status := range []int{http.StatusCreated, http.StatusNoContent, http.StatusMovedPermanently, http.StatusBadRequest, http.StatusInternalServerError} {
		t.Run(fmt.Sprintf("status_%d", status), func(t *testing.T) {
			t.Parallel()
			body := jsonPayload(8 * 1024)
			handler := func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				if status != http.StatusNoContent {
					_, _ = w.Write(body)
				}
			}
			w := serveCompressed(t, handler, nil)
			if w.Code != status {
				t.Fatalf("status = %d, want %d", w.Code, status)
			}
			if w.Header().Get("Content-Encoding") != "" {
				t.Fatalf("status %d must not be compressed", status)
			}
		})
	}
}

func TestCompressRespectsExistingEncodingHeaders(t *testing.T) {
	t.Parallel()
	t.Run("content-encoding already set", func(t *testing.T) {
		t.Parallel()
		body := jsonPayload(8 * 1024)
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Encoding", "br")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(body)
		}
		w := serveCompressed(t, handler, nil)
		if got := w.Header().Get("Content-Encoding"); got != "br" {
			t.Fatalf("Content-Encoding = %q, want the handler's br untouched", got)
		}
		if !bytes.Equal(w.Body.Bytes(), body) {
			t.Fatal("pre-encoded body must pass through unchanged")
		}
	})
	t.Run("content-range present", func(t *testing.T) {
		t.Parallel()
		body := jsonPayload(8 * 1024)
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Range", "bytes 0-8191/65536")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(body)
		}
		w := serveCompressed(t, handler, nil)
		if w.Header().Get("Content-Encoding") != "" {
			t.Fatal("partial content must not be compressed")
		}
	})
}

func TestCompressVaryDeduplicated(t *testing.T) {
	t.Parallel()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin, Accept-Encoding")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("ok"))
	}
	w := serveCompressed(t, handler, nil)
	if got := w.Header().Values("Vary"); len(got) != 1 || got[0] != "Origin, Accept-Encoding" {
		t.Fatalf("Vary = %v, want the handler's value preserved without duplication", got)
	}
}

func TestCompressContentLengthDroppedOnlyWhenCompressing(t *testing.T) {
	t.Parallel()
	body := jsonPayload(8 * 1024)
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", fmt.Sprint(len(body)))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	}
	if w := serveCompressed(t, handler, nil); w.Header().Get("Content-Length") != "" {
		t.Fatal("Content-Length must be dropped when the body is re-encoded")
	}
	identity := serveCompressed(t, handler, func(r *http.Request) { r.Header.Del("Accept-Encoding") })
	if got := identity.Header().Get("Content-Length"); got != fmt.Sprint(len(body)) {
		t.Fatalf("identity Content-Length = %q, want %d preserved", got, len(body))
	}
}

func TestCompressHandlerWritesNothing(t *testing.T) {
	t.Parallel()
	handler := func(w http.ResponseWriter, r *http.Request) {}
	w := serveCompressed(t, handler, nil)
	if w.Code != http.StatusOK || w.Body.Len() != 0 {
		t.Fatalf("silent handler: status %d, %d body bytes", w.Code, w.Body.Len())
	}
	if w.Header().Get("Content-Encoding") != "" {
		t.Fatal("empty response must not be compressed")
	}
}

func TestCompressHeaderOnlyResponse(t *testing.T) {
	t.Parallel()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // eligible, but no body follows
	}
	w := serveCompressed(t, handler, nil)
	if w.Code != http.StatusOK || w.Body.Len() != 0 {
		t.Fatalf("status %d, %d body bytes", w.Code, w.Body.Len())
	}
	if w.Header().Get("Content-Encoding") != "" {
		t.Fatal("empty response must not be compressed")
	}
}

// TestCompressPanicLeavesResponseUnwritten proves the Recovery contract: a
// panic while the compression decision is pending must not flush anything,
// so Recovery (outside) can still answer a clean 500.
func TestCompressPanicLeavesResponseUnwritten(t *testing.T) {
	t.Parallel()
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"partial":`)) // below threshold: still buffered
		panic("boom")
	})
	h := Recovery(logger)(Compress()(handler))

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/x", nil)
	req.Header.Set("Accept-Encoding", "gzip")
	h.ServeHTTP(httpx.WrapResponseWriter(w), req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500 from Recovery", w.Code)
	}
	if w.Body.Len() != 0 {
		t.Fatalf("panic response leaked %d buffered bytes: %q", w.Body.Len(), w.Body.String())
	}
}

// TestCompressWriterContract pins the httpx.ResponseWriter semantics outer
// middleware relies on: status visible before the wire write, handler-
// perspective Written, wire-byte Size after completion.
func TestCompressWriterContract(t *testing.T) {
	t.Parallel()
	under := httpx.WrapResponseWriter(httptest.NewRecorder())
	cw := &compressWriter{rw: under, negotiated: true}

	if cw.Written() || cw.Status() != 0 {
		t.Fatal("fresh writer must be unwritten")
	}
	cw.Header().Set("Content-Type", "application/json")
	cw.WriteHeader(http.StatusOK)
	if !cw.Written() || cw.Status() != http.StatusOK {
		t.Fatalf("after WriteHeader: written=%v status=%d", cw.Written(), cw.Status())
	}
	if under.Written() {
		t.Fatal("the underlying header must still be deferred while deciding")
	}
	if _, err := cw.Write(jsonPayload(4 * 1024)); err != nil {
		t.Fatal(err)
	}
	cw.finish()
	if under.Status() != http.StatusOK {
		t.Fatalf("underlying status = %d", under.Status())
	}
	if cw.Size() == 0 || cw.Size() != under.Size() {
		t.Fatalf("Size() = %d, underlying %d; want equal wire bytes", cw.Size(), under.Size())
	}
	if got := cw.Unwrap(); got != under {
		t.Fatal("Unwrap must return the underlying writer")
	}
}

// TestCompressDuplicateWriteHeader pins first-write-wins semantics: a
// second WriteHeader call must not disturb the committed status.
func TestCompressDuplicateWriteHeader(t *testing.T) {
	t.Parallel()
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.WriteHeader(http.StatusInternalServerError) // ignored
		_, _ = w.Write([]byte(`{"ok":true}`))
	}
	w := serveCompressed(t, handler, nil)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want the first WriteHeader to win", w.Code)
	}
}

// failingWriter refuses every body write, simulating a dead client
// connection under the compressor.
type failingWriter struct {
	httpx.ResponseWriter
}

func (f *failingWriter) Write([]byte) (int, error) {
	return 0, fmt.Errorf("connection reset by test")
}

// TestCompressUnderlyingWriteFailure proves an engage-time write failure
// surfaces to the handler instead of being swallowed.
func TestCompressUnderlyingWriteFailure(t *testing.T) {
	t.Parallel()
	under := &failingWriter{ResponseWriter: httpx.WrapResponseWriter(httptest.NewRecorder())}
	cw := &compressWriter{rw: under, negotiated: true}
	cw.Header().Set("Content-Type", "application/json")
	cw.WriteHeader(http.StatusOK)

	// Random-ish payload large enough that deflate must flush its window
	// to the (failing) underlying writer during engage.
	payload := jsonPayload(128 * 1024)
	if _, err := cw.Write(payload); err == nil {
		t.Fatal("write error was swallowed by the compression buffer")
	}
	cw.finish()
}

// TestCompressConcurrentRequests exercises the gzip writer pool under the
// race detector.
func TestCompressConcurrentRequests(t *testing.T) {
	t.Parallel()
	body := jsonPayload(16 * 1024)
	h := Compress()(jsonHandler(body))
	var wg sync.WaitGroup
	for range 16 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 8 {
				w := httptest.NewRecorder()
				req := httptest.NewRequest(http.MethodGet, "/x", nil)
				req.Header.Set("Accept-Encoding", "gzip")
				h.ServeHTTP(httpx.WrapResponseWriter(w), req)
				if !bytes.Equal(gunzip(t, w.Body.Bytes()), body) {
					t.Error("corrupted concurrent response")
					return
				}
			}
		}()
	}
	wg.Wait()
}

func TestAcceptsGzipNeverPanicsOnControlCharacters(t *testing.T) {
	t.Parallel()
	for _, header := range []string{",,,", ";;;", "q=", "\x00gzip", "gzip\r\n", strings.Repeat(",", 1000)} {
		_ = AcceptsGzip(header) // must not panic
	}
}

// FuzzAcceptEncodingGzip hunts for panics and semantic inconsistencies in
// the Accept-Encoding negotiation: the decision must be deterministic,
// case-insensitive, whitespace-tolerant, and an explicit leading "gzip"
// or "gzip;q=0" element must always dominate whatever follows.
func FuzzAcceptEncodingGzip(f *testing.F) {
	for _, seed := range []string{
		"", "gzip", "GZIP;q=1", "gzip;q=0", "*", "*;q=0.5, gzip;q=0",
		"deflate, br;q=0.9, gzip;q=0.1", "identity;q=0, *;q=0", "x-gzip",
		"gzip;q=0.", ";;q=1,,", "gzip ;q=1", "\tgzip\t", "q=1;gzip",
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, header string) {
		got := AcceptsGzip(header)

		if again := AcceptsGzip(header); again != got {
			t.Fatalf("nondeterministic: %v then %v for %q", got, again, header)
		}
		if upper := AcceptsGzip(strings.ToUpper(header)); upper != got {
			t.Fatalf("case-sensitive: %q=%v but uppercased=%v", header, got, upper)
		}
		spaced := strings.ReplaceAll(header, ",", " , ")
		if s := AcceptsGzip(spaced); s != got {
			t.Fatalf("whitespace-sensitive: %q=%v but %q=%v", header, got, spaced, s)
		}
		// The first explicit gzip element is definitive: nothing after it
		// may change the decision.
		if !AcceptsGzip("gzip, " + header) {
			t.Fatalf("leading gzip element ignored with suffix %q", header)
		}
		if AcceptsGzip("gzip;q=0, " + header) {
			t.Fatalf("leading gzip refusal ignored with suffix %q", header)
		}
	})
}
