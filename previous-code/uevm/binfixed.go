package uevm				// EVM for Uniswap (v3)
import (
	"fmt"
	"math/big"
)
func BinaryFixedToBigFloat(precision int64,value string) *big.Float {

	sqrtP := big.NewInt(0)
	_,success := sqrtP.SetString(value,10)
	if !success {
		panic(fmt.Sprintf("Bad value integer provided at uevm::BinaryFixedToBigFloat() function\n"))
	}

	ninety_six := big.NewInt(precision)
	two := big.NewInt(2)
	exponent := big.NewInt(2)
	exponent.Exp(exponent,ninety_six,nil)
	d := big.NewFloat(0.0)
	f_sqrtP := big.NewFloat(0.0)
	f_sqrtP.SetInt(sqrtP)
	f_exponent := big.NewFloat(0.0)
	f_exponent.SetInt(exponent)
	d.Quo(f_sqrtP,f_exponent)

	f_two := big.NewFloat(0.0)
	f_two.SetInt(two)
	result:=big.NewFloat(0.0)
	result.Set(d)
	return result
}
