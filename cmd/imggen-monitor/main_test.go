package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func imEnv(base string) map[string]string {
	return map[string]string{
		"IM_REQUEST_URL": base + "/generate",
		"IM_IMAGE_URL":   base + "/images/",
		"IM_VIDEO_URL":   base + "/videos/",
	}
}

func envFunc(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

func TestRunFlagAndConfigErrors(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("bad flag", func(t *testing.T) {
		t.Parallel()
		var out, errOut strings.Builder
		err := run(ctx, []string{"-nonsense"}, envFunc(nil), &out, &errOut)
		if err == nil {
			t.Fatal("bad flag must error")
		}
	})

	t.Run("missing env", func(t *testing.T) {
		t.Parallel()
		var out, errOut strings.Builder
		err := run(ctx, nil, envFunc(nil), &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "IM_REQUEST_URL") {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("token without seed", func(t *testing.T) {
		t.Parallel()
		var out, errOut strings.Builder
		err := run(ctx, []string{"-token", "5"}, envFunc(imEnv("http://127.0.0.1:1")), &out, &errOut)
		if err == nil || !strings.Contains(err.Error(), "-token requires -seed") {
			t.Fatalf("err = %v", err)
		}
	})
}

func TestRunOneShotGeneration(t *testing.T) {
	t.Parallel()
	var mu sync.Mutex
	var payloads []map[string]interface{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/generate" || r.Method != http.MethodPost {
			http.NotFound(w, r)
			return
		}
		var payload map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		mu.Lock()
		payloads = append(payloads, payload)
		mu.Unlock()
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	var out, errOut strings.Builder
	err := run(context.Background(), []string{"-token", "123", "-seed", "0xabc"},
		envFunc(imEnv(server.URL)), &out, &errOut)
	if err != nil {
		t.Fatal(err)
	}

	mu.Lock()
	defer mu.Unlock()
	if len(payloads) != 1 || payloads[0]["token_id"] != float64(123) || payloads[0]["seed"] != "0xabc" {
		t.Fatalf("payloads = %v", payloads)
	}
}

func TestRunOneShotGenerationFailure(t *testing.T) {
	t.Parallel()
	var out, errOut strings.Builder
	err := run(context.Background(), []string{"-token", "1", "-seed", "s"},
		envFunc(imEnv("http://127.0.0.1:1")), &out, &errOut)
	if err == nil || !strings.Contains(err.Error(), "submitting generation request") {
		t.Fatalf("err = %v", err)
	}
}
