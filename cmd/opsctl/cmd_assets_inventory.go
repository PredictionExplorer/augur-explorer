package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
	_ "github.com/lib/pq" // postgres driver
	"github.com/spf13/cobra"
)

// assetLoc records where (and whether) one asset was found on disk.
type assetLoc struct {
	found  bool
	path   string
	layout string // e.g. "web" | "source" | "hq" | "legacy" | "flat" | "thumb"
	padded bool   // matched only via the zero-padded 64-char seed name
}

// candTmpl is one place an asset may live. When inner != "" the path is
// 0x<seed>/<inner>; otherwise it is the flat file 0x<seed><flatExt>.
type candTmpl struct {
	inner   string // slash-separated path under the per-seed dir
	flatExt string // extension for the flat sibling file (used when inner == "")
	layout  string
}

// imageCandidates / previewCandidates / videoCandidates list the on-disk
// locations for each asset kind, ordered to match the web server's resolution
// priority (current layout first, then legacy, then flat).
func imageCandidates() []candTmpl {
	return []candTmpl{
		{inner: "images/web/full.webp", layout: "web"},
		{inner: "images/source/master.png", layout: "source"},
		{inner: "image.png", layout: "legacy"},
		{flatExt: ".png", layout: "flat"},
	}
}

func previewCandidates() []candTmpl {
	return []candTmpl{
		{inner: "images/web/preview.webp", layout: "web"},
		{inner: "thumb_card.webp", layout: "thumb"},
		{inner: "thumb_micro.webp", layout: "thumb"},
	}
}

func videoCandidates() []candTmpl {
	return []candTmpl{
		{inner: "videos/web/main.mp4", layout: "web"},
		{inner: "videos/hq/main.mp4", layout: "hq"},
		{inner: "video.mp4", layout: "legacy"},
		{flatExt: ".mp4", layout: "flat"},
	}
}

func candidatesFor(kind string) []candTmpl {
	switch kind {
	case "video":
		return videoCandidates()
	case "preview":
		return previewCandidates()
	default:
		return imageCandidates()
	}
}

// imageLayoutOrder / previewLayoutOrder / videoLayoutOrder give a stable
// display order for the totals.
var (
	imageLayoutOrder   = []string{"web", "source", "legacy", "flat"}
	previewLayoutOrder = []string{"web", "thumb"}
	videoLayoutOrder   = []string{"web", "hq", "legacy", "flat"}
)

