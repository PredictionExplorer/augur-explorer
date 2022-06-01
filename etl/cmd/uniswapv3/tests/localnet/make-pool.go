package main

import (
	"fmt"
	"os"
	"strconv"
	"math/big"
	//"encoding/hex"
	"crypto/ecdsa"
	"context"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	RPC_URL					string
	contract_abi			abi.ABI
	CHAIN_ID				int64 = 1234
)

func main() {

	if len(os.Args) < 6 {
		fmt.Printf(
			"Usage: \n\t\t%v [private_key] [factory_addr] [token0_addr] [token1_addr] [fee]\n\t\t"+
			"Creates a new pool\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 64 characters long\n")
		os.Exit(1)
	}
	factory_addr := common.HexToAddress(os.Args[2])
	token0_addr := common.HexToAddress(os.Args[3])
	token1_addr := common.HexToAddress(os.Args[4])

	fee,err := strconv.ParseInt(os.Args[5],10,64)
	if err != nil {
		fmt.Printf("Error parsing fee field: %v\n",err)
		os.Exit(1)
	}
	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Conected to %v\n",RPC_URL)
	from_PrivateKey, err := crypto.HexToECDSA(from_pkey_str)
	if err!=nil{
		fmt.Printf("Error : %v\n",err)
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
	auth := bind.NewKeyedTransactor(from_PrivateKey)
	auth.Nonce = big.NewInt(int64(from_nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(9500000)
	auth.GasPrice = gasPrice
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
	auth.Signer = signfunc
	factory,err := NewUniswapV3Factory(factory_addr, eclient)
	if err != nil {
		fmt.Printf("Error creating Uniswap Factory instance: %v\n",err)
		os.Exit(1)
	}

	tx,err:=factory.CreatePool(auth, token0_addr,token1_addr,big.NewInt(fee))
	if err != nil {
		fmt.Printf("Error submitting tx: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Transaction submitted, pool address will be provided in PoolCreated event\n")
	fmt.Printf("Tx hash: %v\n",tx.Hash().String())
	_ = tx
}
