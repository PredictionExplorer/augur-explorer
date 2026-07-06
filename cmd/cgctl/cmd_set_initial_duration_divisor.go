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
		Use:   "set-initial-duration-divisor <cosmicgame-addr> <divisor>",
		Short: "Set initialDurationUntilMainPrizeDivisor (owner only)",
		Long: `Set the initial duration until main prize divisor (e.g. 100 = 1% bump on
the first bid). Initial timer after the first bid equals
mainPrizeTimeIncrementInMicroSeconds / divisor.

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSetInitialDurationDivisor(cmd.Context(), ethtx.NewPrinter(info), args)
		},
	}
	c.Flags().BoolVarP(&info, "info", "i", false, "print detailed output")
	register(c)
}

func runSetInitialDurationDivisor(ctx context.Context, out *ethtx.Printer, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	divisor, err := parseInt64("divisor", args[1])
	if err != nil {
		return err
	}
	if divisor <= 0 {
		return fmt.Errorf("divisor must be positive")
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

	currentDivisor, err := game.InitialDurationUntilMainPrizeDivisor(copts)
	if err != nil {
		return fmt.Errorf("getting current divisor: %w", err)
	}
	owner, err := game.Owner(copts)
	if err != nil {
		return fmt.Errorf("getting contract owner: %w", err)
	}

	out.Section("CURRENT STATE")
	out.KeyValue("Contract Owner", owner.String())
	out.KeyValue("Current Divisor", currentDivisor.String())
	out.KeyValue("Current Percentage", ethtx.ConvertToPercentage(currentDivisor))

	out.Section("NEW VALUES")
	out.KeyValue("New Divisor", divisor)
	out.KeyValue("New Percentage", ethtx.ConvertToPercentage(big.NewInt(divisor)))
	out.KeyValue("Formula", "percentage = 100 / divisor")

	if acc.Address != owner {
		out.Section("WARNING")
		out.KeyValue("Your Address", acc.Address.String())
		out.KeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	out.TxSubmitting("SetInitialDurationUntilMainPrizeDivisor", nil, ethtx.GasLimitAdminCall, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall)

	tx, err := game.SetInitialDurationUntilMainPrizeDivisor(txopts, big.NewInt(divisor))
	if err != nil {
		return fmt.Errorf("setInitialDurationUntilMainPrizeDivisor: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}
