package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
	_ "github.com/lib/pq" // postgres driver
	"github.com/spf13/cobra"
)

// thumbSpec describes one thumbnail output size.
type thumbSpec struct {
	name    string // output file (without dir), e.g. "thumb_card.webp"
	edge    int    // long-edge box (preserves aspect ratio)
	quality int
	unsharp string
}

// thumbSpecs are the two WebP thumbnails the frontend uses to keep pages
// light: a 640px gallery card and a 160px table/list micro thumb.
var thumbSpecs = []thumbSpec{
	{name: "thumb_card.webp", edge: 640, quality: 82, unsharp: "0x0.8+0.8+0.005"},
	{name: "thumb_micro.webp", edge: 160, quality: 80, unsharp: "0x0.6+0.7+0.003"},
}

// thumbCounters accumulates per-run results.
type thumbCounters struct {
	generated  int
	skipped    int // up to date
	srcMissing int // no image.png yet (not generated/uploaded)
	tooFresh   int // image.png still being written
	failed     int
}

// newAssetsGenThumbnailsCmd builds `opsctl assets gen-thumbnails`, the
// replacement for the standalone gen_thumbnails tool.
func newAssetsGenThumbnailsCmd() *cobra.Command {
	var (
		dbConn    string
		baseDir   string
		schema    string
		force     bool
		magickBin string
		minAge    int
	)
	cmd := &cobra.Command{
		Use:   "gen-thumbnails",
		Short: "Generate WebP thumbnails for Cosmic Signature per-seed packages",
		Long: `Fetches every minted token seed from the database, then for each on-disk
per-seed package directory (0x<seed>/image.png) generates two WebP thumbnails:

	0x<seed>/thumb_card.webp   640px long edge  (gallery grid, home previews)
	0x<seed>/thumb_micro.webp  160px long edge  (table cells, list rows)

The source art is bright thin lines on black, so the resize is done in linear
light to keep the lines from darkening, followed by a small saturation lift
and an unsharp pass. By default a thumbnail is (re)generated only when it is
missing or older than image.png; pass --force to rewrite everything.

Requires ImageMagick with the WebP delegate: either 7 ('magick') or 6
('convert'); the binary is auto-detected, override with --magick. The asset
directory is $NFT_ASSETS_ROOT/new/cosmicsignature by default (override with
--base). Without --db the connection is built from the PGSQL_* environment
variables.

Cron (every minute) — wrap in flock so runs never overlap:

	* * * * * flock -n /tmp/cs_thumbs.lock nice -n 19 ionice -c3 \
	  bash -lc 'source ~/configs/cosmic-api-config.env && \
	  opsctl assets gen-thumbnails >> ~/ae_logs/thumbnails.log 2>&1'`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenThumbnails(dbConn, baseDir, schema, force, magickBin, minAge)
		},
	}
	cmd.Flags().StringVar(&dbConn, "db", "",
		"PostgreSQL connection string override. If empty, built from PGSQL_HOST, PGSQL_USERNAME, "+
			"PGSQL_DATABASE, PGSQL_PASSWORD.")
	cmd.Flags().StringVar(&baseDir, "base", "", "Override the asset directory. If empty, uses $NFT_ASSETS_ROOT/new/cosmicsignature")
	cmd.Flags().StringVar(&schema, "schema", "public", "Database schema holding cg_mint_event")
	cmd.Flags().BoolVar(&force, "force", false, "Regenerate every thumbnail even if it already exists and is up to date")
	cmd.Flags().StringVar(&magickBin, "magick", "", "ImageMagick binary to use (default: auto-detect 'magick' (IM7) then 'convert' (IM6))")
	cmd.Flags().IntVar(&minAge, "min-age", 10, "Skip a token if image.png was modified within this many seconds (avoids reading a half-uploaded file)")
	return cmd
}

func init() { assetsCmd.AddCommand(newAssetsGenThumbnailsCmd()) }

func runGenThumbnails(dbConn, baseDir, schema string, force bool, magickBin string, minAge int) error {
	base := baseDir
	if base == "" {
		root := strings.TrimSpace(os.Getenv("NFT_ASSETS_ROOT"))
		if root == "" {
			return errors.New("variable NFT_ASSETS_ROOT isn't set")
		}
		base = filepath.Join(root, "new", "cosmicsignature")
	}
	absBase, err := filepath.Abs(base)
	if err != nil {
		return fmt.Errorf("resolve base dir %q: %w", base, err)
	}
	if st, err := os.Stat(absBase); err != nil || !st.IsDir() {
		return fmt.Errorf("base dir %q is not a directory: %v", absBase, err)
	}

	magickPath, err := resolveMagick(magickBin)
	if err != nil {
		return err
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

	log.Printf("thumbnails: base=%s seeds=%d force=%v magick=%s", absBase, len(tokens), force, magickPath)

	var c thumbCounters
	freshCutoff := time.Now().Add(-time.Duration(minAge) * time.Second)

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

		for _, spec := range thumbSpecs {
			dst := filepath.Join(dir, spec.name)
			if !force && upToDate(dst, srcInfo.ModTime()) {
				c.skipped++
				continue
			}
			if err := generateThumbnail(magickPath, src, dst, spec); err != nil {
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
		return fmt.Errorf("%d thumbnail(s) failed to generate", c.failed)
	}
	return nil
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
	return "", errors.New("neither 'magick' (ImageMagick 7) nor 'convert' (ImageMagick 6) found on PATH")
}

// generateThumbnail runs ImageMagick to produce one thumbnail, writing
// atomically via a temp file.
func generateThumbnail(magickPath, src, dst string, spec thumbSpec) error {
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
	for _, c := range seedNameCandidates(seed) {
		d := filepath.Join(base, c.name)
		if isRegularFile(filepath.Join(d, "image.png")) {
			return d, true
		}
	}
	return "", false
}
