// Package verify implements the cmd/freezer-verify engine: it compares the
// events extracted from a node freezer (the JSONL output of freezer-scan)
// with the events stored in a PostgreSQL evt_log table, reporting missing,
// extra and count-mismatched (block, topic0, contract) combinations.
package verify

import (
	"bufio"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"slices"
	"strings"

	"github.com/jackc/pgx/v5"
)

// LogRecord matches the JSONL output format of freezer-scan.
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

// Key identifies events by block, truncated topic0 signature and contract,
// matching the evt_log storage format (topic0_sig holds the first 4 bytes).
type Key struct {
	BlockNum int64
	Topic0   string // first 8 hex chars of topic0, lowercase
	Contract string // 0x-prefixed address, lowercase
}

// validTableName restricts table names to plain SQL identifiers, since the
// name is interpolated into the query text.
var validTableName = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// ValidTableName reports whether name is a plain SQL identifier.
func ValidTableName(name string) bool {
	return validTableName.MatchString(name)
}

// LoadJSONL reads freezer events from r. It returns the count of events per
// key and the set of contract addresses seen (lowercase).
func LoadJSONL(r io.Reader) (map[Key]int, map[string]bool, error) {
	events := make(map[Key]int)
	contracts := make(map[string]bool)
	scanner := bufio.NewScanner(r)

	// Increase buffer size for long lines (dataHex can be large).
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if line == "" {
			continue
		}

		var record LogRecord
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			return nil, nil, fmt.Errorf("line %d: invalid JSON: %w", lineNum, err)
		}

		// Extract the first 4 bytes of topic0 (matching the DB format).
		topic0 := strings.ToLower(record.Topic0)
		if strings.HasPrefix(topic0, "0x") && len(topic0) >= 10 {
			topic0 = topic0[2:10]
		}

		contract := strings.ToLower(record.Contract)
		contracts[contract] = true

		events[Key{
			BlockNum: int64(record.BlockNumber), // #nosec G115 -- scanner output from real freezer files fits int64
			Topic0:   topic0,
			Contract: contract,
		}]++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("scanner error: %w", err)
	}

	return events, contracts, nil
}

// BlockRange returns the inclusive block range covered by the events. ok is
// false when there are no events.
func BlockRange(events map[Key]int) (minBlock, maxBlock int64, ok bool) {
	first := true
	for key := range events {
		if first {
			minBlock, maxBlock = key.BlockNum, key.BlockNum
			first = false
			continue
		}
		if key.BlockNum < minBlock {
			minBlock = key.BlockNum
		}
		if key.BlockNum > maxBlock {
			maxBlock = key.BlockNum
		}
	}
	return minBlock, maxBlock, !first
}

// Querier is the narrow pgx query surface the database readers use.
// *pgx.Conn and *pgxpool.Pool satisfy it.
type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

