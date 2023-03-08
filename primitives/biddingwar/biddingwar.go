package biddingwar

import (

)

type BiddingWarContractAddrs struct {
	BiddingWarAddr				string
	CosmicSignatureAddr			string
	CosmicSignatureTokenAddr	string
	CharityWalletAddr			string

}
type BWPrizeClaimEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	PrizeNum			int64
	WinnerAddr			string
	Amount				string
}
type BWBidEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	RandomWalkTokenId	int64
	ContractAddr		string
	LastBidderAddr		string
	BidPrice			string
}
type BWDonationEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	DonorAddr			string
	Amount				string
}
type BWDonationReceivedEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	DonorAddr			string
	Amount				string
}
type BWDonationSentEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	CharityAddr			string
	Amount				string
}
type BWCharityUpdatedEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	NewCharityAddr		string
}
type BWTokenNameEvent struct {
	EvtId				int64
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	TokenId				int64
	TokenName			string
}
