package main

// Handler-level tests for the static asset routes: the /images mount with
// the cache-control middleware and Cosmic Signature layout fallbacks, and
// the /static ABI directory. These exercise the HTTP layer end to end
// (router + middleware + handler); pure path resolution is covered by
// static_assets_test.go.

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// staticRouter builds a router with only the static asset routes, from a
// temp asset tree with one RandomWalk thumb and one packaged Cosmic seed.
func staticRouter(t *testing.T) (*httpx.Router, string) {
	t.Helper()
	root := t.TempDir()

	// Standard layout: randomwalk/ + new/cosmicsignature/ under the root.
	rwThumb := filepath.Join(root, "randomwalk", "000010_black_thumb.jpg")
	writeFile(t, rwThumb)

	const seed = "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"
	csFull := filepath.Join(root, "new", "cosmicsignature", seed, "images", "web", "full.webp")
	writeFile(t, csFull)

	t.Setenv("NFT_ASSETS_ROOT", root)
	t.Setenv("NFT_ASSETS_FLAT_PATHS", "")
	t.Setenv("WEBSRV_IMAGE_NO_CACHE", "")
	t.Setenv("STATIC_ABI_DIR", "")

	r := httpx.NewRouter()
	registerStaticAssetRoutes(r)
	return r, root
}

func doStatic(r *httpx.Router, method, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(method, path, nil))
	return w
}

func TestStaticImagesServeAndCacheControl(t *testing.T) {
	r, _ := staticRouter(t)

	w := doStatic(r, http.MethodGet, "/images/randomwalk/000010_black_thumb.jpg")
	if w.Code != http.StatusOK || w.Body.String() != "x" {
		t.Fatalf("GET thumb = %d %q, want 200 with file content", w.Code, w.Body.String())
	}
	if cc := w.Header().Get("Cache-Control"); cc != "public, max-age=3600, must-revalidate" {
		t.Errorf("Cache-Control = %q", cc)
	}
}

func TestStaticImagesHead(t *testing.T) {
	r, _ := staticRouter(t)

	w := doStatic(r, http.MethodHead, "/images/randomwalk/000010_black_thumb.jpg")
	if w.Code != http.StatusOK {
		t.Fatalf("HEAD thumb = %d, want 200", w.Code)
	}
	if cl := w.Header().Get("Content-Length"); cl != "1" {
		t.Errorf("HEAD Content-Length = %q, want 1", cl)
	}
}

func TestStaticImagesMissing404WithCacheHeader(t *testing.T) {
	r, _ := staticRouter(t)

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
	r, _ := staticRouter(t)

	// The flat URL resolves onto the packaged full.webp via resolveAssetFile.
	const flatURL = "/images/new/cosmicsignature/0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890.png"
	w := doStatic(r, http.MethodGet, flatURL)
	if w.Code != http.StatusOK {
		t.Fatalf("packaged seed via flat URL = %d, want 200", w.Code)
	}
}

func TestStaticImagesTraversalBlocked(t *testing.T) {
	r, root := staticRouter(t)
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

func TestStaticImagesNoCacheEnv(t *testing.T) {
	r, _ := staticRouter(t)
	t.Setenv("WEBSRV_IMAGE_NO_CACHE", "1")

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
	t.Setenv("NFT_ASSETS_ROOT", "")
	t.Setenv("STATIC_ABI_DIR", abiDir)

	r := httpx.NewRouter()
	registerStaticAssetRoutes(r)

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

func TestStaticRoutesSkippedWithoutEnv(t *testing.T) {
	t.Setenv("NFT_ASSETS_ROOT", "")
	t.Setenv("STATIC_ABI_DIR", "")

	r := httpx.NewRouter()
	registerStaticAssetRoutes(r)
	if n := len(r.Routes()); n != 0 {
		t.Fatalf("expected no static routes without env config, got %d", n)
	}
}
