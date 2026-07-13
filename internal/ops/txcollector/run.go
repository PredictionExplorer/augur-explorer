// Package txcollector implements transaction/receipt backup collection and
// verification independently of command-line and configuration loading.
package txcollector

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer/logscan"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

const (
	// DefaultMinBatch is the smallest range used after FilterLogs failures.
	DefaultMinBatch uint64 = 1_000
	// DefaultRetryDelay is the delay between retries at DefaultMinBatch.
	DefaultRetryDelay = 3 * time.Second
)

// Client is the narrow Ethereum JSON-RPC surface used by Run.
type Client interface {
	logscan.Filterer
	TransactionByHash(context.Context, common.Hash) (*types.Transaction, bool, error)
	TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error)
}

// Logger receives operational warnings and progress. *log.Logger satisfies it.
type Logger interface {
	Printf(string, ...any)
}

// Config contains the stable dependencies and backup destination for Run.
type Config struct {
	Client    Client
	OutputDir string
	Contracts []common.Address
	Logger    Logger
}

// RunOptions bounds the inclusive chain scan and its retry behavior.
type RunOptions struct {
	FromBlock    uint64
	ToBlock      uint64
	InitialBatch uint64
	MinBatch     uint64
	RetryDelay   time.Duration

	Sleep      logscan.SleepFunc
	OnProgress logscan.ProgressFunc
}

// RunStats accumulates counters for one collection run.
type RunStats struct {
	BlocksScanned   uint64
	LogsSeen        uint64
	TxUnique        uint64
	TxWritten       uint64
	ReceiptWritten  uint64
	TxSkippedExists uint64
	TxMissingNode   uint64
	TxFetchErrors   uint64
	FilterLogErrors uint64
	BackupErrors    uint64
	InvalidBackups  uint64
}

type txRef struct {
	hash     common.Hash
	blockNum uint64
}

type pairEncoders struct {
	transaction func(*types.Transaction) ([]byte, error)
	receipt     func(*types.Receipt) ([]byte, error)
}

var defaultPairEncoders = pairEncoders{
	transaction: encodeTransactionRLP,
	receipt:     toolutil.EncodeBackupReceiptRLP,
}

// Run scans for contract activity and backs up every unique transaction and
// receipt. Unavailable historical transactions remain counted warnings, while
// local encoding/filesystem failures are aggregated and returned after all
// discoverable transactions have been attempted.
func Run(ctx context.Context, cfg Config, opts RunOptions) (RunStats, error) {
	var stats RunStats
	if err := validateRunConfig(cfg); err != nil {
		return stats, err
	}
	if err := validateRunOptions(opts); err != nil {
		return stats, err
	}
	if err := ctx.Err(); err != nil {
		return stats, err
	}
	if err := ensureDurableDirectory(cfg.OutputDir, 0o750); err != nil {
		return stats, fmt.Errorf("txcollector: create output directory: %w", err)
	}

	seen := make(map[common.Hash]uint64)
	lastBatch := opts.InitialBatch
	progress := func(ctx context.Context, p logscan.Progress) error {
		if p.BatchSize < lastBatch {
			logf(cfg.Logger, "Reducing batch to %d blocks", p.BatchSize)
		}
		lastBatch = p.BatchSize
		logf(cfg.Logger, "FilterLogs blocks %d .. %d", p.FromBlock, p.ToBlock)
		if opts.OnProgress != nil {
			return opts.OnProgress(ctx, p)
		}
		return nil
	}

	scanStats, err := logscan.Scan(ctx, loggingFilterer{client: cfg.Client, logger: cfg.Logger}, logscan.Options{
		FromBlock:    opts.FromBlock,
		ToBlock:      opts.ToBlock,
		Query:        ethereum.FilterQuery{Addresses: append([]common.Address(nil), cfg.Contracts...)},
		InitialBatch: opts.InitialBatch,
		MinBatch:     opts.MinBatch,
		RetryDelay:   opts.RetryDelay,
		Sleep:        opts.Sleep,
		OnProgress:   progress,
	}, func(_ context.Context, log types.Log) error {
		if _, exists := seen[log.TxHash]; !exists {
			seen[log.TxHash] = log.BlockNumber
		}
		return nil
	})
	stats.BlocksScanned = scanStats.BlocksScanned
	stats.LogsSeen = scanStats.LogsSeen
	stats.FilterLogErrors = scanStats.FilterErrors
	stats.TxUnique = uint64(len(seen))
	if err != nil {
		return stats, err
	}

	refs := make([]txRef, 0, len(seen))
	for hash, blockNum := range seen {
		refs = append(refs, txRef{hash: hash, blockNum: blockNum})
	}
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].blockNum != refs[j].blockNum {
			return refs[i].blockNum < refs[j].blockNum
		}
		return bytes.Compare(refs[i].hash[:], refs[j].hash[:]) < 0
	})

	logf(cfg.Logger, "Collected %d unique transactions from logs; fetching tx + receipt …", len(refs))
	var firstBackupErr error
	for _, ref := range refs {
		if err := ctx.Err(); err != nil {
			return stats, err
		}
		err := storeTransactionPair(ctx, cfg, ref, &stats)
		if err == nil {
			continue
		}
		if ctxErr := ctx.Err(); ctxErr != nil {
			return stats, ctxErr
		}
		logf(cfg.Logger, "store %s: %v", ref.hash.Hex(), err)
		stats.BackupErrors++
		if firstBackupErr == nil {
			firstBackupErr = fmt.Errorf("%s: %w", ref.hash.Hex(), err)
		}
	}
	if stats.BackupErrors > 0 {
		return stats, fmt.Errorf(
			"txcollector: %d transaction backup(s) failed: %w",
			stats.BackupErrors,
			firstBackupErr,
		)
	}
	return stats, nil
}

