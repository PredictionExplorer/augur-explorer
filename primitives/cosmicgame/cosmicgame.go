package cosmicgame

import (

)

type CosmicGameContractAddrs struct {
	CosmicGameAddr				string
	CosmicSignatureAddr			string
	CosmicTokenAddr				string
	CosmicDaoAddr				string
	CharityWalletAddr			string
	RaffleWalletAddr			string
	RandomWalkAddr				string
	StakingWalletCSTAddr			string
	StakingWalletRWalkAddr			string
	MarketingWalletAddr			string
	ImplementationAddr			string
}
type CosmicGameProcStatus struct {
	LastEvtIdProcessed			int64
}
type CGPrizeClaimEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	PrizeNum			int64
	TokenId				int64
	WinnerAddr			string
	Amount				string
	DonationEvtId		int64
}
type CGBidEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	RandomWalkTokenId	int64
	PrizeTime			int64
	RoundNum			int64
	BidType				int64
	NumCSTTokens		string
	ContractAddr		string
	LastBidderAddr		string
	BidPrice			string
	ERC20_Value			string
	Message				string
}
type CGDonationEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	DonorAddr			string
	RoundNum			int64
	Amount				string
}
type CGDonationWithInfoEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	DonorAddr			string
	RoundNum			int64
	RecordId			int64	// record ID at the contract side
	Amount				string
}
type CGDonationReceivedEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	RoundNum			int64
	ContractAddr		string
	DonorAddr			string
	Amount				string
}
type CGDonationSentEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	CharityAddr			string
	Amount				string
}
type CGNFTDonationEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	TokenAddr			string
	RoundNum			int64
	DonorAddr			string
	TokenId				int64
	Index				int64
	BidId				int64	// id of related bid record
	NFTTokenURI			string
}
type CGCharityUpdatedEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	NewCharityAddr		string
}
type CGTokenNameEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	TokenId				int64
	TokenName			string
}
type CGMintEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	TokenId				int64
	OwnerAddr			string
	Seed				string
	RoundNum			int64
	MintType			int64
}
type CGRaffleDeposit struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	Amount				string
}
type CGRaffleWithdrawal struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	WinnerAddr			string
	Amount				string
}
type CGRaffleNFTWinner struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	WinnerIndex			int64
	TokenId				int64
	IsRandomWalk		bool
	IsStaker			bool
}
type CGEnduranceWinner struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	WinnerIndex			int64
	Erc721TokenId		int64
	Erc20Amount			string
}
type CGStellarWinner struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	WinnerIndex			int64
	Erc721TokenId		int64
	Erc20Amount			string
}
/*
type CGRaffleNFTClaimed struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	WinnerAddr			string
	WinnerEvtlogId		int64
	TokenId				int64
}
*/
type CGDonatedNFTClaimed struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	RoundNum			int64
	Index				int64
	TokenAddr			string
	TokenId				string
	WinnerAddr			string
}
type CGStakeActionCST struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	RoundNum			int64
	ActionId			int64
	TokenId				int64
	TotalNfts			int64
	Staker 				string
}
type CGUnstakeActionCST struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	RoundNum			int64
	ActionId			int64
	TokenId				int64
	TotalNfts			int64
	Reward				string
	Staker 				string
}
type CGEthDeposit struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	RoundNum			int64
	DepositTime			int64
	DepositId			int64
	DepositNum			int64
	NumStakedNfts		int64
	Amount				string
	AmountPerStaker		string
	AccumModulo			string
	Modulo				string
}
type CGClaimReward struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	ActionId			int64
	DepositId			int64
	Reward				string	
	Staker				string
}
type CGStakeActionRWalk struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	RoundNum			int64
	ActionId			int64
	TokenId				int64
	TotalNfts			int64
	Staker 				string
}
type CGUnstakeActionRWalk struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	RoundNum			int64
	ActionId			int64
	TokenId				int64
	TotalNfts			int64
	Staker 				string
}
type CGMarketingRewardSent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	ActionId			int64
	DepositId			int64
	Amount				string	
	Marketer			string
}
type CGERC721Transfer struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	From					string
	To						string
	TokenId					int64
}
type CGERC20Transfer struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	From					string
	To						string
	Value					string
}
type CGCharityPercentageChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewCharityPercentage	string
}
type CGPrizePercentageChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewPrizePercentage	string
}
type CGRafflePercentageChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewRafflePercentage		string
}
type CGStakingPercentageChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewStakingPercentage	string
}
type CGNumRaffleETHWinnersBiddingChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewNumRaffleETHWinnersBidding	int64
}
type CGNumRaffleNFTWinnersBiddingChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewNumRaffleNFTWinnersBidding int64
}
type CGNumRaffleNFTWinnersStakingCSTChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewNumRaffleNFTWinnersStakingCST int64
}
type CGNumRaffleNFTWinnersStakingRWalkChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewNumRaffleNFTWinnersStakingRWalk int64
}
type CGSystemModeChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewSystemMode			int64
}
type CGCharityAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewCharity				string
}
type CGRandomWalkAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewRandomWalk				string
}
type CGRaffleWalletAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewRaffleWallet				string
}
type CGStakingWalletCSTAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewStakingWalletCST				string
}
type CGStakingWalletRWalkAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewStakingWalletRWalk				string
}
type CGMarketingWalletAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewMarketingWallet				string
}
type CGCosmicTokenAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewCosmicToken			string
}
type CGCosmicSignatureAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewCosmicSignature		string
}
type CGUpgraded struct {	// openzeppelin proxy Upgraded event
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Implementation			string
}
type CGTimeIncreaseChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewTimeIncrease			string
}
type CGTimeoutClaimPrizeChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewTimeout				int64
}
type CGPriceIncreaseChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewPriceIncrease		string
}
type CGNanoSecondsExtraChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewNanoSecondsExtra		string
}
type CGInitialSecondsUntilPrizeChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewInitialSecondsUntilPrize	string
}
type CGInitialBidAmountFractionChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewInitialBidAmountFraction	string
}
type CGActivationTimeChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewActivationTime		string
}
type CGETHCSTBidRatioChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewETHToCSTBidRatio		string
}
type CGRoundStartCSTAuctionLengthChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewAuctionLength		string
}
type CGERC20RewardMultiplierChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewMultiplier			string
}
type CGStartingBidPriceCSTMinLimitChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewPrice				string
}
type CGMarketingRewardChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewReward				string
}
type CGERC20TokenRewardChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewReward				string
}
type CGMaxMessageLengthChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewMessageLength		string
}
type CGTokenGenerationScriptURL struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewURL					string
}
type CGBaseURIEvent struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewURI					string
}
