package dbverify

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

// Loader separates database reads from pure comparison and reporting logic.
type Loader interface {
	LoadEventRecords(ctx context.Context, contractAddressIDs []int64) (map[string]EventRecord, error)
	TransactionHashesFromEvents(ctx context.Context, contractAddressIDs []int64) ([]string, error)
	LoadTransactions(ctx context.Context, txHashes []string) (map[string]TransactionRecord, error)
	BlockNumbersFromEvents(ctx context.Context, contractAddressIDs []int64) ([]int64, error)
	LoadBlocks(ctx context.Context, blockNumbers []int64) (map[string]BlockRecord, error)
	CountEventLogs(ctx context.Context, contractAddressIDs []int64) (int64, error)
	LoadDetailedEventLogs(ctx context.Context, contractAddressIDs []int64, limit int) ([]EventLogRecord, error)
}

// SQLLoader reads comparison data from one database/sql handle. It does not
// own DB. A nil filter slice means load all rows.
type SQLLoader struct {
	DB *sql.DB
}

var errDatabaseRequired = errors.New("dbverify: database is required")

// LoadRandomWalkContractAddressIDs returns the primary database's RandomWalk
// contract ids in stable order.
func LoadRandomWalkContractAddressIDs(ctx context.Context, db *sql.DB) ([]int64, error) {
	if db == nil {
		return nil, errDatabaseRequired
	}
	rows, err := db.QueryContext(ctx, `
		SELECT DISTINCT a.address_id
		FROM address a
		JOIN rw_contracts rc ON a.addr = rc.randomwalk_addr OR a.addr = rc.marketplace_addr
		ORDER BY a.address_id
	`)
	if err != nil {
		return nil, fmt.Errorf("contract aids: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("scan contract aid: %w", err)
		}
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("contract aids: %w", err)
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("no contract addresses found in rw_contracts")
	}
	return ids, nil
}

func (l *SQLLoader) LoadEventRecords(
	ctx context.Context,
	contractAddressIDs []int64,
) (map[string]EventRecord, error) {
	db, err := l.database()
	if err != nil {
		return nil, err
	}
	query := `
		SELECT e.block_num, t.tx_hash, e.log_rlp
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id`
	var rows *sql.Rows
	if contractAddressIDs != nil {
		rows, err = db.QueryContext(ctx,
			query+" WHERE e.contract_aid = ANY($1) ORDER BY e.log_rlp, e.block_num, t.tx_hash, e.id",
			pq.Array(contractAddressIDs),
		)
	} else {
		rows, err = db.QueryContext(ctx, query+" ORDER BY e.log_rlp, e.block_num, t.tx_hash, e.id")
	}
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	records := make(map[string]EventRecord)
	occurrences := make(map[string]int)
	for rows.Next() {
		var (
			record EventRecord
			logRLP []byte
		)
		if err := rows.Scan(&record.BlockNum, &record.TxHash, &logRLP); err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}
		record.LogRLPHex = hex.EncodeToString(logRLP)
		occurrences[record.LogRLPHex]++
		records[occurrenceKey(record.LogRLPHex, occurrences[record.LogRLPHex])] = record
	}
	return records, rows.Err()
}

func (l *SQLLoader) TransactionHashesFromEvents(
	ctx context.Context,
	contractAddressIDs []int64,
) ([]string, error) {
	db, err := l.database()
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, `
		SELECT DISTINCT t.tx_hash
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
		ORDER BY t.tx_hash
	`, pq.Array(contractAddressIDs))
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var hashes []string
	for rows.Next() {
		var hash string
		if err := rows.Scan(&hash); err != nil {
			return nil, fmt.Errorf("scan tx_hash: %w", err)
		}
		hashes = append(hashes, hash)
	}
	return hashes, rows.Err()
}

