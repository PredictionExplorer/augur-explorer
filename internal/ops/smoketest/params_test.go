package smoketest

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

type fakeParameterSource struct {
	params Params
	err    error
	load   func(context.Context)
}

type parameterTestConnector struct {
	query func(context.Context, string) (driver.Rows, error)
}

func (c parameterTestConnector) Connect(context.Context) (driver.Conn, error) {
	return parameterTestConn(c), nil
}

func (c parameterTestConnector) Driver() driver.Driver { return parameterTestDriver{connector: c} }

type parameterTestDriver struct{ connector parameterTestConnector }

func (d parameterTestDriver) Open(string) (driver.Conn, error) {
	return d.connector.Connect(context.Background())
}

type parameterTestConn struct {
	query func(context.Context, string) (driver.Rows, error)
}

func (c parameterTestConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("Prepare is not supported")
}
func (parameterTestConn) Close() error              { return nil }
func (parameterTestConn) Begin() (driver.Tx, error) { return nil, errors.New("Begin is not supported") }
func (c parameterTestConn) QueryContext(ctx context.Context, query string, _ []driver.NamedValue) (driver.Rows, error) {
	return c.query(ctx, query)
}

type parameterTestRows struct {
	value driver.Value
	empty bool
	read  bool
}

func (*parameterTestRows) Columns() []string { return []string{"value"} }
func (*parameterTestRows) Close() error      { return nil }
func (r *parameterTestRows) Next(dest []driver.Value) error {
	if r.empty || r.read {
		return io.EOF
	}
	r.read = true
	dest[0] = r.value
	return nil
}

func (f fakeParameterSource) Parameters(ctx context.Context) (Params, error) {
	if f.load != nil {
		f.load(ctx)
	}
	return f.params, f.err
}

func TestDefaultParamsAreComplete(t *testing.T) {
	t.Parallel()
	value := reflect.ValueOf(DefaultParams())
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).String() == "" {
			t.Errorf("default %s is empty", value.Type().Field(i).Name)
		}
	}
}

func TestWithDefaultsPreservesAndDerivesValues(t *testing.T) {
	t.Parallel()
	got := WithDefaults(Params{
		UserAddress: "0xUser",
		RoundNumber: "99",
		TokenName:   "named",
	})
	if got.CSTStakerAddress != "0xUser" {
		t.Fatalf("CST staker = %q", got.CSTStakerAddress)
	}
	if got.BidRound != "99" {
		t.Fatalf("bid round = %q", got.BidRound)
	}
	if got.TokenName != "named" {
		t.Fatalf("token name overwritten: %q", got.TokenName)
	}
	if got.NFTDonationID != "1" || got.TimestampMax != "2000000000" {
		t.Fatalf("defaults not filled: %#v", got)
	}
}

func TestSQLParameterSourceRejectsNilDB(t *testing.T) {
	t.Parallel()
	if _, err := (SQLParameterSource{}).Parameters(context.Background()); err == nil {
		t.Fatal("nil DB unexpectedly succeeded")
	}
}

