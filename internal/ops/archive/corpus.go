package archive

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/rlpcorpus"
)

// CorpusExportOptions selects complete archived transactions for one RLP
// replay corpus. Transaction order is preserved; logs within each transaction
// are emitted in canonical log-index order.
type CorpusExportOptions struct {
	Project  string
	TxHashes []string
}

// CorpusExportStats summarizes one archive-to-JSONL corpus export.
type CorpusExportStats struct {
	Transactions int
	EventLogs    int
}

// ExportCorpus writes strict archive-compatible JSONL for every sibling event
// log belonging to the explicitly selected transactions. Database and
// validation errors are reported before any bytes are written.
func ExportCorpus(
	ctx context.Context,
	db Querier,
	options CorpusExportOptions,
	w io.Writer,
) (CorpusExportStats, error) {
	var stats CorpusExportStats
	if db == nil {
		return stats, errors.New("archive corpus export: database is required")
	}
	if w == nil {
		return stats, errors.New("archive corpus export: writer is required")
	}
	project, err := corpusProject(options.Project)
	if err != nil {
		return stats, err
	}
	txHashes, err := corpusTxHashes(options.TxHashes)
	if err != nil {
		return stats, err
	}

	entries := make([]rlpcorpus.Entry, 0)
	for _, txHash := range txHashes {
		if err := ctx.Err(); err != nil {
			return stats, err
		}
		rows, err := db.Query(ctx, `
			SELECT
				block_num,
				COALESCE(evt_id, 0),
				log_index,
				BTRIM(tx_hash),
				BTRIM(contract_addr),
				BTRIM(topic0_sig),
				log_rlp
			FROM arch_evtlog
			WHERE LOWER(BTRIM(tx_hash)) = LOWER($1)
			ORDER BY log_index, evt_id NULLS LAST
		`, txHash)
		if err != nil {
			return stats, fmt.Errorf("query archive transaction %s: %w", txHash, err)
		}

		start := len(entries)
		for rows.Next() {
			var (
				entry rlpcorpus.Entry
				raw   []byte
			)
			if err := rows.Scan(
				&entry.BlockNum,
				&entry.EventID,
				&entry.LogIndex,
				&entry.TxHash,
				&entry.ContractAddress,
				&entry.Topic0Sig,
				&raw,
			); err != nil {
				rows.Close()
				return stats, fmt.Errorf("scan archive transaction %s: %w", txHash, err)
			}
			entry.Project = project
			entry.LogRLP = "0x" + hex.EncodeToString(raw)
			if err := rlpcorpus.Validate(entry); err != nil {
				rows.Close()
				return stats, fmt.Errorf(
					"validate archive transaction %s log index %d: %w",
					txHash,
					entry.LogIndex,
					err,
				)
			}
			if !strings.EqualFold(entry.TxHash, txHash) {
				rows.Close()
				return stats, fmt.Errorf(
					"archive transaction query %s returned row for %s",
					txHash,
					entry.TxHash,
				)
			}
			entries = append(entries, entry)
		}
		if err := rows.Err(); err != nil {
			rows.Close()
			return stats, fmt.Errorf("iterate archive transaction %s: %w", txHash, err)
		}
		rows.Close()
		if len(entries) == start {
			return stats, fmt.Errorf("archive transaction %s has no event logs", txHash)
		}
		stats.Transactions++
	}

	if err := rlpcorpus.Write(w, entries); err != nil {
		return CorpusExportStats{}, fmt.Errorf("write archive RLP corpus: %w", err)
	}
	stats.EventLogs = len(entries)
	return stats, nil
}

func corpusProject(project string) (string, error) {
	projects, err := ResolveProjects(strings.TrimSpace(project))
	if err != nil {
		return "", err
	}
	if len(projects) != 1 {
		return "", errors.New("archive corpus export requires exactly one project")
	}
	return projects[0], nil
}

func corpusTxHashes(values []string) ([]string, error) {
	if len(values) == 0 {
		return nil, errors.New("archive corpus export requires at least one transaction hash")
	}
	out := make([]string, 0, len(values))
	seen := make(map[string]struct{}, len(values))
	for i, value := range values {
		value = strings.TrimSpace(value)
		if !strings.HasPrefix(value, "0x") {
			return nil, fmt.Errorf("transaction hash %d must be 0x-prefixed", i)
		}
		raw, err := hex.DecodeString(value[2:])
		if err != nil || len(raw) != 32 {
			return nil, fmt.Errorf("transaction hash %d is not 32-byte hex", i)
		}
		canonical := "0x" + hex.EncodeToString(raw)
		key := strings.ToLower(canonical)
		if _, ok := seen[key]; ok {
			return nil, fmt.Errorf("duplicate transaction hash %s", canonical)
		}
		seen[key] = struct{}{}
		out = append(out, canonical)
	}
	return out, nil
}
