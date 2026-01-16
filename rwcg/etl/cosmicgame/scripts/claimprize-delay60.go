// Sets delayDurationBeforeRoundActivation, then claims prize
// This ensures the new round activates after the specified delay
package main

import (
	"os"
	"fmt"
	"math/big"
	"time"
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

	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Printf("Usage: \n\t\t%v [priv_key] [contract_addr] [delay_seconds]\n\n"+
			"\t\tSets delayDurationBeforeRoundActivation, then claims prize\n"+
			"\t\tdelay_seconds defaults to 60 if not provided\n",os.Args[0])
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 64 characters long\n")
		os.Exit(1)
	}

	cosmic_game_addr := common.HexToAddress(os.Args[2])

	delay_seconds := int64(60)
	if len(os.Args) == 4 {
		delay_seconds, err = strconv.ParseInt(os.Args[3], 10, 64)
		if err != nil {
			fmt.Printf("Error parsing delay_seconds: %v\n", err)
			os.Exit(1)
		}
	}

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

	fmt.Printf("Using chain_id=%v\n",chain_id.String())
	fmt.Printf("From address: %v\n", from_address.String())
	fmt.Printf("Delay seconds: %v\n", delay_seconds)

	// Step 1: Set delayDurationBeforeRoundActivation
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

	txopts := bind.NewKeyedTransactor(from_PrivateKey)
	txopts.Nonce = big.NewInt(int64(from_nonce))
	txopts.Value = big.NewInt(0)
	txopts.GasLimit = uint64(100000)
	txopts.GasPrice = gasPrice
	txopts.Signer = signfunc

	fmt.Printf("Step 1: Setting delayDurationBeforeRoundActivation to %v seconds...\n", delay_seconds)
	tx,err := cosmic_game_ctrct.SetDelayDurationBeforeRoundActivation(txopts, big.NewInt(delay_seconds))
	if err!=nil {
		fmt.Printf("Error setting delay: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("SetDelayDurationBeforeRoundActivation tx hash: %v\n",tx.Hash().String())

	// Wait a moment for the tx to be mined
	fmt.Printf("Waiting for tx to be mined...\n")
	time.Sleep(2 * time.Second)

	// Step 2: Claim the prize (will use the delay we just set)
	from_nonce, err = eclient.PendingNonceAt(context.Background(), from_address)
	if err != nil {
		fmt.Printf("Error getting account's nonce: %v\n",err)
		os.Exit(1)
	}
	gasPrice, err = eclient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("Error getting suggested gas price: %v\n",err)
		os.Exit(1)
	}

	txopts2 := bind.NewKeyedTransactor(from_PrivateKey)
	txopts2.Nonce = big.NewInt(int64(from_nonce))
	txopts2.Value = big.NewInt(0)
	txopts2.GasLimit = uint64(10000000)
	txopts2.GasPrice = gasPrice
	txopts2.Signer = signfunc

	fmt.Printf("Step 2: Claiming prize...\n")
	tx2,err := cosmic_game_ctrct.ClaimMainPrize(txopts2)
	if err!=nil {
		fmt.Printf("Error claiming prize: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("ClaimMainPrize tx hash: %v\n",tx2.Hash().String())

	fmt.Printf("\nDone! Delay set to %v seconds and prize claimed.\n", delay_seconds)
	fmt.Printf("New round will activate %v seconds after the claim.\n", delay_seconds)
}
