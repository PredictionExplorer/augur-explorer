package assets

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// CommandRunner executes an external command and captures its combined output.
type CommandRunner interface {
	CombinedOutput(ctx context.Context, name string, args ...string) ([]byte, error)
}

// ExecCommandRunner runs commands with os/exec.
type ExecCommandRunner struct{}

// CombinedOutput implements CommandRunner.
func (ExecCommandRunner) CombinedOutput(ctx context.Context, name string, args ...string) ([]byte, error) {
	// #nosec G204 -- the operator selects an ImageMagick executable; no shell is involved.
	return exec.CommandContext(ctx, name, args...).CombinedOutput()
}

// Clock provides the current time to age-sensitive asset operations.
type Clock interface {
	Now() time.Time
}

// Logger is the narrow logging surface used by asset operations.
type Logger interface {
	Printf(format string, args ...any)
}

type wallClock struct{}

func (wallClock) Now() time.Time { return time.Now() }

type discardLogger struct{}

func (discardLogger) Printf(string, ...any) {}

type thumbnailSpec struct {
	name    string
	edge    int
	quality int
	unsharp string
}

var thumbnailSpecs = []thumbnailSpec{
	{name: "thumb_card.webp", edge: 640, quality: 82, unsharp: "0x0.8+0.8+0.005"},
	{name: "thumb_micro.webp", edge: 160, quality: 80, unsharp: "0x0.6+0.7+0.003"},
}

// ThumbnailSummary contains per-run thumbnail counts. Generated and Skipped
// count individual outputs; SourceMissing and TooFresh count tokens.
type ThumbnailSummary struct {
	Generated     int
	Skipped       int
	SourceMissing int
	TooFresh      int
	Failed        int
}

// ThumbnailOptions configures GenerateThumbnails.
type ThumbnailOptions struct {
	Source     TokenSource
	BaseDir    string
	Schema     string
	Force      bool
	MagickPath string
	MinAge     time.Duration
	Runner     CommandRunner
	Logger     Logger
	Clock      Clock
}

// GenerateThumbnails creates the standard card and micro WebP thumbnails for
// all available per-seed image.png sources.
func GenerateThumbnails(ctx context.Context, opts ThumbnailOptions) (ThumbnailSummary, error) {
	var summary ThumbnailSummary
	if err := ctx.Err(); err != nil {
		return summary, err
	}
	if opts.Source == nil {
		return summary, fmt.Errorf("token source is nil")
	}
	if opts.Runner == nil {
		return summary, fmt.Errorf("command runner is nil")
	}
	if opts.MagickPath == "" {
		return summary, fmt.Errorf("ImageMagick path is empty")
	}
	if opts.MinAge < 0 {
		return summary, fmt.Errorf("minimum source age must be non-negative")
	}
	if err := ValidateSchema(opts.Schema); err != nil {
		return summary, err
	}

	absBase, err := filepath.Abs(opts.BaseDir)
	if err != nil {
		return summary, fmt.Errorf("resolve base dir %q: %w", opts.BaseDir, err)
	}
	st, err := os.Stat(absBase)
	if err != nil {
		return summary, fmt.Errorf("base dir %q is not a directory: %w", absBase, err)
	}
	if !st.IsDir() {
		return summary, fmt.Errorf("base dir %q is not a directory", absBase)
	}

	tokens, err := opts.Source.TokenSeeds(ctx, opts.Schema)
	if err != nil {
		return summary, fmt.Errorf("fetch token seeds: %w", err)
	}
	tokens = NormalizeTokens(tokens)
	logger := opts.Logger
	if logger == nil {
		logger = discardLogger{}
	}
	clock := opts.Clock
	if clock == nil {
		clock = wallClock{}
	}

	logger.Printf("thumbnails: base=%s seeds=%d force=%v magick=%s", absBase, len(tokens), opts.Force, opts.MagickPath)
	freshCutoff := clock.Now().Add(-opts.MinAge)
	for _, token := range tokens {
		if err := ctx.Err(); err != nil {
			return summary, err
		}
		dir, ok := seedSourceDir(absBase, token.Seed)
		if !ok {
			summary.SourceMissing++
			continue
		}
		src := filepath.Join(dir, "image.png")
		srcInfo, err := os.Stat(src)
		if err != nil || !srcInfo.Mode().IsRegular() {
			summary.SourceMissing++
			continue
		}
		if srcInfo.ModTime().After(freshCutoff) {
			summary.TooFresh++
			continue
		}

		for _, spec := range thumbnailSpecs {
			if err := ctx.Err(); err != nil {
				return summary, err
			}
			dst := filepath.Join(dir, spec.name)
			if !opts.Force && upToDate(dst, srcInfo.ModTime()) {
				summary.Skipped++
				continue
			}
			if err := generateThumbnail(ctx, opts.Runner, opts.MagickPath, src, dst, spec); err != nil {
				if ctxErr := ctx.Err(); ctxErr != nil {
					return summary, ctxErr
				}
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
					return summary, err
				}
				summary.Failed++
				logger.Printf("FAIL %s: %v", dst, err)
				continue
			}
			summary.Generated++
			logger.Printf("OK   %s", dst)
		}
	}

	logger.Printf("done: generated=%d skipped=%d src-missing=%d too-fresh=%d failed=%d",
		summary.Generated, summary.Skipped, summary.SourceMissing, summary.TooFresh, summary.Failed)
	if summary.Failed > 0 {
		return summary, fmt.Errorf("%d thumbnail(s) failed to generate", summary.Failed)
	}
	return summary, nil
}

