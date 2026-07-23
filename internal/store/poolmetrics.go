// Prometheus visibility for the shared pgx pool: saturation, acquire
// latency and connection churn are the classic missing signals when an API
// slows down under load. The collector reads pgxpool.Stat on scrape, so it
// adds no per-query bookkeeping.

package store

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
)

// poolMetric maps one pgxpool.Stat accessor onto a Prometheus metric.
type poolMetric struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
	value     func(*pgxpool.Stat) float64
}

// PoolCollector exposes pgxpool statistics as rwcg_db_pool_* metrics. It
// implements prometheus.Collector; the binaries that serve /metrics
// (apiserver and both ETLs) register one per process next to their store.
type PoolCollector struct {
	stat    func() *pgxpool.Stat
	metrics []poolMetric
}

// NewPoolCollector builds a collector over the pool's live statistics.
func NewPoolCollector(pool *pgxpool.Pool) *PoolCollector {
	gauge := func(name, help string, value func(*pgxpool.Stat) float64) poolMetric {
		return poolMetric{
			desc:      prometheus.NewDesc(name, help, nil, nil),
			valueType: prometheus.GaugeValue,
			value:     value,
		}
	}
	counter := func(name, help string, value func(*pgxpool.Stat) float64) poolMetric {
		return poolMetric{
			desc:      prometheus.NewDesc(name, help, nil, nil),
			valueType: prometheus.CounterValue,
			value:     value,
		}
	}
	return &PoolCollector{
		stat: pool.Stat,
		metrics: []poolMetric{
			gauge("rwcg_db_pool_max_conns",
				"Configured maximum size of the pgx connection pool.",
				func(s *pgxpool.Stat) float64 { return float64(s.MaxConns()) }),
			gauge("rwcg_db_pool_total_conns",
				"Connections currently open, acquired or constructing.",
				func(s *pgxpool.Stat) float64 { return float64(s.TotalConns()) }),
			gauge("rwcg_db_pool_acquired_conns",
				"Connections currently checked out by queries.",
				func(s *pgxpool.Stat) float64 { return float64(s.AcquiredConns()) }),
			gauge("rwcg_db_pool_idle_conns",
				"Open connections currently idle in the pool.",
				func(s *pgxpool.Stat) float64 { return float64(s.IdleConns()) }),
			gauge("rwcg_db_pool_constructing_conns",
				"Connections currently being established.",
				func(s *pgxpool.Stat) float64 { return float64(s.ConstructingConns()) }),
			counter("rwcg_db_pool_acquires_total",
				"Successful connection acquisitions since process start.",
				func(s *pgxpool.Stat) float64 { return float64(s.AcquireCount()) }),
			counter("rwcg_db_pool_acquire_duration_seconds_total",
				"Cumulative time spent waiting to acquire connections.",
				func(s *pgxpool.Stat) float64 { return s.AcquireDuration().Seconds() }),
			counter("rwcg_db_pool_empty_acquires_total",
				"Acquisitions that had to wait because the pool was empty (saturation signal).",
				func(s *pgxpool.Stat) float64 { return float64(s.EmptyAcquireCount()) }),
			counter("rwcg_db_pool_empty_acquire_wait_seconds_total",
				"Cumulative wait time of acquisitions that found the pool empty.",
				func(s *pgxpool.Stat) float64 { return s.EmptyAcquireWaitTime().Seconds() }),
			counter("rwcg_db_pool_canceled_acquires_total",
				"Acquisitions abandoned because the caller's context was cancelled first.",
				func(s *pgxpool.Stat) float64 { return float64(s.CanceledAcquireCount()) }),
			counter("rwcg_db_pool_new_conns_total",
				"Connections opened since process start (churn signal with the destroy counters).",
				func(s *pgxpool.Stat) float64 { return float64(s.NewConnsCount()) }),
			counter("rwcg_db_pool_max_lifetime_destroys_total",
				"Connections closed for exceeding MaxConnLifetime.",
				func(s *pgxpool.Stat) float64 { return float64(s.MaxLifetimeDestroyCount()) }),
			counter("rwcg_db_pool_max_idle_destroys_total",
				"Connections closed for exceeding MaxConnIdleTime.",
				func(s *pgxpool.Stat) float64 { return float64(s.MaxIdleDestroyCount()) }),
		},
	}
}

// Describe implements prometheus.Collector.
func (c *PoolCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range c.metrics {
		ch <- metric.desc
	}
}

// Collect implements prometheus.Collector: one Stat snapshot per scrape.
func (c *PoolCollector) Collect(ch chan<- prometheus.Metric) {
	stat := c.stat()
	for _, metric := range c.metrics {
		ch <- prometheus.MustNewConstMetric(metric.desc, metric.valueType, metric.value(stat))
	}
}
