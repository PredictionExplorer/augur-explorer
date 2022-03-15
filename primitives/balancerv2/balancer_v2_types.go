package balancerv2

import (

)

type BalV2PoolCreated struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	PoolAddr			string
}
type BalV2PoolRegistered struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAddr			string
	PoolId				int64
	Specialization		int64
}
type BalV2TokensRegistered struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolId				int64
	Tokens				string
	AssetManagers		string
}
type BalV2TokensDeregistered struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolId				int64
	Tokens				string
}
type BalV2InternalBalanceChanged struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	UserAddr			string
	TokenAddr			string
	Delta				string
}
type BalV2ExternalBalanceTransf struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	TokenAddr			string
	SenderAddr			string
	ReceipientAddr		string
	Amount				string
}
type BalV2Swap struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PollId				int64
	TokenInAddr			string
	TokenOutAddr		string
	AmounIn				string
	AmountOut			string
}
