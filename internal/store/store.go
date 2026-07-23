// Package store provides pgx-native database access for rwcg: the
// pool-owning Store with the base-layer queries (addresses, blocks,
// transactions, event logs, archive) plus the domain repositories in the
// randomwalk and cosmicgame subpackages.
package store

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DefaultMaxConns bounds the connection pool when Config.MaxConns is zero.
// The legacy database/sql pool was unbounded, which under load could pile
// hundreds of concurrent queries onto PostgreSQL; a moderate ceiling keeps
// tail latency sane while still allowing plenty of parallelism for the API
// server (the ETLs are sequential and use only one or two connections).
const DefaultMaxConns = 16

// Config holds PostgreSQL connection settings for New. Construct it with
// ConfigFromEnv or fill the fields explicitly (tests, tools with -db flags).
type Config struct {
	// URL is a complete postgres:// connection URL (or libpq DSN). When
	// set it wins over the field-based settings below (12-factor
	// DATABASE_URL support). The store still pins timezone=UTC and
	// search_path=public and applies its keepalive dialer and default
	// connect timeout on top.
	URL string

	// User, Password and Database are the libpq user/password/dbname values.
	User     string
	Password string
	Database string
	// Host is "host" or "host:port" for TCP. Empty selects a local Unix
	// socket with peer/trust auth (like running psql without -h); Password
	// is ignored in that case.
	Host string

	// MaxConns bounds the pool size; zero applies DefaultMaxConns.
	MaxConns int32

	// StatementTimeout, when positive, sets the server-side
	// statement_timeout on every pool connection: PostgreSQL aborts any
	// single statement that runs longer (SQLSTATE 57014). It is defense in
	// depth behind the caller's context deadlines — the backstop for a
	// query whose cancellation signal never arrives. Zero leaves the server
	// default (no limit); operator CLIs with legitimately heavy statements
	// stay unset.
	StatementTimeout time.Duration

	// IdleInTxSessionTimeout, when positive, sets the server-side
	// idle_in_transaction_session_timeout on every pool connection:
	// PostgreSQL terminates a session whose open transaction sits idle
	// longer than this (SQLSTATE 25P03), releasing its locks and snapshot.
	// It bounds the damage of a transaction held open by a stalled non-DB
	// call between statements. Zero leaves the server default (no limit).
	IdleInTxSessionTimeout time.Duration

	// Logger receives query traces (failed and slow queries) and connect
	// retry progress. nil disables query tracing and routes connect retry
	// messages through slog.Default().
	Logger *slog.Logger
}

// ConfigFromEnv reads the connection settings from the environment:
// DATABASE_URL when set (wins), otherwise the legacy PGSQL_* variables
// (PGSQL_USERNAME, PGSQL_PASSWORD, PGSQL_DATABASE, PGSQL_HOST).
func ConfigFromEnv() Config {
	return Config{
		URL:      strings.TrimSpace(os.Getenv("DATABASE_URL")),
		User:     os.Getenv("PGSQL_USERNAME"),
		Password: os.Getenv("PGSQL_PASSWORD"),
		Database: os.Getenv("PGSQL_DATABASE"),
		Host:     os.Getenv("PGSQL_HOST"),
	}
}

// escapeConnParam escapes a value for use inside a single-quoted libpq
// connection parameter.
func escapeConnParam(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "'", "\\'")
	return s
}

// defaultUnixSocketDir returns the PostgreSQL Unix socket directory to use
// when Host is empty. Linux distros put the socket in /var/run/postgresql;
// macOS (Homebrew) and other platforms use /tmp. Set explicitly because
// pgx's fallback default differs across builds.
func defaultUnixSocketDir() string {
	if runtime.GOOS == "linux" {
		return "/var/run/postgresql"
	}
	return "/tmp"
}

// connString renders the config as a libpq keyword/value connection string.
func (cfg Config) connString() string {
	connStr := "user='" + escapeConnParam(cfg.User) +
		"' dbname='" + escapeConnParam(cfg.Database) + "'"
	if cfg.Host == "" {
		// No password: Unix socket + trust/peer auth (like psql -U user).
		connStr += " host='" + escapeConnParam(defaultUnixSocketDir()) + "'"
	} else {
		connStr += " password='" + escapeConnParam(cfg.Password) + "'"
		host, port, err := net.SplitHostPort(cfg.Host)
		if err != nil {
			host = cfg.Host
			port = "5432"
		}
		connStr += " host='" + escapeConnParam(host) + "' port='" + escapeConnParam(port) + "'"
	}
	// connect_timeout bounds each connection attempt. Keepalive +
	// TCP_USER_TIMEOUT are applied at the socket level by keepaliveDialer.
	connStr += fmt.Sprintf(" connect_timeout=%d", int(dbConnectTimeout.Seconds()))
	return connStr
}

