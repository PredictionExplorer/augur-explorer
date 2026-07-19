package testutil

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
)

// Row is one table row with volatile surrogate keys removed and foreign keys
// resolved to natural identifiers (addresses to hex strings, tx ids to tx
// hashes, evt_log ids to "block/logIndex" strings). This makes snapshots
// stable across sequence resets, replays and reorg re-processing.
type Row map[string]any

// Snapshot is the full visible state of a database: every non-empty user
// table with its transformed rows in a canonical order.
type Snapshot map[string][]Row

// Columns dropped from every table (or from one table via a "table.column"
// key): `id` is the volatile surrogate key, `address_id` its equivalent on
// the address table, `log_rlp` is the raw fixture input (pinned implicitly by
// the decoded rows it produces), the block hash columns depend on which fork
// of the fake test chain mined the block, and address.tx_id is first-seen
// bookkeeping that dangles after a reorg deletes the recorded transaction —
// dropping these lets reorg-replay tests assert that re-processing the same
// events on a different fork reproduces identical domain state.
var droppedColumns = map[string]bool{
	"id":            true,
	"address_id":    true,
	"log_rlp":       true,
	"block_hash":    true,
	"parent_hash":   true,
	"address.tx_id": true,
}

// TakeSnapshot reads every non-empty user table (except goose bookkeeping)
// and returns the canonical snapshot. All reads happen on one connection with
// the session time zone pinned to UTC so timestamptz rendering is stable.
func TakeSnapshot(ctx context.Context, db *sql.DB) (Snapshot, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("acquiring snapshot connection: %w", err)
	}
	defer func() { _ = conn.Close() }()

	if _, err := conn.ExecContext(ctx, "SET TIME ZONE 'UTC'"); err != nil {
		return nil, fmt.Errorf("pinning snapshot time zone: %w", err)
	}

	tables, err := listTables(ctx, conn)
	if err != nil {
		return nil, err
	}
	res, err := loadResolvers(ctx, conn)
	if err != nil {
		return nil, err
	}

	snap := make(Snapshot)
	for _, table := range tables {
		rows, err := dumpTable(ctx, conn, table, res)
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			snap[table] = rows
		}
	}
	return snap, nil
}

// listTables returns all base tables in schema public except goose bookkeeping.
func listTables(ctx context.Context, conn *sql.Conn) ([]string, error) {
	rows, err := conn.QueryContext(ctx, `
		SELECT table_name FROM information_schema.tables
		WHERE table_schema = 'public' AND table_type = 'BASE TABLE'
		  AND table_name <> 'goose_db_version'
		ORDER BY table_name`)
	if err != nil {
		return nil, fmt.Errorf("listing tables: %w", err)
	}
	defer func() { _ = rows.Close() }()
	var tables []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("scanning table name: %w", err)
		}
		tables = append(tables, name)
	}
	return tables, rows.Err()
}

// resolvers hold the surrogate-key-to-natural-key maps of one snapshot.
type resolvers struct {
	addr   map[int64]string // address.address_id -> hex address
	txHash map[int64]string // transaction.id      -> tx hash
	evtKey map[int64]string // evt_log.id          -> "evt:<block>/<logIndex>"
	bidKey map[int64]string // cg_bid.id           -> evtKey of the bid's event
}

func loadResolvers(ctx context.Context, conn *sql.Conn) (*resolvers, error) {
	res := &resolvers{
		addr:   make(map[int64]string),
		txHash: make(map[int64]string),
		evtKey: make(map[int64]string),
		bidKey: make(map[int64]string),
	}
	if err := loadMap(ctx, conn, "SELECT address_id, addr FROM address", res.addr); err != nil {
		return nil, err
	}
	if err := loadMap(ctx, conn, "SELECT id, tx_hash FROM transaction", res.txHash); err != nil {
		return nil, err
	}

	rows, err := conn.QueryContext(ctx, "SELECT id, block_num, log_index FROM evt_log")
	if err != nil {
		return nil, fmt.Errorf("loading evt_log resolver: %w", err)
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var id, blockNum int64
		var logIndex int
		if err := rows.Scan(&id, &blockNum, &logIndex); err != nil {
			return nil, fmt.Errorf("scanning evt_log resolver: %w", err)
		}
		res.evtKey[id] = fmt.Sprintf("evt:%d/%d", blockNum, logIndex)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	bidRows, err := conn.QueryContext(ctx, "SELECT id, evtlog_id FROM cg_bid")
	if err != nil {
		return nil, fmt.Errorf("loading cg_bid resolver: %w", err)
	}
	defer func() { _ = bidRows.Close() }()
	for bidRows.Next() {
		var id int64
		var evtlogID sql.NullInt64
		if err := bidRows.Scan(&id, &evtlogID); err != nil {
			return nil, fmt.Errorf("scanning cg_bid resolver: %w", err)
		}
		if evtlogID.Valid {
			res.bidKey[id] = res.evtKey[evtlogID.Int64]
		}
	}
	return res, bidRows.Err()
}

func loadMap(ctx context.Context, conn *sql.Conn, query string, dst map[int64]string) error {
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return fmt.Errorf("loading resolver (%s): %w", query, err)
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var id int64
		var value string
		if err := rows.Scan(&id, &value); err != nil {
			return fmt.Errorf("scanning resolver (%s): %w", query, err)
		}
		dst[id] = value
	}
	return rows.Err()
}

