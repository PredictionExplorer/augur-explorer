//go:build integration

// Test harness for the CosmicGame ETL fixture suite (§4.3 of the
// modernization roadmap): one testcontainers Postgres plus one fake Ethereum
// node per test process, package globals initialized exactly like main(), and
// helpers that push synthetic logs through the real production pipeline
// (EnsureBlockExists -> EnsureTransactionExists -> InsertEventLog ->
// process_single_event).
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

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// Fixture contract addresses: identical to the API parity suite conventions
// (digits only, so checksummed form equals the lowercase form).
const (
	fxGameAddr        = "0x2000000000000000000000000000000000000002"
	fxSignatureAddr   = "0x3000000000000000000000000000000000000003"
	fxTokenAddr       = "0x4000000000000000000000000000000000000004"
	fxDaoAddr         = "0x5000000000000000000000000000000000000005"
	fxCharityAddr     = "0x6000000000000000000000000000000000000006"
	fxPrizesAddr      = "0x7000000000000000000000000000000000000007"
	fxRandomWalkAddr  = "0x8000000000000000000000000000000000000008"
	fxStakingCSTAddr  = "0x9000000000000000000000000000000000000009"
	fxStakingRWKAddr  = "0x1000000000000000000000000000000000000010"
	fxMarketingAddr   = "0x1100000000000000000000000000000000000011"
	fxMarketplaceAddr = "0x1200000000000000000000000000000000000012"
	fxImplAddr        = "0x1300000000000000000000000000000000000013"

	fxAlice      = "0x2100000000000000000000000000000000000021"
	fxBob        = "0x2200000000000000000000000000000000000022"
	fxCarol      = "0x2300000000000000000000000000000000000023"
	fxDave       = "0x2400000000000000000000000000000000000024"
	fxEmma       = "0x2500000000000000000000000000000000000025"
	fxMockERC20  = "0x2600000000000000000000000000000000000026"
	fxDonatedNFT = "0x2700000000000000000000000000000000000027"
	fxCharityRcv = "0x2800000000000000000000000000000000000028"
)

// seedCstRewardForBidding stands in for the ETL's startup chain sync
// (syncContractParamsFromChain), which populates cg_glob_stats from contract
// state; V1 BidPlaced inserts refuse to run without it.
const seedCstRewardForBidding = "100000000000000000000"

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
		fmt.Fprintf(os.Stderr, "cg-etl test: starting test database: %v\n", err)
		return 1
	}
	defer stopDB()
	testDB = db

	chain, stopChain := testchain.Start()
	defer stopChain()
	testChain = chain

	if err := initPackageGlobals(ctx, db); err != nil {
		fmt.Fprintf(os.Stderr, "cg-etl test: initializing globals: %v\n", err)
		return 1
	}

	return m.Run()
}

// initPackageGlobals mirrors the initialization order of main(): loggers,
// RPC clients, storage, ABIs. Contract addresses are (re)established by
// resetDB, which is the per-fixture equivalent of the address bootstrap.
func initPackageGlobals(ctx context.Context, db *testdb.DB) error {
	Info = log.New(io.Discard, "", 0)
	Error = log.New(os.Stderr, "cg-etl ERROR: ", 0)

	rpcclient, err := rpc.DialContext(ctx, testChain.URL())
	if err != nil {
		return fmt.Errorf("dialing fake chain: %w", err)
	}
	eclient = ethclient.NewClient(rpcclient)

	// One Store over the container's pool backs every query, exactly like
	// main().
	dbStore = store.NewFromPool(db.Pool)
	cgRepo = cgstore.NewRepo(dbStore)

	// The fixtures push logs through the same engine pipeline main() runs.
	testIndexer, err = indexer.New(indexer.Config{
		Store:  dbStore,
		Client: eclient,
		Logger: slog.New(slog.DiscardHandler),
	})
	if err != nil {
		return fmt.Errorf("building indexer engine: %w", err)
	}

	cosmic_game_abi = get_abi(cgc.CosmicSignatureGameABI)
	cosmic_game_v2_abi = get_abi(cgc.CosmicSignatureGameV2ABI)
	cosmic_signature_abi = get_abi(cgc.CosmicSignatureNftABI)
	charity_wallet_abi = get_abi(cgc.CharityWalletABI)
	prizes_wallet_abi = get_abi(cgc.PrizesWalletABI)
	staking_wallet_cst_abi = get_abi(cgc.IStakingWalletCosmicSignatureNftABI)
	staking_wallet_rwalk_abi = get_abi(cgc.IStakingWalletRandomWalkNftABI)
	marketing_wallet_abi = get_abi(cgc.MarketingWalletABI)
	erc20_abi = get_abi(cgc.ERC20ABI)
	erc721_abi = get_abi(cgc.ERC721ABI)
	erc1967_abi = get_abi(cgc.IERC1967ABI)

	registerCallHandlers()
	return nil
}

