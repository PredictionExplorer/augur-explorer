// Package rlpcorpus defines the strict, archive-compatible JSONL format used
// to preserve representative Ethereum event-log RLP payloads.
package rlpcorpus

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

const maxLine = 2 << 20

// Entry is one archive-compatible event-log row. Its JSONL shape mirrors
// arch_evtlog; LogRLP is 0x-prefixed hex rather than PostgreSQL bytea text so
// a corpus is portable and reviewable.
type Entry struct {
	Project         string `json:"project"`
	BlockNum        int64  `json:"blockNum"`
	EventID         int64  `json:"eventId"`
	LogIndex        int    `json:"logIndex"`
	TxHash          string `json:"txHash"`
	ContractAddress string `json:"contractAddress"`
	Topic0Sig       string `json:"topic0Sig"`
	LogRLP          string `json:"logRlp"`
}

// RLPBytes validates and decodes the entry's 0x-prefixed byte payload.
func (e Entry) RLPBytes() ([]byte, error) {
	return decodeHex("logRlp", e.LogRLP, -1)
}

// DecodedLog validates and decodes the entry's canonical RLP payload.
func (e Entry) DecodedLog() (types.Log, error) {
	raw, err := e.RLPBytes()
	if err != nil {
		return types.Log{}, err
	}
	var lg types.Log
	if err := rlp.DecodeBytes(raw, &lg); err != nil {
		return types.Log{}, fmt.Errorf("logRlp is not a types.Log RLP: %w", err)
	}
	return lg, nil
}

// Load parses a strict JSONL corpus. Blank lines are ignored; unknown fields,
// malformed identities/RLP, inconsistent derived fields and duplicate
// (txHash, logIndex) identities are rejected with a line number.
func Load(r io.Reader) ([]Entry, error) {
	if r == nil {
		return nil, errors.New("RLP corpus reader is required")
	}
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 64*1024), maxLine)
	entries := make([]Entry, 0)
	seen := make(map[string]int)
	closedTransactions := make(map[string]int)
	var (
		currentTransaction string
		currentGroupLine   int
		lastLogIndex       int
	)
	for line := 1; scanner.Scan(); line++ {
		data := bytes.TrimSpace(scanner.Bytes())
		if len(data) == 0 {
			continue
		}
		dec := json.NewDecoder(bytes.NewReader(data))
		dec.DisallowUnknownFields()
		var entry Entry
		if err := dec.Decode(&entry); err != nil {
			return nil, fmt.Errorf("RLP corpus line %d: decode JSON: %w", line, err)
		}
		if err := requireJSONEOF(dec); err != nil {
			return nil, fmt.Errorf("RLP corpus line %d: %w", line, err)
		}
		if err := Validate(entry); err != nil {
			return nil, fmt.Errorf("RLP corpus line %d: %w", line, err)
		}
		key := strings.ToLower(entry.TxHash) + "/" + strconv.Itoa(entry.LogIndex)
		if first, ok := seen[key]; ok {
			return nil, fmt.Errorf(
				"RLP corpus line %d: duplicate chain identity %s (first seen on line %d)",
				line,
				key,
				first,
			)
		}
		seen[key] = line

		txKey := strings.ToLower(entry.TxHash)
		switch {
		case currentTransaction == "":
			currentTransaction = txKey
			currentGroupLine = line
			lastLogIndex = entry.LogIndex
		case txKey != currentTransaction:
			closedTransactions[currentTransaction] = currentGroupLine
			if first, ok := closedTransactions[txKey]; ok {
				return nil, fmt.Errorf(
					"RLP corpus line %d: transaction %s is not contiguous (first group starts on line %d)",
					line,
					txKey,
					first,
				)
			}
			currentTransaction = txKey
			currentGroupLine = line
			lastLogIndex = entry.LogIndex
		case entry.LogIndex <= lastLogIndex:
			return nil, fmt.Errorf(
				"RLP corpus line %d: transaction %s logIndex %d is not strictly after %d",
				line,
				txKey,
				entry.LogIndex,
				lastLogIndex,
			)
		default:
			lastLogIndex = entry.LogIndex
		}
		entries = append(entries, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read RLP corpus: %w", err)
	}
	return entries, nil
}