// dumpTable reads all rows of one table as JSON documents and canonicalizes them.
func dumpTable(ctx context.Context, conn *sql.Conn, table string, res *resolvers) ([]Row, error) {
	// Table names come from information_schema and are quoted; identifiers
	// cannot be bound as query parameters.
	query := fmt.Sprintf("SELECT row_to_json(t)::text FROM %s t", quoteIdent(table)) //nolint:gosec // quoted identifier from catalog
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("dumping %s: %w", table, err)
	}
	defer func() { _ = rows.Close() }()

	var out []Row
	for rows.Next() {
		var doc string
		if err := rows.Scan(&doc); err != nil {
			return nil, fmt.Errorf("scanning %s row: %w", table, err)
		}
		row, err := decodeRow(doc)
		if err != nil {
			return nil, fmt.Errorf("decoding %s row: %w", table, err)
		}
		out = append(out, transformRow(table, row, res))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	sortRows(out)
	return out, nil
}

func quoteIdent(name string) string {
	return `"` + strings.ReplaceAll(name, `"`, `""`) + `"`
}

func decodeRow(doc string) (Row, error) {
	dec := json.NewDecoder(strings.NewReader(doc))
	dec.UseNumber() // preserve 256-bit wei amounts exactly
	var row Row
	if err := dec.Decode(&row); err != nil {
		return nil, err
	}
	return row, nil
}

// transformRow drops volatile columns and resolves foreign keys.
func transformRow(table string, row Row, res *resolvers) Row {
	out := make(Row, len(row))
	for col, val := range row {
		if droppedColumns[col] || droppedColumns[table+"."+col] {
			continue
		}
		out[col] = resolveValue(col, val, res)
	}
	return out
}

// resolveValue maps surrogate-key column values to natural identifiers.
// Unresolvable values (0, NULL, unknown ids) pass through unchanged.
func resolveValue(col string, val any, res *resolvers) any {
	num, ok := val.(json.Number)
	if !ok {
		return val
	}
	id, err := num.Int64()
	if err != nil {
		return val
	}
	lookup := func(m map[int64]string) any {
		if s, ok := m[id]; ok {
			return s
		}
		return val
	}
	switch {
	case col == "aid" || strings.HasSuffix(col, "_aid"):
		return lookup(res.addr)
	case col == "tx_id" || col == "last_tx_id":
		return lookup(res.txHash)
	case col == "evtlog_id" || col == "withdrawal_id" || col == "last_evtlog_id" || col == "last_evt_id":
		return lookup(res.evtKey)
	case col == "bid_id":
		return lookup(res.bidKey)
	default:
		return val
	}
}

// sortRows orders rows by their canonical JSON encoding, giving every table a
// stable order without knowing its natural key.
func sortRows(rows []Row) {
	slices.SortFunc(rows, func(a, b Row) int {
		return strings.Compare(canonical(a), canonical(b))
	})
}

func canonical(row Row) string {
	b, err := json.Marshal(row) // map keys marshal sorted
	if err != nil {
		panic(fmt.Sprintf("testutil: marshaling row: %v", err))
	}
	return string(b)
}

// TableDiff lists the rows added to and removed from one table between two
// snapshots. An updated row appears as one removal plus one addition.
type TableDiff struct {
	Added   []Row `json:"added,omitempty"`
	Removed []Row `json:"removed,omitempty"`
}

// DiffSnapshots compares two snapshots (multiset semantics per table) and
// renders the differences as stable, indented JSON suitable for golden files.
func DiffSnapshots(before, after Snapshot) ([]byte, error) {
	tables := make(map[string]bool)
	for t := range before {
		tables[t] = true
	}
	for t := range after {
		tables[t] = true
	}

	diff := make(map[string]TableDiff)
	for table := range tables {
		counts := make(map[string]int)
		decoded := make(map[string]Row)
		for _, row := range before[table] {
			key := canonical(row)
			counts[key]--
			decoded[key] = row
		}
		for _, row := range after[table] {
			key := canonical(row)
			counts[key]++
			decoded[key] = row
		}
		var td TableDiff
		keys := make([]string, 0, len(counts))
		for key := range counts {
			keys = append(keys, key)
		}
		slices.Sort(keys)
		for _, key := range keys {
			for range counts[key] {
				td.Added = append(td.Added, decoded[key])
			}
			for range -counts[key] {
				td.Removed = append(td.Removed, decoded[key])
			}
		}
		if len(td.Added) > 0 || len(td.Removed) > 0 {
			diff[table] = td
		}
	}

	rendered, err := json.MarshalIndent(diff, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("rendering snapshot diff: %w", err)
	}
	return append(rendered, '\n'), nil
}
