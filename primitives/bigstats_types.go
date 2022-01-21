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
