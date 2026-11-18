//go:build integration

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5/pgxpool"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

// vcContracts is the deterministic address book used by the validate-chain
// tests; every field lands in cg_contracts (impl2 arrives via cg_adm_upgraded).
type vcContracts struct {
	game, nft, cst, dao      common.Address
	charity, prizes, rwalk   common.Address
	stakingCst, stakingRwalk common.Address
	marketing, implV1, impl2 common.Address
	owner                    common.Address
}

func newVCContracts() vcContracts {
	a := func(last byte) common.Address {
		return common.HexToAddress(fmt.Sprintf("0x%038x%02x", 0x51, last))
	}
	return vcContracts{
		game: a(0x01), nft: a(0x02), cst: a(0x03), dao: a(0x04),
		charity: a(0x05), prizes: a(0x06), rwalk: a(0x07),
		stakingCst: a(0x08), stakingRwalk: a(0x09),
		marketing: a(0x0a), implV1: a(0x0b), impl2: a(0x0c),
		owner: a(0x0d),
	}
}

func lowerHex(a common.Address) string { return strings.ToLower(a.Hex()) }

// vcFixture seeds a database so that the chain stubs from registerVCStubs
// agree with it on every check, and provides primitives for the log-parity
// scenarios.
type vcFixture struct {
	t    *testing.T
	ctx  context.Context
	pool *pgxpool.Pool
	a    vcContracts
	aids map[common.Address]int64
}

