// Package dbs - Blockchain event log utilities for ETL
package dbs

import (
	"database/sql"
	"fmt"
	"os"

	p "github.com/PredictionExplorer/augur-explorer/rwcg/primitives"
)

// Get_event_log retrieves an event log by its ID
func (ss *SQLStorage) Get_event_log(evtlog_id int64) p.EthereumEventLog {

	var evtlog p.EthereumEventLog
	evtlog.EvtId = evtlog_id
	var query string
	query = "SELECT " +
				"e.block_num," +
		"EXTRACT(EPOCH FROM b.ts)::BIGINT AS ts, " +
				"e.tx_id," +
				"tx.tx_hash," +
				"e.contract_aid," +
				"ca.addr, " +
				"e.topic0_sig," +
				"e.log_rlp " +
			"FROM evt_log e " +
				"JOIN block b ON e.block_num=b.block_num " +
				"JOIN transaction tx ON e.tx_id=tx.id " +
				"JOIN address ca ON e.contract_aid=ca.address_id " +
			"WHERE e.id=$1"
	res := ss.db.QueryRow(query, evtlog_id)
	err := res.Scan(
		&evtlog.BlockNum,
		&evtlog.TimeStamp,
		&evtlog.TxId,
		&evtlog.TxHash,
		&evtlog.ContractAid,
		&evtlog.ContractAddress,
		&evtlog.Topic0_Sig,
		&evtlog.RlpLog,
	)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v", err, query))
		os.Exit(1)
	}

	return evtlog
}

// Get_evtlogs_by_signature_in_range retrieves event log IDs by signature in a range
func (ss *SQLStorage) Get_evtlogs_by_signature_in_range(sig string, contract_aids string, from_evt_id, to_evt_id int64) []int64 {

	var query string
	query = "SELECT id FROM evt_log " +
		"WHERE topic0_sig=$1 " +
		"AND contract_aid IN (" + contract_aids + ") " +
		"AND id > $2 AND id <= $3 " +
		"ORDER BY id"

	rows, err := ss.db.Query(query, sig, from_evt_id, to_evt_id)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	output := make([]int64, 0, 256)
	defer rows.Close()
	for rows.Next() {
		var evt_id int64
		err = rows.Scan(&evt_id)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		output = append(output, evt_id)
	}
	return output
}

// Get_evtlogs_by_signature_only_in_range retrieves event log IDs by signature only (no contract filter)
func (ss *SQLStorage) Get_evtlogs_by_signature_only_in_range(sig string, from_evt_id, to_evt_id int64) []int64 {

	var query string
	query = "SELECT id FROM evt_log " +
		"WHERE topic0_sig=$1 " +
		"AND id > $2 AND id <= $3 " +
		"ORDER BY id"

	rows, err := ss.db.Query(query, sig, from_evt_id, to_evt_id)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	output := make([]int64, 0, 256)
	defer rows.Close()
	for rows.Next() {
		var evt_id int64
		err = rows.Scan(&evt_id)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		output = append(output, evt_id)
	}
	return output
}

// Get_last_evtlog_id returns the last event log ID in the database
func (ss *SQLStorage) Get_last_evtlog_id() (int64, error) {

	var query string
	query = "SELECT id FROM evt_log ORDER BY id DESC LIMIT 1"
	res := ss.db.QueryRow(query)
	var null_id sql.NullInt64
	err := res.Scan(&null_id)
		if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v", err, query))
		return 0, err
	}
	return null_id.Int64, nil
}

// Get_events_by_sig_and_tx_id retrieves event logs by signature and transaction ID
func (ss *SQLStorage) Get_events_by_sig_and_tx_id(tx_id int64, sig string) ([]p.EthereumEventLog, error) {

	var query string
	query = "SELECT " +
		"e.id," +
		"e.block_num," +
		"EXTRACT(EPOCH FROM b.ts)::BIGINT AS ts, " +
		"e.tx_id," +
		"tx.tx_hash," +
		"e.contract_aid," +
		"ca.addr, " +
		"e.topic0_sig," +
		"e.log_rlp " +
		"FROM evt_log e " +
		"JOIN block b ON e.block_num=b.block_num " +
		"JOIN transaction tx ON e.tx_id=tx.id " +
		"JOIN address ca ON e.contract_aid=ca.address_id " +
		"WHERE e.tx_id=$1 AND e.topic0_sig=$2 " +
		"ORDER BY e.id"

	rows, err := ss.db.Query(query, tx_id, sig)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		return nil, err
	}
	records := make([]p.EthereumEventLog, 0, 8)
	defer rows.Close()
	for rows.Next() {
		var evtlog p.EthereumEventLog
		err = rows.Scan(
			&evtlog.EvtId,
			&evtlog.BlockNum,
			&evtlog.TimeStamp,
			&evtlog.TxId,
			&evtlog.TxHash,
			&evtlog.ContractAid,
			&evtlog.ContractAddress,
			&evtlog.Topic0_Sig,
			&evtlog.RlpLog,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			return nil, err
		}
		records = append(records, evtlog)
	}
	return records, nil
}

