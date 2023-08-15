package biddingwar

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
type BiddingWarProcStatus struct {
	LastEvtIdProcessed			int64
}
type BWPrizeClaimEvent struct {
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
type BWBidEvent struct {
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
type BWDonationEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	DonorAddr			string
	Amount				string
}
type BWDonationReceivedEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	DonorAddr			string
	Amount				string
}
type BWDonationSentEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	CharityAddr			string
	Amount				string
}
type BWNFTDonationEvent struct {
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
type BWCharityUpdatedEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	NewCharityAddr		string
}
type BWTokenNameEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	TokenId				int64
	TokenName			string
}
type BWMintEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	TokenId				int64
	OwnerAddr			string
	Seed				string
	MintType			int64
}
type BWRaffleDeposit struct {
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
type BWRaffleWithdrawal struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxId				int64
	LogIndex			int64
	ContractAddr		string
	WinnerAddr			string
	Amount				string
}
type BWRaffleNFTWinner struct {
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
type BWRaffleNFTClaimed struct {
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
type BWDonatedNFTClaimed struct {
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
type BWERC721Transfer struct {
	EvtId                   int64
	BlockNum                int64
	TxId                    int64
	TimeStamp               int64
	Contract                string
	From					string
	To						string
	TokenId					int64
}
