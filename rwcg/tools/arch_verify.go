// Archival consistency checks: live evt_log / transaction / block vs arch_* tables.
//
// Build from repo root:
//
//	go build -o rwcg/tools/arch_verify ./rwcg/tools/arch_verify.go
//
// Example:
//
//	./arch_verify -project both -db 'postgres://user:pass@host:5432/dbname?sslmode=disable'
package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/lib/pq"
)

func main() {
	dbConn := flag.String("db", "", "PostgreSQL connection string (same DB holds live + arch_* tables)")
	projectType := flag.String("project", "", "Project: randomwalk | cosmicgame | both (same order as archive_export: cosmicgame then randomwalk)")
	strictOrphans := flag.Bool("strict-orphans", false,
		"If set, fail on any arch_evtlog row whose evt_id is missing from evt_log. Default: only fail on in-range evt_ids (possible deletion); out-of-range ids are reported as legacy imports.")
	strictBlockMeta := flag.Bool("strict-arch-block-metadata", false,
		"If set, require arch_block num_tx, ts, cash_flow to match live block (default: only block_hash and parent_hash must match).")
	strictTxNumLogs := flag.Bool("strict-arch-tx-num-logs", false,
		"If set, require arch_tx.num_logs to match transaction.num_logs (default: ignore num_logs drift; indexer may refresh it after archival).")
	flag.Parse()

	if *dbConn == "" || *projectType == "" {
		log.Fatal("Usage: arch_verify -project <randomwalk|cosmicgame|both> -db 'postgres://...'")
	}

	*projectType = strings.ToLower(*projectType)
	var projects []string
	switch *projectType {
	case "both":
		projects = []string{"cosmicgame", "randomwalk"}
	case "randomwalk", "cosmicgame":
		projects = []string{*projectType}
	default:
		log.Fatalf("Invalid project type %q (want randomwalk, cosmicgame, or both)", *projectType)
	}

	db, err := sql.Open("postgres", *dbConn)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(2)

	allOK := true
	for _, p := range projects {
		ok := verifyProject(db, p, *strictOrphans, *strictBlockMeta, *strictTxNumLogs)
		allOK = allOK && ok
	}

	log.Println("")
	log.Println("=== SUMMARY ===")
	if allOK {
		log.Println("OK — no blocking mismatches for selected project(s). Review any warnings above.")
		return
	}
	log.Println("FAILED — see details above.")
	os.Exit(1)
}

