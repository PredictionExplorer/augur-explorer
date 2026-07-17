// Fuzz targets for the static asset path resolution (MODERNIZATION.md §4.4).
// Security invariant: whatever the request path, a resolved file is always a
// regular file strictly inside the asset root — traversal cannot escape.
package main

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// fuzzAssetRoot builds a one-per-process asset tree with files inside the root
// (flat + package Cosmic Signature layouts) and a bait file OUTSIDE the root
// that a traversal bug would be able to reach.
var fuzzAssetRoot = sync.OnceValue(func() string {
	parent, err := os.MkdirTemp("", "assetfuzz")
	if err != nil {
		panic(err)
	}
	root := filepath.Join(parent, "root")
	seed := "0xf4bcfd7913601331aff2ff6f771f05806cd4983381a217fc02b7e4f2c497f788"
	for _, p := range []string{
		filepath.Join(parent, "outside-secret.txt"), // the bait: must never resolve
		filepath.Join(root, "new", "cosmicsignature", seed+".png"),
		filepath.Join(root, "new", "cosmicsignature", seed, "images", "web", "full.webp"),
		filepath.Join(root, "new", "cosmicsignature", seed, "images", "web", "preview.webp"),
		filepath.Join(root, "new", "cosmicsignature", seed, "videos", "web", "main.mp4"),
		filepath.Join(root, "randomwalk", "000001_black_thumb.jpg"),
	} {
		if err := os.MkdirAll(filepath.Dir(p), 0o750); err != nil {
			panic(err)
		}
		if err := os.WriteFile(p, []byte("x"), 0o600); err != nil {
			panic(err)
		}
	}
	return root
})

// assertInsideRoot fails the test unless full is a regular file strictly under root.
func assertInsideRoot(t *testing.T, root, full, urlRel string) {
	t.Helper()
	rel, err := filepath.Rel(root, full)
	if err != nil || rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		t.Fatalf("resolved path escapes root: urlRel=%q -> %q (rel %q, err %v)", urlRel, full, rel, err)
	}
	st, err := os.Stat(full)
	if err != nil {
		t.Fatalf("resolved path does not exist: urlRel=%q -> %q: %v", urlRel, full, err)
	}
	if st.IsDir() {
		t.Fatalf("resolved path is a directory: urlRel=%q -> %q", urlRel, full)
	}
}

func fuzzPathSeeds(f *testing.F) {
	f.Helper()
	seed := "0xf4bcfd7913601331aff2ff6f771f05806cd4983381a217fc02b7e4f2c497f788"
	for _, s := range []string{
		"new/cosmicsignature/" + seed + ".png",
		"new/cosmicsignature/" + seed + ".mp4",
		"new/cosmicsignature/" + seed + "/thumb_card.webp",
		"new/cosmicsignature/" + strings.TrimPrefix(seed, "0x") + ".png",
		"randomwalk/000001_black_thumb.jpg",
		"../outside-secret.txt",
		"new/cosmicsignature/../../../outside-secret.txt",
		"/etc/passwd",
		"..\\outside-secret.txt",
		"new/cosmicsignature/.png",
		"",
		"new/cosmicsignature/" + seed + "/thumb_/../../x.webp",
	} {
		f.Add(s)
	}
}

func FuzzSafeFileUnderRoot(f *testing.F) {
	fuzzPathSeeds(f)
	f.Fuzz(func(t *testing.T, urlRel string) {
		root := fuzzAssetRoot()
		full, ok := safeFileUnderRoot(root, urlRel)
		if !ok {
			return
		}
		assertInsideRoot(t, root, full, urlRel)
	})
}

func FuzzResolveAssetFile(f *testing.F) {
	fuzzPathSeeds(f)
	f.Fuzz(func(t *testing.T, urlRel string) {
		root := fuzzAssetRoot()
		full, ok := resolveAssetFile(root, urlRel)
		if !ok {
			return
		}
		assertInsideRoot(t, root, full, urlRel)
		// The package-layout fallback is scoped to Cosmic Signature: when it
		// rewrote the path (result differs from the literal request), the
		// request must have been for new/cosmicsignature/.
		if literal, literalOK := safeFileUnderRoot(root, urlRel); !literalOK || literal != full {
			if !strings.HasPrefix(urlRel, cosmicAssetPrefix) {
				t.Fatalf("fallback applied outside %q: urlRel=%q -> %q", cosmicAssetPrefix, urlRel, full)
			}
		}
	})
}

func FuzzNormalizeSeedSegment(f *testing.F) {
	for _, s := range []string{"", "0xabc", "0Xabc", "abc", "ABC", "xyz", "0x", "00", "å"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, seg string) {
		out := normalizeSeedSegment(seg)
		if out != seg && out != "0x"+seg {
			t.Fatalf("normalizeSeedSegment(%q) = %q: must be input or 0x-prefixed input", seg, out)
		}
		if again := normalizeSeedSegment(out); again != out {
			t.Fatalf("normalizeSeedSegment not idempotent: %q -> %q -> %q", seg, out, again)
		}
	})
}

func FuzzIsHex(f *testing.F) {
	for _, s := range []string{"", "0", "abcdefABCDEF0123456789", "g", "0x", " ", "å"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		got := isHex(s)
		want := true
		for _, r := range s {
			if !strings.ContainsRune("0123456789abcdefABCDEF", r) {
				want = false
				break
			}
		}
		if got != want {
			t.Fatalf("isHex(%q) = %v, want %v", s, got, want)
		}
	})
}
