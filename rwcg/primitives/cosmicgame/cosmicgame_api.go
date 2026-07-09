package cosmicgame

// Transaction contains common transaction/event fields
type Transaction struct {
	EvtLogId  int64
	BlockNum  int64
	TxId      int64
	TxHash    string
	TimeStamp int64
	DateTime  string
}

type CGStatistics struct {
	TotalBids					uint64
	CurNumBids					uint64
	TotalPrizes					uint64   // main prize claims (one per round won)
	TotalPrizeAwards				uint64   // SUM(cg_winner.prizes_count); excludes cg_prize rows without winner attribution (e.g. ptype 15 staking)
	CgPrizeRowCount				uint64   // COUNT(*) FROM cg_prize — canonical row count in unified prize table
	NumUniqueBidders			uint64
	NumUniqueWinners			uint64
	NumUniqueStakersCST			uint64
	NumUniqueStakersRWalk		uint64
	NumUniqueStakersBoth		uint64
	TotalPrizesPaidAmountWei	string
	TotalPrizesPaidAmountEth	float64	// divided by 1e18
	NumUniqueDonors				int64
	TotalEthDonatedAmount		string
	TotalEthDonatedAmountEth	float64
	NumVoluntaryDonations		uint64
	SumVoluntaryDonationsEth	float64 // divided by 1e18
	NumCosmicGameDonations		uint64
	SumCosmicGameDonationsEth	float64 // donations from CosmicGame contract
	DirectDonationsEth			float64
	NumDirectDonations			int64
	SumWithdrawals				float64 // total ETH withdrwan from CharityWallet
	NumWithdrawals				uint64	// number of withdrawal operations made
	NumRwalkTokensUsed			uint64
	NumDonatedNFTs				uint64
	NumCSTokenMints				uint64
	TotalNamedTokens			int64
	TotalRaffleEthDeposits		float64
	TotalRaffleEthWithdrawn		float64
	TotalChronoWarriorEthDeposits float64
	TotalCSTGivenInPrizes		float64
	NumWinnersWithPendingRaffleWithdrawal int64
	TotalNFTDonated				int64
	TotalCSTConsumed			string
	TotalCSTConsumedEth			float64
	NumBidsCST					int64
	TotalMktRewards				string	// rewards deposited to marketers (total), from MarketingWallet
	TotalMktRewardsEth			float64
	NumMktRewards				int64
	DonatedTokenDistribution	[]CGDonatedTokenDistrRec 
	StakeStatisticsCST			CGStakeStatsCST
	StakeStatisticsRWalk		CGStakeStatsRWalk
}
type CGBidRec struct {
	Tx							Transaction
	BidderAid					int64
	BidderAddr					string
	EthPrice					string
	EthPriceEth					float64	// divided by 1e18 (or -1 if CST bid)
	CstPrice					string
	CstPriceEth					float64	// divided by 1e18 (or -1 if ETH bid)
	RWalkNFTId					int64
	RoundNum					int64
	BidType						int64
	BidPosition					int64
	PrizeTime					int64
	PrizeTimeDate				string
	TimeUntilPrize				int64	// Seconds until prize (0 if already ended)
	CSTReward				string
	CSTRewardEth				float64
	// V3 bid CST reward 90/10 split (Comment-202607161). PreviousBidder* is the 90% share minted to the
	// outbid (previous last) bidder; ThisBidder* is the ~10% share minted to the bidder placing this bid.
	// For V2/V1 bids the whole reward is the ThisBidder share and PreviousBidder* is 0.
	PreviousBidderCstRewardAmount	string
	PreviousBidderCstRewardAmountEth	float64
	ThisBidderCstRewardAmount		string
	ThisBidderCstRewardAmountEth	float64
	// DEPRECATED (kept for frontend back-compat until the V3 upgrade; remove afterwards). The TOTAL bid CST
	// reward from IBiddingV2 BidPlaced ("-1" = legacy V1). Same value as CSTReward; equals PreviousBidder + ThisBidder.
	BidCstRewardAmount			string
	BidCstRewardAmountEth		float64
	CstDutchAuctionDuration		string	// per-bid auction duration from IBiddingV2 BidPlaced; "-1" = legacy
	CstDutchAuctionDurationInt	int64	// numeric duration when >= 0; else -1
	NFTDonationTokenId			int64
	NFTDonationTokenAddr		string
	NFTTokenURI					string
	ImageURL					string
	Message						string
	DonatedERC20TokenAddr				string
	DonatedERC20TokenAmount		string
	DonatedERC20TokenAmountEth	float64
}

