//go:build integration

package main

// End-to-end wiring test: run() against a real seeded Postgres and the fake
// Ethereum node, Twitter-only, with the notification watermark parked past
// the fixture events so the loop idles without external calls until the
// context cancels.

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/ethereum/go-ethereum/common"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

var (
	sharedRepo   *rwstore.Repo
	sharedPGEnv  map[string]string
	emptyDBName  string
	errSetupSkip error
)

func TestMain(m *testing.M) {
	os.Exit(testMainRun(m))
}

func testMainRun(m *testing.M) int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	db, stop, err := testdb.Start(ctx)
	if err != nil {
		if errors.Is(err, testdb.ErrContainerUnavailable) {
			errSetupSkip = err
			return m.Run()
		}
		fmt.Fprintf(os.Stderr, "notibot: starting test database: %v\n", err)
		return 1
	}
	defer stop()

	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		fmt.Fprintf(os.Stderr, "notibot: seeding fixture data: %v\n", err)
		return 1
	}
	sharedRepo = rwstore.NewRepo(store.NewFromPool(db.Pool))

	// A migration-free database on the same container for the
	// missing-schema failure path.
	if _, err := db.SQL.ExecContext(ctx, "CREATE DATABASE rwcg_empty"); err == nil {
		emptyDBName = "rwcg_empty"
	}

	u, err := url.Parse(db.ConnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "notibot: parsing container conn string: %v\n", err)
		return 1
	}
	password, _ := u.User.Password()
	sharedPGEnv = map[string]string{
		"PGSQL_USERNAME": u.User.Username(),
		"PGSQL_PASSWORD": password,
		"PGSQL_DATABASE": strings.TrimPrefix(u.Path, "/"),
		"PGSQL_HOST":     u.Host,
	}
	return m.Run()
}