func verifyProject(db *sql.DB, project string, strictOrphans, strictBlockMeta, strictTxNumLogs bool) bool {
	log.Printf("")
	log.Printf("=== Verifying project: %s ===", project)

	aids := getContractAddressIds(db, project)
	if len(aids) == 0 {
		log.Fatalf("No contract addresses for project %q", project)
	}
	addrs := getContractAddrsByAids(db, aids)
	if len(addrs) == 0 {
		log.Fatalf("No resolved addresses for project %q", project)
	}
	log.Printf("Contract AIDs (%d): %v", len(aids), aids)

	var loID, hiID int64
	row := db.QueryRow(`SELECT COALESCE(MIN(id), 0), COALESCE(MAX(id), 0) FROM evt_log`)
	if err := row.Scan(&loID, &hiID); err != nil {
		log.Fatalf("evt_log id bounds: %v", err)
	}
	log.Printf("evt_log id range (global): [%d, %d]", loID, hiID)

	ok := true

	c := count(db, `
		SELECT COUNT(*) FROM evt_log e
		WHERE e.contract_aid = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM arch_evtlog ae WHERE ae.evt_id = e.id)
	`, pq.Array(aids))
	log.Printf("evt_log rows missing from arch_evtlog: %d", c)
	if c != 0 {
		ok = false
	}

	allOrphans := count(db, `
		SELECT COUNT(*) FROM arch_evtlog ae
		WHERE ae.contract_addr = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM evt_log e WHERE e.id = ae.evt_id)
	`, pq.Array(addrs))

	legacyOrphans := count(db, `
		SELECT COUNT(*) FROM arch_evtlog ae
		WHERE ae.contract_addr = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM evt_log e WHERE e.id = ae.evt_id)
		AND ( ae.evt_id < $2 OR ae.evt_id > $3 )
	`, pq.Array(addrs), loID, hiID)

	suspectOrphans := count(db, `
		SELECT COUNT(*) FROM arch_evtlog ae
		WHERE ae.contract_addr = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM evt_log e WHERE e.id = ae.evt_id)
		AND ae.evt_id >= $2 AND ae.evt_id <= $3
	`, pq.Array(addrs), loID, hiID)

	log.Printf("arch_evtlog rows with no evt_log row (evt_id): %d total", allOrphans)
	log.Printf("  … legacy (evt_id outside global evt_log id range [%d, %d]): %d — likely rows archived from another DB / old sequence", loID, hiID, legacyOrphans)
	log.Printf("  … suspect (evt_id inside that range but missing): %d — possible deletions or corruption", suspectOrphans)

	if strictOrphans {
		if allOrphans != 0 {
			ok = false
		}
	} else if suspectOrphans != 0 {
		ok = false
	}

	c = count(db, `
		SELECT COUNT(*) FROM arch_evtlog ae
		INNER JOIN evt_log e ON e.id = ae.evt_id
		INNER JOIN transaction t ON e.tx_id = t.id
		INNER JOIN address a ON e.contract_aid = a.address_id
		WHERE ae.contract_addr = ANY($1)
		AND (
			ae.block_num IS DISTINCT FROM e.block_num
			OR ae.tx_hash IS DISTINCT FROM t.tx_hash
			OR ae.contract_addr IS DISTINCT FROM a.addr
			OR ae.topic0_sig IS DISTINCT FROM e.topic0_sig
			OR ae.log_rlp IS DISTINCT FROM e.log_rlp
		)
	`, pq.Array(addrs))
	log.Printf("arch_evtlog rows that disagree with live evt_log/tx/address: %d", c)
	if c != 0 {
		ok = false
	}

	c = count(db, `
		SELECT COUNT(DISTINCT t.tx_hash) FROM evt_log e
		INNER JOIN transaction t ON e.tx_id = t.id
		WHERE e.contract_aid = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM arch_tx at WHERE at.tx_hash = t.tx_hash)
	`, pq.Array(aids))
	log.Printf("Distinct tx_hash from project evt_log missing in arch_tx: %d", c)
	if c != 0 {
		ok = false
	}

	c = count(db, `
		SELECT COUNT(*) FROM arch_evtlog ae
		WHERE ae.contract_addr = ANY($1)
		AND NOT EXISTS (SELECT 1 FROM arch_tx at WHERE at.tx_hash = ae.tx_hash)
	`, pq.Array(addrs))
	log.Printf("arch_evtlog rows whose tx_hash is missing from arch_tx: %d", c)
	if c != 0 {
		ok = false
	}

	txBad, txNumLogsSkew := txMismatchStats(db, aids, strictTxNumLogs)
	log.Printf("arch_tx vs transaction (project tx_hashes): failing rows: %d", txBad)
	if !strictTxNumLogs && txNumLogsSkew > 0 {
		log.Printf("  … rows that match on core fields but differ only on num_logs: %d (not failing; pass -strict-arch-tx-num-logs to count as failure)", txNumLogsSkew)
	}
	if txBad != 0 {
		ok = false
	}

	c = count(db, `
		SELECT COUNT(DISTINCT at.block_num)
		FROM arch_tx at
		WHERE at.tx_hash IN (
			SELECT DISTINCT ae.tx_hash FROM arch_evtlog ae WHERE ae.contract_addr = ANY($1)
		)
		AND NOT EXISTS (SELECT 1 FROM arch_block ab WHERE ab.block_num = at.block_num)
	`, pq.Array(addrs))
	log.Printf("block_num values (via arch_tx for project logs) missing from arch_block: %d", c)
	if c != 0 {
		ok = false
	}

	hashBad, metaBad := blockMismatchStats(db, addrs)
	log.Printf("arch_block vs live block (same block_num): hash/parent mismatch: %d", hashBad)
	log.Printf("arch_block vs live block (same block_num): metadata mismatch (num_tx, ts, cash_flow): %d", metaBad)
	if metaBad > 0 && !strictBlockMeta {
		log.Printf("  … metadata drift not failing (arch is snapshot; live block may be refreshed — pass -strict-arch-block-metadata to fail on this)")
	}
	if hashBad != 0 {
		ok = false
	}
	if strictBlockMeta && metaBad != 0 {
		ok = false
	}

	if ok {
		log.Printf("--- project %s: PASS (under selected strictness) ---", project)
	} else {
		log.Printf("--- project %s: FAIL ---", project)
	}
	return ok
}

