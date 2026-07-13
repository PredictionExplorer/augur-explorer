//go:build integration

package main

import (
	"context"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// setEnv wires HOME, RPC_URL and PGSQL_* for one run.
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
	t.Setenv("PGSQL_USERNAME", u.User.Username())
	t.Setenv("PGSQL_PASSWORD", password)
	t.Setenv("PGSQL_DATABASE", strings.TrimPrefix(u.Path, "/"))
	t.Setenv("PGSQL_HOST", u.Host)
}

// TestRunBootAndGracefulShutdown drives the full production wiring: log
// files, RPC dial, store, contract bootstrap, handler set, param sync,
// engine catch-up and context-cancelled shutdown.
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

	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan error, 1)
	go func() { done <- run(runCtx, os.Getenv, prometheus.NewRegistry(), prometheus.NewRegistry()) }()

	// The startup chain sync allocates a synthetic event log at the current
	// tip, so the engine resumes already caught up. Wait for the engine to
	// resolve its watermark, then advance the chain: completing that batch
	// proves the whole loop works and persists the watermark.
	infoLog := os.Getenv("HOME") + "/ae_logs/cosmicgame_info.log"
	waitForCondition(t, done, cancel, infoLog, "engine start", func() bool {
		info, _ := os.ReadFile(infoLog) //nolint:gosec // test path under t.TempDir
		return strings.Contains(string(info), "resuming after last processed block")
	})
	chain.EnsureBlock(200)

	waitForCondition(t, done, cancel, infoLog, "watermark at chain tip", func() bool {
		var lastBlock int64
		_ = db.Pool.QueryRow(ctx, "SELECT last_block_num FROM cg_proc_status").Scan(&lastBlock)
		return lastBlock == 200
	})

	cancel()
	if err := awaitRun(done); err != nil {
		t.Fatalf("run returned %v, want nil on graceful shutdown", err)
	}

	// The legacy two-file log layout was written under $HOME/ae_logs.
	home := os.Getenv("HOME")
	for _, name := range []string{"cosmicgame_info.log", "cosmicgame_error.log", "cosmicgame_db.log"} {
		if _, err := os.Stat(home + "/ae_logs/" + name); err != nil { //nolint:gosec // test path under t.TempDir
			t.Fatalf("log file %s missing: %v", name, err)
		}
	}
	info, err := os.ReadFile(home + "/ae_logs/cosmicgame_info.log") //nolint:gosec // test path under t.TempDir
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(info), "Connected to ETH node") {
		t.Fatalf("info log = %q", string(info))
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

// waitForCondition polls cond for up to a minute; on timeout it stops the
// run and fails with the info log attached.
func waitForCondition(t *testing.T, done chan error, cancel context.CancelFunc, infoLog, what string, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(60 * time.Second)
	for {
		if cond() {
			return
		}
		if time.Now().After(deadline) {
			cancel()
			runErr := awaitRun(done)
			info, _ := os.ReadFile(infoLog) //nolint:gosec // test path under t.TempDir
			t.Fatalf("timed out waiting for %s (run=%v)\ninfo log:\n%s", what, runErr, info)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func TestRunFailures(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()

	t.Run("bad rpc url", func(t *testing.T) {
		setEnv(t, db.ConnString, "://not-a-url")
		err := run(ctx, os.Getenv, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), "dialing RPC node") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("unusable log dir", func(t *testing.T) {
		chain := testchain.New(t)
		setEnv(t, db.ConnString, chain.URL())
		t.Setenv("HOME", "/dev/null")
		err := run(ctx, os.Getenv, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), "can't open log file") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("individual log files unopenable", func(t *testing.T) {
		// A directory squatting on a log file name makes only that
		// specific open fail, covering each openAppendLog error return.
		chain := testchain.New(t)
		for _, name := range []string{"cosmicgame_error.log", "cosmicgame_db.log"} {
			setEnv(t, db.ConnString, chain.URL())
			home := os.Getenv("HOME")
			if err := os.MkdirAll(home+"/ae_logs/"+name, 0o750); err != nil { //nolint:gosec // test path under t.TempDir
				t.Fatal(err)
			}
			err := run(ctx, os.Getenv, prometheus.NewRegistry(), prometheus.NewRegistry())
			if err == nil || !strings.Contains(err.Error(), "can't open log file") {
				t.Fatalf("%s: err = %v", name, err)
			}
		}
	})

	t.Run("database connect failure", func(t *testing.T) {
		// A pre-cancelled context fails the store's bounded connect retry
		// immediately (HTTP RPC dial is lazy and unaffected).
		chain := testchain.New(t)
		setEnv(t, db.ConnString, chain.URL())
		cancelled, cancel := context.WithCancel(ctx)
		cancel()
		err := run(cancelled, os.Getenv, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), "can't connect to PostgreSQL database") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("empty contract registry", func(t *testing.T) {
		// The database is migrated but not fixture-seeded: cg_contracts is
		// empty, so the bootstrap fails fast.
		chain := testchain.New(t)
		chain.EnsureBlock(1)
		setEnv(t, db.ConnString, chain.URL())
		err := run(ctx, os.Getenv, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), "contract address bootstrap failed") {
			t.Fatalf("err = %v", err)
		}
	})
}

func TestRunMetricsServerFailure(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}
	chain := testchain.New(t)
	chain.EnsureBlock(150)
	setEnv(t, db.ConnString, chain.URL())
	t.Setenv("METRICS_ADDR", "256.256.256.256:1")

	err := run(ctx, os.Getenv, prometheus.NewRegistry(), prometheus.NewRegistry())
	if err == nil || !strings.Contains(err.Error(), "can't start metrics server") {
		t.Fatalf("err = %v", err)
	}
}

// TestProgressAdapter pins the cg_proc_status watermark adapter directly:
// the lazily created zero row, persisted updates, and error propagation.
func TestProgressAdapter(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	p := cgProgress{repo: cgstore.NewRepo(store.NewFromPool(db.Pool))}

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

func TestRunContractSyncFailure(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}
	chain := testchain.New(t)
	chain.EnsureBlock(150)
	setEnv(t, db.ConnString, chain.URL())

	// The startup sync allocates its synthetic event log at the chain tip
	// (a HeaderByNumber read); failing that read makes the whole sync (and
	// run) fail.
	chain.FailNextRPC("eth_getBlockByNumber", "head read refused")

	err := run(ctx, os.Getenv, prometheus.NewRegistry(), prometheus.NewRegistry())
	if err == nil || !strings.Contains(err.Error(), "contract param chain sync failed") {
		t.Fatalf("err = %v", err)
	}
}
