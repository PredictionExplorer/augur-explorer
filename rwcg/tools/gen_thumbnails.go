// Cosmic Signature thumbnail generator.
//
// Fetches every minted token seed from the database, then for each on-disk
// per-seed package directory (0x<seed>/image.png) generates two WebP
// thumbnails used by the frontend to keep pages light:
//
//	0x<seed>/thumb_card.webp   640px long edge  (gallery grid, home previews)
//	0x<seed>/thumb_micro.webp  160px long edge  (table cells, list rows)
//
// The source art is bright thin lines on black, so the resize is done in
// linear light (-colorspace RGB ... -colorspace sRGB) to keep the lines from
// darkening, followed by a small saturation lift and an unsharp pass.
//
// By default a thumbnail is (re)generated only when it is missing or older than
// image.png; pass -force to rewrite everything.
//
// Requires ImageMagick with the WebP delegate: either 7 (`magick`) or
// 6 (`convert`). The binary is auto-detected; override with -magick.
//
// Build from repo root:
//
//	go build -o rwcg/tools/gen_thumbnails ./rwcg/tools/gen_thumbnails.go
//
// The asset directory is $NFT_ASSETS_ROOT/new/cosmicsignature by default
// (override with -base). Source the API env file first so both the PGSQL_*
// variables and NFT_ASSETS_ROOT are set:
//
//	source ~/configs/cosmic-api-config.env
//	./gen_thumbnails
//	./gen_thumbnails -force
//	./gen_thumbnails -base /home/frontend/nft-assets/new/cosmicsignature
//
// Cron (every minute) — wrap in flock so runs never overlap:
//
//   - * * * * flock -n /tmp/cs_thumbs.lock nice -n 19 ionice -c3 \
//     bash -lc 'source ~/configs/cosmic-api-config.env && \
//     ~/backend/rwcg/tools/gen_thumbnails >> ~/ae_logs/thumbnails.log 2>&1'
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/rwcg/tools/toolutil"
	_ "github.com/lib/pq"
)

const seedHexLen = 64 // 32-byte seed as lowercase hex (no 0x prefix)

// thumbSpec describes one output size.
type thumbSpec struct {
	name    string // output file (without dir), e.g. "thumb_card.webp"
	edge    int    // long-edge box (preserves aspect ratio)
	quality int
	unsharp string
}

var specs = []thumbSpec{
	{name: "thumb_card.webp", edge: 640, quality: 82, unsharp: "0x0.8+0.8+0.005"},
	{name: "thumb_micro.webp", edge: 160, quality: 80, unsharp: "0x0.6+0.7+0.003"},
}

type tokenRow struct {
	tokenID int64
	seed    string // normalized: lowercase hex, no 0x prefix
}

type counters struct {
	generated  int
	skipped    int // up to date
	srcMissing int // no image.png yet (not generated/uploaded)
	tooFresh   int // image.png still being written
	failed     int
}

