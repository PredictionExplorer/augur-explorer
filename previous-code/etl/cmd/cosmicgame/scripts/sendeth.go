// Makes a bid
package main

import (
	"os"
	"fmt"
	"math/big"
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"

)
const (
)
var (
	RPC_URL string
	token_addr		common.Address
	bidParamType, _	= abi.NewType("tuple","BidParams",[]abi.ArgumentMarshaling{
		{Name: "message", Type: "string"},
		{Name: "randomWalkNFTId", Type: "int256"},
	})
	params = abi.Arguments{
		{Type: bidParamType, Name: "bp"},
	}
)
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) != 4 {
		fmt.Printf("Usage: \n\t\t%v [priv_key] [to] [value]\n\n\t\tSend transaction\n",os.Args[0])
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 64 characters long\n")
		os.Exit(1)
	}

	to_addr := common.HexToAddress(os.Args[2])
	value := big.NewInt(0)
	_,success := value.SetString(os.Args[3],10)
	if !success {
		fmt.Printf("Error setting value (%v)\n",os.Args[3])
		os.Exit(1)
	}

	fmt.Printf("Sending value %v to %v\n",value.String(),to_addr.String()) 

	from_PrivateKey, err := crypto.HexToECDSA(from_pkey_str)
	if err != nil {
		fmt.Sprintf("Error making private key: %v\n",err)
		os.Exit(1)
	}
	from_publicKey := from_PrivateKey.Public()
	from_publicKeyECDSA, ok := from_publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("Couldn't derive public key for Sender\n")
		os.Exit(1)
	}
	from_address := crypto.PubkeyToAddress(*from_publicKeyECDSA)
	from_nonce, err := eclient.PendingNonceAt(context.Background(), from_address)
	if err != nil {
		fmt.Printf("Error getting account's nonce: %v\n",err)
		os.Exit(1)
	}
	gas_price, err := eclient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested gas price: %v\n",err)
		os.Exit(1)
	}

	big_chain_id, err := eclient.NetworkID(context.Background()) // Get current network ID
	if err != nil {
		fmt.Printf("Error getting network ID: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Using chain_id=%v\n",big_chain_id.String())
	tx := types.NewTransaction(from_nonce, to_addr, value, 21000, gas_price, nil)

	signed_tx, err := types.SignTx(tx, types.NewEIP155Signer(big_chain_id), from_PrivateKey)
/*	signed_tx, err := types.SignTx(tx, types.LatestSignerForChainID(big_chain_id), from_PrivateKey)
	if err != nil {
		fmt.Printf("Error signing transaction: %v\n", err)
		os.Exit(1)
	}*/
	err = eclient.SendTransaction(context.Background(), signed_tx)
	if err != nil {
		fmt.Printf("Error sending transaction: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Tx hash = %v\n",tx.Hash().String())
}
