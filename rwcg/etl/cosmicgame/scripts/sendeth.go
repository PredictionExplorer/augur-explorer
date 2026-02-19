// Sends ETH to an address
package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	cutils.ParseInfoFlag()
	// Usage check
	if len(os.Args) != 3 {
		cutils.PrintUsage(os.Args[0],
			"[to_address] [value_wei]",
			"Sends ETH to an address",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)", "PKEY_HEX": "64-char hex private key, no 0x prefix (required)"},
		)
		os.Exit(1)
	}

	// Connect to network (chainID and gasPrice fetched from network)
	net, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network connection failed: %v", err)
	}
	cutils.PrintNetworkInfo(net)

	// Prepare account
	acc, err := cutils.PrepareAccount(net, cutils.MustGetPkeyHex())
	if err != nil {
		cutils.Fatal("Account setup failed: %v", err)
	}
	cutils.PrintAccountInfo(acc)

	// Parse destination and value
	toAddr := common.HexToAddress(os.Args[1])
	value := big.NewInt(0)
	_, success := value.SetString(os.Args[2], 10)
	if !success {
		cutils.Fatal("Invalid value provided: %s", os.Args[2])
	}

	// Get recipient balance
	recipientBalance, err := cutils.GetBalance(net, toAddr)
	if err != nil {
		cutils.Fatal("Error getting recipient balance: %v", err)
	}

	cutils.Section("TRANSFER INFO")
	cutils.PrintKeyValue("To Address", toAddr.String())
	cutils.PrintKeyValueEth("Recipient Current Balance", recipientBalance)
	cutils.PrintKeyValueEth("Transfer Amount", value)
	cutils.PrintKeyValueEth("Recipient Balance After", new(big.Int).Add(recipientBalance, value))

	// Check if account has enough balance
	gasCost := new(big.Int).Mul(net.GasPrice, big.NewInt(int64(cutils.GasLimitSimpleTransfer)))
	totalNeeded := new(big.Int).Add(value, gasCost)

	if acc.Balance.Cmp(totalNeeded) < 0 {
		cutils.Fatal("Insufficient balance. Need %s ETH (including gas), have %s ETH",
			cutils.WeiToEth(totalNeeded), cutils.WeiToEth(acc.Balance))
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("ETH Transfer", value, cutils.GasLimitSimpleTransfer, net.GasPrice)

	tx, err := cutils.SignAndSendTx(net, acc, toAddr, value, cutils.GasLimitSimpleTransfer, nil)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
