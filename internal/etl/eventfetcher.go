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

// GetCurrentBlockNumber returns the current block number from the chain
func GetCurrentBlockNumber(client *ethclient.Client) (uint64, error) {
	return client.BlockNumber(context.Background())
}
