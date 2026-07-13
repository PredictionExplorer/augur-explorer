package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq" // postgres driver
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/smoketest"
)

type smoketestDeps struct {
	getenv    func(string) string
	openDB    func(string, string) (*sql.DB, error)
	ping      func(context.Context, *sql.DB) error
	newSource func(*sql.DB) smoketest.ParameterSource
	client    smoketest.HTTPClient
	apiBase   func() string
	run       func(context.Context, smoketest.Options) (smoketest.Summary, error)
}

func defaultSmoketestDeps() smoketestDeps {
	return smoketestDeps{
		getenv: os.Getenv,
		openDB: sql.Open,
		ping: func(ctx context.Context, db *sql.DB) error {
			return db.PingContext(ctx)
		},
		newSource: func(db *sql.DB) smoketest.ParameterSource {
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
	cmd := &cobra.Command{
		Use:   "smoketest",
		Short: "Hit every /api/cosmicgame/... endpoint and report failures",
		Long: `Hits every /api/cosmicgame/... endpoint of the cosmicgame websrv and
reports non-200 responses (and in-body error/status:0) as FAILED. A 400/500 or
an "error" body usually means a broken SQL query in the handler.

Environment:

	PGSQL_HOST (host:port), PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD
	           database used to fetch real URL parameter values
	HTTP_PORT  websrv port; API base defaults to http://127.0.0.1:$HTTP_PORT
	API_BASE   optional; overrides the base URL`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSmoketestWithDeps(cmd, deps)
		},
	}
	return cmd
}

func runSmoketest(cmd *cobra.Command) error {
	return runSmoketestWithDeps(cmd, defaultSmoketestDeps())
}

func runSmoketestWithDeps(cmd *cobra.Command, deps smoketestDeps) error {
	db, err := connectSmoketestDBWithDeps(cmd.Context(), deps)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	_, err = deps.run(cmd.Context(), smoketest.Options{
		Source:  deps.newSource(db),
		Client:  deps.client,
		BaseURL: deps.apiBase(),
		Output:  cmd.OutOrStdout(),
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
func connectSmoketestDB(ctx context.Context) (*sql.DB, error) {
	return connectSmoketestDBWithDeps(ctx, defaultSmoketestDeps())
}

func connectSmoketestDBWithDeps(ctx context.Context, deps smoketestDeps) (*sql.DB, error) {
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
	db, err := deps.openDB("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("DB open failed: %w", err)
	}
	if err := deps.ping(ctx, db); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("DB ping failed: %w", err)
	}
	return db, nil
}
