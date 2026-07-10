//go:build integration

// Per-event golden fixtures (§4.3): every event type dispatched by
// select_event_and_process gets at least one synthetic log driven through the
// real pipeline; the resulting database mutation (including all plpgsql
// trigger side effects) is pinned as a golden diff, and re-processing the
// same events must leave the state unchanged (the delete-then-insert pattern
// of the CG handlers).
//
// Fixture block numbers are explicit and unique per fixture (spaced by 10, so
// multi-block fixtures can use +1 offsets). New fixtures take fresh numbers
// at the end; renumbering existing fixtures churns their goldens for nothing.
package cosmicgame

import (
	"context"
	"math/big"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// fixtureLog is one log to ingest: build constructs it, topic0 (a registry
// constant from events_registry.go) cross-checks that the ABI-derived event
// ID matches the constant the dispatcher will compare against. Empty topic0
// skips the check (negative fixtures).
type fixtureLog struct {
	topic0 string
	build  func(t *testing.T) *types.Log
}

// fixtureTx groups logs emitted by one transaction.
type fixtureTx struct {
	blockOffset int64
	to          string // contract the tx calls (transaction.to_aid)
	logs        []fixtureLog
}

type fixture struct {
	name  string
	block int64
	txs   []fixtureTx
	// skipReplay documents why re-processing this fixture's events cannot be
	// asserted state-neutral (empty means the replay assertion runs).
	skipReplay string
}

func TestEventFixtures(t *testing.T) {
	for _, fx := range eventFixtures() {
		t.Run(fx.name, func(t *testing.T) {
			resetDB(t)
			base := snapshot(t)

			var evtIDs []int64
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
				ids := ingestTx(t, block, addr(ftx.to), start, logs)
				logIndexByBlock[block] = start + uint(len(logs))
				evtIDs = append(evtIDs, ids...)
			}

			after := snapshot(t)
			diff, err := testutil.DiffSnapshots(base, after)
			if err != nil {
				t.Fatalf("diffing snapshots: %v", err)
			}
			testutil.CompareGolden(t, filepath.Join("testdata", "golden", fx.name+".json"), diff)

			if fx.skipReplay != "" {
				t.Logf("replay assertion skipped: %s", fx.skipReplay)
				return
			}
			// Replay: the CG handlers delete-then-insert by evtlog id, so
			// re-processing (the reorg/recovery path) must be state-neutral.
			for _, id := range evtIDs {
				if err := testProcess(context.Background(), id); err != nil {
					t.Fatalf("replaying event %d: %v", id, err)
				}
			}
			requireNoDiff(t, after, snapshot(t), "replay of "+fx.name)
		})
	}
}

