package main

import (
	"context"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// verifyTransfersMaxBlocks is the block-range step used by
// verify-erc20-transfers (same as the legacy verify_erc20_transfers script).
const verifyTransfersMaxBlocks = int64(1024)

// checkTransferLog verifies that a chain Transfer log has a matching record
// in the rw_transfer table and reports missing ones on out.
func checkTransferLog(ctx context.Context, repo *rwstore.Repo, out io.Writer, lg *types.Log) error {
	from := common.BytesToAddress(lg.Topics[1][12:]).String()
	to := common.BytesToAddress(lg.Topics[2][12:]).String()
	tokenID := lg.Topics[3].Big().Int64()
	transfers, err := repo.TokenTransfersByTxHash(ctx, lg.TxHash.String())
	if err != nil {
		return fmt.Errorf("transfers for tx %v: %w", lg.TxHash.String(), err)
	}
	found := false
	for i := range transfers {
		item := &transfers[i]
		if from == item.From && to == item.To && tokenID == item.TokenId {
			found = true
		}
	}
	if !found {
		fmt.Fprintf(out,
			"Block %v: Transfer %v -> %v (tok %v) tx %v, transfer missing\n",
			lg.BlockNumber, from, to, tokenID, lg.TxHash.String(),
		)
	}
	return nil
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
			ctx := cmd.Context()
			repo, _, err := connectRWStorage(newInfoLogger())
			if err != nil {
				return err
			}
			caddrs, err := repo.ContractAddrs(ctx)
			if err != nil {
				return fmt.Errorf("resolving contract addresses: %w", err)
			}
			rwalkAddr := common.HexToAddress(caddrs.RandomWalk)

			eclient, err := dialEthClient()
			if err != nil {
				return err
			}
			return scanLogsByRange(ctx, eclient, rwalkAddr, transferEventTopic,
				0, verifyTransfersMaxBlocks, nil,
				func(lg *types.Log) error {
					return checkTransferLog(ctx, repo, cmd.OutOrStdout(), lg)
				},
			)
		},
	}
}

func init() { register(newVerifyTransfersCmd()) }
