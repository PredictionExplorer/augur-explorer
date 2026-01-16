// Sets the round activation time to a specific timestamp or offset from current block time
package main

import (
	"os"
	"fmt"
	"math/big"
	"context"
	"crypto/ecdsa"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
)
var (
	RPC_URL    string
	chain_id   *big.Int
)
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}
	chain_id,err = eclient.ChainID(context.Background())
	if err != nil {
		fmt.Printf("Error getting chain id : %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) != 4 {
		fmt.Printf("Usage: \n\t\t%v [priv_key] [contract_addr] [seconds_offset]\n\n"+
			"\t\tSets round activation time to (current_block_time + seconds_offset)\n"+
			"\t\tUse 0 to activate immediately, or negative to activate in the past\n",os.Args[0])
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 64 characters long\n")
		os.Exit(1)
	}

	cosmic_game_addr := common.HexToAddress(os.Args[2])

	offset,err := strconv.ParseInt(os.Args[3],10,64)
	if err != nil {
		fmt.Printf("Error parsing seconds offset: %v\n",err)
		os.Exit(1)
	}

	// Get current block timestamp
	latest_block, err := eclient.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting latest block: %v\n",err)
		os.Exit(1)
	}
	block_time := int64(latest_block.Time())
	fmt.Printf("Current block timestamp: %v\n", block_time)

	activation_time := big.NewInt(block_time + offset)
	fmt.Printf("Setting round activation time to: %v (offset: %v seconds)\n", activation_time.String(), offset)

	cosmic_game_ctrct,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}

	from_PrivateKey, err := crypto.HexToECDSA(from_pkey_str)
	if err != nil {
		fmt.Printf("Error making private key: %v\n",err)
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
	gasPrice, err := eclient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested gas price: %v\n",err)
		os.Exit(1)
	}

	fmt.Printf("Using chain_id=%v\n",chain_id.String())
	fmt.Printf("From address: %v\n", from_address.String())

	txopts := bind.NewKeyedTransactor(from_PrivateKey)
	txopts.Nonce = big.NewInt(int64(from_nonce))
	txopts.Value = big.NewInt(0)
	txopts.GasLimit = uint64(100000)
	txopts.GasPrice = gasPrice

	signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(chain_id)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), from_PrivateKey)
		if err != nil {
			fmt.Printf("Error signing: %v\n",err)
			os.Exit(1)
			return nil,nil
		}
		return tx.WithSignature(signer, signature)
	}
	txopts.Signer = signfunc

	tx,err := cosmic_game_ctrct.SetRoundActivationTime(txopts, activation_time)
	if err!=nil {
		fmt.Printf("Error sending tx: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Tx hash: %v\n",tx.Hash().String())
	fmt.Printf("Done! Round activation time set.\n")
}

