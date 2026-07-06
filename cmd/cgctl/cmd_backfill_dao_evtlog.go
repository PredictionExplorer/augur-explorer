package main

import (
	"context"
	"fmt"
	"log"
	"os"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/cobra"

	etlcommon "github.com/PredictionExplorer/augur-explorer/internal/etl"
	dbs "github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// backfillDefaultFromBlock is the first block scanned when --from-block is not
// given: the block range the ETL had already passed before cosmic_dao_addr was
// added to FilterLogs.
const backfillDefaultFromBlock uint64 = 455_767_500

func init() {
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
			return runBackfillDaoEvtlog(cmd.Context(), fromBlock, toBlock)
		},
	}
	c.Flags().Uint64Var(&fromBlock, "from-block", backfillDefaultFromBlock, "first block to scan (inclusive)")
	c.Flags().Uint64Var(&toBlock, "to-block", 0, "last block to scan (inclusive); 0 = ETL last processed block")
	register(c)
}

func runBackfillDaoEvtlog(ctx context.Context, fromBlock, toBlock uint64) error {
	logger := log.New(os.Stdout, "", log.Ltime)
	errLogger := log.New(os.Stderr, "ERROR: ", log.Ltime)

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		return fmt.Errorf("RPC_URL must be set")
	}

	rpcclient, err := rpc.DialContext(ctx, rpcURL)
	if err != nil {
		return fmt.Errorf("rpc connect: %w", err)
	}
	client := ethclient.NewClient(rpcclient)
	logger.Printf("RPC: %s", rpcURL)

	storage := dbs.Connect_to_storage(logger)
	if storage == nil {
		return fmt.Errorf("failed to connect to database")
	}
	storage.Db_set_schema_name("public")

	sw := &cgstore.SQLStorageWrapper{S: storage}
	contracts := sw.Get_cosmic_game_contract_addrs()
	daoAddr := ethcommon.HexToAddress(contracts.CosmicDaoAddr)
	if (daoAddr == ethcommon.Address{}) {
		return fmt.Errorf("cosmic_dao_addr is zero in cg_contracts")
	}
	logger.Printf("cosmic_dao: %s", daoAddr.Hex())

	endBlock := toBlock
	if endBlock == 0 {
		status := sw.Get_cosmic_game_processing_status()
		endBlock = uint64(status.LastBlockNum)
		if endBlock == 0 {
			last, err := storage.Get_last_block_num()
			if err != nil {
				return fmt.Errorf("Get_last_block_num: %w", err)
			}
			endBlock = uint64(last)
		}
	}
	if endBlock < fromBlock {
		return fmt.Errorf("to-block %d < from-block %d", endBlock, fromBlock)
	}

	etlCtx := &etlcommon.ETLContext{
		Storage:   storage,
		EthClient: client,
		RpcClient: rpcclient,
		Info:      logger,
		Error:     errLogger,
	}

	logger.Printf("Backfilling cosmic_dao evt_log blocks %d .. %d", fromBlock, endBlock)
	st, err := etlcommon.BackfillContractEvtLogs(
		etlCtx,
		client,
		[]ethcommon.Address{daoAddr},
		fromBlock,
		endBlock,
		100_000,
	)
	if err != nil {
		return fmt.Errorf("backfill failed: %w", err)
	}
	logger.Printf("done: logs_seen=%d inserted=%d skipped=%d", st.LogsSeen, st.Inserted, st.Skipped)

	count, err := storage.Count_evt_log_for_contract(contracts.CosmicDaoAddr)
	if err != nil {
		return fmt.Errorf("Count_evt_log_for_contract: %w", err)
	}
	logger.Printf("cosmic_dao evt_log rows now: %d", count)
	if st.Inserted == 0 && count == 0 {
		return fmt.Errorf("no cosmic_dao evt_log rows found or inserted")
	}
	return nil
}
