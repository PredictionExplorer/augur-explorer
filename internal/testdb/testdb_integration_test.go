//go:build integration

package testdb

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/pressly/goose/v3"
)

// TestMigrationsApply verifies that a fresh container comes up and every
// goose migration applies cleanly — this guards the whole schema.
func TestMigrationsApply(t *testing.T) {
	db := New(t)

	// Spot-check one table per schema group.
	for _, table := range []string{"block", "cg_bid", "rw_mint_evt"} {
		var exists bool
		err := db.SQL.QueryRow(
			`SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)`, table,
		).Scan(&exists)
		if err != nil {
			t.Fatalf("checking table %s: %v", table, err)
		}
		if !exists {
			t.Errorf("expected table %q to exist after migrations", table)
		}
	}
}

func TestBidModerationMigrationDeduplicatesAndEnforcesUniqueness(t *testing.T) {
	db := New(t)
	ctx := context.Background()
	provider, err := goose.NewProvider(goose.DialectPostgres, db.SQL, os.DirFS(migrationsDir()))
	if err != nil {
		t.Fatalf("creating migration provider: %v", err)
	}
	if _, err := provider.DownTo(ctx, 24); err != nil {
		t.Fatalf("rolling back bid-moderation migration: %v", err)
	}

	if _, err := db.SQL.ExecContext(ctx, `INSERT INTO cg_banned_bids(bid_id,user_addr,created_at) VALUES
		(42,'older',100),
		(42,'newer',200)`); err != nil {
		t.Fatalf("seeding duplicate bid bans: %v", err)
	}
	if _, err := provider.UpTo(ctx, 25); err != nil {
		t.Fatalf("reapplying bid-moderation migration: %v", err)
	}

	var count int
	var address string
	if err := db.SQL.QueryRowContext(ctx, `SELECT COUNT(*), MAX(user_addr)
		FROM cg_banned_bids WHERE bid_id=42`).Scan(&count, &address); err != nil {
		t.Fatalf("reading deduplicated bid ban: %v", err)
	}
	if count != 1 || address != "newer" {
		t.Fatalf("deduplicated rows = %d, address=%q; want newest row only", count, address)
	}
	if _, err := db.SQL.ExecContext(ctx,
		`INSERT INTO cg_banned_bids(bid_id,user_addr,created_at) VALUES (42,'duplicate',300)`); err == nil {
		t.Fatal("unique bid-ban index accepted a duplicate bid id")
	}
}

