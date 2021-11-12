package primitives

import (
	//"math/big"

	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"

)
type RW_API_Offer struct {
	Id				int64
	EvtLogId		int64
	BlockNum		int64
	TxId			int64
	TxHash			string
	TimeStamp		int64
	DateTime		string
	ContractAid		int64
	ContractAddr	string	// market addr
	OfferId			int64
	OfferType		int
	SellerAid		int64
	SellerAddr		string
	BuyerAid		int64
	BuyerAddr		string
	TokenId			int64
	Active			bool
	Price			float64
	Profit			float64
	RWalkAid		int64
	RWalkAddr		string
	WasCanceled		bool
}
type RW_API_TokenMint struct {
	TokenId			int64
	BlockNum		int64
	TimeStamp		int64
	DateTime		string
	ContractAid		int64
	ContractAddr	string
	MinterAid		int64
	MinterAddr		string
	Seed			string
	SeedNum			string
	Price			float64
	TxHash			string
}
type RW_API_UserToken struct {
	TokenId			int64
	ContractAid		int64
	ContractAddr	string
	Seed			string
	SeedNum			string
	Price			float64
}
type RW_API_RWalkStats struct {
	TradingVol			float64
	NumTrades			int64
	TokensMinted		int64
	UniqueUsers			int64
	NumWithdrawals		int64
	LastMintedPrice		float64
	MaximumTradedPrice	float64
	CurOwnerAid			int64
	CurOwnerAddr		string
}
type RW_API_MarketStats struct {
	TradingVol			float64
	NumTrades			int64
}
type RW_API_HistEntry_Mint struct { // type = 1
	BlockNum		int64
	TimeStamp		int64
	DateTime		string
	ContractAid		int64
	ContractAddr	string
	TokenId			int64
	OwnerAid		int64
	OwnerAddr		string
	SeedHex			string
	SeedNum			string
	Price			float64
}
type RW_API_HistEntry_Offer struct { // type = 2
	BlockNum		int64
	TimeStamp		int64
	DateTime		string
	ContractAid		int64
	ContractAddr	string
	TokenId			int64
	BuyerAid		int64
	BuyerAddr		string
	SellerAid		int64
	SellerAddr		string
	OfferType		int
	OfferId			int64
	Active			bool
	Price			float64
}
type RW_API_HistEntry_OfferCanceled struct { // type = 3
	BlockNum		int64
	TimeStamp		int64
	DateTime		string
	ContractAid		int64
	ContractAddr	string
	TokenId			int64
	OfferCanceledId	int64
	BuyerAid		int64
	BuyerAddr		string
	SellerAid		int64
	SellerAddr		string
	OfferType		int
	OfferId			int64
	Price			float64
	Aid				int64
	Address			string
}
type RW_API_HistEntry_ItemBought struct { // type = 4
	BlockNum		int64
	TimeStamp		int64
	DateTime		string
	ContractAid		int64
	ContractAddr	string
	TokenId			int64
	ItemBoughtId	int64
	BuyerAid		int64
	BuyerAddr		string
	SellerAid		int64
	SellerAddr		string
	OfferType		int
	OfferId			int64
	Price			float64
	Aid				int64
	Address			string
}
type RW_API_HistEntry_TokenName struct {	// type = 5
	BlockNum		int64
	TimeStamp		int64
	DateTime		string
	ContractAid		int64
	ContractAddr	string
	TokenId			int64
	TokenName		string
}
type RW_API_HistEntry_Transfer struct {	// type = 6
	BlockNum		int64
	TimeStamp		int64
	DateTime		string
	ContractAid		int64
	ContractAddr	string
	TokenId			int64
	FromAid			int64
	FromAddr		string
	ToAid			int64
	ToAddr			string
	TransferId		int64
}
type RW_API_FullHistoryEntry struct {
	RecordType			int		// Mint,or any other event starting with RW_API_HistEntry_*
	Record				interface{}
}
type RW_API_RandomWalkVolumeHistory struct {
	StartTs					int64
	NumOperations			int64
	Volume					float64
	VolumeAccum				float64
}
type RW_API_TokenName struct {
	BlockNum		int64
	TimeStamp		int64
	DateTime		string
	ContractAid		int64
	ContractAddr	string
	TokenId			int64
	TokenName		string
	TxHash			string
	OwnerAid		int64
	OwnerAddr		string
}
type RW_API_UserInfo struct {
	UserAid					int64
	UserAddr				string
	TotalVolume				float64
	TotalNumTrades			int64
	TotalMintedToks			int64
	TotalNumWithdrawals		int64
	IsMarketPlaceContract	bool	// true if the User is not really a user, but is a marketplace contract
}
type RW_API_Top5Toks struct {
	TokenId				int64
	TotalTrades			int64
	SeedHex				string
}
type RW_API_TokenInfo struct {
	TokenId				int64
	CurOwnerAid			int64
	CurOwnerAddr		string
	SeedHex				string
	SeedNum				string
	LastPrice			float64
	TotalVolume			float64
	NumTrades			int64
	CurName				string
	LastNameUpdateTs	int64
	LastNameUpdateDate	string
}
