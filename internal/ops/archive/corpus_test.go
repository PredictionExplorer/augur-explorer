package archive

import (
	"bytes"
	"context"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/rlpcorpus"
	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

const (
	corpusTxA = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	corpusTxB = "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
)

var errCorpusOutput = errors.New("corpus output failed")

type corpusFailingWriter struct{}

func (corpusFailingWriter) Write([]byte) (int, error) { return 0, errCorpusOutput }

func corpusRow(t *testing.T, block, event, index int64, txHash string, data byte) ([]any, []byte) {
	t.Helper()
	address := common.HexToAddress("0x2000000000000000000000000000000000000002")
	topic := common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	raw, err := toolutil.EncodeLogRLP(&types.Log{
		Address: address,
		Topics:  []common.Hash{topic},
		Data:    []byte{data},
	})
	if err != nil {
		t.Fatal(err)
	}
	return []any{
		block,
		event,
		index,
		txHash,
		address.Hex(),
		toolutil.Topic0Sig(&types.Log{Topics: []common.Hash{topic}}),
		raw,
	}, raw
}

func TestExportCorpusCompleteTransactionsAndExactBytes(t *testing.T) {
	rowA0, rawA0 := corpusRow(t, 100, 10, 0, corpusTxA, 0x01)
	rowA1, rawA1 := corpusRow(t, 100, 11, 1, corpusTxA, 0x02)
	rowB0, rawB0 := corpusRow(t, 101, 20, 0, corpusTxB, 0x03)
	db := openScriptDB(t,
		queryOp("FROM arch_evtlog", rowA0, rowA1),
		queryOp("FROM arch_evtlog", rowB0),
	)

	var out bytes.Buffer
	stats, err := ExportCorpus(
		context.Background(),
		db,
		CorpusExportOptions{
			Project:  ProjectCosmicGame,
			TxHashes: []string{"0x" + strings.ToUpper(corpusTxA[2:]), corpusTxB},
		},
		&out,
	)
	if err != nil {
		t.Fatalf("ExportCorpus: %v", err)
	}
	if stats != (CorpusExportStats{Transactions: 2, EventLogs: 3}) {
		t.Fatalf("stats = %+v", stats)
	}

	entries, err := rlpcorpus.Load(&out)
	if err != nil {
		t.Fatalf("Load exported corpus: %v", err)
	}
	if got := []string{entries[0].TxHash, entries[1].TxHash, entries[2].TxHash}; !reflect.DeepEqual(
		got,
		[]string{corpusTxA, corpusTxA, corpusTxB},
	) {
		t.Fatalf("transaction order = %v", got)
	}
	for i, want := range [][]byte{rawA0, rawA1, rawB0} {
		got, err := entries[i].RLPBytes()
		if err != nil {
			t.Fatalf("entry %d RLP: %v", i, err)
		}
		if !bytes.Equal(got, want) {
			t.Errorf("entry %d RLP changed", i)
		}
	}
}

func TestExportCorpusInputValidation(t *testing.T) {
	t.Parallel()
	validDB := testutil.NewScriptedPgx()
	validOptions := CorpusExportOptions{Project: ProjectRandomWalk, TxHashes: []string{corpusTxA}}
	tests := []struct {
		name    string
		db      Querier
		options CorpusExportOptions
		writer  io.Writer
		want    string
	}{
		{"database", nil, validOptions, io.Discard, "database"},
		{"writer", validDB, validOptions, nil, "writer"},
		{"project", validDB, CorpusExportOptions{Project: "other", TxHashes: []string{corpusTxA}}, io.Discard, "invalid project"},
		{"both projects", validDB, CorpusExportOptions{Project: "both", TxHashes: []string{corpusTxA}}, io.Discard, "exactly one"},
		{"no hashes", validDB, CorpusExportOptions{Project: ProjectRandomWalk}, io.Discard, "at least one"},
		{"hash prefix", validDB, CorpusExportOptions{Project: ProjectRandomWalk, TxHashes: []string{"aa"}}, io.Discard, "0x-prefixed"},
		{"hash length", validDB, CorpusExportOptions{Project: ProjectRandomWalk, TxHashes: []string{"0xaa"}}, io.Discard, "32-byte"},
		{"hash hex", validDB, CorpusExportOptions{Project: ProjectRandomWalk, TxHashes: []string{"0x" + strings.Repeat("z", 64)}}, io.Discard, "32-byte"},
		{"duplicate", validDB, CorpusExportOptions{Project: ProjectRandomWalk, TxHashes: []string{corpusTxA, "0x" + strings.ToUpper(corpusTxA[2:])}}, io.Discard, "duplicate"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if _, err := ExportCorpus(context.Background(), test.db, test.options, test.writer); err == nil ||
				!strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func TestExportCorpusDatabaseValidationAndWriteFailures(t *testing.T) {
	sentinel := errors.New("archive read failed")
	validRow, _ := corpusRow(t, 100, 10, 0, corpusTxA, 0x01)
	tests := []struct {
		name   string
		op     scriptOp
		ctx    func() context.Context
		writer io.Writer
		want   error
		text   string
	}{
		{
			name: "query",
			op:   queryErrorOp("FROM arch_evtlog", sentinel),
			ctx:  context.Background,
			want: sentinel,
		},
		{
			name: "iteration",
			op:   queryIterErrorOp("FROM arch_evtlog", sentinel),
			ctx:  context.Background,
			want: sentinel,
		},
		{
			name: "scan",
			op: queryOp("FROM arch_evtlog", []any{
				"not a block", int64(10), int64(0), corpusTxA, "0x0", "", []byte{0x01},
			}),
			ctx:  context.Background,
			text: "scan",
		},
		{
			name: "invalid row",
			op: queryOp("FROM arch_evtlog", []any{
				int64(100), int64(10), int64(0), corpusTxA,
				"0x2000000000000000000000000000000000000002", "aaaaaaaa",
				[]byte{0xff},
			}),
			ctx:  context.Background,
			text: "validate",
		},
		{
			name: "wrong transaction",
			op: queryOp("FROM arch_evtlog", func() []any {
				row := append([]any(nil), validRow...)
				row[3] = corpusTxB
				return row
			}()),
			ctx:  context.Background,
			text: "returned row",
		},
		{
			name:   "writer",
			op:     queryOp("FROM arch_evtlog", validRow),
			ctx:    context.Background,
			writer: corpusFailingWriter{},
			want:   errCorpusOutput,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := test.writer
			if writer == nil {
				writer = io.Discard
			}
			_, err := ExportCorpus(
				test.ctx(),
				openScriptDB(t, test.op),
				CorpusExportOptions{Project: ProjectRandomWalk, TxHashes: []string{corpusTxA}},
				writer,
			)
			if test.want != nil && !errors.Is(err, test.want) {
				t.Fatalf("error = %v, want %v", err, test.want)
			}
			if test.text != "" && (err == nil || !strings.Contains(err.Error(), test.text)) {
				t.Fatalf("error = %v, want containing %q", err, test.text)
			}
		})
	}

	t.Run("missing transaction", func(t *testing.T) {
		_, err := ExportCorpus(
			context.Background(),
			openScriptDB(t, queryOp("FROM arch_evtlog")),
			CorpusExportOptions{Project: ProjectRandomWalk, TxHashes: []string{corpusTxA}},
			io.Discard,
		)
		if err == nil || !strings.Contains(err.Error(), "has no event logs") {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := ExportCorpus(
			ctx,
			openScriptDB(t),
			CorpusExportOptions{Project: ProjectRandomWalk, TxHashes: []string{corpusTxA}},
			io.Discard,
		)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context.Canceled", err)
		}
	})
}
