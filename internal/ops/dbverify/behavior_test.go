package dbverify_test

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/dbverify"
)

func TestComparisonVariants(t *testing.T) {
	t.Parallel()

	t.Run("events report missing extra and every field mismatch", func(t *testing.T) {
		t.Parallel()
		primary := map[string]dbverify.EventRecord{
			"missing-key-that-is-longer-than-sixteen": {
				BlockNum:  1,
				TxHash:    "missing",
				LogRLPHex: "missing-key-that-is-longer-than-sixteen",
			},
			"shared": {BlockNum: 2, TxHash: "primary", LogRLPHex: "shared"},
		}
		secondary := map[string]dbverify.EventRecord{
			"extra":  {BlockNum: 3, TxHash: "extra", LogRLPHex: "extra"},
			"shared": {BlockNum: 20, TxHash: "secondary", LogRLPHex: "shared"},
		}

		got := dbverify.CompareEvents(primary, secondary, 0)

		if got.PrimaryCount != 2 || got.SecondaryCount != 2 {
			t.Fatalf("counts = %d/%d, want 2/2", got.PrimaryCount, got.SecondaryCount)
		}
		if got.MissingTotal != 1 || got.ExtraTotal != 1 || got.MismatchTotal != 1 {
			t.Fatalf("totals = %+v, want one issue in every category", got)
		}
		if !strings.Contains(got.Missing[0].Summary, "rlp=missing-key-that...") {
			t.Errorf("long RLP summary = %q, want a 16-character prefix", got.Missing[0].Summary)
		}
		wantFields := []dbverify.FieldMismatch{
			{Field: "block_num", Primary: "2", Secondary: "20"},
			{Field: "tx_hash", Primary: "primary", Secondary: "secondary"},
		}
		if !reflect.DeepEqual(got.Mismatched[0].Fields, wantFields) {
			t.Errorf("fields = %+v, want %+v", got.Mismatched[0].Fields, wantFields)
		}
	})

	t.Run("transactions report every category and compared field", func(t *testing.T) {
		t.Parallel()
		got := dbverify.CompareTransactions(
			map[string]dbverify.TransactionRecord{
				"missing": {BlockNum: 1, TxHash: "missing"},
				"shared":  {BlockNum: 2, TxHash: "shared", GasUsed: 3, NumLogs: 4},
			},
			map[string]dbverify.TransactionRecord{
				"extra":  {BlockNum: 5, TxHash: "extra"},
				"shared": {BlockNum: 20, TxHash: "shared", GasUsed: 30, NumLogs: 40},
			},
			10,
		)

		if got.MissingTotal != 1 || got.ExtraTotal != 1 || got.MismatchTotal != 1 {
			t.Fatalf("comparison = %+v", got)
		}
		wantNames := []string{"block_num", "gas_used", "num_logs"}
		if names := mismatchFieldNames(got.Mismatched[0]); !reflect.DeepEqual(names, wantNames) {
			t.Errorf("field names = %v, want %v", names, wantNames)
		}
	})

	t.Run("blocks report every category and compared field", func(t *testing.T) {
		t.Parallel()
		got := dbverify.CompareBlocks(
			map[string]dbverify.BlockRecord{
				"missing": {BlockNum: 1, BlockHash: "missing"},
				"shared":  {BlockNum: 2, BlockHash: "shared", ParentHash: "a", NumTx: 3},
			},
			map[string]dbverify.BlockRecord{
				"extra":  {BlockNum: 4, BlockHash: "extra"},
				"shared": {BlockNum: 20, BlockHash: "shared", ParentHash: "b", NumTx: 30},
			},
			10,
		)

		if got.MissingTotal != 1 || got.ExtraTotal != 1 || got.MismatchTotal != 1 {
			t.Fatalf("comparison = %+v", got)
		}
		wantNames := []string{"block_num", "parent_hash", "num_tx"}
		if names := mismatchFieldNames(got.Mismatched[0]); !reflect.DeepEqual(names, wantNames) {
			t.Errorf("field names = %v, want %v", names, wantNames)
		}
	})

	t.Run("identical records match", func(t *testing.T) {
		t.Parallel()
		event := dbverify.EventRecord{BlockNum: 1, TxHash: "tx", LogRLPHex: "aa"}
		transaction := dbverify.TransactionRecord{BlockNum: 1, TxHash: "tx", GasUsed: 2, NumLogs: 3}
		block := dbverify.BlockRecord{BlockNum: 1, BlockHash: "block", ParentHash: "parent", NumTx: 4}

		comparisons := []dbverify.Comparison{
			dbverify.CompareEvents(
				map[string]dbverify.EventRecord{"aa": event},
				map[string]dbverify.EventRecord{"aa": event},
				10,
			),
			dbverify.CompareTransactions(
				map[string]dbverify.TransactionRecord{"tx": transaction},
				map[string]dbverify.TransactionRecord{"tx": transaction},
				10,
			),
			dbverify.CompareBlocks(
				map[string]dbverify.BlockRecord{"block": block},
				map[string]dbverify.BlockRecord{"block": block},
				10,
			),
		}
		for index, comparison := range comparisons {
			if !comparison.Matched() {
				t.Errorf("comparison %d = %+v, want matched", index, comparison)
			}
		}
	})
}

