package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq" // postgres driver
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/cstscan"
)

const (
	// cstAucLenDefaultContract is the CosmicSignatureGame proxy address.
	cstAucLenDefaultContract = "0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2"
	// cstAucLenDefaultFromBlock is the contract deployment block.
	cstAucLenDefaultFromBlock = uint64(455767589)
	// cstAucLenTopic0Hex is the topic0 of
	// ISystemEventsV2.sol:CstDutchAuctionDurationChangeDivisorChanged(uint256).
	// The new value (uint256) is non-indexed, carried in log.Data.
	cstAucLenTopic0Hex = "0xacbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f"
	// Retry behavior retained from the legacy standalone scanner.
	cstAucLenMinBatch   = uint64(1_000)
	cstAucLenRetryDelay = 3 * time.Second
)

type cstAuctionLenRPC interface {
	cstscan.Client
	Close()
}

type cstAuctionLenDeps struct {
	getenv       func(string) string
	dialRPC      func(context.Context, string) (cstAuctionLenRPC, error)
	openDB       func(string, string) (*sql.DB, error)
	newKeySource func(*sql.DB) cstscan.KeySource
	scan         func(context.Context, cstscan.Config, cstscan.Options) (cstscan.Stats, error)
}

func defaultCstAuctionLenDeps() cstAuctionLenDeps {
	return cstAuctionLenDeps{
		getenv: os.Getenv,
		dialRPC: func(ctx context.Context, rpcURL string) (cstAuctionLenRPC, error) {
			return ethclient.DialContext(ctx, rpcURL)
		},
		openDB: sql.Open,
		newKeySource: func(db *sql.DB) cstscan.KeySource {
			return cstscan.PostgresKeySource{DB: db}
		},
		scan: cstscan.Scan,
	}
}

// newScanCstAuctionLenCmd builds `opsctl scan cst-auction-len`, the
// replacement for the standalone scan_cst_auclen_chg_div tool.
func newScanCstAuctionLenCmd() *cobra.Command {
	return newScanCstAuctionLenCmdWithDeps(defaultCstAuctionLenDeps())
}

func newScanCstAuctionLenCmdWithDeps(deps cstAuctionLenDeps) *cobra.Command {
	var (
		contractStr string
		fromBlock   uint64
		toBlock     uint64
		batch       uint64
		dbConn      string
	)
	cmd := &cobra.Command{
		Use:   "cst-auction-len",
		Short: "Scan the chain for CstDutchAuctionDurationChangeDivisorChanged events",
		Long: `Scans the chain for CstDutchAuctionDurationChangeDivisorChanged events on
the CosmicSignatureGame proxy and, when --db is set, reports which on-chain
occurrences are missing from the cg_adm_cst_auclen_chg_div table.

Use this after creating the cg_adm_cst_auclen_chg_div table to find admin
events that may have been emitted while the table did not exist (the ETL
could not insert them then).

Requires the RPC_URL environment variable.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if !ethcommon.IsHexAddress(contractStr) {
				return fmt.Errorf("invalid --contract address %q", contractStr)
			}
			rpcURL := deps.getenv("RPC_URL")
			if rpcURL == "" {
				return errors.New("RPC_URL environment variable is required")
			}

			client, err := deps.dialRPC(cmd.Context(), rpcURL)
			if err != nil {
				return fmt.Errorf("RPC dial: %w", err)
			}
			defer client.Close()

			var (
				keySource cstscan.KeySource
				db        *sql.DB
			)
			if dbConn != "" {
				db, err = deps.openDB("postgres", dbConn)
				if err != nil {
					return fmt.Errorf("db open: %w", err)
				}
				defer func() { _ = db.Close() }()
				keySource = deps.newKeySource(db)
			}

			logger := log.New(cmd.ErrOrStderr(), "", log.LstdFlags)
			_, err = deps.scan(cmd.Context(), cstscan.Config{
				Client:    client,
				Contract:  ethcommon.HexToAddress(contractStr),
				Topic0:    ethcommon.HexToHash(cstAucLenTopic0Hex),
				KeySource: keySource,
				Output:    cmd.OutOrStdout(),
				Logger:    logger,
			}, cstscan.Options{
				FromBlock:    fromBlock,
				ToBlock:      toBlock,
				InitialBatch: batch,
				MinBatch:     min(batch, cstAucLenMinBatch),
				RetryDelay:   cstAucLenRetryDelay,
			})
			return err
		},
	}
	cmd.Flags().StringVar(&contractStr, "contract", cstAucLenDefaultContract, "CosmicSignatureGame proxy address")
	cmd.Flags().Uint64Var(&fromBlock, "from-block", cstAucLenDefaultFromBlock, "start block")
	cmd.Flags().Uint64Var(&toBlock, "to-block", 0, "end block (0 = latest)")
	cmd.Flags().Uint64Var(&batch, "batch", scanDefaultFilterBatchBlocks, "FilterLogs block range size")
	cmd.Flags().StringVar(&dbConn, "db", "", "optional postgres conn string; if set, cross-check cg_adm_cst_auclen_chg_div")
	return cmd
}
