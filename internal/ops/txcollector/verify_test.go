package txcollector

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

func fixtureLog(address common.Address, index uint, data byte) *types.Log {
	return &types.Log{
		Address: address,
		Topics:  []common.Hash{common.HexToHash("0x1234")},
		Data:    []byte{data},
		Index:   index,
	}
}

func fixtureEventRow(t *testing.T, block int64, tx *types.Transaction, log *types.Log) EventRow {
	t.Helper()
	encoded, err := toolutil.EncodeLogRLP(log)
	if err != nil {
		t.Fatal(err)
	}
	return EventRow{
		BlockNum:     block,
		LogIndex:     int(log.Index),
		TxHash:       tx.Hash().Hex(),
		ContractAddr: log.Address.Hex(),
		Topic0Sig:    toolutil.Topic0Sig(log),
		LogRLP:       encoded,
	}
}

func writeTxFixture(t *testing.T, dir string, block uint64, pathHash string, tx *types.Transaction) {
	t.Helper()
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		t.Fatal(err)
	}
	if err := writeFileAtomic(toolutil.TxRLPPath(dir, block, pathHash), data); err != nil {
		t.Fatal(err)
	}
}

func writeReceiptFixture(
	t *testing.T,
	dir string,
	block uint64,
	hash string,
	receipt *types.Receipt,
	legacy bool,
) {
	t.Helper()
	var (
		data []byte
		err  error
	)
	if legacy {
		data, err = rlp.EncodeToBytes(receipt)
	} else {
		data, err = toolutil.EncodeBackupReceiptRLP(receipt)
	}
	if err != nil {
		t.Fatal(err)
	}
	if err := writeFileAtomic(toolutil.ReceiptRLPPath(dir, block, hash), data); err != nil {
		t.Fatal(err)
	}
}

func writeBackupFixture(
	t *testing.T,
	dir string,
	block uint64,
	tx *types.Transaction,
	receipt *types.Receipt,
) {
	t.Helper()
	writeTxFixture(t, dir, block, tx.Hash().Hex(), tx)
	writeReceiptFixture(t, dir, block, tx.Hash().Hex(), receipt, false)
}

func TestVerifyHappyPathWithRLPFixtures(t *testing.T) {
	dir := t.TempDir()
	address := common.HexToAddress("0x3000000000000000000000000000000000000003")
	tx := fixtureTx(100)
	first := fixtureLog(address, 4, 0xaa)
	second := fixtureLog(address, 7, 0xbb)
	receipt := fixtureReceipt(tx, first, second)
	writeBackupFixture(t, dir, 50, tx, receipt)
	rows := []EventRow{
		fixtureEventRow(t, 50, tx, second),
		fixtureEventRow(t, 50, tx, first),
	}
	logger := &testLogger{}

	stats, err := Verify(context.Background(), VerifyConfig{
		OutputDir: dir,
		Rows:      rows,
		MaxReport: 50,
		Logger:    logger,
	})
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	want := VerifyStats{
		EvtRowsTotal:  2,
		TxDistinct:    1,
		LogVerifiedOK: 2,
	}
	if stats != want {
		t.Fatalf("stats = %+v, want %+v", stats, want)
	}
	if !strings.Contains(logger.String(), "RESULT: OK") {
		t.Fatalf("summary output = %q", logger.String())
	}
}

// TestVerifyRejectsNegativeBlockNumber pins the corrupt-input guard: a
// negative block number must fail verification loudly instead of wrapping
// into an astronomical backup path that would report every log missing.
func TestVerifyRejectsNegativeBlockNumber(t *testing.T) {
	dir := t.TempDir()
	address := common.HexToAddress("0x3000000000000000000000000000000000000003")
	tx := fixtureTx(102)
	log := fixtureLog(address, 1, 0xaa)
	row := fixtureEventRow(t, 51, tx, log)
	row.BlockNum = -51

	_, err := Verify(context.Background(), VerifyConfig{
		OutputDir: dir,
		Rows:      []EventRow{row},
		MaxReport: 50,
	})
	if err == nil || !strings.Contains(err.Error(), "negative block number") {
		t.Fatalf("Verify error = %v, want negative-block-number rejection", err)
	}
}

