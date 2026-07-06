package main

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// verifyTransfersMaxBlocks is the block-range step used by
// verify-erc20-transfers (same as the legacy verify_erc20_transfers script).
const verifyTransfersMaxBlocks = int64(1024)

// checkTransferLog verifies that a chain Transfer log has a matching record
// in the rw_transfer table and prints a message when it is missing.
func checkTransferLog(storagew *rwstore.SQLStorageWrapper, lg *types.Log) {
	from := common.BytesToAddress(lg.Topics[1][12:]).String()
	to := common.BytesToAddress(lg.Topics[2][12:]).String()
	tokenID := lg.Topics[3].Big().Int64()
	transfers := storagew.Get_rw_token_transfers_by_tx_hash(lg.TxHash.String())
	found := false
	for i := 0; i < len(transfers); i++ {
		item := &transfers[i]
		if from == item.From && to == item.To && tokenID == item.TokenId {
			found = true
		}
	}
	if !found {
		fmt.Printf(
			"Block %v: Transfer %v -> %v (tok %v) tx %v, transfer missing\n",
			lg.BlockNumber, from, to, tokenID, lg.TxHash.String(),
		)
	}
}

// newVerifyTransfersCmd builds the verify-erc20-transfers subcommand (legacy
// verify_erc20_transfers script; despite the name it checks the ERC-721
// Transfer logs of the RandomWalk NFT against the rw_transfer table).
func newVerifyTransfersCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "verify-erc20-transfers",
		Short: "Verify chain Transfer logs match the rw_transfer table",
		Long: "Verifies that all transfers on chain have a corresponding entry in the rw_transfer table.\n\n" +
			"Environment variables:\n  RPC_URL  Ethereum RPC endpoint (required)\n  PGSQL_*  PostgreSQL connection (required)",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			storagew, err := connectRWStorage(newInfoLogger())
			if err != nil {
				return err
			}
			caddrs := storagew.Get_randomwalk_contract_addresses()
			rwalkAddr := common.HexToAddress(caddrs.RandomWalk)

			eclient, err := dialEthClient()
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
				Topics:    [][]common.Hash{{transferEventTopic}},
				Addresses: []common.Address{rwalkAddr},
			}
			for fromBlock := int64(0); fromBlock <= latestBnum; fromBlock += verifyTransfersMaxBlocks {
				filter.FromBlock = big.NewInt(fromBlock)
				filter.ToBlock = big.NewInt(fromBlock + verifyTransfersMaxBlocks)
				logs, err := eclient.FilterLogs(ctx, filter)
				if err != nil {
					return fmt.Errorf("error querying events: %w", err)
				}
				for i := range logs {
					if logs[i].Removed {
						continue
					}
					checkTransferLog(storagew, &logs[i])
				}
			}
			return nil
		},
	}
}

func init() { register(newVerifyTransfersCmd()) }