func generateThumbnail(
	ctx context.Context,
	runner CommandRunner,
	magickPath, src, dst string,
	spec thumbnailSpec,
) error {
	tempFile, err := os.CreateTemp(
		filepath.Dir(dst),
		"."+filepath.Base(dst)+".tmp-*.webp",
	)
	if err != nil {
		return fmt.Errorf("create thumbnail temp file: %w", err)
	}
	tmp := tempFile.Name()
	if err := tempFile.Close(); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("close thumbnail temp file: %w", err)
	}
	if err := os.Remove(tmp); err != nil {
		return fmt.Errorf("prepare thumbnail temp path: %w", err)
	}
	defer func() { _ = os.Remove(tmp) }()

	args := thumbnailCommandArgs(src, tmp, spec)
	out, err := runner.CombinedOutput(ctx, magickPath, args...)
	if err != nil {
		message := strings.TrimSpace(string(out))
		if message == "" {
			return err
		}
		return fmt.Errorf("%w: %s", err, message)
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := os.Rename(tmp, dst); err != nil {
		return fmt.Errorf("install thumbnail: %w", err)
	}
	return nil
}

func thumbnailCommandArgs(src, dst string, spec thumbnailSpec) []string {
	return []string{
		src,
		"-strip",
		"-colorspace", "RGB",
		"-filter", "Lanczos",
		"-resize", fmt.Sprintf("%dx%d", spec.edge, spec.edge),
		"-colorspace", "sRGB",
		"-modulate", "100,112,100",
		"-unsharp", spec.unsharp,
		"-quality", fmt.Sprintf("%d", spec.quality),
		"-define", "webp:method=6",
		dst,
	}
}

func upToDate(dst string, srcMod time.Time) bool {
	st, err := os.Stat(dst)
	return err == nil && st.Mode().IsRegular() && !st.ModTime().Before(srcMod)
}

func seedSourceDir(base, seed string) (string, bool) {
	for _, candidate := range SeedNameCandidates(seed) {
		dir := filepath.Join(base, candidate.Name)
		sourcePath := filepath.Join(dir, "image.png")
		resolvedSource, ok := pathWithinBase(base, sourcePath)
		if ok && isRegularFile(resolvedSource) {
			return dir, true
		}
	}
	return "", false
}

func pathWithinBase(base, path string) (string, bool) {
	resolvedBase, err := filepath.EvalSymlinks(base)
	if err != nil {
		return "", false
	}
	resolvedPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", false
	}
	relative, err := filepath.Rel(resolvedBase, resolvedPath)
	if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return "", false
	}
	return resolvedPath, true
}
