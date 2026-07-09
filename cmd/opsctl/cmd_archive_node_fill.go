package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	etlcommon "github.com/PredictionExplorer/augur-explorer/internal/etl"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
	"github.com/spf13/cobra"
)

// defaultFilterBatchBlocks is the default FilterLogs block-range size used by
// the chain-scanning subcommands (archive node-fill, tx-collector run).
const defaultFilterBatchBlocks = 100_000

// fillStats accumulates counters for one archive node-fill run.
type fillStats struct {
	BlocksScanned uint64
	LogsFromNode  int64
	LogsSkipped   int64 // already in arch_evtlog
	LogsInserted  int64
	TxInserted    int64
	TxSkipped     int64
	BlockInserted int64
	BlockSkipped  int64
	RPCErrors     int64
	DBErrors      int64
}

// newArchiveNodeFillCmd builds `opsctl archive node-fill`, the replacement
// for the standalone arch_node_fill tool.
func newArchiveNodeFillCmd() *cobra.Command {
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

			rpcURL := os.Getenv("RPC_URL")
			if rpcURL == "" {
				return errors.New("RPC_URL environment variable is required")
			}

			projects, err := resolveProjects(projectType)
			if err != nil {
				return err
			}

			ctx := cmd.Context()

			db, err := sql.Open("postgres", dbConn)
			if err != nil {
				return fmt.Errorf("db connect: %w", err)
			}
			defer db.Close()
			db.SetMaxOpenConns(4)

			// The address lookup/create cache runs on the pgx-native Store
			// (same DSN); the tool-local archive statements stay on db.
			pool, err := pgxpool.New(ctx, dbConn)
			if err != nil {
				return fmt.Errorf("db pool connect: %w", err)
			}
			st := store.NewFromPool(pool)
			defer st.Close()

			if err := requireArchLogIndexPK(db); err != nil {
				return fmt.Errorf("schema check: %w", err)
			}

			eclient, err := ethclient.Dial(rpcURL)
			if err != nil {
				return fmt.Errorf("rpc connect: %w", err)
			}

			head, err := etlcommon.GetCurrentBlockNumber(ctx, eclient)
			if err != nil {
				return fmt.Errorf("chain head: %w", err)
			}
			endBlock := head
			if toBlock > 0 {
				endBlock = toBlock
			}

			log.Printf("RPC: %s", rpcURL)
			log.Printf("Chain head: %d, scanning through block %d", head, endBlock)
			if dryRun {
				log.Println("DRY RUN — no rows will be inserted")
			}

			var total fillStats
			for _, p := range projects {
				log.Printf("")
				log.Printf("========== project: %s ==========", p)
				stats, err := runNodeFillProject(ctx, db, st, eclient, p, flagFrom, endBlock, batchBlocks, dryRun)
				if err != nil {
					return err
				}
				printFillStats(p, &stats)
				mergeFillStats(&total, &stats)
			}

			log.Printf("")
			log.Printf("========== TOTAL ==========")
			printFillStats("all", &total)
			return nil
		},
	}
	cmd.Flags().StringVar(&dbConn, "db", "", "PostgreSQL connection string")
	cmd.Flags().StringVar(&projectType, "project", "", "randomwalk | cosmicgame | both")
	cmd.Flags().Uint64Var(&fromBlock, "from", 0, "Start block (0 = auto: min contract address block_num, else min evt_log block)")
	cmd.Flags().Uint64Var(&startBlock, "start-block", 0, "Start block (same as --from; overrides --from when both are set)")
	cmd.Flags().Uint64Var(&toBlock, "to", 0, "End block inclusive (0 = chain head)")
	cmd.Flags().Uint64Var(&batchBlocks, "batch", defaultFilterBatchBlocks, "FilterLogs block range size")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Scan and report only; do not insert")
	_ = cmd.MarkFlagRequired("db")
	_ = cmd.MarkFlagRequired("project")
	return cmd
}

func init() { archiveCmd.AddCommand(newArchiveNodeFillCmd()) }

