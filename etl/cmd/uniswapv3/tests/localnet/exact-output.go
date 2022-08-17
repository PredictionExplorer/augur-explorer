package main

import (
	"os"
	"fmt"
	"time"
	"math/big"
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"


)
const (
	CHAIN_ID		int64 = 1234
)
var (
	RPC_URL string
	token_addr		common.Address
)
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) != 8 {
		fmt.Printf(
			"Usage: \n\t\t%v [priv_key] [swaprouter_addr] [token_in_addr] [tokne_out_addr] [fee] [amount_out] [price_limit]\n\n"+
			"\t\tSends exactInputSingle() function call to SwapRouter contract\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 66 characters long\n")
		os.Exit(1)
	}

	router_addr := common.HexToAddress(os.Args[2])
	token_in_addr := common.HexToAddress(os.Args[3])
	token_out_addr := common.HexToAddress(os.Args[4])
	fee_str := os.Args[5]
	amount_out_str := os.Args[6]
	price_limit_str := os.Args[7]

	fee := big.NewInt(0)
	_,success := fee.SetString(fee_str,10)
	if !success {
		fmt.Printf("Incorrect fee value")
		os.Exit(1)
	}
	amount_out := big.NewInt(0)
	_,success = amount_out.SetString(amount_out_str,10)
	if !success {
		fmt.Printf("Incorrect amount_out")
		os.Exit(1)
	}
	price_limit := big.NewInt(0)
	_,success = price_limit.SetString(price_limit_str,10)
	if !success {
		fmt.Printf("bad price limit value")
		os.Exit(1)
	}
	router_ctrct,err := NewSwapRouter(router_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Uniswap v3 pool contract: %v\n",err)
		os.Exit(1)
	}

	var params ISwapRouterExactOutputSingleParams

	from_PrivateKey, err := crypto.HexToECDSA(from_pkey_str)
	if err != nil {
		fmt.Sprintf("Error making private key: %v\n",err)
		os.Exit(1)
	}
	from_publicKey := from_PrivateKey.Public()
	from_publicKeyECDSA, ok := from_publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("Couldn't derive public key for Sender")
		os.Exit(1)
	}
	one := big.NewInt(1)
	from_address := crypto.PubkeyToAddress(*from_publicKeyECDSA)
	params.Recipient = from_address
	ts := time.Now().Unix()
	ts = ts + 60*60*24*5    // shift deadline 5 day ahead of current time
	params.Deadline = big.NewInt(ts)
	params.TokenIn = token_in_addr
	params.TokenOut = token_out_addr
	params.Fee = big.NewInt(0).Set(fee)
	params.AmountOut  = big.NewInt(0).Set(amount_out)
	params.AmountInMaximum = big.NewInt(2147483647)
	params.AmountInMaximum.Mul(params.AmountInMaximum,big.NewInt(65536))

	sqrtPriceLimitX96 := big.NewInt(0)
	sqrtPriceLimitX96.Set(price_limit)
	sqrtPriceLimitX96.Add(sqrtPriceLimitX96,one)
	params.SqrtPriceLimitX96 = big.NewInt(0).Set(sqrtPriceLimitX96)

	from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
	if err != nil {
		fmt.Printf("Error getting account's nonce: %v\n",err)
		os.Exit(1)
	}
	gasPrice, err := eclient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested gas price: %v\n",err)
		os.Exit(1)
	}
	big_chain_id := big.NewInt(CHAIN_ID)
	fmt.Printf("Using chain_id=%v\n",big_chain_id.String())
	txopts := bind.NewKeyedTransactor(from_PrivateKey)
	txopts.Nonce = big.NewInt(int64(from_nonce))
	txopts.Value = big.NewInt(0)     // in weia
	txopts.GasLimit = uint64(10000000) // in units
	txopts.GasPrice = gasPrice

	fmt.Printf("Gas price = %v\n",gasPrice.String())
	fmt.Printf("Amount Out  = %v\n",params.AmountOut.String())
	fmt.Printf("Amount in max = %v\n",params.AmountInMaximum.String())

	signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(big_chain_id)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_PrivateKey)
		if err != nil {
			fmt.Printf("Error signing: %v\n",err)
			os.Exit(1)
			return nil,nil
		}
		return tx.WithSignature(signer, signature)
	}
	txopts.Signer = signfunc

	tx,err := router_ctrct.ExactOutputSingle(txopts,params)
	fmt.Printf("Tx hash: %v\n",tx.Hash().String())
	if err!=nil {
		fmt.Printf("Error sending tx: %v\n",err)
		os.Exit(1)
	}
}
