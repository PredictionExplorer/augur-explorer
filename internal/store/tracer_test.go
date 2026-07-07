package store

import (
	"bytes"
	"context"
	"errors"
	"log/slog"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
)

// runTrace pushes one query through the tracer and returns the log output.
func runTrace(t *testing.T, tracer *queryTracer, sql string, queryErr error) string {
	t.Helper()
	var buf bytes.Buffer
	tracer.logger = slog.New(slog.NewTextHandler(&buf, nil))
	ctx := tracer.TraceQueryStart(context.Background(), nil, pgx.TraceQueryStartData{SQL: sql})
	tracer.TraceQueryEnd(ctx, nil, pgx.TraceQueryEndData{Err: queryErr})
	return buf.String()
}

func TestTracerLogsFailedQuery(t *testing.T) {
	out := runTrace(t, &queryTracer{slow: time.Hour}, "SELECT broken", errors.New("relation does not exist"))
	if !strings.Contains(out, "query failed") || !strings.Contains(out, "relation does not exist") {
		t.Errorf("failed query not logged: %q", out)
	}
	if !strings.Contains(out, "SELECT broken") {
		t.Errorf("SQL missing from log: %q", out)
	}
}

func TestTracerSkipsCancelledQuery(t *testing.T) {
	out := runTrace(t, &queryTracer{slow: time.Hour}, "SELECT 1", context.Canceled)
	if out != "" {
		t.Errorf("cancelled query logged as failure: %q", out)
	}
}

func TestTracerLogsSlowQuery(t *testing.T) {
	// slow=0 makes every successful query "slow".
	out := runTrace(t, &queryTracer{slow: 0}, "SELECT pg_sleep(10)", nil)
	if !strings.Contains(out, "slow query") {
		t.Errorf("slow query not logged: %q", out)
	}
}

func TestTracerSilentOnFastSuccess(t *testing.T) {
	out := runTrace(t, &queryTracer{slow: time.Hour}, "SELECT 1", nil)
	if out != "" {
		t.Errorf("fast successful query produced output: %q", out)
	}
}

func TestTracerMissingStartData(t *testing.T) {
	tracer := &queryTracer{slow: 0}
	var buf bytes.Buffer
	tracer.logger = slog.New(slog.NewTextHandler(&buf, nil))
	// End without Start must not panic or log garbage.
	tracer.TraceQueryEnd(context.Background(), nil, pgx.TraceQueryEndData{})
	if buf.String() != "" {
		t.Errorf("unexpected output: %q", buf.String())
	}
}

func TestTruncateSQL(t *testing.T) {
	short := "SELECT 1"
	if got := truncateSQL(short); got != short {
		t.Errorf("truncateSQL(%q) = %q", short, got)
	}
	long := strings.Repeat("x", maxTracedSQLLen+50)
	got := truncateSQL(long)
	if len(got) != maxTracedSQLLen+3 || !strings.HasSuffix(got, "...") {
		t.Errorf("truncateSQL long input: len=%d suffix=%q", len(got), got[len(got)-3:])
	}
}
