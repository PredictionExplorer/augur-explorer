package smoketest

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
)

type fakeParameterSource struct {
	params Params
	err    error
	load   func(context.Context)
}

// queryRowFake adapts a per-query callback to the pgx Querier seam; it
// mirrors a live pool by answering context errors without reaching the
// callback.
type queryRowFake struct {
	query func(query string) pgx.Row
}

func (f queryRowFake) QueryRow(ctx context.Context, query string, _ ...any) pgx.Row {
	if err := ctx.Err(); err != nil {
		return staticRow{err: err}
	}
	return f.query(query)
}

// staticRow is a pgx.Row serving one scripted value (or error, or the
// no-rows sentinel when empty).
type staticRow struct {
	value any
	empty bool
	err   error
}

func (r staticRow) Scan(dest ...any) error {
	if r.err != nil {
		return r.err
	}
	if r.empty {
		return pgx.ErrNoRows
	}
	scanner, ok := dest[0].(sql.Scanner)
	if !ok {
		return fmt.Errorf("unexpected scan destination %T", dest[0])
	}
	return scanner.Scan(r.value)
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
	for i := range value.NumField() {
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
	db := queryRowFake{query: func(query string) pgx.Row {
		queries = append(queries, query)
		return staticRow{value: fmt.Sprintf("value-%02d", len(queries))}
	}}
	got, err := (SQLParameterSource{DB: db}).Parameters(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got.RandomWalkTokenID != "value-09" || got.TokenName != "value-10" ||
		got.RoundNumber != "value-04" || got.ToDate != "value-21" {
		t.Fatalf("params = %#v", got)
	}
	if len(queries) != 21 {
		t.Fatalf("query count = %d", len(queries))
	}
	if !strings.Contains(queries[8], "SELECT token_id FROM rw_mint_evt") {
		t.Fatalf("RandomWalk-token query = %q", queries[8])
	}
	if !strings.Contains(queries[9], "SELECT token_name FROM cg_token_name") ||
		strings.Contains(queries[9], "SELECT name ") {
		t.Fatalf("token-name query = %q", queries[9])
	}
}

func TestSQLParameterSourceFallbacksAndCancellation(t *testing.T) {
	t.Parallel()
	t.Run("query errors use defaults", func(t *testing.T) {
		t.Parallel()
		db := queryRowFake{query: func(string) pgx.Row {
			return staticRow{err: errors.New("missing table")}
		}}
		got, err := (SQLParameterSource{DB: db}).Parameters(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, DefaultParams()) {
			t.Fatalf("params = %#v", got)
		}
	})
	t.Run("null and empty rows use defaults", func(t *testing.T) {
		t.Parallel()
		call := 0
		db := queryRowFake{query: func(string) pgx.Row {
			call++
			if call%2 == 0 {
				return staticRow{empty: true}
			}
			return staticRow{value: nil}
		}}
		got, err := (SQLParameterSource{DB: db}).Parameters(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, DefaultParams()) {
			t.Fatalf("params = %#v", got)
		}
	})
	t.Run("cancellation is returned", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		db := queryRowFake{query: func(string) pgx.Row {
			t.Fatal("query callback should not run for canceled context")
			return staticRow{}
		}}
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
		RandomWalkTokenID:  "value-09",
		TokenName:          "value-10",
		ETHDonationID:      "value-11",
		NFTDonationID:      "value-12",
		ERC20DonationID:    "value-13",
		CSTActionID:        "value-14",
		RandomWalkActionID: "value-15",
		DepositID:          "value-16",
		NFTTokenAddress:    "value-17",
		TimestampMin:       "value-18",
		TimestampMax:       "value-19",
		FromDate:           "value-20",
		ToDate:             "value-21",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("params = %#v, want %#v", got, want)
	}
	if len(calls) != 21 {
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
	for failAt := 1; failAt <= 21; failAt++ {
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
