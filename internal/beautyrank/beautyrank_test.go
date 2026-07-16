package beautyrank

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io"
	"math"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// TestVoteMessage pins the byte-frozen wire format wallets sign. Any change
// here is a breaking frontend change (ADR-0008).
func TestVoteMessage(t *testing.T) {
	t.Parallel()
	got := VoteMessage(42161, "abc123", 7, 9, 9)
	want := "RandomWalk beauty vote\nVersion: 1\nchainId: 42161\nnonce: abc123\nnft1: 7\nnft2: 9\nwinner: 9"
	if got != want {
		t.Fatalf("VoteMessage = %q, want %q", got, want)
	}
	if got := VoteMessage(0, "", 0, 0, 0); !strings.HasPrefix(got, "RandomWalk beauty vote\nVersion: 1\n") {
		t.Fatalf("zero-value message lost its fixed header: %q", got)
	}
}

func TestRecoverSignerRoundTrip(t *testing.T) {
	t.Parallel()
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}
	wantAddr := crypto.PubkeyToAddress(key.PublicKey)
	msg := VoteMessage(1, "nonce", 1, 2, 1)

	sig, err := crypto.Sign(accounts.TextHash([]byte(msg)), key)
	if err != nil {
		t.Fatalf("sign: %v", err)
	}

	// Raw recovery id (0/1), with and without 0x prefix, and with
	// surrounding whitespace as wallets occasionally paste it.
	for _, sigHex := range []string{
		hex.EncodeToString(sig),
		"0x" + hex.EncodeToString(sig),
		"  0x" + hex.EncodeToString(sig) + "\n",
	} {
		got, err := RecoverSigner(msg, sigHex)
		if err != nil {
			t.Fatalf("recover(%q...): %v", sigHex[:8], err)
		}
		if got != wantAddr {
			t.Fatalf("recovered %s, want %s", got.Hex(), wantAddr.Hex())
		}
	}

	// Wallets commonly emit v = 27/28; the helper must normalize it.
	adjusted := make([]byte, len(sig))
	copy(adjusted, sig)
	adjusted[64] += 27
	got, err := RecoverSigner(msg, hex.EncodeToString(adjusted))
	if err != nil {
		t.Fatalf("recover with v+27: %v", err)
	}
	if got != wantAddr {
		t.Fatalf("recovered %s with v+27, want %s", got.Hex(), wantAddr.Hex())
	}

	// A tampered message must not attribute the vote to the signer.
	got, err = RecoverSigner(msg+"tampered", hex.EncodeToString(sig))
	if err == nil && got == wantAddr {
		t.Fatal("tampered message recovered the original signer")
	}
}

func TestRecoverSignerRejectsBadInput(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		sig  string
	}{
		{"empty", ""},
		{"non-hex", "zz"},
		{"odd length hex", "abc"},
		{"too short", "aabb"},
		{"64 bytes", strings.Repeat("ab", 64)},
		{"66 bytes", strings.Repeat("ab", 66)},
		{"0x only", "0x"},
		{"internal whitespace", "ab cd"},
	}
	for _, tc := range cases {
		if _, err := RecoverSigner("msg", tc.sig); err == nil {
			t.Errorf("%s: expected error, got none", tc.name)
		}
	}
}

