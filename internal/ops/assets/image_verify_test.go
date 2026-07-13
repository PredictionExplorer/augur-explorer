package assets

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type fakeTokenCountSource struct {
	count int64
	err   error
}

func (f fakeTokenCountSource) MintedTokenCount(context.Context) (int64, error) {
	return f.count, f.err
}

type httpClientFunc func(*http.Request) (*http.Response, error)

func (f httpClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

type trackingBody struct {
	read   int
	closed bool
}

func (b *trackingBody) Read([]byte) (int, error) {
	b.read++
	return 0, io.EOF
}

func (b *trackingBody) Close() error {
	b.closed = true
	return nil
}

func TestTokenImageURL(t *testing.T) {
	t.Parallel()
	if got, want := TokenImageURL("https://images.example/root/", 42), "https://images.example/root/000042_black.png"; got != want {
		t.Fatalf("TokenImageURL() = %q, want %q", got, want)
	}
}

func TestCheckTokenImageIsStatusOnly(t *testing.T) {
	t.Parallel()
	body := &trackingBody{}
	client := httpClientFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet {
			t.Errorf("method = %s", req.Method)
		}
		if req.Context() == nil {
			t.Error("request has no context")
		}
		return &http.Response{StatusCode: http.StatusOK, Body: body}, nil
	})
	status, err := CheckTokenImage(context.Background(), client, "https://images.example/000001_black.png")
	if err != nil || status != http.StatusOK {
		t.Fatalf("status=%d error=%v", status, err)
	}
	if body.read != 0 {
		t.Fatalf("response body was read %d time(s)", body.read)
	}
	if !body.closed {
		t.Fatal("response body was not closed")
	}
}

func TestCheckTokenImageRejectsMalformedURL(t *testing.T) {
	t.Parallel()
	client := httpClientFunc(func(*http.Request) (*http.Response, error) {
		t.Fatal("client should not be called")
		return nil, nil
	})
	if _, err := CheckTokenImage(context.Background(), client, "://bad-url"); err == nil {
		t.Fatal("malformed URL unexpectedly succeeded")
	}
}

func TestCheckTokenImageHandlesIncompleteClientResponses(t *testing.T) {
	t.Parallel()
	t.Run("nil response", func(t *testing.T) {
		client := httpClientFunc(func(*http.Request) (*http.Response, error) { return nil, nil })
		if _, err := CheckTokenImage(context.Background(), client, "https://images.example/image.png"); err == nil {
			t.Fatal("nil response unexpectedly succeeded")
		}
	})
	t.Run("nil body", func(t *testing.T) {
		client := httpClientFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: http.StatusNoContent}, nil
		})
		status, err := CheckTokenImage(context.Background(), client, "https://images.example/image.png")
		if err != nil || status != http.StatusNoContent {
			t.Fatalf("status=%d error=%v", status, err)
		}
	})
}

func TestVerifyTokenImagesStatusesAndStableSummary(t *testing.T) {
	t.Parallel()
	var paths []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		switch r.URL.Path {
		case "/images/000000_black.png":
			w.WriteHeader(http.StatusOK)
		case "/images/000001_black.png":
			w.WriteHeader(http.StatusForbidden)
		default:
			w.WriteHeader(http.StatusBadGateway)
		}
		_, _ = w.Write([]byte("body is intentionally ignored"))
	}))
	defer server.Close()

	var logs bytes.Buffer
	summary, err := VerifyTokenImages(context.Background(), VerifyTokenImagesOptions{
		Source:  fakeTokenCountSource{count: 3},
		Client:  server.Client(),
		BaseURL: server.URL + "/images/",
		Logger:  log.New(&logs, "", 0),
	})
	if err != nil {
		t.Fatal(err)
	}
	if summary.Tokens != 3 || summary.Checked != 3 || summary.OK != 1 {
		t.Fatalf("summary = %#v", summary)
	}
	if !reflect.DeepEqual(summary.Forbidden, []int64{1}) {
		t.Fatalf("forbidden = %#v", summary.Forbidden)
	}
	if !reflect.DeepEqual(summary.OtherStatuses, []TokenHTTPStatus{{TokenID: 2, Status: http.StatusBadGateway}}) {
		t.Fatalf("other statuses = %#v", summary.OtherStatuses)
	}
	wantPaths := []string{
		"/images/000000_black.png",
		"/images/000001_black.png",
		"/images/000002_black.png",
	}
	if !reflect.DeepEqual(paths, wantPaths) {
		t.Fatalf("paths = %#v, want %#v", paths, wantPaths)
	}
	for _, want := range []string{
		"num_tokens = 3",
		"Image server returns 403 code for token 1",
		"token 1: FAIL",
		"Process ended, failed tokens: 1",
	} {
		if !strings.Contains(logs.String(), want) {
			t.Errorf("logs missing %q:\n%s", want, logs.String())
		}
	}
}

func TestVerifyTokenImagesTransportErrorContinues(t *testing.T) {
	t.Parallel()
	transportErr := errors.New("network down")
	client := httpClientFunc(func(req *http.Request) (*http.Response, error) {
		if strings.Contains(req.URL.Path, "000000") {
			return nil, transportErr
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("ignored")),
		}, nil
	})
	summary, err := VerifyTokenImages(context.Background(), VerifyTokenImagesOptions{
		Source: fakeTokenCountSource{count: 2}, Client: client, BaseURL: "https://images.example",
	})
	if err != nil {
		t.Fatal(err)
	}
	if summary.Checked != 1 || summary.OK != 1 || len(summary.RequestErrors) != 1 {
		t.Fatalf("summary = %#v", summary)
	}
	if summary.RequestErrors[0].TokenID != 0 || !errors.Is(summary.RequestErrors[0].Err, transportErr) {
		t.Fatalf("request error = %#v", summary.RequestErrors[0])
	}
}

func TestVerifyTokenImagesCancellation(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	client := httpClientFunc(func(req *http.Request) (*http.Response, error) {
		cancel()
		<-req.Context().Done()
		return nil, req.Context().Err()
	})
	summary, err := VerifyTokenImages(ctx, VerifyTokenImagesOptions{
		Source: fakeTokenCountSource{count: 2}, Client: client, BaseURL: "https://images.example",
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("summary=%#v error=%v", summary, err)
	}
	if len(summary.RequestErrors) != 0 {
		t.Fatalf("cancellation recorded as request error: %#v", summary)
	}
}

func TestVerifyTokenImagesValidationAndSourceErrors(t *testing.T) {
	t.Parallel()
	client := httpClientFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("unexpected request")
	})
	tests := []struct {
		name string
		opts VerifyTokenImagesOptions
	}{
		{"nil source", VerifyTokenImagesOptions{Client: client, BaseURL: "https://example.test"}},
		{"nil client", VerifyTokenImagesOptions{Source: fakeTokenCountSource{}, BaseURL: "https://example.test"}},
		{"empty base", VerifyTokenImagesOptions{Source: fakeTokenCountSource{}, Client: client}},
		{"negative count", VerifyTokenImagesOptions{Source: fakeTokenCountSource{count: -1}, Client: client, BaseURL: "https://example.test"}},
		{"source error", VerifyTokenImagesOptions{Source: fakeTokenCountSource{err: errors.New("stats failed")}, Client: client, BaseURL: "https://example.test"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if _, err := VerifyTokenImages(context.Background(), test.opts); err == nil {
				t.Fatal("expected error")
			}
		})
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := VerifyTokenImages(ctx, VerifyTokenImagesOptions{}); !errors.Is(err, context.Canceled) {
		t.Fatalf("pre-canceled error = %v", err)
	}
}
