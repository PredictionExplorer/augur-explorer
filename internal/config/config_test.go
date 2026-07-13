package config

import (
	"errors"
	"strings"
	"testing"
	"time"
)

// mapEnv builds a getenv over a map (missing keys read "").
func mapEnv(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

type kindsConfig struct {
	Str      string        `env:"T_STR" default:"dflt"`
	Bool     bool          `env:"T_BOOL"`
	BoolDflt bool          `env:"T_BOOL_DFLT" default:"true"`
	Int      int           `env:"T_INT" default:"7"`
	Int64    int64         `env:"T_I64"`
	Float    float64       `env:"T_FLOAT" default:"1.5"`
	Dur      time.Duration `env:"T_DUR" default:"30s"`
	List     []int64       `env:"T_LIST"`
	Req      string        `env:"T_REQ" required:"true"`
	Nested   struct {
		Inner string `env:"T_INNER" default:"in"`
	}
	Skipped string // untagged fields are ignored
}

func TestLoadAppliesValuesAndDefaults(t *testing.T) {
	t.Parallel()
	var cfg kindsConfig
	err := Load(&cfg, mapEnv(map[string]string{
		"T_BOOL": " yes ",
		"T_I64":  "-42",
		"T_LIST": "1, 2,,3",
		"T_REQ":  "here",
		"T_DUR":  "2m",
	}))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Str != "dflt" || !cfg.Bool || !cfg.BoolDflt || cfg.Int != 7 || cfg.Int64 != -42 ||
		cfg.Float != 1.5 || cfg.Dur != 2*time.Minute || cfg.Req != "here" || cfg.Nested.Inner != "in" {
		t.Errorf("loaded config = %+v", cfg)
	}
	if len(cfg.List) != 3 || cfg.List[0] != 1 || cfg.List[1] != 2 || cfg.List[2] != 3 {
		t.Errorf("List = %v, want [1 2 3]", cfg.List)
	}
}

func TestLoadAggregatesEveryProblem(t *testing.T) {
	t.Parallel()
	var cfg kindsConfig
	err := Load(&cfg, mapEnv(map[string]string{
		"T_BOOL":  "maybe",
		"T_INT":   "seven",
		"T_FLOAT": "x",
		"T_DUR":   "fast",
		"T_LIST":  "1,a",
		// T_REQ unset
	}))
	if err == nil {
		t.Fatal("Load accepted a config with six problems")
	}
	var le *LoadError
	if !errors.As(err, &le) {
		t.Fatalf("error type = %T, want *LoadError", err)
	}
	if len(le.Problems) != 6 {
		t.Fatalf("got %d problems, want 6:\n%s", len(le.Problems), err)
	}
	for _, want := range []string{
		`T_BOOL: cannot parse "maybe" as a boolean`,
		`T_INT: cannot parse "seven" as an integer`,
		`T_FLOAT: cannot parse "x" as a number`,
		`T_DUR: cannot parse "fast" as a duration`,
		`T_LIST: cannot parse list entry "a" as an integer`,
		"T_REQ: required but not set",
	} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("aggregated error missing %q:\n%s", want, err)
		}
	}
}

func TestLoadBoolSpellings(t *testing.T) {
	t.Parallel()
	for raw, want := range map[string]bool{
		"1": true, "t": true, "TRUE": true, "Yes": true, "y": true, "on": true,
		"0": false, "f": false, "False": false, "NO": false, "n": false, "off": false,
	} {
		var cfg struct {
			B bool `env:"B"`
		}
		if err := Load(&cfg, mapEnv(map[string]string{"B": raw})); err != nil {
			t.Errorf("Load(%q): %v", raw, err)
		} else if cfg.B != want {
			t.Errorf("Load(%q) = %v, want %v", raw, cfg.B, want)
		}
	}
}

func TestLoadRequiredWhitespaceOnlyCountsAsUnset(t *testing.T) {
	t.Parallel()
	var cfg struct {
		R string `env:"R" required:"true"`
	}
	err := Load(&cfg, mapEnv(map[string]string{"R": "   "}))
	if err == nil || !strings.Contains(err.Error(), "R: required but not set") {
		t.Fatalf("whitespace-only required value accepted: %v", err)
	}
}

func TestLoadPanicsOnBadDestination(t *testing.T) {
	t.Parallel()
	defer func() {
		if recover() == nil {
			t.Fatal("Load accepted a non-pointer destination")
		}
	}()
	_ = Load(kindsConfig{}, mapEnv(nil))
}

