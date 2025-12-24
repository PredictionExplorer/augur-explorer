// Package dbs - Archive table operations for reading historical data
// when RPC nodes have pruned transaction indices
package dbs

import (
	"database/sql"
	"fmt"
)

// ArchivedTransaction holds transaction data from archive
type ArchivedTransaction struct {
	BlockNum    int64
	TxHash      string
	TxIndex     int
	FromAid     int64
	ToAid       int64
	Value       string
	GasUsed     int64
	GasPrice    string
	InputSig    string
	NumLogs     int
	CtrctCreate bool
}

// ArchivedEventLog holds event log data from archive
type ArchivedEventLog struct {
	BlockNum     int64
	EvtId        int64
	TxHash       string
	ContractAddr string
	Topic0Sig    string
	LogRlp       []byte
}

// Get_archived_transaction reads a transaction from rw_arch_tx table
// Returns nil if not found
func (ss *SQLStorage) Get_archived_transaction(txHash string) (*ArchivedTransaction, error) {
	var tx ArchivedTransaction
	var query string
	query = `SELECT 
		block_num, tx_hash, tx_index,
		from_aid, to_aid, value,
		gas_used, gas_price, input_sig,
		num_logs, ctrct_create
	FROM rw_arch_tx WHERE tx_hash = $1`

	var inputSig sql.NullString
	err := ss.db.QueryRow(query, txHash).Scan(
		&tx.BlockNum,
		&tx.TxHash,
		&tx.TxIndex,
		&tx.FromAid,
		&tx.ToAid,
		&tx.Value,
		&tx.GasUsed,
		&tx.GasPrice,
		&inputSig,
		&tx.NumLogs,
		&tx.CtrctCreate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found in archive
		}
		return nil, fmt.Errorf("archive query failed: %v", err)
	}
	if inputSig.Valid {
		tx.InputSig = inputSig.String
	}
	return &tx, nil
}

// Get_archived_event_logs reads all event logs for a transaction from rw_arch_evtlog
func (ss *SQLStorage) Get_archived_event_logs(txHash string) ([]ArchivedEventLog, error) {
	var query string
	query = `SELECT 
		block_num, evt_id, tx_hash, contract_addr, topic0_sig, log_rlp
	FROM rw_arch_evtlog WHERE tx_hash = $1
	ORDER BY evt_id`

	rows, err := ss.db.Query(query, txHash)
	if err != nil {
		return nil, fmt.Errorf("archive event log query failed: %v", err)
	}
	defer rows.Close()

	var logs []ArchivedEventLog
	for rows.Next() {
		var log ArchivedEventLog
		err := rows.Scan(
			&log.BlockNum,
			&log.EvtId,
			&log.TxHash,
			&log.ContractAddr,
			&log.Topic0Sig,
			&log.LogRlp,
		)
		if err != nil {
			return nil, fmt.Errorf("archive event log scan failed: %v", err)
		}
		logs = append(logs, log)
	}
	return logs, nil
}

// Insert_transaction_from_archive inserts a transaction record using archived data
// Returns the transaction ID
func (ss *SQLStorage) Insert_transaction_from_archive(arch *ArchivedTransaction) (int64, error) {
	var query string
	query = `INSERT INTO transaction (
		block_num, tx_hash, tx_index, 
		from_aid, to_aid, value, 
		gas_used, gas_price,
		input_sig, num_logs, ctrct_create
	) VALUES (
		$1, $2, $3, 
		$4, $5, $6, 
		$7, $8,
		$9, $10, $11
	) ON CONFLICT (tx_hash) DO UPDATE SET id = transaction.id
	RETURNING id`

	var txId int64
	err := ss.db.QueryRow(query,
		arch.BlockNum,
		arch.TxHash,
		arch.TxIndex,
		arch.FromAid,
		arch.ToAid,
		arch.Value,
		arch.GasUsed,
		arch.GasPrice,
		arch.InputSig,
		arch.NumLogs,
		arch.CtrctCreate,
	).Scan(&txId)

	if err != nil {
		ss.Log_msg(fmt.Sprintf("Insert_transaction_from_archive failed: %v", err))
		return 0, err
	}

	return txId, nil
}

