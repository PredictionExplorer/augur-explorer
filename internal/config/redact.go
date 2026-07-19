package config

import (
	"fmt"
	"log/slog"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Attrs renders the effective configuration (defaults applied) as one slog
// attribute per env-tagged field, keyed by the variable name, for the
// startup "effective configuration" record. Secrets never appear:
// `secret:"true"` fields render as "[set]"/"[unset]" and `secret:"url"`
// fields keep only scheme and host (RPC URLs embed API keys in the path).
func Attrs(cfg any) []slog.Attr {
	rv := reflect.ValueOf(cfg)
	if rv.Kind() != reflect.Pointer || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		panic("config.Attrs: cfg must be a non-nil pointer to a struct")
	}
	var attrs []slog.Attr
	visitFields(rv.Elem(), func(f reflect.StructField, fv reflect.Value) {
		attrs = append(attrs, slog.String(f.Tag.Get("env"), renderField(f, fv)))
	})
	return attrs
}

// renderField renders one field value for logging, applying the field's
// secret redaction mode.
func renderField(f reflect.StructField, fv reflect.Value) string {
	switch f.Tag.Get("secret") {
	case "true":
		if fv.Kind() == reflect.String && fv.String() == "" {
			return "[unset]"
		}
		return "[set]"
	case "url":
		return RedactURL(fv.String())
	}
	if fv.Type() == reflect.TypeFor[time.Duration]() {
		return time.Duration(fv.Int()).String()
	}
	if fv.Type() == reflect.TypeFor[time.Time]() {
		t, _ := fv.Interface().(time.Time)
		if t.IsZero() {
			return "[unset]"
		}
		return t.Format(time.RFC3339)
	}
	switch fv.Kind() {
	case reflect.String:
		return fv.String()
	case reflect.Bool:
		return strconv.FormatBool(fv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(fv.Int(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(fv.Float(), 'g', -1, 64)
	case reflect.Slice:
		parts := make([]string, 0, fv.Len())
		for i := range fv.Len() {
			parts = append(parts, fmt.Sprintf("%v", fv.Index(i).Interface()))
		}
		return strings.Join(parts, ",")
	default:
		panic(fmt.Sprintf("config: unsupported field kind %s", fv.Type()))
	}
}

// RedactURL reduces a URL to scheme://host for logging: the path and query
// commonly carry credentials (https://provider/v2/<api-key>). Empty input
// renders "[unset]"; anything unparsable renders "[set]" rather than risking
// a leak.
func RedactURL(raw string) string {
	if raw == "" {
		return "[unset]"
	}
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return "[set]"
	}
	redacted := u.Scheme + "://" + u.Host
	if u.Path != "" || u.RawQuery != "" || u.User != nil {
		redacted += "/[redacted]"
	}
	return redacted
}
