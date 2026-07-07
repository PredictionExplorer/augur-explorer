package store

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

// slowQueryThreshold is the elapsed time above which a successful query is
// logged as slow.
const slowQueryThreshold = 500 * time.Millisecond

// maxTracedSQLLen bounds the SQL text captured in log records; the store's
// hand-written UNION queries run to several kilobytes and the head is enough
// to identify them.
const maxTracedSQLLen = 200

// queryTracer is a pgx.QueryTracer that reports failed and slow queries
// through slog. It replaces the legacy Log_msg file logger for converted
// query code: routine successes stay silent, problems become structured
// log records.
type queryTracer struct {
	logger *slog.Logger
	slow   time.Duration
}

func newQueryTracer(logger *slog.Logger) *queryTracer {
	return &queryTracer{logger: logger, slow: slowQueryThreshold}
}

type traceQueryKey struct{}

type traceQueryData struct {
	start time.Time
	sql   string
}

// TraceQueryStart implements pgx.QueryTracer.
func (t *queryTracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	return context.WithValue(ctx, traceQueryKey{}, traceQueryData{start: time.Now(), sql: data.SQL})
}

// TraceQueryEnd implements pgx.QueryTracer.
func (t *queryTracer) TraceQueryEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryEndData) {
	qd, ok := ctx.Value(traceQueryKey{}).(traceQueryData)
	if !ok {
		return
	}
	elapsed := time.Since(qd.start)
	switch {
	// A cancelled context is the client going away, not a database problem.
	case data.Err != nil && !errors.Is(data.Err, context.Canceled):
		t.logger.Error("query failed",
			"sql", truncateSQL(qd.sql),
			"elapsed_ms", elapsed.Milliseconds(),
			"err", data.Err)
	case data.Err == nil && elapsed >= t.slow:
		t.logger.Warn("slow query",
			"sql", truncateSQL(qd.sql),
			"elapsed_ms", elapsed.Milliseconds())
	}
}

func truncateSQL(sql string) string {
	if len(sql) <= maxTracedSQLLen {
		return sql
	}
	return sql[:maxTracedSQLLen] + "..."
}
