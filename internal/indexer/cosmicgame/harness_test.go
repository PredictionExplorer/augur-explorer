//go:build integration

// Test harness for the CosmicGame handler fixture suite (§4.3 of the
// modernization roadmap): one testcontainers Postgres plus one fake Ethereum
// node per test process, a Handlers set built exactly like main() builds it,
// and helpers that push synthetic logs through the real production pipeline
// (EnsureBlockExists -> EnsureTransactionExists -> InsertEventLog ->
// LogProcessor -> Registry).
package cosmicgame

import (
	"context"
	"errors"
	"flag"
	"fmt"
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
// (SyncContractParams), which populates cg_glob_stats from contract
// state; V1 BidPlaced inserts refuse to run without it.
const seedCstRewardForBidding = "100000000000000000000"

var (
	testDB       *testdb.DB
	testChain    *testchain.Chain
	testIndexer  *indexer.Engine // the production pipeline the fixtures run through
	errSetupSkip error           // non-nil: integration environment unavailable, skip

	dbStore *store.Store
	cgRepo  *cgstore.Repo
	eclient *ethclient.Client

	// Rebuilt by resetDB over the freshly re-registered contract addresses.
	testHandlers  *Handlers
	testProcess   indexer.ProcessFunc
	testContracts Contracts

	// ABI handles for the fixture-log builders (parsed once in TestMain).
	gameABI            *abi.ABI
	gameV2ABI          *abi.ABI
	signatureABI       *abi.ABI
	charityWalletABI   *abi.ABI
	prizesWalletABI    *abi.ABI
	stakingCSTABI      *abi.ABI
	stakingRWalkABI    *abi.ABI
	marketingWalletABI *abi.ABI
	erc20ABI           *abi.ABI
	erc721ABI          *abi.ABI
	erc1967ABI         *abi.ABI
)

// TestMain owns the database container, the fake chain and the harness state
// (storage, ABIs, addresses) that main() would normally initialize. Fixtures
// reset the database, not the process, between cases.
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
		fmt.Fprintf(os.Stderr, "cosmicgame handler test: starting test database: %v\n", err)
		return 1
	}
	defer stopDB()
	testDB = db

	chain, stopChain := testchain.Start()
	defer stopChain()
	testChain = chain

	if err := initHarness(ctx, db); err != nil {
		fmt.Fprintf(os.Stderr, "cosmicgame handler test: initializing harness: %v\n", err)
		return 1
	}

	return m.Run()
}

// initHarness mirrors the initialization order of main(): RPC client,
// storage, ABIs. The contract addresses and the Handlers set are
// (re)established by resetDB, which is the per-fixture equivalent of the
// address bootstrap.
func initHarness(ctx context.Context, db *testdb.DB) error {
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

	for _, a := range []struct {
		dst *(*abi.ABI)
		raw string
	}{
		{&gameABI, cgc.CosmicSignatureGameABI},
		{&gameV2ABI, cgc.CosmicSignatureGameV2ABI},
		{&signatureABI, cgc.CosmicSignatureNftABI},
		{&charityWalletABI, cgc.CharityWalletABI},
		{&prizesWalletABI, cgc.PrizesWalletABI},
		{&stakingCSTABI, cgc.IStakingWalletCosmicSignatureNftABI},
		{&stakingRWalkABI, cgc.IStakingWalletRandomWalkNftABI},
		{&marketingWalletABI, cgc.MarketingWalletABI},
		{&erc20ABI, cgc.ERC20ABI},
		{&erc721ABI, cgc.ERC721ABI},
		{&erc1967ABI, cgc.IERC1967ABI},
	} {
		parsed, err := abi.JSON(strings.NewReader(a.raw))
		if err != nil {
			return fmt.Errorf("parsing fixture ABI: %w", err)
		}
		*a.dst = &parsed
	}

	registerCallHandlers()
	return nil
}

// registerCallHandlers wires the two contract reads the event handlers
// perform: EthDonationWithInfoRecords on the game and tokenURI on donated
// NFT collections. Responses are deterministic functions of the arguments.
func registerCallHandlers() {
	game := ethcommon.HexToAddress(fxGameAddr)
	donationRecords := gameABI.Methods["ethDonationWithInfoRecords"]
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

	tokenURI := signatureABI.Methods["tokenURI"]
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
		t.Fatal("cosmicgame handler test harness not initialized (TestMain did not run?)")
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
// seed, re-registers the contract addresses (the same BootstrapContracts
// call main() makes) and rebuilds the Handlers set over them, leaving the
// database exactly as fixtures expect it.
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

	// Same bootstrap as main(): register all contract addresses, resolve the
	// ids the handlers need, build the handler set.
	testContracts, _, err = BootstrapContracts(ctx, cgRepo, dbStore)
	if err != nil {
		t.Fatalf("bootstrapping contract addresses: %v", err)
	}
	testHandlers = newTestHandlers(t, cgRepo, dbStore, testContracts)
	testProcess = indexer.LogProcessor(dbStore, testHandlers.Registry())
}

