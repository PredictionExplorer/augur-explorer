package balancerv2

import (
)


type BalV2PoolInfo struct {
	PoolId				string
	BlockNum			int64
	Specialization		int
	PoolAddr			string
}
type BalV2PoolTokBalance struct {
	TimeStamp			int64
	DateTime			string
	Amount				float64
	Balance				float64
}
type BalV2PoolToken struct {
	TokenAid			int64
	TokenAddr			string
	CurBalance			string
	Balances			[]BalV2PoolTokBalance
}
