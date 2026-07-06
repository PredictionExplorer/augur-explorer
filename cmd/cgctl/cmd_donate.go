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
		Use:   "donate <cosmicgame-addr> <amount-wei>",
		Short: "Donate ETH to the CosmicGame contract",
		Long: `Donate ETH to the CosmicGame contract (DonateEth).

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDonate(cmd.Context(), ethtx.NewPrinter(info), args)
		},
	}
	c.Flags().BoolVarP(&info, "info", "i", false, "print detailed output")
	register(c)
}

func runDonate(ctx context.Context, out *ethtx.Printer, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", args[0])
	if err != nil {
		return err
	}
	donationAmount, err := parseBigInt("amount", args[1])
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

	contractBalance, err := net.Balance(ctx, gameAddr)
	if err != nil {
		return fmt.Errorf("getting contract balance: %w", err)
	}

	out.Section("DONATION INFO")
	out.KeyValueEth("Contract Current Balance", contractBalance)
	out.KeyValueEth("Donation Amount", donationAmount)
	out.KeyValueEth("Contract Balance After", new(big.Int).Add(contractBalance, donationAmount))

	if acc.Balance.Cmp(donationAmount) < 0 {
		return fmt.Errorf("insufficient balance: need %s ETH, have %s ETH",
			ethtx.WeiToEth(donationAmount), ethtx.WeiToEth(acc.Balance))
	}

	out.TxSubmitting("DonateEth", donationAmount, ethtx.GasLimitDonate, net.GasPrice)
	txopts := net.TransactOpts(acc, donationAmount, ethtx.GasLimitDonate)

	tx, err := game.DonateEth(txopts)
	if err != nil {
		return fmt.Errorf("donateEth: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}
