package toolutil

import (
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

// IsHexTxHash reports whether s is a 0x-prefixed 32-byte transaction hash.
func IsHexTxHash(s string) bool {
	_, err := ethcommon.ParseHexOrString(strings.TrimSpace(s))
	if err != nil {
		return false
	}
	return len(strings.TrimSpace(s)) == 2+2*ethcommon.HashLength
}

// NormalizeTxHash returns canonical 0x + 64-hex for a transaction hash.
func NormalizeTxHash(hash string) string {
	return ethcommon.HexToHash(hash).Hex()
}
