// Package ethtx provides the shared plumbing for cgctl subcommands that talk
// to an Ethereum JSON-RPC endpoint: connecting to the network, deriving the
// signing account from PKEY_HEX, building EIP-155 transaction options with the
// configured gas policy, and printing standardized human-readable output.
//
// Configuration comes from environment variables:
//
//	RPC_URL               JSON-RPC endpoint (required by Connect)
//	PKEY_HEX              64-char hex private key, no 0x prefix (transaction commands)
//	GAS_PRICE_MULTIPLIER  multiplier applied to the suggested gas price (default 2.0)
//
// Chain ID and gas price are always fetched from the network, never hardcoded.
package ethtx

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NetworkInfo holds the RPC client together with chain data fetched from the
// network at connection time.
type NetworkInfo struct {
	Client    *ethclient.Client
	ChainID   *big.Int
	GasPrice  *big.Int
	BlockNum  *big.Int
	BlockTime uint64
	RPCURL    string
}

// AccountInfo holds the signing key and the account state (nonce, balance)
// fetched from the network.
type AccountInfo struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
	Nonce      uint64
	Balance    *big.Int
}

// Connect dials the endpoint in the RPC_URL environment variable and fetches
// chain ID, suggested gas price, and the latest block header.
func Connect(ctx context.Context) (*NetworkInfo, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		return nil, fmt.Errorf("RPC_URL environment variable not set")
	}

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to ETH RPC at %s: %w", rpcURL, err)
	}

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting chain ID from network: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting suggested gas price: %w", err)
	}

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("getting latest block: %w", err)
	}

	return &NetworkInfo{
		Client:    client,
		ChainID:   chainID,
		GasPrice:  gasPrice,
		BlockNum:  header.Number,
		BlockTime: header.Time,
		RPCURL:    rpcURL,
	}, nil
}

// PrivateKeyHexFromEnv returns the value of the PKEY_HEX environment variable
// after validating that it is exactly 64 hex characters (no 0x prefix).
func PrivateKeyHexFromEnv() (string, error) {
	s := os.Getenv("PKEY_HEX")
	if s == "" {
		return "", fmt.Errorf("PKEY_HEX environment variable is required (64 hex characters, no 0x prefix)")
	}
	if len(s) != 64 {
		return "", fmt.Errorf("PKEY_HEX must be 64 hex characters (got %d)", len(s))
	}
	return s, nil
}

// PrepareAccount parses the private key and fetches the account's pending
// nonce and balance from the network.
func (n *NetworkInfo) PrepareAccount(ctx context.Context, pkeyHex string) (*AccountInfo, error) {
	if len(pkeyHex) != 64 {
		return nil, fmt.Errorf("private key must be 64 hex characters (got %d)", len(pkeyHex))
	}

	privateKey, err := crypto.HexToECDSA(pkeyHex)
	if err != nil {
		return nil, fmt.Errorf("parsing private key: %w", err)
	}

	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("couldn't derive public key from private key")
	}
	address := crypto.PubkeyToAddress(*publicKey)

	nonce, err := n.Client.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("getting account nonce: %w", err)
	}

	balance, err := n.Client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("getting account balance: %w", err)
	}

	return &AccountInfo{
		PrivateKey: privateKey,
		Address:    address,
		Nonce:      nonce,
		Balance:    balance,
	}, nil
}

// TransactOpts builds transaction options with an EIP-155 signer. The chain ID
// comes from the network and the gas price is the suggested price adjusted by
// the GAS_PRICE_MULTIPLIER policy.
func (n *NetworkInfo) TransactOpts(acc *AccountInfo, value *big.Int, gasLimit uint64) *bind.TransactOpts {
	txopts := bind.NewKeyedTransactor(acc.PrivateKey)
	txopts.Nonce = big.NewInt(int64(acc.Nonce))
	txopts.Value = value
	txopts.GasLimit = gasLimit
	txopts.GasPrice = AdjustGasPrice(n.GasPrice)

	txopts.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(n.ChainID)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), acc.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("signing transaction: %w", err)
		}
		return tx.WithSignature(signer, signature)
	}

	return txopts
}

// CallOpts returns options for read-only contract calls.
func CallOpts() *bind.CallOpts {
	return &bind.CallOpts{}
}

// Balance fetches the ETH balance of an address.
func (n *NetworkInfo) Balance(ctx context.Context, addr common.Address) (*big.Int, error) {
	return n.Client.BalanceAt(ctx, addr, nil)
}

// AdvanceHardhatTime advances block time on Hardhat/Ganache development nodes
// (evm_increaseTime + evm_mine) and refreshes the cached block number and
// timestamp on success.
func (n *NetworkInfo) AdvanceHardhatTime(ctx context.Context, seconds int64) error {
	if seconds <= 0 {
		return nil
	}
	rpc := n.Client.Client()
	var result interface{}
	if err := rpc.CallContext(ctx, &result, "evm_increaseTime", seconds); err != nil {
		return fmt.Errorf("evm_increaseTime(%d): %w", seconds, err)
	}
	if err := rpc.CallContext(ctx, &result, "evm_mine"); err != nil {
		return fmt.Errorf("evm_mine: %w", err)
	}
	header, err := n.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("refresh block header: %w", err)
	}
	n.BlockNum = header.Number
	n.BlockTime = header.Time
	return nil
}

// DefaultReceiptWaitTimeout is how long WaitForReceipt waits for a transaction
// to be mined before giving up.
const DefaultReceiptWaitTimeout = 2 * time.Minute

// WaitForReceipt blocks until the transaction is mined and returns its receipt
// (receipt.Status is 0 for reverted, 1 for success). It gives up after
// DefaultReceiptWaitTimeout.
func WaitForReceipt(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	ctx2, cancel := context.WithTimeout(ctx, DefaultReceiptWaitTimeout)
	defer cancel()
	return bind.WaitMined(ctx2, client, tx)
}
