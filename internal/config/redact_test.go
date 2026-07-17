package config

import (
	"log/slog"
	"strings"
	"testing"
	"time"
)

func TestAttrsRendersEffectiveValuesAndRedactsSecrets(t *testing.T) {
	t.Parallel()
	cfg := struct {
		Plain     string        `env:"PLAIN"`
		Secret    string        `env:"SECRET" secret:"true"`
		Unset     string        `env:"UNSET_SECRET" secret:"true"`
		RPC       string        `env:"RPC" secret:"url"`
		B         bool          `env:"B" default:"true"`
		N         int64         `env:"N" default:"42"`
		F         float64       `env:"F" default:"2.5"`
		D         time.Duration `env:"D" default:"90s"`
		T         time.Time     `env:"T"`
		UnsetTime time.Time     `env:"UNSET_TIME"`
		L         []int64       `env:"L"`
	}{}
	err := Load(&cfg, mapEnv(map[string]string{
		"PLAIN":  "hello",
		"SECRET": "hunter2",
		"RPC":    "https://arb.example.com/v2/APIKEY123",
		"T":      "2027-01-01T00:00:00Z",
		"L":      "5,6",
	}))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	got := map[string]string{}
	for _, a := range Attrs(&cfg) {
		got[a.Key] = a.Value.String()
	}
	want := map[string]string{
		"PLAIN":        "hello",
		"SECRET":       "[set]",
		"UNSET_SECRET": "[unset]",
		"RPC":          "https://arb.example.com/[redacted]",
		"B":            "true",
		"N":            "42",
		"F":            "2.5",
		"D":            "1m30s",
		"T":            "2027-01-01T00:00:00Z",
		"UNSET_TIME":   "[unset]",
		"L":            "5,6",
	}
	for k, w := range want {
		if got[k] != w {
			t.Errorf("attr %s = %q, want %q", k, got[k], w)
		}
	}
	for _, a := range Attrs(&cfg) {
		if strings.Contains(a.Value.String(), "hunter2") || strings.Contains(a.Value.String(), "APIKEY123") {
			t.Errorf("secret leaked through attr %s=%s", a.Key, a.Value)
		}
	}
}

func TestRedactURL(t *testing.T) {
	t.Parallel()
	cases := map[string]string{ //nolint:gosec // G101: fake credentials proving the redaction behavior
		"":                                     "[unset]",
		"https://host.example.com":             "https://host.example.com",
		"https://host.example.com/v2/KEY":      "https://host.example.com/[redacted]",
		"https://host.example.com?apikey=KEY":  "https://host.example.com/[redacted]",
		"https://user:pass@host.example.com":   "https://host.example.com/[redacted]",
		"http://127.0.0.1:8545":                "http://127.0.0.1:8545",
		"not a url at all":                     "[set]",
		"postgres://u:hunter2@db:5432/gamedb":  "postgres://db:5432/[redacted]",
		"wss://relay.example.com/ws?token=abc": "wss://relay.example.com/[redacted]",
	}
	for in, want := range cases {
		if got := RedactURL(in); got != want {
			t.Errorf("RedactURL(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestAttrsPanicsOnBadArgument(t *testing.T) {
	t.Parallel()
	defer func() {
		if recover() == nil {
			t.Fatal("Attrs accepted a non-pointer argument")
		}
	}()
	_ = Attrs(struct{}{})
}

func TestAttrsPanicsOnUnsupportedFieldKind(t *testing.T) {
	t.Parallel()
	defer func() {
		if recover() == nil {
			t.Fatal("Attrs rendered an unsupported field kind")
		}
	}()
	var cfg struct {
		U uint `env:"U"`
	}
	_ = Attrs(&cfg)
}

func TestAttrsUsableWithSlog(t *testing.T) {
	t.Parallel()
	cfg := struct {
		A string `env:"A" default:"v"`
	}{}
	if err := Load(&cfg, mapEnv(nil)); err != nil {
		t.Fatalf("Load: %v", err)
	}
	var sb strings.Builder
	logger := slog.New(slog.NewTextHandler(&sb, nil))
	logger.LogAttrs(t.Context(), slog.LevelInfo, "effective configuration", Attrs(&cfg)...)
	if !strings.Contains(sb.String(), "A=v") {
		t.Errorf("log line missing attr: %s", sb.String())
	}
}
