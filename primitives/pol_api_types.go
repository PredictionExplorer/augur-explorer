package primitives

type API_Pol_BuySell_Op struct {
	BlockNum			int64
	TimeStamp			int64
	MarketId			int64
	MarketMakerAid		int64
	OperationType		int		// 0: BUY	1: SELL
	OutcomeIdx			int		// Outcome index
	CollateralAmount	float64		// How many cash were swapped for tokens
	FeeAmount			float64		// commission
	TokenAmount			float64		// How many tokens were received for cash
	UserAid				int64
	UserAddr			string
	DateTime			string
	MarketMakerAddr		string
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