// requireArchLogIndexPK verifies that arch_evtlog carries the log_index
// column of the (tx_hash, log_index) natural key.
func requireArchLogIndexPK(db *sql.DB) error {
	var n int
	err := db.QueryRow(`
		SELECT COUNT(*) FROM information_schema.columns
		WHERE table_name = 'arch_evtlog' AND column_name = 'log_index'
	`).Scan(&n)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("arch_evtlog.log_index missing — create the archive tables from db/layer1/archive_tables.sql first")
	}
	return nil
}

// resolveNodeFillFromBlock picks the scan start block: the flag value when
// given, otherwise the earliest block seen for the project contracts.
func resolveNodeFillFromBlock(db *sql.DB, aids []int64, addrs []string, flagFrom uint64) (uint64, error) {
	if flagFrom > 0 {
		log.Printf("Start block: %d (from --start-block / --from)", flagFrom)
		return flagFrom, nil
	}
	var fromAddr, fromEvt sql.NullInt64
	_ = db.QueryRow(`
		SELECT MIN(block_num) FROM address WHERE addr = ANY($1)
	`, pq.Array(addrs)).Scan(&fromAddr)
	_ = db.QueryRow(`
		SELECT MIN(block_num) FROM evt_log WHERE contract_aid = ANY($1)
	`, pq.Array(aids)).Scan(&fromEvt)

	candidates := []int64{}
	if fromAddr.Valid && fromAddr.Int64 > 0 {
		candidates = append(candidates, fromAddr.Int64)
	}
	if fromEvt.Valid && fromEvt.Int64 > 0 {
		candidates = append(candidates, fromEvt.Int64)
	}
	if len(candidates) == 0 {
		return 0, errors.New("could not auto-detect start block; pass --from <deployment_block>")
	}
	minBlock := candidates[0]
	for _, c := range candidates[1:] {
		if c < minBlock {
			minBlock = c
		}
	}
	log.Printf("Auto start block: %d (from address/evt_log metadata)", minBlock)
	return uint64(minBlock), nil
}

