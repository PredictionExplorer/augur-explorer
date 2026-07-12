package main

import (
	"fmt"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newOwnerCmd builds the owner subcommand.
func newOwnerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "owner <contract-addr>",
		Short: "Show the owner of an Ownable contract",
		Long: `Show the owner of an Ownable contract and the owner's ETH balance.

` + readEnvHelp,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runOwner(cmd, args[0])
		},
	}
}

func init() { register(newOwnerCmd()) }

func runOwner(cmd *cobra.Command, addrArg string) error {
	contractAddr, err := parseAddress("contract-addr", addrArg)
	if err != nil {
		return err
	}

	net, out, err := connectNetwork(cmd)
	if err != nil {
		return err
	}
	out.ContractInfo("Contract Address", contractAddr)

	ownable, err := cgcontracts.NewOwnable(contractAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate Ownable contract: %w", err)
	}

	owner, err := ownable.Owner(ethtx.CallOpts())
	if err != nil {
		return fmt.Errorf("calling Owner(): %w", err)
	}

	ownerBalance, err := net.Balance(cmd.Context(), owner)
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
