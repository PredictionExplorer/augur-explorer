package main

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// Scan parameters preserved from the legacy scan_transfers script. The token
// filter was hardcoded there and is kept as the default for compatibility.
const (
	scanTransfersMaxBlocks    = int64(1024)
	scanTransfersDefaultToken = int64(3601)
)

// newScanTransfersCmd builds the scan-transfers subcommand (legacy
// scan_transfers script).
func newScanTransfersCmd() *cobra.Command {
	tokenID := scanTransfersDefaultToken
	c := &cobra.Command{
		Use:   "scan-transfers [randomwalk_addr]",
		Short: "Scan the chain for Transfer logs of a RandomWalk token",
		Long: "Scans the chain for ERC-721 Transfer logs of the RandomWalk contract and prints those of the selected token.\n\n" +
			"Environment variables:\n  RPC_URL  Ethereum RPC endpoint (required)",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			rwalkAddr := common.HexToAddress(args[0])
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
			for fromBlock := int64(0); fromBlock <= latestBnum; fromBlock += scanTransfersMaxBlocks {
				filter.FromBlock = big.NewInt(fromBlock)
				filter.ToBlock = big.NewInt(fromBlock + scanTransfersMaxBlocks)
				logs, err := eclient.FilterLogs(ctx, filter)
				if err != nil {
					return fmt.Errorf("error querying events: %w", err)
				}
				for i := range logs {
					lg := &logs[i]
					if lg.Removed {
						continue
					}
					from := common.BytesToAddress(lg.Topics[1][12:]).String()
					to := common.BytesToAddress(lg.Topics[2][12:]).String()
					tok := lg.Topics[3].Big().Int64()
					if tok != tokenID {
						continue
					}
					fmt.Printf(
						"Block %v: Transfer %v -> %v (tok %v) tx %v\n",
						lg.BlockNumber, from, to, tok, lg.TxHash.String(),
					)
				}
			}
			return nil
		},
	}
	c.Flags().Int64Var(&tokenID, "token-id", scanTransfersDefaultToken,
		"token ID whose transfers are printed (legacy script hardcoded 3601)")
	return c
}

func init() { register(newScanTransfersCmd()) }
