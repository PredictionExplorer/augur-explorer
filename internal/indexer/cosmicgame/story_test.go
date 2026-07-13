//go:build integration

// Scripted-round golden test and reorg simulation (§4.3): a realistic event
// sequence for one full game round is processed through the production
// pipeline and its cumulative database state pinned; then a chain split rolls
// part of it back (golden) and re-processing the replacement fork must
// reproduce the exact original state.
package cosmicgame

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// storyTx is one transaction of the scripted round; unlike the per-event
// fixtures the story spans many blocks, so block numbers are absolute.
type storyTx struct {
	block int64
	to    string
	logs  []fixtureLog
}

// scriptedRound returns the event sequence of one complete round: admin
// bootstrap, three bid types, donations of every kind, the multi-log claim
// transaction (main prize, raffles, special prizes, charity flows), prize
// withdrawal, donated-prize claims, secondary transfers and staking activity.
// startBlock shifts the whole story so multiple tests can use disjoint ranges.
func scriptedRound(startBlock int64) []storyTx {
	b := func(offset int64) int64 { return startBlock + offset }
	zero := "0x0000000000000000000000000000000000000000"

	return []storyTx{
		// Admin bootstrap: reward parameter + round activation.
		{block: b(0), to: fxGameAddr, logs: []fixtureLog{
			{TopicCstRewardForBiddingChanged, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "CstRewardAmountForBiddingChanged", addr(fxGameAddr), nil, []any{eth(100)})
			}},
			{TopicRoundActivationTimeChanged, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "RoundActivationTimeChanged", addr(fxGameAddr), nil, []any{bigInt(1767225600)})
			}},
		}},
		// Round 0 starts: alice places the first (ETH) bid, reward minted.
		{block: b(1), to: fxGameAddr, logs: []fixtureLog{
			{TopicFirstBidEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "FirstBidPlacedInRound", addr(fxGameAddr), []any{bigInt(0)}, []any{bigInt(1767225700)})
			}},
			{TopicTransferEvt, func(t *testing.T) *types.Log {
				return buildLog(t, erc20ABI, "Transfer", addr(fxTokenAddr), []any{addr(zero), addr(fxAlice)}, []any{eth(100)})
			}},
			{TopicBidEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "BidPlaced", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxAlice), bigInt(-1)},
					[]any{bigInt(100000000000000000), bigInt(-1), "alice opens round 0", bigInt(1767229200)})
			}},
		}},
		// Bob bids with a RandomWalk NFT.
		{block: b(2), to: fxGameAddr, logs: []fixtureLog{
			{TopicTransferEvt, func(t *testing.T) *types.Log {
				return buildLog(t, erc20ABI, "Transfer", addr(fxTokenAddr), []any{addr(zero), addr(fxBob)}, []any{eth(100)})
			}},
			{TopicBidEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "BidPlaced", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxBob), bigInt(7)},
					[]any{bigInt(50000000000000000), bigInt(-1), "bob rwalk bid", bigInt(1767230000)})
			}},
		}},
		// Carol bids with CST (V2 event: burn of the price + reward mint).
		{block: b(3), to: fxGameAddr, logs: []fixtureLog{
			{TopicTransferEvt, func(t *testing.T) *types.Log {
				return buildLog(t, erc20ABI, "Transfer", addr(fxTokenAddr), []any{addr(fxCarol), addr(zero)}, []any{eth(5)})
			}},
			{TopicTransferEvt, func(t *testing.T) *types.Log {
				return buildLog(t, erc20ABI, "Transfer", addr(fxTokenAddr), []any{addr(zero), addr(fxCarol)}, []any{eth(110)})
			}},
			{TopicBidEventV2, func(t *testing.T) *types.Log {
				return buildLog(t, gameV2ABI, "BidPlaced", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxCarol), bigInt(-1)},
					[]any{bigInt(-1), eth(5), "carol cst bid", eth(110), bigInt(1800), bigInt(1767230800)})
			}},
		}},
		// Donations: plain ETH, ETH-with-info, ERC20 after a bid, an NFT.
		{block: b(4), to: fxGameAddr, logs: []fixtureLog{
			{TopicEthDonatedEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "EthDonated", addr(fxGameAddr), []any{bigInt(0), addr(fxDave)}, []any{eth(2)})
			}},
		}},
		{block: b(5), to: fxGameAddr, logs: []fixtureLog{
			{TopicEthDonatedWIEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "EthDonatedWithInfo", addr(fxGameAddr), []any{bigInt(0), addr(fxEmma), bigInt(1)}, []any{eth(3)})
			}},
		}},
		{block: b(6), to: fxGameAddr, logs: []fixtureLog{
			{TopicBidEventV2, func(t *testing.T) *types.Log {
				return buildLog(t, gameV2ABI, "BidPlaced", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxAlice), bigInt(-1)},
					[]any{bigInt(110000000000000000), bigInt(-1), "alice bidAndDonateToken", eth(1), bigInt(900), bigInt(1767231600)})
			}},
			{TopicERC20Donated, func(t *testing.T) *types.Log {
				return buildLog(t, prizesWalletABI, "TokenDonated", addr(fxPrizesAddr),
					[]any{bigInt(0), addr(fxAlice), addr(fxMockERC20)}, []any{eth(50)})
			}},
		}},
		{block: b(7), to: fxPrizesAddr, logs: []fixtureLog{
			{TopicNftEthDonatedEvent, func(t *testing.T) *types.Log {
				return buildLog(t, prizesWalletABI, "NftDonated", addr(fxPrizesAddr),
					[]any{bigInt(0), addr(fxBob), addr(fxDonatedNFT)}, []any{bigInt(5), bigInt(0)})
			}},
		}},
		// Marketing reward for emma.
		{block: b(8), to: fxMarketingAddr, logs: []fixtureLog{
			{TopicMarketingRewardPaid, func(t *testing.T) *types.Log {
				return buildLog(t, marketingWalletABI, "RewardPaid", addr(fxMarketingAddr), []any{addr(fxEmma)}, []any{eth(12)})
			}},
		}},
		// Alice claims round 0: the big multi-log transaction.
		{block: b(9), to: fxGameAddr, logs: []fixtureLog{
			{TopicPrizeClaimEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "MainPrizeClaimed", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxAlice), bigInt(1)}, []any{eth(3), eth(10), bigInt(600)})
			}},
			{TopicMintEvent, func(t *testing.T) *types.Log {
				return buildLog(t, signatureABI, "NftMinted", addr(fxSignatureAddr),
					[]any{bigInt(0), addr(fxAlice), bigInt(1)}, []any{bigInt(0xa11ce)})
			}},
			{TopicTransferEvt, func(t *testing.T) *types.Log {
				return buildLog(t, erc721ABI, "Transfer", addr(fxSignatureAddr),
					[]any{addr(zero), addr(fxAlice), bigInt(1)}, nil)
			}},
			{TopicRaffleEthPrizeEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "RaffleWinnerBidderEthPrizeAllocated", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxBob)}, []any{bigInt(0), eth(1)})
			}},
			{TopicEthPrizeDepositEvent, func(t *testing.T) *types.Log {
				return buildLog(t, prizesWalletABI, "EthReceived", addr(fxPrizesAddr),
					[]any{bigInt(0), addr(fxBob)}, []any{bigInt(0), eth(1)})
			}},
			{TopicRaffleNftPrizeEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "RaffleWinnerPrizePaid", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxDave), bigInt(2)}, []any{false, bigInt(1), eth(5)})
			}},
			{TopicMintEvent, func(t *testing.T) *types.Log {
				return buildLog(t, signatureABI, "NftMinted", addr(fxSignatureAddr),
					[]any{bigInt(0), addr(fxDave), bigInt(2)}, []any{bigInt(0xda7e)})
			}},
			{TopicTransferEvt, func(t *testing.T) *types.Log {
				return buildLog(t, erc721ABI, "Transfer", addr(fxSignatureAddr),
					[]any{addr(zero), addr(fxDave), bigInt(2)}, nil)
			}},
			{TopicEndurancePrizeEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "EnduranceChampionPrizePaid", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxBob), bigInt(3)}, []any{eth(7)})
			}},
			{TopicChronoWarriorPrizeEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "ChronoWarriorPrizePaid", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxCarol), bigInt(4)}, []any{bigInt(0), eth(2), eth(8)})
			}},
			{TopicLastcstBidderPrizeEvent, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "LastCstBidderPrizePaid", addr(fxGameAddr),
					[]any{bigInt(0), addr(fxCarol), bigInt(5)}, []any{eth(6)})
			}},
			{TopicDonationReceivedEvent, func(t *testing.T) *types.Log {
				// Emitted by the charity wallet within the claim tx: the
				// handler resolves the round via find_prize_num.
				return buildLog(t, charityWalletABI, "DonationReceived", addr(fxCharityAddr),
					[]any{addr(fxGameAddr)}, []any{eth(1)})
			}},
			{TopicFundsToCharity, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "FundsTransferredToCharity", addr(fxMarketingAddr),
					[]any{addr(fxCharityRcv)}, []any{eth(1)})
			}},
		}},
		// Bob withdraws his raffle ETH.
		{block: b(10), to: fxPrizesAddr, logs: []fixtureLog{
			{TopicEthPrizeWithdrawalEvent, func(t *testing.T) *types.Log {
				return buildLog(t, prizesWalletABI, "EthWithdrawn", addr(fxPrizesAddr),
					[]any{bigInt(0), addr(fxBob), addr(fxBob)}, []any{eth(1)})
			}},
		}},
		// Alice claims the donated prizes.
		{block: b(11), to: fxPrizesAddr, logs: []fixtureLog{
			{TopicDonatedNftClaimed, func(t *testing.T) *types.Log {
				return buildLog(t, prizesWalletABI, "DonatedNftClaimed", addr(fxPrizesAddr),
					[]any{bigInt(0), addr(fxAlice), addr(fxDonatedNFT)}, []any{bigInt(5), bigInt(0)})
			}},
			{TopicDonatedTokenClaimed, func(t *testing.T) *types.Log {
				return buildLog(t, prizesWalletABI, "DonatedTokenClaimed", addr(fxPrizesAddr),
					[]any{bigInt(0), addr(fxAlice), addr(fxMockERC20)}, []any{eth(50)})
			}},
		}},
		// Alice names her NFT and transfers it to bob.
		{block: b(12), to: fxSignatureAddr, logs: []fixtureLog{
			{TopicTokenNameEvent, func(t *testing.T) *types.Log {
				return buildLog(t, signatureABI, "NftNameChanged", addr(fxSignatureAddr),
					[]any{bigInt(1)}, []any{"Round Zero Genesis"})
			}},
			{TopicTransferEvt, func(t *testing.T) *types.Log {
				return buildLog(t, erc721ABI, "Transfer", addr(fxSignatureAddr),
					[]any{addr(fxAlice), addr(fxBob), bigInt(1)}, nil)
			}},
		}},
		// Staking: alice stakes her CS NFT, carol stakes a RandomWalk NFT.
		{block: b(13), to: fxStakingCSTAddr, logs: []fixtureLog{
			{TopicCstNftStakedEvent, func(t *testing.T) *types.Log {
				return buildLog(t, stakingCSTABI, "NftStaked", addr(fxStakingCSTAddr),
					[]any{bigInt(1), bigInt(1), addr(fxAlice)}, []any{bigInt(1), bigInt(0)})
			}},
			{TopicRwalkNftStakedEvent, func(t *testing.T) *types.Log {
				return buildLog(t, stakingRWalkABI, "NftStaked", addr(fxStakingRWKAddr),
					[]any{bigInt(2), bigInt(10), addr(fxCarol)}, []any{bigInt(1)})
			}},
		}},
		// Staking reward deposit, then alice unstakes.
		{block: b(14), to: fxStakingCSTAddr, logs: []fixtureLog{
			{TopicStakingEthDepositEvent, func(t *testing.T) *types.Log {
				return buildLog(t, stakingCSTABI, "EthDepositReceived", addr(fxStakingCSTAddr),
					[]any{bigInt(0)}, []any{bigInt(3), eth(2), eth(2), bigInt(1)})
			}},
		}},
		{block: b(15), to: fxStakingCSTAddr, logs: []fixtureLog{
			{TopicNftUnstakedCst, func(t *testing.T) *types.Log {
				return buildLog(t, stakingCSTABI, "NftUnstaked", addr(fxStakingCSTAddr),
					[]any{bigInt(1), bigInt(1), addr(fxAlice)}, []any{bigInt(4), bigInt(0), eth(2), eth(2)})
			}},
		}},
	}
}

