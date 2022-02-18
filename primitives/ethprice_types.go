package primitives

type EthpriceSwap struct {
	TxHash				string
	TimeStamp			int64
	BlockNum			int64
	TxIdx				int32
	LogIdx				int32
	TokenCode			int8
	Sender				string
	Recipient			string
	Amount0				string
	Amount1				string
	SqrtPrice			string
	Liquidity			string
	Tick				string
	EthUsdPrice			float64
}
