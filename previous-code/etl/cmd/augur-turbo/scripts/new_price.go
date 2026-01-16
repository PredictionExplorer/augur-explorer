package main

import (
	"fmt"
	"os"
	"time"
	"math/big"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)

var (
	PKEY_HEX = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	RPC_URL string

	_ = abi.U256
)

func main() {

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Conected to %v\n",RPC_URL)

	factory_addr := common.HexToAddress("0x0165878a594ca255338adfa4d48449f69242eb8f")
	cur_timestamp:=time.Now().Unix()
	ts := big.NewInt(cur_timestamp)
	description := "Trusted market example"
	outcome_names := [3]string{"Invalid","One","Two"}
	outcome_symbols := [3]string{"I","1","2"}

	tm_factory,err := NewPrieMarketFactory(factory_addr,eclient)

	var acct_nonce int64 = 9
	from_PrivateKey, err := crypto.HexToECDSA(PKEY_HEX)
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
	auth := bind.NewKeyedTransactor(from_PrivateKey)
	auth.Nonce = big.NewInt(int64(acct_nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(9500000)
	auth.GasPrice = big.NewInt(1000000000)
	fmt.Printf("Waiting for output..\n")
	tx,err:=tm_factory.CreateMarket(auth,from_address,ts,description,outcome_names[:],outcome_symbols[:])
	if err!=nil {
		fmt.Printf("Error on Deploy: %v\n",err)
		os.Exit(1)
	}
	_ = tx
	fmt.Printf("Done\n")
}
