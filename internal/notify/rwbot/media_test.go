package rwbot

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPFetcherReturnsBodyOn200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/img/000042_black.png" {
			t.Errorf("path = %q", r.URL.Path)
		}
		_, _ = w.Write([]byte("png-bytes"))
	}))
	defer srv.Close()

	status, body, err := HTTPFetcher{}.Fetch(context.Background(), srv.URL+"/img/000042_black.png")
	if err != nil {
		t.Fatalf("Fetch: %v", err)
	}
	if status != 200 || string(body) != "png-bytes" {
		t.Errorf("Fetch = (%d, %q)", status, body)
	}
}

func TestHTTPFetcherReturnsStatusWithoutBodyOnNon200(t *testing.T) {
	for _, code := range []int{403, 404, 500} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(code)
			_, _ = w.Write([]byte("error page"))
		}))
		status, body, err := HTTPFetcher{Client: srv.Client()}.Fetch(context.Background(), srv.URL)
		srv.Close()
		if err != nil {
			t.Fatalf("Fetch(%d): %v", code, err)
		}
		if status != code || body != nil {
			t.Errorf("Fetch(%d) = (%d, %q), want status only", code, status, body)
		}
	}
}

func TestHTTPFetcherTransportError(t *testing.T) {
	_, _, err := HTTPFetcher{}.Fetch(context.Background(), "http://127.0.0.1:1/nope")
	if err == nil {
		t.Error("unreachable host produced no error")
	}
}

func TestHTTPFetcherHonorsContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("late"))
	}))
	defer srv.Close()
	if _, _, err := (HTTPFetcher{}).Fetch(ctx, srv.URL); err == nil {
		t.Error("cancelled context produced no error")
	}
}

func TestHTTPFetcherRejectsBadURL(t *testing.T) {
	if _, _, err := (HTTPFetcher{}).Fetch(context.Background(), "http://\x00bad"); err == nil {
		t.Error("bad URL produced no error")
	}
}

func TestHTTPFetcherBodyReadFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// Declare a body longer than what is sent, then abort: the client's
		// read fails mid-body.
		w.Header().Set("Content-Length", "1000")
		_, _ = w.Write([]byte("partial"))
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		panic(http.ErrAbortHandler)
	}))
	defer srv.Close()

	_, _, err := (HTTPFetcher{Client: srv.Client()}).Fetch(context.Background(), srv.URL)
	if err == nil {
		t.Error("aborted body read produced no error")
	}
}
