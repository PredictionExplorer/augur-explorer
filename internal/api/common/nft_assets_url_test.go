package common

import (
	"strings"
	"testing"
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
