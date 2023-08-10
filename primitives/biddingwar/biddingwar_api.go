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
	NumCSTokenMints				uint64
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
	RoundNum					int64
	ERC20_Amount				string
	ERC20_AmountEth				float64	// divided by 1e18
	NFTDonationTokenId			int64
	NFTDonationTokenAddr		string
	NFTTokenURI					string
	ImageURL					string
	Message						string
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
	CharityAddress				string
	CharityAmount				string
	CharityAmountETH			float64
	RoundStats					BwRoundStats
	RaffleNFTWinners			[]BwRaffleNFTWinnerRec
	RaffleETHDeposits			[]BwRaffleDepositRec
}
type BwUserInfo struct {
	AddressId					int64
	Address						string
	NumPrizes					int64
	NumBids						int64
	MaxWinAmount				float64
	MaxBidAmount				float64
	SumRaffleEthWinnings		float64
	NumRaffleEthWinnings		int64
	RaffleNFTWon				int64
	RaffleNFTClaimed			int64
	UnclaimedNFTs				int64
	TotalCSTokensWon			int64	// prizes + raffles
}
type BwCharityDonation struct {
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
type BwBiddingwarDonation struct {
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
type BwRaffleDepositRec struct {
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
	DepositId					int64
	Amount						float64
}
type BwRaffleWithdrawalRec struct {
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
type BwRaffleNFTWinnerRec struct {
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
	WinnerIndex					int64
	ClaimTimestamp				int64
	ClaimDateTime				string
	ClaimTokenId				int64
}
type BwRaffleNFTClaimRec struct {
	RecordId					int64
	EvtLogId					int64
	BlockNum					int64
	TxId						int64
	TxHash						string
	TimeStamp					int64
	DateTime					string
	WinnerAddr					string
	WinnerAid					int64
	TokenId						int64
	WinningRoundNum				int64
	WinningTimestamp			int64
	WinningDateTime				string
	WinningIndex				int64
}
type BwDonatedNFTClaimRec struct {
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
	TokenId						string
	WinnerIndex					int64
	WinnerAid					int64
	WinnerAddr					string
}
type BwCosmicSignatureMintRec struct {
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
	MintType					int64
	PrizeNum					int64	// -1 if minted via Raffle , > -1 if MintType = 0
	ClaimTimestamp				int64
	ClaimDateTime				string
}
type BwRoundStats struct {
	RoundNum					int64
	TotalBids					int64
	TotalDonatedNFTs			int64
	TotalRaffleEthDeposits		string
	TotalRaffleEthDepositsEth	float64 // deposits of ETH (same as above) but divided by 1^18
	TotalRaffleNFTs				int64
}
type BwClaimInfo struct {
	ETHRaffleToClaim			float64
	ETHRaffleToClaimWei			string
	NumCSNFTRaffleToClaim		int64		// CosmicSignature NFT tokens to claim (counter)
	NumDonatedNFTToClaim		int64		// Pending unclaimed donated tokens (counter)
}