func TestVerifyDuplicatesRemainDistinctRows(t *testing.T) {
	dir := t.TempDir()
	address := common.HexToAddress("0x3000000000000000000000000000000000000003")
	tx := fixtureTx(101)
	log := fixtureLog(address, 1, 0xaa)
	writeBackupFixture(t, dir, 51, tx, fixtureReceipt(tx, log))
	row := fixtureEventRow(t, 51, tx, log)

	stats, err := Verify(context.Background(), VerifyConfig{
		OutputDir: dir,
		Rows:      []EventRow{row, row},
		MaxReport: 50,
	})
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if stats.EvtRowsTotal != 2 || stats.TxDistinct != 1 || stats.LogVerifiedOK != 2 {
		t.Fatalf("stats = %+v", stats)
	}
}

func TestVerifyLegacyAndCorruptBackups(t *testing.T) {
	address := common.HexToAddress("0x3000000000000000000000000000000000000003")

	t.Run("legacy receipt is accepted", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(110)
		log := fixtureLog(address, 9, 0xab)
		writeTxFixture(t, dir, 60, tx.Hash().Hex(), tx)
		writeReceiptFixture(t, dir, 60, tx.Hash().Hex(), fixtureReceipt(tx, log), true)

		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir,
			Rows:      []EventRow{fixtureEventRow(t, 60, tx, log)},
			MaxReport: 50,
		})
		if err != nil {
			t.Fatalf("Verify: %v", err)
		}
		if stats.LegacyReceiptBlobs != 1 || stats.LogVerifiedOK != 1 {
			t.Fatalf("stats = %+v", stats)
		}
	})

	t.Run("corrupt receipt is blocking", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(111)
		log := fixtureLog(address, 2, 0xac)
		writeTxFixture(t, dir, 61, tx.Hash().Hex(), tx)
		receiptPath := toolutil.ReceiptRLPPath(dir, 61, tx.Hash().Hex())
		if err := writeFileAtomic(receiptPath, []byte{0xff}); err != nil {
			t.Fatal(err)
		}

		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir,
			Rows:      []EventRow{fixtureEventRow(t, 61, tx, log)},
			MaxReport: 50,
		})
		if !errors.Is(err, ErrVerificationFailed) {
			t.Fatalf("error = %v, want verification failure", err)
		}
		if stats.ReceiptDecodeErr != 1 || !stats.HasFailures() {
			t.Fatalf("stats = %+v", stats)
		}
	})

	t.Run("corrupt transaction is blocking", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(112)
		log := fixtureLog(address, 3, 0xad)
		txPath := toolutil.TxRLPPath(dir, 62, tx.Hash().Hex())
		if err := writeFileAtomic(txPath, []byte{0xff}); err != nil {
			t.Fatal(err)
		}
		writeReceiptFixture(t, dir, 62, tx.Hash().Hex(), fixtureReceipt(tx, log), false)

		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir,
			Rows:      []EventRow{fixtureEventRow(t, 62, tx, log)},
			MaxReport: 50,
		})
		if !errors.Is(err, ErrVerificationFailed) {
			t.Fatalf("error = %v, want verification failure", err)
		}
		if stats.TxDecodeErr != 1 || stats.LogVerifiedOK != 1 || !stats.HasFailures() {
			t.Fatalf("stats = %+v", stats)
		}
	})
}