// eventFixtures returns one fixture per dispatched event type plus negative
// cases. Values are arbitrary but distinctive, so a golden mismatch points at
// the field that moved.
func eventFixtures() []fixture {
	game := fxGameAddr
	sig := fxSignatureAddr
	token := fxTokenAddr
	charity := fxCharityAddr
	prizes := fxPrizesAddr
	stakeCST := fxStakingCSTAddr
	stakeRWK := fxStakingRWKAddr
	marketing := fxMarketingAddr

	return []fixture{
		// --- Admin: single-uint256 parameter changes on the game contract ---
		{name: "admin_charity_percentage", block: 1000, txs: []fixtureTx{{to: game, logs: []fixtureLog{{CHARITY_PERCENTAGE_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "CharityEthDonationAmountPercentageChanged", addr(game), nil, []any{bigInt(11)})
		}}}}}},
		{name: "admin_prize_percentage", block: 1010, txs: []fixtureTx{{to: game, logs: []fixtureLog{{PRIZE_PERCENTAGE_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "MainEthPrizeAmountPercentageChanged", addr(game), nil, []any{bigInt(26)})
		}}}}}},
		{name: "admin_raffle_percentage", block: 1020, txs: []fixtureTx{{to: game, logs: []fixtureLog{{RAFFLE_PERCENTAGE_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", addr(game), nil, []any{bigInt(6)})
		}}}}}},
		{name: "admin_staking_percentage", block: 1030, txs: []fixtureTx{{to: game, logs: []fixtureLog{{STAKE_PERCENTAGE_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", addr(game), nil, []any{bigInt(9)})
		}}}}}},
		{name: "admin_chrono_percentage", block: 1040, txs: []fixtureTx{{to: game, logs: []fixtureLog{{CHRONO_PERCENTAGE_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "ChronoWarriorEthPrizeAmountPercentageChanged", addr(game), nil, []any{bigInt(8)})
		}}}}}},
		{name: "admin_num_raffle_eth_winners", block: 1050, txs: []fixtureTx{{to: game, logs: []fixtureLog{{NUM_RAFFLE_ETH_PRIZE_EVENTS_BIDDING_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "NumRaffleEthPrizesForBiddersChanged", addr(game), nil, []any{bigInt(3)})
		}}}}}},
		{name: "admin_num_raffle_nft_winners", block: 1060, txs: []fixtureTx{{to: game, logs: []fixtureLog{{NUM_RAFFLE_NFT_PRIZE_EVENTS_BIDDING_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "NumRaffleCosmicSignatureNftsForBiddersChanged", addr(game), nil, []any{bigInt(5)})
		}}}}}},
		{name: "admin_num_raffle_nft_stakers", block: 1070, txs: []fixtureTx{{to: game, logs: []fixtureLog{{NUM_RAFFLE_NFT_PRIZE_EVENTS_STAKING_RWALK_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", addr(game), nil, []any{bigInt(4)})
		}}}}}},
		{name: "admin_timeout_claim_prize", block: 1080, txs: []fixtureTx{{to: game, logs: []fixtureLog{{TIMEOUT_CLAIMPRIZE_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "TimeoutDurationToClaimMainPrizeChanged", addr(game), nil, []any{bigInt(86400)})
		}}}}}},
		{name: "admin_timeout_withdraw_prizes", block: 1090, txs: []fixtureTx{{to: prizes, logs: []fixtureLog{{TIMEOUT_TO_WITHDRAW_PRIZE, func(t *testing.T) *types.Log {
			return buildLog(t, prizesWalletABI, "TimeoutDurationToWithdrawPrizesChanged", addr(prizes), nil, []any{bigInt(600)})
		}}}}}},
		{name: "admin_price_increase_divisor", block: 1100, txs: []fixtureTx{{to: game, logs: []fixtureLog{{PRICE_INCREASE_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "EthBidPriceIncreaseDivisorChanged", addr(game), nil, []any{bigInt(100)})
		}}}}}},
		{name: "admin_mainprize_time_increment", block: 1110, txs: []fixtureTx{{to: game, logs: []fixtureLog{{MAIN_PRIZE_MICROSECOND_INCREASE, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "MainPrizeTimeIncrementInMicroSecondsChanged", addr(game), nil, []any{bigInt(3600000000)})
		}}}}}},
		{name: "admin_initial_duration_divisor", block: 1120, txs: []fixtureTx{{to: game, logs: []fixtureLog{{INITIAL_SECONDS_UNTIL_PRIZE_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "InitialDurationUntilMainPrizeDivisorChanged", addr(game), nil, []any{bigInt(2)})
		}}}}}},
		{name: "admin_round_activation_time", block: 1130, txs: []fixtureTx{{to: game, logs: []fixtureLog{{ROUND_ACTIVATION_TIME_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "RoundActivationTimeChanged", addr(game), nil, []any{bigInt(1767225600)})
		}}}}}},
		{name: "admin_cst_auction_duration_divisor", block: 1140, txs: []fixtureTx{{to: game, logs: []fixtureLog{{CST_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "CstDutchAuctionDurationDivisorChanged", addr(game), nil, []any{bigInt(11)})
		}}}}}},
		{name: "admin_cst_auction_duration_v2", block: 1150, txs: []fixtureTx{{to: game, logs: []fixtureLog{{CST_DUTCH_AUCTION_DURATION_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameV2ABI, "CstDutchAuctionDurationChanged", addr(game), nil, []any{bigInt(7200)})
		}}}}}},
		{name: "admin_cst_auction_change_divisor_v2", block: 1160, txs: []fixtureTx{{to: game, logs: []fixtureLog{{CST_DUTCH_AUCTION_DURATION_CHANGE_DIVISOR_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameV2ABI, "CstDutchAuctionDurationChangeDivisorChanged", addr(game), nil, []any{bigInt(3)})
		}}}}}},
		{name: "admin_eth_auction_duration_divisor", block: 1170, txs: []fixtureTx{{to: game, logs: []fixtureLog{{ETH_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "EthDutchAuctionDurationDivisorChanged", addr(game), nil, []any{bigInt(20)})
		}}}}}},
		{name: "admin_eth_auction_ending_price_divisor", block: 1180, txs: []fixtureTx{{to: game, logs: []fixtureLog{{ETH_DUTCH_AUCTION_ENDING_PRICE_DIVISOR_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "EthDutchAuctionEndingBidPriceDivisorChanged", addr(game), nil, []any{bigInt(10)})
		}}}}}},
		{name: "admin_cst_reward_for_bidding", block: 1190, txs: []fixtureTx{{to: game, logs: []fixtureLog{{CST_REWARD_FOR_BIDDING_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "CstRewardAmountForBiddingChanged", addr(game), nil, []any{eth(150)})
		}}}}}},
		{name: "admin_bid_cst_reward_raw", block: 1200, txs: []fixtureTx{{to: game, logs: []fixtureLog{{BID_CST_REWARD_AMOUNT_CHANGED, func(t *testing.T) *types.Log {
			return buildRawLog(t, BID_CST_REWARD_AMOUNT_CHANGED, addr(game), nil, eth(99))
		}}}}}},
		{name: "admin_bid_cst_reward_multiplier_v2", block: 1210, txs: []fixtureTx{{to: game, logs: []fixtureLog{{BID_CST_REWARD_AMOUNT_MULTIPLIER_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameV2ABI, "BidCstRewardAmountMultiplierChanged", addr(game), nil, []any{bigInt(3)})
		}}}}}},
		{name: "admin_static_cst_reward", block: 1220, txs: []fixtureTx{{to: game, logs: []fixtureLog{{STATIC_CST_REWARD, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "CstPrizeAmountChanged", addr(game), nil, []any{eth(10)})
		}}}}}},
		{name: "admin_max_message_length", block: 1230, txs: []fixtureTx{{to: game, logs: []fixtureLog{{MAX_MESSAGE_LENGTH, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "BidMessageLengthMaxLimitChanged", addr(game), nil, []any{bigInt(280)})
		}}}}}},
		{name: "admin_marketing_reward_changed", block: 1240, txs: []fixtureTx{{to: game, logs: []fixtureLog{{MARKETING_REWARD_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "MarketingWalletCstContributionAmountChanged", addr(game), nil, []any{eth(300)})
		}}}}}},
		{name: "admin_delay_duration_round", block: 1250, txs: []fixtureTx{{to: game, logs: []fixtureLog{{DELAY_DURATION_ROUND, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "DelayDurationBeforeRoundActivationChanged", addr(game), nil, []any{bigInt(1800)})
		}}}}}},
		{name: "admin_cst_min_limit", block: 1260, txs: []fixtureTx{{to: game, logs: []fixtureLog{{STARTING_CST_MIN_LIM, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "CstDutchAuctionBeginningBidPriceMinLimitChanged", addr(game), nil, []any{eth(200)})
		}}}}}},
		{name: "admin_time_increase_raw", block: 1270, txs: []fixtureTx{{to: game, logs: []fixtureLog{{TIME_INCREASE_CHANGED, func(t *testing.T) *types.Log {
			return buildRawLog(t, TIME_INCREASE_CHANGED, addr(game), nil, bigInt(3600000001))
		}}}}}},
		{name: "admin_token_script_url", block: 1280, txs: []fixtureTx{{to: sig, logs: []fixtureLog{{TOKEN_SCRIPT_URL, func(t *testing.T) *types.Log {
			return buildLog(t, signatureABI, "NftGenerationScriptUriChanged", addr(sig), nil, []any{"https://fixture.example/script.js"})
		}}}}}},
		{name: "admin_base_uri", block: 1290, txs: []fixtureTx{{to: sig, logs: []fixtureLog{{BASE_URI, func(t *testing.T) *types.Log {
			return buildLog(t, signatureABI, "NftBaseUriChanged", addr(sig), nil, []any{"https://fixture.example/base/"})
		}}}}}},

		// --- Admin: address changes (indexed address, empty data) ---
		{name: "admin_charity_wallet_changed_game", block: 1300, txs: []fixtureTx{{to: game, logs: []fixtureLog{{CHARITY_WALLET_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "CharityAddressChanged", addr(game), []any{addr(fxCharityAddr)}, nil)
		}}}}}},
		{name: "admin_charity_receiver_changed_wallet", block: 1310, txs: []fixtureTx{{to: charity, logs: []fixtureLog{{CHARITY_RECEIVER_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, charityWalletABI, "CharityAddressChanged", addr(charity), []any{addr(fxCharityRcv)}, nil)
		}}}}}},
		{name: "admin_random_walk_address", block: 1320, txs: []fixtureTx{{to: game, logs: []fixtureLog{{RWALK_ADDRESS_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "RandomWalkNftAddressChanged", addr(game), []any{addr(fxRandomWalkAddr)}, nil)
		}}}}}},
		{name: "admin_prizes_wallet_address", block: 1330, txs: []fixtureTx{{to: game, logs: []fixtureLog{{PRIZE_WALLET_ADDRESS_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "PrizesWalletAddressChanged", addr(game), []any{addr(fxPrizesAddr)}, nil)
		}}}}}},
		{name: "admin_staking_cst_address", block: 1340, txs: []fixtureTx{{to: game, logs: []fixtureLog{{STAKING_WALLET_CST_ADDRESS_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "StakingWalletCosmicSignatureNftAddressChanged", addr(game), []any{addr(fxStakingCSTAddr)}, nil)
		}}}}}},
		{name: "admin_staking_rwalk_address", block: 1350, txs: []fixtureTx{{to: game, logs: []fixtureLog{{STAKING_WALLET_RWALK_ADDRESS_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "StakingWalletRandomWalkNftAddressChanged", addr(game), []any{addr(fxStakingRWKAddr)}, nil)
		}}}}}},
		{name: "admin_marketing_address", block: 1360, txs: []fixtureTx{{to: game, logs: []fixtureLog{{MARKETING_ADDRESS_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "MarketingWalletAddressChanged", addr(game), []any{addr(fxMarketingAddr)}, nil)
		}}}}}},
		{name: "admin_cosmic_token_address", block: 1370, txs: []fixtureTx{{to: game, logs: []fixtureLog{{COSMIC_TOKEN_ADDRESS_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "CosmicSignatureTokenAddressChanged", addr(game), []any{addr(fxTokenAddr)}, nil)
		}}}}}},
		{name: "admin_cosmic_signature_address", block: 1380, txs: []fixtureTx{{to: game, logs: []fixtureLog{{COSMIC_SIGNATURE_ADDRESS_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "CosmicSignatureNftAddressChanged", addr(game), []any{addr(fxSignatureAddr)}, nil)
		}}}}}},
		{name: "admin_treasurer_changed", block: 1390, txs: []fixtureTx{{to: marketing, logs: []fixtureLog{{TREASURER_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, marketingWalletABI, "TreasurerAddressChanged", addr(marketing), []any{addr(fxDave)}, nil)
		}}}}}},
		{name: "admin_proxy_upgraded", block: 1400, txs: []fixtureTx{{to: game, logs: []fixtureLog{{PROXY_UPGRADED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "Upgraded", addr(game), []any{addr(fxImplAddr)}, nil)
		}}}}}},
		{name: "admin_admin_changed", block: 1410, txs: []fixtureTx{{to: game, logs: []fixtureLog{{ADMIN_CHANGED, func(t *testing.T) *types.Log {
			return buildLog(t, erc1967ABI, "AdminChanged", addr(game), nil, []any{addr(fxAlice), addr(fxBob)})
		}}}}}},
		{name: "admin_ownership_transferred", block: 1420, txs: []fixtureTx{{to: game, logs: []fixtureLog{{OWNERSHIP_TRANSFERRED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "OwnershipTransferred", addr(game), []any{addr(fxAlice), addr(fxBob)}, nil)
		}}}}}},
		{name: "admin_initialized", block: 1430, txs: []fixtureTx{{to: game, logs: []fixtureLog{{INITIALIZED, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "Initialized", addr(game), nil, []any{uint64(1)})
		}}}}}},
		{name: "admin_fund_transfer_failed", block: 1440, txs: []fixtureTx{{to: game, logs: []fixtureLog{{FUND_TRANSFER_ERR, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "FundTransferFailed", addr(game), []any{addr(fxCarol)}, []any{"fixture transfer failure", eth(1)})
		}}}}}},

		// --- Bidding ---
		{name: "bid_eth_v1", block: 1450, txs: []fixtureTx{{to: game, logs: []fixtureLog{
			{TRANSFER_EVT, func(t *testing.T) *types.Log {
				return buildLog(t, erc20ABI, "Transfer", addr(token), []any{addr("0x0000000000000000000000000000000000000000"), addr(fxAlice)}, []any{eth(100)})
			}},
			{BID_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "BidPlaced", addr(game),
					[]any{bigInt(0), addr(fxAlice), bigInt(-1)},
					[]any{bigInt(100000000000000000), bigInt(-1), "eth bid fixture", bigInt(1767226700)})
			}},
		}}}},
		{name: "bid_randomwalk_v1", block: 1460, txs: []fixtureTx{{to: game, logs: []fixtureLog{
			{TRANSFER_EVT, func(t *testing.T) *types.Log {
				return buildLog(t, erc20ABI, "Transfer", addr(token), []any{addr("0x0000000000000000000000000000000000000000"), addr(fxBob)}, []any{eth(100)})
			}},
			{BID_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "BidPlaced", addr(game),
					[]any{bigInt(0), addr(fxBob), bigInt(7)},
					[]any{bigInt(50000000000000000), bigInt(-1), "randomwalk bid fixture", bigInt(1767226800)})
			}},
		}}}},
		{name: "bid_cst_v1", block: 1470, txs: []fixtureTx{{to: game, logs: []fixtureLog{
			{TRANSFER_EVT, func(t *testing.T) *types.Log { // burn of the paid CST price (ignored by the reward matcher)
				return buildLog(t, erc20ABI, "Transfer", addr(token), []any{addr(fxCarol), addr("0x0000000000000000000000000000000000000000")}, []any{eth(5)})
			}},
			{TRANSFER_EVT, func(t *testing.T) *types.Log { // reward mint to the bidder
				return buildLog(t, erc20ABI, "Transfer", addr(token), []any{addr("0x0000000000000000000000000000000000000000"), addr(fxCarol)}, []any{eth(100)})
			}},
			{BID_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "BidPlaced", addr(game),
					[]any{bigInt(0), addr(fxCarol), bigInt(-1)},
					[]any{bigInt(-1), eth(5), "cst bid fixture", bigInt(1767226900)})
			}},
		}}}},
		{name: "bid_eth_v2", block: 1480, txs: []fixtureTx{{to: game, logs: []fixtureLog{
			{BID_EVENT_V2, func(t *testing.T) *types.Log {
				// No ERC20 mint in the tx: the dynamic V2 reward is zero and
				// find_cosmic_token_transfer must report "0".
				return buildLog(t, gameV2ABI, "BidPlaced", addr(game),
					[]any{bigInt(0), addr(fxAlice), bigInt(-1)},
					[]any{bigInt(200000000000000000), bigInt(-1), "v2 bid fixture", eth(12), bigInt(3600), bigInt(1767227000)})
			}},
		}}}},
		{name: "bid_first_in_round", block: 1490, txs: []fixtureTx{{to: game, logs: []fixtureLog{
			{FIRST_BID_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "FirstBidPlacedInRound", addr(game), []any{bigInt(0)}, []any{bigInt(1767225700)})
			}},
			{BID_EVENT_V2, func(t *testing.T) *types.Log {
				return buildLog(t, gameV2ABI, "BidPlaced", addr(game),
					[]any{bigInt(0), addr(fxAlice), bigInt(-1)},
					[]any{bigInt(100000000000000000), bigInt(-1), "round starter", eth(1), bigInt(1800), bigInt(1767227100)})
			}},
		}}}},
		{name: "bid_wrong_address_skipped", block: 1500, txs: []fixtureTx{{to: game, logs: []fixtureLog{
			{BID_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, gameABI, "BidPlaced", addr(fxAlice),
					[]any{bigInt(0), addr(fxAlice), bigInt(-1)},
					[]any{bigInt(1), bigInt(-1), "wrong contract", bigInt(1767227200)})
			}},
		}}}},

		// --- Donations ---
		{name: "donation_eth", block: 1510, txs: []fixtureTx{{to: game, logs: []fixtureLog{{ETH_DONATED_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "EthDonated", addr(game), []any{bigInt(0), addr(fxDave)}, []any{eth(2)})
		}}}}}},
		{name: "donation_eth_with_info", block: 1520, txs: []fixtureTx{{to: game, logs: []fixtureLog{{ETH_DONATED_WI_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "EthDonatedWithInfo", addr(game), []any{bigInt(0), addr(fxEmma), bigInt(1)}, []any{eth(3)})
		}}}}}},
		{name: "donation_received_standalone", block: 1530, txs: []fixtureTx{{to: charity, logs: []fixtureLog{{DONATION_RECEIVED_EVENT, func(t *testing.T) *types.Log {
			// No MainPrizeClaimed in the tx: find_prize_num pins round -1.
			return buildLog(t, charityWalletABI, "DonationReceived", addr(charity), []any{addr(fxDave)}, []any{eth(2)})
		}}}}}},
		{name: "donation_sent_charity", block: 1540, txs: []fixtureTx{{to: charity, logs: []fixtureLog{{DONATION_SENT_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, charityWalletABI, "FundsTransferredToCharity", addr(charity), []any{addr(fxCharityRcv)}, []any{eth(4)})
		}}}}}},
		{name: "funds_to_charity_marketing", block: 1550, txs: []fixtureTx{{to: marketing, logs: []fixtureLog{{FUNDS_TO_CHARITY, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "FundsTransferredToCharity", addr(marketing), []any{addr(fxCharityRcv)}, []any{eth(1)})
		}}}}}},
		{name: "donation_erc20_after_bid", block: 1560, txs: []fixtureTx{{to: game, logs: []fixtureLog{
			{BID_EVENT_V2, func(t *testing.T) *types.Log {
				return buildLog(t, gameV2ABI, "BidPlaced", addr(game),
					[]any{bigInt(0), addr(fxAlice), bigInt(-1)},
					[]any{bigInt(100000000000000000), bigInt(-1), "bidAndDonateToken", eth(1), bigInt(60), bigInt(1767227300)})
			}},
			{ERC20_DONATED, func(t *testing.T) *types.Log {
				return buildLog(t, prizesWalletABI, "TokenDonated", addr(prizes),
					[]any{bigInt(0), addr(fxAlice), addr(fxMockERC20)}, []any{eth(50)})
			}},
		}}}},
		{name: "donation_erc20_without_bid", block: 1570, txs: []fixtureTx{{to: prizes, logs: []fixtureLog{{ERC20_DONATED, func(t *testing.T) *types.Log {
			return buildLog(t, prizesWalletABI, "TokenDonated", addr(prizes),
				[]any{bigInt(0), addr(fxBob), addr(fxMockERC20)}, []any{eth(25)})
		}}}}}},
		{name: "donation_nft", block: 1580, txs: []fixtureTx{{to: prizes, logs: []fixtureLog{{NFT_ETH_DONATED_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, prizesWalletABI, "NftDonated", addr(prizes),
				[]any{bigInt(0), addr(fxBob), addr(fxDonatedNFT)}, []any{bigInt(5), bigInt(0)})
		}}}}}},
		{name: "donated_token_claimed", block: 1590, txs: []fixtureTx{{to: prizes, logs: []fixtureLog{{DONATED_TOKEN_CLAIMED, func(t *testing.T) *types.Log {
			return buildLog(t, prizesWalletABI, "DonatedTokenClaimed", addr(prizes),
				[]any{bigInt(0), addr(fxAlice), addr(fxMockERC20)}, []any{eth(50)})
		}}}}}},
		{name: "donated_nft_claimed", block: 1600, txs: []fixtureTx{{to: prizes, logs: []fixtureLog{{DONATED_NFT_CLAIMED, func(t *testing.T) *types.Log {
			return buildLog(t, prizesWalletABI, "DonatedNftClaimed", addr(prizes),
				[]any{bigInt(0), addr(fxAlice), addr(fxDonatedNFT)}, []any{bigInt(5), bigInt(0)})
		}}}}}},

		// --- Prizes ---
		{name: "prize_claim_main", block: 1610, txs: []fixtureTx{
			{blockOffset: 0, to: game, logs: []fixtureLog{
				{FIRST_BID_EVENT, func(t *testing.T) *types.Log {
					return buildLog(t, gameABI, "FirstBidPlacedInRound", addr(game), []any{bigInt(0)}, []any{bigInt(1767225800)})
				}},
				{BID_EVENT_V2, func(t *testing.T) *types.Log {
					return buildLog(t, gameV2ABI, "BidPlaced", addr(game),
						[]any{bigInt(0), addr(fxAlice), bigInt(-1)},
						[]any{bigInt(100000000000000000), bigInt(-1), "round 0 bid", eth(1), bigInt(60), bigInt(1767227400)})
				}},
			}},
			{blockOffset: 1, to: game, logs: []fixtureLog{
				{PRIZE_CLAIM_EVENT, func(t *testing.T) *types.Log {
					return buildLog(t, gameABI, "MainPrizeClaimed", addr(game),
						[]any{bigInt(0), addr(fxAlice), bigInt(1)},
						[]any{eth(3), eth(10), bigInt(600)})
				}},
			}},
		}},
		{name: "prizes_eth_deposit", block: 1620, txs: []fixtureTx{{to: prizes, logs: []fixtureLog{{ETH_PRIZE_DEPOSIT_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, prizesWalletABI, "EthReceived", addr(prizes),
				[]any{bigInt(0), addr(fxBob)}, []any{bigInt(2), eth(1)})
		}}}}}},
		{name: "prizes_eth_withdrawal", block: 1630, txs: []fixtureTx{
			{blockOffset: 0, to: prizes, logs: []fixtureLog{{ETH_PRIZE_DEPOSIT_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, prizesWalletABI, "EthReceived", addr(prizes),
					[]any{bigInt(0), addr(fxBob)}, []any{bigInt(2), eth(1)})
			}}}},
			{blockOffset: 1, to: prizes, logs: []fixtureLog{{ETH_PRIZE_WITHDRAWAL_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, prizesWalletABI, "EthWithdrawn", addr(prizes),
					[]any{bigInt(0), addr(fxBob), addr(fxBob)}, []any{eth(1)})
			}}}},
		}},
		{name: "raffle_eth_winner", block: 1640, txs: []fixtureTx{{to: game, logs: []fixtureLog{{RAFFLE_ETH_PRIZE_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "RaffleWinnerBidderEthPrizeAllocated", addr(game),
				[]any{bigInt(0), addr(fxBob)}, []any{bigInt(0), eth(1)})
		}}}}}},
		{name: "raffle_nft_winner_bidder", block: 1650, txs: []fixtureTx{{to: game, logs: []fixtureLog{{RAFFLE_NFT_PRIZE_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "RaffleWinnerPrizePaid", addr(game),
				[]any{bigInt(0), addr(fxDave), bigInt(2)}, []any{false, bigInt(1), eth(5)})
		}}}}}},
		{name: "raffle_nft_winner_staker", block: 1660, txs: []fixtureTx{{to: game, logs: []fixtureLog{{RAFFLE_NFT_PRIZE_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "RaffleWinnerPrizePaid", addr(game),
				[]any{bigInt(0), addr(fxCarol), bigInt(3)}, []any{true, bigInt(0), eth(5)})
		}}}}}},
		{name: "endurance_champion_prize", block: 1670, txs: []fixtureTx{{to: game, logs: []fixtureLog{{ENDURANCE_PRIZE_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "EnduranceChampionPrizePaid", addr(game),
				[]any{bigInt(0), addr(fxBob), bigInt(4)}, []any{eth(7)})
		}}}}}},
		{name: "last_cst_bidder_prize", block: 1680, txs: []fixtureTx{{to: game, logs: []fixtureLog{{LASTCST_BIDDER_PRIZE_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "LastCstBidderPrizePaid", addr(game),
				[]any{bigInt(0), addr(fxCarol), bigInt(5)}, []any{eth(6)})
		}}}}}},
		{name: "chrono_warrior_prize", block: 1690, txs: []fixtureTx{{to: game, logs: []fixtureLog{{CHRONO_WARRIOR_PRIZE_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, gameABI, "ChronoWarriorPrizePaid", addr(game),
				[]any{bigInt(0), addr(fxAlice), bigInt(6)}, []any{bigInt(0), eth(2), eth(8)})
		}}}}}},

		// --- Tokens ---
		{name: "token_mint", block: 1700, txs: []fixtureTx{{to: sig, logs: []fixtureLog{{MINT_EVENT, func(t *testing.T) *types.Log {
			// Small seed value exercises the %064x zero-padding of NftSeed.
			return buildLog(t, signatureABI, "NftMinted", addr(sig),
				[]any{bigInt(0), addr(fxAlice), bigInt(1)}, []any{bigInt(0xc0ffee)})
		}}}}}},
		{name: "token_name_after_mint", block: 1710, txs: []fixtureTx{
			{blockOffset: 0, to: sig, logs: []fixtureLog{{MINT_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, signatureABI, "NftMinted", addr(sig),
					[]any{bigInt(0), addr(fxAlice), bigInt(1)}, []any{bigInt(0xbeef)})
			}}}},
			{blockOffset: 1, to: sig, logs: []fixtureLog{{TOKEN_NAME_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, signatureABI, "NftNameChanged", addr(sig),
					[]any{bigInt(1)}, []any{"Genesis Fixture"})
			}}}},
		}},
		{name: "token_transfer_erc721", block: 1720, txs: []fixtureTx{{to: sig, logs: []fixtureLog{{TRANSFER_EVT, func(t *testing.T) *types.Log {
			return buildLog(t, erc721ABI, "Transfer", addr(sig),
				[]any{addr(fxAlice), addr(fxBob), bigInt(1)}, nil)
		}}}}}},
		{name: "token_transfer_erc20", block: 1730, txs: []fixtureTx{{to: token, logs: []fixtureLog{{TRANSFER_EVT, func(t *testing.T) *types.Log {
			return buildLog(t, erc20ABI, "Transfer", addr(token),
				[]any{addr(fxAlice), addr(fxBob)}, []any{eth(25)})
		}}}}}},
		{name: "marketing_reward_paid", block: 1740, txs: []fixtureTx{{to: marketing, logs: []fixtureLog{{MARKETING_REWARD_PAID, func(t *testing.T) *types.Log {
			return buildLog(t, marketingWalletABI, "RewardPaid", addr(marketing),
				[]any{addr(fxEmma)}, []any{eth(12)})
		}}}}}},

		// --- Staking ---
		{name: "staking_cst_stake", block: 1750, txs: []fixtureTx{{to: stakeCST, logs: []fixtureLog{{CST_NFT_STAKED_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, stakingCSTABI, "NftStaked", addr(stakeCST),
				[]any{bigInt(1), bigInt(1), addr(fxAlice)}, []any{bigInt(1), bigInt(0)})
		}}}}}},
		{name: "staking_cst_stake_unstake", block: 1760, txs: []fixtureTx{
			{blockOffset: 0, to: stakeCST, logs: []fixtureLog{{CST_NFT_STAKED_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, stakingCSTABI, "NftStaked", addr(stakeCST),
					[]any{bigInt(1), bigInt(1), addr(fxAlice)}, []any{bigInt(1), bigInt(0)})
			}}}},
			{blockOffset: 1, to: stakeCST, logs: []fixtureLog{{NFT_UNSTAKED_CST, func(t *testing.T) *types.Log {
				return buildLog(t, stakingCSTABI, "NftUnstaked", addr(stakeCST),
					[]any{bigInt(1), bigInt(1), addr(fxAlice)}, []any{bigInt(2), bigInt(0), bigInt(0), bigInt(0)})
			}}}},
		}},
		{name: "staking_rwalk_stake", block: 1770, txs: []fixtureTx{{to: stakeRWK, logs: []fixtureLog{{RWALK_NFT_STAKED_EVENT, func(t *testing.T) *types.Log {
			return buildLog(t, stakingRWalkABI, "NftStaked", addr(stakeRWK),
				[]any{bigInt(5), bigInt(10), addr(fxCarol)}, []any{bigInt(1)})
		}}}}}},
		{name: "staking_rwalk_stake_unstake", block: 1780, txs: []fixtureTx{
			{blockOffset: 0, to: stakeRWK, logs: []fixtureLog{{RWALK_NFT_STAKED_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, stakingRWalkABI, "NftStaked", addr(stakeRWK),
					[]any{bigInt(5), bigInt(10), addr(fxCarol)}, []any{bigInt(1)})
			}}}},
			{blockOffset: 1, to: stakeRWK, logs: []fixtureLog{{NFT_UNSTAKED_RWALK, func(t *testing.T) *types.Log {
				return buildLog(t, stakingRWalkABI, "NftUnstaked", addr(stakeRWK),
					[]any{bigInt(5), bigInt(10), addr(fxCarol)}, []any{bigInt(6), bigInt(0)})
			}}}},
		}},
		{name: "staking_eth_deposit", block: 1790, txs: []fixtureTx{
			{blockOffset: 0, to: stakeCST, logs: []fixtureLog{{CST_NFT_STAKED_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, stakingCSTABI, "NftStaked", addr(stakeCST),
					[]any{bigInt(1), bigInt(1), addr(fxAlice)}, []any{bigInt(1), bigInt(0)})
			}}}},
			{blockOffset: 1, to: stakeCST, logs: []fixtureLog{{STAKING_ETH_DEPOSIT_EVENT, func(t *testing.T) *types.Log {
				return buildLog(t, stakingCSTABI, "EthDepositReceived", addr(stakeCST),
					[]any{bigInt(0)}, []any{bigInt(3), eth(2), eth(2), bigInt(1)})
			}}}},
		}},

		// --- Negative / dispatch edge cases ---
		{name: "unknown_topic_noop", block: 1800, txs: []fixtureTx{{to: game, logs: []fixtureLog{{"", func(t *testing.T) *types.Log {
			return &types.Log{
				Address: addr(game),
				Topics:  []ethcommon.Hash{ethcommon.HexToHash("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")},
				Data:    []byte{0x01},
			}
		}}}}}},
		{name: "no_topics_noop", block: 1810, txs: []fixtureTx{{to: game, logs: []fixtureLog{{"", func(t *testing.T) *types.Log {
			return &types.Log{Address: addr(game), Topics: nil, Data: []byte{0x02}}
		}}}}}},

		// ERC20TransferFailed lives in ICosmicSignatureErrors.sol and is
		// absent from every generated ABI, so the log is hand-packed
		// (string errStr + uint256 amount tail, destination indexed).
		{name: "admin_erc20_transfer_failed", block: 1820, txs: []fixtureTx{{to: game, logs: []fixtureLog{{ERC20_TRANSFER_ERR, func(t *testing.T) *types.Log {
			return buildERC20TransferFailedLog(t, addr(game), addr(fxCarol), "fixture erc20 transfer failure", eth(2))
		}}}}}},
	}
}

// buildERC20TransferFailedLog packs an ERC20TransferFailed(string,address
// indexed,uint256) log without an ABI: the data body carries the canonical
// encoding of (errStr, amount), the destination rides in topic 1.
func buildERC20TransferFailedLog(t *testing.T, contract, destination ethcommon.Address, errStr string, amount *big.Int) *types.Log {
	t.Helper()
	stringT, err := abi.NewType("string", "", nil)
	if err != nil {
		t.Fatalf("string abi type: %v", err)
	}
	uint256T, err := abi.NewType("uint256", "", nil)
	if err != nil {
		t.Fatalf("uint256 abi type: %v", err)
	}
	data, err := abi.Arguments{{Type: stringT}, {Type: uint256T}}.Pack(errStr, amount)
	if err != nil {
		t.Fatalf("packing ERC20TransferFailed data: %v", err)
	}
	return &types.Log{
		Address: contract,
		Topics: []ethcommon.Hash{
			ethcommon.HexToHash("0x" + ERC20_TRANSFER_ERR),
			ethcommon.BytesToHash(destination.Bytes()),
		},
		Data: data,
	}
}
