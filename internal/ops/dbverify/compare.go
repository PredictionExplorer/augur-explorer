// Package dbverify contains context-aware database comparison loaders and
// deterministic, pure comparison/reporting helpers used by opsctl db.
package dbverify

import (
	"encoding/hex"
	"fmt"
	"slices"
	"strings"
)

// EventRecord is one evt_log row reduced to the fields used by db verify.
type EventRecord struct {
	BlockNum  int64
	TxHash    string
	LogRLPHex string
}

// TransactionRecord is one transaction row reduced to compared fields.
type TransactionRecord struct {
	BlockNum int64
	TxHash   string
	GasUsed  int64
	NumLogs  int
}

// BlockRecord is one block row reduced to compared fields.
type BlockRecord struct {
	BlockNum   int64
	BlockHash  string
	ParentHash string
	NumTx      int64
}

// EventLogRecord is the detailed evt_log representation used by evtlog-diff.
type EventLogRecord struct {
	BlockNum        int64
	TxHash          string
	ContractAddress string
	Topic0Sig       string
	LogRLP          []byte
}

// RecordIssue describes a missing or extra keyed record.
type RecordIssue struct {
	Key     string
	Summary string
}

// FieldMismatch describes one differing field on a record with a shared key.
type FieldMismatch struct {
	Field     string
	Primary   string
	Secondary string
}

// RecordMismatch describes a shared-key record with differing fields.
type RecordMismatch struct {
	Key     string
	Summary string
	Fields  []FieldMismatch
}

// Comparison is a deterministic set/content comparison. Issue slices are
// sorted by key and capped by ReportLimit; totals always describe all issues.
type Comparison struct {
	PrimaryCount   int
	SecondaryCount int
	MissingTotal   int
	ExtraTotal     int
	MismatchTotal  int
	Missing        []RecordIssue
	Extra          []RecordIssue
	Mismatched     []RecordMismatch
	ReportLimit    int
}

// Matched reports whether no category contains a difference.
func (c Comparison) Matched() bool {
	return c.MissingTotal == 0 && c.ExtraTotal == 0 && c.MismatchTotal == 0
}

// CompareEvents compares records keyed by log_rlp content.
func CompareEvents(primary, secondary map[string]EventRecord, reportLimit int) Comparison {
	return compareMaps(
		len(primary),
		len(secondary),
		sortedKeys(primary),
		sortedKeys(secondary),
		reportLimit,
		func(key string) string {
			record := primary[key]
			return fmt.Sprintf("block=%d tx=%s rlp=%s...", record.BlockNum, record.TxHash, prefix(key, 16))
		},
		func(key string) string {
			record := secondary[key]
			return fmt.Sprintf("block=%d tx=%s rlp=%s...", record.BlockNum, record.TxHash, prefix(key, 16))
		},
		func(key string) (string, []FieldMismatch) {
			left, right := primary[key], secondary[key]
			fields := make([]FieldMismatch, 0, 2)
			appendFieldMismatch(&fields, "block_num", left.BlockNum, right.BlockNum)
			appendFieldMismatch(&fields, "tx_hash", left.TxHash, right.TxHash)
			return fmt.Sprintf("block=%d tx=%s", left.BlockNum, left.TxHash), fields
		},
	)
}

// CompareTransactions compares records keyed by tx_hash.
func CompareTransactions(primary, secondary map[string]TransactionRecord, reportLimit int) Comparison {
	return compareMaps(
		len(primary),
		len(secondary),
		sortedKeys(primary),
		sortedKeys(secondary),
		reportLimit,
		func(key string) string {
			record := primary[key]
			return fmt.Sprintf("tx=%s block=%d", key, record.BlockNum)
		},
		func(key string) string {
			record := secondary[key]
			return fmt.Sprintf("tx=%s block=%d", key, record.BlockNum)
		},
		func(key string) (string, []FieldMismatch) {
			left, right := primary[key], secondary[key]
			fields := make([]FieldMismatch, 0, 3)
			appendFieldMismatch(&fields, "block_num", left.BlockNum, right.BlockNum)
			appendFieldMismatch(&fields, "gas_used", left.GasUsed, right.GasUsed)
			appendFieldMismatch(&fields, "num_logs", left.NumLogs, right.NumLogs)
			return fmt.Sprintf("tx=%s block=%d", key, left.BlockNum), fields
		},
	)
}

// CompareBlocks compares records keyed by block_hash.
func CompareBlocks(primary, secondary map[string]BlockRecord, reportLimit int) Comparison {
	return compareMaps(
		len(primary),
		len(secondary),
		sortedKeys(primary),
		sortedKeys(secondary),
		reportLimit,
		func(key string) string {
			record := primary[key]
			return fmt.Sprintf("block_num=%d hash=%s", record.BlockNum, key)
		},
		func(key string) string {
			record := secondary[key]
			return fmt.Sprintf("block_num=%d hash=%s", record.BlockNum, key)
		},
		func(key string) (string, []FieldMismatch) {
			left, right := primary[key], secondary[key]
			fields := make([]FieldMismatch, 0, 3)
			appendFieldMismatch(&fields, "block_num", left.BlockNum, right.BlockNum)
			appendFieldMismatch(&fields, "parent_hash", left.ParentHash, right.ParentHash)
			appendFieldMismatch(&fields, "num_tx", left.NumTx, right.NumTx)
			return fmt.Sprintf("block_num=%d hash=%s", left.BlockNum, key), fields
		},
	)
}

