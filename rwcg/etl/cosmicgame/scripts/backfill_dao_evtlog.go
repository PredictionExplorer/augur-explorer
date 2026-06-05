// backfill_dao_evtlog inserts missing cosmic_dao evt_log rows for blocks the ETL
// already passed before cosmic_dao_addr was added to FilterLogs.
//
// Requires env: RPC_URL, PGSQL_HOST, PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD
//
// Usage:
//   source ~/configs/cg-prod.env
//   ./backfill_dao_evtlog
//   ./backfill-dao-evtlog.sh
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	cgdb "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/cosmicgame"
	etlcommon "github.com/PredictionExplorer/augur-explorer/rwcg/etl/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const defaultFromBlock uint64 = 455_767_500

func main() {
	fromBlock := flag.Uint64("from-block", defaultFromBlock, "First block to scan (inclusive)")
	toBlock := flag.Uint64("to-block", 0, "Last block to scan (inclusive); 0 = ETL last processed block")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ltime)
	errLogger := log.New(os.Stderr, "ERROR: ", log.Ltime)

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		errLogger.Fatal("RPC_URL must be set")
	}

	rpcclient, err := rpc.DialContext(context.Background(), rpcURL)
	if err != nil {
		errLogger.Fatalf("rpc connect: %v", err)
	}
	client := ethclient.NewClient(rpcclient)
	logger.Printf("RPC: %s", rpcURL)

	storage := dbs.Connect_to_storage(logger)
	if storage == nil {
		errLogger.Fatal("failed to connect to database")
	}
	storage.Db_set_schema_name("public")

	sw := &cgdb.SQLStorageWrapper{S: storage}
	contracts := sw.Get_cosmic_game_contract_addrs()
	daoAddr := ethcommon.HexToAddress(contracts.CosmicDaoAddr)
	if (daoAddr == ethcommon.Address{}) {
		errLogger.Fatal("cosmic_dao_addr is zero in cg_contracts")
	}
	logger.Printf("cosmic_dao: %s", daoAddr.Hex())

	endBlock := *toBlock
	if endBlock == 0 {
		status := sw.Get_cosmic_game_processing_status()
		endBlock = uint64(status.LastBlockNum)
		if endBlock == 0 {
			last, err := storage.Get_last_block_num()
			if err != nil {
				errLogger.Fatalf("Get_last_block_num: %v", err)
			}
			endBlock = uint64(last)
		}
	}
	if endBlock < *fromBlock {
		errLogger.Fatalf("to-block %d < from-block %d", endBlock, *fromBlock)
	}

	ctx := &etlcommon.ETLContext{
		Storage:   storage,
		EthClient: client,
		RpcClient: rpcclient,
		Info:      logger,
		Error:     errLogger,
	}

	logger.Printf("Backfilling cosmic_dao evt_log blocks %d .. %d", *fromBlock, endBlock)
	st, err := etlcommon.BackfillContractEvtLogs(
		ctx,
		client,
		[]ethcommon.Address{daoAddr},
		*fromBlock,
		endBlock,
		100_000,
	)
	if err != nil {
		errLogger.Fatalf("backfill failed: %v", err)
	}
	logger.Printf("done: logs_seen=%d inserted=%d skipped=%d", st.LogsSeen, st.Inserted, st.Skipped)

	count, err := storage.Count_evt_log_for_contract(contracts.CosmicDaoAddr)
	if err != nil {
		errLogger.Fatalf("Count_evt_log_for_contract: %v", err)
	}
	logger.Printf("cosmic_dao evt_log rows now: %d", count)
	if st.Inserted == 0 && count == 0 {
		fmt.Fprintln(os.Stderr, "warning: no cosmic_dao evt_log rows found or inserted")
		os.Exit(1)
	}
}
