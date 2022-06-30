// 
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

	if len(os.Args) != 7 {
		fmt.Printf(
			"Usage: \n\t\t%v [priv_key] [pos_mgr_addr] [pool_addr] [tick_lower] [tick_upper] [amount(delta)]\n\n"+
			"\t\tAdds liquidity to the pool through position manager contract (periphery)\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 66 characters long\n")
		os.Exit(1)
	}

	pos_mgr_addr := common.HexToAddress(os.Args[2])
	pool_addr := common.HexToAddress(os.Args[3])

	tick_lower_str := os.Args[4]
	tick_lower:= big.NewInt(0)
	_,success := tick_lower.SetString(tick_lower_str,10)
	if !success {
		fmt.Printf("Incorrect tick_lower number provided on the command line")
		os.Exit(1)
	}
	tick_upper_str := os.Args[5]
	tick_upper:= big.NewInt(0)
	_,success = tick_upper.SetString(tick_upper_str,10)
	if !success {
		fmt.Printf("Incorrect tick_upper number provided on the command line")
		os.Exit(1)
	}

	amount_str := os.Args[6]
	desired_amount := big.NewInt(0)
	_,success = desired_amount.SetString(amount_str,10)
	if !success {
		fmt.Printf("Incorrect amount provided on the command line")
		os.Exit(1)
	}

	pool_ctrct,err := NewUniswapV3Pool(pool_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Wrapped ETH contract: %v\n",err)
		os.Exit(1)
	}

	var copts = new(bind.CallOpts)
	token0_addr,err := pool_ctrct.Token0(copts)
	if err != nil {
		fmt.Printf("Error getting token0 address: %v\n",err)
		os.Exit(1)
	}
	token1_addr,err := pool_ctrct.Token1(copts)
	if err != nil {
		fmt.Printf("Error getting token1 address: %v\n",err)
		os.Exit(1)
	}
	fee,err := pool_ctrct.Fee(copts)
	if err != nil {
		fmt.Printf("Error getting fee: %v\n",err)
		os.Exit(1)
	}

	pos_mgr,err := NewNonfungiblePositionManager(pos_mgr_addr,eclient)
	if err !=  nil {
		fmt.Printf("Error instantiating NonFungiblePositionManager: %v\n",err)
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

	var mint_params INonfungiblePositionManagerMintParams
	mint_params.Token0 = token0_addr
	mint_params.Token1 = token1_addr
	mint_params.Fee = big.NewInt(0).Set(fee)
	mint_params.TickLower = tick_lower
	mint_params.TickUpper = tick_upper
	mint_params.Amount0Desired = big.NewInt(0).Set(desired_amount)
	mint_params.Amount1Desired = big.NewInt(0).Set(desired_amount)
	mint_params.Amount0Min = big.NewInt(0)
	mint_params.Amount1Min = big.NewInt(0)
	//mint_params.Amount0Min = big.NewInt(0).Set(mint_params.Amount0Desired)
	//mint_params.Amount1Min = big.NewInt(0).Set(mint_params.Amount1Desired)
	mint_params.Recipient = from_address
	ts := time.Now().Unix()
	ts = ts + 60*60*24*5	// shift deadline 5 day ahead of current time
	mint_params.Deadline = big.NewInt(ts)

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

	tx,err := pos_mgr.Mint(txopts,mint_params)
	fmt.Printf("Tx hash: %v\n",tx.Hash().String())
	if err!=nil {
		fmt.Printf("Error sending tx: %v\n",err)
		os.Exit(1)
	}
}
