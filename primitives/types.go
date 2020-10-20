// Global types, used anywhere in the package
package primitives
import (
	"math/big"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
)
const (
	MAX_BLOCKS_CHAIN_SPLIT = 128
	OWNER_FIELD_OFFSET int = 2	// offset to the 'owner' field in WalletContract in EVM (contract storage)
	CATEGORICAL_MULTIPLIER int = 1000
	SCALAR_MULTIPLIER int = 10
)
const (
	MktTypeYesNo		= iota
	MktTypeCategorical
	MktTypeScalar
)
const (
	OOOpCodeNone= iota
	OOOpCodeCreated
	OOOpCodeFill
	OOOpCodeCancelledByUser
	OOOpCodeExpired
	OOOpCodeSyncProcess		// when no other reason exist, this one is used (this is a kind of a bugfix)
)
type MeshEvtCode int
const (
	MeshEvtGetOrders MeshEvtCode = iota
	MeshEvtInvalid
	MeshEvtAdded				// 2
	MeshEvtFilled				// 3
	MeshEvtFullyFilled			// 4
	MeshEvtCancelled			// 5
	MeshEvtExpired				// 6
	MeshEvtUnexpired
	MeshEvtBecameUnfunded
	MeshEvtFillabilityIncreased
	MeshEvtStoppedWatching
)
var (
	ErrChainSplit error = errors.New("Chainsplit detected")
)