func FuzzRecoverSigner(f *testing.F) {
	f.Add("hello", strings.Repeat("ab", 65))
	f.Add("", "")
	f.Add("msg", "0x"+strings.Repeat("00", 65))
	f.Add("msg", strings.Repeat("ff", 65))
	f.Fuzz(func(t *testing.T, message, sigHex string) {
		addr, err := RecoverSigner(message, sigHex)
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

func TestEloUpdate(t *testing.T) {
	t.Parallel()
	// First match ever (k=250), equal ratings, player A wins:
	// expected score 0.5, so A gains 125 and B loses 125.
	newA, newB := EloUpdate(1000, 1000, 1.0, 0)
	if newA != 1125 || newB != 875 {
		t.Fatalf("EloUpdate(1000,1000,1,0) = (%v,%v), want (1125,875)", newA, newB)
	}
	// A loss mirrors the win.
	newA, newB = EloUpdate(1000, 1000, 0.0, 0)
	if newA != 875 || newB != 1125 {
		t.Fatalf("EloUpdate(1000,1000,0,0) = (%v,%v), want (875,1125)", newA, newB)
	}
	// K-factor floors at 1 once enough matches have been played.
	newA, newB = EloUpdate(1000, 1000, 1.0, 1_000_000)
	if newA != 1000.5 || newB != 999.5 {
		t.Fatalf("k-floor case = (%v,%v), want (1000.5,999.5)", newA, newB)
	}
	// The decay boundary: at 47,429 matches K is still above 1; the exact
	// floor engagement point is (250-1)/0.00525 ≈ 47,428.57.
	beforeFloorA, _ := EloUpdate(1000, 1000, 1.0, 47_428)
	atFloorA, _ := EloUpdate(1000, 1000, 1.0, 47_429)
	if gain := beforeFloorA - 1000; gain <= 0.5 {
		t.Fatalf("pre-floor K gain = %v, want > 0.5", gain)
	}
	if gain := atFloorA - 1000; gain != 0.5 {
		t.Fatalf("floored K gain = %v, want exactly 0.5 (K=1, expectation 0.5)", gain)
	}
	// A 400-point favorite that wins gains little: expectation ~0.909.
	newA, _ = EloUpdate(1600, 1200, 1.0, 0)
	if gain := newA - 1600; gain < 22 || gain > 23 {
		t.Fatalf("favorite win gain = %v, want ~22.7", gain)
	}
}

func FuzzEloUpdate(f *testing.F) {
	f.Add(1000.0, 1000.0, true, int64(0))
	f.Add(875.5, 1223.25, false, int64(52000))
	f.Add(0.0, 0.0, true, int64(1_000_000_000))
	f.Add(-500.0, 3000.0, false, int64(-1))
	f.Fuzz(func(t *testing.T, ratingA, ratingB float64, aWins bool, totalMatches int64) {
		if math.IsNaN(ratingA) || math.IsInf(ratingA, 0) || math.IsNaN(ratingB) || math.IsInf(ratingB, 0) {
			t.Skip("ratings from the DB are always finite")
		}
		score := 0.0
		if aWins {
			score = 1.0
		}
		newA, newB := EloUpdate(ratingA, ratingB, score, totalMatches)
		if math.IsNaN(newA) || math.IsNaN(newB) {
			t.Fatalf("EloUpdate(%v,%v,%v,%d) produced NaN: (%v,%v)", ratingA, ratingB, score, totalMatches, newA, newB)
		}
		if aWins {
			if newA < ratingA {
				t.Fatalf("winner's rating decreased: %v -> %v", ratingA, newA)
			}
			if newB > ratingB {
				t.Fatalf("loser's rating increased: %v -> %v", ratingB, newB)
			}
		} else {
			if newB < ratingB {
				t.Fatalf("winner's rating decreased: %v -> %v", ratingB, newB)
			}
			if newA > ratingA {
				t.Fatalf("loser's rating increased: %v -> %v", ratingA, newA)
			}
		}
		// Elo is zero-sum: the pair's total rating is conserved (up to float rounding).
		sumBefore := ratingA + ratingB
		sumAfter := newA + newB
		tol := 1e-9 + 1e-12*(math.Abs(ratingA)+math.Abs(ratingB)+math.Abs(newA)+math.Abs(newB))
		if diff := math.Abs(sumAfter - sumBefore); diff > tol && !math.IsInf(sumAfter, 0) {
			t.Fatalf("rating sum not conserved: before %v, after %v (diff %v)", sumBefore, sumAfter, diff)
		}
	})
}

type errorReader struct {
	err error
}

func (r errorReader) Read([]byte) (int, error) {
	return 0, r.err
}

func TestNewNonce(t *testing.T) {
	t.Parallel()
	input := make([]byte, 32)
	for i := range input {
		input[i] = byte(i)
	}
	got, err := NewNonce(bytes.NewReader(input))
	if err != nil {
		t.Fatalf("NewNonce: %v", err)
	}
	if want := hex.EncodeToString(input); got != want {
		t.Fatalf("NewNonce = %q, want %q", got, want)
	}
	if len(got) != 64 {
		t.Fatalf("nonce length = %d, want 64 hex characters", len(got))
	}

	t.Run("short entropy source", func(t *testing.T) {
		t.Parallel()
		nonce, err := NewNonce(bytes.NewReader(input[:31]))
		if !errors.Is(err, io.ErrUnexpectedEOF) {
			t.Fatalf("error = %v, want io.ErrUnexpectedEOF", err)
		}
		if nonce != "" {
			t.Fatalf("nonce = %q on entropy failure, want empty", nonce)
		}
	})

	t.Run("reader failure", func(t *testing.T) {
		t.Parallel()
		wantErr := errors.New("entropy unavailable")
		nonce, err := NewNonce(errorReader{err: wantErr})
		if !errors.Is(err, wantErr) {
			t.Fatalf("error = %v, want %v", err, wantErr)
		}
		if nonce != "" {
			t.Fatalf("nonce = %q on entropy failure, want empty", nonce)
		}
	})
}