// newTestHandlers builds a Handlers set with the harness's eth client and a
// discarding logger.
func newTestHandlers(t *testing.T, repo *cgstore.Repo, st *store.Store, contracts Contracts) *Handlers {
	t.Helper()
	h, err := New(Config{
		Repo:      repo,
		Store:     st,
		Caller:    eclient,
		Contracts: contracts,
		Logger:    slog.New(slog.DiscardHandler),
	})
	if err != nil {
		t.Fatalf("building handlers: %v", err)
	}
	return h
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

// harnessProgress is the production cg-etl watermark adapter with an
// optional post-write hook. Tests use the hook to cancel Engine.Run only
// after the target block's watermark write has joined its transaction.
type harnessProgress struct {
	repo  *cgstore.Repo
	onSet func(block int64)
}

func (p *harnessProgress) LastBlock(ctx context.Context) (int64, error) {
	status, err := p.repo.ProcessingStatus(ctx)
	if err != nil {
		return 0, err
	}
	return status.LastBlockNum, nil
}

func (p *harnessProgress) SetLastBlock(ctx context.Context, block int64) error {
	status, err := p.repo.ProcessingStatus(ctx)
	if err != nil {
		return err
	}
	status.LastBlockNum = block
	if err := p.repo.UpdateProcessingStatus(ctx, &status); err != nil {
		return err
	}
	if p.onSet != nil {
		p.onSet(block)
	}
	return nil
}

// primeHarnessProgress makes the next Engine.Run start at firstBlock without
// changing the layer-1 last_block watermark being tested.
func primeHarnessProgress(t *testing.T, firstBlock int64) *harnessProgress {
	t.Helper()
	progress := &harnessProgress{repo: cgRepo}
	status, err := cgRepo.ProcessingStatus(context.Background())
	if err != nil {
		t.Fatalf("reading processing status: %v", err)
	}
	status.LastBlockNum = firstBlock - 1
	if err := cgRepo.UpdateProcessingStatus(context.Background(), &status); err != nil {
		t.Fatalf("priming processing status: %v", err)
	}
	return progress
}

// newHarnessRunEngine builds the same Engine configuration as cmd/cg-etl,
// with a fixed batch width and one-attempt breaker for deterministic
// integration tests.
func newHarnessRunEngine(
	t *testing.T,
	progress indexer.Progress,
	process indexer.ProcessFunc,
	batchSize uint64,
) *indexer.Engine {
	t.Helper()
	if process == nil {
		process = testProcess
	}
	engine, err := indexer.New(indexer.Config{
		Store:     dbStore,
		Client:    eclient,
		Progress:  progress,
		Process:   process,
		Contracts: testContracts.All(),
		Logger:    slog.New(slog.DiscardHandler),
		TopicName: testHandlers.Registry().TopicName,
		Batch: indexer.BatchConfig{
			Initial: batchSize,
			Min:     batchSize,
			Max:     batchSize,
		},
		Retry: indexer.RetryConfig{
			MaxConsecutiveFailures: 1,
			MinDelay:               time.Nanosecond,
			MaxDelay:               time.Nanosecond,
		},
		CaughtUpDelay: time.Hour,
	})
	if err != nil {
		t.Fatalf("building run engine: %v", err)
	}
	return engine
}

// runHarnessRange runs the production polling loop from firstBlock through
// targetBlock and cancels immediately after the target watermark commits.
// A returned error is the engine's real breaker error.
func runHarnessRange(
	t *testing.T,
	firstBlock, targetBlock int64,
	batchSize uint64,
) error {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	progress := primeHarnessProgress(t, firstBlock)
	progress.onSet = func(block int64) {
		if block == targetBlock {
			cancel()
		}
	}
	err := newHarnessRunEngine(t, progress, nil, batchSize).Run(ctx)
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		t.Fatal("Engine.Run timed out")
	}
	return err
}

// stageTx records one transaction and its logs on the fake chain without
// touching PostgreSQL. Engine.Run can then fetch and ingest the block through
// the exact production processBatch path.
func stageTx(t *testing.T, blockNum int64, to ethcommon.Address, startLogIndex uint, logs []*types.Log) {
	t.Helper()
	tx := testChain.AddTx(blockNum, to, nil)
	for i, l := range logs {
		l.BlockNumber = uint64(blockNum) // #nosec G115 -- positive test block constant
		l.BlockHash = testChain.BlockHash(blockNum)
		l.TxHash = tx.Hash()
		l.TxIndex = 0
		l.Index = startLogIndex + uint(i)
	}
	testChain.AttachLogs(tx.Hash(), logs)
}

// ingestTx records one transaction with its logs on the fake chain and runs
// every log through the production pipeline, returning the evt_log ids.
// Log indexes continue across transactions of the same block via startLogIndex.
func ingestTx(t *testing.T, blockNum int64, to ethcommon.Address, startLogIndex uint, logs []*types.Log) []int64 {
	t.Helper()

	stageTx(t, blockNum, to, startLogIndex, logs)

	// One transaction per ingested chain transaction, mirroring the engine's
	// per-block InTx: the fixtures prove every handler behaves identically
	// inside the production transaction scope.
	evtIDs := make([]int64, 0, len(logs))
	err := dbStore.InTx(context.Background(), func(ctx context.Context) error {
		for _, l := range logs {
			if _, err := testIndexer.EnsureBlockExists(ctx, blockNum, l.BlockHash.Hex()); err != nil {
				return fmt.Errorf("EnsureBlockExists(%d): %w", blockNum, err)
			}
			txID, _, err := testIndexer.EnsureTransactionExists(ctx, l.TxHash, blockNum)
			if err != nil {
				return fmt.Errorf("EnsureTransactionExists(%s): %w", l.TxHash, err)
			}
			evtID, err := testIndexer.InsertEventLog(ctx, *l, txID)
			if err != nil {
				return fmt.Errorf("InsertEventLog: %w", err)
			}
			if err := testProcess(ctx, evtID); err != nil {
				return fmt.Errorf("processing event %d: %w", evtID, err)
			}
			evtIDs = append(evtIDs, evtID)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("ingesting logs: %v", err)
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