func newVCFixture(t *testing.T, db *testdb.DB, a vcContracts) *vcFixture {
	t.Helper()
	f := &vcFixture{
		t: t, ctx: context.Background(), pool: db.Pool, a: a,
		aids: map[common.Address]int64{},
	}
	for _, addr := range []common.Address{
		a.game, a.nft, a.cst, a.dao, a.charity, a.prizes, a.rwalk,
		a.stakingCst, a.stakingRwalk, a.marketing, a.implV1, a.impl2, a.owner,
	} {
		f.aids[addr] = f.insertAddress(addr)
	}
	f.exec(`INSERT INTO cg_contracts(
			cosmic_game_addr, cosmic_signature_addr, cosmic_token_addr,
			cosmic_dao_addr, charity_wallet_addr, prizes_wallet_addr,
			random_walk_addr, staking_wallet_cst_addr, staking_wallet_rwalk_addr,
			marketing_wallet_addr, implementation_addr)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		lowerHex(a.game), lowerHex(a.nft), lowerHex(a.cst), lowerHex(a.dao),
		lowerHex(a.charity), lowerHex(a.prizes), lowerHex(a.rwalk),
		lowerHex(a.stakingCst), lowerHex(a.stakingRwalk),
		lowerHex(a.marketing), lowerHex(a.implV1))
	f.exec(`INSERT INTO cg_proc_status(last_evt_id, last_block_num) VALUES (0, 200)`)
	// A proxy upgrade: impl2 joins the tracked-contract list.
	f.exec(`INSERT INTO cg_adm_upgraded(
			block_num, tx_id, time_stamp, contract_aid, implementation_aid)
		VALUES (120, 0, TO_TIMESTAMP(1700000120), $1, $2)`,
		f.aids[a.game], f.aids[a.impl2])
	return f
}

func (f *vcFixture) exec(sql string, args ...any) {
	f.t.Helper()
	if _, err := f.pool.Exec(f.ctx, sql, args...); err != nil {
		f.t.Fatalf("fixture %q: %v", sql[:min(60, len(sql))], err)
	}
}

func (f *vcFixture) insertAddress(addr common.Address) int64 {
	f.t.Helper()
	var aid int64
	if err := f.pool.QueryRow(f.ctx,
		`INSERT INTO address(block_num, tx_id, addr) VALUES (0, 0, $1)
		 RETURNING address_id`, lowerHex(addr)).Scan(&aid); err != nil {
		f.t.Fatalf("fixture address %s: %v", addr, err)
	}
	return aid
}

func (f *vcFixture) insertBlock(blockNum int64) {
	f.t.Helper()
	f.exec(`INSERT INTO block(block_num, block_hash, parent_hash, ts)
		VALUES ($1, $2, $3, TO_TIMESTAMP($4)) ON CONFLICT DO NOTHING`,
		blockNum,
		fmt.Sprintf("0x%063xb", blockNum),
		fmt.Sprintf("0x%063xa", blockNum),
		1700000000+blockNum)
}

func (f *vcFixture) insertTx(blockNum int64, txHash string) int64 {
	f.t.Helper()
	f.insertBlock(blockNum)
	var txID int64
	if err := f.pool.QueryRow(f.ctx,
		`INSERT INTO transaction(block_num, from_aid, to_aid, tx_hash)
		 VALUES ($1, $2, $2, $3) RETURNING id`,
		blockNum, f.aids[f.a.owner], txHash).Scan(&txID); err != nil {
		f.t.Fatalf("fixture transaction %s: %v", txHash, err)
	}
	return txID
}

// insertEvtLogs writes n evt_log rows for contract at blockNum on txHash.
func (f *vcFixture) insertEvtLogs(blockNum int64, txHash string, contract common.Address, n int) {
	f.t.Helper()
	txID := f.insertTx(blockNum, txHash)
	for i := 0; i < n; i++ {
		f.exec(`INSERT INTO evt_log(block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp)
			VALUES ($1, $2, $3, '81f6a773', $4, '\x00')`,
			blockNum, txID, f.aids[contract], i)
	}
}

// seedStateChecks adds rows so every contract-state check matches the stubs
// from registerVCStubs: CST supply 600, 2 NFTs owned by owner, 1 staked NFT
// per wallet, prizes balance 60 wei, round 0 legacy (2 bids) and round 1 V3
// (one 7-wei ETH bid).
func (f *vcFixture) seedStateChecks() {
	f.t.Helper()
	a := f.a
	ts := int64(1700000050)
	f.exec(`INSERT INTO cg_erc20_transfer(
			block_num, tx_id, time_stamp, contract_aid, value, from_aid, to_aid, otype)
		VALUES (50, 0, TO_TIMESTAMP($1), $2, 1000, $3, $3, 1),
		       (51, 0, TO_TIMESTAMP($1), $2, 400,  $3, $3, 2)`,
		ts, f.aids[a.cst], f.aids[a.owner])
	f.exec(`INSERT INTO cg_mint_event(
			block_num, tx_id, time_stamp, contract_aid, owner_aid, token_id,
			cur_owner_aid, round_num, seed)
		VALUES (52, 0, TO_TIMESTAMP($1), $2, $3, 0, $3, 0, 'seed0'),
		       (53, 0, TO_TIMESTAMP($1), $2, $3, 1, $3, 0, 'seed1')`,
		ts, f.aids[a.nft], f.aids[a.owner])
	f.exec(`INSERT INTO cg_nft_staked_cst(
			block_num, tx_id, time_stamp, contract_aid, action_id, token_id,
			num_staked_nfts, reward_per_staker, staker_aid)
		VALUES (54, 0, TO_TIMESTAMP($1), $2, 1, 0, 1, 0, $3),
		       (55, 0, TO_TIMESTAMP($1), $2, 2, 1, 2, 0, $3)`,
		ts, f.aids[a.stakingCst], f.aids[a.owner])
	f.exec(`INSERT INTO cg_nft_unstaked_cst(
			block_num, tx_id, time_stamp, contract_aid, action_id, action_counter,
			token_id, num_staked_nfts, staker_aid, reward, reward_per_tok)
		VALUES (56, 0, TO_TIMESTAMP($1), $2, 1, 1, 0, 1, $3, 0, 0)`,
		ts, f.aids[a.stakingCst], f.aids[a.owner])
	f.exec(`INSERT INTO cg_nft_staked_rwalk(
			block_num, tx_id, time_stamp, contract_aid, action_id, token_id,
			num_staked_nfts, staker_aid)
		VALUES (57, 0, TO_TIMESTAMP($1), $2, 3, 5, 1, $3)`,
		ts, f.aids[a.stakingRwalk], f.aids[a.owner])
	f.exec(`INSERT INTO cg_prize_deposit(
			block_num, tx_id, time_stamp, contract_aid, winner_aid, round_num,
			winner_index, amount)
		VALUES (58, 0, TO_TIMESTAMP($1), $2, $3, 0, 0, 100)`,
		ts, f.aids[a.prizes], f.aids[a.owner])
	f.exec(`INSERT INTO cg_prize_withdrawal(
			block_num, tx_id, time_stamp, contract_aid, round_num, winner_aid,
			beneficiary_aid, amount)
		VALUES (59, 0, TO_TIMESTAMP($1), $2, 0, $3, $3, 40)`,
		ts, f.aids[a.prizes], f.aids[a.owner])
	// The V3 reinitialize landed at block 100: round 0 (first bid at 50) is
	// legacy numBids-only, round 1 (first bid at 150) carries full stats.
	f.exec(`INSERT INTO cg_adm_initialized(
			block_num, tx_id, time_stamp, contract_aid, version)
		VALUES (100, 0, TO_TIMESTAMP($1), $2, 3)`, ts, f.aids[a.game])
	f.exec(`INSERT INTO cg_bid(
			block_num, tx_id, time_stamp, contract_aid, bidder_aid, rwalk_nft_id,
			round_num, bid_type, bid_position, prize_time, eth_price, cst_price)
		VALUES (50,  0, TO_TIMESTAMP($1), $2, $3, -1, 0, 0, 1, TO_TIMESTAMP($1), 10, -1),
		       (51,  0, TO_TIMESTAMP($1), $2, $3, -1, 0, 2, 2, TO_TIMESTAMP($1), -1, 5),
		       (150, 0, TO_TIMESTAMP($1), $2, $3, -1, 1, 0, 1, TO_TIMESTAMP($1), 7, -1)`,
		ts, f.aids[a.game], f.aids[a.owner])
}

// registerVCStubs scripts the chain so every contract-state check agrees with
// seedStateChecks.
func registerVCStubs(chain *testchain.Chain, a vcContracts) {
	cstStub := testchain.MustContractStub(cgc.CosmicSignatureTokenABI).
		Return("totalSupply", big.NewInt(600))
	chain.RegisterCall(a.cst, cstStub.Handler())

	nftStub := testchain.MustContractStub(cgc.CosmicSignatureNftABI).
		Return("totalSupply", big.NewInt(2)).
		Handle("ownerOf", func(args []any) ([]any, error) {
			return []any{a.owner}, nil
		})
	chain.RegisterCall(a.nft, nftStub.Handler())

	for _, wallet := range []common.Address{a.stakingCst, a.stakingRwalk} {
		stub := testchain.MustContractStub(cgc.StakingWalletNftBaseABI).
			Return("numStakedNfts", big.NewInt(1))
		chain.RegisterCall(wallet, stub.Handler())
	}
	chain.SetBalance(a.prizes, big.NewInt(60))

	gameStub := testchain.MustContractStub(cgc.CosmicSignatureGameV3ABI).
		Return("roundNum", big.NewInt(1)).
		Handle("roundStats", func(args []any) ([]any, error) {
			round := args[0].(*big.Int).Int64()
			z := big.NewInt(0)
			if round == 0 { // legacy round: only numBids is compared
				return []any{big.NewInt(2), z, z, z, z, z, z, z, z}, nil
			}
			return []any{
				big.NewInt(1), z, big.NewInt(7), z, big.NewInt(7), z, z, z, z,
			}, nil
		})
	chain.RegisterCall(a.game, gameStub.Handler())
}

// runValidateChain executes `opsctl validate-chain` against the given chain
// and database, returning the error and the combined output.
func runValidateChain(t *testing.T, chain *testchain.Chain, db *testdb.DB, args ...string) (error, string) {
	t.Helper()
	t.Setenv("RPC_URL", chain.URL())
	t.Setenv("DATABASE_URL", db.ConnString)
	// Flag defaults read these at command construction; isolate from the
	// developer's environment.
	for _, v := range []string{"ARCHIVE_DIR", "ARCHIVE_CONFIG", "FULL_LOG_PARITY",
		"VALIDATE_CHAIN_STATE", "LOG_WORKERS", "LOG_CHUNK"} {
		t.Setenv(v, "")
	}
	var out strings.Builder
	root := newRootCmd()
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs(append([]string{"validate-chain"}, args...))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	err := root.ExecuteContext(ctx)
	return err, out.String()
}

func TestValidateChainAllChecksPass(t *testing.T) {
	a := newVCContracts()
	db := testdb.New(t)
	f := newVCFixture(t, db, a)
	f.seedStateChecks()
	chain := testchain.New(t)
	chain.EnsureBlock(200)
	registerVCStubs(chain, a)

	err, out := runValidateChain(t, chain, db)
	if err != nil {
		t.Fatalf("validate-chain: %v\n%s", err, out)
	}
	for _, want := range []string{
		"log parity skipped",
		"CST totalSupply vs mints-burns | chain=600 db=600",
		"CS NFT totalSupply vs cg_mint_event count | chain=2 db=2",
		"ownerOf() vs cur_owner for 2 NFTs",
		"numStakedNfts (CST) vs stake-unstake events | chain=1 db=1",
		"numStakedNfts (RandomWalk) vs stake-unstake events | chain=1 db=1",
		"PrizesWallet ETH balance vs deposits-withdrawals | chain=60 db=60",
		"roundStats(0) vs cg_bid (legacy numBids-only)",
		"roundStats(1) vs cg_bid (V3 full)",
		"CHAIN VALIDATION OK",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "FAIL") {
		t.Errorf("unexpected FAIL:\n%s", out)
	}
}

func TestValidateChainDetectsSupplyMismatch(t *testing.T) {
	a := newVCContracts()
	db := testdb.New(t)
	f := newVCFixture(t, db, a)
	f.seedStateChecks()
	chain := testchain.New(t)
	chain.EnsureBlock(200)
	registerVCStubs(chain, a)
	// Overwrite the CST stub: the chain now claims one token too many.
	badCst := testchain.MustContractStub(cgc.CosmicSignatureTokenABI).
		Return("totalSupply", big.NewInt(601))
	chain.RegisterCall(a.cst, badCst.Handler())

	err, out := runValidateChain(t, chain, db)
	if !errors.Is(err, ErrChainValidationFailed) {
		t.Fatalf("expected ErrChainValidationFailed, got %v\n%s", err, out)
	}
	if !strings.Contains(out, "FAIL | CST totalSupply vs mints-burns | chain=601 db=600") ||
		!strings.Contains(out, "CHAIN VALIDATION FAILED") {
		t.Errorf("missing failure report:\n%s", out)
	}
}

func TestValidateChainRequiresRPCURL(t *testing.T) {
	t.Setenv("RPC_URL", "")
	root := newRootCmd()
	root.SetOut(&strings.Builder{})
	root.SetErr(&strings.Builder{})
	root.SetArgs([]string{"validate-chain"})
	if err := root.Execute(); err == nil ||
		!strings.Contains(err.Error(), "RPC_URL must be set") {
		t.Fatalf("expected RPC_URL error, got %v", err)
	}
}

func writeVCCollectorConfig(t *testing.T, a vcContracts) string {
	t.Helper()
	// Everything except the marketing wallet is in archive scope; the
	// validator reports the gap as a WARN.
	cfg := map[string]any{"output_dir": t.TempDir(), "contract_addresses": []string{
		lowerHex(a.game), lowerHex(a.nft), lowerHex(a.cst), lowerHex(a.dao),
		lowerHex(a.charity), lowerHex(a.prizes), lowerHex(a.rwalk),
		lowerHex(a.stakingCst), lowerHex(a.stakingRwalk),
		lowerHex(a.implV1), lowerHex(a.impl2),
	}}
	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "collector.json")
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestValidateChainArchiveLogParity(t *testing.T) {
	a := newVCContracts()
	db := testdb.New(t)
	f := newVCFixture(t, db, a)
	f.seedStateChecks()
	chain := testchain.New(t)
	chain.EnsureBlock(200)
	registerVCStubs(chain, a)

	// Archive: tx 0xaa at block 150 emitted two game logs, tx 0xbb at 160
	// one RandomWalk log. The database agrees.
	dir := t.TempDir()
	writeReceiptFixture(t, dir, "150", "0xaa", receiptBlob(t, a.game, a.game))
	writeReceiptFixture(t, dir, "160", "0xbb", receiptBlob(t, a.rwalk))
	f.insertEvtLogs(150, "0xaa", a.game, 2)
	f.insertEvtLogs(160, "0xbb", a.rwalk, 1)

	err, out := runValidateChain(t, chain, db,
		"--archive-dir", dir, "--archive-config", writeVCCollectorConfig(t, a))
	if err != nil {
		t.Fatalf("validate-chain --archive-dir: %v\n%s", err, out)
	}
	for _, want := range []string{
		"archive readable (2 receipts <= block 200)",
		"log parity (archive): game proxy | archive=2 db=2",
		"log parity (archive): RandomWalkNFT | archive=1 db=1",
		"WARN | log parity (archive): MarketingWallet | not in collector config",
		"CHAIN VALIDATION OK",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestValidateChainArchiveLogParityMismatch(t *testing.T) {
	a := newVCContracts()
	db := testdb.New(t)
	f := newVCFixture(t, db, a)
	f.seedStateChecks()
	chain := testchain.New(t)
	chain.EnsureBlock(200)
	registerVCStubs(chain, a)

	// The database has an NFT log the archive never captured.
	f.insertEvtLogs(150, "0xdd", a.nft, 1)
	dir := t.TempDir()
	writeReceiptFixture(t, dir, "150", "0xaa", receiptBlob(t, a.game))
	f.insertEvtLogs(150, "0xaa", a.game, 1)

	err, out := runValidateChain(t, chain, db,
		"--archive-dir", dir, "--archive-config", writeVCCollectorConfig(t, a))
	if !errors.Is(err, ErrChainValidationFailed) {
		t.Fatalf("expected ErrChainValidationFailed, got %v\n%s", err, out)
	}
	if !strings.Contains(out, "log parity (archive): CosmicSignatureNft | archive=0 db=1") ||
		!strings.Contains(out, "tx 0xdd: archive=0 db=1") {
		t.Errorf("missing archive mismatch report:\n%s", out)
	}
}

func TestValidateChainArchiveRequiresConfig(t *testing.T) {
	a := newVCContracts()
	db := testdb.New(t)
	f := newVCFixture(t, db, a)
	f.seedStateChecks()
	chain := testchain.New(t)
	chain.EnsureBlock(200)

	err, out := runValidateChain(t, chain, db, "--archive-dir", t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "--archive-config") {
		t.Fatalf("expected archive-config error, got %v\n%s", err, out)
	}
}

// attachChainLog mines a transaction at blockNum emitting one log for addr,
// so eth_getLogs sees it.
func attachChainLog(chain *testchain.Chain, blockNum int64, addr common.Address) {
	tx := chain.AddTx(blockNum, addr, nil)
	chain.AttachLogs(tx.Hash(), []*types.Log{{
		Address:     addr,
		BlockNumber: uint64(blockNum),
		BlockHash:   chain.BlockHash(blockNum),
		TxHash:      tx.Hash(),
	}})
}

func TestValidateChainNetworkLogParity(t *testing.T) {
	a := newVCContracts()
	db := testdb.New(t)
	f := newVCFixture(t, db, a)
	f.seedStateChecks()
	chain := testchain.New(t)
	chain.EnsureBlock(200)
	registerVCStubs(chain, a)

	// RandomWalk predates the platform (block 50 vs the first game log at
	// 150): the scan needs the separate RandomWalk prefix pass.
	attachChainLog(chain, 50, a.rwalk)
	attachChainLog(chain, 150, a.game)
	attachChainLog(chain, 160, a.rwalk)
	f.insertEvtLogs(50, "0x50aa", a.rwalk, 1)
	f.insertEvtLogs(150, "0x150aa", a.game, 1)
	f.insertEvtLogs(160, "0x160aa", a.rwalk, 1)

	stateFile := filepath.Join(t.TempDir(), "scan-state.json")
	// A garbage pre-existing state file is ignored, not fatal.
	if err := os.WriteFile(stateFile, []byte("{not json"), 0o600); err != nil {
		t.Fatal(err)
	}
	// One failing eth_getLogs answer exercises the bisecting retry.
	chain.FailNextRPC("eth_getLogs", "too many results")

	args := []string{"--full-log-parity", "--state-file", stateFile,
		"--log-chunk", "30", "--log-workers", "2"}
	err, out := runValidateChain(t, chain, db, args...)
	if err != nil {
		t.Fatalf("validate-chain --full-log-parity: %v\n%s", err, out)
	}
	for _, want := range []string{
		"log parity: game proxy | chain=1 db=1",
		"log parity: RandomWalkNFT | chain=2 db=2",
		"log parity: MarketingWallet | chain=0 db=0",
		"CHAIN VALIDATION OK",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}

	// Second run: the checkpoint already covers the whole range, so the scan
	// resumes with zero chunks and reports the same counts.
	err, out = runValidateChain(t, chain, db, args...)
	if err != nil {
		t.Fatalf("checkpoint resume: %v\n%s", err, out)
	}
	if !strings.Contains(out, "log parity: RandomWalkNFT | chain=2 db=2") {
		t.Errorf("resume lost checkpointed counts:\n%s", out)
	}

	// Push the checkpoint past the DB head: every contract degrades to WARN.
	state := loadChainScanState(stateFile)
	for _, entry := range state {
		entry.ScannedTo = 10_000
	}
	if err := saveChainScanState(stateFile, state); err != nil {
		t.Fatal(err)
	}
	err, out = runValidateChain(t, chain, db, args...)
	if err != nil {
		t.Fatalf("checkpoint-ahead run: %v\n%s", err, out)
	}
	if !strings.Contains(out, "WARN | log parity: game proxy | checkpoint ahead of DB head") {
		t.Errorf("missing checkpoint-ahead WARN:\n%s", out)
	}
}

func TestValidateChainNetworkLogParityNoRwalkLogs(t *testing.T) {
	a := newVCContracts()
	db := testdb.New(t)
	f := newVCFixture(t, db, a)
	f.seedStateChecks()
	chain := testchain.New(t)
	chain.EnsureBlock(200)
	registerVCStubs(chain, a)

	attachChainLog(chain, 150, a.game)
	f.insertEvtLogs(150, "0x150aa", a.game, 1)

	stateFile := filepath.Join(t.TempDir(), "scan-state.json")
	err, out := runValidateChain(t, chain, db, "--full-log-parity",
		"--state-file", stateFile, "--log-chunk", "100")
	if err != nil {
		t.Fatalf("validate-chain: %v\n%s", err, out)
	}
	if !strings.Contains(out, "WARN | log parity: RandomWalkNFT | no indexed logs; skipped") {
		t.Errorf("missing no-logs WARN:\n%s", out)
	}
}

func TestValidateChainNetworkScanPropagatesGetLogsError(t *testing.T) {
	a := newVCContracts()
	db := testdb.New(t)
	f := newVCFixture(t, db, a)
	f.seedStateChecks()
	chain := testchain.New(t)
	chain.EnsureBlock(200)
	registerVCStubs(chain, a)

	attachChainLog(chain, 150, a.game)
	f.insertEvtLogs(150, "0x150aa", a.game, 1)

	// With one-block chunks the bisection cannot split a failing range, so
	// the scan reports the error.
	chain.FailNextRPC("eth_getLogs", "node unavailable")
	stateFile := filepath.Join(t.TempDir(), "scan-state.json")
	err, out := runValidateChain(t, chain, db, "--full-log-parity",
		"--state-file", stateFile, "--log-chunk", "1", "--log-workers", "1")
	if err == nil || !strings.Contains(err.Error(), "getLogs") {
		t.Fatalf("expected getLogs error, got %v\n%s", err, out)
	}
}
