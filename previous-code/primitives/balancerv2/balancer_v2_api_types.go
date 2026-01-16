package balancerv2

import (
)


type BalV2PoolInfo struct {
	PoolId				string
	BlockNum			int64
	PoolAid				int64
	Specialization		int
	PoolAddr			string
	Unhandled			bool
	UnhandledComments	string
	FirstSwapTs			int64
	LastSwapTs			int64
}
type BalV2PoolTokBalanceHistory struct {
	BlockNum			int64
	TimeStamp			int64
	TokenInAid			int64
	TokenOutAid			int64
	OpSign				int
	IsSwap				bool
	DateTime			string
	Amount				string
	Balance				string
}
type BalV2TokenInfo		struct {
	TokenAid			int64
	TokenAddr			string
	CurBalance			float64
	CurBalanceUSD		float64	// converted to USD using ethusd price and availableswap records
	USDBalanceAvailable	bool
	Symbol				string
	Name				string
}
type BalV2PoolToken struct {
	Token				BalV2TokenInfo
	Balances			[]BalV2PoolTokBalanceHistory
}
type BalV2SwapRecordInfo struct {
	BlockNum			int64
	TimeStamp			int64
	DateTime			string
	TokenInAid			int64
	TokenOutAid			int64
	TokenInAddr			string
	TokenInAddrShort	string
	TokenOutAddr		string
	TokenOutAddrShort	string
	AmountIn			string
	AmountOut			string
	DecimalsTokIn		int64
	DecimalsTokOut		int64
	AmountInFmt			string	// formatted for visualization (User)
	AmountOutFmt		string
	USDValue			float64	// profit from swap fees
	SymbolIn			string
	SymbolOut			string
}
type BalV2FeeReturns struct {
	TimeStamp			int64
	DateTime			string
	FeeReturnsUSD		float64

}
type BalV2LiqProvDistrib struct {	// Distribution of funds funded by liquidity providers into the pool
	FunderAid			int64
	FunderAddr			string
	Balance				float64		// this balance is not money, but the share of the funder in the pool
	Percentage			float64		// percent of the total liquidity funds
}
