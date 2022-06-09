// 
package main

import (
	"os"
	"fmt"
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

	if len(os.Args) != 6 {
		fmt.Printf(
			"Usage: \n\t\t%v [priv_key] [pool_addr] [tick_lower] [tick_upper] [amount(delta)]\n\n"+
			"\t\tAdds liquidity to the pool\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 66 characters long\n")
		os.Exit(1)
	}

	pool_addr := common.HexToAddress(os.Args[2])

	tick_lower_str := os.Args[3]
	tick_lower:= big.NewInt(0)
	_,success := tick_lower.SetString(tick_lower_str,10)
	if !success {
		fmt.Printf("Incorrect tick_lower number provided on the command line")
		os.Exit(1)
	}
	tick_upper_str := os.Args[4]
	tick_upper:= big.NewInt(0)
	_,success = tick_upper.SetString(tick_upper_str,10)
	if !success {
		fmt.Printf("Incorrect tick_upper number provided on the command line")
		os.Exit(1)
	}

	amount_str := os.Args[5]
	delta_amount := big.NewInt(0)
	_,success = delta_amount.SetString(amount_str,10)
	if !success {
		fmt.Printf("Incorrect amount provided on the command line")
		os.Exit(1)
	}

	pool_ctrct,err := NewUniswapV3Pool(pool_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Wrapped ETH contract: %v\n",err)
		os.Exit(1)
	}

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
	big_chain_id := big.NewInt(CHAIN_ID)
	fmt.Printf("Using chain_id=%v\n",big_chain_id.String())
	txopts := bind.NewKeyedTransactor(from_PrivateKey)
	txopts.Nonce = big.NewInt(int64(from_nonce))
	txopts.Value = big.NewInt(0)     // in weia
	txopts.GasLimit = uint64(10000000) // in units
	txopts.GasPrice = gasPrice

	fmt.Printf("Gas price = %v\n",gasPrice.String())

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

	tx,err := pool_ctrct.Mint(txopts,from_address,tick_lower,tick_upper,delta_amount,nil)
	fmt.Printf("Tx hash: %v\n",tx.Hash().String())
	if err!=nil {
		fmt.Printf("Error sending tx: %v\n",err)
		os.Exit(1)
	}
}
