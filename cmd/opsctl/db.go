package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// opsDB is the PostgreSQL handle the opsctl commands hand to their engines.
// *pgxpool.Pool satisfies it, and it satisfies every ops package's narrow
// Querier interface; wiring tests inject fakes.
type opsDB interface {
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
	Close()
}

// openOpsDB returns an opener for a bounded pgx pool. Like the retired
// database/sql opener it validates the connection string without dialing;
// the first query (or an explicit Ping) connects. pgconn redacts the
// password when a malformed connection string is reported.
func openOpsDB(maxConns int32) func(ctx context.Context, conn string) (opsDB, error) {
	return func(ctx context.Context, conn string) (opsDB, error) {
		poolConfig, err := pgxpool.ParseConfig(conn)
		if err != nil {
			return nil, fmt.Errorf("db pool config: %w", err)
		}
		poolConfig.MaxConns = maxConns
		pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
		if err != nil {
			return nil, fmt.Errorf("db pool connect: %w", err)
		}
		return pool, nil
	}
}
