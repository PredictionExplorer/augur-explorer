package main

// Handler-level tests for the static asset routes: the /images mount with
// the cache-control middleware and Cosmic Signature layout fallbacks, and
// the /static ABI directory. These exercise the HTTP layer end to end
// (router + middleware + handler); pure path resolution is covered by
// static_assets_test.go.

import (
	"bytes"
	"compress/gzip"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
)

// testStaticAssets fills the logger for a static asset configuration under
// test.
func testStaticAssets(sa staticAssets) staticAssets {
	sa.logger = slog.New(slog.DiscardHandler)
	return sa
}

// staticRouter builds a router with only the static asset routes, from a
// temp asset tree with one RandomWalk thumb and one packaged Cosmic seed.
func staticRouter(t *testing.T, opts staticAssets) (*httpx.Router, string) {
	t.Helper()
	root := t.TempDir()

	// Standard layout: randomwalk/ + new/cosmicsignature/ under the root.
	rwThumb := filepath.Join(root, "randomwalk", "000010_black_thumb.jpg")
	writeFile(t, rwThumb)

	const seed = "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"
	csFull := filepath.Join(root, "new", "cosmicsignature", seed, "images", "web", "full.webp")
	writeFile(t, csFull)

	opts.nftAssetsRoot = root
	r := httpx.NewRouter()
	registerStaticAssetRoutes(testStaticAssets(opts))(r)
	return r, root
}

func doStatic(r *httpx.Router, method, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(method, path, nil))
	return w
}

func TestStaticImagesServeAndCacheControl(t *testing.T) {
	r, _ := staticRouter(t, staticAssets{})

	w := doStatic(r, http.MethodGet, "/images/randomwalk/000010_black_thumb.jpg")
	if w.Code != http.StatusOK || w.Body.String() != "x" {
		t.Fatalf("GET thumb = %d %q, want 200 with file content", w.Code, w.Body.String())
	}
	if cc := w.Header().Get("Cache-Control"); cc != "public, max-age=3600, must-revalidate" {
		t.Errorf("Cache-Control = %q", cc)
	}
}

func TestStaticImagesHead(t *testing.T) {
	r, _ := staticRouter(t, staticAssets{})

	w := doStatic(r, http.MethodHead, "/images/randomwalk/000010_black_thumb.jpg")
	if w.Code != http.StatusOK {
		t.Fatalf("HEAD thumb = %d, want 200", w.Code)
	}
	if cl := w.Header().Get("Content-Length"); cl != "1" {
		t.Errorf("HEAD Content-Length = %q, want 1", cl)
	}
}

func TestStaticImagesMissing404WithCacheHeader(t *testing.T) {
	r, _ := staticRouter(t, staticAssets{})

	w := doStatic(r, http.MethodGet, "/images/randomwalk/does_not_exist.jpg")
	if w.Code != http.StatusNotFound {
		t.Fatalf("missing file = %d, want 404", w.Code)
	}
	// The cache middleware runs before the handler, so even 404s carry it
	// (mirrors the legacy chain; browsers cache the miss briefly).
	if cc := w.Header().Get("Cache-Control"); cc == "" {
		t.Error("404 lost the Cache-Control header")
	}
}

func TestStaticImagesCosmicPackageFallback(t *testing.T) {
	r, _ := staticRouter(t, staticAssets{})

	// The flat URL resolves onto the packaged full.webp via resolveAssetFile.
	const flatURL = "/images/new/cosmicsignature/0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890.png"
	w := doStatic(r, http.MethodGet, flatURL)
	if w.Code != http.StatusOK {
		t.Fatalf("packaged seed via flat URL = %d, want 200", w.Code)
	}
}