// ContractAids resolves the address_id values of the given contract
// addresses (matched case-insensitively).
func ContractAids(ctx context.Context, q Querier, contracts []string) ([]int64, error) {
	rows, err := q.Query(ctx, `
		SELECT address_id FROM address
		WHERE LOWER(addr) = ANY($1)
	`, contracts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aids []int64
	for rows.Next() {
		var aid int64
		if err := rows.Scan(&aid); err != nil {
			return nil, err
		}
		aids = append(aids, aid)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return aids, nil
}

// DBEvents queries the event counts per (block, topic0, contract) from the
// given evt_log table, restricted to the block range and contract ids.
func DBEvents(ctx context.Context, q Querier, table string, minBlock, maxBlock int64, aids []int64) (map[Key]int, error) {
	if !ValidTableName(table) {
		return nil, fmt.Errorf("invalid table name %q: must be a plain SQL identifier", table)
	}
	// The table name is validated above and quoted here; values are bound.
	query := `
		SELECT e.block_num, e.topic0_sig, a.addr, COUNT(*) as cnt
		FROM ` + pgx.Identifier{table}.Sanitize() + ` e
		JOIN address a ON e.contract_aid = a.address_id
		WHERE e.block_num BETWEEN $1 AND $2
		  AND e.contract_aid = ANY($3)
		GROUP BY e.block_num, e.topic0_sig, a.addr
		ORDER BY e.block_num
	`
	rows, err := q.Query(ctx, query, minBlock, maxBlock, aids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make(map[Key]int)
	for rows.Next() {
		var blockNum int64
		var topic0, contract string
		var cnt int
		if err := rows.Scan(&blockNum, &topic0, &contract, &cnt); err != nil {
			return nil, err
		}
		events[Key{
			BlockNum: blockNum,
			Topic0:   strings.ToLower(strings.TrimSpace(topic0)),
			Contract: strings.ToLower(contract),
		}] = cnt
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// Report is the outcome of comparing freezer events with database events.
type Report struct {
	FreezerKeys int // distinct keys in the freezer extract
	DBKeys      int // distinct keys in the database

	Match         int // keys present on both sides with equal counts
	CountMismatch int // keys present on both sides with different counts
	Missing       int // keys in the database but not in the freezer
	Extra         int // keys in the freezer but not in the database

	MissingDetails  []string // up to maxMissing samples
	ExtraDetails    []string // up to maxExtra samples
	MismatchDetails []string // every count mismatch
}

// Compare builds the report. maxMissing/maxExtra bound the detail samples.
func Compare(freezer, db map[Key]int, maxMissing, maxExtra int) Report {
	r := Report{FreezerKeys: len(freezer), DBKeys: len(db)}

	for _, key := range sortedKeys(db) {
		dbCount := db[key]
		fCount, ok := freezer[key]
		switch {
		case !ok:
			r.Missing++
			if len(r.MissingDetails) < maxMissing {
				r.MissingDetails = append(r.MissingDetails,
					fmt.Sprintf("block=%d topic0=%s contract=%s (db count=%d)",
						key.BlockNum, key.Topic0, key.Contract, dbCount))
			}
		case fCount == dbCount:
			r.Match++
		default:
			r.CountMismatch++
			r.MismatchDetails = append(r.MismatchDetails,
				fmt.Sprintf("block=%d topic0=%s contract=%s: db=%d freezer=%d",
					key.BlockNum, key.Topic0, key.Contract, dbCount, fCount))
		}
	}

	for _, key := range sortedKeys(freezer) {
		if _, ok := db[key]; !ok {
			r.Extra++
			if len(r.ExtraDetails) < maxExtra {
				r.ExtraDetails = append(r.ExtraDetails,
					fmt.Sprintf("block=%d topic0=%s contract=%s (freezer count=%d)",
						key.BlockNum, key.Topic0, key.Contract, freezer[key]))
			}
		}
	}

	return r
}

// sortedKeys returns the map keys in deterministic order, so detail samples
// and reports are stable across runs.
func sortedKeys(m map[Key]int) []Key {
	keys := make([]Key, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	slices.SortFunc(keys, func(a, b Key) int {
		return cmp.Or(
			cmp.Compare(a.BlockNum, b.BlockNum),
			cmp.Compare(a.Topic0, b.Topic0),
			cmp.Compare(a.Contract, b.Contract),
		)
	})
	return keys
}

// MatchRate returns the fraction of database keys the freezer matched
// exactly. It is 1 when the database had no keys.
func (r Report) MatchRate() float64 {
	if r.DBKeys == 0 {
		return 1
	}
	return float64(r.Match) / float64(r.DBKeys)
}

// Passed reports whether the verification met the 99% match threshold.
func (r Report) Passed() bool {
	return r.MatchRate() >= 0.99
}

// Write renders the report. Mismatch details render only when verbose.
func (r Report) Write(w io.Writer, verbose bool) {
	fmt.Fprintln(w)
	fmt.Fprintln(w, "=== Verification Results ===")
	fmt.Fprintf(w, "Freezer events:  %d distinct (block, topic0, contract)\n", r.FreezerKeys)
	fmt.Fprintf(w, "Database events: %d distinct (block, topic0, contract)\n", r.DBKeys)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Matching:          %d\n", r.Match)
	fmt.Fprintf(w, "Missing (in DB, not in freezer): %d\n", r.Missing)
	fmt.Fprintf(w, "Extra (in freezer, not in DB):   %d\n", r.Extra)

	if r.DBKeys > 0 {
		fmt.Fprintf(w, "\nMatch rate: %.2f%%\n", r.MatchRate()*100)
	}

	if verbose && len(r.MismatchDetails) > 0 {
		fmt.Fprintln(w, "\nCount mismatches:")
		for _, d := range r.MismatchDetails {
			fmt.Fprintf(w, "  %s\n", d)
		}
	}

	if len(r.MissingDetails) > 0 {
		fmt.Fprintln(w, "\nMissing from freezer (sample):")
		for _, d := range r.MissingDetails {
			fmt.Fprintf(w, "  %s\n", d)
		}
		if r.Missing > len(r.MissingDetails) {
			fmt.Fprintf(w, "  ... and %d more\n", r.Missing-len(r.MissingDetails))
		}
	}

	if len(r.ExtraDetails) > 0 {
		fmt.Fprintln(w, "\nExtra in freezer (sample):")
		for _, d := range r.ExtraDetails {
			fmt.Fprintf(w, "  %s\n", d)
		}
		if r.Extra > len(r.ExtraDetails) {
			fmt.Fprintf(w, "  ... and %d more\n", r.Extra-len(r.ExtraDetails))
		}
	}
}
