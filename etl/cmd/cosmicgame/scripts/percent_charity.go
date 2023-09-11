// Sets charity percentage
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

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	CHAIN_ID		int64 = 31337
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

	if len(os.Args) < 4 {
		fmt.Printf("Usage: \n\t\t%v [priv_key] [contract_addr] [percentage]\n\n\t\tChanges percentage of charity eth deposits\n",os.Args[0])
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 66 characters long\n")
		os.Exit(1)
	}

	cosmic_game_addr := common.HexToAddress(os.Args[2])
	percentage_str := os.Args[3]

	percentage_val := big.NewInt(0)
	_,success := percentage_val.SetString(percentage_str,10)
	if !success {
		fmt.Printf("Incorrect percentage value provided on the command line")
		os.Exit(1)
	}

	cosmic_game_ctrct,err := NewCosmicGame(cosmic_game_addr,eclient)
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

	tx,err := cosmic_game_ctrct.SetCharityPercentage(txopts,percentage_val)
	fmt.Printf("Tx hash: %v\n",tx.Hash().String())
	if err!=nil {
		fmt.Printf("Error sending tx: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Tx hash = %v\n",tx.Hash().String())
}
