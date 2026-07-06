package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	var info bool
	var delaySeconds int64 = -1
	c := &cobra.Command{
		Use:   "claim-prize <cosmicgame-addr>",
		Short: "Claim the CosmicGame main prize",
		Long: `Claim the main prize from CosmicGame.

With --delay, the command first sets delayDurationBeforeRoundActivation to the
given number of seconds and then claims, so the next round activates only after
that delay (this replaces the old claimprize-delay60 script; --delay without a
value uses 60 seconds).

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			out := ethtx.NewPrinter(info)
			if cmd.Flags().Changed("delay") {
				return runClaimPrizeWithDelay(cmd.Context(), out, args[0], delaySeconds)
			}
			return runClaimPrize(cmd.Context(), out, args[0])
		},
	}
	c.Flags().BoolVarP(&info, "info", "i", false, "print detailed output")
	c.Flags().Int64Var(&delaySeconds, "delay", 60, "set delayDurationBeforeRoundActivation to this many seconds before claiming")
	c.Flags().Lookup("delay").NoOptDefVal = "60"
	register(c)
}

func runClaimPrize(ctx context.Context, out *ethtx.Printer, addrArg string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", addrArg)
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

	roundNum, err := game.RoundNum(copts)
	if err != nil {
		return fmt.Errorf("getting round number: %w", err)
	}
	lastBidder, err := game.LastBidderAddress(copts)
	if err != nil {
		return fmt.Errorf("getting last bidder: %w", err)
	}
	prizeAmount, err := game.GetMainEthPrizeAmount(copts)
	if err != nil {
		return fmt.Errorf("getting prize amount: %w", err)
	}
	durationUntilPrize, err := game.GetDurationUntilMainPrize(copts)
	if err != nil {
		return fmt.Errorf("getting duration until prize: %w", err)
	}

	out.Section("PRIZE INFO")
	out.KeyValue("Round Number", roundNum.String())
	out.KeyValue("Last Bidder", lastBidder.String())
	out.KeyValueEth("Prize Amount", prizeAmount)
	out.KeyValueDuration("Time Until Prize", durationUntilPrize.Int64())

	if acc.Address != lastBidder {
		out.Section("WARNING")
		out.KeyValue("Your Address", acc.Address.String())
		out.KeyValue("Last Bidder", lastBidder.String())
		out.KeyValue("Note", "You are NOT the last bidder. Claim may fail unless timeout has passed.")
	}
	if durationUntilPrize.Int64() > 0 {
		out.Section("WARNING")
		out.KeyValue("Status", "Prize is NOT yet claimable")
		out.KeyValueDuration("Wait Time Remaining", durationUntilPrize.Int64())
	}

	out.TxSubmitting("ClaimMainPrize", nil, ethtx.GasLimitClaimPrize, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitClaimPrize)

	tx, err := game.ClaimMainPrize(txopts)
	if err != nil {
		return fmt.Errorf("claimMainPrize: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}

// runClaimPrizeWithDelay sets delayDurationBeforeRoundActivation and then
// claims the prize, so the new round activates after the requested delay.
func runClaimPrizeWithDelay(ctx context.Context, out *ethtx.Printer, addrArg string, delaySeconds int64) error {
	gameAddr, err := parseAddress("cosmicgame-addr", addrArg)
	if err != nil {
		return err
	}
	pkeyHex, err := ethtx.PrivateKeyHexFromEnv()
	if err != nil {
		return err
	}

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)

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
	prizeAmount, err := game.GetMainEthPrizeAmount(copts)
	if err != nil {
		return fmt.Errorf("getting prize amount: %w", err)
	}

	out.Section("CURRENT STATE")
	out.KeyValueDuration("Current Delay", currentDelay.Int64())
	out.KeyValueDuration("New Delay To Set", delaySeconds)
	out.KeyValueEth("Prize Amount", prizeAmount)

	out.Section("STEP 1: SET DELAY")
	out.TxSubmitting("SetDelayDurationBeforeRoundActivation", nil, ethtx.GasLimitAdminCall, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall)

	tx1, err := game.SetDelayDurationBeforeRoundActivation(txopts, big.NewInt(delaySeconds))
	if err != nil {
		return fmt.Errorf("failed to set delay, aborting: %w", err)
	}
	out.TxSubmitted(tx1)

	fmt.Printf("\nWaiting 2 seconds for tx to be mined...\n")
	time.Sleep(2 * time.Second)

	// Refresh account nonce and gas price for the second transaction.
	net2, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network refresh failed: %w", err)
	}
	acc, err = net2.PrepareAccount(ctx, pkeyHex)
	if err != nil {
		return fmt.Errorf("account refresh failed: %w", err)
	}

	out.Section("STEP 2: CLAIM PRIZE")
	out.TxSubmitting("ClaimMainPrize", nil, ethtx.GasLimitClaimPrize, net2.GasPrice)
	txopts2 := net2.TransactOpts(acc, nil, ethtx.GasLimitClaimPrize)

	tx2, err := game.ClaimMainPrize(txopts2)
	if err != nil {
		return fmt.Errorf("claimMainPrize: %w", err)
	}
	out.TxSubmitted(tx2)

	out.Section("SUMMARY")
	out.KeyValueDuration("Delay set to", delaySeconds)
	out.KeyValue("Status", "Prize claimed successfully")
	out.KeyValue("Note", fmt.Sprintf("New round will activate %d seconds after the claim", delaySeconds))
	return nil
}
