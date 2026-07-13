package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/lib/pq" // postgres driver
	"github.com/spf13/cobra"

	opsassets "github.com/PredictionExplorer/augur-explorer/internal/ops/assets"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

type clockFunc func() time.Time

func (f clockFunc) Now() time.Time { return f() }

type assetsThumbnailsDeps struct {
	getenv        func(string) string
	postgresConn  func() (string, error)
	resolveMagick func(string) (string, error)
	openDB        func(string, string) (*sql.DB, error)
	ping          func(context.Context, *sql.DB) error
	newSource     func(*sql.DB) opsassets.TokenSource
	runner        opsassets.CommandRunner
	clock         opsassets.Clock
	generate      func(context.Context, opsassets.ThumbnailOptions) (opsassets.ThumbnailSummary, error)
}

func defaultAssetsThumbnailsDeps() assetsThumbnailsDeps {
	return assetsThumbnailsDeps{
		getenv:        os.Getenv,
		postgresConn:  toolutil.PostgresConnStringFromEnv,
		resolveMagick: resolveMagick,
		openDB:        sql.Open,
		ping: func(ctx context.Context, db *sql.DB) error {
			return db.PingContext(ctx)
		},
		newSource: func(db *sql.DB) opsassets.TokenSource {
			return opsassets.SQLTokenSource{DB: db}
		},
		runner:   opsassets.ExecCommandRunner{},
		clock:    clockFunc(time.Now),
		generate: opsassets.GenerateThumbnails,
	}
}

// newAssetsGenThumbnailsCmd builds `opsctl assets gen-thumbnails`, the
// replacement for the standalone gen_thumbnails tool.
func newAssetsGenThumbnailsCmd() *cobra.Command {
	return newAssetsGenThumbnailsCmdWithDeps(defaultAssetsThumbnailsDeps())
}

func newAssetsGenThumbnailsCmdWithDeps(deps assetsThumbnailsDeps) *cobra.Command {
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
			return runGenThumbnailsWithDeps(cmd, dbConn, baseDir, schema, force, magickBin, minAge, deps)
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

func runGenThumbnails(cmd *cobra.Command, dbConn, baseDir, schema string, force bool, magickBin string, minAge int) error {
	return runGenThumbnailsWithDeps(
		cmd,
		dbConn,
		baseDir,
		schema,
		force,
		magickBin,
		minAge,
		defaultAssetsThumbnailsDeps(),
	)
}

func runGenThumbnailsWithDeps(
	cmd *cobra.Command,
	dbConn, baseDir, schema string,
	force bool,
	magickBin string,
	minAge int,
	deps assetsThumbnailsDeps,
) error {
	if minAge < 0 {
		return errors.New("--min-age must be non-negative")
	}
	base := baseDir
	if base == "" {
		root := strings.TrimSpace(deps.getenv("NFT_ASSETS_ROOT"))
		if root == "" {
			return errors.New("variable NFT_ASSETS_ROOT isn't set")
		}
		base = filepath.Join(root, "new", "cosmicsignature")
	}
	if err := opsassets.ValidateSchema(schema); err != nil {
		return err
	}

	magickPath, err := deps.resolveMagick(magickBin)
	if err != nil {
		return err
	}

	conn := dbConn
	if conn == "" {
		conn, err = deps.postgresConn()
		if err != nil {
			return fmt.Errorf("no --db flag and %w", err)
		}
	}

	db, err := deps.openDB("postgres", conn)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer func() { _ = db.Close() }()
	db.SetMaxOpenConns(2)
	if err := deps.ping(cmd.Context(), db); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	logger := log.New(cmd.ErrOrStderr(), "", log.LstdFlags)
	_, err = deps.generate(cmd.Context(), opsassets.ThumbnailOptions{
		Source:     deps.newSource(db),
		BaseDir:    base,
		Schema:     schema,
		Force:      force,
		MagickPath: magickPath,
		MinAge:     time.Duration(minAge) * time.Second,
		Runner:     deps.runner,
		Logger:     logger,
		Clock:      deps.clock,
	})
	return err
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