// Get_specific_event_logs_by_tx_backwards_from_id retrieves event log RLP data
func (ss *SQLStorage) Get_specific_event_logs_by_tx_backwards_from_id(tx_id, contract_aid, starting_id int64, signature string) [][]byte {

	records := make([][]byte, 0, 4)
	var query string
	query = "SELECT log_rlp FROM evt_log " +
		"WHERE tx_id=$1 AND contract_aid=$2 AND id<$3 AND topic0_sig=$4 " +
		"ORDER BY id DESC"

	rows, err := ss.db.Query(query, tx_id, contract_aid, starting_id, signature)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rlp []byte
		err = rows.Scan(&rlp)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		records = append(records, rlp)
	}
	return records
}

// Get_last_block_timestamp returns the timestamp of the last block
// Returns 0 if no blocks exist (graceful handling for empty/new database)
func (ss *SQLStorage) Get_last_block_timestamp() int64 {
	var query string
	query = "SELECT FLOOR(EXTRACT(EPOCH FROM block.ts))::BIGINT AS ts " +
		"FROM block,last_block WHERE last_block.block_num=block.block_num"
	row := ss.db.QueryRow(query)
	var ts int64
	var err error
	err = row.Scan(&ts)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0 // No blocks yet, return 0
		}
		ss.Log_msg(fmt.Sprintf("Error in Get_last_block_timestamp(): %v, q=%v", err, query))
		return 0
	}
	return ts
}

// Get_first_block_timestamp returns the timestamp of the first block
// Returns 0 if no blocks exist (graceful handling for empty/new database)
func (ss *SQLStorage) Get_first_block_timestamp() int64 {
	var query string
	query = "SELECT FLOOR(EXTRACT(EPOCH FROM block.ts))::BIGINT AS ts " +
		"FROM block ORDER BY block_num LIMIT 1"
	row := ss.db.QueryRow(query)
	var ts int64
	var err error
	err = row.Scan(&ts)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0 // No blocks yet, return 0
		}
		ss.Log_msg(fmt.Sprintf("Error in Get_first_block_timestamp(): %v, q=%v", err, query))
		return 0
	}
	return ts
}

// =============================================================================
// BLOCK OPERATIONS (for new FilterLogs-based ETL)
// =============================================================================

// Get_block_hash returns the block hash for a given block number
// Returns error if block doesn't exist
func (ss *SQLStorage) Get_block_hash(blockNum int64) (string, error) {
	var query string
	query = "SELECT block_hash FROM block WHERE block_num = $1"
	row := ss.db.QueryRow(query, blockNum)
	var blockHash string
	err := row.Scan(&blockHash)
	if err != nil {
		return "", err
		}
	return blockHash, nil
	}

// Get_last_block_num returns the last block number from last_block table
func (ss *SQLStorage) Get_last_block_num() (int64, error) {
	var query string
	query = "SELECT block_num FROM last_block LIMIT 1"
	row := ss.db.QueryRow(query)
	var null_block sql.NullInt64
	err := row.Scan(&null_block)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return null_block.Int64, nil
}

// Set_last_block_num updates the last block number in last_block table
func (ss *SQLStorage) Set_last_block_num(blockNum int64) error {
	var query string 
	query = "UPDATE last_block SET block_num = $1"
	_, err := ss.db.Exec(query, blockNum)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Set_last_block_num failed: %v", err))
		return err
	}
	return nil
}

// Delete_block deletes a block and all its associated data (cascades via foreign keys)
func (ss *SQLStorage) Delete_block(blockNum int64) error {
	var query string 
	query = "DELETE FROM block WHERE block_num = $1"
	_, err := ss.db.Exec(query, blockNum)
	if err != nil {
		ss.Log_msg(fmt.Sprintf("Delete_block failed for block %d: %v", blockNum, err))
		return err
	}
	return nil
}

// =============================================================================
// TRANSACTION OPERATIONS
// =============================================================================

// Get_transaction_id_by_hash returns the transaction ID for a given tx hash
func (ss *SQLStorage) Get_transaction_id_by_hash(txHash string) (int64, error) {
	var query string
	query = "SELECT id FROM transaction WHERE tx_hash = $1"
	row := ss.db.QueryRow(query, txHash)
	var txId int64
	err := row.Scan(&txId)
	if err != nil {
		return 0, err
	}
	return txId, nil
}
