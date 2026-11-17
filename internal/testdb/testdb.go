// Package testdb provides a disposable, fully-migrated PostgreSQL database
// for integration tests.
//
// Typical use (in a test file guarded by the `integration` build tag):
//
//	func TestSomething(t *testing.T) {
//		db := testdb.New(t)          // throwaway database + schema, cleaned up automatically
//		// ... run queries against db.SQL, db.Pool or db.ConnString ...
//	}
//
// Two backends, selected by the TEST_PGSQL_URL environment variable:
//
//   - TEST_PGSQL_URL set: an already-running PostgreSQL server (13+), no
//     Docker involved. The value is a URL such as
//     postgres://user:password@localhost:5432/postgres?sslmode=disable
//     naming a maintenance database the role can connect to; the role needs
//     the CREATEDB privilege. Every Start provisions a uniquely-named
//     throwaway database on that server and drops it on cleanup.
//     Misconfiguration (unreachable server, missing privilege) fails the
//     tests instead of skipping them.
//   - TEST_PGSQL_URL unset: a PostgreSQL testcontainer, which requires a
//     running Docker daemon; without one the tests skip.
//
// Run with `go test -tags=integration` (see `make test-integration`).
package testdb

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// DB is a running, migrated PostgreSQL instance dedicated to one test.
type DB struct {
	// Pool is the native pgx connection pool (store.NewFromPool wraps it).
	Pool *pgxpool.Pool
	// SQL is a database/sql view over the same pool.
	SQL *sql.DB
	// ConnString is a keyword/value or URL connection string for the database.
	ConnString string
}

const containerImage = "postgres:17-alpine"

// New starts a throwaway PostgreSQL database (see the package comment for
// the TEST_PGSQL_URL / testcontainer backends), applies all goose migrations
// from db/migrations, and registers cleanup with t. In testcontainer mode it
// skips the test when the environment cannot run containers (no Docker
// daemon); in TEST_PGSQL_URL mode every failure is fatal.
func New(t *testing.T) *DB {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	db, stop, err := Start(ctx)
	if err != nil {
		if errors.Is(err, ErrContainerUnavailable) {
			t.Skipf("skipping: cannot start postgres container (is Docker running?): %v", err)
		}
		t.Fatalf("starting test database: %v", err)
	}
	t.Cleanup(stop)
	return db
}

// ErrContainerUnavailable wraps container startup failures (typically: Docker
// is not running). Callers that manage the database from TestMain can detect
// it to skip integration tests instead of failing them.
var ErrContainerUnavailable = errors.New("test database container unavailable")

// Start boots a throwaway PostgreSQL database and applies all goose
// migrations: on the server named by TEST_PGSQL_URL when that variable is
// set, otherwise in a testcontainer. It is intended for TestMain-style
// callers that need the database to outlive a single test; the returned stop
// function drops the database (or terminates the container) and closes the
// pool. Prefer New in ordinary tests.
func Start(ctx context.Context) (*DB, func(), error) {
	if serverURL := os.Getenv("TEST_PGSQL_URL"); serverURL != "" {
		return startExternal(ctx, serverURL)
	}
	container, err := postgres.Run(ctx, containerImage,
		postgres.WithDatabase("rwcg_test"),
		postgres.WithUsername("rwcg"),
		postgres.WithPassword("rwcg"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %w", ErrContainerUnavailable, err)
	}
	terminate := func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			fmt.Fprintf(os.Stderr, "terminating postgres container: %v\n", err)
		}
	}

	connString, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		terminate()
		return nil, nil, fmt.Errorf("postgres connection string: %w", err)
	}

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		terminate()
		return nil, nil, fmt.Errorf("open postgres pool: %w", err)
	}
	db := stdlib.OpenDBFromPool(pool)
	stop := func() {
		_ = db.Close()
		pool.Close()
		terminate()
	}

	if err := Migrate(ctx, db); err != nil {
		stop()
		return nil, nil, fmt.Errorf("applying migrations: %w", err)
	}

	return &DB{Pool: pool, SQL: db, ConnString: connString}, stop, nil
}

