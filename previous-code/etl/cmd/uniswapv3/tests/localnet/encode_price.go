package main
import (
	"fmt"
	"os"
	"math/big"
	"strconv"
)
func main() {

	if len(os.Args) != 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [token_amount0] [token_amount1]\n\n"+
			"\t\tConverts balances of two tokens into SqrtPrice96\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	a0,err := strconv.ParseInt(os.Args[1],10,64)
	if err != nil {
		fmt.Printf("Error parsing amount0: %v\n",err)
		os.Exit(1)
	}
	a1,err := strconv.ParseInt(os.Args[1],10,64)
	if err != nil {
		fmt.Printf("Error parsing amount1: %v\n",err)
		os.Exit(1)
	}
	amount0:=big.NewInt(a0)
	amount1:=big.NewInt(a1)
	num := big.NewInt(0).Set(amount1)
	num.Quo(num,amount0)
	pow := big.NewInt(2)
	pow.Exp(pow,big.NewInt(96),nil)
	num.Mul(num,pow)
	fmt.Printf("%v\n",num.String())
}




