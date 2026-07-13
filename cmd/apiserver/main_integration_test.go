//go:build integration

package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

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
	t.Setenv("HTTP_PORT", "")
	t.Setenv("HTTPS_HOSTNAME", "")
	t.Setenv("HTTPS_EXTRA_LISTEN_ADDR", "")
	t.Setenv("NFT_ASSETS_DIR", "")
	t.Setenv("PGSQL_USERNAME", u.User.Username())
	t.Setenv("PGSQL_PASSWORD", password)
	t.Setenv("PGSQL_DATABASE", strings.TrimPrefix(u.Path, "/"))
	t.Setenv("PGSQL_HOST", u.Host)
}

// freePort reserves an ephemeral TCP port for the server under test.
func freePort(t *testing.T) string {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	_, port, _ := net.SplitHostPort(ln.Addr().String())
	_ = ln.Close()
	return port
}

// TestRunBootServeAndGracefulShutdown drives the full production wiring: log
// files, RPC dial, store, module construction, contract-state load, listener
// startup, live traffic on the plain listener, and the context-cancelled
// drain with the readiness flip.
func TestRunBootServeAndGracefulShutdown(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}

	chain := testchain.New(t)
	registerBootChainState(chain)
	setEnv(t, db.ConnString, chain.URL())
	port := freePort(t)
	t.Setenv("HTTP_PORT", port)
	t.Setenv("METRICS_ADDR", "127.0.0.1:0")

	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan error, 1)
	go func() { done <- run(runCtx, os.Getenv) }()

	base := "http://127.0.0.1:" + port
	waitForServer(t, done, cancel, base+"/healthz")

	// The full route table answers through the real listeners.
	for path, wantStatus := range map[string]int{
		"/healthz":                            http.StatusOK,
		"/readyz":                             http.StatusOK,
		"/api/cosmicgame/statistics/counters": http.StatusOK,
		"/api/randomwalk/floor_price":         http.StatusOK,
		"/api/v2/cosmicgame/rounds/1":         http.StatusOK,
	} {
		resp, err := http.Get(base + path) //nolint:noctx // short-lived test client
		if err != nil {
			t.Fatalf("GET %s: %v", path, err)
		}
		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if resp.StatusCode != wantStatus {
			t.Fatalf("GET %s = %d, want %d\n%s", path, resp.StatusCode, wantStatus, body)
		}
	}

	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("run returned %v, want nil on graceful shutdown", err)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("run did not return after cancellation")
	}

	// The legacy log layout was written under $HOME/ae_logs.
	home := os.Getenv("HOME")
	for _, name := range []string{"webserver_info.log", "webserver_error.log", "webserver-db.log"} {
		if _, err := os.Stat(home + "/ae_logs/" + name); err != nil { //nolint:gosec // test path under t.TempDir
			t.Fatalf("log file %s missing: %v", name, err)
		}
	}
}

// waitForServer polls the health probe until the listener answers; on
// timeout it stops the run and fails with its error.
func waitForServer(t *testing.T, done chan error, cancel context.CancelFunc, probe string) {
	t.Helper()
	deadline := time.Now().Add(60 * time.Second)
	for time.Now().Before(deadline) {
		select {
		case err := <-done:
			t.Fatalf("run exited during startup: %v", err)
		default:
		}
		resp, err := http.Get(probe) //nolint:noctx // startup poll
		if err == nil {
			_ = resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return
			}
		}
		time.Sleep(20 * time.Millisecond)
	}
	cancel()
	t.Fatalf("server never answered %s", probe)
}

// TestRunDisabledModulesStillServeMetadataEnvelopes boots with every module
// disabled: the process must come up, module routes must 404, and the bare
// /metadata dispatch answers the legacy unavailable envelopes.
func TestRunDisabledModulesStillServeMetadataEnvelopes(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}
	chain := testchain.New(t)
	setEnv(t, db.ConnString, chain.URL())
	port := freePort(t)
	t.Setenv("HTTP_PORT", port)
	t.Setenv("ENABLE_ROUTES_COSMICGAME", "false")
	t.Setenv("ENABLE_ROUTES_RANDOMWALK", "false")
	t.Setenv("ENABLE_ROUTES_FAQ", "false")

	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan error, 1)
	go func() { done <- run(runCtx, os.Getenv) }()

	base := "http://127.0.0.1:" + port
	waitForServer(t, done, cancel, base+"/healthz")

	get := func(path, host string) (int, string) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, base+path, nil)
		if err != nil {
			t.Fatal(err)
		}
		if host != "" {
			req.Host = host
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("GET %s: %v", path, err)
		}
		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		return resp.StatusCode, string(body)
	}

	if status, _ := get("/api/cosmicgame/statistics/counters", ""); status != http.StatusNotFound {
		t.Fatalf("disabled cosmicgame route = %d, want 404", status)
	}
	if status, _ := get("/api/randomwalk/floor_price", ""); status != http.StatusNotFound {
		t.Fatalf("disabled randomwalk route = %d, want 404", status)
	}
	if status, _ := get("/api/cosmicgame/faq/health", ""); status != http.StatusNotFound {
		t.Fatalf("disabled faq route = %d, want 404", status)
	}
	if status, body := get("/metadata/1", "nfts.cosmicsignature.com"); status != http.StatusBadRequest ||
		!strings.Contains(body, "CosmicGame module or database not available") {
		t.Fatalf("CS metadata dispatch = %d %s", status, body)
	}
	if status, body := get("/metadata/1", "www.randomwalknft.com"); status != http.StatusBadRequest ||
		!strings.Contains(body, "Database link wasn't configured") {
		t.Fatalf("RW metadata dispatch = %d %s", status, body)
	}

	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("run returned %v", err)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("run did not return after cancellation")
	}
}

