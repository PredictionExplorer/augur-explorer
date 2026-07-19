package assets

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// AssetKind identifies one asset resolved by the inventory.
type AssetKind string

// The asset kinds the inventory scans for.
const (
	AssetImage   AssetKind = "image"
	AssetPreview AssetKind = "preview"
	AssetVideo   AssetKind = "video"
)

// AssetLocation records the first matching on-disk asset candidate.
type AssetLocation struct {
	Found  bool
	Path   string
	Layout string
	Padded bool
}

type candidateTemplate struct {
	inner   string
	flatExt string
	layout  string
}

var (
	imageCandidates = []candidateTemplate{
		{inner: "images/web/full.webp", layout: "web"},
		{inner: "images/source/master.png", layout: "source"},
		{inner: "image.png", layout: "legacy"},
		{flatExt: ".png", layout: "flat"},
	}
	previewCandidates = []candidateTemplate{
		{inner: "images/web/preview.webp", layout: "web"},
		{inner: "thumb_card.webp", layout: "thumb"},
		{inner: "thumb_micro.webp", layout: "thumb"},
	}
	videoCandidates = []candidateTemplate{
		{inner: "videos/web/main.mp4", layout: "web"},
		{inner: "videos/hq/main.mp4", layout: "hq"},
		{inner: "video.mp4", layout: "legacy"},
		{flatExt: ".mp4", layout: "flat"},
	}
	imageLayoutOrder   = []string{"web", "source", "legacy", "flat"}
	previewLayoutOrder = []string{"web", "thumb"}
	videoLayoutOrder   = []string{"web", "hq", "legacy", "flat"}
)

// AssetTotals holds presence and layout counts for one kind of asset.
type AssetTotals struct {
	Present int
	Missing int
	Layouts map[string]int
}

// MissingToken records the absent asset kinds for one token.
type MissingToken struct {
	Token   Token
	Missing []AssetKind
}

// InventorySummary is the deterministic result of one inventory run.
type InventorySummary struct {
	BaseDir       string
	DBSeeds       int
	Images        AssetTotals
	Previews      AssetTotals
	Videos        AssetTotals
	BothPresent   int
	BothMissing   int
	Partial       int
	MissingTokens []MissingToken
	PaddedTokens  []Token
}

// InventoryOptions configures RunInventory.
type InventoryOptions struct {
	Source      TokenSource
	BaseDir     string
	Schema      string
	Database    string
	MissingOnly bool
	ShowAll     bool
	Output      io.Writer
}