func validateRunConfig(cfg Config) error {
	switch {
	case cfg.Client == nil:
		return errors.New("txcollector: client is required")
	case cfg.OutputDir == "":
		return errors.New("txcollector: output directory is required")
	case len(cfg.Contracts) == 0:
		return errors.New("txcollector: at least one contract is required")
	default:
		return nil
	}
}

func validateRunOptions(opts RunOptions) error {
	switch {
	case opts.FromBlock > opts.ToBlock:
		return fmt.Errorf("txcollector: invalid range %d..%d", opts.FromBlock, opts.ToBlock)
	case opts.InitialBatch == 0:
		return errors.New("txcollector: initial batch must be greater than zero")
	case opts.MinBatch == 0:
		return errors.New("txcollector: minimum batch must be greater than zero")
	case opts.InitialBatch < opts.MinBatch:
		return fmt.Errorf(
			"txcollector: initial batch %d is smaller than minimum batch %d",
			opts.InitialBatch,
			opts.MinBatch,
		)
	case opts.RetryDelay <= 0:
		return errors.New("txcollector: retry delay must be greater than zero")
	default:
		return nil
	}
}

type loggingFilterer struct {
	client logscan.Filterer
	logger Logger
}

func (f loggingFilterer) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	logs, err := f.client.FilterLogs(ctx, query)
	if err != nil {
		logf(
			f.logger,
			"FilterLogs error [%d..%d]: %v",
			query.FromBlock.Uint64(),
			query.ToBlock.Uint64(),
			err,
		)
	}
	return logs, err
}

func storeTransactionPair(ctx context.Context, cfg Config, ref txRef, stats *RunStats) error {
	return storeTransactionPairWithEncoders(ctx, cfg, ref, stats, defaultPairEncoders)
}

