package testutil

import (
	"bufio"
	"bytes"
	"context"
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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

const maxRLPCorpusLine = 2 << 20

// RLPCorpusEntry is one archive-compatible event-log row. Its JSONL shape
// mirrors arch_evtlog; logRlp is 0x-prefixed hex rather than PostgreSQL bytea
// text so a corpus is portable and reviewable.
type RLPCorpusEntry struct {
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
func (e RLPCorpusEntry) RLPBytes() ([]byte, error) {
	return decodeCorpusHex("logRlp", e.LogRLP, -1)
}

// DecodedLog validates and decodes the entry's canonical RLP payload.
func (e RLPCorpusEntry) DecodedLog() (types.Log, error) {
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

// LoadRLPCorpus parses a strict JSONL corpus. Blank lines are ignored;
// unknown fields, malformed identities/RLP, inconsistent derived fields and
// duplicate (txHash, logIndex) identities are rejected with a line number.
func LoadRLPCorpus(r io.Reader) ([]RLPCorpusEntry, error) {
	if r == nil {
		return nil, errors.New("RLP corpus reader is required")
	}
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 64*1024), maxRLPCorpusLine)
	entries := make([]RLPCorpusEntry, 0)
	seen := make(map[string]int)
	for line := 1; scanner.Scan(); line++ {
		data := bytes.TrimSpace(scanner.Bytes())
		if len(data) == 0 {
			continue
		}
		dec := json.NewDecoder(bytes.NewReader(data))
		dec.DisallowUnknownFields()
		var entry RLPCorpusEntry
		if err := dec.Decode(&entry); err != nil {
			return nil, fmt.Errorf("RLP corpus line %d: decode JSON: %w", line, err)
		}
		if err := requireJSONEOF(dec); err != nil {
			return nil, fmt.Errorf("RLP corpus line %d: %w", line, err)
		}
		if err := validateRLPCorpusEntry(entry); err != nil {
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
		entries = append(entries, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read RLP corpus: %w", err)
	}
	return entries, nil
}

// WriteRLPCorpus validates entries and writes deterministic one-object-per-line
// JSON. Loading the output applies the same duplicate-identity policy.
func WriteRLPCorpus(w io.Writer, entries []RLPCorpusEntry) error {
	if w == nil {
		return errors.New("RLP corpus writer is required")
	}
	var validated bytes.Buffer
	enc := json.NewEncoder(&validated)
	enc.SetEscapeHTML(false)
	for i, entry := range entries {
		if err := validateRLPCorpusEntry(entry); err != nil {
			return fmt.Errorf("RLP corpus entry %d: %w", i, err)
		}
		if err := enc.Encode(entry); err != nil {
			return fmt.Errorf("encode RLP corpus entry %d: %w", i, err)
		}
	}
	if _, err := LoadRLPCorpus(bytes.NewReader(validated.Bytes())); err != nil {
		return err
	}
	if _, err := io.Copy(w, &validated); err != nil {
		return fmt.Errorf("write RLP corpus: %w", err)
	}
	return nil
}

// InstallRLPCorpus inserts archive-compatible layer-1 scaffolding and the
// corpus's exact log_rlp bytes into a test Store. It deliberately bypasses
// InsertEventLog's re-encoding so integration tests exercise exported bytes.
// The returned evt_log IDs preserve corpus order; callers dispatch them only
// after all sibling logs have been installed.
func InstallRLPCorpus(ctx context.Context, st *store.Store, entries []RLPCorpusEntry) ([]int64, error) {
	if st == nil || st.Pool() == nil {
		return nil, errors.New("RLP corpus Store is required")
	}
	if err := WriteRLPCorpus(io.Discard, entries); err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(entries))
	for i, entry := range entries {
		if err := ctx.Err(); err != nil {
			return ids, err
		}
		blockHash := crypto.Keccak256Hash(fmt.Appendf(nil, "rlp-corpus:block:%d", entry.BlockNum))
		parentHash := crypto.Keccak256Hash(fmt.Appendf(nil, "rlp-corpus:block:%d", entry.BlockNum-1))
		if _, err := st.Pool().Exec(ctx, `INSERT INTO block (
				block_num, ts, block_hash, parent_hash
			) VALUES ($1, TO_TIMESTAMP($2), $3, $4)
			ON CONFLICT (block_num) DO NOTHING`,
			entry.BlockNum,
			entry.BlockNum,
			blockHash.Hex(),
			parentHash.Hex(),
		); err != nil {
			return ids, fmt.Errorf("install RLP corpus entry %d block: %w", i, err)
		}
		txID, err := st.InsertMinimalTransaction(ctx, entry.TxHash, entry.BlockNum)
		if err != nil {
			return ids, fmt.Errorf("install RLP corpus entry %d transaction: %w", i, err)
		}
		contractAID, err := st.LookupOrCreateAddress(
			ctx,
			common.HexToAddress(entry.ContractAddress).Hex(),
			entry.BlockNum,
			txID,
		)
		if err != nil {
			return ids, fmt.Errorf("install RLP corpus entry %d contract: %w", i, err)
		}
		raw, err := entry.RLPBytes()
		if err != nil {
			return ids, fmt.Errorf("install RLP corpus entry %d: %w", i, err)
		}
		var evtID int64
		err = st.Pool().QueryRow(ctx, `INSERT INTO evt_log (
				block_num, tx_id, contract_aid, topic0_sig, log_index, log_rlp
			) VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id`,
			entry.BlockNum,
			txID,
			contractAID,
			entry.Topic0Sig,
			entry.LogIndex,
			raw,
		).Scan(&evtID)
		if err != nil {
			return ids, fmt.Errorf("install RLP corpus entry %d event log: %w", i, err)
		}
		ids = append(ids, evtID)
	}
	return ids, nil
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

func validateRLPCorpusEntry(entry RLPCorpusEntry) error {
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
	if _, err := decodeCorpusHex("txHash", entry.TxHash, common.HashLength); err != nil {
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

func decodeCorpusHex(field, value string, byteLen int) ([]byte, error) {
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
