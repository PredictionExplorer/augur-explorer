package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/txcollector"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
	_ "github.com/lib/pq" // postgres driver
	"github.com/spf13/cobra"
)

type txCollectorVerifyDeps struct {
	loadConfig   func(string) (*toolutil.CollectorConfig, error)
	postgresConn func() (string, error)
	openDB       func(string, string) (*sql.DB, error)
	loadRows     func(context.Context, *sql.DB, []string, uint64) ([]txcollector.EventRow, error)
	verify       func(context.Context, txcollector.VerifyConfig) (txcollector.VerifyStats, error)
}

func defaultTxCollectorVerifyDeps() txCollectorVerifyDeps {
	return txCollectorVerifyDeps{
		loadConfig:   toolutil.LoadCollectorConfig,
		postgresConn: toolutil.PostgresConnStringFromEnv,
		openDB:       sql.Open,
		loadRows:     txcollector.LoadEventRows,
		verify:       txcollector.Verify,
	}
}

// newTxCollectorVerifyCmd builds `opsctl tx-collector verify`, the
// replacement for the standalone transaction-collector-verify tool.
func newTxCollectorVerifyCmd() *cobra.Command {
	return newTxCollectorVerifyCmdWithDeps(defaultTxCollectorVerifyDeps())
}

func newTxCollectorVerifyCmdWithDeps(deps txCollectorVerifyDeps) *cobra.Command {
	var (
		configPath string
		dbConn     string
		startBlock uint64
		maxReport  int
	)
	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Check RLP backup blobs against evt_log in PostgreSQL",
		Long: `For each evt_log row (scoped by config contract addresses and start_block),
verifies that:

 1. the tx hash matches the tx RLP blob (receipt RLP does not store tx hash)
 2. log_index matches the backed-up log (v1 format) or is matched via
    log_rlp (legacy blobs)
 3. the RLP-encoded log bytes match evt_log.log_rlp

Without --db the PostgreSQL connection is built from the PGSQL_HOST,
PGSQL_USERNAME, PGSQL_DATABASE and PGSQL_PASSWORD environment variables.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.New(cmd.ErrOrStderr(), "", log.LstdFlags)
			cfg, err := deps.loadConfig(configPath)
			if err != nil {
				return fmt.Errorf("config: %w", err)
			}
			contractAddrs, err := cfg.ResolveContractAddresses()
			if err != nil {
				return fmt.Errorf("contracts: %w", err)
			}

			fromBlock := cfg.StartBlock
			if startBlock > 0 {
				fromBlock = startBlock
			}

			conn := dbConn
			if conn == "" {
				conn, err = deps.postgresConn()
				if err != nil {
					return fmt.Errorf("db: %w", err)
				}
			}

			db, err := deps.openDB("postgres", conn)
			if err != nil {
				return fmt.Errorf("connect: %w", err)
			}
			defer func() { _ = db.Close() }()
			db.SetMaxOpenConns(4)

			logger.Printf("Output dir: %s", cfg.OutputDir)
			logger.Printf("Contracts (%d): %v", len(contractAddrs), contractAddrs)
			logger.Printf("evt_log block_num >= %d", fromBlock)

			rows, err := deps.loadRows(cmd.Context(), db, contractAddrs, fromBlock)
			if err != nil {
				return fmt.Errorf("load evt_log: %w", err)
			}
			logger.Printf("Loaded %d evt_log rows", len(rows))

			_, err = deps.verify(cmd.Context(), txcollector.VerifyConfig{
				OutputDir: cfg.OutputDir,
				Rows:      rows,
				MaxReport: maxReport,
				Logger:    logger,
			})
			return err
		},
	}
	cmd.Flags().StringVar(&configPath, "config", "", "transaction-collector JSON config (output_dir, contracts, start_block)")
	cmd.Flags().StringVar(&dbConn, "db", "", "PostgreSQL URL (default: PGSQL_* from environment)")
	cmd.Flags().Uint64Var(&startBlock, "start-block", 0, "Override config start_block for evt_log filter")
	cmd.Flags().IntVar(&maxReport, "max-report", 50, "Max detailed mismatch lines per category (0 = unlimited)")
	_ = cmd.MarkFlagRequired("config")
	return cmd
}
