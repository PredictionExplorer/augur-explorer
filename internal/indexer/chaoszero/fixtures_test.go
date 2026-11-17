//go:build integration

// Per-event golden fixtures for the ChaosZero ETL: each of the thirteen
// dispatched event types is driven through the real pipeline and its
// database mutation pinned, plus a full round-lifecycle story and a reorg
// rollback/replay test.
//
// Like the RandomWalk handlers, the ChaosZero handlers insert without a
// preceding delete-by-evtlog-id: re-running process_single_event on the same
// evt_log row violates the tables' UNIQUE(evtlog_id) constraints and
// terminates the process. Re-processing therefore only happens through the
// reorg path (cascade delete + fresh evt_log rows), which is what
// TestReorgRollbackAndReplay exercises.
package chaoszero

import (
	"context"
	"math/big"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// fixtureLog and fixtureTx mirror the RandomWalk fixture DSL.
type fixtureLog struct {
	topic0 string
	build  func(t *testing.T) *types.Log
}

type fixtureTx struct {
	blockOffset int64
	to          string
	logs        []fixtureLog
}

type fixture struct {
	name  string
	block int64
	txs   []fixtureTx
}

// positionID encodes an ERC-1155 position token id the way
// IndependentOutcomeMarket.positionId does: (roundId << 4) | (question << 1) | yes.
func positionID(round, question int64, yes bool) *big.Int {
	id := round<<4 | question<<1
	if yes {
		id |= 1
	}
	return big.NewInt(id)
}

func initLog(round int64) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "RoundInitialized", addr(fxMarketAddr),
			[]any{bigInt(round)}, nil)
	}
}

func fundedLog(round int64, caller string, donated, transfersUsed, carryUsed int64, allocations [8]*big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "RoundFunded", addr(fxMarketAddr),
			[]any{bigInt(round), addr(caller)},
			[]any{cst(donated), cst(transfersUsed), cst(carryUsed), allocations})
	}
}

func tradeLog(round int64, user string, question uint8, yes, buy bool, cash, units *big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "Trade", addr(fxMarketAddr),
			[]any{bigInt(round), addr(user), question},
			[]any{yes, buy, cash, units})
	}
}

func splitLog(round int64, user string, question uint8, amount *big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "PositionSplit", addr(fxMarketAddr),
			[]any{bigInt(round), addr(user), question},
			[]any{amount})
	}
}

func mergeLog(round int64, user string, amounts [8]*big.Int, collateralOut *big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "PositionsMerged", addr(fxMarketAddr),
			[]any{bigInt(round), addr(user)},
			[]any{amounts, collateralOut})
	}
}

func settledLog(round int64, valuesMask uint8, escrowed [8]*big.Int, carried *big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "RoundSettled", addr(fxMarketAddr),
			[]any{bigInt(round)},
			[]any{valuesMask, escrowed, carried})
	}
}

func claimLog(round int64, user string, amounts [8]*big.Int, total *big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "Claimed", addr(fxMarketAddr),
			[]any{bigInt(round), addr(user)},
			[]any{amounts, total})
	}
}

func answersLog(round int64, knownMask, valuesMask uint8) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "AnswersObserved", addr(fxMarketAddr),
			[]any{bigInt(round)},
			[]any{knownMask, valuesMask})
	}
}

func escrowLog(round int64, amounts [8]*big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxVaultABI, "CycleEscrowed", addr(fxVaultAddr),
			[]any{bigInt(round)},
			[]any{amounts})
	}
}

func paidLog(round int64, question uint8, recipient string, amount *big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxVaultABI, "Paid", addr(fxVaultAddr),
			[]any{bigInt(round), question, addr(recipient)},
			[]any{amount})
	}
}

func transferSingleLog(operator, from, to string, id, value *big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "TransferSingle", addr(fxMarketAddr),
			[]any{addr(operator), addr(from), addr(to)},
			[]any{id, value})
	}
}

func transferBatchLog(operator, from, to string, ids, values []*big.Int) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "TransferBatch", addr(fxMarketAddr),
			[]any{addr(operator), addr(from), addr(to)},
			[]any{ids, values})
	}
}

func approvalLog(account, operator string, approved bool) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "ApprovalForAll", addr(fxMarketAddr),
			[]any{addr(account), addr(operator)},
			[]any{approved})
	}
}