// CGBannedBidRec is one row from cg_banned_bids (API: get_banned_bids)
type CGBannedBidRec struct {
	Id        int64  `json:"id"`
	BidId     int64  `json:"bid_id"`
	UserAddr  string `json:"user_addr"`
	CreatedAt int64  `json:"created_at"`
}

type CGClaimPrizeTx struct {
	Tx							Transaction
}
type CGMainPrizeInfo struct {
	WinnerAid					int64
	WinnerAddr					string
	TimeoutTs					int64
	EthAmount					string
	EthAmountEth				float64
	CstAmount					string
	CstAmountEth				float64
	NftTokenId					uint64	// V3: first of NumCSNfts sequential NFT IDs awarded to the winner
	NumCSNfts					int64	// number of Cosmic Signature NFTs awarded (V2 = 1, V3 default 3)
	NftTokenIds					[]int64	// all main-prize NFT IDs (NftTokenId .. NftTokenId+NumCSNfts-1); empty if no claim yet
	Seed						string
}
type CGCharityDeposit struct {
	CharityAddress				string
	CharityAmount				string
	CharityAmountETH			float64
}
type CGStakingDeposit struct {
	StakingDepositId			int64
	StakingDepositAmount		string
	StakingDepositAmountEth		float64
	StakingPerToken				string
	StakingPerTokenEth			float64
	StakingNumStakedTokens		int64
}
type CGEnduranceChampionPrize struct {
	WinnerAddr					string
	NftTokenId					int64
	CstAmount					string
	CstAmountEth				float64
}
type CGLastCSTBidderPrize struct {
	WinnerAddr					string
	NftTokenId					int64
	CstAmount					string
	CstAmountEth				float64
}
type CGChronoWarriorPrize struct {
	WinnerAddr					string
	EthAmount					string
	EthAmountEth				float64
	CstAmount					string
	CstAmountEth				float64
	NftTokenId					int64
}
type CGRoundRec struct {
	RoundNum					uint64
	ClaimPrizeTx				CGClaimPrizeTx
	MainPrize					CGMainPrizeInfo
	CharityDeposit				CGCharityDeposit
	StakingDeposit				CGStakingDeposit
	EnduranceChampion			CGEnduranceChampionPrize
	LastCstBidder				CGLastCSTBidderPrize
	ChronoWarrior				CGChronoWarriorPrize
	RoundStats					CGRoundStats
	RaffleNFTWinners			[]CGRaffleNFTWinnerRec
	StakingNFTWinners			[]CGRaffleNFTWinnerRec
	RaffleETHDeposits			[]CGPrizeDepositRec
	AllPrizes					[]CGPrizeHistory
}
type CGUserInfo struct {
	AddressId					int64
	Address						string
	NumPrizes					int64
	NumBids						int64
	MaxWinAmount				float64
	MaxBidAmount				float64
	SumRaffleEthWinnings		float64
	SumRaffleEthWithdrawal		float64
	NumRaffleEthWinnings		int64
	RaffleNFTsCount				int64
	RewardNFTsCount				int64	// RaffleNftsCount + endurance count + chrono + lastcstbidder count + main prize NFT
	UnclaimedNFTs				int64
	TotalCSTokensWon			int64	// prizes + raffles (ERC721 NFT count, not CST!)
	CosmicSignatureNumTransfers	int64
	TotalDonatedCount			int64
	TotalDonatedAmountEth		float64
	StakingStatisticsRWalk		CGStakeStatsRWalk
}
type CGCharityDonation struct {
	Tx							Transaction
	DonorAid					int64
	DonorAddr					string
	Amount						string
	AmountEth					float64
	IsVoluntary					bool	// true - made by direct send, false=made by CosmicGame contract
	RoundNum					int64
}
type CGCosmicGameDonationSimple struct {
	Tx							Transaction
	DonorAid					int64
	DonorAddr					string
	Amount						string
	AmountEth					float64
	RoundNum					int64
}
type CGCosmicGameDonationWithInfo struct {
	Tx							Transaction
	DonorAid					int64
	DonorAddr					string
	Amount						string
	AmountEth					float64
	RoundNum					int64
	CGRecordId					int64	// CosmicGame contract's record id
	DataJson					string
}
type CGDonationCombinedRec struct {
	RecordType					int64	// 0 - simple donation, 1 - donation with info
	Tx							Transaction
	DonorAid					int64
	DonorAddr					string
	Amount						string
	AmountEth					float64
	RoundNum					int64
	CGRecordId					int64	// CosmicGame contract's record id
	DataJson					string
}
type CGUniqueBidder struct {
	BidderAid					int64
	BidderAddr					string
	NumBids						int64
	MaxBidAmount				string
	MaxBidAmountEth				float64	// same as above but with 18 decimal places (i.e. in ETH )
}
type CGWinnerStats struct {
	MaxWinAmount				string
	MaxWinAmountEth				float64
	PrizesCount					int64
	PrizesSum					string
	PrizesSumEth				float64
	TokensCount					int64
	ERC20Count					int64
	ERC721Count					int64
	UnclaimedNfts				int64
	TotalSpent					string
	TotalSpentEth				float64
}
type CGUniqueWinner struct {
	WinnerAid					int64
	WinnerAddr					string
	PrizesCount					int64
	MaxWinAmount				string
	MaxWinAmountEth				float64	// same as above but with 18 decimal places (i.e. in ETH )
	PrizesSum					float64	// all winnings in ETH
	WinnerStats					CGWinnerStats
}
type CGUniqueDonor struct {
	DonorAid					int64
	DonorAddr					string
	CountDonations				int64
	TotalDonated				string
	TotalDonatedEth				float64
}
type CGRoiLeaderboardEntry struct {
	BidderAid					int64
	BidderAddr					string
	NumBids						int64
	RoundsParticipated			int64	// distinct rounds the player bid in
	RoundsWon					int64	// distinct rounds the player won any prize
	WinRate						float64	// RoundsWon / RoundsParticipated (0..1)
	TotalEthSpent				string	// wei
	TotalEthSpentEth			float64
	TotalCstSpent				string	// wei
	TotalCstSpentEth			float64
	EthWon						string	// cg_winner.prizes_sum (main+raffle+chrono ETH), wei
	EthWonEth					float64
	PrizesCount					int64	// all prize types
	CstPrizesCount				int64	// count of CST (ERC20) prizes
	NftPrizesCount				int64	// count of CS NFT (ERC721) prizes
	NetPlEth					float64	// (EthWon - TotalEthSpent) in ETH
	Roi							float64	// fraction; multiply by 100 for percent. 0 when no ETH spent
}
// A single not-yet-claimed claimable asset held in PrizesWallet (for the per-cycle dialog).
type CGClaimUnclaimedItem struct {
	AssetType					string	// "ETH" | "ERC721" | "ERC20"
	RecipientAddr				string	// who is entitled to claim it
	AmountEth					float64	// ETH prize, or ERC20 token amount /1e18; 0 for ERC721
	TokenAddr					string	// contract address for ERC721 / ERC20; "" for ETH
	TokenId						int64	// token id for ERC721; -1 otherwise
}
// Per-cycle summary of claimable assets awarded via PrizesWallet and their claim status.
type CGRoundClaimSummary struct {
	RoundNum					int64
	ClaimWindowTimeout			int64	// unix ts after which unclaimed assets can be swept by anyone
	AwardedTs					int64	// unix ts when the cycle finalized (assets became claimable)
	Expired						bool	// now >= ClaimWindowTimeout
	EthAwarded					int64
	EthUnclaimed				int64
	EthUnclaimedEth				float64
	NftAwarded					int64
	NftUnclaimed				int64
	Erc20Awarded				int64
	Erc20Unclaimed				int64
	TotalAwarded				int64
	TotalUnclaimed				int64
	AvgClaimPeriodSecs			int64	// avg seconds from cycle finalize to claim (over claimed assets)
	UnclaimedItems				[]CGClaimUnclaimedItem
}
// A single claim transaction (a recipient withdrawing a claimable asset from PrizesWallet).
type CGClaimTxn struct {
	AssetType					string	// "ETH" | "ERC721" | "ERC20"
	RecipientAddr				string	// entitled recipient
	BeneficiaryAddr				string	// who actually claimed (ETH: can differ from recipient after expiry)
	AmountEth					float64	// ETH, or ERC20 amount /1e18; 0 for ERC721
	TokenAddr					string
	TokenId						int64	// ERC721 token id; -1 otherwise
	ClaimedAfterSecs			int64	// seconds from cycle finalize to this claim
	ClaimTs						int64	// unix ts of the claim
	TxHash						string
}
// A token attached (donated) during a cycle, held in PrizesWallet for the recipient to claim.
type CGAttachedToken struct {
	AssetType					string	// "ERC721" | "ERC20"
	ContributorAddr				string	// who attached it
	TokenAddr					string
	TokenId						int64	// ERC721 token id; -1 otherwise
	AmountEth					float64	// ERC20 amount /1e18; 0 for ERC721
	Ts							int64	// unix ts when attached
	TxHash						string
}
// Per-cycle claim drill-down: the claim transactions (with latency) and the tokens attached that cycle.
type CGRoundClaimDetail struct {
	RoundNum					int64
	ClaimTransactions			[]CGClaimTxn
	AttachedTokens				[]CGAttachedToken
}
type CGERC20Donation struct {
	RecordId					int64
	Tx							Transaction
	RoundNum					int64
	DonorAid					int64
	DonorAddr					string
	TokenAid					int64
	TokenAddr					string
	Amount						string
	AmountEth					float64
	WinnerAid					int64
	WinnerAddr					string
//	Claimed						bool
}
type CGSummarizedERC20Donation struct {
	RecordId					int64
	Tx							Transaction
	RoundNum					int64
	TokenAid					int64
	TokenAddr					string
	AmountDonated				string
	AmountDonatedEth			float64
	AmountClaimed				string
	AmountClaimedEth			float64
	DonateClaimDiff				string
	DonateClaimDiffEth			float64
	WinnerAid					int64
	WinnerAddr					string
	Claimed						bool
}
type CGNFTDonation struct {
	RecordId					int64
	Tx							Transaction
	RoundNum					int64
	DonorAid					int64
	DonorAddr					string
	TokenAddressId				int64	// the 'aid' of TokenAddr
	TokenAddr					string
	NFTTokenId					int64
	NFTTokenURI					string
	Index						int64
}
type CGNFTDonationStats struct {
	TokenAddressId				int64
	TokenAddress				string
	NumDonations				int64	// total number of donated tokens per this contract
}
type CGRecordCounters struct {
	TotalBids					int64
	TotalPrizes					int64
	TotalDonatedNFTs			int64
}
type CGPrizeDepositRec struct {
	RecordId					int64
	Tx							Transaction
	RecordType					int64		// 0 - undefined ; 1 - Raffle Deposit ; 2 - Chronor Warrior
	WinnerAddr					string
	WinnerAid					int64
	WinnerIndex					int64
	RoundNum					int64
	Amount						float64
	Claimed						bool
	ClaimTimeStamp				int64
	ClaimDateTime				string
}
type CGRaffleNFTWinnerRec struct {
	RecordId					int64
	Tx							Transaction
	WinnerAddr					string
	WinnerAid					int64
	RoundNum					int64
	TokenId						int64
	CstAmount					string
	CstAmountEth				float64
	WinnerIndex					int64
	IsRWalk						bool
	IsStaker					bool
}
type CGDonatedNFTClaimRec struct {
	RecordId					int64
	Tx							Transaction
	RoundNum					int64
	Index						int64
	TokenAddr					string
	NFTTokenId					int64
	NFTTokenURI					string
	WinnerIndex					int64
	WinnerAid					int64
	WinnerAddr					string
	DonorAddr					string
}
type CGERC20ClaimRec struct {
	RecordId					int64
	Tx							Transaction
	RoundNum					int64
	Index						int64
	TokenAid					int64
	TokenAddr					string
	Amount						string
	AmountEth					float64
	WinnerAid					int64
	WinnerAddr					string
	DonorAddr					string
}
type CGCosmicSignatureMintRec struct {
	RecordId					int64
	Tx							Transaction
	ContractAddr				string
	TokenId						int64
	WinnerAid					int64
	WinnerAddr					string
	CurOwnerAid					int64
	CurOwnerAddr				string
	Seed						string
	RoundNum					int64
	RecordType					int64	// 0 - undefined 1 Raffle NFT, 2 - Staking Rwalk, 3 - Main Prize, 4 - Endurance Champ, 5 - Last CST Bidder
	TokenName					string
	Staked						bool
	StakedOwnerAid				int64
	StakedOwnerAddr				string
	StakeActionId				int64
	StakeTimeStamp				int64
	StakeDateTime				string
	UnstakeActionId				int64
	WasUnstaked					bool
	ActualUnstakeTimeStamp		int64	// if there is unstake record, these fields hold dates
	ActualUnstakeDateTime		string
}
type CGRoundStats struct {
	RoundNum					int64
	TotalBids					int64
	TotalDonatedNFTs			int64
	NumERC20Donations			int64
	TotalRaffleEthDeposits		string
	TotalRaffleEthDepositsEth	float64 // deposits of ETH (same as above) but divided by 1^18
	TotalRaffleNFTs				int64
	TotalDonatedCount			int64
	TotalDonatedAmount			string
	TotalDonatedAmountEth		float64
	TotalCstInBidsEth			float64	// CST consumed in gestures during this cycle (cg_round_stats.total_cst_in_bids /1e18)
	TotalEthInBidsEth			float64	// ETH wagered in gestures during this cycle (cg_round_stats.total_eth_in_bids /1e18)
	// Round timing fields (added 2025-11-06)
	ParamWindowStartTime		string	// ISO 8601 format
	ActivationTime				int64	// Unix seconds (0 = not set); matches contract roundActivationTime()
	// Seconds; contract delayDurationBeforeRoundActivation() (global config, not per-round DB column)
	DelayDurationBeforeRoundActivation	int64
	ParamWindowDurationSeconds	int64
	RoundStartTime				string	// ISO 8601 format
	RoundEndTime				string	// ISO 8601 format
	RoundDurationSeconds		int64
}
type CGClaimInfo struct {
	ETHRaffleToClaim			float64
	ETHRaffleToClaimWei			string
	ETHChronoWarriorToClaim		float64
	ETHChronoWarriorToClaimWei	string
	NumDonatedNFTToClaim		int64		// Pending unclaimed donated tokens (counter)
	UnclaimedStakingReward		float64
	DonatedERC20Tokens			[]ERC20DonatedTokensInfo
}
type CGPrizeHistory struct {
	Tx							Transaction
	RecordType					int64		// 0-ETH raffle, 1-CS NFT raffle, 2-Donated NFT, 3-Main Prize, 4 - StakingDeposit (at StakingWallet CST), 5 CST Mint for RandomWalk staker , 6 CST Mint for CST staker, 7 - Endurance NFT winner, 8 - LastCst Bid NFT winner, 9 - Endurance ERC20 winner, 10 - LastCst Bid ERC20 winner , 11 - Donated ERC20 token , 12 - Chrono Warrior, 16 - Donated NFT (timeout), 17 - Donated ERC20 (timeout), 18 - Raffle ETH (timeout)
	RoundNum					int64
	Amount						string
	AmountEth					float64
	WinnerIndex					int64
	TokenAddress				string
	TokenId						int64
	TokenURI					string
	Claimed						bool
	IsTimeoutClaim				bool		// True if prize was claimed after timeout by non-winner
	WinnerAddr					string
	WinnerAid					int64
}
type CGTokenName struct {
	Tx							Transaction
	TokenId						int64
	TokenName					string
	ChangedByAid				int64
	ChangedBy					string
}
type CGTransfer struct {
	RecordId					int64
	Tx							Transaction
	TokenId						int64
	FromAddr					string
	ToAddr						string
	FromAid						int64
	ToAid						int64
	TransferType				int64 // 0 - regular transfer , 1 - mint, 2 - burn (there are no burns in CST)
}
type CGCharityWithdrawal struct {
	RecordId					int64
	Tx							Transaction
	DestinationAddr				string
	Amount						string
	AmountEth					float64
}
type CGTokenSearchResult struct {
	MintTimeStamp				int64
	MintDateTime				string
	TokenId						int64
	TokenName					string
}
type CGDonatedTokenDistrRec struct {
	ContractAddr				string
	NumDonatedTokens			int64
}
type CGCSTokenDistributionRec struct {
	OwnerAid					int64
	OwnerAddr					string
	NumTokens					int64
}
type CGCosmicTokenHolderRec struct {
	OwnerAid					int64
	OwnerAddr					string
	Balance						string
	BalanceFloat				float64
	PercentOfSupply				float64
}
// CGTotalSupplyHistoryByDateRec is one calendar day's aggregated CST supply change from bids.
type CGTotalSupplyHistoryByDateRec struct {
	Date			string	// YYYYMMDD
	TimeStamp		int64	// epoch at start of day (UTC)
	DateTime		string
	NumBids			int64
	MintAmount		string
	MintAmountEth	float64
	BurnAmount		string
	BurnAmountEth	float64
	Amount			string	// net mint minus burn for the day
	AmountEth		float64
	TotalSupply		string
	TotalSupplyEth	float64
}
// CGTotalSupplyHistoryRec is one bid's net CST supply change (mint minus burn on that bid).
type CGTotalSupplyHistoryRec struct {
	Tx				Transaction
	BidInfoId		int64	// evtlog_id for /bid/info/:evtlog_id
	BidType			int64	// 0 = ETH, 1 = RandomWalk, 2 = CST
	BidderAddr		string
	MintAmount		string	// cst_reward minted for this bid
	MintAmountEth	float64
	BurnAmount		string	// cst_price burned (0 for ETH / RandomWalk bids)
	BurnAmountEth	float64
	Amount			string	// net: MintAmount - BurnAmount (wei string)
	AmountEth		float64
	TotalSupply		string
	TotalSupplyEth	float64
}
type CGCosmicTokenStats struct {
	// Supply metrics
	TotalSupply					string
	TotalSupplyEth				float64
	TotalHolders				int64
	
	// How CST (ERC20) enters the game (sources)
	EarnedFromBidding			string
	EarnedFromBiddingEth		float64
	DistributedToMarketers		string
	DistributedToMarketersEth	float64
	GivenAsMainPrizes			string
	GivenAsMainPrizesEth		float64
	GivenAsRafflePrizes			string
	GivenAsRafflePrizesEth		float64
	GivenAsChronoWarriorPrizes	string
	GivenAsChronoWarriorPrizesEth	float64
	
	// How CST leaves the game (burns)
	ConsumedInBids				string
	ConsumedInBidsEth			float64
	
	// Activity metrics
	TotalMints					int64
	TotalBurns					int64
	TotalTransfers				int64
	
	// Top holders
	TopHolders					[]CGCosmicTokenHolderRec
}
type CGUserCosmicTokenSummary struct {
	UserAddr					string
	UserAid						int64
	CurrentBalance				string
	CurrentBalanceEth			float64
	
	// CST (ERC20) Earnings breakdown
	TotalEarned					string
	TotalEarnedEth				float64
	EarnedFromBidding			string
	EarnedFromBiddingEth		float64
	EarnedFromMainPrizes		string
	EarnedFromMainPrizesEth		float64
	EarnedFromRafflePrizes		string
	EarnedFromRafflePrizesEth	float64
	EarnedFromChronoWarrior		string
	EarnedFromChronoWarriorEth	float64
	EarnedFromMarketing			string
	EarnedFromMarketingEth		float64
	
	// CST (ERC20) Spending
	ConsumedInBids				string
	ConsumedInBidsEth			float64
	
	// Net CST flow
	NetCSTFlow					string
	NetCSTFlowEth				float64
	
	// Activity
	NumTransfers				int64
	NumMints					int64
	NumBurns					int64
}
type CGERC20TransferRec struct {
	RecordId					int64
	Tx							Transaction
	FromAddr					string
	ToAddr						string
	FromAid						int64
	ToAid						int64
	TransferType				int64 // 0 - regular transfer , 1 - mint, 2 - burn (there are no burns in CST)
	Value						string
	ValueFloat					float64
}
type CGRWalkUsed struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	RoundNum					int64
	BidderAid					int64
	BidderAddr					string
	RWalkTokenId				int64
}
type CGMarketingRewardRec struct {
	RecordId					int64
	Tx							Transaction
	Amount						string
	AmountEth					float64
	MarketerAid					int64
	MarketerAddr				string
}
type CGSystemModeRec struct {
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	NextEvtLogId				int64
	RoundNum					int64
}
type CGAdminEvent struct {
	RecordType					int64	// Type codes:
										// 			0		Undefined
										//			1		CharityPercentageChanged
										//			2		PrizePercentageChanged
										//			3		RafflePercentageChanged
										//			4		StakingPercentageChanged
										//			5		numRaffleETHWinnersBidding
										//			6		numRaffleNFTWinnersBidding
										//			7		DelayDurationBeforeRoundActivationChanged
										//			8		NumRaffleNFTWinnersStakingRWalkChanged
										//			9		CharityAddressChanged
										//			10		RandomWalkAddressChanged
										//			11		PrizeWalletAddressChanged
										//			12		StakingWalletCSTAddressChanged
										//			13		StakingWalletRWalkAddressChanged
										//			14		MarketingWalletAddressChanged
										//			15		CosmicTokenAddressChanged
										//			16		CosmicSignatureAddressChanged
										//			17		Upgraded
										//			18		TimeIncreaseChanged
										//			19		TimeoutClaimPrizeChanged
										//			20		PriceIncreaseChanged
										//			21		NanoSecondsExtraChanged
										//			22		InitialSecondsUntilPrizeChanged
										//			23		TreasurerAddressChanged
										//			24		ActivationTimeChanged
										//			25		CstDutchAuctionDurationDivisorChanged / CstDutchAuctionDurationChanged (V2)
										//			26		Erc20RewardMultiplierChanged
										//			27		StartingBidPriceCSTMinLimitChanged
										//			28		MarketingRewardChanged
										//			29		CstRewardAmountForBiddingChanged / BidCstRewardAmountChanged / BidCstRewardAmountMultiplierChanged (V2)
										//			30		MaxMessageLengthChanged
										//			31		TokenGenerationScriptURLEvent
										//			32		BaseURI (CosmicSignature)
										//			33		Initialized (Initialized event, openzeppelin)
										//			34		OwnershipTransferred
										//			35		TimeoutDurationToWithdrawPrizesChanged
										//			36		EthDutchAuctionDurationDivisorChanged
										//			37		EthDutchAuctionEndingBidPriceDivisorChanged
										//			38		ChronoWarriorEthPrizeAmountPercentageChanged
										//			39		CstDutchAuctionDurationChangeDivisorChanged (V2)
										//			40		RoundLateBidDurationDivisorChanged (V3)
										//			41		RoundLateBidPricePremiumAmountBaseMultiplierChanged (V3)
										//			42		RoundLateBidPricePremiumAmountExponentChanged (V3)
										//			43		LastBidderBidCstRewardAmountPercentageChanged (V3)
										//			44		MainPrizeNumCosmicSignatureNftsChanged (V3)
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	AddressValue				string
	IntegerValue				int64
	FloatValue					float64
	StringValue					string
	ResolvedValue				string	// Human-readable value when IntegerValue is a divisor (or direct unit conversion)
}
type ERC20DonatedTokensInfo struct {
	RoundNum					int64
	TokenAid					int64
	TokenAddr					string
	Amount						string
	AmountEth					float64
}

