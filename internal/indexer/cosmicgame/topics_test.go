// Unit tests (no Docker) pinning the hand-maintained topic-hash registry in
// events_registry.go against the ABI-derived event IDs: if a contract ABI is
// regenerated and an event signature changes, the constant here goes stale
// and the ETL silently stops dispatching that event. This test turns silent
// drift into a build-time failure.
package cosmicgame

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func mustABI(t *testing.T, raw string) *abi.ABI {
	t.Helper()
	parsed, err := abi.JSON(strings.NewReader(raw))
	if err != nil {
		t.Fatalf("parsing ABI: %v", err)
	}
	return &parsed
}

func TestRegistryConstantsMatchABIEventIDs(t *testing.T) {
	game := mustABI(t, cgc.CosmicSignatureGameABI)
	gameV2 := mustABI(t, cgc.CosmicSignatureGameV2ABI)
	nft := mustABI(t, cgc.CosmicSignatureNftABI)
	charity := mustABI(t, cgc.CharityWalletABI)
	prizes := mustABI(t, cgc.PrizesWalletABI)
	stakingCST := mustABI(t, cgc.IStakingWalletCosmicSignatureNftABI)
	stakingRWK := mustABI(t, cgc.IStakingWalletRandomWalkNftABI)
	marketing := mustABI(t, cgc.MarketingWalletABI)
	erc20 := mustABI(t, cgc.ERC20ABI)
	erc1967 := mustABI(t, cgc.IERC1967ABI)

	cases := []struct {
		constant string
		abi      *abi.ABI
		event    string
	}{
		{PRIZE_CLAIM_EVENT, game, "MainPrizeClaimed"},
		{BID_EVENT, game, "BidPlaced"},
		{BID_EVENT_V2, gameV2, "BidPlaced"},
		{ETH_DONATED_EVENT, game, "EthDonated"},
		{ETH_DONATED_WI_EVENT, game, "EthDonatedWithInfo"},
		{DONATION_RECEIVED_EVENT, charity, "DonationReceived"},
		{DONATION_SENT_EVENT, charity, "FundsTransferredToCharity"},
		{CHARITY_RECEIVER_CHANGED, charity, "CharityAddressChanged"},
		{CHARITY_WALLET_CHANGED, game, "CharityAddressChanged"},
		{TOKEN_NAME_EVENT, nft, "NftNameChanged"},
		{MINT_EVENT, nft, "NftMinted"},
		{NFT_ETH_DONATED_EVENT, prizes, "NftDonated"},
		{ERC20_DONATED, prizes, "TokenDonated"},
		{DONATED_TOKEN_CLAIMED, prizes, "DonatedTokenClaimed"},
		{DONATED_NFT_CLAIMED, prizes, "DonatedNftClaimed"},
		{ETH_PRIZE_DEPOSIT_EVENT, prizes, "EthReceived"},
		{ETH_PRIZE_WITHDRAWAL_EVENT, prizes, "EthWithdrawn"},
		{RAFFLE_ETH_PRIZE_EVENT, game, "RaffleWinnerBidderEthPrizeAllocated"},
		{RAFFLE_NFT_PRIZE_EVENT, game, "RaffleWinnerPrizePaid"},
		{ENDURANCE_PRIZE_EVENT, game, "EnduranceChampionPrizePaid"},
		{LASTCST_BIDDER_PRIZE_EVENT, game, "LastCstBidderPrizePaid"},
		{CHRONO_WARRIOR_PRIZE_EVENT, game, "ChronoWarriorPrizePaid"},
		{TRANSFER_EVT, erc20, "Transfer"},
		{CST_NFT_STAKED_EVENT, stakingCST, "NftStaked"},
		{RWALK_NFT_STAKED_EVENT, stakingRWK, "NftStaked"},
		{NFT_UNSTAKED_RWALK, stakingRWK, "NftUnstaked"},
		{NFT_UNSTAKED_CST, stakingCST, "NftUnstaked"},
		{STAKING_ETH_DEPOSIT_EVENT, stakingCST, "EthDepositReceived"},
		{FUND_TRANSFER_ERR, game, "FundTransferFailed"},
		{FIRST_BID_EVENT, game, "FirstBidPlacedInRound"},
		{PROXY_UPGRADED, game, "Upgraded"},
		{ADMIN_CHANGED, erc1967, "AdminChanged"},
		{TREASURER_CHANGED, marketing, "TreasurerAddressChanged"},
		{INITIALIZED, game, "Initialized"},
		{CHARITY_PERCENTAGE_CHANGED, game, "CharityEthDonationAmountPercentageChanged"},
		{PRIZE_PERCENTAGE_CHANGED, game, "MainEthPrizeAmountPercentageChanged"},
		{RAFFLE_PERCENTAGE_CHANGED, game, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged"},
		{STAKE_PERCENTAGE_CHANGED, game, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged"},
		{CHRONO_PERCENTAGE_CHANGED, game, "ChronoWarriorEthPrizeAmountPercentageChanged"},
		{NUM_RAFFLE_ETH_PRIZE_EVENTS_BIDDING_CHANGED, game, "NumRaffleEthPrizesForBiddersChanged"},
		{NUM_RAFFLE_NFT_PRIZE_EVENTS_BIDDING_CHANGED, game, "NumRaffleCosmicSignatureNftsForBiddersChanged"},
		{NUM_RAFFLE_NFT_PRIZE_EVENTS_STAKING_RWALK_CHANGED, game, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged"},
		{RWALK_ADDRESS_CHANGED, game, "RandomWalkNftAddressChanged"},
		{PRIZE_WALLET_ADDRESS_CHANGED, game, "PrizesWalletAddressChanged"},
		{STAKING_WALLET_CST_ADDRESS_CHANGED, game, "StakingWalletCosmicSignatureNftAddressChanged"},
		{STAKING_WALLET_RWALK_ADDRESS_CHANGED, game, "StakingWalletRandomWalkNftAddressChanged"},
		{MARKETING_ADDRESS_CHANGED, game, "MarketingWalletAddressChanged"},
		{COSMIC_TOKEN_ADDRESS_CHANGED, game, "CosmicSignatureTokenAddressChanged"},
		{COSMIC_SIGNATURE_ADDRESS_CHANGED, game, "CosmicSignatureNftAddressChanged"},
		{TIMEOUT_CLAIMPRIZE_CHANGED, game, "TimeoutDurationToClaimMainPrizeChanged"},
		{TIMEOUT_TO_WITHDRAW_PRIZE, prizes, "TimeoutDurationToWithdrawPrizesChanged"},
		{PRICE_INCREASE_CHANGED, game, "EthBidPriceIncreaseDivisorChanged"},
		{MAIN_PRIZE_MICROSECOND_INCREASE, game, "MainPrizeTimeIncrementInMicroSecondsChanged"},
		{INITIAL_SECONDS_UNTIL_PRIZE_CHANGED, game, "InitialDurationUntilMainPrizeDivisorChanged"},
		{ROUND_ACTIVATION_TIME_CHANGED, game, "RoundActivationTimeChanged"},
		{ETH_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED, game, "EthDutchAuctionDurationDivisorChanged"},
		{CST_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED, game, "CstDutchAuctionDurationDivisorChanged"},
		{ETH_DUTCH_AUCTION_ENDING_PRICE_DIVISOR_CHANGED, game, "EthDutchAuctionEndingBidPriceDivisorChanged"},
		{MARKETING_REWARD_PAID, marketing, "RewardPaid"},
		{MARKETING_REWARD_CHANGED, game, "MarketingWalletCstContributionAmountChanged"},
		{CST_REWARD_FOR_BIDDING_CHANGED, game, "CstRewardAmountForBiddingChanged"},
		{BID_CST_REWARD_AMOUNT_MULTIPLIER_CHANGED, gameV2, "BidCstRewardAmountMultiplierChanged"},
		{CST_DUTCH_AUCTION_DURATION_CHANGED, gameV2, "CstDutchAuctionDurationChanged"},
		{CST_DUTCH_AUCTION_DURATION_CHANGE_DIVISOR_CHANGED, gameV2, "CstDutchAuctionDurationChangeDivisorChanged"},
		{STATIC_CST_REWARD, game, "CstPrizeAmountChanged"},
		{MAX_MESSAGE_LENGTH, game, "BidMessageLengthMaxLimitChanged"},
		{TOKEN_SCRIPT_URL, nft, "NftGenerationScriptUriChanged"},
		{BASE_URI, nft, "NftBaseUriChanged"},
		{OWNERSHIP_TRANSFERRED, game, "OwnershipTransferred"},
		{STARTING_CST_MIN_LIM, game, "CstDutchAuctionBeginningBidPriceMinLimitChanged"},
		{FUNDS_TO_CHARITY, game, "FundsTransferredToCharity"},
		{DELAY_DURATION_ROUND, game, "DelayDurationBeforeRoundActivationChanged"},
	}

	for _, c := range cases {
		event, ok := c.abi.Events[c.event]
		if !ok {
			t.Errorf("event %s not present in ABI (constant %s)", c.event, c.constant)
			continue
		}
		if got := event.ID.Hex()[2:]; got != c.constant {
			t.Errorf("registry constant for %s = %s, but ABI event ID is %s", c.event, c.constant, got)
		}
	}
}

// TestLegacyConstantsHaveNoABIEvent documents the registry entries that no
// current ABI defines: their handlers must decode the raw data words instead
// of unpacking by name (doing the latter killed the process; see
// proc_time_increase_changed_event and proc_bid_cst_reward_amount_changed_event).
func TestLegacyConstantsHaveNoABIEvent(t *testing.T) {
	game := mustABI(t, cgc.CosmicSignatureGameABI)
	gameV2 := mustABI(t, cgc.CosmicSignatureGameV2ABI)
	for name, constant := range map[string]string{
		"TIME_INCREASE_CHANGED":         TIME_INCREASE_CHANGED,
		"BID_CST_REWARD_AMOUNT_CHANGED": BID_CST_REWARD_AMOUNT_CHANGED,
		"ERC20_TRANSFER_ERR":            ERC20_TRANSFER_ERR,
	} {
		for _, a := range []*abi.ABI{game, gameV2} {
			for _, ev := range a.Events {
				if ev.ID.Hex()[2:] == constant {
					t.Errorf("%s (%s) now exists in an ABI as %s: switch its handler from raw decoding to ABI unpacking",
						name, constant, ev.Name)
				}
			}
		}
	}
}

// TestERC20TransferFailedConstantMatchesSignature pins the ERC20_TRANSFER_ERR
// registry constant to the ICosmicSignatureErrors.sol event signature it is
// derived from. No generated ABI carries the event, so — unlike the entries
// covered by TestRegistryConstantsMatchABIEventIDs — the Solidity signature
// itself is the only cross-check, and it also documents the raw data layout
// proc_erc20_transfer_failed_event decodes (string errStr, address indexed
// destination, uint256 amount).
func TestERC20TransferFailedConstantMatchesSignature(t *testing.T) {
	got := crypto.Keccak256Hash([]byte("ERC20TransferFailed(string,address,uint256)")).Hex()[2:]
	if got != ERC20_TRANSFER_ERR {
		t.Fatalf("keccak256 of the event signature = %s, registry constant is %s", got, ERC20_TRANSFER_ERR)
	}
}
