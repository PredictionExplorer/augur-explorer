//go:build integration

// Per-event golden fixtures for the RandomWalk ETL (§4.3): each of the seven
// dispatched event types is driven through the real pipeline and its database
// mutation pinned, plus a marketplace story and a reorg rollback/replay test.
//
// Unlike the CosmicGame handlers, the RandomWalk handlers insert without a
// preceding delete-by-evtlog-id: re-running process_single_event on the same
// evt_log row violates the tables' UNIQUE(evtlog_id) constraints and
// terminates the process. Re-processing therefore only happens through the
// reorg path (cascade delete + fresh evt_log rows), which is what
// TestReorgRollbackAndReplay exercises.
package randomwalk

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// fixtureLog and fixtureTx mirror the cg-etl fixture DSL.
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

// mintLog builds the MintEvent for tokenID owned by owner. Seed bytes are a
// deterministic function of the token id.
func mintLog(tokenID int64, owner string) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		var seed [32]byte
		seed[0] = 0x5e
		seed[31] = byte(tokenID)
		return buildLog(t, fxRwalkABI, "MintEvent", addr(fxRandomWalkAddr),
			[]any{bigInt(tokenID), addr(owner)}, []any{seed, eth(1)})
	}
}

// transferLog builds the RandomWalk ERC721 Transfer.
func transferLog(from, to string, tokenID int64) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxRwalkABI, "Transfer", addr(fxRandomWalkAddr),
			[]any{addr(from), addr(to), bigInt(tokenID)}, nil)
	}
}

// offerLog builds a marketplace NewOffer (sell offer when seller != 0x0).
func offerLog(offerID, tokenID int64, seller, buyer string, priceEth int64) func(t *testing.T) *types.Log {
	return func(t *testing.T) *types.Log {
		return buildLog(t, fxMarketABI, "NewOffer", addr(fxMarketplaceAddr),
			[]any{addr(fxRandomWalkAddr), bigInt(offerID), bigInt(tokenID)},
			[]any{addr(seller), addr(buyer), eth(priceEth)})
	}
}

