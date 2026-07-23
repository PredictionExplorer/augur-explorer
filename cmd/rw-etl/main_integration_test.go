//go:build integration

package main

import (
	"bytes"
	"context"
	"io"
	"net/url"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// setEnv wires RPC_URL and PGSQL_* for one run and clears everything else
// the typed configuration reads.
func setEnv(t *testing.T, connString, rpcURL string) {
	t.Helper()
	u, err := url.Parse(connString)
	if err != nil {
		t.Fatal(err)
	}
	password, _ := u.User.Password()
	t.Setenv("HOME", t.TempDir())
	t.Setenv("RPC_URL", rpcURL)
	t.Setenv("METRICS_ADDR", "")
	t.Setenv("LOG_FORMAT", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("DATABASE_URL", "")
	t.Setenv("PGSQL_USERNAME", u.User.Username())
	t.Setenv("PGSQL_PASSWORD", password)
	t.Setenv("PGSQL_DATABASE", strings.TrimPrefix(u.Path, "/"))
	t.Setenv("PGSQL_HOST", u.Host)
}

// syncBuffer is a mutex-guarded log sink: the engine goroutines log
// concurrently with the test's readers.
type syncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *syncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

// TestRunBootAndGracefulShutdown drives the full production wiring: typed
// configuration, structured logging, RPC dial, store, contract bootstrap,
// handler set, engine catch-up and context-cancelled shutdown.
func TestRunBootAndGracefulShutdown(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}

	chain := testchain.New(t)
	chain.EnsureBlock(150)
	setEnv(t, db.ConnString, chain.URL())
	t.Setenv("METRICS_ADDR", "127.0.0.1:0")

	var logBuf syncBuffer
	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan error, 1)
	reg := prometheus.NewRegistry()
	go func() { done <- run(runCtx, os.Getenv, &logBuf, reg, reg) }()

	// The engine has fully booted once the watermark reaches the chain tip.
	deadline := time.Now().Add(60 * time.Second)
	for {
		var lastBlock int64
		_ = db.Pool.QueryRow(ctx, "SELECT last_block FROM rw_proc_status").Scan(&lastBlock)
		if lastBlock == 150 {
			break
		}
		if time.Now().After(deadline) {
			cancel()
			runErr := awaitRun(done)
			t.Fatalf("watermark never reached the chain tip (last=%d, run=%v)\nlog stream:\n%s", lastBlock, runErr, logBuf.String())
		}
		time.Sleep(20 * time.Millisecond)
	}

	// The production registry carries the DB-pool series while the process
	// runs, advanced by the real ingestion workload.
	if got := gatheredValue(t, reg, "rwcg_db_pool_acquires_total"); got < 1 {
		t.Errorf("rwcg_db_pool_acquires_total = %v, want at least 1", got)
	}

	cancel()
	if err := awaitRun(done); err != nil {
		t.Fatalf("run returned %v, want nil on graceful shutdown", err)
	}

	// Shutdown unregisters the pool collector (no stale series for a
	// closed pool).
	if families, err := reg.Gather(); err == nil {
		for _, family := range families {
			if strings.HasPrefix(family.GetName(), "rwcg_db_pool_") {
				t.Errorf("%s still registered after shutdown", family.GetName())
			}
		}
	}

	// §8.3: structured records on the log stream replace the legacy
	// $HOME/ae_logs two-file layout.
	if _, err := os.Stat(os.Getenv("HOME") + "/ae_logs"); !os.IsNotExist(err) { //nolint:gosec // test path under t.TempDir
		t.Errorf("legacy ae_logs directory was created (stat err=%v)", err)
	}
	logs := logBuf.String()
	for _, want := range []string{"build info", "effective configuration", "Connected to ETH node", "PGSQL_PASSWORD=[set]"} {
		if !strings.Contains(logs, want) {
			t.Errorf("log stream missing %q:\n%s", want, logs)
		}
	}
}

func awaitRun(done chan error) error {
	select {
	case err := <-done:
		return err
	case <-time.After(60 * time.Second):
		return context.DeadlineExceeded
	}
}

