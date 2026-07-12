package cosmicgame

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func TestRegisteredDatabaseHandlersFailCleanlyBeforeInitialization(t *testing.T) {
	previousContext := common.Ctx
	previousEnabled := Enabled
	common.Ctx = nil
	Enabled = true
	t.Cleanup(func() {
		common.Ctx = previousContext
		Enabled = previousEnabled
	})

	router := httpx.NewRouter()
	RegisterAPIRoutes(router)

	contractOnly := map[string]bool{
		"/api/cosmicgame/statistics/dashboard": true,
		"/api/cosmicgame/bid/cst_price":        true,
		"/api/cosmicgame/bid/eth_price":        true,
	}
	// #nosec G101 -- these are public response texts, not credentials.
	specialError := map[string]string{
		"/api/cosmicgame/cst/metadata/{token_id}": "CosmicGame module or database not available",
		"/cg/metadata/{token_id}":                 "CosmicGame module or database not available",
	}
	placeholder := regexp.MustCompile(`\{[^}]+\}`)
	tested := 0
	for _, route := range router.Routes() {
		if route.Method != http.MethodGet || contractOnly[route.Pattern] {
			continue
		}
		tested++
		path := placeholder.ReplaceAllString(route.Pattern, "0")
		t.Run(route.Pattern, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, path, nil)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)

			if response.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
			}
			var envelope struct {
				Status int    `json:"status"`
				Error  string `json:"error"`
			}
			if err := json.Unmarshal(response.Body.Bytes(), &envelope); err != nil {
				t.Fatalf("decode response: %v\n%s", err, response.Body.String())
			}
			wantError := specialError[route.Pattern]
			if wantError == "" {
				wantError = "Database link wasn't configured"
			}
			if envelope.Status != 0 || envelope.Error != wantError {
				t.Fatalf("envelope = %+v", envelope)
			}
		})
	}
	if tested < 100 {
		t.Fatalf("tested only %d guarded GET routes; expected the complete v1 surface", tested)
	}
}

func TestFloatSanitizersRemoveNonFiniteValuesRecursively(t *testing.T) {
	t.Run("safe scalar", func(t *testing.T) {
		for name, input := range map[string]float64{
			"nan":      math.NaN(),
			"positive": math.Inf(1),
			"negative": math.Inf(-1),
		} {
			if got := safeFloat64(input); got != 0 {
				t.Errorf("%s = %v, want 0", name, got)
			}
		}
		if got := safeFloat64(1.25); got != 1.25 {
			t.Fatalf("finite value = %v", got)
		}
	})

	t.Run("maps and interface slices", func(t *testing.T) {
		value := map[string]interface{}{
			"nil":    nil,
			"finite": 3.5,
			"nan":    math.NaN(),
			"map": map[string]interface{}{
				"positiveInfinity": math.Inf(1),
			},
			"slice": []interface{}{
				math.Inf(-1),
				map[string]interface{}{"nestedNaN": math.NaN()},
				[]interface{}{math.Inf(1), "unchanged"},
			},
			"other": "unchanged",
		}
		sanitizeMapFloatsForJSON(value)
		assertFiniteJSON(t, value)
		if value["finite"] != 3.5 || value["other"] != "unchanged" || value["nil"] != nil {
			t.Fatalf("sanitizer changed finite/non-float values: %#v", value)
		}
	})

	t.Run("struct pointers and slices", func(t *testing.T) {
		type nested struct {
			Value float64
		}
		type payload struct {
			Top        float64
			Nested     nested
			Pointer    *nested
			Nil        *nested
			Collection []nested
			private    float64
		}
		value := payload{
			Top:        math.NaN(),
			Nested:     nested{Value: math.Inf(1)},
			Pointer:    &nested{Value: math.Inf(-1)},
			Collection: []nested{{Value: math.NaN()}, {Value: 4.5}},
			private:    math.Inf(1),
		}
		sanitizeFloatsForJSON(&value)
		if value.Top != 0 || value.Nested.Value != 0 || value.Pointer.Value != 0 ||
			value.Collection[0].Value != 0 || value.Collection[1].Value != 4.5 {
			t.Fatalf("struct sanitizer result = %#v", value)
		}
		if !math.IsInf(value.private, 1) {
			t.Fatal("unexported field should remain untouched")
		}

		scalars := []float64{math.NaN(), math.Inf(1), 7}
		sanitizeFloatsForJSON(scalars)
		if scalars[0] != 0 || scalars[1] != 0 || scalars[2] != 7 {
			t.Fatalf("slice sanitizer result = %#v", scalars)
		}

		pointer := &value
		sanitizeFloatsForJSON(&pointer)
		var nilPayload *payload
		sanitizeFloatsForJSON(nilPayload)
		sanitizeFloatsForJSON(nil)
	})
}

func TestRespondStoreErrorKeepsDetailsOutOfResponse(t *testing.T) {
	previousInfo, previousError := Info, Error
	var infoLog, errorLog bytes.Buffer
	Info = log.New(&infoLog, "", 0)
	Error = log.New(&errorLog, "", 0)
	t.Cleanup(func() {
		Info = previousInfo
		Error = previousError
	})

	t.Run("server error is logged", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/api/cosmicgame/test", nil)
		request.Pattern = "GET /api/cosmicgame/test"
		respondStoreError(httpx.NewContext(response, request), errors.New("database password secret"))

		if response.Code != http.StatusInternalServerError {
			t.Fatalf("status = %d", response.Code)
		}
		if strings.Contains(response.Body.String(), "database password secret") {
			t.Fatalf("response leaked internal error: %s", response.Body.String())
		}
		if !strings.Contains(infoLog.String(), "database password secret") ||
			!strings.Contains(errorLog.String(), "database password secret") {
			t.Fatalf("underlying error was not logged: info=%q error=%q", infoLog.String(), errorLog.String())
		}
	})

	t.Run("client cancellation is not logged", func(t *testing.T) {
		infoLog.Reset()
		errorLog.Reset()
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/api/cosmicgame/test", nil)
		respondStoreError(httpx.NewContext(response, request), context.Canceled)
		if response.Code != http.StatusInternalServerError {
			t.Fatalf("status = %d", response.Code)
		}
		if infoLog.Len() != 0 || errorLog.Len() != 0 {
			t.Fatalf("cancellation was logged: info=%q error=%q", infoLog.String(), errorLog.String())
		}
	})
}

func assertFiniteJSON(t *testing.T, value interface{}) {
	t.Helper()
	if _, err := json.Marshal(value); err != nil {
		t.Fatalf("value is not valid JSON after sanitizing: %v\n%#v", err, value)
	}
}
