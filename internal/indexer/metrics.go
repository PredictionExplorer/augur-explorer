// Prometheus instrumentation of the engine. A nil *Metrics disables all
// recording, so tests and tools can run without a registry.

package indexer

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics holds the engine's Prometheus instruments. Create one per process
// with NewMetrics and pass it through Config.
type Metrics struct {
	lastBlock     prometheus.Gauge
	eventsTotal   *prometheus.CounterVec
	batchDuration prometheus.Histogram
	reorgsTotal   prometheus.Counter
	failuresTotal *prometheus.CounterVec
}

// NewMetrics builds and registers the engine metrics on reg
// (prometheus.DefaultRegisterer in the binaries, a fresh registry in tests).
func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		lastBlock: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "rwcg_etl_last_block",
			Help: "Last fully processed block (the processing watermark).",
		}),
		eventsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "rwcg_etl_events_total",
			Help: "Event logs processed, by event type.",
		}, []string{"type"}),
		batchDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: "rwcg_etl_batch_duration_seconds",
			Help: "Time to fetch, store and process one event batch.",
			// Batches range from millisecond no-op scans to multi-minute
			// backfill chunks; default buckets top out at 10s.
			Buckets: prometheus.ExponentialBuckets(0.01, 2, 16),
		}),
		reorgsTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "rwcg_etl_reorgs_total",
			Help: "Chain reorganizations detected and rolled back.",
		}),
		failuresTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "rwcg_etl_batch_failures_total",
			Help: "Batch attempts that failed, by pipeline stage.",
		}, []string{"stage"}),
	}
	reg.MustRegister(m.lastBlock, m.eventsTotal, m.batchDuration, m.reorgsTotal, m.failuresTotal)
	return m
}

func (m *Metrics) watermark(block int64) {
	if m == nil {
		return
	}
	m.lastBlock.Set(float64(block))
}

func (m *Metrics) eventProcessed(eventType string) {
	if m == nil {
		return
	}
	m.eventsTotal.WithLabelValues(eventType).Inc()
}

func (m *Metrics) batchProcessed(seconds float64) {
	if m == nil {
		return
	}
	m.batchDuration.Observe(seconds)
}

func (m *Metrics) reorgHandled() {
	if m == nil {
		return
	}
	m.reorgsTotal.Inc()
}

func (m *Metrics) batchFailed(stage string) {
	if m == nil {
		return
	}
	m.failuresTotal.WithLabelValues(stage).Inc()
}
