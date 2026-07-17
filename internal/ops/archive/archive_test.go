package archive_test

import (
	"context"
	"database/sql"
	"errors"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/archive"
)

func TestResolveProjects(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{name: "both has documented order", input: "both", want: []string{archive.ProjectCosmicGame, archive.ProjectRandomWalk}},
		{name: "case insensitive both", input: "BoTh", want: []string{archive.ProjectCosmicGame, archive.ProjectRandomWalk}},
		{name: "randomwalk", input: "RANDOMWALK", want: []string{archive.ProjectRandomWalk}},
		{name: "cosmicgame", input: "cosmicgame", want: []string{archive.ProjectCosmicGame}},
		{name: "invalid", input: "augur", wantErr: true},
		{name: "whitespace remains invalid", input: " both ", wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := archive.ResolveProjects(test.input)
			if (err != nil) != test.wantErr {
				t.Fatalf("ResolveProjects(%q) error = %v, wantErr %v", test.input, err, test.wantErr)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("ResolveProjects(%q) = %v, want %v", test.input, got, test.want)
			}
		})
	}
}

func TestResumeFloor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		positions []archive.ResumePosition
		want      int64
	}{
		{name: "empty", want: 0},
		{
			name: "minimum contract watermark",
			positions: []archive.ResumePosition{
				{ContractAddress: "a", MaxEventID: 30},
				{ContractAddress: "b", MaxEventID: 4},
				{ContractAddress: "c", MaxEventID: 19},
			},
			want: 4,
		},
		{
			name: "unarchived contract forces zero",
			positions: []archive.ResumePosition{
				{ContractAddress: "a", MaxEventID: 30},
				{ContractAddress: "b", MaxEventID: 0},
			},
			want: 0,
		},
		{
			name:      "defensive negative watermark",
			positions: []archive.ResumePosition{{MaxEventID: -1}, {MaxEventID: 10}},
			want:      0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := archive.ResumeFloor(test.positions); got != test.want {
				t.Errorf("ResumeFloor() = %d, want %d", got, test.want)
			}
		})
	}
}

type fakeExporter struct {
	projects []string
	errFor   string
	cancel   context.CancelFunc
}

func (f *fakeExporter) ExportProject(_ context.Context, project string) (archive.ExportStats, error) {
	f.projects = append(f.projects, project)
	if f.cancel != nil && len(f.projects) == 1 {
		f.cancel()
	}
	if project == f.errFor {
		return archive.ExportStats{}, errors.New("export failed")
	}
	return archive.ExportStats{EventLogsProcessed: 1}, nil
}

func TestExportProjectsOrderCancellationAndErrors(t *testing.T) {
	t.Parallel()
	t.Run("ordered", func(t *testing.T) {
		exporter := &fakeExporter{}
		results, err := archive.ExportProjects(context.Background(), []string{"cg", "rw"}, exporter)
		if err != nil {
			t.Fatalf("ExportProjects() error = %v", err)
		}
		if !reflect.DeepEqual(exporter.projects, []string{"cg", "rw"}) {
			t.Errorf("calls = %v, want ordered projects", exporter.projects)
		}
		if len(results) != 2 || results[0].Project != "cg" || results[1].Project != "rw" {
			t.Errorf("results = %+v", results)
		}
	})
	t.Run("project error stops later projects", func(t *testing.T) {
		exporter := &fakeExporter{errFor: "bad"}
		results, err := archive.ExportProjects(context.Background(), []string{"ok", "bad", "never"}, exporter)
		if err == nil {
			t.Fatal("ExportProjects() error = nil")
		}
		if len(results) != 1 || !reflect.DeepEqual(exporter.projects, []string{"ok", "bad"}) {
			t.Errorf("partial results/calls = %+v / %v", results, exporter.projects)
		}
	})
	t.Run("cancellation stops later projects", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		exporter := &fakeExporter{cancel: cancel}
		results, err := archive.ExportProjects(ctx, []string{"first", "never"}, exporter)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context.Canceled", err)
		}
		if len(results) != 1 || !reflect.DeepEqual(exporter.projects, []string{"first"}) {
			t.Errorf("partial results/calls = %+v / %v", results, exporter.projects)
		}
	})
}

