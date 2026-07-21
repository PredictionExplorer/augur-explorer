//go:build integration

package indexer

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/jackc/pgx/v5"
)

var failedBackfillContract = ethcommon.HexToAddress("0x3000000000000000000000000000000000000003")

func addBackfillLog(
	t *testing.T,
	e *env,
	block int64,
	contract ethcommon.Address,
	topic ethcommon.Hash,
	logIndex uint,
) types.Log {
	t.Helper()
	tx := e.chain.AddTx(block, contract, nil)
	log := types.Log{
		Address:     contract,
		Topics:      []ethcommon.Hash{topic},
		Data:        []byte{byte(logIndex)}, // #nosec G115 -- tiny fixed test indexes
		BlockNumber: uint64(block),          // #nosec G115 -- positive test block constants
		BlockHash:   e.chain.BlockHash(block),
		TxHash:      tx.Hash(),
		Index:       logIndex,
	}
	e.chain.AttachLogs(tx.Hash(), []*types.Log{&log})
	return log
}

func TestBackfillContractEvtLogsCommitsPerBlockAndRetryConverges(t *testing.T) {
	e := newEnv(t)
	ctx := context.Background()
	prior := addBackfillLog(t, e, 100, watchedContract, topicA, 0)
	failedFirst := addBackfillLog(t, e, 101, failedBackfillContract, topicA, 0)
	failedSecond := addBackfillLog(t, e, 101, failedBackfillContract, topicB, 1)

	if _, err := e.db.Pool.Exec(ctx, `
		CREATE FUNCTION fail_backfill_later_log() RETURNS trigger
		LANGUAGE plpgsql AS $$
		BEGIN
			IF NEW.block_num = 101 AND NEW.log_index = 1 THEN
				RAISE EXCEPTION 'injected later-log failure';
			END IF;
			RETURN NEW;
		END
		$$;
		CREATE TRIGGER fail_backfill_later_log
		BEFORE INSERT ON evt_log
		FOR EACH ROW EXECUTE FUNCTION fail_backfill_later_log()
	`); err != nil {
		t.Fatalf("install failure trigger: %v", err)
	}

	stats, err := e.engine.BackfillContractEvtLogs(
		ctx,
		[]ethcommon.Address{watchedContract, failedBackfillContract},
		100,
		101,
		100,
	)
	if err == nil || !strings.Contains(err.Error(), "injected later-log failure") {
		t.Fatalf("BackfillContractEvtLogs error = %v", err)
	}
	if stats != (BackfillStats{LogsSeen: 1, Inserted: 1}) {
		t.Fatalf("durable stats after failed block = %+v", stats)
	}
	requireBackfillCount(t, e, "SELECT COUNT(*) FROM block WHERE block_num=100", 1)
	requireBackfillCount(t, e, "SELECT COUNT(*) FROM evt_log WHERE block_num=100", 1)
	requireBackfillCount(t, e, "SELECT COUNT(*) FROM block WHERE block_num=101", 0)
	requireBackfillCount(t, e, "SELECT COUNT(*) FROM transaction WHERE block_num=101", 0)
	requireBackfillCount(t, e, "SELECT COUNT(*) FROM evt_log WHERE block_num=101", 0)
	requireBackfillCount(
		t,
		e,
		"SELECT COUNT(*) FROM address WHERE LOWER(BTRIM(addr))=LOWER($1)",
		0,
		failedBackfillContract.Hex(),
	)
	var lastBlock int64
	if err := e.db.Pool.QueryRow(ctx, "SELECT block_num FROM last_block").Scan(&lastBlock); err != nil {
		t.Fatalf("read last_block: %v", err)
	}
	if lastBlock != 100 {
		t.Fatalf("last_block = %d, want prior committed block 100", lastBlock)
	}
	requireStoredBackfillRLP(t, e, prior)

	if _, err := e.db.Pool.Exec(ctx, `
		DROP TRIGGER fail_backfill_later_log ON evt_log;
		DROP FUNCTION fail_backfill_later_log()
	`); err != nil {
		t.Fatalf("remove failure trigger: %v", err)
	}

	stats, err = e.engine.BackfillContractEvtLogs(
		ctx,
		[]ethcommon.Address{watchedContract, failedBackfillContract},
		100,
		101,
		1,
	)
	if err != nil {
		t.Fatalf("retry BackfillContractEvtLogs: %v", err)
	}
	if stats != (BackfillStats{LogsSeen: 3, Inserted: 2, Skipped: 1}) {
		t.Fatalf("retry stats = %+v", stats)
	}
	requireBackfillCount(t, e, "SELECT COUNT(*) FROM block WHERE block_num=101", 1)
	requireBackfillCount(t, e, "SELECT COUNT(*) FROM transaction WHERE block_num=101", 2)
	requireBackfillCount(t, e, "SELECT COUNT(*) FROM evt_log WHERE block_num=101", 2)
	requireStoredBackfillRLP(t, e, failedFirst)
	requireStoredBackfillRLP(t, e, failedSecond)
}

