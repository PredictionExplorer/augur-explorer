// The event-handler table: one registration per dispatched event type,
// pairing the topic-hash constant, the metric label, the emitting contracts
// and the decode/store implementation. The three shared-signature cases
// (CharityAddressChanged from two contracts, FundsTransferredToCharity from
// the CharityWallet versus the game and the CST staking wallet, Transfer from
// the ERC721 and the ERC20) are separate registrations distinguished by
// source.

package cosmicgame

import (
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

// eventHandlers returns every CosmicGame event handler in registration order.
func (h *Handlers) eventHandlers() []indexer.EventHandler {
	one := func(addr ethcommon.Address) []ethcommon.Address { return []ethcommon.Address{addr} }
	game := one(h.c.Game)
	signature := one(h.c.Signature)
	charity := one(h.c.CharityWallet)
	prizes := one(h.c.PrizesWallet)
	marketing := one(h.c.MarketingWallet)
	dao := one(h.c.Dao)
	// The legacy NftStaked guards accepted either staking wallet for both
	// event variants; preserved verbatim (the topics differ per wallet ABI,
	// so only one variant can arrive from each wallet in practice).
	stakingEither := []ethcommon.Address{h.c.StakingCST, h.c.StakingRWalk}
	// The contracts that hand ETH to the charity address directly: the game
	// forwarding the charity share in MainPrize._distributePrizes, and the CST
	// staking wallet sweeping its balance in tryPerformMaintenance. The
	// CharityWallet emits the same topic0 when it forwards its own balance;
	// that one is registered separately into cg_donation_sent.
	charitySenders := []ethcommon.Address{h.c.Game, h.c.StakingCST}

	return []indexer.EventHandler{
		indexer.NewHandler(topicHash(TopicPrizeClaimEvent), "MainPrizeClaimed", game, h.decodeMainPrizeClaimed, h.storeMainPrizeClaimed),
		indexer.NewHandler(topicHash(TopicPrizeClaimEventV3), "MainPrizeClaimedV3", game, h.decodeMainPrizeClaimedV3, h.storeMainPrizeClaimedV3),
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
		indexer.NewHandler(topicHash(TopicEthBidRefundGasMaxLimitChanged), "EthBidRefundAmountInGasToSwallowMaxLimitChanged", game, h.decodeEthBidRefundGasMaxLimitChanged, h.storeEthBidRefundGasMaxLimitChanged),
		indexer.NewHandler(topicHash(TopicTimeoutToWithdrawPrize), "TimeoutDurationToWithdrawPrizesChanged", prizes, h.decodeTimeoutToWithdrawPrizesChanged, h.storeTimeoutToWithdrawPrizesChanged),
		indexer.NewHandler(topicHash(TopicPriceIncreaseChanged), "EthBidPriceIncreaseDivisorChanged", game, h.decodePriceIncreaseChanged, h.storePriceIncreaseChanged),
		indexer.NewHandler(topicHash(TopicMainPrizeMicrosecondIncrease), "MainPrizeTimeIncrementInMicroSecondsChanged", game, h.decodeMainPrizeMicrosecondsChanged, h.storeMainPrizeMicrosecondsChanged),
		indexer.NewHandler(topicHash(TopicInitialSecondsUntilPrizeChanged), "InitialDurationUntilMainPrizeDivisorChanged", game, h.decodeInitialSecondsUntilPrizeChanged, h.storeInitialSecondsUntilPrizeChanged),
		indexer.NewHandler(topicHash(TopicRoundActivationTimeChanged), "RoundActivationTimeChanged", game, h.decodeRoundActivationTimeChanged, h.storeRoundActivationTimeChanged),
		indexer.NewHandler(topicHash(TopicCstDutchAuctionDurationDivisorChanged), "CstDutchAuctionDurationDivisorChanged", game, h.decodeCstDutchAuctionDurationDivisorChanged, h.storeCstAuctionLengthChange("CstDutchAuctionDurationDivisorChanged")),
		indexer.NewHandler(topicHash(TopicCstDutchAuctionDurationChanged), "CstDutchAuctionDurationChanged", game, h.decodeCstDutchAuctionDurationChanged, h.storeCstAuctionLengthChange("CstDutchAuctionDurationChanged")),
		indexer.NewHandler(topicHash(TopicCstDutchAuctionDurationChangeDivisorChanged), "CstDutchAuctionDurationChangeDivisorChanged", game, h.decodeCstAuctionDurationChangeDivisorChanged, h.storeCstAuctionDurationChangeDivisorChanged),
		indexer.NewHandler(topicHash(TopicRoundLateBidDurationDivisorChanged), "RoundLateBidDurationDivisorChanged", game, h.decodeRoundLateBidDurationDivisorChanged, h.storeRoundLateBidDurationDivisorChanged),
		indexer.NewHandler(topicHash(TopicRoundLateBidPremiumBaseMultiplierChanged), "RoundLateBidPricePremiumAmountBaseMultiplierChanged", game, h.decodeRoundLateBidPremiumBaseMultiplierChanged, h.storeRoundLateBidPremiumBaseMultiplierChanged),
		indexer.NewHandler(topicHash(TopicRoundLateBidPremiumExponentChanged), "RoundLateBidPricePremiumAmountExponentChanged", game, h.decodeRoundLateBidPremiumExponentChanged, h.storeRoundLateBidPremiumExponentChanged),
		indexer.NewHandler(topicHash(TopicCstBidPriceDeclineMultiplierChanged), "CstBidPriceDeclineMultiplierChanged", game, h.decodeCstBidPriceDeclineMultiplierChanged, h.storeCstBidPriceDeclineMultiplierChanged),
		indexer.NewHandler(topicHash(TopicCstBidPriceDeclineMultiplierChangeDivisorChanged), "CstBidPriceDeclineMultiplierChangeDivisorChanged", game, h.decodeCstBidPriceDeclineMultiplierChangeDivisorChanged, h.storeCstBidPriceDeclineMultiplierChangeDivisorChanged),
		indexer.NewHandler(topicHash(TopicMainPrizeNumNftsChanged), "MainPrizeNumCosmicSignatureNftsChanged", game, h.decodeMainPrizeNumNftsChanged, h.storeMainPrizeNumNftsChanged),
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
		// Initialized is emitted by every platform contract's initializer and
		// by each game implementation's constructor; the implementation set
		// grows as Upgraded events arrive, so the sources are dynamic.
		indexer.NewDynamicHandler(topicHash(TopicInitialized), "Initialized", h.initializedSources, h.decodeInitialized, h.storeInitialized),
		indexer.NewHandler(topicHash(TopicStartingCstMinLim), "CstDutchAuctionBeginningBidPriceMinLimitChanged", game, h.decodeCstMinLimitChanged, h.storeCstMinLimitChanged),
		// FundTransferFailed: the game emitted it on charity-donation failure
		// before V3.1; the CST staking wallet still emits it when
		// tryPerformMaintenance cannot forward its balance to charity.
		indexer.NewHandler(topicHash(TopicFundTransferErr), "FundTransferFailed", charitySenders, h.decodeFundTransferFailed, h.storeFundTransferFailed),
		indexer.NewHandler(topicHash(TopicERC20TransferErr), "ERC20TransferFailed", game, h.decodeERC20TransferFailed, h.storeERC20TransferFailed),
		// Only MainPrize/MainPrizeV2 reach ArbitrumHelpers, so the game is the
		// sole emitter of ArbitrumError.
		indexer.NewHandler(topicHash(TopicArbitrumError), "ArbitrumError", game, h.decodeArbitrumError, h.storeArbitrumError),
		// V3.1 replaced ArbitrumError(string) with one parameterless event per
		// failing precompile call; all four land in cg_arbitrum_error with the
		// legacy message wording.
		indexer.NewHandler(topicHash(TopicArbSysArbBlockNumberCallFailed), "ArbSysArbBlockNumberCallFailed", game,
			h.decodeArbitrumCallFailed("ArbSys.arbBlockNumber call failed."), h.storeArbitrumError),
		indexer.NewHandler(topicHash(TopicArbSysArbBlockHashCallFailed), "ArbSysArbBlockHashCallFailed", game,
			h.decodeArbitrumCallFailed("ArbSys.arbBlockHash call failed."), h.storeArbitrumError),
		indexer.NewHandler(topicHash(TopicArbGasInfoGetGasBacklogCallFailed), "ArbGasInfoGetGasBacklogCallFailed", game,
			h.decodeArbitrumCallFailed("ArbGasInfo.getGasBacklog call failed."), h.storeArbitrumError),
		indexer.NewHandler(topicHash(TopicArbGasInfoGetL1PricingUnitsCallFailed), "ArbGasInfoGetL1PricingUnitsSinceUpdateCallFailed", game,
			h.decodeArbitrumCallFailed("ArbGasInfo.getL1PricingUnitsSinceUpdate call failed."), h.storeArbitrumError),
		// V3.1 charity-donation failure event (replaces the game's
		// FundTransferFailed emission on that path).
		indexer.NewHandler(topicHash(TopicEthTransferToCharityFailed), "EthTransferToCharityFailed", game,
			h.decodeEthTransferToCharityFailed, h.storeEthTransferToCharityFailed),
		indexer.NewHandler(topicHash(TopicFundsToCharity), "FundsTransferredToCharity", charitySenders, h.decodeFundsToCharity, h.storeFundsToCharity),
		indexer.NewHandler(topicHash(TopicDelayDurationRound), "DelayDurationBeforeRoundActivationChanged", game, h.decodeDelayDurationChanged, h.storeDelayDurationChanged),
		indexer.NewHandler(topicHash(TopicFirstBidEvent), "FirstBidPlacedInRound", game, h.decodeFirstBidPlacedInRound, h.storeFirstBidPlacedInRound),

		// OpenZeppelin events inherited by the NFT and the token. Approval
		// shares its topic0 between ERC721 (tokenId indexed) and ERC20 (value
		// in data); the source filter separates them like Transfer.
		indexer.NewHandler(topicHash(TopicApproval), "Approval", signature, h.decodeNftApproval, h.storeNftApproval),
		indexer.NewHandler(topicHash(TopicApprovalForAll), "NftApprovalForAll", signature, h.decodeNftApprovalForAll, h.storeNftApprovalForAll),
		indexer.NewHandler(topicHash(TopicApproval), "Approval", one(h.c.Token), h.decodeTokenApproval, h.storeTokenApproval),
		indexer.NewHandler(topicHash(TopicDelegateChanged), "DelegateChanged", one(h.c.Token), h.decodeDelegateChanged, h.storeDelegateChanged),
		indexer.NewHandler(topicHash(TopicDelegateVotesChanged), "DelegateVotesChanged", one(h.c.Token), h.decodeDelegateVotesChanged, h.storeDelegateVotesChanged),
		indexer.NewHandler(topicHash(TopicEIP712DomainChanged), "EIP712DomainChanged", []ethcommon.Address{h.c.Token, h.c.Dao}, h.decodeEIP712DomainChanged, h.storeEIP712DomainChanged),

		// CosmicSignatureDao (OpenZeppelin Governor).
		indexer.NewHandler(topicHash(TopicDaoProposalCreated), "DaoProposalCreated", dao, h.decodeDaoProposalCreated, h.storeDaoProposalCreated),
		indexer.NewHandler(topicHash(TopicDaoProposalCanceled), "DaoProposalCanceled", dao, h.decodeDaoProposalCanceled, h.storeDaoProposalStateChange),
		indexer.NewHandler(topicHash(TopicDaoProposalExecuted), "DaoProposalExecuted", dao, h.decodeDaoProposalExecuted, h.storeDaoProposalStateChange),
		indexer.NewHandler(topicHash(TopicDaoProposalQueued), "DaoProposalQueued", dao, h.decodeDaoProposalQueued, h.storeDaoProposalStateChange),
		indexer.NewHandler(topicHash(TopicDaoVoteCast), "DaoVoteCast", dao, h.decodeDaoVoteCast, h.storeDaoVoteCast),
		indexer.NewHandler(topicHash(TopicDaoVoteCastWithParams), "DaoVoteCastWithParams", dao, h.decodeDaoVoteCastWithParams, h.storeDaoVoteCast),
		indexer.NewHandler(topicHash(TopicDaoProposalThresholdSet), "DaoProposalThresholdSet", dao, h.decodeDaoSettingChanged("ProposalThresholdSet", cgmodel.DaoSettingProposalThreshold), h.storeDaoSettingChanged),
		indexer.NewHandler(topicHash(TopicDaoVotingDelaySet), "DaoVotingDelaySet", dao, h.decodeDaoSettingChanged("VotingDelaySet", cgmodel.DaoSettingVotingDelay), h.storeDaoSettingChanged),
		indexer.NewHandler(topicHash(TopicDaoVotingPeriodSet), "DaoVotingPeriodSet", dao, h.decodeDaoSettingChanged("VotingPeriodSet", cgmodel.DaoSettingVotingPeriod), h.storeDaoSettingChanged),
		indexer.NewHandler(topicHash(TopicDaoQuorumNumeratorSet), "DaoQuorumNumeratorUpdated", dao, h.decodeDaoSettingChanged("QuorumNumeratorUpdated", cgmodel.DaoSettingQuorumNumerator), h.storeDaoSettingChanged),
	}
}
