package main

// Transaction plumbing shared by the state-changing subcommands: network
// connection, account preparation from PKEY_HEX, transact options with an
// EIP-155 signer, gas policy, and receipt waiting. Ported from the legacy
// script helpers so subcommand output and behavior stay identical.

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
	"github.com/spf13/cobra"
)

// txEnvHelp documents the environment variables shared by all transaction
// subcommands; it is appended to their long help text.
const txEnvHelp = `Environment variables:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`

// addInfoFlag registers the -i/--info flag used by transaction subcommands to
// switch from quiet output (only success or error) to detailed output.
func addInfoFlag(c *cobra.Command, verbose *bool) {
	c.Flags().BoolVarP(verbose, "info", "i", false, "print detailed network, account and transaction information")
}

// Gas limits by operation type (same values the legacy scripts used).
const (
	// gasLimitERC721Approve covers setApprovalForAll calls.
	gasLimitERC721Approve = uint64(100000)
	// gasLimitContractCall covers ordinary contract calls (transfer, set name).
	gasLimitContractCall = uint64(300000)
	// gasLimitHighComplexity covers minting and marketplace operations.
	gasLimitHighComplexity = uint64(5000000)
)

// receiptWaitTimeout bounds how long finishTx waits for a transaction receipt.
const receiptWaitTimeout = 2 * time.Minute

// networkInfo holds chain and RPC information fetched from the network.
type networkInfo struct {
	client    *ethclient.Client
	chainID   *big.Int
	gasPrice  *big.Int
	blockNum  *big.Int
	blockTime uint64
	rpcURL    string
}

// accountInfo holds the signing account state fetched from the network.
type accountInfo struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
	nonce      uint64
	balance    *big.Int
}

// connectToRPC connects to the endpoint in RPC_URL and fetches the chain ID,
// suggested gas price and latest block. The chain ID is always taken from the
// network, never hardcoded.
func connectToRPC() (*networkInfo, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		return nil, fmt.Errorf("RPC_URL environment variable not set")
	}
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to ETH RPC at %s: %w", rpcURL, err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting chain ID from network: %w", err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting suggested gas price: %w", err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error getting latest block: %w", err)
	}
	return &networkInfo{
		client:    client,
		chainID:   chainID,
		gasPrice:  gasPrice,
		blockNum:  header.Number,
		blockTime: header.Time,
		rpcURL:    rpcURL,
	}, nil
}

// prepareAccount parses the private key and fetches the account nonce and
// balance from the network.
func prepareAccount(net *networkInfo, pkeyHex string) (*accountInfo, error) {
	privateKey, err := crypto.HexToECDSA(pkeyHex)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %w", err)
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("couldn't derive public key from private key")
	}
	address := crypto.PubkeyToAddress(*publicKey)
	nonce, err := net.client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, fmt.Errorf("error getting account nonce: %w", err)
	}
	balance, err := net.client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting account balance: %w", err)
	}
	return &accountInfo{
		privateKey: privateKey,
		address:    address,
		nonce:      nonce,
		balance:    balance,
	}, nil
}

// adjustGasPrice doubles the node-suggested gas price for faster inclusion,
// the same fixed 2.0 multiplier every legacy RandomWalk script applied.
func adjustGasPrice(basePrice *big.Int) *big.Int {
	if basePrice == nil {
		return big.NewInt(0)
	}
	adjusted := new(big.Float).Mul(new(big.Float).SetInt(basePrice), big.NewFloat(2.0))
	result := new(big.Int)
	adjusted.Int(result)
	return result
}

// transactOpts builds transaction options signing with an EIP-155 signer for
// the network's chain ID.
func transactOpts(net *networkInfo, acc *accountInfo, value *big.Int, gasLimit uint64) *bind.TransactOpts {
	txopts := bind.NewKeyedTransactor(acc.privateKey)
	txopts.Nonce = big.NewInt(int64(acc.nonce))
	txopts.Value = value
	txopts.GasLimit = gasLimit
	txopts.GasPrice = adjustGasPrice(net.gasPrice)
	txopts.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(net.chainID)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), acc.privateKey)
		if err != nil {
			return nil, fmt.Errorf("error signing transaction: %w", err)
		}
		return tx.WithSignature(signer, signature)
	}
	return txopts
}

// callOpts returns options for read-only contract calls.
func callOpts() *bind.CallOpts {
	return &bind.CallOpts{}
}

// waitForReceipt waits (bounded by receiptWaitTimeout) until the transaction
// is mined and returns its receipt.
func waitForReceipt(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
	ctx2, cancel := context.WithTimeout(ctx, receiptWaitTimeout)
	defer cancel()
	return bind.WaitMined(ctx2, client, tx)
}

// txSession bundles the connected network, the signing account and the output
// policy used by a transaction subcommand.
type txSession struct {
	net *networkInfo
	acc *accountInfo
	out txOutput
}

// newTxSession connects to the RPC endpoint (RPC_URL), loads the signer from
// PKEY_HEX and prints network/account details when verbose is enabled.
func newTxSession(verbose bool) (*txSession, error) {
	out := txOutput{verbose: verbose}
	net, err := connectToRPC()
	if err != nil {
		return nil, fmt.Errorf("network connection failed: %w", err)
	}
	out.networkInfo(net)
	pkey, err := pkeyHexFromEnv()
	if err != nil {
		return nil, err
	}
	acc, err := prepareAccount(net, pkey)
	if err != nil {
		return nil, fmt.Errorf("account setup failed: %w", err)
	}
	out.accountInfo(acc)
	return &txSession{net: net, acc: acc, out: out}, nil
}