type OrderType uint8
const (
	OrderTypeBid		OrderType = 0
	OrderTypeAsk		OrderType = 1
)
type OrderAction uint8
const(
	OrderActionCreate	OrderAction = 0
	OrderActionCancel	OrderAction = 1
	OrderActionFill		OrderAction = 2
)
type TokenType uint8
const(
	ReputationToken		TokenType = 0
	DisputeCrowdsourcer TokenType = 1
	ParticipationToken	TokenType = 2
)
type MarketStatus uint8
const (
	MktStatusTraded		MarketStatus = 0
	MktStatusReporting	MarketStatus = 1
	MktStatusReported	MarketStatus = 2
	MktStatusDisputing	MarketStatus = 3
	MktStatusFinalized	MarketStatus = 4
	MktStatusFinInvalid	MarketStatus = 5
)
//type BlockNumber int64	// -1 is used to mark 'block not set' for the database DISCONTINUED: 
type AugurTx struct {	// just a wrapper for Ethereum Transaction object, but in our own format
	TxId				int64		// once inserted tx_id is stored here
	BlockNum			int64
	GasUsed				int64
	TxIndex				int32
	CtrctCreate			bool
//	TxMsg				*types.Message	Discontinued , to be deleted
	GasPrice			string
	TxHash				string
	From				string
	To					string
	Value				string
	Input				[]byte
}
type ExtraInfo struct {
	Categories			[]string	`json:"categories"`
	Description			string		`json:"description"`
	Tags				[]string	`json:"tags"`
	LongDescription		string		`json:"longDescription"`
}
type InfoMarket struct {
	MktAid			int64
	Fee				float64
	OpenInterest	float64
	CurVolume		float64
	MoneyAtStake	float64
	VolTraded		float64
	NoShowBond		float64
	ValidityBond	float64
	LowPriceLimit	float64
	TotalTrades		int64
	NumTicks		int64
	CreatedTs		int64
	EndTs			int64
	FinTs			int64
	MktType			int
	MktStatus		int
	WinOutcomeIdx	int
	MktAddr			string
	MktAddrSh		string	// short address (with .. in the middle)
	Signer			string
	SignerSh		string
	MktCreator		string
	MktCreatorSh	string	// short address (with .. in the middle)
	Reporter		string
	ReporterSh		string
	EndDate			string
	Description		string
	LongDesc		string
	CategoryStr		string
	Outcomes		string
	MktTypeStr		string
	Status			string
	CurOutcome		string	// calculated only if the query is made on specific outcome
	Subcategories	[]string	// splitted string of subcategories
	OutcomeVolumes	[]OutcomeVol
	PriceEstimates	[]PriceEstimate
}
type InfoCategories struct {
	CatId			int64
	TotalMarkets	int64
	Category		string
	Subcategories	[]string
}
type MarketTrade struct {
	OrderId			int64
	Price			float64
	Amount			float64
	Outcome			int
	OrderHash		string
	OrderHashSh		string
	MktAddr			string
	MktAddrSh		string	// short address (with .. in the middle)
	CreatorAddr		string
	CreatorAddrSh	string	// short address (with .. in the middle)
	FillerAddr		string
	FillerAddrSh	string	// short address (with .. in the middle)
	Type			string
	Direction		string
	Date			string
	OutcomeStr		string
}
type OutcomeVol struct {
	Outcome			int
	MktType			int
	Volume			float64
	LastPrice		float64
//	HighestBid		float64
//	LowestAsk		float64
//	CurSpread		float64
//	PriceEstimate	float64
	TotalTrades		int64
	TotalOpenOrders	int64
	MktAddr			string
	OutcomeStr		string
}
type DepthEntry struct {
	MktAid			int64
	ExpiresTs		int64
	Price			float64
	Volume			float64
	AccumVol		float64
	TotalBids		int32
	TotalAsks		int32
	TotalCancel		int32
	OutcomeIdx		int32
	WalletAddr		string
	WalletAddrSh	string	// short version of the addr
	EOAAddr			string
	EOAAddrSh		string	// short version of the addr
	DateCreated		string
	Expires			string
}
type MarketDepth struct {
	LastOOID		int64
	Bids			[]DepthEntry
	Asks			[]DepthEntry
}
type UserInfo struct {
	EOAAid			int64
	WalletAid		int64
	BlockNum		int64
	TimeStamp		int64	// user registration timestamp (from block table)
	ProfitLoss		float64 // profit/loss for the (account) lifetime
	TradeFreq		float64	// trade frequency as percentil of all users (ex: top 15% of all users)
	ReportProfits	float64	// amount of money user has made in profits in outcome reporting
	AffProfits		float64	// profits made in affiliate commissions
	MoneyAtStake	float64	// how much money User has invested
	ValidityBonds	float64	// amount of validity bonds for all the markets user created
	TotalWithdrawn	float64	// amount of money User has deposited
	TotalDeposited	float64	// amount of money User has withdrawn
	TopTrades		float64
	TopProfit		float64
	UnclaimedProfit	float64
	HedgingProfits	bool	// Flag to indicate negative 'MoneyAtStake' field
	NotAugur		bool	// True if doesn't have entry in 'ustats' table
	TotalTrades		uint32	// how many trades were made by this User
	MarketsCreated	uint32	// how many markets this User has created
	MarketsTraded	uint32	// how many markets this User has traded
	WithdrawReqs	uint32	// number of withdrawal requests
	DepositReqs		uint32	// number of Deposit requests
	TotalReports	uint32	// amount of reports User has made
	TotalDesignated	uint32	// total reports submitted as designated reporter
	EOAAddr			string	// User's Ethereum address (Externally Owned Account)
	EOAAddrSh		string	// short version of the above addr
	WalletAddr		string	// User's Wallet Contract Address
	WalletAddrSh	string	// short version of the above addr
}
type MainStats struct {
	LastBlockNum	int64
	MarketsCount	int64
	YesNoCount		int64
	CategCount		int64
	ScalarCount		int64
	ActiveCount		int64
	FinalizedCount	int64
	InvalidCount	int64
	MoneyAtStake	float64
	TradesCount		int64
}
/* DISCONTINUED, new object used instead is 'OrderInfo'. Removal pending
type MarketOrder struct {	// this is a short order info, to show in tables
	OrderId				int64
	MktAid				int64
	TradeTs				int64
	Price				float64
	Volume				float64
	AccumVol			float64
	CreatedTs			int64
	OutcomeIdx			int32
	OType				int32
	CreatorBuyer		bool
	FillerBuyer			bool
	OrderHash			string
	CreatorWalletAddr	string
	CreatorWalletAddrSh	string	// short version of the addr
	CreatorEOAAddr		string
	CreatorEOAAddrSh	string	// short version of the addr
	FillerWalletAddr	string
	FillerWalletAddrSh	string
	FillerEOAAddr		string
	FillerEOAAddrSh		string
	Direction			string
	Date				string
}*/
type PLEntry struct {	// profit loss entry
	Id					int64
	MktAid				int64
	Timestamp			int64
	BlockNum			int64
	NetPosition			float64
	AccumPl				float64
	AccumFrozen			float64
	AvgPrice			float64
	FrozenFunds			float64
	RealizedProfit		float64
	RealizedCost		float64
	ImmediateProfit		float64
	MktType				int
	OutcomeIdx			int
	ClaimStatus			int
	CreatorBuyer		bool
	FillerBuyer			bool
	Date				string
	EOAAddr				string
	EOAAddrSh			string
	WalletAddr			string
	WalletAddrSh		string
	MktAddr				string
	MktAddrSh			string
	OutcomeStr			string
	MktDescr			string
	OrderHash			string
	CounterPAddr		string
	CounterPAddrSh		string
}
type OpenOrder struct {		// the Order on 0x Mesh network, that is yet to be filled
	Id					int64
	MktAid				int64
	Amount				float64
	InitialAmount		float64
	Price				float64
	Timestamp			int64
	MktExpirationTs		int64
	OrderExpirationTs	int64
	MktOrderId			int64
	MktStatus			int
	MktType				int
	MarketStatus		int
	Outcome				int
	OrderType			int
	OpCode				int
	OrderDate			string
	Direction			string
	MktDescr			string
	OutcomeStr			string
	MktStatusStr		string
	MktTypeStr			string
	OrderHash			string
	OrderHashSh			string
	CreatorAddr			string
	CreatorAddrSh		string	// shortened address
	MktAddr				string
	MktAddrSh			string
}
type RankStats struct {
	EoaAid				int64
	TotalTrades			int64
	ProfitLoss			float64
	VolumeTraded		float64
}
type ProfitMaker struct {
	Percentage			float64
	ProfitLoss			float64
	EOAAddr				string
}
type TradeMaker struct {
	Percentage			float64
	TotalTrades			int64
	EOAAddr				string
}
type VolumeMaker struct {
	Percentage			float64
	Volume				float64
	EOAAddr				string
}
type UserRank struct {
	EoaAid				int64
	WalletAid			int64
	ProfitLoss			float64
	TotalTrades			int64
	VolumeTraded		float64
	NumMktCreated		int64
	EOAAddr				string
	WalletAddr			string
}
type OrderInfo struct {		// this is a full order information, to show in dedicated webpage
	OrderId				int64
	MktAid				int64
	TradeTs				int64
	Price				float64
	Amount				float64
	AccumVol            float64
	CreatedTs			int64
	MktType				int32
	OType				int32
	OutcomeIdx			int32
	CreatorBuyer		bool	// true if the Creator is the buyer
	FillerBuyer			bool	// true if the Filler is the buyer
	OrderHashSh			string
	OrderHash			string
	OTypeStr			string
	OutcomeStr			string
	CreatorWalletAddr	string
	CreatorWalletAddrSh	string	// short version of the addr
	CreatorEOAAddr		string
	CreatorEOAAddrSh	string	// short version of the addr
	FillerWalletAddr	string
	FillerWalletAddrSh	string
	FillerEOAAddr		string
	FillerEOAAddrSh		string
	Date				string
	MarketAddr			string
	MarketAddrSh		string
	Direction			string
}
type Report struct {
	MktAid				int64
	RepStake			float64
	Round				int
	OutcomeIdx			int
	MktType				int
	IsInitial			bool
	IsDesignated		bool
	MktAddr				string
	MktAddrSh			string
	MktDescription		string
	OutcomeStr			string
	Date				string
	ReportType			string
	WinStart			string
	WinEnd				string
}
type BlockInfo struct {
	BlockNum			int64
	NumTx				int64
	NumAddresses		int64
	NumMarkets			int64
	Addresses			[]string	//list of addresses participated in this block
	Transactions		[]string
	Markets				[]string	// list of market addresses created at this block
}
type TxInfo struct {
	BlockNum			int64
	Value				float64
	Hash				string
	From				string
	To					string
}
type FrontPageStats struct {
	MoneyAtStake		float64
	MarketsCreated		float64
	TradesCount			int64
}
type DaiB struct {
	Id					int64
	Aid					int64
	DaiTransfId			int64
	BlockNum			int64
	Processed			bool
	Address				string
	Amount				string
	Balance				string
	BlockHash			string
}
type DaiOp struct {
	BlockNum			int64
	Date				string
	Deposit				string
	Withdrawal			string
	FromAddr			string
	ToAddr				string
}
type BlockCash struct {
	BlockNum			int64
	Ts					int64
	CashFlow			float64
	AccumCashFlow		float64
}
type ContractAddresses struct {
	ChainId					int64
	Augur					common.Address	// Main Augur contract
	AugurTrading			common.Address	// Augur trading contract
	PL						common.Address	// ProfitLoss contract
	ZeroxTrade				common.Address	// ZeroX Trade contract
	ZeroxXchg				common.Address	// ZeroX Exchange contract
	Dai						common.Address	// Shows DAI balance and also to fill dai_transf table 
	Reputation				common.Address	// used to query REP token balance when showing User info
	WalletReg				common.Address	// used to get the link between EOA and Wallet contract
	WalletReg2				common.Address	// same as WalletReg but with GSN, used for EOA-Wallet link
	FillOrder				common.Address	// used to identify if DAI transfer is internal or not
	EthXchg					common.Address	// used to identify if DAI transfer is internal or not
	ShareToken				common.Address  // used to identify if DAI transfer is internal or not
	GenesisUniverse			common.Address	// used to identify if DAI transfer is internal or not
	CreateOrder				common.Address	// CreateOrder contract, used to detect wallet creation pattern
	LegacyReputationToken	common.Address
	BuyParticipationTokens	common.Address
	RedeemStake				common.Address
	WarpSync				common.Address
	HotLoading				common.Address
	Affiliates				common.Address
	AffiliateValidator		common.Address
	Time					common.Address
	CancelOrder				common.Address
	Orders					common.Address
	SimulateTrade			common.Address
	Trade					common.Address
	OICash					common.Address
	UniswapV2Factory		common.Address
	UniswapV2Router02		common.Address
	AuditFunds				common.Address
	WETH9					common.Address
	USDC					common.Address
	USDT					common.Address
	RelayHubV2				common.Address
	AccountLoader			common.Address
}
type UniqueAddrEntry struct {
	Ts					int64
	NumAddrs			int64
	NumAddrsAccum		int64
	Day					string
}
type MktDepthStatus struct {
	NumOrders			int64	// used to catch deletes
	LastOOID			int64	// used to catch new inserts
}
type PosChg struct {		// change in positon for logging/debugging purposes
	Mkt_addr			common.Address
	Wallet_addr			common.Address
	BlockNum			int64
	Outcome				*big.Int
	ProfitLoss			*big.Int
	FrozenFunds			*big.Int
	NetPos				*big.Int
	AvgPrice			*big.Int
}
type GasSpent struct {	// used to pass values of Statistics of Gas Usage
	Day					int64
	Ts					int64
	Num_trading			int64
	Num_reporting		int64
	Num_markets			int64
	Num_total			int64
	Trading				string
	Reporting			string
	Markets				string
	Total				string
	EthTrading			string
	EthReporting		string
	EthMarkets			string
	EthTotal			string
}
func (obj *GasSpent) Has_rows() bool {
	if (obj.Num_trading==0) && (obj.Num_reporting==0) && (obj.Num_markets==0) && (obj.Num_total==0) {
		return false
	}
	return true
}
type PriceHistory struct {
	OutcomeIdx			int
	OutcomeStr			string
	Trades				[]OrderInfo
}
type FullPriceHistory struct {
	Outcomes			[]PriceHistory
}
type ZHistT1Entry struct {		// the Order on 0x Mesh network, that is yet to be filled
	Id						int64
	MktAid				int64
	Amount				float64
	FillableAmount		float64
	Price				float64
	PriceEstimate		float64
	WeightedPriceEst	float64
	Spread				float64
	MaxBid				float64
	MinAsk				float64
	WMaxBid				float64
	WMinAsk				float64
	Timestamp			int64
	MktExpirationTs		int64
	OrderExpirationTs	int64
	MktOrderId			int64
	MktStatus			int
	MktType				int
	MarketStatus		int
	OutcomeIdx			int
	OrderType			int
	EvtCode				int
	OrderDate			string
	Direction			string
	MktDescr			string
	OutcomeStr			string
	MktStatusStr		string
	MktTypeStr			string
	OrderHash			string
	OrderHashSh			string
	MakerAddr			string
	MakerAddrSh			string	// shortened address
	RelatedAddr			string	// address of the filler or the one who cancles, or nil
	EOAAddr				string
	EOAAddrSh			string
	WalletAddr			string
	WalletAddrSh		string
	RelatedAddrSh		string
	MktAddr				string
	MktAddrSh			string
}
type ZHistT2Entry struct { // Type2 entry, summarized data
	Timestamp				int64
	PriceEstimate			float64
	WeightedPriceEstimate	float64
}
type ZoomedPriceHist struct {
	OutcomeIdx			int
	InitTs				int
	FinTs				int
	IntervalSecs		int
	OutcomeStr			string
	Type1Entries		[]ZHistT1Entry
	Type2Entries		[]ZHistT2Entry
}
type FullZoomedPriceHist struct {
	Outcomes			[]ZoomedPriceHist
}
type StatementEntry struct {
	Id					int64
	BlockNum			int64
	MktAid				int64
	Amount				float64
	Balance				float64
	Date				string
	From				string
	FromSh				string
	To					string
	ToSh				string
	Info				string
	MktAddr				string
}
type ExecuteWalletTx struct {
	RevertOnFailure			bool
	To						string
	CallData				string	// hex-encoded bytecode of the Input to the contract in 'to'
	Value					string
	Payment					string
	ReferralAddress			string
	Fingerprint				string	// the fingerprint of the Browser of the account that does referrals
	DesiredSignerBalance	string
	MaxExchangeRateInDAI	string
	InputSig				string	// first 4 bytes of CallData, extracted for indexing
}
type MeshProcStatus struct {
	LastIdProcessed			int64
}
type MeshEvent struct {
	Id						int64
	Timestamp				int64
	FillableAmount			string
	EvtCode					int
	OrderHash				string
	ChainId					int
	ExchangeAddress			string
	MakerAddress			string
	MakerAssetData			string
	MakerFeeAssetData		string
	MakerAssetAmount		string
	MakerFee				string
	TakerAddress			string
	TakerAssetData			string
	TakerFeeAssetData		string
	TakerAssetAmount		string
	TakerFee				string
	SenderAddress			string
	FeeRecipientAddress		string
	ExpirationTime			int64
	Salt					string
	Signature				string
}
type DepthState struct {
	Id						int64
	MeshEvtId				int64
	MarketAid				int64
	OutcomeIdx				int64
	OrderType				int
	OrderHash				string
	Price					float64
	Amount					float64
	IniTs					int64
	FinTs					int64
	IniDate					string
	FinDate					string
}
type PriceEstimate struct {
	Id						int64
	MarketAid				int64
	MeshEvtId				int64
	TimeStamp				int64
	BidStateId				int64
	AskStateId				int64
	OutcomeIdx				int64
	Spread					float64
	PriceEst				float64
	WeightedPriceEst		float64
	MaxBid					float64
	MinAsk					float64
	WMaxBid					float64
	WMinAsk					float64
	EvtCode					int
	Date					string
	MatchingBids			[]DepthState
	MatchingAsks			[]DepthState
}
type OIAccum struct {
	TimeStamp				int
	AccumOpenInterest		float64
}
type TradesByInterval struct {
	TimeStamp				int64
	NumTrades				int64
	AccumNumTrades			int64
	Volume					float64
	AccumVolume				float64
}
