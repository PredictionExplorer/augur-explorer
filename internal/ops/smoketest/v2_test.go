package smoketest

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/PredictionExplorer/augur-explorer/internal/api/policy"
	apiv2 "github.com/PredictionExplorer/augur-explorer/internal/api/v2"
)

func TestParseSuite(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		input string
		want  Suite
	}{
		{input: "v1", want: SuiteV1},
		{input: " V2 ", want: SuiteV2},
		{input: "BOTH", want: SuiteBoth},
		{input: "operational", want: SuiteOperational},
	} {
		got, err := ParseSuite(test.input)
		if err != nil || got != test.want {
			t.Errorf("ParseSuite(%q) = %q, %v; want %q", test.input, got, err, test.want)
		}
	}
	if _, err := ParseSuite("all"); err == nil || !strings.Contains(err.Error(), "v1, v2, both, or operational") {
		t.Fatalf("invalid suite error = %v", err)
	}
}

func TestSuiteParameterRequirementsAndMembers(t *testing.T) {
	t.Parallel()
	if !SuiteV1.RequiresParameters() || !SuiteV2.RequiresParameters() ||
		!SuiteBoth.RequiresParameters() || SuiteOperational.RequiresParameters() {
		t.Fatal("suite parameter requirements changed")
	}
	members, err := SuiteBoth.members()
	if err != nil || !slices.Equal(members, []Suite{SuiteV2, SuiteV1}) {
		t.Fatalf("both members = %v, %v", members, err)
	}
	if _, err := Suite("invalid").members(); err == nil {
		t.Fatal("invalid suite unexpectedly has members")
	}
	if _, err := buildProbes(SuiteBoth, DefaultParams()); err == nil ||
		!strings.Contains(err.Error(), "cannot build probes") {
		t.Fatalf("build invalid member error = %v", err)
	}
}

func TestBuildV2ProbesMatchesEveryOpenAPIGet(t *testing.T) {
	t.Parallel()
	probes, err := BuildV2Probes(DefaultParams())
	if err != nil {
		t.Fatal(err)
	}
	if len(probes) != 101 {
		t.Fatalf("v2 probe count = %d, want 101", len(probes))
	}

	spec, err := apiv2.GetSpec()
	if err != nil {
		t.Fatal(err)
	}
	want := make(map[string]string)
	for template, item := range spec.Paths.Map() {
		if item.Get != nil {
			want[item.Get.OperationID] = template
		}
	}
	if len(want) != len(probes) {
		t.Fatalf("OpenAPI GET count = %d, probes = %d", len(want), len(probes))
	}

	seen := make(map[string]bool, len(probes))
	previous := ""
	for _, probe := range probes {
		if probe.Suite != SuiteV2 || probe.OperationID == "" || probe.operation == nil {
			t.Errorf("incomplete probe: %#v", probe)
		}
		if template, ok := want[probe.OperationID]; !ok || template != probe.Template {
			t.Errorf("probe %s template = %q, OpenAPI = %q/%v",
				probe.OperationID, probe.Template, template, ok)
		}
		if seen[probe.OperationID] {
			t.Errorf("duplicate operation probe %q", probe.OperationID)
		}
		seen[probe.OperationID] = true
		if previous != "" && probe.Template < previous {
			t.Errorf("probe order is unstable: %q before %q", previous, probe.Template)
		}
		previous = probe.Template
		if strings.Contains(probe.Endpoint, "{") {
			t.Errorf("%s has unbound endpoint %q", probe.OperationID, probe.Endpoint)
		}
		if policy.V1Deprecated(probe.Endpoint) {
			t.Errorf("%s is D6-unsafe: %q", probe.OperationID, probe.Endpoint)
		}
		req, requestErr := http.NewRequest(http.MethodGet, "https://api.example"+probe.Endpoint, nil)
		if requestErr != nil {
			t.Errorf("%s request: %v", probe.OperationID, requestErr)
			continue
		}
		if requestErr := validateV2Request(context.Background(), req, probe); requestErr != nil {
			t.Errorf("%s %s violates OpenAPI: %v", probe.OperationID, probe.Endpoint, requestErr)
		}
	}
}

