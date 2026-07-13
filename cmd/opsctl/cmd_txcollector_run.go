package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/txcollector"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

type txCollectorRunRPC interface {
	txcollector.Client
	BlockNumber(context.Context) (uint64, error)
	Close()
}

type txCollectorRunDeps struct {
	loadConfig func(string) (*toolutil.CollectorConfig, error)
	dialRPC    func(context.Context, string) (txCollectorRunRPC, error)
	run        func(context.Context, txcollector.Config, txcollector.RunOptions) (txcollector.RunStats, error)
}

func defaultTxCollectorRunDeps() txCollectorRunDeps {
	return txCollectorRunDeps{
		loadConfig: toolutil.LoadCollectorConfig,
		dialRPC: func(ctx context.Context, rpcURL string) (txCollectorRunRPC, error) {
			return ethclient.DialContext(ctx, rpcURL)
		},
		run: txcollector.Run,
	}
}

// newTxCollectorRunCmd builds `opsctl tx-collector run`, the replacement for
// the standalone transaction-collector tool.
func newTxCollectorRunCmd() *cobra.Command {
	return newTxCollectorRunCmdWithDeps(defaultTxCollectorRunDeps())
}

func newTxCollectorRunCmdWithDeps(deps txCollectorRunDeps) *cobra.Command {
	var (
		configPath  string
		startBlock  uint64
		batchBlocks uint64
		toBlock     uint64
	)
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Back up RLP-encoded transactions and receipts for contract activity",
		Long: `Discovers contract activity via eth_getLogs (FilterLogs) and writes each
transaction and receipt as RLP blobs under the config's output_dir
(<output_dir>/<block_num>/<tx_hash>_tx.rlp and ..._receipt.rlp).

The JSON config carries rpc_url, output_dir, start_block and the contract
addresses; see cmd/opsctl/tx-collector.example.json for the format.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(cmd.ErrOrStderr(), "", log.LstdFlags)
			cfg, err := deps.loadConfig(configPath)
			if err != nil {
				return fmt.Errorf("config: %w", err)
			}
			if cfg.RPCURL == "" {
				return errors.New("rpc_url is required in config")
			}
			fromBlock := cfg.StartBlock
			if startBlock > 0 {
				fromBlock = startBlock
			}

			addrs, err := cfg.ResolveContractAddresses()
			if err != nil {
				return fmt.Errorf("contracts: %w", err)
			}
			contracts := make([]ethcommon.Address, 0, len(addrs))
			for _, a := range addrs {
				contracts = append(contracts, ethcommon.HexToAddress(a))
			}

			client, err := deps.dialRPC(cmd.Context(), cfg.RPCURL)
			if err != nil {
				return fmt.Errorf("rpc connect: %w", err)
			}
			defer client.Close()

			head, err := client.BlockNumber(cmd.Context())
			if err != nil {
				return fmt.Errorf("chain head: %w", err)
			}
			endBlock := head
			if toBlock > 0 {
				endBlock = toBlock
			}
			if fromBlock > endBlock {
				logger.Printf("start block %d > end %d — nothing to scan", fromBlock, endBlock)
				return nil
			}

			logger.Printf("RPC: %s", redactRPCURL(cfg.RPCURL))
			logger.Printf("Output: %s", cfg.OutputDir)
			logger.Printf("Contracts (%d): %v", len(contracts), addrs)
			logger.Printf("Scanning blocks %d .. %d (batch %d)", fromBlock, endBlock, batchBlocks)

			minBatch := min(batchBlocks, txcollector.DefaultMinBatch)
			stats, err := deps.run(cmd.Context(), txcollector.Config{
				Client:    client,
				OutputDir: cfg.OutputDir,
				Contracts: contracts,
				Logger:    logger,
			}, txcollector.RunOptions{
				FromBlock:    fromBlock,
				ToBlock:      endBlock,
				InitialBatch: batchBlocks,
				MinBatch:     minBatch,
				RetryDelay:   txcollector.DefaultRetryDelay,
			})
			logger.Printf(
				"done: blocks_scanned=%d logs=%d unique_tx=%d tx_written=%d receipt_written=%d skipped_exists=%d missing_on_node=%d tx_errors=%d filter_errors=%d invalid_backups=%d backup_errors=%d",
				stats.BlocksScanned,
				stats.LogsSeen,
				stats.TxUnique,
				stats.TxWritten,
				stats.ReceiptWritten,
				stats.TxSkippedExists,
				stats.TxMissingNode,
				stats.TxFetchErrors,
				stats.FilterLogErrors,
				stats.InvalidBackups,
				stats.BackupErrors,
			)
			return err
		},
	}
	cmd.Flags().StringVar(&configPath, "config", "", "Path to JSON config (rpc_url, output_dir, start_block, contract addresses)")
	cmd.Flags().Uint64Var(&startBlock, "start-block", 0, "Override config start_block (scan from this block inclusive)")
	cmd.Flags().Uint64Var(&batchBlocks, "batch", txCollectorDefaultFilterBatchBlocks, "FilterLogs block range size")
	cmd.Flags().Uint64Var(&toBlock, "to", 0, "End block inclusive (0 = chain head)")
	_ = cmd.MarkFlagRequired("config")
	return cmd
}
