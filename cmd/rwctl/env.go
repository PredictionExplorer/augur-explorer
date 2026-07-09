package main

// Shared configuration helpers used by multiple rwctl subcommands:
// environment-variable access, RPC and database connection helpers,
// well-known contract addresses, event topics, and argument parsing.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// Well-known RandomWalk contract addresses on Arbitrum One, kept for operator
// reference (subcommands take contract addresses as arguments).
const (
	// rwalkContractAddr is the RandomWalk NFT contract.
	rwalkContractAddr = "0x895a6F444BE4ba9d124F61DF736605792B35D66b"
	// rwMarketContractAddr is the RandomWalk marketplace contract.
	rwMarketContractAddr = "0x47eF85Dfb775aCE0934fBa9EEd09D22e6eC0Cc08"
)

// Event signature topics used by the scan/verify subcommands.
var (
	// mintEventTopic is keccak256 of the RandomWalk MintEvent signature.
	mintEventTopic = common.HexToHash("0xad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec")
	// transferEventTopic is keccak256 of the ERC-721 Transfer event signature.
	transferEventTopic = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
)

// dialEthClient connects to the Ethereum JSON-RPC endpoint given by the
// RPC_URL environment variable.
func dialEthClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		return nil, fmt.Errorf("can't connect to ETH RPC: %w", err)
	}
	return client, nil
}

// pkeyHexFromEnv returns the signer private key from the PKEY_HEX environment
// variable (64 hex characters, no 0x prefix).
func pkeyHexFromEnv() (string, error) {
	s := os.Getenv("PKEY_HEX")
	if s == "" {
		return "", errors.New("PKEY_HEX environment variable is required (64 hex characters, no 0x prefix)")
	}
	if len(s) != 64 {
		return "", fmt.Errorf("PKEY_HEX must be 64 hex characters (got %d)", len(s))
	}
	return s, nil
}

// newInfoLogger returns the stdout logger passed to database helpers by the
// scan/verify and ranking subcommands.
func newInfoLogger() *log.Logger {
	return log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// connectRWStorage connects to PostgreSQL using the PGSQL_* environment
// variables and returns the RandomWalk repository plus the base Store for
// address lookups. The pool lives for the remainder of the process (rwctl
// runs one command and exits). The logger receives connect-retry and query
// traces.
func connectRWStorage(info *log.Logger) (*rwstore.Repo, *store.Store, error) {
	cfg := store.ConfigFromEnv()
	cfg.Logger = slog.New(slog.NewTextHandler(info.Writer(), nil))
	st, err := store.New(context.Background(), cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to storage: %w", err)
	}
	return rwstore.NewRepo(st), st, nil
}

// parseInt64 parses a base-10 int64 command argument, reporting the argument
// name on failure.
func parseInt64(name, value string) (int64, error) {
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %w", name, err)
	}
	return n, nil
}

// parseBigInt parses a base-10 big-integer command argument (wei amounts).
func parseBigInt(name, value string) (*big.Int, error) {
	n := new(big.Int)
	if _, ok := n.SetString(value, 10); !ok {
		return nil, fmt.Errorf("invalid %s: %s", name, value)
	}
	return n, nil
}

// weiToEthText formats a wei amount as a decimal ETH string with 18 decimal
// places, matching Ethereum precision.
func weiToEthText(wei *big.Int) string {
	if wei == nil {
		return "0.000000000000000000"
	}
	ether := new(big.Float).SetInt(wei)
	eth := new(big.Float).Quo(ether, big.NewFloat(1e18))
	return eth.Text('f', 18)
}

// weiToGwei converts a wei amount to gwei as float64 for display.
func weiToGwei(wei *big.Int) float64 {
	if wei == nil {
		return 0
	}
	gwei := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e9))
	f, _ := gwei.Float64()
	return f
}