// startExternal provisions a uniquely-named throwaway database on the
// PostgreSQL server named by TEST_PGSQL_URL, migrates it, and drops it on
// stop. Every failure is a hard error (never ErrContainerUnavailable): the
// variable being set is an explicit statement that this server must work.
func startExternal(ctx context.Context, serverURL string) (*DB, func(), error) {
	parsed, err := url.Parse(serverURL)
	if err != nil || parsed.Host == "" {
		return nil, nil, fmt.Errorf("TEST_PGSQL_URL must be a URL like postgres://user:password@host:port/maintenancedb (got %q): %w", serverURL, err)
	}

	adminCfg, err := pgxpool.ParseConfig(serverURL)
	if err != nil {
		return nil, nil, fmt.Errorf("TEST_PGSQL_URL: parsing server URL: %w", err)
	}
	// The admin pool only ever runs one CREATE/DROP DATABASE at a time, and
	// many test packages share the server in parallel: keep it tiny so the
	// suite stays well under max_connections.
	adminCfg.MaxConns = 2
	admin, err := pgxpool.NewWithConfig(ctx, adminCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("TEST_PGSQL_URL: opening admin pool: %w", err)
	}
	if err := admin.Ping(ctx); err != nil {
		admin.Close()
		return nil, nil, fmt.Errorf("TEST_PGSQL_URL: cannot connect to PostgreSQL server: %w", err)
	}

	var suffix [8]byte
	if _, err := rand.Read(suffix[:]); err != nil {
		admin.Close()
		return nil, nil, fmt.Errorf("generating test database name: %w", err)
	}
	name := "rwcg_test_" + hex.EncodeToString(suffix[:])
	quoted := pgx.Identifier{name}.Sanitize()

	if _, err := admin.Exec(ctx, "CREATE DATABASE "+quoted); err != nil {
		admin.Close()
		return nil, nil, fmt.Errorf("TEST_PGSQL_URL: creating database %s (does the role have CREATEDB?): %w", name, err)
	}
	dropDB := func() {
		// Fresh context: stop runs during cleanup, when the test context
		// may already be done. FORCE (PostgreSQL 13+) evicts stragglers.
		dropCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if _, err := admin.Exec(dropCtx, "DROP DATABASE "+quoted+" WITH (FORCE)"); err != nil {
			fmt.Fprintf(os.Stderr, "dropping test database %s: %v\n", name, err)
		}
		admin.Close()
	}

	parsed.Path = "/" + name
	// The testcontainer image runs with timezone=UTC; pin the session
	// timezone on the external server too (pgx forwards it as a runtime
	// parameter), so timestamp rendering and date-bucketed queries in
	// golden files are machine-independent.
	query := parsed.Query()
	query.Set("timezone", "UTC")
	parsed.RawQuery = query.Encode()
	connString := parsed.String()
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		dropDB()
		return nil, nil, fmt.Errorf("open postgres pool: %w", err)
	}
	db := stdlib.OpenDBFromPool(pool)
	stop := func() {
		_ = db.Close()
		pool.Close()
		dropDB()
	}

	if err := Migrate(ctx, db); err != nil {
		stop()
		return nil, nil, fmt.Errorf("applying migrations: %w", err)
	}

	return &DB{Pool: pool, SQL: db, ConnString: connString}, stop, nil
}

// Migrate applies all goose migrations from db/migrations to the given database.
func Migrate(ctx context.Context, db *sql.DB) error {
	provider, err := goose.NewProvider(goose.DialectPostgres, db, os.DirFS(migrationsDir()))
	if err != nil {
		return fmt.Errorf("creating goose provider: %w", err)
	}
	if _, err := provider.Up(ctx); err != nil {
		return fmt.Errorf("running migrations: %w", err)
	}
	return nil
}

// migrationsDir locates db/migrations relative to this source file, so tests
// work regardless of the package they run from.
func migrationsDir() string {
	_, thisFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(thisFile), "..", "..", "db", "migrations")
}
