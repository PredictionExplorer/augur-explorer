package common

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func parserContext(params map[string]string) (*httpx.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	for name, value := range params {
		req.SetPathValue(name, value)
	}
	w := httptest.NewRecorder()
	return httpx.NewContext(w, req), w
}

func assertParserError(t *testing.T, w *httptest.ResponseRecorder, want string) {
	t.Helper()
	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
	if got := w.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Errorf("Content-Type = %q, want JSON", got)
	}
	wantBytes, err := json.Marshal(httpx.H{"error": want, "status": 0})
	if err != nil {
		t.Fatal(err)
	}
	wantBody := string(wantBytes)
	if got := w.Body.String(); got != wantBody {
		t.Errorf("body = %q, want %q", got, wantBody)
	}
}

func TestParseTimeframeParams(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	const maxTimestamp = 2147483647

	tests := []struct {
		name         string
		params       map[string]string
		wantOK       bool
		wantInit     int
		wantFin      int
		wantInterval int
		wantError    string
	}{
		{
			name:         "valid",
			params:       map[string]string{"init_ts": "10", "fin_ts": "20", "interval_secs": "5"},
			wantOK:       true,
			wantInit:     10,
			wantFin:      20,
			wantInterval: 5,
		},
		{
			name:         "zero final timestamp and interval use defaults",
			params:       map[string]string{"init_ts": "10", "fin_ts": "0", "interval_secs": "0"},
			wantOK:       true,
			wantInit:     10,
			wantFin:      maxTimestamp,
			wantInterval: maxTimestamp - 10,
		},
		{
			name:         "zero interval spans requested range",
			params:       map[string]string{"init_ts": "5", "fin_ts": "12", "interval_secs": "0"},
			wantOK:       true,
			wantInit:     5,
			wantFin:      12,
			wantInterval: 7,
		},
		{
			name:         "native integer boundaries",
			params:       map[string]string{"init_ts": strconv.Itoa(minInt), "fin_ts": strconv.Itoa(maxInt), "interval_secs": "-1"},
			wantOK:       true,
			wantInit:     minInt,
			wantFin:      maxInt,
			wantInterval: -1,
		},
		{
			name:      "missing init timestamp",
			params:    map[string]string{"fin_ts": "20", "interval_secs": "5"},
			wantError: "'init_ts' parameter wasn't provided: <nil>",
		},
		{
			name:      "invalid init timestamp",
			params:    map[string]string{"init_ts": "bad", "fin_ts": "20", "interval_secs": "5"},
			wantError: `Bad 'init_ts' parameter: strconv.Atoi: parsing "bad": invalid syntax`,
		},
		{
			name:      "missing final timestamp",
			params:    map[string]string{"init_ts": "10", "interval_secs": "5"},
			wantError: "'fin_ts' parameter wasn't provided: <nil>",
		},
		{
			name:      "invalid final timestamp",
			params:    map[string]string{"init_ts": "10", "fin_ts": "bad", "interval_secs": "5"},
			wantError: `'fin_ts' parameter: strconv.Atoi: parsing "bad": invalid syntax`,
		},
		{
			name:      "missing interval",
			params:    map[string]string{"init_ts": "10", "fin_ts": "20"},
			wantError: "'interval_secs' parameter wasn't provided: <nil>",
		},
		{
			name:      "invalid interval",
			params:    map[string]string{"init_ts": "10", "fin_ts": "20", "interval_secs": "bad"},
			wantError: `Bad 'interval_secs' parameter: strconv.Atoi: parsing "bad": invalid syntax`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := parserContext(tt.params)
			ok, initTs, finTs, interval := ParseTimeframeParams(c)
			if ok != tt.wantOK || initTs != tt.wantInit || finTs != tt.wantFin || interval != tt.wantInterval {
				t.Errorf("result = (%v, %d, %d, %d), want (%v, %d, %d, %d)",
					ok, initTs, finTs, interval,
					tt.wantOK, tt.wantInit, tt.wantFin, tt.wantInterval)
			}
			if tt.wantOK {
				if w.Body.Len() != 0 {
					t.Errorf("successful parse wrote response %q", w.Body.String())
				}
				return
			}
			assertParserError(t, w, tt.wantError)
		})
	}
}

