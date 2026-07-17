package archive

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lib/pq"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer/logscan"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

const (
	// DefaultFilterBatchBlocks is the default eth_getLogs block range.
	DefaultFilterBatchBlocks = uint64(100_000)
	minRetryBatchBlocks      = uint64(1_000)
	defaultRetryDelay        = 3 * time.Second
)

// FillStats accumulates counters for one or more node-fill runs.
type FillStats struct {
	BlocksScanned uint64
	LogsFromNode  int64
	LogsSkipped   int64
	LogsInserted  int64
	TxInserted    int64
	TxSkipped     int64
	BlockInserted int64
	BlockSkipped  int64
	FilterRetries int64
	RPCErrors     int64
	DBErrors      int64
}

// Merge adds other into s.
func (s *FillStats) Merge(other FillStats) {
	s.BlocksScanned += other.BlocksScanned
	s.LogsFromNode += other.LogsFromNode
	s.LogsSkipped += other.LogsSkipped
	s.LogsInserted += other.LogsInserted
	s.TxInserted += other.TxInserted
	s.TxSkipped += other.TxSkipped
	s.BlockInserted += other.BlockInserted
	s.BlockSkipped += other.BlockSkipped
	s.FilterRetries += other.FilterRetries
	s.RPCErrors += other.RPCErrors
	s.DBErrors += other.DBErrors
}

// NodeFillOptions controls one project's node scan.
type NodeFillOptions struct {
	FromBlock uint64
	EndBlock  uint64
	BatchSize uint64
	DryRun    bool
}

// AddressStore is the address-cache operation node-fill needs from store.Store.
type AddressStore interface {
	LookupOrCreateAddress(ctx context.Context, address string, blockNum, txID int64) (int64, error)
}

// NodeClient is the narrow RPC surface used by node-fill.
type NodeClient interface {
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error)
	TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
}

// NodeFillRepository owns project metadata, resume metadata, and a prepared
// writer session. Production uses SQLNodeFillRepository; tests can fake it.
type NodeFillRepository interface {
	ProjectContracts(ctx context.Context, project string) (Contracts, error)
	ResolveStartBlock(ctx context.Context, contracts Contracts, flagFrom uint64) (uint64, error)
	ArchivedBlockNumbers(
		ctx context.Context,
		contracts Contracts,
		fromBlock uint64,
		toBlock uint64,
	) ([]int64, error)
	PrepareWriter(ctx context.Context) (NodeFillWriter, error)
}

// NodeFillWriter is the per-row archive persistence surface.
type NodeFillWriter interface {
	EventLogExists(ctx context.Context, txHash string, logIndex int) (bool, error)
	InsertEventLog(ctx context.Context, event EventLog) (bool, error)
	TransactionExists(ctx context.Context, txHash string) (bool, error)
	InsertTransaction(ctx context.Context, tx Transaction) (bool, error)
	BlockExists(ctx context.Context, blockNum int64, blockHash string) (bool, error)
	InsertBlock(
		ctx context.Context,
		block Block,
		projectAddresses []string,
		forceProjectCleanup bool,
	) (bool, error)
	Close() error
}

// EventLog is the row written to arch_evtlog by node-fill.
type EventLog struct {
	BlockNum        int64
	LogIndex        int
	TxHash          string
	ContractAddress string
	Topic0Sig       string
	LogRLP          []byte
}

// Transaction is the row written to arch_tx by node-fill.
type Transaction struct {
	BlockNum       int64
	FromAddressID  int64
	ToAddressID    int64
	GasUsed        int64
	TxIndex        int
	NumLogs        int
	ContractCreate bool
	Value          string
	GasPrice       string
	TxHash         string
	InputSig       string
}

// Block is the row written to arch_block by node-fill.
type Block struct {
	BlockNum   int64
	NumTx      int64
	Timestamp  uint64
	BlockHash  string
	ParentHash string
}

// NodeFiller runs context-aware node-fill operations.
type NodeFiller struct {
	Repository   NodeFillRepository
	AddressStore AddressStore
	Client       NodeClient
	Logger       Logger
	RetryDelay   time.Duration
	Sleep        func(ctx context.Context, delay time.Duration) error
	EncodeLog    func(logEntry *types.Log) ([]byte, error)
}