func TestSQLParameterSourceLoadsValuesAndQueriesTokenNameColumn(t *testing.T) {
	t.Parallel()
	var queries []string
	db := sql.OpenDB(parameterTestConnector{query: func(_ context.Context, query string) (driver.Rows, error) {
		queries = append(queries, query)
		return &parameterTestRows{value: fmt.Sprintf("value-%02d", len(queries))}, nil
	}})
	t.Cleanup(func() { _ = db.Close() })
	got, err := (SQLParameterSource{DB: db}).Parameters(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got.TokenName != "value-09" || got.RoundNumber != "value-04" || got.ToDate != "value-20" {
		t.Fatalf("params = %#v", got)
	}
	if len(queries) != 20 {
		t.Fatalf("query count = %d", len(queries))
	}
	if !strings.Contains(queries[8], "SELECT token_name FROM cg_token_name") ||
		strings.Contains(queries[8], "SELECT name ") {
		t.Fatalf("token-name query = %q", queries[8])
	}
}

func TestSQLParameterSourceFallbacksAndCancellation(t *testing.T) {
	t.Parallel()
	t.Run("query errors use defaults", func(t *testing.T) {
		db := sql.OpenDB(parameterTestConnector{query: func(context.Context, string) (driver.Rows, error) {
			return nil, errors.New("missing table")
		}})
		defer func() { _ = db.Close() }()
		got, err := (SQLParameterSource{DB: db}).Parameters(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, DefaultParams()) {
			t.Fatalf("params = %#v", got)
		}
	})
	t.Run("null and empty rows use defaults", func(t *testing.T) {
		call := 0
		db := sql.OpenDB(parameterTestConnector{query: func(context.Context, string) (driver.Rows, error) {
			call++
			if call%2 == 0 {
				return &parameterTestRows{empty: true}, nil
			}
			return &parameterTestRows{value: nil}, nil
		}})
		defer func() { _ = db.Close() }()
		got, err := (SQLParameterSource{DB: db}).Parameters(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, DefaultParams()) {
			t.Fatalf("params = %#v", got)
		}
	})
	t.Run("cancellation is returned", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		db := sql.OpenDB(parameterTestConnector{query: func(context.Context, string) (driver.Rows, error) {
			t.Fatal("query callback should not run for canceled context")
			return nil, nil
		}})
		defer func() { _ = db.Close() }()
		_, err := (SQLParameterSource{DB: db}).Parameters(ctx)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v", err)
		}
	})
}

func TestFakeParameterSourceErrorContract(t *testing.T) {
	t.Parallel()
	want := errors.New("parameters unavailable")
	_, err := (fakeParameterSource{err: want}).Parameters(context.Background())
	if !errors.Is(err, want) {
		t.Fatalf("error = %v, want %v", err, want)
	}
}

func TestLoadParametersMapsEveryQuery(t *testing.T) {
	t.Parallel()
	var calls, fallbacks []string
	got, err := loadParameters(func(query, fallback string) (string, error) {
		calls = append(calls, query)
		fallbacks = append(fallbacks, fallback)
		return fmt.Sprintf("value-%02d", len(calls)), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	want := Params{
		UserAddress:        "value-01",
		CSTStakerAddress:   "value-02",
		RoundNumber:        "value-04",
		BidEventLogID:      "value-05",
		BidRound:           "value-06",
		BidPosition:        "value-07",
		TokenID:            "value-08",
		TokenName:          "value-09",
		ETHDonationID:      "value-10",
		NFTDonationID:      "value-11",
		ERC20DonationID:    "value-12",
		CSTActionID:        "value-13",
		RandomWalkActionID: "value-14",
		DepositID:          "value-15",
		NFTTokenAddress:    "value-16",
		TimestampMin:       "value-17",
		TimestampMax:       "value-18",
		FromDate:           "value-19",
		ToDate:             "value-20",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("params = %#v, want %#v", got, want)
	}
	if len(calls) != 20 {
		t.Fatalf("query count = %d", len(calls))
	}
	if fallbacks[1] != "value-01" {
		t.Fatalf("staker fallback = %q, want loaded user", fallbacks[1])
	}
	if fallbacks[3] != "value-03" {
		t.Fatalf("round fallback = %q, want latest bid round", fallbacks[3])
	}
	if fallbacks[5] != "value-04" {
		t.Fatalf("bid-round fallback = %q, want selected round", fallbacks[5])
	}
}

func TestLoadParametersFallbacksAndErrors(t *testing.T) {
	t.Parallel()
	t.Run("fallbacks", func(t *testing.T) {
		got, err := loadParameters(func(_ string, fallback string) (string, error) {
			return fallback, nil
		})
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, DefaultParams()) {
			t.Fatalf("fallback params = %#v", got)
		}
	})
	for failAt := 1; failAt <= 20; failAt++ {
		t.Run(fmt.Sprintf("stops on error %02d", failAt), func(t *testing.T) {
			t.Parallel()
			want := errors.New("query canceled")
			calls := 0
			_, err := loadParameters(func(_ string, fallback string) (string, error) {
				calls++
				if calls == failAt {
					return "", want
				}
				return fallback, nil
			})
			if !errors.Is(err, want) || calls != failAt {
				t.Fatalf("calls=%d error=%v", calls, err)
			}
		})
	}
}
