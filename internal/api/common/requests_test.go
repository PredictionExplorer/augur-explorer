package common

import (
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func TestErrorResponders(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*httpx.Context)
		wantStatus int
		wantBody   httpx.H
	}{
		{
			name:       "legacy responder",
			call:       func(c *httpx.Context) { RespondError(c, `bad "request"`) },
			wantStatus: http.StatusBadRequest,
			wantBody:   httpx.H{"status": 0, "error": `bad "request"`},
		},
		{
			name:       "JSON responder",
			call:       func(c *httpx.Context) { RespondErrorJSON(c, "invalid <value>\n") },
			wantStatus: http.StatusBadRequest,
			wantBody:   httpx.H{"status": 0, "error": "invalid <value>\n"},
		},
		{
			name:       "internal responder",
			call:       RespondInternalErrorJSON,
			wantStatus: http.StatusInternalServerError,
			wantBody:   httpx.H{"status": 0, "error": "Internal server error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c := httpx.NewContext(w, httptest.NewRequest(http.MethodGet, "/", nil))
			tt.call(c)

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}
			if got := w.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
				t.Errorf("Content-Type = %q, want JSON", got)
			}
			want, err := json.Marshal(tt.wantBody)
			if err != nil {
				t.Fatal(err)
			}
			if got := w.Body.String(); got != string(want) {
				t.Errorf("body = %q, want %q", got, want)
			}
		})
	}
}

func TestParseIntFromRemoteOrError(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		jsonOutput bool
		want       int64
		wantOK     bool
		wantError  string
	}{
		{name: "zero JSON", input: "0", jsonOutput: true, wantOK: true},
		{name: "positive legacy", input: "+42", jsonOutput: false, want: 42, wantOK: true},
		{name: "negative JSON", input: "-42", jsonOutput: true, want: -42, wantOK: true},
		{name: "maximum int64", input: strconv.FormatInt(math.MaxInt64, 10), jsonOutput: false, want: math.MaxInt64, wantOK: true},
		{name: "minimum int64", input: strconv.FormatInt(math.MinInt64, 10), jsonOutput: true, want: math.MinInt64, wantOK: true},
		{
			name:       "invalid JSON",
			input:      "bad",
			jsonOutput: true,
			wantError:  `Can't parse integer parameter: strconv.ParseInt: parsing "bad": invalid syntax`,
		},
		{
			name:       "invalid legacy",
			input:      "",
			jsonOutput: false,
			wantError:  `Can't parse integer parameter: strconv.ParseInt: parsing "": invalid syntax`,
		},
		{
			name:       "leading whitespace rejected",
			input:      " 1",
			jsonOutput: true,
			wantError:  `Can't parse integer parameter: strconv.ParseInt: parsing " 1": invalid syntax`,
		},
		{
			name:       "int64 overflow rejected",
			input:      "9223372036854775808",
			jsonOutput: false,
			wantError:  `Can't parse integer parameter: strconv.ParseInt: parsing "9223372036854775808": value out of range`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c := httpx.NewContext(w, httptest.NewRequest(http.MethodGet, "/", nil))
			input := tt.input
			got, ok := ParseIntFromRemoteOrError(c, tt.jsonOutput, &input)
			if got != tt.want || ok != tt.wantOK {
				t.Errorf("result = (%d, %v), want (%d, %v)", got, ok, tt.want, tt.wantOK)
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

func TestParseTimeframeIniFin(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	const maxTimestamp = 2147483647

	tests := []struct {
		name       string
		params     map[string]string
		jsonOutput bool
		wantOK     bool
		wantInit   int
		wantFin    int
		wantError  string
	}{
		{
			name:       "valid JSON",
			params:     map[string]string{"init_ts": "10", "fin_ts": "20"},
			jsonOutput: true,
			wantOK:     true,
			wantInit:   10,
			wantFin:    20,
		},
		{
			name:       "valid legacy",
			params:     map[string]string{"init_ts": "-10", "fin_ts": "20"},
			jsonOutput: false,
			wantOK:     true,
			wantInit:   -10,
			wantFin:    20,
		},
		{
			name:       "zero final timestamp uses maximum",
			params:     map[string]string{"init_ts": "10", "fin_ts": "0"},
			jsonOutput: true,
			wantOK:     true,
			wantInit:   10,
			wantFin:    maxTimestamp,
		},
		{
			name:       "native integer boundaries",
			params:     map[string]string{"init_ts": strconv.Itoa(minInt), "fin_ts": strconv.Itoa(maxInt)},
			jsonOutput: false,
			wantOK:     true,
			wantInit:   minInt,
			wantFin:    maxInt,
		},
		{
			name:       "missing init timestamp JSON",
			params:     map[string]string{"fin_ts": "20"},
			jsonOutput: true,
			wantError:  "'init_ts' parameter wasn't provided: <nil>",
		},
		{
			name:       "invalid init timestamp legacy",
			params:     map[string]string{"init_ts": "bad", "fin_ts": "20"},
			jsonOutput: false,
			wantError:  `Bad 'init_ts' parameter: strconv.Atoi: parsing "bad": invalid syntax`,
		},
		{
			name:       "missing final timestamp JSON",
			params:     map[string]string{"init_ts": "10"},
			jsonOutput: true,
			wantError:  "'fin_ts' parameter wasn't provided: <nil>",
		},
		{
			name:       "invalid final timestamp legacy",
			params:     map[string]string{"init_ts": "10", "fin_ts": "bad"},
			jsonOutput: false,
			wantError:  `'fin_ts' parameter: strconv.Atoi: parsing "bad": invalid syntax`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, w := parserContext(tt.params)
			ok, initTs, finTs := ParseTimeframeIniFin(c, tt.jsonOutput)
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
