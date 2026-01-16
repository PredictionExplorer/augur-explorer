// Approves ERC20 token spending (sets MAX_UINT256 allowance for test networks)
package main

import (
	"os"
	"fmt"
	"time"
	"math/big"
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
)
const (
	CHAIN_ID		int64 = 31337  // Local testnet; change to 421614 for Arbitrum Sepolia
)
var (
	RPC_URL string
)
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 4 {
		fmt.Printf(
			"Usage: \n\t\t%v [private_key] [erc20_token_addr] [spender_addr]\n\t\t"+
			"Approves MAX_UINT256 allowance for spender to spend your ERC20 tokens\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Private key is not 64 characters long\n")
		os.Exit(1)
	}

	token_addr := common.HexToAddress(os.Args[2])
	spender_addr := common.HexToAddress(os.Args[3])

	// MAX_UINT256 = 2^256 - 1
	max_uint256 := new(big.Int)
	max_uint256.Exp(big.NewInt(2), big.NewInt(256), nil)
	max_uint256.Sub(max_uint256, big.NewInt(1))

	erc20_ctrct,err := NewERC20(token_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract: %v\n",err)
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
		fmt.Printf("Couldn't derive public key\n")
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
	fmt.Printf("From address: %v\n", from_address.String())
	fmt.Printf("Token address: %v\n", token_addr.String())
	fmt.Printf("Spender address: %v\n", spender_addr.String())
	fmt.Printf("Allowance amount: MAX_UINT256\n")

	txopts := bind.NewKeyedTransactor(from_PrivateKey)
	txopts.Nonce = big.NewInt(int64(from_nonce))
	txopts.Value = big.NewInt(0)
	txopts.GasLimit = uint64(100000)
	txopts.GasPrice = gasPrice

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

	tx,err := erc20_ctrct.Approve(txopts, spender_addr, max_uint256)
	if err!=nil {
		fmt.Printf("Error sending approve tx: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Tx hash = %v\n",tx.Hash().String())
	fmt.Printf("Waiting for confirmation...\n")
	time.Sleep(2 * time.Second)
	fmt.Printf("Done. Approval should be set.\n")
}


