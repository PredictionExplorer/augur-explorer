// Cosmic Signature image/video asset inventory.
//
// Connects to the SQL database, fetches every minted Cosmic Signature token
// seed, then checks the local filesystem (current directory as base by default)
// for the presence of each token's image, preview (thumbnail) and video.
//
// The current per-seed package layout (uploaded by the generator) is:
//
//	0x<seed>/images/web/full.webp        (full image)
//	0x<seed>/images/web/preview.webp     (small preview / thumbnail source)
//	0x<seed>/images/source/master.png    (full-res master)
//	0x<seed>/videos/web/main.mp4         (web video)
//	0x<seed>/videos/hq/main.mp4          (high-quality video)
//
// Earlier layouts are still accepted as "present" because the web server serves
// any of them (see resolveAssetFile in websrv), so the missing counts reflect
// assets that are genuinely not on disk:
//
//	0x<seed>/image.png  0x<seed>/video.mp4        (legacy package)
//	0x<seed>/thumb_card.webp  thumb_micro.webp    (legacy generated thumbs)
//	0x<seed>.png        0x<seed>.mp4              (flat)
//
// Detection order per asset mirrors the web server's fallback priority, and the
// matched layout is reported so you can see which tokens still rely on an old
// layout. A missing preview is non-fatal (the server falls back to the full
// image), so previews are reported separately from the image/video pivot.
//
// Because some DB seeds drop leading hex zeros while the generator names its
// package directory with a full 64-hex-char seed, the tool also looks up a
// zero-padded variant and reports any token found only via the padded name.
//
// Build from repo root:
//
//	go build -o rwcg/tools/asset_inventory ./rwcg/tools/asset_inventory.go
//
// Run from the asset directory (e.g. ~/nft-assets/new/cosmicsignature):
//
//	./asset_inventory
//	./asset_inventory -db 'dbname=cgprod sslmode=disable' -base . -missing-only
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/rwcg/tools/toolutil"
	_ "github.com/lib/pq"
)

const seedHexLen = 64 // 32-byte seed as lowercase hex (no 0x prefix)

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

// imageLayoutOrder / videoLayoutOrder give a stable display order for totals.
var (
	imageLayoutOrder   = []string{"web", "source", "legacy", "flat"}
	previewLayoutOrder = []string{"web", "thumb"}
	videoLayoutOrder   = []string{"web", "hq", "legacy", "flat"}
)

type tokenRow struct {
	tokenID int64
	seed    string // normalized: lowercase hex, no 0x prefix
}

func main() {
	dbConn := flag.String("db", "",
		"PostgreSQL connection string override. If empty, built from PGSQL_HOST, PGSQL_USERNAME, "+
			"PGSQL_DATABASE, PGSQL_PASSWORD (source your cosmic-etl-config.env first).")
	baseDir := flag.String("base", ".", "Base directory holding the per-seed assets (default: current directory)")
	schema := flag.String("schema", "public", "Database schema holding cg_mint_event")
	missingOnly := flag.Bool("missing-only", false, "Only list tokens with at least one missing asset")
	showAll := flag.Bool("all", false, "List every token with its image/video status")
	flag.Parse()

	absBase, err := filepath.Abs(*baseDir)
	if err != nil {
		log.Fatalf("resolve base dir %q: %v", *baseDir, err)
	}
	if st, err := os.Stat(absBase); err != nil || !st.IsDir() {
		log.Fatalf("base dir %q is not a directory: %v", absBase, err)
	}

	conn := *dbConn
	if conn == "" {
		conn, err = toolutil.PostgresConnStringFromEnv()
		if err != nil {
			log.Fatalf("no -db flag and %v", err)
		}
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(2)
	if err := db.Ping(); err != nil {
		log.Fatalf("ping database: %v", err)
	}

	tokens, err := fetchTokens(db, *schema)
	if err != nil {
		log.Fatalf("fetch token seeds: %v", err)
	}

	fmt.Println("Cosmic Signature asset inventory")
	fmt.Printf("  base dir : %s\n", absBase)
	fmt.Printf("  database : %s\n", redactConn(conn))
	fmt.Printf("  schema   : %s\n", *schema)
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
				fmt.Sprintf("  #%-5d db=0x%s  file=0x%s", tk.tokenID, tk.seed, leftPad(tk.seed)))
		}

		if *showAll {
			allLines = append(allLines,
				fmt.Sprintf("  #%-5d 0x%s  image=%s  preview=%s  video=%s", tk.tokenID, tk.seed,
					status(img), status(prev), status(vid)))
		}
	}

	if *showAll {
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
	} else if !*missingOnly {
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

	if *missingOnly {
		return
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

// fetchTokens returns one row per minted token (deduped by seed), normalized.
func fetchTokens(db *sql.DB, schema string) ([]tokenRow, error) {
	query := fmt.Sprintf(
		"SELECT token_id, seed FROM %s.cg_mint_event WHERE seed IS NOT NULL AND seed <> '' ORDER BY token_id",
		schema,
	)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	seen := make(map[string]bool)
	var out []tokenRow
	for rows.Next() {
		var tokenID int64
		var seed string
		if err := rows.Scan(&tokenID, &seed); err != nil {
			return nil, err
		}
		seed = normalizeSeed(seed)
		if seed == "" || seen[seed] {
			continue
		}
		seen[seed] = true
		out = append(out, tokenRow{tokenID: tokenID, seed: seed})
	}
	return out, rows.Err()
}

// normalizeSeed lowercases and strips an optional 0x prefix.
func normalizeSeed(seed string) string {
	seed = strings.ToLower(strings.TrimSpace(seed))
	seed = strings.TrimPrefix(seed, "0x")
	return seed
}

// leftPad returns the seed zero-padded to the full 64-hex-char width.
func leftPad(seed string) string {
	if len(seed) >= seedHexLen {
		return seed
	}
	return strings.Repeat("0", seedHexLen-len(seed)) + seed
}

// seedNames returns candidate "0x<seed>" directory/file basenames: the raw DB
// seed and, when shorter than 64 hex chars, the zero-padded variant.
func seedNames(seed string) []struct {
	name   string
	padded bool
} {
	out := []struct {
		name   string
		padded bool
	}{{name: "0x" + seed, padded: false}}
	if padded := leftPad(seed); padded != seed {
		out = append(out, struct {
			name   string
			padded bool
		}{name: "0x" + padded, padded: true})
	}
	return out
}

// locateAsset returns the first on-disk match across all candidate layouts,
// trying the exact seed name before the zero-padded variant. kind is "image",
// "preview" or "video".
func locateAsset(base, seed, kind string) assetLoc {
	cands := candidatesFor(kind)
	for _, c := range seedNames(seed) {
		for _, t := range cands {
			var p string
			if t.inner != "" {
				p = filepath.Join(base, c.name, filepath.FromSlash(t.inner))
			} else {
				p = filepath.Join(base, c.name+t.flatExt)
			}
			if isFile(p) {
				return assetLoc{found: true, path: p, layout: t.layout, padded: c.padded}
			}
		}
	}
	return assetLoc{found: false}
}

func isFile(path string) bool {
	st, err := os.Stat(path)
	return err == nil && !st.IsDir()
}

func status(a assetLoc) string {
	if !a.found {
		return "MISSING"
	}
	if a.padded {
		return a.layout + "(padded)"
	}
	return a.layout
}
