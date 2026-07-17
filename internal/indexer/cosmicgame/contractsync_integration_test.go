//go:build integration

// Integration tests for the startup contract-parameter sync
// (syncContractParamsFromChain): against the real database and a stubbed
// game/prizes-wallet contract, the first run inserts correction rows for
// every readable parameter, a repeated run with unchanged chain state inserts
// nothing (and leaves the address table untouched), and a changed chain value
// produces exactly one new correction.
package cosmicgame

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// syncedAdminTables lists every table the V1-mechanics sync writes.
var syncedAdminTables = []string{
	"cg_adm_erc20_reward",
	"cg_adm_cst_auclen",
	"cg_adm_timeout_claimprize",
	"cg_adm_price_inc",
	"cg_adm_time_inc",
	"cg_adm_prize_microsec",
	"cg_adm_inisecprize",
	"cg_adm_eth_auclen",
	"cg_adm_eth_auc_endprice",
	"cg_adm_cst_min_limit",
	"cg_adm_mkt_reward",
	"cg_delay_duration",
	"cg_adm_timeout_withdraw",
}

func tableCounts(t *testing.T, tables []string) map[string]int {
	t.Helper()
	out := make(map[string]int, len(tables))
	for _, table := range tables {
		var n int
		if err := testDB.SQL.QueryRow("SELECT COUNT(*) FROM " + table).Scan(&n); err != nil {
			t.Fatalf("counting %s: %v", table, err)
		}
		out[table] = n
	}
	return out
}

// v1GameStub stubs every V1 parameter read the sync performs. The CST reward
// matches the harness seed value so the glob-stats side stays consistent.
func v1GameStub() *testchain.ContractStub {
	return testchain.MustContractStub(cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI).
		Return("cstDutchAuctionDurationDivisor", big.NewInt(400)). // mechanics probe -> V1
		Return("cstRewardAmountForBidding", eth(100)).
		Return("timeoutDurationToClaimMainPrize", big.NewInt(86400)).
		Return("ethBidPriceIncreaseDivisor", big.NewInt(100)).
		Return("mainPrizeTimeIncrementIncreaseDivisor", big.NewInt(50)).
		Return("mainPrizeTimeIncrementInMicroSeconds", big.NewInt(3_600_000_000)).
		Return("initialDurationUntilMainPrizeDivisor", big.NewInt(200)).
		Return("ethDutchAuctionDurationDivisor", big.NewInt(40)).
		Return("ethDutchAuctionEndingBidPriceDivisor", big.NewInt(10)).
		Return("cstDutchAuctionBeginningBidPriceMinLimit", eth(200)).
		Return("marketingWalletCstContributionAmount", eth(300)).
		Return("delayDurationBeforeRoundActivation", big.NewInt(1800))
}

func prizesWalletStub() *testchain.ContractStub {
	return testchain.MustContractStub(cgc.PrizesWalletABI).
		Return("timeoutDurationToWithdrawPrizes", big.NewInt(604800))
}

func requireLatestParam(t *testing.T, table, column, want string) {
	t.Helper()
	got, hasRow, err := cgRepo.LatestDecimalParam(context.Background(), table, column)
	if err != nil {
		t.Fatalf("LatestDecimalParam(%s.%s): %v", table, column, err)
	}
	if !hasRow {
		t.Fatalf("%s has no correction row", table)
	}
	if got != want {
		t.Errorf("%s.%s = %s, want %s", table, column, got, want)
	}
}

func TestSyncContractParamsFromChain(t *testing.T) {
	resetDB(t)
	// The stub replaces the harness's game-contract call handler; restore it
	// for the fixture tests that follow.
	t.Cleanup(registerCallHandlers)
	// The sync anchors its correction rows to the latest chain block.
	testChain.EnsureBlock(500)

	gameStub := v1GameStub()
	testChain.RegisterCall(addr(fxGameAddr), gameStub.Handler())
	testChain.RegisterCall(addr(fxPrizesAddr), prizesWalletStub().Handler())

	logger := slog.New(slog.DiscardHandler)
	ctx := context.Background()

	// First run on a fresh database: every parameter gets a correction row.
	if err := SyncContractParams(ctx, cgRepo, dbStore, eclient, fxGameAddr, fxPrizesAddr, logger); err != nil {
		t.Fatalf("first sync run: %v", err)
	}

	requireLatestParam(t, "cg_adm_erc20_reward", "new_reward", eth(100).String())
	requireLatestParam(t, "cg_adm_cst_auclen", "new_len", "400")
	requireLatestParam(t, "cg_adm_timeout_claimprize", "new_timeout", "86400")
	requireLatestParam(t, "cg_adm_price_inc", "new_price_increase", "100")
	requireLatestParam(t, "cg_adm_mkt_reward", "new_reward", eth(300).String())
	requireLatestParam(t, "cg_delay_duration", "new_value", "1800")
	requireLatestParam(t, "cg_adm_timeout_withdraw", "new_timeout", "604800")

	globReward, err := cgRepo.GlobStatsCstRewardForBidding(ctx)
	if err != nil {
		t.Fatalf("GlobStatsCstRewardForBidding: %v", err)
	}
	if globReward != eth(100).String() {
		t.Errorf("glob stats CST reward = %s, want %s", globReward, eth(100).String())
	}

	// V1 mechanics: the V2-only auction-duration-change divisor is skipped.
	var chgDivRows int
	if err := testDB.SQL.QueryRow("SELECT COUNT(*) FROM cg_adm_cst_auclen_chg_div").Scan(&chgDivRows); err != nil {
		t.Fatalf("counting cg_adm_cst_auclen_chg_div: %v", err)
	}
	if chgDivRows != 0 {
		t.Errorf("cg_adm_cst_auclen_chg_div rows = %d, want 0 under V1 mechanics", chgDivRows)
	}

	for _, table := range syncedAdminTables {
		if n := tableCounts(t, []string{table})[table]; n != 1 {
			t.Errorf("%s rows after first run = %d, want 1", table, n)
		}
	}

	// Second run with identical chain state: check-then-correct inserts
	// nothing and the address table stays untouched.
	watched := append([]string{"address"}, syncedAdminTables...)
	before := tableCounts(t, watched)
	if err := SyncContractParams(ctx, cgRepo, dbStore, eclient, fxGameAddr, fxPrizesAddr, logger); err != nil {
		t.Fatalf("second sync run: %v", err)
	}
	after := tableCounts(t, watched)
	for _, table := range watched {
		if before[table] != after[table] {
			t.Errorf("%s rows changed on a clean re-run: %d -> %d", table, before[table], after[table])
		}
	}

	// An on-chain admin change: exactly that parameter gains a correction.
	gameStub.Return("ethBidPriceIncreaseDivisor", big.NewInt(125))
	if err := SyncContractParams(ctx, cgRepo, dbStore, eclient, fxGameAddr, fxPrizesAddr, logger); err != nil {
		t.Fatalf("third sync run: %v", err)
	}
	requireLatestParam(t, "cg_adm_price_inc", "new_price_increase", "125")
	final := tableCounts(t, watched)
	for _, table := range watched {
		want := after[table]
		if table == "cg_adm_price_inc" {
			want++
		}
		if final[table] != want {
			t.Errorf("%s rows after targeted change = %d, want %d", table, final[table], want)
		}
	}
}

