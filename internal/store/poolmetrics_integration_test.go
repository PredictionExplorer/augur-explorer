//go:build integration

package store_test

import (
	"context"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

// TestPoolCollectorObservesRealWorkload proves the collector reports live
// pool activity: after real queries the acquire counters advance, the
// configured maximum matches the production default and connections stay
// accounted for across the gauges.
func TestPoolCollectorObservesRealWorkload(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()

	st, err := store.New(ctx, configFromConnString(t, db.ConnString))
	if err != nil {
		t.Fatalf("store.New: %v", err)
	}
	defer st.Close()

	for range 3 {
		var one int
		if err := st.Pool().QueryRow(ctx, "SELECT 1").Scan(&one); err != nil {
			t.Fatalf("SELECT 1: %v", err)
		}
	}

	values := collectPoolMetrics(t, store.NewPoolCollector(st.Pool()))
	// store.New's startup ping plus the three queries all acquire.
	if got := values["rwcg_db_pool_acquires_total"]; got < 4 {
		t.Errorf("rwcg_db_pool_acquires_total = %v, want at least 4", got)
	}
	if got := values["rwcg_db_pool_new_conns_total"]; got < 1 {
		t.Errorf("rwcg_db_pool_new_conns_total = %v, want at least 1", got)
	}
	if got := values["rwcg_db_pool_max_conns"]; got != store.DefaultMaxConns {
		t.Errorf("rwcg_db_pool_max_conns = %v, want the production default %d", got, store.DefaultMaxConns)
	}
	if got := values["rwcg_db_pool_total_conns"]; got < 1 {
		t.Errorf("rwcg_db_pool_total_conns = %v, want at least 1", got)
	}
	if total, parts := values["rwcg_db_pool_total_conns"],
		values["rwcg_db_pool_acquired_conns"]+values["rwcg_db_pool_idle_conns"]+values["rwcg_db_pool_constructing_conns"]; total != parts {
		t.Errorf("total_conns = %v, want acquired+idle+constructing = %v", total, parts)
	}
}

// collectPoolMetrics gathers one scrape into name → value.
func collectPoolMetrics(t *testing.T, collector *store.PoolCollector) map[string]float64 {
	t.Helper()
	registry := prometheus.NewPedanticRegistry()
	if err := registry.Register(collector); err != nil {
		t.Fatalf("Register: %v", err)
	}
	families, err := registry.Gather()
	if err != nil {
		t.Fatalf("Gather: %v", err)
	}
	values := make(map[string]float64, len(families))
	for _, family := range families {
		for _, metric := range family.GetMetric() {
			values[family.GetName()] = metricValue(family.GetType(), metric)
		}
	}
	return values
}

func metricValue(kind dto.MetricType, metric *dto.Metric) float64 {
	if kind == dto.MetricType_COUNTER {
		return metric.GetCounter().GetValue()
	}
	return metric.GetGauge().GetValue()
}
