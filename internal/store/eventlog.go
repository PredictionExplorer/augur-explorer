package store

// EthereumEventLog is the evt_log row as the readers return it: the stored
// log identity, its block/transaction context resolved to natural keys, and
// the raw RLP-encoded log payload the indexer decodes.
type EthereumEventLog struct {
	EvtID           int64
	BlockNum        int64
	TxID            int64
	ContractAid     int64
	TimeStamp       int64
	ContractAddress string
	TxHash          string
	Topic0Sig       string
	RlpLog          []byte
}
