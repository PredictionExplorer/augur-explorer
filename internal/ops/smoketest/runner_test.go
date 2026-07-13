package smoketest

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type testHTTPClientFunc func(*http.Request) (*http.Response, error)

func (f testHTTPClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func response(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestBodyErrorStructuredDetection(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		body       string
		wantBad    bool
		wantReason string
	}{
		{"empty", "", false, ""},
		{"array carrying object error", `[{"error":"not top level"}]`, false, ""},
		{"ordinary object", `{"message":"the word error is harmless","status":1}`, false, ""},
		{"nested error", `{"data":{"error":"nested"}}`, false, ""},
		{"empty error", `{"error":""}`, false, ""},
		{"whitespace error", `{"error":"   "}`, false, ""},
		{"null error", `{"error":null}`, false, ""},
		{"false error", `{"error":false}`, false, ""},
		{"empty object error", `{"error":{}}`, false, ""},
		{"empty array error", `{"error":[]}`, false, ""},
		{"string error", `{"error":"database failed"}`, true, "error=database failed"},
		{"case insensitive error", `{"ERROR":"boom"}`, true, "error=boom"},
		{"deterministic duplicate casing", `{"error":"lower","ERROR":"upper"}`, true, "error=upper"},
		{"object error", `{"error":{"message":"boom"}}`, true, `error={"message":"boom"}`},
		{"zero status", "{\n  \"status\" : 0,\n  \"data\": []\n}", true, "status=0;"},
		{"float zero status", `{"status":0.0}`, true, "status=0;"},
		{"string zero status", `{"status":"0"}`, false, ""},
		{"nonzero status", `{"status":1}`, false, ""},
		{"malformed", `{"error":"boom"`, false, ""},
		{"trailing JSON", `{"error":"boom"} {"status":1}`, false, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			reason, bad := BodyError([]byte(test.body))
			if bad != test.wantBad {
				t.Fatalf("BodyError(%q) bad=%v reason=%q", test.body, bad, reason)
			}
			if test.wantReason != "" && !strings.Contains(reason, test.wantReason) {
				t.Fatalf("reason = %q, want substring %q", reason, test.wantReason)
			}
		})
	}
}

func TestCheckEndpointResponseCases(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		status     int
		body       string
		limit      int64
		wantFailed bool
		wantReason string
	}{
		{"200 array", http.StatusOK, `[1,2,3]`, 1024, false, ""},
		{"200 object", http.StatusOK, `{"status":1,"data":[]}`, 1024, false, ""},
		{"200 malformed", http.StatusOK, `{"status":`, 1024, false, ""},
		{"200 error", http.StatusOK, `{"error":"query failed"}`, 1024, true, "200 but error body: error=query failed"},
		{"200 status zero", http.StatusOK, `{"status":0}`, 1024, true, "200 but error body: status=0"},
		{"non-200", http.StatusBadRequest, "bad\nrequest", 1024, true, "non-200; body: bad request"},
		{"oversized", http.StatusOK, "123456789", 8, true, "response body exceeds 8 bytes"},
		{"exact limit", http.StatusOK, "12345678", 8, false, ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			client := testHTTPClientFunc(func(req *http.Request) (*http.Response, error) {
				if req.Method != http.MethodGet {
					t.Errorf("method = %s", req.Method)
				}
				return response(test.status, test.body), nil
			})
			got := CheckEndpoint(context.Background(), client, "https://api.example.test/path", test.limit)
			if got.StatusCode != test.status || got.Failed != test.wantFailed {
				t.Fatalf("result = %#v", got)
			}
			if test.wantReason != "" && !strings.Contains(got.Reason, test.wantReason) {
				t.Fatalf("reason = %q, want %q", got.Reason, test.wantReason)
			}
		})
	}
}

type errorReadCloser struct {
	err    error
	closed bool
}

func (r *errorReadCloser) Read([]byte) (int, error) { return 0, r.err }
func (r *errorReadCloser) Close() error {
	r.closed = true
	return nil
}

func TestCheckEndpointTransportReadAndRequestErrors(t *testing.T) {
	t.Parallel()
	t.Run("transport", func(t *testing.T) {
		want := errors.New("connection refused")
		result := CheckEndpoint(context.Background(), testHTTPClientFunc(func(*http.Request) (*http.Response, error) {
			return nil, want
		}), "https://api.example.test/path", 10)
		if !result.Failed || result.StatusCode != 0 || !strings.Contains(result.Reason, want.Error()) {
			t.Fatalf("result = %#v", result)
		}
	})
	t.Run("nil response", func(t *testing.T) {
		result := CheckEndpoint(context.Background(), testHTTPClientFunc(func(*http.Request) (*http.Response, error) {
			return nil, nil
		}), "https://api.example.test/path", 10)
		if !result.Failed || !strings.Contains(result.Reason, "nil response") {
			t.Fatalf("result = %#v", result)
		}
	})
	t.Run("body read", func(t *testing.T) {
		want := errors.New("body interrupted")
		body := &errorReadCloser{err: want}
		result := CheckEndpoint(context.Background(), testHTTPClientFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: http.StatusOK, Body: body}, nil
		}), "https://api.example.test/path", 10)
		if !result.Failed || !strings.Contains(result.Reason, want.Error()) || !body.closed {
			t.Fatalf("result=%#v closed=%v", result, body.closed)
		}
	})
	t.Run("malformed URL", func(t *testing.T) {
		result := CheckEndpoint(context.Background(), testHTTPClientFunc(func(*http.Request) (*http.Response, error) {
			t.Fatal("client should not be called")
			return nil, nil
		}), "://bad-url", 10)
		if !result.Failed || !strings.Contains(result.Reason, "request error") {
			t.Fatalf("result = %#v", result)
		}
	})
}