// registerCallHandlers wires the two contract reads the event handlers
// perform: EthDonationWithInfoRecords on the game and tokenURI on donated
// NFT collections. Responses are deterministic functions of the arguments.
func registerCallHandlers() {
	game := ethcommon.HexToAddress(fxGameAddr)
	donationRecords := cosmic_game_abi.Methods["ethDonationWithInfoRecords"]
	testChain.RegisterCall(game, func(input []byte) ([]byte, error) {
		if len(input) < 4 || string(input[:4]) != string(donationRecords.ID) {
			return nil, fmt.Errorf("unexpected call to game contract: %x", input)
		}
		args, err := donationRecords.Inputs.Unpack(input[4:])
		if err != nil {
			return nil, fmt.Errorf("unpacking donation record args: %w", err)
		}
		recordID := args[0].(*big.Int)
		return donationRecords.Outputs.Pack(
			big.NewInt(0),
			ethcommon.Address{},
			big.NewInt(0),
			fmt.Sprintf(`{"fixture":"donation record %v"}`, recordID),
		)
	})

	tokenURI := cosmic_signature_abi.Methods["tokenURI"]
	testChain.RegisterCall(ethcommon.HexToAddress(fxDonatedNFT), func(input []byte) ([]byte, error) {
		if len(input) < 4 || string(input[:4]) != string(tokenURI.ID) {
			return nil, fmt.Errorf("unexpected call to donated NFT contract: %x", input)
		}
		args, err := tokenURI.Inputs.Unpack(input[4:])
		if err != nil {
			return nil, fmt.Errorf("unpacking tokenURI args: %w", err)
		}
		return tokenURI.Outputs.Pack(fmt.Sprintf("ipfs://fixture/%v", args[0].(*big.Int)))
	})
}

// requireHarness skips the test when Docker is unavailable.
func requireHarness(t *testing.T) {
	t.Helper()
	if errSetupSkip != nil {
		t.Skipf("skipping: %v", errSetupSkip)
	}
	if testDB == nil {
		t.Fatal("cg-etl test harness not initialized (TestMain did not run?)")
	}
}

// resetSeedSQL restores the rows the goose migrations and operators provide
// on a fresh production database: the contract registry, the singleton stats
// rows and the last_block watermark.
const resetSeedSQL = `
INSERT INTO cg_contracts(
  cosmic_game_addr, cosmic_signature_addr, cosmic_token_addr, cosmic_dao_addr,
  charity_wallet_addr, prizes_wallet_addr, random_walk_addr,
  staking_wallet_cst_addr, staking_wallet_rwalk_addr, marketing_wallet_addr,
  implementation_addr
) VALUES (
  '` + fxGameAddr + `', '` + fxSignatureAddr + `', '` + fxTokenAddr + `', '` + fxDaoAddr + `',
  '` + fxCharityAddr + `', '` + fxPrizesAddr + `', '` + fxRandomWalkAddr + `',
  '` + fxStakingCSTAddr + `', '` + fxStakingRWKAddr + `', '` + fxMarketingAddr + `',
  '` + fxImplAddr + `'
);
INSERT INTO rw_contracts(marketplace_addr, randomwalk_addr)
  VALUES ('` + fxMarketplaceAddr + `', '` + fxRandomWalkAddr + `');
INSERT INTO cg_glob_stats DEFAULT VALUES;
INSERT INTO cg_stake_stats_cst DEFAULT VALUES;
INSERT INTO cg_stake_stats_rwalk DEFAULT VALUES;
INSERT INTO last_block(block_num) VALUES (0);
UPDATE cg_glob_stats SET cst_reward_for_bidding = '` + seedCstRewardForBidding + `';
`