func TestParseOffsetLimitParamsJSON(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	tests := []struct {
		name      string
		params    map[string]string
		wantOK    bool
		wantOff   int
		wantLimit int
		wantError string
	}{
		{
			name:      "valid",
			params:    map[string]string{"offset": "7", "limit": "25"},
			wantOK:    true,
			wantOff:   7,
			wantLimit: 25,
		},
		{
			name:      "zero limit uses default",
			params:    map[string]string{"offset": "0", "limit": "0"},
			wantOK:    true,
			wantLimit: 20,
		},
		{
			name:      "negative values are preserved",
			params:    map[string]string{"offset": "-7", "limit": "-25"},
			wantOK:    true,
			wantOff:   -7,
			wantLimit: -25,
		},
		{
			name:      "native integer boundaries",
			params:    map[string]string{"offset": strconv.Itoa(minInt), "limit": strconv.Itoa(maxInt)},
			wantOK:    true,
			wantOff:   minInt,
			wantLimit: maxInt,
		},
		{
			name:      "missing offset",
			params:    map[string]string{"limit": "20"},
			wantError: "'offset' parameter wasn't provided: <nil>",
		},
		{
			name:      "invalid offset",
			params:    map[string]string{"offset": "bad", "limit": "20"},
			wantError: `Bad 'offset' parameter: strconv.Atoi: parsing "bad": invalid syntax`,
		},
		{
			name:      "missing limit",
			params:    map[string]string{"offset": "0"},
			wantError: "'limit' parameter wasn't provided: <nil>",
		},
		{
			name:      "invalid limit",
			params:    map[string]string{"offset": "0", "limit": "bad"},
			wantError: `'limit' parameter: strconv.Atoi: parsing "bad": invalid syntax`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := parserContext(tt.params)
			ok, offset, limit := ParseOffsetLimitParamsJSON(c)
			if ok != tt.wantOK || offset != tt.wantOff || limit != tt.wantLimit {
				t.Errorf("result = (%v, %d, %d), want (%v, %d, %d)",
					ok, offset, limit, tt.wantOK, tt.wantOff, tt.wantLimit)
			}
			if tt.wantOK {
				if w.Body.Len() != 0 {
					t.Errorf("successful parse wrote response %q", w.Body.String())
				}
				return
			}
			assertParserError(t, w, tt.wantError)
		})
	}
}

func TestParseInitFinTsParams(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	const maxTimestamp = 2147483647

	tests := []struct {
		name      string
		params    map[string]string
		wantOK    bool
		wantInit  int
		wantFin   int
		wantError string
	}{
		{
			name:     "valid",
			params:   map[string]string{"init_ts": "10", "fin_ts": "20"},
			wantOK:   true,
			wantInit: 10,
			wantFin:  20,
		},
		{
			name:     "zero final timestamp uses maximum",
			params:   map[string]string{"init_ts": "10", "fin_ts": "0"},
			wantOK:   true,
			wantInit: 10,
			wantFin:  maxTimestamp,
		},
		{
			name:     "native integer boundaries",
			params:   map[string]string{"init_ts": strconv.Itoa(minInt), "fin_ts": strconv.Itoa(maxInt)},
			wantOK:   true,
			wantInit: minInt,
			wantFin:  maxInt,
		},
		{
			name:      "missing init timestamp",
			params:    map[string]string{"fin_ts": "20"},
			wantError: `Bad 'init_ts' parameter: strconv.Atoi: parsing "": invalid syntax`,
		},
		{
			name:      "invalid init timestamp",
			params:    map[string]string{"init_ts": "bad", "fin_ts": "20"},
			wantError: `Bad 'init_ts' parameter: strconv.Atoi: parsing "bad": invalid syntax`,
		},
		{
			name:      "missing final timestamp",
			params:    map[string]string{"init_ts": "10"},
			wantError: `Bad 'fin_ts' parameter: strconv.Atoi: parsing "": invalid syntax`,
		},
		{
			name:      "invalid final timestamp",
			params:    map[string]string{"init_ts": "10", "fin_ts": "bad"},
			wantError: `Bad 'fin_ts' parameter: strconv.Atoi: parsing "bad": invalid syntax`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := parserContext(tt.params)
			ok, initTs, finTs := ParseInitFinTsParams(c)
			if ok != tt.wantOK || initTs != tt.wantInit || finTs != tt.wantFin {
				t.Errorf("result = (%v, %d, %d), want (%v, %d, %d)",
					ok, initTs, finTs, tt.wantOK, tt.wantInit, tt.wantFin)
			}
			if tt.wantOK {
				if w.Body.Len() != 0 {
					t.Errorf("successful parse wrote response %q", w.Body.String())
				}
				return
			}
			assertParserError(t, w, tt.wantError)
		})
	}
}