// IndexEventLogsByRLP builds the content-based identity map used by
// evtlog-diff. Repeated RLP values receive stable occurrence suffixes so
// duplicate records retain their multiplicity.
func IndexEventLogsByRLP(events []EventLogRecord) map[string]EventLogRecord {
	ordered := append([]EventLogRecord(nil), events...)
	slices.SortStableFunc(ordered, func(a, b EventLogRecord) int {
		return strings.Compare(eventLogSortKey(a), eventLogSortKey(b))
	})

	indexed := make(map[string]EventLogRecord, len(events))
	occurrences := make(map[string]int)
	for _, event := range ordered {
		baseKey := hex.EncodeToString(event.LogRLP)
		occurrences[baseKey]++
		indexed[occurrenceKey(baseKey, occurrences[baseKey])] = event
	}
	return indexed
}

// CompareDetailedEventLogs performs the pure field-level evtlog comparison.
func CompareDetailedEventLogs(primary, secondary []EventLogRecord, reportLimit int) Comparison {
	primaryByRLP := IndexEventLogsByRLP(primary)
	secondaryByRLP := IndexEventLogsByRLP(secondary)
	return compareMaps(
		len(primary),
		len(secondary),
		sortedKeys(primaryByRLP),
		sortedKeys(secondaryByRLP),
		reportLimit,
		func(key string) string {
			event := primaryByRLP[key]
			return fmt.Sprintf("block=%d tx=%s topic0=%s", event.BlockNum, event.TxHash, event.Topic0Sig)
		},
		func(key string) string {
			event := secondaryByRLP[key]
			return fmt.Sprintf("block=%d tx=%s topic0=%s", event.BlockNum, event.TxHash, event.Topic0Sig)
		},
		func(key string) (string, []FieldMismatch) {
			left, right := primaryByRLP[key], secondaryByRLP[key]
			fields := make([]FieldMismatch, 0, 4)
			appendFieldMismatch(&fields, "block_num", left.BlockNum, right.BlockNum)
			appendFieldMismatch(&fields, "tx_hash", left.TxHash, right.TxHash)
			appendFieldMismatch(&fields, "contract_addr", left.ContractAddress, right.ContractAddress)
			appendFieldMismatch(&fields, "topic0_sig", left.Topic0Sig, right.Topic0Sig)
			return fmt.Sprintf("block=%d tx=%s", left.BlockNum, left.TxHash), fields
		},
	)
}

func eventLogSortKey(event EventLogRecord) string {
	return fmt.Sprintf(
		"%x\x00%d\x00%q\x00%q\x00%q",
		event.LogRLP,
		event.BlockNum,
		event.TxHash,
		event.ContractAddress,
		event.Topic0Sig,
	)
}

func occurrenceKey(base string, occurrence int) string {
	if occurrence <= 1 {
		return base
	}
	return fmt.Sprintf("%s#%d", base, occurrence)
}

type (
	summaryFunc  func(key string) string
	mismatchFunc func(key string) (summary string, fields []FieldMismatch)
)

func compareMaps(
	primaryCount int,
	secondaryCount int,
	primaryKeys []string,
	secondaryKeys []string,
	reportLimit int,
	primarySummary summaryFunc,
	secondarySummary summaryFunc,
	mismatch mismatchFunc,
) Comparison {
	limit := normalizedLimit(reportLimit)
	comparison := Comparison{
		PrimaryCount:   primaryCount,
		SecondaryCount: secondaryCount,
		ReportLimit:    reportLimit,
	}
	primarySet := keysSet(primaryKeys)
	secondarySet := keysSet(secondaryKeys)

	for _, key := range primaryKeys {
		if _, exists := secondarySet[key]; !exists {
			comparison.MissingTotal++
			if len(comparison.Missing) < limit {
				comparison.Missing = append(comparison.Missing, RecordIssue{
					Key:     key,
					Summary: primarySummary(key),
				})
			}
			continue
		}
		summary, fields := mismatch(key)
		if len(fields) == 0 {
			continue
		}
		comparison.MismatchTotal++
		if len(comparison.Mismatched) < limit {
			comparison.Mismatched = append(comparison.Mismatched, RecordMismatch{
				Key:     key,
				Summary: summary,
				Fields:  fields,
			})
		}
	}
	for _, key := range secondaryKeys {
		if _, exists := primarySet[key]; exists {
			continue
		}
		comparison.ExtraTotal++
		if len(comparison.Extra) < limit {
			comparison.Extra = append(comparison.Extra, RecordIssue{
				Key:     key,
				Summary: secondarySummary(key),
			})
		}
	}
	return comparison
}

func appendFieldMismatch(fields *[]FieldMismatch, name string, primary, secondary any) {
	if fmt.Sprint(primary) == fmt.Sprint(secondary) {
		return
	}
	*fields = append(*fields, FieldMismatch{
		Field:     name,
		Primary:   fmt.Sprint(primary),
		Secondary: fmt.Sprint(secondary),
	})
}

func sortedKeys[T any](records map[string]T) []string {
	keys := make([]string, 0, len(records))
	for key := range records {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

func keysSet(keys []string) map[string]struct{} {
	set := make(map[string]struct{}, len(keys))
	for _, key := range keys {
		set[key] = struct{}{}
	}
	return set
}

func normalizedLimit(limit int) int {
	if limit <= 0 {
		return int(^uint(0) >> 1)
	}
	return limit
}

func prefix(value string, length int) string {
	if len(value) <= length {
		return value
	}
	return value[:length]
}

func formatFields(fields []FieldMismatch) string {
	parts := make([]string, 0, len(fields))
	for _, field := range fields {
		parts = append(parts, fmt.Sprintf("%s: %s vs %s", field.Field, field.Primary, field.Secondary))
	}
	return "[" + strings.Join(parts, " ") + "]"
}
