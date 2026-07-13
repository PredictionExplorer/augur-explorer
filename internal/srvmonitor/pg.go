package srvmonitor

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

// dbConnectTimeout bounds one connection attempt to a monitored database.
// The legacy lib/pq connector had no limit, so a black-holed host could hang
// a monitor's check cycle forever.
const dbConnectTimeout = 10 * time.Second

// DBConn is the narrow database handle the monitors use for one check cycle.
// *pgx.Conn satisfies it; tests inject scripted fakes.
type DBConn interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Close(ctx context.Context) error
}

// Connector opens a short-lived database connection for one check cycle.
type Connector func(ctx context.Context, cfg DatabaseConfig) (DBConn, error)

// ConnectPostgres is the production Connector: a single pgx connection in
// UTC with a bounded connect timeout. Credentials are set as config fields,
// never interpolated into a connection string.
func ConnectPostgres(ctx context.Context, cfg DatabaseConfig) (DBConn, error) {
	host, portStr, err := net.SplitHostPort(cfg.Host)
	if err != nil {
		host = cfg.Host
		portStr = "5432"
	}
	port, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid port in host %q: %w", cfg.Host, err)
	}

	// Parse an empty string to pick up libpq-compatible defaults (including
	// PG* environment variables, which the legacy lib/pq connector honored),
	// then pin the monitor-specific fields.
	connCfg, err := pgx.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("building connection config: %w", err)
	}
	connCfg.Host = host
	connCfg.Port = uint16(port)
	connCfg.User = cfg.User
	connCfg.Password = cfg.Pass
	connCfg.Database = cfg.DBName
	connCfg.ConnectTimeout = dbConnectTimeout
	// Monitored values are compared across servers; pin the session to UTC
	// like every other database consumer in this repository.
	connCfg.RuntimeParams["timezone"] = "UTC"

	conn, err := pgx.ConnectConfig(ctx, connCfg)
	if err != nil {
		return nil, fmt.Errorf("error connecting: %w", err)
	}
	return conn, nil
}

// isNoRows reports whether err is an empty query result.
func isNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
