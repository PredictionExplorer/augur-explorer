// Fuzz targets for the pure formatting/date helpers (MODERNIZATION.md §4.4).
// Invariants: never panic; output shape matches the documented contract.
package primitives

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func FuzzShortAddress(f *testing.F) {
	f.Add("0x1234567890abcdef1234567890abcdef12345678")
	f.Add("")
	f.Add("0x")
	f.Add("0x123456789012345678901234567890123456789")   // 41 chars
	f.Add("0x12345678901234567890123456789012345678901") // 43 chars
	f.Add("ααααααααααααααααααααα")                       // multibyte runes, 42 bytes
	f.Fuzz(func(t *testing.T, s string) {
		out := Short_address(s)
		if len(s) != 42 {
			if out != "inval_addr" {
				t.Fatalf("Short_address(%q): len %d input must yield inval_addr, got %q", s, len(s), out)
			}
			return
		}
		// 6 bytes + "…" (3 bytes) + 6 bytes
		if len(out) != 15 {
			t.Fatalf("Short_address(%q) = %q: want 15 bytes, got %d", s, out, len(out))
		}
		if !strings.Contains(out, "…") {
			t.Fatalf("Short_address(%q) = %q: missing ellipsis", s, out)
		}
		if out[:6] != s[2:8] || out[9:] != s[36:42] {
			t.Fatalf("Short_address(%q) = %q: head/tail bytes don't match input", s, out)
		}
	})
}

func FuzzShortHash(f *testing.F) {
	f.Add("0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789")
	f.Add("")
	f.Add("0x1234")
	f.Fuzz(func(t *testing.T, s string) {
		out := Short_hash(s)
		if len(s) != 66 {
			if !strings.HasPrefix(out, "inval_hash: ") {
				t.Fatalf("Short_hash(%q): len %d input must yield inval_hash, got %q", s, len(s), out)
			}
			return
		}
		if len(out) != 15 {
			t.Fatalf("Short_hash(%q) = %q: want 15 bytes, got %d", s, out, len(out))
		}
		if out[:6] != s[2:8] || out[9:] != s[59:65] {
			t.Fatalf("Short_hash(%q) = %q: head/tail bytes don't match input", s, out)
		}
	})
}

func FuzzThousandsFormat(f *testing.F) {
	for _, seed := range []int64{0, 1, -1, 999, 1000, -1000, 999999, 1000000,
		-9223372036854775808, 9223372036854775807} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, n int64) {
		out := ThousandsFormat(n)
		// Round-trip: stripping separators recovers the decimal representation.
		plain := strings.ReplaceAll(out, ",", "")
		if plain != strconv.FormatInt(n, 10) {
			t.Fatalf("ThousandsFormat(%d) = %q: strips to %q, want %q", n, out, plain, strconv.FormatInt(n, 10))
		}
		// Grouping: first group 1-3 digits (optionally signed), the rest exactly 3.
		groups := strings.Split(out, ",")
		for i, g := range groups {
			digits := g
			if i == 0 {
				digits = strings.TrimPrefix(digits, "-")
				if len(digits) < 1 || len(digits) > 3 {
					t.Fatalf("ThousandsFormat(%d) = %q: leading group %q not 1-3 digits", n, out, g)
				}
			} else if len(digits) != 3 {
				t.Fatalf("ThousandsFormat(%d) = %q: group %q not exactly 3 digits", n, out, g)
			}
		}
	})
}

func FuzzDateUtils(f *testing.F) {
	f.Add(int64(0), int64(0))
	f.Add(int64(0), int64(1))
	f.Add(int64(0), int64(86400))
	f.Add(int64(1580428800), int64(1583020800)) // 2020-01-31 → 2020-03-01
	f.Add(int64(-62135596800), int64(253402300799))
	f.Fuzz(func(t *testing.T, aSec, bSec int64) {
		a := time.Unix(aSec, 0).UTC()
		b := time.Unix(bSec, 0).UTC()
		y, mo, d, h, mi, s := TimeDifference(a, b)
		// All components normalized and non-negative.
		if y < 0 || mo < 0 || d < 0 || h < 0 || mi < 0 || s < 0 {
			t.Fatalf("TimeDifference(%v, %v) = (%d,%d,%d,%d,%d,%d): negative component", a, b, y, mo, d, h, mi, s)
		}
		if mo > 11 || d > 30 || h > 23 || mi > 59 || s > 59 {
			t.Fatalf("TimeDifference(%v, %v) = (%d,%d,%d,%d,%d,%d): component out of range", a, b, y, mo, d, h, mi, s)
		}
		// Symmetric in argument order.
		y2, mo2, d2, h2, mi2, s2 := TimeDifference(b, a)
		if y != y2 || mo != mo2 || d != d2 || h != h2 || mi != mi2 || s != s2 {
			t.Fatalf("TimeDifference not symmetric for (%v, %v)", a, b)
		}
		if aSec == bSec && (y|mo|d|h|mi|s) != 0 {
			t.Fatalf("TimeDifference of equal times = (%d,%d,%d,%d,%d,%d), want all zero", y, mo, d, h, mi, s)
		}
		// DurationToString must never panic on TimeDifference output.
		_ = DurationToString(y, mo, d, h, mi, s)
	})
}