func TestDetailedEventLogComparisonVariantsAndMultiplicity(t *testing.T) {
	t.Parallel()

	t.Run("missing extra and all metadata fields", func(t *testing.T) {
		t.Parallel()
		primary := []dbverify.EventLogRecord{
			{BlockNum: 1, TxHash: "missing", LogRLP: []byte{0x01}},
			{
				BlockNum:        2,
				TxHash:          "primary",
				ContractAddress: "contract-a",
				Topic0Sig:       "topic-a",
				LogRLP:          []byte{0x02},
			},
		}
		secondary := []dbverify.EventLogRecord{
			{BlockNum: 3, TxHash: "extra", LogRLP: []byte{0x03}},
			{
				BlockNum:        20,
				TxHash:          "secondary",
				ContractAddress: "contract-b",
				Topic0Sig:       "topic-b",
				LogRLP:          []byte{0x02},
			},
		}

		got := dbverify.CompareDetailedEventLogs(primary, secondary, 10)

		if got.MissingTotal != 1 || got.ExtraTotal != 1 || got.MismatchTotal != 1 {
			t.Fatalf("comparison = %+v", got)
		}
		wantNames := []string{"block_num", "tx_hash", "contract_addr", "topic0_sig"}
		if names := mismatchFieldNames(got.Mismatched[0]); !reflect.DeepEqual(names, wantNames) {
			t.Errorf("field names = %v, want %v", names, wantNames)
		}
	})

	t.Run("unequal duplicate counts do not match", func(t *testing.T) {
		t.Parallel()
		event := dbverify.EventLogRecord{BlockNum: 1, TxHash: "tx", LogRLP: []byte{0xaa}}
		got := dbverify.CompareDetailedEventLogs(
			[]dbverify.EventLogRecord{event, event},
			[]dbverify.EventLogRecord{event},
			10,
		)
		if got.Matched() || got.MissingTotal != 1 {
			t.Fatalf("duplicate comparison = %+v, want one missing occurrence", got)
		}
	})

	t.Run("equal duplicate multisets match regardless of input order", func(t *testing.T) {
		t.Parallel()
		first := dbverify.EventLogRecord{BlockNum: 1, TxHash: "a", LogRLP: []byte{0xaa}}
		second := dbverify.EventLogRecord{BlockNum: 2, TxHash: "b", LogRLP: []byte{0xaa}}
		got := dbverify.CompareDetailedEventLogs(
			[]dbverify.EventLogRecord{second, first},
			[]dbverify.EventLogRecord{first, second},
			10,
		)
		if !got.Matched() {
			t.Fatalf("reordered duplicate multiset = %+v, want matched", got)
		}

		indexed := dbverify.IndexEventLogsByRLP([]dbverify.EventLogRecord{second, first, first})
		if len(indexed) != 3 {
			t.Fatalf("indexed duplicate count = %d, want 3", len(indexed))
		}
		if indexed["aa"].TxHash != "a" || indexed["aa#2"].TxHash != "a" || indexed["aa#3"].TxHash != "b" {
			t.Errorf("stable duplicate index = %+v", indexed)
		}
	})
}

