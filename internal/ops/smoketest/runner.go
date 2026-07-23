package smoketest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[1;31m"
	colorGreen = "\033[32m"

	// DefaultMaxBodyBytes bounds each response body read to one MiB.
	DefaultMaxBodyBytes int64 = 1 << 20
	// DefaultProbeInterval respects the public API's sustained 50 rps
	// limiter, preventing a local smoke run from creating its own 429s.
	DefaultProbeInterval = 20 * time.Millisecond
)

// HTTPClient is the request surface required by the smoke-test engine.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// EndpointResult is the result of checking one endpoint.
type EndpointResult struct {
	StatusCode int
	Failed     bool
	Reason     string
}

// Failure records one failed endpoint in stable execution order.
type Failure struct {
	Suite       Suite
	OperationID string
	StatusCode  int
	Endpoint    string
	Reason      string
}

// SuiteSummary is the deterministic result for one selected API surface.
type SuiteSummary struct {
	Suite    Suite
	Total    int
	OK       int
	Failures []Failure
}

// Summary is the aggregate deterministic result of a smoke-test run.
type Summary struct {
	Total    int
	OK       int
	Failures []Failure
	Suites   []SuiteSummary
}

// FailuresError reports that one or more endpoints failed.
type FailuresError struct {
	Count int
}

func (e FailuresError) Error() string {
	return fmt.Sprintf("%d endpoint(s) failed", e.Count)
}

// Options configures Run.
type Options struct {
	Source        ParameterSource
	Client        HTTPClient
	BaseURL       string
	Output        io.Writer
	MaxBodyBytes  int64
	Suite         Suite
	ProbeInterval time.Duration
	DisablePacing bool
}

// Run loads parameters when needed, checks the selected suites in stable
// order, and writes a deterministic opsctl report.
func Run(ctx context.Context, opts Options) (Summary, error) {
	var summary Summary
	if err := ctx.Err(); err != nil {
		return summary, err
	}
	if opts.Client == nil {
		return summary, errors.New("HTTP client is nil")
	}
	baseURL := strings.TrimRight(strings.TrimSpace(opts.BaseURL), "/")
	if baseURL == "" {
		return summary, errors.New("API base URL is empty")
	}
	suite := opts.Suite
	if suite == "" {
		suite = SuiteV2
	}
	members, err := suite.members()
	if err != nil {
		return summary, err
	}

	params := DefaultParams()
	if suite.RequiresParameters() {
		if opts.Source == nil {
			return summary, errors.New("parameter source is nil")
		}
		params, err = opts.Source.Parameters(ctx)
		if err != nil {
			return summary, fmt.Errorf("load smoketest parameters: %w", err)
		}
		params = WithDefaults(params)
	}

	out := opts.Output
	if out == nil {
		out = io.Discard
	}
	maxBodyBytes := opts.MaxBodyBytes
	if maxBodyBytes <= 0 {
		maxBodyBytes = DefaultMaxBodyBytes
	}
	interval := opts.ProbeInterval
	if !opts.DisablePacing && interval <= 0 {
		interval = DefaultProbeInterval
	}
	if opts.DisablePacing {
		interval = 0
	}

	fmt.Fprintf(out, "API base : %s\n", baseURL)
	fmt.Fprintf(out, "Suite    : %s\n", suite)
	if suite.RequiresParameters() {
		fmt.Fprintf(out, "Params   : userAddr=%s round=%s bidEvtlog=%s tokenId=%s rwalkTokenId=%s nftDonId=%s erc20DonId=%s cstAction=%s rwalkAction=%s deposit=%s\n",
			params.UserAddress, params.RoundNumber, params.BidEventLogID, params.TokenID, params.RandomWalkTokenID,
			params.NFTDonationID, params.ERC20DonationID, params.CSTActionID, params.RandomWalkActionID, params.DepositID)
	}

	requestsStarted := 0
	for _, member := range members {
		probes, buildErr := buildProbes(member, params)
		if buildErr != nil {
			return summary, fmt.Errorf("build %s smoke-test probes: %w", member, buildErr)
		}
		fmt.Fprintf(out, "\n==================== %s ====================\n", strings.ToUpper(string(member)))
		suiteSummary := SuiteSummary{Suite: member, Total: len(probes)}
		for _, probe := range probes {
			if err := ctx.Err(); err != nil {
				return summary, err
			}
			if requestsStarted > 0 && interval > 0 {
				if err := waitProbeInterval(ctx, interval); err != nil {
					return summary, err
				}
			}
			requestsStarted++

			result := CheckProbe(ctx, opts.Client, baseURL, probe, maxBodyBytes)
			if err := ctx.Err(); err != nil {
				return summary, err
			}
			if result.Failed {
				fmt.Fprintf(out, "%sFAILED%s [%s] %s  %s\n",
					colorRed, colorReset, statusString(result.StatusCode), probe.Endpoint, result.Reason)
				failure := Failure{
					Suite:       member,
					OperationID: probe.OperationID,
					StatusCode:  result.StatusCode,
					Endpoint:    probe.Endpoint,
					Reason:      result.Reason,
				}
				suiteSummary.Failures = append(suiteSummary.Failures, failure)
				summary.Failures = append(summary.Failures, failure)
				continue
			}
			suiteSummary.OK++
			fmt.Fprintf(out, "%sOK%s     [200] %s\n", colorGreen, colorReset, probe.Endpoint)
		}
		summary.Total += suiteSummary.Total
		summary.OK += suiteSummary.OK
		summary.Suites = append(summary.Suites, suiteSummary)
	}

	fmt.Fprintln(out)
	fmt.Fprintln(out, "==================== SUMMARY ====================")
	fmt.Fprintf(out, "Total: %d   %sOK: %d%s   %sFAILED: %d%s\n",
		summary.Total, colorGreen, summary.OK, colorReset, colorRed, len(summary.Failures), colorReset)
	if len(summary.Failures) > 0 {
		fmt.Fprintf(out, "\n%sFailures:%s\n", colorRed, colorReset)
		for _, failure := range summary.Failures {
			operation := ""
			if failure.OperationID != "" {
				operation = " (" + failure.OperationID + ")"
			}
			fmt.Fprintf(out, "  %s[%s] %s %s%s  %s%s\n",
				colorRed, statusString(failure.StatusCode), failure.Suite, failure.Endpoint,
				operation, failure.Reason, colorReset)
		}
		return summary, FailuresError{Count: len(summary.Failures)}
	}
	fmt.Fprintf(out, "%sAll selected endpoints returned contract-valid HTTP 200 responses.%s\n", colorGreen, colorReset)
	return summary, nil
}

