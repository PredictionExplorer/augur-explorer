//go:build integration

package apitest

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"

	apiv2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

const (
	v2ListRoundsPath    = "/api/v2/cosmicgame/rounds"
	v2CurrentRoundPath  = "/api/v2/cosmicgame/rounds/current"
	v2GetRoundPath      = "/api/v2/cosmicgame/rounds/{round}"
	v2ListRoundBidsPath = "/api/v2/cosmicgame/rounds/{round}/bids"
	v2GetRoundBidPath   = "/api/v2/cosmicgame/rounds/{round}/bids/{position}"
)

type v2GoldenCase struct {
	name       string
	target     string
	template   string
	pathParams map[string]string
	ctx        context.Context
}

func TestAPIV2CurrentRound(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []v2GoldenCase{
		{
			name:       "current_round_get",
			target:     v2CurrentRoundPath,
			template:   v2CurrentRoundPath,
			pathParams: map[string]string{},
		},
		{
			name:       "current_round_error_internal",
			target:     v2CurrentRoundPath,
			template:   v2CurrentRoundPath,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
	}

	runV2GoldenCases(t, h, spec, cases)

	original := h.state.Snapshot()
	h.state.SetBidPrice("error", 0)
	defer h.state.SetBidPrice(original.BidPrice, original.BidPriceEth)
	unavailable := h.get(t, v2CurrentRoundPath)
	if unavailable.Code != http.StatusServiceUnavailable ||
		unavailable.Header().Get("Retry-After") != "5" {
		t.Fatalf("unavailable live state = status %d Retry-After %q, want 503/5",
			unavailable.Code, unavailable.Header().Get("Retry-After"))
	}
	runV2GoldenCases(t, h, spec, []v2GoldenCase{
		{
			name:       "current_round_error_unavailable",
			target:     v2CurrentRoundPath,
			template:   v2CurrentRoundPath,
			pathParams: map[string]string{},
		},
	})
}

func TestAPIV2RoundBids(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	firstPath := "/api/v2/cosmicgame/rounds/0/bids?limit=2"
	firstResponse := h.get(t, firstPath)
	if firstResponse.Code != http.StatusOK {
		t.Fatalf("first page: status=%d body=%s", firstResponse.Code, firstResponse.Body.String())
	}
	var firstPage apiv2.RoundBidPage
	if err := json.Unmarshal(firstResponse.Body.Bytes(), &firstPage); err != nil {
		t.Fatalf("decoding first page: %v", err)
	}
	if firstPage.Meta.NextCursor == nil {
		t.Fatal("fixture first page did not return a continuation cursor")
	}

	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	cases := []v2GoldenCase{
		{
			name:       "list_first_page",
			target:     firstPath,
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "list_next_page",
			target:     "/api/v2/cosmicgame/rounds/0/bids?limit=2&cursor=" + *firstPage.Meta.NextCursor,
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "list_empty_page",
			target:     "/api/v2/cosmicgame/rounds/999/bids?limit=2",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "999"},
		},
		{
			name:       "get_bid",
			target:     "/api/v2/cosmicgame/rounds/0/bids/3",
			template:   v2GetRoundBidPath,
			pathParams: map[string]string{"round": "0", "position": "3"},
		},
		{
			name:       "error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds/0/bids?cursor=not-a-cursor",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds/0/bids?limit=201",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds/0/bids?limit=wat",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "error_cross_round_cursor",
			target:     "/api/v2/cosmicgame/rounds/1/bids?cursor=" + *firstPage.Meta.NextCursor,
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "1"},
		},
		{
			name:       "error_bid_not_found",
			target:     "/api/v2/cosmicgame/rounds/0/bids/99",
			template:   v2GetRoundBidPath,
			pathParams: map[string]string{"round": "0", "position": "99"},
		},
		{
			name:       "error_internal",
			target:     "/api/v2/cosmicgame/rounds/0/bids?limit=2",
			template:   v2ListRoundBidsPath,
			pathParams: map[string]string{"round": "0"},
			ctx:        cancelledCtx,
		},
	}

	runV2GoldenCases(t, h, spec, cases)
}