// resetDB truncates every user table, restarts all sequences, re-applies the
// seed and re-registers the contract addresses (mirroring main()'s address
// bootstrap), leaving the database exactly as fixtures expect it.
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

	// Same bootstrap as main(): register all contract addresses, then resolve
	// the two address ids the decode helpers need.
	cg_contracts, err = cgRepo.ContractAddrs(ctx)
	if err != nil {
		t.Fatalf("reading contract addresses: %v", err)
	}
	for _, contractAddr := range []string{
		cg_contracts.CosmicGameAddr,
		cg_contracts.CosmicSignatureAddr,
		cg_contracts.CosmicTokenAddr,
		cg_contracts.CosmicDaoAddr,
		cg_contracts.CharityWalletAddr,
		cg_contracts.PrizesWalletAddr,
		cg_contracts.RandomWalkAddr,
		cg_contracts.StakingWalletCSTAddr,
		cg_contracts.StakingWalletRWalkAddr,
		cg_contracts.MarketingWalletAddr,
		cg_contracts.ImplementationAddr,
	} {
		if _, err := dbStore.LookupOrCreateAddress(ctx, contractAddr, 0, 0); err != nil {
			t.Fatalf("registering contract address %v: %v", contractAddr, err)
		}
	}
	cosmic_tok_aid, err = dbStore.LookupAddressID(ctx, cg_contracts.CosmicTokenAddr)
	if err != nil {
		t.Fatalf("looking up CosmicToken aid: %v", err)
	}
	cosmic_game_addr = ethcommon.HexToAddress(cg_contracts.CosmicGameAddr)
	cosmic_signature_addr = ethcommon.HexToAddress(cg_contracts.CosmicSignatureAddr)
	cosmic_token_addr = ethcommon.HexToAddress(cg_contracts.CosmicTokenAddr)
	cosmic_dao_addr = ethcommon.HexToAddress(cg_contracts.CosmicDaoAddr)
	charity_wallet_addr = ethcommon.HexToAddress(cg_contracts.CharityWalletAddr)
	prizes_wallet_addr = ethcommon.HexToAddress(cg_contracts.PrizesWalletAddr)
	staking_wallet_cst_addr = ethcommon.HexToAddress(cg_contracts.StakingWalletCSTAddr)
	staking_wallet_rwalk_addr = ethcommon.HexToAddress(cg_contracts.StakingWalletRWalkAddr)
	marketing_wallet_addr = ethcommon.HexToAddress(cg_contracts.MarketingWalletAddr)
	implementation_addr = ethcommon.HexToAddress(cg_contracts.ImplementationAddr)
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
// Log indexes continue across transactions of the same block via startLogIndex.
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

// buildLog and buildRawLog alias the shared fixture-log builders.
func buildLog(t *testing.T, contractABI *abi.ABI, eventName string, address ethcommon.Address, indexed []any, nonIndexed []any) *types.Log {
	t.Helper()
	return testutil.BuildEventLog(t, contractABI, eventName, address, indexed, nonIndexed)
}

func buildRawLog(t *testing.T, topic0Hex string, address ethcommon.Address, indexedTopics []ethcommon.Hash, dataWords ...*big.Int) *types.Log {
	t.Helper()
	return testutil.BuildRawLog(t, topic0Hex, address, indexedTopics, dataWords...)
}

// Common big.Int shorthands used across fixtures.
func bigInt(v int64) *big.Int { return big.NewInt(v) }

// eth converts whole ether to wei.
func eth(n int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(n), big.NewInt(1_000_000_000_000_000_000))
}

func addr(hex string) ethcommon.Address { return ethcommon.HexToAddress(hex) }
