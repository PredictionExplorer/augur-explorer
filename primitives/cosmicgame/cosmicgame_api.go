package cosmicgame

type CGStatistics struct {
	TotalBids					uint64
	CurNumBids					uint64
	TotalPrizes					uint64
	NumUniqueBidders			uint64
	NumUniqueWinners			uint64
	NumUniqueStakers			uint64
	TotalPrizesPaidAmountWei	string
	TotalPrizesPaidAmountEth	float64	// divided by 1e18
	NumVoluntaryDonations		uint64
	SumVoluntaryDonationsEth	float64 // divided by 1e18
	NumCosmicGameDonations		uint64
	SumCosmicGameDonationsEth	float64 // donations from CosmicGame contract
	SumWithdrawals				float64 // total ETH withdrwan from CharityWallet
	NumWithdrawals				uint64	// number of withdrawal operations made
	NumRwalkTokensUsed			uint64
	NumDonatedNFTs				uint64
	NumCSTokenMints				uint64
	TotalNamedTokens			int64
	TotalRaffleEthDeposits		float64
	TotalRaffleEthWithdrawn		float64
	NumWinnersWithPendingRaffleWithdrawal int64
	TotalNFTDonated				int64
	TotalCSTConsumed			string
	TotalCSTConsumedEth			float64
	NumBidsCST					int64
	TotalMktRewards				string	// rewards deposited to marketers (total), from MarketingWallet
	TotalMktRewardsEth			float64
	NumMktRewards				int64
	DonatedTokenDistribution	[]CGDonatedTokenDistrRec 
	StakeStatistics				CGStakeStats
}
type CGStakeStats struct {
	TotalTokensStaked			int64
	TotalReward					string
	TotalRewardEth				float64
	UnclaimedReward				string
	UnclaimedRewardEth			float64
	NumActiveStakers			int64
	NumDeposits					int64
}
type CGBidRec struct {
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	BidderAid					int64
	BidderAddr					string
	BidPrice					string
	BidPriceEth					float64	// divided by 1e18
	RWalkNFTId					int64
	RoundNum					int64
	BidType						int64
	NumCSTTokens				string
	NumCSTTokensEth				float64
	ERC20_Amount				string
	ERC20_AmountEth				float64	// divided by 1e18
	NFTDonationTokenId			int64
	NFTDonationTokenAddr		string
	NFTTokenURI					string
	ImageURL					string
	Message						string
}
type CGPrizeRec struct {
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	WinnerAid					int64
	WinnerAddr					string
	Amount						string
	AmountEth					float64	// divided by 1e18
	PrizeNum					uint64
	TokenId						uint64
	Seed						string
	CharityAddress				string
	CharityAmount				string
	CharityAmountETH			float64
	StakingDepositNum			int64
	StakingDepositAmount		string
	StakingDepositAmountEth		float64
	StakingPerToken				string
	StakingPerTokenEth			float64
	RoundStats					CGRoundStats
	RaffleNFTWinners			[]CGRaffleNFTWinnerRec
	RaffleETHDeposits			[]CGRaffleDepositRec
}
type UserStakingInfo struct {
	TotalTokensStaked			int64
	TotalNumStakeActions		int64
	TotalReward					string
	TotalRewardEth				float64
	UnclaimedReward				string
	UnclaimedRewardEth			float64
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
	RaffleNFTWon				int64
	RaffleNFTClaimed			int64
	UnclaimedNFTs				int64
	TotalCSTokensWon			int64	// prizes + raffles
	CosmicTokenNumTransfers		int64
	CosmicSignatureNumTransfers	int64
	StakingStatistics		UserStakingInfo
}
type CGCharityDonation struct {
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DonorAid					int64
	DonorAddr					string
	Amount						string
	AmountEth					float64
	IsVoluntary					bool	// true - made by direct send, false=made by CosmicGame contract
	RoundNum					int64
}
type CGCosmicGameDonation struct {
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DonorAid					int64
	DonorAddr					string
	Amount						string
	AmountEth					float64
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
type CGUniqueStaker struct {
	StakerAid					int64
	StakerAddr					string
	TotalTokensStaked			int64
	NumStakeActions				int64
	NumUnstakeActions			int64
	TotalReward					string
	TotalRewardEth				float64
	UnclaimedReward				string
	UnclaimedRewardEth			float64
}
type CGNFTDonation struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
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
type CGRaffleDepositRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	WinnerAddr					string
	WinnerAid					int64
	RoundNum					int64
	Amount						float64
	Claimed						bool
	ClaimTimeStamp				int64
	ClaimDateTime				string
}
type CGRaffleWithdrawalRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	WinnerAddr					string
	WinnerAid					int64
	Amount						float64
}
type CGRaffleNFTWinnerRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	WinnerAddr					string
	WinnerAid					int64
	RoundNum					int64
	TokenId						int64
	WinnerIndex					int64
}
type CGDonatedNFTClaimRec struct {
	RecordId					int64
	EvtId						int64
	BlockNum					int64
	TimeStamp					int64
	DateTime					string
	TxId						int64
	TxHash						string
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
type CGCosmicSignatureMintRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TimeStamp					int64
	DateTime					string
	TxId						int64
	TxHash						string
	LogIndex					int64
	ContractAddr				string
	TokenId						int64
	WinnerAid					int64
	WinnerAddr					string
	CurOwnerAid					int64
	CurOwnerAddr				string
	Seed						string
	RoundNum					int64
	RecordType					int64
	TokenName					string
	Staked						bool
	StakedOwnerAid				int64
	StakedOwnerAddr				string
	StakeActionId				int64
	UnstakeElligibleTimeStamp	int64
	UnstakeElligibleDateTime	string
	WasUnstaked					bool
	ActualUnstakeTimeStamp		int64	// if there is unstake record, these fields hold dates
	ActualUnstakeDateTime		string
}
type CGRoundStats struct {
	RoundNum					int64
	TotalBids					int64
	TotalDonatedNFTs			int64
	TotalRaffleEthDeposits		string
	TotalRaffleEthDepositsEth	float64 // deposits of ETH (same as above) but divided by 1^18
	TotalRaffleNFTs				int64
}
type CGClaimInfo struct {
	ETHRaffleToClaim			float64
	ETHRaffleToClaimWei			string
	NumDonatedNFTToClaim		int64		// Pending unclaimed donated tokens (counter)
	UnclaimedStakingReward		float64
}
type CGRaffleHistory struct {
	EvtLogId					int64
	RecordType					int64		// 0-ETH raffle, 1-CS NFT raffle, 2-Donated NFT, 3-Main Prize, 4 - SatkingDeposit (at StakingWallet)
	TimeStamp					int64
	DateTime					string
	BlockNum					int64
	TxId						int64
	TxHash						string
	RoundNum					int64
	Amount						string
	AmountEth					float64
	WinnerIndex					int64
	TokenAddress				string
	TokenId						int64
	TokenURI					string
	Claimed						bool
	WinnerAddr					string
	WinnerAid					int64
}
type CGTokenName struct {
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	TokenId						int64
	TokenName					string
}
type CGTransfer struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	TokenId						int64
	FromAddr					string
	ToAddr						string
	FromAid						int64
	ToAid						int64
	TransferType				int64 // 0 - regular transfer , 1 - mint, 2 - burn (there are no burns in CST)
}
type CGCharityWithdrawal struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
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
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
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
type CGEthDepositRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositDate					string
	DepositTimeStamp			int64
	NumStakedNFTs				int64
	Amount						string
	AmountEth					float64
	AmountPerHolder				string
	AmountPerHolderEth			float64
	Modulo						string
	ModuloF64					float64
}
type CGStakeActionRec struct {
	ActionType					int64
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	UnstakeDate					string
	UnstakeTimeStamp			int64
	ActionId					int64
	TokenId						int64
	NumStakedNFTs				int64
	AmountPerHolder				string
	AmountPerHolderEth			float64
	Modulo						string
	ModuloF64					float64
	IsRandomWalk				bool
	Claimed						bool
}
type CGStakingHistoryRec struct {
	ActionType					int64
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	UnstakeEvtLogId				int64
	UnstakeBlockNum				int64
	UnstakeDate					string
	UnstakeTimeStamp			int64
	ActionId					int64
	TokenId						int64
	NumStakedNFTs				int64
	AccumNumStakedNFTs			int64
	AmountPerHolder				string
	AmountPerHolderEth			float64
	Modulo						string
	ModuloF64					float64
	Claimed						bool
	StakerAid					int64
	StakerAddr					string
	LastBlockTS					int64
	UnstakeExpirationDiff		int64
	IsRandomWalk				bool
}
type CGStakeActionInfoRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	ActionId					int64
	TokenId						int64
	NumStakedNFTs				int64
	UnstakeTimeStamp			int64
	UnstakeDate					string
	StakerAid					int64
	IsRandomWalk				bool
	StakerAddr					string
}
type CGUnstakeActionInfoRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	ActionId					int64
	TokenId						int64
	NumStakedNFTs				int64
	StakerAid					int64
	StakerAddr					string
}
type CGStakeUnstakeCombined struct {
	Stake						CGStakeActionInfoRec
	Unstake						CGUnstakeActionInfoRec
}
type CGRewardToClaim struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositDate					string
	DepositTimeStamp			int64
	DepositId					int64
	NumStakedNFTs				int64
	DepositAmount				string
	DepositAmountEth			float64
	YourTokensStaked			int64
	YourClaimableAmount			string
	YourClaimableAmountEth		float64
	AmountPerToken				string
	AmountPerTokenEth			float64
	Modulo						string
	ModuloF64					float64
}
type CGCollectedReward struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositDate					string
	DepositTimeStamp			int64
	DepositId					int64
	NumStakedNFTs				int64
	TotalDepositAmount			string
	TotalDepositAmountEth		float64
	YourTokensStaked			int64
	YourAmountToClaim			string
	YourAmountToClaimEth		float64
	DepositAmountPerToken		string
	DepositAmountPerTokenEth	float64
	NumTokensCollected			int64
	YourCollectedAmount			string
	YourCollectedAmountEth		float64
	Modulo						string
	ModuloF64					float64
	RoundNum					int64
	StakerAid					int64
	StakerAddr					string
	FullyClaimed				bool
}
type CGMarketingRewardRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	Amount						string
	AmountEth					float64
	MarketerAid					int64
	MarketerAddr				string
}
type CGStakedTokenRec struct {
	TokenInfo					CGCosmicSignatureMintRec
	StakeEvtLogId				int64
	StakeBlockNum				int64
	StakeActionId				int64
	StakeTimeStamp				int64
	StakeDateTime				string
	StakedIsRandomWalk			bool
	UnstakeTimeStamp			int64
	UnstakeDateTime				string
	UserAddr					string
	UserAid						int64

}
type CGActionIdsForDeposit struct {
	RecordId					int64
	DepositId					int64
	UserAid						int64
	StakeActionId				int64
	TokenId						int64
	Claimed						bool
	StakeActionTimeStamp		int64
	UnstakeEligibleTimeStamp	int64
	CurChainTimeStamp			int64
	TimeStampDiff				int64	// subtraction of UnstakeEligibleTimestamp from CurChainTimestamp
	Amount						string
	AmountEth					float64
}
type CGActionIdsForDepositWithClaimInfo struct {
	RecordId					int64
	DepositId					int64
	UserAid						int64
	StakeActionId				int64
	TokenId						int64
	Claimed						bool
	ClaimBlockNum				int64
	ClaimTimeStamp				int64
	ClaimDateTime				string
	ClaimTxId					int64
	ClaimTxHash					string
	ClaimRewardAmount			string
	ClaimRewardAmountEth		float64
}
type CGStakingRewardRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	RoundNum					int64
	Amount						string
	AmountEth					float64
	StakerAid					int64
	StakerAddr					string
}
type CGEthDepositAsReward struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DepositDate					string
	DepositTimeStamp			int64
	NumStakedNFTsTotal			int64
	Amount						string
	AmountEth					float64
	AmountPerToken				string
	AmountPerTokenEth			float64
	StakerAid					int64
	StakerAddr					string
	StakerAmount				string
	StakerAmountEth				float64
	StakerNumStakedNFTs			int64
}
type CGSystemModeRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	SystemMode					int64
}