func TestVerifyMismatchStatsAndFailureSemantics(t *testing.T) {
	address := common.HexToAddress("0x3000000000000000000000000000000000000003")

	t.Run("missing receipt", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(120)
		log := fixtureLog(address, 1, 1)
		writeTxFixture(t, dir, 70, tx.Hash().Hex(), tx)
		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir, Rows: []EventRow{fixtureEventRow(t, 70, tx, log)}, MaxReport: 50,
		})
		if !errors.Is(err, ErrVerificationFailed) || stats.MissingReceiptFile != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("missing transaction also fails coverage", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(121)
		log := fixtureLog(address, 1, 2)
		writeReceiptFixture(t, dir, 71, tx.Hash().Hex(), fixtureReceipt(tx, log), false)
		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir, Rows: []EventRow{fixtureEventRow(t, 71, tx, log)}, MaxReport: 50,
		})
		if !errors.Is(err, ErrVerificationFailed) ||
			stats.MissingTxFile != 1 ||
			stats.SQLTxMissingOnDisk != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("transaction hash mismatch", func(t *testing.T) {
		dir := t.TempDir()
		expected := fixtureTx(122)
		different := fixtureTx(123)
		log := fixtureLog(address, 1, 3)
		writeTxFixture(t, dir, 72, expected.Hash().Hex(), different)
		writeReceiptFixture(t, dir, 72, expected.Hash().Hex(), fixtureReceipt(expected, log), false)
		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir, Rows: []EventRow{fixtureEventRow(t, 72, expected, log)}, MaxReport: 50,
		})
		if !errors.Is(err, ErrVerificationFailed) || stats.TxHashMismatch != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("log absent from receipt", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(124)
		log := fixtureLog(address, 1, 4)
		writeBackupFixture(t, dir, 73, tx, fixtureReceipt(tx))
		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir, Rows: []EventRow{fixtureEventRow(t, 73, tx, log)}, MaxReport: 50,
		})
		if !errors.Is(err, ErrVerificationFailed) || stats.LogNotInReceipt != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("log index mismatch", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(125)
		backupLog := fixtureLog(address, 5, 5)
		writeBackupFixture(t, dir, 74, tx, fixtureReceipt(tx, backupLog))
		row := fixtureEventRow(t, 74, tx, backupLog)
		row.LogIndex = 6
		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir, Rows: []EventRow{row}, MaxReport: 50,
		})
		if !errors.Is(err, ErrVerificationFailed) || stats.LogIndexMismatch != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("log RLP mismatch", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(126)
		backupLog := fixtureLog(address, 5, 6)
		writeBackupFixture(t, dir, 75, tx, fixtureReceipt(tx, backupLog))
		row := fixtureEventRow(t, 75, tx, fixtureLog(address, 5, 0xff))
		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir, Rows: []EventRow{row}, MaxReport: 50,
		})
		if !errors.Is(err, ErrVerificationFailed) || stats.LogRLPMismatch != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("backup-only transaction is non-blocking", func(t *testing.T) {
		dir := t.TempDir()
		sqlTx := fixtureTx(127)
		sqlLog := fixtureLog(address, 1, 7)
		writeBackupFixture(t, dir, 76, sqlTx, fixtureReceipt(sqlTx, sqlLog))
		extraTx := fixtureTx(128)
		writeTxFixture(t, dir, 77, extraTx.Hash().Hex(), extraTx)

		stats, err := Verify(context.Background(), VerifyConfig{
			OutputDir: dir, Rows: []EventRow{fixtureEventRow(t, 76, sqlTx, sqlLog)}, MaxReport: 50,
		})
		if err != nil {
			t.Fatalf("Verify: %v", err)
		}
		if stats.DiskTxExtra != 1 || !stats.HasMismatches() || stats.HasFailures() {
			t.Fatalf("stats = %+v", stats)
		}
	})
}

type callbackLogger struct {
	onLog func()
}

func (l callbackLogger) Printf(string, ...any) {
	if l.onLog != nil {
		l.onLog()
	}
}

