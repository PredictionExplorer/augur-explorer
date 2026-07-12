package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

			return scanLogsByRange(cmd.Context(), eclient, rwalkAddr, transferEventTopic,
				0, scanTransfersMaxBlocks, nil,
				func(lg *types.Log) error {
					from := common.BytesToAddress(lg.Topics[1][12:]).String()
					to := common.BytesToAddress(lg.Topics[2][12:]).String()
					tok := lg.Topics[3].Big().Int64()
					if tok != tokenID {
						return nil
					}
					fmt.Fprintf(cmd.OutOrStdout(),
						"Block %v: Transfer %v -> %v (tok %v) tx %v\n",
						lg.BlockNumber, from, to, tok, lg.TxHash.String(),
					)
					return nil
				},
			)
		},
	}
	c.Flags().Int64Var(&tokenID, "token-id", scanTransfersDefaultToken,
		"token ID whose transfers are printed (legacy script hardcoded 3601)")
	return c
}

func init() { register(newScanTransfersCmd()) }
