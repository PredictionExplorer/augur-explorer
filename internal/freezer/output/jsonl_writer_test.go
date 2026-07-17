package output

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/PredictionExplorer/augur-explorer/internal/freezer/decode"
)

func TestJSONLWriterCreateAndAppend(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "logs.jsonl")
	writer, err := NewJSONLWriter(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := writer.Write(testLogRecord(1)); err != nil {
		t.Fatal(err)
	}
	if err := writer.Flush(); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	appender, err := NewJSONLWriterWithMode(path, true)
	if err != nil {
		t.Fatal(err)
	}
	if err := appender.Write(testLogRecord(2)); err != nil {
		t.Fatal(err)
	}
	if err := appender.Close(); err != nil {
		t.Fatal(err)
	}

	// #nosec G304 -- path is inside this test's temporary directory.
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) != 2 ||
		!strings.Contains(lines[0], `"blockNumber":1`) ||
		!strings.Contains(lines[1], `"blockNumber":2`) {
		t.Fatalf("JSONL = %q", data)
	}
}

func TestCSVWriterCreateAndAppend(t *testing.T) {
	t.Parallel()
	path := filepath.Join(t.TempDir(), "logs.csv")
	writer, err := NewCSVWriter(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := writer.Write(testLogRecord(1)); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	appender, err := NewCSVWriterAppend(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := appender.Write(testLogRecord(2)); err != nil {
		t.Fatal(err)
	}
	if err := appender.Close(); err != nil {
		t.Fatal(err)
	}

	file, err := os.Open(path) // #nosec G304 -- path is inside t.TempDir.
	if err != nil {
		t.Fatal(err)
	}
	rows, err := csv.NewReader(file).ReadAll()
	closeErr := file.Close()
	if err != nil {
		t.Fatal(err)
	}
	if closeErr != nil {
		t.Fatal(closeErr)
	}
	if len(rows) != 3 || rows[0][0] != "blockNumber" ||
		rows[1][0] != "1" || rows[2][0] != "2" ||
		rows[1][6] != `["0x01","0x02"]` {
		t.Fatalf("CSV rows = %#v", rows)
	}
}

func TestWriterFactoryAndOpenFailures(t *testing.T) {
	t.Parallel()
	for _, format := range []string{"", "json", "jsonl", "csv"} {
		writer, err := NewWriter(format, "-")
		if err != nil {
			t.Fatalf("NewWriter(%q): %v", format, err)
		}
		if err := writer.Close(); err != nil {
			t.Fatalf("Close(%q): %v", format, err)
		}
	}
	if _, err := NewWriter("xml", "-"); err == nil {
		t.Fatal("unknown output format was accepted")
	}

	badPath := filepath.Join(t.TempDir(), "missing", "output")
	for name, open := range map[string]func() error{
		"JSON": func() error {
			_, err := NewJSONLWriter(badPath)
			return err
		},
		"JSON mode": func() error {
			_, err := NewJSONLWriterWithMode(badPath, false)
			return err
		},
		"CSV": func() error {
			_, err := NewCSVWriter(badPath)
			return err
		},
		"CSV append": func() error {
			_, err := NewCSVWriterAppend(badPath)
			return err
		},
	} {
		if err := open(); err == nil {
			t.Errorf("%s writer unexpectedly opened missing parent", name)
		}
	}
}

func TestWriterErrorPropagation(t *testing.T) {
	t.Parallel()
	writeErr := errors.New("write failed")
	closeErr := errors.New("close failed")
	jsonWriter := &JSONLWriter{
		w:       errorWriter{err: writeErr},
		encoder: json.NewEncoder(errorWriter{err: writeErr}),
		closer:  errorCloser{err: closeErr},
	}
	if err := jsonWriter.Write(testLogRecord(1)); !errors.Is(err, writeErr) {
		t.Fatalf("JSON write error = %v", err)
	}
	if err := jsonWriter.Close(); !errors.Is(err, closeErr) {
		t.Fatalf("JSON close error = %v", err)
	}

	csvWriter := &CSVWriter{
		w:      csv.NewWriter(errorWriter{err: writeErr}),
		closer: errorCloser{err: closeErr},
	}
	if err := csvWriter.Write(testLogRecord(1)); err != nil {
		t.Fatalf("CSV Write buffers and should not fail immediately: %v", err)
	}
	if err := csvWriter.Flush(); !errors.Is(err, writeErr) {
		t.Fatalf("CSV flush error = %v", err)
	}
	if err := csvWriter.Close(); !errors.Is(err, writeErr) {
		t.Fatalf("CSV close should return flush error first: %v", err)
	}
}

func TestJSONLWriterSerializesConcurrentWrites(t *testing.T) {
	t.Parallel()
	var buffer bytes.Buffer
	writer := &JSONLWriter{w: &buffer, encoder: json.NewEncoder(&buffer)}
	const records = 100
	var wait sync.WaitGroup
	for i := range uint64(records) {
		wait.Add(1)
		go func(index uint64) {
			defer wait.Done()
			if err := writer.Write(testLogRecord(index)); err != nil {
				t.Errorf("Write(%d): %v", index, err)
			}
		}(i)
	}
	wait.Wait()
	if got := strings.Count(strings.TrimSpace(buffer.String()), "\n") + 1; got != records {
		t.Fatalf("serialized records = %d, want %d", got, records)
	}
}

func TestLogConversions(t *testing.T) {
	t.Parallel()
	address := ethcommon.HexToAddress("0x2100000000000000000000000000000000000021")
	topics := []ethcommon.Hash{
		ethcommon.HexToHash("0x01"),
		ethcommon.HexToHash("0x02"),
	}
	data := []byte{0xde, 0xad, 0xbe, 0xef}
	decoded := &decode.DecodedLog{
		Address: address, Topics: topics, Data: data,
		ReceiptIndex: 3, LogIndex: 4, DataKeccak: crypto.Keccak256Hash(data),
	}
	record := DecodedLogToRecord(10, 2, decoded, true)
	assertConvertedRecord(t, record, 10, 2, 3, 4, address, topics, data)

	entry := &decode.Log{Address: address, Topics: topics, Data: data}
	record = LogEntryToRecord(11, 5, entry, true)
	assertConvertedRecord(t, record, 11, 5, 5, 5, address, topics, data)

	empty := DecodedLogToRecord(1, 0, &decode.DecodedLog{Address: address}, false)
	if empty.Topic0 != "" || empty.DataHex != "" ||
		empty.DataKeccak != (ethcommon.Hash{}).Hex() ||
		empty.Topics == nil {
		t.Fatalf("empty conversion = %+v", empty)
	}
}

func testLogRecord(block uint64) *LogRecord {
	return &LogRecord{
		BlockNumber: block,
		Contract:    "0x2100000000000000000000000000000000000021",
		Topic0:      "0x01",
		Topics:      []string{"0x01", "0x02"},
		DataKeccak:  "0x03",
		DataLen:     2,
		DataHex:     "0x0405",
	}
}

func assertConvertedRecord(
	t *testing.T,
	record *LogRecord,
	block uint64,
	txIndex, receiptIndex, logIndex uint,
	address ethcommon.Address,
	topics []ethcommon.Hash,
	data []byte,
) {
	t.Helper()
	if record.BlockNumber != block || record.TxIndex != txIndex ||
		record.ReceiptIndex != receiptIndex || record.LogIndex != logIndex ||
		record.Contract != address.Hex() || record.Topic0 != topics[0].Hex() ||
		len(record.Topics) != len(topics) ||
		record.DataKeccak != crypto.Keccak256Hash(data).Hex() ||
		record.DataHex != "0xdeadbeef" || record.DataLen != len(data) {
		t.Fatalf("converted record = %+v", record)
	}
}

type errorWriter struct{ err error }

func (writer errorWriter) Write([]byte) (int, error) { return 0, writer.err }

type errorCloser struct{ err error }

func (closer errorCloser) Close() error { return closer.err }