// RunProject scans and fills archive rows for one project. Setup errors abort.
// Per-log RPC and DB errors are counted and processing continues.
func (f *NodeFiller) RunProject(
	ctx context.Context,
	project string,
	options NodeFillOptions,
) (FillStats, error) {
	var stats FillStats
	if f.Repository == nil {
		return stats, errors.New("archive node-fill: repository is required")
	}
	if f.Client == nil {
		return stats, errors.New("archive node-fill: client is required")
	}
	if f.AddressStore == nil {
		return stats, errors.New("archive node-fill: address store is required")
	}
	if options.BatchSize == 0 {
		return stats, errors.New("archive node-fill: batch size must be greater than zero")
	}
	if err := ctx.Err(); err != nil {
		return stats, err
	}

	contracts, err := f.Repository.ProjectContracts(ctx, project)
	if err != nil {
		return stats, err
	}
	chainContracts := make([]common.Address, 0, len(contracts.Addresses))
	for _, address := range contracts.Addresses {
		chainContracts = append(chainContracts, common.HexToAddress(address))
	}
	f.printf("Contracts (%d): %v", len(chainContracts), contracts.Addresses)

	start, err := f.Repository.ResolveStartBlock(ctx, contracts, options.FromBlock)
	if err != nil {
		return stats, err
	}
	if options.FromBlock > 0 {
		f.printf("Start block: %d (from --start-block / --from)", start)
	} else {
		f.printf("Auto start block: %d (from address/evt_log metadata)", start)
	}
	if start > options.EndBlock {
		f.printf("start block %d > end %d — nothing to scan", start, options.EndBlock)
		return stats, nil
	}
	archivedBlocks, err := f.Repository.ArchivedBlockNumbers(
		ctx,
		contracts,
		start,
		options.EndBlock,
	)
	if err != nil {
		return stats, err
	}

	writer, err := f.Repository.PrepareWriter(ctx)
	if err != nil {
		return stats, err
	}
	writerClosed := false
	defer func() {
		if !writerClosed {
			_ = writer.Close()
		}
	}()

	retryDelay := f.RetryDelay
	if retryDelay <= 0 {
		retryDelay = defaultRetryDelay
	}
	lastBatch := options.BatchSize
	ensuredBlocks := make(map[int64]struct{})
	scanStats, scanErr := logscan.Scan(
		ctx,
		nodeFillFilterer{client: f.Client, logger: f.Logger},
		logscan.Options{
			FromBlock: start,
			ToBlock:   options.EndBlock,
			Query: ethereum.FilterQuery{
				Addresses: chainContracts,
			},
			InitialBatch: options.BatchSize,
			MinBatch:     min(options.BatchSize, minRetryBatchBlocks),
			RetryDelay:   retryDelay,
			Sleep:        f.Sleep,
			OnProgress: func(_ context.Context, progress logscan.Progress) error {
				if progress.BatchSize < lastBatch {
					f.printf("Reducing batch to %d blocks", progress.BatchSize)
				}
				lastBatch = progress.BatchSize
				f.printf("FilterLogs blocks %d .. %d", progress.FromBlock, progress.ToBlock)
				return nil
			},
		},
		func(ctx context.Context, logEntry types.Log) error {
			stats.LogsFromNode++

			blockNum := int64(logEntry.BlockNumber)
			if !options.DryRun {
				if _, done := ensuredBlocks[blockNum]; !done {
					insertedCount, skippedCount, kind, err := f.ensureBlock(
						ctx,
						writer,
						blockNum,
						contracts.Addresses,
					)
					if err != nil {
						if ctxErr := contextError(ctx, err); ctxErr != nil {
							return ctxErr
						}
						stats.countError(kind)
						f.printf("arch_block %d: %v", logEntry.BlockNumber, err)
						// Avoid persisting rows until this block's canonical
						// identity has been established.
						return nil
					}
					stats.BlockInserted += insertedCount
					stats.BlockSkipped += skippedCount
					ensuredBlocks[blockNum] = struct{}{}
				}
			}

			txHash := logEntry.TxHash.Hex()
			logIndex := int(logEntry.Index)
			exists, err := writer.EventLogExists(ctx, txHash, logIndex)
			if err != nil {
				if ctxErr := contextError(ctx, err); ctxErr != nil {
					return ctxErr
				}
				stats.DBErrors++
				f.printf("exists check: %v", err)
				return nil
			}
			if exists {
				stats.LogsSkipped++
			} else {
				if options.DryRun {
					stats.LogsInserted++
					return nil
				}

				encodeLog := f.EncodeLog
				if encodeLog == nil {
					encodeLog = toolutil.EncodeLogRLP
				}
				encodedLog, err := encodeLog(&logEntry)
				if err != nil {
					stats.DBErrors++
					f.printf("encode log: %v", err)
					return nil
				}
				inserted, err := writer.InsertEventLog(ctx, EventLog{
					BlockNum:        int64(logEntry.BlockNumber),
					LogIndex:        logIndex,
					TxHash:          txHash,
					ContractAddress: logEntry.Address.Hex(),
					Topic0Sig:       toolutil.Topic0Sig(&logEntry),
					LogRLP:          encodedLog,
				})
				if err != nil {
					if ctxErr := contextError(ctx, err); ctxErr != nil {
						return ctxErr
					}
					stats.DBErrors++
					f.printf("insert arch_evtlog %s:%d: %v", txHash, logIndex, err)
					return nil
				}
				if !inserted {
					stats.LogsSkipped++
				} else {
					stats.LogsInserted++
				}
			}
			if options.DryRun {
				return nil
			}

			// Event-log presence and archive dependencies are independent.
			// A previous run may have inserted arch_evtlog before an RPC or DB
			// failure prevented arch_tx/arch_block from being filled, so always
			// repair those rows even when this log was already present.
			insertedCount, skippedCount, kind, err := f.ensureTransaction(
				ctx, writer, txHash, int64(logEntry.BlockNumber),
			)
			if err != nil {
				if ctxErr := contextError(ctx, err); ctxErr != nil {
					return ctxErr
				}
				stats.countError(kind)
				f.printf("arch_tx %s: %v", txHash, err)
			} else {
				stats.TxInserted += insertedCount
				stats.TxSkipped += skippedCount
			}
			return nil
		},
	)
	if scanErr == nil && !options.DryRun {
		for _, blockNum := range archivedBlocks {
			if _, done := ensuredBlocks[blockNum]; done {
				continue
			}
			insertedCount, skippedCount, kind, err := f.ensureBlock(
				ctx,
				writer,
				blockNum,
				contracts.Addresses,
			)
			if err != nil {
				if ctxErr := contextError(ctx, err); ctxErr != nil {
					scanErr = ctxErr
					break
				}
				stats.countError(kind)
				f.printf("arch_block %d: %v", blockNum, err)
				continue
			}
			stats.BlockInserted += insertedCount
			stats.BlockSkipped += skippedCount
			ensuredBlocks[blockNum] = struct{}{}
		}
	}
	stats.BlocksScanned = scanStats.BlocksScanned
	stats.FilterRetries += int64(scanStats.FilterErrors)
	closeErr := writer.Close()
	writerClosed = true
	if scanErr != nil {
		return stats, errors.Join(scanErr, closeErr)
	}
	if closeErr != nil {
		return stats, fmt.Errorf("archive node-fill: close writer: %w", closeErr)
	}
	return stats, nil
}

