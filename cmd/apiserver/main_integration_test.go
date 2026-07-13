//go:build integration

package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// setEnv wires HOME, RPC_URL and PGSQL_* for one run (and clears every other
// variable the typed configuration reads, so ambient environment cannot
// leak into a test run).
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
	t.Setenv("NFT_ASSETS_ROOT", "")
	t.Setenv("STATIC_ABI_DIR", "")
	t.Setenv("ENABLE_ROUTES_COSMICGAME", "")
	t.Setenv("ENABLE_ROUTES_RANDOMWALK", "")
	t.Setenv("ENABLE_ROUTES_FAQ", "")
	t.Setenv("LOG_FORMAT", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("DATABASE_URL", "")
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

// TestRunBootServeAndGracefulShutdown drives the full production wiring:
// typed configuration, structured logging, RPC dial, store, module
// construction, contract-state load, listener startup, live traffic on the
// plain listener, and the context-cancelled drain with the readiness flip.
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
	t.Setenv("LOG_FORMAT", "json")

	var logBuf syncBuffer
	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan error, 1)
	go func() { done <- run(runCtx, os.Getenv, &logBuf) }()

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
		resp, err := http.Get(base + path)
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

	// §8.3: no ae_logs files anymore — the process emits structured JSON
	// records on its log stream (journald owns persistence in production).
	if _, err := os.Stat(os.Getenv("HOME") + "/ae_logs"); !os.IsNotExist(err) { //nolint:gosec // test path under t.TempDir
		t.Errorf("legacy ae_logs directory was created (stat err=%v)", err)
	}
	assertJSONLogRecords(t, logBuf.String(),
		"effective configuration", // startup config record (secrets redacted)
		"request",                 // access log through the shared middleware
		"shutdown: complete",
	)
	// The startup record must redact the database password.
	for line := range strings.Lines(logBuf.String()) {
		var rec map[string]any
		if json.Unmarshal([]byte(strings.TrimSpace(line)), &rec) != nil || rec["msg"] != "effective configuration" {
			continue
		}
		if got := rec["PGSQL_PASSWORD"]; got != "[set]" {
			t.Errorf("PGSQL_PASSWORD rendered %v, want [set]", got)
		}
	}
}

// syncBuffer is a mutex-guarded bytes.Buffer: run's HTTP goroutines log
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

