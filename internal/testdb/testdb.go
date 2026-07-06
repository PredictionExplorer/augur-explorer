// Package testdb provides a disposable, fully-migrated PostgreSQL database
// for integration tests, backed by testcontainers.
//
// Typical use (in a test file guarded by the `integration` build tag):
//
//	func TestSomething(t *testing.T) {
//		db := testdb.New(t)          // container + schema, cleaned up automatically
//		// ... run queries against db.SQL or db.ConnString ...
//	}
//
// Tests using this package require a running Docker daemon and should be
// executed with `go test -tags=integration` (see `make test-integration`).
package testdb

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // database/sql driver ("pgx")
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// DB is a running, migrated PostgreSQL instance dedicated to one test.
type DB struct {
	// SQL is an open connection pool to the database.
	SQL *sql.DB
	// ConnString is a keyword/value or URL connection string for the database.
	ConnString string
}

const containerImage = "postgres:17-alpine"

// New starts a PostgreSQL container, applies all goose migrations from
// db/migrations, and registers cleanup with t. It skips the test when the
// environment cannot run containers (no Docker daemon).
func New(t *testing.T) *DB {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	container, err := postgres.Run(ctx, containerImage,
		postgres.WithDatabase("rwcg_test"),
		postgres.WithUsername("rwcg"),
		postgres.WithPassword("rwcg"),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		t.Skipf("skipping: cannot start postgres container (is Docker running?): %v", err)
	}
	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			t.Logf("terminating postgres container: %v", err)
		}
	})

	connString, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("postgres connection string: %v", err)
	}

	db, err := sql.Open("pgx", connString)
	if err != nil {
		t.Fatalf("open postgres: %v", err)
	}
	t.Cleanup(func() { db.Close() })

	if err := Migrate(ctx, db); err != nil {
		t.Fatalf("applying migrations: %v", err)
	}

	return &DB{SQL: db, ConnString: connString}
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

// SetLegacyEnv points the legacy PGSQL_* environment variables at this test
// database so code paths that still read them (internal/store) can connect.
// It uses t.Setenv, so the variables are restored automatically and the test
// is excluded from t.Parallel.
func (d *DB) SetLegacyEnv(t *testing.T) {
	t.Helper()
	u, err := url.Parse(d.ConnString)
	if err != nil {
		t.Fatalf("parsing connection string: %v", err)
	}
	pw, _ := u.User.Password()
	t.Setenv("PGSQL_USERNAME", u.User.Username())
	t.Setenv("PGSQL_PASSWORD", pw)
	t.Setenv("PGSQL_HOST", u.Host)
	t.Setenv("PGSQL_DATABASE", u.Path[1:])
}

// migrationsDir locates db/migrations relative to this source file, so tests
// work regardless of the package they run from.
func migrationsDir() string {
	_, thisFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(thisFile), "..", "..", "db", "migrations")
}