// TestSyncContractParamsRejectsOverflowingHeaderTime pins the chain-boundary
// guard in allocChainSyncEvtlog: a header timestamp beyond int64 (corrupt
// node data) aborts the sync before any row is written, instead of wrapping
// into a negative correction timestamp. A private fake chain keeps the huge
// time offset away from the shared harness chain.
func TestSyncContractParamsRejectsOverflowingHeaderTime(t *testing.T) {
	resetDB(t)
	ctx := context.Background()

	chain := testchain.New(t)
	chain.EnsureBlock(500)
	chain.RegisterCall(addr(fxGameAddr), v1GameStub().Handler())
	chain.RegisterCall(addr(fxPrizesAddr), prizesWalletStub().Handler())

	rpcClient, err := ethrpc.DialContext(ctx, chain.URL())
	if err != nil {
		t.Fatalf("dialing fake chain: %v", err)
	}
	t.Cleanup(rpcClient.Close)
	var total uint64
	if err := rpcClient.CallContext(ctx, &total, "evm_increaseTime", int64(math.MaxInt64)); err != nil {
		t.Fatalf("evm_increaseTime: %v", err)
	}
	// A block built now carries BaseTime + offset, past math.MaxInt64.
	chain.EnsureBlock(501)

	client := ethclient.NewClient(rpcClient)
	err = SyncContractParams(ctx, cgRepo, dbStore, client, fxGameAddr, fxPrizesAddr, slog.New(slog.DiscardHandler))
	if err == nil || !strings.Contains(err.Error(), "overflows int64") {
		t.Fatalf("sync error = %v, want header-timestamp overflow", err)
	}

	var blocks int
	if err := testDB.SQL.QueryRow("SELECT COUNT(*) FROM block").Scan(&blocks); err != nil {
		t.Fatalf("counting blocks: %v", err)
	}
	if blocks != 0 {
		t.Errorf("block rows = %d, want 0 (guard must fire before InsertBlock)", blocks)
	}
}

// TestSyncContractParamsSkipsUnreadableParams pins the degraded mode: when a
// parameter read reverts, the sync logs and skips it instead of failing the
// whole startup.
func TestSyncContractParamsSkipsUnreadableParams(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)
	testChain.EnsureBlock(500)

	// Only two parameters answer; everything else reverts.
	stub := testchain.MustContractStub(cgc.CosmicSignatureGameABI, cgc.CosmicSignatureGameV2ABI).
		Return("cstDutchAuctionDurationDivisor", big.NewInt(400)).
		Return("cstRewardAmountForBidding", eth(100))
	testChain.RegisterCall(addr(fxGameAddr), stub.Handler())
	testChain.RegisterCall(addr(fxPrizesAddr), testchain.MustContractStub(cgc.PrizesWalletABI).Handler())

	ctx := context.Background()
	if err := SyncContractParams(ctx, cgRepo, dbStore, eclient, fxGameAddr, fxPrizesAddr, slog.New(slog.DiscardHandler)); err != nil {
		t.Fatalf("sync with unreadable params: %v", err)
	}

	requireLatestParam(t, "cg_adm_erc20_reward", "new_reward", eth(100).String())
	requireLatestParam(t, "cg_adm_cst_auclen", "new_len", "400")
	for _, table := range []string{"cg_adm_timeout_claimprize", "cg_adm_price_inc", "cg_adm_timeout_withdraw", "cg_delay_duration"} {
		var n int
		if err := testDB.SQL.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&n); err != nil {
			t.Fatalf("counting %s: %v", table, err)
		}
		if n != 0 {
			t.Errorf("%s rows = %d, want 0 (parameter was unreadable)", table, n)
		}
	}
}