func TestStaticImagesTraversalBlocked(t *testing.T) {
	r, root := staticRouter(t, staticAssets{})
	writeFile(t, filepath.Join(root, "..", "outside.txt"))

	// Encoded traversal: the raw path dodges the server's path cleaning and
	// must be stopped by safeFileUnderRoot.
	req := httptest.NewRequest(http.MethodGet, "/images/randomwalk/x", nil)
	req.URL.Path = "/images/randomwalk/../../outside.txt"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code == http.StatusOK {
		t.Fatalf("traversal escaped the asset root (status %d)", w.Code)
	}
}

func TestStaticImagesNoCacheOption(t *testing.T) {
	r, _ := staticRouter(t, staticAssets{noCache: true})

	w := doStatic(r, http.MethodGet, "/images/randomwalk/000010_black_thumb.jpg")
	if cc := w.Header().Get("Cache-Control"); cc != "no-store" {
		t.Errorf("Cache-Control = %q, want no-store", cc)
	}
}

func TestStaticABIDirServesFilesOnly(t *testing.T) {
	abiDir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(abiDir, "abi"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(abiDir, "abi", "RandomWalkNFT.json"), []byte(`{"abi":[]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	r := httpx.NewRouter()
	registerStaticAssetRoutes(testStaticAssets(staticAssets{abiDir: abiDir}))(r)

	if w := doStatic(r, http.MethodGet, "/static/abi/RandomWalkNFT.json"); w.Code != http.StatusOK || w.Body.String() != `{"abi":[]}` {
		t.Fatalf("GET abi = %d %q", w.Code, w.Body.String())
	}
	// Directory requests must not produce listings.
	if w := doStatic(r, http.MethodGet, "/static/abi/"); w.Code != http.StatusNotFound {
		t.Errorf("directory request = %d, want 404", w.Code)
	}
	if w := doStatic(r, http.MethodGet, "/static/missing.json"); w.Code != http.StatusNotFound {
		t.Errorf("missing file = %d, want 404", w.Code)
	}
}

func TestStaticRoutesSkippedWithoutConfig(t *testing.T) {
	t.Parallel()
	r := httpx.NewRouter()
	registerStaticAssetRoutes(testStaticAssets(staticAssets{}))(r)
	if n := len(r.Routes()); n != 0 {
		t.Fatalf("expected no static routes without config, got %d", n)
	}
}

// TestStaticAssetsThroughProductionChain composes the full production
// middleware chain (routes.New + RegisterExtra) and proves the response-edge
// middleware leaves file serving alone: images are never compressed and keep
// their cache policy, and ABI JSON files compress while http.ServeFile's
// Last-Modified validation — not a weak ETag — owns their revalidation.
func TestStaticAssetsThroughProductionChain(t *testing.T) {
	root := t.TempDir()
	// A "large" image (fake JPEG bytes) that would cross the compression
	// threshold if the content-type gate failed.
	imgPath := filepath.Join(root, "randomwalk", "000010_black_thumb.jpg")
	writeFile(t, imgPath)
	if err := os.WriteFile(imgPath, bytes.Repeat([]byte{0xff, 0xd8, 0x00, 0x01}, 1024), 0o600); err != nil {
		t.Fatal(err)
	}
	abiDir := t.TempDir()
	abiBody := `{"abi":[` + strings.Repeat(`{"name":"Transfer","type":"event"},`, 100) + `{}]}`
	if err := os.WriteFile(filepath.Join(abiDir, "RandomWalkNFT.json"), []byte(abiBody), 0o600); err != nil {
		t.Fatal(err)
	}

	r := routes.New(nil, routes.Options{
		RegisterExtra: registerStaticAssetRoutes(testStaticAssets(staticAssets{
			nftAssetsRoot: root,
			abiDir:        abiDir,
		})),
	})
	get := func(path string, headers map[string]string) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		req.RemoteAddr = "10.9.9.9:4242"
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		return w
	}

	t.Run("images stay identity with their cache policy", func(t *testing.T) {
		w := get("/images/randomwalk/000010_black_thumb.jpg", map[string]string{"Accept-Encoding": "gzip"})
		if w.Code != http.StatusOK {
			t.Fatalf("status = %d", w.Code)
		}
		if w.Header().Get("Content-Encoding") != "" {
			t.Fatal("image response was compressed")
		}
		if cc := w.Header().Get("Cache-Control"); cc != "public, max-age=3600, must-revalidate" {
			t.Fatalf("Cache-Control = %q; the image policy must win over the API default", cc)
		}
		if w.Header().Get("ETag") != "" {
			t.Fatal("file responses must not gain a weak ETag (ServeFile owns validation)")
		}
	})

	t.Run("abi json compresses and keeps ServeFile validation", func(t *testing.T) {
		w := get("/static/RandomWalkNFT.json", map[string]string{"Accept-Encoding": "gzip"})
		if w.Code != http.StatusOK {
			t.Fatalf("status = %d", w.Code)
		}
		if got := w.Header().Get("Content-Encoding"); got != "gzip" {
			t.Fatalf("Content-Encoding = %q, want gzip for a large JSON file", got)
		}
		lastModified := w.Header().Get("Last-Modified")
		if lastModified == "" {
			t.Fatal("ServeFile response lost Last-Modified")
		}
		if w.Header().Get("ETag") != "" {
			t.Fatal("file responses must not gain a weak ETag")
		}
		zr, err := gzip.NewReader(bytes.NewReader(w.Body.Bytes()))
		if err != nil {
			t.Fatalf("gzip reader: %v", err)
		}
		defer func() { _ = zr.Close() }() // checksum already verified by ReadAll
		decompressed, err := io.ReadAll(zr)
		if err != nil {
			t.Fatalf("decompress: %v", err)
		}
		if string(decompressed) != abiBody {
			t.Fatal("compressed ABI file does not round-trip")
		}

		// ServeFile's native conditional handling still answers 304.
		notModified := get("/static/RandomWalkNFT.json", map[string]string{
			"Accept-Encoding":   "gzip",
			"If-Modified-Since": lastModified,
		})
		if notModified.Code != http.StatusNotModified {
			t.Fatalf("If-Modified-Since revalidation = %d, want 304", notModified.Code)
		}
		if notModified.Header().Get("Content-Encoding") != "" {
			t.Fatal("304 must not claim an encoding")
		}
	})
}

// TestResolveNFTStaticMountLayouts drives every on-disk layout the mount
// resolver understands, in both flat and nested URL modes.
func TestResolveNFTStaticMountLayouts(t *testing.T) {
	t.Parallel()
	nestedThumbs := func(t *testing.T) string {
		t.Helper()
		root := t.TempDir()
		writeFile(t, filepath.Join(root, "randomwalk", "000010_black_thumb.jpg"))
		return root
	}
	rootThumbs := func(t *testing.T) string {
		t.Helper()
		root := t.TempDir()
		writeFile(t, filepath.Join(root, "000010_black_thumb.jpg"))
		return root
	}
	emptyNested := func(t *testing.T) string {
		t.Helper()
		root := t.TempDir()
		if err := os.MkdirAll(filepath.Join(root, "randomwalk"), 0o750); err != nil {
			t.Fatal(err)
		}
		return root
	}
	cases := []struct {
		name      string
		layout    func(*testing.T) string
		flat      bool
		wantMount string
		wantRoot  func(abs string) string
	}{
		{"nested thumbs standard", nestedThumbs, false, "/images", func(abs string) string { return abs }},
		{"nested thumbs flat", nestedThumbs, true, "/images", func(abs string) string { return filepath.Join(abs, "randomwalk") }},
		{"root thumbs compact", rootThumbs, false, "/images/randomwalk", func(abs string) string { return abs }},
		{"root thumbs flat", rootThumbs, true, "/images", func(abs string) string { return abs }},
		{"empty nested dir standard", emptyNested, false, "/images", func(abs string) string { return abs }},
		{"empty nested dir flat", emptyNested, true, "/images", func(abs string) string { return filepath.Join(abs, "randomwalk") }},
		{"bare dir standard", func(t *testing.T) string { t.Helper(); return t.TempDir() }, false, "/images", func(abs string) string { return abs }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			abs := tc.layout(t)
			sa := testStaticAssets(staticAssets{flatPaths: tc.flat})
			mount, fsRoot := sa.resolveNFTStaticMount(abs)
			if mount != tc.wantMount || fsRoot != tc.wantRoot(abs) {
				t.Fatalf("mount, fsRoot = %q, %q; want %q, %q", mount, fsRoot, tc.wantMount, tc.wantRoot(abs))
			}
		})
	}

	t.Run("dir named randomwalk", func(t *testing.T) {
		t.Parallel()
		parent := t.TempDir()
		abs := filepath.Join(parent, "randomwalk")
		if err := os.MkdirAll(abs, 0o750); err != nil {
			t.Fatal(err)
		}
		sa := testStaticAssets(staticAssets{})
		if mount, fsRoot := sa.resolveNFTStaticMount(abs); mount != "/images/randomwalk" || fsRoot != abs {
			t.Fatalf("mount, fsRoot = %q, %q", mount, fsRoot)
		}
		saFlat := testStaticAssets(staticAssets{flatPaths: true})
		if mount, fsRoot := saFlat.resolveNFTStaticMount(abs); mount != "/images" || fsRoot != abs {
			t.Fatalf("flat mount, fsRoot = %q, %q", mount, fsRoot)
		}
	})
}

// TestStaticRegistrationLogsLayouts pins the registration-time operator
// diagnostics for each layout, including the missing-thumbs warning.
func TestStaticRegistrationLogsLayouts(t *testing.T) {
	t.Parallel()
	register := func(t *testing.T, root string, flat bool) string {
		t.Helper()
		var buf strings.Builder
		sa := staticAssets{
			nftAssetsRoot: root,
			flatPaths:     flat,
			logger:        slog.New(slog.NewTextHandler(&buf, nil)),
		}
		r := httpx.NewRouter()
		registerStaticAssetRoutes(sa)(r)
		return buf.String()
	}

	t.Run("compact layout logs the randomwalk mount", func(t *testing.T) {
		t.Parallel()
		root := t.TempDir()
		writeFile(t, filepath.Join(root, "000010_black_thumb.jpg"))
		log := register(t, root, false)
		if !strings.Contains(log, "token assets live directly in NFT_ASSETS_ROOT") ||
			!strings.Contains(log, "/images/new/...") {
			t.Fatalf("log = %q", log)
		}
	})

	t.Run("flat layout logs the nested root", func(t *testing.T) {
		t.Parallel()
		root := t.TempDir()
		writeFile(t, filepath.Join(root, "randomwalk", "000010_black_thumb.jpg"))
		log := register(t, root, true)
		if !strings.Contains(log, "flat URL layout, files under randomwalk/") {
			t.Fatalf("log = %q", log)
		}
	})

	t.Run("missing thumbs warns", func(t *testing.T) {
		t.Parallel()
		log := register(t, t.TempDir(), false)
		if !strings.Contains(log, "no *_black_thumb.jpg under") {
			t.Fatalf("log = %q", log)
		}
	})

	t.Run("no-cache note is logged", func(t *testing.T) {
		t.Parallel()
		var buf strings.Builder
		root := t.TempDir()
		writeFile(t, filepath.Join(root, "randomwalk", "000010_black_thumb.jpg"))
		sa := staticAssets{
			nftAssetsRoot: root,
			noCache:       true,
			logger:        slog.New(slog.NewTextHandler(&buf, nil)),
		}
		r := httpx.NewRouter()
		registerStaticAssetRoutes(sa)(r)
		if !strings.Contains(buf.String(), "WEBSRV_IMAGE_NO_CACHE: /images responses use Cache-Control: no-store") {
			t.Fatalf("log = %q", buf.String())
		}
	})
}