func waitProbeInterval(ctx context.Context, interval time.Duration) error {
	timer := time.NewTimer(interval)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// CheckEndpoint retains the frozen v1 one-endpoint API used by existing
// callers and tests.
func CheckEndpoint(ctx context.Context, client HTTPClient, fullURL string, maxBodyBytes int64) EndpointResult {
	baseURL, endpoint := splitFullURL(fullURL)
	return CheckProbe(ctx, client, baseURL, Probe{
		Suite:    SuiteV1,
		Template: endpoint,
		Endpoint: endpoint,
	}, maxBodyBytes)
}

func splitFullURL(fullURL string) (string, string) {
	index := strings.Index(fullURL, "://")
	if index < 0 {
		return "", fullURL
	}
	pathStart := strings.Index(fullURL[index+3:], "/")
	if pathStart < 0 {
		return fullURL, ""
	}
	pathStart += index + 3
	return fullURL[:pathStart], fullURL[pathStart:]
}

// CheckProbe performs one context-bound GET with a bounded body read and the
// suite-specific response contract.
func CheckProbe(
	ctx context.Context,
	client HTTPClient,
	baseURL string,
	probe Probe,
	maxBodyBytes int64,
) EndpointResult {
	if maxBodyBytes <= 0 {
		maxBodyBytes = DefaultMaxBodyBytes
	}
	req, err := http.NewRequestWithContext(ctx, probe.method(), baseURL+probe.Endpoint, nil)
	if err != nil {
		return EndpointResult{Failed: true, Reason: "request error: " + err.Error()}
	}
	if probe.operation != nil {
		if err := validateV2Request(ctx, req, probe); err != nil {
			return EndpointResult{Failed: true, Reason: "request violates OpenAPI v2: " + err.Error()}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return EndpointResult{Failed: true, Reason: "request error: " + err.Error()}
	}
	if resp == nil {
		return EndpointResult{Failed: true, Reason: "request error: HTTP client returned a nil response"}
	}
	if resp.Body == nil {
		resp.Body = http.NoBody
	}
	defer func() { _ = resp.Body.Close() }()

	body, readErr := io.ReadAll(io.LimitReader(resp.Body, maxBodyBytes+1))
	if readErr != nil {
		return EndpointResult{
			StatusCode: resp.StatusCode,
			Failed:     true,
			Reason:     "body read error: " + readErr.Error(),
		}
	}
	if int64(len(body)) > maxBodyBytes {
		return EndpointResult{
			StatusCode: resp.StatusCode,
			Failed:     true,
			Reason:     fmt.Sprintf("response body exceeds %d bytes", maxBodyBytes),
		}
	}
	if probe.Suite != SuiteV1 {
		for _, header := range []string{"Deprecation", "Sunset"} {
			if value := resp.Header.Get(header); value != "" {
				return EndpointResult{
					StatusCode: resp.StatusCode,
					Failed:     true,
					Reason:     fmt.Sprintf("D6-unsafe %s header: %s", header, value),
				}
			}
		}
	}
	if resp.StatusCode != http.StatusOK {
		return EndpointResult{
			StatusCode: resp.StatusCode,
			Failed:     true,
			Reason:     "non-200; body: " + snippet(body),
		}
	}
	if probe.Suite == SuiteV1 {
		if reason, ok := BodyError(body); ok {
			return EndpointResult{
				StatusCode: resp.StatusCode,
				Failed:     true,
				Reason:     "200 but error body: " + reason,
			}
		}
		return EndpointResult{StatusCode: resp.StatusCode}
	}
	if strings.HasPrefix(resp.Header.Get("Content-Type"), "application/problem+json") {
		return EndpointResult{
			StatusCode: resp.StatusCode,
			Failed:     true,
			Reason:     "200 carried application/problem+json: " + snippet(body),
		}
	}
	if probe.operation != nil {
		if err := validateV2Response(ctx, req, resp, body, probe); err != nil {
			return EndpointResult{
				StatusCode: resp.StatusCode,
				Failed:     true,
				Reason:     "response violates OpenAPI v2: " + err.Error(),
			}
		}
	}
	return EndpointResult{StatusCode: resp.StatusCode}
}

func v2ValidationInput(req *http.Request, probe Probe) *openapi3filter.RequestValidationInput {
	return &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: probe.PathParams,
		Route: &routers.Route{
			Spec:      probe.spec,
			Path:      probe.Template,
			PathItem:  probe.pathItem,
			Method:    http.MethodGet,
			Operation: probe.operation,
		},
		Options: &openapi3filter.Options{
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		},
	}
}

func validateV2Request(ctx context.Context, req *http.Request, probe Probe) error {
	return openapi3filter.ValidateRequest(ctx, v2ValidationInput(req, probe))
}

func validateV2Response(
	ctx context.Context,
	req *http.Request,
	resp *http.Response,
	body []byte,
	probe Probe,
) error {
	input := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: v2ValidationInput(req, probe),
		Status:                 resp.StatusCode,
		Header:                 resp.Header,
	}
	input.SetBodyBytes(body)
	return openapi3filter.ValidateResponse(ctx, input)
}