// ingestStory pushes every story transaction through the pipeline.
func ingestStory(t *testing.T, story []storyTx) {
	t.Helper()
	logIndexByBlock := make(map[int64]uint)
	for _, stx := range story {
		logs := make([]*types.Log, 0, len(stx.logs))
		for _, fl := range stx.logs {
			l := fl.build(t)
			if fl.topic0 != "" && len(l.Topics) > 0 {
				if got := l.Topics[0].Hex()[2:]; got != fl.topic0 {
					t.Fatalf("story builds topic0 %s, registry constant is %s", got, fl.topic0)
				}
			}
			logs = append(logs, l)
		}
		start := logIndexByBlock[stx.block]
		ingestTx(t, stx.block, addr(stx.to), start, logs)
		logIndexByBlock[stx.block] = start + uint(len(logs))
	}
}

// TestScriptedRoundGolden pins the cumulative database state of one complete
// round, exercising the insert paths of every major event family plus all the
// aggregate triggers in one realistic interleaving.
func TestScriptedRoundGolden(t *testing.T) {
	resetDB(t)
	base := snapshot(t)

	ingestStory(t, scriptedRound(2000))

	after := snapshot(t)
	diff, err := testutil.DiffSnapshots(base, after)
	if err != nil {
		t.Fatalf("diffing snapshots: %v", err)
	}
	testutil.CompareGolden(t, filepath.Join("testdata", "golden", "story_scripted_round.json"), diff)
}