func TestAPIV2Rounds(t *testing.T) {
	h := server(t)
	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatalf("loading embedded v2 spec: %v", err)
	}
	if err := spec.Validate(context.Background()); err != nil {
		t.Fatalf("validating embedded v2 spec: %v", err)
	}

	firstPath := "/api/v2/cosmicgame/rounds?limit=2"
	firstResponse := h.get(t, firstPath)
	if firstResponse.Code != http.StatusOK {
		t.Fatalf("first page: status=%d body=%s", firstResponse.Code, firstResponse.Body.String())
	}
	var firstPage apiv2.RoundPage
	if err := json.Unmarshal(firstResponse.Body.Bytes(), &firstPage); err != nil {
		t.Fatalf("decoding first page: %v", err)
	}
	if firstPage.Meta.NextCursor == nil {
		t.Fatal("fixture first page did not return a continuation cursor")
	}

	afterRoundZero := base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"r":0,"e":5018}`))
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()

	cases := []v2GoldenCase{
		{
			name:       "rounds_list_first_page",
			target:     firstPath,
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_list_next_page",
			target:     "/api/v2/cosmicgame/rounds?limit=2&cursor=" + *firstPage.Meta.NextCursor,
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_list_empty_page",
			target:     "/api/v2/cosmicgame/rounds?limit=2&cursor=" + afterRoundZero,
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_get_0",
			target:     "/api/v2/cosmicgame/rounds/0",
			template:   v2GetRoundPath,
			pathParams: map[string]string{"round": "0"},
		},
		{
			name:       "rounds_get_2",
			target:     "/api/v2/cosmicgame/rounds/2",
			template:   v2GetRoundPath,
			pathParams: map[string]string{"round": "2"},
		},
		{
			name:       "rounds_error_malformed_cursor",
			target:     "/api/v2/cosmicgame/rounds?cursor=not-a-cursor",
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_error_invalid_limit",
			target:     "/api/v2/cosmicgame/rounds?limit=201",
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_error_bind_limit",
			target:     "/api/v2/cosmicgame/rounds?limit=wat",
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
		},
		{
			name:       "rounds_error_not_found",
			target:     "/api/v2/cosmicgame/rounds/999",
			template:   v2GetRoundPath,
			pathParams: map[string]string{"round": "999"},
		},
		{
			name:       "rounds_error_open_round",
			target:     "/api/v2/cosmicgame/rounds/3",
			template:   v2GetRoundPath,
			pathParams: map[string]string{"round": "3"},
		},
		{
			name:       "rounds_error_internal",
			target:     "/api/v2/cosmicgame/rounds?limit=2",
			template:   v2ListRoundsPath,
			pathParams: map[string]string{},
			ctx:        cancelledCtx,
		},
	}

	runV2GoldenCases(t, h, spec, cases)
}

func runV2GoldenCases(t *testing.T, h *harness, spec *openapi3.T, cases []v2GoldenCase) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			request := request{path: tc.target, ctx: tc.ctx}
			first := h.do(t, request)
			second := h.do(t, request)

			firstBody := canonicalJSON(t, first.Body.Bytes())
			secondBody := canonicalJSON(t, second.Body.Bytes())
			if first.Code != second.Code || !reflect.DeepEqual(firstBody, secondBody) {
				t.Fatalf("nondeterministic v2 response: first=%d %#v second=%d %#v",
					first.Code, firstBody, second.Code, secondBody)
			}

			validateV2Response(t, spec, tc, first)
			compareV2Golden(t, tc.name, response{
				Status:      first.Code,
				ContentType: contentTypeOf(first),
				Body:        firstBody,
			})
		})
	}
}

func validateV2Response(t *testing.T, spec *openapi3.T, tc v2GoldenCase, response *httptest.ResponseRecorder) {
	t.Helper()

	pathItem := spec.Paths.Value(tc.template)
	if pathItem == nil || pathItem.Get == nil {
		t.Fatalf("spec has no GET operation for %s", tc.template)
	}
	request := httptest.NewRequest(http.MethodGet, tc.target, nil)
	route := &routers.Route{
		Spec:      spec,
		Path:      tc.template,
		PathItem:  pathItem,
		Method:    http.MethodGet,
		Operation: pathItem.Get,
	}
	requestInput := &openapi3filter.RequestValidationInput{
		Request:    request,
		PathParams: tc.pathParams,
		Route:      route,
	}
	responseInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestInput,
		Status:                 response.Code,
		Header:                 response.Header(),
	}
	responseInput.SetBodyBytes(response.Body.Bytes())
	if err := openapi3filter.ValidateResponse(context.Background(), responseInput); err != nil {
		t.Fatalf("%s response violates OpenAPI v2: %v\n%s", tc.name, err, response.Body.String())
	}
}