// BodyError detects top-level JSON objects carrying a non-empty error value or
// a numeric status value equal to zero. Arrays, malformed JSON, and ordinary
// objects are not classified as failures.
func BodyError(body []byte) (string, bool) {
	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" || trimmed[0] != '{' {
		return "", false
	}
	var object map[string]any
	decoder := json.NewDecoder(strings.NewReader(trimmed))
	decoder.UseNumber()
	if err := decoder.Decode(&object); err != nil {
		return "", false
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return "", false
	}

	keys := make([]string, 0, len(object))
	for key := range object {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	for _, key := range keys {
		if strings.EqualFold(key, "error") {
			if rendered, ok := nonEmptyError(object[key]); ok {
				return "error=" + rendered, true
			}
		}
	}
	for _, key := range keys {
		if !strings.EqualFold(key, "status") {
			continue
		}
		number, ok := object[key].(json.Number)
		if !ok {
			continue
		}
		parsed, err := strconv.ParseFloat(number.String(), 64)
		if err == nil && parsed == 0 {
			return "status=0; " + snippet(body), true
		}
	}
	return "", false
}

func nonEmptyError(value any) (string, bool) {
	switch typed := value.(type) {
	case nil:
		return "", false
	case string:
		typed = strings.TrimSpace(typed)
		return typed, typed != ""
	case bool:
		if !typed {
			return "", false
		}
	case json.Number:
		if typed.String() == "0" {
			return "", false
		}
	case []any:
		if len(typed) == 0 {
			return "", false
		}
	case map[string]any:
		if len(typed) == 0 {
			return "", false
		}
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		return fmt.Sprint(value), true
	}
	return string(encoded), true
}

func snippet(body []byte) string {
	value := strings.TrimSpace(string(body))
	value = strings.ReplaceAll(value, "\n", " ")
	if len(value) > 160 {
		value = value[:160] + "..."
	}
	return value
}

func statusString(code int) string {
	if code == 0 {
		return "ERR"
	}
	return strconv.Itoa(code)
}
