package dbverify_test

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/dbverify"
)

func TestCompareEventsCategoriesOrderingAndLimit(t *testing.T) {
	t.Parallel()
	primary := map[string]dbverify.EventRecord{
		"b":    {BlockNum: 2, TxHash: "tx-b", LogRLPHex: "b"},
		"a":    {BlockNum: 1, TxHash: "tx-a", LogRLPHex: "a"},
		"same": {BlockNum: 3, TxHash: "primary", LogRLPHex: "same"},
	}
	secondary := map[string]dbverify.EventRecord{
		"c":    {BlockNum: 4, TxHash: "tx-c", LogRLPHex: "c"},
		"same": {BlockNum: 30, TxHash: "secondary", LogRLPHex: "same"},
	}
	comparison := dbverify.CompareEvents(primary, secondary, 1)
	if comparison.MissingTotal != 2 || comparison.ExtraTotal != 1 || comparison.MismatchTotal != 1 {
		t.Fatalf("comparison totals = %+v", comparison)
	}
	if len(comparison.Missing) != 1 || comparison.Missing[0].Key != "a" {
		t.Errorf("limited/sorted missing = %+v, want key a", comparison.Missing)
	}
	if len(comparison.Extra) != 1 || comparison.Extra[0].Key != "c" {
		t.Errorf("extra = %+v, want key c", comparison.Extra)
	}
	if len(comparison.Mismatched) != 1 || comparison.Mismatched[0].Key != "same" {
		t.Fatalf("mismatched = %+v", comparison.Mismatched)
	}
	fields := comparison.Mismatched[0].Fields
	if len(fields) != 2 || fields[0].Field != "block_num" || fields[1].Field != "tx_hash" {
		t.Errorf("mismatch fields = %+v", fields)
	}
	if comparison.Matched() {
		t.Error("comparison with differences reported matched")
	}
}

func TestCompareTransactionsAndBlocksDetectFieldMismatches(t *testing.T) {
	t.Parallel()
	tx := dbverify.CompareTransactions(
		map[string]dbverify.TransactionRecord{
			"tx": {BlockNum: 1, TxHash: "tx", GasUsed: 10, NumLogs: 2},
		},
		map[string]dbverify.TransactionRecord{
			"tx": {BlockNum: 1, TxHash: "tx", GasUsed: 11, NumLogs: 3},
		},
		10,
	)
	if tx.MismatchTotal != 1 || len(tx.Mismatched[0].Fields) != 2 {
		t.Errorf("transaction comparison = %+v", tx)
	}

	blocks := dbverify.CompareBlocks(
		map[string]dbverify.BlockRecord{
			"hash": {BlockNum: 1, BlockHash: "hash", ParentHash: "parent-a", NumTx: 4},
		},
		map[string]dbverify.BlockRecord{
			"hash": {BlockNum: 2, BlockHash: "hash", ParentHash: "parent-b", NumTx: 5},
		},
		10,
	)
	if blocks.MismatchTotal != 1 || len(blocks.Mismatched[0].Fields) != 3 {
		t.Errorf("block comparison = %+v", blocks)
	}
}

func TestCompareDetailedEventLogsMismatchCategories(t *testing.T) {
	t.Parallel()
	primary := []dbverify.EventLogRecord{{
		BlockNum:        1,
		TxHash:          "tx-a",
		ContractAddress: "contract-a",
		Topic0Sig:       "topic-a",
		LogRLP:          []byte{0x01},
	}}
	secondary := []dbverify.EventLogRecord{{
		BlockNum:        2,
		TxHash:          "tx-b",
		ContractAddress: "contract-b",
		Topic0Sig:       "topic-b",
		LogRLP:          []byte{0x01},
	}}
	comparison := dbverify.CompareDetailedEventLogs(primary, secondary, 20)
	if comparison.MissingTotal != 0 || comparison.ExtraTotal != 0 || comparison.MismatchTotal != 1 {
		t.Fatalf("comparison = %+v", comparison)
	}
	gotFields := make([]string, 0, len(comparison.Mismatched[0].Fields))
	for _, field := range comparison.Mismatched[0].Fields {
		gotFields = append(gotFields, field.Field)
	}
	wantFields := []string{"block_num", "tx_hash", "contract_addr", "topic0_sig"}
	if !reflect.DeepEqual(gotFields, wantFields) {
		t.Errorf("fields = %v, want %v", gotFields, wantFields)
	}
}