func storeTransactionPairWithEncoders(
	ctx context.Context,
	cfg Config,
	ref txRef,
	stats *RunStats,
	encoders pairEncoders,
) error {
	txPath := toolutil.TxRLPPath(cfg.OutputDir, ref.blockNum, ref.hash.Hex())
	receiptPath := toolutil.ReceiptRLPPath(cfg.OutputDir, ref.blockNum, ref.hash.Hex())

	txExists, err := fileExists(txPath)
	if err != nil {
		return fmt.Errorf("stat tx backup: %w", err)
	}
	receiptExists, err := fileExists(receiptPath)
	if err != nil {
		return fmt.Errorf("stat receipt backup: %w", err)
	}
	repairNeeded := false
	if txExists {
		if err := validateTransactionBackup(txPath, ref.hash); err != nil {
			logf(cfg.Logger, "invalid tx backup %s: %v; replacing", txPath, err)
			stats.InvalidBackups++
			txExists = false
			repairNeeded = true
		}
	}
	if receiptExists {
		if err := validateReceiptBackup(receiptPath); err != nil {
			logf(cfg.Logger, "invalid receipt backup %s: %v; replacing", receiptPath, err)
			stats.InvalidBackups++
			receiptExists = false
			repairNeeded = true
		}
	}
	if txExists && receiptExists {
		stats.TxSkippedExists++
		return nil
	}

	tx, pending, err := cfg.Client.TransactionByHash(ctx, ref.hash)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return ctxErr
		}
		if isMissingOnNodeError(err) {
			stats.TxMissingNode++
			logf(cfg.Logger, "tx %s block %d: not on node (pruned?): %v", ref.hash.Hex(), ref.blockNum, err)
			if repairNeeded {
				return fmt.Errorf("repair invalid backup: transaction unavailable: %w", err)
			}
			return nil
		}
		stats.TxFetchErrors++
		logf(cfg.Logger, "tx %s: TransactionByHash: %v", ref.hash.Hex(), err)
		if repairNeeded {
			return fmt.Errorf("repair invalid backup: TransactionByHash: %w", err)
		}
		return nil
	}
	if pending {
		stats.TxFetchErrors++
		logf(cfg.Logger, "tx %s: still pending", ref.hash.Hex())
		if repairNeeded {
			return errors.New("repair invalid backup: transaction is still pending")
		}
		return nil
	}
	if tx == nil {
		stats.TxFetchErrors++
		logf(cfg.Logger, "tx %s: TransactionByHash returned no transaction", ref.hash.Hex())
		if repairNeeded {
			return errors.New("repair invalid backup: TransactionByHash returned no transaction")
		}
		return nil
	}

	receipt, err := cfg.Client.TransactionReceipt(ctx, ref.hash)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return ctxErr
		}
		if isMissingOnNodeError(err) {
			stats.TxMissingNode++
			logf(cfg.Logger, "tx %s block %d: receipt not on node: %v", ref.hash.Hex(), ref.blockNum, err)
			if repairNeeded {
				return fmt.Errorf("repair invalid backup: receipt unavailable: %w", err)
			}
			return nil
		}
		stats.TxFetchErrors++
		logf(cfg.Logger, "tx %s: TransactionReceipt: %v", ref.hash.Hex(), err)
		if repairNeeded {
			return fmt.Errorf("repair invalid backup: TransactionReceipt: %w", err)
		}
		return nil
	}
	if receipt == nil {
		stats.TxFetchErrors++
		logf(cfg.Logger, "tx %s: TransactionReceipt returned no receipt", ref.hash.Hex())
		if repairNeeded {
			return errors.New("repair invalid backup: TransactionReceipt returned no receipt")
		}
		return nil
	}

	if !txExists {
		txRLP, err := encoders.transaction(tx)
		if err != nil {
			return fmt.Errorf("encode transaction RLP: %w", err)
		}
		if err := writeFileAtomic(txPath, txRLP); err != nil {
			return fmt.Errorf("write transaction backup: %w", err)
		}
		stats.TxWritten++
	}

	if !receiptExists {
		receiptRLP, err := encoders.receipt(receipt)
		if err != nil {
			return fmt.Errorf("encode receipt RLP: %w", err)
		}
		if err := writeFileAtomic(receiptPath, receiptRLP); err != nil {
			return fmt.Errorf("write receipt backup: %w", err)
		}
		stats.ReceiptWritten++
	}
	return nil
}

