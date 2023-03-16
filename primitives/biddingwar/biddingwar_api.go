package biddingwar;

type BWStatistics struct {
	TotalBids					uint64
	CurNumBids					uint64
	TotalPrizes					uint64
	NumUniqueBidders			uint64
	NumUniqueWinners			uint64
	TotalPrizesPaidAmountWei	string
	TotalPrizesPaidAmountEth	float64
	NumVoluntaryDonations		uint64
	SumVoluntaryDonationsEth	float64
	NumRwalkTokensUsed			uint64
}
