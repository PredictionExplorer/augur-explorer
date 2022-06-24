package main

import (
	"fmt"
	"os"
	"strconv"
	"math/big"
	"bytes"
	//"encoding/hex"
	"crypto/ecdsa"
	"context"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	RPC_URL					string
	contract_abi			abi.ABI
	CHAIN_ID				int64 = 1234
)

func main() {

	if len(os.Args) < 5 {
		fmt.Printf(
			"Usage: \n\t\t%v [private_key] [pool_addr] [token0_in_addr] [amount_in]\n\t\t"+
			"Swaps token_in for the other token\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 64 characters long\n")
		os.Exit(1)
	}
	pool_addr := common.HexToAddress(os.Args[2])
	token_in_addr := common.HexToAddress(os.Args[3])

	amount_in_int64,err := strconv.ParseInt(os.Args[4],10,64)
	if err != nil {
		fmt.Printf("Error parsing amount_in field: %v\n",err)
		os.Exit(1)
	}
	amount_in := big.NewInt(amount_in_int64)
	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Conected to %v\n",RPC_URL)
	var copts = new(bind.CallOpts)
	from_PrivateKey, err := crypto.HexToECDSA(from_pkey_str)
	if err!=nil{
		fmt.Printf("Error : %v\n",err)
		os.Exit(1)
	}
	from_publicKey := from_PrivateKey.Public()
	from_publicKeyECDSA, ok := from_publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("Couldn't derive public key for Sender")
		os.Exit(1)
	}
	from_address := crypto.PubkeyToAddress(*from_publicKeyECDSA)
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

	pool_ctrct,err := NewUniswapV3Pool(pool_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Uniswap v3 pool contract: %v\n",err)
		os.Exit(1)
	}

	token0_addr,err := pool_ctrct.Token0(copts)
	if err != nil {
		fmt.Printf("Error calling Token0(): %v\n",err)
		os.Exit(1)
	}
	token1_addr,err := pool_ctrct.Token1(copts)
	if err != nil {
		fmt.Printf("Error calling Token1(): %v\n",err)
		os.Exit(1)
	}
	slot0,err := pool_ctrct.Slot0(copts)
	if err != nil {
		fmt.Printf("Error calling Slot()\n",err)
		os.Exit(1)
	}
	one := big.NewInt(1)
	sqrtPriceLimitX96 := big.NewInt(0)
	sqrtPriceLimitX96.Set(slot0.SqrtPriceX96)
	sqrtPriceLimitX96.Add(sqrtPriceLimitX96,one)
	zero_for_one := false
	if bytes.Equal(token0_addr.Bytes(),token_in_addr.Bytes()) {
		zero_for_one = true
	} else {
		if bytes.Equal(token1_addr.Bytes(),token_in_addr.Bytes()) {
			/// nothing, it is already 0
		} else {
			fmt.Printf("Error, token_in address doesn't belong to tokens the pool was registered\n")
			os.Exit(1)
		}
	}
	big_chain_id := big.NewInt(CHAIN_ID)
	auth := bind.NewKeyedTransactor(from_PrivateKey)
	auth.Nonce = big.NewInt(int64(from_nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(9500000)
	auth.GasPrice = gasPrice
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
	auth.Signer = signfunc

	tx,err:=pool_ctrct.Swap(auth, from_address,zero_for_one,amount_in,sqrtPriceLimitX96,nil)
	if err != nil {
		fmt.Printf("Error submitting tx: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Tx hash: %v\n",tx.Hash().String())
	_ = tx
}
