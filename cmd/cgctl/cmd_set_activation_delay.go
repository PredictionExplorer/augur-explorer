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
		Use:   "set-activation-delay <cosmicgame-addr> <seconds>",
		Short: "Set delayDurationBeforeRoundActivation (owner only)",
		Long: `Set the delay duration before the next round activates after a prize claim.

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSetActivationDelay(cmd.Context(), ethtx.NewPrinter(info), args)
		},
	}
	c.Flags().BoolVarP(&info, "info", "i", false, "print detailed output")
	register(c)
}

func runSetActivationDelay(ctx context.Context, out *ethtx.Printer, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	seconds, err := parseInt64("seconds", args[1])
	if err != nil {
		return err
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

	currentDelay, err := game.DelayDurationBeforeRoundActivation(copts)
	if err != nil {
		return fmt.Errorf("getting current delay: %w", err)
	}
	owner, err := game.Owner(copts)
	if err != nil {
		return fmt.Errorf("getting contract owner: %w", err)
	}

	out.Section("CURRENT STATE")
	out.KeyValue("Contract Owner", owner.String())
	out.KeyValueDuration("Current Delay", currentDelay.Int64())
	out.KeyValueDuration("New Delay", seconds)

	if acc.Address != owner {
		out.Section("WARNING")
		out.KeyValue("Your Address", acc.Address.String())
		out.KeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	out.TxSubmitting("SetDelayDurationBeforeRoundActivation", nil, ethtx.GasLimitAdminCall, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall)

	tx, err := game.SetDelayDurationBeforeRoundActivation(txopts, big.NewInt(seconds))
	if err != nil {
		return fmt.Errorf("setDelayDurationBeforeRoundActivation: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}
