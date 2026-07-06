package main

// txOutput reproduces the output format of the legacy transaction scripts:
// quiet mode prints only "Success. Tx hash = ..." or the error, verbose mode
// (-i/--info) prints network, account, and transaction detail sections.

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// txOutput controls how much detail a transaction subcommand prints.
type txOutput struct {
	verbose bool
}

// section prints a section header when verbose output is enabled.
func (o txOutput) section(title string) {
	if !o.verbose {
		return
	}
	fmt.Printf("\n==================== %s ====================\n", title)
}

// keyValue prints an aligned key-value pair when verbose output is enabled.
func (o txOutput) keyValue(key string, value interface{}) {
	if !o.verbose {
		return
	}
	fmt.Printf("%-28s= %v\n", key, value)
}

// keyValueEth prints a wei amount formatted as ETH when verbose output is enabled.
func (o txOutput) keyValueEth(key string, wei *big.Int) {
	if !o.verbose {
		return
	}
	fmt.Printf("%-28s= %s ETH\n", key, weiToEthText(wei))
}

// networkInfo prints the connected-network details when verbose output is enabled.
func (o txOutput) networkInfo(net *networkInfo) {
	if !o.verbose {
		return
	}
	o.section("NETWORK INFO")
	fmt.Printf("RPC URL             = %s\n", net.rpcURL)
	fmt.Printf("Chain ID            = %s\n", net.chainID.String())
	fmt.Printf("Gas Price (wei)     = %s\n", net.gasPrice.String())
	fmt.Printf("Gas Price (gwei)    = %.4f\n", weiToGwei(net.gasPrice))
	fmt.Printf("Latest Block        = %s\n", net.blockNum.String())
	fmt.Printf("Block Timestamp     = %d\n", net.blockTime)
}

// accountInfo prints the signing-account details when verbose output is enabled.
func (o txOutput) accountInfo(acc *accountInfo) {
	if !o.verbose {
		return
	}
	o.section("ACCOUNT INFO")
	fmt.Printf("Address             = %s\n", acc.address.String())
	fmt.Printf("Nonce               = %d\n", acc.nonce)
	fmt.Printf("Balance (wei)       = %s\n", acc.balance.String())
	fmt.Printf("Balance (ETH)       = %s\n", weiToEthText(acc.balance))
}

// txSubmitting announces the transaction about to be sent when verbose output
// is enabled.
func (o txOutput) txSubmitting(action string, value *big.Int, gasLimit uint64, gasPrice *big.Int) {
	if !o.verbose {
		return
	}
	o.section("SUBMITTING TRANSACTION")
	fmt.Printf("Action              = %s\n", action)
	if value != nil && value.Sign() > 0 {
		fmt.Printf("Value (ETH)         = %s\n", weiToEthText(value))
	}
	fmt.Printf("Gas Limit           = %d\n", gasLimit)
	fmt.Printf("Gas Price (gwei)    = %.4f\n", weiToGwei(gasPrice))
	maxCostWei := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasLimit)))
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(maxCostWei), big.NewFloat(1e18))
	fmt.Printf("Max Gas Cost (ETH)  = %s\n", ethValue.Text('f', 6))
}

// txResultAndWait reports the send outcome, waits for the transaction to be
// mined, and reports on-chain success or revert. It returns true only when
// the transaction was mined with a successful status.
func (o txOutput) txResultAndWait(client *ethclient.Client, tx *types.Transaction, err error) bool {
	if err != nil {
		if !o.verbose {
			fmt.Printf("%v\n", err)
		} else {
			o.section("TRANSACTION RESULT")
			fmt.Printf("Status              = FAILED\n")
			fmt.Printf("Error               = %v\n", err)
		}
		return false
	}
	if tx == nil {
		if !o.verbose {
			fmt.Printf("transaction is nil\n")
		} else {
			o.section("TRANSACTION RESULT")
			fmt.Printf("Status              = FAILED\n")
			fmt.Printf("Error               = transaction is nil\n")
		}
		return false
	}
	receipt, waitErr := waitForReceipt(context.Background(), client, tx)
	if waitErr != nil {
		if !o.verbose {
			fmt.Printf("Transaction submitted but receipt not received: %v. Tx hash = %s\n", waitErr, tx.Hash().String())
		} else {
			o.section("TRANSACTION RESULT")
			fmt.Printf("Status              = SUBMITTED (receipt wait failed)\n")
			fmt.Printf("Tx Hash             = %s\n", tx.Hash().String())
			fmt.Printf("Error               = %v\n", waitErr)
			fmt.Printf("\nCheck block explorer to see if the transaction was mined.\n")
		}
		return false
	}
	if receipt.Status == 0 {
		if !o.verbose {
			fmt.Printf("Transaction reverted on-chain. Tx hash = %s (check block explorer for revert reason)\n", tx.Hash().String())
		} else {
			o.section("TRANSACTION RESULT")
			fmt.Printf("Status              = REVERTED\n")
			fmt.Printf("Tx Hash             = %s\n", tx.Hash().String())
			fmt.Printf("Block               = %s\n", receipt.BlockNumber.String())
			fmt.Printf("\nTransaction was mined but reverted. Check the block explorer for the revert reason.\n")
		}
		return false
	}
	if !o.verbose {
		fmt.Printf("Success. Tx hash = %s\n", tx.Hash().String())
	} else {
		o.section("TRANSACTION RESULT")
		fmt.Printf("Status              = SUCCESS\n")
		fmt.Printf("Tx Hash             = %s\n", tx.Hash().String())
		fmt.Printf("Block               = %s\n", receipt.BlockNumber.String())
		fmt.Printf("Gas Used            = %d\n", receipt.GasUsed)
	}
	return true
}

// finishTx converts the txResultAndWait outcome into an error for cobra RunE.
func (s *txSession) finishTx(tx *types.Transaction, err error) error {
	if !s.out.txResultAndWait(s.net.client, tx, err) {
		return fmt.Errorf("transaction did not succeed")
	}
	return nil
}
