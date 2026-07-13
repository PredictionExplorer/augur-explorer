package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/archive"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

type archiveNodeFillRPC interface {
	archive.NodeClient
	BlockNumber(context.Context) (uint64, error)
	Close()
}

type archiveNodeFillStorage struct {
	db           *sql.DB
	addressStore archive.AddressStore
	close        func()
}

type archiveNodeFillDeps struct {
	getenv          func(string) string
	resolveProjects func(string) ([]string, error)
	openStorage     func(context.Context, string) (archiveNodeFillStorage, error)
	requireSchema   func(context.Context, *sql.DB) error
	dialRPC         func(context.Context, string) (archiveNodeFillRPC, error)
	newRepository   func(*sql.DB) archive.NodeFillRepository
	runProject      func(context.Context, *archive.NodeFiller, string, archive.NodeFillOptions) (archive.FillStats, error)
}

func defaultArchiveNodeFillDeps() archiveNodeFillDeps {
	return archiveNodeFillDeps{
		getenv:          os.Getenv,
		resolveProjects: archive.ResolveProjects,
		openStorage:     openArchiveNodeFillStorage,
		requireSchema:   archive.RequireArchLogIndex,
		dialRPC: func(ctx context.Context, rpcURL string) (archiveNodeFillRPC, error) {
			return ethclient.DialContext(ctx, rpcURL)
		},
		newRepository: func(db *sql.DB) archive.NodeFillRepository {
			return &archive.SQLNodeFillRepository{DB: db}
		},
		runProject: func(
			ctx context.Context,
			filler *archive.NodeFiller,
			project string,
			options archive.NodeFillOptions,
		) (archive.FillStats, error) {
			return filler.RunProject(ctx, project, options)
		},
	}
}

func openArchiveNodeFillStorage(ctx context.Context, dbConn string) (archiveNodeFillStorage, error) {
	poolConfig, err := pgxpool.ParseConfig(dbConn)
	if err != nil {
		return archiveNodeFillStorage{}, fmt.Errorf("db pool config: %w", err)
	}
	poolConfig.MaxConns = 4
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return archiveNodeFillStorage{}, fmt.Errorf("db pool connect: %w", err)
	}
	sqlDB := stdlib.OpenDBFromPool(pool)
	st := store.NewFromPool(pool)
	return archiveNodeFillStorage{
		db:           sqlDB,
		addressStore: st,
		close: func() {
			_ = sqlDB.Close()
			st.Close()
		},
	}, nil
}

// newArchiveNodeFillCmd builds `opsctl archive node-fill`, the replacement
// for the standalone arch_node_fill tool.
func newArchiveNodeFillCmd() *cobra.Command {
	return newArchiveNodeFillCmdWithDeps(defaultArchiveNodeFillDeps())
}

func newArchiveNodeFillCmdWithDeps(deps archiveNodeFillDeps) *cobra.Command {
	var (
		dbConn      string
		projectType string
		fromBlock   uint64
		startBlock  uint64
		toBlock     uint64
		batchBlocks uint64
		dryRun      bool
	)
	cmd := &cobra.Command{
		Use:   "node-fill",
		Short: "Backfill arch_evtlog / arch_tx / arch_block from an Ethereum node via FilterLogs",
		Long: `Scans the chain for the project contracts' logs and inserts only rows
missing from arch_evtlog (keyed by tx_hash + log_index), fetching the matching
transactions and blocks from the node as needed.

Requires the RPC_URL environment variable (Ethereum/Arbitrum JSON-RPC
endpoint) and arch_evtlog with PRIMARY KEY (tx_hash, log_index) — created by
the goose migrations under db/migrations.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			flagFrom := fromBlock
			if startBlock > 0 {
				if fromBlock > 0 && fromBlock != startBlock {
					return fmt.Errorf("--start-block (%d) and --from (%d) disagree; pass only one", startBlock, fromBlock)
				}
				flagFrom = startBlock
			}

			rpcURL := deps.getenv("RPC_URL")
			if rpcURL == "" {
				return errors.New("RPC_URL environment variable is required")
			}
			projects, err := deps.resolveProjects(projectType)
			if err != nil {
				return err
			}
			ctx := cmd.Context()

			// One pgx pool backs both store.Store and the SQL-heavy archive
			// repository. Closing the SQL adapter does not close the pool;
			// store owns and closes it after the adapter is closed.
			storage, err := deps.openStorage(ctx, dbConn)
			if err != nil {
				return err
			}
			if storage.close != nil {
				defer storage.close()
			}

			if err := deps.requireSchema(ctx, storage.db); err != nil {
				return fmt.Errorf("schema check: %w", err)
			}
			client, err := deps.dialRPC(ctx, rpcURL)
			if err != nil {
				return fmt.Errorf("rpc connect: %w", err)
			}
			defer client.Close()

			head, err := client.BlockNumber(ctx)
			if err != nil {
				return fmt.Errorf("chain head: %w", err)
			}
			endBlock := head
			if toBlock > 0 {
				endBlock = toBlock
			}

			logger := log.New(cmd.ErrOrStderr(), "", log.LstdFlags)
			logger.Printf("RPC: %s", redactRPCURL(rpcURL))
			logger.Printf("Chain head: %d, scanning through block %d", head, endBlock)
			if dryRun {
				logger.Println("DRY RUN — no rows will be inserted")
			}

			filler := &archive.NodeFiller{
				Repository:   deps.newRepository(storage.db),
				AddressStore: storage.addressStore,
				Client:       client,
				Logger:       logger,
			}
			var total archive.FillStats
			for _, project := range projects {
				logger.Println("")
				logger.Printf("========== project: %s ==========", project)
				stats, err := deps.runProject(ctx, filler, project, archive.NodeFillOptions{
					FromBlock: flagFrom,
					EndBlock:  endBlock,
					BatchSize: batchBlocks,
					DryRun:    dryRun,
				})
				if err != nil {
					return err
				}
				archive.LogFillStats(logger, project, stats)
				total.Merge(stats)
			}

			logger.Println("")
			logger.Println("========== TOTAL ==========")
			archive.LogFillStats(logger, "all", total)
			if total.RPCErrors > 0 || total.DBErrors > 0 {
				return fmt.Errorf(
					"archive node-fill completed with unresolved errors (rpc=%d db=%d)",
					total.RPCErrors,
					total.DBErrors,
				)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&dbConn, "db", "", "PostgreSQL connection string")
	cmd.Flags().StringVar(&projectType, "project", "", "randomwalk | cosmicgame | both")
	cmd.Flags().Uint64Var(&fromBlock, "from", 0, "Start block (0 = auto: min contract address block_num, else min evt_log block)")
	cmd.Flags().Uint64Var(&startBlock, "start-block", 0, "Start block (same as --from; overrides --from when both are set)")
	cmd.Flags().Uint64Var(&toBlock, "to", 0, "End block inclusive (0 = chain head)")
	cmd.Flags().Uint64Var(&batchBlocks, "batch", archive.DefaultFilterBatchBlocks, "FilterLogs block range size")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Scan and report only; do not insert")
	_ = cmd.MarkFlagRequired("db")
	_ = cmd.MarkFlagRequired("project")
	return cmd
}
