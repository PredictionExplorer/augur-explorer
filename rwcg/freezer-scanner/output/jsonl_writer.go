// Package output provides output formatters for the freezer scanner
package output

import (
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/PredictionExplorer/augur-explorer/rwcg/freezer-scanner/decode"
)

// LogRecord represents a single log entry for output
type LogRecord struct {
	BlockNumber  uint64   `json:"blockNumber"`
	TxIndex      uint     `json:"txIndex"`
	ReceiptIndex uint     `json:"receiptIndex"`
	LogIndex     uint     `json:"logIndex"`
	Contract     string   `json:"contract"`
	Topic0       string   `json:"topic0"`
	Topics       []string `json:"topics"`
	DataKeccak   string   `json:"dataKeccak"`
	DataLen      int      `json:"dataLen"`
	DataHex      string   `json:"dataHex,omitempty"`
}

// Writer is the interface for output writers
type Writer interface {
	Write(record *LogRecord) error
	Flush() error
	Close() error
}

// JSONLWriter writes records in JSON Lines format
type JSONLWriter struct {
	w       io.Writer
	encoder *json.Encoder
	mu      sync.Mutex
	closer  io.Closer
}

// NewJSONLWriter creates a new JSONL writer
func NewJSONLWriter(path string) (*JSONLWriter, error) {
	var w io.Writer
	var closer io.Closer

	if path == "" || path == "-" {
		w = os.Stdout
	} else {
		f, err := os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("failed to create output file: %w", err)
		}
		w = f
		closer = f
	}

	return &JSONLWriter{
		w:       w,
		encoder: json.NewEncoder(w),
		closer:  closer,
	}, nil
}

func (jw *JSONLWriter) Write(record *LogRecord) error {
	jw.mu.Lock()
	defer jw.mu.Unlock()
	return jw.encoder.Encode(record)
}

func (jw *JSONLWriter) Flush() error {
	return nil // json.Encoder doesn't buffer
}

func (jw *JSONLWriter) Close() error {
	if jw.closer != nil {
		return jw.closer.Close()
	}
	return nil
}

// CSVWriter writes records in CSV format
type CSVWriter struct {
	w      *csv.Writer
	mu     sync.Mutex
	closer io.Closer
	header bool
}

// NewCSVWriter creates a new CSV writer
func NewCSVWriter(path string) (*CSVWriter, error) {
	var w io.Writer
	var closer io.Closer

	if path == "" || path == "-" {
		w = os.Stdout
	} else {
		f, err := os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("failed to create output file: %w", err)
		}
		w = f
		closer = f
	}

	cw := &CSVWriter{
		w:      csv.NewWriter(w),
		closer: closer,
	}

	// Write header
	if err := cw.w.Write([]string{
		"blockNumber", "txIndex", "receiptIndex", "logIndex",
		"contract", "topic0", "topics", "dataKeccak", "dataLen", "dataHex",
	}); err != nil {
		if closer != nil {
			closer.Close()
		}
		return nil, fmt.Errorf("failed to write CSV header: %w", err)
	}
	cw.header = true

	return cw, nil
}

func (cw *CSVWriter) Write(record *LogRecord) error {
	cw.mu.Lock()
	defer cw.mu.Unlock()

	// Convert topics array to JSON string
	topicsJSON, _ := json.Marshal(record.Topics)

	return cw.w.Write([]string{
		strconv.FormatUint(record.BlockNumber, 10),
		strconv.FormatUint(uint64(record.TxIndex), 10),
		strconv.FormatUint(uint64(record.ReceiptIndex), 10),
		strconv.FormatUint(uint64(record.LogIndex), 10),
		record.Contract,
		record.Topic0,
		string(topicsJSON),
		record.DataKeccak,
		strconv.Itoa(record.DataLen),
		record.DataHex,
	})
}

func (cw *CSVWriter) Flush() error {
	cw.mu.Lock()
	defer cw.mu.Unlock()
	cw.w.Flush()
	return cw.w.Error()
}

