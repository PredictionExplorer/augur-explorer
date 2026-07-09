//go:build integration

// Test harness for the RandomWalk ETL fixture suite (§4.3): one testcontainers
// Postgres plus one fake Ethereum node per test process, package globals
// initialized exactly like main(), and helpers that push synthetic logs
// through the real production pipeline.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	rwc "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// Fixture contract and user addresses (digits only: checksummed form equals
// the lowercase form), matching the conventions of the other suites.
const (
	fxMarketplaceAddr = "0x1200000000000000000000000000000000000012"
	fxRandomWalkAddr  = "0x8000000000000000000000000000000000000008"

	fxAlice = "0x2100000000000000000000000000000000000021"
	fxBob   = "0x2200000000000000000000000000000000000022"
	fxCarol = "0x2300000000000000000000000000000000000023"
	fxDave  = "0x2400000000000000000000000000000000000024"
	fxZero  = "0x0000000000000000000000000000000000000000"
)

var (
	testDB       *testdb.DB
	testChain    *testchain.Chain
	testIndexer  *indexer.Engine // the production pipeline the fixtures run through
	errSetupSkip error           // non-nil: integration environment unavailable, skip
)

// TestMain owns the database container, the fake chain and the package
// globals (storage, ABIs, addresses, loggers) that main() would normally
// initialize. Fixtures reset the database, not the process, between cases.
func TestMain(m *testing.M) {
	os.Exit(runMain(m))
}

func runMain(m *testing.M) int {
	flag.Parse()
	time.Local = time.UTC

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	db, stopDB, err := testdb.Start(ctx)
	if err != nil {
		if errors.Is(err, testdb.ErrContainerUnavailable) {
			errSetupSkip = err
			return m.Run() // every test skips with the reason
		}
		fmt.Fprintf(os.Stderr, "rw-etl test: starting test database: %v\n", err)
		return 1
	}
	defer stopDB()
	testDB = db

	chain, stopChain := testchain.Start()
	defer stopChain()
	testChain = chain

	if err := initPackageGlobals(ctx, db); err != nil {
		fmt.Fprintf(os.Stderr, "rw-etl test: initializing globals: %v\n", err)
		return 1
	}

	return m.Run()
}

// initPackageGlobals mirrors the initialization order of main(): loggers,
// RPC clients, storage, ABIs. The contract addresses are (re)established by
// resetDB.
func initPackageGlobals(ctx context.Context, db *testdb.DB) error {
	Info = log.New(io.Discard, "", 0)
	Error = log.New(os.Stderr, "rw-etl ERROR: ", 0)

	rpcclient, err := rpc.DialContext(ctx, testChain.URL())
	if err != nil {
		return fmt.Errorf("dialing fake chain: %w", err)
	}
	eclient = ethclient.NewClient(rpcclient)

	dbStore = store.NewFromPool(db.Pool)
	rwRepo = rwdb.NewRepo(dbStore)

	// The fixtures push logs through the same engine pipeline main() runs.
	testIndexer, err = indexer.New(indexer.Config{
		Store:  dbStore,
		Client: eclient,
		Logger: slog.New(slog.DiscardHandler),
	})
	if err != nil {
		return fmt.Errorf("building indexer engine: %w", err)
	}

	marketABI, err := abi.JSON(strings.NewReader(rwc.RWMarketABI))
	if err != nil {
		return fmt.Errorf("parsing marketplace ABI: %w", err)
	}
	marketplace_abi = &marketABI
	rwalkABI, err := abi.JSON(strings.NewReader(rwc.RWalkABI))
	if err != nil {
		return fmt.Errorf("parsing randomwalk ABI: %w", err)
	}
	randomwalk_abi = &rwalkABI
	return nil
}

// requireHarness skips the test when Docker is unavailable.
func requireHarness(t *testing.T) {
	t.Helper()
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	if testDB == nil {
		t.Fatal("rw-etl test harness not initialized (TestMain did not run?)")
	}
}

// resetSeedSQL restores what a fresh production database provides: the
// contract registry row and the last_block watermark.
const resetSeedSQL = `
INSERT INTO rw_contracts(marketplace_addr, randomwalk_addr)
  VALUES ('` + fxMarketplaceAddr + `', '` + fxRandomWalkAddr + `');
INSERT INTO last_block(block_num) VALUES (0);
`