func TestBuildV2ProbesBindsProductionValues(t *testing.T) {
	t.Parallel()
	params := DefaultParams()
	params.UserAddress = "0xAb/c?"
	params.RoundNumber = "77"
	params.BidRound = "77"
	params.BidPosition = "8"
	params.CSTActionID = "51"
	params.RandomWalkActionID = "52"
	params.RandomWalkTokenID = "123"
	params.TimestampMin = "100"
	params.TimestampMax = "200"
	params.FromDate = "20260101"
	params.ToDate = "20260131"

	probes, err := BuildV2Probes(params)
	if err != nil {
		t.Fatal(err)
	}
	byTemplate := make(map[string]Probe, len(probes))
	for _, probe := range probes {
		byTemplate[probe.Template] = probe
	}
	for template, wantPart := range map[string]string{
		"/api/v2/cosmicgame/users/{address}":                        "0xAb%2Fc%3F",
		"/api/v2/cosmicgame/rounds/{round}/bids/{position}":         "/77/bids/8",
		"/api/v2/cosmicgame/staking/cst/actions/{actionId}":         "/actions/51",
		"/api/v2/cosmicgame/staking/random-walk/actions/{actionId}": "/actions/52",
		"/api/v2/randomwalk/tokens/{tokenId}":                       "/tokens/123",
	} {
		if got := byTemplate[template].Endpoint; !strings.Contains(got, wantPart) {
			t.Errorf("%s endpoint = %q, want substring %q", template, got, wantPart)
		}
	}

	for _, template := range []string{
		"/api/v2/cosmicgame/cosmic-token/supply-history/daily",
	} {
		query := mustProbeQuery(t, byTemplate[template])
		if query.Get("from") != "2026-01-01" || query.Get("to") != "2026-01-31" {
			t.Errorf("%s query = %v", template, query)
		}
	}
	for _, template := range []string{
		"/api/v2/cosmicgame/statistics/bidding/activity",
		"/api/v2/randomwalk/statistics/trading-volume",
	} {
		query := mustProbeQuery(t, byTemplate[template])
		if query.Get("from") != "100" || query.Get("to") != "200" {
			t.Errorf("%s query = %v", template, query)
		}
	}
}

func mustProbeQuery(t *testing.T, probe Probe) url.Values {
	t.Helper()
	parsed, err := url.Parse(probe.Endpoint)
	if err != nil {
		t.Fatal(err)
	}
	return parsed.Query()
}

func TestBuildOperationalProbesAreStableAndD6Safe(t *testing.T) {
	t.Parallel()
	probes, err := BuildOperationalProbes()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"/healthz",
		"/readyz",
		"/version",
		"/api/v2/cosmicgame/rounds?limit=1",
		"/api/v2/cosmicgame/statistics/counters",
		"/api/v2/randomwalk/statistics",
	}
	got := make([]string, 0, len(probes))
	for _, probe := range probes {
		got = append(got, probe.Endpoint)
		if probe.Suite != SuiteOperational || policy.V1Deprecated(probe.Endpoint) {
			t.Errorf("unsafe operational probe: %#v", probe)
		}
		for _, forbidden := range []string{
			"/contracts/configuration",
			"/contracts/balances",
			"/rounds/current",
		} {
			if strings.Contains(probe.Endpoint, forbidden) {
				t.Errorf("degradable contract-state route in operational suite: %s", probe.Endpoint)
			}
		}
	}
	if !slices.Equal(got, want) {
		t.Fatalf("operational endpoints = %v, want %v", got, want)
	}
}

func TestCompactDateToISORejectsNonCanonicalInput(t *testing.T) {
	t.Parallel()
	if got, err := compactDateToISO("20260723"); err != nil || got != "2026-07-23" {
		t.Fatalf("valid date = %q, %v", got, err)
	}
	for _, value := range []string{"", "2026-07-23", "20260230", "password-secret"} {
		if _, err := compactDateToISO(value); err == nil {
			t.Errorf("compactDateToISO(%q) succeeded", value)
		}
	}
}

func TestV2QueryValueFailsClosedForNewRequiredParameters(t *testing.T) {
	t.Parallel()
	params := DefaultParams()
	unknown := &openapi3.Parameter{
		Name: "futureRequired", In: openapi3.ParameterInQuery, Required: true,
	}
	if _, _, err := v2QueryValue("/api/v2/future", unknown, params); err == nil ||
		!strings.Contains(err.Error(), "unmapped required query parameter") {
		t.Fatalf("unknown required parameter error = %v", err)
	}
	pool := &openapi3.Parameter{
		Name: "pool", In: openapi3.ParameterInQuery, Required: true,
	}
	if _, _, err := v2QueryValue("/api/v2/future", pool, params); err == nil ||
		!strings.Contains(err.Error(), "has no string enum") {
		t.Fatalf("pool without enum error = %v", err)
	}
	optional := &openapi3.Parameter{
		Name: "futureOptional", In: openapi3.ParameterInQuery,
	}
	if value, include, err := v2QueryValue("/api/v2/future", optional, params); err != nil ||
		include || value != "" {
		t.Fatalf("optional parameter = %q/%v/%v", value, include, err)
	}
}

