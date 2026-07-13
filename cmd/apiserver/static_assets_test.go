package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeFile creates a file (and parent dirs) with trivial content.
func writeFile(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte("x"), 0o600); err != nil {
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

	// Legacy package layout for another seed (image.png / video.mp4).
	pkgSeed := "0x95c22c676296e4cca545427249c75a2cc99e4e3a43abd24869c9a09ca5960f8a"
	pkgPNG := filepath.Join(csDir, pkgSeed, "image.png")
	pkgMP4 := filepath.Join(csDir, pkgSeed, "video.mp4")
	writeFile(t, pkgPNG)
	writeFile(t, pkgMP4)

	// Current package layout (nested images/ + videos/ subdirectories).
	newSeed := "0xe29887e5f8aea85d6b775ab8dc95df16a5a0ad2979ace2b539058a0040aca67d"
	newFull := filepath.Join(csDir, newSeed, "images", "web", "full.webp")
	newPreview := filepath.Join(csDir, newSeed, "images", "web", "preview.webp")
	newMaster := filepath.Join(csDir, newSeed, "images", "source", "master.png")
	newVideoWeb := filepath.Join(csDir, newSeed, "videos", "web", "main.mp4")
	newVideoHQ := filepath.Join(csDir, newSeed, "videos", "hq", "main.mp4")
	writeFile(t, newFull)
	writeFile(t, newPreview)
	writeFile(t, newMaster)
	writeFile(t, newVideoWeb)
	writeFile(t, newVideoHQ)

	// Current layout for a seed that only has a master.png (no web/full.webp),
	// to exercise the source/master.png fallback.
	masterOnlySeed := "0x7f12ef4506339e9ccc3c665a5ae8c1878fcf241b7dcbdc4209a95a164622e204"
	masterOnly := filepath.Join(csDir, masterOnlySeed, "images", "source", "master.png")
	writeFile(t, masterOnly)

	// Current layout for a seed that only has the hq video (no web/main.mp4).
	hqOnlySeed := "0xae7e4b937e44eddb5693b29148b6bd03f65417f7a363b60829a562ca2370ec0d"
	hqOnly := filepath.Join(csDir, hqOnlySeed, "videos", "hq", "main.mp4")
	writeFile(t, hqOnly)

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
			name:   "new layout png -> images/web/full.webp",
			urlRel: "new/cosmicsignature/" + newSeed + ".png",
			want:   newFull,
			wantOK: true,
		},
		{
			name:   "new layout mp4 -> videos/web/main.mp4",
			urlRel: "new/cosmicsignature/" + newSeed + ".mp4",
			want:   newVideoWeb,
			wantOK: true,
		},
		{
			name:   "new layout thumb_card -> images/web/preview.webp",
			urlRel: "new/cosmicsignature/" + newSeed + "/thumb_card.webp",
			want:   newPreview,
			wantOK: true,
		},
		{
			name:   "new layout thumb_micro -> images/web/preview.webp",
			urlRel: "new/cosmicsignature/" + newSeed + "/thumb_micro.webp",
			want:   newPreview,
			wantOK: true,
		},
		{
			name:   "png falls back to source/master.png when no web image",
			urlRel: "new/cosmicsignature/" + masterOnlySeed + ".png",
			want:   masterOnly,
			wantOK: true,
		},
		{
			name:   "mp4 falls back to videos/hq/main.mp4 when no web video",
			urlRel: "new/cosmicsignature/" + hqOnlySeed + ".mp4",
			want:   hqOnly,
			wantOK: true,
		},
		{
			name:   "bare-hex seed (no 0x) normalizes to 0x<seed> dir",
			urlRel: "new/cosmicsignature/" + strings.TrimPrefix(newSeed, "0x") + ".png",
			want:   newFull,
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
