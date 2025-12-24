package randomwalk

// =====================================================================
// API Response Types
// =====================================================================

type API_Offer struct {
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
	Profit       float64
	RWalkAid     int64
	RWalkAddr    string
	WasCanceled  bool
}

type API_TokenMint struct {
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

type API_TokenMint_CSV struct {
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

type API_UserToken struct {
	TokenId      int64
	ContractAid  int64
	ContractAddr string
	Seed         string
	SeedNum      string
	Price        float64
}

type API_RWalkStats struct {
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

type API_MarketStats struct {
	TradingVol float64
	NumTrades  int64
}

type API_HistEntry_Mint struct { // type = 1
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

type API_HistEntry_Offer struct { // type = 2
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

type API_HistEntry_OfferCanceled struct { // type = 3
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

type API_HistEntry_ItemBought struct { // type = 4
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

type API_HistEntry_TokenName struct { // type = 5
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	TokenId      int64
	TokenName    string
}

type API_HistEntry_Transfer struct { // type = 6
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

type API_FullHistoryEntry struct {
	RecordType int // Mint, or any other event starting with API_HistEntry_*
	Record     interface{}
}

type API_VolumeHistory struct {
	StartTs       int64
	NumOperations int64
	Volume        float64
	VolumeAccum   float64
}

type API_TokenName struct {
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

type API_UserInfo struct {
	UserAid               int64
	UserAddr              string
	TotalVolume           float64
	TotalNumTrades        int64
	TotalMintedToks       int64
	TotalNumWithdrawals   int64
	IsMarketPlaceContract bool // true if the User is not really a user, but is a marketplace contract
}

type API_Top5Toks struct {
	TokenId     int64
	TotalTrades int64
	SeedHex     string
}

type API_TokenInfo struct {
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

type API_MintInterval struct {
	MintNumber int64 // sequential number of mint (same as token_id)
	TimeStamp  int64
	Interval   int64 // Duration from previous mint
	TokenId    int64
}

type API_WithdrawalChartEntry struct {
	TimeStamp        int64
	DateTime         string
	WithdrawalAmount float64
}

type API_TradingHistoryLog struct {
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
	Profit           float64
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

type API_FloorPrice struct {
	TimeStamp int64
	Price     float64
}

type API_MintReportRec struct {
	Year         int64
	Month        int64
	MonthStr     string
	TotalMinted  int64 // number of tokens minted for the period
	TotalWei     string
	TotalEth     float64
	RedeemAmount float64 // sum of the price accumulated / 2 in ETH
}

