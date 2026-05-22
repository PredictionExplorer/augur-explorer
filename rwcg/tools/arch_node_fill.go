// Backfill arch_evtlog / arch_tx / arch_block from an Ethereum node via FilterLogs.
// Inserts only rows missing from arch_evtlog (keyed by tx_hash + log_index).
//
// Build: go build -o rwcg/tools/arch_node_fill ./rwcg/tools/arch_node_fill.go
//
// Requires arch_evtlog migrated to PRIMARY KEY (tx_hash, log_index) — see
// rwcg/sql/layer1/migrate_arch_evtlog_natural_key.sql
//
// Example:
//
//	RPC_URL=https://... ./arch_node_fill -project both -db 'postgres://...' -start-block 9292392
package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	etlcommon "github.com/PredictionExplorer/augur-explorer/rwcg/etl/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/lib/pq"
)

const defaultBatchBlocks = 100_000

type fillStats struct {
	BlocksScanned   uint64
	LogsFromNode    int64
	LogsSkipped     int64 // already in arch_evtlog
	LogsInserted    int64
	TxInserted      int64
	TxSkipped       int64
	BlockInserted   int64
	BlockSkipped    int64
	RPCErrors       int64
	DBErrors        int64
}

func main() {
	dbConn := flag.String("db", "", "PostgreSQL connection string")
	projectType := flag.String("project", "", "randomwalk | cosmicgame | both")
	fromBlock := flag.Uint64("from", 0, "Start block (0 = auto: min contract address block_num, else min evt_log block)")
	startBlock := flag.Uint64("start-block", 0, "Start block (same as -from; overrides -from when both are set)")
	toBlock := flag.Uint64("to", 0, "End block inclusive (0 = chain head)")
	batchBlocks := flag.Uint64("batch", defaultBatchBlocks, "FilterLogs block range size")
	dryRun := flag.Bool("dry-run", false, "Scan and report only; do not insert")
	flag.Parse()

	flagFrom := *fromBlock
	if *startBlock > 0 {
		if *fromBlock > 0 && *fromBlock != *startBlock {
			log.Fatalf("-start-block (%d) and -from (%d) disagree; pass only one", *startBlock, *fromBlock)
		}
		flagFrom = *startBlock
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatal("RPC_URL environment variable is required")
	}
	if *dbConn == "" || *projectType == "" {
		log.Fatal("Usage: arch_node_fill -project <randomwalk|cosmicgame|both> -db 'postgres://...' [-start-block N] [-from N] [-to N] [-batch N] [-dry-run]")
	}

	*projectType = strings.ToLower(*projectType)
	projects, err := parseProjects(*projectType)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", *dbConn)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(4)

	if err := requireArchLogIndexPK(db); err != nil {
		log.Fatalf("schema check: %v", err)
	}

	eclient, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("rpc connect: %v", err)
	}

	info := log.New(os.Stdout, "", log.LstdFlags)
	storage := dbs.NewSQLStorageFromDB(db, info)

	head, err := etlcommon.GetCurrentBlockNumber(eclient)
	if err != nil {
		log.Fatalf("chain head: %v", err)
	}
	endBlock := head
	if *toBlock > 0 {
		endBlock = *toBlock
	}

	log.Printf("RPC: %s", rpcURL)
	log.Printf("Chain head: %d, scanning through block %d", head, endBlock)
	if *dryRun {
		log.Println("DRY RUN — no rows will be inserted")
	}

	var total fillStats
	for _, p := range projects {
		log.Printf("")
		log.Printf("========== project: %s ==========", p)
		st := runProject(db, storage, eclient, p, flagFrom, endBlock, *batchBlocks, *dryRun)
		printStats(p, &st)
		mergeStats(&total, &st)
	}

	log.Printf("")
	log.Printf("========== TOTAL ==========")
	printStats("all", &total)
}

func parseProjects(s string) ([]string, error) {
	switch s {
	case "both":
		return []string{"cosmicgame", "randomwalk"}, nil
	case "randomwalk", "cosmicgame":
		return []string{s}, nil
	default:
		return nil, fmt.Errorf("invalid project %q", s)
	}
}

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
		return fmt.Errorf("arch_evtlog.log_index missing — run rwcg/sql/layer1/migrate_arch_evtlog_natural_key.sql first")
	}
	return nil
}

func resolveFromBlock(db *sql.DB, aids []int64, addrs []string, flagFrom uint64) (uint64, error) {
	if flagFrom > 0 {
		log.Printf("Start block: %d (from -start-block / -from)", flagFrom)
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
		return 0, fmt.Errorf("could not auto-detect start block; pass -from <deployment_block>")
	}
	min := candidates[0]
	for _, c := range candidates[1:] {
		if c < min {
			min = c
		}
	}
	log.Printf("Auto start block: %d (from address/evt_log metadata)", min)
	return uint64(min), nil
}

