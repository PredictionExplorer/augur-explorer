package main

import (
	"errors"
	"fmt"
	"log/slog"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// backfillDefaultFromBlock is the first block scanned when --from-block is not
// given: the block range the ETL had already passed before cosmic_dao_addr was
// added to FilterLogs.
const backfillDefaultFromBlock uint64 = 455_767_500

// newBackfillDaoEvtlogCmd builds the backfill-dao-evtlog subcommand.
func newBackfillDaoEvtlogCmd() *cobra.Command {
	var fromBlock, toBlock uint64
	c := &cobra.Command{
		Use:   "backfill-dao-evtlog",
		Short: "Backfill missing cosmic_dao evt_log rows",
		Long: `Insert missing cosmic_dao evt_log rows for blocks the ETL already passed
before cosmic_dao_addr was added to FilterLogs.

Environment:
  RPC_URL  Ethereum RPC endpoint (required)
  PGSQL_*  PostgreSQL connection (PGSQL_HOST, PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD)`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBackfillDaoEvtlog(cmd, fromBlock, toBlock)
		},
	}
	c.Flags().Uint64Var(&fromBlock, "from-block", backfillDefaultFromBlock, "first block to scan (inclusive)")
	c.Flags().Uint64Var(&toBlock, "to-block", 0, "last block to scan (inclusive); 0 = ETL last processed block")
	return c
}

func init() { register(newBackfillDaoEvtlogCmd()) }

// blockNumFromWatermark converts a database block watermark to a block
// number, rejecting negative values instead of letting them wrap into an
// astronomically large scan end.
func blockNumFromWatermark(source string, n int64) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("%s watermark is negative: %d", source, n)
	}
	return uint64(n), nil
}

func runBackfillDaoEvtlog(cmd *cobra.Command, fromBlock, toBlock uint64) error {
	ctx := cmd.Context()
	logger := slog.New(slog.NewTextHandler(cmd.OutOrStdout(), nil))

	rpcURL, err := rpcURLFromEnv()
	if err != nil {
		return err
	}

	rpcclient, err := rpc.DialContext(ctx, rpcURL)
	if err != nil {
		return fmt.Errorf("rpc connect: %w", err)
	}
	client := ethclient.NewClient(rpcclient)
	logger.Info("connected", "rpc", rpcURL)

	st, err := store.New(ctx, store.ConfigFromEnv())
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer st.Close()

	repo := cgstore.NewRepo(st)
	contracts, err := repo.ContractAddrs(ctx)
	if err != nil {
		return fmt.Errorf("reading contract addresses: %w", err)
	}
	daoAddr := ethcommon.HexToAddress(contracts.CosmicDaoAddr)
	if (daoAddr == ethcommon.Address{}) {
		return errors.New("cosmic_dao_addr is zero in cg_contracts")
	}
	logger.Info("resolved contract", "cosmic_dao", daoAddr.Hex())

	endBlock := toBlock
	if endBlock == 0 {
		status, err := repo.ProcessingStatus(ctx)
		if err != nil {
			return fmt.Errorf("reading processing status: %w", err)
		}
		endBlock, err = blockNumFromWatermark("processing status", status.LastBlockNum)
		if err != nil {
			return err
		}
		if endBlock == 0 {
			last, err := st.LastBlockNum(ctx)
			if err != nil {
				return fmt.Errorf("reading last block watermark: %w", err)
			}
			endBlock, err = blockNumFromWatermark("last_block", last)
			if err != nil {
				return err
			}
		}
	}
	if endBlock < fromBlock {
		return fmt.Errorf("to-block %d < from-block %d", endBlock, fromBlock)
	}

	engine, err := indexer.New(indexer.Config{Store: st, Client: client, Logger: logger})
	if err != nil {
		return fmt.Errorf("building indexer engine: %w", err)
	}

	logger.Info("backfilling cosmic_dao evt_log", "from_block", fromBlock, "to_block", endBlock)
	stats, err := engine.BackfillContractEvtLogs(ctx, []ethcommon.Address{daoAddr}, fromBlock, endBlock, 100_000)
	if err != nil {
		return fmt.Errorf("backfill failed: %w", err)
	}
	logger.Info("backfill done",
		"logs_seen", stats.LogsSeen, "inserted", stats.Inserted, "skipped", stats.Skipped)

	count, err := st.CountEvtLogsForContract(ctx, contracts.CosmicDaoAddr)
	if err != nil {
		return fmt.Errorf("counting cosmic_dao evt_log rows: %w", err)
	}
	logger.Info("cosmic_dao evt_log total", "rows", count)
	if stats.Inserted == 0 && count == 0 {
		return errors.New("no cosmic_dao evt_log rows found or inserted")
	}
	return nil
}