type nodeFillFilterer struct {
	client NodeClient
	logger Logger
}

func (f nodeFillFilterer) FilterLogs(
	ctx context.Context,
	query ethereum.FilterQuery,
) ([]types.Log, error) {
	logs, err := f.client.FilterLogs(ctx, query)
	if err != nil && f.logger != nil {
		f.logger.Printf(
			"FilterLogs error [%d..%d]: %v",
			query.FromBlock.Uint64(),
			query.ToBlock.Uint64(),
			err,
		)
	}
	return logs, err
}

type errorKind uint8

const (
	dbError errorKind = iota
	rpcError
)

func (s *FillStats) countError(kind errorKind) {
	if kind == rpcError {
		s.RPCErrors++
		return
	}
	s.DBErrors++
}

func (f *NodeFiller) ensureTransaction(
	ctx context.Context,
	writer NodeFillWriter,
	txHash string,
	blockNum int64,
) (inserted, skipped int64, kind errorKind, err error) {
	exists, err := writer.TransactionExists(ctx, txHash)
	if err != nil {
		return 0, 0, dbError, err
	}
	if exists {
		return 0, 1, dbError, nil
	}

	hash := common.HexToHash(txHash)
	tx, pending, err := f.Client.TransactionByHash(ctx, hash)
	if err != nil {
		return 0, 0, rpcError, err
	}
	if tx == nil {
		return 0, 0, rpcError, fmt.Errorf("transaction %s returned no data", txHash)
	}
	if pending {
		return 0, 0, rpcError, fmt.Errorf("transaction %s still pending", txHash)
	}
	receipt, err := f.Client.TransactionReceipt(ctx, hash)
	if err != nil {
		return 0, 0, rpcError, err
	}
	if receipt == nil {
		return 0, 0, rpcError, fmt.Errorf("transaction receipt %s returned no data", txHash)
	}

	from, err := transactionSender(tx)
	if err != nil {
		return 0, 0, rpcError, err
	}
	fromAddressID, err := f.AddressStore.LookupOrCreateAddress(ctx, from.Hex(), blockNum, 0)
	if err != nil {
		return 0, 0, dbError, fmt.Errorf("from address %s: %w", from.Hex(), err)
	}
	var toAddressID int64
	if tx.To() != nil {
		toAddressID, err = f.AddressStore.LookupOrCreateAddress(ctx, tx.To().Hex(), blockNum, 0)
		if err != nil {
			return 0, 0, dbError, fmt.Errorf("to address %s: %w", tx.To().Hex(), err)
		}
	}
	gasPrice := tx.GasPrice()
	if tx.Type() == types.DynamicFeeTxType {
		gasPrice = tx.GasFeeCap()
	}
	inputSig := ""
	if len(tx.Data()) >= 4 {
		inputSig = "0x" + hex.EncodeToString(tx.Data()[:4])
	}

	insertedRow, err := writer.InsertTransaction(ctx, Transaction{
		BlockNum:       blockNum,
		FromAddressID:  fromAddressID,
		ToAddressID:    toAddressID,
		GasUsed:        int64(receipt.GasUsed),
		TxIndex:        int(receipt.TransactionIndex),
		NumLogs:        len(receipt.Logs),
		ContractCreate: tx.To() == nil,
		Value:          tx.Value().String(),
		GasPrice:       gasPrice.String(),
		TxHash:         txHash,
		InputSig:       inputSig,
	})
	if err != nil {
		return 0, 0, dbError, err
	}
	if insertedRow {
		return 1, 0, dbError, nil
	}
	return 0, 1, dbError, nil
}