// poolConfig parses the connection settings — the URL when present,
// otherwise the rendered PGSQL_* keyword string — and applies the store's
// invariants on top: keepalive dialer, UTC timezone, public search_path,
// bounded pool size and (unless the URL sets its own) the default connect
// timeout.
func (cfg Config) poolConfig() (*pgxpool.Config, error) {
	connStr := cfg.URL
	if connStr == "" {
		connStr = cfg.connString()
	}
	poolCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		// Never echo the failing string: a URL embeds the password.
		return nil, fmt.Errorf("parse db config: %w", err)
	}
	if cfg.URL != "" && !strings.Contains(cfg.URL, "connect_timeout") {
		poolCfg.ConnConfig.ConnectTimeout = dbConnectTimeout
	}
	poolCfg.ConnConfig.DialFunc = newKeepaliveDialer().DialContext
	// Every connection works in UTC. The legacy connector ran
	// "SET timezone TO 0" on a single connection of the pool; a runtime
	// parameter applies it uniformly. Queries and scan helpers depend on
	// it, so a URL-provided timezone is deliberately overridden.
	poolCfg.ConnConfig.RuntimeParams["timezone"] = "UTC"
	// The schema is always public (§5.1): pinning search_path makes the
	// converted queries' bare table names resolve exactly like the
	// "public."-qualified names the legacy SQL used.
	poolCfg.ConnConfig.RuntimeParams["search_path"] = "public"
	// Server-side time bounds (D22 defense in depth): PostgreSQL enforces
	// them even when a client-side context deadline never fires. Rendered
	// in milliseconds, the GUCs' unit-less integer form.
	if cfg.StatementTimeout > 0 {
		poolCfg.ConnConfig.RuntimeParams["statement_timeout"] = strconv.FormatInt(cfg.StatementTimeout.Milliseconds(), 10)
	}
	if cfg.IdleInTxSessionTimeout > 0 {
		poolCfg.ConnConfig.RuntimeParams["idle_in_transaction_session_timeout"] = strconv.FormatInt(cfg.IdleInTxSessionTimeout.Milliseconds(), 10)
	}
	poolCfg.MaxConns = DefaultMaxConns
	if cfg.MaxConns > 0 {
		poolCfg.MaxConns = cfg.MaxConns
	}
	if cfg.Logger != nil {
		poolCfg.ConnConfig.Tracer = newQueryTracer(cfg.Logger)
	}
	return poolCfg, nil
}

// Store is the process-wide database handle: it owns a pgx connection pool
// plus the bounded address-id cache, and every query method (directly on
// Store or through the domain repos) runs on it. Create one per process with
// New and share it.
type Store struct {
	pool      *pgxpool.Pool
	addrCache *addressCache
}

// New connects to PostgreSQL and returns a ready Store. The initial
// connection is retried with capped backoff so a short link blip at process
// startup does not kill the service; after dbConnectMaxAttempts failures the
// last error is returned.
func New(ctx context.Context, cfg Config) (*Store, error) {
	poolCfg, err := cfg.poolConfig()
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("create db pool: %w", err)
	}
	if err := pingWithRetry(ctx, pool, cfg.Logger); err != nil {
		pool.Close()
		return nil, err
	}
	return NewFromPool(pool), nil
}

// NewFromPool wraps an existing pool (test harnesses, callers that build
// their own pool config). The Store takes ownership: Close closes the pool.
func NewFromPool(pool *pgxpool.Pool) *Store {
	return &Store{
		pool:      pool,
		addrCache: newAddressCache(DefaultAddressCacheSize),
	}
}

// Pool exposes the native pgx pool the query methods run on.
func (s *Store) Pool() *pgxpool.Pool { return s.pool }

// Close releases the underlying pool.
func (s *Store) Close() {
	s.pool.Close()
}

// pingWithRetry forces an initial connection, absorbing short link failures
// at process startup: dbConnectMaxAttempts attempts with exponential backoff
// capped at dbConnectRetryMaxWait.
func pingWithRetry(ctx context.Context, pool *pgxpool.Pool, logger *slog.Logger) error {
	if logger == nil {
		logger = slog.Default()
	}
	delay := dbConnectRetryDelay
	var lastErr error
	for attempt := 1; attempt <= dbConnectMaxAttempts; attempt++ {
		pingCtx, cancel := context.WithTimeout(ctx, dbConnectTimeout)
		err := pool.Ping(pingCtx)
		cancel()
		if err == nil {
			if attempt > 1 {
				logger.Info("database connected after retry", "attempt", attempt, "max_attempts", dbConnectMaxAttempts)
			}
			return nil
		}
		lastErr = err
		logger.Warn("database connect attempt failed", "attempt", attempt, "max_attempts", dbConnectMaxAttempts, "err", err)
		if attempt == dbConnectMaxAttempts {
			break
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}
		if delay *= 2; delay > dbConnectRetryMaxWait {
			delay = dbConnectRetryMaxWait
		}
	}
	return fmt.Errorf("db connect failed after %d attempts: %w", dbConnectMaxAttempts, lastErr)
}
