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
		Use:   "set-round-activation <cosmicgame-addr> <timestamp>",
		Short: "Set roundActivationTime (owner only)",
		Long: `Set the round activation time to the given Unix timestamp.

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSetRoundActivation(cmd.Context(), ethtx.NewPrinter(info), args)
		},
	}
	c.Flags().BoolVarP(&info, "info", "i", false, "print detailed output")
	register(c)
}

func runSetRoundActivation(ctx context.Context, out *ethtx.Printer, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	timestamp, err := parseInt64("timestamp", args[1])
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

	currentActivation, err := game.RoundActivationTime(copts)
	if err != nil {
		return fmt.Errorf("getting current activation time: %w", err)
	}
	owner, err := game.Owner(copts)
	if err != nil {
		return fmt.Errorf("getting contract owner: %w", err)
	}

	secsUntilCurrent := currentActivation.Int64() - int64(net.BlockTime)
	secsUntilNew := timestamp - int64(net.BlockTime)

	out.Section("CURRENT STATE")
	out.KeyValue("Contract Owner", owner.String())
	out.KeyValue("Current Block Time", net.BlockTime)
	out.KeyValue("Current Activation Time", currentActivation.String())
	out.KeyValueDuration("Time Until Current Activation", secsUntilCurrent)

	out.Section("NEW VALUES")
	out.KeyValue("New Activation Time", timestamp)
	out.KeyValueDuration("Time Until New Activation", secsUntilNew)

	if acc.Address != owner {
		out.Section("WARNING")
		out.KeyValue("Your Address", acc.Address.String())
		out.KeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	out.TxSubmitting("SetRoundActivationTime", nil, ethtx.GasLimitAdminCall, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall)

	tx, err := game.SetRoundActivationTime(txopts, big.NewInt(timestamp))
	if err != nil {
		return fmt.Errorf("setRoundActivationTime: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}
