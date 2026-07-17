package assets

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func writeTestFile(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		t.Fatalf("MkdirAll(%q): %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte("asset"), 0o600); err != nil {
		t.Fatalf("WriteFile(%q): %v", path, err)
	}
}

func TestLocateAssetEveryLayout(t *testing.T) {
	t.Parallel()
	seed := strings.Repeat("a", SeedHexLength)
	tests := []struct {
		name   string
		kind   AssetKind
		path   func(string) string
		layout string
	}{
		{"image web", AssetImage, func(base string) string {
			return filepath.Join(base, "0x"+seed, "images", "web", "full.webp")
		}, "web"},
		{"image source", AssetImage, func(base string) string {
			return filepath.Join(base, "0x"+seed, "images", "source", "master.png")
		}, "source"},
		{"image legacy", AssetImage, func(base string) string {
			return filepath.Join(base, "0x"+seed, "image.png")
		}, "legacy"},
		{"image flat", AssetImage, func(base string) string {
			return filepath.Join(base, "0x"+seed+".png")
		}, "flat"},
		{"preview web", AssetPreview, func(base string) string {
			return filepath.Join(base, "0x"+seed, "images", "web", "preview.webp")
		}, "web"},
		{"preview card", AssetPreview, func(base string) string {
			return filepath.Join(base, "0x"+seed, "thumb_card.webp")
		}, "thumb"},
		{"preview micro", AssetPreview, func(base string) string {
			return filepath.Join(base, "0x"+seed, "thumb_micro.webp")
		}, "thumb"},
		{"video web", AssetVideo, func(base string) string {
			return filepath.Join(base, "0x"+seed, "videos", "web", "main.mp4")
		}, "web"},
		{"video hq", AssetVideo, func(base string) string {
			return filepath.Join(base, "0x"+seed, "videos", "hq", "main.mp4")
		}, "hq"},
		{"video legacy", AssetVideo, func(base string) string {
			return filepath.Join(base, "0x"+seed, "video.mp4")
		}, "legacy"},
		{"video flat", AssetVideo, func(base string) string {
			return filepath.Join(base, "0x"+seed+".mp4")
		}, "flat"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			base := t.TempDir()
			path := test.path(base)
			writeTestFile(t, path)
			got := LocateAsset(base, seed, test.kind)
			if !got.Found || got.Path != path || got.Layout != test.layout || got.Padded {
				t.Fatalf("LocateAsset() = %#v", got)
			}
		})
	}
}

func TestLocateAssetPriorityAndPaddedFallback(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	seed := "ab"
	padded := LeftPadSeed(seed)

	paddedWeb := filepath.Join(base, "0x"+padded, "images", "web", "full.webp")
	rawFlat := filepath.Join(base, "0x"+seed+".png")
	writeTestFile(t, paddedWeb)
	writeTestFile(t, rawFlat)
	got := LocateAsset(base, seed, AssetImage)
	if got.Path != rawFlat || got.Layout != "flat" || got.Padded {
		t.Fatalf("raw-name priority = %#v", got)
	}

	if err := os.Remove(rawFlat); err != nil {
		t.Fatal(err)
	}
	got = LocateAsset(base, seed, AssetImage)
	if got.Path != paddedWeb || got.Layout != "web" || !got.Padded {
		t.Fatalf("padded fallback = %#v", got)
	}

	if err := os.Remove(paddedWeb); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(paddedWeb, 0o750); err != nil {
		t.Fatal(err)
	}
	if got := LocateAsset(base, seed, AssetImage); got.Found {
		t.Fatalf("directory accepted as asset: %#v", got)
	}
}

func TestLocateAssetRejectsTraversalAndSymlinkEscapes(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	outside := t.TempDir()
	writeTestFile(t, filepath.Join(outside, "image.png"))
	if err := os.Symlink(outside, filepath.Join(base, "0xdead")); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
	if got := LocateAsset(base, "../escape", AssetImage); got.Found {
		t.Fatalf("traversal seed resolved outside base: %#v", got)
	}
	if got := LocateAsset(base, "dead", AssetImage); got.Found {
		t.Fatalf("symlink seed resolved outside base: %#v", got)
	}
}

