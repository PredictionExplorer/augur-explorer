package common

import (
	"math/big"
	"os"
	"strconv"
)

// Gas limit constants by transaction type.
// These are reasonable defaults based on typical gas usage.
const (
	// GasLimitSimpleTransfer is for plain ETH transfers (21000 is the exact cost)
	GasLimitSimpleTransfer = uint64(21000)

	// GasLimitERC20Approve is for ERC20 approve() calls
	GasLimitERC20Approve = uint64(100000)

	// GasLimitERC20Transfer is for ERC20 transfer() calls
	GasLimitERC20Transfer = uint64(100000)

	// GasLimitERC721Approve is for ERC721 setApprovalForAll() calls
	GasLimitERC721Approve = uint64(100000)

	// GasLimitBid is for CosmicGame bidding operations
	GasLimitBid = uint64(500000)

	// GasLimitClaimPrize is for CosmicGame prize claiming (complex operation)
	GasLimitClaimPrize = uint64(2000000)

	// GasLimitDonate is for CosmicGame donation operations
	GasLimitDonate = uint64(300000)

	// GasLimitContractCall is default for miscellaneous contract calls
	GasLimitContractCall = uint64(300000)

	// GasLimitAdminCall is for admin/setter operations
	GasLimitAdminCall = uint64(100000)

	// GasLimitHighComplexity is for very complex operations
	GasLimitHighComplexity = uint64(5000000)
)

// GasPriceMultiplier bumps gas price for faster inclusion and to stay above block base fee.
// Default is 2.0 (double suggested price). Override with env GAS_PRICE_MULTIPLIER (e.g. 1.5, 2, 3).
var GasPriceMultiplier = defaultGasPriceMultiplier()

func defaultGasPriceMultiplier() *big.Float {
	if s := os.Getenv("GAS_PRICE_MULTIPLIER"); s != "" {
		if f, err := strconv.ParseFloat(s, 64); err == nil && f > 0 {
			return big.NewFloat(f)
		}
	}
	return big.NewFloat(2.0)
}

// AdjustGasPrice applies the GasPriceMultiplier to the base price.
// Returns the adjusted gas price.
func AdjustGasPrice(basePrice *big.Int) *big.Int {
	if basePrice == nil {
		return big.NewInt(0)
	}
	one := big.NewFloat(1.0)
	if GasPriceMultiplier.Cmp(one) == 0 {
		return basePrice
	}
	adjusted := new(big.Float).SetInt(basePrice)
	adjusted.Mul(adjusted, GasPriceMultiplier)
	result := new(big.Int)
	adjusted.Int(result)
	return result
}

// BumpGasPrice adds a fixed amount to the gas price.
// Useful for avoiding "replacement transaction underpriced" errors.
func BumpGasPrice(basePrice *big.Int, bumpWei int64) *big.Int {
	if basePrice == nil {
		return big.NewInt(bumpWei)
	}
	return new(big.Int).Add(basePrice, big.NewInt(bumpWei))
}