// runNodeFillProject scans the chain for one project and inserts the missing
// archive rows. Per-log RPC/DB problems are counted in the stats; only setup
// failures abort the run.
func runNodeFillProject(ctx context.Context, db *sql.DB, addrStore *store.Store, client *ethclient.Client, project string, flagFrom, endBlock, batchSize uint64, dryRun bool) (fillStats, error) {
	var st fillStats

	aids, addrs, err := projectContracts(db, project)
	if err != nil {
		return st, err
	}
	contracts := make([]ethcommon.Address, 0, len(addrs))
	for _, a := range addrs {
		contracts = append(contracts, ethcommon.HexToAddress(a))
	}
	log.Printf("Contracts (%d): %v", len(contracts), addrs)

	start, err := resolveNodeFillFromBlock(db, aids, addrs, flagFrom)
	if err != nil {
		return st, err
	}
	if start > endBlock {
		log.Printf("start block %d > end %d — nothing to scan", start, endBlock)
		return st, nil
	}

	archLogStmt, err := db.Prepare(`
		INSERT INTO arch_evtlog (block_num, evt_id, log_index, tx_hash, contract_addr, topic0_sig, log_rlp)
		VALUES ($1, NULL, $2, $3, $4, $5, $6)
		ON CONFLICT (tx_hash, log_index) DO NOTHING
	`)
	if err != nil {
		return st, fmt.Errorf("prepare arch_evtlog: %w", err)
	}
	defer archLogStmt.Close()

	existsStmt, err := db.Prepare(`SELECT 1 FROM arch_evtlog WHERE tx_hash = $1 AND log_index = $2`)
	if err != nil {
		return st, fmt.Errorf("prepare exists check: %w", err)
	}
	defer existsStmt.Close()

	archTxStmt, err := db.Prepare(`
		INSERT INTO arch_tx (block_num, from_aid, to_aid, gas_used, tx_index, num_logs, ctrct_create, value, gas_price, tx_hash, input_sig)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (tx_hash) DO NOTHING
	`)
	if err != nil {
		return st, fmt.Errorf("prepare arch_tx: %w", err)
	}
	defer archTxStmt.Close()

	archBlockStmt, err := db.Prepare(`
		INSERT INTO arch_block (block_num, num_tx, ts, cash_flow, block_hash, parent_hash)
		VALUES ($1, $2, TO_TIMESTAMP($3), 0, $4, $5)
		ON CONFLICT (block_hash) DO NOTHING
	`)
	if err != nil {
		return st, fmt.Errorf("prepare arch_block: %w", err)
	}
	defer archBlockStmt.Close()

	txExistsStmt, err := db.Prepare(`SELECT 1 FROM arch_tx WHERE tx_hash = $1`)
	if err != nil {
		return st, fmt.Errorf("prepare tx exists: %w", err)
	}
	defer txExistsStmt.Close()

	blockExistsStmt, err := db.Prepare(`SELECT 1 FROM arch_block WHERE block_num = $1`)
	if err != nil {
		return st, fmt.Errorf("prepare block exists: %w", err)
	}
	defer blockExistsStmt.Close()

	for from := start; from <= endBlock; {
		to := from + batchSize - 1
		if to > endBlock {
			to = endBlock
		}
		st.BlocksScanned += to - from + 1
		log.Printf("FilterLogs blocks %d .. %d", from, to)

		logs, err := etlcommon.FetchEvents(ctx, client, from, to, contracts)
		if err != nil {
			log.Printf("FilterLogs error [%d..%d]: %v", from, to, err)
			st.RPCErrors++
			if batchSize > 1000 {
				batchSize /= 2
				log.Printf("Reducing batch to %d blocks", batchSize)
				continue
			}
			time.Sleep(3 * time.Second)
			continue
		}

		for i := range logs {
			lg := &logs[i]
			if lg.Removed {
				continue
			}
			st.LogsFromNode++

			txHash := lg.TxHash.Hex()
			logIndex := int(lg.Index)

			var one int
			if err := existsStmt.QueryRow(txHash, logIndex).Scan(&one); err == nil {
				st.LogsSkipped++
				continue
			} else if err != sql.ErrNoRows {
				st.DBErrors++
				log.Printf("exists check: %v", err)
				continue
			}

			if dryRun {
				st.LogsInserted++
				continue
			}

			topic0 := toolutil.Topic0Sig(lg)
			rlpBytes, err := toolutil.EncodeLogRLP(lg)
			if err != nil {
				st.DBErrors++
				log.Printf("encode log: %v", err)
				continue
			}
			contractAddr := lg.Address.Hex()

			res, err := archLogStmt.Exec(int64(lg.BlockNumber), logIndex, txHash, contractAddr, topic0, rlpBytes)
			if err != nil {
				st.DBErrors++
				log.Printf("insert arch_evtlog %s:%d: %v", txHash, logIndex, err)
				continue
			}
			n, _ := res.RowsAffected()
			if n == 0 {
				st.LogsSkipped++
				continue
			}
			st.LogsInserted++

			if inserted, skipped, err := ensureArchTx(ctx, client, addrStore, archTxStmt, txExistsStmt, txHash, int64(lg.BlockNumber)); err != nil {
				st.RPCErrors++
				log.Printf("arch_tx %s: %v", txHash, err)
			} else {
				st.TxInserted += inserted
				st.TxSkipped += skipped
			}

			if inserted, skipped, err := ensureArchBlock(ctx, client, archBlockStmt, blockExistsStmt, int64(lg.BlockNumber)); err != nil {
				st.RPCErrors++
				log.Printf("arch_block %d: %v", lg.BlockNumber, err)
			} else {
				st.BlockInserted += inserted
				st.BlockSkipped += skipped
			}
		}

		from = to + 1
	}

	return st, nil
}