func TestMismatchStatsStrictness(t *testing.T) {
	t.Parallel()
	blocking := []func(*archive.MismatchStats){
		func(s *archive.MismatchStats) { s.EventLogsMissingFromArchive = 1 },
		func(s *archive.MismatchStats) { s.ArchiveEventLogOrphans = 1 },
		func(s *archive.MismatchStats) { s.EventLogDataMismatches = 1 },
		func(s *archive.MismatchStats) { s.ProjectTransactionsMissing = 1 },
		func(s *archive.MismatchStats) { s.ArchiveLogsMissingTx = 1 },
		func(s *archive.MismatchStats) { s.TxCoreMismatches = 1 },
		func(s *archive.MismatchStats) { s.ProjectBlocksMissing = 1 },
		func(s *archive.MismatchStats) { s.BlockHashMismatches = 1 },
	}
	for i, setMismatch := range blocking {
		var stats archive.MismatchStats
		setMismatch(&stats)
		if stats.Passed(archive.VerifyOptions{}) {
			t.Errorf("blocking mismatch category %d passed", i)
		}
	}

	warnings := archive.MismatchStats{
		TxNumLogsOnlyMismatches: 2,
		BlockMetadataMismatches: 3,
	}
	if !warnings.Passed(archive.VerifyOptions{}) {
		t.Error("non-strict metadata-only mismatches should pass")
	}
	if warnings.Passed(archive.VerifyOptions{StrictTxNumLogs: true}) {
		t.Error("strict tx num_logs should fail")
	}
	if warnings.Passed(archive.VerifyOptions{StrictBlockMetadata: true}) {
		t.Error("strict block metadata should fail")
	}
}

type fakeVerifier struct {
	calls  []string
	failAt string
}

func (f *fakeVerifier) VerifyProject(
	_ context.Context,
	project string,
	_ archive.VerifyOptions,
) (archive.ProjectVerification, error) {
	f.calls = append(f.calls, project)
	if project == f.failAt {
		return archive.ProjectVerification{}, errors.New("query failed")
	}
	return archive.ProjectVerification{Project: project, Passed: project != "mismatch"}, nil
}

func TestVerifyProjectsAggregatesAndPropagates(t *testing.T) {
	t.Parallel()
	verifier := &fakeVerifier{}
	report, err := archive.VerifyProjects(
		context.Background(),
		[]string{"pass", "mismatch", "pass-again"},
		archive.VerifyOptions{},
		verifier,
	)
	if err != nil {
		t.Fatalf("VerifyProjects() error = %v", err)
	}
	if report.Passed {
		t.Error("report passed despite one project mismatch")
	}
	if !reflect.DeepEqual(verifier.calls, []string{"pass", "mismatch", "pass-again"}) {
		t.Errorf("calls = %v", verifier.calls)
	}

	verifier = &fakeVerifier{failAt: "query-error"}
	_, err = archive.VerifyProjects(
		context.Background(),
		[]string{"pass", "query-error", "never"},
		archive.VerifyOptions{},
		verifier,
	)
	if err == nil || !reflect.DeepEqual(verifier.calls, []string{"pass", "query-error"}) {
		t.Errorf("error/calls = %v / %v", err, verifier.calls)
	}
}

func TestSelectStartBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		flag        uint64
		fromAddress sql.NullInt64
		fromEvent   sql.NullInt64
		want        uint64
		wantErr     bool
	}{
		{name: "flag wins", flag: 99, fromAddress: sql.NullInt64{Int64: 1, Valid: true}, want: 99},
		{
			name:        "minimum metadata",
			fromAddress: sql.NullInt64{Int64: 50, Valid: true},
			fromEvent:   sql.NullInt64{Int64: 20, Valid: true},
			want:        20,
		},
		{
			name:        "valid positive candidate",
			fromAddress: sql.NullInt64{Int64: 0, Valid: true},
			fromEvent:   sql.NullInt64{Int64: 7, Valid: true},
			want:        7,
		},
		{name: "no candidate", wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := archive.SelectStartBlock(test.flag, test.fromAddress, test.fromEvent)
			if (err != nil) != test.wantErr {
				t.Fatalf("SelectStartBlock() error = %v, wantErr %v", err, test.wantErr)
			}
			if got != test.want {
				t.Errorf("SelectStartBlock() = %d, want %d", got, test.want)
			}
		})
	}
}