func TestVerifyCancellation(t *testing.T) {
	dir := t.TempDir()
	address := common.HexToAddress("0x3000000000000000000000000000000000000003")
	first := fixtureTx(130)
	second := fixtureTx(131)
	rows := []EventRow{
		fixtureEventRow(t, 80, first, fixtureLog(address, 1, 1)),
		fixtureEventRow(t, 81, second, fixtureLog(address, 1, 2)),
	}
	ctx, cancel := context.WithCancel(context.Background())
	logger := callbackLogger{onLog: cancel}

	stats, err := Verify(ctx, VerifyConfig{
		OutputDir: dir,
		Rows:      rows,
		MaxReport: 50,
		Logger:    logger,
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context canceled", err)
	}
	if stats.MissingReceiptFile != 1 {
		t.Fatalf("partial stats = %+v, want first missing receipt", stats)
	}
}

func TestVerifyCancellationCheckpoints(t *testing.T) {
	t.Run("before verification", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		stats, err := Verify(ctx, VerifyConfig{OutputDir: t.TempDir()})
		if !errors.Is(err, context.Canceled) || stats != (VerifyStats{}) {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("after row verification", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(132)
		row := fixtureEventRow(
			t,
			82,
			tx,
			fixtureLog(common.HexToAddress("0x3000000000000000000000000000000000000003"), 1, 1),
		)
		ctx, cancel := context.WithCancel(context.Background())
		logger := hookLogger{hook: func(line string) {
			if strings.HasPrefix(line, "MISSING receipt file") {
				cancel()
			}
		}}
		stats, err := Verify(ctx, VerifyConfig{
			OutputDir: dir, Rows: []EventRow{row}, MaxReport: 50, Logger: logger,
		})
		if !errors.Is(err, context.Canceled) || stats.MissingReceiptFile != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("during coverage", func(t *testing.T) {
		dir := t.TempDir()
		address := common.HexToAddress("0x3000000000000000000000000000000000000003")
		rows := []EventRow{
			fixtureEventRow(t, 83, fixtureTx(133), fixtureLog(address, 1, 1)),
			fixtureEventRow(t, 84, fixtureTx(134), fixtureLog(address, 1, 2)),
		}
		ctx, cancel := context.WithCancel(context.Background())
		logger := hookLogger{hook: func(line string) {
			if strings.HasPrefix(line, "SQL tx missing on disk") {
				cancel()
			}
		}}
		stats, err := Verify(ctx, VerifyConfig{
			OutputDir: dir, Rows: rows, MaxReport: 50, Logger: logger,
		})
		if !errors.Is(err, context.Canceled) || stats.SQLTxMissingOnDisk != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})
}

func TestVerifyRowsCancellationAndReportLimit(t *testing.T) {
	address := common.HexToAddress("0x3000000000000000000000000000000000000003")

	t.Run("cancellation before grouping", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		stats, err := verifyRows(ctx, VerifyConfig{
			OutputDir: t.TempDir(),
			Rows: []EventRow{
				fixtureEventRow(t, 1, fixtureTx(135), fixtureLog(address, 1, 1)),
			},
		})
		if !errors.Is(err, context.Canceled) || stats.EvtRowsTotal != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("cancellation before event comparison", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(136)
		log := fixtureLog(address, 1, 1)
		writeReceiptFixture(t, dir, 85, tx.Hash().Hex(), fixtureReceipt(tx, log), false)
		ctx, cancel := context.WithCancel(context.Background())
		logger := hookLogger{hook: func(line string) {
			if strings.HasPrefix(line, "MISSING tx file") {
				cancel()
			}
		}}
		stats, err := verifyRows(ctx, VerifyConfig{
			OutputDir: dir,
			Rows:      []EventRow{fixtureEventRow(t, 85, tx, log)},
			MaxReport: 50,
			Logger:    logger,
		})
		if !errors.Is(err, context.Canceled) || stats.MissingTxFile != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("report limit and same-block ordering", func(t *testing.T) {
		first := fixtureTx(137)
		second := fixtureTx(138)
		stats, err := verifyRows(context.Background(), VerifyConfig{
			OutputDir: t.TempDir(),
			Rows: []EventRow{
				fixtureEventRow(t, 86, second, fixtureLog(address, 1, 2)),
				fixtureEventRow(t, 86, first, fixtureLog(address, 1, 1)),
			},
			MaxReport: 1,
			Logger:    &testLogger{},
		})
		if err != nil {
			t.Fatalf("verifyRows: %v", err)
		}
		if stats.MissingReceiptFile != 2 || stats.Reported != 1 {
			t.Fatalf("stats = %+v", stats)
		}
	})

	t.Run("log RLP encoder error", func(t *testing.T) {
		dir := t.TempDir()
		tx := fixtureTx(139)
		log := fixtureLog(address, 1, 1)
		writeBackupFixture(t, dir, 87, tx, fixtureReceipt(tx, log))
		boom := errors.New("log encoding failed")
		logger := &testLogger{}
		stats, err := verifyRowsWithEncoder(context.Background(), VerifyConfig{
			OutputDir: dir,
			Rows:      []EventRow{fixtureEventRow(t, 87, tx, log)},
			MaxReport: 50,
			Logger:    logger,
		}, func(*types.Log) ([]byte, error) {
			return nil, boom
		})
		if err != nil {
			t.Fatalf("verifyRowsWithEncoder: %v", err)
		}
		if stats.LogVerifiedOK != 0 || stats.Reported != 1 ||
			!strings.Contains(logger.String(), "RLP encode log") {
			t.Fatalf("stats=%+v logger=%q", stats, logger.String())
		}
	})
}

func TestVerifyReportsTransactionsDeterministically(t *testing.T) {
	dir := t.TempDir()
	address := common.HexToAddress("0x3000000000000000000000000000000000000003")
	early := fixtureTx(140)
	late := fixtureTx(141)
	writeTxFixture(t, dir, 10, early.Hash().Hex(), early)
	writeTxFixture(t, dir, 20, late.Hash().Hex(), late)
	rows := []EventRow{
		fixtureEventRow(t, 20, late, fixtureLog(address, 1, 2)),
		fixtureEventRow(t, 10, early, fixtureLog(address, 1, 1)),
	}
	logger := &testLogger{}

	_, err := Verify(context.Background(), VerifyConfig{
		OutputDir: dir,
		Rows:      rows,
		MaxReport: 50,
		Logger:    logger,
	})
	if !errors.Is(err, ErrVerificationFailed) {
		t.Fatalf("error = %v", err)
	}
	output := logger.String()
	earlyPath := toolutil.ReceiptRLPPath(dir, 10, early.Hash().Hex())
	latePath := toolutil.ReceiptRLPPath(dir, 20, late.Hash().Hex())
	if strings.Index(output, earlyPath) > strings.Index(output, latePath) {
		t.Fatalf("reports are not sorted by block:\n%s", output)
	}
}

func TestVerifyInputValidation(t *testing.T) {
	stats, err := Verify(context.Background(), VerifyConfig{})
	if err == nil || stats != (VerifyStats{}) {
		t.Fatalf("empty output: stats=%+v err=%v", stats, err)
	}
	stats, err = Verify(context.Background(), VerifyConfig{
		OutputDir: t.TempDir(),
		MaxReport: -1,
	})
	if err == nil || stats != (VerifyStats{}) {
		t.Fatalf("negative report limit: stats=%+v err=%v", stats, err)
	}
}

func TestFindBackupLogSelection(t *testing.T) {
	address := common.HexToAddress("0x3000000000000000000000000000000000000003")
	exact := fixtureLog(address, 3, 1)
	sameRLPDifferentIndex := fixtureLog(address, 4, 1)
	receipt := toolutil.ReceiptToBackup(fixtureReceipt(fixtureTx(150), sameRLPDifferentIndex, exact))
	row := fixtureEventRow(t, 1, fixtureTx(151), exact)

	got, ok := findBackupLog(receipt, row, false)
	if !ok || got.Index != exact.Index {
		t.Fatalf("got = %+v, ok=%v; want exact index", got, ok)
	}
	row.LogIndex = 5
	got, ok = findBackupLog(receipt, row, false)
	if !ok || got.Index != sameRLPDifferentIndex.Index {
		t.Fatalf("got = %+v, ok=%v; want deterministic RLP candidate", got, ok)
	}

	wrongAddress := fixtureLog(common.HexToAddress("0x9999999999999999999999999999999999999999"), 3, 1)
	receipt = toolutil.ReceiptToBackup(fixtureReceipt(fixtureTx(152), wrongAddress))
	if got, ok := findBackupLog(receipt, row, false); ok || got != nil {
		t.Fatalf("wrong-address log matched: got=%+v ok=%v", got, ok)
	}
}

func TestCoverageStatsAreOrderIndependent(t *testing.T) {
	rows := []EventRow{
		{BlockNum: 2, TxHash: fixtureTx(160).Hash().Hex()},
		{BlockNum: 1, TxHash: fixtureTx(161).Hash().Hex()},
	}
	onDisk := map[string]uint64{
		fixtureTx(162).Hash().Hex(): 3,
	}
	var first VerifyStats
	if err := checkBackupCoverage(context.Background(), rows, onDisk, &first, 0, nil); err != nil {
		t.Fatal(err)
	}
	reversed := append([]EventRow(nil), rows...)
	reversed[0], reversed[1] = reversed[1], reversed[0]
	var second VerifyStats
	if err := checkBackupCoverage(context.Background(), reversed, onDisk, &second, 0, nil); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(first, second) {
		t.Fatalf("stats differ: %+v vs %+v", first, second)
	}
	if first.SQLTxMissingOnDisk != 2 || first.DiskTxExtra != 1 {
		t.Fatalf("stats = %+v", first)
	}
}

func TestCoverageCancellationAndReportLimit(t *testing.T) {
	t.Run("canceled before SQL grouping", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		stats := VerifyStats{}
		err := checkBackupCoverage(ctx, []EventRow{{TxHash: fixtureTx(170).Hash().Hex()}}, nil, &stats, 0, nil)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
	})

	t.Run("canceled during disk scan", func(t *testing.T) {
		first := fixtureTx(171).Hash().Hex()
		second := fixtureTx(172).Hash().Hex()
		ctx, cancel := context.WithCancel(context.Background())
		logger := hookLogger{hook: func(line string) {
			if strings.HasPrefix(line, "MISMATCH backup-only tx") {
				cancel()
			}
		}}
		stats := VerifyStats{}
		err := checkBackupCoverage(ctx, nil, map[string]uint64{
			first:  1,
			second: 2,
		}, &stats, 0, logger)
		if !errors.Is(err, context.Canceled) || stats.DiskTxExtra != 1 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("report limit", func(t *testing.T) {
		logger := &testLogger{}
		stats := VerifyStats{}
		err := checkBackupCoverage(context.Background(), []EventRow{
			{BlockNum: 1, TxHash: fixtureTx(173).Hash().Hex()},
			{BlockNum: 2, TxHash: fixtureTx(174).Hash().Hex()},
		}, nil, &stats, 1, logger)
		if err != nil {
			t.Fatalf("checkBackupCoverage: %v", err)
		}
		if stats.SQLTxMissingOnDisk != 2 || strings.Count(logger.String(), "SQL tx missing on disk") != 1 {
			t.Fatalf("stats=%+v logger=%q", stats, logger.String())
		}
	})
}

func TestVerifyMissingOutputDirectory(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	_, err := Verify(context.Background(), VerifyConfig{OutputDir: missing, MaxReport: 50})
	if err == nil || !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("error = %v, want missing-directory walk error", err)
	}
}
