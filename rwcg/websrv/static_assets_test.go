package main

import (
	"os"
	"path/filepath"
	"testing"
)

// writeFile creates a file (and parent dirs) with trivial content.
func writeFile(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte("x"), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func TestResolveAssetFile(t *testing.T) {
	root := t.TempDir()
	csDir := filepath.Join(root, "new", "cosmicsignature")

	// Flat legacy layout for one seed.
	flatSeed := "0xf4bcfd7913601331aff2ff6f771f05806cd4983381a217fc02b7e4f2c497f788"
	flatPNG := filepath.Join(csDir, flatSeed+".png")
	flatMP4 := filepath.Join(csDir, flatSeed+".mp4")
	writeFile(t, flatPNG)
	writeFile(t, flatMP4)

	// Package layout for another seed.
	pkgSeed := "0x95c22c676296e4cca545427249c75a2cc99e4e3a43abd24869c9a09ca5960f8a"
	pkgPNG := filepath.Join(csDir, pkgSeed, "image.png")
	pkgMP4 := filepath.Join(csDir, pkgSeed, "video.mp4")
	writeFile(t, pkgPNG)
	writeFile(t, pkgMP4)

	// A RandomWalk-style package dir must NOT be served via the fallback:
	// the package fallback is scoped to new/cosmicsignature/ only.
	rwPkgImage := filepath.Join(root, "randomwalk", "0xabc", "image.png")
	writeFile(t, rwPkgImage)

	cases := []struct {
		name   string
		urlRel string
		want   string
		wantOK bool
	}{
		{
			name:   "flat png served directly",
			urlRel: "new/cosmicsignature/" + flatSeed + ".png",
			want:   flatPNG,
			wantOK: true,
		},
		{
			name:   "flat mp4 served directly",
			urlRel: "new/cosmicsignature/" + flatSeed + ".mp4",
			want:   flatMP4,
			wantOK: true,
		},
		{
			name:   "package png fallback",
			urlRel: "new/cosmicsignature/" + pkgSeed + ".png",
			want:   pkgPNG,
			wantOK: true,
		},
		{
			name:   "package mp4 fallback",
			urlRel: "new/cosmicsignature/" + pkgSeed + ".mp4",
			want:   pkgMP4,
			wantOK: true,
		},
		{
			name:   "missing seed",
			urlRel: "new/cosmicsignature/0xdeadbeef.png",
			wantOK: false,
		},
		{
			name:   "unsupported extension has no fallback",
			urlRel: "new/cosmicsignature/" + pkgSeed + ".json",
			wantOK: false,
		},
		{
			name:   "randomwalk package layout is not served via fallback",
			urlRel: "randomwalk/0xabc.png",
			wantOK: false,
		},
		{
			name:   "path traversal rejected",
			urlRel: "../../../etc/passwd",
			wantOK: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := resolveAssetFile(root, tc.urlRel)
			if ok != tc.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tc.wantOK)
			}
			if tc.wantOK && got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