func TestFillStatsMerge(t *testing.T) {
	t.Parallel()
	left := archive.FillStats{
		BlocksScanned: 1,
		LogsFromNode:  2,
		LogsSkipped:   3,
		LogsInserted:  4,
		TxInserted:    5,
	}
	right := archive.FillStats{
		BlocksScanned: 10,
		LogsFromNode:  20,
		LogsSkipped:   30,
		LogsInserted:  40,
		TxInserted:    50,
		TxSkipped:     60,
		BlockInserted: 70,
		BlockSkipped:  80,
		FilterRetries: 90,
		RPCErrors:     100,
		DBErrors:      110,
	}
	left.Merge(right)
	want := archive.FillStats{
		BlocksScanned: 11,
		LogsFromNode:  22,
		LogsSkipped:   33,
		LogsInserted:  44,
		TxInserted:    55,
		TxSkipped:     60,
		BlockInserted: 70,
		BlockSkipped:  80,
		FilterRetries: 90,
		RPCErrors:     100,
		DBErrors:      110,
	}
	if left != want {
		t.Errorf("merged stats = %+v, want %+v", left, want)
	}
}

type fakeNodeFillRepository struct {
	contracts archive.Contracts
	start     uint64
	writer    *fakeNodeFillWriter
	err       error
}

func (r *fakeNodeFillRepository) ProjectContracts(context.Context, string) (archive.Contracts, error) {
	if r.err != nil {
		return archive.Contracts{}, r.err
	}
	return r.contracts, nil
}

func (r *fakeNodeFillRepository) ResolveStartBlock(context.Context, archive.Contracts, uint64) (uint64, error) {
	return r.start, nil
}

func (*fakeNodeFillRepository) ArchivedBlockNumbers(
	context.Context,
	archive.Contracts,
	uint64,
	uint64,
) ([]int64, error) {
	return nil, nil
}

func (r *fakeNodeFillRepository) PrepareWriter(context.Context) (archive.NodeFillWriter, error) {
	return r.writer, nil
}

type fakeNodeFillWriter struct {
	eventExists    bool
	eventExistsErr error
	insertedEvents int
	eventInserted  bool
	eventInsertErr error
	closed         bool
	closeErr       error
}

func (w *fakeNodeFillWriter) EventLogExists(context.Context, string, int) (bool, error) {
	return w.eventExists, w.eventExistsErr
}

func (w *fakeNodeFillWriter) InsertEventLog(context.Context, archive.EventLog) (bool, error) {
	w.insertedEvents++
	return w.eventInserted, w.eventInsertErr
}

func (w *fakeNodeFillWriter) TransactionExists(context.Context, string) (bool, error) {
	return true, nil
}

func (w *fakeNodeFillWriter) InsertTransaction(context.Context, archive.Transaction) (bool, error) {
	return true, nil
}

func (w *fakeNodeFillWriter) BlockExists(context.Context, int64, string) (bool, error) {
	return true, nil
}

func (w *fakeNodeFillWriter) InsertBlock(
	context.Context,
	archive.Block,
	[]string,
	bool,
) (bool, error) {
	return true, nil
}

func (w *fakeNodeFillWriter) Close() error {
	w.closed = true
	return w.closeErr
}

type fakeNodeClient struct {
	logs      []types.Log
	filterErr error
}

func (c *fakeNodeClient) FilterLogs(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
	return c.logs, c.filterErr
}

func (*fakeNodeClient) TransactionByHash(context.Context, common.Hash) (*types.Transaction, bool, error) {
	return nil, false, errors.New("unexpected transaction request")
}

func (*fakeNodeClient) TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error) {
	return nil, errors.New("unexpected receipt request")
}

func (*fakeNodeClient) BlockByNumber(context.Context, *big.Int) (*types.Block, error) {
	return types.NewBlockWithHeader(&types.Header{
		Number:     big.NewInt(10),
		Time:       123,
		ParentHash: common.HexToHash("0x01"),
	}), nil
}

type fakeAddressStore struct{}

func (fakeAddressStore) LookupOrCreateAddress(context.Context, string, int64, int64) (int64, error) {
	return 1, nil
}

func testLog() types.Log {
	return types.Log{
		Address:     common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Topics:      []common.Hash{common.HexToHash("0x01")},
		BlockNumber: 10,
		TxHash:      common.HexToHash("0x02"),
		Index:       3,
	}
}

