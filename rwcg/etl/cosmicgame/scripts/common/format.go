package common

import (
	"fmt"
	"math/big"
	"strings"
)

// WeiToEth converts wei to ETH as a formatted string with 18 decimal places
func WeiToEth(wei *big.Int) string {
	if wei == nil {
		return "0.000000000000000000"
	}
	ether := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(ether, big.NewFloat(1e18))
	return ethValue.Text('f', 18)
}

// WeiToEthCompact converts wei to ETH with fewer decimal places for display
func WeiToEthCompact(wei *big.Int) string {
	if wei == nil {
		return "0.0000"
	}
	ether := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(ether, big.NewFloat(1e18))
	return ethValue.Text('f', 6)
}

// WeiToGwei converts wei to gwei as a float64
func WeiToGwei(wei *big.Int) float64 {
	if wei == nil {
		return 0
	}
	gwei := new(big.Float).SetInt(wei)
	gweiValue := new(big.Float).Quo(gwei, big.NewFloat(1e9))
	result, _ := gweiValue.Float64()
	return result
}

// FormatTokenAmount formats a token amount with the given decimals
func FormatTokenAmount(amount *big.Int, decimals uint8, symbol string) string {
	if amount == nil {
		return fmt.Sprintf("0 %s", symbol)
	}
	if decimals == 0 {
		return fmt.Sprintf("%s %s", amount.String(), symbol)
	}

	divisorStr := "1" + strings.Repeat("0", int(decimals))
	divisor := new(big.Int)
	divisor.SetString(divisorStr, 10)

	whole := new(big.Int)
	remainder := new(big.Int)
	whole.QuoRem(amount, divisor, remainder)

	// Format with decimal point
	format := fmt.Sprintf("%%s.%%0%ds %%s", decimals)
	return fmt.Sprintf(format, whole.String(), remainder.String(), symbol)
}

// FmtDuration formats seconds into human readable duration
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

// ConvertToPercentage converts a divisor to a percentage
// e.g., divisor of 100 = 1%, divisor of 10 = 10%
func ConvertToPercentage(divisor *big.Int) float64 {
	if divisor == nil || divisor.Cmp(big.NewInt(0)) == 0 {
		return 0
	}
	one := big.NewFloat(1)
	hundred := big.NewFloat(100)
	divisorFloat := new(big.Float).SetInt(divisor)
	fraction := new(big.Float).Quo(one, divisorFloat)
	percent := new(big.Float).Mul(fraction, hundred)
	result, _ := percent.Float64()
	return result
}

// TruncateAddress returns a shortened address like 0x1234...5678
func TruncateAddress(addr string) string {
	if len(addr) <= 12 {
		return addr
	}
	return addr[:6] + "..." + addr[len(addr)-4:]
}

// MaxUint256 returns 2^256 - 1
func MaxUint256() *big.Int {
	maxUint256 := new(big.Int)
	maxUint256.Exp(big.NewInt(2), big.NewInt(256), nil)
	maxUint256.Sub(maxUint256, big.NewInt(1))
	return maxUint256
}
