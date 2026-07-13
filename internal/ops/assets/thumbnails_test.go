package assets

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type commandCall struct {
	name string
	args []string
}

type fakeCommandRunner struct {
	calls []commandCall
	run   func(context.Context, string, []string) ([]byte, error)
}

func (f *fakeCommandRunner) CombinedOutput(ctx context.Context, name string, args ...string) ([]byte, error) {
	f.calls = append(f.calls, commandCall{name: name, args: append([]string(nil), args...)})
	if f.run != nil {
		return f.run(ctx, name, args)
	}
	if err := os.WriteFile(args[len(args)-1], []byte("webp"), 0o600); err != nil {
		return nil, err
	}
	return nil, nil
}

type fixedClock struct{ now time.Time }

func (f fixedClock) Now() time.Time { return f.now }

func setModTime(t *testing.T, path string, when time.Time) {
	t.Helper()
	if err := os.Chtimes(path, when, when); err != nil {
		t.Fatalf("Chtimes(%q): %v", path, err)
	}
}

func thumbnailTestOptions(base string, source TokenSource, runner CommandRunner, now time.Time) ThumbnailOptions {
	return ThumbnailOptions{
		Source:     source,
		BaseDir:    base,
		Schema:     "public",
		MagickPath: "/test/bin/magick",
		MinAge:     10 * time.Second,
		Runner:     runner,
		Clock:      fixedClock{now: now},
	}
}

func TestGenerateThumbnailsCommandArguments(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	now := time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)
	dir := filepath.Join(base, "0xabc")
	src := filepath.Join(dir, "image.png")
	writeTestFile(t, src)
	setModTime(t, src, now.Add(-time.Hour))

	runner := &fakeCommandRunner{}
	var logs bytes.Buffer
	opts := thumbnailTestOptions(base, &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "0xABC"}}}, runner, now)
	opts.Logger = log.New(&logs, "", 0)
	summary, err := GenerateThumbnails(context.Background(), opts)
	if err != nil {
		t.Fatalf("GenerateThumbnails(): %v", err)
	}
	if summary != (ThumbnailSummary{Generated: 2}) {
		t.Fatalf("summary = %#v", summary)
	}
	if len(runner.calls) != 2 {
		t.Fatalf("command count = %d", len(runner.calls))
	}
	wantCardPrefix := []string{
		src,
		"-strip",
		"-colorspace", "RGB",
		"-filter", "Lanczos",
		"-resize", "640x640",
		"-colorspace", "sRGB",
		"-modulate", "100,112,100",
		"-unsharp", "0x0.8+0.8+0.005",
		"-quality", "82",
		"-define", "webp:method=6",
	}
	cardArgs := runner.calls[0].args
	if runner.calls[0].name != opts.MagickPath ||
		strings.Join(cardArgs[:len(cardArgs)-1], "\x00") != strings.Join(wantCardPrefix, "\x00") {
		t.Fatalf("card command = %#v", runner.calls[0])
	}
	assertThumbnailTempPath(t, cardArgs[len(cardArgs)-1], dir, "thumb_card.webp")
	if got := runner.calls[1].args; got[7] != "160x160" || got[13] != "0x0.6+0.7+0.003" || got[15] != "80" {
		t.Fatalf("micro command args = %#v", got)
	}
	for _, name := range []string{"thumb_card.webp", "thumb_micro.webp"} {
		if !isRegularFile(filepath.Join(dir, name)) {
			t.Errorf("%s was not installed", name)
		}
		assertNoThumbnailTemps(t, dir, name)
	}
	if !strings.Contains(logs.String(), "generated=2 skipped=0") {
		t.Fatalf("logs = %q", logs.String())
	}
}

