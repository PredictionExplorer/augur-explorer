package ethtx

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Printer emits the standardized sectioned output shared by all cgctl
// subcommands. When Verbose is false only the final result line is printed
// (transaction commands default to quiet; pass -i/--info for full output).
// Read-only commands construct a verbose Printer so state dumps always show.
type Printer struct {
	// Verbose controls whether detailed output (network, account, sections,
	// key/value pairs) is printed.
	Verbose bool
}

// NewPrinter returns a Printer with the given verbosity.
func NewPrinter(verbose bool) *Printer {
	return &Printer{Verbose: verbose}
}

// Section prints a section header for organized output (only when Verbose).
func (p *Printer) Section(title string) {
	if !p.Verbose {
		return
	}
	fmt.Printf("\n==================== %s ====================\n", title)
}

// NetworkInfo displays network information (only when Verbose).
func (p *Printer) NetworkInfo(n *NetworkInfo) {
	if !p.Verbose {
		return
	}
	p.Section("NETWORK INFO")
	fmt.Printf("RPC URL             = %s\n", n.RPCURL)
	fmt.Printf("Chain ID            = %s\n", n.ChainID.String())
	fmt.Printf("Gas Price (wei)     = %s\n", n.GasPrice.String())
	fmt.Printf("Gas Price (gwei)    = %.4f\n", WeiToGwei(n.GasPrice))
	fmt.Printf("Latest Block        = %s\n", n.BlockNum.String())
	fmt.Printf("Block Timestamp     = %d\n", n.BlockTime)
}

// AccountInfo displays account information (only when Verbose).
func (p *Printer) AccountInfo(acc *AccountInfo) {
	if !p.Verbose {
		return
	}
	p.Section("ACCOUNT INFO")
	fmt.Printf("Address             = %s\n", acc.Address.String())
	fmt.Printf("Nonce               = %d\n", acc.Nonce)
	fmt.Printf("Balance (wei)       = %s\n", acc.Balance.String())
	fmt.Printf("Balance (ETH)       = %s\n", WeiToEth(acc.Balance))
}

// ContractInfo displays a contract address (only when Verbose).
func (p *Printer) ContractInfo(name string, addr common.Address) {
	if !p.Verbose {
		return
	}
	p.Section("CONTRACT")
	fmt.Printf("%-20s= %s\n", name, addr.String())
}

// TxSubmitting shows what transaction is about to be sent (only when Verbose).
func (p *Printer) TxSubmitting(action string, value *big.Int, gasLimit uint64, gasPrice *big.Int) {
	if !p.Verbose {
		return
	}
	p.Section("SUBMITTING TRANSACTION")
	fmt.Printf("Action              = %s\n", action)
	if value != nil && value.Sign() > 0 {
		fmt.Printf("Value (ETH)         = %s\n", WeiToEth(value))
	}
	fmt.Printf("Gas Limit           = %d\n", gasLimit)
	fmt.Printf("Gas Price (gwei)    = %.4f\n", WeiToGwei(gasPrice))
	maxCostWei := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))
	fmt.Printf("Max Gas Cost (ETH)  = %s\n", WeiToEthCompact(maxCostWei))
}

// TxSubmitted reports a successfully submitted transaction. It does not wait
// for the transaction to be mined; use TxMined to confirm on-chain success.
func (p *Printer) TxSubmitted(tx *types.Transaction) {
	if !p.Verbose {
		fmt.Printf("Success. Tx hash = %s\n", tx.Hash().String())
		return
	}
	p.Section("TRANSACTION RESULT")
	fmt.Printf("Status              = SUBMITTED\n")
	fmt.Printf("Tx Hash             = %s\n", tx.Hash().String())
	fmt.Printf("Gas Limit           = %d\n", tx.Gas())
	fmt.Printf("Gas Price (gwei)    = %.4f\n", WeiToGwei(tx.GasPrice()))
	if tx.Value() != nil && tx.Value().Sign() > 0 {
		fmt.Printf("Value (ETH)         = %s\n", WeiToEth(tx.Value()))
	}
	fmt.Printf("\nNote: Transaction submitted. Use a block explorer to verify confirmation.\n")
}

// TxMined waits for the transaction receipt and reports success only if the
// transaction was mined and did not revert. It returns an error when the
// receipt cannot be fetched or the transaction reverted on-chain.
func (p *Printer) TxMined(ctx context.Context, client *ethclient.Client, tx *types.Transaction) error {
	receipt, err := WaitForReceipt(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("transaction submitted but receipt not received (tx hash = %s): %w", tx.Hash().String(), err)
	}
	if receipt.Status == 0 {
		return fmt.Errorf("transaction reverted on-chain, tx hash = %s (check block explorer for revert reason)", tx.Hash().String())
	}
	if !p.Verbose {
		fmt.Printf("Success. Tx hash = %s\n", tx.Hash().String())
		return nil
	}
	p.Section("TRANSACTION RESULT")
	fmt.Printf("Status              = SUCCESS\n")
	fmt.Printf("Tx Hash             = %s\n", tx.Hash().String())
	fmt.Printf("Block               = %s\n", receipt.BlockNumber.String())
	fmt.Printf("Gas Used            = %d\n", receipt.GasUsed)
	return nil
}

// KeyValue prints a key/value pair with consistent formatting (only when Verbose).
func (p *Printer) KeyValue(key string, value interface{}) {
	if !p.Verbose {
		return
	}
	fmt.Printf("%-28s= %v\n", key, value)
}

// KeyValueEth prints a key/value pair formatted as ETH (only when Verbose).
func (p *Printer) KeyValueEth(key string, wei *big.Int) {
	if !p.Verbose {
		return
	}
	fmt.Printf("%-28s= %s ETH\n", key, WeiToEth(wei))
}

// KeyValueDuration prints a key/value pair formatted as a duration (only when Verbose).
func (p *Printer) KeyValueDuration(key string, secs int64) {
	if !p.Verbose {
		return
	}
	fmt.Printf("%-28s= %d (%s)\n", key, secs, FmtDuration(secs))
}
