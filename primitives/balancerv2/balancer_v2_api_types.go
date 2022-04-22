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
	IsSwap				bool
	DateTime			string
	Amount				string
	Balance				string
}
type BalV2TokenInfo		struct {
	TokenAid			int64
	TokenAddr			string
	CurBalance			float64
}
type BalV2PoolToken struct {
	Token				BalV2TokenInfo
	Balances			[]BalV2PoolTokBalanceHistory
}