// resetDB truncates every user table, restarts all sequences, re-applies the
// seed and re-registers the contract addresses (mirroring main()'s bootstrap).
func resetDB(t *testing.T) {
	t.Helper()
	requireHarness(t)
	ctx := context.Background()

	var tables string
	err := testDB.SQL.QueryRowContext(ctx, `
		SELECT string_agg(quote_ident(table_name), ', ')
		FROM information_schema.tables
		WHERE table_schema = 'public' AND table_type = 'BASE TABLE'
		  AND table_name <> 'goose_db_version'`).Scan(&tables)
	if err != nil {
		t.Fatalf("listing tables for reset: %v", err)
	}
	if _, err := testDB.SQL.ExecContext(ctx, "TRUNCATE "+tables+" RESTART IDENTITY CASCADE"); err != nil {
		t.Fatalf("truncating tables: %v", err)
	}
	dbStore.ResetAddressCache()

	if _, err := testDB.SQL.ExecContext(ctx, resetSeedSQL); err != nil {
		t.Fatalf("re-seeding database: %v", err)
	}

	for _, contractAddr := range []string{fxMarketplaceAddr, fxRandomWalkAddr} {
		if _, err := dbStore.LookupOrCreateAddress(ctx, contractAddr, 0, 0); err != nil {
			t.Fatalf("registering contract address %v: %v", contractAddr, err)
		}
	}
	var err2 error
	rw_contracts, err2 = rwRepo.ContractAddrs(ctx)
	if err2 != nil {
		t.Fatalf("resolving contract addresses: %v", err2)
	}
	rwalk_addr = ethcommon.HexToAddress(rw_contracts.RandomWalk)
	market_addr = ethcommon.HexToAddress(rw_contracts.MarketPlace)
}

// snapshot captures the canonical database state.
func snapshot(t *testing.T) testutil.Snapshot {
	t.Helper()
	snap, err := testutil.TakeSnapshot(context.Background(), testDB.SQL)
	if err != nil {
		t.Fatalf("taking snapshot: %v", err)
	}
	return snap
}

// requireNoDiff asserts two snapshots are identical.
func requireNoDiff(t *testing.T, before, after testutil.Snapshot, context string) {
	t.Helper()
	diff, err := testutil.DiffSnapshots(before, after)
	if err != nil {
		t.Fatalf("%s: diffing snapshots: %v", context, err)
	}
	if strings.TrimSpace(string(diff)) != "{}" {
		t.Errorf("%s: state changed unexpectedly:\n%s", context, diff)
	}
}

// ingestTx records one transaction with its logs on the fake chain and runs
// every log through the production pipeline, returning the evt_log ids.
func ingestTx(t *testing.T, blockNum int64, to ethcommon.Address, startLogIndex uint, logs []*types.Log) []int64 {
	t.Helper()

	tx := testChain.AddTx(blockNum, to, nil)
	for i, l := range logs {
		l.BlockNumber = uint64(blockNum)
		l.BlockHash = testChain.BlockHash(blockNum)
		l.TxHash = tx.Hash()
		l.TxIndex = 0
		l.Index = startLogIndex + uint(i)
	}
	testChain.AttachLogs(tx.Hash(), logs)

	evtIDs := make([]int64, 0, len(logs))
	for _, l := range logs {
		ctx := context.Background()
		if _, err := testIndexer.EnsureBlockExists(ctx, blockNum, l.BlockHash.Hex()); err != nil {
			t.Fatalf("EnsureBlockExists(%d): %v", blockNum, err)
		}
		txID, _, err := testIndexer.EnsureTransactionExists(ctx, l.TxHash, blockNum)
		if err != nil {
			t.Fatalf("EnsureTransactionExists(%s): %v", l.TxHash, err)
		}
		evtID, err := testIndexer.InsertEventLog(ctx, *l, txID)
		if err != nil {
			t.Fatalf("InsertEventLog: %v", err)
		}
		if err := process_single_event(context.Background(), evtID); err != nil {
			t.Fatalf("process_single_event(%d): %v", evtID, err)
		}
		evtIDs = append(evtIDs, evtID)
	}
	return evtIDs
}

func buildLog(t *testing.T, contractABI *abi.ABI, eventName string, address ethcommon.Address, indexed []any, nonIndexed []any) *types.Log {
	t.Helper()
	return testutil.BuildEventLog(t, contractABI, eventName, address, indexed, nonIndexed)
}

func bigInt(v int64) *big.Int { return big.NewInt(v) }

// eth converts whole ether to wei.
func eth(n int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(n), big.NewInt(1_000_000_000_000_000_000))
}

func addr(hex string) ethcommon.Address { return ethcommon.HexToAddress(hex) }
