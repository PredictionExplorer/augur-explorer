package common

import (
	"net/http/httptest"
	"strings"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// validateAddr runs IsAddressValid against a throwaway gin context and also
// returns the response body written on rejection.
func validateAddr(addr string) (result string, ok bool, body string) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	result, ok = IsAddressValid(c, true, addr)
	return result, ok, w.Body.String()
}

func TestIsAddressValid(t *testing.T) {
	const lower = "0xc0ffee254729296a45a3885639ac7e10f9d54979"
	const checksummed = "0xc0ffee254729296a45a3885639AC7E10F9d54979"

	cases := []struct {
		name   string
		in     string
		wantOK bool
		want   string
	}{
		{"lowercase with 0x", lower, true, checksummed},
		{"already checksummed", checksummed, true, checksummed},
		{"no 0x prefix", lower[2:], true, checksummed},
		{"empty", "", false, ""},
		{"too short", "0x1234", false, ""},
		{"too long", lower + "00", false, ""},
		{"41 chars", lower[:41], false, ""},
		{"non-hex 40 chars", strings.Repeat("zz", 20), false, ""},
	}
	for _, tc := range cases {
		got, ok, body := validateAddr(tc.in)
		if ok != tc.wantOK {
			t.Errorf("%s: IsAddressValid(%q) ok = %v, want %v", tc.name, tc.in, ok, tc.wantOK)
			continue
		}
		if ok && got != tc.want {
			t.Errorf("%s: IsAddressValid(%q) = %q, want %q", tc.name, tc.in, got, tc.want)
		}
		if !ok && !strings.Contains(body, `"error"`) {
			t.Errorf("%s: rejection must write the JSON error envelope, body = %q", tc.name, body)
		}
	}
}

func FuzzIsAddressValid(f *testing.F) {
	f.Add("0xC0ffee254729296a45a3885639AC7E10F9d54979")
	f.Add("c0ffee254729296a45a3885639ac7e10f9d54979")
	f.Add("")
	f.Add("0x")
	f.Add(strings.Repeat("0", 42))
	f.Add("0X" + strings.Repeat("a", 40)) // uppercase 0X prefix is NOT stripped
	f.Fuzz(func(t *testing.T, addr string) {
		result, ok, body := validateAddr(addr)
		if !ok {
			if body == "" {
				t.Fatalf("IsAddressValid(%q): rejected without writing an error response", addr)
			}
			return
		}
		// Accepted values must come back as the canonical EIP-55 form.
		if len(result) != 42 {
			t.Fatalf("IsAddressValid(%q) accepted with non-address result %q", addr, result)
		}
		if ethcommon.HexToAddress(result).Hex() != result {
			t.Fatalf("IsAddressValid(%q) = %q which is not EIP-55 checksummed", addr, result)
		}
	})
}
