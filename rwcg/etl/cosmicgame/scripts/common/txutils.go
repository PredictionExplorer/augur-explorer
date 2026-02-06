// Package common provides shared utilities for CosmicGame development scripts.
// It standardizes network connection, account handling, transaction creation,
// and verbose output across all scripts.
package common

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NetworkInfo holds chain and RPC information fetched from the network
type NetworkInfo struct {
	Client    *ethclient.Client
	ChainID   *big.Int
	GasPrice  *big.Int
	BlockNum  *big.Int
	BlockTime uint64
	RPCURL    string
}

// AccountInfo holds account-related transaction info
type AccountInfo struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
	Nonce      uint64
	Balance    *big.Int
}

// ConnectToRPC establishes connection and fetches network info.
// Reads RPC_URL from environment variable.
// Returns NetworkInfo with chainID, gasPrice, and latest block info.
func ConnectToRPC() (*NetworkInfo, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		return nil, fmt.Errorf("RPC_URL environment variable not set")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to ETH RPC at %s: %v", rpcURL, err)
	}

	// Always fetch chainID from network (never hardcode)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting chain ID from network: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting suggested gas price: %v", err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error getting latest block: %v", err)
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

// PrepareAccount parses private key and fetches account info from network.
// The pkeyHex must be a 64-character hex string (without 0x prefix).
func PrepareAccount(net *NetworkInfo, pkeyHex string) (*AccountInfo, error) {
	if len(pkeyHex) != 64 {
		return nil, fmt.Errorf("private key must be 64 hex characters (got %d)", len(pkeyHex))
	}

	privateKey, err := crypto.HexToECDSA(pkeyHex)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("couldn't derive public key from private key")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := net.Client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, fmt.Errorf("error getting account nonce: %v", err)
	}

	balance, err := net.Client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting account balance: %v", err)
	}

	return &AccountInfo{
		PrivateKey: privateKey,
		Address:    address,
		Nonce:      nonce,
		Balance:    balance,
	}, nil
}

// CreateTransactOpts builds transaction options with proper EIP-155 signer.
// The chainID is always read from the network (via NetworkInfo).
func CreateTransactOpts(net *NetworkInfo, acc *AccountInfo, value *big.Int, gasLimit uint64) *bind.TransactOpts {
	txopts := bind.NewKeyedTransactor(acc.PrivateKey)
	txopts.Nonce = big.NewInt(int64(acc.Nonce))
	txopts.Value = value
	txopts.GasLimit = gasLimit
	txopts.GasPrice = net.GasPrice

	txopts.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(net.ChainID)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), acc.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("error signing transaction: %v", err)
		}
		return tx.WithSignature(signer, signature)
	}

	return txopts
}

// CreateCallOpts returns a CallOpts for read-only contract calls
func CreateCallOpts() *bind.CallOpts {
	return &bind.CallOpts{}
}

// SignAndSendTx signs a raw transaction and sends it to the network.
// Useful for simple ETH transfers.
func SignAndSendTx(net *NetworkInfo, acc *AccountInfo, to common.Address, value *big.Int, gasLimit uint64, data []byte) (*types.Transaction, error) {
	tx := types.NewTransaction(acc.Nonce, to, value, gasLimit, net.GasPrice, data)

	signer := types.NewEIP155Signer(net.ChainID)
	signedTx, err := types.SignTx(tx, signer, acc.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error signing transaction: %v", err)
	}

	err = net.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return signedTx, fmt.Errorf("error sending transaction: %v", err)
	}

	return signedTx, nil
}

// GetBalance fetches the ETH balance of an address
func GetBalance(net *NetworkInfo, addr common.Address) (*big.Int, error) {
	return net.Client.BalanceAt(context.Background(), addr, nil)
}
