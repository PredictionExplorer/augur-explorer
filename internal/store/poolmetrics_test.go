package store_test

import (
	"context"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/testutil"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// lazyPool builds a pool that never dials (pgxpool connects on first
// acquire), so the collector's shape is testable without a database.
func lazyPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	cfg, err := pgxpool.ParseConfig("postgres://user:unused@127.0.0.1:1/unused?pool_max_conns=7")
	if err != nil {
		t.Fatalf("ParseConfig: %v", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		t.Fatalf("NewWithConfig: %v", err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// TestPoolCollectorExposition pins every metric name, help string, type and
// idle-pool value. CollectAndCompare parses the expectation like a scrape,
// so a renamed or retyped series fails here before it breaks dashboards.
func TestPoolCollectorExposition(t *testing.T) {
	t.Parallel()
	collector := store.NewPoolCollector(lazyPool(t))

	problems, err := testutil.CollectAndLint(collector)
	if err != nil {
		t.Fatalf("CollectAndLint: %v", err)
	}
	if len(problems) > 0 {
		t.Fatalf("metric lint problems: %v", problems)
	}

	const want = `
# HELP rwcg_db_pool_max_conns Configured maximum size of the pgx connection pool.
# TYPE rwcg_db_pool_max_conns gauge
rwcg_db_pool_max_conns 7
# HELP rwcg_db_pool_total_conns Connections currently open, acquired or constructing.
# TYPE rwcg_db_pool_total_conns gauge
rwcg_db_pool_total_conns 0
# HELP rwcg_db_pool_acquired_conns Connections currently checked out by queries.
# TYPE rwcg_db_pool_acquired_conns gauge
rwcg_db_pool_acquired_conns 0
# HELP rwcg_db_pool_idle_conns Open connections currently idle in the pool.
# TYPE rwcg_db_pool_idle_conns gauge
rwcg_db_pool_idle_conns 0
# HELP rwcg_db_pool_constructing_conns Connections currently being established.
# TYPE rwcg_db_pool_constructing_conns gauge
rwcg_db_pool_constructing_conns 0
# HELP rwcg_db_pool_acquires_total Successful connection acquisitions since process start.
# TYPE rwcg_db_pool_acquires_total counter
rwcg_db_pool_acquires_total 0
# HELP rwcg_db_pool_acquire_duration_seconds_total Cumulative time spent waiting to acquire connections.
# TYPE rwcg_db_pool_acquire_duration_seconds_total counter
rwcg_db_pool_acquire_duration_seconds_total 0
# HELP rwcg_db_pool_empty_acquires_total Acquisitions that had to wait because the pool was empty (saturation signal).
# TYPE rwcg_db_pool_empty_acquires_total counter
rwcg_db_pool_empty_acquires_total 0
# HELP rwcg_db_pool_empty_acquire_wait_seconds_total Cumulative wait time of acquisitions that found the pool empty.
# TYPE rwcg_db_pool_empty_acquire_wait_seconds_total counter
rwcg_db_pool_empty_acquire_wait_seconds_total 0
# HELP rwcg_db_pool_canceled_acquires_total Acquisitions abandoned because the caller's context was cancelled first.
# TYPE rwcg_db_pool_canceled_acquires_total counter
rwcg_db_pool_canceled_acquires_total 0
# HELP rwcg_db_pool_new_conns_total Connections opened since process start (churn signal with the destroy counters).
# TYPE rwcg_db_pool_new_conns_total counter
rwcg_db_pool_new_conns_total 0
# HELP rwcg_db_pool_max_lifetime_destroys_total Connections closed for exceeding MaxConnLifetime.
# TYPE rwcg_db_pool_max_lifetime_destroys_total counter
rwcg_db_pool_max_lifetime_destroys_total 0
# HELP rwcg_db_pool_max_idle_destroys_total Connections closed for exceeding MaxConnIdleTime.
# TYPE rwcg_db_pool_max_idle_destroys_total counter
rwcg_db_pool_max_idle_destroys_total 0
`
	if err := testutil.CollectAndCompare(collector, strings.NewReader(want)); err != nil {
		t.Fatalf("CollectAndCompare: %v", err)
	}

	if got := testutil.CollectAndCount(collector); got != 13 {
		t.Fatalf("collected %d metrics, want 13", got)
	}
}
