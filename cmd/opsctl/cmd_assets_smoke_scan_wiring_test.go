package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	opsassets "github.com/PredictionExplorer/augur-explorer/internal/ops/assets"
	"github.com/PredictionExplorer/augur-explorer/internal/ops/cstscan"
	"github.com/PredictionExplorer/augur-explorer/internal/ops/smoketest"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/ethereum/go-ethereum/common"
)

func TestAssetsInventoryCommandWiring(t *testing.T) {
	db := newCommandTestDB(t)
	source := tokenSourceFunc(func(context.Context, string) ([]opsassets.Token, error) {
		return nil, nil
	})
	deps := defaultAssetsInventoryDeps()
	deps.postgresConn = func() (string, error) {
		t.Fatal("environment connection used with --db")
		return "", nil
	}
	deps.openDB = func(driverName, conn string) (*sql.DB, error) {
		// #nosec G101 -- deliberately fake URL verifies password redaction.
		if driverName != "postgres" || conn != "postgres://user:secret@db.example/app" {
			t.Fatalf("open = %q/%q", driverName, conn)
		}
		return db, nil
	}
	deps.ping = func(ctx context.Context, gotDB *sql.DB) error {
		if err := ctx.Err(); err != nil || gotDB != db {
			t.Fatalf("ping = %v/%p", err, gotDB)
		}
		return nil
	}
	deps.newSource = func(gotDB *sql.DB) opsassets.TokenSource {
		if gotDB != db {
			t.Fatal("wrong inventory source database")
		}
		return source
	}
	deps.run = func(ctx context.Context, options opsassets.InventoryOptions) (opsassets.InventorySummary, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if _, ok := options.Source.(tokenSourceFunc); !ok {
			t.Fatalf("source = %T", options.Source)
		}
		if options.BaseDir != "/asset/base" ||
			options.Schema != "tenant_7" || !options.MissingOnly || !options.ShowAll {
			t.Fatalf("options = %#v", options)
		}
		// #nosec G101 -- deliberately fake URL is the expected redacted value.
		if options.Database != "postgres://user:%2A%2A%2A@db.example/app" {
			t.Fatalf("redacted database = %q", options.Database)
		}
		fmt.Fprintln(options.Output, "inventory wiring complete")
		return opsassets.InventorySummary{DBSeeds: 2}, nil
	}
	result := executeCommand(
		newAssetsInventoryCmdWithDeps(deps),
		"--db", "postgres://user:secret@db.example/app",
		"--base", "/asset/base",
		"--schema", "tenant_7",
		"--missing-only",
		"--all",
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	if !strings.Contains(result.stdout, "inventory wiring complete") {
		t.Fatalf("stdout = %q", result.stdout)
	}
	if db.Stats().MaxOpenConnections != 2 {
		t.Fatalf("max connections = %d", db.Stats().MaxOpenConnections)
	}
	assertCommandDBClosed(t, db)
}

func TestAssetsInventoryCommandSetupAndErrors(t *testing.T) {
	t.Run("environment connection", func(t *testing.T) {
		db := newCommandTestDB(t)
		deps := defaultAssetsInventoryDeps()
		deps.postgresConn = func() (string, error) { return "environment-dsn", nil }
		deps.openDB = func(_ string, conn string) (*sql.DB, error) {
			if conn != "environment-dsn" {
				t.Fatalf("conn = %q", conn)
			}
			return db, nil
		}
		deps.ping = func(context.Context, *sql.DB) error { return nil }
		deps.newSource = func(*sql.DB) opsassets.TokenSource {
			return tokenSourceFunc(func(context.Context, string) ([]opsassets.Token, error) { return nil, nil })
		}
		deps.run = func(context.Context, opsassets.InventoryOptions) (opsassets.InventorySummary, error) {
			return opsassets.InventorySummary{}, nil
		}
		result := executeCommand(newAssetsInventoryCmdWithDeps(deps))
		if result.err != nil {
			t.Fatal(result.err)
		}
	})

	t.Run("invalid schema precedes setup", func(t *testing.T) {
		deps := defaultAssetsInventoryDeps()
		deps.postgresConn = func() (string, error) {
			t.Fatal("connection resolved after invalid schema")
			return "", nil
		}
		result := executeCommand(newAssetsInventoryCmdWithDeps(deps), "--schema", "bad.schema")
		if result.err == nil || !strings.Contains(result.err.Error(), "invalid database schema") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("environment error", func(t *testing.T) {
		want := errors.New("environment missing")
		deps := defaultAssetsInventoryDeps()
		deps.postgresConn = func() (string, error) { return "", want }
		result := executeCommand(newAssetsInventoryCmdWithDeps(deps))
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "no --db flag") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("open", func(t *testing.T) {
		want := errors.New("database unavailable")
		deps := defaultAssetsInventoryDeps()
		deps.openDB = func(string, string) (*sql.DB, error) { return nil, want }
		result := executeCommand(newAssetsInventoryCmdWithDeps(deps), "--db", "database")
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("ping", func(t *testing.T) {
		db := newCommandTestDB(t)
		want := errors.New("ping failed")
		deps := defaultAssetsInventoryDeps()
		deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
		deps.ping = func(context.Context, *sql.DB) error { return want }
		result := executeCommand(newAssetsInventoryCmdWithDeps(deps), "--db", "database")
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "ping database") {
			t.Fatalf("error = %v", result.err)
		}
		assertCommandDBClosed(t, db)
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "inventory", err: errors.New("inventory failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			db := newCommandTestDB(t)
			deps := defaultAssetsInventoryDeps()
			deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
			deps.ping = func(context.Context, *sql.DB) error { return nil }
			deps.newSource = func(*sql.DB) opsassets.TokenSource { return nil }
			deps.run = func(context.Context, opsassets.InventoryOptions) (opsassets.InventorySummary, error) {
				return opsassets.InventorySummary{}, test.err
			}
			result := executeCommand(newAssetsInventoryCmdWithDeps(deps), "--db", "database")
			if !errors.Is(result.err, test.err) {
				t.Fatalf("error = %v", result.err)
			}
		})
	}
}

func TestAssetsThumbnailsCommandWiring(t *testing.T) {
	db := newCommandTestDB(t)
	source := tokenSourceFunc(func(context.Context, string) ([]opsassets.Token, error) { return nil, nil })
	runner := commandRunnerFunc(func(context.Context, string, ...string) ([]byte, error) { return nil, nil })
	now := time.Unix(1_700_000_000, 0)
	clock := clockFunc(func() time.Time { return now })
	deps := defaultAssetsThumbnailsDeps()
	deps.getenv = func(name string) string {
		if name != "NFT_ASSETS_ROOT" {
			t.Fatalf("environment key = %q", name)
		}
		return " /srv/assets "
	}
	deps.postgresConn = func() (string, error) {
		t.Fatal("environment database used with --db")
		return "", nil
	}
	deps.resolveMagick = func(explicit string) (string, error) {
		if explicit != "custom-magick" {
			t.Fatalf("explicit magick = %q", explicit)
		}
		return "/opt/bin/magick", nil
	}
	deps.openDB = func(driverName, conn string) (*sql.DB, error) {
		if driverName != "postgres" || conn != "database" {
			t.Fatalf("open = %q/%q", driverName, conn)
		}
		return db, nil
	}
	deps.ping = func(context.Context, *sql.DB) error { return nil }
	deps.newSource = func(*sql.DB) opsassets.TokenSource { return source }
	deps.runner = runner
	deps.clock = clock
	deps.generate = func(
		ctx context.Context,
		options opsassets.ThumbnailOptions,
	) (opsassets.ThumbnailSummary, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if _, ok := options.Source.(tokenSourceFunc); !ok {
			t.Fatalf("source = %T", options.Source)
		}
		if options.BaseDir != filepath.Join("/srv/assets", "new", "cosmicsignature") ||
			options.Schema != "tenant_9" || !options.Force ||
			options.MagickPath != "/opt/bin/magick" || options.MinAge != 3*time.Second ||
			options.Clock.Now() != now {
			t.Fatalf("options = %#v", options)
		}
		if _, ok := options.Runner.(commandRunnerFunc); !ok {
			t.Fatalf("runner = %T", options.Runner)
		}
		options.Logger.Printf("thumbnail wiring complete")
		return opsassets.ThumbnailSummary{Generated: 2}, nil
	}
	result := executeCommand(
		newAssetsGenThumbnailsCmdWithDeps(deps),
		"--db", "database",
		"--schema", "tenant_9",
		"--force",
		"--magick", "custom-magick",
		"--min-age", "3",
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	if !strings.Contains(result.stderr, "thumbnail wiring complete") {
		t.Fatalf("stderr = %q", result.stderr)
	}
	if clock.Now() != now {
		t.Fatalf("clock = %v", clock.Now())
	}
	assertCommandDBClosed(t, db)
}

func TestAssetsThumbnailsCommandValidationSetupAndErrors(t *testing.T) {
	t.Run("negative minimum age", func(t *testing.T) {
		deps := defaultAssetsThumbnailsDeps()
		deps.getenv = func(string) string {
			t.Fatal("environment read")
			return ""
		}
		result := executeCommand(newAssetsGenThumbnailsCmdWithDeps(deps), "--min-age", "-1")
		if result.err == nil || !strings.Contains(result.err.Error(), "--min-age must be non-negative") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("missing asset root", func(t *testing.T) {
		deps := defaultAssetsThumbnailsDeps()
		deps.getenv = func(string) string { return " " }
		result := executeCommand(newAssetsGenThumbnailsCmdWithDeps(deps))
		if result.err == nil || !strings.Contains(result.err.Error(), "NFT_ASSETS_ROOT") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("invalid schema", func(t *testing.T) {
		deps := defaultAssetsThumbnailsDeps()
		deps.resolveMagick = func(string) (string, error) {
			t.Fatal("ImageMagick resolved")
			return "", nil
		}
		result := executeCommand(
			newAssetsGenThumbnailsCmdWithDeps(deps),
			"--base", t.TempDir(),
			"--schema", "bad.schema",
		)
		if result.err == nil || !strings.Contains(result.err.Error(), "invalid database schema") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("ImageMagick", func(t *testing.T) {
		want := errors.New("magick unavailable")
		deps := defaultAssetsThumbnailsDeps()
		deps.resolveMagick = func(string) (string, error) { return "", want }
		result := executeCommand(
			newAssetsGenThumbnailsCmdWithDeps(deps),
			"--base", t.TempDir(),
		)
		if !errors.Is(result.err, want) {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("environment database", func(t *testing.T) {
		want := errors.New("database environment missing")
		deps := defaultAssetsThumbnailsDeps()
		deps.resolveMagick = func(string) (string, error) { return "/magick", nil }
		deps.postgresConn = func() (string, error) { return "", want }
		result := executeCommand(
			newAssetsGenThumbnailsCmdWithDeps(deps),
			"--base", t.TempDir(),
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "no --db flag") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("open", func(t *testing.T) {
		want := errors.New("database unavailable")
		deps := defaultAssetsThumbnailsDeps()
		deps.resolveMagick = func(string) (string, error) { return "/magick", nil }
		deps.openDB = func(string, string) (*sql.DB, error) { return nil, want }
		result := executeCommand(
			newAssetsGenThumbnailsCmdWithDeps(deps),
			"--base", t.TempDir(),
			"--db", "database",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("ping", func(t *testing.T) {
		db := newCommandTestDB(t)
		want := errors.New("ping failed")
		deps := defaultAssetsThumbnailsDeps()
		deps.resolveMagick = func(string) (string, error) { return "/magick", nil }
		deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
		deps.ping = func(context.Context, *sql.DB) error { return want }
		result := executeCommand(
			newAssetsGenThumbnailsCmdWithDeps(deps),
			"--base", t.TempDir(),
			"--db", "database",
		)
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "ping database") {
			t.Fatalf("error = %v", result.err)
		}
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "generator", err: errors.New("generation failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			db := newCommandTestDB(t)
			deps := defaultAssetsThumbnailsDeps()
			deps.resolveMagick = func(string) (string, error) { return "/magick", nil }
			deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
			deps.ping = func(context.Context, *sql.DB) error { return nil }
			deps.newSource = func(*sql.DB) opsassets.TokenSource { return nil }
			deps.generate = func(context.Context, opsassets.ThumbnailOptions) (opsassets.ThumbnailSummary, error) {
				return opsassets.ThumbnailSummary{}, test.err
			}
			result := executeCommand(
				newAssetsGenThumbnailsCmdWithDeps(deps),
				"--base", t.TempDir(),
				"--db", "database",
			)
			if !errors.Is(result.err, test.err) {
				t.Fatalf("error = %v", result.err)
			}
		})
	}
}

func TestResolveMagickSearchOrder(t *testing.T) {
	binDir := t.TempDir()
	magick := filepath.Join(binDir, "magick")
	convert := filepath.Join(binDir, "convert")
	for _, path := range []string{magick, convert} {
		// #nosec G306 -- executable fixture must be runnable by exec.LookPath.
		if err := os.WriteFile(path, []byte("#!/bin/sh\nexit 0\n"), 0o750); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("PATH", binDir)

	got, err := resolveMagick("")
	if err != nil || got != magick {
		t.Fatalf("auto-detected path/error = %q/%v", got, err)
	}
	got, err = resolveMagick("convert")
	if err != nil || got != convert {
		t.Fatalf("explicit path/error = %q/%v", got, err)
	}
	if err := os.Remove(magick); err != nil {
		t.Fatal(err)
	}
	got, err = resolveMagick("")
	if err != nil || got != convert {
		t.Fatalf("fallback path/error = %q/%v", got, err)
	}
	if err := os.Remove(convert); err != nil {
		t.Fatal(err)
	}
	if _, err := resolveMagick(""); err == nil || !strings.Contains(err.Error(), "neither") {
		t.Fatalf("missing ImageMagick error = %v", err)
	}
}

func TestVerifyTokenImagesCommandWiring(t *testing.T) {
	source := tokenCountSourceFunc(func(context.Context) (int64, error) { return 2, nil })
	client := httpClientFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}, nil
	})
	var (
		closed    bool
		gotConfig store.Config
	)
	deps := defaultAssetsVerifyTokenImagesDeps()
	deps.storeConfig = func() store.Config {
		return store.Config{Host: "db:5432", User: "operator", Database: "explorer"}
	}
	deps.openStore = func(ctx context.Context, cfg store.Config) (*store.Store, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		gotConfig = cfg
		return nil, nil
	}
	deps.closeStore = func(*store.Store) { closed = true }
	deps.newSource = func(*store.Store) opsassets.TokenCountSource { return source }
	deps.client = client
	deps.baseURL = "https://images.example"
	deps.verify = func(
		ctx context.Context,
		options opsassets.VerifyTokenImagesOptions,
	) (opsassets.ImageVerificationSummary, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if _, ok := options.Source.(tokenCountSourceFunc); !ok {
			t.Fatalf("source = %T", options.Source)
		}
		if _, ok := options.Client.(httpClientFunc); !ok {
			t.Fatalf("client = %T", options.Client)
		}
		if options.BaseURL != "https://images.example" {
			t.Fatalf("options = %#v", options)
		}
		options.Logger.Printf("image verification wiring complete")
		return opsassets.ImageVerificationSummary{Tokens: 2, Checked: 2, OK: 2}, nil
	}
	result := executeCommand(newAssetsVerifyTokenImagesCmdWithDeps(deps))
	if result.err != nil {
		t.Fatal(result.err)
	}
	if gotConfig.Host != "db:5432" || !closed {
		t.Fatalf("config/closed = %#v/%v", gotConfig, closed)
	}
	if !strings.Contains(result.stdout, "image verification wiring complete") {
		t.Fatalf("stdout = %q", result.stdout)
	}
}

func TestVerifyTokenImagesCommandErrorsAndCancellation(t *testing.T) {
	t.Run("store open", func(t *testing.T) {
		want := errors.New("store unavailable")
		deps := defaultAssetsVerifyTokenImagesDeps()
		deps.openStore = func(context.Context, store.Config) (*store.Store, error) {
			return nil, want
		}
		result := executeCommand(newAssetsVerifyTokenImagesCmdWithDeps(deps))
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "connect to storage") {
			t.Fatalf("error = %v", result.err)
		}
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "verification", err: errors.New("HTTP verification failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			closed := false
			deps := defaultAssetsVerifyTokenImagesDeps()
			deps.openStore = func(context.Context, store.Config) (*store.Store, error) { return nil, nil }
			deps.closeStore = func(*store.Store) { closed = true }
			deps.newSource = func(*store.Store) opsassets.TokenCountSource { return nil }
			deps.verify = func(
				context.Context,
				opsassets.VerifyTokenImagesOptions,
			) (opsassets.ImageVerificationSummary, error) {
				return opsassets.ImageVerificationSummary{}, test.err
			}
			result := executeCommand(newAssetsVerifyTokenImagesCmdWithDeps(deps))
			if !errors.Is(result.err, test.err) || !closed {
				t.Fatalf("error/closed = %v/%v", result.err, closed)
			}
		})
	}
}

func TestRandomWalkTokenCountSource(t *testing.T) {
	t.Run("lookup error", func(t *testing.T) {
		want := errors.New("lookup failed")
		source := randomWalkTokenCountSource{
			lookupAddressID: func(context.Context, string) (int64, error) { return 0, want },
			mintedTokens: func(context.Context, int64) (int64, error) {
				t.Fatal("stats read after lookup failure")
				return 0, nil
			},
		}
		_, err := source.MintedTokenCount(context.Background())
		if !errors.Is(err, want) || !strings.Contains(err.Error(), "resolve RandomWalk") {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("stats error", func(t *testing.T) {
		want := errors.New("stats failed")
		source := randomWalkTokenCountSource{
			lookupAddressID: func(_ context.Context, address string) (int64, error) {
				if address != rwalkNFTAddr {
					t.Fatalf("address = %q", address)
				}
				return 77, nil
			},
			mintedTokens: func(_ context.Context, addressID int64) (int64, error) {
				if addressID != 77 {
					t.Fatalf("address ID = %d", addressID)
				}
				return 0, want
			},
		}
		_, err := source.MintedTokenCount(context.Background())
		if !errors.Is(err, want) || !strings.Contains(err.Error(), "read RandomWalk stats") {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("success", func(t *testing.T) {
		source := randomWalkTokenCountSource{
			lookupAddressID: func(context.Context, string) (int64, error) { return 4, nil },
			mintedTokens:    func(context.Context, int64) (int64, error) { return 123, nil },
		}
		got, err := source.MintedTokenCount(context.Background())
		if err != nil || got != 123 {
			t.Fatalf("count/error = %d/%v", got, err)
		}
	})
}

func TestSmoketestCommandWiring(t *testing.T) {
	db := newCommandTestDB(t)
	source := parameterSourceFunc(func(context.Context) (smoketest.Params, error) {
		return smoketest.DefaultParams(), nil
	})
	client := httpClientFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}, nil
	})
	values := map[string]string{
		"PGSQL_HOST":     "db.example:5432",
		"PGSQL_USERNAME": "operator",
		"PGSQL_DATABASE": "explorer",
		"PGSQL_PASSWORD": "p@ss word",
	}
	deps := defaultSmoketestDeps()
	deps.getenv = func(name string) string { return values[name] }
	deps.openDB = func(driverName, dsn string) (*sql.DB, error) {
		if driverName != "postgres" {
			t.Fatalf("driver = %q", driverName)
		}
		parsed, err := url.Parse(dsn)
		if err != nil {
			t.Fatal(err)
		}
		password, _ := parsed.User.Password()
		if parsed.Host != "db.example:5432" || parsed.User.Username() != "operator" ||
			password != "p@ss word" || parsed.Path != "/explorer" ||
			parsed.Query().Get("sslmode") != "disable" {
			t.Fatalf("DSN = %q", dsn)
		}
		return db, nil
	}
	deps.ping = func(context.Context, *sql.DB) error { return nil }
	deps.newSource = func(*sql.DB) smoketest.ParameterSource { return source }
	deps.client = client
	deps.apiBase = func() string { return "https://api.example/root" }
	deps.run = func(ctx context.Context, options smoketest.Options) (smoketest.Summary, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if _, ok := options.Source.(parameterSourceFunc); !ok {
			t.Fatalf("source = %T", options.Source)
		}
		if _, ok := options.Client.(httpClientFunc); !ok {
			t.Fatalf("client = %T", options.Client)
		}
		if options.BaseURL != "https://api.example/root" {
			t.Fatalf("options = %#v", options)
		}
		fmt.Fprintln(options.Output, "smoketest wiring complete")
		return smoketest.Summary{Total: 1, OK: 1}, nil
	}
	result := executeCommand(newSmoketestCmdWithDeps(deps))
	if result.err != nil {
		t.Fatal(result.err)
	}
	if !strings.Contains(result.stdout, "smoketest wiring complete") {
		t.Fatalf("stdout = %q", result.stdout)
	}
	assertCommandDBClosed(t, db)
}

func TestSmoketestCommandSetupAndErrors(t *testing.T) {
	requiredEnv := map[string]string{
		"PGSQL_HOST":     "db:5432",
		"PGSQL_USERNAME": "user",
		"PGSQL_DATABASE": "database",
		"PGSQL_PASSWORD": "password",
	}

	t.Run("missing environment", func(t *testing.T) {
		deps := defaultSmoketestDeps()
		deps.getenv = func(string) string { return "" }
		deps.openDB = func(string, string) (*sql.DB, error) {
			t.Fatal("database opened")
			return nil, nil
		}
		result := executeCommand(newSmoketestCmdWithDeps(deps))
		if result.err == nil || !strings.Contains(result.err.Error(), "PGSQL_HOST") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("open", func(t *testing.T) {
		want := errors.New("open failed")
		deps := defaultSmoketestDeps()
		deps.getenv = func(name string) string { return requiredEnv[name] }
		deps.openDB = func(string, string) (*sql.DB, error) { return nil, want }
		result := executeCommand(newSmoketestCmdWithDeps(deps))
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "DB open failed") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("ping", func(t *testing.T) {
		db := newCommandTestDB(t)
		want := errors.New("ping failed")
		deps := defaultSmoketestDeps()
		deps.getenv = func(name string) string { return requiredEnv[name] }
		deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
		deps.ping = func(context.Context, *sql.DB) error { return want }
		result := executeCommand(newSmoketestCmdWithDeps(deps))
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "DB ping failed") {
			t.Fatalf("error = %v", result.err)
		}
		assertCommandDBClosed(t, db)
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "runner", err: errors.New("endpoint failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			db := newCommandTestDB(t)
			deps := defaultSmoketestDeps()
			deps.getenv = func(name string) string { return requiredEnv[name] }
			deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
			deps.ping = func(context.Context, *sql.DB) error { return nil }
			deps.newSource = func(*sql.DB) smoketest.ParameterSource { return nil }
			deps.apiBase = func() string { return "http://api" }
			deps.run = func(context.Context, smoketest.Options) (smoketest.Summary, error) {
				return smoketest.Summary{}, test.err
			}
			result := executeCommand(newSmoketestCmdWithDeps(deps))
			if !errors.Is(result.err, test.err) {
				t.Fatalf("error = %v", result.err)
			}
		})
	}
}

func TestCstAuctionLengthCommandWiring(t *testing.T) {
	db := newCommandTestDB(t)
	rpc := &fakeOpsRPC{head: 99}
	keySource := cstscan.KeySourceFunc(func(context.Context) (map[cstscan.EventKey]struct{}, error) {
		return nil, nil
	})
	deps := defaultCstAuctionLenDeps()
	deps.getenv = func(name string) string {
		if name != "RPC_URL" {
			t.Fatalf("environment key = %q", name)
		}
		return "http://rpc.example"
	}
	deps.dialRPC = func(ctx context.Context, rpcURL string) (cstAuctionLenRPC, error) {
		if err := ctx.Err(); err != nil || rpcURL != "http://rpc.example" {
			t.Fatalf("dial = %v/%q", err, rpcURL)
		}
		return rpc, nil
	}
	deps.openDB = func(driverName, conn string) (*sql.DB, error) {
		if driverName != "postgres" || conn != "database" {
			t.Fatalf("open = %q/%q", driverName, conn)
		}
		return db, nil
	}
	deps.newKeySource = func(gotDB *sql.DB) cstscan.KeySource {
		if gotDB != db {
			t.Fatal("wrong key-source database")
		}
		return keySource
	}
	deps.scan = func(
		ctx context.Context,
		config cstscan.Config,
		options cstscan.Options,
	) (cstscan.Stats, error) {
		if err := ctx.Err(); err != nil {
			t.Fatal(err)
		}
		if config.Client != rpc ||
			config.Contract != common.HexToAddress("0x0000000000000000000000000000000000000002") ||
			config.Topic0 != common.HexToHash(cstAucLenTopic0Hex) {
			t.Fatalf("config = %#v", config)
		}
		if _, ok := config.KeySource.(cstscan.KeySourceFunc); !ok {
			t.Fatalf("key source = %T", config.KeySource)
		}
		if options.FromBlock != 12 || options.ToBlock != 44 ||
			options.InitialBatch != 500 || options.MinBatch != 500 ||
			options.RetryDelay != cstAucLenRetryDelay {
			t.Fatalf("options = %#v", options)
		}
		fmt.Fprintln(config.Output, "scan wiring complete")
		config.Logger.Printf("scan logger complete")
		return cstscan.Stats{Events: 1}, nil
	}
	result := executeCommand(
		newScanCstAuctionLenCmdWithDeps(deps),
		"--contract", "0x0000000000000000000000000000000000000002",
		"--from-block", "12",
		"--to-block", "44",
		"--batch", "500",
		"--db", "database",
	)
	if result.err != nil {
		t.Fatal(result.err)
	}
	if !strings.Contains(result.stdout, "scan wiring complete") ||
		!strings.Contains(result.stderr, "scan logger complete") {
		t.Fatalf("stdout/stderr = %q/%q", result.stdout, result.stderr)
	}
	if !rpc.closed.Load() {
		t.Fatal("RPC client was not closed")
	}
	assertCommandDBClosed(t, db)
}

func TestCstAuctionLengthCommandWithoutDatabase(t *testing.T) {
	rpc := &fakeOpsRPC{}
	deps := defaultCstAuctionLenDeps()
	deps.getenv = func(string) string { return "rpc" }
	deps.dialRPC = func(context.Context, string) (cstAuctionLenRPC, error) { return rpc, nil }
	deps.openDB = func(string, string) (*sql.DB, error) {
		t.Fatal("database opened without --db")
		return nil, nil
	}
	deps.scan = func(_ context.Context, config cstscan.Config, _ cstscan.Options) (cstscan.Stats, error) {
		if config.KeySource != nil {
			t.Fatalf("key source = %#v", config.KeySource)
		}
		return cstscan.Stats{}, nil
	}
	result := executeCommand(newScanCstAuctionLenCmdWithDeps(deps), "--to-block", "1")
	if result.err != nil {
		t.Fatal(result.err)
	}
}

func TestCstAuctionLengthValidationSetupAndErrors(t *testing.T) {
	t.Run("invalid contract", func(t *testing.T) {
		deps := defaultCstAuctionLenDeps()
		deps.getenv = func(string) string {
			t.Fatal("environment read")
			return ""
		}
		result := executeCommand(
			newScanCstAuctionLenCmdWithDeps(deps),
			"--contract", "invalid",
		)
		if result.err == nil || !strings.Contains(result.err.Error(), "invalid --contract") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("missing RPC URL", func(t *testing.T) {
		deps := defaultCstAuctionLenDeps()
		deps.getenv = func(string) string { return "" }
		result := executeCommand(newScanCstAuctionLenCmdWithDeps(deps))
		if result.err == nil || !strings.Contains(result.err.Error(), "RPC_URL") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("RPC dial", func(t *testing.T) {
		want := errors.New("dial failed")
		deps := defaultCstAuctionLenDeps()
		deps.getenv = func(string) string { return "rpc" }
		deps.dialRPC = func(context.Context, string) (cstAuctionLenRPC, error) { return nil, want }
		result := executeCommand(newScanCstAuctionLenCmdWithDeps(deps))
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "RPC dial") {
			t.Fatalf("error = %v", result.err)
		}
	})

	t.Run("database open closes RPC", func(t *testing.T) {
		rpc := &fakeOpsRPC{}
		want := errors.New("database unavailable")
		deps := defaultCstAuctionLenDeps()
		deps.getenv = func(string) string { return "rpc" }
		deps.dialRPC = func(context.Context, string) (cstAuctionLenRPC, error) { return rpc, nil }
		deps.openDB = func(string, string) (*sql.DB, error) { return nil, want }
		result := executeCommand(newScanCstAuctionLenCmdWithDeps(deps), "--db", "database")
		if !errors.Is(result.err, want) || !strings.Contains(result.err.Error(), "db open") {
			t.Fatalf("error = %v", result.err)
		}
		if !rpc.closed.Load() {
			t.Fatal("RPC client was not closed")
		}
	})

	for _, test := range []struct {
		name string
		err  error
	}{
		{name: "scan", err: errors.New("scan failed")},
		{name: "cancellation", err: context.Canceled},
	} {
		t.Run(test.name, func(t *testing.T) {
			db := newCommandTestDB(t)
			rpc := &fakeOpsRPC{}
			deps := defaultCstAuctionLenDeps()
			deps.getenv = func(string) string { return "rpc" }
			deps.dialRPC = func(context.Context, string) (cstAuctionLenRPC, error) { return rpc, nil }
			deps.openDB = func(string, string) (*sql.DB, error) { return db, nil }
			deps.newKeySource = func(*sql.DB) cstscan.KeySource { return nil }
			deps.scan = func(context.Context, cstscan.Config, cstscan.Options) (cstscan.Stats, error) {
				return cstscan.Stats{}, test.err
			}
			result := executeCommand(newScanCstAuctionLenCmdWithDeps(deps), "--db", "database")
			if !errors.Is(result.err, test.err) {
				t.Fatalf("error = %v", result.err)
			}
			assertCommandDBClosed(t, db)
		})
	}
}

func TestAssetSmokeAndScanDefaultDependencyAdapters(t *testing.T) {
	_ = runVerifyTokenImages
	_ = connectSmoketestDB

	db := newCommandTestDB(t)
	if _, ok := defaultAssetsInventoryDeps().newSource(db).(opsassets.SQLTokenSource); !ok {
		t.Fatal("default inventory source is not SQLTokenSource")
	}
	if _, ok := defaultAssetsThumbnailsDeps().newSource(db).(opsassets.SQLTokenSource); !ok {
		t.Fatal("default thumbnail source is not SQLTokenSource")
	}
	imageClient, ok := defaultAssetsVerifyTokenImagesDeps().client.(*http.Client)
	if !ok || imageClient.Timeout != rwalkImageRequestTimeout {
		t.Fatalf("default image client = %#v", imageClient)
	}
	if _, ok := defaultSmoketestDeps().newSource(db).(smoketest.SQLParameterSource); !ok {
		t.Fatal("default smoketest source is not SQLParameterSource")
	}
	if _, ok := defaultCstAuctionLenDeps().newKeySource(db).(cstscan.PostgresKeySource); !ok {
		t.Fatal("default scan key source is not PostgresKeySource")
	}
	if _, err := defaultCstAuctionLenDeps().dialRPC(context.Background(), "://bad-rpc-url"); err == nil {
		t.Fatal("default scan RPC dialer accepted malformed URL")
	}
}
