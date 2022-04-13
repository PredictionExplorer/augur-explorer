// Global types, used anywhere in the package
package primitives
import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
)
//type BlockNumber int64	// -1 is used to mark 'block not set' for the database DISCONTINUED: 
type AugurTx struct {	// just a wrapper for Ethereum Transaction object, but in our own format
	TxId				int64		// once inserted tx_id is stored here
	BlockNum			int64
	GasUsed				int64
	TimeStamp			int64
	TxIndex				int32
	NumLogs				int32
	CtrctCreate			bool
	CumulativeGasUsed	uint64
	EffectiveGasPrice	*big.Int
//	TxMsg				*types.Message	Discontinued , to be deleted
	BlockHash			string
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
type MarketVeryShortInfo struct {
	MktAddr				string
	MktDesc				string
}
type FrontPageStats struct {
	MoneyAtStake		float64
	MarketsCreated		float64
	TradesCount			int64
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
	EthUsdPrice			float64
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
	TimeStamp				int64
	ContractAddress			string
	TxHash					string
	Topic0_Sig				string
	RlpLog					[]byte
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
type AugurProcessStatus struct {
	LastTxId				int64
}
type AugurFoundryStatus struct {
	LastEvtId				int64
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
type TradingVolume struct {
	TimeStamp				int64
	NumRecords				int64
	Amount					float64
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
type SearchResultObject struct {
	SRType					SearchResultType
	Found					bool
	ErrStr					string
	Query					string
	Object					interface{}
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
type EthUsdPrice struct {
	TimeStamp				int64
	Price					float64
}
type PayoutNumerator struct {
	IsInvalid				bool
	WinningOutcomeYesNo		int		// 0-reserved 1- No, 2-Yes
	WinningValueScalar	float64		// the value that has won
}
type ValidityBondPrice struct {
	TimeStamp				int64
	Price					float64
	DateTime				string
}
