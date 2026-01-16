package main
import (
	"fmt"
	"os"
	"math/big"
)
func Pow(a *big.Float, e uint64) *big.Float {
	result := big.NewFloat(0.0)
	result.Set(a)
	for i:=uint64(0); i<e-1; i++ {
		result.Mul(result, a)
	}
	return result
}
func main() {

	if len(os.Args) != 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [sqrt_priceX96]\n\n"+
			"\t\tconverts sartPriceX96 to price\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	sqrtP_str := os.Args[1]
	sqrtP := big.NewInt(0)
	_,success := sqrtP.SetString(sqrtP_str,10)
	if !success {
		fmt.Printf("Bad price integer provided\n")
		os.Exit(1)
	}

	ninety_six := big.NewInt(96)
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
	final_result := Pow(result,2)
	fmt.Printf("Price: %v\n",final_result.String())
}

