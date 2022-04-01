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
type BalV2PoolBalanceChanged struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	PoolId				string
	LiqProvAddr			string
	Tokens				string
	Deltas				string
	ProtocolFeeAmounts	string
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
	ContractAid			int64
	TokenInAid			int64
	TokenOutAid			int64
	PoolId				string
	ContractAddr		string
	TokenInAddr			string
	TokenOutAddr		string
	AmountIn			string
	AmountOut			string
}
type BalV2SwapFeePercentageChanged struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolId				string
	ContractAddr		string
	SwapFeePercentage	string
}
type BalV2PoolBalanceManaged struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolId				string
	ContractAddr		string
	AssetManagerAddr	string
	TokenAddr			string
	CashDelta			string
	ManagedDelta		string
}
type BalV2FlashLoan struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	RecipientAddr		string
	TokenAddr			string
	Amount				string
	FeeAmount			string
}
type BalV2ContractAddrs struct {
	VaultAddr			string
	FactoryAddr			string
}
type BalV2SwapHist struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAid			int64
	PoolId				string
	SwapFee				string
	ProtocolFee			string
	AccumSwapFee		string
	AccumProtoFee		string
}
