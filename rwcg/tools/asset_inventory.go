// Cosmic Signature image/video asset inventory.
//
// Connects to the SQL database, fetches every minted Cosmic Signature token
// seed, then checks the local filesystem (current directory as base by default)
// for the presence of each token's image and video. The "new" per-seed package
// layout is the standard:
//
//	0x<seed>/image.png
//	0x<seed>/video.mp4
//
// The legacy flat layout is also accepted as "present" (the web server serves
// either), so the missing counts reflect assets that are genuinely not on disk:
//
//	0x<seed>.png
//	0x<seed>.mp4
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
	layout string // "package" | "flat"
	padded bool   // matched only via the zero-padded 64-char seed name
}

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
	fmt.Printf("  layout   : package 0x<seed>/{image.png,video.mp4}; flat fallback 0x<seed>.{png,mp4}\n")
	fmt.Printf("  DB seeds : %d\n\n", len(tokens))

	var (
		imgPresent, imgMissing   int
		vidPresent, vidMissing   int
		imgPkg, imgFlat          int
		vidPkg, vidFlat          int
		bothPresent, bothMissing int
		partial                  int
		missingLines             []string
		paddedLines              []string
		allLines                 []string
	)

	for _, tk := range tokens {
		img := locateAsset(absBase, tk.seed, "image")
		vid := locateAsset(absBase, tk.seed, "video")

		if img.found {
			imgPresent++
			if img.layout == "package" {
				imgPkg++
			} else {
				imgFlat++
			}
		} else {
			imgMissing++
		}
		if vid.found {
			vidPresent++
			if vid.layout == "package" {
				vidPkg++
			} else {
				vidFlat++
			}
		} else {
			vidMissing++
		}

		switch {
		case img.found && vid.found:
			bothPresent++
		case !img.found && !vid.found:
			bothMissing++
		default:
			partial++
		}

		if !img.found || !vid.found {
			var miss []string
			if !img.found {
				miss = append(miss, "image")
			}
			if !vid.found {
				miss = append(miss, "video")
			}
			missingLines = append(missingLines,
				fmt.Sprintf("  #%-5d 0x%s  missing: %s", tk.tokenID, tk.seed, strings.Join(miss, ",")))
		}

		if (img.found && img.padded) || (vid.found && vid.padded) {
			paddedLines = append(paddedLines,
				fmt.Sprintf("  #%-5d db=0x%s  file=0x%s", tk.tokenID, tk.seed, leftPad(tk.seed)))
		}

		if *showAll {
			allLines = append(allLines,
				fmt.Sprintf("  #%-5d 0x%s  image=%s  video=%s", tk.tokenID, tk.seed,
					status(img), status(vid)))
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
		fmt.Println("No missing assets. Every DB seed has an image and a video on disk.")
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
	fmt.Printf("Images present    : %-5d (package %d, flat %d)\n", imgPresent, imgPkg, imgFlat)
	fmt.Printf("Images missing    : %d\n", imgMissing)
	fmt.Printf("Videos present    : %-5d (package %d, flat %d)\n", vidPresent, vidPkg, vidFlat)
	fmt.Printf("Videos missing    : %d\n", vidMissing)
	fmt.Printf("Both present      : %d\n", bothPresent)
	fmt.Printf("Both missing      : %d\n", bothMissing)
	fmt.Printf("Partial (one only): %d\n", partial)
	if len(paddedLines) > 0 {
		fmt.Printf("Padded-name only  : %d\n", len(paddedLines))
	}
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

// locateAsset checks the package layout first, then the flat fallback, for each
// candidate seed name. kind is "image" or "video".
func locateAsset(base, seed, kind string) assetLoc {
	var pkgInner, flatExt string
	if kind == "video" {
		pkgInner, flatExt = "video.mp4", ".mp4"
	} else {
		pkgInner, flatExt = "image.png", ".png"
	}

	for _, c := range seedNames(seed) {
		if p := filepath.Join(base, c.name, pkgInner); isFile(p) {
			return assetLoc{found: true, path: p, layout: "package", padded: c.padded}
		}
		if f := filepath.Join(base, c.name+flatExt); isFile(f) {
			return assetLoc{found: true, path: f, layout: "flat", padded: c.padded}
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
