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
		indexer.NewHandler(topicHash(PRIZE_CLAIM_EVENT), "MainPrizeClaimed", game, h.decodeMainPrizeClaimed, h.storeMainPrizeClaimed),
		indexer.NewHandler(topicHash(BID_EVENT), "BidPlaced", game, h.decodeBidPlacedV1, h.storeBid),
		indexer.NewHandler(topicHash(BID_EVENT_V2), "BidPlacedV2", game, h.decodeBidPlacedV2, h.storeBid),
		indexer.NewHandler(topicHash(ETH_DONATED_EVENT), "EthDonated", game, h.decodeEthDonated, h.storeEthDonated),
		indexer.NewHandler(topicHash(ETH_DONATED_WI_EVENT), "EthDonatedWithInfo", game, h.decodeEthDonatedWithInfo, h.storeEthDonatedWithInfo),
		indexer.NewHandler(topicHash(DONATION_RECEIVED_EVENT), "DonationReceived", charity, h.decodeDonationReceived, h.storeDonationReceived),
		indexer.NewHandler(topicHash(DONATION_SENT_EVENT), "FundsTransferredToCharity", charity, h.decodeDonationSent, h.storeDonationSent),
		indexer.NewHandler(topicHash(NFT_ETH_DONATED_EVENT), "NftDonated", prizes, h.decodeNftDonated, h.storeNftDonated),
		indexer.NewHandler(topicHash(ERC20_DONATED), "TokenDonated", prizes, h.decodeTokenDonated, h.storeTokenDonated),
		indexer.NewHandler(topicHash(CHARITY_RECEIVER_CHANGED), "CharityAddressChanged", charity, h.decodeCharityReceiverChanged, h.storeCharityReceiverChanged),
		indexer.NewHandler(topicHash(CHARITY_WALLET_CHANGED), "CharityAddressChanged", game, h.decodeCharityWalletChanged, h.storeCharityWalletChanged),
		indexer.NewHandler(topicHash(TOKEN_NAME_EVENT), "NftNameChanged", signature, h.decodeNftNameChanged, h.storeNftNameChanged),
		indexer.NewHandler(topicHash(MINT_EVENT), "NftMinted", signature, h.decodeNftMinted, h.storeNftMinted),
		indexer.NewHandler(topicHash(ETH_PRIZE_DEPOSIT_EVENT), "EthReceived", prizes, h.decodePrizesEthReceived, h.storePrizesEthReceived),
		indexer.NewHandler(topicHash(ETH_PRIZE_WITHDRAWAL_EVENT), "EthWithdrawn", prizes, h.decodePrizesEthWithdrawn, h.storePrizesEthWithdrawn),
		indexer.NewHandler(topicHash(RAFFLE_ETH_PRIZE_EVENT), "RaffleWinnerBidderEthPrizeAllocated", game, h.decodeRaffleEthAllocated, h.storeRaffleEthAllocated),
		indexer.NewHandler(topicHash(RAFFLE_NFT_PRIZE_EVENT), "RaffleWinnerPrizePaid", game, h.decodeRaffleWinnerPrizePaid, h.storeRaffleWinnerPrizePaid),
		indexer.NewHandler(topicHash(ENDURANCE_PRIZE_EVENT), "EnduranceChampionPrizePaid", game, h.decodeEnduranceChampionPrizePaid, h.storeEnduranceChampionPrizePaid),
		indexer.NewHandler(topicHash(LASTCST_BIDDER_PRIZE_EVENT), "LastCstBidderPrizePaid", game, h.decodeLastCstBidderPrizePaid, h.storeLastCstBidderPrizePaid),
		indexer.NewHandler(topicHash(CHRONO_WARRIOR_PRIZE_EVENT), "ChronoWarriorPrizePaid", game, h.decodeChronoWarriorPrizePaid, h.storeChronoWarriorPrizePaid),
		indexer.NewHandler(topicHash(DONATED_TOKEN_CLAIMED), "DonatedTokenClaimed", prizes, h.decodeDonatedTokenClaimed, h.storeDonatedTokenClaimed),
		indexer.NewHandler(topicHash(DONATED_NFT_CLAIMED), "DonatedNftClaimed", prizes, h.decodeDonatedNftClaimed, h.storeDonatedNftClaimed),
		indexer.NewHandler(topicHash(TRANSFER_EVT), "Transfer", signature, h.decodeCosmicSignatureTransfer, h.storeCosmicSignatureTransfer),
		indexer.NewHandler(topicHash(TRANSFER_EVT), "Transfer", one(h.c.Token), h.decodeCosmicTokenTransfer, h.storeCosmicTokenTransfer),
		indexer.NewHandler(topicHash(CST_NFT_STAKED_EVENT), "NftStakedCST", stakingEither, h.decodeNftStakedCST, h.storeNftStakedCST),
		indexer.NewHandler(topicHash(RWALK_NFT_STAKED_EVENT), "NftStakedRWalk", stakingEither, h.decodeNftStakedRWalk, h.storeNftStakedRWalk),
		indexer.NewHandler(topicHash(NFT_UNSTAKED_RWALK), "NftUnstakedRWalk", one(h.c.StakingRWalk), h.decodeNftUnstakedRWalk, h.storeNftUnstakedRWalk),
		indexer.NewHandler(topicHash(NFT_UNSTAKED_CST), "NftUnstakedCST", one(h.c.StakingCST), h.decodeNftUnstakedCST, h.storeNftUnstakedCST),
		indexer.NewHandler(topicHash(STAKING_ETH_DEPOSIT_EVENT), "EthDepositReceived", one(h.c.StakingCST), h.decodeStakingEthDeposit, h.storeStakingEthDeposit),
		indexer.NewHandler(topicHash(MARKETING_REWARD_PAID), "RewardPaid", marketing, h.decodeMarketingRewardPaid, h.storeMarketingRewardPaid),
		indexer.NewHandler(topicHash(CHARITY_PERCENTAGE_CHANGED), "CharityEthDonationAmountPercentageChanged", game, h.decodeCharityPercentageChanged, h.storeCharityPercentageChanged),
		indexer.NewHandler(topicHash(PRIZE_PERCENTAGE_CHANGED), "MainEthPrizeAmountPercentageChanged", game, h.decodePrizePercentageChanged, h.storePrizePercentageChanged),
		indexer.NewHandler(topicHash(RAFFLE_PERCENTAGE_CHANGED), "RaffleTotalEthPrizeAmountForBiddersPercentageChanged", game, h.decodeRafflePercentageChanged, h.storeRafflePercentageChanged),
		indexer.NewHandler(topicHash(STAKE_PERCENTAGE_CHANGED), "CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged", game, h.decodeStakingPercentageChanged, h.storeStakingPercentageChanged),
		indexer.NewHandler(topicHash(CHRONO_PERCENTAGE_CHANGED), "ChronoWarriorEthPrizeAmountPercentageChanged", game, h.decodeChronoPercentageChanged, h.storeChronoPercentageChanged),
		indexer.NewHandler(topicHash(NUM_RAFFLE_ETH_PRIZE_EVENTS_BIDDING_CHANGED), "NumRaffleEthPrizesForBiddersChanged", game, h.decodeNumRaffleETHWinnersBiddingChanged, h.storeNumRaffleETHWinnersBiddingChanged),
		indexer.NewHandler(topicHash(NUM_RAFFLE_NFT_PRIZE_EVENTS_BIDDING_CHANGED), "NumRaffleCosmicSignatureNftsForBiddersChanged", game, h.decodeNumRaffleNFTWinnersBiddingChanged, h.storeNumRaffleNFTWinnersBiddingChanged),
		indexer.NewHandler(topicHash(NUM_RAFFLE_NFT_PRIZE_EVENTS_STAKING_RWALK_CHANGED), "NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged", game, h.decodeNumRaffleNFTWinnersStakingRWalkChanged, h.storeNumRaffleNFTWinnersStakingRWalkChanged),
		indexer.NewHandler(topicHash(RWALK_ADDRESS_CHANGED), "RandomWalkNftAddressChanged", game, h.decodeRandomWalkAddressChanged, h.storeRandomWalkAddressChanged),
		indexer.NewHandler(topicHash(PRIZE_WALLET_ADDRESS_CHANGED), "PrizesWalletAddressChanged", game, h.decodePrizesWalletAddressChanged, h.storePrizesWalletAddressChanged),
		indexer.NewHandler(topicHash(STAKING_WALLET_CST_ADDRESS_CHANGED), "StakingWalletCosmicSignatureNftAddressChanged", game, h.decodeStakingWalletCSTAddressChanged, h.storeStakingWalletCSTAddressChanged),
		indexer.NewHandler(topicHash(STAKING_WALLET_RWALK_ADDRESS_CHANGED), "StakingWalletRandomWalkNftAddressChanged", game, h.decodeStakingWalletRWalkAddressChanged, h.storeStakingWalletRWalkAddressChanged),
		indexer.NewHandler(topicHash(MARKETING_ADDRESS_CHANGED), "MarketingWalletAddressChanged", game, h.decodeMarketingWalletAddressChanged, h.storeMarketingWalletAddressChanged),
		indexer.NewHandler(topicHash(TREASURER_CHANGED), "TreasurerAddressChanged", marketing, h.decodeTreasurerAddressChanged, h.storeTreasurerAddressChanged),
		indexer.NewHandler(topicHash(COSMIC_TOKEN_ADDRESS_CHANGED), "CosmicSignatureTokenAddressChanged", game, h.decodeCosmicTokenAddressChanged, h.storeCosmicTokenAddressChanged),
		indexer.NewHandler(topicHash(COSMIC_SIGNATURE_ADDRESS_CHANGED), "CosmicSignatureNftAddressChanged", game, h.decodeCosmicSignatureAddressChanged, h.storeCosmicSignatureAddressChanged),
		indexer.NewHandler(topicHash(PROXY_UPGRADED), "Upgraded", game, h.decodeUpgraded, h.storeUpgraded),
		indexer.NewHandler(topicHash(ADMIN_CHANGED), "AdminChanged", game, h.decodeAdminChanged, h.storeAdminChanged),
		indexer.NewHandler(topicHash(TIME_INCREASE_CHANGED), "MainPrizeTimeIncrementIncreaseDivisorChanged", game, h.decodeTimeIncreaseChanged, h.storeTimeIncreaseChanged),
		indexer.NewHandler(topicHash(TIMEOUT_CLAIMPRIZE_CHANGED), "TimeoutDurationToClaimMainPrizeChanged", game, h.decodeTimeoutClaimPrizeChanged, h.storeTimeoutClaimPrizeChanged),
		indexer.NewHandler(topicHash(TIMEOUT_TO_WITHDRAW_PRIZE), "TimeoutDurationToWithdrawPrizesChanged", prizes, h.decodeTimeoutToWithdrawPrizesChanged, h.storeTimeoutToWithdrawPrizesChanged),
		indexer.NewHandler(topicHash(PRICE_INCREASE_CHANGED), "EthBidPriceIncreaseDivisorChanged", game, h.decodePriceIncreaseChanged, h.storePriceIncreaseChanged),
		indexer.NewHandler(topicHash(MAIN_PRIZE_MICROSECOND_INCREASE), "MainPrizeTimeIncrementInMicroSecondsChanged", game, h.decodeMainPrizeMicrosecondsChanged, h.storeMainPrizeMicrosecondsChanged),
		indexer.NewHandler(topicHash(INITIAL_SECONDS_UNTIL_PRIZE_CHANGED), "InitialDurationUntilMainPrizeDivisorChanged", game, h.decodeInitialSecondsUntilPrizeChanged, h.storeInitialSecondsUntilPrizeChanged),
		indexer.NewHandler(topicHash(ROUND_ACTIVATION_TIME_CHANGED), "RoundActivationTimeChanged", game, h.decodeRoundActivationTimeChanged, h.storeRoundActivationTimeChanged),
		indexer.NewHandler(topicHash(CST_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED), "CstDutchAuctionDurationDivisorChanged", game, h.decodeCstDutchAuctionDurationDivisorChanged, h.storeCstAuctionLengthChange("CstDutchAuctionDurationDivisorChanged")),
		indexer.NewHandler(topicHash(CST_DUTCH_AUCTION_DURATION_CHANGED), "CstDutchAuctionDurationChanged", game, h.decodeCstDutchAuctionDurationChanged, h.storeCstAuctionLengthChange("CstDutchAuctionDurationChanged")),
		indexer.NewHandler(topicHash(CST_DUTCH_AUCTION_DURATION_CHANGE_DIVISOR_CHANGED), "CstDutchAuctionDurationChangeDivisorChanged", game, h.decodeCstAuctionDurationChangeDivisorChanged, h.storeCstAuctionDurationChangeDivisorChanged),
		indexer.NewHandler(topicHash(ETH_DUTCH_AUCTION_DURATION_DIVISOR_CHANGED), "EthDutchAuctionDurationDivisorChanged", game, h.decodeEthAuctionDurationDivisorChanged, h.storeEthAuctionDurationDivisorChanged),
		indexer.NewHandler(topicHash(ETH_DUTCH_AUCTION_ENDING_PRICE_DIVISOR_CHANGED), "EthDutchAuctionEndingBidPriceDivisorChanged", game, h.decodeEthAuctionEndingBidPriceDivisorChanged, h.storeEthAuctionEndingBidPriceDivisorChanged),
		indexer.NewHandler(topicHash(CST_REWARD_FOR_BIDDING_CHANGED), "CstRewardAmountForBiddingChanged", game, h.decodeCstRewardAmountForBiddingChanged, h.storeCstRewardForBiddingChange("CstRewardAmountForBiddingChanged")),
		indexer.NewHandler(topicHash(BID_CST_REWARD_AMOUNT_CHANGED), "BidCstRewardAmountChanged", game, h.decodeBidCstRewardAmountChanged, h.storeCstRewardForBiddingChange("BidCstRewardAmountChanged")),
		indexer.NewHandler(topicHash(BID_CST_REWARD_AMOUNT_MULTIPLIER_CHANGED), "BidCstRewardAmountMultiplierChanged", game, h.decodeBidCstRewardAmountMultiplierChanged, h.storeCstRewardForBiddingChange("BidCstRewardAmountMultiplierChanged")),
		indexer.NewHandler(topicHash(STATIC_CST_REWARD), "CstPrizeAmountChanged", game, h.decodeStaticCstRewardChanged, h.storeStaticCstRewardChanged),
		indexer.NewHandler(topicHash(MAX_MESSAGE_LENGTH), "BidMessageLengthMaxLimitChanged", game, h.decodeMaxMessageLengthChanged, h.storeMaxMessageLengthChanged),
		indexer.NewHandler(topicHash(TOKEN_SCRIPT_URL), "NftGenerationScriptUriChanged", signature, h.decodeNftGenerationScriptURLChanged, h.storeNftGenerationScriptURLChanged),
		indexer.NewHandler(topicHash(BASE_URI), "NftBaseUriChanged", signature, h.decodeNftBaseURIChanged, h.storeNftBaseURIChanged),
		indexer.NewHandler(topicHash(MARKETING_REWARD_CHANGED), "MarketingWalletCstContributionAmountChanged", game, h.decodeMarketingRewardChanged, h.storeMarketingRewardChanged),
		indexer.NewHandler(topicHash(OWNERSHIP_TRANSFERRED), "OwnershipTransferred", h.ownershipSources(), h.decodeOwnershipTransferred, h.storeOwnershipTransferred),
		indexer.NewHandler(topicHash(INITIALIZED), "Initialized", h.initializedSources(), h.decodeInitialized, h.storeInitialized),
		indexer.NewHandler(topicHash(STARTING_CST_MIN_LIM), "CstDutchAuctionBeginningBidPriceMinLimitChanged", game, h.decodeCstMinLimitChanged, h.storeCstMinLimitChanged),
		indexer.NewHandler(topicHash(FUND_TRANSFER_ERR), "FundTransferFailed", game, h.decodeFundTransferFailed, h.storeFundTransferFailed),
		indexer.NewHandler(topicHash(ERC20_TRANSFER_ERR), "ERC20TransferFailed", game, h.decodeERC20TransferFailed, h.storeERC20TransferFailed),
		indexer.NewHandler(topicHash(FUNDS_TO_CHARITY), "FundsTransferredToCharity", marketing, h.decodeFundsToCharity, h.storeFundsToCharity),
		indexer.NewHandler(topicHash(DELAY_DURATION_ROUND), "DelayDurationBeforeRoundActivationChanged", game, h.decodeDelayDurationChanged, h.storeDelayDurationChanged),
		indexer.NewHandler(topicHash(FIRST_BID_EVENT), "FirstBidPlacedInRound", game, h.decodeFirstBidPlacedInRound, h.storeFirstBidPlacedInRound),
	}
}
