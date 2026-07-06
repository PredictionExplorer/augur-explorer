package ethtx

import (
	"math/big"
	"os"
	"strconv"
)

// Gas limit defaults by transaction type, based on typical gas usage of the
// CosmicGame contracts.
const (
	// GasLimitERC20Approve is for ERC-20 approve() calls.
	GasLimitERC20Approve = uint64(100000)

	// GasLimitBid is for CosmicGame bidding operations.
	GasLimitBid = uint64(500000)

	// GasLimitClaimPrize is for CosmicGame prize claiming (complex operation; V2 needs ~3M).
	GasLimitClaimPrize = uint64(3500000)

	// GasLimitDonate is for CosmicGame donation operations.
	GasLimitDonate = uint64(300000)

	// GasLimitAdminCall is for owner-only setter operations.
	GasLimitAdminCall = uint64(100000)
)

// gasPriceMultiplier returns the multiplier applied to the RPC-suggested gas
// price so transactions stay above the block base fee. The default is 2.0;
// override with the GAS_PRICE_MULTIPLIER environment variable (e.g. 1.5, 3).
func gasPriceMultiplier() *big.Float {
	if s := os.Getenv("GAS_PRICE_MULTIPLIER"); s != "" {
		if f, err := strconv.ParseFloat(s, 64); err == nil && f > 0 {
			return big.NewFloat(f)
		}
	}
	return big.NewFloat(2.0)
}

// AdjustGasPrice applies the GAS_PRICE_MULTIPLIER policy to the base price and
// returns the adjusted gas price.
func AdjustGasPrice(basePrice *big.Int) *big.Int {
	if basePrice == nil {
		return big.NewInt(0)
	}
	multiplier := gasPriceMultiplier()
	if multiplier.Cmp(big.NewFloat(1.0)) == 0 {
		return basePrice
	}
	adjusted := new(big.Float).SetInt(basePrice)
	adjusted.Mul(adjusted, multiplier)
	result := new(big.Int)
	adjusted.Int(result)
	return result
}