func TestVerifyDatabasesEveryLoaderFailure(t *testing.T) {
	t.Parallel()
	errStage := errors.New("stage failed")
	tests := []struct {
		stage       string
		wantMessage string
	}{
		{stage: "primary.events", wantMessage: "query evt_log (primary)"},
		{stage: "secondary.events", wantMessage: "query evt_log (secondary)"},
		{stage: "primary.tx-hashes", wantMessage: "get tx_hashes"},
		{stage: "primary.transactions", wantMessage: "query transactions (primary)"},
		{stage: "secondary.transactions", wantMessage: "query transactions (secondary)"},
		{stage: "primary.block-numbers", wantMessage: "get block_nums"},
		{stage: "primary.blocks", wantMessage: "query blocks (primary)"},
		{stage: "secondary.blocks", wantMessage: "query blocks (secondary)"},
	}
	for _, test := range tests {
		t.Run(test.stage, func(t *testing.T) {
			t.Parallel()
			primary, secondary := newStagedMatchingLoaders(test.stage, errStage, "", nil)

			_, err := dbverify.VerifyDatabases(
				context.Background(),
				primary,
				secondary,
				[]int64{1},
				dbverify.DefaultVerifyReportLimit,
			)

			if !errors.Is(err, errStage) || !strings.Contains(err.Error(), test.wantMessage) {
				t.Fatalf("error = %v, want wrapped %q", err, test.wantMessage)
			}
		})
	}
}

func TestVerifyDatabasesEveryCancellationStage(t *testing.T) {
	t.Parallel()
	stages := []string{
		"primary.events",
		"secondary.events",
		"primary.tx-hashes",
		"primary.transactions",
		"secondary.transactions",
		"primary.block-numbers",
		"primary.blocks",
		"secondary.blocks",
	}
	for _, stage := range stages {
		t.Run(stage, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			primary, secondary := newStagedMatchingLoaders("", nil, stage, cancel)

			_, err := dbverify.VerifyDatabases(ctx, primary, secondary, []int64{1}, 0)

			if !errors.Is(err, context.Canceled) {
				t.Fatalf("error = %v, want context cancellation after %s", err, stage)
			}
		})
	}
}

func TestVerifyDatabasesDefaultsAndRequiresLoaders(t *testing.T) {
	t.Parallel()
	primary, secondary := newStagedMatchingLoaders("", nil, "", nil)

	report, err := dbverify.VerifyDatabases(context.Background(), primary, secondary, []int64{1}, 0)
	if err != nil {
		t.Fatalf("VerifyDatabases() error = %v", err)
	}
	if !report.Matched() || report.Events.ReportLimit != dbverify.DefaultVerifyReportLimit {
		t.Fatalf("report = %+v, want a matched report using the default limit", report)
	}

	for name, loaders := range map[string][2]dbverify.Loader{
		"missing primary":   {nil, secondary},
		"missing secondary": {primary, nil},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := dbverify.VerifyDatabases(context.Background(), loaders[0], loaders[1], nil, 1)
			if err == nil || !strings.Contains(err.Error(), "loaders are required") {
				t.Fatalf("error = %v, want required-loader error", err)
			}
		})
	}
}

