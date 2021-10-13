package primitives

type API_Pol_BuySell_Op struct {
	Id					int64
	BlockNum			int64
	TimeStamp			int64
	MarketId			int64
	MarketMakerAid		int64
	OperationType		int		// 0: BUY	1: SELL
	OutcomeIdx			int		// Outcome index
	CollateralAmount	float64		// How many cash were swapped for tokens
	FeeAmount			float64		// commission
	FeeInCollateral		float64		// fee convert to USDC units
	TokenAmount			float64		// How many tokens were received for cash
	Price				float64		// price in collateral units (i.e. in USDC)
	UserAid				int64
	UserAddr			string
	DateTime			string
	MarketMakerAddr		string
}
type API_Pol_OutcomePriceHistoryEntry struct {
	OperationId			int64
	TimeStamp			int64
	OperationType		int32
	Price				float64
}
type API_OutcomePriceHistory struct {
	OutcomePriceHistory		[]API_Pol_OutcomePriceHistoryEntry
}
type API_Pol_Liquidity_Op struct {
	BlockNum			int64
	TimeStamp			int64
	MarketId			int64
	MarketMakerAid		int64
	OperationType		int		// 0: ADD 1: REMOVE
	CollateralAmount	float64		// How many cash were swapped for tokens
	FunderAid			int64
	FunderAddr			string
	DateTime			string
}
type API_Pol_MarketInfo struct {
	MarketId			int64
	Question			string
	QuestionId			string
	ConditionId			string
	Slug				string
	ResolutionSource	string
	CreatedAtTs			int64
	CreatedAtDate		string
	EndDateTs			int64
	EndDate				string
	ResolvedTs			int64
	ResolvedDate		string
	Category			string
	Image				string
	Icon				string
	Description			string
	Tags				string
	Outcomes			string
	MarketType			string
	MarketTypeCode		int
	MarketMakerAid		int64
	MarketMakerAddr		string
	OutcomeSlotCount	int64	// number of outcomes as reported by Prepare Condition event
	WasResolved			bool	// true if there is a ConditionResolution event
	Volume				float64
	OpenInterest		float64	// Liquidity Added + BUY operations - Fees
	Liquidity			float64	// Only the amount deposited by investors
	TotalTrades			int64
	TotalFeesCollected	float64	// Fees collected. Fees are not part of liquidity total (they are just held temporarily)
	NumTrades			int64	// total number of buy/sell operations
	NumLiquidityOps		int64
	QuestionIdHashValid		bool
	CondPrepTxHash		string	// Condition preparation transaction hash
	PayoutNumerators	[]float64 // Percentages
	PayoutNumeratorsStr string
}
type API_Pol_MarketStats struct {
	OpenInterest		float64
	NumLiquidityOps		int64	// total number of add/remove liquidity
	NumTrades			int64	// total number of buy/sell operations
	TotalVolume			float64
	TotalFeesCollected	float64
}
type API_Pol_CondPrepInfo struct {

}
type API_Pol_GlobalLiquidityHistoryEntry struct {
	StartTs					int64
	NumOperations			int64
	//SumAmounts				float64
	//SumShares				float64
	//SumCollateralRemoved	float64
	Liquidity				float64
	LiquidityAccum			float64
}
type API_Pol_MarketLiquidityHistoryEntry struct {
	StartTs					int64
	NumOperations			int64
	Liquidity				float64
	LiquidityAccum			float64
}
type API_Pol_GlobalTradingHistoryEntry struct {
	StartTs					int64
	NumOperations			int64
	TradingVol				float64
	TradingVolAccum			float64
}
type API_Pol_MarketTradingHistoryEntry struct {
	StartTs					int64
	NumOperations			int64
	TradingVol				float64
	TradingVolAccum			float64
}
type API_Pol_Unique_Users struct {
	TimeStamp				int64
	NumFunders				int64
	NumTraders				int64
	NumTotal				int64
}
type API_Pol_DataFeed struct {
	EvtlogId				int64
	TimeStamp				int64
	OperationType			int		// 0: ADD 1: REMOVE
	OutcomeIdx				int
	Collateral				float64
	Fee						float64
	UserAid					int64
	MarketId				int64
	MarketMakerAid			int64
	MarketMakerAddr			string
	UserAddr				string
	DateTime				string
	MarketQuestion			string

}
type API_Pol_TraderOp struct {
	BlockNum			int64
	TimeStamp			int64
	OperationType		int		// 0: BUY	1: SELL
	OutcomeIdx			int		// Outcome index
	CollateralAmount	float64		// How many cash were swapped for tokens
	FeeAmount			float64		// commission
	TokenAmount			float64		// How many tokens were received for cash
	ProfitLoss			float64
	AccumProfitLoss		float64
	DateTime			string
}
type API_Pol_TraderListEntry struct {
	UserAid					int64
	NumTrades				int64
	NumLiquidityOps			int64
	TotalTradeVolume		float64
	TotalLiquidityVol		float64
	TotalFeesPaid			float64
	TotalProfitLoss			float64
	UserAddr				string
}
type API_Pol_MarketInfoEntry struct {

}
type API_Pol_MarketOpenPosition struct {
	CurrentBalance			float64
	OutcomeIdx				int32
	TokenId					string
	UserAid					int64
	UserAddr				string
	NumTrades				int64
	TotalVolume				float64
	NumLiquidityOps			int64
	TotalFeesPaid			float64
	RealizedProfit			float64
	UnrealizedProfit		float64
	TotalProfit				float64
	CurrentPrice			float64
	PositionValue			float64 // Price * amount
}
type API_Pol_MarketUserOpenPosition struct {
	CollateralInvested		float64
	OutcomeIdx				int32
	TokenId					string
	MarketId				int64
	MarketQuestion			string
	NumTrades				int64
	TotalVolume				float64
	NumLiquidityOps			int64
	TotalFeesPaid			float64
	TotalProfit				float64
}
type API_Pol_LiquidityShareRatio struct {
	FunderAid				int64
	TotalLiquidityOps		int64		// total number of liquidity addition/removal operations
	Balance					float64
	FunderAddr				string
	ShareRatio				float64
}
type API_Pol_MarketRedemption struct {
	Id					int64
	BlockNum			int64
	TimeStamp			int64
	DateTime			string
	RedeemerAid			int64
	RedeemerAddr		string
	Outcomes			string
	Payout				float64
}
type API_Pol_UserRedemption struct {
	Id					int64
	BlockNum			int64
	TimeStamp			int64
	DateTime			string
	MarketId			int64
	Outcomes			string
	Payout				float64
}
type API_Pol_MarketCategory struct {
	Category			string
	NumMarkets			int64
}
type API_Pol_MarketERC1155Transfer struct {
	BalOpId				int64
	TokenId				int64
	TokenIdHex			string
	TimeStamp			int64
	DateTime			string
	ParentId			int64
	BatchId				int64
	IsBatch				bool
	BuySellOpType		int32
	FundOpType			int32
	Amount				float64
	Balance				float64
	BalChgAid			int64	// Address id of the account that changes balance
	BalChgAddr			string
	BuySellAid			int64	// Address id of the user who is doing Buy/Sell operation
	BuySellAddr			string
	BuySellAmount		float64
	FunderAid			int64	// Address id of the funder who is doing Fund Add/Remove operation
	FunderAddr			string
	FunderAmount		float64
	TxHash				string
}
type API_Pol_OpenInterestHistory struct {
	TimeStamp				int64
	DateTime				string
	TxId					int64	// if -1 then it means we have condition resolution record
	TxHash					string
	FromAid					int64
	FromAddr				string
	ToAid					int64
	UserAid					int64
	ToAddr					string
	PayoutNumerators		string
	BalChgId				int64		// id of erc20 balance transfer
	BuySellOpId				int64		// id for buy/sell operation
	BuySellOpType			int32
	FundOpId				int64		// id for fund add/remove operation
	FundOpType				int32
	RedeemId				int64		// id for Payout Redeemed event
	Amount					float64
	IntegerAmount			float64
	Balance					float64
	AdjustedBalance			float64
	IntegerBalance			float64
	OpenInterest			float64
	OIVerif					float64	// OI calculated copy for verification
	IntegerFee				float64
	Fee						float64
	FeeAccum				float64
	IntegerFeeAccum			float64
	ContractBalance			float64
	ContractBalanceAccum	float64
}
type API_Pol_OI_HistoryTotals struct {
	FinalOpenInterest		float64
	FinalFees				float64
}
type PolTextSearchResult struct {
	ObjType					int
	MarketId				int64
	ContractAid				int64
	Volume					float64
	ContractAddr			string	// aka Market Maker address
	Description				string
	Title					string
}