func eventFixtures() []fixture {
	market := fxMarketAddr
	vault := fxVaultAddr
	even := cstArray(12, 12, 12, 12, 12, 12, 12, 13) // 97 CST across the eight pools

	return []fixture{
		{name: "round_initialized", block: 1000, txs: []fixtureTx{{to: market, logs: []fixtureLog{
			{TopicRoundInitialized, initLog(0)},
		}}}},
		{name: "round_funded", block: 1010, txs: []fixtureTx{
			{blockOffset: 0, to: market, logs: []fixtureLog{
				{TopicRoundInitialized, initLog(0)},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				{TopicRoundFunded, fundedLog(0, fxAlice, 97, 0, 0, even)},
			}},
		}},
		{name: "trades_buy_and_sell", block: 1020, txs: []fixtureTx{
			{blockOffset: 0, to: market, logs: []fixtureLog{
				{TopicRoundInitialized, initLog(0)},
				{TopicRoundFunded, fundedLog(0, fxAlice, 97, 0, 0, even)},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				{TopicTrade, tradeLog(0, fxAlice, 0, true, true, cst(12), cst(16))},
			}},
			{blockOffset: 2, to: market, logs: []fixtureLog{
				{TopicTrade, tradeLog(0, fxBob, 1, false, false, cst(5), cst(8))},
			}},
		}},
		{name: "position_split", block: 1030, txs: []fixtureTx{{to: market, logs: []fixtureLog{
			{TopicRoundInitialized, initLog(0)},
			{TopicPositionSplit, splitLog(0, fxCarol, 2, cst(20))},
		}}}},
		{name: "positions_merged", block: 1040, txs: []fixtureTx{
			{blockOffset: 0, to: market, logs: []fixtureLog{
				{TopicRoundInitialized, initLog(0)},
				{TopicPositionSplit, splitLog(0, fxCarol, 2, cst(20))},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				{TopicPositionsMerged, mergeLog(0, fxCarol, cstArray(0, 0, 5), cst(5))},
			}},
		}},
		{name: "round_settled", block: 1050, txs: []fixtureTx{
			{blockOffset: 0, to: market, logs: []fixtureLog{
				{TopicRoundInitialized, initLog(0)},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				{TopicRoundSettled, settledLog(0, 0b00000101, cstArray(20, 0, 15), cst(90))},
			}},
		}},
		{name: "claimed", block: 1060, txs: []fixtureTx{
			{blockOffset: 0, to: market, logs: []fixtureLog{
				{TopicRoundInitialized, initLog(0)},
				{TopicRoundSettled, settledLog(0, 0b00000001, cstArray(20), cst(77))},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				{TopicClaimed, claimLog(0, fxAlice, cstArray(17), cst(17))},
			}},
		}},
		{name: "answers_observed", block: 1070, txs: []fixtureTx{{to: market, logs: []fixtureLog{
			{TopicRoundInitialized, initLog(0)},
			{TopicAnswersObserved, answersLog(0, 0b00000011, 0b00000001)},
		}}}},
		{name: "cycle_escrowed", block: 1080, txs: []fixtureTx{{to: vault, logs: []fixtureLog{
			{TopicCycleEscrowed, escrowLog(0, cstArray(20, 0, 15))},
		}}}},
		{name: "paid", block: 1090, txs: []fixtureTx{
			{blockOffset: 0, to: vault, logs: []fixtureLog{
				{TopicCycleEscrowed, escrowLog(0, cstArray(20))},
			}},
			{blockOffset: 1, to: vault, logs: []fixtureLog{
				{TopicPaid, paidLog(0, 0, fxAlice, cst(17))},
			}},
		}},
		{name: "transfer_single_mint", block: 1100, txs: []fixtureTx{{to: market, logs: []fixtureLog{
			{TopicTransferSingle, transferSingleLog(market, fxZero, fxAlice, positionID(0, 0, true), cst(16))},
		}}}},
		{name: "transfer_batch", block: 1110, txs: []fixtureTx{{to: market, logs: []fixtureLog{
			{TopicTransferBatch, transferBatchLog(fxCarol, fxCarol, fxDave,
				[]*big.Int{positionID(0, 2, true), positionID(0, 2, false)},
				[]*big.Int{cst(2), cst(3)})},
		}}}},
		{name: "approval_for_all", block: 1120, txs: []fixtureTx{
			{blockOffset: 0, to: market, logs: []fixtureLog{
				{TopicApprovalForAll, approvalLog(fxDave, fxBob, true)},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				{TopicApprovalForAll, approvalLog(fxDave, fxBob, false)},
			}},
		}},
		{name: "trade_wrong_address_skipped", block: 1130, txs: []fixtureTx{{to: market, logs: []fixtureLog{
			// A Trade-shaped log emitted by a non-market address: the
			// registry's source filter skips it (layer-1 rows still land).
			{TopicTrade, func(t *testing.T) *types.Log {
				return buildLog(t, fxMarketABI, "Trade", addr(fxAlice),
					[]any{bigInt(0), addr(fxBob), uint8(1)},
					[]any{true, true, cst(1), cst(1)})
			}},
		}}}},
		{name: "vault_event_from_market_skipped", block: 1140, txs: []fixtureTx{{to: market, logs: []fixtureLog{
			// CycleEscrowed emitted by the market address: the handler
			// declares the vault as its only source, so it is skipped.
			{TopicCycleEscrowed, func(t *testing.T) *types.Log {
				return buildLog(t, fxVaultABI, "CycleEscrowed", addr(fxMarketAddr),
					[]any{bigInt(0)}, []any{cstArray(1)})
			}},
		}}}},
	}
}

