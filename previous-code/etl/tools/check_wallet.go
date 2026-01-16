// Checks if an EOA (signing account) has Wallet contract deplyed.
package main

import (
	"os"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"


	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	// WalletRegistry contract address in MainNet
	WALLET_REG_ADDR_PROD = "0x9Fa160f92A10b431F255BF1a70a1c1e5808E5128"
)
var (
	wallet_reg_addr common.Address = common.HexToAddress(WALLET_REG_ADDR_PROD)
	RPC_URL string
)
func main() {

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		Fatalf("Can't connect to ETH RPC: %v\n",err)
	}

	ctrct_wallet_reg,err := NewAugurWalletRegistry(wallet_reg_addr,eclient)
	if err!=nil {
		Fatalf("Failed to instantiate AugurWalletRegistry contract: %v\n",err)
	}

	if len(os.Args) < 2 {
		fmt.Printf("Usage: \n\t\t%v [EOA_address]\n",os.Args[0])
		os.Exit(1)
	}

	eoa_addr := common.HexToAddress(os.Args[1])

	var copts = new(bind.CallOpts)
	wallet_addr,err:=ctrct_wallet_reg.GetWallet(copts,eoa_addr)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
		os.Exit(2)
	} else {
		if !Eth_addr_is_zero(&wallet_addr) {
			fmt.Printf("Creation status: exists at %v\n",wallet_addr.String())
			os.Exit(0)
		} else {
			fmt.Printf("Creation status: DOES NOT EXIST\n")
			wallet_addr,err = ctrct_wallet_reg.GetCreate2WalletAddress(copts,eoa_addr)
			if err!=nil {
				fmt.Printf("Couldn't calculate wallet addr: %v\n",err)
			} else {
				fmt.Printf("Contract address when created will be: %v\n",wallet_addr.String())
			}
			os.Exit(2)
		}
	}
}
