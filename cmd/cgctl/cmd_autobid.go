package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Default values for the user-configurable autobid parameters.
const (
	defaultMaxEthBidEther      = 5    // in ETH (converted to wei)
	defaultMaxCstBidAmount     = 9    // in CST tokens (converted to wei)
	defaultRWalkBidStartPrice  = 0.1  // in ETH - only use RWALK when bid price above this
	defaultTimeUntilPrizeLimit = 15   // seconds before prize to start bidding
	defaultCstBidAnyway        = true // keep bidding with CST even when last bidder
)

func init() {
	register(&cobra.Command{
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
  AT_STARTUP_BID_UP_TO_PRICE_LEVEL  bid at round start until price reaches this level (ETH)`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := loadAutobidConfig()
			if err != nil {
				return err
			}
			bot, err := newBiddingBot(cfg)
			if err != nil {
				return fmt.Errorf("failed to initialize bot: %w", err)
			}
			return bot.run()
		},
	})
}

// autobidConfig holds the user-configurable autobid parameters.
type autobidConfig struct {
	maxEthBid        *big.Int // maximum ETH to spend on bidding (wei)
	maxCstBid        *big.Int // max CST price for bidding (wei)
	rwalkMinPrice    *big.Int // only use RWALK when bid price above this (wei)
	timeBeforePrize  int64    // seconds before prize to start bidding
	cstBidAnyway     bool     // keep bidding with CST even when last bidder
	initialBidPrice  *big.Int // initial bid price level (optional)
	rpcURL           string
	privateKeyHex    string
	gameContractAddr string
}

// loadAutobidConfig reads the autobid configuration from the environment.
func loadAutobidConfig() (autobidConfig, error) {
	if err := validateAutobidEnv(); err != nil {
		return autobidConfig{}, err
	}
	return autobidConfig{
		rpcURL:           os.Getenv("RPC_URL"),
		privateKeyHex:    os.Getenv("PKEY_HEX"),
		gameContractAddr: os.Getenv("CGAME_ADDR"),
		maxEthBid:        envBigIntEth("MAX_ETH_BID", defaultMaxEthBidEther),
		maxCstBid:        envBigIntEth("MAX_CST_BID", defaultMaxCstBidAmount),
		rwalkMinPrice:    envBigIntEth("RWALK_MIN_PRICE", defaultRWalkBidStartPrice),
		timeBeforePrize:  envInt64("TIME_BEFORE_PRIZE", defaultTimeUntilPrizeLimit),
		cstBidAnyway:     envBool("CST_BID_ANYWAY", defaultCstBidAnyway),
		initialBidPrice:  envBigIntEthOptional("AT_STARTUP_BID_UP_TO_PRICE_LEVEL"),
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
	} else if len(addr) != 40 {
		problems = append(problems, fmt.Sprintf("CGAME_ADDR must be 40 chars (got %d)", len(addr)))
	}

	if len(problems) > 0 {
		return fmt.Errorf("configuration errors:\n  - %s\n\nRequired: RPC_URL, PKEY_HEX, CGAME_ADDR\nOptional: MAX_ETH_BID, MAX_CST_BID, RWALK_MIN_PRICE, TIME_BEFORE_PRIZE, CST_BID_ANYWAY",
			strings.Join(problems, "\n  - "))
	}
	return nil
}

// envBigIntEth reads an ETH-denominated float from the environment and returns
// it in wei, falling back to defaultVal (also ETH) when unset or invalid.
func envBigIntEth(key string, defaultVal float64) *big.Int {
	if val := os.Getenv(key); val != "" {
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			wei := new(big.Float).Mul(big.NewFloat(f), big.NewFloat(1e18))
			result, _ := wei.Int(nil)
			return result
		}
		botLog("Warning: invalid %s, using default", key)
	}
	wei := new(big.Float).Mul(big.NewFloat(defaultVal), big.NewFloat(1e18))
	result, _ := wei.Int(nil)
	return result
}

// envBigIntEthOptional is like envBigIntEth but returns nil when unset.
func envBigIntEthOptional(key string) *big.Int {
	if val := os.Getenv(key); val != "" {
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			wei := new(big.Float).Mul(big.NewFloat(f), big.NewFloat(1e18))
			result, _ := wei.Int(nil)
			return result
		}
		botLog("Warning: invalid %s, ignoring", key)
	}
	return nil
}

func envInt64(key string, defaultVal int64) int64 {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			return i
		}
	}
	return defaultVal
}

func envBool(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		if b, err := strconv.ParseBool(val); err == nil {
			return b
		}
	}
	return defaultVal
}