// CGBidFrequencyBucket is bid count in a fixed time bucket.
type CGBidFrequencyBucket struct {
	BucketTs      int64
	NumBids       int64
	UniqueBidders int64
}

// CGBidTypeRatioBucket holds the bid-type composition of a single sampling
// window. Counts are the raw number of bids of each type that fell within the
// window [BucketTs, BucketTs+interval); the *Pct fields are those counts
// normalized to 100% of TotalBids (windowed, not cumulative). When a window has
// no bids, TotalBids is 0 and all *Pct fields are 0.
// bid_type mapping: 0=ETH, 1=RandomWalk (ETH-paid), 2=CST.
type CGBidTypeRatioBucket struct {
	BucketTs   int64
	EthBids    int64
	RwalkBids  int64
	CstBids    int64
	TotalBids  int64
	EthPct     float64
	RwalkPct   float64
	CstPct     float64
}

// CGBidSpike is a merged run of above-normal bid frequency buckets.
type CGBidSpike struct {
	Index       int
	StartTs     int64
	EndTs       int64
	PeakTs      int64
	PeakNumBids int64
	TotalBids   int64
	BucketCount int64
}

// CGTopBidderInfo ranks bidders by lifetime gesture count.
type CGTopBidderInfo struct {
	BidderAid  int64
	BidderAddr string
	NumBids    int64
}

// CGBidderActivePeriod is a contiguous burst of bids by one address (gap-separated).
type CGBidderActivePeriod struct {
	BidderAid    int64
	BidderAddr   string
	PeriodStart  int64
	PeriodEnd    int64
	NumBids      int64
	DurationSecs int64
}
