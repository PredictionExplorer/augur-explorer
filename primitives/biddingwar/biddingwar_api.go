package biddingwar;

type BWStatistics struct {
	TotalBids					uint64
	CurNumBids					uint64
	TotalPrizes					uint64
	NumUniqueBidders			uint64
	NumUniqueWinners			uint64
	TotalPrizesPaidAmountWei	string
	TotalPrizesPaidAmountEth	float64	// divided by 1e18
	NumVoluntaryDonations		uint64
	SumVoluntaryDonationsEth	float64 // divided by 1e18
	NumRwalkTokensUsed			uint64
}
type BwBidRec struct {
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
	ERC20_Amount				string
	ERC20_AmountEth				float64	// divided by 1e18
}
type BwPrizeRec struct {
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
}
type BwUserInfo struct {
	AddressId					int64
	Address						string
	NumPrizes					int64
	NumBids						int64
	MaxWinAmount				float64
	MaxBidAmount				float64
}
type BwDonation struct {
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
	IsVoluntary					bool	// true - made by direct send, false=made by BiddingWar contract
}
