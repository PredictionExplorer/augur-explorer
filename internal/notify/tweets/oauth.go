// Copyright 2010 Gary Burd
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package tweets contains an OAuth 1.0a consumer (RFC 5849) trimmed to the
// surface this repository uses, together with helpers for sending tweets
// through the Twitter API.
//
// The OAuth client supports exactly the production flows:
//
//   - HMAC-SHA1 request signing via the Get and Post methods (twsend.go and
//     the rwctl tweet commands). Twitter only accepts HMAC-SHA1 for these
//     endpoints; the upstream library's RSA/PLAINTEXT/HMAC-SHA256 methods,
//     Put/Delete verbs, xAuth and session-renewal flows had no callers and
//     were removed.
//   - The out-of-band (PIN) authorization flow used by `rwctl twitter-auth`:
//     RequestTemporaryCredentials, AuthorizationURL and RequestToken.
//
// A nil *http.Client selects the package's bounded default client (a
// two-minute per-request timeout sized to the chunked video uploads — never
// the timeout-less http.DefaultClient, so no Twitter call can hang forever);
// tests inject an httptest.Server client. The context-aware variants accept
// a client via the HTTPClient context key.
package tweets

// PredictionExplorer note:
//		The original oauth code has to be modified to send correct 'oauth_consumer_key' field
//		which is Twitter's API key, because the default package doesn't send it, instead it
//		sends the token the User authorized