func main() {
	dbConn := flag.String("db", "",
		"PostgreSQL connection string override. If empty, built from PGSQL_HOST, PGSQL_USERNAME, "+
			"PGSQL_DATABASE, PGSQL_PASSWORD.")
	baseDir := flag.String("base", "", "Override the asset directory. If empty, uses $NFT_ASSETS_ROOT/new/cosmicsignature")
	schema := flag.String("schema", "public", "Database schema holding cg_mint_event")
	force := flag.Bool("force", false, "Regenerate every thumbnail even if it already exists and is up to date")
	magickBin := flag.String("magick", "", "ImageMagick binary to use (default: auto-detect 'magick' (IM7) then 'convert' (IM6))")
	minAge := flag.Int("min-age", 10, "Skip a token if image.png was modified within this many seconds (avoids reading a half-uploaded file)")
	flag.Parse()

	base := *baseDir
	if base == "" {
		root := strings.TrimSpace(os.Getenv("NFT_ASSETS_ROOT"))
		if root == "" {
			log.Fatal("variable NFT_ASSETS_ROOT isn't set")
		}
		base = filepath.Join(root, "new", "cosmicsignature")
	}
	absBase, err := filepath.Abs(base)
	if err != nil {
		log.Fatalf("resolve base dir %q: %v", base, err)
	}
	if st, err := os.Stat(absBase); err != nil || !st.IsDir() {
		log.Fatalf("base dir %q is not a directory: %v", absBase, err)
	}

	magickPath, err := resolveMagick(*magickBin)
	if err != nil {
		log.Fatal(err)
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

	log.Printf("thumbnails: base=%s seeds=%d force=%v magick=%s", absBase, len(tokens), *force, magickPath)

	var c counters
	freshCutoff := time.Now().Add(-time.Duration(*minAge) * time.Second)

	for _, tk := range tokens {
		dir, ok := seedDir(absBase, tk.seed)
		if !ok {
			c.srcMissing++
			continue
		}
		src := filepath.Join(dir, "image.png")

		srcInfo, err := os.Stat(src)
		if err != nil {
			c.srcMissing++
			continue
		}
		if srcInfo.ModTime().After(freshCutoff) {
			// image.png may still be uploading; try again next run.
			c.tooFresh++
			continue
		}

		for _, spec := range specs {
			dst := filepath.Join(dir, spec.name)
			if !*force && upToDate(dst, srcInfo.ModTime()) {
				c.skipped++
				continue
			}
			if err := generate(magickPath, src, dst, spec); err != nil {
				c.failed++
				log.Printf("FAIL %s: %v", dst, err)
				continue
			}
			c.generated++
			log.Printf("OK   %s", dst)
		}
	}

	log.Printf("done: generated=%d skipped=%d src-missing=%d too-fresh=%d failed=%d",
		c.generated, c.skipped, c.srcMissing, c.tooFresh, c.failed)

	if c.failed > 0 {
		os.Exit(1)
	}
}

// resolveMagick returns the ImageMagick binary path. If explicit is empty it
// auto-detects 'magick' (IM7) then 'convert' (IM6); both accept the same
// input-first argument order used here.
func resolveMagick(explicit string) (string, error) {
	if explicit != "" {
		path, err := exec.LookPath(explicit)
		if err != nil {
			return "", fmt.Errorf("ImageMagick %q not found on PATH: %w", explicit, err)
		}
		return path, nil
	}
	for _, cand := range []string{"magick", "convert"} {
		if path, err := exec.LookPath(cand); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("neither 'magick' (ImageMagick 7) nor 'convert' (ImageMagick 6) found on PATH")
}

// generate runs ImageMagick to produce one thumbnail, writing atomically.
func generate(magickPath, src, dst string, spec thumbSpec) error {
	tmp := dst + ".tmp"
	args := []string{
		src,
		"-strip",
		"-colorspace", "RGB", // linear light: keeps bright-on-black lines from darkening
		"-filter", "Lanczos",
		"-resize", fmt.Sprintf("%dx%d", spec.edge, spec.edge), // fits within box, preserves aspect
		"-colorspace", "sRGB",
		"-modulate", "100,112,100", // +12% saturation so colors stay vivid when small
		"-unsharp", spec.unsharp,
		"-quality", fmt.Sprintf("%d", spec.quality),
		"-define", "webp:method=6",
		tmp,
	}
	out, err := exec.Command(magickPath, args...).CombinedOutput()
	if err != nil {
		_ = os.Remove(tmp)
		msg := strings.TrimSpace(string(out))
		if msg == "" {
			return err
		}
		return fmt.Errorf("%v: %s", err, msg)
	}
	return os.Rename(tmp, dst)
}

// upToDate reports whether dst exists and is at least as new as the source.
func upToDate(dst string, srcMod time.Time) bool {
	st, err := os.Stat(dst)
	if err != nil {
		return false
	}
	return !st.ModTime().Before(srcMod)
}

// seedDir returns the per-seed package directory that holds image.png, trying
// the raw DB seed name and the zero-padded 64-char variant.
func seedDir(base, seed string) (string, bool) {
	for _, name := range seedNames(seed) {
		d := filepath.Join(base, name)
		if isFile(filepath.Join(d, "image.png")) {
			return d, true
		}
	}
	return "", false
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

func normalizeSeed(seed string) string {
	seed = strings.ToLower(strings.TrimSpace(seed))
	return strings.TrimPrefix(seed, "0x")
}

func leftPad(seed string) string {
	if len(seed) >= seedHexLen {
		return seed
	}
	return strings.Repeat("0", seedHexLen-len(seed)) + seed
}

// seedNames returns candidate "0x<seed>" directory basenames: the raw DB seed
// and, when shorter than 64 hex chars, the zero-padded variant.
func seedNames(seed string) []string {
	out := []string{"0x" + seed}
	if padded := leftPad(seed); padded != seed {
		out = append(out, "0x"+padded)
	}
	return out
}

func isFile(path string) bool {
	st, err := os.Stat(path)
	return err == nil && !st.IsDir()
}