// integrationEnv skips without Docker and wires HOME, PGSQL_* and RPC_URL
// (clearing the other configuration variables so ambient environment cannot
// leak in).
func integrationEnv(t *testing.T) string {
	t.Helper()
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	for k, v := range sharedPGEnv {
		t.Setenv(k, v)
	}
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("DATABASE_URL", "")
	t.Setenv("LOG_FORMAT", "")
	t.Setenv("LOG_LEVEL", "")
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	t.Setenv("RPC_URL", chain.URL())
	return home
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

func TestRunTwitterWiringEndToEnd(t *testing.T) {
	home := integrationEnv(t)

	// Park the watermark past all fixture events so the engine idles.
	status, err := sharedRepo.MessagingStatus(context.Background())
	if err != nil {
		t.Fatalf("MessagingStatus: %v", err)
	}
	parked := rwmodel.MsgStatus{TxId: status.TxId, EvtLogId: 1 << 40, BlockNum: status.BlockNum, TimeStamp: status.TimeStamp}
	if err := sharedRepo.UpdateMessagingStatus(context.Background(), &parked); err != nil {
		t.Fatalf("parking watermark: %v", err)
	}

	if err := os.MkdirAll(filepath.Join(home, "configs"), 0o750); err != nil {
		t.Fatal(err)
	}
	keysPath := filepath.Join(home, "configs", "tw.json")
	if err := os.WriteFile(keysPath, []byte(`{"ApiKey":"a","ApiSecret":"b","TokenKey":"c","TokenSecret":"d"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TWITTER_KEYS_FILE", "tw.json")

	var logBuf syncBuffer
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := run(ctx, os.Getenv, &logBuf, true, false); err != nil {
		t.Fatalf("run: %v", err)
	}

	// §8.3: the engine logs structured records on the stream (no ae_logs).
	if _, err := os.Stat(filepath.Join(home, "ae_logs")); !os.IsNotExist(err) {
		t.Errorf("legacy ae_logs directory was created (stat err=%v)", err)
	}
	logs := logBuf.String()
	for _, want := range []string{"effective configuration", "connected to ETH node", "resolved contracts", "loaded twitter keys", "rwbot starting"} {
		if !strings.Contains(logs, want) {
			t.Errorf("log stream missing %q\nlog:\n%s", want, logs)
		}
	}
}

func TestRunMissingTwitterKeysFails(t *testing.T) {
	integrationEnv(t)
	t.Setenv("TWITTER_KEYS_FILE", "does-not-exist.json")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := run(ctx, os.Getenv, io.Discard, true, false)
	if err == nil || !strings.Contains(err.Error(), "twitter keys") {
		t.Errorf("run = %v, want twitter keys failure", err)
	}
}

func TestRunDiscordWiringEndToEnd(t *testing.T) {
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	for k, v := range sharedPGEnv {
		t.Setenv(k, v)
	}
	home := t.TempDir()
	t.Setenv("HOME", home)

	// Fake chain serving the RandomWalk withdrawalAmount read the Discord
	// startup path performs.
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	addrs, err := sharedRepo.ContractAddrs(context.Background())
	if err != nil {
		t.Fatalf("ContractAddrs: %v", err)
	}
	wei, _ := new(big.Int).SetString("1500000000000000000", 10)
	chain.RegisterCall(common.HexToAddress(addrs.RandomWalk),
		testchain.MustContractStub(rwcontracts.RWalkMetaData.ABI).Return("withdrawalAmount", wei).Handler())
	t.Setenv("RPC_URL", chain.URL())

	// Park the watermark so no events are announced.
	status, err := sharedRepo.MessagingStatus(context.Background())
	if err != nil {
		t.Fatalf("MessagingStatus: %v", err)
	}
	parked := status
	parked.EvtLogId = 1 << 40
	if err := sharedRepo.UpdateMessagingStatus(context.Background(), &parked); err != nil {
		t.Fatalf("parking watermark: %v", err)
	}
	t.Cleanup(func() {
		if err := sharedRepo.UpdateMessagingStatus(context.Background(), &status); err != nil {
			t.Errorf("restoring watermark: %v", err)
		}
	})

	// Stub Discord REST API + factory override.
	stub := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/users/@me") {
			_, _ = w.Write([]byte(`{"id":"1","username":"bot"}`))
			return
		}
		_, _ = w.Write([]byte(`{"id":"1"}`))
	}))
	defer stub.Close()
	target, err := url.Parse(stub.URL)
	if err != nil {
		t.Fatal(err)
	}
	prevFactory := newDiscordClient
	newDiscordClient = func(cfg disgord.Config) *disgord.Client {
		cfg.HTTPClient = &http.Client{Transport: rewriteTransport{target: target}} //nolint:staticcheck // SA1019: the identity probe uses the deprecated field
		return disgord.New(cfg)
	}
	t.Cleanup(func() { newDiscordClient = prevFactory })

	if err := os.MkdirAll(filepath.Join(home, "configs"), 0o750); err != nil {
		t.Fatal(err)
	}
	discordKeysJSON := `{"TokenKey":"tok","ChannelId":101,"MainChannelId":100,` +
		`"MintStatsChanId":102,"PriceStatsChanId":103,"DateStatsChanId":104,"RewardStatsChanId":105}`
	if err := os.WriteFile(filepath.Join(home, "configs", "dc.json"), []byte(discordKeysJSON), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("DISCORD_KEYS_FILE", "dc.json")

	var logBuf syncBuffer
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := run(ctx, os.Getenv, &logBuf, false, true); err != nil {
		t.Fatalf("run --discord: %v", err)
	}

	logs := logBuf.String()
	for _, want := range []string{"loaded discord keys", "rwbot starting", "renamed discord channel"} {
		if !strings.Contains(logs, want) {
			t.Errorf("log stream missing %q\nlog:\n%s", want, logs)
		}
	}
}

func TestRunFailsOnEmptyDatabase(t *testing.T) {
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	if emptyDBName == "" {
		t.Skip("empty database not provisioned")
	}
	for k, v := range sharedPGEnv {
		t.Setenv(k, v)
	}
	t.Setenv("PGSQL_DATABASE", emptyDBName)
	home := t.TempDir()
	t.Setenv("HOME", home)
	chain := testchain.New(t)
	chain.EnsureBlock(1)
	t.Setenv("RPC_URL", chain.URL())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := run(ctx, os.Getenv, io.Discard, true, false)
	if err == nil || !strings.Contains(err.Error(), "contract addresses") {
		t.Errorf("run against an empty database = %v, want contract-address failure", err)
	}
}