func TestBackfillBlockRejectsMixedIdentityAndCancellation(t *testing.T) {
	e := newEnv(t)
	first := addBackfillLog(t, e, 110, watchedContract, topicA, 0)
	second := first
	second.BlockHash = ethcommon.HexToHash(
		"0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
	)
	if _, err := e.engine.backfillBlock(context.Background(), []types.Log{first, second}); err == nil ||
		!strings.Contains(err.Error(), "inconsistent block identity") {
		t.Fatalf("mixed block identity error = %v", err)
	}
	requireBackfillCount(t, e, "SELECT COUNT(*) FROM block WHERE block_num=110", 0)

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	stats, err := e.engine.BackfillContractEvtLogs(
		cancelled,
		[]ethcommon.Address{watchedContract},
		110,
		110,
		1,
	)
	if err == nil || !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled backfill = %+v, %v", stats, err)
	}
	if stats != (BackfillStats{}) {
		t.Fatalf("cancelled stats = %+v", stats)
	}
}

func TestBackfillContractEvtLogsPropagatesStoreFailures(t *testing.T) {
	t.Run("block verification", func(t *testing.T) {
		e := newEnv(t)
		log := addBackfillLog(t, e, 120, watchedContract, topicA, 0)
		log.BlockHash = ethcommon.HexToHash(
			"0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
		)
		if _, err := e.engine.backfillBlock(context.Background(), []types.Log{log}); err == nil ||
			!strings.Contains(err.Error(), "hash mismatch") {
			t.Fatalf("block verification error = %v", err)
		}
	})

	t.Run("transaction lookup", func(t *testing.T) {
		e := newEnv(t)
		addBackfillLog(t, e, 121, watchedContract, topicA, 0)
		restore := hideBackfillTable(t, e, "transaction")
		defer restore()
		if _, err := e.engine.BackfillContractEvtLogs(
			context.Background(),
			[]ethcommon.Address{watchedContract},
			121,
			121,
			1,
		); err == nil || !strings.Contains(err.Error(), "transaction id lookup") {
			t.Fatalf("transaction lookup error = %v", err)
		}
	})

	t.Run("transaction insertion", func(t *testing.T) {
		e := newEnv(t)
		addBackfillLog(t, e, 122, watchedContract, topicA, 0)
		restore := hideBackfillTable(t, e, "address")
		defer restore()
		if _, err := e.engine.BackfillContractEvtLogs(
			context.Background(),
			[]ethcommon.Address{watchedContract},
			122,
			122,
			1,
		); err == nil || !strings.Contains(err.Error(), "EnsureTransactionExists") {
			t.Fatalf("transaction insertion error = %v", err)
		}
	})

	t.Run("event existence", func(t *testing.T) {
		e := newEnv(t)
		log := addBackfillLog(t, e, 123, watchedContract, topicA, 0)
		ctx := context.Background()
		if _, err := e.engine.EnsureBlockExists(ctx, 123, log.BlockHash.Hex()); err != nil {
			t.Fatal(err)
		}
		if _, _, err := e.engine.EnsureTransactionExists(ctx, log.TxHash, 123); err != nil {
			t.Fatal(err)
		}
		restore := hideBackfillTable(t, e, "evt_log")
		defer restore()
		if _, err := e.engine.BackfillContractEvtLogs(
			ctx,
			[]ethcommon.Address{watchedContract},
			123,
			123,
			1,
		); err == nil || !strings.Contains(err.Error(), "evt_log existence check") {
			t.Fatalf("event existence error = %v", err)
		}
	})
}

func hideBackfillTable(t *testing.T, e *env, table string) func() {
	t.Helper()
	backup := table + "_backfill_test_backup"
	tableID := pgx.Identifier{table}.Sanitize()
	backupID := pgx.Identifier{backup}.Sanitize()
	rename := func(from, to string) {
		t.Helper()
		if _, err := e.db.Pool.Exec(context.Background(), "ALTER TABLE "+from+" RENAME TO "+to); err != nil {
			t.Fatalf("rename %s to %s: %v", from, to, err)
		}
		e.db.Pool.Reset()
	}
	rename(tableID, backupID)
	restored := false
	return func() {
		if !restored {
			restored = true
			rename(backupID, tableID)
		}
	}
}

func requireBackfillCount(t *testing.T, e *env, query string, want int, args ...any) {
	t.Helper()
	var got int
	if err := e.db.Pool.QueryRow(context.Background(), query, args...).Scan(&got); err != nil {
		t.Fatalf("count query %q: %v", query, err)
	}
	if got != want {
		t.Fatalf("count query %q = %d, want %d", query, got, want)
	}
}

func requireStoredBackfillRLP(t *testing.T, e *env, log types.Log) {
	t.Helper()
	want, err := rlp.EncodeToBytes(&log)
	if err != nil {
		t.Fatalf("encode expected RLP: %v", err)
	}
	var got []byte
	if err := e.db.Pool.QueryRow(
		context.Background(),
		"SELECT log_rlp FROM evt_log WHERE block_num=$1 AND log_index=$2",
		int64(log.BlockNumber), // #nosec G115 -- positive test block constants
		int(log.Index),
	).Scan(&got); err != nil {
		t.Fatalf("read stored RLP: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("stored RLP changed for block=%d index=%d", log.BlockNumber, log.Index)
	}
}
