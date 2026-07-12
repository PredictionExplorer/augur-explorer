package main

// Shared block-range log scanning used by the scan-mints, scan-transfers and
// verify-erc20-transfers subcommands: fetch logs for one contract + topic0
// in fixed-size ranges from a start block to the chain tip.

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// logScanClient is the narrow Ethereum client surface the scans need;
// *ethclient.Client satisfies it and tests fake it.
type logScanClient interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
}

// scanLogsByRange walks the chain from startBlock to the tip in step-sized
// ranges, fetching logs emitted by addr with the given topic0 and calling
// handle for each non-removed log. onRange, when non-nil, reports each range
// before it is fetched (progress output).
func scanLogsByRange(
	ctx context.Context,
	client logScanClient,
	addr common.Address,
	topic0 common.Hash,
	startBlock, step int64,
	onRange func(from, to int64),
	handle func(*types.Log) error,
) error {
	latestBlock, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("error getting latest block: %w", err)
	}
	latestBnum := latestBlock.Number.Int64()

	filter := ethereum.FilterQuery{
		Topics:    [][]common.Hash{{topic0}},
		Addresses: []common.Address{addr},
	}
	for fromBlock := startBlock; fromBlock <= latestBnum; fromBlock += step {
		filter.FromBlock = big.NewInt(fromBlock)
		filter.ToBlock = big.NewInt(fromBlock + step)
		if onRange != nil {
			onRange(fromBlock, fromBlock+step)
		}
		logs, err := client.FilterLogs(ctx, filter)
		if err != nil {
			return fmt.Errorf("error querying events: %w", err)
		}
		for i := range logs {
			if logs[i].Removed {
				continue
			}
			if err := handle(&logs[i]); err != nil {
				return err
			}
		}
	}
	return nil
}
