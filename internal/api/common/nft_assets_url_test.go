package common

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func TestNormalizeNFTAssetsPublicBase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"   ", ""},
		{"/", ""},
		{"https://x.example", "https://x.example/images"},
		{"https://x.example/", "https://x.example/images"},
		{"https://x.example//", "https://x.example/images"},
		{"https://x.example/images", "https://x.example/images"},
		{"https://x.example/images/", "https://x.example/images"},
		{"https://x.example/randomwalk", "https://x.example/images"},
		{"https://x.example/randomwalk/", "https://x.example/images"},
		{" https://x.example ", "https://x.example/images"},
		{"http://cdn.example/nested/images", "http://cdn.example/nested/images"},
	}
	for _, tc := range cases {
		if got := NormalizeNFTAssetsPublicBase(tc.in); got != tc.want {
			t.Errorf("NormalizeNFTAssetsPublicBase(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func FuzzNFTAssetsPublicBase(f *testing.F) {
	for _, seed := range []string{
		"", "/", "https://x.example", "https://x.example/images/",
		"https://x.example/randomwalk", "images", "///", " padded ",
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, b string) {
		out := NormalizeNFTAssetsPublicBase(b)
		// Result is either empty or a "/images"-suffixed prefix without a trailing slash.
		if out != "" && !strings.HasSuffix(out, "/images") {
			t.Fatalf("NormalizeNFTAssetsPublicBase(%q) = %q: not empty and not /images-suffixed", b, out)
		}
		// Normalization is idempotent.
		if again := NormalizeNFTAssetsPublicBase(out); again != out {
			t.Fatalf("not idempotent: first %q, second %q", out, again)
		}
	})
}

func TestNFTImagePublicBase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		configured     string
		host           string
		tls            bool
		forwardedProto string
		forwardedHost  string
		want           string
	}{
		{
			name: "request HTTP origin",
			host: "api.example:8080",
			want: "http://api.example:8080/images",
		},
		{
			name: "request HTTPS origin",
			host: "api.example",
			tls:  true,
			want: "https://api.example/images",
		},
		{
			name:           "forwarded proto overrides TLS",
			host:           "internal.example",
			tls:            true,
			forwardedProto: " HTTP ",
			want:           "http://internal.example/images",
		},
		{
			name:           "first forwarded proto wins",
			host:           "internal.example",
			forwardedProto: " HTTPS, http ",
			want:           "https://internal.example/images",
		},
		{
			name:           "empty first forwarded proto falls back",
			host:           "internal.example",
			tls:            true,
			forwardedProto: ", http",
			want:           "https://internal.example/images",
		},
		{
			name:          "first forwarded host wins",
			host:          "internal.example",
			forwardedHost: " public.example:1443, proxy.example ",
			want:          "http://public.example:1443/images",
		},
		{
			name: "missing host",
			host: "",
			want: "",
		},
		{
			name:       "configured base overrides request origin",
			configured: " https://cdn.example/randomwalk/ ",
			host:       "ignored.example",
			want:       "https://cdn.example/images",
		},
		{
			name:       "configured base accepts images prefix",
			configured: "https://cdn.example/nested/images/",
			host:       "ignored.example",
			want:       "https://cdn.example/nested/images",
		},
		{
			name:       "configured root normalizes empty",
			configured: "/",
			host:       "ignored.example",
			want:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, "http://placeholder/", nil)
			req.Host = tt.host
			if tt.tls {
				req.TLS = &tls.ConnectionState{}
			}
			req.Header.Set("X-Forwarded-Proto", tt.forwardedProto)
			req.Header.Set("X-Forwarded-Host", tt.forwardedHost)
			c := httpx.NewContext(httptest.NewRecorder(), req)

			if got := NFTImagePublicBase(c, tt.configured); got != tt.want {
				t.Errorf("NFTImagePublicBase() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestMetadataRandomWalkImagePublicBase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		configured string
		want       string
	}{
		{
			name: "unset uses public API",
			want: "https://api.randomwalknft.com:1443/images",
		},
		{
			name:       "whitespace uses public API",
			configured: "   ",
			want:       "https://api.randomwalknft.com:1443/images",
		},
		{
			name:       "custom origin is normalized",
			configured: "https://assets.example/",
			want:       "https://assets.example/images",
		},
		{
			name:       "legacy randomwalk suffix is corrected",
			configured: " https://assets.example/randomwalk/ ",
			want:       "https://assets.example/images",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := MetadataRandomWalkImagePublicBase(tt.configured); got != tt.want {
				t.Errorf("MetadataRandomWalkImagePublicBase() = %q, want %q", got, tt.want)
			}
		})
	}
}
