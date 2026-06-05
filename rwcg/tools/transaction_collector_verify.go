// transaction-collector-verify checks RLP backup blobs against evt_log in PostgreSQL.
//
// For each evt_log row (scoped by config contract addresses and start_block), verifies:
//   1) tx hash matches the tx RLP blob (receipt RLP does not store tx hash)
//   2) log_index matches the backed-up log (v1 format) or is matched via log_rlp (legacy blobs)
//   3) RLP-encoded log bytes match evt_log.log_rlp
//
// Build:
//
//	go build -o rwcg/tools/transaction-collector-verify ./rwcg/tools/transaction_collector_verify.go
//
// Example (after sourcing ~/configs/cg-prod.env):
//
//	./transaction-collector-verify -config ~/configs/transaction-collector.cosmicgame.json
package main

import (
	"bytes"
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/PredictionExplorer/augur-explorer/rwcg/tools/toolutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/lib/pq"
)

type sqlEvtRow struct {
	BlockNum     int64
	LogIndex     int
	TxHash       string
	ContractAddr string
	Topic0Sig    string
	LogRLP       []byte
}

type verifyStats struct {
	EvtRowsTotal       int64
	TxDistinct         int64
	MissingReceiptFile int64
	MissingTxFile      int64
	ReceiptDecodeErr   int64
	TxDecodeErr        int64
	TxHashMismatch     int64
	LogNotInReceipt    int64
	LogIndexMismatch   int64
	LogRLPMismatch     int64
	LogVerifiedOK      int64
	LegacyReceiptBlobs int64
	SQLTxMissingOnDisk int64
	DiskTxExtra        int64
	Reported           int64
}