func TestBuildV2ProbeFailsClosedOnContractDrift(t *testing.T) {
	t.Parallel()
	spec := new(openapi3.T)
	params := DefaultParams()
	operation := func(id string, parameters ...*openapi3.ParameterRef) *openapi3.PathItem {
		return &openapi3.PathItem{Get: &openapi3.Operation{
			OperationID: id,
			Parameters:  parameters,
		}}
	}
	pathParameter := func(name string) *openapi3.ParameterRef {
		return &openapi3.ParameterRef{Value: &openapi3.Parameter{
			Name: name, In: openapi3.ParameterInPath, Required: true,
		}}
	}
	queryParameter := func(name string) *openapi3.ParameterRef {
		return &openapi3.ParameterRef{Value: &openapi3.Parameter{
			Name: name, In: openapi3.ParameterInQuery, Required: true,
		}}
	}

	tests := []struct {
		name     string
		template string
		item     *openapi3.PathItem
		want     string
	}{
		{"missing operation ID", "/api/v2/future", operation(""), "has no operationId"},
		{"deprecated path", "/api/cosmicgame/future", operation("future"), "uses deprecated path"},
		{"unresolved parameter", "/api/v2/future", operation("future", &openapi3.ParameterRef{}), "unresolved parameter"},
		{"unknown path parameter", "/api/v2/future/{future}", operation("future", pathParameter("future")), "unmapped path parameter"},
		{"unknown query parameter", "/api/v2/future", operation("future", queryParameter("future")), "unmapped required query parameter"},
		{"undeclared placeholder", "/api/v2/future/{future}", operation("future"), "unbound path parameter"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := buildV2Probe(spec, test.template, test.item, params)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestCheckProbeV2ContractAndD6Failures(t *testing.T) {
	t.Parallel()
	probes, err := BuildV2Probes(DefaultParams())
	if err != nil {
		t.Fatal(err)
	}
	var rounds Probe
	for _, probe := range probes {
		if probe.Template == "/api/v2/cosmicgame/rounds" {
			rounds = probe
			break
		}
	}
	if rounds.OperationID == "" {
		t.Fatal("rounds probe not found")
	}
	invalidRequest := rounds
	invalidRequest.Endpoint = strings.Replace(invalidRequest.Endpoint, "limit=1", "limit=invalid", 1)
	preflight := CheckProbe(
		context.Background(),
		testHTTPClientFunc(func(*http.Request) (*http.Response, error) {
			t.Fatal("client called after OpenAPI request validation failed")
			return nil, nil
		}),
		"https://api.example",
		invalidRequest,
		1024,
	)
	if !preflight.Failed || !strings.Contains(preflight.Reason, "request violates OpenAPI v2") {
		t.Fatalf("preflight result = %#v", preflight)
	}

	tests := []struct {
		name        string
		status      int
		contentType string
		body        string
		headers     http.Header
		wantFailed  bool
		wantReason  string
	}{
		{
			name:        "valid",
			status:      http.StatusOK,
			contentType: "application/json",
			body:        `{"data":[],"meta":{"limit":1}}`,
		},
		{
			name:        "schema invalid",
			status:      http.StatusOK,
			contentType: "application/json",
			body:        `{"data":"wrong","meta":{"limit":1}}`,
			wantFailed:  true,
			wantReason:  "violates OpenAPI v2",
		},
		{
			name:        "problem with 200",
			status:      http.StatusOK,
			contentType: "application/problem+json",
			body:        `{"type":"about:blank","title":"bad","status":500}`,
			wantFailed:  true,
			wantReason:  "application/problem+json",
		},
		{
			name:        "degraded",
			status:      http.StatusServiceUnavailable,
			contentType: "application/problem+json",
			body:        `{"type":"about:blank","title":"Unavailable","status":503}`,
			wantFailed:  true,
			wantReason:  "non-200",
		},
		{
			name:        "deprecated header",
			status:      http.StatusOK,
			contentType: "application/json",
			body:        `{"data":[],"meta":{"limit":1}}`,
			headers:     http.Header{"Deprecation": []string{"@1784160000"}},
			wantFailed:  true,
			wantReason:  "D6-unsafe Deprecation",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			client := testHTTPClientFunc(func(*http.Request) (*http.Response, error) {
				headers := test.headers.Clone()
				if headers == nil {
					headers = make(http.Header)
				}
				headers.Set("Content-Type", test.contentType)
				return &http.Response{
					StatusCode: test.status,
					Header:     headers,
					Body:       io.NopCloser(strings.NewReader(test.body)),
				}, nil
			})
			got := CheckProbe(context.Background(), client, "https://api.example", rounds, 1024)
			if got.Failed != test.wantFailed {
				t.Fatalf("result = %#v", got)
			}
			if test.wantReason != "" && !strings.Contains(got.Reason, test.wantReason) {
				t.Fatalf("reason = %q, want %q", got.Reason, test.wantReason)
			}
		})
	}
}

func FuzzBuildV2ProbePathEscaping(f *testing.F) {
	f.Add("0x1111111111111111111111111111111111111111")
	f.Add("slash/space ? percent%")
	f.Fuzz(func(t *testing.T, address string) {
		params := DefaultParams()
		params.UserAddress = address
		probes, err := BuildV2Probes(params)
		if err != nil {
			t.Fatal(err)
		}
		for _, probe := range probes {
			if strings.Contains(probe.Endpoint, "{") {
				t.Fatalf("unbound endpoint %q", probe.Endpoint)
			}
			if _, err := url.ParseRequestURI(probe.Endpoint); err != nil {
				t.Fatalf("invalid endpoint %q: %v", probe.Endpoint, err)
			}
		}
	})
}
