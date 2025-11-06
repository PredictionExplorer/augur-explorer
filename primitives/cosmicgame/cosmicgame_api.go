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
	TotalPrizes					uint64
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
	NFTDonationTokenId			int64
	NFTDonationTokenAddr		string
	NFTTokenURI					string
	ImageURL					string
	Message						string
	DonatedERC20TokenAddr				string
	DonatedERC20TokenAmount		string
	DonatedERC20TokenAmountEth	float64
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
	NftTokenId					uint64
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
	TotalCSTokensWon			int64	// prizes + raffles
	CosmicTokenNumTransfers		int64
	CosmicSignatureNumTransfers	int64
	TotalDonatedCount			int64
	TotalDonatedAmountEth		float64
	StakingStatistics		UserStakingInfo
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
type CGUniqueWinner struct {
	WinnerAid					int64
	WinnerAddr					string
	PrizesCount					int64
	MaxWinAmount				string
	MaxWinAmountEth				float64	// same as above but with 18 decimal places (i.e. in ETH )
	PrizesSum					float64	// all winnings in ETH
}
type CGUniqueDonor struct {
	DonorAid					int64
	DonorAddr					string
	CountDonations				int64
	TotalDonated				string
	TotalDonatedEth				float64
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
	// Round timing fields (added 2025-11-06)
	ParamWindowStartTime		string	// ISO 8601 format
	ActivationTime				string	// ISO 8601 format
	ParamWindowDurationSeconds	int64
	RoundStartTime				string	// ISO 8601 format
	RoundEndTime				string	// ISO 8601 format
	RoundDurationSeconds		int64
}
type CGClaimInfo struct {
	ETHRaffleToClaim			float64
	ETHRaffleToClaimWei			string
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
										//			25		RoundStartCSTAuctionLengthChanged
										//			26		Erc20RewardMultiplierChanged
										//			27		StartingBidPriceCSTMinLimitChanged
										//			28		MarketingRewardChanged
										//			29		TokenRewardChanged
										//			30		MaxMessageLengthChanged
										//			31		TokenGenerationScriptURLEvent
										//			32		BaseURI (CosmicSignature)
										//			33		Initialized (Initialized event, openzeppelin)
										//			34		OwnershipTransferred
										//			35		TimeoutDurationToWithdrawPrizesChanged
										//			36		EthDutchAuctionDurationDivisorChanged
										//			37		EthDutchAuctionEndingBidPriceDivisorChanged
										//			38		ChronoWarriorEthPrizeAmountPercentageChanged
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
}
type ERC20DonatedTokensInfo struct {
	RoundNum					int64
	TokenAid					int64
	TokenAddr					string
	Amount						string
	AmountEth					float64
}