// Write validates entries and writes deterministic one-object-per-line JSON.
// Loading the output applies the same duplicate-identity policy.
func Write(w io.Writer, entries []Entry) error {
	if w == nil {
		return errors.New("RLP corpus writer is required")
	}
	var validated bytes.Buffer
	enc := json.NewEncoder(&validated)
	enc.SetEscapeHTML(false)
	for i, entry := range entries {
		if err := Validate(entry); err != nil {
			return fmt.Errorf("RLP corpus entry %d: %w", i, err)
		}
		if err := enc.Encode(entry); err != nil {
			return fmt.Errorf("encode RLP corpus entry %d: %w", i, err)
		}
	}
	if _, err := Load(bytes.NewReader(validated.Bytes())); err != nil {
		return err
	}
	if _, err := io.Copy(w, &validated); err != nil {
		return fmt.Errorf("write RLP corpus: %w", err)
	}
	return nil
}

// Validate checks one entry's archive identity and RLP-derived fields.
func Validate(entry Entry) error {
	switch entry.Project {
	case "cosmicgame", "randomwalk":
	default:
		return fmt.Errorf("project %q must be cosmicgame or randomwalk", entry.Project)
	}
	if entry.BlockNum < 0 {
		return fmt.Errorf("blockNum %d is negative", entry.BlockNum)
	}
	if entry.EventID < 0 {
		return fmt.Errorf("eventId %d is negative", entry.EventID)
	}
	if entry.LogIndex < 0 {
		return fmt.Errorf("logIndex %d is negative", entry.LogIndex)
	}
	if entry.LogIndex > math.MaxInt32 {
		return fmt.Errorf("logIndex %d exceeds PostgreSQL integer range", entry.LogIndex)
	}
	if _, err := decodeHex("txHash", entry.TxHash, common.HashLength); err != nil {
		return err
	}
	if !common.IsHexAddress(entry.ContractAddress) {
		return fmt.Errorf("contractAddress %q is not a 20-byte hex address", entry.ContractAddress)
	}
	if entry.Topic0Sig != "" {
		topic, err := hex.DecodeString(entry.Topic0Sig)
		if err != nil {
			return fmt.Errorf("topic0Sig is invalid hex: %w", err)
		}
		if len(topic) != 4 {
			return fmt.Errorf("topic0Sig is %d bytes, want 4", len(topic))
		}
	}
	lg, err := entry.DecodedLog()
	if err != nil {
		return err
	}
	if lg.Address != common.HexToAddress(entry.ContractAddress) {
		return fmt.Errorf(
			"contractAddress %s disagrees with logRlp address %s",
			entry.ContractAddress,
			lg.Address.Hex(),
		)
	}
	if got := toolutil.Topic0Sig(&lg); !strings.EqualFold(got, entry.Topic0Sig) {
		return fmt.Errorf("topic0Sig %q disagrees with logRlp topic %q", entry.Topic0Sig, got)
	}
	return nil
}

func requireJSONEOF(dec *json.Decoder) error {
	var extra any
	err := dec.Decode(&extra)
	if errors.Is(err, io.EOF) {
		return nil
	}
	if err == nil {
		return errors.New("multiple JSON values")
	}
	return fmt.Errorf("trailing JSON: %w", err)
}

func decodeHex(field, value string, byteLen int) ([]byte, error) {
	if !strings.HasPrefix(value, "0x") {
		return nil, fmt.Errorf("%s must be 0x-prefixed hex", field)
	}
	raw, err := hex.DecodeString(value[2:])
	if err != nil {
		return nil, fmt.Errorf("%s is invalid hex: %w", field, err)
	}
	if byteLen >= 0 && len(raw) != byteLen {
		return nil, fmt.Errorf("%s is %d bytes, want %d", field, len(raw), byteLen)
	}
	if field == "logRlp" && len(raw) == 0 {
		return nil, errors.New("logRlp is empty")
	}
	return raw, nil
}