func TestDiffEventLogsEveryLoaderFailureAndCancellation(t *testing.T) {
	t.Parallel()
	errStage := errors.New("stage failed")
	tests := []struct {
		stage       string
		wantMessage string
	}{
		{stage: "primary.count", wantMessage: "count primary events"},
		{stage: "secondary.count", wantMessage: "count secondary events"},
		{stage: "primary.detailed", wantMessage: "query primary events"},
		{stage: "secondary.detailed", wantMessage: "query secondary events"},
	}
	for _, test := range tests {
		t.Run("failure "+test.stage, func(t *testing.T) {
			t.Parallel()
			primary, secondary := newStagedMatchingLoaders(test.stage, errStage, "", nil)

			_, err := dbverify.DiffEventLogs(context.Background(), primary, secondary, []int64{1}, 5, 20)

			if !errors.Is(err, errStage) || !strings.Contains(err.Error(), test.wantMessage) {
				t.Fatalf("error = %v, want wrapped %q", err, test.wantMessage)
			}
		})

		t.Run("cancellation "+test.stage, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			primary, secondary := newStagedMatchingLoaders("", nil, test.stage, cancel)

			_, err := dbverify.DiffEventLogs(ctx, primary, secondary, []int64{1}, 5, 0)

			if !errors.Is(err, context.Canceled) {
				t.Fatalf("error = %v, want context cancellation after %s", err, test.stage)
			}
		})
	}
}

func TestDiffEventLogsDefaultsInformationalSemanticsAndRequiredLoaders(t *testing.T) {
	t.Parallel()
	primary, secondary := newStagedMatchingLoaders("", nil, "", nil)
	secondary.detailed = []dbverify.EventLogRecord{{
		BlockNum: 2,
		TxHash:   "different",
		LogRLP:   []byte{0xaa},
	}}

	report, err := dbverify.DiffEventLogs(context.Background(), primary, secondary, []int64{1}, 0, 0)
	if err != nil {
		t.Fatalf("DiffEventLogs() treated data differences as an error: %v", err)
	}
	if report.Comparison.Matched() || report.Comparison.ReportLimit != dbverify.DefaultDiffReportLimit {
		t.Fatalf("report = %+v, want informational mismatch using default limit", report)
	}

	for name, loaders := range map[string][2]dbverify.Loader{
		"missing primary":   {nil, secondary},
		"missing secondary": {primary, nil},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := dbverify.DiffEventLogs(context.Background(), loaders[0], loaders[1], nil, 0, 1)
			if err == nil || !strings.Contains(err.Error(), "loaders are required") {
				t.Fatalf("error = %v, want required-loader error", err)
			}
		})
	}
}

func TestFormatVerifyReportExactOutput(t *testing.T) {
	t.Parallel()
	report := dbverify.VerifyReport{
		Events: dbverify.Comparison{
			PrimaryCount:   3,
			SecondaryCount: 3,
			MissingTotal:   2,
			ExtraTotal:     2,
			MismatchTotal:  2,
			Missing:        []dbverify.RecordIssue{{Summary: "missing event"}},
			Extra:          []dbverify.RecordIssue{{Summary: "extra event"}},
			Mismatched: []dbverify.RecordMismatch{{
				Summary: "event mismatch",
				Fields: []dbverify.FieldMismatch{
					{Field: "block_num", Primary: "1", Secondary: "2"},
					{Field: "tx_hash", Primary: "a", Secondary: "b"},
				},
			}},
		},
		Transactions:       dbverify.Comparison{PrimaryCount: 1, SecondaryCount: 1},
		Blocks:             dbverify.Comparison{PrimaryCount: 1, SecondaryCount: 2, ExtraTotal: 1, Extra: []dbverify.RecordIssue{{Summary: "extra block"}}},
		PrimaryTxHashCount: 1,
		PrimaryBlockCount:  1,
	}
	want := []string{
		"",
		"=== Comparing evt_log tables ===",
		"Primary: 3 events",
		"Secondary: 3 events",
		"MISSING in secondary: missing event",
		"... and 1 more missing records",
		"EXTRA in secondary: extra event",
		"... and 1 more extra records",
		"MISMATCH in secondary: event mismatch: [block_num: 1 vs 2 tx_hash: a vs b]",
		"... and 1 more mismatched records",
		"evt_log: Missing=2, Extra=2, Mismatched=2",
		"",
		"=== Comparing transaction tables ===",
		"Primary has 1 unique transactions for our contracts",
		"Primary: 1 transactions loaded",
		"Secondary: 1 transactions",
		"transaction: Missing=0, Extra=0, Mismatched=0",
		"",
		"=== Comparing block tables ===",
		"Primary has 1 unique blocks for our contracts",
		"Primary: 1 blocks loaded",
		"Secondary: 2 blocks",
		"EXTRA in secondary: extra block",
		"block: Missing=0, Extra=1, Mismatched=0",
		"",
		"=== FINAL SUMMARY ===",
		"✗ evt_log: MISMATCH",
		"✗ block: MISMATCH",
	}

	first := dbverify.FormatVerifyReport(report)
	second := dbverify.FormatVerifyReport(report)
	if !reflect.DeepEqual(first, want) {
		t.Fatalf("FormatVerifyReport() =\n%q\nwant\n%q", first, want)
	}
	if !reflect.DeepEqual(first, second) {
		t.Fatal("FormatVerifyReport() was not deterministic")
	}

	matched := report
	matched.Events = dbverify.Comparison{}
	matched.Blocks = dbverify.Comparison{}
	if got := dbverify.FormatVerifyReport(matched); got[len(got)-1] != "✓ All tables match perfectly!" {
		t.Errorf("matched final line = %q", got[len(got)-1])
	}

	transactionsOnly := report
	transactionsOnly.Events = dbverify.Comparison{}
	transactionsOnly.Blocks = dbverify.Comparison{}
	transactionsOnly.Transactions = dbverify.Comparison{MissingTotal: 1}
	got := dbverify.FormatVerifyReport(transactionsOnly)
	if got[len(got)-1] != "✗ transaction: MISMATCH" {
		t.Errorf("transaction-only final line = %q", got[len(got)-1])
	}
}

