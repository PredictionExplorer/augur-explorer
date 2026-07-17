package verify

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func TestLoadJSONL(t *testing.T) {
	t.Parallel()

	t.Run("valid records with normalization", func(t *testing.T) {
		t.Parallel()
		input := strings.Join([]string{
			`{"blockNumber":100,"contract":"0xABCDEF0000000000000000000000000000000001","topic0":"0xDDF252AD1BE2C89B69C2B068FC378DAA952BA7F163C4A11628F55A4DF523B3EF"}`,
			``,
			`{"blockNumber":100,"contract":"0xabcdef0000000000000000000000000000000001","topic0":"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"}`,
			`{"blockNumber":101,"contract":"0xabcdef0000000000000000000000000000000002","topic0":"0x8c5be1e5"}`,
		}, "\n") + "\n"

		events, contracts, err := LoadJSONL(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		// The two case-variant records collapse into one key with count 2.
		key := Key{BlockNum: 100, Topic0: "ddf252ad", Contract: "0xabcdef0000000000000000000000000000000001"}
		if events[key] != 2 {
			t.Fatalf("events = %v", events)
		}
		// Short topic0 (0x + exactly 8 chars): 0x stripped, 8 chars kept.
		short := Key{BlockNum: 101, Topic0: "8c5be1e5", Contract: "0xabcdef0000000000000000000000000000000002"}
		if events[short] != 1 {
			t.Fatalf("events = %v", events)
		}
		if len(contracts) != 2 {
			t.Fatalf("contracts = %v", contracts)
		}
	})

	t.Run("topic0 shorter than 10 chars kept as-is", func(t *testing.T) {
		t.Parallel()
		events, _, err := LoadJSONL(strings.NewReader(`{"blockNumber":5,"contract":"0xc","topic0":"0xABCD"}` + "\n"))
		if err != nil {
			t.Fatal(err)
		}
		key := Key{BlockNum: 5, Topic0: "0xabcd", Contract: "0xc"}
		if events[key] != 1 {
			t.Fatalf("events = %v", events)
		}
	})

	t.Run("invalid JSON reports line number", func(t *testing.T) {
		t.Parallel()
		input := `{"blockNumber":1,"contract":"0xa","topic0":"0x12345678aa"}` + "\n" + `{broken` + "\n"
		_, _, err := LoadJSONL(strings.NewReader(input))
		if err == nil || !strings.Contains(err.Error(), "line 2: invalid JSON") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("long lines within buffer", func(t *testing.T) {
		t.Parallel()
		dataHex := strings.Repeat("ab", 100_000) // 200KB line
		input := `{"blockNumber":1,"contract":"0xa","topic0":"0x1234567890","dataHex":"` + dataHex + `"}` + "\n"
		events, _, err := LoadJSONL(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if len(events) != 1 {
			t.Fatalf("events = %v", events)
		}
	})

	t.Run("line exceeding max buffer errors", func(t *testing.T) {
		t.Parallel()
		dataHex := strings.Repeat("ab", 600_000) // 1.2MB line
		input := `{"blockNumber":1,"contract":"0xa","topic0":"0x12","dataHex":"` + dataHex + `"}` + "\n"
		_, _, err := LoadJSONL(strings.NewReader(input))
		if err == nil || !strings.Contains(err.Error(), "scanner error") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		t.Parallel()
		events, contracts, err := LoadJSONL(strings.NewReader(""))
		if err != nil || len(events) != 0 || len(contracts) != 0 {
			t.Fatalf("events=%v contracts=%v err=%v", events, contracts, err)
		}
	})
}

func TestBlockRange(t *testing.T) {
	t.Parallel()
	// Map iteration order is randomized; repeat so both the "lower than
	// min" and "higher than max" branches are exercised regardless of
	// which key comes first.
	for range 32 {
		events := map[Key]int{
			{BlockNum: 50, Topic0: "a", Contract: "c"}:  1,
			{BlockNum: 7, Topic0: "b", Contract: "c"}:   1,
			{BlockNum: 900, Topic0: "a", Contract: "d"}: 1,
		}
		minBlock, maxBlock, ok := BlockRange(events)
		if !ok || minBlock != 7 || maxBlock != 900 {
			t.Fatalf("range = %d-%d ok=%v", minBlock, maxBlock, ok)
		}
	}

	if _, _, ok := BlockRange(nil); ok {
		t.Fatal("empty map must report ok=false")
	}
}

func TestValidTableName(t *testing.T) {
	t.Parallel()
	for name, want := range map[string]bool{
		"evt_log":        true,
		"_private":       true,
		"Evt2":           true,
		"1bad":           false,
		"bad-name":       false,
		"drop table; --": false,
		"":               false,
	} {
		if got := ValidTableName(name); got != want {
			t.Fatalf("ValidTableName(%q) = %v, want %v", name, got, want)
		}
	}
}

func key(block int64, topic, contract string) Key {
	return Key{BlockNum: block, Topic0: topic, Contract: contract}
}

func TestCompare(t *testing.T) {
	t.Parallel()
	freezer := map[Key]int{
		key(1, "aaaa", "0xc1"): 2, // match
		key(2, "bbbb", "0xc1"): 3, // count mismatch (db has 1)
		key(3, "cccc", "0xc2"): 1, // extra (not in db)
	}
	db := map[Key]int{
		key(1, "aaaa", "0xc1"): 2,
		key(2, "bbbb", "0xc1"): 1,
		key(4, "dddd", "0xc2"): 5, // missing from freezer
	}

	r := Compare(freezer, db, 10, 10)
	if r.FreezerKeys != 3 || r.DBKeys != 3 {
		t.Fatalf("report = %+v", r)
	}
	if r.Match != 1 || r.CountMismatch != 1 || r.Missing != 1 || r.Extra != 1 {
		t.Fatalf("report = %+v", r)
	}
	if len(r.MissingDetails) != 1 || !strings.Contains(r.MissingDetails[0], "block=4 topic0=dddd contract=0xc2 (db count=5)") {
		t.Fatalf("missing details = %v", r.MissingDetails)
	}
	if len(r.ExtraDetails) != 1 || !strings.Contains(r.ExtraDetails[0], "block=3 topic0=cccc contract=0xc2 (freezer count=1)") {
		t.Fatalf("extra details = %v", r.ExtraDetails)
	}
	if len(r.MismatchDetails) != 1 || !strings.Contains(r.MismatchDetails[0], "db=1 freezer=3") {
		t.Fatalf("mismatch details = %v", r.MismatchDetails)
	}
}

func TestCompareDetailCaps(t *testing.T) {
	t.Parallel()
	db := make(map[Key]int)
	freezer := make(map[Key]int)
	for i := range int64(20) {
		db[key(i, "aaaa", "0xc")] = 1          // all missing from freezer
		freezer[key(100+i, "bbbb", "0xc")] = 1 // all extra
	}

	r := Compare(freezer, db, 3, 5)
	if r.Missing != 20 || len(r.MissingDetails) != 3 {
		t.Fatalf("missing = %d details = %d", r.Missing, len(r.MissingDetails))
	}
	if r.Extra != 20 || len(r.ExtraDetails) != 5 {
		t.Fatalf("extra = %d details = %d", r.Extra, len(r.ExtraDetails))
	}
	// Details are deterministic: lowest blocks first.
	if !strings.Contains(r.MissingDetails[0], "block=0 ") {
		t.Fatalf("details = %v", r.MissingDetails)
	}
}

func TestCompareDetailOrderTieBreakers(t *testing.T) {
	t.Parallel()
	// Same block: topic breaks the tie; same block and topic: contract does.
	db := map[Key]int{
		key(1, "bbbb", "0xc1"): 1,
		key(1, "aaaa", "0xc2"): 1,
		key(1, "aaaa", "0xc1"): 1,
	}
	r := Compare(map[Key]int{}, db, 10, 10)
	if len(r.MissingDetails) != 3 {
		t.Fatalf("details = %v", r.MissingDetails)
	}
	if !strings.Contains(r.MissingDetails[0], "topic0=aaaa contract=0xc1") ||
		!strings.Contains(r.MissingDetails[1], "topic0=aaaa contract=0xc2") ||
		!strings.Contains(r.MissingDetails[2], "topic0=bbbb contract=0xc1") {
		t.Fatalf("details = %v", r.MissingDetails)
	}
}

func TestMatchRateAndPassed(t *testing.T) {
	t.Parallel()
	cases := []struct {
		match, dbKeys int
		wantRate      float64
		wantPassed    bool
	}{
		{match: 100, dbKeys: 100, wantRate: 1, wantPassed: true},
		{match: 99, dbKeys: 100, wantRate: 0.99, wantPassed: true},
		{match: 98, dbKeys: 100, wantRate: 0.98, wantPassed: false},
		{match: 0, dbKeys: 0, wantRate: 1, wantPassed: true},
	}
	for _, tc := range cases {
		r := Report{Match: tc.match, DBKeys: tc.dbKeys}
		if got := r.MatchRate(); got != tc.wantRate {
			t.Fatalf("MatchRate() = %v, want %v", got, tc.wantRate)
		}
		if got := r.Passed(); got != tc.wantPassed {
			t.Fatalf("Passed() = %v, want %v", got, tc.wantPassed)
		}
	}
}

// fakeRows serves scripted rows and a scripted iteration error through the
// pgx.Rows interface, covering the scan/err branches a healthy PostgreSQL
// cannot produce.
type fakeRows struct {
	rows    [][]any
	current int
	scanErr error
	err     error
}

func (r *fakeRows) Close()                                       {}
func (r *fakeRows) Err() error                                   { return r.err }
func (r *fakeRows) CommandTag() pgconn.CommandTag                { return pgconn.CommandTag{} }
func (r *fakeRows) FieldDescriptions() []pgconn.FieldDescription { return nil }
func (r *fakeRows) Values() ([]any, error)                       { return nil, nil }
func (r *fakeRows) RawValues() [][]byte                          { return nil }
func (r *fakeRows) Conn() *pgx.Conn                              { return nil }
func (r *fakeRows) Next() bool                                   { return r.current < len(r.rows) }

func (r *fakeRows) Scan(dest ...any) error {
	if r.scanErr != nil {
		return r.scanErr
	}
	row := r.rows[r.current]
	r.current++
	for i := range dest {
		switch d := dest[i].(type) {
		case *int64:
			*d = row[i].(int64)
		case *string:
			*d = row[i].(string)
		case *int:
			*d = row[i].(int)
		}
	}
	return nil
}

// fakeQuerier returns scripted rows or an error for every query.
type fakeQuerier struct {
	rows *fakeRows
	err  error
}

func (q fakeQuerier) Query(context.Context, string, ...any) (pgx.Rows, error) {
	if q.err != nil {
		return nil, q.err
	}
	return q.rows, nil
}

func TestContractAidsRowFailures(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	if _, err := ContractAids(ctx, fakeQuerier{rows: &fakeRows{
		rows:    [][]any{{int64(1)}},
		scanErr: errors.New("scan blew up"),
	}}, []string{"0xa"}); err == nil || !strings.Contains(err.Error(), "scan blew up") {
		t.Fatalf("err = %v", err)
	}

	if _, err := ContractAids(ctx, fakeQuerier{rows: &fakeRows{
		err: errors.New("stream cut"),
	}}, []string{"0xa"}); err == nil || !strings.Contains(err.Error(), "stream cut") {
		t.Fatalf("err = %v", err)
	}
}

func TestDBEventsRowFailures(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	if _, err := DBEvents(ctx, fakeQuerier{rows: &fakeRows{
		rows:    [][]any{{int64(1), "aaaa", "0xc", 1}},
		scanErr: errors.New("scan blew up"),
	}}, "evt_log", 0, 10, []int64{1}); err == nil || !strings.Contains(err.Error(), "scan blew up") {
		t.Fatalf("err = %v", err)
	}

	if _, err := DBEvents(ctx, fakeQuerier{rows: &fakeRows{
		err: errors.New("stream cut"),
	}}, "evt_log", 0, 10, []int64{1}); err == nil || !strings.Contains(err.Error(), "stream cut") {
		t.Fatalf("err = %v", err)
	}

	// A scripted happy row proves the normalization (trim + lowercase).
	events, err := DBEvents(ctx, fakeQuerier{rows: &fakeRows{
		rows: [][]any{{int64(9), "AAAA    ", "0xC", 4}},
	}}, "evt_log", 0, 10, []int64{1})
	if err != nil {
		t.Fatal(err)
	}
	if events[Key{BlockNum: 9, Topic0: "aaaa", Contract: "0xc"}] != 4 {
		t.Fatalf("events = %v", events)
	}
}

func TestReportWrite(t *testing.T) {
	t.Parallel()
	r := Report{
		FreezerKeys:     3,
		DBKeys:          4,
		Match:           2,
		CountMismatch:   1,
		Missing:         12,
		Extra:           7,
		MissingDetails:  []string{"m1", "m2"},
		ExtraDetails:    []string{"e1"},
		MismatchDetails: []string{"mm1"},
	}

	var quiet strings.Builder
	r.Write(&quiet, false)
	out := quiet.String()
	for _, want := range []string{
		"=== Verification Results ===",
		"Freezer events:  3 distinct (block, topic0, contract)",
		"Database events: 4 distinct (block, topic0, contract)",
		"Matching:          2",
		"Missing (in DB, not in freezer): 12",
		"Extra (in freezer, not in DB):   7",
		"Match rate: 50.00%",
		"Missing from freezer (sample):",
		"  ... and 10 more",
		"Extra in freezer (sample):",
		"  ... and 6 more",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("output missing %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "Count mismatches:") {
		t.Fatalf("quiet output must omit mismatches:\n%s", out)
	}

	var verbose strings.Builder
	r.Write(&verbose, true)
	if !strings.Contains(verbose.String(), "Count mismatches:") || !strings.Contains(verbose.String(), "  mm1") {
		t.Fatalf("verbose output = %s", verbose.String())
	}

	// Zero DB keys: no match rate line.
	var empty strings.Builder
	Report{}.Write(&empty, false)
	if strings.Contains(empty.String(), "Match rate") {
		t.Fatalf("empty report output = %s", empty.String())
	}
}
