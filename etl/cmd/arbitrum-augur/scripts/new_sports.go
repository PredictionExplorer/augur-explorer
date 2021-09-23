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

	factory_addr := common.HexToAddress("0x5fc8d32690cc91d4c39d9d3abcbd16989f875707")
	cur_timestamp:=time.Now().Unix()
	start_time:=big.NewInt(cur_timestamp)
	cur_timestamp = cur_timestamp + 60*60*24*5	// 5 days
	ts := big.NewInt(cur_timestamp)
	home_team_id := big.NewInt(cur_timestamp)
	away_team_id := big.NewInt(cur_timestamp)
	home_spread := big.NewInt(32)
	event_id := ts
	over_under_total := ts

	tm_factory,err := NewSportsLinkMarketFactory(factory_addr,eclient)

	var acct_nonce int64 = 11
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
	tx,err:=tm_factory.CreateMarket(auth,from_address,ts,event_id,home_team_id,away_team_id,home_spread,over_under_total,start_time)
	if err!=nil {
		fmt.Printf("Error on Deploy: %v\n",err)
		os.Exit(1)
	}
	_ = tx
	fmt.Printf("Done\n")
}