// txMismatchStats returns (failing row count under chosen rules, rows that differ on num_logs only when strictTxNumLogs is false).
func txMismatchStats(db *sql.DB, aids []int64, strictTxNumLogs bool) (failCount, numLogsSkew int64) {
	base := `
		FROM (
			SELECT DISTINCT t.tx_hash AS tx_hash
			FROM evt_log e
			INNER JOIN transaction t ON e.tx_id = t.id
			WHERE e.contract_aid = ANY($1)
		) x
		INNER JOIN transaction t ON t.tx_hash = x.tx_hash
		INNER JOIN arch_tx at ON at.tx_hash = t.tx_hash
	`
	whereStrict := `
		at.block_num IS DISTINCT FROM t.block_num
			OR at.from_aid IS DISTINCT FROM t.from_aid
			OR at.to_aid IS DISTINCT FROM t.to_aid
			OR at.gas_used IS DISTINCT FROM t.gas_used
			OR at.tx_index IS DISTINCT FROM t.tx_index
			OR at.num_logs IS DISTINCT FROM t.num_logs
			OR at.ctrct_create IS DISTINCT FROM t.ctrct_create
			OR at.input_sig IS DISTINCT FROM t.input_sig
			OR at.value IS DISTINCT FROM t.value
			OR at.gas_price IS DISTINCT FROM t.gas_price`
	whereCore := `
		at.block_num IS DISTINCT FROM t.block_num
			OR at.from_aid IS DISTINCT FROM t.from_aid
			OR at.to_aid IS DISTINCT FROM t.to_aid
			OR at.gas_used IS DISTINCT FROM t.gas_used
			OR at.tx_index IS DISTINCT FROM t.tx_index
			OR at.ctrct_create IS DISTINCT FROM t.ctrct_create
			OR at.input_sig IS DISTINCT FROM t.input_sig
			OR at.value IS DISTINCT FROM t.value
			OR at.gas_price IS DISTINCT FROM t.gas_price`

	var nStrict, nCore int64
	if err := db.QueryRow(`SELECT COUNT(*) `+base+` WHERE `+whereStrict, pq.Array(aids)).Scan(&nStrict); err != nil {
		log.Fatalf("tx mismatch (strict): %v", err)
	}
	if err := db.QueryRow(`SELECT COUNT(*) `+base+` WHERE `+whereCore, pq.Array(aids)).Scan(&nCore); err != nil {
		log.Fatalf("tx mismatch (core): %v", err)
	}
	if strictTxNumLogs {
		return nStrict, 0
	}
	// Rows failing full compare but matching on everything except possibly num_logs.
	skew := nStrict - nCore
	if skew < 0 {
		skew = 0
	}
	return nCore, skew
}

// blockMismatchStats counts arch_block rows for project-related blocks joined to live block on block_num.
func blockMismatchStats(db *sql.DB, addrs []string) (hashBad, metaBad int64) {
	sub := `
		SELECT DISTINCT at.block_num FROM arch_tx at
		WHERE at.tx_hash IN (
			SELECT DISTINCT ae.tx_hash FROM arch_evtlog ae WHERE ae.contract_addr = ANY($1)
		)`
	qHash := `
		SELECT COUNT(*) FROM arch_block ab
		INNER JOIN block b ON b.block_num = ab.block_num
		WHERE ab.block_num IN (` + sub + `)
		AND (
			ab.block_hash IS DISTINCT FROM b.block_hash
			OR ab.parent_hash IS DISTINCT FROM b.parent_hash
		)`
	qMeta := `
		SELECT COUNT(*) FROM arch_block ab
		INNER JOIN block b ON b.block_num = ab.block_num
		WHERE ab.block_num IN (` + sub + `)
		AND (
			ab.num_tx IS DISTINCT FROM b.num_tx
			OR ab.ts IS DISTINCT FROM b.ts
			OR ab.cash_flow IS DISTINCT FROM b.cash_flow
		)
		AND ab.block_hash IS NOT DISTINCT FROM b.block_hash
		AND ab.parent_hash IS NOT DISTINCT FROM b.parent_hash`

	if err := db.QueryRow(qHash, pq.Array(addrs)).Scan(&hashBad); err != nil {
		log.Fatalf("block hash mismatch: %v", err)
	}
	if err := db.QueryRow(qMeta, pq.Array(addrs)).Scan(&metaBad); err != nil {
		log.Fatalf("block meta mismatch: %v", err)
	}
	return hashBad, metaBad
}

func count(db *sql.DB, query string, args ...interface{}) int64 {
	var n int64
	if err := db.QueryRow(query, args...).Scan(&n); err != nil {
		log.Fatalf("query failed: %v\n%s", err, query)
	}
	return n
}

func getContractAddressIds(db *sql.DB, projectType string) []int64 {
	var query string
	if projectType == "randomwalk" {
		query = `
			SELECT a.address_id
			FROM address a
			JOIN rw_contracts rc ON a.addr = rc.randomwalk_addr OR a.addr = rc.marketplace_addr
		`
	} else {
		query = `
			SELECT a.address_id
			FROM address a
			JOIN cg_contracts cc ON
				a.addr = cc.cosmic_game_addr OR
				a.addr = cc.cosmic_signature_addr OR
				a.addr = cc.cosmic_token_addr OR
				a.addr = cc.cosmic_dao_addr OR
				a.addr = cc.charity_wallet_addr OR
				a.addr = cc.prizes_wallet_addr OR
				a.addr = cc.random_walk_addr OR
				a.addr = cc.staking_wallet_cst_addr OR
				a.addr = cc.staking_wallet_rwalk_addr OR
				a.addr = cc.marketing_wallet_addr OR
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
	rows, err := db.Query(
		`SELECT addr FROM address WHERE address_id = ANY($1) ORDER BY address_id`,
		pq.Array(contractAids),
	)
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
