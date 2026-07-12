// Package ethtx is the transaction plumbing shared by the operator CLIs
// (rwctl, and cgctl once its local copy merges): network connection, account
// preparation from a hex private key, transact options with an EIP-155
// signer, the legacy 2.0x gas-price policy, receipt waiting, and the
// quiet/verbose output format the legacy scripts printed.
//
// Nothing in this package reads environment variables or writes to
// os.Stdout directly: endpoints, keys and the output writer are injected
// through Options so the whole surface is testable against
// internal/testchain.
package ethtx

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Gas limits by operation type (same values the legacy scripts used).
const (
	// GasLimitApprove covers ERC-721 setApprovalForAll and ERC-20 approve calls.
	GasLimitApprove = uint64(100000)
	// GasLimitContractCall covers ordinary contract calls (transfer, set name).
	GasLimitContractCall = uint64(300000)
	// GasLimitHighComplexity covers minting and marketplace operations.
	GasLimitHighComplexity = uint64(5000000)
	// GasLimitBid covers CosmicGame bidding operations.
	GasLimitBid = uint64(500000)
	// GasLimitClaimPrize covers CosmicGame prize claiming (complex operation;
	// V2 needs ~3M).
	GasLimitClaimPrize = uint64(3500000)
	// GasLimitDonate covers CosmicGame donation operations.
	GasLimitDonate = uint64(300000)
	// GasLimitAdminCall covers owner-only setter operations.
	GasLimitAdminCall = uint64(100000)
)

// DefaultGasPriceMultiplier is the factor applied to the node-suggested gas
// price when Options.GasPriceMultiplier is unset — the fixed 2.0 policy every
// legacy script used.
const DefaultGasPriceMultiplier = 2.0

// DefaultReceiptTimeout bounds how long FinishTx waits for a transaction
// receipt unless Options.ReceiptTimeout overrides it.
const DefaultReceiptTimeout = 2 * time.Minute

// Options configures a Session. RPCURL and PrivateKeyHex are required.
type Options struct {
	// RPCURL is the Ethereum JSON-RPC endpoint.
	RPCURL string
	// PrivateKeyHex is the signer key: 64 hex characters, no 0x prefix.
	PrivateKeyHex string
	// Verbose switches from quiet output (only success or error) to the
	// detailed network/account/transaction sections.
	Verbose bool
	// Out receives all human-readable output; defaults to os.Stdout.
	Out io.Writer
	// ReceiptTimeout bounds the receipt wait; defaults to
	// DefaultReceiptTimeout.
	ReceiptTimeout time.Duration
	// GasPriceMultiplier scales the node-suggested gas price on every
	// transaction the session builds. Zero or negative applies
	// DefaultGasPriceMultiplier.
	GasPriceMultiplier float64
}

// Network holds chain and RPC information fetched from the network.
type Network struct {
	Client    *ethclient.Client
	ChainID   *big.Int
	GasPrice  *big.Int
	BlockNum  *big.Int
	BlockTime uint64
	RPCURL    string
}

// Account holds the signing account state fetched from the network.
type Account struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
	Nonce      uint64
	Balance    *big.Int
}

// Session bundles the connected network, the signing account and the output
// policy used by a transaction subcommand.
type Session struct {
	Net *Network
	Acc *Account
	Out Output

	receiptTimeout time.Duration
	gasMultiplier  float64
}

// New connects to the RPC endpoint, loads the signer key and prints
// network/account details when verbose output is enabled.
func New(ctx context.Context, opts Options) (*Session, error) {
	out := Output{Verbose: opts.Verbose, W: opts.Out}
	if out.W == nil {
		out.W = os.Stdout
	}
	net, err := Connect(ctx, opts.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)
	acc, err := PrepareAccount(ctx, net, opts.PrivateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("account setup failed: %w", err)
	}
	out.AccountInfo(acc)
	timeout := opts.ReceiptTimeout
	if timeout <= 0 {
		timeout = DefaultReceiptTimeout
	}
	multiplier := opts.GasPriceMultiplier
	if multiplier <= 0 {
		multiplier = DefaultGasPriceMultiplier
	}
	return &Session{
		Net:            net,
		Acc:            acc,
		Out:            out,
		receiptTimeout: timeout,
		gasMultiplier:  multiplier,
	}, nil
}

// Connect dials the endpoint and fetches the chain ID, suggested gas price
// and latest block. The chain ID is always taken from the network, never
// hardcoded.
func Connect(ctx context.Context, rpcURL string) (*Network, error) {
	if rpcURL == "" {
		return nil, errors.New("RPC endpoint not set")
	}
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to ETH RPC at %s: %w", rpcURL, err)
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting chain ID from network: %w", err)
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting suggested gas price: %w", err)
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting latest block: %w", err)
	}
	return &Network{
		Client:    client,
		ChainID:   chainID,
		GasPrice:  gasPrice,
		BlockNum:  header.Number,
		BlockTime: header.Time,
		RPCURL:    rpcURL,
	}, nil
}