// ensureArchTx inserts the transaction into arch_tx when missing, fetching it
// (and its receipt) from the node.
func ensureArchTx(ctx context.Context, client *ethclient.Client, addrStore *store.Store, ins *sql.Stmt, exists *sql.Stmt, txHash string, blockNum int64) (inserted, skipped int64, err error) {
	var one int
	if err := exists.QueryRow(txHash).Scan(&one); err == nil {
		return 0, 1, nil
	} else if err != sql.ErrNoRows {
		return 0, 0, err
	}

	h := ethcommon.HexToHash(txHash)
	tx, pending, err := client.TransactionByHash(ctx, h)
	if err != nil {
		return 0, 0, err
	}
	if pending {
		return 0, 0, fmt.Errorf("transaction %s still pending", txHash)
	}
	receipt, err := client.TransactionReceipt(ctx, h)
	if err != nil {
		return 0, 0, err
	}

	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		from, err = types.Sender(types.HomesteadSigner{}, tx)
		if err != nil {
			return 0, 0, err
		}
	}
	fromAid, err := addrStore.LookupOrCreateAddress(ctx, from.Hex(), blockNum, 0)
	if err != nil {
		return 0, 0, fmt.Errorf("from address %s: %w", from.Hex(), err)
	}
	var toAid int64
	if tx.To() != nil {
		toAid, err = addrStore.LookupOrCreateAddress(ctx, tx.To().Hex(), blockNum, 0)
		if err != nil {
			return 0, 0, fmt.Errorf("to address %s: %w", tx.To().Hex(), err)
		}
	}
	var gasPrice *big.Int
	if tx.Type() == types.DynamicFeeTxType {
		gasPrice = tx.GasFeeCap()
	} else {
		gasPrice = tx.GasPrice()
	}
	var inputSig string
	if len(tx.Data()) >= 4 {
		inputSig = "0x" + hex.EncodeToString(tx.Data()[:4])
	}

	res, err := ins.Exec(
		blockNum,
		fromAid,
		toAid,
		int64(receipt.GasUsed),
		int(receipt.TransactionIndex),
		len(receipt.Logs),
		tx.To() == nil,
		tx.Value().String(),
		gasPrice.String(),
		txHash,
		inputSig,
	)
	if err != nil {
		return 0, 0, err
	}
	n, _ := res.RowsAffected()
	if n > 0 {
		return 1, 0, nil
	}
	return 0, 1, nil
}

// ensureArchBlock inserts the block into arch_block when missing, fetching it
// from the node.
func ensureArchBlock(ctx context.Context, client *ethclient.Client, ins *sql.Stmt, exists *sql.Stmt, blockNum int64) (inserted, skipped int64, err error) {
	var one int
	if err := exists.QueryRow(blockNum).Scan(&one); err == nil {
		return 0, 1, nil
	} else if err != sql.ErrNoRows {
		return 0, 0, err
	}

	blk, err := client.BlockByNumber(ctx, big.NewInt(blockNum))
	if err != nil {
		return 0, 0, err
	}
	header := blk.Header()
	numTx := int64(len(blk.Transactions()))

	res, err := ins.Exec(
		blockNum,
		numTx,
		header.Time,
		header.Hash().Hex(),
		header.ParentHash.Hex(),
	)
	if err != nil {
		return 0, 0, err
	}
	n, _ := res.RowsAffected()
	if n > 0 {
		return 1, 0, nil
	}
	return 0, 1, nil
}

// mergeFillStats adds src counters into dst.
func mergeFillStats(dst, src *fillStats) {
	dst.BlocksScanned += src.BlocksScanned
	dst.LogsFromNode += src.LogsFromNode
	dst.LogsSkipped += src.LogsSkipped
	dst.LogsInserted += src.LogsInserted
	dst.TxInserted += src.TxInserted
	dst.TxSkipped += src.TxSkipped
	dst.BlockInserted += src.BlockInserted
	dst.BlockSkipped += src.BlockSkipped
	dst.RPCErrors += src.RPCErrors
	dst.DBErrors += src.DBErrors
}

// printFillStats logs one stats block with the given label.
func printFillStats(label string, st *fillStats) {
	log.Printf("[%s] blocks scanned: %d", label, st.BlocksScanned)
	log.Printf("[%s] logs from node: %d", label, st.LogsFromNode)
	log.Printf("[%s] arch_evtlog already present (skipped): %d", label, st.LogsSkipped)
	log.Printf("[%s] arch_evtlog inserted (or would insert): %d", label, st.LogsInserted)
	log.Printf("[%s] arch_tx inserted: %d, skipped (existed): %d", label, st.TxInserted, st.TxSkipped)
	log.Printf("[%s] arch_block inserted: %d, skipped (existed): %d", label, st.BlockInserted, st.BlockSkipped)
	if st.RPCErrors > 0 || st.DBErrors > 0 {
		log.Printf("[%s] errors — rpc: %d, db: %d", label, st.RPCErrors, st.DBErrors)
	}
}