func TestEventFixtures(t *testing.T) {
	for _, fx := range eventFixtures() {
		t.Run(fx.name, func(t *testing.T) {
			testChain.ClearTransactionsFrom(fx.block)
			resetDB(t)
			base := snapshot(t)

			logIndexByBlock := make(map[int64]uint)
			for _, ftx := range fx.txs {
				block := fx.block + ftx.blockOffset
				logs := make([]*types.Log, 0, len(ftx.logs))
				for _, fl := range ftx.logs {
					l := fl.build(t)
					if fl.topic0 != "" && len(l.Topics) > 0 {
						if got := l.Topics[0].Hex()[2:]; got != fl.topic0 {
							t.Fatalf("fixture builds topic0 %s, registry constant is %s", got, fl.topic0)
						}
					}
					logs = append(logs, l)
				}
				start := logIndexByBlock[block]
				ingestTx(t, block, addr(ftx.to), start, logs)
				logIndexByBlock[block] = start + uint(len(logs))
			}

			after := snapshot(t)
			diff, err := testutil.DiffSnapshots(base, after)
			if err != nil {
				t.Fatalf("diffing snapshots: %v", err)
			}
			testutil.CompareGolden(t, filepath.Join("testdata", "golden", fx.name+".json"), diff)
		})
	}
}

// marketStory is the multi-block narrative used by the story golden and the
// reorg test: one complete round lifecycle — initialize, fund, three trades,
// a split, a merge, an answer observation, settlement with its vault escrow,
// a claim with its vault payout — plus position transfers and approvals.
func marketStory() []fixtureTx {
	even := cstArray(12, 12, 12, 12, 12, 12, 12, 13)
	return []fixtureTx{
		{blockOffset: 0, to: fxMarketAddr, logs: []fixtureLog{
			{TopicRoundInitialized, initLog(0)},
		}},
		{blockOffset: 1, to: fxMarketAddr, logs: []fixtureLog{
			{TopicRoundFunded, fundedLog(0, fxAlice, 97, 0, 0, even)},
		}},
		{blockOffset: 2, to: fxMarketAddr, logs: []fixtureLog{
			{TopicTrade, tradeLog(0, fxAlice, 0, true, true, cst(12), cst(16))},
			{TopicTransferSingle, transferSingleLog(fxMarketAddr, fxZero, fxAlice, positionID(0, 0, true), cst(16))},
		}},
		{blockOffset: 3, to: fxMarketAddr, logs: []fixtureLog{
			{TopicTrade, tradeLog(0, fxBob, 1, false, true, cst(5), cst(8))},
		}},
		{blockOffset: 4, to: fxMarketAddr, logs: []fixtureLog{
			{TopicPositionSplit, splitLog(0, fxCarol, 2, cst(20))},
		}},
		{blockOffset: 5, to: fxMarketAddr, logs: []fixtureLog{
			{TopicTrade, tradeLog(0, fxCarol, 2, true, false, cst(7), cst(10))},
		}},
		{blockOffset: 6, to: fxMarketAddr, logs: []fixtureLog{
			{TopicPositionsMerged, mergeLog(0, fxCarol, cstArray(0, 0, 5), cst(5))},
		}},
		{blockOffset: 7, to: fxMarketAddr, logs: []fixtureLog{
			{TopicAnswersObserved, answersLog(0, 0b00000001, 0b00000001)},
		}},
		{blockOffset: 8, to: fxMarketAddr, logs: []fixtureLog{
			{TopicRoundSettled, settledLog(0, 0b00000101, cstArray(20, 0, 15), cst(90))},
			{TopicCycleEscrowed, func(t *testing.T) *types.Log {
				// Settlement escrows through the vault in the same tx.
				return escrowLog(0, cstArray(20, 0, 15))(t)
			}},
		}},
		{blockOffset: 9, to: fxMarketAddr, logs: []fixtureLog{
			{TopicClaimed, claimLog(0, fxAlice, cstArray(17), cst(17))},
			{TopicPaid, paidLog(0, 0, fxAlice, cst(17))},
		}},
		{blockOffset: 10, to: fxMarketAddr, logs: []fixtureLog{
			{TopicTransferSingle, transferSingleLog(fxAlice, fxAlice, fxDave, positionID(0, 0, true), cst(3))},
		}},
		{blockOffset: 11, to: fxMarketAddr, logs: []fixtureLog{
			{TopicTransferBatch, transferBatchLog(fxCarol, fxCarol, fxDave,
				[]*big.Int{positionID(0, 2, true), positionID(0, 2, false)},
				[]*big.Int{cst(2), cst(3)})},
		}},
		{blockOffset: 12, to: fxMarketAddr, logs: []fixtureLog{
			{TopicApprovalForAll, approvalLog(fxDave, fxBob, true)},
		}},
	}
}