// gatheredValue returns the single sample of the named series (counter or
// gauge), or fails if it is not currently registered.
func gatheredValue(t *testing.T, reg *prometheus.Registry, name string) float64 {
	t.Helper()
	families, err := reg.Gather()
	if err != nil {
		t.Fatalf("Gather: %v", err)
	}
	for _, family := range families {
		if family.GetName() != name {
			continue
		}
		metric := family.GetMetric()[0]
		if counter := metric.GetCounter(); counter != nil {
			return counter.GetValue()
		}
		return metric.GetGauge().GetValue()
	}
	t.Fatalf("series %s not registered", name)
	return 0
}

// TestProgressAdapter pins the rw_proc_status watermark adapter directly:
// the persisted round trip and error propagation.
func TestProgressAdapter(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	p := rwProgress{repo: rwstore.NewRepo(store.NewFromPool(db.Pool))}

	last, err := p.LastBlock(ctx)
	if err != nil || last != 0 {
		t.Fatalf("LastBlock = %d, %v", last, err)
	}
	if err := p.SetLastBlock(ctx, 42); err != nil {
		t.Fatal(err)
	}
	last, err = p.LastBlock(ctx)
	if err != nil || last != 42 {
		t.Fatalf("LastBlock after set = %d, %v", last, err)
	}

	cancelled, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := p.LastBlock(cancelled); err == nil {
		t.Fatal("cancelled LastBlock must error")
	}
	if err := p.SetLastBlock(cancelled, 43); err == nil {
		t.Fatal("cancelled SetLastBlock must error")
	}
}

func TestRunFailures(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()

	t.Run("missing RPC_URL fails at config load", func(t *testing.T) {
		setEnv(t, db.ConnString, "")
		err := run(ctx, os.Getenv, io.Discard, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), "RPC_URL: required but not set") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("bad rpc url", func(t *testing.T) {
		setEnv(t, db.ConnString, "://not-a-url")
		err := run(ctx, os.Getenv, io.Discard, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), "dialing RPC node") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("empty contract registry", func(t *testing.T) {
		// Migrated but not fixture-seeded: rw_contracts is empty, so the
		// bootstrap fails fast.
		chain := testchain.New(t)
		chain.EnsureBlock(1)
		setEnv(t, db.ConnString, chain.URL())
		err := run(ctx, os.Getenv, io.Discard, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), "contract address bootstrap failed") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("database connect failure", func(t *testing.T) {
		// A pre-cancelled context fails the store's bounded connect retry
		// immediately (HTTP RPC dial is lazy and unaffected).
		chain := testchain.New(t)
		setEnv(t, db.ConnString, chain.URL())
		cancelled, cancel := context.WithCancel(ctx)
		cancel()
		err := run(cancelled, os.Getenv, io.Discard, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), "can't connect to PostgreSQL database") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("metrics server failure", func(t *testing.T) {
		if err := testfixtures.Apply(ctx, db.SQL); err != nil {
			t.Fatalf("applying fixtures: %v", err)
		}
		chain := testchain.New(t)
		chain.EnsureBlock(150)
		setEnv(t, db.ConnString, chain.URL())
		t.Setenv("METRICS_ADDR", "256.256.256.256:1")
		err := run(ctx, os.Getenv, io.Discard, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), "can't start metrics server") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("pool metrics registration conflict", func(t *testing.T) {
		// A registry already carrying one of the pool series makes the
		// collector registration fail loudly instead of serving a partial
		// metric set.
		chain := testchain.New(t)
		setEnv(t, db.ConnString, chain.URL())
		reg := prometheus.NewRegistry()
		reg.MustRegister(prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "rwcg_db_pool_max_conns",
			Help: "squats the pool collector's series",
		}))
		err := run(ctx, os.Getenv, io.Discard, reg, reg)
		if err == nil || !strings.Contains(err.Error(), "registering db pool metrics") {
			t.Fatalf("err = %v", err)
		}
	})
}