// PrepareAccount parses the private key and fetches the account nonce and
// balance from the network.
func PrepareAccount(ctx context.Context, net *Network, pkeyHex string) (*Account, error) {
	privateKey, err := crypto.HexToECDSA(pkeyHex)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %w", err)
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("couldn't derive public key from private key")
	}
	address := crypto.PubkeyToAddress(*publicKey)
	nonce, err := net.Client.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("error getting account nonce: %w", err)
	}
	balance, err := net.Client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting account balance: %w", err)
	}
	return &Account{
		PrivateKey: privateKey,
		Address:    address,
		Nonce:      nonce,
		Balance:    balance,
	}, nil
}

// AdjustGasPrice scales the node-suggested gas price by the fixed legacy
// default of 2.0 for faster inclusion. Sessions with a configured multiplier
// use AdjustedGasPrice instead.
func AdjustGasPrice(basePrice *big.Int) *big.Int {
	return AdjustGasPriceBy(basePrice, DefaultGasPriceMultiplier)
}

// AdjustGasPriceBy scales the node-suggested gas price by an arbitrary
// positive multiplier. Multipliers of exactly 1.0 return the base price
// untouched; nil base prices yield zero.
func AdjustGasPriceBy(basePrice *big.Int, multiplier float64) *big.Int {
	if basePrice == nil {
		return big.NewInt(0)
	}
	if multiplier == 1.0 {
		return new(big.Int).Set(basePrice)
	}
	adjusted := new(big.Float).Mul(new(big.Float).SetInt(basePrice), big.NewFloat(multiplier))
	result := new(big.Int)
	adjusted.Int(result)
	return result
}

// AdjustedGasPrice returns the network's suggested gas price scaled by the
// session's gas-price multiplier — the price TransactOpts attaches to
// transactions.
func (s *Session) AdjustedGasPrice() *big.Int {
	return AdjustGasPriceBy(s.Net.GasPrice, s.gasMultiplier)
}

// CallOpts returns options for read-only contract calls.
func CallOpts() *bind.CallOpts {
	return &bind.CallOpts{}
}

// TransactOpts builds transaction options signing with an EIP-155 signer for
// the session network's chain ID.
func (s *Session) TransactOpts(value *big.Int, gasLimit uint64) *bind.TransactOpts {
	chainID := s.Net.ChainID
	key := s.Acc.PrivateKey
	return &bind.TransactOpts{
		From:     s.Acc.Address,
		Nonce:    big.NewInt(int64(s.Acc.Nonce)),
		Value:    value,
		GasLimit: gasLimit,
		GasPrice: s.AdjustedGasPrice(),
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.NewEIP155Signer(chainID)
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
			if err != nil {
				return nil, fmt.Errorf("error signing transaction: %w", err)
			}
			return tx.WithSignature(signer, signature)
		},
	}
}

// Balance fetches the current ETH balance of an address.
func (n *Network) Balance(ctx context.Context, addr common.Address) (*big.Int, error) {
	return n.Client.BalanceAt(ctx, addr, nil)
}

// Refresh re-reads the network state (gas price, latest block) and the
// account state (nonce, balance) over the existing connection. Multi-
// transaction commands call it between transactions so the next
// TransactOpts carries a fresh nonce and gas price.
func (s *Session) Refresh(ctx context.Context) error {
	gasPrice, err := s.Net.Client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("error refreshing gas price: %w", err)
	}
	header, err := s.Net.Client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("error refreshing latest block: %w", err)
	}
	nonce, err := s.Net.Client.PendingNonceAt(ctx, s.Acc.Address)
	if err != nil {
		return fmt.Errorf("error refreshing account nonce: %w", err)
	}
	balance, err := s.Net.Client.BalanceAt(ctx, s.Acc.Address, nil)
	if err != nil {
		return fmt.Errorf("error refreshing account balance: %w", err)
	}
	s.Net.GasPrice = gasPrice
	s.Net.BlockNum = header.Number
	s.Net.BlockTime = header.Time
	s.Acc.Nonce = nonce
	s.Acc.Balance = balance
	return nil
}

// WaitForReceipt waits (bounded by the session's receipt timeout) until the
// transaction is mined and returns its receipt.
func (s *Session) WaitForReceipt(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	ctx2, cancel := context.WithTimeout(ctx, s.receiptTimeout)
	defer cancel()
	return bind.WaitMined(ctx2, s.Net.Client, tx)
}

// FinishTx reports the send outcome, waits for the transaction to be mined,
// and converts the result into an error suitable for a CLI RunE: nil only
// when the transaction was mined with a successful status.
func (s *Session) FinishTx(ctx context.Context, tx *types.Transaction, err error) error {
	if !s.Out.txResultAndWait(ctx, s, tx, err) {
		return errors.New("transaction did not succeed")
	}
	return nil
}
