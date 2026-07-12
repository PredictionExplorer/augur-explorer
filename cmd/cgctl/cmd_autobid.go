package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/autobid"
)

// Default values for the user-configurable autobid parameters.
const (
	defaultMaxEthBidEther      = 5    // in ETH (converted to wei)
	defaultMaxCstBidAmount     = 9    // in CST tokens (converted to wei)
	defaultRWalkBidStartPrice  = 0.1  // in ETH - only use RWALK when bid price above this
	defaultTimeUntilPrizeLimit = 15   // seconds before prize to start bidding
	defaultCstBidAnyway        = true // keep bidding with CST even when last bidder
)

// newAutobidCmd builds the autobid subcommand.
func newAutobidCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "autobid",
		Short: "Run the automated CosmicGame bidding bot",
		Long: `Run the automated CosmicGame bidding bot: an event-based state machine
that bids with ETH, CST, or RandomWalk NFTs, claims the prize when it wins,
and exits when the round ends (unless initial bidding is configured for the
next round).

Environment (required):
  RPC_URL     Ethereum RPC endpoint
  PKEY_HEX    64-char hex private key, no 0x prefix
  CGAME_ADDR  CosmicGame contract address, 40 hex chars without 0x prefix

Environment (optional):
  MAX_ETH_BID                       max ETH per bid (default 5)
  MAX_CST_BID                       max CST price for CST bids (default 9)
  RWALK_MIN_PRICE                   min ETH bid price before RandomWalk bidding is considered (default 0.1)
  TIME_BEFORE_PRIZE                 seconds before prize to start bidding (default 15)
  CST_BID_ANYWAY                    keep bidding with CST even when last bidder (default true)
  AT_STARTUP_BID_UP_TO_PRICE_LEVEL  bid at round start until price reaches this level (ETH)
  GAS_PRICE_MULTIPLIER              multiplier applied to the suggested gas price (default 2.0)`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := loadAutobidConfig()
			if err != nil {
				return err
			}
			cfg.Out = cmd.OutOrStdout()
			engine, err := autobid.New(cmd.Context(), cfg)
			if err != nil {
				return fmt.Errorf("failed to initialize bot: %w", err)
			}
			return engine.Run(cmd.Context())
		},
	}
}

func init() { register(newAutobidCmd()) }

// loadAutobidConfig reads the autobid configuration from the environment.
// Malformed numeric values are configuration errors — the legacy bot
// silently substituted its defaults, which could bid with 50x the intended
// limit on a typo.
func loadAutobidConfig() (autobid.Config, error) {
	if err := validateAutobidEnv(); err != nil {
		return autobid.Config{}, err
	}
	maxEthBid, err := envBigIntEth("MAX_ETH_BID", defaultMaxEthBidEther)
	if err != nil {
		return autobid.Config{}, err
	}
	maxCstBid, err := envBigIntEth("MAX_CST_BID", defaultMaxCstBidAmount)
	if err != nil {
		return autobid.Config{}, err
	}
	rwalkMinPrice, err := envBigIntEth("RWALK_MIN_PRICE", defaultRWalkBidStartPrice)
	if err != nil {
		return autobid.Config{}, err
	}
	timeBeforePrize, err := envInt64("TIME_BEFORE_PRIZE", defaultTimeUntilPrizeLimit)
	if err != nil {
		return autobid.Config{}, err
	}
	cstBidAnyway, err := envBool("CST_BID_ANYWAY", defaultCstBidAnyway)
	if err != nil {
		return autobid.Config{}, err
	}
	initialBidPrice, err := envBigIntEthOptional("AT_STARTUP_BID_UP_TO_PRICE_LEVEL")
	if err != nil {
		return autobid.Config{}, err
	}
	multiplier, err := gasMultiplierFromEnv()
	if err != nil {
		return autobid.Config{}, err
	}
	return autobid.Config{
		RPCURL:        os.Getenv("RPC_URL"),
		PrivateKeyHex: os.Getenv("PKEY_HEX"),
		GameAddr:      common.HexToAddress(os.Getenv("CGAME_ADDR")),
		Limits: autobid.Limits{
			MaxEthBid:       maxEthBid,
			MaxCstBid:       maxCstBid,
			RWalkMinPrice:   rwalkMinPrice,
			TimeBeforePrize: timeBeforePrize,
			CstBidAnyway:    cstBidAnyway,
		},
		InitialBidPrice:    initialBidPrice,
		GasPriceMultiplier: multiplier,
	}, nil
}

func validateAutobidEnv() error {
	var problems []string

	if os.Getenv("RPC_URL") == "" {
		problems = append(problems, "RPC_URL is required")
	}

	pkey := os.Getenv("PKEY_HEX")
	if pkey == "" {
		problems = append(problems, "PKEY_HEX is required")
	} else if len(pkey) != 64 {
		problems = append(problems, fmt.Sprintf("PKEY_HEX must be 64 chars (got %d)", len(pkey)))
	}

	addr := os.Getenv("CGAME_ADDR")
	if addr == "" {
		problems = append(problems, "CGAME_ADDR is required")
	} else if len(strings.TrimPrefix(addr, "0x")) != 40 {
		problems = append(problems, fmt.Sprintf("CGAME_ADDR must be 40 hex chars without 0x prefix (got %d)", len(addr)))
	}

	if len(problems) > 0 {
		return fmt.Errorf("configuration errors:\n  - %s\n\nRequired: RPC_URL, PKEY_HEX, CGAME_ADDR\nOptional: MAX_ETH_BID, MAX_CST_BID, RWALK_MIN_PRICE, TIME_BEFORE_PRIZE, CST_BID_ANYWAY",
			strings.Join(problems, "\n  - "))
	}
	return nil
}

// ethFloatToWei converts an ETH-denominated float to wei.
func ethFloatToWei(f float64) *big.Int {
	wei := new(big.Float).Mul(big.NewFloat(f), big.NewFloat(1e18))
	result, _ := wei.Int(nil)
	return result
}

// envBigIntEth reads an ETH-denominated float from the environment and
// returns it in wei, using defaultVal (also ETH) when unset.
func envBigIntEth(key string, defaultVal float64) (*big.Int, error) {
	val := os.Getenv(key)
	if val == "" {
		return ethFloatToWei(defaultVal), nil
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil || f < 0 {
		return nil, fmt.Errorf("%s must be a non-negative number (got %q)", key, val)
	}
	return ethFloatToWei(f), nil
}

// envBigIntEthOptional is like envBigIntEth but returns nil when unset.
func envBigIntEthOptional(key string) (*big.Int, error) {
	val := os.Getenv(key)
	if val == "" {
		return nil, nil
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil || f < 0 {
		return nil, fmt.Errorf("%s must be a non-negative number (got %q)", key, val)
	}
	return ethFloatToWei(f), nil
}

func envInt64(key string, defaultVal int64) (int64, error) {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal, nil
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer (got %q)", key, val)
	}
	return i, nil
}

func envBool(key string, defaultVal bool) (bool, error) {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal, nil
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return false, fmt.Errorf("%s must be a boolean (got %q)", key, val)
	}
	return b, nil
}
