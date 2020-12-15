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
	MeshEvtUnexpired			// 7
	MeshEvtBecameUnfunded		// 8
	MeshEvtFillabilityIncreased
	MeshEvtStoppedWatching
)
type SearchResultType int
const (
	SR_Unknown SearchResultType = iota
	SR_MarketOrders				// 1
	SR_Address					// 2
	SR_Hash						// 3
	SR_Transaction				// 4
	SR_Block					// 5
	SR_UserInfo					// 6
	SR_WalletContractInfo		// 7
	SR_AugurMarketInfo			// 8
	SR_AugurUniverseInfo		// 9
	SR_ShareTokenWrapper		// 10
	SR_BalancerPool				// 11
	SR_UniswapPair				// 12
	SR_TextSearchResults		// 13
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
	TimeStamp			int64
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
	MktAid						int64
	Fee							float64
	OpenInterest				float64
	CurVolume					float64
	MoneyAtStake				float64
	VolTraded					float64
	NoShowBond					float64
	ValidityBond				float64
	LowPriceLimit				float64
	HighPriceLimit				float64
	TotalTrades					int64
	NumTicks					int64
	CreatedTs					int64
	EndTs						int64
	FinTs						int64
	DesignRepTimeLeft			int64
	MktType						int
	MktStatus					int
	WinOutcomeIdx				int
	OutsideAugurBalanceChanges	bool
	MktAddr						string
	MktAddrSh					string	// short address (with .. in the middle)
	Signer						string
	SignerSh					string
	MktCreator					string
	MktCreatorSh				string	// short address (with .. in the middle)
	Reporter					string
	ReporterSh					string
	EndDate						string
	Description					string
	LongDesc					string
	CategoryStr					string
	Outcomes					string
	MktTypeStr					string
	Status						string
	CurOutcome					string	// calculated only if the query is made on specific outcome
	ScalarUnits					string
	Subcategories				[]string	// splitted string of subcategories
	OutcomeVolumes				[]OutcomeVol
	PriceEstimates				[]PriceEstimate
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
	Addr			string
	AddrSh			string	// short version of the addr
	DateCreated		string
	Expires			string
	OrderHash		string
}
type MarketDepth struct {
	LastOOID		int64
	Bids			[]DepthEntry
	Asks			[]DepthEntry
}
type UserInfo struct {
	Aid						int64
	BlockNum				int64
	TimeStamp				int64	// user registration timestamp (from block table)
	ProfitLoss				float64 // profit/loss for the (account) lifetime
	TradeFreq				float64	// trade frequency as percentil of all users (ex: top 15% of all users)
	ReportProfits			float64	// amount of money user has made in profits in outcome reporting
	AffProfits				float64	// profits made in affiliate commissions
	MoneyAtStake			float64	// how much money User has invested
	ValidityBonds			float64	// amount of validity bonds for all the markets user created
	TotalWithdrawn			float64	// amount of money User has deposited
	TotalDeposited			float64	// amount of money User has withdrawn
	TopTrades				float64
	TopProfit				float64
	UnclaimedProfit			float64
	WalletAid				int64	// Filled only if present
	EOAAid					int64	// Filled only if present
	BalancerNumSwaps		int64	// statistics: how many swaps at Balancer
	UniswapNumSwaps			int64	// statistics: how many swaps at Uniswap
	HedgingProfits			bool	// Flag to indicate negative 'MoneyAtStake' field
	NoActivity				bool	// True if doesn't have entry in 'ustats' table
	TotalTrades				int32	// how many trades were made by this User
	MarketsCreated			int32	// how many markets this User has created
	MarketsTraded			int32	// how many markets this User has traded
	WithdrawReqs			int32	// number of withdrawal requests
	DepositReqs				int32	// number of Deposit requests
	TotalReports			int32	// amount of reports User has made
	TotalDesignated			int32	// total reports submitted as designated reporter
	AugurFlags				AugurAcctFlags
	Addr					string	// User's Ethereum address (Externally Owned Account)
	AddrSh					string	// short version of the above addr
	WalletAddr				string	// Wallet contract address, filled only if present
	EOAAddr					string	// EOA address (controlling account) for wallet contract (if present)
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
	Addr				string
	AddrSh				string
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
	Aid					int64
	TotalTrades			int64
	ProfitLoss			float64
	VolumeTraded		float64
}
type ProfitMaker struct {
	Percentage			float64
	ProfitLoss			float64
	Addr				string
}
type TradeMaker struct {
	Percentage			float64
	TotalTrades			int64
	Addr				string
}
type VolumeMaker struct {
	Percentage			float64
	Volume				float64
	Addr				string
}
type UserRank struct {
	Aid					int64
	ProfitLoss			float64
	TotalTrades			int64
	VolumeTraded		float64
	NumMktCreated		int64
	Addr				string
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
	CreatorAddr			string
	CreatorAddrSh		string	// short version of the addr
	FillerAddr			string
	FillerAddrSh		string
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
	Reporter			string
	MktAddr				string
	MktAddrSh			string
	MktDescription		string
	OutcomeStr			string
	Date				string
	ReportType			string
	WinStart			string
	WinEnd				string
}
type AgtxInBlock struct {
	TxId				int64
	BlockNum			int64
	ContextAid			int64
	MarketAid			int64
	PoolAid				int64
	PairAid				int64
	TimeStamp			int64
	TxType				int
	ContextAddr			string		// Address related to transaction type
	MarketAddr			string		// Market address related to tx
	TxHash				string
	OrderHash			string		// if available, hash of the Order
	PoolAddr			string
	PairAddr			string
}
type MarketVeryShortInfo struct {
	MktAddr				string
	MktDesc				string
}
type BlockInfo struct {
	BlockNumFrom		int64
	BlockNumTo			int64
	FromTimeStamp		int64
	ToTimeStamp			int64
	NumBlocks			int64
	NumAugurTx			int64		// Only Augur-related transaction counter
	NumEvents			int64
	NumAugurEvents		int64
	NumDefiEvents		int64
	NumOtherEvents		int64
	NumBalSwaps			int64		// Num swaps at Balancer
	NumUniSwaps			int64		// Num swaps at Uniswap
	NumMarketsTraded	int64
	NumMarketsCreated	int64
	NumUniqueAddresses	int64
	NumUniqueOrders		int64
	GasUsed				int64
	TxCostEth			float64
	FromDate			string
	ToDate				string
	ActiveAddresses		[]string	//list of addresses participated in this block
	Transactions		[]string
	Orders				[]string
	MarketsTraded		[]MarketVeryShortInfo // list of market addresses created at this block
	MarketsCreated		[]MarketVeryShortInfo
//	BlockTransactions	[]AgtxInBlock	DISCONTINUED

}
type PoolVeryShortInfo struct {
	PoolAid				int64
	NumSwaps			int64
	NumHolders			int64
	PoolAddr			string
}
type PairVeryShortInfo struct {
	PairAid				int64
	TotalSwaps			int64
	PairAddr			string
}
type TokenVeryShortInfo struct {
	TokenAid			int64
	TokenAddr			string
	Name				string
	Symbol				string
}
type AgtxEvent struct {
	EvtType				int
	DeFiPlatformCode	int
	ReferenceId			int64	// Could be Market Order ID, or Event Log id
	Aid					int64
	MktAid				int64
	DeFiSwapId			int64
	Addr				string
	MktAddr				string
	OrderHash			string
	MktDescr			string
}
type TxInfo struct {
	TotalEvents			int
	NumAugurEvents		int
	NumDeFiEvents		int
	NumOtherEvents		int
	NumBalancerSwaps	int
	NumUniswapSwaps		int
	TxId				int64
	GasUsed				int64
	BlockNum			int64
	FromAid				int64
	ToAid				int64
	Value				float64
	TxFeeEth			float64
	Hash				string
	From				string
	To					string
	BalancerSwaps		[]BalancerSwap
	UniswapSwaps		[]UniswapSwap
	BalancerPools		[]PoolVeryShortInfo
	UniswapPairs		[]PairVeryShortInfo
	TokensTraded		[]TokenVeryShortInfo
	MarketsTraded		[]MarketVeryShortInfo
	FullEventList		[]AgtxEvent
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
	Addr				common.Address
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
type GasUsed struct {	// used to pass values of Statistics of Gas Usage
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
}
type TxCost struct {	// used to pass values of Statistics of Gas Usage
	Day					int64
	Ts					int64
	Num_trading			int64
	Num_reporting		int64
	Num_markets			int64
	Num_total			int64
	EthTrading			string
	EthReporting		string
	EthMarkets			string
	EthTotal			string
}
type AccumGasUsed struct {
	Day					int64
	Ts					int64
	Num_trading			int64
	Num_reporting		int64
	Num_markets			int64
	Num_total			int64
	Trading				int64
	Reporting			int64
	Markets				int64
	Total				int64
}
type AccumTxCost struct {
	Day					int64
	Ts					int64
	Num_trading			int64
	Num_reporting		int64
	Num_markets			int64
	Num_total			int64
	EthTrading			float64
	EthReporting		float64
	EthMarkets			float64
	EthTotal			float64
}
type GasCounter struct {
	TimeStamp			int64
	GasUsed				int64
	TxCost				float64
	NumRecs				int64
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
	// helping fields
	MarketAddress			string
	NumTicks				int
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
type EthereumEventTopic struct {
	BlockNum				int64
	TxId					int64
	EventLogId				int64
	ContractAid				int64
	Pos						int
	Value					string	// Important note: this isn't 0x-prefixed
}
type EthereumEventLog struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	ContractAid				int64
//	ContractAddress			string
	Topic0_Sig				string
	RlpLog					[]byte
}
type ChainReorg struct {
	BlockNum				int64
	Hash					string
}
type EvtLogEntry struct { // Layer1 entry (event)
	BlockNum				int64
	TxId					int64
	EvtId					int64
//	TxHash					*string
}
type ETLTokenConfig struct {
	LastIdDAI				int64
	LastIdREP				int64
	LastIdShareTokTransf	int64		// ShareToken.sol::Transfer event
	LastIdShareTokBalChg	int64		// ShareToken.sol::ShareTokenBalanceChanged event
}
type InspectedEvent struct {
	ContractAid				int64
	Signature				string
}
type ShortEvtLog struct {
	EvtId					int64
	TxId					int64
}
type TokProcessStatus struct {
	LastEvtId				int64
}
type AugurProcessStatus struct {
	LastTxId				int64
}
type AugurFoundryStatus struct {
	LastEvtId				int64
}
type ERC20ShTokContract struct {
	TimeStamp				int64
	WrapperAid				int64
	LastEvtId				int64
	MarketAid				int64
	OutcomeIdx				int
	Decimals				int
	Address					string
	Symbol					string
	Name					string
	MktAddr					string
}
type WShTokTransfer struct { // (ERC20) Wrapped ShareToken Transfer
	TimeStamp				int64
	EvtLogId				int64
	WrapperAid				int64
	BlockNum				int64
	TxId					int64
	FromAid					int64
	ToAid					int64
	Amount					float64
	FromPool				bool
	ToPool					bool
	NonPoolTransfer			bool
	AmountStr				string
	From					string
	To						string
}
type BalancerStatus struct {
	LastEvtId				int64
}
type BalancerNewPool struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	CallerAid				int64
	NumSwaps				int64
	NumHolders				int64
	NumTokens				int64
	NumAugurTokens			int64
	SwapFee					float64
	PoolAddr				string
	CallerAddr				string
	CreatedDate				string
	ControllerAddr			string
	Tokens					[]BalancerToken
}
type BalancerToken struct {
	TimeStampAdded			int64
	TokenAid				int64
	Denorm					float64
	Weight					float64
	Balance					float64
	TokenAddr				string
	DateAdded				string
	WrappingContract		ERC20ShTokContract 
}
type BalancerJoin struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	CallerAid				int64
	TokenAid				int64
	PoolAddr				string
	CallerAddr				string
	TokenInAddr				string
	AmountIn				string
}
type BalancerExit struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	CallerAid				int64
	TokenAid				int64
	PoolAddr				string
	CallerAddr				string
	TokenOutAddr			string
	AmountOut				string
}
type BalancerSwap struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	CallerAid				int64
	TokenInAid				int64
	TokenOutAid				int64
	AmountInF				float64
	AmountOutF				float64
	DecimalsIn				int
	DecimalsOut				int
	PoolAddr				string
	CallerAddr				string
	TokenInAddr				string
	TokenOutAddr			string
	SymbolIn				string
	SymbolOut				string
	AmountIn				string
	AmountOut				string
	Date					string
}
type SetSwapFee struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	FeeStr					string
}
type SetController struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	ControllerAddr			string
}
type SetPublic struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	Public					bool
}
type Finalize struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
}
type PoolBind struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	Denorm					string
	PoolAddr				string
	TokenAddr				string
	Balance					string
}
type PoolUnBind struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	TokenAddr				string
}
type PoolReBind struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	Denorm					string
	PoolAddr				string
	TokenAddr				string
	Balance					string
}
type PoolGulp struct {
	Id						int64
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	PoolAid					int64
	PoolAddr				string
	TokenAddr				string
	AbsorbedBalance			string
}
type AugurAcctFlags struct { // indicates if Account is Augur - enabled or not (by approvals)
	BlockNum				int64
	CreatedTs				int64
	AugurEnabled			bool
	ZeroXOnCash				bool
	FillOnCash				bool
	FillOnShareToken		bool
	SetReferrer				bool
}
type UpdatePriceEst struct {
	MktAid					int64
	TimeStamp				int64
	MktType					int
	Outcomes				string
}
type OutsideAugurSBChg struct {
	AccountAid				int64
	TimeStamp				int64
	MktAid					int64
	BlockNum				int64
	Balance					float64
	OutcomeIdx				int
	Address					string
	DateTime				string
	TxHash					string
	TxHashSh				string
}
type PoolInfo struct {
	PoolAid					int64
	NumHolders				int64
	NumSwaps				int64
	CreatedBlockNum			int64
	WentPublicTs			int64
	CreatedTs				int64
	FinalizedTs				int64
	NumTokens				int
	IsPublic				bool
	WasFinalized			bool
	UsdLiquidity			float64
	SwapFee					float64
	PoolAddr				string
}
type MarketPool struct {
	OutcomeIdx				int
	OutcomeStr				string
	MktAddress				string
}
type TradingVolume struct {
	TimeStamp				int64
	NumRecords				int64
	Amount					float64
}
type ERC20Info struct {
	Id						int64
	Aid						int64
	TotalSupplyF			float64
	Decimals				int
	TotalSupply				string
	Address					string
	Name					string
	Symbol					string
}
type UniswapStatus struct {
	LastEvtId				int64
}
type BasicChainInfo struct { // piece of common information for storing in tables
	BlockNum				int64
	TxId					int64
	EvtId					int64
	TimeStamp				int64
}
type MarketUPair struct { // Uniswap Pair where the Market can be traded
	MktAid					int64
	OutcomeIdx				int64
	PairAid					int64
	Token0Aid				int64
	Token1Aid				int64
	TotalSwaps				int64
	CreatedTs				int64
	NumAugurTokens			int64
	Token0Decimals			int
	Token1Decimals			int
	Outcome					string
	MktAddr					string
	CreatedDate				string
	PairAddr				string
	Token0Addr				string
	Token1Addr				string
	Token0Name				string
	Token1Name				string
	Token0Symbol			string
	Token1Symbol			string
}
type UniswapSwap struct {
	Id						int64
	PairAid					int64
	BlockNum				int64
	RequesterAid			int64
	CreatedTs				int64
	Amount0_In				float64
	Amount1_In				float64
	Amount0_Out				float64
	Amount1_Out				float64
	CreatedDate				string
	RequesterAddr			string
	PairAddr				string
	Symbol0					string
	Symbol1					string
}
type TextSearchResult struct {
	ObjType					int
	CatId					int64
	MktAid					int64
	TotalMarkets			int64
	Volume					float64
	MktAddr					string
	MktDescription			string
	Category				string
}
type UPairTokens struct {
	Decimals0				int
	Decimals1				int
	Token0Addr				common.Address
	Token1Addr				common.Address
}
type SearchResultObject struct {
	SRType					SearchResultType
	Found					bool
	ErrStr					string
	Query					string
	Object					interface{}
}
type BSwapPrice struct {
	Id						int64
	TimeStamp				int64
	NumRecords				int64
	Price					float64
	Date					string
}
type UPairPrice struct {
	Id						int64
	TimeStamp				int64
	NumRecords				int64
	Price					float64
	Date					string
}
type AddressInfo struct {
	Aid						int64
	Addr					string
}
type TokenSlippage struct {
	Token1Addr				string
	Token2Addr				string
	Token1Symbol			string
	Token2Symbol			string
	PoolAddr				string
	NumSwaps				int64
	Decimals1				int
	Decimals2				int
	ReservesTok1			float64
	ReservesTok2			float64
	Slippage				float64
	AmountIn				float64
	AmountOut				float64
}
type UserShTokens struct {
	WrapperAid				int64
	NumTransfers			int64
	MarketAid				int64
	Balance					float64
	OutcomeIdx				int
	Symbol					string
	Name					string
	WrapperAddr				string
	MarketAddr				string
}
type EthUsdPriceEvt struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Amount0In				float64
	Amount1In				float64
	Amount0Out				float64
	Amount1Out				float64
	EthUsd					float64
}
type EthUsdProcessStatus struct {
	LastEvtId				int64
}
