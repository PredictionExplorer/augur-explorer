// Archive table operations for reading historical data when RPC nodes have
// pruned transaction indices.

package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// ArchivedTransaction holds transaction data from archive.
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

// ArchivedTransactionByHash reads a transaction from the arch_tx table.
// A transaction absent from the archive yields (nil, nil).
func (s *Store) ArchivedTransactionByHash(ctx context.Context, txHash string) (*ArchivedTransaction, error) {
	var tx ArchivedTransaction
	var inputSig *string
	err := s.pool.QueryRow(ctx, `SELECT
			block_num, tx_hash, tx_index,
			from_aid, to_aid, value,
			gas_used, gas_price, input_sig,
			num_logs, ctrct_create
		FROM arch_tx WHERE tx_hash = $1`, txHash).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // not found in archive
		}
		return nil, WrapError(fmt.Sprintf("archive lookup for %v", txHash), err)
	}
	if inputSig != nil {
		tx.InputSig = *inputSig
	}
	return &tx, nil
}

// InsertTransactionFromArchive inserts a transaction record using archived
// data (idempotent on tx_hash) and returns the transaction id.
func (s *Store) InsertTransactionFromArchive(ctx context.Context, arch *ArchivedTransaction) (int64, error) {
	var txID int64
	err := s.pool.QueryRow(ctx, `INSERT INTO transaction (
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
		RETURNING id`,
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
	).Scan(&txID)
	if err != nil {
		return 0, WrapError(fmt.Sprintf("insert transaction from archive %v", arch.TxHash), err)
	}
	return txID, nil
}
