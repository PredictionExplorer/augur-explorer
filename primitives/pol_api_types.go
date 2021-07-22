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


