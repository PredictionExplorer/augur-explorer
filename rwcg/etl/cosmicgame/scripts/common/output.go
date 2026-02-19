package common

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Verbose controls whether detailed output (network, account, sections, etc.) is printed.
// Default is true so read-only scripts keep current behavior. Tx scripts call ParseInfoFlag()
// at start: with -i they get Verbose=true, without -i Verbose=false (quiet: only success/error line).
var Verbose = true

// ParseInfoFlag must be called at the start of main() by scripts that send transactions.
// It removes "-i" from os.Args if present and sets Verbose accordingly: Verbose=true if -i was
// passed (detailed output), Verbose=false otherwise (only "Success. Tx hash = ..." or error).
func ParseInfoFlag() {
	var newArgs []string
	for _, a := range os.Args {
		if a == "-i" {
			Verbose = true
			continue
		}
		newArgs = append(newArgs, a)
	}
	if len(newArgs) == len(os.Args) {
		// -i was not present; caller is a tx script so default to quiet
		Verbose = false
	}
	os.Args = newArgs
}

// Section prints a section header for organized output (only when Verbose).
func Section(title string) {
	if !Verbose {
		return
	}
	fmt.Printf("\n==================== %s ====================\n", title)
}

// PrintNetworkInfo displays verbose network information (only when Verbose).
func PrintNetworkInfo(net *NetworkInfo) {
	if !Verbose {
		return
	}
	Section("NETWORK INFO")
	fmt.Printf("RPC URL             = %s\n", net.RPCURL)
	fmt.Printf("Chain ID            = %s\n", net.ChainID.String())
	fmt.Printf("Gas Price (wei)     = %s\n", net.GasPrice.String())
	fmt.Printf("Gas Price (gwei)    = %.4f\n", WeiToGwei(net.GasPrice))
	fmt.Printf("Latest Block        = %s\n", net.BlockNum.String())
	fmt.Printf("Block Timestamp     = %d\n", net.BlockTime)
}

// PrintAccountInfo displays verbose account information (only when Verbose).
func PrintAccountInfo(acc *AccountInfo) {
	if !Verbose {
		return
	}
	Section("ACCOUNT INFO")
	fmt.Printf("Address             = %s\n", acc.Address.String())
	fmt.Printf("Nonce               = %d\n", acc.Nonce)
	fmt.Printf("Balance (wei)       = %s\n", acc.Balance.String())
	fmt.Printf("Balance (ETH)       = %s\n", WeiToEth(acc.Balance))
}

// PrintContractInfo displays contract address information (only when Verbose).
func PrintContractInfo(name string, addr common.Address) {
	if !Verbose {
		return
	}
	Section("CONTRACT")
	fmt.Printf("%-20s= %s\n", name, addr.String())
}

// PrintTxSubmitting shows what transaction is about to be sent (only when Verbose).
func PrintTxSubmitting(action string, value *big.Int, gasLimit uint64, gasPrice *big.Int) {
	if !Verbose {
		return
	}
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

// PrintTxResult displays transaction result. When Verbose is false, prints only
// "Success. Tx hash = <hash>" or the error; when Verbose is true, prints full details.
func PrintTxResult(tx *types.Transaction, err error) {
	if !Verbose {
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		if tx == nil {
			fmt.Printf("transaction is nil\n")
			return
		}
		fmt.Printf("Success. Tx hash = %s\n", tx.Hash().String())
		return
	}
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

// PrintCallResult displays a read-only call result (only when Verbose).
func PrintCallResult(name string, value interface{}) {
	if !Verbose {
		return
	}
	fmt.Printf("%-28s= %v\n", name, value)
}

// PrintKeyValue prints a key-value pair with consistent formatting (only when Verbose).
func PrintKeyValue(key string, value interface{}) {
	if !Verbose {
		return
	}
	fmt.Printf("%-28s= %v\n", key, value)
}

// PrintKeyValueEth prints a key-value pair formatted as ETH (only when Verbose).
func PrintKeyValueEth(key string, wei *big.Int) {
	if !Verbose {
		return
	}
	fmt.Printf("%-28s= %s ETH\n", key, WeiToEth(wei))
}

// PrintKeyValueDuration prints a key-value pair formatted as duration (only when Verbose).
func PrintKeyValueDuration(key string, secs int64) {
	if !Verbose {
		return
	}
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
