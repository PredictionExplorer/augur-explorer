package primitives

type AddrStatsLog struct {
	BlockNum			int64
	TxIndex				int64
	Aid					int64		// Address ID
}
type TxShort struct {
	BlockNum			int64
	TxIndex				int64
	TxFee				string
}
type BigStatRec struct {
	TimeStamp					int64
	DateTime					string
	NumUniqHumanAccts			int64
	NumUniqContractAccts		int64
	EthTransferred				float64
	TxFeesEth					float64	// fees in ETH
}