func eventFixtures() []fixture {
	market := fxMarketplaceAddr
	rwalk := fxRandomWalkAddr

	return []fixture{
		{name: "mint_event", block: 1000, txs: []fixtureTx{{to: rwalk, logs: []fixtureLog{
			{TopicTransferEvt, transferLog(fxZero, fxCarol, 10)},
			{TopicMintEvent, mintLog(10, fxCarol)},
		}}}},
		{name: "token_name_after_mint", block: 1010, txs: []fixtureTx{
			{blockOffset: 0, to: rwalk, logs: []fixtureLog{
				{TopicTransferEvt, transferLog(fxZero, fxCarol, 10)},
				{TopicMintEvent, mintLog(10, fxCarol)},
			}},
			{blockOffset: 1, to: rwalk, logs: []fixtureLog{
				{TopicTokenNameEvt, func(t *testing.T) *types.Log {
					return buildLog(t, fxRwalkABI, "TokenNameEvent", addr(rwalk),
						nil, []any{bigInt(10), "Fixture Walker"})
				}},
			}},
		}},
		{name: "token_name_unknown_token_skipped", block: 1020, txs: []fixtureTx{{to: rwalk, logs: []fixtureLog{
			{TopicTokenNameEvt, func(t *testing.T) *types.Log {
				// No mint for token 99: RWalk_token_exists fails and the
				// handler skips the insert (layer-1 rows still land).
				return buildLog(t, fxRwalkABI, "TokenNameEvent", addr(rwalk),
					nil, []any{bigInt(99), "Ghost Token"})
			}},
		}}}},
		{name: "transfer_secondary", block: 1030, txs: []fixtureTx{
			{blockOffset: 0, to: rwalk, logs: []fixtureLog{
				{TopicTransferEvt, transferLog(fxZero, fxCarol, 10)},
				{TopicMintEvent, mintLog(10, fxCarol)},
			}},
			{blockOffset: 1, to: rwalk, logs: []fixtureLog{
				{TopicTransferEvt, transferLog(fxCarol, fxDave, 10)},
			}},
		}},
		{name: "transfer_wrong_address_skipped", block: 1040, txs: []fixtureTx{{to: rwalk, logs: []fixtureLog{
			{TopicTransferEvt, func(t *testing.T) *types.Log {
				return buildLog(t, fxRwalkABI, "Transfer", addr(fxAlice),
					[]any{addr(fxCarol), addr(fxDave), bigInt(10)}, nil)
			}},
		}}}},
		{name: "new_offer_sell", block: 1050, txs: []fixtureTx{
			{blockOffset: 0, to: rwalk, logs: []fixtureLog{
				{TopicTransferEvt, transferLog(fxZero, fxCarol, 10)},
				{TopicMintEvent, mintLog(10, fxCarol)},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				{TopicNewOffer, offerLog(1, 10, fxCarol, fxZero, 3)},
			}},
		}},
		{name: "new_offer_buy", block: 1060, txs: []fixtureTx{
			{blockOffset: 0, to: rwalk, logs: []fixtureLog{
				{TopicTransferEvt, transferLog(fxZero, fxCarol, 10)},
				{TopicMintEvent, mintLog(10, fxCarol)},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				// Buy offer: seller is the zero address, buyer bids for the token.
				{TopicNewOffer, offerLog(2, 10, fxZero, fxDave, 2)},
			}},
		}},
		{name: "item_bought", block: 1070, txs: []fixtureTx{
			{blockOffset: 0, to: rwalk, logs: []fixtureLog{
				{TopicTransferEvt, transferLog(fxZero, fxCarol, 10)},
				{TopicMintEvent, mintLog(10, fxCarol)},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				{TopicNewOffer, offerLog(1, 10, fxCarol, fxZero, 3)},
			}},
			{blockOffset: 2, to: market, logs: []fixtureLog{
				{TopicItemBought, func(t *testing.T) *types.Log {
					return buildLog(t, fxMarketABI, "ItemBought", addr(market),
						[]any{bigInt(1), addr(fxCarol), addr(fxDave)}, nil)
				}},
				{TopicTransferEvt, transferLog(fxCarol, fxDave, 10)},
			}},
		}},
		{name: "offer_canceled", block: 1080, txs: []fixtureTx{
			{blockOffset: 0, to: rwalk, logs: []fixtureLog{
				{TopicTransferEvt, transferLog(fxZero, fxCarol, 10)},
				{TopicMintEvent, mintLog(10, fxCarol)},
			}},
			{blockOffset: 1, to: market, logs: []fixtureLog{
				{TopicNewOffer, offerLog(1, 10, fxCarol, fxZero, 3)},
			}},
			{blockOffset: 2, to: market, logs: []fixtureLog{
				{TopicOfferCanceled, func(t *testing.T) *types.Log {
					return buildLog(t, fxMarketABI, "OfferCanceled", addr(market),
						[]any{bigInt(1)}, nil)
				}},
			}},
		}},
		{name: "offer_canceled_unknown_offer_skipped", block: 1090, txs: []fixtureTx{{to: market, logs: []fixtureLog{
			{TopicOfferCanceled, func(t *testing.T) *types.Log {
				// Offer 77 was never created: Offer_exists fails and the
				// handler skips the insert.
				return buildLog(t, fxMarketABI, "OfferCanceled", addr(market),
					[]any{bigInt(77)}, nil)
			}},
		}}}},
		{name: "withdrawal_event", block: 1100, txs: []fixtureTx{
			{blockOffset: 0, to: rwalk, logs: []fixtureLog{
				{TopicTransferEvt, transferLog(fxZero, fxCarol, 10)},
				{TopicMintEvent, mintLog(10, fxCarol)},
			}},
			{blockOffset: 1, to: rwalk, logs: []fixtureLog{
				{TopicWithdrawalEvt, func(t *testing.T) *types.Log {
					return buildLog(t, fxRwalkABI, "WithdrawalEvent", addr(rwalk),
						[]any{bigInt(10)}, []any{addr(fxCarol), eth(1)})
				}},
			}},
		}},
	}
}