// assertJSONLogRecords proves the stream is one JSON record per line and
// that records with the wanted msg values are present.
func assertJSONLogRecords(t *testing.T, stream string, wantMsgs ...string) {
	t.Helper()
	found := map[string]bool{}
	for line := range strings.Lines(stream) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var rec map[string]any
		if err := json.Unmarshal([]byte(line), &rec); err != nil {
			t.Fatalf("log line is not JSON: %v (%s)", err, line)
		}
		if msg, ok := rec["msg"].(string); ok {
			found[msg] = true
		}
	}
	for _, want := range wantMsgs {
		if !found[want] {
			t.Errorf("log stream missing %q record:\n%s", want, stream)
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
		resp, err := http.Get(probe) //nolint:gosec // G107: startup poll against a local test listener URL
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
	go func() { done <- run(runCtx, os.Getenv, io.Discard) }()

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
		t.Setenv("HTTP_PORT", "0")
		err := run(ctx, os.Getenv, io.Discard)
		if err == nil || !strings.Contains(err.Error(), "RPC_URL: required but not set") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("configuration problems are aggregated", func(t *testing.T) {
		setEnv(t, db.ConnString, "")
		t.Setenv("ENABLE_ROUTES_FAQ", "maybe")
		err := run(ctx, os.Getenv, io.Discard)
		if err == nil {
			t.Fatal("run accepted a broken configuration")
		}
		for _, want := range []string{
			"RPC_URL: required but not set",
			"ENABLE_ROUTES_FAQ: cannot parse",
			"HTTP_PORT: no listeners configured",
		} {
			if !strings.Contains(err.Error(), want) {
				t.Errorf("aggregated error missing %q:\n%v", want, err)
			}
		}
	})

	t.Run("bad RPC URL", func(t *testing.T) {
		setEnv(t, db.ConnString, "://not-a-url")
		t.Setenv("HTTP_PORT", "0")
		err := run(ctx, os.Getenv, io.Discard)
		if err == nil || !strings.Contains(err.Error(), "connection to RPC service") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("database connect failure", func(t *testing.T) {
		chain := testchain.New(t)
		setEnv(t, db.ConnString, chain.URL())
		t.Setenv("HTTP_PORT", "0")
		cancelled, cancel := context.WithCancel(ctx)
		cancel()
		err := run(cancelled, os.Getenv, io.Discard)
		if err == nil || !strings.Contains(err.Error(), "can't connect to PostgreSQL database") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("empty contract registry is fatal with hint", func(t *testing.T) {
		// Migrated but not fixture-seeded: cg_contracts is empty.
		chain := testchain.New(t)
		setEnv(t, db.ConnString, chain.URL())
		t.Setenv("HTTP_PORT", "0")
		err := run(ctx, os.Getenv, io.Discard)
		if err == nil || !strings.Contains(err.Error(), "CosmicGame module init failed") ||
			!strings.Contains(err.Error(), "ENABLE_ROUTES_COSMICGAME=false") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("no listeners configured fails at load", func(t *testing.T) {
		if err := testfixtures.Apply(ctx, db.SQL); err != nil {
			t.Fatalf("applying fixtures: %v", err)
		}
		chain := testchain.New(t)
		registerBootChainState(chain)
		setEnv(t, db.ConnString, chain.URL())
		err := run(ctx, os.Getenv, io.Discard)
		if err == nil || !strings.Contains(err.Error(), "no listeners configured") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("plain listener bind failure", func(t *testing.T) {
		chain := testchain.New(t)
		registerBootChainState(chain)
		setEnv(t, db.ConnString, chain.URL())
		// Squat the wildcard port first so the server's bind fails.
		ln, err := net.Listen("tcp", ":0") // #nosec G102 -- deliberately squats the wildcard port for the bind-failure case
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { _ = ln.Close() })
		_, port, _ := net.SplitHostPort(ln.Addr().String())
		t.Setenv("HTTP_PORT", port)
		err = run(ctx, os.Getenv, io.Discard)
		if err == nil || !strings.Contains(err.Error(), "HTTP bind failed") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("HTTPS-only with unreadable certificates starts no listener", func(t *testing.T) {
		chain := testchain.New(t)
		registerBootChainState(chain)
		setEnv(t, db.ConnString, chain.URL())
		t.Setenv("HTTPS_HOSTNAME", "127.0.0.1:0")
		t.Setenv("TLS_CERT_FILE", "/nonexistent/cert.pem")
		t.Setenv("TLS_KEY_FILE", "/nonexistent/key.pem")
		t.Setenv("TLS_CERT_FILE_2", "/nonexistent/cert2.pem")
		t.Setenv("TLS_KEY_FILE_2", "/nonexistent/key2.pem")
		err := run(ctx, os.Getenv, io.Discard)
		if err == nil || !strings.Contains(err.Error(), "no listeners started") {
			t.Fatalf("err = %v", err)
		}
	})
}

// writeSelfSignedCert generates a self-signed localhost certificate pair for
// the TLS listener tests and returns the cert/key file paths.
func writeSelfSignedCert(t *testing.T, dir, name string) (certFile, keyFile string) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	tmpl := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "127.0.0.1"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}
	der, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	keyDER, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	certFile = filepath.Join(dir, name+".crt")
	keyFile = filepath.Join(dir, name+".key")
	if err := os.WriteFile(certFile, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(keyFile, pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyDER}), 0o600); err != nil {
		t.Fatal(err)
	}
	return certFile, keyFile
}

// TestRunServesTLSListeners boots the HTTPS path end to end: both certificate
// pairs load, the primary and extra TLS listeners bind, answer requests, and
// drain gracefully.
func TestRunServesTLSListeners(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}
	chain := testchain.New(t)
	registerBootChainState(chain)
	setEnv(t, db.ConnString, chain.URL())

	certDir := t.TempDir()
	cert1, key1 := writeSelfSignedCert(t, certDir, "primary")
	cert2, key2 := writeSelfSignedCert(t, certDir, "second")
	t.Setenv("HTTPS_HOSTNAME", "127.0.0.1:0")
	t.Setenv("HTTPS_EXTRA_LISTEN_ADDR", "127.0.0.1:0")
	t.Setenv("TLS_CERT_FILE", cert1)
	t.Setenv("TLS_KEY_FILE", key1)
	t.Setenv("TLS_CERT_FILE_2", cert2)
	t.Setenv("TLS_KEY_FILE_2", key2)
	t.Setenv("LOG_FORMAT", "json")

	var logBuf syncBuffer
	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan error, 1)
	go func() { done <- run(runCtx, os.Getenv, &logBuf) }()

	// The bound addresses come out of the structured startup records.
	tlsAddrs := func() []string {
		var addrs []string
		for line := range strings.Lines(logBuf.String()) {
			var rec map[string]any
			if json.Unmarshal([]byte(strings.TrimSpace(line)), &rec) != nil {
				continue
			}
			if rec["msg"] == "HTTPS bound and listening" {
				if addr, ok := rec["addr"].(string); ok {
					addrs = append(addrs, addr)
				}
			}
		}
		return addrs
	}
	deadline := time.Now().Add(60 * time.Second)
	for len(tlsAddrs()) < 2 {
		select {
		case err := <-done:
			t.Fatalf("run exited during TLS startup: %v\n%s", err, logBuf.String())
		default:
		}
		if time.Now().After(deadline) {
			cancel()
			t.Fatalf("TLS listeners never bound:\n%s", logBuf.String())
		}
		time.Sleep(20 * time.Millisecond)
	}

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec // G402: self-signed test certificate
	}}
	for _, addr := range tlsAddrs() {
		resp, err := client.Get("https://" + addr + "/healthz")
		if err != nil {
			t.Fatalf("GET over TLS on %s: %v", addr, err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("TLS /healthz on %s = %d", addr, resp.StatusCode)
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
}

// TestRunExtraTLSListenerBindFailureIsNonFatal squats the extra TLS address:
// the failed bind is logged, the primary TLS listener still serves, and the
// process drains cleanly.
func TestRunExtraTLSListenerBindFailureIsNonFatal(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}
	chain := testchain.New(t)
	registerBootChainState(chain)
	setEnv(t, db.ConnString, chain.URL())

	cert1, key1 := writeSelfSignedCert(t, t.TempDir(), "primary")
	squatted, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = squatted.Close() })

	t.Setenv("HTTPS_HOSTNAME", "127.0.0.1:0")
	t.Setenv("HTTPS_EXTRA_LISTEN_ADDR", squatted.Addr().String())
	t.Setenv("TLS_CERT_FILE", cert1)
	t.Setenv("TLS_KEY_FILE", key1)

	var logBuf syncBuffer
	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan error, 1)
	go func() { done <- run(runCtx, os.Getenv, &logBuf) }()

	deadline := time.Now().Add(60 * time.Second)
	for !strings.Contains(logBuf.String(), "HTTPS bind failed") ||
		!strings.Contains(logBuf.String(), "HTTPS bound and listening") {
		select {
		case err := <-done:
			t.Fatalf("run exited during startup: %v\n%s", err, logBuf.String())
		default:
		}
		if time.Now().After(deadline) {
			cancel()
			t.Fatalf("expected one bound and one failed TLS listener:\n%s", logBuf.String())
		}
		time.Sleep(20 * time.Millisecond)
	}

	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("run returned %v, want nil (primary listener served)", err)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("run did not return after cancellation")
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
