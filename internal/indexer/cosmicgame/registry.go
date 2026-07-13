// The event-handler table: one registration per dispatched event type,
// pairing the topic-hash constant, the metric label, the emitting contracts
// and the decode/store implementation. The three shared-signature cases
// (CharityAddressChanged from two contracts, FundsTransferredToCharity from
// two contracts, Transfer from the ERC721 and the ERC20) are separate
// registrations distinguished by source.

package cosmicgame

import (
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
)

// eventHandlers returns every CosmicGame event handler in registration order.
func (h *Handlers) eventHandlers() []indexer.EventHandler {
	one := func(addr ethcommon.Address) []ethcommon.Address { return []ethcommon.Address{addr} }
	game := one(h.c.Game)
	signature := one(h.c.Signature)
	charity := one(h.c.CharityWallet)
	prizes := one(h.c.PrizesWallet)
	marketing := one(h.c.MarketingWallet)
	// The legacy NftStaked guards accepted either staking wallet for both
	// event variants; preserved verbatim (the topics differ per wallet ABI,
	// so only one variant can arrive from each wallet in practice).
	stakingEither := []ethcommon.Address{h.c.StakingCST, h.c.StakingRWalk}

	return []indexer.EventHandler{
		indexer.NewHandler(topicHash(TopicPrizeClaimEvent), "MainPrizeClaimed", game, h.decodeMainPrizeClaimed, h.storeMainPrizeClaimed),
		indexer.NewHandler(topicHash(TopicBidEvent), "BidPlaced", game, h.decodeBidPlacedV1, h.storeBid),
		indexer.NewHandler(topicHash(TopicBidEventV2), "BidPlacedV2", game, h.decodeBidPlacedV2, h.storeBid),
		indexer.NewHandler(topicHash(TopicEthDonatedEvent), "EthDonated", game, h.decodeEthDonated, h.storeEthDonated),
		indexer.NewHandler(topicHash(TopicEthDonatedWIEvent), "EthDonatedWithInfo", game, h.decodeEthDonatedWithInfo, h.storeEthDonatedWithInfo),
		indexer.NewHandler(topicHash(TopicDonationReceivedEvent), "DonationReceived", charity, h.decodeDonationReceived, h.storeDonationReceived),
		indexer.NewHandler(topicHash(TopicDonationSentEvent), "FundsTransferredToCharity", charity, h.decodeDonationSent, h.storeDonationSent),
		indexer.NewHandler(topicHash(TopicNftEthDonatedEvent), "NftDonated", prizes, h.decodeNftDonated, h.storeNftDonated),
		indexer.NewHandler(topicHash(TopicERC20Donated), "TokenDonated", prizes, h.decodeTokenDonated, h.storeTokenDonated),
		indexer.NewHandler(topicHash(TopicCharityReceiverChanged), "CharityAddressChanged", charity, h.decodeCharityReceiverChanged, h.storeCharityReceiverChanged),
		indexer.NewHandler(topicHash(TopicCharityWalletChanged), "CharityAddressChanged", game, h.decodeCharityWalletChanged, h.storeCharityWalletChanged),
		indexer.NewHandler(topicHash(TopicTokenNameEvent), "NftNameChanged", signature, h.decodeNftNameChanged, h.storeNftNameChanged),
		indexer.NewHandler(topicHash(TopicMintEvent), "NftMinted", signature, h.decodeNftMinted, h.storeNftMinted),
		indexer.NewHandler(topicHash(TopicEthPrizeDepositEvent), "EthReceived", prizes, h.decodePrizesEthReceived, h.storePrizesEthReceived),
		indexer.NewHandler(topicHash(TopicEthPrizeWithdrawalEvent), "EthWithdrawn", prizes, h.decodePrizesEthWithdrawn, h.storePrizesEthWithdrawn),
		indexer.NewHandler(topicHash(TopicRaffleEthPrizeEvent), "RaffleWinnerBidderEthPrizeAllocated", game, h.decodeRaffleEthAllocated, h.storeRaffleEthAllocated),
		indexer.NewHandler(topicHash(TopicRaffleNftPrizeEvent), "RaffleWinnerPrizePaid", game, h.decodeRaffleWinnerPrizePaid, h.storeRaffleWinnerPrizePaid),
		indexer.NewHandler(topicHash(TopicEndurancePrizeEvent), "EnduranceChampionPrizePaid", game, h.decodeEnduranceChampionPrizePaid, h.storeEnduranceChampionPrizePaid),
		indexer.NewHandler(topicHash(TopicLastcstBidderPrizeEvent), "LastCstBidderPrizePaid", game, h.decodeLastCstBidderPrizePaid, h.storeLastCstBidderPrizePaid),
		indexer.NewHandler(topicHash(TopicChronoWarriorPrizeEvent), "ChronoWarriorPrizePaid", game, h.decodeChronoWarriorPrizePaid, h.storeChronoWarriorPrizePaid),
		indexer.NewHandler(topicHash(TopicDonatedTokenClaimed), "DonatedTokenClaimed", prizes, h.decodeDonatedTokenClaimed, h.storeDonatedTokenClaimed),
		indexer.NewHandler(topicHash(TopicDonatedNftClaimed), "DonatedNftClaimed", prizes, h.decodeDonatedNftClaimed, h.storeDonatedNftClaimed),
		indexer.NewHandler(topicHash(TopicTransferEvt), "Transfer", signature, h.decodeCosmicSignatureTransfer, h.storeCosmicSignatureTransfer),
		indexer.NewHandler(topicHash(TopicTransferEvt), "Transfer", one(h.c.Token), h.decodeCosmicTokenTransfer, h.storeCosmicTokenTransfer),
		indexer.NewHandler(topicHash(TopicCstNftStakedEvent), "NftStakedCST", stakingEither, h.decodeNftStakedCST, h.storeNftStakedCST),
		indexer.NewHandler(topicHash(TopicRwalkNftStakedEvent), "NftStakedRWalk", stakingEither, h.decodeNftStakedRWalk, h.storeNftStakedRWalk),
		indexer.NewHandler(topicHash(TopicNftUnstakedRwalk), "NftUnstakedRWalk", one(h.c.StakingRWalk), h.decodeNftUnstakedRWalk, h.storeNftUnstakedRWalk),
		indexer.NewHandler(topicHash(TopicNftUnstakedCst), "NftUnstakedCST", one(h.c.StakingCST), h.decodeNftUnstakedCST, h.storeNftUnstakedCST),
		indexer.NewHandler(topicHash(TopicStakingEthDepositEvent), "EthDepositReceived", one(h.c.StakingCST), h.decodeStakingEthDeposit, h.storeStakingEthDeposit),
		indexer.NewHandler(topicHash(TopicMarketingRewardPaid), "RewardPaid", marketing, h.decodeMarketingRewardPaid, h.storeMarketingRewardPaid),
		indexer.NewHandler(topicHash(TopicCharityPercentageChanged), "CharityEthDonationAmountPercentageChanged", game, h.decodeCharityPercentageChanged, h.storeCharityPercentageChanged),
		indexer.NewHandler(topicHash(TopicPrizePercentageChanged), "MainEthPrizeAmountPercentageChanged", game, h.decodePrizePercentageChanged, h.storePrizePercentageChanged),
		indexer.NewHandler(topicHash(TopicRafflePercentageChanged), "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", game, h.decodeRafflePercentageChanged, h.storeRafflePercentageChanged),
		indexer.NewHandler(topicHash(TopicStakePercentageChanged), "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", game, h.decodeStakingPercentageChanged, h.storeStakingPercentageChanged),
		indexer.NewHandler(topicHash(TopicChronoPercentageChanged), "ChronoWarriorEthPrizeAmountPercentageChanged", game, h.decodeChronoPercentageChanged, h.storeChronoPercentageChanged),
		indexer.NewHandler(topicHash(TopicNumRaffleEthPrizeEventsBiddingChanged), "NumRaffleEthPrizesForBiddersChanged", game, h.decodeNumRaffleETHWinnersBiddingChanged, h.storeNumRaffleETHWinnersBiddingChanged),
		indexer.NewHandler(topicHash(TopicNumRaffleNftPrizeEventsBiddingChanged), "NumRaffleCosmicSignatureNftsForBiddersChanged", game, h.decodeNumRaffleNFTWinnersBiddingChanged, h.storeNumRaffleNFTWinnersBiddingChanged),
		indexer.NewHandler(topicHash(TopicNumRaffleNftPrizeEventsStakingRwalkChanged), "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", game, h.decodeNumRaffleNFTWinnersStakingRWalkChanged, h.storeNumRaffleNFTWinnersStakingRWalkChanged),
		indexer.NewHandler(topicHash(TopicRwalkAddressChanged), "RandomWalkNftAddressChanged", game, h.decodeRandomWalkAddressChanged, h.storeRandomWalkAddressChanged),
		indexer.NewHandler(topicHash(TopicPrizeWalletAddressChanged), "PrizesWalletAddressChanged", game, h.decodePrizesWalletAddressChanged, h.storePrizesWalletAddressChanged),
		indexer.NewHandler(topicHash(TopicStakingWalletCstAddressChanged), "StakingWalletCosmicSignatureNftAddressChanged", game, h.decodeStakingWalletCSTAddressChanged, h.storeStakingWalletCSTAddressChanged),
		indexer.NewHandler(topicHash(TopicStakingWalletRwalkAddressChanged), "StakingWalletRandomWalkNftAddressChanged", game, h.decodeStakingWalletRWalkAddressChanged, h.storeStakingWalletRWalkAddressChanged),
		indexer.NewHandler(topicHash(TopicMarketingAddressChanged), "MarketingWalletAddressChanged", game, h.decodeMarketingWalletAddressChanged, h.storeMarketingWalletAddressChanged),
		indexer.NewHandler(topicHash(TopicTreasurerChanged), "TreasurerAddressChanged", marketing, h.decodeTreasurerAddressChanged, h.storeTreasurerAddressChanged),
		indexer.NewHandler(topicHash(TopicCosmicTokenAddressChanged), "CosmicSignatureTokenAddressChanged", game, h.decodeCosmicTokenAddressChanged, h.storeCosmicTokenAddressChanged),
		indexer.NewHandler(topicHash(TopicCosmicSignatureAddressChanged), "CosmicSignatureNftAddressChanged", game, h.decodeCosmicSignatureAddressChanged, h.storeCosmicSignatureAddressChanged),
		indexer.NewHandler(topicHash(TopicProxyUpgraded), "Upgraded", game, h.decodeUpgraded, h.storeUpgraded),
		indexer.NewHandler(topicHash(TopicAdminChanged), "AdminChanged", game, h.decodeAdminChanged, h.storeAdminChanged),
		indexer.NewHandler(topicHash(TopicTimeIncreaseChanged), "MainPrizeTimeIncrementIncreaseDivisorChanged", game, h.decodeTimeIncreaseChanged, h.storeTimeIncreaseChanged),
		indexer.NewHandler(topicHash(TopicTimeoutClaimprizeChanged), "TimeoutDurationToClaimMainPrizeChanged", game, h.decodeTimeoutClaimPrizeChanged, h.storeTimeoutClaimPrizeChanged),
		indexer.NewHandler(topicHash(TopicTimeoutToWithdrawPrize), "TimeoutDurationToWithdrawPrizesChanged", prizes, h.decodeTimeoutToWithdrawPrizesChanged, h.storeTimeoutToWithdrawPrizesChanged),
		indexer.NewHandler(topicHash(TopicPriceIncreaseChanged), "EthBidPriceIncreaseDivisorChanged", game, h.decodePriceIncreaseChanged, h.storePriceIncreaseChanged),
		indexer.NewHandler(topicHash(TopicMainPrizeMicrosecondIncrease), "MainPrizeTimeIncrementInMicroSecondsChanged", game, h.decodeMainPrizeMicrosecondsChanged, h.storeMainPrizeMicrosecondsChanged),
		indexer.NewHandler(topicHash(TopicInitialSecondsUntilPrizeChanged), "InitialDurationUntilMainPrizeDivisorChanged", game, h.decodeInitialSecondsUntilPrizeChanged, h.storeInitialSecondsUntilPrizeChanged),
		indexer.NewHandler(topicHash(TopicRoundActivationTimeChanged), "RoundActivationTimeChanged", game, h.decodeRoundActivationTimeChanged, h.storeRoundActivationTimeChanged),
		indexer.NewHandler(topicHash(TopicCstDutchAuctionDurationDivisorChanged), "CstDutchAuctionDurationDivisorChanged", game, h.decodeCstDutchAuctionDurationDivisorChanged, h.storeCstAuctionLengthChange("CstDutchAuctionDurationDivisorChanged")),
		indexer.NewHandler(topicHash(TopicCstDutchAuctionDurationChanged), "CstDutchAuctionDurationChanged", game, h.decodeCstDutchAuctionDurationChanged, h.storeCstAuctionLengthChange("CstDutchAuctionDurationChanged")),
		indexer.NewHandler(topicHash(TopicCstDutchAuctionDurationChangeDivisorChanged), "CstDutchAuctionDurationChangeDivisorChanged", game, h.decodeCstAuctionDurationChangeDivisorChanged, h.storeCstAuctionDurationChangeDivisorChanged),
		indexer.NewHandler(topicHash(TopicEthDutchAuctionDurationDivisorChanged), "EthDutchAuctionDurationDivisorChanged", game, h.decodeEthAuctionDurationDivisorChanged, h.storeEthAuctionDurationDivisorChanged),
		indexer.NewHandler(topicHash(TopicEthDutchAuctionEndingPriceDivisorChanged), "EthDutchAuctionEndingBidPriceDivisorChanged", game, h.decodeEthAuctionEndingBidPriceDivisorChanged, h.storeEthAuctionEndingBidPriceDivisorChanged),
		indexer.NewHandler(topicHash(TopicCstRewardForBiddingChanged), "CstRewardAmountForBiddingChanged", game, h.decodeCstRewardAmountForBiddingChanged, h.storeCstRewardForBiddingChange("CstRewardAmountForBiddingChanged")),
		indexer.NewHandler(topicHash(TopicBidCstRewardAmountChanged), "BidCstRewardAmountChanged", game, h.decodeBidCstRewardAmountChanged, h.storeCstRewardForBiddingChange("BidCstRewardAmountChanged")),
		indexer.NewHandler(topicHash(TopicBidCstRewardAmountMultiplierChanged), "BidCstRewardAmountMultiplierChanged", game, h.decodeBidCstRewardAmountMultiplierChanged, h.storeCstRewardForBiddingChange("BidCstRewardAmountMultiplierChanged")),
		indexer.NewHandler(topicHash(TopicStaticCstReward), "CstPrizeAmountChanged", game, h.decodeStaticCstRewardChanged, h.storeStaticCstRewardChanged),
		indexer.NewHandler(topicHash(TopicMaxMessageLength), "BidMessageLengthMaxLimitChanged", game, h.decodeMaxMessageLengthChanged, h.storeMaxMessageLengthChanged),
		indexer.NewHandler(topicHash(TopicTokenScriptURL), "NftGenerationScriptUriChanged", signature, h.decodeNftGenerationScriptURLChanged, h.storeNftGenerationScriptURLChanged),
		indexer.NewHandler(topicHash(TopicBaseURI), "NftBaseUriChanged", signature, h.decodeNftBaseURIChanged, h.storeNftBaseURIChanged),
		indexer.NewHandler(topicHash(TopicMarketingRewardChanged), "MarketingWalletCstContributionAmountChanged", game, h.decodeMarketingRewardChanged, h.storeMarketingRewardChanged),
		indexer.NewHandler(topicHash(TopicOwnershipTransferred), "OwnershipTransferred", h.ownershipSources(), h.decodeOwnershipTransferred, h.storeOwnershipTransferred),
		indexer.NewHandler(topicHash(TopicInitialized), "Initialized", h.initializedSources(), h.decodeInitialized, h.storeInitialized),
		indexer.NewHandler(topicHash(TopicStartingCstMinLim), "CstDutchAuctionBeginningBidPriceMinLimitChanged", game, h.decodeCstMinLimitChanged, h.storeCstMinLimitChanged),
		indexer.NewHandler(topicHash(TopicFundTransferErr), "FundTransferFailed", game, h.decodeFundTransferFailed, h.storeFundTransferFailed),
		indexer.NewHandler(topicHash(TopicERC20TransferErr), "ERC20TransferFailed", game, h.decodeERC20TransferFailed, h.storeERC20TransferFailed),
		indexer.NewHandler(topicHash(TopicFundsToCharity), "FundsTransferredToCharity", marketing, h.decodeFundsToCharity, h.storeFundsToCharity),
		indexer.NewHandler(topicHash(TopicDelayDurationRound), "DelayDurationBeforeRoundActivationChanged", game, h.decodeDelayDurationChanged, h.storeDelayDurationChanged),
		indexer.NewHandler(topicHash(TopicFirstBidEvent), "FirstBidPlacedInRound", game, h.decodeFirstBidPlacedInRound, h.storeFirstBidPlacedInRound),
	}
}
