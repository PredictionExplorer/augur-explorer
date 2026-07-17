package ethtx

// Output reproduces the output format of the legacy transaction scripts:
// quiet mode prints only "Success. Tx hash = ..." or the error, verbose mode
// (-i/--info) prints network, account, and transaction detail sections.

import (
	"context"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Output controls how much detail a transaction subcommand prints and where.
type Output struct {
	// Verbose enables the detailed sections; quiet mode prints only the
	// final result line.
	Verbose bool
	// W receives all output.
	W io.Writer
}

// WeiToEthText formats a wei amount as a decimal ETH string with 18 decimal
// places, matching Ethereum precision.
func WeiToEthText(wei *big.Int) string {
	if wei == nil {
		return "0.000000000000000000"
	}
	ether := new(big.Float).SetInt(wei)
	eth := new(big.Float).Quo(ether, big.NewFloat(1e18))
	return eth.Text('f', 18)
}

// WeiToGwei converts a wei amount to gwei as float64 for display.
func WeiToGwei(wei *big.Int) float64 {
	if wei == nil {
		return 0
	}
	gwei := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e9))
	f, _ := gwei.Float64()
	return f
}

func (o Output) printf(format string, args ...any) {
	_, _ = fmt.Fprintf(o.W, format, args...)
}

// Section prints a section header when verbose output is enabled.
func (o Output) Section(title string) {
	if !o.Verbose {
		return
	}
	o.printf("\n==================== %s ====================\n", title)
}

// KeyValue prints an aligned key-value pair when verbose output is enabled.
func (o Output) KeyValue(key string, value any) {
	if !o.Verbose {
		return
	}
	o.printf("%-28s= %v\n", key, value)
}

// KeyValueEth prints a wei amount formatted as ETH when verbose output is
// enabled.
func (o Output) KeyValueEth(key string, wei *big.Int) {
	if !o.Verbose {
		return
	}
	o.printf("%-28s= %s ETH\n", key, WeiToEthText(wei))
}

// KeyValueDuration prints a seconds value with its human-readable rendering
// when verbose output is enabled.
func (o Output) KeyValueDuration(key string, secs int64) {
	if !o.Verbose {
		return
	}
	o.printf("%-28s= %d (%s)\n", key, secs, FmtDuration(secs))
}

// ContractInfo prints a contract address section when verbose output is
// enabled.
func (o Output) ContractInfo(name string, addr common.Address) {
	if !o.Verbose {
		return
	}
	o.Section("CONTRACT")
	o.printf("%-20s= %s\n", name, addr.String())
}

// NetworkInfo prints the connected-network details when verbose output is
// enabled.
func (o Output) NetworkInfo(net *Network) {
	if !o.Verbose {
		return
	}
	o.Section("NETWORK INFO")
	o.printf("RPC URL             = %s\n", net.RPCURL)
	o.printf("Chain ID            = %s\n", net.ChainID.String())
	o.printf("Gas Price (wei)     = %s\n", net.GasPrice.String())
	o.printf("Gas Price (gwei)    = %.4f\n", WeiToGwei(net.GasPrice))
	o.printf("Latest Block        = %s\n", net.BlockNum.String())
	o.printf("Block Timestamp     = %d\n", net.BlockTime)
}

// AccountInfo prints the signing-account details when verbose output is
// enabled.
func (o Output) AccountInfo(acc *Account) {
	if !o.Verbose {
		return
	}
	o.Section("ACCOUNT INFO")
	o.printf("Address             = %s\n", acc.Address.String())
	o.printf("Nonce               = %d\n", acc.Nonce)
	o.printf("Balance (wei)       = %s\n", acc.Balance.String())
	o.printf("Balance (ETH)       = %s\n", WeiToEthText(acc.Balance))
}

// TxSubmitting announces the transaction about to be sent when verbose
// output is enabled.
func (o Output) TxSubmitting(action string, value *big.Int, gasLimit uint64, gasPrice *big.Int) {
	if !o.Verbose {
		return
	}
	o.Section("SUBMITTING TRANSACTION")
	o.printf("Action              = %s\n", action)
	if value != nil && value.Sign() > 0 {
		o.printf("Value (ETH)         = %s\n", WeiToEthText(value))
	}
	o.printf("Gas Limit           = %d\n", gasLimit)
	o.printf("Gas Price (gwei)    = %.4f\n", WeiToGwei(gasPrice))
	maxCostWei := new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gasLimit))
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(maxCostWei), big.NewFloat(1e18))
	o.printf("Max Gas Cost (ETH)  = %s\n", ethValue.Text('f', 6))
}

// txResultAndWait reports the send outcome, waits for the transaction to be
// mined, and reports on-chain success or revert. It returns true only when
// the transaction was mined with a successful status.
func (o Output) txResultAndWait(ctx context.Context, s *Session, tx *types.Transaction, err error) bool {
	if err != nil {
		if !o.Verbose {
			o.printf("%v\n", err)
		} else {
			o.Section("TRANSACTION RESULT")
			o.printf("Status              = FAILED\n")
			o.printf("Error               = %v\n", err)
		}
		return false
	}
	if tx == nil {
		if !o.Verbose {
			o.printf("transaction is nil\n")
		} else {
			o.Section("TRANSACTION RESULT")
			o.printf("Status              = FAILED\n")
			o.printf("Error               = transaction is nil\n")
		}
		return false
	}
	receipt, waitErr := s.WaitForReceipt(ctx, tx)
	if waitErr != nil {
		if !o.Verbose {
			o.printf("Transaction submitted but receipt not received: %v. Tx hash = %s\n", waitErr, tx.Hash().String())
		} else {
			o.Section("TRANSACTION RESULT")
			o.printf("Status              = SUBMITTED (receipt wait failed)\n")
			o.printf("Tx Hash             = %s\n", tx.Hash().String())
			o.printf("Error               = %v\n", waitErr)
			o.printf("\nCheck block explorer to see if the transaction was mined.\n")
		}
		return false
	}
	if receipt.Status == 0 {
		if !o.Verbose {
			o.printf("Transaction reverted on-chain. Tx hash = %s (check block explorer for revert reason)\n", tx.Hash().String())
		} else {
			o.Section("TRANSACTION RESULT")
			o.printf("Status              = REVERTED\n")
			o.printf("Tx Hash             = %s\n", tx.Hash().String())
			o.printf("Block               = %s\n", receipt.BlockNumber.String())
			o.printf("\nTransaction was mined but reverted. Check the block explorer for the revert reason.\n")
		}
		return false
	}
	if !o.Verbose {
		o.printf("Success. Tx hash = %s\n", tx.Hash().String())
	} else {
		o.Section("TRANSACTION RESULT")
		o.printf("Status              = SUCCESS\n")
		o.printf("Tx Hash             = %s\n", tx.Hash().String())
		o.printf("Block               = %s\n", receipt.BlockNumber.String())
		o.printf("Gas Used            = %d\n", receipt.GasUsed)
	}
	return true
}
