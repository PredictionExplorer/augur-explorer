package traits

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewFetcherValidatesBase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		base    string
		want    string
		wantErr bool
	}{
		{"https host", "https://nfts.cosmicsignature.com", "https://nfts.cosmicsignature.com", false},
		{"trailing slash trimmed", "https://example.com/", "https://example.com", false},
		{"repeated slashes trimmed", "https://example.com///", "https://example.com", false},
		{"surrounding space trimmed", "  http://example.com  ", "http://example.com", false},
		{"sub path kept", "https://example.com/assets", "https://example.com/assets", false},
		{"empty", "", "", true},
		{"only slashes", "///", "", true},
		{"no scheme", "nfts.cosmicsignature.com", "", true},
		{"unsupported scheme", "ftp://example.com", "", true},
		{"file scheme", "file:///tmp/x", "", true},
		{"no host", "https://", "", true},
		{"unparseable", "https://exa mple.com\x7f", "", true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			f, err := newFetcher(http.DefaultClient, tc.base)
			if (err != nil) != tc.wantErr {
				t.Fatalf("newFetcher(%q) error = %v, wantErr %v", tc.base, err, tc.wantErr)
			}
			if err == nil && f.base != tc.want {
				t.Errorf("base = %q, want %q", f.base, tc.want)
			}
		})
	}
}

func TestFetcherURLs(t *testing.T) {
	t.Parallel()
	f, err := newFetcher(http.DefaultClient, "https://nfts.cosmicsignature.com/")
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	if got, want := f.traitsURL("0x100033"), "https://nfts.cosmicsignature.com/traits/0x100033.json"; got != want {
		t.Errorf("traitsURL = %q, want %q", got, want)
	}
	if got, want := f.manifestURL("0x100033"), "https://nfts.cosmicsignature.com/asset-manifests/0x100033.json"; got != want {
		t.Errorf("manifestURL = %q, want %q", got, want)
	}
}

func TestFetcherClassifiesStatuses(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		status     int
		body       string
		want       fetchStatus
		wantErr    bool
		wantBody   string
		wantHasTag bool
	}{
		{name: "200 carries a body", status: http.StatusOK, body: `{"ok":true}`, want: fetchOK, wantBody: `{"ok":true}`, wantHasTag: true},
		{name: "304 is not modified", status: http.StatusNotModified, want: fetchNotModified, wantHasTag: true},
		{name: "404 is missing", status: http.StatusNotFound, want: fetchMissing},
		{name: "410 is missing", status: http.StatusGone, want: fetchMissing},
		{name: "500 is an error", status: http.StatusInternalServerError, wantErr: true},
		{name: "403 is an error", status: http.StatusForbidden, wantErr: true},
		{name: "302 is an unfollowed error", status: http.StatusFound, wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("ETag", `"v1"`)
				if tc.status == http.StatusFound {
					// Location-less 302: the client cannot follow it.
					w.WriteHeader(tc.status)
					return
				}
				w.WriteHeader(tc.status)
				if tc.body != "" {
					_, _ = w.Write([]byte(tc.body))
				}
			}))
			defer srv.Close()

			f, err := newFetcher(srv.Client(), srv.URL)
			if err != nil {
				t.Fatalf("newFetcher: %v", err)
			}
			got, err := f.get(t.Context(), f.traitsURL("0x1"), "")
			if (err != nil) != tc.wantErr {
				t.Fatalf("get() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if got.Status != tc.want {
				t.Errorf("status = %v, want %v", got.Status, tc.want)
			}
			if string(got.Body) != tc.wantBody {
				t.Errorf("body = %q, want %q", got.Body, tc.wantBody)
			}
			if tc.wantHasTag && got.ETag != `"v1"` {
				t.Errorf("etag = %q, want %q", got.ETag, `"v1"`)
			}
		})
	}
}