// TestRunStartupFailures pins every startup error return of run.
func TestRunStartupFailures(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()

	t.Run("missing RPC_URL", func(t *testing.T) {
		setEnv(t, db.ConnString, "")
		err := run(ctx, os.Getenv)
		if err == nil || !strings.Contains(err.Error(), "RPC URL of Ethereum node is not set") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("bad RPC URL", func(t *testing.T) {
		setEnv(t, db.ConnString, "://not-a-url")
		err := run(ctx, os.Getenv)
		if err == nil || !strings.Contains(err.Error(), "connection to RPC service") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("individual log files unopenable", func(t *testing.T) {
		chain := testchain.New(t)
		for _, name := range []string{"webserver_info.log", "webserver_error.log", "webserver-db.log"} {
			setEnv(t, db.ConnString, chain.URL())
			home := os.Getenv("HOME")
			if err := os.MkdirAll(home+"/ae_logs/"+name, 0o750); err != nil { //nolint:gosec // test path under t.TempDir
				t.Fatal(err)
			}
			err := run(ctx, os.Getenv)
			if err == nil || !strings.Contains(err.Error(), "can't open log file") {
				t.Fatalf("%s: err = %v", name, err)
			}
		}
	})

	t.Run("database connect failure", func(t *testing.T) {
		chain := testchain.New(t)
		setEnv(t, db.ConnString, chain.URL())
		cancelled, cancel := context.WithCancel(ctx)
		cancel()
		err := run(cancelled, os.Getenv)
		if err == nil || !strings.Contains(err.Error(), "can't connect to PostgreSQL database") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("empty contract registry is fatal with hint", func(t *testing.T) {
		// Migrated but not fixture-seeded: cg_contracts is empty.
		chain := testchain.New(t)
		setEnv(t, db.ConnString, chain.URL())
		err := run(ctx, os.Getenv)
		if err == nil || !strings.Contains(err.Error(), "CosmicGame module init failed") ||
			!strings.Contains(err.Error(), "ENABLE_ROUTES_COSMICGAME=false") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("no listeners configured", func(t *testing.T) {
		if err := testfixtures.Apply(ctx, db.SQL); err != nil {
			t.Fatalf("applying fixtures: %v", err)
		}
		chain := testchain.New(t)
		registerBootChainState(chain)
		setEnv(t, db.ConnString, chain.URL())
		err := run(ctx, os.Getenv)
		if err == nil || !strings.Contains(err.Error(), "no listeners configured") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("plain listener bind failure", func(t *testing.T) {
		chain := testchain.New(t)
		registerBootChainState(chain)
		setEnv(t, db.ConnString, chain.URL())
		// Squat the wildcard port first so the server's bind fails.
		ln, err := net.Listen("tcp", ":0")
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { _ = ln.Close() })
		_, port, _ := net.SplitHostPort(ln.Addr().String())
		t.Setenv("HTTP_PORT", port)
		err = run(ctx, os.Getenv)
		if err == nil || !strings.Contains(err.Error(), "HTTP bind failed") {
			t.Fatalf("err = %v", err)
		}
	})
}

// TestEnvBoolDefaultTrue pins the tri-state parsing main relies on.
func TestEnvBoolDefaultTrue(t *testing.T) {
	cases := []struct {
		value string
		want  bool
	}{
		{"", true}, {"  ", true}, {"true", true}, {"1", true}, {"yes", true},
		{"anything", true}, {"false", false}, {"FALSE", false}, {"0", false},
		{"no", false}, {"off", false}, {" Off ", false},
	}
	for _, tc := range cases {
		getenv := func(string) string { return tc.value }
		if got := envBoolDefaultTrue(getenv, "K"); got != tc.want {
			t.Errorf("envBoolDefaultTrue(%q) = %v, want %v", tc.value, got, tc.want)
		}
	}
}

// registerBootChainState registers the minimal eth_call surface the
// contract-state initial load touches, so module construction succeeds
// against the fake chain (values need not be fixture-coherent here; the
// parity suite pins response bodies).
func registerBootChainState(chain *testchain.Chain) {
	chain.EnsureBlock(142)
	fmt.Println("boot chain state registered at", chain.URL())
}
