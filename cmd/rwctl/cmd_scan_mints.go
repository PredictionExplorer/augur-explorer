package main

import (
	"context"
	"fmt"
	"io"
	"time"

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

// mintChecker verifies minted tokens against the database, retrying
// transient database errors.
type mintChecker struct {
	repo  *rwstore.Repo
	out   io.Writer
	sleep func(time.Duration) // injected in tests
}

// checkMintLog verifies that the minted token from a MintEvent log exists in
// the database, retrying on transient database errors.
func (m *mintChecker) checkMintLog(ctx context.Context, lg *types.Log) error {
	tokenID := lg.Topics[1].Big().Int64()
	for {
		exists, err := m.repo.TokenMinted(ctx, tokenID)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			fmt.Fprintf(m.out, "Error accessing database: %v\n", err)
			m.sleep(1 * time.Second)
			continue
		}
		if !exists {
			fmt.Fprintf(m.out, "Token %v DOES NOT exist in the DB\n", tokenID)
		}
		return nil
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
			repo, _, err := connectRWStorage(newInfoLogger())
			if err != nil {
				return err
			}

			ctx := cmd.Context()
			checker := &mintChecker{repo: repo, out: cmd.OutOrStdout(), sleep: time.Sleep}
			return scanLogsByRange(ctx, eclient, rwalkAddr, mintEventTopic,
				scanMintsStartBlock, scanMintsMaxBlocks,
				func(from, to int64) { fmt.Fprintf(cmd.OutOrStdout(), "From %v , to %v\n", from, to) },
				func(lg *types.Log) error { return checker.checkMintLog(ctx, lg) },
			)
		},
	}
}

func init() { register(newScanMintsCmd()) }
