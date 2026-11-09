package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// insertCheck exercises the ChaosZero insert parameter shapes against the
// scratch database: the []string -> TEXT[]::NUMERIC[] array path and the
// []bool -> BOOLEAN[] path, over real block/transaction/evt_log/address
// fixture rows. Enabled by CZCHECK_INSERTS=1.
func insertCheck(ctx context.Context, dsn string) error {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	// Everything rolls back: the scratch database stays clean.
	defer func() { _ = tx.Rollback(ctx) }()

	hash66 := func(last byte) string {
		return fmt.Sprintf("0x%063x%x", 0, last)
	}
	var aid int64
	if err := tx.QueryRow(ctx,
		"INSERT INTO address(block_num,tx_id,addr) VALUES(0,0,'0x00000000000000000000000000000000000000AA') RETURNING address_id").Scan(&aid); err != nil {
		return fmt.Errorf("fixture address: %w", err)
	}
	if _, err := tx.Exec(ctx,
		"INSERT INTO block(block_num,block_hash,parent_hash,ts) VALUES(517037400,$1,$2,TO_TIMESTAMP(1760000000))",
		hash66(0xb), hash66(0xa)); err != nil {
		return fmt.Errorf("fixture block: %w", err)
	}
	var txID int64
	if err := tx.QueryRow(ctx,
		"INSERT INTO transaction(block_num,from_aid,to_aid,tx_hash) VALUES(517037400,$1,$1,$2) RETURNING id",
		aid, hash66(0xc)).Scan(&txID); err != nil {
		return fmt.Errorf("fixture transaction: %w", err)
	}
	var evtlogID int64
	if err := tx.QueryRow(ctx,
		"INSERT INTO evt_log(block_num,tx_id,contract_aid,topic0_sig,log_rlp) VALUES(517037400,$1,$2,'81f6a773','\\x00') RETURNING id",
		txID, aid).Scan(&evtlogID); err != nil {
		return fmt.Errorf("fixture evt_log: %w", err)
	}

	allocations := []string{"1", "2", "3", "4", "5", "6", "7", "115792089237316195423570985008687907853269984665640564039457584007913129639935"}
	if _, err := tx.Exec(ctx, `INSERT INTO cz_round_funded(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,caller_aid,donated,transfers_used,carry_used,allocations
		) VALUES ($1,$2,$3,TO_TIMESTAMP($4),$5,$6,$7,$8,$9,$10,$11::TEXT[]::NUMERIC[])`,
		evtlogID, int64(517037400), txID, int64(1760000000), aid,
		int64(3), aid, "100", "0", "50", allocations); err != nil {
		return fmt.Errorf("cz_round_funded insert: %w", err)
	}

	answers := []bool{true, false, true, false, false, true, false, true}
	escrowed := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	if _, err := tx.Exec(ctx, `INSERT INTO cz_round_settled(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,values_mask,answers,escrowed,carried
		) VALUES (NULL,$1,$2,TO_TIMESTAMP($3),$4,$5,$6,$7,$8::TEXT[]::NUMERIC[],$9)`,
		int64(517037400), txID, int64(1760000000), aid,
		int64(3), int16(165), answers, escrowed, "777"); err != nil {
		return fmt.Errorf("cz_round_settled insert: %w", err)
	}

	var roundTripped string
	if err := tx.QueryRow(ctx,
		"SELECT allocations[8]::TEXT FROM cz_round_funded WHERE evtlog_id=$1", evtlogID).Scan(&roundTripped); err != nil {
		return fmt.Errorf("reading back allocations: %w", err)
	}
	if roundTripped != allocations[7] {
		return fmt.Errorf("allocations[8] round-trip mismatch: %s", roundTripped)
	}
	fmt.Println("insert shapes verified (uint256 max survives NUMERIC[])")
	return nil
}

func maybeInsertCheck(ctx context.Context, dsn string) {
	if os.Getenv("CZCHECK_INSERTS") != "1" {
		return
	}
	if err := insertCheck(ctx, dsn); err != nil {
		fmt.Fprintln(os.Stderr, "insert check:", err)
		os.Exit(1)
	}
}
