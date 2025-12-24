// Package common - Event fetching using FilterLogs
package common

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// FetchEvents retrieves events from the blockchain using FilterLogs
// for the specified contract addresses in the given block range
func FetchEvents(client *ethclient.Client, fromBlock, toBlock uint64, contracts []common.Address) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: contracts,
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("FilterLogs failed: %v", err)
	}

	return logs, nil
}

// FetchEventsFromBlock retrieves events from a single block
func FetchEventsFromBlock(client *ethclient.Client, blockNum uint64, contracts []common.Address) ([]types.Log, error) {
	return FetchEvents(client, blockNum, blockNum, contracts)
}

// FetchEventsBatch retrieves events in batches to avoid timeout
// maxBlocksPerBatch limits how many blocks are queried at once
func FetchEventsBatch(client *ethclient.Client, fromBlock, toBlock uint64, contracts []common.Address, maxBlocksPerBatch uint64) ([]types.Log, error) {
	var allLogs []types.Log

	for start := fromBlock; start <= toBlock; start += maxBlocksPerBatch {
		end := start + maxBlocksPerBatch - 1
		if end > toBlock {
			end = toBlock
		}

		logs, err := FetchEvents(client, start, end, contracts)
		if err != nil {
			return nil, fmt.Errorf("FetchEvents failed for range %d-%d: %v", start, end, err)
		}

		allLogs = append(allLogs, logs...)
	}

	return allLogs, nil
}

// GetCurrentBlockNumber returns the current block number from the chain
func GetCurrentBlockNumber(client *ethclient.Client) (uint64, error) {
	return client.BlockNumber(context.Background())
}
