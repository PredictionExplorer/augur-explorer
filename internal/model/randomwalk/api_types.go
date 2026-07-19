package randomwalk

// =====================================================================
// API Response Types
// =====================================================================

// Offer is one marketplace offer row (sell or buy side, per OfferType) with
// its outcome flags and the seller's profit when resolved.
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

// TokenMint is one RandomWalk token mint with its generation seed and price.
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

// TokenMintCSV is the extended mint row of the CSV export: the mint joined
// with lifetime trading, naming and current-owner columns.
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

// UserToken is one RandomWalk token currently owned by a wallet.
type UserToken struct {
	TokenId      int64
	ContractAid  int64
	ContractAddr string
	Seed         string
	SeedNum      string
	Price        float64
}

// RWalkStats is the global RandomWalk statistics response: mint, trading and
// withdrawal aggregates (rw_stats row).
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

// MarketStats is the marketplace trading aggregate (volume and trade count).
type MarketStats struct {
	TradingVol float64
	NumTrades  int64
}

// HistEntryMint is the token-history entry for the mint (RecordType 1).
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

// HistEntryOffer is the token-history entry for an offer creation
// (RecordType 2).
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

// HistEntryOfferCanceled is the token-history entry for an offer
// cancellation (RecordType 3).
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

// HistEntryItemBought is the token-history entry for a marketplace sale
// (RecordType 4).
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

// HistEntryTokenName is the token-history entry for a rename (RecordType 5).
type HistEntryTokenName struct { // type = 5
	BlockNum     int64
	TimeStamp    int64
	DateTime     string
	ContractAid  int64
	ContractAddr string
	TokenId      int64
	TokenName    string
}

// HistEntryTransfer is the token-history entry for an ERC-721 transfer
// (RecordType 6).
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

// FullHistoryEntry is one polymorphic token-history row: RecordType selects
// which HistEntry* shape Record carries.
type FullHistoryEntry struct {
	RecordType int // Mint, or any other event starting with HistEntry*
	Record     any
}

// VolumeHistory is one bucket of the trading-volume time series with the
// running total.
type VolumeHistory struct {
	StartTs       int64
	NumOperations int64
	Volume        float64
	VolumeAccum   float64
}

// TokenNameRec is one token rename with the owner at rename time.
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

// UserInfo is one wallet's RandomWalk activity aggregate (rw_users row).
type UserInfo struct {
	UserAid               int64
	UserAddr              string
	TotalVolume           float64
	TotalNumTrades        int64
	TotalMintedToks       int64
	TotalNumWithdrawals   int64
	IsMarketPlaceContract bool // true if the User is not really a user, but is a marketplace contract
}

// TopTradedToken is one row of the most-traded-tokens ranking.
type TopTradedToken struct {
	TokenId     int64
	TotalTrades int64
	SeedHex     string
}

// TokenInfo is the per-token detail: current owner, seed, trading aggregates
// and naming state.
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

// MintInterval is one mint's spacing from the previous mint (the retired
// mint-intervals chart).
type MintInterval struct {
	MintNumber int64 // sequential number of mint (same as token_id)
	TimeStamp  int64
	Interval   int64 // Duration from previous mint
	TokenId    int64
}

// WithdrawalChartEntry is one point of the withdrawable-amount chart.
type WithdrawalChartEntry struct {
	TimeStamp        int64
	DateTime         string
	WithdrawalAmount float64
}

// TradingHistoryLog is one offer joined with its outcome (sale or
// cancellation) for the trading-history listing, with display durations.
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

// FloorPrice is one point of the sell-side listing-floor history.
type FloorPrice struct {
	TimeStamp int64
	Price     float64
}

// MintReportRec is one calendar month's mint totals in the mint report.
type MintReportRec struct {
	Year         int64
	Month        int64
	MonthStr     string
	TotalMinted  int64 // number of tokens minted for the period
	TotalWei     string
	TotalEth     float64
	RedeemAmount float64 // sum of the price accumulated / 2 in ETH
}