func validateTransactionBackup(path string, expectedHash common.Hash) error {
	// #nosec G304 -- path is rooted in operator config and uses a fixed-format hash filename.
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var tx types.Transaction
	if err := rlp.DecodeBytes(data, &tx); err != nil {
		return fmt.Errorf("decode transaction RLP: %w", err)
	}
	if tx.Hash() != expectedHash {
		return fmt.Errorf(
			"transaction hash mismatch: decoded=%s expected=%s",
			tx.Hash().Hex(),
			expectedHash.Hex(),
		)
	}
	return nil
}

func validateReceiptBackup(path string) error {
	// #nosec G304 -- path is rooted in operator config and uses a fixed-format hash filename.
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if _, _, err := toolutil.TryDecodeReceiptRLP(data); err != nil {
		return fmt.Errorf("decode receipt RLP: %w", err)
	}
	return nil
}

func encodeTransactionRLP(tx *types.Transaction) (data []byte, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("malformed transaction: %v", recovered)
		}
	}()
	return rlp.EncodeToBytes(tx)
}

func isMissingOnNodeError(err error) bool {
	if err == nil {
		return false
	}
	message := strings.ToLower(err.Error())
	return strings.Contains(message, "not found") ||
		strings.Contains(message, "missing") ||
		strings.Contains(message, "unknown transaction") ||
		strings.Contains(message, "transaction indexing") ||
		strings.Contains(message, "header not found")
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	default:
		return false, err
	}
}

type atomicFile interface {
	Name() string
	Write([]byte) (int, error)
	Chmod(fs.FileMode) error
	Sync() error
	Close() error
}

type atomicWriteOps struct {
	mkdirAll   func(string, fs.FileMode) error
	createTemp func(string, string) (atomicFile, error)
	remove     func(string) error
	rename     func(string, string) error
	syncParent func(string) error
	syncDir    func(string) error
}

var defaultAtomicWriteOps = atomicWriteOps{
	mkdirAll: os.MkdirAll,
	createTemp: func(dir, pattern string) (atomicFile, error) {
		return os.CreateTemp(dir, pattern)
	},
	remove:     os.Remove,
	rename:     os.Rename,
	syncParent: syncDirectory,
	syncDir:    syncDirectory,
}

func writeFileAtomic(path string, data []byte) error {
	return writeFileAtomicWithOps(path, data, defaultAtomicWriteOps)
}

func writeFileAtomicWithOps(path string, data []byte, ops atomicWriteOps) (err error) {
	dir := filepath.Dir(path)
	if err := ops.mkdirAll(dir, 0o750); err != nil {
		return err
	}
	if err := ops.syncParent(filepath.Dir(dir)); err != nil {
		return err
	}

	file, err := ops.createTemp(dir, "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return err
	}
	tempPath := file.Name()
	closed := false
	defer func() {
		if !closed {
			_ = file.Close()
		}
		_ = ops.remove(tempPath)
	}()

	if _, err := file.Write(data); err != nil {
		return err
	}
	if err := file.Chmod(0o640); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	closed = true
	if err := ops.rename(tempPath, path); err != nil {
		return err
	}
	return ops.syncDir(dir)
}

func syncDirectory(path string) error {
	directory, err := os.Open(path) // #nosec G304 -- path is the configured backup directory.
	if err != nil {
		return err
	}
	defer func() { _ = directory.Close() }()
	return directory.Sync()
}

func ensureDurableDirectory(path string, mode fs.FileMode) error {
	cleanPath := filepath.Clean(path)
	var missing []string
	for current := cleanPath; ; current = filepath.Dir(current) {
		_, err := os.Stat(current)
		if err == nil {
			break
		}
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		missing = append(missing, current)
		parent := filepath.Dir(current)
		if parent == current {
			break
		}
	}
	if err := os.MkdirAll(cleanPath, mode); err != nil {
		return err
	}
	// Persist each newly created directory entry from the outermost parent
	// inward, so first-run backup roots survive a sudden power loss.
	for index := len(missing) - 1; index >= 0; index-- {
		if err := syncDirectory(filepath.Dir(missing[index])); err != nil {
			return err
		}
	}
	return nil
}

func logf(logger Logger, format string, args ...any) {
	if logger != nil {
		logger.Printf(format, args...)
	}
}
