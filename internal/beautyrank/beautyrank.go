// Package beautyrank implements the mechanics of the RandomWalk
// beauty-contest ranking mini-app shared by the frozen v1 API and the v2
// ranking slice: the Elo rating update, the canonical EIP-191 vote message
// with its signer recovery, and the one-time challenge nonce format.
//
// The vote message layout and the Elo parameters are wire and data
// contracts: wallets sign the exact message text, and every recorded match
// row was produced by this update rule. Neither may change shape without a
// coordinated frontend migration (ADR-0008).
package beautyrank

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// DefaultRating is the Elo rating of a token that has never been part of a
// recorded match. The store queries embed the same value as their COALESCE
// default.
const DefaultRating = 1200

// ChallengeTTL is how long an issued vote challenge nonce stays valid.
const ChallengeTTL = 15 * time.Minute

// VoteMessage builds the canonical text a wallet signs for one beauty vote.
// The format is byte-frozen: deployed frontends construct the identical
// string, so any change would invalidate every in-flight vote.
func VoteMessage(chainID int64, nonce string, first, second, winner int64) string {
	return fmt.Sprintf(
		"RandomWalk beauty vote\nVersion: 1\nchainId: %d\nnonce: %s\nnft1: %d\nnft2: %d\nwinner: %d",
		chainID, nonce, first, second, winner,
	)
}

// RecoverSigner recovers the wallet that EIP-191 personal_sign-ed message.
// It accepts the common wallet signature encodings: optional 0x prefix and
// a recovery id of 0/1 or 27/28. Every malformed input fails closed with an
// error and the zero address.
func RecoverSigner(message, signatureHex string) (ethcommon.Address, error) {
	raw := strings.TrimSpace(signatureHex)
	raw = strings.TrimPrefix(raw, "0x")
	sig, err := hex.DecodeString(raw)
	if err != nil {
		return ethcommon.Address{}, fmt.Errorf("decode signature: %w", err)
	}
	if len(sig) != 65 {
		return ethcommon.Address{}, errors.New("signature must be 65 bytes")
	}
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}
	h := accounts.TextHash([]byte(message))
	pub, err := crypto.SigToPub(h, sig)
	if err != nil {
		return ethcommon.Address{}, err
	}
	return crypto.PubkeyToAddress(*pub), nil
}

// EloUpdate computes both new ratings after one match. scoreA is 1 when the
// first token won and 0 when it lost. The K factor starts at 250 and decays
// by 0.00525 per already-recorded match with a floor of 1, so early votes
// converge the ranking quickly and a mature ranking moves slowly. The
// update is zero-sum around the standard 400-point logistic expectation.
func EloUpdate(ratingA, ratingB, scoreA float64, totalMatches int64) (newA, newB float64) {
	k := 250.0 - float64(totalMatches)*0.00525
	if k < 1 {
		k = 1
	}
	expectedA := 1.0 / (1.0 + math.Pow(10, (ratingB-ratingA)/400))
	newA = ratingA + k*(scoreA-expectedA)
	newB = ratingB - k*(scoreA-expectedA)
	return newA, newB
}

// NewNonce draws a 32-byte challenge nonce from entropy and returns it as
// 64 lowercase hex characters.
func NewNonce(entropy io.Reader) (string, error) {
	var b [32]byte
	if _, err := io.ReadFull(entropy, b[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(b[:]), nil
}
