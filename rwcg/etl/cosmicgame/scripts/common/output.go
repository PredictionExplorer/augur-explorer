package common

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Section prints a section header for organized output
func Section(title string) {
	fmt.Printf("\n==================== %s ====================\n", title)
}

// PrintNetworkInfo displays verbose network information
func PrintNetworkInfo(net *NetworkInfo) {
	Section("NETWORK INFO")
	fmt.Printf("RPC URL             = %s\n", net.RPCURL)
	fmt.Printf("Chain ID            = %s\n", net.ChainID.String())
	fmt.Printf("Gas Price (wei)     = %s\n", net.GasPrice.String())
	fmt.Printf("Gas Price (gwei)    = %.4f\n", WeiToGwei(net.GasPrice))
	fmt.Printf("Latest Block        = %s\n", net.BlockNum.String())
	fmt.Printf("Block Timestamp     = %d\n", net.BlockTime)
}

// PrintAccountInfo displays verbose account information
func PrintAccountInfo(acc *AccountInfo) {
	Section("ACCOUNT INFO")
	fmt.Printf("Address             = %s\n", acc.Address.String())
	fmt.Printf("Nonce               = %d\n", acc.Nonce)
	fmt.Printf("Balance (wei)       = %s\n", acc.Balance.String())
	fmt.Printf("Balance (ETH)       = %s\n", WeiToEth(acc.Balance))
}

// PrintContractInfo displays contract address information
func PrintContractInfo(name string, addr common.Address) {
	Section("CONTRACT")
	fmt.Printf("%-20s= %s\n", name, addr.String())
}

// PrintTxSubmitting shows what transaction is about to be sent
func PrintTxSubmitting(action string, value *big.Int, gasLimit uint64, gasPrice *big.Int) {
	Section("SUBMITTING TRANSACTION")
	fmt.Printf("Action              = %s\n", action)
	if value != nil && value.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("Value (ETH)         = %s\n", WeiToEth(value))
	}
	fmt.Printf("Gas Limit           = %d\n", gasLimit)
	fmt.Printf("Gas Price (gwei)    = %.4f\n", WeiToGwei(gasPrice))
	maxCostWei := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))
	fmt.Printf("Max Gas Cost (ETH)  = %s\n", WeiToEthCompact(maxCostWei))
}

// PrintTxResult displays transaction result
func PrintTxResult(tx *types.Transaction, err error) {
	Section("TRANSACTION RESULT")
	if err != nil {
		fmt.Printf("Status              = FAILED\n")
		fmt.Printf("Error               = %v\n", err)
		return
	}
	if tx == nil {
		fmt.Printf("Status              = FAILED\n")
		fmt.Printf("Error               = transaction is nil\n")
		return
	}
	fmt.Printf("Status              = SUBMITTED\n")
	fmt.Printf("Tx Hash             = %s\n", tx.Hash().String())
	fmt.Printf("Gas Limit           = %d\n", tx.Gas())
	fmt.Printf("Gas Price (gwei)    = %.4f\n", WeiToGwei(tx.GasPrice()))
	if tx.Value() != nil && tx.Value().Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("Value (ETH)         = %s\n", WeiToEth(tx.Value()))
	}
	fmt.Printf("\nNote: Transaction submitted. Use a block explorer to verify confirmation.\n")
}

// PrintCallResult displays a read-only call result
func PrintCallResult(name string, value interface{}) {
	fmt.Printf("%-28s= %v\n", name, value)
}

// PrintKeyValue prints a key-value pair with consistent formatting
func PrintKeyValue(key string, value interface{}) {
	fmt.Printf("%-28s= %v\n", key, value)
}

// PrintKeyValueEth prints a key-value pair formatted as ETH
func PrintKeyValueEth(key string, wei *big.Int) {
	fmt.Printf("%-28s= %s ETH\n", key, WeiToEth(wei))
}

// PrintKeyValueDuration prints a key-value pair formatted as duration
func PrintKeyValueDuration(key string, secs int64) {
	fmt.Printf("%-28s= %d (%s)\n", key, secs, FmtDuration(secs))
}

// Fatal prints error message and exits with code 1
func Fatal(format string, args ...interface{}) {
	fmt.Printf("\nERROR: "+format+"\n", args...)
	os.Exit(1)
}

// FatalIf exits with error if err is not nil
func FatalIf(err error, format string, args ...interface{}) {
	if err != nil {
		fullArgs := append(args, err)
		fmt.Printf("\nERROR: "+format+": %v\n", fullArgs...)
		os.Exit(1)
	}
}

// PrintUsage prints a standardized usage message
func PrintUsage(programName string, args string, description string, envVars map[string]string) {
	fmt.Printf("Usage: %s %s\n\n", programName, args)
	fmt.Printf("  %s\n\n", description)
	if len(envVars) > 0 {
		fmt.Printf("Environment Variables:\n")
		for name, desc := range envVars {
			fmt.Printf("  %-12s %s\n", name, desc)
		}
		fmt.Println()
	}
}
