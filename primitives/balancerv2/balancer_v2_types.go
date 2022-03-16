package balancerv2

import (

)

type BalV2PoolCreated struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	PoolAid				int64
	PoolAddr			string
}
type BalV2PoolRegistered struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	PoolAddr			string
	PoolId				string
	Specialization		int64
}
type BalV2TokensRegistered struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	PoolId				string
	Tokens				string
	AssetManagers		string
}
type BalV2TokensDeregistered struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	PoolId				string
	Tokens				string
}
type BalV2InternalBalanceChanged struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	UserAddr			string
	TokenAddr			string
	Delta				string
}
type BalV2ExternalBalanceTransfer struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	TokenAddr			string
	SenderAddr			string
	RecipientAddr		string
	Amount				string
}
type BalV2Swap struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolId				string
	ContractAddr		string
	TokenInAddr			string
	TokenOutAddr		string
	AmountIn			string
	AmountOut			string
}
