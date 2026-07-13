package main

// Shared configuration helpers used by multiple rwctl subcommands:
// environment-variable access, RPC and database connection helpers,
// well-known contract addresses, event topics, and argument parsing.

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
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
// scan/verify and ranking subcommands (connect-retry and query traces).
func newInfoLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}

// connectRWStorage connects to PostgreSQL using the DATABASE_URL / PGSQL_*
// environment variables and returns the RandomWalk repository plus the base
// Store for address lookups. The pool lives for the remainder of the process
// (rwctl runs one command and exits). The logger receives connect-retry and
// query traces.
func connectRWStorage(logger *slog.Logger) (*rwstore.Repo, *store.Store, error) {
	cfg := store.ConfigFromEnv()
	cfg.Logger = logger
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

// newTxSession connects to the RPC endpoint (RPC_URL), loads the signer from
// PKEY_HEX and prints network/account details on the command's output when
// verbose is enabled. The plumbing lives in internal/ethtx.
func newTxSession(cmd *cobra.Command, verbose bool) (*ethtx.Session, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		return nil, errors.New("RPC_URL environment variable not set")
	}
	pkey, err := pkeyHexFromEnv()
	if err != nil {
		return nil, err
	}
	return ethtx.New(cmd.Context(), ethtx.Options{
		RPCURL:        rpcURL,
		PrivateKeyHex: pkey,
		Verbose:       verbose,
		Out:           cmd.OutOrStdout(),
	})
}

// callOpts returns options for read-only contract calls.
func callOpts() *bind.CallOpts {
	return &bind.CallOpts{}
}
