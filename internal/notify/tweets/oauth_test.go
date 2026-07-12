package tweets

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// TestHmacSHA1SignatureKnownAnswer pins the HMAC-SHA1 signing algorithm
// against the worked example in Twitter's "Creating a signature" developer
// documentation (the de-facto conformance vector for OAuth 1.0a requests to
// the statuses/update endpoint).
func TestHmacSHA1SignatureKnownAnswer(t *testing.T) {
	c := Client{ //nolint:gosec // public example vector from Twitter's "creating a signature" docs, not a credential
		APIKey:      "xvz1evFS4wEEPTGEFPHBog",
		ClientToken: "370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb",
		Nonce:       "kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg",
		Ts:          1318622958,
		Credentials: Credentials{ //nolint:gosec // public example vector, not a credential
			Token:  "xvz1evFS4wEEPTGEFPHBog",
			Secret: "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw",
		},
	}
	tokenCreds := &Credentials{ //nolint:gosec // public example vector, not a credential
		Token:  "370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb",
		Secret: "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE",
	}
	// The documented example signs the v1 endpoint URL.
	u, err := url.Parse("https://api.twitter.com/1/statuses/update.json")
	if err != nil {
		t.Fatal(err)
	}
	form := url.Values{
		"status":           {"Hello Ladies + Gentlemen, a signed OAuth request!"},
		"include_entities": {"true"},
	}
	params := c.oauthParams(&request{credentials: tokenCreds, method: "POST", u: u, form: form})

	const wantSig = "tnnArxj06cWHq44gCs1OSKk/jLY="
	if got := params["oauth_signature"]; got != wantSig {
		t.Errorf("oauth_signature = %q, want %q", got, wantSig)
	}
	if got := params["oauth_signature_method"]; got != "HMAC-SHA1" {
		t.Errorf("oauth_signature_method = %q, want HMAC-SHA1", got)
	}
	if got := params["oauth_consumer_key"]; got != c.APIKey {
		t.Errorf("oauth_consumer_key = %q, want %q", got, c.APIKey)
	}
	if got := params["oauth_token"]; got != c.ClientToken {
		t.Errorf("oauth_token = %q, want %q", got, c.ClientToken)
	}
}

// TestOauthParamsWithoutClientToken pins the fallback used by the PIN flow:
// with no ClientToken the request credentials' token is sent, and verifier /
// callback parameters appear only when set.
func TestOauthParamsWithoutClientToken(t *testing.T) {
	c := Client{APIKey: "app-key", Credentials: Credentials{Token: "app-key", Secret: "app-secret"}}
	u, _ := url.Parse("https://example.com/oauth/access_token")

	params := c.oauthParams(&request{
		credentials: &Credentials{Token: "temp-token", Secret: "temp-secret"},
		method:      "POST",
		u:           u,
		verifier:    "pin123",
		callbackURL: "oob",
	})
	if got := params["oauth_token"]; got != "temp-token" {
		t.Errorf("oauth_token = %q, want temp-token", got)
	}
	if got := params["oauth_verifier"]; got != "pin123" {
		t.Errorf("oauth_verifier = %q, want pin123", got)
	}
	if got := params["oauth_callback"]; got != "oob" {
		t.Errorf("oauth_callback = %q, want oob", got)
	}
	if params["oauth_nonce"] == "" || params["oauth_timestamp"] == "" {
		t.Errorf("generated nonce/timestamp missing: %v", params)
	}

	noExtras := c.oauthParams(&request{method: "POST", u: u})
	for _, k := range []string{"oauth_token", "oauth_verifier", "oauth_callback"} {
		if _, ok := noExtras[k]; ok {
			t.Errorf("%s should be absent without credentials/verifier/callback", k)
		}
	}
}

