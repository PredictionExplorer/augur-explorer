// transaction-collector backs up RLP-encoded transactions and receipts for contract
// activity discovered via eth_getLogs (FilterLogs).
//
// Build: go build -o rwcg/tools/transaction-collector ./rwcg/tools/transaction_collector.go
//
// Example:
//
//	./transaction-collector -config ~/configs/transaction-collector.cosmicgame.json
//	./transaction-collector -config config.json -start-block 9292392
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	etlcommon "github.com/PredictionExplorer/augur-explorer/rwcg/etl/common"
	"github.com/PredictionExplorer/augur-explorer/rwcg/tools/toolutil"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

const defaultBatchBlocks = 100_000

type collectorStats struct {
	BlocksScanned   uint64
	LogsSeen        int64
	TxUnique        int64
	TxWritten       int64
	ReceiptWritten  int64
	TxSkippedExists int64
	TxMissingNode   int64
	TxFetchErrors   int64
	FilterLogErrors int64
}

func main() {
	configPath := flag.String("config", "", "Path to JSON config (rpc_url, output_dir, start_block, contract addresses)")
	startBlock := flag.Uint64("start-block", 0, "Override config start_block (scan from this block inclusive)")
	batchBlocks := flag.Uint64("batch", defaultBatchBlocks, "FilterLogs block range size")
	toBlock := flag.Uint64("to", 0, "End block inclusive (0 = chain head)")
	flag.Parse()

	if *configPath == "" {
		log.Fatal("Usage: transaction-collector -config <path.json> [-start-block N] [-to N] [-batch N]")
	}

	cfg, err := toolutil.LoadCollectorConfig(*configPath)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	if cfg.RPCURL == "" {
		log.Fatal("rpc_url is required in config")
	}
	fromBlock := cfg.StartBlock
	if *startBlock > 0 {
		fromBlock = *startBlock
	}

	addrs, err := cfg.ResolveContractAddresses()
	if err != nil {
		log.Fatalf("contracts: %v", err)
	}
	contracts := make([]ethcommon.Address, 0, len(addrs))
	for _, a := range addrs {
		contracts = append(contracts, ethcommon.HexToAddress(a))
	}

	if err := os.MkdirAll(cfg.OutputDir, 0o755); err != nil {
		log.Fatalf("output dir: %v", err)
	}

	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		log.Fatalf("rpc connect: %v", err)
	}

	head, err := etlcommon.GetCurrentBlockNumber(client)
	if err != nil {
		log.Fatalf("chain head: %v", err)
	}
	endBlock := head
	if *toBlock > 0 {
		endBlock = *toBlock
	}
	if fromBlock > endBlock {
		log.Printf("start block %d > end %d — nothing to scan", fromBlock, endBlock)
		return
	}

	log.Printf("RPC: %s", cfg.RPCURL)
	log.Printf("Output: %s", cfg.OutputDir)
	log.Printf("Contracts (%d): %v", len(contracts), addrs)
	log.Printf("Scanning blocks %d .. %d (batch %d)", fromBlock, endBlock, *batchBlocks)

	st := runCollector(client, cfg.OutputDir, contracts, fromBlock, endBlock, *batchBlocks)
	log.Printf("done: blocks_scanned=%d logs=%d unique_tx=%d tx_written=%d receipt_written=%d skipped_exists=%d missing_on_node=%d tx_errors=%d filter_errors=%d",
		st.BlocksScanned, st.LogsSeen, st.TxUnique, st.TxWritten, st.ReceiptWritten,
		st.TxSkippedExists, st.TxMissingNode, st.TxFetchErrors, st.FilterLogErrors)
}

func runCollector(client *ethclient.Client, outputDir string, contracts []ethcommon.Address, start, endBlock, batchSize uint64) collectorStats {
	var st collectorStats
	ctx := context.Background()
	seenTx := make(map[ethcommon.Hash]uint64)

	for from := start; from <= endBlock; {
		to := from + batchSize - 1
		if to > endBlock {
			to = endBlock
		}
		st.BlocksScanned += to - from + 1
		log.Printf("FilterLogs blocks %d .. %d", from, to)

		logs, err := etlcommon.FetchEvents(client, from, to, contracts)
		if err != nil {
			log.Printf("FilterLogs error [%d..%d]: %v", from, to, err)
			st.FilterLogErrors++
			if batchSize > 1000 {
				batchSize /= 2
				log.Printf("Reducing batch to %d blocks", batchSize)
				continue
			}
			time.Sleep(3 * time.Second)
			continue
		}

		for i := range logs {
			lg := &logs[i]
			if lg.Removed {
				continue
			}
			st.LogsSeen++
			h := lg.TxHash
			if _, ok := seenTx[h]; !ok {
				seenTx[h] = lg.BlockNumber
				st.TxUnique++
			}
		}

		from = to + 1
	}

	log.Printf("Collected %d unique transactions from logs; fetching tx + receipt …", len(seenTx))

	for h, blockNum := range seenTx {
		if err := storeTransactionPair(ctx, client, outputDir, h, blockNum, &st); err != nil {
			log.Printf("store %s: %v", h.Hex(), err)
		}
	}
	return st
}

func storeTransactionPair(ctx context.Context, client *ethclient.Client, outputDir string, txHash ethcommon.Hash, blockNum uint64, st *collectorStats) error {
	txPath := toolutil.TxRLPPath(outputDir, blockNum, txHash.Hex())
	rcptPath := toolutil.ReceiptRLPPath(outputDir, blockNum, txHash.Hex())

	txExists := fileExists(txPath)
	rcptExists := fileExists(rcptPath)
	if txExists && rcptExists {
		st.TxSkippedExists++
		return nil
	}

	tx, pending, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		if isMissingOnNodeError(err) {
			st.TxMissingNode++
			log.Printf("tx %s block %d: not on node (pruned?): %v", txHash.Hex(), blockNum, err)
			return nil
		}
		st.TxFetchErrors++
		log.Printf("tx %s: TransactionByHash: %v", txHash.Hex(), err)
		return nil
	}
	if pending {
		st.TxFetchErrors++
		log.Printf("tx %s: still pending", txHash.Hex())
		return nil
	}

	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		if isMissingOnNodeError(err) {
			st.TxMissingNode++
			log.Printf("tx %s block %d: receipt not on node: %v", txHash.Hex(), blockNum, err)
			return nil
		}
		st.TxFetchErrors++
		log.Printf("tx %s: TransactionReceipt: %v", txHash.Hex(), err)
		return nil
	}

	if !txExists {
		txRLP, err := rlp.EncodeToBytes(tx)
		if err != nil {
			return fmt.Errorf("rlp tx: %w", err)
		}
		if err := writeFileAtomic(txPath, txRLP); err != nil {
			return err
		}
		st.TxWritten++
	}

	if !rcptExists {
		rcptRLP, err := toolutil.EncodeBackupReceiptRLP(receipt)
		if err != nil {
			return fmt.Errorf("rlp receipt: %w", err)
		}
		if err := writeFileAtomic(rcptPath, rcptRLP); err != nil {
			return err
		}
		st.ReceiptWritten++
	}
	return nil
}

func isMissingOnNodeError(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "not found") ||
		strings.Contains(msg, "missing") ||
		strings.Contains(msg, "unknown transaction") ||
		strings.Contains(msg, "transaction indexing") ||
		strings.Contains(msg, "header not found")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func writeFileAtomic(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