func (l *SQLLoader) LoadTransactions(
	ctx context.Context,
	txHashes []string,
) (map[string]TransactionRecord, error) {
	db, err := l.database()
	if err != nil {
		return nil, err
	}
	records := make(map[string]TransactionRecord)
	query := `
		SELECT block_num, tx_hash, gas_used, num_logs
		FROM transaction`
	var rows *sql.Rows
	if txHashes != nil {
		if len(txHashes) == 0 {
			return records, nil
		}
		rows, err = db.QueryContext(ctx,
			query+" WHERE tx_hash = ANY($1) ORDER BY tx_hash",
			pq.Array(txHashes),
		)
	} else {
		rows, err = db.QueryContext(ctx, query+" ORDER BY tx_hash")
	}
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var record TransactionRecord
		if err := rows.Scan(&record.BlockNum, &record.TxHash, &record.GasUsed, &record.NumLogs); err != nil {
			return nil, fmt.Errorf("scan transaction: %w", err)
		}
		records[record.TxHash] = record
	}
	return records, rows.Err()
}

func (l *SQLLoader) BlockNumbersFromEvents(
	ctx context.Context,
	contractAddressIDs []int64,
) ([]int64, error) {
	db, err := l.database()
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, `
		SELECT DISTINCT t.block_num
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
		ORDER BY t.block_num
	`, pq.Array(contractAddressIDs))
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var numbers []int64
	for rows.Next() {
		var number int64
		if err := rows.Scan(&number); err != nil {
			return nil, fmt.Errorf("scan block_num: %w", err)
		}
		numbers = append(numbers, number)
	}
	return numbers, rows.Err()
}

func (l *SQLLoader) LoadBlocks(
	ctx context.Context,
	blockNumbers []int64,
) (map[string]BlockRecord, error) {
	db, err := l.database()
	if err != nil {
		return nil, err
	}
	records := make(map[string]BlockRecord)
	query := `
		SELECT block_num, block_hash, parent_hash, num_tx
		FROM block`
	var rows *sql.Rows
	if blockNumbers != nil {
		if len(blockNumbers) == 0 {
			return records, nil
		}
		rows, err = db.QueryContext(ctx,
			query+" WHERE block_num = ANY($1) ORDER BY block_num",
			pq.Array(blockNumbers),
		)
	} else {
		rows, err = db.QueryContext(ctx, query+" ORDER BY block_num")
	}
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var record BlockRecord
		if err := rows.Scan(&record.BlockNum, &record.BlockHash, &record.ParentHash, &record.NumTx); err != nil {
			return nil, fmt.Errorf("scan block: %w", err)
		}
		records[record.BlockHash] = record
	}
	return records, rows.Err()
}

func (l *SQLLoader) CountEventLogs(
	ctx context.Context,
	contractAddressIDs []int64,
) (int64, error) {
	db, err := l.database()
	if err != nil {
		return 0, err
	}
	query := "SELECT COUNT(*) FROM evt_log"
	var count int64
	if contractAddressIDs != nil {
		err = db.QueryRowContext(ctx,
			query+" WHERE contract_aid = ANY($1)",
			pq.Array(contractAddressIDs),
		).Scan(&count)
	} else {
		err = db.QueryRowContext(ctx, query).Scan(&count)
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (l *SQLLoader) LoadDetailedEventLogs(
	ctx context.Context,
	contractAddressIDs []int64,
	limit int,
) ([]EventLogRecord, error) {
	db, err := l.database()
	if err != nil {
		return nil, err
	}
	query := `
		SELECT e.block_num, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
		FROM evt_log e
		JOIN transaction t ON e.tx_id = t.id
		JOIN address a ON e.contract_aid = a.address_id`
	args := make([]any, 0, 2)
	if contractAddressIDs != nil {
		args = append(args, pq.Array(contractAddressIDs))
		query += " WHERE e.contract_aid = ANY($1)"
	}
	query += " ORDER BY e.block_num, e.id"
	if limit > 0 {
		args = append(args, limit)
		// #nosec G202 -- only the generated placeholder ordinal is interpolated.
		query += fmt.Sprintf(" LIMIT $%d", len(args))
	}
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var events []EventLogRecord
	for rows.Next() {
		var event EventLogRecord
		if err := rows.Scan(
			&event.BlockNum,
			&event.TxHash,
			&event.ContractAddress,
			&event.Topic0Sig,
			&event.LogRLP,
		); err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}
		events = append(events, event)
	}
	return events, rows.Err()
}

func (l *SQLLoader) database() (*sql.DB, error) {
	if l == nil || l.DB == nil {
		return nil, errDatabaseRequired
	}
	return l.DB, nil
}
