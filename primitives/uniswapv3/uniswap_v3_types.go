package uniswapv3

import (
)

type UniV3ContractAddrs struct {
	FactoryAddr			string
	NFTPosMgrAddr		string
}
type UniV3PoolCreated struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	Token0Addr			string
	Token1Addr			string
	Fee					string
	TickSpacing			string
	PoolAddr			string
}
type UniV3Initialize struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	SqrtPriceX96		string
	Tick				string
}
type UniV3Mint struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	SenderAddr			string
	OwnerAddr			string
	TickLower			string
	TickUpper			string
	Amount				string
	Amount0				string
	Amount1				string
}
type UniV3Collect struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	OwnerAddr			string
	RecipientAddr		string
	TickLower			string
	TickUpper			string
	Amount0				string
	Amount1				string
}
type UniV3CollectPeriphery struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	RecipientAddr		string
	TokenId				string
	Amount0				string
	Amount1				string
}
type UniV3Burn struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	OwnerAddr			string
	TickLower			string
	TickUpper			string
	Amount				string
	Amount0				string
	Amount1				string
}
type UniV3Swap struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	SenderAddr			string
	RecipientAddr		string
	Amount0				string
	Amount1				string
	Fee					string
	SqrtPriceX96		string
	Liquidity			string
	Tick				string
}
type UniV3Flash struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	SenderAddr			string
	RecipientAddr		string
	Amount0				string
	Amount1				string
	Paid0				string
	Paid1				string
}
type UniV3IncObservCardinNext struct{
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	ObservationCardinalityNextOld		uint16
	ObservationCardinalityNextNew		uint16
}
type UniV3SetFeeProtocol struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	FeeProtocol0Old		uint8
	FeeProtocol0New		uint8
	FeeProtocol1Old		uint8
	FeeProtocol1New		uint8
}
type UniV3PoolCollectProtocol struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	SenderAddr			string
	RecipientAddr		string
	Amount0				string
	Amount1				string
}
type UniV3IncreaseLiquidity struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	TokenId				string
	Liquidity			string
	Amount0				string
	Amount1				string
}
type UniV3DecreaseLiquidity struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	ContractAddr		string
	TokenId				string
	Liquidity			string
	Amount0				string
	Amount1				string
}
type UniV3DBGSwapLoop struct {
	BlockNum			int64
	TimeStamp			int64
	TxIndex				int64
	LogIndex			int64
	PoolAid				int64
	ContractAddr		string
	Tick				int64
	SqrtPrice			string
	Liquidity			string
	StepAmountIn		string
	StepAmountOut		string
	FeeGrowthGlobal		string
	FeeAmount			string
}
