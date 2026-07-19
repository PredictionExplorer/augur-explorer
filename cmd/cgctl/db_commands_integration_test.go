//go:build integration

package main

// Integration tests against a real seeded Postgres: the database-backed
// subcommands (total-tokens, token-seed, backfill-dao-evtlog) run through
// the production wiring — cobra args, PGSQL_* environment, the shared store
// pool and, for the backfill, a fake chain serving cosmic_dao logs.

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// fixtureDaoAddr is cosmic_dao_addr in the shared fixture dataset.
const fixtureDaoAddr = "0x5000000000000000000000000000000000000005"

var (
	sharedStore  *store.Store
	sharedRepo   *cgstore.Repo
	sharedPGEnv  map[string]string
	errSetupSkip error
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
		fmt.Fprintf(os.Stderr, "cgctl: starting test database: %v\n", err)
		return 1
	}
	defer stop()

	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		fmt.Fprintf(os.Stderr, "cgctl: seeding fixture data: %v\n", err)
		return 1
	}
	sharedStore = store.NewFromPool(db.Pool)
	sharedRepo = cgstore.NewRepo(sharedStore)

	u, err := url.Parse(db.ConnString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cgctl: parsing container conn string: %v\n", err)
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

// setPGEnv points the production PGSQL_* variables at the container,
// skipping when Docker is unavailable.
func setPGEnv(t *testing.T) {
	t.Helper()
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	if sharedPGEnv == nil {
		t.Fatal("cgctl harness not initialized (TestMain did not run?)")
	}
	for k, v := range sharedPGEnv {
		t.Setenv(k, v)
	}
}

func TestTotalTokensCommand(t *testing.T) {
	setPGEnv(t)

	want, err := sharedRepo.CosmicSignatureTokenCount(context.Background())
	if err != nil {
		t.Fatalf("reference count: %v", err)
	}
	if want == 0 {
		t.Fatal("fixture dataset has no CosmicSignature tokens; test is vacuous")
	}

	out, err := executeCmd(t, newTotalTokensCmd())
	if err != nil {
		t.Fatalf("total-tokens: %v\noutput: %s", err, out)
	}
	// The value is printed without a trailing newline for shell capture.
	if out != strconv.FormatInt(want, 10) {
		t.Errorf("total-tokens output = %q, want %d", out, want)
	}
}

func TestTokenSeedCommand(t *testing.T) {
	setPGEnv(t)

	t.Run("known token", func(t *testing.T) {
		out, err := executeCmd(t, newTokenSeedCmd(), "1")
		if err != nil {
			t.Fatalf("token-seed: %v\noutput: %s", err, out)
		}
		if out != "seed0000000000000000000000000000000000000000000000000000000001" {
			t.Errorf("token-seed output = %q", out)
		}
	})

	t.Run("unknown token prints empty", func(t *testing.T) {
		out, err := executeCmd(t, newTokenSeedCmd(), "999999")
		if err != nil {
			t.Fatalf("token-seed unknown: %v\noutput: %s", err, out)
		}
		if out != "" {
			t.Errorf("unknown token output = %q, want empty (imgcheck.sh contract)", out)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		_, err := executeCmd(t, newTokenSeedCmd(), "one")
		if err == nil || !strings.Contains(err.Error(), "token-id") {
			t.Errorf("invalid id = %v", err)
		}
	})
}

func TestDBCommandsConnectionFailure(t *testing.T) {
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	t.Setenv("PGSQL_HOST", "127.0.0.1:1")
	t.Setenv("PGSQL_USERNAME", "nobody")
	t.Setenv("PGSQL_DATABASE", "nothing")
	t.Setenv("PGSQL_PASSWORD", "wrong")

	_, err := executeCmd(t, newTotalTokensCmd())
	if err == nil || !strings.Contains(err.Error(), "failed to connect to storage") {
		t.Errorf("connection failure = %v", err)
	}
}

func TestBackfillDaoEvtlogCommand(t *testing.T) {
	setPGEnv(t)

	// A fake chain serving two cosmic_dao logs in blocks 200 and 201.
	chain := testchain.New(t)
	daoAddr := common.HexToAddress(fixtureDaoAddr)
	topic := common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
	for i, blockNum := range []int64{200, 201} {
		tx := chain.AddTx(blockNum, daoAddr, []byte{byte(i)})
		chain.AttachLogs(tx.Hash(), []*types.Log{{
			Address:     daoAddr,
			Topics:      []common.Hash{topic},
			BlockNumber: uint64(blockNum), // #nosec G115 -- positive test block constant
			BlockHash:   chain.BlockHash(blockNum),
			TxHash:      tx.Hash(),
			Index:       0,
		}})
	}
	chain.EnsureBlock(210)
	t.Setenv("RPC_URL", chain.URL())

	countRows := func() int64 {
		n, err := sharedStore.CountEvtLogsForContract(context.Background(), fixtureDaoAddr)
		if err != nil {
			t.Fatalf("counting dao evt_log rows: %v", err)
		}
		return n
	}
	before := countRows()

	out, err := executeCmd(t, newBackfillDaoEvtlogCmd(), "--from-block=200", "--to-block=210")
	if err != nil {
		t.Fatalf("backfill-dao-evtlog: %v\noutput: %s", err, out)
	}
	if got := countRows(); got != before+2 {
		t.Errorf("dao evt_log rows = %d, want %d", got, before+2)
	}
	for _, want := range []string{"backfilling cosmic_dao evt_log", "inserted=2", "backfill done"} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}

	// Re-running is idempotent: everything is skipped.
	out, err = executeCmd(t, newBackfillDaoEvtlogCmd(), "--from-block=200", "--to-block=210")
	if err != nil {
		t.Fatalf("backfill re-run: %v\noutput: %s", err, out)
	}
	if got := countRows(); got != before+2 {
		t.Errorf("dao evt_log rows after re-run = %d, want %d", got, before+2)
	}
	if !strings.Contains(out, "skipped=2") {
		t.Errorf("re-run output missing skip report:\n%s", out)
	}
}

func TestBackfillDaoEvtlogValidatesRange(t *testing.T) {
	setPGEnv(t)
	chain := testchain.New(t)
	chain.EnsureBlock(10)
	t.Setenv("RPC_URL", chain.URL())

	_, err := executeCmd(t, newBackfillDaoEvtlogCmd(), "--from-block=100", "--to-block=50")
	if err == nil || !strings.Contains(err.Error(), "to-block 50 < from-block 100") {
		t.Errorf("inverted range = %v", err)
	}
}

// TestBackfillDaoEvtlogRejectsNegativeWatermarks pins the corrupt-watermark
// guards: without --to-block the scan end comes from the database, and a
// negative value there must abort instead of wrapping into an astronomical
// block range.
func TestBackfillDaoEvtlogRejectsNegativeWatermarks(t *testing.T) {
	setPGEnv(t)
	chain := testchain.New(t)
	chain.EnsureBlock(10)
	t.Setenv("RPC_URL", chain.URL())
	ctx := context.Background()

	setWatermarks := func(procStatus, lastBlock int64) {
		t.Helper()
		if _, err := sharedStore.Pool().Exec(ctx,
			"UPDATE cg_proc_status SET last_block_num = $1", procStatus); err != nil {
			t.Fatalf("setting cg_proc_status watermark: %v", err)
		}
		if _, err := sharedStore.Pool().Exec(ctx,
			"UPDATE last_block SET block_num = $1", lastBlock); err != nil {
			t.Fatalf("setting last_block watermark: %v", err)
		}
	}
	// ProcessingStatus lazily creates the singleton row on a fresh database.
	status, err := sharedRepo.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("bootstrapping cg_proc_status: %v", err)
	}
	origLast, err := sharedStore.LastBlockNum(ctx)
	if err != nil {
		t.Fatalf("reading last_block watermark: %v", err)
	}
	t.Cleanup(func() { setWatermarks(status.LastBlockNum, origLast) })

	setWatermarks(-7, origLast)
	_, err = executeCmd(t, newBackfillDaoEvtlogCmd())
	if err == nil || !strings.Contains(err.Error(), "processing status watermark is negative") {
		t.Errorf("negative processing status = %v", err)
	}

	setWatermarks(0, -9)
	_, err = executeCmd(t, newBackfillDaoEvtlogCmd())
	if err == nil || !strings.Contains(err.Error(), "last_block watermark is negative") {
		t.Errorf("negative last_block = %v", err)
	}
}

func TestBackfillDaoEvtlogRequiresRPC(t *testing.T) {
	setPGEnv(t)
	t.Setenv("RPC_URL", "")
	_, err := executeCmd(t, newBackfillDaoEvtlogCmd())
	if err == nil || !strings.Contains(err.Error(), "RPC_URL") {
		t.Errorf("missing RPC_URL = %v", err)
	}
}
