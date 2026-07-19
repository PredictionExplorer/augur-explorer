package main

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	opsassets "github.com/PredictionExplorer/augur-explorer/internal/ops/assets"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

// assetsMaxConns bounds the asset command pools; the token source runs one
// query.
const assetsMaxConns = 2

type assetsInventoryDeps struct {
	postgresConn func() (string, error)
	openDB       func(context.Context, string) (opsDB, error)
	ping         func(context.Context, opsDB) error
	newSource    func(opsDB) opsassets.TokenSource
	run          func(context.Context, opsassets.InventoryOptions) (opsassets.InventorySummary, error)
}

func defaultAssetsInventoryDeps() assetsInventoryDeps {
	return assetsInventoryDeps{
		postgresConn: toolutil.PostgresConnStringFromEnv,
		openDB:       openOpsDB(assetsMaxConns),
		ping: func(ctx context.Context, db opsDB) error {
			return db.Ping(ctx)
		},
		newSource: func(db opsDB) opsassets.TokenSource {
			return opsassets.SQLTokenSource{DB: db}
		},
		run: opsassets.RunInventory,
	}
}

// newAssetsInventoryCmd builds `opsctl assets inventory`, the replacement for
// the standalone asset_inventory tool.
func newAssetsInventoryCmd() *cobra.Command {
	return newAssetsInventoryCmdWithDeps(defaultAssetsInventoryDeps())
}

func newAssetsInventoryCmdWithDeps(deps assetsInventoryDeps) *cobra.Command {
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
			return runAssetInventoryWithDeps(cmd, dbConn, baseDir, schema, missingOnly, showAll, deps)
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

func runAssetInventory(cmd *cobra.Command, dbConn, baseDir, schema string, missingOnly, showAll bool) error {
	return runAssetInventoryWithDeps(
		cmd,
		dbConn,
		baseDir,
		schema,
		missingOnly,
		showAll,
		defaultAssetsInventoryDeps(),
	)
}

func runAssetInventoryWithDeps(
	cmd *cobra.Command,
	dbConn, baseDir, schema string,
	missingOnly, showAll bool,
	deps assetsInventoryDeps,
) error {
	if err := opsassets.ValidateSchema(schema); err != nil {
		return err
	}
	conn := dbConn
	if conn == "" {
		var err error
		conn, err = deps.postgresConn()
		if err != nil {
			return fmt.Errorf("no --db flag and %w", err)
		}
	}

	db, err := deps.openDB(cmd.Context(), conn)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer db.Close()
	if err := deps.ping(cmd.Context(), db); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	_, err = deps.run(cmd.Context(), opsassets.InventoryOptions{
		Source:      deps.newSource(db),
		BaseDir:     baseDir,
		Schema:      schema,
		Database:    redactConn(conn),
		MissingOnly: missingOnly,
		ShowAll:     showAll,
		Output:      cmd.OutOrStdout(),
	})
	return err
}

var keywordPasswordRe = regexp.MustCompile(`password\s*=\s*('[^']*'|"[^"]*"|\S+)`)

// redactConn hides credentials in database and RPC URLs as well as PostgreSQL
// keyword connection strings.
func redactConn(conn string) string {
	if u, err := url.Parse(conn); err == nil && u.Scheme != "" && u.Host != "" {
		changed := false
		if _, hasPW := u.User.Password(); hasPW {
			u.User = url.UserPassword(u.User.Username(), "***")
			changed = true
		}
		query := u.Query()
		for key := range query {
			if sensitiveURLQueryKey(key) {
				query.Set(key, "***")
				changed = true
			}
		}
		if changed {
			u.RawQuery = query.Encode()
			return u.String()
		}
		return conn
	}
	return keywordPasswordRe.ReplaceAllString(conn, "password=***")
}

func sensitiveURLQueryKey(key string) bool {
	switch strings.ToLower(key) {
	case "access_token", "apikey", "api_key", "key", "pass", "passwd", "password", "secret", "token":
		return true
	default:
		return false
	}
}

func redactRPCURL(rawURL string) string {
	redacted := redactConn(rawURL)
	parsed, err := url.Parse(redacted)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return redacted
	}
	if parsed.Path != "" && parsed.Path != "/" {
		parsed.Path = "/<redacted>"
		parsed.RawPath = ""
	}
	return parsed.String()
}