// TestFetcherSendsConditionalHeaders pins the bandwidth contract: a stored
// validator must travel with the request, and every request must announce
// that it wants JSON.
func TestFetcherSendsConditionalHeaders(t *testing.T) {
	t.Parallel()
	var gotIfNoneMatch, gotAccept string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotIfNoneMatch = r.Header.Get("If-None-Match")
		gotAccept = r.Header.Get("Accept")
		w.WriteHeader(http.StatusNotModified)
	}))
	defer srv.Close()

	f, err := newFetcher(srv.Client(), srv.URL)
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	if _, err := f.get(t.Context(), f.traitsURL("0x1"), `W/"abc"`); err != nil {
		t.Fatalf("get: %v", err)
	}
	if gotIfNoneMatch != `W/"abc"` {
		t.Errorf("If-None-Match = %q, want %q", gotIfNoneMatch, `W/"abc"`)
	}
	if gotAccept != "application/json" {
		t.Errorf("Accept = %q, want application/json", gotAccept)
	}
}

func TestFetcherOmitsEmptyValidator(t *testing.T) {
	t.Parallel()
	var present bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, present = r.Header["If-None-Match"]
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	f, err := newFetcher(srv.Client(), srv.URL)
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	if _, err := f.get(t.Context(), f.traitsURL("0x1"), ""); err != nil {
		t.Fatalf("get: %v", err)
	}
	if present {
		t.Error("an empty validator was still sent as If-None-Match")
	}
}

func TestFetcherRejectsOversizeBody(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(strings.Repeat("a", maxBodyBytes+1)))
	}))
	defer srv.Close()

	f, err := newFetcher(srv.Client(), srv.URL)
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	_, err = f.get(t.Context(), f.traitsURL("0x1"), "")
	if !errors.Is(err, errBodyTooLarge) {
		t.Fatalf("get() error = %v, want %v", err, errBodyTooLarge)
	}
}

func TestFetcherAcceptsBodyAtTheLimit(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(strings.Repeat("a", maxBodyBytes)))
	}))
	defer srv.Close()

	f, err := newFetcher(srv.Client(), srv.URL)
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	got, err := f.get(t.Context(), f.traitsURL("0x1"), "")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if len(got.Body) != maxBodyBytes {
		t.Errorf("len(body) = %d, want %d", len(got.Body), maxBodyBytes)
	}
}

// TestFetcherReportsTruncatedBody covers a host that announces more than it
// delivers: the read must fail rather than hand a half a document to the
// parser, which could conceivably still parse.
func TestFetcherReportsTruncatedBody(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Length", "4096")
		_, _ = w.Write([]byte(`{"schema_version":"1.0.0"`))
		// Dropping the connection mid-body surfaces as a read error.
		if hijacker, ok := w.(http.Hijacker); ok {
			conn, _, err := hijacker.Hijack()
			if err == nil {
				_ = conn.Close()
			}
		}
	}))
	defer srv.Close()

	f, err := newFetcher(srv.Client(), srv.URL)
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	if _, err := f.get(t.Context(), f.traitsURL("0x1"), ""); err == nil {
		t.Fatal("get() accepted a truncated body")
	}
}

func TestFetcherPropagatesCancellation(t *testing.T) {
	t.Parallel()
	release := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		<-release
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	defer close(release)

	f, err := newFetcher(srv.Client(), srv.URL)
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if _, err := f.get(ctx, f.traitsURL("0x1"), ""); !errors.Is(err, context.Canceled) {
		t.Fatalf("get() error = %v, want context.Canceled", err)
	}
}

func TestFetcherReportsUnreachableHost(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	client := srv.Client()
	base := srv.URL
	srv.Close()

	f, err := newFetcher(client, base)
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	if _, err := f.get(t.Context(), f.traitsURL("0x1"), ""); err == nil {
		t.Fatal("get() succeeded against a closed server")
	}
}

func TestFetcherRejectsUnbuildableRequest(t *testing.T) {
	t.Parallel()
	f, err := newFetcher(http.DefaultClient, "https://example.com")
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	if _, err := f.get(t.Context(), "https://exa\x7fmple.com/x", ""); err == nil {
		t.Fatal("get() accepted an unbuildable request URL")
	}
}

func TestFetcherHonoursClientTimeout(t *testing.T) {
	t.Parallel()
	release := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		<-release
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	defer close(release)

	client := srv.Client()
	client.Timeout = 20 * time.Millisecond
	f, err := newFetcher(client, srv.URL)
	if err != nil {
		t.Fatalf("newFetcher: %v", err)
	}
	if _, err := f.get(t.Context(), f.traitsURL("0x1"), ""); err == nil {
		t.Fatal("get() ignored the client timeout")
	}
}
