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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
)
const (
	
//	CHAIN_ID		int64 = 31337
	CHAIN_ID		int64 = 421614
//	CHAIN_ID		int64 = 11155111
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
		fmt.Printf("Usage: \n\t\t%v [priv_key] [contract_addr] [new_owner]\n\n\t\tChanges owner of a contract\n",os.Args[0])
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 66 characters long\n")
		os.Exit(1)
	}

	ctrct_addr := common.HexToAddress(os.Args[2])
	new_owner_addr := common.HexToAddress(os.Args[3])

	ownable_ctrct,err := NewOwnable(ctrct_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
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
	txopts.GasPrice.Add(txopts.GasPrice,gasPrice);

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

	fmt.Printf("Tranfering ownership to %v\n",new_owner_addr.String())
	tx,err := ownable_ctrct.TransferOwnership(txopts,new_owner_addr)
	fmt.Printf("Tx hash: %v\n",tx.Hash().String())
	if err!=nil {
		fmt.Printf("Error sending tx: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Tx hash = %v\n",tx.Hash().String())
}
