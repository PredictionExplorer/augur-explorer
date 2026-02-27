// Approves an address on RandomWalk token
package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/randomwalk"
)

func main() {
	cutils.ParseInfoFlag()
	cutils.GasPriceMultiplier = big.NewFloat(2.0)
	if len(os.Args) != 3 {
		cutils.PrintUsage(os.Args[0],
			"[rwalk_addr] [operator_addr]",
			"Sets ERC721 approval for all (SetApprovalForAll) for the given operator.",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)", "PKEY_HEX": "64-char hex private key, no 0x prefix (required)"},
		)
		os.Exit(1)
	}

	rwalkAddr := common.HexToAddress(os.Args[1])
	operatorAddr := common.HexToAddress(os.Args[2])

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

	cutils.Section("APPROVAL INFO")
	cutils.PrintKeyValue("Operator", operatorAddr.String())

	gasLimit := cutils.GasLimitERC721Approve
	txopts := cutils.CreateTransactOpts(net, acc, big.NewInt(0), gasLimit)
	cutils.PrintTxSubmitting("SetApprovalForAll", big.NewInt(0), gasLimit, cutils.AdjustGasPrice(net.GasPrice))

	tx, err := rwalkCtrct.SetApprovalForAll(txopts, operatorAddr, true)
	if !cutils.PrintTxResultAndWait(net.Client, tx, err) {
		os.Exit(1)
	}
}