func TestGenerateThumbnailsStaleForceAndMinAge(t *testing.T) {
	t.Parallel()
	now := time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)
	t.Run("stale and current", func(t *testing.T) {
		base := t.TempDir()
		dir := filepath.Join(base, "0x1")
		src := filepath.Join(dir, "image.png")
		card := filepath.Join(dir, "thumb_card.webp")
		micro := filepath.Join(dir, "thumb_micro.webp")
		writeTestFile(t, src)
		writeTestFile(t, card)
		writeTestFile(t, micro)
		setModTime(t, src, now.Add(-time.Hour))
		setModTime(t, card, now.Add(-30*time.Minute))
		setModTime(t, micro, now.Add(-2*time.Hour))
		runner := &fakeCommandRunner{}

		summary, err := GenerateThumbnails(context.Background(), thumbnailTestOptions(
			base, &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "1"}}}, runner, now,
		))
		if err != nil {
			t.Fatal(err)
		}
		if summary.Generated != 1 || summary.Skipped != 1 || len(runner.calls) != 1 {
			t.Fatalf("summary=%#v calls=%#v", summary, runner.calls)
		}
		assertThumbnailTempPath(
			t,
			runner.calls[0].args[len(runner.calls[0].args)-1],
			dir,
			"thumb_micro.webp",
		)
	})

	t.Run("force", func(t *testing.T) {
		base := t.TempDir()
		dir := filepath.Join(base, "0x1")
		src := filepath.Join(dir, "image.png")
		writeTestFile(t, src)
		setModTime(t, src, now.Add(-time.Hour))
		for _, name := range []string{"thumb_card.webp", "thumb_micro.webp"} {
			path := filepath.Join(dir, name)
			writeTestFile(t, path)
			setModTime(t, path, now)
		}
		runner := &fakeCommandRunner{}
		opts := thumbnailTestOptions(base, &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "1"}}}, runner, now)
		opts.Force = true
		summary, err := GenerateThumbnails(context.Background(), opts)
		if err != nil {
			t.Fatal(err)
		}
		if summary.Generated != 2 || summary.Skipped != 0 || len(runner.calls) != 2 {
			t.Fatalf("summary=%#v calls=%d", summary, len(runner.calls))
		}
	})

	t.Run("minimum age", func(t *testing.T) {
		base := t.TempDir()
		src := filepath.Join(base, "0x1", "image.png")
		writeTestFile(t, src)
		setModTime(t, src, now.Add(-5*time.Second))
		runner := &fakeCommandRunner{}
		summary, err := GenerateThumbnails(context.Background(), thumbnailTestOptions(
			base, &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "1"}}}, runner, now,
		))
		if err != nil {
			t.Fatal(err)
		}
		if summary.TooFresh != 1 || len(runner.calls) != 0 {
			t.Fatalf("summary=%#v calls=%d", summary, len(runner.calls))
		}
	})
}

func TestGenerateThumbnailsPaddedAndMissingSources(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	now := time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)
	src := filepath.Join(base, "0x"+LeftPadSeed("a"), "image.png")
	writeTestFile(t, src)
	setModTime(t, src, now.Add(-time.Hour))
	runner := &fakeCommandRunner{}
	summary, err := GenerateThumbnails(context.Background(), thumbnailTestOptions(
		base,
		&fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "a"}, {TokenID: 2, Seed: "missing"}}},
		runner,
		now,
	))
	if err != nil {
		t.Fatal(err)
	}
	if summary.Generated != 2 || summary.SourceMissing != 1 {
		t.Fatalf("summary = %#v", summary)
	}
	if !strings.Contains(runner.calls[0].args[0], LeftPadSeed("a")) {
		t.Fatalf("padded source not used: %#v", runner.calls[0])
	}
}

func TestGenerateThumbnailsRejectsSymlinkEscape(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	outside := t.TempDir()
	src := filepath.Join(outside, "image.png")
	writeTestFile(t, src)
	if err := os.Symlink(outside, filepath.Join(base, "0xdead")); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}
	runner := &fakeCommandRunner{}
	summary, err := GenerateThumbnails(context.Background(), thumbnailTestOptions(
		base,
		&fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "dead"}}},
		runner,
		time.Now(),
	))
	if err != nil {
		t.Fatal(err)
	}
	if summary.SourceMissing != 1 || len(runner.calls) != 0 {
		t.Fatalf("summary=%#v calls=%#v", summary, runner.calls)
	}
	if isRegularFile(filepath.Join(outside, "thumb_card.webp")) {
		t.Fatal("thumbnail escaped the configured asset root")
	}

	insideDir := filepath.Join(base, "0xbeef")
	if err := os.MkdirAll(insideDir, 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(src, filepath.Join(insideDir, "image.png")); err != nil {
		t.Skipf("file symlinks unavailable: %v", err)
	}
	runner = &fakeCommandRunner{}
	summary, err = GenerateThumbnails(context.Background(), thumbnailTestOptions(
		base,
		&fakeTokenSource{tokens: []Token{{TokenID: 2, Seed: "beef"}}},
		runner,
		time.Now(),
	))
	if err != nil {
		t.Fatal(err)
	}
	if summary.SourceMissing != 1 || len(runner.calls) != 0 {
		t.Fatalf("file symlink summary=%#v calls=%#v", summary, runner.calls)
	}
}