func TestRunInventorySummaryAndStableOutput(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	source := &fakeTokenSource{tokens: []Token{
		{TokenID: 4, Seed: "0X4"},
		{TokenID: 2, Seed: "2"},
		{TokenID: 1, Seed: "1"},
		{TokenID: 3, Seed: "3"},
		{TokenID: 99, Seed: " 0x1 "},
	}}

	writeTestFile(t, filepath.Join(base, "0x1", "images", "web", "full.webp"))
	writeTestFile(t, filepath.Join(base, "0x1", "images", "web", "preview.webp"))
	writeTestFile(t, filepath.Join(base, "0x1", "videos", "web", "main.mp4"))
	writeTestFile(t, filepath.Join(base, "0x2", "image.png"))

	padded4 := "0x" + LeftPadSeed("4")
	writeTestFile(t, filepath.Join(base, padded4, "images", "source", "master.png"))
	writeTestFile(t, filepath.Join(base, padded4, "thumb_card.webp"))
	writeTestFile(t, filepath.Join(base, padded4, "videos", "hq", "main.mp4"))

	run := func() (InventorySummary, string) {
		var out bytes.Buffer
		summary, err := RunInventory(context.Background(), InventoryOptions{
			Source:   source,
			BaseDir:  base,
			Schema:   "public",
			Database: "postgres://test/***",
			ShowAll:  true,
			Output:   &out,
		})
		if err != nil {
			t.Fatalf("RunInventory(): %v", err)
		}
		return summary, out.String()
	}
	summary, output := run()
	_, secondOutput := run()
	if output != secondOutput {
		t.Fatal("inventory output changed between identical runs")
	}

	if summary.DBSeeds != 4 {
		t.Fatalf("DBSeeds = %d, want 4", summary.DBSeeds)
	}
	if summary.Images.Present != 3 || summary.Images.Missing != 1 {
		t.Fatalf("image totals = %#v", summary.Images)
	}
	if summary.Previews.Present != 2 || summary.Previews.Missing != 2 {
		t.Fatalf("preview totals = %#v", summary.Previews)
	}
	if summary.Videos.Present != 2 || summary.Videos.Missing != 2 {
		t.Fatalf("video totals = %#v", summary.Videos)
	}
	if summary.BothPresent != 2 || summary.BothMissing != 1 || summary.Partial != 1 {
		t.Fatalf("pair totals = present %d missing %d partial %d",
			summary.BothPresent, summary.BothMissing, summary.Partial)
	}
	if got := summary.Images.Layouts; !reflect.DeepEqual(got, map[string]int{"web": 1, "legacy": 1, "source": 1}) {
		t.Fatalf("image layouts = %#v", got)
	}
	if len(summary.MissingTokens) != 2 ||
		!reflect.DeepEqual(summary.MissingTokens[0].Missing, []AssetKind{AssetPreview, AssetVideo}) ||
		!reflect.DeepEqual(summary.MissingTokens[1].Missing, []AssetKind{AssetImage, AssetPreview, AssetVideo}) {
		t.Fatalf("missing tokens = %#v", summary.MissingTokens)
	}
	if !reflect.DeepEqual(summary.PaddedTokens, []Token{{TokenID: 4, Seed: "4"}}) {
		t.Fatalf("padded tokens = %#v", summary.PaddedTokens)
	}

	for _, want := range []string{
		"Cosmic Signature asset inventory",
		"All tokens:\n  #1     0x1  image=web  preview=web  video=web",
		"#2     0x2  missing: preview,video",
		"#3     0x3  missing: image,preview,video",
		"file=0x" + LeftPadSeed("4"),
		"Images present    : 3     (web 1, source 1, legacy 1)",
		"Previews present  : 2     (web 1, thumb 1)",
		"Videos present    : 2     (web 1, hq 1)",
		"Both (img+vid)    : present 2, missing 1, partial 1",
	} {
		if !strings.Contains(output, want) {
			t.Errorf("output missing %q:\n%s", want, output)
		}
	}
}