// TestClientPostAndGetAgainstServer sends signed requests to an httptest
// server and checks the wire shape end to end: Authorization header contents,
// form transport (body for POST, query string for GET) and the injected
// client plumbing.
func TestClientPostAndGetAgainstServer(t *testing.T) {
	type seen struct {
		method string
		auth   string
		body   string
		query  string
		ctype  string
	}
	var got seen
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		got = seen{
			method: r.Method,
			auth:   r.Header.Get("Authorization"),
			body:   string(b),
			query:  r.URL.RawQuery,
			ctype:  r.Header.Get("Content-Type"),
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c := Client{
		APIKey:      "api-key",
		ClientToken: "access-token",
		Credentials: Credentials{Token: "api-key", Secret: "api-secret"},
	}
	creds := &Credentials{Token: "access-token", Secret: "token-secret"}

	resp, err := c.Post(srv.Client(), creds, srv.URL+"/post", url.Values{"status": {"hello world"}})
	if err != nil {
		t.Fatalf("Post: %v", err)
	}
	_ = resp.Body.Close()
	if got.method != http.MethodPost {
		t.Errorf("method = %q, want POST", got.method)
	}
	if !strings.HasPrefix(got.auth, "OAuth ") {
		t.Errorf("Authorization header %q does not start with OAuth", got.auth)
	}
	for _, part := range []string{`oauth_consumer_key="api-key"`, `oauth_token="access-token"`, `oauth_signature_method="HMAC-SHA1"`, `oauth_signature="`} {
		if !strings.Contains(got.auth, part) {
			t.Errorf("Authorization header %q missing %q", got.auth, part)
		}
	}
	if got.body != "status=hello+world" {
		t.Errorf("POST body = %q, want status=hello+world", got.body)
	}
	if got.ctype != "application/x-www-form-urlencoded" {
		t.Errorf("Content-Type = %q", got.ctype)
	}

	resp, err = c.Get(srv.Client(), creds, srv.URL+"/get", url.Values{"command": {"STATUS"}})
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	_ = resp.Body.Close()
	if got.method != http.MethodGet {
		t.Errorf("method = %q, want GET", got.method)
	}
	if got.query != "command=STATUS" {
		t.Errorf("GET query = %q, want command=STATUS", got.query)
	}
	if got.body != "" {
		t.Errorf("GET body = %q, want empty", got.body)
	}
}

// TestClientExtraHeadersForwarded pins the Header field used by the video
// APPEND stage: custom headers are forwarded, but for POSTs the
// form-encoding Content-Type deliberately wins because the body really is
// form-encoded (media_data field) — the behavior Twitter's chunked upload
// has accepted from this code in production.
func TestClientExtraHeadersForwarded(t *testing.T) {
	var gotCT, gotCTE string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotCT = r.Header.Get("Content-Type")
		gotCTE = r.Header.Get("Content-Transfer-Encoding")
	}))
	defer srv.Close()

	c := Client{Credentials: Credentials{Token: "k", Secret: "s"}}
	c.Header = make(http.Header)
	c.Header.Set("Content-Type", "application/octet-stream")
	c.Header.Set("Content-Transfer-Encoding", "base64")

	resp, err := c.Post(srv.Client(), nil, srv.URL, url.Values{"command": {"APPEND"}})
	if err != nil {
		t.Fatalf("Post: %v", err)
	}
	_ = resp.Body.Close()
	// do sets the form Content-Type after copying c.Header, so the
	// form-encoding value wins for POST bodies (which are form-encoded).
	if gotCT != "application/x-www-form-urlencoded" {
		t.Errorf("Content-Type = %q, want application/x-www-form-urlencoded", gotCT)
	}
	if gotCTE != "base64" {
		t.Errorf("Content-Transfer-Encoding = %q, want base64", gotCTE)
	}
}

func TestDoRejectsQueryString(t *testing.T) {
	c := Client{}
	// Both rejections happen before any request is issued, so no response
	// body ever exists to close.
	resp, err := c.Post(nil, nil, "https://example.com/path?x=1", nil)
	if err == nil {
		_ = resp.Body.Close()
		t.Fatal("Post with query string should fail")
	}
	resp, err = c.Post(nil, nil, "://bad-url", nil)
	if err == nil {
		_ = resp.Body.Close()
		t.Fatal("Post with unparseable URL should fail")
	}
}

