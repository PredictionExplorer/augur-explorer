package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/smoketest"
)

// smoketestMaxConns bounds the parameter-database pool; the source runs one
// query at a time.
const smoketestMaxConns = 2

type smoketestDeps struct {
	getenv    func(string) string
	openDB    func(context.Context, string) (opsDB, error)
	ping      func(context.Context, opsDB) error
	newSource func(opsDB) smoketest.ParameterSource
	client    smoketest.HTTPClient
	apiBase   func() string
	run       func(context.Context, smoketest.Options) (smoketest.Summary, error)
}

func defaultSmoketestDeps() smoketestDeps {
	return smoketestDeps{
		getenv: os.Getenv,
		openDB: openOpsDB(smoketestMaxConns),
		ping: func(ctx context.Context, db opsDB) error {
			return db.Ping(ctx)
		},
		newSource: func(db opsDB) smoketest.ParameterSource {
			return smoketest.SQLParameterSource{DB: db}
		},
		client:  &http.Client{Timeout: 60 * time.Second},
		apiBase: smoketestAPIBase,
		run:     smoketest.Run,
	}
}

// newSmoketestCmd builds `opsctl smoketest`, the replacement for the
// standalone api_smoketest tool.
func newSmoketestCmd() *cobra.Command {
	return newSmoketestCmdWithDeps(defaultSmoketestDeps())
}

func newSmoketestCmdWithDeps(deps smoketestDeps) *cobra.Command {
	suite := string(smoketest.SuiteV2)
	cmd := &cobra.Command{
		Use:   "smoketest",
		Short: "Validate API surfaces and operational probes",
		Long: `Validates the canonical v2 API by default. The v2 suite hits every
documented GET operation, requires HTTP 200, validates each response against
the embedded OpenAPI contract and rejects deprecated response headers.

The frozen v1 regression remains available with --suite=v1, and --suite=both
runs v2 followed by v1. The lightweight --suite=operational set checks health,
readiness, version and stable v2 resources without opening PostgreSQL.

Environment:

	PGSQL_HOST (host:port), PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD
	           database used to fetch real URL parameter values (not needed
	           by --suite=operational)
	HTTP_PORT  websrv port; API base defaults to http://127.0.0.1:$HTTP_PORT
	API_BASE   optional; overrides the base URL`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSmoketestWithDeps(cmd, deps, suite)
		},
	}
	cmd.Flags().StringVar(
		&suite,
		"suite",
		string(smoketest.SuiteV2),
		"probe suite: v2, v1, both, or operational",
	)
	return cmd
}

func runSmoketest(cmd *cobra.Command) error {
	return runSmoketestWithDeps(cmd, defaultSmoketestDeps(), string(smoketest.SuiteV2))
}

func runSmoketestWithDeps(cmd *cobra.Command, deps smoketestDeps, suiteValue string) error {
	suite, err := smoketest.ParseSuite(suiteValue)
	if err != nil {
		return err
	}

	var source smoketest.ParameterSource
	if suite.RequiresParameters() {
		db, err := connectSmoketestDBWithDeps(cmd.Context(), deps)
		if err != nil {
			return err
		}
		defer db.Close()
		source = deps.newSource(db)
	}
	_, err = deps.run(cmd.Context(), smoketest.Options{
		Source:  source,
		Client:  deps.client,
		BaseURL: deps.apiBase(),
		Output:  cmd.OutOrStdout(),
		Suite:   suite,
	})
	return err
}

// smoketestAPIBase returns the base URL: API_BASE when set, otherwise
// http://127.0.0.1:$HTTP_PORT (port defaults to 9090).
func smoketestAPIBase() string {
	if base := strings.TrimSpace(os.Getenv("API_BASE")); base != "" {
		return strings.TrimRight(base, "/")
	}
	port := strings.TrimSpace(os.Getenv("HTTP_PORT"))
	if port == "" {
		port = "9090"
	}
	return "http://127.0.0.1:" + port
}

// connectSmoketestDB opens the parameter database from the PGSQL_* env vars.
func connectSmoketestDB(ctx context.Context) (opsDB, error) {
	return connectSmoketestDBWithDeps(ctx, defaultSmoketestDeps())
}

func connectSmoketestDBWithDeps(ctx context.Context, deps smoketestDeps) (opsDB, error) {
	dsn := strings.TrimSpace(deps.getenv("DATABASE_URL"))
	if dsn == "" {
		host := deps.getenv("PGSQL_HOST")
		user := deps.getenv("PGSQL_USERNAME")
		dbName := deps.getenv("PGSQL_DATABASE")
		pass := deps.getenv("PGSQL_PASSWORD")
		if host == "" || user == "" || dbName == "" {
			return nil, errors.New("DATABASE_URL or PGSQL_HOST / PGSQL_USERNAME / PGSQL_DATABASE are required")
		}
		dsn = (&url.URL{
			Scheme:   "postgres",
			User:     url.UserPassword(user, pass),
			Host:     host,
			Path:     "/" + dbName,
			RawQuery: "sslmode=disable",
		}).String()
	}
	db, err := deps.openDB(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("DB open failed: %w", err)
	}
	if err := deps.ping(ctx, db); err != nil {
		db.Close()
		return nil, fmt.Errorf("DB ping failed: %w", err)
	}
	return db, nil
}
