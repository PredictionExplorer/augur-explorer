// FilterLogs polling loop: fetches event batches from the chain, stores them and
// dispatches each log to the event processors, updating processing status per batch.
package main

import (
	"context"
	"fmt"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	etlcommon "github.com/PredictionExplorer/augur-explorer/internal/etl"
)

// getContractAddresses returns all contract addresses for FilterLogs
func getContractAddresses() []ethcommon.Address {
	return []ethcommon.Address{
		rwalk_addr,
		market_addr,
	}
}

// process_events_filterlog uses FilterLogs to get events directly from blockchain.
// It returns nil on user-requested shutdown and an error when a batch hits an
// unrecoverable DB failure (the caller logs it and exits with status 1).
func process_events_filterlog(ctx context.Context) error {
	// Create ETL context for common operations
	etl_ctx := &etlcommon.ETLContext{
		Storage:   storage,
		EthClient: eclient,
		RpcClient: rpcclient,
		Info:      Info,
		Error:     Error,
	}

	// Adaptive batch sizing: start large, reduce if we get events
	var batchSize uint64 = 100000     // Start with 100k blocks
	var minBatchSize uint64 = 1000    // Minimum when processing events
	var maxBatchSize uint64 = 1000000 // Maximum when scanning empty ranges
	contracts := getContractAddresses()

	// Debug: log the addresses being used for FilterLogs
	Info.Printf("FilterLogs will query these contract addresses:\n")
	for i, addr := range contracts {
		Info.Printf("  [%d] %s\n", i, addr.Hex())
	}

	for {
		select {
		case <-ctx.Done():
			Info.Println("Exiting by user request.")
			return nil
		default:
		}

		// Get last processed block from status
		status := storagew.Get_randomwalk_processing_status()
		lastProcessedBlock := status.LastBlockNum
		if lastProcessedBlock == 0 {
			// If no blocks processed yet, start from the block where contracts were deployed
			lastProcessedBlock, _ = storage.Get_last_block_num()
		}

		// Get current block from chain
		currentBlock, err := etlcommon.GetCurrentBlockNumber(eclient)
		if err != nil {
			Error.Printf("Failed to get current block number: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Calculate block range to process
		fromBlock := uint64(lastProcessedBlock + 1)
		toBlock := fromBlock + batchSize - 1
		if toBlock > currentBlock {
			toBlock = currentBlock
		}

		if fromBlock > currentBlock {
			// Already caught up, wait for new blocks
			time.Sleep(2 * time.Second)
			batchSize = minBatchSize // Reset to small batch for real-time
			continue
		}

		Info.Printf("Fetching events from block %d to %d (batch size: %d)\n", fromBlock, toBlock, batchSize)

		// Fetch events using FilterLogs
		logs, err := etlcommon.FetchEvents(eclient, fromBlock, toBlock, contracts)
		if err != nil {
			Error.Printf("FetchEvents failed: %v", err)
			// Reduce batch size on error (might be too large)
			batchSize = batchSize / 2
			if batchSize < minBatchSize {
				batchSize = minBatchSize
			}
			time.Sleep(5 * time.Second)
			continue
		}

		Info.Printf("Received %d events\n", len(logs))

		// Process each event
		var processingFailed bool
		var lastSuccessfulBlock uint64
		for _, log := range logs {
			// Ensure block exists with correct hash (chain split detection)
			_, err := etlcommon.EnsureBlockExists(etl_ctx, int64(log.BlockNumber), log.BlockHash.Hex())
			if err != nil {
				Error.Printf("EnsureBlockExists failed for block %d: %v", log.BlockNumber, err)
				processingFailed = true
				time.Sleep(5 * time.Second)
				break
			}

			// Ensure transaction exists
			txId, _, err := etlcommon.EnsureTransactionExists(etl_ctx, log.TxHash, int64(log.BlockNumber))
			if err != nil {
				Error.Printf("EnsureTransactionExists failed for tx %s: %v", log.TxHash.Hex(), err)
				processingFailed = true
				time.Sleep(5 * time.Second)
				break
			}

			// Insert event log
			evtId, err := etlcommon.InsertEventLog(etl_ctx, log, txId)
			if err != nil {
				Error.Printf("InsertEventLog failed: %v", err)
				processingFailed = true
				time.Sleep(5 * time.Second)
				break
			}

			// Process the event using existing logic. Errors here are DB
			// failures propagated from the store layer (which previously
			// terminated the process from inside the store); terminate the
			// loop without updating the processing status so the batch is
			// re-processed on restart, exactly as before.
			err = process_single_event(evtId)
			if err != nil {
				Error.Printf("process_single_event failed for evt %d: %v", evtId, err)
				return fmt.Errorf("processing event %d: %w", evtId, err)
			}

			// Track last successfully processed block
			lastSuccessfulBlock = log.BlockNumber
		}

		// Only update status if processing succeeded
		if !processingFailed {
			status.LastBlockNum = int64(toBlock)
			storagew.Update_randomwalk_process_status(&status)
		} else if lastSuccessfulBlock > 0 {
			// Update to last successfully processed block
			status.LastBlockNum = int64(lastSuccessfulBlock)
			storagew.Update_randomwalk_process_status(&status)
		}
		// If processingFailed and lastSuccessfulBlock==0, don't update - will retry same batch

		// Adaptive batch sizing
		if len(logs) == 0 {
			// No events - increase batch size for faster scanning
			batchSize = batchSize * 2
			if batchSize > maxBatchSize {
				batchSize = maxBatchSize
			}
		} else {
			// Found events - use smaller batch for granularity
			batchSize = minBatchSize
		}
	}
}
