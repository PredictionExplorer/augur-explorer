package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/notify/tweets"
)

func TestReadOAuthAppCredentials(t *testing.T) {
	dir := t.TempDir()

	if _, err := readOAuthAppCredentials(filepath.Join(dir, "missing.json")); err == nil ||
		!strings.Contains(err.Error(), "can't read") {
		t.Errorf("missing config error = %v", err)
	}

	bad := filepath.Join(dir, "bad.json")
	if err := os.WriteFile(bad, []byte("{"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := readOAuthAppCredentials(bad); err == nil || !strings.Contains(err.Error(), "can't parse") {
		t.Errorf("bad config error = %v", err)
	}

	good := filepath.Join(dir, "good.json")
	if err := os.WriteFile(good, []byte(`{"Credentials":{"Token":"app-key","Secret":"app-secret"}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	creds, err := readOAuthAppCredentials(good)
	if err != nil || creds.Token != "app-key" || creds.Secret != "app-secret" {
		t.Errorf("readOAuthAppCredentials = (%+v, %v)", creds, err)
	}
}

func TestRunTwitterAuthFlowPinExchange(t *testing.T) {
	var tokenAuth, tweetBody string
	mux := http.NewServeMux()
	mux.HandleFunc("/oauth/request_token", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("oauth_token=temp-tok&oauth_token_secret=temp-sec&oauth_callback_confirmed=true"))
	})
	mux.HandleFunc("/oauth/access_token", func(w http.ResponseWriter, r *http.Request) {
		tokenAuth = r.Header.Get("Authorization")
		_, _ = w.Write([]byte("oauth_token=final-tok&oauth_token_secret=final-sec&screen_name=tester"))
	})
	mux.HandleFunc("/1.1/statuses/update.json", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Errorf("parsing tweet form: %v", err)
		}
		tweetBody = r.PostForm.Get("status")
		_, _ = w.Write([]byte(`{"id_str":"777"}`))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	endpoints := twitterAuthEndpoints{
		requestTokenURI: srv.URL + "/oauth/request_token",
		authorizeURI:    srv.URL + "/oauth/authorize",
		accessTokenURI:  srv.URL + "/oauth/access_token",
		statusUpdateURI: srv.URL + "/1.1/statuses/update.json",
	}
	appCreds := tweets.Credentials{Token: "app-key", Secret: "app-secret"}

	var out bytes.Buffer
	in := strings.NewReader("123456\n")
	if err := runTwitterAuthFlow(appCreds, endpoints, srv.Client(), in, &out); err != nil {
		t.Fatalf("runTwitterAuthFlow: %v", err)
	}

	text := out.String()
	for _, want := range []string{
		"Token: temp-tok",
		"Secret: temp-sec",
		"1. Go to " + srv.URL + "/oauth/authorize?oauth_token=temp-tok",
		"Token: final-tok",
		"Token Secret: final-sec",
		`{"id_str":"777"}`,
	} {
		if !strings.Contains(text, want) {
			t.Errorf("flow output missing %q\noutput:\n%s", want, text)
		}
	}
	if !strings.Contains(tokenAuth, `oauth_verifier="123456"`) {
		t.Errorf("access_token Authorization %q missing the PIN verifier", tokenAuth)
	}
	if tweetBody != "got authorization from account owner" {
		t.Errorf("test tweet status = %q", tweetBody)
	}
}

func TestRunTwitterAuthFlowErrors(t *testing.T) {
	// okRequestToken serves a valid temporary-credentials answer.
	okRequestToken := func(mux *http.ServeMux) {
		mux.HandleFunc("/oauth/request_token", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte("oauth_token=t&oauth_token_secret=s&oauth_callback_confirmed=true"))
		})
	}

	t.Run("request token failure", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "nope", http.StatusUnauthorized)
		}))
		defer srv.Close()
		endpoints := twitterAuthEndpoints{requestTokenURI: srv.URL + "/oauth/request_token"}
		err := runTwitterAuthFlow(tweets.Credentials{}, endpoints, srv.Client(), strings.NewReader(""), &bytes.Buffer{})
		if err == nil || !strings.Contains(err.Error(), "RequestTemporaryCredentials") {
			t.Errorf("err = %v", err)
		}
	})
	t.Run("empty verification code", func(t *testing.T) {
		mux := http.NewServeMux()
		okRequestToken(mux)
		srv := httptest.NewServer(mux)
		defer srv.Close()
		endpoints := twitterAuthEndpoints{
			requestTokenURI: srv.URL + "/oauth/request_token",
			authorizeURI:    srv.URL + "/oauth/authorize",
		}
		err := runTwitterAuthFlow(tweets.Credentials{}, endpoints, srv.Client(), strings.NewReader("\n"), &bytes.Buffer{})
		if err == nil || !strings.Contains(err.Error(), "verification code") {
			t.Errorf("err = %v", err)
		}
	})
	t.Run("token exchange failure", func(t *testing.T) {
		mux := http.NewServeMux()
		okRequestToken(mux)
		mux.HandleFunc("/oauth/access_token", func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "denied", http.StatusUnauthorized)
		})
		srv := httptest.NewServer(mux)
		defer srv.Close()
		endpoints := twitterAuthEndpoints{
			requestTokenURI: srv.URL + "/oauth/request_token",
			authorizeURI:    srv.URL + "/oauth/authorize",
			accessTokenURI:  srv.URL + "/oauth/access_token",
		}
		err := runTwitterAuthFlow(tweets.Credentials{}, endpoints, srv.Client(), strings.NewReader("123\n"), &bytes.Buffer{})
		if err == nil || !strings.Contains(err.Error(), "RequestToken") {
			t.Errorf("err = %v", err)
		}
	})
	t.Run("test tweet failure", func(t *testing.T) {
		mux := http.NewServeMux()
		okRequestToken(mux)
		mux.HandleFunc("/oauth/access_token", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte("oauth_token=ft&oauth_token_secret=fs"))
		})
		srv := httptest.NewServer(mux)
		defer srv.Close()
		endpoints := twitterAuthEndpoints{
			requestTokenURI: srv.URL + "/oauth/request_token",
			authorizeURI:    srv.URL + "/oauth/authorize",
			accessTokenURI:  srv.URL + "/oauth/access_token",
			statusUpdateURI: "http://127.0.0.1:1/unreachable",
		}
		err := runTwitterAuthFlow(tweets.Credentials{}, endpoints, nil, strings.NewReader("123\n"), &bytes.Buffer{})
		if err == nil || !strings.Contains(err.Error(), "test tweet failed") {
			t.Errorf("err = %v", err)
		}
	})
	t.Run("response body read failure", func(t *testing.T) {
		mux := http.NewServeMux()
		okRequestToken(mux)
		mux.HandleFunc("/oauth/access_token", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte("oauth_token=ft&oauth_token_secret=fs"))
		})
		mux.HandleFunc("/1.1/statuses/update.json", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Length", "1000")
			_, _ = w.Write([]byte("partial"))
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			panic(http.ErrAbortHandler)
		})
		srv := httptest.NewServer(mux)
		defer srv.Close()
		endpoints := twitterAuthEndpoints{
			requestTokenURI: srv.URL + "/oauth/request_token",
			authorizeURI:    srv.URL + "/oauth/authorize",
			accessTokenURI:  srv.URL + "/oauth/access_token",
			statusUpdateURI: srv.URL + "/1.1/statuses/update.json",
		}
		err := runTwitterAuthFlow(tweets.Credentials{}, endpoints, srv.Client(), strings.NewReader("123\n"), &bytes.Buffer{})
		if err == nil || !strings.Contains(err.Error(), "error reading response") {
			t.Errorf("err = %v", err)
		}
	})
}

func TestProductionTwitterEndpoints(t *testing.T) {
	e := productionTwitterEndpoints()
	for name, uri := range map[string]string{
		"request": e.requestTokenURI,
		"auth":    e.authorizeURI,
		"access":  e.accessTokenURI,
		"status":  e.statusUpdateURI,
	} {
		if !strings.HasPrefix(uri, "https://api.twitter.com/") {
			t.Errorf("%s endpoint = %q, want twitter.com", name, uri)
		}
	}
}