func TestFormatEventLogDiffReportExactOutput(t *testing.T) {
	t.Parallel()
	report := dbverify.EventLogDiffReport{
		PrimaryTotal:    10,
		SecondaryTotal:  11,
		PrimaryLoaded:   3,
		SecondaryLoaded: 4,
		Comparison: dbverify.Comparison{
			MissingTotal:  2,
			ExtraTotal:    2,
			MismatchTotal: 2,
			Missing:       []dbverify.RecordIssue{{Summary: "missing event"}},
			Extra:         []dbverify.RecordIssue{{Summary: "extra event"}},
			Mismatched: []dbverify.RecordMismatch{{
				Summary: "event mismatch",
				Fields:  []dbverify.FieldMismatch{{Field: "topic0_sig", Primary: "a", Secondary: "b"}},
			}},
		},
	}
	want := []string{
		"Primary has 10 records for our contracts, Secondary has 11 records total",
		"",
		"=== Loading events from primary (filtered by contract) ===",
		"Loaded 3 events from primary",
		"",
		"=== Loading events from secondary ===",
		"Loaded 4 events from secondary",
		"",
		"=== Comparing events ===",
		"ERROR: Secondary missing event: missing event",
		"... (showing first 1 missing; 1 more)",
		"ERROR: Secondary has extra event: extra event",
		"... (showing first 1 extra; 1 more)",
		"ERROR: Mismatch: event mismatch: [topic0_sig: a vs b]",
		"... (showing first 1 mismatches; 1 more)",
		"",
		"=== SUMMARY ===",
		"Primary events (for our contracts): 3",
		"Secondary events: 4",
		"Missing in secondary: 2 records",
		"Extra in secondary: 2 records",
		"Field mismatches: 2 records",
		"✗ Databases have differences!",
	}

	first := dbverify.FormatEventLogDiffReport(report)
	second := dbverify.FormatEventLogDiffReport(report)
	if !reflect.DeepEqual(first, want) {
		t.Fatalf("FormatEventLogDiffReport() =\n%q\nwant\n%q", first, want)
	}
	if !reflect.DeepEqual(first, second) {
		t.Fatal("FormatEventLogDiffReport() was not deterministic")
	}

	report.Comparison = dbverify.Comparison{}
	got := dbverify.FormatEventLogDiffReport(report)
	if got[len(got)-1] != "✓ Databases match perfectly!" {
		t.Errorf("matched final line = %q", got[len(got)-1])
	}
}

