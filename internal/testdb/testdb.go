// Package testdb provides a disposable, fully-migrated PostgreSQL database
// for integration tests, backed by testcontainers.
//
// Typical use (in a test file guarded by the `integration` build tag):
//
//	func TestSomething(t *testing.T) {
//		db := testdb.New(t)          // container + schema, cleaned up automatically
//		// ... run queries against db.SQL, db.Pool or db.ConnString ...
//	}
//
// Tests using this package require a running Docker daemon and should be
// executed with `go test -tags=integration` (see `make test-integration`).
package testdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

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

// New starts a PostgreSQL container, applies all goose migrations from
// db/migrations, and registers cleanup with t. It skips the test when the
// environment cannot run containers (no Docker daemon).
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

// Start boots a PostgreSQL container and applies all goose migrations.
// It is intended for TestMain-style callers that need the database to outlive
// a single test; the returned stop function terminates the container and
// closes the pool. Prefer New in ordinary tests.
func Start(ctx context.Context) (*DB, func(), error) {
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