func TestCheckEndpointPropagatesRequestContext(t *testing.T) {
	t.Parallel()
	type contextKey struct{}
	ctx := context.WithValue(context.Background(), contextKey{}, "marker")
	client := testHTTPClientFunc(func(req *http.Request) (*http.Response, error) {
		if got := req.Context().Value(contextKey{}); got != "marker" {
			t.Fatalf("context value = %#v", got)
		}
		return response(http.StatusOK, `[]`), nil
	})
	if result := CheckEndpoint(ctx, client, "https://api.example.test/path", 10); result.Failed {
		t.Fatalf("result = %#v", result)
	}
}

func TestRunStableSummaryAndOutput(t *testing.T) {
	t.Parallel()
	run := func() (Summary, string, []string, error) {
		var requested []string
		client := testHTTPClientFunc(func(req *http.Request) (*http.Response, error) {
			requested = append(requested, req.URL.String())
			switch req.URL.Path {
			case "/api/cosmicgame/statistics/counters":
				return response(http.StatusServiceUnavailable, "temporarily unavailable"), nil
			case "/api/cosmicgame/time/current":
				return response(http.StatusOK, `{"status":0,"error":""}`), nil
			default:
				return response(http.StatusOK, `[]`), nil
			}
		})
		var out bytes.Buffer
		summary, err := Run(context.Background(), Options{
			Source:  fakeParameterSource{params: Params{UserAddress: "0xUser", RoundNumber: "8"}},
			Client:  client,
			BaseURL: "https://api.example.test/",
			Output:  &out,
		})
		return summary, out.String(), requested, err
	}

	first, firstOutput, firstRequests, firstErr := run()
	second, secondOutput, secondRequests, secondErr := run()
	var firstFailuresErr, secondFailuresErr FailuresError
	if !errors.As(firstErr, &firstFailuresErr) || !errors.As(secondErr, &secondFailuresErr) ||
		firstFailuresErr.Count != 2 || secondFailuresErr.Count != 2 {
		t.Fatalf("errors = %v / %v", firstErr, secondErr)
	}
	if !reflect.DeepEqual(first, second) || firstOutput != secondOutput || !reflect.DeepEqual(firstRequests, secondRequests) {
		t.Fatal("identical runs produced unstable results")
	}
	if first.Total != 142 || first.OK != 140 || len(first.Failures) != 2 {
		t.Fatalf("summary = %#v", first)
	}
	if first.Failures[0].Endpoint != "/api/cosmicgame/statistics/counters" ||
		first.Failures[1].Endpoint != "/api/cosmicgame/time/current" {
		t.Fatalf("failure order = %#v", first.Failures)
	}
	if len(firstRequests) != 142 ||
		firstRequests[0] != "https://api.example.test/api/cosmicgame/statistics/dashboard" {
		t.Fatalf("requests = %#v", firstRequests)
	}
	for _, want := range []string{
		"API base : https://api.example.test",
		"FAILED",
		"[503] /api/cosmicgame/statistics/counters",
		"[200] /api/cosmicgame/time/current",
		"Total: 142",
		"FAILED: 2",
	} {
		if !strings.Contains(firstOutput, want) {
			t.Errorf("output missing %q", want)
		}
	}
}

func TestRunAllSuccessful(t *testing.T) {
	t.Parallel()
	client := testHTTPClientFunc(func(*http.Request) (*http.Response, error) {
		return response(http.StatusOK, `{"status":1}`), nil
	})
	var out bytes.Buffer
	summary, err := Run(context.Background(), Options{
		Source: fakeParameterSource{}, Client: client, BaseURL: "https://api.example.test", Output: &out,
	})
	if err != nil {
		t.Fatal(err)
	}
	if summary.Total != 142 || summary.OK != 142 || len(summary.Failures) != 0 {
		t.Fatalf("summary = %#v", summary)
	}
	if !strings.Contains(out.String(), "All endpoints returned 200 with no error body.") {
		t.Fatalf("output missing success message")
	}
}

func TestRunCancellationStopsWithoutFailureSummary(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	requests := 0
	client := testHTTPClientFunc(func(req *http.Request) (*http.Response, error) {
		requests++
		cancel()
		<-req.Context().Done()
		return nil, req.Context().Err()
	})
	summary, err := Run(ctx, Options{
		Source: fakeParameterSource{}, Client: client, BaseURL: "https://api.example.test",
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("summary=%#v error=%v", summary, err)
	}
	if requests != 1 || summary.OK != 0 || len(summary.Failures) != 0 {
		t.Fatalf("requests=%d summary=%#v", requests, summary)
	}
}

func TestRunValidationAndSourceErrors(t *testing.T) {
	t.Parallel()
	client := testHTTPClientFunc(func(*http.Request) (*http.Response, error) {
		return response(http.StatusOK, `[]`), nil
	})
	sourceErr := errors.New("database unavailable")
	tests := []struct {
		name string
		opts Options
	}{
		{"nil source", Options{Client: client, BaseURL: "https://api.example.test"}},
		{"nil client", Options{Source: fakeParameterSource{}, BaseURL: "https://api.example.test"}},
		{"empty base", Options{Source: fakeParameterSource{}, Client: client}},
		{"source error", Options{Source: fakeParameterSource{err: sourceErr}, Client: client, BaseURL: "https://api.example.test"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Run(context.Background(), test.opts)
			if err == nil {
				t.Fatal("expected error")
			}
			if test.name == "source error" && !errors.Is(err, sourceErr) {
				t.Fatalf("error = %v", err)
			}
		})
	}
}