func TestCosmicGameV3MigrationUpDownAndBackfill(t *testing.T) {
	db := New(t)
	ctx := context.Background()
	provider, err := goose.NewProvider(goose.DialectPostgres, db.SQL, os.DirFS(migrationsDir()))
	if err != nil {
		t.Fatalf("creating migration provider: %v", err)
	}
	if _, err := provider.DownTo(ctx, 25); err != nil {
		t.Fatalf("rolling back V3 migration: %v", err)
	}

	const (
		blockHash  = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		parentHash = "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
		txHash     = "0xdddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"
		sentinel   = "0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc"
		gameAddr   = "0x2000000000000000000000000000000000000002"
		bidderAddr = "0x2100000000000000000000000000000000000021"
	)
	if _, err := db.SQL.ExecContext(ctx,
		`INSERT INTO block(block_num,ts,block_hash,parent_hash)
		 VALUES(100,TO_TIMESTAMP(1000),$1,$2)`, blockHash, parentHash); err != nil {
		t.Fatal(err)
	}
	var gameAid, bidderAid int64
	if err := db.SQL.QueryRowContext(ctx,
		`INSERT INTO address(block_num,tx_id,addr) VALUES(100,0,$1) RETURNING address_id`,
		gameAddr).Scan(&gameAid); err != nil {
		t.Fatal(err)
	}
	if err := db.SQL.QueryRowContext(ctx,
		`INSERT INTO address(block_num,tx_id,addr) VALUES(100,0,$1) RETURNING address_id`,
		bidderAddr).Scan(&bidderAid); err != nil {
		t.Fatal(err)
	}
	var txID, sentinelTxID int64
	if err := db.SQL.QueryRowContext(ctx,
		`INSERT INTO transaction(block_num,tx_hash) VALUES(100,$1) RETURNING id`,
		txHash).Scan(&txID); err != nil {
		t.Fatal(err)
	}
	if err := db.SQL.QueryRowContext(ctx,
		`INSERT INTO transaction(block_num,tx_hash) VALUES(100,$1) RETURNING id`,
		sentinel).Scan(&sentinelTxID); err != nil {
		t.Fatal(err)
	}
	var eventID int64
	if err := db.SQL.QueryRowContext(ctx,
		`INSERT INTO evt_log(block_num,tx_id,contract_aid,topic0_sig,log_index,log_rlp)
		 VALUES(100,$1,$2,'1d1f406c',0,$3) RETURNING id`,
		txID, gameAid, []byte{0}).Scan(&eventID); err != nil {
		t.Fatal(err)
	}
	if _, err := db.SQL.ExecContext(ctx,
		`INSERT INTO evt_log(block_num,tx_id,contract_aid,topic0_sig,log_index,log_rlp)
		 VALUES(100,$1,$2,'00000000',1,$3)`,
		sentinelTxID, gameAid, []byte{0}); err != nil {
		t.Fatal(err)
	}
	if _, err := db.SQL.ExecContext(ctx,
		`INSERT INTO cg_bid(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,bidder_aid,
			rwalk_nft_id,round_num,bid_type,bid_position,prize_time,
			eth_price,cst_price,cst_reward,bid_cst_reward_amount,
			cst_dutch_auction_duration,msg
		 ) VALUES($1,100,$2,TO_TIMESTAMP(1000),$3,$4,-1,0,0,1,
			TO_TIMESTAMP(1100),1,-1,100,100,60,'pre-v3')`,
		eventID, txID, gameAid, bidderAid); err != nil {
		t.Fatal(err)
	}

	if _, err := provider.UpTo(ctx, 26); err != nil {
		t.Fatalf("applying V3 migration: %v", err)
	}
	var rewardAmount string
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT amount::TEXT FROM cg_bid_reward WHERE reward_type=0`).Scan(&rewardAmount); err != nil {
		t.Fatal(err)
	}
	if rewardAmount != "100" {
		t.Fatalf("backfilled reward = %s, want 100", rewardAmount)
	}
	var sentinelRows int
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM transaction WHERE tx_hash=$1`, sentinel).Scan(&sentinelRows); err != nil {
		t.Fatal(err)
	}
	if sentinelRows != 1 {
		t.Fatalf("sentinel rows after Up = %d, want 1", sentinelRows)
	}

	if _, err := provider.DownTo(ctx, 25); err != nil {
		t.Fatalf("rolling V3 migration down: %v", err)
	}
	var v3TableExists bool
	if err := db.SQL.QueryRowContext(ctx, `SELECT EXISTS(
		SELECT 1 FROM information_schema.tables WHERE table_name='cg_live_state_updates'
	)`).Scan(&v3TableExists); err != nil {
		t.Fatal(err)
	}
	if v3TableExists {
		t.Fatal("cg_live_state_updates still exists after Down")
	}
	var v3ColumnExists bool
	if err := db.SQL.QueryRowContext(ctx, `SELECT EXISTS(
		SELECT 1 FROM information_schema.columns
		WHERE table_name='cg_prize_claim' AND column_name='num_cs_nfts'
	)`).Scan(&v3ColumnExists); err != nil {
		t.Fatal(err)
	}
	if v3ColumnExists {
		t.Fatal("num_cs_nfts still exists after Down")
	}
	var insertFunction string
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT pg_get_functiondef('on_prize_claim_insert()'::regprocedure)`).Scan(&insertFunction); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(insertFunction, "num_cs_nfts") {
		t.Fatal("Down left the V3 prize trigger body installed")
	}
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM transaction WHERE tx_hash=$1`, sentinel).Scan(&sentinelRows); err != nil {
		t.Fatal(err)
	}
	if sentinelRows != 1 {
		t.Fatalf("sentinel rows after Down = %d, want 1", sentinelRows)
	}
}
