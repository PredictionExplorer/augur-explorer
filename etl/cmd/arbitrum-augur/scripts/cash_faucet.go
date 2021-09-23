package main

import (
	"fmt"
	"os"
	"math/big"
	"crypto/ecdsa"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)

var (
	RPC_URL string

	_ = abi.U256
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v [private_key]\n")
		os.Exit(1)
	}

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Conected to %v\n",RPC_URL)

	pkey_hex := os.Args[1]

	cash_addr := common.HexToAddress("0x7a5c71af12bd0bb639e6a67d19c8cbaa57273be0")

	cash_ctrct,err := NewAMMCash(cash_addr,eclient)

	from_PrivateKey, err := crypto.HexToECDSA(pkey_hex)
	if err!=nil{
		fmt.Printf("Error : %v\n",err)
		os.Exit(1)
	}
	from_publicKey := from_PrivateKey.Public()
	from_publicKeyECDSA, ok := from_publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("Cant derive public key\n")
		os.Exit(1)
	}

	from_address := crypto.PubkeyToAddress(*from_publicKeyECDSA)
	fmt.Printf("Sending Faucet() call from addr %v\n",from_address.String())
	from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
	if err != nil {
		fmt.Printf("Error getting nonce: %v\n",err)
		os.Exit(1)
	}

	auth := bind.NewKeyedTransactor(from_PrivateKey)
	auth.Nonce = big.NewInt(int64(from_nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(9500000)
	auth.GasPrice = big.NewInt(1200000000)
	fmt.Printf("Waiting for output..\n")
	amount := big.NewInt(0)
	amount.SetString("100000000",10)	// Cash is 6 decimals
	fmt.Printf("Amount requested: %v\n",amount.String())
	tx,err:=cash_ctrct.Faucet(auth,amount)
	if err!=nil {
		fmt.Printf("Error on Deploy: %v\n",err)
		os.Exit(1)
	}
	tx_hash := tx.Hash().String()
	fmt.Printf("Tx hash %v\n",tx_hash)
	_ = tx
	fmt.Printf("Done\n")
}