// newAssetsInventoryCmd builds `opsctl assets inventory`, the replacement for
// the standalone asset_inventory tool.
func newAssetsInventoryCmd() *cobra.Command {
	var (
		dbConn      string
		baseDir     string
		schema      string
		missingOnly bool
		showAll     bool
	)
	cmd := &cobra.Command{
		Use:   "inventory",
		Short: "Inventory Cosmic Signature image/video assets on disk against DB seeds",
		Long: `Fetches every minted Cosmic Signature token seed from the database, then
checks the local filesystem for each token's image, preview (thumbnail) and
video.

The current per-seed package layout is 0x<seed>/{images/web/full.webp,
images/web/preview.webp,videos/web/main.mp4}; earlier layouts (legacy package,
generated thumbs, flat files) are still accepted as present, mirroring the web
server's fallback priority. A missing preview is non-fatal (the server falls
back to the full image), so previews are reported separately from the
image/video pivot. Tokens found only via the zero-padded 64-char seed name are
reported as padded-name matches.

Run from the asset directory (e.g. ~/nft-assets/new/cosmicsignature) or pass
--base. Without --db the connection is built from the PGSQL_* environment
variables.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAssetInventory(dbConn, baseDir, schema, missingOnly, showAll)
		},
	}
	cmd.Flags().StringVar(&dbConn, "db", "",
		"PostgreSQL connection string override. If empty, built from PGSQL_HOST, PGSQL_USERNAME, "+
			"PGSQL_DATABASE, PGSQL_PASSWORD (source your cosmic-etl-config.env first).")
	cmd.Flags().StringVar(&baseDir, "base", ".", "Base directory holding the per-seed assets (default: current directory)")
	cmd.Flags().StringVar(&schema, "schema", "public", "Database schema holding cg_mint_event")
	cmd.Flags().BoolVar(&missingOnly, "missing-only", false, "Only list tokens with at least one missing asset")
	cmd.Flags().BoolVar(&showAll, "all", false, "List every token with its image/video status")
	return cmd
}

func init() { assetsCmd.AddCommand(newAssetsInventoryCmd()) }

func runAssetInventory(dbConn, baseDir, schema string, missingOnly, showAll bool) error {
	absBase, err := filepath.Abs(baseDir)
	if err != nil {
		return fmt.Errorf("resolve base dir %q: %w", baseDir, err)
	}
	if st, err := os.Stat(absBase); err != nil || !st.IsDir() {
		return fmt.Errorf("base dir %q is not a directory: %v", absBase, err)
	}

	conn := dbConn
	if conn == "" {
		conn, err = toolutil.PostgresConnStringFromEnv()
		if err != nil {
			return fmt.Errorf("no --db flag and %w", err)
		}
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(2)
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	tokens, err := fetchTokenSeeds(db, schema)
	if err != nil {
		return fmt.Errorf("fetch token seeds: %w", err)
	}

	fmt.Println("Cosmic Signature asset inventory")
	fmt.Printf("  base dir : %s\n", absBase)
	fmt.Printf("  database : %s\n", redactConn(conn))
	fmt.Printf("  schema   : %s\n", schema)
	fmt.Printf("  layout   : 0x<seed>/{images/web/full.webp,images/web/preview.webp,videos/web/main.mp4} (+legacy/flat fallbacks)\n")
	fmt.Printf("  DB seeds : %d\n\n", len(tokens))

	var (
		imgPresent, imgMissing   int
		prevPresent, prevMissing int
		vidPresent, vidMissing   int
		bothPresent, bothMissing int
		partial                  int
		missingLines             []string
		paddedLines              []string
		allLines                 []string
	)
	imgLayouts := map[string]int{}
	prevLayouts := map[string]int{}
	vidLayouts := map[string]int{}

	for _, tk := range tokens {
		img := locateAsset(absBase, tk.seed, "image")
		prev := locateAsset(absBase, tk.seed, "preview")
		vid := locateAsset(absBase, tk.seed, "video")

		if img.found {
			imgPresent++
			imgLayouts[img.layout]++
		} else {
			imgMissing++
		}
		if prev.found {
			prevPresent++
			prevLayouts[prev.layout]++
		} else {
			prevMissing++
		}
		if vid.found {
			vidPresent++
			vidLayouts[vid.layout]++
		} else {
			vidMissing++
		}

		// "both" pivots on the renderable image+video pair; a missing preview is
		// non-fatal (server falls back to the full image) and reported on its own.
		switch {
		case img.found && vid.found:
			bothPresent++
		case !img.found && !vid.found:
			bothMissing++
		default:
			partial++
		}

		if !img.found || !vid.found || !prev.found {
			var miss []string
			if !img.found {
				miss = append(miss, "image")
			}
			if !prev.found {
				miss = append(miss, "preview")
			}
			if !vid.found {
				miss = append(miss, "video")
			}
			missingLines = append(missingLines,
				fmt.Sprintf("  #%-5d 0x%s  missing: %s", tk.tokenID, tk.seed, strings.Join(miss, ",")))
		}

		if (img.found && img.padded) || (prev.found && prev.padded) || (vid.found && vid.padded) {
			paddedLines = append(paddedLines,
				fmt.Sprintf("  #%-5d db=0x%s  file=0x%s", tk.tokenID, tk.seed, leftPadSeed(tk.seed)))
		}

		if showAll {
			allLines = append(allLines,
				fmt.Sprintf("  #%-5d 0x%s  image=%s  preview=%s  video=%s", tk.tokenID, tk.seed,
					assetStatus(img), assetStatus(prev), assetStatus(vid)))
		}
	}

	if showAll {
		fmt.Println("All tokens:")
		for _, l := range allLines {
			fmt.Println(l)
		}
		fmt.Println()
	}

	if len(missingLines) > 0 {
		fmt.Printf("Tokens with missing assets (%d):\n", len(missingLines))
		sort.Strings(missingLines)
		for _, l := range missingLines {
			fmt.Println(l)
		}
		fmt.Println()
	} else if !missingOnly {
		fmt.Println("No missing assets. Every DB seed has an image, preview and video on disk.")
		fmt.Println()
	}

	if len(paddedLines) > 0 {
		fmt.Printf("Seed-padding mismatches — found only via zero-padded 64-char name (%d):\n", len(paddedLines))
		for _, l := range paddedLines {
			fmt.Println(l)
		}
		fmt.Println()
	}

	if missingOnly {
		return nil
	}

	fmt.Println("==== Totals ====")
	fmt.Printf("DB seeds          : %d\n", len(tokens))
	fmt.Printf("Images present    : %-5d%s\n", imgPresent, layoutBreakdown(imgLayouts, imageLayoutOrder))
	fmt.Printf("Images missing    : %d\n", imgMissing)
	fmt.Printf("Previews present  : %-5d%s\n", prevPresent, layoutBreakdown(prevLayouts, previewLayoutOrder))
	fmt.Printf("Previews missing  : %d\n", prevMissing)
	fmt.Printf("Videos present    : %-5d%s\n", vidPresent, layoutBreakdown(vidLayouts, videoLayoutOrder))
	fmt.Printf("Videos missing    : %d\n", vidMissing)
	fmt.Printf("Both (img+vid)    : present %d, missing %d, partial %d\n", bothPresent, bothMissing, partial)
	if len(paddedLines) > 0 {
		fmt.Printf("Padded-name only  : %d\n", len(paddedLines))
	}
	return nil
}

// layoutBreakdown formats per-layout counts in a stable order, e.g.
// " (web 20, source 3, flat 1)". Returns "" when there is nothing to show.
func layoutBreakdown(counts map[string]int, order []string) string {
	var parts []string
	for _, k := range order {
		if counts[k] > 0 {
			parts = append(parts, fmt.Sprintf("%s %d", k, counts[k]))
		}
	}
	// Include any layout not in the known order (defensive).
	for k, v := range counts {
		known := false
		for _, o := range order {
			if o == k {
				known = true
				break
			}
		}
		if !known && v > 0 {
			parts = append(parts, fmt.Sprintf("%s %d", k, v))
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return " (" + strings.Join(parts, ", ") + ")"
}

var keywordPasswordRe = regexp.MustCompile(`password\s*=\s*('[^']*'|"[^"]*"|\S+)`)

// redactConn hides the password in a connection string for safe logging,
// handling both URL form (postgres://user:pass@host/db) and keyword form
// (host=... password=...).
func redactConn(conn string) string {
	if u, err := url.Parse(conn); err == nil && strings.HasPrefix(u.Scheme, "postgres") {
		if _, hasPW := u.User.Password(); hasPW {
			u.User = url.UserPassword(u.User.Username(), "***")
			return u.String()
		}
		return conn
	}
	return keywordPasswordRe.ReplaceAllString(conn, "password=***")
}

// locateAsset returns the first on-disk match across all candidate layouts,
// trying the exact seed name before the zero-padded variant. kind is "image",
// "preview" or "video".
func locateAsset(base, seed, kind string) assetLoc {
	cands := candidatesFor(kind)
	for _, c := range seedNameCandidates(seed) {
		for _, t := range cands {
			var p string
			if t.inner != "" {
				p = filepath.Join(base, c.name, filepath.FromSlash(t.inner))
			} else {
				p = filepath.Join(base, c.name+t.flatExt)
			}
			if isRegularFile(p) {
				return assetLoc{found: true, path: p, layout: t.layout, padded: c.padded}
			}
		}
	}
	return assetLoc{found: false}
}

// assetStatus renders one asset's presence for the --all listing.
func assetStatus(a assetLoc) string {
	if !a.found {
		return "MISSING"
	}
	if a.padded {
		return a.layout + "(padded)"
	}
	return a.layout
}
