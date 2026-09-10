package cosmicgame

import "math/big"

var (
	twoTo255 = new(big.Int).Lsh(big.NewInt(1), 255)
	twoTo256 = new(big.Int).Lsh(big.NewInt(1), 256)
)

// AsSignedInt256 reinterprets a *big.Int that was ABI-decoded as uint256 as
// the two's-complement int256 the contract actually returned.
//
// V3.1 (Solidity commit 9c04656e) removed getDurationUntilMainPrizeRaw() and
// re-typed getDurationUntilMainPrize() to return int256 — same selector, but
// the value goes negative once the main prize becomes claimable (the old
// version clamped at zero). Callers that go through the V1/V2 bindings
// (whose ABI still says uint256, matching the originally deployed contracts)
// would otherwise see a negative duration as a number near 2^256.
//
// Values below 2^255 are returned unchanged, so this is a no-op for genuine
// pre-V3.1 unsigned responses and for non-negative V3.1 values.
func AsSignedInt256(v *big.Int) *big.Int {
	if v == nil || v.Cmp(twoTo255) < 0 {
		return v
	}
	return new(big.Int).Sub(v, twoTo256)
}

// ClampNonNegative returns v, or zero if v is negative or nil. It restores
// the pre-V3.1 "duration is never negative" semantics for display and wait
// calculations after AsSignedInt256.
func ClampNonNegative(v *big.Int) *big.Int {
	if v == nil || v.Sign() < 0 {
		return big.NewInt(0)
	}
	return v
}
