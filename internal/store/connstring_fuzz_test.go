// Fuzz target for escapeConnParam (MODERNIZATION.md §4.4): a value embedded in
// a single-quoted libpq connection-string parameter must never be able to
// break out of the quotes, and must round-trip through pgx's DSN parser.
package dbs

import (
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestEscapeConnParam(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"plain", "plain"},
		{"it's", `it\'s`},
		{`back\slash`, `back\\slash`},
		{`\`, `\\`},
		{"'", `\'`},
		{`\'`, `\\\'`},
		{`pass'word\`, `pass\'word\\`},
	}
	for _, tc := range cases {
		if got := escapeConnParam(tc.in); got != tc.want {
			t.Errorf("escapeConnParam(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// assertNoQuoteBreakout scans an escaped value as a libpq quoted-string parser
// would: a backslash consumes the next byte; a bare single quote would
// terminate the quoted value early (i.e. the input broke out of the quotes).
func assertNoQuoteBreakout(t *testing.T, original, escaped string) {
	t.Helper()
	i := 0
	for i < len(escaped) {
		switch escaped[i] {
		case '\\':
			if i+1 >= len(escaped) {
				t.Fatalf("escapeConnParam(%q) = %q: trailing lone backslash escapes the closing quote", original, escaped)
			}
			i += 2
		case '\'':
			t.Fatalf("escapeConnParam(%q) = %q: unescaped quote at byte %d breaks out of the quoted value", original, escaped, i)
		default:
			i++
		}
	}
}

func FuzzConnStringEscape(f *testing.F) {
	for _, seed := range []string{
		"", "plain", "it's", `back\slash`, `\`, "'", `\'`, `'; host=evil `,
		`pass'word\`, "unicode✓", " spaces and\ttabs ",
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, s string) {
		// Env vars (the only source of these values) can never contain NUL,
		// and multi-megabyte values are not a realistic configuration; the cap
		// keeps the pgx.ParseConfig round trip below the fuzz worker deadline.
		if strings.ContainsRune(s, 0) || len(s) > 4096 {
			t.Skip()
		}
		escaped := escapeConnParam(s)
		assertNoQuoteBreakout(t, s, escaped)

		// End to end: pgx must parse the quoted value back to exactly s, and
		// the crafted value must not have injected any other parameter.
		cfg, err := pgx.ParseConfig("user='" + escaped + "' host='h' port='5432'")
		if err != nil {
			t.Fatalf("pgx.ParseConfig failed for escaped %q (original %q): %v", escaped, s, err)
		}
		// pgx substitutes the OS username for an empty user value, so the
		// round-trip equality only holds for non-empty inputs.
		if s != "" && cfg.User != s {
			t.Fatalf("round-trip mismatch: original %q, parsed user %q", s, cfg.User)
		}
		if cfg.Host != "h" {
			t.Fatalf("parameter injection: host = %q after embedding %q", cfg.Host, s)
		}
	})
}
