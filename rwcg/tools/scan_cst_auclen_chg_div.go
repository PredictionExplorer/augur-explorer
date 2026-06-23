// Scan the chain for CstDutchAuctionDurationChangeDivisorChanged events on the
// CosmicSignatureGame proxy, and optionally report which on-chain occurrences are
// missing from the cg_adm_cst_auclen_chg_div table.
//
// Use this after creating the cg_adm_cst_auclen_chg_div table to find admin events
// that may have been emitted while the table did not exist (the ETL could not
// insert them then).
//
// Event (ISystemEventsV2.sol): CstDutchAuctionDurationChangeDivisorChanged(uint256)
//   topic0 = 0xacbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f
//   The new value (uint256) is non-indexed, carried in log.Data.
//
// Build (single-file, like the other tools in this dir):
//   go build scan_cst_auclen_chg_div.go
//
// Usage:
//   RPC_URL=https://arb1.arbitrum.io/rpc ./scan_cst_auclen_chg_div
//   RPC_URL=https://arb1.arbitrum.io/rpc ./scan_cst_auclen_chg_div -db 'postgres://cgprod@localhost/cgprod?sslmode=disable'
//   RPC_URL=https://... ./scan_cst_auclen_chg_div -from-block 455767589 -to-block 0 -batch 100000
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
)

const (
	defaultContract   = "0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2"
	defaultFromBlock  = uint64(455767589) // contract deployment block
	defaultBatchBlock = uint64(100000)
	// ISystemEventsV2.sol:CstDutchAuctionDurationChangeDivisorChanged(uint256)
	topic0Hex = "0xacbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f"
)

func main() {
	contractStr := flag.String("contract", defaultContract, "CosmicSignatureGame proxy address")
	fromBlock := flag.Uint64("from-block", defaultFromBlock, "start block")
	toBlock := flag.Uint64("to-block", 0, "end block (0 = latest)")
	batch := flag.Uint64("batch", defaultBatchBlock, "FilterLogs block range size")
	dbConn := flag.String("db", "", "optional postgres conn string; if set, cross-check cg_adm_cst_auclen_chg_div")
	flag.Parse()

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatal("RPC_URL environment variable is required")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("RPC dial: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	end := *toBlock
	if end == 0 {
		end, err = client.BlockNumber(ctx)
		if err != nil {
			log.Fatalf("BlockNumber: %v", err)
		}
	}
	if *fromBlock > end {
		log.Fatalf("invalid range: %d..%d", *fromBlock, end)
	}

	contract := common.HexToAddress(*contractStr)
	topic0 := common.HexToHash(topic0Hex)

	// Optional DB set of already-stored (tx_hash, log_index).
	var dbKeys map[string]bool
	if *dbConn != "" {
		dbKeys = loadDBKeys(*dbConn)
		log.Printf("Loaded %d existing rows from cg_adm_cst_auclen_chg_div", len(dbKeys))
	}

	log.Printf("Scanning %s for topic %s, blocks %d..%d", contract.Hex(), topic0.Hex(), *fromBlock, end)
	if dbKeys != nil {
		fmt.Println("block_num\ttx_hash\tlog_index\tnew_len\tin_db")
	} else {
		fmt.Println("block_num\ttx_hash\tlog_index\tnew_len")
	}

	total, missing := 0, 0
	batchSize := *batch
	for from := *fromBlock; from <= end; {
		to := from + batchSize - 1
		if to > end {
			to = end
		}
		query := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(from),
			ToBlock:   new(big.Int).SetUint64(to),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{topic0}},
		}
		logs, err := client.FilterLogs(ctx, query)
		if err != nil {
			log.Printf("FilterLogs error [%d..%d]: %v", from, to, err)
			if batchSize > 1000 {
				batchSize /= 2
				log.Printf("Reducing batch to %d blocks", batchSize)
				continue
			}
			time.Sleep(3 * time.Second)
			continue
		}
		log.Printf("scanned %d..%d (%d events)", from, to, len(logs))

		for i := range logs {
			lg := &logs[i]
			if lg.Removed {
				continue
			}
			total++
			newLen := new(big.Int).SetBytes(lg.Data) // single non-indexed uint256
			txHash := lg.TxHash.Hex()
			line := fmt.Sprintf("%d\t%s\t%d\t%s", lg.BlockNumber, txHash, lg.Index, newLen.String())
			if dbKeys != nil {
				key := fmt.Sprintf("%s|%d", txHash, lg.Index)
				if dbKeys[key] {
					line += "\tyes"
				} else {
					line += "\tNO"
					missing++
				}
			}
			fmt.Println(line)
		}
		from = to + 1
	}

	if dbKeys != nil {
		log.Printf("Done. on_chain_events=%d in_db=%d MISSING_FROM_DB=%d", total, total-missing, missing)
	} else {
		log.Printf("Done. on_chain_events=%d blocks=%d..%d", total, *fromBlock, end)
	}
}

// loadDBKeys returns the set of (lower(tx_hash)|log_index) already recorded in
// cg_adm_cst_auclen_chg_div, joined through evt_log and transaction.
func loadDBKeys(connStr string) map[string]bool {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("db open: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT lower(t.tx_hash), e.log_index
		FROM cg_adm_cst_auclen_chg_div r
		JOIN evt_log e ON e.id = r.evtlog_id
		JOIN transaction t ON t.id = e.tx_id`)
	if err != nil {
		log.Fatalf("db query: %v", err)
	}
	defer rows.Close()

	keys := make(map[string]bool)
	for rows.Next() {
		var txHash string
		var logIndex int64
		if err := rows.Scan(&txHash, &logIndex); err != nil {
			log.Fatalf("db scan: %v", err)
		}
		keys[fmt.Sprintf("%s|%d", txHash, logIndex)] = true
	}
	return keys
}
