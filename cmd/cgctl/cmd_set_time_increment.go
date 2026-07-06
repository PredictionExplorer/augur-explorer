package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	var info bool
	c := &cobra.Command{
		Use:   "set-time-increment <cosmicgame-addr> <time-increment-seconds>",
		Short: "Set mainPrizeTimeIncrementInMicroSeconds (owner only)",
		Long: `Set mainPrizeTimeIncrementInMicroSeconds so that each bid extends the time
until the main prize by the given number of seconds. Requires an inactive
round; see claim-and-set-time-increment for a variant that opens the inactive
window automatically.

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSetTimeIncrement(cmd.Context(), ethtx.NewPrinter(info), args)
		},
	}
	c.Flags().BoolVarP(&info, "info", "i", false, "print detailed output")
	register(c)
}

func runSetTimeIncrement(ctx context.Context, out *ethtx.Printer, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	desiredSeconds, err := parseInt64("time_increment_seconds", args[1])
	if err != nil {
		return err
	}
	if desiredSeconds <= 0 {
		return fmt.Errorf("time_increment_seconds must be positive")
	}

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)

	pkeyHex, err := ethtx.PrivateKeyHexFromEnv()
	if err != nil {
		return err
	}
	acc, err := net.PrepareAccount(ctx, pkeyHex)
	if err != nil {
		return fmt.Errorf("account setup failed: %w", err)
	}
	out.AccountInfo(acc)

	out.ContractInfo("CosmicGame Address", gameAddr)
	game, err := cgcontracts.NewCosmicSignatureGame(gameAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicGame: %w", err)
	}

	copts := ethtx.CallOpts()

	currentMicroseconds, err := game.MainPrizeTimeIncrementInMicroSeconds(copts)
	if err != nil {
		return fmt.Errorf("reading mainPrizeTimeIncrementInMicroSeconds: %w", err)
	}
	currentSeconds := new(big.Int).Div(currentMicroseconds, big.NewInt(1000000))

	owner, err := game.Owner(copts)
	if err != nil {
		return fmt.Errorf("getting contract owner: %w", err)
	}

	newMicroseconds := new(big.Int).Mul(big.NewInt(desiredSeconds), big.NewInt(1000000))

	out.Section("CURRENT STATE")
	out.KeyValue("Contract Owner", owner.String())
	out.KeyValue("Current Microseconds", currentMicroseconds.String())
	out.KeyValueDuration("Current Time Increment", currentSeconds.Int64())

	out.Section("NEW VALUES")
	out.KeyValue("New Microseconds", newMicroseconds.String())
	out.KeyValueDuration("New Time Increment", desiredSeconds)
	out.KeyValue("Formula", "timeIncrement (seconds) = microseconds / 1,000,000")

	if acc.Address != owner {
		out.Section("WARNING")
		out.KeyValue("Your Address", acc.Address.String())
		out.KeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	out.TxSubmitting("SetMainPrizeTimeIncrementInMicroSeconds", nil, ethtx.GasLimitAdminCall, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall)

	tx, err := game.SetMainPrizeTimeIncrementInMicroSeconds(txopts, newMicroseconds)
	if err != nil {
		return fmt.Errorf("setMainPrizeTimeIncrementInMicroSeconds: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}