import (
	"bytes"
	"cmp"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1" //nolint:gosec // SHA-1 is mandated by the OAuth 1.0a HMAC-SHA1 signature method (RFC 5849)
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// noscape[b] is true if b should not be escaped per section 3.6 of the RFC.
var noEscape = [256]bool{
	'A': true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
	'a': true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
	'0': true, true, true, true, true, true, true, true, true, true,
	'-': true,
	'.': true,
	'_': true,
	'~': true,
}

// encode encodes string per section 3.6 of the RFC. If double is true, then
// the encoding is applied twice.
func encode(s string, double bool) []byte {
	// Compute size of result.
	m := 3
	if double {
		m = 5
	}
	n := 0
	for i := range len(s) {
		if noEscape[s[i]] {
			n++
		} else {
			n += m
		}
	}

	p := make([]byte, n)

	// Encode it.
	j := 0
	for i := range len(s) {
		b := s[i]
		switch {
		case noEscape[b]:
			p[j] = b
			j++
		case double:
			p[j] = '%'
			p[j+1] = '2'
			p[j+2] = '5'
			p[j+3] = "0123456789ABCDEF"[b>>4]
			p[j+4] = "0123456789ABCDEF"[b&15]
			j += 5
		default:
			p[j] = '%'
			p[j+1] = "0123456789ABCDEF"[b>>4]
			p[j+2] = "0123456789ABCDEF"[b&15]
			j += 3
		}
	}
	return p
}

type keyValue struct{ key, value []byte }

// compareKeyValue orders parameters by key, then by value, as RFC 5849
// section 3.4.1.3.2 requires for the signature base string.
func compareKeyValue(a, b keyValue) int {
	return cmp.Or(bytes.Compare(a.key, b.key), bytes.Compare(a.value, b.value))
}

type keyValues []keyValue

func (p keyValues) appendValues(values url.Values) keyValues {
	for k, vs := range values {
		k := encode(k, true)
		for _, v := range vs {
			v := encode(v, true)
			p = append(p, keyValue{k, v})
		}
	}
	return p
}

// writeBaseString writes method, url, and params to w using the OAuth signature
// base string computation described in section 3.4.1 of the RFC.
// Write errors are ignored: w is always an in-memory hash, whose Write never fails.
func writeBaseString(w io.Writer, method string, u *url.URL, form url.Values, oauthParams map[string]string) {
	// Method
	_, _ = w.Write(encode(strings.ToUpper(method), false))
	_, _ = w.Write([]byte{'&'})

	// URL
	scheme := strings.ToLower(u.Scheme)
	host := strings.ToLower(u.Host)

	uNoQuery := *u
	uNoQuery.RawQuery = ""
	path := uNoQuery.RequestURI()

	switch {
	case scheme == "http" && strings.HasSuffix(host, ":80"):
		host = host[:len(host)-len(":80")]
	case scheme == "https" && strings.HasSuffix(host, ":443"):
		host = host[:len(host)-len(":443")]
	}

	_, _ = w.Write(encode(scheme, false))
	_, _ = w.Write(encode("://", false))
	_, _ = w.Write(encode(host, false))
	_, _ = w.Write(encode(path, false))
	_, _ = w.Write([]byte{'&'})

	// Create sorted slice of encoded parameters. Parameter keys and values are
	// double encoded in a single step. This is safe because double encoding
	// does not change the sort order.
	queryParams := u.Query()
	p := make(keyValues, 0, len(form)+len(queryParams)+len(oauthParams))
	p = p.appendValues(form)
	p = p.appendValues(queryParams)
	for k, v := range oauthParams {
		p = append(p, keyValue{encode(k, true), encode(v, true)})
	}
	slices.SortFunc(p, compareKeyValue)

	// Write the parameters.
	encodedAmp := encode("&", false)
	encodedEqual := encode("=", false)
	sep := false
	for _, kv := range p {
		if sep {
			_, _ = w.Write(encodedAmp)
		} else {
			sep = true
		}
		_, _ = w.Write(kv.key)
		_, _ = w.Write(encodedEqual)
		_, _ = w.Write(kv.value)
	}
}

var nonceCounter uint64

func init() {
	if err := binary.Read(rand.Reader, binary.BigEndian, &nonceCounter); err != nil {
		// fallback to time if rand reader is broken
		nonceCounter = uint64(time.Now().UnixNano())
	}
}

// nonce returns a unique string.
func nonce() string {
	return strconv.FormatUint(atomic.AddUint64(&nonceCounter, 1), 16)
}

// signatureMethod is the OAuth signature method sent with every request.
// Twitter's v1.1 endpoints require HMAC-SHA1; the other RFC 5849 methods had
// no callers in this repository and were removed.
const signatureMethod = "HMAC-SHA1"

// Credentials represents client, temporary and token credentials.
type Credentials struct {
	Token  string // Also known as consumer key or access token.
	Secret string // Also known as consumer secret or access token secret.
}

// Client represents an OAuth client.
type Client struct {
	APIKey      string // twitter api key
	Nonce       string // nonce set in previous session
	Ts          int64  // timestamp
	ClientToken string

	// Credentials specifies the client key and secret.
	// Also known as the consumer key and secret
	Credentials Credentials

	// TemporaryCredentialRequestURI is the endpoint used by the client to
	// obtain a set of temporary credentials. Also known as the request token
	// URL.
	TemporaryCredentialRequestURI string

	// ResourceOwnerAuthorizationURI is the endpoint to which the resource
	// owner is redirected to grant authorization. Also known as authorization
	// URL.
	ResourceOwnerAuthorizationURI string

	// TokenRequestURI is the endpoint used by the client to request a set of
	// token credentials using a set of temporary credentials. Also known as
	// access token URL.
	TokenRequestURI string

	// TemporaryCredentialsMethod is the HTTP method used by the client to
	// obtain a set of temporary credentials. If this field is the empty
	// string, then POST is used.
	TemporaryCredentialsMethod string

	// TokenCredentailsMethod is the HTTP method used by the client to request
	// a set of token credentials. If this field is the empty string, then POST
	// is used.
	TokenCredentailsMethod string

	// Header specifies optional extra headers for requests.
	Header http.Header
}

type request struct {
	credentials *Credentials
	method      string
	u           *url.URL
	form        url.Values
	verifier    string
	callbackURL string
}

// oauthParams returns the OAuth request parameters for the given credentials,
// method, URL and application params. See
// http://tools.ietf.org/html/rfc5849#section-3.4 for more information about
// signatures.
func (c *Client) oauthParams(r *request) map[string]string {
	oauthParams := map[string]string{
		// PredictionExplorer: send the Twitter API key here, not the user-authorized token.
		"oauth_consumer_key":     c.APIKey,
		"oauth_signature_method": signatureMethod,
		"oauth_version":          "1.0",
	}

	if c.Ts == 0 {
		oauthParams["oauth_timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	} else {
		oauthParams["oauth_timestamp"] = strconv.FormatInt(c.Ts, 10)
	}
	oauthParams["oauth_nonce"] = nonce()

	if r.credentials != nil {
		if len(c.ClientToken) > 0 {
			oauthParams["oauth_token"] = c.ClientToken
		} else {
			oauthParams["oauth_token"] = r.credentials.Token
		}
	}

	if r.verifier != "" {
		oauthParams["oauth_verifier"] = r.verifier
	}

	if r.callbackURL != "" {
		oauthParams["oauth_callback"] = r.callbackURL
	}
	if len(c.Nonce) > 0 {
		oauthParams["oauth_nonce"] = c.Nonce
	}

	oauthParams["oauth_signature"] = c.hmacSHA1Signature(r, oauthParams)
	return oauthParams
}

// hmacSHA1Signature computes the HMAC-SHA1 request signature per section
// 3.4.2 of the RFC: the key is consumerSecret&tokenSecret and the message is
// the signature base string.
func (c *Client) hmacSHA1Signature(r *request, oauthParams map[string]string) string {
	key := encode(c.Credentials.Secret, false)
	key = append(key, '&')
	if r.credentials != nil {
		key = append(key, encode(r.credentials.Secret, false)...)
	}
	hm := hmac.New(sha1.New, key)
	writeBaseString(hm, r.method, r.u, r.form, oauthParams)

	return base64.StdEncoding.EncodeToString(hm.Sum(key[:0]))
}

var oauthKeys = []string{
	"oauth_consumer_key",
	"oauth_nonce",
	"oauth_signature",
	"oauth_signature_method",
	"oauth_timestamp",
	"oauth_token",
	"oauth_version",
	"oauth_callback",
	"oauth_verifier",
}

func (c *Client) authorizationHeader(r *request) string {
	p := c.oauthParams(r)
	var h []byte
	// Append parameters in a fixed order to support testing.
	for _, k := range oauthKeys {
		if v, ok := p[k]; ok {
			if h == nil {
				h = []byte(`OAuth `)
			} else {
				h = append(h, ", "...)
			}
			h = append(h, k...)
			h = append(h, `="`...)
			h = append(h, encode(v, false)...)
			h = append(h, '"')
		}
	}
	return string(h)
}

func (c *Client) do(ctx context.Context, urlStr string, r *request) (*http.Response, error) {
	var body io.Reader
	if r.method != http.MethodGet {
		body = strings.NewReader(r.form.Encode())
	}
	req, err := http.NewRequestWithContext(ctx, r.method, urlStr, body)
	if err != nil {
		return nil, err
	}
	if req.URL.RawQuery != "" {
		return nil, errors.New("oauth: url must not contain a query string")
	}
	maps.Copy(req.Header, c.Header)
	r.u = req.URL
	req.Header.Set("Authorization", c.authorizationHeader(r))
	if r.method == http.MethodGet {
		req.URL.RawQuery = r.form.Encode()
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	req = req.WithContext(ctx)
	client := contextClient(ctx)
	return client.Do(req)
}

// Get issues a GET to the specified URL with form added as a query string.
func (c *Client) Get(client *http.Client, credentials *Credentials, urlStr string, form url.Values) (*http.Response, error) {
	ctx := context.WithValue(context.Background(), HTTPClient, client)
	return c.GetContext(ctx, credentials, urlStr, form)
}

// GetContext uses Context to perform Get.
func (c *Client) GetContext(ctx context.Context, credentials *Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return c.do(ctx, urlStr, &request{method: http.MethodGet, credentials: credentials, form: form})
}

// Post issues a POST with the specified form.
func (c *Client) Post(client *http.Client, credentials *Credentials, urlStr string, form url.Values) (*http.Response, error) {
	ctx := context.WithValue(context.Background(), HTTPClient, client)
	return c.PostContext(ctx, credentials, urlStr, form)
}

// PostContext uses Context to perform Post.
func (c *Client) PostContext(ctx context.Context, credentials *Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return c.do(ctx, urlStr, &request{method: http.MethodPost, credentials: credentials, form: form})
}

func (c *Client) requestCredentials(ctx context.Context, u string, r *request) (*Credentials, url.Values, error) {
	if r.method == "" {
		r.method = http.MethodPost
	}
	resp, err := c.do(ctx, u, r)
	if err != nil {
		return nil, nil, err
	}
	p, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close() // best-effort: body already fully read
	if err != nil {
		return nil, nil, RequestCredentialsError{
			StatusCode: resp.StatusCode, Header: resp.Header,
			Body: p, msg: err.Error(),
		}
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, nil, RequestCredentialsError{
			StatusCode: resp.StatusCode, Header: resp.Header,
			Body: p, msg: fmt.Sprintf("OAuth server status %d, %s", resp.StatusCode, string(p)),
		}
	}
	m, err := url.ParseQuery(string(p))
	if err != nil {
		return nil, nil, RequestCredentialsError{
			StatusCode: resp.StatusCode, Header: resp.Header,
			Body: p, msg: err.Error(),
		}
	}
	tokens := m["oauth_token"]
	if len(tokens) == 0 || tokens[0] == "" {
		return nil, nil, RequestCredentialsError{
			StatusCode: resp.StatusCode, Header: resp.Header,
			Body: p, msg: "oauth: token missing from server result",
		}
	}
	secrets := m["oauth_token_secret"]
	if len(secrets) == 0 { // allow "" as a valid secret.
		return nil, nil, RequestCredentialsError{
			StatusCode: resp.StatusCode, Header: resp.Header,
			Body: p, msg: "oauth: secret missing from server result",
		}
	}
	return &Credentials{Token: tokens[0], Secret: secrets[0]}, m, nil
}

// RequestTemporaryCredentials requests temporary credentials from the server.
// See http://tools.ietf.org/html/rfc5849#section-2.1 for information about
// temporary credentials.
func (c *Client) RequestTemporaryCredentials(client *http.Client, callbackURL string, additionalParams url.Values) (*Credentials, error) {
	ctx := context.WithValue(context.Background(), HTTPClient, client)
	return c.RequestTemporaryCredentialsContext(ctx, callbackURL, additionalParams)
}

// RequestTemporaryCredentialsContext uses Context to perform RequestTemporaryCredentials.
func (c *Client) RequestTemporaryCredentialsContext(ctx context.Context, callbackURL string, additionalParams url.Values) (*Credentials, error) {
	credentials, _, err := c.requestCredentials(ctx, c.TemporaryCredentialRequestURI,
		&request{method: c.TemporaryCredentialsMethod, form: additionalParams, callbackURL: callbackURL})
	return credentials, err
}

// RequestToken requests token credentials from the server. See
// http://tools.ietf.org/html/rfc5849#section-2.3 for information about token
// credentials.
func (c *Client) RequestToken(client *http.Client, temporaryCredentials *Credentials, verifier string) (*Credentials, url.Values, error) {
	ctx := context.WithValue(context.Background(), HTTPClient, client)
	return c.RequestTokenContext(ctx, temporaryCredentials, verifier)
}

// RequestTokenContext uses Context to perform RequestToken.
func (c *Client) RequestTokenContext(ctx context.Context, temporaryCredentials *Credentials, verifier string) (*Credentials, url.Values, error) {
	return c.requestCredentials(ctx, c.TokenRequestURI,
		&request{credentials: temporaryCredentials, method: c.TokenCredentailsMethod, verifier: verifier})
}

// AuthorizationURL returns the URL for resource owner authorization. See
// http://tools.ietf.org/html/rfc5849#section-2.2 for information about
// resource owner authorization.
func (c *Client) AuthorizationURL(temporaryCredentials *Credentials, additionalParams url.Values) string {
	params := make(url.Values)
	maps.Copy(params, additionalParams)
	params.Set("oauth_token", temporaryCredentials.Token)
	return c.ResourceOwnerAuthorizationURI + "?" + params.Encode()
}

// HTTPClient is the context key to use with context's
// WithValue function to associate an *http.Client value with a context.
var HTTPClient contextKey

type contextKey struct{}

// defaultClientTimeout bounds one Twitter API exchange made through the
// package's fallback client (D22: no outbound call may wait forever). Two
// minutes is sized to the slowest legitimate exchange — a chunked video
// APPEND carrying a base64-encoded MP4 — while a plain status update
// finishes in a second.
const defaultClientTimeout = 2 * time.Minute

// defaultHTTPClient replaces the timeout-less http.DefaultClient as the
// fallback: every request without an injected client is still bounded.
var defaultHTTPClient = &http.Client{Timeout: defaultClientTimeout}

func contextClient(ctx context.Context) *http.Client {
	if ctx != nil {
		if hc, ok := ctx.Value(HTTPClient).(*http.Client); ok && hc != nil {
			return hc
		}
	}
	return defaultHTTPClient
}

// RequestCredentialsError is an error containing
// response information when requesting credentials.
type RequestCredentialsError struct {
	StatusCode int
	Header     http.Header
	Body       []byte
	msg        string
}

func (e RequestCredentialsError) Error() string {
	return e.msg
}
