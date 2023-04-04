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
	NumDonatedNFTs				uint64
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
	RWalkNFTId					int64
	ERC20_Amount				string
	ERC20_AmountEth				float64	// divided by 1e18
	NFTDonationTokenId			int64
	NFTDonationTokenAddr		string
	NFTTokenURI					string
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
type BwUniqueBidder struct {
	BidderAid					int64
	BidderAddr					string
	NumBids						int64
	MaxBidAmount				string
	MaxBidAmountEth				float64	// same as above but with 18 decimal places (i.e. in ETH )
}
type BwUniqueWinner struct {
	WinnerAid					int64
	WinnerAddr					string
	PrizesCount					int64
	MaxWinAmount				string
	MaxWinAmountEth				float64	// same as above but with 18 decimal places (i.e. in ETH )
	PrizesSum					float64	// all winnings in ETH
}
type BwNFTDonation struct {
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	DonorAid					int64
	DonorAddr					string
	TokenAddressId				int64	// the 'aid' of TokenAddr
	TokenAddr					string
	NFTTokenId					int64
	NFTTokenURI					string
}
type BwNFTDonationStats struct {
	TokenAddressId				int64
	TokenAddress				string
	NumDonations				int64	// total number of donated tokens per this contract
}
type BwRecordCounters struct {
	TotalBids					int64
	TotalPrizes					int64
	TotalDonatedNFTs			int64
}
