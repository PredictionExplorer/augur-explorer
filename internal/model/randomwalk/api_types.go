package randomwalk

// =====================================================================
// API Response Types
// =====================================================================

type Offer struct {
	Id           int64
	EvtLogId     int64
	BlockNum     int64
	TxId         int64
	TxHash       string
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string // market addr
	OfferId      int64
	OfferType    int
	SellerAid    int64
	SellerAddr   string
	BuyerAid     int64
	BuyerAddr    string
	TokenId      int64
	Active       bool
	Price        float64
	Profit       JSONNullFloat64
	RWalkAid     int64
	RWalkAddr    string
	WasCanceled  bool
}

type TokenMint struct {
	TokenId      int64
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	MinterAid    int64
	MinterAddr   string
	Seed         string
	SeedNum      string
	Price        float64
	TxHash       string
}

type TokenMintCSV struct {
	TokenId      int64
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	MinterAid    int64
	MinterAddr   string
	Seed         string
	SeedNum      string
	Price        float64
	TxHash       string
	NumTrades    int64
	TotalVolume  float64
	LastPrice    float64
	LastName     string
	LastOwner    string
}

type UserToken struct {
	TokenId      int64
	ContractAid  int64
	ContractAddr string
	Seed         string
	SeedNum      string
	Price        float64
}

type RWalkStats struct {
	TradingVol         float64
	NumTrades          int64
	TokensMinted       int64
	UniqueUsers        int64
	NumWithdrawals     int64
	LastMintedPrice    float64
	MaximumTradedPrice float64
	CurOwnerAid        int64
	CurOwnerAddr       string
	WithdrawalAmount   float64 // the latest amount
}

type MarketStats struct {
	TradingVol float64
	NumTrades  int64
}

type HistEntryMint struct { // type = 1
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	TokenId      int64
	OwnerAid     int64
	OwnerAddr    string
	SeedHex      string
	SeedNum      string
	Price        float64
}

type HistEntryOffer struct { // type = 2
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	TokenId      int64
	BuyerAid     int64
	BuyerAddr    string
	SellerAid    int64
	SellerAddr   string
	OfferType    int
	OfferId      int64
	Active       bool
	Price        float64
}

type HistEntryOfferCanceled struct { // type = 3
	BlockNum        int64
	TimeStamp       int64
	DateTime        string
	ContractAid     int64
	ContractAddr    string
	TokenId         int64
	OfferCanceledId int64
	BuyerAid        int64
	BuyerAddr       string
	SellerAid       int64
	SellerAddr      string
	OfferType       int
	OfferId         int64
	Price           float64
	Aid             int64
	Address         string
}

type HistEntryItemBought struct { // type = 4
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	TokenId      int64
	ItemBoughtId int64
	BuyerAid     int64
	BuyerAddr    string
	SellerAid    int64
	SellerAddr   string
	OfferType    int
	OfferId      int64
	Price        float64
	Aid          int64
	Address      string
}

type HistEntryTokenName struct { // type = 5
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	TokenId      int64
	TokenName    string
}

type HistEntryTransfer struct { // type = 6
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	TokenId      int64
	FromAid      int64
	FromAddr     string
	ToAid        int64
	ToAddr       string
	TransferId   int64
}

type FullHistoryEntry struct {
	RecordType int // Mint, or any other event starting with HistEntry*
	Record     interface{}
}

type VolumeHistory struct {
	StartTs       int64
	NumOperations int64
	Volume        float64
	VolumeAccum   float64
}

type TokenNameRec struct {
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	TokenId      int64
	TokenName    string
	TxHash       string
	OwnerAid     int64
	OwnerAddr    string
}

type UserInfo struct {
	UserAid               int64
	UserAddr              string
	TotalVolume           float64
	TotalNumTrades        int64
	TotalMintedToks       int64
	TotalNumWithdrawals   int64
	IsMarketPlaceContract bool // true if the User is not really a user, but is a marketplace contract
}

type TopTradedToken struct {
	TokenId     int64
	TotalTrades int64
	SeedHex     string
}

type TokenInfo struct {
	TokenId            int64
	CurOwnerAid        int64
	CurOwnerAddr       string
	SeedHex            string
	SeedNum            string
	LastPrice          float64
	TotalVolume        float64
	NumTrades          int64
	CurName            string
	LastNameUpdateTs   int64
	LastNameUpdateDate string
}

type MintInterval struct {
	MintNumber int64 // sequential number of mint (same as token_id)
	TimeStamp  int64
	Interval   int64 // Duration from previous mint
	TokenId    int64
}

type WithdrawalChartEntry struct {
	TimeStamp        int64
	DateTime         string
	WithdrawalAmount float64
}

type TradingHistoryLog struct {
	Id               int64
	EvtLogId         int64
	BlockNum         int64
	TxId             int64
	TimeStamp        int64
	DateTime         string
	ContractAid      int64
	ContractAddr     string // market addr
	OfferId          int64
	OfferType        int
	SellerAid        int64
	SellerAddr       string
	BuyerAid         int64
	BuyerAddr        string
	TokenId          int64
	Active           bool
	Price            float64
	Profit           JSONNullFloat64
	RWalkAid         int64
	RWalkAddr        string
	WasCanceled      bool
	WasBought        bool
	ItemBoughtTs     int64
	ItemBoughtDate   string
	BoughtDuration   string // how quickly the item as bought
	CanceledTs       int64
	CanceledDate     string
	CanceledDuration string
	RealDate         string // date corresponding to the event (newoffer/itembought/canceled)
	RealTs           int64
}

type FloorPrice struct {
	TimeStamp int64
	Price     float64
}

type MintReportRec struct {
	Year         int64
	Month        int64
	MonthStr     string
	TotalMinted  int64 // number of tokens minted for the period
	TotalWei     string
	TotalEth     float64
	RedeemAmount float64 // sum of the price accumulated / 2 in ETH
}
