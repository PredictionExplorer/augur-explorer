// Transfers CosmicSignatureToken (CST) to another address
package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	cutils.ParseInfoFlag()
	// Usage check
	if len(os.Args) < 4 {
		cutils.PrintUsage(os.Args[0],
			"[cosmictoken_addr] [destination_addr] [amount]",
			"Transfers CosmicSignatureToken (CST) to another address",
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

	// Parse parameters
	tokenAddr := common.HexToAddress(os.Args[1])
	destAddr := common.HexToAddress(os.Args[2])

	amount := big.NewInt(0)
	_, success := amount.SetString(os.Args[3], 10)
	if !success {
		cutils.Fatal("Invalid amount provided: %s", os.Args[3])
	}

	// Contract setup
	cstToken, err := NewCosmicSignatureToken(tokenAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicSignatureToken: %v", err)
	}

	// Get token info
	copts := cutils.CreateCallOpts()

	symbol, err := cstToken.Symbol(copts)
	if err != nil {
		symbol = "CST"
	}

	senderBalance, err := cstToken.BalanceOf(copts, acc.Address)
	if err != nil {
		cutils.Fatal("Error getting sender balance: %v", err)
	}

	destBalance, err := cstToken.BalanceOf(copts, destAddr)
	if err != nil {
		cutils.Fatal("Error getting destination balance: %v", err)
	}

	cutils.Section("TOKEN INFO")
	cutils.PrintKeyValue("Token Address", tokenAddr.String())
	cutils.PrintKeyValue("Token Symbol", symbol)

	cutils.Section("TRANSFER INFO")
	cutils.PrintKeyValue("From", acc.Address.String())
	cutils.PrintKeyValue("To", destAddr.String())
	cutils.PrintKeyValue("Amount (raw)", amount.String())
	cutils.PrintKeyValue("Amount", cutils.WeiToEth(amount)+" "+symbol)
	cutils.PrintKeyValue("Sender Balance Before", cutils.WeiToEth(senderBalance)+" "+symbol)
	cutils.PrintKeyValue("Dest Balance Before", cutils.WeiToEth(destBalance)+" "+symbol)

	// Check balance
	if senderBalance.Cmp(amount) < 0 {
		cutils.Fatal("Insufficient balance. Need %s %s, have %s %s",
			cutils.WeiToEth(amount), symbol, cutils.WeiToEth(senderBalance), symbol)
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("Transfer", nil, cutils.GasLimitERC20Transfer, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitERC20Transfer)

	tx, err := cstToken.Transfer(txopts, destAddr, amount)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
