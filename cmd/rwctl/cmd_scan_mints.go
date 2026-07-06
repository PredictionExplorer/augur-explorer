package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// Block-range scanning parameters used by scan-mints (same as the legacy
// scan_rwmints script).
const (
	scanMintsStartBlock = int64(2000000)
	scanMintsMaxBlocks  = int64(1024 * 1024)
)

// checkMintLog verifies that the minted token from a MintEvent log exists in
// the database, retrying on transient database errors.
func checkMintLog(storagew *rwstore.SQLStorageWrapper, lg *types.Log) {
	tokenID := lg.Topics[1].Big().Int64()
	for {
		exists, err := storagew.Check_rwalk_token_exists(tokenID)
		if err != nil {
			fmt.Printf("Error accessing database: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if !exists {
			fmt.Printf("Token %v DOES NOT exist in the DB\n", tokenID)
		}
		return
	}
}

// newScanMintsCmd builds the scan-mints subcommand (legacy scan_rwmints script).
func newScanMintsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "scan-mints [randomwalk_addr]",
		Short: "Scan the chain for MintEvent logs and check DB presence",
		Long: "Scans the chain for RandomWalk MintEvent logs and reports tokens missing from the database.\n\n" +
			"Environment variables:\n  RPC_URL  Ethereum RPC endpoint (required)\n  PGSQL_*  PostgreSQL connection (required)",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			rwalkAddr := common.HexToAddress(args[0])
			eclient, err := dialEthClient()
			if err != nil {
				return err
			}
			storagew, err := connectRWStorage(newInfoLogger())
			if err != nil {
				return err
			}

			ctx := context.Background()
			latestBlock, err := eclient.HeaderByNumber(ctx, nil)
			if err != nil {
				return fmt.Errorf("error getting latest block: %w", err)
			}
			latestBnum := latestBlock.Number.Int64()

			filter := ethereum.FilterQuery{
				Topics:    [][]common.Hash{{mintEventTopic}},
				Addresses: []common.Address{rwalkAddr},
			}
			for fromBlock := scanMintsStartBlock; fromBlock <= latestBnum; fromBlock += scanMintsMaxBlocks {
				filter.FromBlock = big.NewInt(fromBlock)
				filter.ToBlock = big.NewInt(fromBlock + scanMintsMaxBlocks)
				fmt.Printf("From %v , to %v\n", filter.FromBlock.Int64(), filter.ToBlock.Int64())
				logs, err := eclient.FilterLogs(ctx, filter)
				if err != nil {
					return fmt.Errorf("error querying events: %w", err)
				}
				for i := range logs {
					if logs[i].Removed {
						continue
					}
					checkMintLog(storagew, &logs[i])
				}
			}
			return nil
		},
	}
}

func init() { register(newScanMintsCmd()) }
