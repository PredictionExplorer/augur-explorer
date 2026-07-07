package randomwalk

import (
	"encoding/hex"
	"math"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestBeautyVoteSignMessage(t *testing.T) {
	got := beautyVoteSignMessage(42161, "abc123", 7, 9, 9)
	want := "RandomWalk beauty vote\nVersion: 1\nchainId: 42161\nnonce: abc123\nnft1: 7\nnft2: 9\nwinner: 9"
	if got != want {
		t.Fatalf("beautyVoteSignMessage = %q, want %q", got, want)
	}
}

func TestRecoverPersonalSignSignerRoundTrip(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	wantAddr := crypto.PubkeyToAddress(key.PublicKey)
	msg := beautyVoteSignMessage(1, "nonce", 1, 2, 1)

	sig, err := crypto.Sign(accounts.TextHash([]byte(msg)), key)
	if err != nil {
		t.Fatalf("sign: %v", err)
	}

	// Raw recovery id (0/1), with and without 0x prefix.
	for _, sigHex := range []string{hex.EncodeToString(sig), "0x" + hex.EncodeToString(sig)} {
		got, err := recoverPersonalSignSigner(msg, sigHex)
		if err != nil {
			t.Fatalf("recover(%q): %v", sigHex[:8], err)
		}
		if got != wantAddr {
			t.Fatalf("recovered %s, want %s", got.Hex(), wantAddr.Hex())
		}
	}

	// Wallets commonly emit v = 27/28; the helper must normalize it.
	adjusted := make([]byte, len(sig))
	copy(adjusted, sig)
	adjusted[64] += 27
	got, err := recoverPersonalSignSigner(msg, hex.EncodeToString(adjusted))
	if err != nil {
		t.Fatalf("recover with v+27: %v", err)
	}
	if got != wantAddr {
		t.Fatalf("recovered %s with v+27, want %s", got.Hex(), wantAddr.Hex())
	}

	// A different message must not recover the same signer.
	got, err = recoverPersonalSignSigner(msg+"tampered", hex.EncodeToString(sig))
	if err == nil && got == wantAddr {
		t.Fatal("tampered message recovered the original signer")
	}
}

func TestRecoverPersonalSignSignerRejectsBadInput(t *testing.T) {
	cases := []struct {
		name string
		sig  string
	}{
		{"empty", ""},
		{"non-hex", "zz"},
		{"too short", "aabb"},
		{"64 bytes", strings.Repeat("ab", 64)},
		{"66 bytes", strings.Repeat("ab", 66)},
	}
	for _, tc := range cases {
		if _, err := recoverPersonalSignSigner("msg", tc.sig); err == nil {
			t.Errorf("%s: expected error, got none", tc.name)
		}
	}
}

func FuzzRecoverPersonalSignSigner(f *testing.F) {
	f.Add("hello", strings.Repeat("ab", 65))
	f.Add("", "")
	f.Add("msg", "0x"+strings.Repeat("00", 65))
	f.Add("msg", strings.Repeat("ff", 65))
	f.Fuzz(func(t *testing.T, message, sigHex string) {
		addr, err := recoverPersonalSignSigner(message, sigHex)
		if err != nil {
			return
		}
		// Success is only possible for well-formed 65-byte signatures.
		raw := strings.TrimPrefix(strings.TrimSpace(sigHex), "0x")
		decoded, decErr := hex.DecodeString(raw)
		if decErr != nil || len(decoded) != 65 {
			t.Fatalf("recover succeeded for invalid signature input %q", sigHex)
		}
		if addr == (ethcommon.Address{}) {
			t.Fatalf("recover succeeded but returned the zero address for %q", sigHex)
		}
	})
}

func TestComputeEloUpdate(t *testing.T) {
	// First match ever (k=250), equal ratings, player A wins:
	// expected score 0.5, so A gains 125 and B loses 125.
	raNew, rbNew := computeEloUpdate(1000, 1000, 1.0, 0)
	if raNew != 1125 || rbNew != 875 {
		t.Fatalf("computeEloUpdate(1000,1000,1,0) = (%v,%v), want (1125,875)", raNew, rbNew)
	}
	// K-factor floors at 1 once enough matches have been played.
	raNew, rbNew = computeEloUpdate(1000, 1000, 1.0, 1_000_000)
	if raNew != 1000.5 || rbNew != 999.5 {
		t.Fatalf("k-floor case = (%v,%v), want (1000.5,999.5)", raNew, rbNew)
	}
}

func FuzzEloUpdate(f *testing.F) {
	f.Add(1000.0, 1000.0, true, int64(0))
	f.Add(875.5, 1223.25, false, int64(52000))
	f.Add(0.0, 0.0, true, int64(1_000_000_000))
	f.Add(-500.0, 3000.0, false, int64(-1))
	f.Fuzz(func(t *testing.T, ra, rb float64, aWins bool, totalMatches int64) {
		if math.IsNaN(ra) || math.IsInf(ra, 0) || math.IsNaN(rb) || math.IsInf(rb, 0) {
			t.Skip("ratings from the DB are always finite")
		}
		score := 0.0
		if aWins {
			score = 1.0
		}
		raNew, rbNew := computeEloUpdate(ra, rb, score, totalMatches)
		if math.IsNaN(raNew) || math.IsNaN(rbNew) {
			t.Fatalf("computeEloUpdate(%v,%v,%v,%d) produced NaN: (%v,%v)", ra, rb, score, totalMatches, raNew, rbNew)
		}
		if aWins {
			if raNew < ra {
				t.Fatalf("winner's rating decreased: %v -> %v", ra, raNew)
			}
			if rbNew > rb {
				t.Fatalf("loser's rating increased: %v -> %v", rb, rbNew)
			}
		} else {
			if rbNew < rb {
				t.Fatalf("winner's rating decreased: %v -> %v", rb, rbNew)
			}
			if raNew > ra {
				t.Fatalf("loser's rating increased: %v -> %v", ra, raNew)
			}
		}
		// Elo is zero-sum: the pair's total rating is conserved (up to float rounding).
		sumBefore := ra + rb
		sumAfter := raNew + rbNew
		tol := 1e-9 + 1e-12*(math.Abs(ra)+math.Abs(rb)+math.Abs(raNew)+math.Abs(rbNew))
		if diff := math.Abs(sumAfter - sumBefore); diff > tol && !math.IsInf(sumAfter, 0) {
			t.Fatalf("rating sum not conserved: before %v, after %v (diff %v)", sumBefore, sumAfter, diff)
		}
	})
}
