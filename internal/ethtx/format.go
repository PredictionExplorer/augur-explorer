package ethtx

// Display formatting shared by the operator CLIs: token amounts, durations,
// percentages. Wei-to-ETH conversion lives in output.go (WeiToEthText).

import (
	"fmt"
	"math/big"
	"strings"
)

// WeiToEthCompact converts wei to ETH with six decimal places for compact
// display (gas-cost estimates).
func WeiToEthCompact(wei *big.Int) string {
	if wei == nil {
		return "0.0000"
	}
	ether := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(ether, big.NewFloat(1e18))
	return ethValue.Text('f', 6)
}

// FormatTokenAmount formats a base-unit token amount using the given decimals
// and symbol, e.g. FormatTokenAmount(1500000, 6, "USDC") = "1.500000 USDC".
func FormatTokenAmount(amount *big.Int, decimals uint8, symbol string) string {
	if amount == nil {
		return fmt.Sprintf("0 %s", symbol)
	}
	if decimals == 0 {
		return fmt.Sprintf("%s %s", amount.String(), symbol)
	}

	divisor := new(big.Int)
	divisor.SetString("1"+strings.Repeat("0", int(decimals)), 10)

	whole := new(big.Int)
	remainder := new(big.Int)
	whole.QuoRem(amount, divisor, remainder)

	format := fmt.Sprintf("%%s.%%0%ds %%s", decimals)
	return fmt.Sprintf(format, whole.String(), remainder.String(), symbol)
}

// FmtDuration formats seconds into a human-readable duration.
func FmtDuration(secs int64) string {
	if secs < 0 {
		return fmt.Sprintf("%d sec (negative)", secs)
	}
	if secs == 0 {
		return "0 sec"
	}
	if secs < 60 {
		return fmt.Sprintf("%d sec", secs)
	}
	if secs < 3600 {
		return fmt.Sprintf("%d min %d sec", secs/60, secs%60)
	}
	hours := secs / 3600
	mins := (secs % 3600) / 60
	sec := secs % 60
	return fmt.Sprintf("%dh %dm %ds", hours, mins, sec)
}

// ConvertToPercentage converts a divisor to the percentage it represents,
// e.g. a divisor of 100 is 1% and a divisor of 10 is 10%.
func ConvertToPercentage(divisor *big.Int) float64 {
	if divisor == nil || divisor.Sign() == 0 {
		return 0
	}
	fraction := new(big.Float).Quo(big.NewFloat(1), new(big.Float).SetInt(divisor))
	percent := new(big.Float).Mul(fraction, big.NewFloat(100))
	result, _ := percent.Float64()
	return result
}

// MaxUint256 returns 2^256 - 1, the unlimited ERC-20 allowance value.
func MaxUint256() *big.Int {
	limit := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	return limit.Sub(limit, big.NewInt(1))
}