// TestReorgRollbackAndReplay drives the production chain-split path:
// HandleChainSplit (via EnsureBlockExists hash mismatch) must cascade-delete
// everything from the divergent block up, the trigger delete paths must roll
// the aggregates back (pinned as a golden), and re-processing the replacement
// fork must reproduce the pre-reorg state exactly.
func TestReorgRollbackAndReplay(t *testing.T) {
	resetDB(t)
	const startBlock = 3000
	const divergentBlock = startBlock + 9 // the multi-log claim transaction

	story := scriptedRound(startBlock)
	ingestStory(t, story)
	full := snapshot(t)

	// The competing fork replaces every block from the claim tx onward.
	testChain.Reorg(divergentBlock)
	newHash := testChain.BlockHash(divergentBlock)

	// The polling loop detects the split when a fetched log's block hash
	// disagrees with the stored one.
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
	// Golden view: everything the rollback removed (relative to the full
	// story) — the claim, withdrawals, staking and their trigger effects.
	testutil.CompareGolden(t, filepath.Join("testdata", "golden", "story_reorg_rollback.json"), diff)

	lastBlock, err := dbStore.LastBlockNum(context.Background())
	if err != nil {
		t.Fatalf("Get_last_block_num: %v", err)
	}
	if lastBlock != divergentBlock {
		// HandleChainSplit sets the watermark to divergent-1, then the
		// re-inserted divergent block advances it again.
		t.Errorf("last_block after reorg = %d, want %d", lastBlock, divergentBlock)
	}

	// Re-process the replacement fork: same logical events, new fork.
	var replay []storyTx
	for _, stx := range story {
		if stx.block >= divergentBlock {
			replay = append(replay, stx)
		}
	}
	ingestStory(t, replay)

	requireNoDiff(t, full, snapshot(t), "state after reorg replay")
}
