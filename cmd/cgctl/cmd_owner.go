package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	c := &cobra.Command{
		Use:   "owner <contract-addr>",
		Short: "Show the owner of an Ownable contract",
		Long: `Show the owner of an Ownable contract and the owner's ETH balance.

Environment:
  RPC_URL  Ethereum RPC endpoint (required)`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runOwner(cmd.Context(), ethtx.NewPrinter(true), args[0])
		},
	}
	register(c)
}

func runOwner(ctx context.Context, out *ethtx.Printer, addrArg string) error {
	contractAddr, err := parseAddress("contract-addr", addrArg)
	if err != nil {
		return err
	}

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)
	out.ContractInfo("Contract Address", contractAddr)

	ownable, err := cgcontracts.NewOwnable(contractAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate Ownable contract: %w", err)
	}

	owner, err := ownable.Owner(ethtx.CallOpts())
	if err != nil {
		return fmt.Errorf("calling Owner(): %w", err)
	}

	ownerBalance, err := net.Balance(ctx, owner)
	if err != nil {
		ownerBalance = nil
	}

	out.Section("OWNERSHIP INFO")
	out.KeyValue("Owner Address", owner.String())
	if ownerBalance != nil {
		out.KeyValueEth("Owner Balance", ownerBalance)
	}
	return nil
}
