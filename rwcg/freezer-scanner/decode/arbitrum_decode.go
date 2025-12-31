// Package decode provides decoders for Arbitrum Nitro freezer receipts
package decode

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/snappy"
)

// Log represents an Ethereum log entry
type Log struct {
	Address common.Address
	Topics  []common.Hash
	Data    []byte
}

// DecodeArbitrumLog decodes a single log from raw RLP
func DecodeArbitrumLog(rawLog []byte) (*Log, error) {
	var fields []rlp.RawValue
	if err := rlp.DecodeBytes(rawLog, &fields); err != nil {
		return nil, err
	}
	if len(fields) < 3 {
		return nil, fmt.Errorf("expected 3 fields, got %d", len(fields))
	}

	log := &Log{}
	if err := rlp.DecodeBytes(fields[0], &log.Address); err != nil {
		return nil, fmt.Errorf("decode address: %w", err)
	}
	if err := rlp.DecodeBytes(fields[1], &log.Topics); err != nil {
		return nil, fmt.Errorf("decode topics: %w", err)
	}
	if err := rlp.DecodeBytes(fields[2], &log.Data); err != nil {
		return nil, fmt.Errorf("decode data: %w", err)
	}
	return log, nil
}

// DecodeArbitrumReceipts decodes snappy-compressed Arbitrum Nitro receipts
// and returns all logs from all receipts in the block
func DecodeArbitrumReceipts(compressedData []byte) ([]*Log, error) {
	// Snappy decompress
	data, err := snappy.Decode(nil, compressedData)
	if err != nil {
		return nil, fmt.Errorf("snappy decode: %w", err)
	}

	// Decode as list of receipts
	var rawReceipts []rlp.RawValue
	if err := rlp.DecodeBytes(data, &rawReceipts); err != nil {
		return nil, fmt.Errorf("decode receipts list: %w", err)
	}

	var allLogs []*Log
	for _, rawReceipt := range rawReceipts {
		// Each receipt has 7 fields in Arbitrum Nitro format
		var fields []rlp.RawValue
		if err := rlp.DecodeBytes(rawReceipt, &fields); err != nil || len(fields) < 7 {
			continue
		}

		// Field 6 is the logs list
		var rawLogs []rlp.RawValue
		if err := rlp.DecodeBytes(fields[6], &rawLogs); err != nil {
			continue
		}

		for _, rawLog := range rawLogs {
			log, err := DecodeArbitrumLog(rawLog)
			if err == nil {
				allLogs = append(allLogs, log)
			}
		}
	}

	return allLogs, nil
}


