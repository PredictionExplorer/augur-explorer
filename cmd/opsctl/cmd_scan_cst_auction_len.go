package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq" // postgres driver
	"github.com/spf13/cobra"
)

const (
	// cstAucLenDefaultContract is the CosmicSignatureGame proxy address.
	cstAucLenDefaultContract = "0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2"
	// cstAucLenDefaultFromBlock is the contract deployment block.
	cstAucLenDefaultFromBlock = uint64(455767589)
	// cstAucLenTopic0Hex is the topic0 of
	// ISystemEventsV2.sol:CstDutchAuctionDurationChangeDivisorChanged(uint256).
	// The new value (uint256) is non-indexed, carried in log.Data.
	cstAucLenTopic0Hex = "0xacbc6b6929088e4b2d043625fa7248e00fb6658a425eb9d9dc8c37b18c7f3e6f"
)

// newScanCstAuctionLenCmd builds `opsctl scan cst-auction-len`, the
// replacement for the standalone scan_cst_auclen_chg_div tool.
func newScanCstAuctionLenCmd() *cobra.Command {
	var (
		contractStr string
		fromBlock   uint64
		toBlock     uint64
		batch       uint64
		dbConn      string
	)
	cmd := &cobra.Command{
		Use:   "cst-auction-len",
		Short: "Scan the chain for CstDutchAuctionDurationChangeDivisorChanged events",
		Long: `Scans the chain for CstDutchAuctionDurationChangeDivisorChanged events on
the CosmicSignatureGame proxy and, when --db is set, reports which on-chain
occurrences are missing from the cg_adm_cst_auclen_chg_div table.

Use this after creating the cg_adm_cst_auclen_chg_div table to find admin
events that may have been emitted while the table did not exist (the ETL
could not insert them then).

Requires the RPC_URL environment variable.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runScanCstAuctionLen(contractStr, fromBlock, toBlock, batch, dbConn)
		},
	}
	cmd.Flags().StringVar(&contractStr, "contract", cstAucLenDefaultContract, "CosmicSignatureGame proxy address")
	cmd.Flags().Uint64Var(&fromBlock, "from-block", cstAucLenDefaultFromBlock, "start block")
	cmd.Flags().Uint64Var(&toBlock, "to-block", 0, "end block (0 = latest)")
	cmd.Flags().Uint64Var(&batch, "batch", defaultFilterBatchBlocks, "FilterLogs block range size")
	cmd.Flags().StringVar(&dbConn, "db", "", "optional postgres conn string; if set, cross-check cg_adm_cst_auclen_chg_div")
	return cmd
}

func init() { scanCmd.AddCommand(newScanCstAuctionLenCmd()) }

func runScanCstAuctionLen(contractStr string, fromBlock, toBlock, batch uint64, dbConn string) error {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		return errors.New("RPC_URL environment variable is required")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return fmt.Errorf("RPC dial: %w", err)
	}
	defer client.Close()

	ctx := context.Background()
	end := toBlock
	if end == 0 {
		end, err = client.BlockNumber(ctx)
		if err != nil {
			return fmt.Errorf("BlockNumber: %w", err)
		}
	}
	if fromBlock > end {
		return fmt.Errorf("invalid range: %d..%d", fromBlock, end)
	}

	contract := ethcommon.HexToAddress(contractStr)
	topic0 := ethcommon.HexToHash(cstAucLenTopic0Hex)

	// Optional DB set of already-stored (tx_hash, log_index).
	var dbKeys map[string]bool
	if dbConn != "" {
		dbKeys, err = loadCstAucLenDBKeys(dbConn)
		if err != nil {
			return err
		}
		log.Printf("Loaded %d existing rows from cg_adm_cst_auclen_chg_div", len(dbKeys))
	}

	log.Printf("Scanning %s for topic %s, blocks %d..%d", contract.Hex(), topic0.Hex(), fromBlock, end)
	if dbKeys != nil {
		fmt.Println("block_num\ttx_hash\tlog_index\tnew_len\tin_db")
	} else {
		fmt.Println("block_num\ttx_hash\tlog_index\tnew_len")
	}

	total, missing := 0, 0
	batchSize := batch
	for from := fromBlock; from <= end; {
		to := from + batchSize - 1
		if to > end {
			to = end
		}
		query := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(from),
			ToBlock:   new(big.Int).SetUint64(to),
			Addresses: []ethcommon.Address{contract},
			Topics:    [][]ethcommon.Hash{{topic0}},
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
		log.Printf("Done. on_chain_events=%d blocks=%d..%d", total, fromBlock, end)
	}
	return nil
}

// loadCstAucLenDBKeys returns the set of (lower(tx_hash)|log_index) already
// recorded in cg_adm_cst_auclen_chg_div, joined through evt_log and
// transaction.
func loadCstAucLenDBKeys(connStr string) (map[string]bool, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("db open: %w", err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT lower(t.tx_hash), e.log_index
		FROM cg_adm_cst_auclen_chg_div r
		JOIN evt_log e ON e.id = r.evtlog_id
		JOIN transaction t ON t.id = e.tx_id`)
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	keys := make(map[string]bool)
	for rows.Next() {
		var txHash string
		var logIndex int64
		if err := rows.Scan(&txHash, &logIndex); err != nil {
			return nil, fmt.Errorf("db scan: %w", err)
		}
		keys[fmt.Sprintf("%s|%d", txHash, logIndex)] = true
	}
	return keys, rows.Err()
}