func TestRunInventoryMissingOnlyAndCompleteMessage(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	source := &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "1"}}}

	var missingOut bytes.Buffer
	if _, err := RunInventory(context.Background(), InventoryOptions{
		Source: source, BaseDir: base, Schema: "public", MissingOnly: true, Output: &missingOut,
	}); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(missingOut.String(), "==== Totals ====") {
		t.Fatalf("missing-only output included totals:\n%s", missingOut.String())
	}

	writeTestFile(t, filepath.Join(base, "0x1", "images", "web", "full.webp"))
	writeTestFile(t, filepath.Join(base, "0x1", "images", "web", "preview.webp"))
	writeTestFile(t, filepath.Join(base, "0x1", "videos", "web", "main.mp4"))
	var completeOut bytes.Buffer
	if _, err := RunInventory(context.Background(), InventoryOptions{
		Source: source, BaseDir: base, Schema: "public", Output: &completeOut,
	}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(completeOut.String(), "No missing assets.") {
		t.Fatalf("complete output:\n%s", completeOut.String())
	}
}

func TestRunInventoryErrorsAndCancellation(t *testing.T) {
	t.Parallel()
	t.Run("already canceled", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := RunInventory(ctx, InventoryOptions{})
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("nil source", func(t *testing.T) {
		t.Parallel()
		_, err := RunInventory(context.Background(), InventoryOptions{BaseDir: t.TempDir(), Schema: "public"})
		if err == nil || !strings.Contains(err.Error(), "source") {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("invalid schema", func(t *testing.T) {
		t.Parallel()
		_, err := RunInventory(context.Background(), InventoryOptions{
			Source: &fakeTokenSource{}, BaseDir: t.TempDir(), Schema: "bad.schema",
		})
		if err == nil || !strings.Contains(err.Error(), "invalid database schema") {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("source error", func(t *testing.T) {
		t.Parallel()
		base := t.TempDir()
		want := errors.New("database unavailable")
		_, err := RunInventory(context.Background(), InventoryOptions{
			Source: &fakeTokenSource{err: want}, BaseDir: base, Schema: "public",
		})
		if !errors.Is(err, want) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("invalid base", func(t *testing.T) {
		t.Parallel()
		path := filepath.Join(t.TempDir(), "file")
		writeTestFile(t, path)
		_, err := RunInventory(context.Background(), InventoryOptions{
			Source: &fakeTokenSource{}, BaseDir: path, Schema: "public",
		})
		if err == nil || !strings.Contains(err.Error(), "not a directory") {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("missing base", func(t *testing.T) {
		t.Parallel()
		path := filepath.Join(t.TempDir(), "missing")
		_, err := RunInventory(context.Background(), InventoryOptions{
			Source: &fakeTokenSource{}, BaseDir: path, Schema: "public",
		})
		if err == nil || !errors.Is(err, os.ErrNotExist) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("cancel after load", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(context.Background())
		source := &fakeTokenSource{
			tokens: []Token{{TokenID: 1, Seed: "1"}},
			load:   func(context.Context) { cancel() },
		}
		_, err := RunInventory(ctx, InventoryOptions{
			Source: source, BaseDir: t.TempDir(), Schema: "public",
		})
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
	})
}

func TestLayoutBreakdownIncludesUnknownLayoutsStably(t *testing.T) {
	t.Parallel()
	if got := layoutBreakdown(nil, imageLayoutOrder); got != "" {
		t.Fatalf("empty breakdown = %q", got)
	}
	got := layoutBreakdown(map[string]int{
		"web": 2, "z-layout": 1, "a-layout": 3, "ignored": 0,
	}, imageLayoutOrder)
	if got != " (web 2, a-layout 3, z-layout 1)" {
		t.Fatalf("breakdown = %q", got)
	}
}
