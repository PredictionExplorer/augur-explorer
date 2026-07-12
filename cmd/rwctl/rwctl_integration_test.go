//go:build integration

package main

// Integration tests against a real seeded Postgres: the database-backed
// helpers (mint/transfer verification, rank recomputation) and a full
// scan-mints command run wired to both the fake chain and the container
// database through the production environment variables.

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5/pgxpool"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// Fixture handles (see internal/testfixtures/seed/*.sql).
const fixtureRWalkAddr = "0x8000000000000000000000000000000000000008"

var (
	sharedRepo       *rwstore.Repo
	sharedPGEnv      map[string]string
	sharedConnString string
	errSetupSkip     error
)

func TestMain(m *testing.M) {
	os.Exit(runMain(m))
}

func runMain(m *testing.M) int {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	db, stop, err := testdb.Start(ctx)
	if err != nil {
		if errors.Is(err, testdb.ErrContainerUnavailable) {
			errSetupSkip = err
			return m.Run()
		}
		fmt.Fprintf(os.Stderr, "rwctl: starting test database: %v\n", err)
		return 1
	}
	defer stop()

	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		fmt.Fprintf(os.Stderr, "rwctl: seeding fixture data: %v\n", err)
		return 1
	}
	sharedRepo = rwstore.NewRepo(store.NewFromPool(db.Pool))
	sharedConnString = db.ConnString

	// The PGSQL_* environment the production connectRWStorage reads, derived
	// from the container's connection URL.
	u, err := url.Parse(db.ConnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rwctl: parsing container conn string: %v\n", err)
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

// integrationRepo skips the test when Docker is unavailable.
func integrationRepo(t *testing.T) *rwstore.Repo {
	t.Helper()
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	if sharedRepo == nil {
		t.Fatal("rwctl harness not initialized (TestMain did not run?)")
	}
	return sharedRepo
}

// setPGEnv points the production PGSQL_* variables at the container.
func setPGEnv(t *testing.T) {
	t.Helper()
	integrationRepo(t)
	for k, v := range sharedPGEnv {
		t.Setenv(k, v)
	}
}

// mintLog builds a MintEvent log for the given token id.
func mintLog(blockNum uint64, tokenID int64, txHash common.Hash) *types.Log {
	return &types.Log{
		Address: common.HexToAddress(fixtureRWalkAddr),
		Topics: []common.Hash{
			mintEventTopic,
			common.BigToHash(big.NewInt(tokenID)),
		},
		BlockNumber: blockNum,
		TxHash:      txHash,
	}
}

func TestCheckMintLogAgainstFixtures(t *testing.T) {
	repo := integrationRepo(t)
	ctx := context.Background()

	t.Run("existing token stays quiet", func(t *testing.T) {
		var out bytes.Buffer
		checker := &mintChecker{repo: repo, out: &out, sleep: func(time.Duration) {}}
		if err := checker.checkMintLog(ctx, mintLog(1, 10, common.Hash{})); err != nil {
			t.Fatalf("checkMintLog: %v", err)
		}
		if out.Len() != 0 {
			t.Errorf("output = %q, want none for a token present in the DB", out.String())
		}
	})

	t.Run("missing token is reported", func(t *testing.T) {
		var out bytes.Buffer
		checker := &mintChecker{repo: repo, out: &out, sleep: func(time.Duration) {}}
		if err := checker.checkMintLog(ctx, mintLog(1, 999, common.Hash{})); err != nil {
			t.Fatalf("checkMintLog: %v", err)
		}
		if !strings.Contains(out.String(), "Token 999 DOES NOT exist in the DB") {
			t.Errorf("output = %q", out.String())
		}
	})

	t.Run("cancelled context aborts the retry loop", func(t *testing.T) {
		cancelled, cancel := context.WithCancel(ctx)
		cancel()
		var out bytes.Buffer
		checker := &mintChecker{repo: repo, out: &out, sleep: func(time.Duration) {}}
		if err := checker.checkMintLog(cancelled, mintLog(1, 10, common.Hash{})); err == nil {
			t.Error("cancelled context did not abort")
		}
	})
}

func TestCheckTransferLogAgainstFixtures(t *testing.T) {
	repo := integrationRepo(t)
	ctx := context.Background()

	// The fixture rw_transfer rows exist for token 10; find one via the full
	// history and reuse its identity for the positive case.
	transfers, err := repo.TokenTransfersByTxHash(ctx, "0xf000000000000000000000000000000000000000000000000000000000001036")
	if err != nil {
		t.Fatalf("fixture transfer lookup: %v", err)
	}
	if len(transfers) == 0 {
		t.Skip("fixture tx hash not found; transfer fixtures changed")
	}

	item := transfers[0]
	matching := &types.Log{
		Topics: []common.Hash{
			transferEventTopic,
			common.HexToHash(item.From),
			common.HexToHash(item.To),
			common.BigToHash(big.NewInt(item.TokenId)),
		},
		BlockNumber: 1,
		TxHash:      common.HexToHash("0xf000000000000000000000000000000000000000000000000000000000001036"),
	}
	var out bytes.Buffer
	if err := checkTransferLog(ctx, repo, &out, matching); err != nil {
		t.Fatalf("checkTransferLog: %v", err)
	}
	if out.Len() != 0 {
		t.Errorf("output = %q, want none for a recorded transfer", out.String())
	}

	// A mismatching token id on the same tx is reported missing.
	missing := &types.Log{
		Topics: []common.Hash{
			transferEventTopic,
			common.HexToHash(item.From),
			common.HexToHash(item.To),
			common.BigToHash(big.NewInt(item.TokenId + 500)),
		},
		BlockNumber: 1,
		TxHash:      common.HexToHash("0xf000000000000000000000000000000000000000000000000000000000001036"),
	}
	out.Reset()
	if err := checkTransferLog(ctx, repo, &out, missing); err != nil {
		t.Fatalf("checkTransferLog: %v", err)
	}
	if !strings.Contains(out.String(), "transfer missing") {
		t.Errorf("output = %q, want a missing-transfer report", out.String())
	}
}

func TestScanMintsCommandEndToEnd(t *testing.T) {
	setPGEnv(t)

	chain := testchain.New(t)
	// The scan starts at block 2,000,000; put one recorded and one missing
	// mint above it.
	rwalkAddr := common.HexToAddress(fixtureRWalkAddr)
	tx1 := chain.AddTx(2_000_001, rwalkAddr, nil)
	chain.AttachLogs(tx1.Hash(), []*types.Log{mintLog(2_000_001, 10, tx1.Hash())})
	tx2 := chain.AddTx(2_000_002, rwalkAddr, nil)
	chain.AttachLogs(tx2.Hash(), []*types.Log{mintLog(2_000_002, 999, tx2.Hash())})
	t.Setenv("RPC_URL", chain.URL())

	out, err := executeCmd(t, newScanMintsCmd(), fixtureRWalkAddr)
	if err != nil {
		t.Fatalf("scan-mints: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "From 2000000 , to") {
		t.Errorf("output missing the range progress line:\n%s", out)
	}
	if !strings.Contains(out, "Token 999 DOES NOT exist in the DB") {
		t.Errorf("output missing the missing-token report:\n%s", out)
	}
	if strings.Contains(out, "Token 10 DOES NOT exist") {
		t.Errorf("recorded token 10 reported missing:\n%s", out)
	}
}

func TestNotifyBotCommandWiringEndToEnd(t *testing.T) {
	setPGEnv(t)
	repo := integrationRepo(t)
	ctx := context.Background()

	// Park the watermark past all fixture events so the engine idles without
	// touching Twitter or the image server; restore it afterwards.
	status, err := repo.MessagingStatus(ctx)
	if err != nil {
		t.Fatalf("MessagingStatus: %v", err)
	}
	parked := status
	parked.EvtLogId = 1 << 40
	if err := repo.UpdateMessagingStatus(ctx, &parked); err != nil {
		t.Fatalf("parking watermark: %v", err)
	}
	t.Cleanup(func() {
		if err := repo.UpdateMessagingStatus(context.Background(), &status); err != nil {
			t.Errorf("restoring watermark: %v", err)
		}
	})

	chain := testchain.New(t)
	chain.EnsureBlock(1)
	t.Setenv("RPC_URL", chain.URL())

	home := t.TempDir()
	t.Setenv("HOME", home)
	if err := os.MkdirAll(filepath.Join(home, "configs"), 0o750); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(home, "configs", "tw.json"),
		[]byte(`{"ApiKey":"a","ApiSecret":"b","TokenKey":"c","TokenSecret":"d"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TWITTER_KEYS_FILE", "tw.json")

	runCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	cmd := newNotifyBotCmd()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs(nil)
	if err := cmd.ExecuteContext(runCtx); err != nil {
		t.Fatalf("notify-bot: %v\noutput: %s", err, out.String())
	}
}

func TestNotifyBotCommandMissingKeys(t *testing.T) {
	setPGEnv(t)
	t.Setenv("HOME", t.TempDir())
	t.Setenv("TWITTER_KEYS_FILE", "missing.json")

	_, err := executeCmd(t, newNotifyBotCmd())
	if err == nil || !strings.Contains(err.Error(), "twitter account keys") {
		t.Errorf("notify-bot without keys = %v", err)
	}
}

func TestVerifyOwnerCommand(t *testing.T) {
	setPGEnv(t)

	t.Run("count mismatch reported without token loop", func(t *testing.T) {
		chain := testchain.New(t)
		chain.EnsureBlock(1)
		stub := testchain.MustContractStub(rwcontracts.RWalkMetaData.ABI).
			Return("nextTokenId", big.NewInt(0))
		chain.RegisterCall(common.HexToAddress(fixtureRWalkAddr), stub.Handler())
		t.Setenv("RPC_URL", chain.URL())

		out, err := executeCmd(t, newVerifyOwnerCmd())
		if err != nil {
			t.Fatalf("verify-owner: %v\noutput: %s", err, out)
		}
		if !strings.Contains(out, "num tokens doesn't match") {
			t.Errorf("output missing the count mismatch:\n%s", out)
		}
	})

	t.Run("token loop surfaces database gaps", func(t *testing.T) {
		chain := testchain.New(t)
		chain.EnsureBlock(1)
		// Four tokens on chain matches the fixture count (so the "set
		// correctly" branch runs), but chain token ids 0..3 are not the
		// fixture's tokens 10..13: the loop must fail on TokenInfo.
		stub := testchain.MustContractStub(rwcontracts.RWalkMetaData.ABI).
			Return("nextTokenId", big.NewInt(4)).
			Return("ownerOf", common.HexToAddress("0x2100000000000000000000000000000000000021"))
		chain.RegisterCall(common.HexToAddress(fixtureRWalkAddr), stub.Handler())
		t.Setenv("RPC_URL", chain.URL())

		out, err := executeCmd(t, newVerifyOwnerCmd())
		if err == nil || !strings.Contains(err.Error(), "token info") {
			t.Errorf("verify-owner over unindexed token = %v", err)
		}
		if !strings.Contains(out, "set correctly") {
			t.Errorf("output missing the matching-count line:\n%s", out)
		}
	})
}

func TestVerifyErc20TransfersCommand(t *testing.T) {
	setPGEnv(t)
	repo := integrationRepo(t)
	ctx := context.Background()

	// A recorded fixture transfer plus a fabricated one on the same tx.
	transfers, err := repo.TokenTransfersByTxHash(ctx, "0xf000000000000000000000000000000000000000000000000000000000001036")
	if err != nil || len(transfers) == 0 {
		t.Skipf("fixture transfer lookup failed: %v", err)
	}
	item := transfers[0]

	chain := testchain.New(t)
	rwalkAddr := common.HexToAddress(fixtureRWalkAddr)
	tx := chain.AddTx(3, rwalkAddr, nil)
	transferLog := func(tokenID int64) *types.Log {
		return &types.Log{
			Address: rwalkAddr,
			Topics: []common.Hash{
				transferEventTopic,
				common.HexToHash(item.From),
				common.HexToHash(item.To),
				common.BigToHash(big.NewInt(tokenID)),
			},
			BlockNumber: 3,
			// The DB check keys on the log's tx hash: point both logs at the
			// fixture transaction.
			TxHash: common.HexToHash("0xf000000000000000000000000000000000000000000000000000000000001036"),
		}
	}
	chain.AttachLogs(tx.Hash(), []*types.Log{transferLog(item.TokenId), transferLog(item.TokenId + 500)})
	t.Setenv("RPC_URL", chain.URL())

	out, err := executeCmd(t, newVerifyTransfersCmd())
	if err != nil {
		t.Fatalf("verify-erc20-transfers: %v\noutput: %s", err, out)
	}
	if !strings.Contains(out, "transfer missing") {
		t.Errorf("output missing the fabricated-transfer report:\n%s", out)
	}
}

func TestMintCheckerRetriesTransientDBErrors(t *testing.T) {
	integrationRepo(t)

	// A repo over a closed pool: every read fails; the injected sleep
	// cancels the context so the retry loop ends deterministically.
	pool, err := pgxpool.New(context.Background(), sharedConnString)
	if err != nil {
		t.Fatalf("opening spare pool: %v", err)
	}
	closedRepo := rwstore.NewRepo(store.NewFromPool(pool))
	pool.Close()

	ctx, cancel := context.WithCancel(context.Background())
	var out bytes.Buffer
	checker := &mintChecker{
		repo:  closedRepo,
		out:   &out,
		sleep: func(time.Duration) { cancel() },
	}
	err = checker.checkMintLog(ctx, mintLog(1, 10, common.Hash{}))
	if err == nil {
		t.Fatal("checkMintLog over a dead pool with cancelled retry returned nil")
	}
	if !strings.Contains(out.String(), "Error accessing database") {
		t.Errorf("output = %q, want the transient-error report", out.String())
	}
}

func TestTopRatedCommandRecomputesRanks(t *testing.T) {
	setPGEnv(t)
	repo := integrationRepo(t)
	ctx := context.Background()

	out, err := executeCmd(t, newTopRatedCmd())
	if err != nil {
		t.Fatalf("top-rated: %v\noutput: %s", err, out)
	}

	stats, err := repo.RankingDataForAllUsers(ctx)
	if err != nil {
		t.Fatalf("RankingDataForAllUsers: %v", err)
	}
	if len(stats) == 0 {
		t.Fatal("fixture dataset has no user trade stats")
	}
	profitMakers, err := repo.TopProfitMakers(ctx)
	if err != nil {
		t.Fatalf("TopProfitMakers: %v", err)
	}
	if len(profitMakers) != len(stats) {
		t.Errorf("profit leaderboard has %d rows, want one per user with stats (%d)", len(profitMakers), len(stats))
	}
	// The best-ranked entry carries the top percentile value 1.0.
	if len(profitMakers) > 0 && profitMakers[0].Percentage != 1.0 {
		t.Errorf("top profit rank = %v, want 1.0", profitMakers[0].Percentage)
	}
}
