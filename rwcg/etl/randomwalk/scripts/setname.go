// Sets name on a token
package main

import (
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/randomwalk"
)

func main() {
	cutils.ParseInfoFlag()
	cutils.GasPriceMultiplier = big.NewFloat(2.0)
	if len(os.Args) != 4 {
		cutils.PrintUsage(os.Args[0],
			"[rwalk_addr] [token_id] [new_name]",
			"Sets the display name for a RandomWalk token.",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)", "PKEY_HEX": "64-char hex private key, no 0x prefix (required)"},
		)
		os.Exit(1)
	}

	rwalkAddr := common.HexToAddress(os.Args[1])
	tokenID, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		cutils.Fatal("Invalid token_id: %v", err)
	}
	newName := os.Args[3]

	net, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network connection failed: %v", err)
	}
	cutils.PrintNetworkInfo(net)

	acc, err := cutils.PrepareAccount(net, cutils.MustGetPkeyHex())
	if err != nil {
		cutils.Fatal("Account setup failed: %v", err)
	}
	cutils.PrintAccountInfo(acc)

	rwalkCtrct, err := NewRWalk(rwalkAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate RWalk contract: %v", err)
	}

	cutils.Section("SET NAME INFO")
	cutils.PrintKeyValue("Token ID", tokenID)
	cutils.PrintKeyValue("New Name", newName)

	gasLimit := cutils.GasLimitContractCall
	txopts := cutils.CreateTransactOpts(net, acc, big.NewInt(0), gasLimit)
	cutils.PrintTxSubmitting("SetTokenName", big.NewInt(0), gasLimit, cutils.AdjustGasPrice(net.GasPrice))

	tx, err := rwalkCtrct.SetTokenName(txopts, big.NewInt(tokenID), newName)
	if !cutils.PrintTxResultAndWait(net.Client, tx, err) {
		os.Exit(1)
	}
}
