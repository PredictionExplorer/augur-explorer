package store

// EthereumEventLog is the evt_log row as the readers return it: the stored
// log identity, its block/transaction context resolved to natural keys, and
// the raw RLP-encoded log payload the indexer decodes.
type EthereumEventLog struct {
	EvtId           int64
	BlockNum        int64
	TxId            int64
	ContractAid     int64
	TimeStamp       int64
	ContractAddress string
	TxHash          string
	Topic0Sig       string
	RlpLog          []byte
}
