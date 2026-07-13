package indexer

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestMetricsRecording(t *testing.T) {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	m.watermark(1234)
	m.eventProcessed("BidPlaced")
	m.eventProcessed("BidPlaced")
	m.eventProcessed("other")
	m.batchProcessed(0.25)
	m.reorgHandled()
	m.batchFailed("fetch")

	if got := testutil.ToFloat64(m.lastBlock); got != 1234 {
		t.Errorf("rwcg_etl_last_block = %v, want 1234", got)
	}
	if got := testutil.ToFloat64(m.eventsTotal.WithLabelValues("BidPlaced")); got != 2 {
		t.Errorf("rwcg_etl_events_total{type=BidPlaced} = %v, want 2", got)
	}
	if got := testutil.ToFloat64(m.eventsTotal.WithLabelValues("other")); got != 1 {
		t.Errorf("rwcg_etl_events_total{type=other} = %v, want 1", got)
	}
	if got := testutil.ToFloat64(m.reorgsTotal); got != 1 {
		t.Errorf("rwcg_etl_reorgs_total = %v, want 1", got)
	}
	if got := testutil.ToFloat64(m.failuresTotal.WithLabelValues("fetch")); got != 1 {
		t.Errorf("rwcg_etl_batch_failures_total{stage=fetch} = %v, want 1", got)
	}
	if got := testutil.CollectAndCount(m.batchDuration, "rwcg_etl_batch_duration_seconds"); got != 1 {
		t.Errorf("batch duration histogram series = %d, want 1", got)
	}
}

func TestNilMetricsAreSafe(t *testing.T) {
	var m *Metrics
	m.watermark(1)
	m.eventProcessed("x")
	m.batchProcessed(0.1)
	m.reorgHandled()
	m.batchFailed("fetch")
}

func TestStartMetricsServerServesMetricsAndPprof(t *testing.T) {
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)
	m.watermark(77)

	srv, addr, err := StartMetricsServer(t.Context(), "127.0.0.1:0", reg, slog.New(slog.DiscardHandler))
	if err != nil {
		t.Fatalf("StartMetricsServer: %v", err)
	}
	defer func() { _ = srv.Close() }()

	client := &http.Client{Timeout: 5 * time.Second}
	body := httpGet(t, client, fmt.Sprintf("http://%s/metrics", addr))
	if !strings.Contains(body, "rwcg_etl_last_block 77") {
		t.Errorf("/metrics missing the watermark gauge:\n%s", body)
	}

	if got := httpGet(t, client, fmt.Sprintf("http://%s/debug/pprof/cmdline", addr)); got == "" {
		t.Error("/debug/pprof/cmdline returned an empty body")
	}
}

func TestStartMetricsServerRejectsBadAddr(t *testing.T) {
	if _, _, err := StartMetricsServer(t.Context(), "256.256.256.256:99999", prometheus.NewRegistry(), slog.New(slog.DiscardHandler)); err == nil {
		t.Fatal("expected a listen error for an invalid address")
	}
}

func httpGet(t *testing.T, client *http.Client, url string) string {
	t.Helper()
	resp, err := client.Get(url)
	if err != nil {
		t.Fatalf("GET %s: %v", url, err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET %s: status %d", url, resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading %s body: %v", url, err)
	}
	return string(b)
}