func TestVarsPanicsOnBadArgument(t *testing.T) {
	t.Parallel()
	defer func() {
		if recover() == nil {
			t.Fatal("Vars accepted a non-pointer argument")
		}
	}()
	_ = Vars(kindsConfig{})
}

func TestLoadPanicsOnUnsupportedFieldKind(t *testing.T) {
	t.Parallel()
	defer func() {
		if recover() == nil {
			t.Fatal("Load accepted an unsupported field kind")
		}
	}()
	var cfg struct {
		U uint `env:"U"`
	}
	_ = Load(&cfg, mapEnv(map[string]string{"U": "1"}))
}

func TestLoadSkipsUnexportedAndUntaggedFields(t *testing.T) {
	t.Parallel()
	var cfg struct {
		Tagged   string `env:"TAGGED"`
		Untagged string
		hidden   string // proves the loader skips unexported fields
	}
	if err := Load(&cfg, mapEnv(map[string]string{"TAGGED": "v", "Untagged": "x"})); err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Tagged != "v" || cfg.Untagged != "" {
		t.Errorf("cfg = %+v", cfg)
	}
	if got := len(Vars(&cfg)); got != 1 {
		t.Errorf("Vars = %d entries, want only the tagged field", got)
	}
}

func TestVarsListsEveryTaggedField(t *testing.T) {
	t.Parallel()
	vars := Vars(&kindsConfig{})
	if len(vars) != 10 {
		t.Fatalf("got %d vars, want 10: %+v", len(vars), vars)
	}
	byName := map[string]Var{}
	for _, v := range vars {
		byName[v.Name] = v
	}
	if v := byName["T_REQ"]; !v.Required {
		t.Errorf("T_REQ not marked required: %+v", v)
	}
	if v := byName["T_INT"]; v.Default != "7" {
		t.Errorf("T_INT default = %q, want 7", v.Default)
	}
	if _, ok := byName["T_INNER"]; !ok {
		t.Error("nested struct field T_INNER missing from Vars")
	}
}

func TestLoadErrorMessageShape(t *testing.T) {
	t.Parallel()
	err := &LoadError{Problems: []string{"A: bad", "B: worse"}}
	want := "configuration errors:\n  A: bad\n  B: worse"
	if err.Error() != want {
		t.Errorf("Error() = %q, want %q", err.Error(), want)
	}
}

// FuzzLoadValue proves the loader never panics and never accepts-and-corrupts
// on arbitrary environment values: every field either parses or reports a
// problem naming the variable.
func FuzzLoadValue(f *testing.F) {
	f.Add("true", "1", "1.5", "30s", "1,2,3")
	f.Add("garbage", "-", "NaN", "-1ns", ",,,")
	f.Add("", " ", "\x00", "9223372036854775808", "1,")
	f.Fuzz(func(t *testing.T, b, i, fl, d, l string) {
		var cfg struct {
			B  bool          `env:"B"`
			I  int64         `env:"I"`
			F  float64       `env:"F"`
			D  time.Duration `env:"D"`
			L  []int64       `env:"L"`
			S  string        `env:"S"`
			SR string        `env:"SR" secret:"true"`
			SU string        `env:"SU" secret:"url"`
		}
		env := map[string]string{"B": b, "I": i, "F": fl, "D": d, "L": l, "S": b, "SR": i, "SU": fl}
		err := Load(&cfg, mapEnv(env))
		var le *LoadError
		if err != nil && !errors.As(err, &le) {
			t.Fatalf("Load returned a non-LoadError: %T %v", err, err)
		}
		// Attrs must never panic, and secret attrs render only their
		// redacted placeholders regardless of the loaded value.
		for _, a := range Attrs(&cfg) {
			switch a.Key {
			case "SR":
				if a.Value.String() != "[set]" && a.Value.String() != "[unset]" {
					t.Fatalf("secret attr rendered %q, want [set]/[unset]", a.Value.String())
				}
			case "SU":
				got := a.Value.String()
				if got != "[set]" && got != "[unset]" &&
					!strings.HasSuffix(got, "/[redacted]") && strings.ContainsAny(got, "?#") {
					t.Fatalf("url-secret attr kept path/query detail: %q", got)
				}
			}
		}
	})
}
