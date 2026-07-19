// Package config loads and validates the typed per-service configuration
// from environment variables (docs/MODERNIZATION.md §8.3).
//
// Every service owns one struct (APIServer, ETL, Notibot, ImggenMonitor,
// RwalkAlarm) whose fields carry `env` tags naming the variable, optional
// `default` values, `required` markers and `secret` redaction modes. Load
// fails fast at startup and aggregates every bad variable into one error, so
// a deployment with three problems reports all three at once. Attrs renders
// the effective configuration for the startup log with secrets redacted.
//
// The srvmonitor TUI keeps its own loader (internal/srvmonitor.LoadFromEnv):
// its indexed variable families (RPC{n}_URL, SSL_CERT{n}_HOST, ...) do not
// fit flat struct tags.
package config

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// LoadError aggregates every configuration problem found by a Load call.
type LoadError struct {
	// Problems holds one human-readable line per bad variable, in struct
	// declaration order.
	Problems []string
}

// Error renders the problems as one multiline message suitable for the
// startup failure output of a service.
func (e *LoadError) Error() string {
	return "configuration errors:\n  " + strings.Join(e.Problems, "\n  ")
}

// Load fills dst (a pointer to a struct) from environment variables named by
// `env` field tags, reading through getenv. Values are whitespace-trimmed;
// an empty value falls back to the `default` tag (or the zero value), and
// `required:"true"` fields report a problem when unset. Untagged nested
// structs are recursed into. All problems are aggregated into one *LoadError.
//
// Field kinds: string, bool (true/false, 1/0, yes/no, on/off), integers,
// float64, time.Duration, time.Time (RFC 3339) and []int64
// (comma-separated). Anything else is a programmer error and panics.
func Load(dst any, getenv func(string) string) error {
	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Pointer || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		panic("config.Load: dst must be a non-nil pointer to a struct")
	}
	var problems []string
	visitFields(rv.Elem(), func(f reflect.StructField, fv reflect.Value) {
		name := f.Tag.Get("env")
		raw := strings.TrimSpace(getenv(name))
		if raw == "" {
			if f.Tag.Get("required") == "true" {
				problems = append(problems, name+": required but not set")
				return
			}
			raw = f.Tag.Get("default")
			if raw == "" {
				return
			}
		}
		if err := setField(fv, raw); err != nil {
			problems = append(problems, fmt.Sprintf("%s: %v", name, err))
		}
	})
	if len(problems) > 0 {
		return &LoadError{Problems: problems}
	}
	return nil
}

// Var describes one environment variable declared by a configuration struct.
// The .env.example verification test walks these.
type Var struct {
	// Name is the environment variable name from the `env` tag.
	Name string
	// Default is the `default` tag value ("" when none).
	Default string
	// Required reports the `required:"true"` tag.
	Required bool
	// Secret is the `secret` tag mode: "", "true" or "url".
	Secret string
}

// Vars lists every env-tagged field of cfg (a pointer to a struct) in
// declaration order, recursing into untagged nested structs.
func Vars(cfg any) []Var {
	rv := reflect.ValueOf(cfg)
	if rv.Kind() != reflect.Pointer || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		panic("config.Vars: cfg must be a non-nil pointer to a struct")
	}
	var vars []Var
	visitFields(rv.Elem(), func(f reflect.StructField, _ reflect.Value) {
		vars = append(vars, Var{
			Name:     f.Tag.Get("env"),
			Default:  f.Tag.Get("default"),
			Required: f.Tag.Get("required") == "true",
			Secret:   f.Tag.Get("secret"),
		})
	})
	return vars
}

// visitFields calls fn for every exported field of the struct value v that
// carries an `env` tag, recursing into untagged nested structs in
// declaration order.
func visitFields(v reflect.Value, fn func(f reflect.StructField, fv reflect.Value)) {
	t := v.Type()
	for i := range t.NumField() {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		if f.Tag.Get("env") == "" {
			if f.Type.Kind() == reflect.Struct {
				visitFields(v.Field(i), fn)
			}
			continue
		}
		fn(f, v.Field(i))
	}
}

// setField parses raw into the field value. The supported kinds mirror the
// Load documentation; an unsupported kind panics (programmer error caught by
// the loader tests).
func setField(v reflect.Value, raw string) error {
	if v.Type() == reflect.TypeFor[time.Duration]() {
		d, err := time.ParseDuration(raw)
		if err != nil {
			return fmt.Errorf("cannot parse %q as a duration (examples: 30s, 5m)", raw)
		}
		v.SetInt(int64(d))
		return nil
	}
	if v.Type() == reflect.TypeFor[time.Time]() {
		t, err := time.Parse(time.RFC3339, raw)
		if err != nil {
			return fmt.Errorf("cannot parse %q as an RFC 3339 timestamp (example: 2027-01-01T00:00:00Z)", raw)
		}
		v.Set(reflect.ValueOf(t))
		return nil
	}
	switch v.Kind() {
	case reflect.String:
		v.SetString(raw)
		return nil
	case reflect.Bool:
		b, err := parseBool(raw)
		if err != nil {
			return err
		}
		v.SetBool(b)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(raw, 10, v.Type().Bits())
		if err != nil {
			return fmt.Errorf("cannot parse %q as an integer", raw)
		}
		v.SetInt(n)
		return nil
	case reflect.Float64:
		f, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return fmt.Errorf("cannot parse %q as a number", raw)
		}
		v.SetFloat(f)
		return nil
	case reflect.Slice:
		if v.Type().Elem().Kind() != reflect.Int64 {
			panic(fmt.Sprintf("config: unsupported field kind %s", v.Type()))
		}
		list, err := parseInt64List(raw)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(list))
		return nil
	default:
		panic(fmt.Sprintf("config: unsupported field kind %s", v.Type()))
	}
}

// parseBool accepts the common spellings of a boolean flag. Unlike the
// legacy per-variable parsers (some checked == "1", some fell back to a
// default on garbage), a malformed value is a startup error.
func parseBool(raw string) (bool, error) {
	switch strings.ToLower(raw) {
	case "1", "t", "true", "yes", "y", "on":
		return true, nil
	case "0", "f", "false", "no", "n", "off":
		return false, nil
	}
	return false, fmt.Errorf("cannot parse %q as a boolean (use true/false, 1/0, yes/no, on/off)", raw)
}

// parseInt64List parses a comma-separated integer list, ignoring empty
// entries ("1, 2," parses to [1 2]).
func parseInt64List(raw string) ([]int64, error) {
	parts := strings.Split(raw, ",")
	out := make([]int64, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		n, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse list entry %q as an integer", p)
		}
		out = append(out, n)
	}
	return out, nil
}
