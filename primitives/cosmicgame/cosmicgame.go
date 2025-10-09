package cosmicgame

import (

)

type CosmicGameContractAddrs struct {
	CosmicGameAddr				string
	CosmicSignatureAddr			string
	CosmicTokenAddr				string
	CosmicDaoAddr				string
	CharityWalletAddr			string
	PrizesWalletAddr				string
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
	ContractAddr		string
	RoundNum			int64
	TokenId				int64
	WinnerAddr			string
	Timeout				int64
	Amount				string
	DonationEvtId		int64
}
type CGBidEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	RandomWalkTokenId	int64
	PrizeTime			int64
	RoundNum			int64
	BidType				int64
	CstPrice			string
	ContractAddr		string
	LastBidderAddr		string
	BidPrice			string
	ERC20_Value			string	// reward of CST tokens earned for bidding
	Message				string
}
type CGDonationEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
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
	ContractAddr		string
	CharityAddr			string
	Amount				string
}
type CGERC20DonationEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	TokenAddr			string
	RoundNum			int64
	DonorAddr			string
	Amount				string
	BidId				int64	// id of related bid record
}
type CGNFTDonationEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
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
	ContractAddr		string
	NewCharityAddr		string
}
type CGTokenNameEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	TokenId				int64
	TokenName			string
}
type CGMintEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	TokenId				int64
	OwnerAddr			string
	Seed				string
	RoundNum			int64
}
type CGPrizesEthDeposit struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	WinnerIndex			int64
	Amount				string
}
type CGPrizesEthWithdrawal struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	WinnerAddr			string
	Amount				string
}
type CGRaffleNFTWinner struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	WinnerIndex			int64
	TokenId				int64
	IsRandomWalk		bool
	IsStaker			bool
}
type CGRaffleETHWinner struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	WinnerIndex			int64
	Amount 				string
}
type CGEnduranceWinner struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	WinnerIndex			int64
	Erc721TokenId		int64
	Erc20Amount			string
}
type CGLastBidderWinner struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	WinnerIndex			int64
	Erc721TokenId		int64
	Erc20Amount			string
}
type CGChronoWarrior struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	WinnerAddr			string
	Round				int64
	Amount				string
}
type CGDonatedTokenClaimed struct {	// ERC20 tokens
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	RoundNum			int64
	Index				int64
	TokenAddr			string
	Amount				string
	BeneficiaryAddr		string
}
type CGDonatedNFTClaimed struct {	// ERC721 tokens
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	RoundNum			int64
	Index				int64
	TokenAddr			string
	TokenId				string
	BeneficiaryAddr		string
}
type CGNftStakedCst struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	ActionId			int64
	NftId				int64
	NumStakedNfts		int64
	StakerAddress		string
	RewardPerStaker		string
}
type CGNftStakedRWalk struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	ActionId			int64
	NftId				int64
	NumStakedNfts		int64
	StakerAddress		string
}
type CGEthDeposit struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	RoundNum			int64
	DepositTime			int64
	DepositId			int64
	NumStakedNfts		int64
	Amount				string
	AmountPerStaker		string
	AccumModulo			string
	Modulo				string
}
type CGNftUnstakedRWalk struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	RoundNum			int64
	ActionId			int64
	NftId				int64
	NumStakedNfts		int64
	StakerAddress		string
}
type CGNftUnstakedCst struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	ContractAddr		string
	RoundNum			int64
	ActionId			int64
	ActionCounter		int64
	NftId				int64
	NumStakedNfts		int64
	StakerAddress		string
	RewardAmount		string
	RewardPerToken		string
}
type CGMarketingRewardSent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
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
type CGChronoPercentageChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewChronoPercentage		string
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
type CGPrizeWalletAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewPrizeWallet			string
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
type CGTreasurerAddressChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewTreasurer			string
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
type CGAdminChanged struct {	// openzeppelin proxy AdminChanged event
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	OldAdmin				string
	NewAdmin				string
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
type CGTimeoutToWithdrawPrizeChanged  struct {
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
type CGMainPrizeMicroSecondsIncreaseChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewMicroseconds			string
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
type CGCstDutchAuctionDurationDivisorChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewValue				string
}
type CGEthDutchAuctionDurationDivisorChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewValue				string
}
type CGEthDutchAuctionEndingBidPriceDivisorChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewValue				string
}
type CGERC20RewardMultiplierChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewMultiplier			string
}
type CGMarketingRewardChanged struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewReward				string
}
type CGCstRewardForBiddingChanged struct {
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
type CGOwnershipTransferred struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	PrevOwner				string
	NewOwner				string
	ContractCode			int64
}
type CGInitialized struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Version					int64
}
type CGCstMinLimit struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	CstMinLimit				string
}
type CGFundTransferFailed struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Destination				string
	Amount					string
}
type CGErc20TransferFailed struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	Destination				string
	Amount					string
}
type CGFundsToCharity struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	CharityAddr				string
	Amount					string
}
type CGNextRoundDelayDuration struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	NewValue				int64
}
type CGRoundStarted	struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	RoundNum				int64
	StartTimestamp			int64
}
