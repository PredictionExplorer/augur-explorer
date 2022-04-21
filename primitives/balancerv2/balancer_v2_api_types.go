package balancerv2

import (
)


type BalV2PoolInfo struct {
	PoolId				string
	BlockNum			int64
	Specialization		int
	PoolAddr			string
}
type BalV2PoolTokBalanceHistory struct {
	BlockNum			int64
	TimeStamp			int64
	FromAid				int64
	ToAid				int64
	FromAddr			string
	ToAddr				string
	DateTime			string
	Amount				float64
	Balance				float64
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