type fakeLoader struct {
	events       map[string]dbverify.EventRecord
	txHashes     []string
	transactions map[string]dbverify.TransactionRecord
	blockNumbers []int64
	blocks       map[string]dbverify.BlockRecord
	count        int64
	detailed     []dbverify.EventLogRecord
	failMethod   string
	calls        []string
}

func (l *fakeLoader) call(ctx context.Context, method string) error {
	l.calls = append(l.calls, method)
	if err := ctx.Err(); err != nil {
		return err
	}
	if l.failMethod == method {
		return errors.New("loader failure")
	}
	return nil
}

func (l *fakeLoader) LoadEventRecords(ctx context.Context, _ []int64) (map[string]dbverify.EventRecord, error) {
	if err := l.call(ctx, "events"); err != nil {
		return nil, err
	}
	return l.events, nil
}

func (l *fakeLoader) TransactionHashesFromEvents(ctx context.Context, _ []int64) ([]string, error) {
	if err := l.call(ctx, "tx-hashes"); err != nil {
		return nil, err
	}
	return l.txHashes, nil
}

func (l *fakeLoader) LoadTransactions(ctx context.Context, _ []string) (map[string]dbverify.TransactionRecord, error) {
	if err := l.call(ctx, "transactions"); err != nil {
		return nil, err
	}
	return l.transactions, nil
}

func (l *fakeLoader) BlockNumbersFromEvents(ctx context.Context, _ []int64) ([]int64, error) {
	if err := l.call(ctx, "block-numbers"); err != nil {
		return nil, err
	}
	return l.blockNumbers, nil
}

func (l *fakeLoader) LoadBlocks(ctx context.Context, _ []int64) (map[string]dbverify.BlockRecord, error) {
	if err := l.call(ctx, "blocks"); err != nil {
		return nil, err
	}
	return l.blocks, nil
}

func (l *fakeLoader) CountEventLogs(ctx context.Context, _ []int64) (int64, error) {
	if err := l.call(ctx, "count"); err != nil {
		return 0, err
	}
	return l.count, nil
}

func (l *fakeLoader) LoadDetailedEventLogs(
	ctx context.Context,
	_ []int64,
	_ int,
) ([]dbverify.EventLogRecord, error) {
	if err := l.call(ctx, "detailed"); err != nil {
		return nil, err
	}
	return l.detailed, nil
}

func matchingLoaders() (*fakeLoader, *fakeLoader) {
	event := dbverify.EventRecord{BlockNum: 1, TxHash: "tx", LogRLPHex: "aa"}
	transaction := dbverify.TransactionRecord{BlockNum: 1, TxHash: "tx", GasUsed: 10, NumLogs: 1}
	block := dbverify.BlockRecord{BlockNum: 1, BlockHash: "block", ParentHash: "parent", NumTx: 1}
	primary := &fakeLoader{
		events:       map[string]dbverify.EventRecord{"aa": event},
		txHashes:     []string{"tx"},
		transactions: map[string]dbverify.TransactionRecord{"tx": transaction},
		blockNumbers: []int64{1},
		blocks:       map[string]dbverify.BlockRecord{"block": block},
	}
	secondary := &fakeLoader{
		events:       map[string]dbverify.EventRecord{"aa": event},
		transactions: map[string]dbverify.TransactionRecord{"tx": transaction},
		blocks:       map[string]dbverify.BlockRecord{"block": block},
	}
	return primary, secondary
}