func ingestStory(t *testing.T, startBlock int64, story []fixtureTx) {
	t.Helper()
	logIndexByBlock := make(map[int64]uint)
	for _, stx := range story {
		block := startBlock + stx.blockOffset
		logs := make([]*types.Log, 0, len(stx.logs))
		for _, fl := range stx.logs {
			logs = append(logs, fl.build(t))
		}
		start := logIndexByBlock[block]
		ingestTx(t, block, addr(stx.to), start, logs)
		logIndexByBlock[block] = start + uint(len(logs))
	}
}

// TestMarketStoryGolden pins the cumulative state of the round-lifecycle
// narrative.
func TestMarketStoryGolden(t *testing.T) {
	resetDB(t)
	base := snapshot(t)

	ingestStory(t, 2000, marketStory())

	after := snapshot(t)
	diff, err := testutil.DiffSnapshots(base, after)
	if err != nil {
		t.Fatalf("diffing snapshots: %v", err)
	}
	testutil.CompareGolden(t, filepath.Join("testdata", "golden", "story_market.json"), diff)
}

// TestReorgRollbackAndReplay exercises the chain-split path for the
// ChaosZero domain: rollback from the settlement onward (golden), then
// replay of the replacement fork must reproduce the original state byte for
// byte.
func TestReorgRollbackAndReplay(t *testing.T) {
	resetDB(t)
	const startBlock = 3000
	const divergentOffset = 8 // the settlement + escrow transaction
	const divergentBlock = startBlock + divergentOffset

	story := marketStory()
	ingestStory(t, startBlock, story)
	full := snapshot(t)

	testChain.Reorg(divergentBlock)
	newHash := testChain.BlockHash(divergentBlock)

	inserted, err := testIndexer.EnsureBlockExists(context.Background(), divergentBlock, newHash.Hex())
	if err != nil {
		t.Fatalf("EnsureBlockExists after reorg: %v", err)
	}
	if !inserted {
		t.Error("EnsureBlockExists did not re-insert the divergent block")
	}

	rolledBack := snapshot(t)
	diff, err := testutil.DiffSnapshots(rolledBack, full)
	if err != nil {
		t.Fatalf("diffing rollback state: %v", err)
	}
	testutil.CompareGolden(t, filepath.Join("testdata", "golden", "story_reorg_rollback.json"), diff)

	var replay []fixtureTx
	for _, stx := range story {
		if startBlock+stx.blockOffset >= divergentBlock {
			replay = append(replay, stx)
		}
	}
	ingestStory(t, startBlock, replay)

	requireNoDiff(t, full, snapshot(t), "state after reorg replay")
}
