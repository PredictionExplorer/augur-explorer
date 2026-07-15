//go:build integration

package apitest

// End-to-end coverage for the response-edge middleware through the real
// production chain (routes.New): gzip compression negotiated per request,
// weak-ETag conditional requests, the default Cache-Control policy and the
// /version build-identity endpoint. The parity goldens stay byte-identical
// because the golden harness never negotiates gzip or sends validators;
// these tests prove the negotiated representations are equivalent.

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

// fatV1Route is a large deterministic v1 response (well above the 1 KiB
// compression threshold); fatV2Route is its v2 counterpart.
const (
	fatV1Route = "/api/cosmicgame/user/info/0x2100000000000000000000000000000000000021"
	fatV2Route = "/api/v2/cosmicgame/rounds/0/claims"
)

func gunzipBody(t *testing.T, data []byte) []byte {
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

// TestResponseCompression proves that a gzip-negotiating client receives a
// compressed representation that decompresses byte-identically to the
// identity representation the goldens pin — on both API generations.
func TestResponseCompression(t *testing.T) {
	h := server(t)

	for _, route := range []string{fatV1Route, fatV2Route} {
		t.Run(route, func(t *testing.T) {
			identity := h.get(t, route)
			if identity.Code != http.StatusOK {
				t.Fatalf("identity status = %d", identity.Code)
			}
			if identity.Header().Get("Content-Encoding") != "" {
				t.Fatal("response compressed without Accept-Encoding")
			}
			if identity.Body.Len() < 1024 {
				t.Fatalf("test route too small to exercise compression: %d bytes", identity.Body.Len())
			}

			compressed := h.do(t, request{path: route, headers: map[string]string{
				"Accept-Encoding": "gzip, deflate, br",
			}})
			if compressed.Code != http.StatusOK {
				t.Fatalf("negotiated status = %d", compressed.Code)
			}
			if got := compressed.Header().Get("Content-Encoding"); got != "gzip" {
				t.Fatalf("Content-Encoding = %q, want gzip", got)
			}
			if !strings.Contains(compressed.Header().Get("Vary"), "Accept-Encoding") {
				t.Fatalf("Vary = %q, want Accept-Encoding listed", compressed.Header().Get("Vary"))
			}
			if compressed.Body.Len() >= identity.Body.Len() {
				t.Fatalf("compressed %d bytes >= identity %d bytes", compressed.Body.Len(), identity.Body.Len())
			}
			if !bytes.Equal(gunzipBody(t, compressed.Body.Bytes()), identity.Body.Bytes()) {
				t.Fatal("compressed representation does not decompress to the identity body")
			}
		})
	}

	t.Run("small bodies stay identity", func(t *testing.T) {
		w := h.do(t, request{path: "/healthz", headers: map[string]string{"Accept-Encoding": "gzip"}})
		if w.Code != http.StatusOK || w.Header().Get("Content-Encoding") != "" {
			t.Fatalf("healthz status=%d encoding=%q; tiny responses must not be compressed",
				w.Code, w.Header().Get("Content-Encoding"))
		}
	})

	t.Run("gzip refused via quality zero", func(t *testing.T) {
		w := h.do(t, request{path: fatV1Route, headers: map[string]string{"Accept-Encoding": "gzip;q=0"}})
		if w.Header().Get("Content-Encoding") != "" {
			t.Fatal("gzip;q=0 must refuse compression")
		}
	})
}

// TestConditionalRequests proves the full revalidation loop through the
// production chain: first response carries a weak validator and the default
// Cache-Control, a matching If-None-Match answers an empty 304 (also when
// the client negotiates gzip), and a stale validator gets the full body.
func TestConditionalRequests(t *testing.T) {
	h := server(t)

	first := h.get(t, fatV1Route)
	if first.Code != http.StatusOK {
		t.Fatalf("status = %d", first.Code)
	}
	etag := first.Header().Get("ETag")
	if !strings.HasPrefix(etag, `W/"`) {
		t.Fatalf("ETag = %q, want a weak validator", etag)
	}
	if got := first.Header().Get("Cache-Control"); got != "no-cache" {
		t.Fatalf("Cache-Control = %q, want the no-cache default", got)
	}

	t.Run("matching validator yields empty 304", func(t *testing.T) {
		w := h.do(t, request{path: fatV1Route, headers: map[string]string{"If-None-Match": etag}})
		if w.Code != http.StatusNotModified {
			t.Fatalf("status = %d, want 304", w.Code)
		}
		if w.Body.Len() != 0 {
			t.Fatalf("304 carried %d body bytes", w.Body.Len())
		}
		if got := w.Header().Get("ETag"); got != etag {
			t.Fatalf("304 ETag = %q, want %q", got, etag)
		}
	})

	t.Run("revalidation works for gzip clients", func(t *testing.T) {
		w := h.do(t, request{path: fatV1Route, headers: map[string]string{
			"If-None-Match":   etag,
			"Accept-Encoding": "gzip",
		}})
		if w.Code != http.StatusNotModified || w.Body.Len() != 0 {
			t.Fatalf("status=%d bodyBytes=%d, want empty 304", w.Code, w.Body.Len())
		}
		if w.Header().Get("Content-Encoding") != "" {
			t.Fatal("304 must not claim an encoding")
		}
	})

	t.Run("stale validator receives the full body", func(t *testing.T) {
		w := h.do(t, request{path: fatV1Route, headers: map[string]string{
			"If-None-Match": `W/"00000000000000000000000000000000"`,
		}})
		if w.Code != http.StatusOK || !bytes.Equal(w.Body.Bytes(), first.Body.Bytes()) {
			t.Fatalf("status=%d, body equal=%v; want the full 200", w.Code, bytes.Equal(w.Body.Bytes(), first.Body.Bytes()))
		}
	})

	t.Run("v2 responses carry validators too", func(t *testing.T) {
		w := h.get(t, fatV2Route)
		v2Tag := w.Header().Get("ETag")
		if w.Code != http.StatusOK || !strings.HasPrefix(v2Tag, `W/"`) {
			t.Fatalf("status=%d etag=%q", w.Code, v2Tag)
		}
		notModified := h.do(t, request{path: fatV2Route, headers: map[string]string{"If-None-Match": v2Tag}})
		if notModified.Code != http.StatusNotModified {
			t.Fatalf("v2 revalidation = %d, want 304", notModified.Code)
		}
	})

	t.Run("error responses are not validated", func(t *testing.T) {
		w := h.do(t, request{path: "/api/does/not/exist", headers: map[string]string{"If-None-Match": "*"}})
		if w.Code != http.StatusNotFound {
			t.Fatalf("status = %d, want 404", w.Code)
		}
		if w.Header().Get("ETag") != "" || w.Header().Get("Cache-Control") != "" {
			t.Fatal("404 must not gain cache metadata")
		}
	})
}

// TestVersionEndpoint pins the /version shape: the values are
// build-dependent (hence no golden), but every field is present, non-empty
// and stable within a process.
func TestVersionEndpoint(t *testing.T) {
	h := server(t)

	first := h.get(t, "/version")
	second := h.get(t, "/version")
	if first.Code != http.StatusOK {
		t.Fatalf("status = %d\n%s", first.Code, first.Body.String())
	}
	if got := contentTypeOf(first); got != "application/json" {
		t.Fatalf("Content-Type = %q", got)
	}
	if !bytes.Equal(first.Body.Bytes(), second.Body.Bytes()) {
		t.Fatal("/version must be deterministic within a process")
	}

	var payload map[string]string
	if err := json.Unmarshal(first.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode: %v\n%s", err, first.Body.String())
	}
	for _, key := range []string{"version", "commit", "buildDate", "goVersion"} {
		if payload[key] == "" {
			t.Errorf("field %q missing or empty: %s", key, first.Body.String())
		}
	}
	if len(payload) != 4 {
		t.Errorf("payload has %d fields, want 4: %s", len(payload), first.Body.String())
	}
	if !strings.HasPrefix(payload["goVersion"], "go") {
		t.Errorf("goVersion = %q", payload["goVersion"])
	}

	// Build identity is metadata about the process, not a cacheable API
	// read — but it flows through the same chain, so it gets a validator
	// like every JSON 200. Revalidation must work.
	if etag := first.Header().Get("ETag"); etag != "" {
		w := h.do(t, request{path: "/version", headers: map[string]string{"If-None-Match": etag}})
		if w.Code != http.StatusNotModified {
			t.Errorf("/version revalidation = %d, want 304", w.Code)
		}
	}
}
