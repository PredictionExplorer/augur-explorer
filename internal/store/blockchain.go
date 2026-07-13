// Blockchain event-log, block and transaction reads used by the ETL
// pipeline and its operator tools.

package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// evtLogColumns is the SELECT list shared by the event-log readers (without
// the leading id column, which only EventsBySigAndTx needs).
const evtLogColumns = `
		e.block_num,
		EXTRACT(EPOCH FROM b.ts)::BIGINT AS ts,
		e.tx_id,
		tx.tx_hash,
		e.contract_aid,
		ca.addr,
		e.topic0_sig,
		e.log_rlp
	FROM evt_log e
	JOIN block b ON e.block_num=b.block_num
	JOIN transaction tx ON e.tx_id=tx.id
	JOIN address ca ON e.contract_aid=ca.address_id`

// EventLog retrieves a stored event log by its id; a missing row yields a
// wrapped ErrNotFound.
func (s *Store) EventLog(ctx context.Context, evtlogID int64) (EthereumEventLog, error) {
	var evtlog EthereumEventLog
	evtlog.EvtId = evtlogID
	err := s.pool.QueryRow(ctx, "SELECT "+evtLogColumns+" WHERE e.id=$1", evtlogID).Scan(
		&evtlog.BlockNum,
		&evtlog.TimeStamp,
		&evtlog.TxId,
		&evtlog.TxHash,
		&evtlog.ContractAid,
		&evtlog.ContractAddress,
		&evtlog.Topic0Sig,
		&evtlog.RlpLog,
	)
	if err != nil {
		return evtlog, WrapError(fmt.Sprintf("event log lookup for id %v", evtlogID), err)
	}
	return evtlog, nil
}

// EventsBySigAndTx retrieves the event logs of one transaction carrying the
// given topic-0 signature, in insertion order.
func (s *Store) EventsBySigAndTx(ctx context.Context, txID int64, sig string) ([]EthereumEventLog, error) {
	query := "SELECT e.id," + evtLogColumns + " WHERE e.tx_id=$1 AND e.topic0_sig=$2 ORDER BY e.id"
	return QueryList(ctx, s.pool, fmt.Sprintf("event logs for tx %v sig %v", txID, sig), 8, query,
		func(rows pgx.Rows, evtlog *EthereumEventLog) error {
			return rows.Scan(
				&evtlog.EvtId,
				&evtlog.BlockNum,
				&evtlog.TimeStamp,
				&evtlog.TxId,
				&evtlog.TxHash,
				&evtlog.ContractAid,
				&evtlog.ContractAddress,
				&evtlog.Topic0Sig,
				&evtlog.RlpLog,
			)
		}, txID, sig)
}

// EventLogRLPsBefore returns the raw RLP payloads of a transaction's event
// logs with the given signature and contract, at ids below beforeID, newest
// first. The CosmicGame donation decoder uses it to pair a claim with its
// original donation event.
func (s *Store) EventLogRLPsBefore(ctx context.Context, txID, contractAid, beforeID int64, sig string) ([][]byte, error) {
	query := `SELECT log_rlp FROM evt_log
		WHERE tx_id=$1 AND contract_aid=$2 AND id<$3 AND topic0_sig=$4
		ORDER BY id DESC`
	return QueryList(ctx, s.pool, fmt.Sprintf("event log RLPs for tx %v sig %v", txID, sig), 4, query,
		func(rows pgx.Rows, rlp *[]byte) error {
			return rows.Scan(rlp)
		}, txID, contractAid, beforeID, sig)
}

// =============================================================================
// BLOCK OPERATIONS (for the FilterLogs-based ETL)
// =============================================================================

// BlockHash returns the stored hash of a block; a block that was never
// inserted yields a wrapped ErrNotFound.
func (s *Store) BlockHash(ctx context.Context, blockNum int64) (string, error) {
	var blockHash string
	err := s.pool.QueryRow(ctx, "SELECT block_hash FROM block WHERE block_num = $1", blockNum).Scan(&blockHash)
	if err != nil {
		return "", WrapError(fmt.Sprintf("block hash lookup for block %d", blockNum), err)
	}
	return blockHash, nil
}

// LastBlockNum returns the ETL block watermark from the last_block table
// (0 when the singleton row is missing or NULL).
func (s *Store) LastBlockNum(ctx context.Context) (int64, error) {
	var blockNum *int64
	err := s.pool.QueryRow(ctx, "SELECT block_num FROM last_block LIMIT 1").Scan(&blockNum)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, WrapError("last block num", err)
	}
	if blockNum == nil {
		return 0, nil
	}
	return *blockNum, nil
}

// SetLastBlockNum updates the ETL block watermark.
func (s *Store) SetLastBlockNum(ctx context.Context, blockNum int64) error {
	_, err := s.pool.Exec(ctx, "UPDATE last_block SET block_num = $1", blockNum)
	return WrapError("set last block num", err)
}

// DeleteBlock deletes a block and all its associated data (cascades via
// foreign keys; the plpgsql delete triggers reverse the aggregates).
func (s *Store) DeleteBlock(ctx context.Context, blockNum int64) error {
	_, err := s.pool.Exec(ctx, "DELETE FROM block WHERE block_num = $1", blockNum)
	return WrapError(fmt.Sprintf("delete block %d", blockNum), err)
}

// =============================================================================
// TRANSACTION OPERATIONS
// =============================================================================

// EvtLogExists reports whether evt_log already has (block_num, tx_id,
// log_index).
func (s *Store) EvtLogExists(ctx context.Context, blockNum, txID int64, logIndex int) (bool, error) {
	var evtID int64
	err := s.pool.QueryRow(ctx,
		"SELECT id FROM evt_log WHERE block_num=$1 AND tx_id=$2 AND log_index=$3 LIMIT 1",
		blockNum, txID, logIndex).Scan(&evtID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, WrapError("evt_log existence check", err)
	}
	return true, nil
}

// CountEvtLogsForContract returns the number of evt_log rows emitted by the
// contract address (case-insensitive).
func (s *Store) CountEvtLogsForContract(ctx context.Context, contractAddr string) (int64, error) {
	var count int64
	err := s.pool.QueryRow(ctx, `SELECT COUNT(*) FROM evt_log e
		JOIN address a ON e.contract_aid=a.address_id
		WHERE lower(a.addr)=lower($1)`, contractAddr).Scan(&count)
	if err != nil {
		return 0, WrapError(fmt.Sprintf("evt_log count for contract %v", contractAddr), err)
	}
	return count, nil
}

// TransactionIDByHash returns the transaction id for a tx hash; a hash that
// was never stored yields a wrapped ErrNotFound.
func (s *Store) TransactionIDByHash(ctx context.Context, txHash string) (int64, error) {
	var txID int64
	err := s.pool.QueryRow(ctx, "SELECT id FROM transaction WHERE tx_hash = $1", txHash).Scan(&txID)
	if err != nil {
		return 0, WrapError(fmt.Sprintf("transaction id lookup for %v", txHash), err)
	}
	return txID, nil
}
