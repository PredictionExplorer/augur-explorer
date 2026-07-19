package assets

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

type fakeTokenSource struct {
	tokens []Token
	err    error
	load   func(context.Context)
	schema string
}

// tokenQueryFake adapts a per-query callback to the pgx Querier seam.
type tokenQueryFake struct {
	query func(query string) (pgx.Rows, error)
}

func (f tokenQueryFake) Query(_ context.Context, query string, _ ...any) (pgx.Rows, error) {
	return f.query(query)
}

func tokenRows(finalErr error, values ...[]any) pgx.Rows {
	ops := testutil.PgxOp{Kind: "query", Rows: values, RowsErr: finalErr}
	script := testutil.NewScriptedPgx(ops)
	rows, err := script.Query(context.Background(), "")
	if err != nil {
		panic(err)
	}
	return rows
}

func (f *fakeTokenSource) TokenSeeds(ctx context.Context, schema string) ([]Token, error) {
	f.schema = schema
	if f.load != nil {
		f.load(ctx)
	}
	return append([]Token(nil), f.tokens...), f.err
}

func TestNormalizeSeed(t *testing.T) {
	t.Parallel()
	tests := map[string]string{
		"":             "",
		"  0XAbC123  ": "abc123",
		"abcdef":       "abcdef",
		"0x0xAB":       "0xab",
	}
	for input, want := range tests {
		if got := NormalizeSeed(input); got != want {
			t.Errorf("NormalizeSeed(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestSeedNameCandidates(t *testing.T) {
	t.Parallel()
	short := SeedNameCandidates(" 0xAb ")
	if len(short) != 2 {
		t.Fatalf("short candidate count = %d, want 2", len(short))
	}
	if short[0] != (SeedName{Name: "0xab"}) {
		t.Fatalf("raw candidate = %#v", short[0])
	}
	if !short[1].Padded || short[1].Name != "0x"+strings.Repeat("0", 62)+"ab" {
		t.Fatalf("padded candidate = %#v", short[1])
	}

	fullSeed := strings.Repeat("a", SeedHexLength)
	full := SeedNameCandidates(fullSeed)
	if !reflect.DeepEqual(full, []SeedName{{Name: "0x" + fullSeed}}) {
		t.Fatalf("full candidates = %#v", full)
	}
	if got := LeftPadSeed(strings.Repeat("b", SeedHexLength+1)); len(got) != SeedHexLength+1 {
		t.Fatalf("overlong seed was changed: %q", got)
	}
	for _, invalid := range []string{"", "../escape", "ab/cd", "xyz", strings.Repeat("a", SeedHexLength+1)} {
		if got := SeedNameCandidates(invalid); len(got) != 0 {
			t.Errorf("SeedNameCandidates(%q) = %#v, want no filesystem candidates", invalid, got)
		}
	}
}

func TestNormalizeTokensOrdersAndDeduplicates(t *testing.T) {
	t.Parallel()
	got := NormalizeTokens([]Token{
		{TokenID: 9, Seed: " 0xBB "},
		{TokenID: 4, Seed: ""},
		{TokenID: 3, Seed: "aa"},
		{TokenID: 2, Seed: "0xBB"},
		{TokenID: 1, Seed: "CC"},
	})
	want := []Token{
		{TokenID: 1, Seed: "cc"},
		{TokenID: 2, Seed: "bb"},
		{TokenID: 3, Seed: "aa"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("NormalizeTokens() = %#v, want %#v", got, want)
	}

	got = NormalizeTokens([]Token{{TokenID: 1, Seed: "b"}, {TokenID: 1, Seed: "a"}})
	want = []Token{{TokenID: 1, Seed: "a"}, {TokenID: 1, Seed: "b"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("same-id order = %#v, want %#v", got, want)
	}
}

func TestValidateSchema(t *testing.T) {
	t.Parallel()
	for _, schema := range []string{"public", "_private", "tenant_42", "schema$1"} {
		if err := ValidateSchema(schema); err != nil {
			t.Errorf("ValidateSchema(%q): %v", schema, err)
		}
	}
	for _, schema := range []string{"", "42tenant", "public; DROP TABLE cg_mint_event", "a.b", `"public"`, "two words"} {
		if err := ValidateSchema(schema); err == nil {
			t.Errorf("ValidateSchema(%q) unexpectedly succeeded", schema)
		}
	}
}

func TestSQLTokenSourceValidation(t *testing.T) {
	t.Parallel()
	_, err := (SQLTokenSource{}).TokenSeeds(context.Background(), "public")
	if err == nil || !strings.Contains(err.Error(), "nil") {
		t.Fatalf("nil DB error = %v", err)
	}

	_, err = (SQLTokenSource{}).TokenSeeds(context.Background(), "bad.schema")
	if err == nil || !strings.Contains(err.Error(), "invalid database schema") {
		t.Fatalf("invalid schema error = %v", err)
	}
}

func TestSQLTokenSourceLoadsNormalizesAndDeduplicates(t *testing.T) {
	t.Parallel()
	var gotQuery string
	db := tokenQueryFake{query: func(query string) (pgx.Rows, error) {
		gotQuery = query
		return tokenRows(nil,
			[]any{int64(2), " 0xBB "},
			[]any{int64(1), "AA"},
			[]any{int64(3), "bb"},
			[]any{int64(4), ""},
		), nil
	}}
	got, err := (SQLTokenSource{DB: db}).TokenSeeds(context.Background(), "tenant_1")
	if err != nil {
		t.Fatal(err)
	}
	want := []Token{{TokenID: 1, Seed: "aa"}, {TokenID: 2, Seed: "bb"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("tokens = %#v, want %#v", got, want)
	}
	if !strings.Contains(gotQuery, "FROM tenant_1.cg_mint_event") ||
		!strings.Contains(gotQuery, "ORDER BY token_id") {
		t.Fatalf("query = %q", gotQuery)
	}
}

func TestSQLTokenSourceDatabaseErrors(t *testing.T) {
	t.Parallel()
	t.Run("query error", func(t *testing.T) {
		t.Parallel()
		want := errors.New("query failed")
		db := tokenQueryFake{query: func(string) (pgx.Rows, error) {
			return nil, want
		}}
		_, err := (SQLTokenSource{DB: db}).TokenSeeds(context.Background(), "public")
		if !errors.Is(err, want) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("scan error", func(t *testing.T) {
		t.Parallel()
		db := tokenQueryFake{query: func(string) (pgx.Rows, error) {
			return tokenRows(nil, []any{"not-an-integer", "seed"}), nil
		}}
		if _, err := (SQLTokenSource{DB: db}).TokenSeeds(context.Background(), "public"); err == nil {
			t.Fatal("scan error was ignored")
		}
	})
	t.Run("rows error", func(t *testing.T) {
		t.Parallel()
		want := errors.New("row stream failed")
		db := tokenQueryFake{query: func(string) (pgx.Rows, error) {
			return tokenRows(want, []any{int64(1), "seed"}), nil
		}}
		_, err := (SQLTokenSource{DB: db}).TokenSeeds(context.Background(), "public")
		if !errors.Is(err, want) {
			t.Fatalf("error = %v", err)
		}
	})
}

func TestFakeTokenSourceErrorContract(t *testing.T) {
	t.Parallel()
	want := errors.New("load failed")
	source := &fakeTokenSource{err: want}
	_, err := source.TokenSeeds(context.Background(), "public")
	if !errors.Is(err, want) {
		t.Fatalf("error = %v, want %v", err, want)
	}
}
