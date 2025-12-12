// Sends a new offer to the marketplace
package main

import (
	"os"
	"fmt"
	"time"
	"math/big"
	"strconv"
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/randomwalk"
)
const (
	CHAIN_ID		int64 = 421611
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

	if len(os.Args) < 7 {
		fmt.Printf(
			"Usage: \n\t\t%v [private_key] [BUY|SELL] [market_addr] [nft_addr] [token_id] [price]\n\t\t"+
			"Sends [amount] for making BUY or SELL offer to MarketPlace contract\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 64 characters long\n")
		os.Exit(1)
	}
	method := os.Args[2]
	if (method != "BUY") && (method != "SELL") {
		fmt.Printf("Invalid operation, BUY or SELL is the only valid code\n")
		os.Exit(1)
	}
	market_addr := common.HexToAddress(os.Args[3])
	nft_addr := common.HexToAddress(os.Args[4])
	token_id,err := strconv.ParseInt(os.Args[5],10,64)
	if err != nil {
		fmt.Printf("Error parsing token_id field: %v\n",err)
		os.Exit(1)
	}

	amount_str := os.Args[6]
	amount := big.NewInt(0)
	_,success := amount.SetString(amount_str,10)
	if !success {
		fmt.Printf("Couldn't parse amount, bad integer\n")
		os.Exit(1)
	}

	market_ctrct,err := NewRWMarket(market_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate RWMarket contract: %v\n",err)
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
	fmt.Printf("Market Addr %v\n",market_addr.String())
	fmt.Printf("Token ID: %v\n",token_id)
	txopts := bind.NewKeyedTransactor(from_PrivateKey)
	txopts.Nonce = big.NewInt(int64(from_nonce))
	txopts.Value = big.NewInt(0)     // in wei
	txopts.GasLimit = uint64(10000000) // in units
	txopts.GasPrice = gasPrice
	fmt.Printf("Gas price = %v\n",gasPrice.String())

	signfunc := func(signer_disabled types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
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

	var tx *types.Transaction
	if method == "BUY" {
		txopts.Value.Set(amount)
		fmt.Printf("Setting msg.value to %v\n",txopts.Value.String())
		tx,err = market_ctrct.MakeBuyOffer(txopts,nft_addr,big.NewInt(token_id))
	} else {
		if method == "SELL" {
			tx,err = market_ctrct.MakeSellOffer(txopts,nft_addr,big.NewInt(token_id),big.NewInt(0).Set(amount))
		} else {
			fmt.Printf("Unknown method: %v\n",method)
			os.Exit(1)
		}
	}

	if err!=nil {
		fmt.Printf("Error sending tx: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Tx hash = %v\n",tx.Hash().String())
	time.Sleep(1 * time.Second)

}