func runProject(db *sql.DB, storage *dbs.SQLStorage, client *ethclient.Client, project string, flagFrom, endBlock, batchSize uint64, dryRun bool) fillStats {
	var st fillStats

	aids := getContractAddressIds(db, project)
	if len(aids) == 0 {
		log.Fatalf("no contracts for project %q", project)
	}
	addrs := getContractAddrsByAids(db, aids)
	contracts := make([]ethcommon.Address, 0, len(addrs))
	for _, a := range addrs {
		contracts = append(contracts, ethcommon.HexToAddress(a))
	}
	log.Printf("Contracts (%d): %v", len(contracts), addrs)

	start, err := resolveFromBlock(db, aids, addrs, flagFrom)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if start > endBlock {
		log.Printf("start block %d > end %d — nothing to scan", start, endBlock)
		return st
	}

	archLogStmt, err := db.Prepare(`
		INSERT INTO arch_evtlog (block_num, evt_id, log_index, tx_hash, contract_addr, topic0_sig, log_rlp)
		VALUES ($1, NULL, $2, $3, $4, $5, $6)
		ON CONFLICT (tx_hash, log_index) DO NOTHING
	`)
	if err != nil {
		log.Fatalf("prepare arch_evtlog: %v", err)
	}
	defer archLogStmt.Close()

	existsStmt, err := db.Prepare(`SELECT 1 FROM arch_evtlog WHERE tx_hash = $1 AND log_index = $2`)
	if err != nil {
		log.Fatalf("prepare exists check: %v", err)
	}
	defer existsStmt.Close()

	archTxStmt, err := db.Prepare(`
		INSERT INTO arch_tx (block_num, from_aid, to_aid, gas_used, tx_index, num_logs, ctrct_create, value, gas_price, tx_hash, input_sig)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (tx_hash) DO NOTHING
	`)
	if err != nil {
		log.Fatalf("prepare arch_tx: %v", err)
	}
	defer archTxStmt.Close()

	archBlockStmt, err := db.Prepare(`
		INSERT INTO arch_block (block_num, num_tx, ts, cash_flow, block_hash, parent_hash)
		VALUES ($1, $2, TO_TIMESTAMP($3), 0, $4, $5)
		ON CONFLICT (block_hash) DO NOTHING
	`)
	if err != nil {
		log.Fatalf("prepare arch_block: %v", err)
	}
	defer archBlockStmt.Close()

	txExistsStmt, err := db.Prepare(`SELECT 1 FROM arch_tx WHERE tx_hash = $1`)
	if err != nil {
		log.Fatalf("prepare tx exists: %v", err)
	}
	defer txExistsStmt.Close()

	blockExistsStmt, err := db.Prepare(`SELECT 1 FROM arch_block WHERE block_num = $1`)
	if err != nil {
		log.Fatalf("prepare block exists: %v", err)
	}
	defer blockExistsStmt.Close()

	ctx := context.Background()
	for from := start; from <= endBlock; {
		to := from + batchSize - 1
		if to > endBlock {
			to = endBlock
		}
		st.BlocksScanned += to - from + 1
		log.Printf("FilterLogs blocks %d .. %d", from, to)

		logs, err := etlcommon.FetchEvents(client, from, to, contracts)
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

			topic0, rlpBytes, err := encodeLog(lg)
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

			if inserted, skipped, err := ensureArchTx(ctx, client, storage, db, archTxStmt, txExistsStmt, txHash, int64(lg.BlockNumber)); err != nil {
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

	return st
}

func encodeLog(lg *types.Log) (topic0 string, rlpBytes []byte, err error) {
	if len(lg.Topics) > 0 {
		full := lg.Topics[0].Hex()[2:]
		if len(full) >= 8 {
			topic0 = full[:8]
		} else {
			topic0 = full
		}
	}
	rlpBytes, err = rlp.EncodeToBytes(lg)
	return topic0, rlpBytes, err
}

func ensureArchTx(ctx context.Context, client *ethclient.Client, storage *dbs.SQLStorage, db *sql.DB, ins *sql.Stmt, exists *sql.Stmt, txHash string, blockNum int64) (inserted, skipped int64, err error) {
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
	fromAid := storage.Lookup_or_create_address(from.Hex(), blockNum, 0)
	var toAid int64
	if tx.To() != nil {
		toAid = storage.Lookup_or_create_address(tx.To().Hex(), blockNum, 0)
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

func mergeStats(dst, src *fillStats) {
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

func printStats(label string, st *fillStats) {
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

func getContractAddressIds(db *sql.DB, projectType string) []int64 {
	var query string
	if projectType == "randomwalk" {
		query = `
			SELECT a.address_id FROM address a
			JOIN rw_contracts rc ON a.addr = rc.randomwalk_addr OR a.addr = rc.marketplace_addr
		`
	} else {
		query = `
			SELECT a.address_id FROM address a
			JOIN cg_contracts cc ON
				a.addr = cc.cosmic_game_addr OR a.addr = cc.cosmic_signature_addr OR
				a.addr = cc.cosmic_token_addr OR a.addr = cc.cosmic_dao_addr OR
				a.addr = cc.charity_wallet_addr OR a.addr = cc.prizes_wallet_addr OR
				a.addr = cc.random_walk_addr OR a.addr = cc.staking_wallet_cst_addr OR
				a.addr = cc.staking_wallet_rwalk_addr OR a.addr = cc.marketing_wallet_addr OR
				a.addr = cc.implementation_addr
		`
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("contract aids: %v", err)
	}
	defer rows.Close()
	var aids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			log.Fatalf("scan: %v", err)
		}
		aids = append(aids, id)
	}
	return aids
}

func getContractAddrsByAids(db *sql.DB, contractAids []int64) []string {
	rows, err := db.Query(`SELECT addr FROM address WHERE address_id = ANY($1) ORDER BY address_id`, pq.Array(contractAids))
	if err != nil {
		log.Fatalf("resolve addrs: %v", err)
	}
	defer rows.Close()
	var addrs []string
	for rows.Next() {
		var addr string
		if err := rows.Scan(&addr); err != nil {
			log.Fatalf("scan addr: %v", err)
		}
		addrs = append(addrs, addr)
	}
	return addrs
}