func transactionSender(tx *types.Transaction) (common.Address, error) {
	if !tx.Protected() || tx.ChainId() == nil || tx.ChainId().Sign() <= 0 {
		return types.Sender(types.HomesteadSigner{}, tx)
	}
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err == nil {
		return from, nil
	}
	return types.Sender(types.HomesteadSigner{}, tx)
}

func (f *NodeFiller) ensureBlock(
	ctx context.Context,
	writer NodeFillWriter,
	blockNum int64,
	projectAddresses []string,
) (inserted, skipped int64, kind errorKind, err error) {
	block, err := f.Client.BlockByNumber(ctx, big.NewInt(blockNum))
	if err != nil {
		return 0, 0, rpcError, err
	}
	if block == nil {
		return 0, 0, rpcError, fmt.Errorf("block %d returned no data", blockNum)
	}
	header := block.Header()
	blockHash := header.Hash().Hex()
	exists, err := writer.BlockExists(ctx, blockNum, blockHash)
	if err != nil {
		return 0, 0, dbError, err
	}
	forceProjectCleanup := exists
	insertedRow, err := writer.InsertBlock(ctx, Block{
		BlockNum:   blockNum,
		NumTx:      int64(len(block.Transactions())),
		Timestamp:  header.Time,
		BlockHash:  blockHash,
		ParentHash: header.ParentHash.Hex(),
	}, projectAddresses, forceProjectCleanup)
	if err != nil {
		return 0, 0, dbError, err
	}
	if exists {
		return 0, 1, dbError, nil
	}
	if insertedRow {
		return 1, 0, dbError, nil
	}
	return 0, 1, dbError, nil
}

func (f *NodeFiller) printf(format string, args ...any) {
	if f.Logger != nil {
		f.Logger.Printf(format, args...)
	}
}

// SQLNodeFillRepository implements NodeFillRepository over one database/sql
// adapter. It does not own DB.
type SQLNodeFillRepository struct {
	DB *sql.DB
}

// ProjectContracts resolves the registered contract addresses for project.
func (r *SQLNodeFillRepository) ProjectContracts(ctx context.Context, project string) (Contracts, error) {
	return LoadProjectContracts(ctx, r.DB, project)
}

