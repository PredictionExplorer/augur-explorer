// structs used to return reponses of the API v1
package primitives
import (
)
type API_AMM_AbstractMarketInfo struct {
	BlockNum			int64
	ContractAid			int64
	BlockTimeStamp		int64
	MarketCreatedTs		int64
	MarketEndTimeTs		int64
	MarketId			int64
	FactoryAid			int64
	ShareFactor			float64
	SettlementFee		float64
	ProtocolFee			float64
	StakerFee			float64
	Liquidity			float64
	BlockDateTime		string
	TxHash				string
	ContractAddr		string
	MarketCreatedDate	string
	MarketEndDate		string
	FactoryAddr			string
	CollateralAddr		string
	ProtocolAddr		string
	SettlementAddr		string
	FeePotAddr			string
	WinnerAddr			string
}
type API_AMM_SportsMarket struct {
	BlockNum				int64
	CreatedTs				int64
	ResolvedTs				int64
	Score					int64
	MarketTypeCode			int64
	HomeTeamId				int64
	AwayTeamId				int64
	EventId					int64
	CreatorAid				int64
	EstimatedStartTs		int64
	Liquidity				float64
	CreatorAddr				string
	HomeTeam				string
	AwayTeam				string
	Title					string
	Description				string
	ResolvedDate			string
	EstimatedStartDate		string
	MarketRules				[]string
	AbstractMarketInfo		API_AMM_AbstractMarketInfo
}
type API_AMM_Out_SharesBurned struct {// Arbitrum Augur SharesBurned event
	BlockNum				int64
	TimeStamp				int64
	MarketId				int64
	FactoryAid				int64
	CallerAid				int64
	Amount					float64
	FactoryAddr				string
	CallerAddr				string
	DateTime				string
	TxHash					string
}
type API_AMM_Out_SharesMinted struct {// Arbitrum Augur SharesMinted event
	BlockNum				int64
	TimeStamp				int64
	MarketId				int64
	FactoryAid				int64
	CallerAid				int64
	Amount					float64
	FactoryAddr				string
	CallerAddr				string
	DateTime				string
	TxHash					string
}
type API_AMM_Out_BalancerSwap struct {// Arbitrum Augur Balancer Swap of ShareTokens
	BlockNum				int64
	TimeStamp				int64
	CallerAid				int64
	TokenInAid				int64
	TokenOutAid				int64
	AmountIn				float64
	AmountOut				float64
	DateTime				string
	TxHash					string
	TokenOutAddr			string
	TokenInAddr				string
	CallerAddr				string
	TokenInSymbol			string
	TokenOutSymbol			string
	TokenInName				string
	TokenOutName			string
}
