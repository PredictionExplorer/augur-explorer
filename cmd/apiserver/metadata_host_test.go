package main

import (
	"strings"
	"testing"
)

func TestMetadataHostServesCosmicSignature(t *testing.T) {
	cases := []struct {
		name string
		host string
		xfh  string
		want bool
	}{
		{"cosmic host", "api.cosmicsignature.com", "", true},
		{"cosmic host uppercase", "API.COSMICSIGNATURE.COM", "", true},
		{"randomwalk host", "api.randomwalknft.com", "", false},
		{"empty host", "", "", false},
		{"xfh overrides host to cosmic", "api.randomwalknft.com", "www.cosmicsignature.com", true},
		{"xfh overrides host to randomwalk", "cosmicsignature.com", "randomwalknft.com", false},
		{"xfh first entry wins", "x", "cosmicsignature.com, randomwalknft.com", true},
		{"xfh first entry wins (reversed)", "x", "randomwalknft.com, cosmicsignature.com", false},
		{"xfh trimmed", "x", "  cosmicsignature.com  ", true},
		{"xfh uppercase", "x", "COSMICSIGNATURE.COM", true},
		// Existing quirk, pinned: a malformed X-Forwarded-Host consisting of
		// just a comma blanks the effective host instead of falling back.
		{"xfh lone comma blanks host", "cosmicsignature.com", ",", false},
	}
	for _, tc := range cases {
		if got := metadataHostServesCosmicSignature(tc.host, tc.xfh); got != tc.want {
			t.Errorf("%s: metadataHostServesCosmicSignature(%q, %q) = %v, want %v",
				tc.name, tc.host, tc.xfh, got, tc.want)
		}
	}
}

func FuzzMetadataHostDispatch(f *testing.F) {
	f.Add("api.cosmicsignature.com", "")
	f.Add("api.randomwalknft.com", "cosmicsignature.com, other.example")
	f.Add("", ",")
	f.Add("HOST", "  spaced.example  ")
	f.Fuzz(func(t *testing.T, host, xfh string) {
		got := metadataHostServesCosmicSignature(host, xfh)

		// Case-insensitive in both inputs.
		if upper := metadataHostServesCosmicSignature(strings.ToUpper(host), strings.ToUpper(xfh)); upper != got {
			t.Fatalf("case sensitivity: (%q,%q)=%v but uppercased=%v", host, xfh, got, upper)
		}
		// Only the first comma-separated X-Forwarded-Host entry may matter:
		// appending further entries can never change the decision.
		if xfh != "" && !strings.Contains(xfh, ",") {
			if smuggled := metadataHostServesCosmicSignature(host, xfh+",evil.cosmicsignature.example"); smuggled != got {
				t.Fatalf("appended forwarded-host entry changed dispatch: (%q,%q)=%v, smuggled=%v",
					host, xfh, got, smuggled)
			}
		}
		// When no forwarded host is present, the decision depends on Host alone.
		if xfh == "" {
			want := strings.Contains(strings.ToLower(host), "cosmicsignature")
			if got != want {
				t.Fatalf("host-only dispatch mismatch for %q: got %v, want %v", host, got, want)
			}
		}
	})
}