func TestVerifyDatabasesMatchingAndDivergent(t *testing.T) {
	t.Parallel()
	primary, secondary := matchingLoaders()
	report, err := dbverify.VerifyDatabases(context.Background(), primary, secondary, []int64{1}, 10)
	if err != nil {
		t.Fatalf("VerifyDatabases() error = %v", err)
	}
	if !report.Matched() {
		t.Errorf("matching report = %+v", report)
	}

	secondary.transactions["tx"] = dbverify.TransactionRecord{
		BlockNum: 1,
		TxHash:   "tx",
		GasUsed:  999,
		NumLogs:  1,
	}
	report, err = dbverify.VerifyDatabases(context.Background(), primary, secondary, []int64{1}, 10)
	if err != nil {
		t.Fatalf("VerifyDatabases(divergent) error = %v", err)
	}
	if report.Matched() || report.Transactions.MismatchTotal != 1 {
		t.Errorf("divergent report = %+v", report)
	}
}

func TestVerifyDatabasesErrorAndCancellationPropagation(t *testing.T) {
	t.Parallel()
	primary, secondary := matchingLoaders()
	secondary.failMethod = "transactions"
	_, err := dbverify.VerifyDatabases(context.Background(), primary, secondary, []int64{1}, 10)
	if err == nil || !strings.Contains(err.Error(), "query transactions (secondary)") {
		t.Fatalf("error = %v", err)
	}
	if contains(secondary.calls, "blocks") {
		t.Errorf("continued loading after failure: %v", secondary.calls)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	primary, secondary = matchingLoaders()
	_, err = dbverify.VerifyDatabases(ctx, primary, secondary, []int64{1}, 10)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("canceled error = %v", err)
	}
}

func TestDiffEventLogsIsInformationalAndPropagatesQueries(t *testing.T) {
	t.Parallel()
	primary := &fakeLoader{
		count: 1,
		detailed: []dbverify.EventLogRecord{{
			BlockNum: 1,
			TxHash:   "primary",
			LogRLP:   []byte{0x01},
		}},
	}
	secondary := &fakeLoader{
		count: 1,
		detailed: []dbverify.EventLogRecord{{
			BlockNum: 2,
			TxHash:   "secondary",
			LogRLP:   []byte{0x01},
		}},
	}
	report, err := dbverify.DiffEventLogs(context.Background(), primary, secondary, []int64{1}, 0, 20)
	if err != nil {
		t.Fatalf("DiffEventLogs() returned mismatch as error: %v", err)
	}
	if report.Comparison.Matched() || report.Comparison.MismatchTotal != 1 {
		t.Errorf("informational report = %+v", report)
	}

	secondary.failMethod = "count"
	_, err = dbverify.DiffEventLogs(context.Background(), primary, secondary, []int64{1}, 0, 20)
	if err == nil || !strings.Contains(err.Error(), "count secondary events") {
		t.Errorf("query error = %v", err)
	}
}

func TestFormattedReportsAreStableAndRespectLimits(t *testing.T) {
	t.Parallel()
	comparison := dbverify.CompareEvents(
		map[string]dbverify.EventRecord{
			"b": {BlockNum: 2, TxHash: "b"},
			"a": {BlockNum: 1, TxHash: "a"},
		},
		map[string]dbverify.EventRecord{},
		1,
	)
	report := dbverify.VerifyReport{
		Events:       comparison,
		Transactions: dbverify.CompareTransactions(nil, nil, 1),
		Blocks:       dbverify.CompareBlocks(nil, nil, 1),
	}
	first := dbverify.FormatVerifyReport(report)
	second := dbverify.FormatVerifyReport(report)
	if !reflect.DeepEqual(first, second) {
		t.Fatalf("report formatting is not deterministic:\n%v\n%v", first, second)
	}
	joined := strings.Join(first, "\n")
	if !strings.Contains(joined, "tx=a") || strings.Contains(joined, "tx=b block=2") {
		t.Errorf("report did not apply sorted limit:\n%s", joined)
	}
	if !strings.Contains(joined, "... and 1 more missing records") {
		t.Errorf("report omitted truncation count:\n%s", joined)
	}
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
