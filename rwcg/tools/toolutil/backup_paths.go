package toolutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

)

// TxRLPPath returns the path for a backed-up transaction RLP blob.
func TxRLPPath(outputDir string, blockNum uint64, txHash string) string {
	return filepath.Join(outputDir, fmt.Sprintf("%d", blockNum), txHash+"_tx.rlp")
}

// ReceiptRLPPath returns the path for a backed-up receipt RLP blob.
func ReceiptRLPPath(outputDir string, blockNum uint64, txHash string) string {
	return filepath.Join(outputDir, fmt.Sprintf("%d", blockNum), txHash+"_receipt.rlp")
}

// BackupTxOnDisk maps tx_hash -> block_num for every *_tx.rlp under outputDir.
func BackupTxOnDisk(outputDir string) (map[string]uint64, error) {
	out := make(map[string]uint64)
	err := filepath.WalkDir(outputDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), "_tx.rlp") {
			return nil
		}
		blockNum, txHash, ok := parseBackupFile(path, outputDir, "_tx.rlp")
		if !ok {
			return nil
		}
		out[txHash] = blockNum
		return nil
	})
	return out, err
}

func parseBackupFile(path, outputDir, suffix string) (blockNum uint64, txHash string, ok bool) {
	rel, err := filepath.Rel(outputDir, path)
	if err != nil {
		return 0, "", false
	}
	parts := strings.Split(rel, string(os.PathSeparator))
	if len(parts) != 2 || !strings.HasSuffix(parts[1], suffix) {
		return 0, "", false
	}
	bn, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return 0, "", false
	}
	name := strings.TrimSuffix(parts[1], suffix)
	if !IsHexTxHash(name) {
		return 0, "", false
	}
	return bn, NormalizeTxHash(name), true
}
