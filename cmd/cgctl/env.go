package main

// Shared configuration helpers used by the cgctl subcommands: environment
// variable access, transaction-session and read-only network construction,
// and the -i/--info flag. All transaction plumbing lives in internal/ethtx;
// this file is the only place cgctl reads the environment.

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// txEnvHelp documents the environment variables shared by all transaction
// subcommands; it is appended to their long help text.
const txEnvHelp = `Environment:
  RPC_URL               Ethereum RPC endpoint (required)
  PKEY_HEX              64-char hex private key, no 0x prefix (required)
  GAS_PRICE_MULTIPLIER  multiplier applied to the suggested gas price (default 2.0)`

// readEnvHelp documents the environment of read-only subcommands.
const readEnvHelp = `Environment:
  RPC_URL  Ethereum RPC endpoint (required)`

// rpcURLFromEnv returns the RPC endpoint from the RPC_URL environment
// variable.
func rpcURLFromEnv() (string, error) {
	url := os.Getenv("RPC_URL")
	if url == "" {
		return "", errors.New("RPC_URL environment variable not set")
	}
	return url, nil
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

// gasMultiplierFromEnv parses the GAS_PRICE_MULTIPLIER environment variable.
// Unset returns 0, which makes the session apply its default of 2.0. A value
// that does not parse to a positive number is a configuration error — the
// legacy code silently fell back to the default, which could double the gas
// cost an operator thought they had lowered.
func gasMultiplierFromEnv() (float64, error) {
	s := os.Getenv("GAS_PRICE_MULTIPLIER")
	if s == "" {
		return 0, nil
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil || f <= 0 {
		return 0, fmt.Errorf("GAS_PRICE_MULTIPLIER must be a positive number (got %q)", s)
	}
	return f, nil
}

// newTxSession connects to the RPC endpoint (RPC_URL), loads the signer from
// PKEY_HEX, applies the GAS_PRICE_MULTIPLIER policy and prints network and
// account details on the command's output when verbose is enabled.
func newTxSession(cmd *cobra.Command, verbose bool) (*ethtx.Session, error) {
	rpcURL, err := rpcURLFromEnv()
	if err != nil {
		return nil, err
	}
	pkey, err := pkeyHexFromEnv()
	if err != nil {
		return nil, err
	}
	multiplier, err := gasMultiplierFromEnv()
	if err != nil {
		return nil, err
	}
	return ethtx.New(cmd.Context(), ethtx.Options{
		RPCURL:             rpcURL,
		PrivateKeyHex:      pkey,
		Verbose:            verbose,
		Out:                cmd.OutOrStdout(),
		GasPriceMultiplier: multiplier,
	})
}

// connectNetwork connects to the RPC endpoint for a read-only subcommand and
// returns the network together with an always-verbose Output on the command's
// writer (read commands exist to print state).
func connectNetwork(cmd *cobra.Command) (*ethtx.Network, ethtx.Output, error) {
	out := ethtx.Output{Verbose: true, W: cmd.OutOrStdout()}
	rpcURL, err := rpcURLFromEnv()
	if err != nil {
		return nil, out, err
	}
	net, err := ethtx.Connect(cmd.Context(), rpcURL)
	if err != nil {
		return nil, out, fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)
	return net, out, nil
}

// addInfoFlag registers the -i/--info flag used by transaction subcommands to
// switch from quiet output (only success or error) to detailed output.
func addInfoFlag(c *cobra.Command, verbose *bool) {
	c.Flags().BoolVarP(verbose, "info", "i", false, "print detailed network, account and transaction information")
}
