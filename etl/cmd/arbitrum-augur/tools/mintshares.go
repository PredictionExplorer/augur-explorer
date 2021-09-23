// Mints for market/outcome (in exchange for liquidity (collateral))
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

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	CHAIN_ID		int64 = 80001
)
var (
	RPC_URL string
	token_addr		common.Address
)
func main() {

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 4 {
		fmt.Printf(
			"Usage: \n\t\t%v [private_key] [market_factory] [market_id] [amount_to_mint]\n\t\t"+
			"mints shares\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	//var copts = new(bind.CallOpts)

	from_pkey_str := os.Args[1]
	if len(from_pkey_str) != 64 {
		fmt.Printf("Sender's private key is not 66 characters long (including 0x prefix)\n")
		os.Exit(1)
	}

	market_factory_addr := common.HexToAddress(os.Args[2])
	market_id,err := strconv.ParseInt(os.Args[3],10,64)
	if err != nil {
		fmt.Printf("Error parsing market_id field: %v\n",err)
		os.Exit(1)
	}
	mint_amount_str := os.Args[4]
	mint_amount := big.NewInt(0)
	_,success := mint_amount.SetString(mint_amount_str,10)
	if !success {
		fmt.Printf("Couldn't parse mint amount, bad integer\n")
		os.Exit(1)
	}

	market_factory,err := NewSportsLinkMarketFactory(market_factory_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Market Factory contract: %v\n",err)
		os.Exit(1)
	}

	big_market_id := big.NewInt(market_id)
	/*market_obj,err:=market_factory.GetMarket(copts,big_market_id)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}*/

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
	fmt.Printf("mint shares, market_id=%v\n",big_market_id.String())
	fmt.Printf("Using chain_id=%v\n",big_chain_id.String())
	fmt.Printf("Mint amount = %v\n",mint_amount.String())
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

	tx,err := market_factory.MintShares(txopts,big_market_id,mint_amount,from_address)
	if err!=nil {
		fmt.Printf("Error sending tx: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Tx hash = %v\n",tx.Hash().String())
	time.Sleep(1 * time.Second)

}