func mismatchFieldNames(mismatch dbverify.RecordMismatch) []string {
	names := make([]string, 0, len(mismatch.Fields))
	for _, field := range mismatch.Fields {
		names = append(names, field.Field)
	}
	return names
}

type stagedLoader struct {
	name        string
	errorAt     string
	errorValue  error
	cancelAfter string
	cancel      context.CancelFunc

	events       map[string]dbverify.EventRecord
	txHashes     []string
	transactions map[string]dbverify.TransactionRecord
	blockNumbers []int64
	blocks       map[string]dbverify.BlockRecord
	count        int64
	detailed     []dbverify.EventLogRecord
}

func newStagedMatchingLoaders(
	errorAt string,
	errorValue error,
	cancelAfter string,
	cancel context.CancelFunc,
) (*stagedLoader, *stagedLoader) {
	event := dbverify.EventRecord{BlockNum: 1, TxHash: "tx", LogRLPHex: "aa"}
	transaction := dbverify.TransactionRecord{BlockNum: 1, TxHash: "tx", GasUsed: 10, NumLogs: 1}
	block := dbverify.BlockRecord{BlockNum: 1, BlockHash: "block", ParentHash: "parent", NumTx: 1}
	detailed := []dbverify.EventLogRecord{{BlockNum: 1, TxHash: "tx", LogRLP: []byte{0xaa}}}
	makeLoader := func(name string) *stagedLoader {
		return &stagedLoader{
			name:         name,
			errorAt:      errorAt,
			errorValue:   errorValue,
			cancelAfter:  cancelAfter,
			cancel:       cancel,
			events:       map[string]dbverify.EventRecord{"aa": event},
			txHashes:     []string{"tx"},
			transactions: map[string]dbverify.TransactionRecord{"tx": transaction},
			blockNumbers: []int64{1},
			blocks:       map[string]dbverify.BlockRecord{"block": block},
			count:        1,
			detailed:     detailed,
		}
	}
	return makeLoader("primary"), makeLoader("secondary")
}

func (l *stagedLoader) step(ctx context.Context, method string) error {
	stage := l.name + "." + method
	if err := ctx.Err(); err != nil {
		return err
	}
	if stage == l.errorAt {
		return l.errorValue
	}
	if stage == l.cancelAfter {
		l.cancel()
	}
	return nil
}

func (l *stagedLoader) LoadEventRecords(
	ctx context.Context,
	_ []int64,
) (map[string]dbverify.EventRecord, error) {
	if err := l.step(ctx, "events"); err != nil {
		return nil, err
	}
	return l.events, nil
}

func (l *stagedLoader) TransactionHashesFromEvents(
	ctx context.Context,
	_ []int64,
) ([]string, error) {
	if err := l.step(ctx, "tx-hashes"); err != nil {
		return nil, err
	}
	return l.txHashes, nil
}

func (l *stagedLoader) LoadTransactions(
	ctx context.Context,
	_ []string,
) (map[string]dbverify.TransactionRecord, error) {
	if err := l.step(ctx, "transactions"); err != nil {
		return nil, err
	}
	return l.transactions, nil
}

func (l *stagedLoader) BlockNumbersFromEvents(
	ctx context.Context,
	_ []int64,
) ([]int64, error) {
	if err := l.step(ctx, "block-numbers"); err != nil {
		return nil, err
	}
	return l.blockNumbers, nil
}

func (l *stagedLoader) LoadBlocks(
	ctx context.Context,
	_ []int64,
) (map[string]dbverify.BlockRecord, error) {
	if err := l.step(ctx, "blocks"); err != nil {
		return nil, err
	}
	return l.blocks, nil
}

func (l *stagedLoader) CountEventLogs(ctx context.Context, _ []int64) (int64, error) {
	if err := l.step(ctx, "count"); err != nil {
		return 0, err
	}
	return l.count, nil
}

func (l *stagedLoader) LoadDetailedEventLogs(
	ctx context.Context,
	_ []int64,
	_ int,
) ([]dbverify.EventLogRecord, error) {
	if err := l.step(ctx, "detailed"); err != nil {
		return nil, err
	}
	return l.detailed, nil
}