func (cw *CSVWriter) Close() error {
	if err := cw.Flush(); err != nil {
		return err
	}
	if cw.closer != nil {
		return cw.closer.Close()
	}
	return nil
}

// NewWriter creates a writer based on format
func NewWriter(format, path string) (Writer, error) {
	return NewWriterWithMode(format, path, false)
}

// NewWriterWithMode creates a writer with optional append mode
func NewWriterWithMode(format, path string, appendMode bool) (Writer, error) {
	switch format {
	case "jsonl", "json", "":
		return NewJSONLWriterWithMode(path, appendMode)
	case "csv":
		if appendMode {
			return NewCSVWriterAppend(path)
		}
		return NewCSVWriter(path)
	default:
		return nil, fmt.Errorf("unknown output format: %s", format)
	}
}

// NewJSONLWriterWithMode creates a JSONL writer with optional append mode
func NewJSONLWriterWithMode(path string, appendMode bool) (*JSONLWriter, error) {
	var w io.Writer
	var closer io.Closer

	if path == "" || path == "-" {
		w = os.Stdout
	} else {
		flags := os.O_CREATE | os.O_WRONLY
		if appendMode {
			flags |= os.O_APPEND
		} else {
			flags |= os.O_TRUNC
		}
		f, err := os.OpenFile(path, flags, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open output file: %w", err)
		}
		w = f
		closer = f
	}

	return &JSONLWriter{
		w:       w,
		encoder: json.NewEncoder(w),
		closer:  closer,
	}, nil
}

// NewCSVWriterAppend creates a CSV writer in append mode (no header)
func NewCSVWriterAppend(path string) (*CSVWriter, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open output file: %w", err)
	}

	return &CSVWriter{
		w:      csv.NewWriter(f),
		closer: f,
		header: true, // Skip header in append mode
	}, nil
}

// DecodedLogToRecord converts a decoded log to an output record
func DecodedLogToRecord(blockNum uint64, txIndex uint, log *decode.DecodedLog, includeData bool) *LogRecord {
	record := &LogRecord{
		BlockNumber:  blockNum,
		TxIndex:      txIndex,
		ReceiptIndex: log.ReceiptIndex,
		LogIndex:     log.LogIndex,
		Contract:     log.Address.Hex(),
		DataLen:      len(log.Data),
	}

	if len(log.Data) > 0 {
		record.DataKeccak = log.DataKeccak.Hex()
	} else {
		record.DataKeccak = common.Hash{}.Hex()
	}

	if len(log.Topics) > 0 {
		record.Topic0 = log.Topics[0].Hex()
	}

	record.Topics = make([]string, len(log.Topics))
	for i, topic := range log.Topics {
		record.Topics[i] = topic.Hex()
	}

	if includeData && len(log.Data) > 0 {
		record.DataHex = "0x" + hex.EncodeToString(log.Data)
	}

	return record
}

// LogEntryToRecord converts a decode.Log to an output record
func LogEntryToRecord(blockNum uint64, logIdx uint, log *decode.Log, includeData bool) *LogRecord {
	record := &LogRecord{
		BlockNumber:  blockNum,
		TxIndex:      uint(logIdx), // Best approximation for Arbitrum
		ReceiptIndex: uint(logIdx),
		LogIndex:     uint(logIdx),
		Contract:     log.Address.Hex(),
		DataLen:      len(log.Data),
	}

	if len(log.Data) > 0 {
		record.DataKeccak = common.BytesToHash(crypto.Keccak256(log.Data)).Hex()
	} else {
		record.DataKeccak = common.Hash{}.Hex()
	}

	if len(log.Topics) > 0 {
		record.Topic0 = log.Topics[0].Hex()
	}

	record.Topics = make([]string, len(log.Topics))
	for i, topic := range log.Topics {
		record.Topics[i] = topic.Hex()
	}

	if includeData && len(log.Data) > 0 {
		record.DataHex = "0x" + hex.EncodeToString(log.Data)
	}

	return record
}


