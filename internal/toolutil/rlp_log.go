package toolutil

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// Topic0Sig returns the first 8 hex chars of topic0 (event signature), matching evt_log.topic0_sig.
func Topic0Sig(lg *types.Log) string {
	if len(lg.Topics) == 0 {
		return ""
	}
	full := lg.Topics[0].Hex()[2:]
	if len(full) >= 8 {
		return full[:8]
	}
	return full
}

// EncodeLogRLP RLP-encodes a log the same way as Layer1 ETL (evt_log.log_rlp).
func EncodeLogRLP(lg *types.Log) ([]byte, error) {
	return rlp.EncodeToBytes(lg)
}

// NormalizeAddr returns the canonical hex address for comparisons.
func NormalizeAddr(addr string) string {
	return common.HexToAddress(addr).Hex()
}

// LogRLPEqual compares stored log RLP with the canonical encoding of lg.
func LogRLPEqual(dbBytes []byte, lg *types.Log) (bool, error) {
	encoded, err := EncodeLogRLP(lg)
	if err != nil {
		return false, err
	}
	if len(encoded) != len(dbBytes) {
		return false, nil
	}
	for i := range encoded {
		if encoded[i] != dbBytes[i] {
			return false, nil
		}
	}
	return true, nil
}

// Topic0SigHex is a debug helper: full topic0 without 0x prefix.
func Topic0SigHex(lg *types.Log) string {
	if len(lg.Topics) == 0 {
		return ""
	}
	return hex.EncodeToString(lg.Topics[0].Bytes())
}
