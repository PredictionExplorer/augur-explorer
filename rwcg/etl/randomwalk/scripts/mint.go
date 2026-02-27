// Mints a token
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
			"[rwalk_addr] [amount_wei]",
			"Mints a RandomWalk token by sending amount_wei to the contract.",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)", "PKEY_HEX": "64-char hex private key, no 0x prefix (required)"},
		)
		os.Exit(1)
	}

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

	rwalkAddr := common.HexToAddress(os.Args[1])
	amount := big.NewInt(0)
	if _, ok := amount.SetString(os.Args[2], 10); !ok {
		cutils.Fatal("Invalid amount: %s", os.Args[2])
	}

	rwalkCtrct, err := NewRWalk(rwalkAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate RWalk contract: %v", err)
	}

	cutils.Section("MINT INFO")
	cutils.PrintKeyValueEth("Amount (wei)", amount)
	gasLimit := cutils.GasLimitHighComplexity
	totalNeeded := new(big.Int).Mul(net.GasPrice, big.NewInt(int64(gasLimit)))
	totalNeeded.Add(totalNeeded, amount)
	if acc.Balance.Cmp(totalNeeded) < 0 {
		cutils.Fatal("Insufficient balance. Need %s ETH (amount + gas), have %s ETH",
			cutils.WeiToEth(totalNeeded), cutils.WeiToEth(acc.Balance))
	}

	txopts := cutils.CreateTransactOpts(net, acc, amount, gasLimit)
	cutils.PrintTxSubmitting("Mint", amount, gasLimit, cutils.AdjustGasPrice(net.GasPrice))

	tx, err := rwalkCtrct.Mint(txopts)
	if !cutils.PrintTxResultAndWait(net.Client, tx, err) {
		os.Exit(1)
	}
}