// RunInventory compares database seeds with every supported on-disk asset
// layout and writes the legacy opsctl report.
func RunInventory(ctx context.Context, opts InventoryOptions) (InventorySummary, error) {
	var summary InventorySummary
	if err := ctx.Err(); err != nil {
		return summary, err
	}
	if opts.Source == nil {
		return summary, errors.New("token source is nil")
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
	summary.BaseDir = absBase

	tokens, err := opts.Source.TokenSeeds(ctx, opts.Schema)
	if err != nil {
		return summary, fmt.Errorf("fetch token seeds: %w", err)
	}
	tokens = NormalizeTokens(tokens)
	summary.DBSeeds = len(tokens)
	summary.Images.Layouts = make(map[string]int)
	summary.Previews.Layouts = make(map[string]int)
	summary.Videos.Layouts = make(map[string]int)

	out := opts.Output
	if out == nil {
		out = io.Discard
	}
	fmt.Fprintln(out, "Cosmic Signature asset inventory")
	fmt.Fprintf(out, "  base dir : %s\n", absBase)
	fmt.Fprintf(out, "  database : %s\n", opts.Database)
	fmt.Fprintf(out, "  schema   : %s\n", opts.Schema)
	fmt.Fprintln(out, "  layout   : 0x<seed>/{images/web/full.webp,images/web/preview.webp,videos/web/main.mp4} (+legacy/flat fallbacks)")
	fmt.Fprintf(out, "  DB seeds : %d\n\n", len(tokens))

	var missingLines, paddedLines, allLines []string
	for _, token := range tokens {
		if err := ctx.Err(); err != nil {
			return summary, err
		}
		image := LocateAsset(absBase, token.Seed, AssetImage)
		preview := LocateAsset(absBase, token.Seed, AssetPreview)
		video := LocateAsset(absBase, token.Seed, AssetVideo)

		accumulateAsset(&summary.Images, image)
		accumulateAsset(&summary.Previews, preview)
		accumulateAsset(&summary.Videos, video)
		switch {
		case image.Found && video.Found:
			summary.BothPresent++
		case !image.Found && !video.Found:
			summary.BothMissing++
		default:
			summary.Partial++
		}

		missing := missingKinds(image, preview, video)
		if len(missing) > 0 {
			summary.MissingTokens = append(summary.MissingTokens, MissingToken{
				Token:   token,
				Missing: append([]AssetKind(nil), missing...),
			})
			names := make([]string, len(missing))
			for i, kind := range missing {
				names[i] = string(kind)
			}
			missingLines = append(missingLines,
				fmt.Sprintf("  #%-5d 0x%s  missing: %s", token.TokenID, token.Seed, strings.Join(names, ",")))
		}

		if (image.Found && image.Padded) || (preview.Found && preview.Padded) || (video.Found && video.Padded) {
			summary.PaddedTokens = append(summary.PaddedTokens, token)
			paddedLines = append(paddedLines,
				fmt.Sprintf("  #%-5d db=0x%s  file=0x%s", token.TokenID, token.Seed, LeftPadSeed(token.Seed)))
		}
		if opts.ShowAll {
			allLines = append(allLines,
				fmt.Sprintf("  #%-5d 0x%s  image=%s  preview=%s  video=%s",
					token.TokenID, token.Seed, assetStatus(image), assetStatus(preview), assetStatus(video)))
		}
	}

	if opts.ShowAll {
		fmt.Fprintln(out, "All tokens:")
		for _, line := range allLines {
			fmt.Fprintln(out, line)
		}
		fmt.Fprintln(out)
	}

	if len(missingLines) > 0 {
		slices.Sort(missingLines)
		fmt.Fprintf(out, "Tokens with missing assets (%d):\n", len(missingLines))
		for _, line := range missingLines {
			fmt.Fprintln(out, line)
		}
		fmt.Fprintln(out)
	} else if !opts.MissingOnly {
		fmt.Fprintln(out, "No missing assets. Every DB seed has an image, preview and video on disk.")
		fmt.Fprintln(out)
	}

	if len(paddedLines) > 0 {
		slices.Sort(paddedLines)
		fmt.Fprintf(out, "Seed-padding mismatches — found only via zero-padded 64-char name (%d):\n", len(paddedLines))
		for _, line := range paddedLines {
			fmt.Fprintln(out, line)
		}
		fmt.Fprintln(out)
	}
	if opts.MissingOnly {
		return summary, nil
	}

	fmt.Fprintln(out, "==== Totals ====")
	fmt.Fprintf(out, "DB seeds          : %d\n", summary.DBSeeds)
	fmt.Fprintf(out, "Images present    : %-5d%s\n", summary.Images.Present, layoutBreakdown(summary.Images.Layouts, imageLayoutOrder))
	fmt.Fprintf(out, "Images missing    : %d\n", summary.Images.Missing)
	fmt.Fprintf(out, "Previews present  : %-5d%s\n", summary.Previews.Present, layoutBreakdown(summary.Previews.Layouts, previewLayoutOrder))
	fmt.Fprintf(out, "Previews missing  : %d\n", summary.Previews.Missing)
	fmt.Fprintf(out, "Videos present    : %-5d%s\n", summary.Videos.Present, layoutBreakdown(summary.Videos.Layouts, videoLayoutOrder))
	fmt.Fprintf(out, "Videos missing    : %d\n", summary.Videos.Missing)
	fmt.Fprintf(out, "Both (img+vid)    : present %d, missing %d, partial %d\n",
		summary.BothPresent, summary.BothMissing, summary.Partial)
	if len(summary.PaddedTokens) > 0 {
		fmt.Fprintf(out, "Padded-name only  : %d\n", len(summary.PaddedTokens))
	}
	return summary, nil
}

func candidatesFor(kind AssetKind) []candidateTemplate {
	switch kind {
	case AssetVideo:
		return videoCandidates
	case AssetPreview:
		return previewCandidates
	default:
		return imageCandidates
	}
}

// LocateAsset returns the first regular-file match across the supported
// layouts, trying the normalized seed name before its padded variant.
func LocateAsset(base, seed string, kind AssetKind) AssetLocation {
	for _, name := range SeedNameCandidates(seed) {
		for _, candidate := range candidatesFor(kind) {
			var path string
			if candidate.inner != "" {
				path = filepath.Join(base, name.Name, filepath.FromSlash(candidate.inner))
			} else {
				path = filepath.Join(base, name.Name+candidate.flatExt)
			}
			resolvedPath, safe := pathWithinBase(base, path)
			if safe && isRegularFile(resolvedPath) {
				return AssetLocation{
					Found:  true,
					Path:   path,
					Layout: candidate.layout,
					Padded: name.Padded,
				}
			}
		}
	}
	return AssetLocation{}
}

func isRegularFile(path string) bool {
	st, err := os.Stat(path)
	return err == nil && st.Mode().IsRegular()
}

func accumulateAsset(total *AssetTotals, location AssetLocation) {
	if location.Found {
		total.Present++
		total.Layouts[location.Layout]++
		return
	}
	total.Missing++
}

func missingKinds(image, preview, video AssetLocation) []AssetKind {
	var missing []AssetKind
	if !image.Found {
		missing = append(missing, AssetImage)
	}
	if !preview.Found {
		missing = append(missing, AssetPreview)
	}
	if !video.Found {
		missing = append(missing, AssetVideo)
	}
	return missing
}

func assetStatus(location AssetLocation) string {
	if !location.Found {
		return "MISSING"
	}
	if location.Padded {
		return location.Layout + "(padded)"
	}
	return location.Layout
}

func layoutBreakdown(counts map[string]int, order []string) string {
	parts := make([]string, 0, len(counts))
	known := make(map[string]struct{}, len(order))
	for _, layout := range order {
		known[layout] = struct{}{}
		if counts[layout] > 0 {
			parts = append(parts, fmt.Sprintf("%s %d", layout, counts[layout]))
		}
	}
	var extra []string
	for layout, count := range counts {
		if _, ok := known[layout]; !ok && count > 0 {
			extra = append(extra, fmt.Sprintf("%s %d", layout, count))
		}
	}
	slices.Sort(extra)
	parts = append(parts, extra...)
	if len(parts) == 0 {
		return ""
	}
	return " (" + strings.Join(parts, ", ") + ")"
}