func TestNodeFillerDryRunAndIdempotentSkip(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		eventExists  bool
		dryRun       bool
		wantInserted int64
		wantSkipped  int64
		wantWrites   int
	}{
		{name: "dry run counts would insert", dryRun: true, wantInserted: 1},
		{name: "existing row is skipped", eventExists: true, wantSkipped: 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			writer := &fakeNodeFillWriter{eventExists: test.eventExists}
			filler := archive.NodeFiller{
				Repository: &fakeNodeFillRepository{
					contracts: archive.Contracts{Addresses: []string{testLog().Address.Hex()}},
					start:     10,
					writer:    writer,
				},
				AddressStore: fakeAddressStore{},
				Client:       &fakeNodeClient{logs: []types.Log{testLog()}},
			}
			stats, err := filler.RunProject(context.Background(), archive.ProjectRandomWalk, archive.NodeFillOptions{
				EndBlock:  10,
				BatchSize: 1,
				DryRun:    test.dryRun,
			})
			if err != nil {
				t.Fatalf("RunProject() error = %v", err)
			}
			if stats.LogsInserted != test.wantInserted || stats.LogsSkipped != test.wantSkipped {
				t.Errorf("stats = %+v", stats)
			}
			if writer.insertedEvents != test.wantWrites {
				t.Errorf("event writes = %d, want %d", writer.insertedEvents, test.wantWrites)
			}
			if !writer.closed {
				t.Error("writer was not closed")
			}
		})
	}
}

func TestNodeFillerCountsPerRowDBErrors(t *testing.T) {
	t.Parallel()
	writer := &fakeNodeFillWriter{eventExistsErr: errors.New("db unavailable")}
	filler := archive.NodeFiller{
		Repository: &fakeNodeFillRepository{
			contracts: archive.Contracts{Addresses: []string{testLog().Address.Hex()}},
			start:     10,
			writer:    writer,
		},
		AddressStore: fakeAddressStore{},
		Client:       &fakeNodeClient{logs: []types.Log{testLog()}},
	}
	stats, err := filler.RunProject(context.Background(), archive.ProjectRandomWalk, archive.NodeFillOptions{
		EndBlock:  10,
		BatchSize: 1,
	})
	if err != nil {
		t.Fatalf("RunProject() error = %v", err)
	}
	if stats.DBErrors != 1 || stats.RPCErrors != 0 || stats.LogsFromNode != 1 {
		t.Errorf("stats = %+v", stats)
	}
}

func TestNodeFillerReturnsWriterCloseError(t *testing.T) {
	t.Parallel()
	closeErr := errors.New("close statements")
	writer := &fakeNodeFillWriter{eventExists: true, closeErr: closeErr}
	filler := archive.NodeFiller{
		Repository: &fakeNodeFillRepository{
			contracts: archive.Contracts{Addresses: []string{testLog().Address.Hex()}},
			start:     10,
			writer:    writer,
		},
		AddressStore: fakeAddressStore{},
		Client:       &fakeNodeClient{logs: []types.Log{testLog()}},
	}
	_, err := filler.RunProject(context.Background(), archive.ProjectRandomWalk, archive.NodeFillOptions{
		EndBlock:  10,
		BatchSize: 1,
	})
	if !errors.Is(err, closeErr) {
		t.Fatalf("RunProject() error = %v, want close error", err)
	}
	if !writer.closed {
		t.Fatal("writer was not closed")
	}
}

func TestNodeFillerRPCErrorHonorsCancellation(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	writer := &fakeNodeFillWriter{}
	filler := archive.NodeFiller{
		Repository: &fakeNodeFillRepository{
			contracts: archive.Contracts{Addresses: []string{testLog().Address.Hex()}},
			start:     10,
			writer:    writer,
		},
		AddressStore: fakeAddressStore{},
		Client:       &fakeNodeClient{filterErr: errors.New("rpc unavailable")},
		RetryDelay:   time.Millisecond,
		Sleep: func(context.Context, time.Duration) error {
			cancel()
			return context.Canceled
		},
	}
	stats, err := filler.RunProject(ctx, archive.ProjectRandomWalk, archive.NodeFillOptions{
		EndBlock:  10,
		BatchSize: 1,
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("RunProject() error = %v, want context.Canceled", err)
	}
	if stats.FilterRetries != 1 || stats.RPCErrors != 0 {
		t.Errorf("filter retries / unresolved RPC errors = %d / %d, want 1 / 0", stats.FilterRetries, stats.RPCErrors)
	}
}