// ResolveStartBlock returns flagFrom when set, otherwise auto-detects the
// earliest block referenced by the project's addresses and events.
func (r *SQLNodeFillRepository) ResolveStartBlock(
	ctx context.Context,
	contracts Contracts,
	flagFrom uint64,
) (uint64, error) {
	if flagFrom > 0 {
		return flagFrom, nil
	}
	var fromAddress, fromEvent sql.NullInt64
	if err := r.DB.QueryRowContext(ctx, `
		SELECT MIN(block_num) FROM address WHERE addr = ANY($1)
	`, pq.Array(contracts.Addresses)).Scan(&fromAddress); err != nil {
		return 0, fmt.Errorf("read minimum contract block: %w", err)
	}
	if err := r.DB.QueryRowContext(ctx, `
		SELECT MIN(block_num) FROM evt_log WHERE contract_aid = ANY($1)
	`, pq.Array(contracts.AddressIDs)).Scan(&fromEvent); err != nil {
		return 0, fmt.Errorf("read minimum event block: %w", err)
	}
	return SelectStartBlock(flagFrom, fromAddress, fromEvent)
}

// ArchivedBlockNumbers lists the distinct block numbers already archived
// for the project's contracts within [fromBlock, toBlock].
func (r *SQLNodeFillRepository) ArchivedBlockNumbers(
	ctx context.Context,
	contracts Contracts,
	fromBlock uint64,
	toBlock uint64,
) ([]int64, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT DISTINCT block_num
		FROM arch_evtlog
		WHERE contract_addr = ANY($1)
		  AND block_num BETWEEN $2 AND $3
		ORDER BY block_num
	`, pq.Array(contracts.Addresses), fromBlock, toBlock)
	if err != nil {
		return nil, fmt.Errorf("read archived project blocks: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var blocks []int64
	for rows.Next() {
		var blockNum int64
		if err := rows.Scan(&blockNum); err != nil {
			return nil, fmt.Errorf("scan archived project block: %w", err)
		}
		blocks = append(blocks, blockNum)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("read archived project blocks: %w", err)
	}
	return blocks, nil
}

// SelectStartBlock applies the node-fill start policy to loaded metadata.
func SelectStartBlock(flagFrom uint64, fromAddress, fromEvent sql.NullInt64) (uint64, error) {
	if flagFrom > 0 {
		return flagFrom, nil
	}
	var candidates []int64
	if fromAddress.Valid && fromAddress.Int64 > 0 {
		candidates = append(candidates, fromAddress.Int64)
	}
	if fromEvent.Valid && fromEvent.Int64 > 0 {
		candidates = append(candidates, fromEvent.Int64)
	}
	if len(candidates) == 0 {
		return 0, errors.New("could not auto-detect start block; pass --from <deployment_block>")
	}
	minimum := candidates[0]
	for _, candidate := range candidates[1:] {
		if candidate < minimum {
			minimum = candidate
		}
	}
	return uint64(minimum), nil
}

// RequireArchLogIndex verifies the natural-key column required by node-fill.
func RequireArchLogIndex(ctx context.Context, db *sql.DB) error {
	var count int
	err := db.QueryRowContext(ctx, `
		SELECT COUNT(*) FROM information_schema.columns
		WHERE table_name = 'arch_evtlog' AND column_name = 'log_index'
	`).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("arch_evtlog.log_index missing — create the archive tables from db/layer1/archive_tables.sql first")
	}
	return nil
}

// PrepareWriter prepares the archive insert/exists statements and returns
// the writer that owns them; callers must Close it.
func (r *SQLNodeFillRepository) PrepareWriter(ctx context.Context) (NodeFillWriter, error) {
	writer := &sqlNodeFillWriter{}
	var err error
	if writer.insertEvent, err = r.DB.PrepareContext(ctx, `
		INSERT INTO arch_evtlog (block_num, evt_id, log_index, tx_hash, contract_addr, topic0_sig, log_rlp)
		VALUES ($1, NULL, $2, $3, $4, $5, $6)
		ON CONFLICT (tx_hash, log_index) DO NOTHING
	`); err != nil {
		return nil, fmt.Errorf("prepare arch_evtlog: %w", err)
	}
	if writer.eventExists, err = r.DB.PrepareContext(ctx,
		`SELECT 1 FROM arch_evtlog WHERE tx_hash = $1 AND log_index = $2`,
	); err != nil {
		_ = writer.Close()
		return nil, fmt.Errorf("prepare exists check: %w", err)
	}
	if writer.insertTx, err = r.DB.PrepareContext(ctx, `
		INSERT INTO arch_tx (block_num, from_aid, to_aid, gas_used, tx_index, num_logs, ctrct_create, value, gas_price, tx_hash, input_sig)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (tx_hash) DO NOTHING
	`); err != nil {
		_ = writer.Close()
		return nil, fmt.Errorf("prepare arch_tx: %w", err)
	}
	if writer.txExists, err = r.DB.PrepareContext(ctx,
		`SELECT 1 FROM arch_tx WHERE tx_hash = $1`,
	); err != nil {
		_ = writer.Close()
		return nil, fmt.Errorf("prepare tx exists: %w", err)
	}
	if writer.cleanupTx, err = r.DB.PrepareContext(ctx, `
		DELETE FROM arch_tx at
		WHERE at.block_num = $1
		  AND EXISTS (
			SELECT 1 FROM arch_evtlog stale
			WHERE stale.tx_hash = at.tx_hash
			  AND stale.block_num = $1
			  AND stale.contract_addr = ANY($3)
		  )
		  AND NOT EXISTS (
			SELECT 1 FROM arch_evtlog keep
			WHERE keep.tx_hash = at.tx_hash
			  AND NOT (keep.block_num = $1 AND keep.contract_addr = ANY($3))
		  )
		  AND ($4 OR EXISTS (
			SELECT 1 FROM arch_block
			WHERE block_num = $1 AND block_hash IS DISTINCT FROM $2
		  ))
	`); err != nil {
		_ = writer.Close()
		return nil, fmt.Errorf("prepare stale transaction cleanup: %w", err)
	}
	if writer.cleanupLogs, err = r.DB.PrepareContext(ctx, `
		DELETE FROM arch_evtlog
		WHERE block_num = $1
		  AND contract_addr = ANY($3)
		  AND ($4 OR EXISTS (
			SELECT 1 FROM arch_block
			WHERE block_num = $1 AND block_hash IS DISTINCT FROM $2
		  ))
	`); err != nil {
		_ = writer.Close()
		return nil, fmt.Errorf("prepare stale log cleanup: %w", err)
	}
	if writer.cleanupBlock, err = r.DB.PrepareContext(ctx, `
		DELETE FROM arch_block
		WHERE block_num = $1 AND block_hash IS DISTINCT FROM $2
	`); err != nil {
		_ = writer.Close()
		return nil, fmt.Errorf("prepare stale block cleanup: %w", err)
	}
	if writer.insertBlock, err = r.DB.PrepareContext(ctx, `
		INSERT INTO arch_block (block_num, num_tx, ts, cash_flow, block_hash, parent_hash)
		VALUES ($1, $2, TO_TIMESTAMP($3), 0, $4, $5)
		ON CONFLICT (block_hash) DO UPDATE SET
			block_num = EXCLUDED.block_num,
			num_tx = EXCLUDED.num_tx,
			ts = EXCLUDED.ts,
			parent_hash = EXCLUDED.parent_hash
	`); err != nil {
		_ = writer.Close()
		return nil, fmt.Errorf("prepare arch_block: %w", err)
	}
	if writer.blockExists, err = r.DB.PrepareContext(ctx,
		`SELECT 1 FROM arch_block WHERE block_num = $1 AND block_hash = $2`,
	); err != nil {
		_ = writer.Close()
		return nil, fmt.Errorf("prepare block exists: %w", err)
	}
	return writer, nil
}

type sqlNodeFillWriter struct {
	insertEvent  *sql.Stmt
	eventExists  *sql.Stmt
	insertTx     *sql.Stmt
	txExists     *sql.Stmt
	cleanupTx    *sql.Stmt
	cleanupLogs  *sql.Stmt
	cleanupBlock *sql.Stmt
	insertBlock  *sql.Stmt
	blockExists  *sql.Stmt
}

func (w *sqlNodeFillWriter) EventLogExists(ctx context.Context, txHash string, logIndex int) (bool, error) {
	return statementExists(ctx, w.eventExists, txHash, logIndex)
}

func (w *sqlNodeFillWriter) InsertEventLog(ctx context.Context, event EventLog) (bool, error) {
	result, err := w.insertEvent.ExecContext(ctx,
		event.BlockNum,
		event.LogIndex,
		event.TxHash,
		event.ContractAddress,
		event.Topic0Sig,
		event.LogRLP,
	)
	return rowInserted(result, err)
}

func (w *sqlNodeFillWriter) TransactionExists(ctx context.Context, txHash string) (bool, error) {
	return statementExists(ctx, w.txExists, txHash)
}

func (w *sqlNodeFillWriter) InsertTransaction(ctx context.Context, tx Transaction) (bool, error) {
	result, err := w.insertTx.ExecContext(ctx,
		tx.BlockNum,
		tx.FromAddressID,
		tx.ToAddressID,
		tx.GasUsed,
		tx.TxIndex,
		tx.NumLogs,
		tx.ContractCreate,
		tx.Value,
		tx.GasPrice,
		tx.TxHash,
		tx.InputSig,
	)
	return rowInserted(result, err)
}

func (w *sqlNodeFillWriter) BlockExists(ctx context.Context, blockNum int64, blockHash string) (bool, error) {
	return statementExists(ctx, w.blockExists, blockNum, blockHash)
}

func (w *sqlNodeFillWriter) InsertBlock(
	ctx context.Context,
	block Block,
	projectAddresses []string,
	forceProjectCleanup bool,
) (bool, error) {
	for _, statement := range []*sql.Stmt{w.cleanupTx, w.cleanupLogs} {
		if _, err := statement.ExecContext(
			ctx,
			block.BlockNum,
			block.BlockHash,
			pq.Array(projectAddresses),
			forceProjectCleanup,
		); err != nil {
			return false, err
		}
	}
	if _, err := w.cleanupBlock.ExecContext(ctx, block.BlockNum, block.BlockHash); err != nil {
		return false, err
	}
	result, err := w.insertBlock.ExecContext(ctx,
		block.BlockNum,
		block.NumTx,
		block.Timestamp,
		block.BlockHash,
		block.ParentHash,
	)
	return rowInserted(result, err)
}

func (w *sqlNodeFillWriter) Close() error {
	var statements []statementCloser
	for _, statement := range []*sql.Stmt{
		w.insertEvent,
		w.eventExists,
		w.insertTx,
		w.txExists,
		w.cleanupTx,
		w.cleanupLogs,
		w.cleanupBlock,
		w.insertBlock,
		w.blockExists,
	} {
		if statement != nil {
			statements = append(statements, statement)
		}
	}
	return closeStatements(statements)
}

type statementCloser interface {
	Close() error
}

func closeStatements(statements []statementCloser) error {
	var first error
	for _, statement := range statements {
		if err := statement.Close(); err != nil && first == nil {
			first = err
		}
	}
	return first
}

func statementExists(ctx context.Context, statement *sql.Stmt, args ...any) (bool, error) {
	var one int
	err := statement.QueryRowContext(ctx, args...).Scan(&one)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func rowInserted(result sql.Result, err error) (bool, error) {
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

func contextError(ctx context.Context, err error) error {
	if ctxErr := ctx.Err(); ctxErr != nil {
		return ctxErr
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return err
	}
	return nil
}

// LogFillStats writes the established opsctl stats block.
func LogFillStats(logger Logger, label string, stats FillStats) {
	if logger == nil {
		return
	}
	logger.Printf("[%s] blocks scanned: %d", label, stats.BlocksScanned)
	logger.Printf("[%s] logs from node: %d", label, stats.LogsFromNode)
	logger.Printf("[%s] arch_evtlog already present (skipped): %d", label, stats.LogsSkipped)
	logger.Printf("[%s] arch_evtlog inserted (or would insert): %d", label, stats.LogsInserted)
	logger.Printf("[%s] arch_tx inserted: %d, skipped (existed): %d", label, stats.TxInserted, stats.TxSkipped)
	logger.Printf("[%s] arch_block inserted: %d, skipped (existed): %d", label, stats.BlockInserted, stats.BlockSkipped)
	if stats.FilterRetries > 0 {
		logger.Printf("[%s] recovered FilterLogs retries: %d", label, stats.FilterRetries)
	}
	if stats.RPCErrors > 0 || stats.DBErrors > 0 {
		logger.Printf("[%s] errors — rpc: %d, db: %d", label, stats.RPCErrors, stats.DBErrors)
	}
}
