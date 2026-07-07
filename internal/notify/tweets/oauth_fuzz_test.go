// Fuzz targets for the OAuth 1.0 request signing helpers (MODERNIZATION.md
// §4.4 FuzzTwitterRequestBuild): percent-encoding must be reversible and the
// signature base string must keep its three-component shape for any input.
package tweets

import (
	"bytes"
	"net/url"
	"strings"
	"testing"
)

func TestEncodeKnown(t *testing.T) {
	cases := []struct {
		in     string
		double bool
		want   string
	}{
		{"abcXYZ019-._~", false, "abcXYZ019-._~"},
		{"a b", false, "a%20b"},
		{"a b", true, "a%2520b"},
		{"100%", false, "100%25"},
		{"", false, ""},
		{"√", false, "%E2%88%9A"},
	}
	for _, tc := range cases {
		if got := string(encode(tc.in, tc.double)); got != tc.want {
			t.Errorf("encode(%q, %v) = %q, want %q", tc.in, tc.double, got, tc.want)
		}
	}
}

func FuzzTwitterRequestBuild(f *testing.F) {
	f.Add("POST", "https://api.twitter.com/2/tweets", "status", "hello world & more", "oauth_nonce", "abc123")
	f.Add("get", "http://example.com:80/path?q=1", "k", "√unicode", "oauth_token", "")
	f.Add("", "https://example.com", "", "", "", "")
	f.Fuzz(func(t *testing.T, method, rawURL, formKey, formVal, oauthKey, oauthVal string) {
		// encode: output must percent-decode back to the input.
		for _, double := range []bool{false, true} {
			enc := string(encode(formVal, double))
			dec, err := url.PathUnescape(enc)
			if err != nil {
				t.Fatalf("encode(%q, %v) = %q: not valid percent-encoding: %v", formVal, double, enc, err)
			}
			if double {
				if dec, err = url.PathUnescape(dec); err != nil {
					t.Fatalf("encode(%q, true) = %q: second decode failed: %v", formVal, enc, err)
				}
			}
			if dec != formVal {
				t.Fatalf("encode(%q, %v) round-tripped to %q", formVal, double, dec)
			}
		}

		// writeBaseString: base string is method&url&params with exactly the
		// two top-level ampersands; everything else must be escaped.
		u, err := url.Parse(rawURL)
		if err != nil {
			t.Skip("unparseable URL")
		}
		form := url.Values{}
		if formKey != "" {
			form.Set(formKey, formVal)
		}
		oauthParams := map[string]string{}
		if oauthKey != "" {
			oauthParams[oauthKey] = oauthVal
		}
		var buf bytes.Buffer
		writeBaseString(&buf, method, u, form, oauthParams)
		base := buf.String()
		if got := strings.Count(base, "&"); got != 2 {
			t.Fatalf("base string has %d raw '&' separators, want exactly 2: %q", got, base)
		}
	})
}
