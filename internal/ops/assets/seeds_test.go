package assets

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

type fakeTokenSource struct {
	tokens []Token
	err    error
	load   func(context.Context)
	schema string
}

type tokenTestConnector struct {
	query func(context.Context, string) (driver.Rows, error)
}

func (c tokenTestConnector) Connect(context.Context) (driver.Conn, error) {
	return tokenTestConn(c), nil
}

func (c tokenTestConnector) Driver() driver.Driver { return tokenTestDriver{connector: c} }

type tokenTestDriver struct{ connector tokenTestConnector }

func (d tokenTestDriver) Open(string) (driver.Conn, error) {
	return d.connector.Connect(context.Background())
}

type tokenTestConn struct {
	query func(context.Context, string) (driver.Rows, error)
}

func (c tokenTestConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("Prepare is not supported")
}
func (tokenTestConn) Close() error              { return nil }
func (tokenTestConn) Begin() (driver.Tx, error) { return nil, errors.New("Begin is not supported") }
func (c tokenTestConn) QueryContext(ctx context.Context, query string, _ []driver.NamedValue) (driver.Rows, error) {
	return c.query(ctx, query)
}

type tokenTestRows struct {
	values   [][]driver.Value
	index    int
	finalErr error
}

func (*tokenTestRows) Columns() []string { return []string{"token_id", "seed"} }
func (*tokenTestRows) Close() error      { return nil }
func (r *tokenTestRows) Next(dest []driver.Value) error {
	if r.index < len(r.values) {
		copy(dest, r.values[r.index])
		r.index++
		return nil
	}
	if r.finalErr != nil {
		err := r.finalErr
		r.finalErr = nil
		return err
	}
	return io.EOF
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
	db := sql.OpenDB(tokenTestConnector{query: func(_ context.Context, query string) (driver.Rows, error) {
		gotQuery = query
		return &tokenTestRows{values: [][]driver.Value{
			{int64(2), " 0xBB "},
			{int64(1), "AA"},
			{int64(3), "bb"},
			{int64(4), ""},
		}}, nil
	}})
	t.Cleanup(func() { _ = db.Close() })
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
		want := errors.New("query failed")
		db := sql.OpenDB(tokenTestConnector{query: func(context.Context, string) (driver.Rows, error) {
			return nil, want
		}})
		defer func() { _ = db.Close() }()
		_, err := (SQLTokenSource{DB: db}).TokenSeeds(context.Background(), "public")
		if !errors.Is(err, want) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("scan error", func(t *testing.T) {
		db := sql.OpenDB(tokenTestConnector{query: func(context.Context, string) (driver.Rows, error) {
			return &tokenTestRows{values: [][]driver.Value{{"not-an-integer", "seed"}}}, nil
		}})
		defer func() { _ = db.Close() }()
		if _, err := (SQLTokenSource{DB: db}).TokenSeeds(context.Background(), "public"); err == nil {
			t.Fatal("scan error was ignored")
		}
	})
	t.Run("rows error", func(t *testing.T) {
		want := errors.New("row stream failed")
		db := sql.OpenDB(tokenTestConnector{query: func(context.Context, string) (driver.Rows, error) {
			return &tokenTestRows{
				values:   [][]driver.Value{{int64(1), "seed"}},
				finalErr: want,
			}, nil
		}})
		defer func() { _ = db.Close() }()
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
