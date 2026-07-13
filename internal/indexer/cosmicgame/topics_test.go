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
		{TopicPrizeClaimEvent, game, "MainPrizeClaimed"},
		{TopicBidEvent, game, "BidPlaced"},
		{TopicBidEventV2, gameV2, "BidPlaced"},
		{TopicEthDonatedEvent, game, "EthDonated"},
		{TopicEthDonatedWIEvent, game, "EthDonatedWithInfo"},
		{TopicDonationReceivedEvent, charity, "DonationReceived"},
		{TopicDonationSentEvent, charity, "FundsTransferredToCharity"},
		{TopicCharityReceiverChanged, charity, "CharityAddressChanged"},
		{TopicCharityWalletChanged, game, "CharityAddressChanged"},
		{TopicTokenNameEvent, nft, "NftNameChanged"},
		{TopicMintEvent, nft, "NftMinted"},
		{TopicNftEthDonatedEvent, prizes, "NftDonated"},
		{TopicERC20Donated, prizes, "TokenDonated"},
		{TopicDonatedTokenClaimed, prizes, "DonatedTokenClaimed"},
		{TopicDonatedNftClaimed, prizes, "DonatedNftClaimed"},
		{TopicEthPrizeDepositEvent, prizes, "EthReceived"},
		{TopicEthPrizeWithdrawalEvent, prizes, "EthWithdrawn"},
		{TopicRaffleEthPrizeEvent, game, "RaffleWinnerBidderEthPrizeAllocated"},
		{TopicRaffleNftPrizeEvent, game, "RaffleWinnerPrizePaid"},
		{TopicEndurancePrizeEvent, game, "EnduranceChampionPrizePaid"},
		{TopicLastcstBidderPrizeEvent, game, "LastCstBidderPrizePaid"},
		{TopicChronoWarriorPrizeEvent, game, "ChronoWarriorPrizePaid"},
		{TopicTransferEvt, erc20, "Transfer"},
		{TopicCstNftStakedEvent, stakingCST, "NftStaked"},
		{TopicRwalkNftStakedEvent, stakingRWK, "NftStaked"},
		{TopicNftUnstakedRwalk, stakingRWK, "NftUnstaked"},
		{TopicNftUnstakedCst, stakingCST, "NftUnstaked"},
		{TopicStakingEthDepositEvent, stakingCST, "EthDepositReceived"},
		{TopicFundTransferErr, game, "FundTransferFailed"},
		{TopicFirstBidEvent, game, "FirstBidPlacedInRound"},
		{TopicProxyUpgraded, game, "Upgraded"},
		{TopicAdminChanged, erc1967, "AdminChanged"},
		{TopicTreasurerChanged, marketing, "TreasurerAddressChanged"},
		{TopicInitialized, game, "Initialized"},
		{TopicCharityPercentageChanged, game, "CharityEthDonationAmountPercentageChanged"},
		{TopicPrizePercentageChanged, game, "MainEthPrizeAmountPercentageChanged"},
		{TopicRafflePercentageChanged, game, "RaffleTotalEthPrizeAmountForBiddersPercentageChanged"},
		{TopicStakePercentageChanged, game, "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged"},
		{TopicChronoPercentageChanged, game, "ChronoWarriorEthPrizeAmountPercentageChanged"},
		{TopicNumRaffleEthPrizeEventsBiddingChanged, game, "NumRaffleEthPrizesForBiddersChanged"},
		{TopicNumRaffleNftPrizeEventsBiddingChanged, game, "NumRaffleCosmicSignatureNftsForBiddersChanged"},
		{TopicNumRaffleNftPrizeEventsStakingRwalkChanged, game, "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged"},
		{TopicRwalkAddressChanged, game, "RandomWalkNftAddressChanged"},
		{TopicPrizeWalletAddressChanged, game, "PrizesWalletAddressChanged"},
		{TopicStakingWalletCstAddressChanged, game, "StakingWalletCosmicSignatureNftAddressChanged"},
		{TopicStakingWalletRwalkAddressChanged, game, "StakingWalletRandomWalkNftAddressChanged"},
		{TopicMarketingAddressChanged, game, "MarketingWalletAddressChanged"},
		{TopicCosmicTokenAddressChanged, game, "CosmicSignatureTokenAddressChanged"},
		{TopicCosmicSignatureAddressChanged, game, "CosmicSignatureNftAddressChanged"},
		{TopicTimeoutClaimprizeChanged, game, "TimeoutDurationToClaimMainPrizeChanged"},
		{TopicTimeoutToWithdrawPrize, prizes, "TimeoutDurationToWithdrawPrizesChanged"},
		{TopicPriceIncreaseChanged, game, "EthBidPriceIncreaseDivisorChanged"},
		{TopicMainPrizeMicrosecondIncrease, game, "MainPrizeTimeIncrementInMicroSecondsChanged"},
		{TopicInitialSecondsUntilPrizeChanged, game, "InitialDurationUntilMainPrizeDivisorChanged"},
		{TopicRoundActivationTimeChanged, game, "RoundActivationTimeChanged"},
		{TopicEthDutchAuctionDurationDivisorChanged, game, "EthDutchAuctionDurationDivisorChanged"},
		{TopicCstDutchAuctionDurationDivisorChanged, game, "CstDutchAuctionDurationDivisorChanged"},
		{TopicEthDutchAuctionEndingPriceDivisorChanged, game, "EthDutchAuctionEndingBidPriceDivisorChanged"},
		{TopicMarketingRewardPaid, marketing, "RewardPaid"},
		{TopicMarketingRewardChanged, game, "MarketingWalletCstContributionAmountChanged"},
		{TopicCstRewardForBiddingChanged, game, "CstRewardAmountForBiddingChanged"},
		{TopicBidCstRewardAmountMultiplierChanged, gameV2, "BidCstRewardAmountMultiplierChanged"},
		{TopicCstDutchAuctionDurationChanged, gameV2, "CstDutchAuctionDurationChanged"},
		{TopicCstDutchAuctionDurationChangeDivisorChanged, gameV2, "CstDutchAuctionDurationChangeDivisorChanged"},
		{TopicStaticCstReward, game, "CstPrizeAmountChanged"},
		{TopicMaxMessageLength, game, "BidMessageLengthMaxLimitChanged"},
		{TopicTokenScriptURL, nft, "NftGenerationScriptUriChanged"},
		{TopicBaseURI, nft, "NftBaseUriChanged"},
		{TopicOwnershipTransferred, game, "OwnershipTransferred"},
		{TopicStartingCstMinLim, game, "CstDutchAuctionBeginningBidPriceMinLimitChanged"},
		{TopicFundsToCharity, game, "FundsTransferredToCharity"},
		{TopicDelayDurationRound, game, "DelayDurationBeforeRoundActivationChanged"},
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
		"TopicTimeIncreaseChanged":       TopicTimeIncreaseChanged,
		"TopicBidCstRewardAmountChanged": TopicBidCstRewardAmountChanged,
		"TopicERC20TransferErr":          TopicERC20TransferErr,
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

// TestERC20TransferFailedConstantMatchesSignature pins the TopicERC20TransferErr
// registry constant to the ICosmicSignatureErrors.sol event signature it is
// derived from. No generated ABI carries the event, so — unlike the entries
// covered by TestRegistryConstantsMatchABIEventIDs — the Solidity signature
// itself is the only cross-check, and it also documents the raw data layout
// proc_erc20_transfer_failed_event decodes (string errStr, address indexed
// destination, uint256 amount).
func TestERC20TransferFailedConstantMatchesSignature(t *testing.T) {
	got := crypto.Keccak256Hash([]byte("ERC20TransferFailed(string,address,uint256)")).Hex()[2:]
	if got != TopicERC20TransferErr {
		t.Fatalf("keccak256 of the event signature = %s, registry constant is %s", got, TopicERC20TransferErr)
	}
}
