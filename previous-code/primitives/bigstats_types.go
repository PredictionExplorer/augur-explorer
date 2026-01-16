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
	TimeStampTo					int64
	BlockNum					int64
	NumUniqHumanAccts			int64
	NumUniqContractAccts		int64
	EthTransferred				float64
	TxFeesEth					float64	// fees in ETH
	TxFeesUsd					float64
}
type BigStatsTimeframeRange struct {
	TsIni						int64
	TsFin						int64
}
