package smoketest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[1;31m"
	colorGreen = "\033[32m"

	// DefaultMaxBodyBytes bounds each response body read to one MiB.
	DefaultMaxBodyBytes int64 = 1 << 20
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
	StatusCode int
	Endpoint   string
	Reason     string
}

// Summary is the deterministic result of a smoke-test run.
type Summary struct {
	Total    int
	OK       int
	Failures []Failure
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
	Source       ParameterSource
	Client       HTTPClient
	BaseURL      string
	Output       io.Writer
	MaxBodyBytes int64
}

// Run loads parameters, checks every endpoint in stable order, and writes the
// legacy colored opsctl report.
func Run(ctx context.Context, opts Options) (Summary, error) {
	var summary Summary
	if err := ctx.Err(); err != nil {
		return summary, err
	}
	if opts.Source == nil {
		return summary, fmt.Errorf("parameter source is nil")
	}
	if opts.Client == nil {
		return summary, fmt.Errorf("HTTP client is nil")
	}
	baseURL := strings.TrimRight(strings.TrimSpace(opts.BaseURL), "/")
	if baseURL == "" {
		return summary, fmt.Errorf("API base URL is empty")
	}
	out := opts.Output
	if out == nil {
		out = io.Discard
	}
	maxBodyBytes := opts.MaxBodyBytes
	if maxBodyBytes <= 0 {
		maxBodyBytes = DefaultMaxBodyBytes
	}

	params, err := opts.Source.Parameters(ctx)
	if err != nil {
		return summary, fmt.Errorf("load smoketest parameters: %w", err)
	}
	params = WithDefaults(params)
	endpoints := BuildEndpoints(params)
	summary.Total = len(endpoints)

	fmt.Fprintf(out, "API base : %s\n", baseURL)
	fmt.Fprintf(out, "Params   : userAddr=%s round=%s bidEvtlog=%s tokenId=%s nftDonId=%s erc20DonId=%s cstAction=%s rwalkAction=%s deposit=%s\n\n",
		params.UserAddress, params.RoundNumber, params.BidEventLogID, params.TokenID, params.NFTDonationID,
		params.ERC20DonationID, params.CSTActionID, params.RandomWalkActionID, params.DepositID)

	for _, endpoint := range endpoints {
		if err := ctx.Err(); err != nil {
			return summary, err
		}
		result := CheckEndpoint(ctx, opts.Client, baseURL+endpoint, maxBodyBytes)
		if err := ctx.Err(); err != nil {
			return summary, err
		}
		if result.Failed {
			fmt.Fprintf(out, "%sFAILED%s [%s] %s  %s\n",
				colorRed, colorReset, statusString(result.StatusCode), endpoint, result.Reason)
			summary.Failures = append(summary.Failures, Failure{
				StatusCode: result.StatusCode,
				Endpoint:   endpoint,
				Reason:     result.Reason,
			})
			continue
		}
		summary.OK++
		fmt.Fprintf(out, "%sOK%s     [200] %s\n", colorGreen, colorReset, endpoint)
	}

	fmt.Fprintln(out)
	fmt.Fprintln(out, "==================== SUMMARY ====================")
	fmt.Fprintf(out, "Total: %d   %sOK: %d%s   %sFAILED: %d%s\n",
		summary.Total, colorGreen, summary.OK, colorReset, colorRed, len(summary.Failures), colorReset)
	if len(summary.Failures) > 0 {
		fmt.Fprintf(out, "\n%sFailures:%s\n", colorRed, colorReset)
		for _, failure := range summary.Failures {
			fmt.Fprintf(out, "  %s[%s] %s  %s%s\n", colorRed, statusString(failure.StatusCode),
				failure.Endpoint, failure.Reason, colorReset)
		}
		return summary, FailuresError{Count: len(summary.Failures)}
	}
	fmt.Fprintf(out, "%sAll endpoints returned 200 with no error body.%s\n", colorGreen, colorReset)
	return summary, nil
}

// CheckEndpoint performs one context-bound GET with a bounded body read.
func CheckEndpoint(ctx context.Context, client HTTPClient, fullURL string, maxBodyBytes int64) EndpointResult {
	if maxBodyBytes <= 0 {
		maxBodyBytes = DefaultMaxBodyBytes
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return EndpointResult{Failed: true, Reason: "request error: " + err.Error()}
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
	oversized := int64(len(body)) > maxBodyBytes
	if oversized {
		body = body[:maxBodyBytes]
	}
	if resp.StatusCode != http.StatusOK {
		return EndpointResult{
			StatusCode: resp.StatusCode,
			Failed:     true,
			Reason:     "non-200; body: " + snippet(body),
		}
	}
	if oversized {
		return EndpointResult{
			StatusCode: resp.StatusCode,
			Failed:     true,
			Reason:     fmt.Sprintf("response body exceeds %d bytes", maxBodyBytes),
		}
	}
	if reason, ok := BodyError(body); ok {
		return EndpointResult{
			StatusCode: resp.StatusCode,
			Failed:     true,
			Reason:     "200 but error body: " + reason,
		}
	}
	return EndpointResult{StatusCode: resp.StatusCode}
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
	sort.Strings(keys)
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