func TestGenerateThumbnailsCommandFailures(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	now := time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)
	dir := filepath.Join(base, "0x1")
	src := filepath.Join(dir, "image.png")
	writeTestFile(t, src)
	setModTime(t, src, now.Add(-time.Hour))

	runner := &fakeCommandRunner{run: func(_ context.Context, _ string, args []string) ([]byte, error) {
		dst := args[len(args)-1]
		if strings.Contains(dst, "thumb_card") {
			if err := os.WriteFile(dst, []byte("partial"), 0o600); err != nil {
				t.Fatal(err)
			}
			return []byte("missing WebP delegate\n"), errors.New("exit status 1")
		}
		if err := os.WriteFile(dst, []byte("webp"), 0o600); err != nil {
			t.Fatal(err)
		}
		return nil, nil
	}}
	var logs bytes.Buffer
	opts := thumbnailTestOptions(base, &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "1"}}}, runner, now)
	opts.Logger = log.New(&logs, "", 0)
	summary, err := GenerateThumbnails(context.Background(), opts)
	if err == nil || !strings.Contains(err.Error(), "1 thumbnail(s)") {
		t.Fatalf("error = %v", err)
	}
	if summary.Failed != 1 || summary.Generated != 1 {
		t.Fatalf("summary = %#v", summary)
	}
	assertNoThumbnailTemps(t, dir, "thumb_card.webp")
	if !strings.Contains(logs.String(), "missing WebP delegate") {
		t.Fatalf("logs = %q", logs.String())
	}
}

func TestGenerateThumbnailsMissingRunnerOutputFails(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	now := time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)
	src := filepath.Join(base, "0x1", "image.png")
	writeTestFile(t, src)
	setModTime(t, src, now.Add(-time.Hour))
	runner := &fakeCommandRunner{run: func(context.Context, string, []string) ([]byte, error) {
		return nil, nil
	}}
	summary, err := GenerateThumbnails(context.Background(), thumbnailTestOptions(
		base, &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "1"}}}, runner, now,
	))
	if err == nil || summary.Failed != 2 {
		t.Fatalf("summary=%#v error=%v", summary, err)
	}
}

func TestGenerateThumbnailsCancellation(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	now := time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)
	src := filepath.Join(base, "0x1", "image.png")
	writeTestFile(t, src)
	setModTime(t, src, now.Add(-time.Hour))
	ctx, cancel := context.WithCancel(context.Background())
	runner := &fakeCommandRunner{run: func(ctx context.Context, _ string, _ []string) ([]byte, error) {
		cancel()
		<-ctx.Done()
		return nil, ctx.Err()
	}}
	summary, err := GenerateThumbnails(ctx, thumbnailTestOptions(
		base, &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "1"}}}, runner, now,
	))
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("summary=%#v error=%v", summary, err)
	}
	if summary.Failed != 0 {
		t.Fatalf("cancellation counted as failure: %#v", summary)
	}
}

func TestGenerateThumbnailsCancellationAfterLoadAndRunnerDeadline(t *testing.T) {
	t.Parallel()
	t.Run("after source load", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		source := &fakeTokenSource{
			tokens: []Token{{TokenID: 1, Seed: "1"}},
			load:   func(context.Context) { cancel() },
		}
		opts := thumbnailTestOptions(t.TempDir(), source, &fakeCommandRunner{}, time.Now())
		_, err := GenerateThumbnails(ctx, opts)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("runner deadline", func(t *testing.T) {
		base := t.TempDir()
		now := time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)
		src := filepath.Join(base, "0x1", "image.png")
		writeTestFile(t, src)
		setModTime(t, src, now.Add(-time.Hour))
		runner := &fakeCommandRunner{run: func(context.Context, string, []string) ([]byte, error) {
			return nil, context.DeadlineExceeded
		}}
		summary, err := GenerateThumbnails(context.Background(), thumbnailTestOptions(
			base, &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "1"}}}, runner, now,
		))
		if !errors.Is(err, context.DeadlineExceeded) || summary.Failed != 0 {
			t.Fatalf("summary=%#v error=%v", summary, err)
		}
	})
	t.Run("after runner success", func(t *testing.T) {
		base := t.TempDir()
		now := time.Date(2026, 7, 12, 12, 0, 0, 0, time.UTC)
		dir := filepath.Join(base, "0x1")
		src := filepath.Join(dir, "image.png")
		writeTestFile(t, src)
		setModTime(t, src, now.Add(-time.Hour))
		ctx, cancel := context.WithCancel(context.Background())
		runner := &fakeCommandRunner{run: func(_ context.Context, _ string, args []string) ([]byte, error) {
			if err := os.WriteFile(args[len(args)-1], []byte("webp"), 0o600); err != nil {
				t.Fatal(err)
			}
			cancel()
			return nil, nil
		}}
		summary, err := GenerateThumbnails(ctx, thumbnailTestOptions(
			base, &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "1"}}}, runner, now,
		))
		if !errors.Is(err, context.Canceled) || summary.Generated != 0 {
			t.Fatalf("summary=%#v error=%v", summary, err)
		}
		if _, statErr := os.Stat(filepath.Join(dir, "thumb_card.webp")); !errors.Is(statErr, os.ErrNotExist) {
			t.Fatalf("canceled output remains: %v", statErr)
		}
		assertNoThumbnailTemps(t, dir, "thumb_card.webp")
	})
}

