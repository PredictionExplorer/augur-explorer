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
}
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