// TestPinAuthorizationFlow drives the complete out-of-band flow that
// `rwctl twitter-auth` runs: temporary credentials, the user-facing
// authorization URL and the verifier exchange for token credentials.
func TestPinAuthorizationFlow(t *testing.T) {
	var tempAuth, tokenAuth string
	mux := http.NewServeMux()
	mux.HandleFunc("/oauth/request_token", func(w http.ResponseWriter, r *http.Request) {
		tempAuth = r.Header.Get("Authorization")
		_, _ = w.Write([]byte("oauth_token=temp-tok&oauth_token_secret=temp-sec&oauth_callback_confirmed=true"))
	})
	mux.HandleFunc("/oauth/access_token", func(w http.ResponseWriter, r *http.Request) {
		tokenAuth = r.Header.Get("Authorization")
		_, _ = w.Write([]byte("oauth_token=final-tok&oauth_token_secret=final-sec&screen_name=tester"))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	c := Client{
		APIKey:                        "app-key",
		Credentials:                   Credentials{Token: "app-key", Secret: "app-secret"},
		TemporaryCredentialRequestURI: srv.URL + "/oauth/request_token",
		ResourceOwnerAuthorizationURI: srv.URL + "/oauth/authorize",
		TokenRequestURI:               srv.URL + "/oauth/access_token",
	}

	tempCred, err := c.RequestTemporaryCredentials(srv.Client(), "oob", nil)
	if err != nil {
		t.Fatalf("RequestTemporaryCredentials: %v", err)
	}
	if tempCred.Token != "temp-tok" || tempCred.Secret != "temp-sec" {
		t.Errorf("temporary credentials = %+v", tempCred)
	}
	if !strings.Contains(tempAuth, `oauth_callback="oob"`) {
		t.Errorf("request_token Authorization %q missing oob callback", tempAuth)
	}

	authURL := c.AuthorizationURL(tempCred, url.Values{"extra": {"1"}})
	parsed, err := url.Parse(authURL)
	if err != nil {
		t.Fatalf("AuthorizationURL produced unparseable URL %q: %v", authURL, err)
	}
	if got := parsed.Query().Get("oauth_token"); got != "temp-tok" {
		t.Errorf("authorization URL oauth_token = %q", got)
	}
	if got := parsed.Query().Get("extra"); got != "1" {
		t.Errorf("authorization URL extra param = %q", got)
	}

	tokenCred, extras, err := c.RequestToken(srv.Client(), tempCred, "123456")
	if err != nil {
		t.Fatalf("RequestToken: %v", err)
	}
	if tokenCred.Token != "final-tok" || tokenCred.Secret != "final-sec" {
		t.Errorf("token credentials = %+v", tokenCred)
	}
	if got := extras.Get("screen_name"); got != "tester" {
		t.Errorf("extras screen_name = %q", got)
	}
	if !strings.Contains(tokenAuth, `oauth_verifier="123456"`) {
		t.Errorf("access_token Authorization %q missing verifier", tokenAuth)
	}
	if !strings.Contains(tokenAuth, `oauth_token="temp-tok"`) {
		t.Errorf("access_token Authorization %q missing temporary token", tokenAuth)
	}
}

// TestRequestCredentialsErrors pins every failure branch of the credentials
// exchange: transport errors, non-200 statuses (with status/header/body
// preserved on the typed error), unparseable bodies and missing fields.
func TestRequestCredentialsErrors(t *testing.T) {
	newClient := func(srvURL string) Client {
		return Client{
			APIKey:                        "app-key",
			Credentials:                   Credentials{Token: "app-key", Secret: "app-secret"},
			TemporaryCredentialRequestURI: srvURL + "/oauth/request_token",
			TokenRequestURI:               srvURL + "/oauth/access_token",
		}
	}

	t.Run("transport error", func(t *testing.T) {
		c := newClient("http://127.0.0.1:1")
		if _, err := c.RequestTemporaryCredentials(nil, "oob", nil); err == nil {
			t.Fatal("expected transport error")
		}
	})

	cases := []struct {
		name    string
		status  int
		body    string
		wantMsg string
	}{
		{"server error status", http.StatusUnauthorized, "denied", "OAuth server status 401, denied"},
		{"unparseable body", http.StatusOK, "%zz", ""},
		{"missing token", http.StatusOK, "oauth_token_secret=s", "oauth: token missing from server result"},
		{"empty token", http.StatusOK, "oauth_token=&oauth_token_secret=s", "oauth: token missing from server result"},
		{"missing secret", http.StatusOK, "oauth_token=tok", "oauth: secret missing from server result"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer srv.Close()
			c := newClient(srv.URL)
			_, err := c.RequestTemporaryCredentials(srv.Client(), "oob", nil)
			if err == nil {
				t.Fatal("expected error")
			}
			var rce RequestCredentialsError
			if !errors.As(err, &rce) {
				t.Fatalf("error %T is not RequestCredentialsError", err)
			}
			if rce.StatusCode != tc.status {
				t.Errorf("StatusCode = %d, want %d", rce.StatusCode, tc.status)
			}
			if string(rce.Body) != tc.body {
				t.Errorf("Body = %q, want %q", rce.Body, tc.body)
			}
			if tc.wantMsg != "" && rce.Error() != tc.wantMsg {
				t.Errorf("Error() = %q, want %q", rce.Error(), tc.wantMsg)
			}
		})
	}
}

// TestRequestTokenEmptySecretAllowed pins the documented "allow empty secret"
// carve-out: oauth_token_secret= (present but empty) is a valid response.
func TestRequestTokenEmptySecretAllowed(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("oauth_token=tok&oauth_token_secret="))
	}))
	defer srv.Close()
	c := Client{TokenRequestURI: srv.URL + "/oauth/access_token"}
	cred, _, err := c.RequestToken(srv.Client(), &Credentials{Token: "t", Secret: "s"}, "v")
	if err != nil {
		t.Fatalf("RequestToken: %v", err)
	}
	if cred.Token != "tok" || cred.Secret != "" {
		t.Errorf("credentials = %+v, want token=tok secret empty", cred)
	}
}

func TestContextClientFallback(t *testing.T) {
	if got := contextClient(nil); got != http.DefaultClient { //nolint:staticcheck // nil context is the documented fallback input
		t.Error("nil context should fall back to http.DefaultClient")
	}
	ctx := t.Context()
	if got := contextClient(ctx); got != http.DefaultClient {
		t.Error("context without HTTPClient should fall back to http.DefaultClient")
	}
}

func TestNonceUnique(t *testing.T) {
	a, b := nonce(), nonce()
	if a == b {
		t.Errorf("nonce() returned duplicate %q", a)
	}
}
