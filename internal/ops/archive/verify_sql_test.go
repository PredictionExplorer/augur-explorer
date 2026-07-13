package archive

import (
	"context"
	"database/sql/driver"
	"errors"
	"strings"
	"testing"
)

func verifierCountOps(counts ...int64) []scriptOp {
	ops := []scriptOp{
		queryOp("JOIN rw_contracts", []string{"address_id"}, []driver.Value{int64(8)}),
		queryOp("SELECT addr FROM address", []string{"addr"}, []driver.Value{"0x08"}),
	}
	for _, count := range counts {
		ops = append(ops, queryOp("", []string{"count"}, []driver.Value{count}))
	}
	return ops
}

func TestVerifyProjectsValidationAndCancellation(t *testing.T) {
	if _, err := VerifyProjects(context.Background(), nil, VerifyOptions{}, nil); err == nil {
		t.Fatal("VerifyProjects accepted a nil verifier")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	report, err := VerifyProjects(ctx, []string{ProjectRandomWalk}, VerifyOptions{}, &SQLVerifier{})
	if !errors.Is(err, context.Canceled) || len(report.Projects) != 0 {
		t.Fatalf("report/error = %+v / %v", report, err)
	}
}

func TestSQLVerifierValidationAndContractFailure(t *testing.T) {
	var nilVerifier *SQLVerifier
	if _, err := nilVerifier.VerifyProject(context.Background(), ProjectRandomWalk, VerifyOptions{}); err == nil {
		t.Fatal("nil verifier was accepted")
	}
	if _, err := (&SQLVerifier{}).VerifyProject(context.Background(), ProjectRandomWalk, VerifyOptions{}); err == nil {
		t.Fatal("nil database was accepted")
	}
	verifier := &SQLVerifier{DB: openScriptDB(t)}
	if _, err := verifier.VerifyProject(context.Background(), "unknown", VerifyOptions{}); err == nil {
		t.Fatal("unknown project was accepted")
	}
}

func TestSQLVerifierStrictnessAndWarnings(t *testing.T) {
	// Query order after project metadata:
	// five basic categories, tx strict/core, missing blocks, block hash/meta.
	counts := []int64{0, 0, 0, 0, 0, 2, 0, 0, 0, 3}
	logger := &recordingLogger{}
	report, err := (&SQLVerifier{
		DB:     openScriptDB(t, verifierCountOps(counts...)...),
		Logger: logger,
	}).VerifyProject(context.Background(), ProjectRandomWalk, VerifyOptions{})
	if err != nil {
		t.Fatalf("non-strict verification error = %v", err)
	}
	if !report.Passed ||
		report.Stats.TxNumLogsOnlyMismatches != 2 ||
		report.Stats.BlockMetadataMismatches != 3 {
		t.Fatalf("non-strict report = %+v", report)
	}
	if !logger.contains("num_logs") || !logger.contains("metadata drift") || !logger.contains("PASS") {
		t.Fatalf("warning log lines = %v", logger.lines)
	}

	logger = &recordingLogger{}
	report, err = (&SQLVerifier{
		DB:     openScriptDB(t, verifierCountOps(counts...)...),
		Logger: logger,
	}).VerifyProject(context.Background(), ProjectRandomWalk, VerifyOptions{
		StrictTxNumLogs:     true,
		StrictBlockMetadata: true,
	})
	if err != nil {
		t.Fatalf("strict verification error = %v", err)
	}
	if report.Passed || !logger.contains("failing rows: 2") || !logger.contains("FAIL") {
		t.Fatalf("strict report/logs = %+v / %v", report, logger.lines)
	}
}

func TestSQLVerifierEveryQueryFailure(t *testing.T) {
	sentinel := errors.New("verification query failed")
	const queryCount = 10
	for failAt := 0; failAt < queryCount; failAt++ {
		t.Run(strings.ReplaceAll(string(rune('A'+failAt)), " ", "_"), func(t *testing.T) {
			ops := verifierCountOps()
			for index := 0; index <= failAt; index++ {
				if index == failAt {
					ops = append(ops, queryErrorOp("", sentinel))
				} else {
					ops = append(ops, queryOp("", []string{"count"}, []driver.Value{int64(0)}))
				}
			}
			_, err := (&SQLVerifier{DB: openScriptDB(t, ops...)}).VerifyProject(
				context.Background(),
				ProjectRandomWalk,
				VerifyOptions{},
			)
			if err == nil {
				t.Fatalf("query %d failure was ignored", failAt)
			}
		})
	}
}

func TestMismatchQueryHelpersBranches(t *testing.T) {
	sentinel := errors.New("mismatch query failed")
	t.Run("negative num logs skew clamps", func(t *testing.T) {
		core, skew, err := txMismatchStats(
			context.Background(),
			openScriptDB(t,
				queryOp("", []string{"count"}, []driver.Value{int64(0)}),
				queryOp("", []string{"count"}, []driver.Value{int64(1)}),
			),
			[]int64{8},
		)
		if err != nil || core != 1 || skew != 0 {
			t.Fatalf("core/skew/error = %d/%d/%v", core, skew, err)
		}
	})
	t.Run("strict tx query", func(t *testing.T) {
		_, _, err := txMismatchStats(
			context.Background(),
			openScriptDB(t, queryErrorOp("", sentinel)),
			[]int64{8},
		)
		if !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("core tx query", func(t *testing.T) {
		_, _, err := txMismatchStats(
			context.Background(),
			openScriptDB(t,
				queryOp("", []string{"count"}, []driver.Value{int64(0)}),
				queryErrorOp("", sentinel),
			),
			[]int64{8},
		)
		if !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("hash block query", func(t *testing.T) {
		_, _, err := blockMismatchStats(
			context.Background(),
			openScriptDB(t, queryErrorOp("", sentinel)),
			[]string{"0x08"},
		)
		if !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("metadata block query", func(t *testing.T) {
		_, _, err := blockMismatchStats(
			context.Background(),
			openScriptDB(t,
				queryOp("", []string{"count"}, []driver.Value{int64(0)}),
				queryErrorOp("", sentinel),
			),
			[]string{"0x08"},
		)
		if !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("count query", func(t *testing.T) {
		_, err := countRow(context.Background(), openScriptDB(t, queryErrorOp("", sentinel)), "SELECT count")
		if !errors.Is(err, sentinel) || !strings.Contains(err.Error(), "SELECT count") {
			t.Fatalf("error = %v", err)
		}
	})
}