func TestGenerateThumbnailsUsesDefaultClockAndLogger(t *testing.T) {
	t.Parallel()
	opts := ThumbnailOptions{
		Source:     &fakeTokenSource{tokens: []Token{{TokenID: 1, Seed: "missing"}}},
		BaseDir:    t.TempDir(),
		Schema:     "public",
		MagickPath: "magick",
		MinAge:     time.Second,
		Runner:     &fakeCommandRunner{},
	}
	summary, err := GenerateThumbnails(context.Background(), opts)
	if err != nil || summary.SourceMissing != 1 {
		t.Fatalf("summary=%#v error=%v", summary, err)
	}
}

func TestGenerateThumbnailsValidation(t *testing.T) {
	t.Parallel()
	base := t.TempDir()
	baseFile := filepath.Join(base, "not-a-directory")
	writeTestFile(t, baseFile)
	missingBase := filepath.Join(base, "missing")
	tests := []struct {
		name string
		opts ThumbnailOptions
	}{
		{"nil source", ThumbnailOptions{BaseDir: base, Schema: "public", MagickPath: "magick", Runner: &fakeCommandRunner{}}},
		{"nil runner", ThumbnailOptions{Source: &fakeTokenSource{}, BaseDir: base, Schema: "public", MagickPath: "magick"}},
		{"empty magick", ThumbnailOptions{Source: &fakeTokenSource{}, BaseDir: base, Schema: "public", Runner: &fakeCommandRunner{}}},
		{"negative age", ThumbnailOptions{Source: &fakeTokenSource{}, BaseDir: base, Schema: "public", MagickPath: "magick", MinAge: -time.Second, Runner: &fakeCommandRunner{}}},
		{"invalid schema", ThumbnailOptions{Source: &fakeTokenSource{}, BaseDir: base, Schema: "bad.schema", MagickPath: "magick", Runner: &fakeCommandRunner{}}},
		{"base file", ThumbnailOptions{Source: &fakeTokenSource{}, BaseDir: baseFile, Schema: "public", MagickPath: "magick", Runner: &fakeCommandRunner{}}},
		{"missing base", ThumbnailOptions{Source: &fakeTokenSource{}, BaseDir: missingBase, Schema: "public", MagickPath: "magick", Runner: &fakeCommandRunner{}}},
		{"source error", ThumbnailOptions{Source: &fakeTokenSource{err: errors.New("load failed")}, BaseDir: base, Schema: "public", MagickPath: "magick", Runner: &fakeCommandRunner{}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := GenerateThumbnails(context.Background(), test.opts); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := GenerateThumbnails(ctx, ThumbnailOptions{}); !errors.Is(err, context.Canceled) {
		t.Fatalf("pre-canceled error = %v", err)
	}
}

func assertThumbnailTempPath(t *testing.T, path, dir, outputName string) {
	t.Helper()
	if filepath.Dir(path) != dir {
		t.Fatalf("temp output directory = %q, want %q", filepath.Dir(path), dir)
	}
	name := filepath.Base(path)
	if !strings.HasPrefix(name, "."+outputName+".tmp-") || !strings.HasSuffix(name, ".webp") {
		t.Fatalf("temp output name = %q, want unique WebP temp for %q", name, outputName)
	}
}

func assertNoThumbnailTemps(t *testing.T, dir, outputName string) {
	t.Helper()
	matches, err := filepath.Glob(filepath.Join(dir, "."+outputName+".tmp-*.webp"))
	if err != nil {
		t.Fatalf("matching thumbnail temp files: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("thumbnail temp files remain: %v", matches)
	}
}