func TestEventFixtures(t *testing.T) {
	for _, fx := range eventFixtures() {
		t.Run(fx.name, func(t *testing.T) {
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

// marketplaceStory is the multi-block narrative used by the story golden and
// the reorg test: four mints, a name change, a sale, competing offers, a
// cancellation and a withdrawal. Block numbers are startBlock+blockOffset,
// applied by ingestStory.
func marketplaceStory() []fixtureTx {
	return []fixtureTx{
		{blockOffset: 0, to: fxRandomWalkAddr, logs: []fixtureLog{
			{TopicTransferEvt, transferLog(fxZero, fxCarol, 10)},
			{TopicMintEvent, mintLog(10, fxCarol)},
		}},
		{blockOffset: 1, to: fxRandomWalkAddr, logs: []fixtureLog{
			{TopicTransferEvt, transferLog(fxZero, fxDave, 11)},
			{TopicMintEvent, mintLog(11, fxDave)},
		}},
		{blockOffset: 2, to: fxRandomWalkAddr, logs: []fixtureLog{
			{TopicTransferEvt, transferLog(fxZero, fxAlice, 12)},
			{TopicMintEvent, mintLog(12, fxAlice)},
		}},
		{blockOffset: 3, to: fxRandomWalkAddr, logs: []fixtureLog{
			{TopicTransferEvt, transferLog(fxZero, fxBob, 13)},
			{TopicMintEvent, mintLog(13, fxBob)},
		}},
		{blockOffset: 4, to: fxRandomWalkAddr, logs: []fixtureLog{
			{TopicTokenNameEvt, func(t *testing.T) *types.Log {
				return buildLog(t, fxRwalkABI, "TokenNameEvent", addr(fxRandomWalkAddr),
					nil, []any{bigInt(10), "Story Walker"})
			}},
		}},
		{blockOffset: 5, to: fxMarketplaceAddr, logs: []fixtureLog{
			{TopicNewOffer, offerLog(1, 10, fxCarol, fxZero, 3)},
		}},
		{blockOffset: 6, to: fxMarketplaceAddr, logs: []fixtureLog{
			{TopicItemBought, func(t *testing.T) *types.Log {
				return buildLog(t, fxMarketABI, "ItemBought", addr(fxMarketplaceAddr),
					[]any{bigInt(1), addr(fxCarol), addr(fxDave)}, nil)
			}},
			{TopicTransferEvt, transferLog(fxCarol, fxDave, 10)},
		}},
		{blockOffset: 7, to: fxMarketplaceAddr, logs: []fixtureLog{
			{TopicNewOffer, offerLog(2, 11, fxDave, fxZero, 5)},
		}},
		{blockOffset: 8, to: fxMarketplaceAddr, logs: []fixtureLog{
			{TopicNewOffer, offerLog(3, 13, fxBob, fxZero, 4)},
		}},
		{blockOffset: 9, to: fxMarketplaceAddr, logs: []fixtureLog{
			{TopicNewOffer, offerLog(4, 10, fxDave, fxZero, 6)},
		}},
		{blockOffset: 10, to: fxMarketplaceAddr, logs: []fixtureLog{
			{TopicOfferCanceled, func(t *testing.T) *types.Log {
				return buildLog(t, fxMarketABI, "OfferCanceled", addr(fxMarketplaceAddr),
					[]any{bigInt(4)}, nil)
			}},
		}},
		{blockOffset: 11, to: fxRandomWalkAddr, logs: []fixtureLog{
			{TopicWithdrawalEvt, func(t *testing.T) *types.Log {
				return buildLog(t, fxRwalkABI, "WithdrawalEvent", addr(fxRandomWalkAddr),
					[]any{bigInt(10)}, []any{addr(fxCarol), eth(1)})
			}},
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

// TestMarketplaceStoryGolden pins the cumulative state of the marketplace
// narrative: mints, names, offers, a sale, a cancellation and a withdrawal.
func TestMarketplaceStoryGolden(t *testing.T) {
	resetDB(t)
	base := snapshot(t)

	ingestStory(t, 2000, marketplaceStory())

	after := snapshot(t)
	diff, err := testutil.DiffSnapshots(base, after)
	if err != nil {
		t.Fatalf("diffing snapshots: %v", err)
	}
	testutil.CompareGolden(t, filepath.Join("testdata", "golden", "story_marketplace.json"), diff)
}

// TestReorgRollbackAndReplay exercises the chain-split path for the
// RandomWalk domain: rollback from the sale onward (golden), then replay of
// the replacement fork must reproduce the original state byte for byte.
func TestReorgRollbackAndReplay(t *testing.T) {
	resetDB(t)
	const startBlock = 3000
	const divergentOffset = 6 // the ItemBought sale transaction
	const divergentBlock = startBlock + divergentOffset

	story := marketplaceStory()
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