func main() {
	configPath := flag.String("config", "", "transaction-collector JSON config (output_dir, contracts, start_block)")
	dbConn := flag.String("db", "", "PostgreSQL URL (default: PGSQL_* from environment)")
	startBlock := flag.Uint64("start-block", 0, "Override config start_block for evt_log filter")
	maxReport := flag.Int("max-report", 50, "Max detailed mismatch lines per category (0 = unlimited)")
	flag.Parse()

	if *configPath == "" {
		log.Fatal("Usage: transaction-collector-verify -config <path.json> [-db 'postgres://...'] [-start-block N]")
	}

	cfg, err := toolutil.LoadCollectorConfig(*configPath)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	contractAddrs, err := cfg.ResolveContractAddresses()
	if err != nil {
		log.Fatalf("contracts: %v", err)
	}

	fromBlock := cfg.StartBlock
	if *startBlock > 0 {
		fromBlock = *startBlock
	}

	conn := *dbConn
	if conn == "" {
		conn, err = toolutil.PostgresConnStringFromEnv()
		if err != nil {
			log.Fatalf("db: %v", err)
		}
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(4)

	log.Printf("Output dir: %s", cfg.OutputDir)
	log.Printf("Contracts (%d): %v", len(contractAddrs), contractAddrs)
	log.Printf("evt_log block_num >= %d", fromBlock)

	rows, err := loadEvtRows(db, contractAddrs, fromBlock)
	if err != nil {
		log.Fatalf("load evt_log: %v", err)
	}
	log.Printf("Loaded %d evt_log rows", len(rows))

	st := verifyAgainstBackup(cfg.OutputDir, rows, *maxReport)

	onDisk, err := toolutil.BackupTxOnDisk(cfg.OutputDir)
	if err != nil {
		log.Fatalf("walk backup: %v", err)
	}
	log.Printf("On-disk tx blobs: %d", len(onDisk))
	checkCoverage(rows, onDisk, &st, *maxReport)

	printSummary(st)
	if st.hasFailures() {
		os.Exit(1)
	}
}

func loadEvtRows(db *sql.DB, contractAddrs []string, fromBlock uint64) ([]sqlEvtRow, error) {
	q := `
		SELECT e.block_num, e.log_index, t.tx_hash, a.addr, e.topic0_sig, e.log_rlp
		FROM evt_log e
		INNER JOIN transaction t ON e.tx_id = t.id
		INNER JOIN address a ON e.contract_aid = a.address_id
		WHERE a.addr = ANY($1)
		AND e.block_num >= $2
		ORDER BY t.tx_hash, e.log_index
	`
	r, err := db.Query(q, pq.Array(contractAddrs), fromBlock)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var out []sqlEvtRow
	for r.Next() {
		var row sqlEvtRow
		if err := r.Scan(&row.BlockNum, &row.LogIndex, &row.TxHash, &row.ContractAddr, &row.Topic0Sig, &row.LogRLP); err != nil {
			return nil, err
		}
		row.TxHash = toolutil.NormalizeTxHash(row.TxHash)
		row.ContractAddr = toolutil.NormalizeAddr(row.ContractAddr)
		out = append(out, row)
	}
	return out, r.Err()
}

func verifyAgainstBackup(outputDir string, rows []sqlEvtRow, maxReport int) verifyStats {
	var st verifyStats
	st.EvtRowsTotal = int64(len(rows))

	type txKey struct {
		hash     string
		blockNum uint64
	}
	byTx := make(map[txKey][]sqlEvtRow)
	for _, row := range rows {
		k := txKey{row.TxHash, uint64(row.BlockNum)}
		byTx[k] = append(byTx[k], row)
	}
	st.TxDistinct = int64(len(byTx))

	report := func(format string, args ...interface{}) {
		if maxReport > 0 && st.Reported >= int64(maxReport) {
			return
		}
		log.Printf(format, args...)
		st.Reported++
	}

	for k, evtRows := range byTx {
		rcptPath := toolutil.ReceiptRLPPath(outputDir, k.blockNum, k.hash)
		txPath := toolutil.TxRLPPath(outputDir, k.blockNum, k.hash)

		rcptData, err := os.ReadFile(rcptPath)
		if err != nil {
			st.MissingReceiptFile++
			report("MISSING receipt file: %s (evt_log rows=%d)", rcptPath, len(evtRows))
			continue
		}

		backupRcpt, legacyFmt, err := toolutil.TryDecodeReceiptRLP(rcptData)
		if err != nil {
			st.ReceiptDecodeErr++
			report("DECODE receipt %s: %v", rcptPath, err)
			continue
		}
		if legacyFmt {
			st.LegacyReceiptBlobs++
		}

		if txData, err := os.ReadFile(txPath); err != nil {
			st.MissingTxFile++
			report("MISSING tx file: %s", txPath)
		} else {
			var tx types.Transaction
			if err := rlp.DecodeBytes(txData, &tx); err != nil {
				st.TxDecodeErr++
				report("DECODE tx %s: %v", txPath, err)
			} else if tx.Hash().Hex() != k.hash {
				st.TxHashMismatch++
				report("TX HASH tx blob %s: blob=%s sql=%s", txPath, tx.Hash().Hex(), k.hash)
			}
		}

		for _, row := range evtRows {
			lg, ok := findBackupLog(backupRcpt, row, legacyFmt)
			if !ok {
				st.LogNotInReceipt++
				report("LOG missing in receipt: tx=%s log_index=%d contract=%s topic0=%s",
					row.TxHash, row.LogIndex, row.ContractAddr, row.Topic0Sig)
				continue
			}
			if !legacyFmt && int(lg.Index) != row.LogIndex {
				st.LogIndexMismatch++
				report("LOG INDEX tx=%s sql_index=%d receipt_index=%d contract=%s",
					row.TxHash, row.LogIndex, lg.Index, row.ContractAddr)
				continue
			}
			encoded, err := toolutil.EncodeLogRLP(lg)
			if err != nil {
				report("RLP encode log tx=%s index=%d: %v", row.TxHash, row.LogIndex, err)
				continue
			}
			if !bytes.Equal(encoded, row.LogRLP) {
				st.LogRLPMismatch++
				report("LOG RLP mismatch: tx=%s log_index=%d contract=%s topic0=%s",
					row.TxHash, row.LogIndex, row.ContractAddr, row.Topic0Sig)
				continue
			}
			st.LogVerifiedOK++
		}
	}
	return st
}

// findBackupLog locates the SQL evt_log row in the decoded backup receipt.
func findBackupLog(br *toolutil.BackupReceipt, row sqlEvtRow, legacyFmt bool) (*types.Log, bool) {
	for i := range br.Logs {
		bl := &br.Logs[i]
		if toolutil.NormalizeAddr(bl.Address.Hex()) != row.ContractAddr {
			continue
		}
		lg := bl.ToTypesLog()
		encoded, err := toolutil.EncodeLogRLP(lg)
		if err != nil {
			continue
		}
		if !bytes.Equal(encoded, row.LogRLP) {
			continue
		}
		if legacyFmt {
			return lg, true
		}
		if int(bl.Index) == row.LogIndex {
			return lg, true
		}
	}
	return nil, false
}

func checkCoverage(rows []sqlEvtRow, onDisk map[string]uint64, st *verifyStats, maxReport int) {
	sqlTxs := make(map[string]uint64)
	for _, row := range rows {
		if _, ok := sqlTxs[row.TxHash]; !ok {
			sqlTxs[row.TxHash] = uint64(row.BlockNum)
		}
	}

	reported := 0
	report := func(format string, args ...interface{}) {
		if maxReport > 0 && reported >= maxReport {
			return
		}
		log.Printf(format, args...)
		reported++
	}

	for txHash, blockNum := range sqlTxs {
		if _, ok := onDisk[txHash]; !ok {
			st.SQLTxMissingOnDisk++
			report("SQL tx missing on disk: %s block=%d", txHash, blockNum)
		}
	}

	for txHash, blockNum := range onDisk {
		if _, ok := sqlTxs[txHash]; !ok {
			st.DiskTxExtra++
			report("MISMATCH backup-only tx (on disk, not in evt_log): %s block=%d", txHash, blockNum)
		}
	}
}

func (s verifyStats) hasFailures() bool {
	return s.MissingReceiptFile > 0 ||
		s.ReceiptDecodeErr > 0 ||
		s.TxHashMismatch > 0 ||
		s.LogNotInReceipt > 0 ||
		s.LogIndexMismatch > 0 ||
		s.LogRLPMismatch > 0 ||
		s.SQLTxMissingOnDisk > 0
}

func (s verifyStats) hasMismatches() bool {
	return s.DiskTxExtra > 0
}

func printSummary(s verifyStats) {
	log.Println("")
	log.Println("=== SUMMARY ===")
	log.Printf("evt_log rows:              %d", s.EvtRowsTotal)
	log.Printf("distinct tx (SQL):         %d", s.TxDistinct)
	log.Printf("logs verified OK:          %d", s.LogVerifiedOK)
	log.Printf("legacy receipt blobs:      %d (log index verified via log_rlp only)", s.LegacyReceiptBlobs)
	log.Printf("missing receipt file:      %d", s.MissingReceiptFile)
	log.Printf("missing tx file:           %d (warning only)", s.MissingTxFile)
	log.Printf("receipt decode errors:     %d", s.ReceiptDecodeErr)
	log.Printf("tx decode errors:          %d", s.TxDecodeErr)
	log.Printf("tx hash mismatches:        %d", s.TxHashMismatch)
	log.Printf("log not in receipt:        %d", s.LogNotInReceipt)
	log.Printf("log index mismatches:      %d", s.LogIndexMismatch)
	log.Printf("log_rlp mismatches:        %d", s.LogRLPMismatch)
	log.Printf("SQL tx missing on disk:    %d", s.SQLTxMissingOnDisk)
	log.Printf("backup-only tx (mismatch): %d (on disk, no matching evt_log — often DAO/admin deploy events)", s.DiskTxExtra)
	if s.hasFailures() {
		log.Println("RESULT: FAILED")
	} else if s.hasMismatches() {
		log.Println("RESULT: OK — SQL evt_log matches backup; backup-only txs noted above (expected for unindexed contracts/events)")
	} else {
		log.Println("RESULT: OK")
	}
}
