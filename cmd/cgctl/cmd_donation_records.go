package main

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newDonationRecordsCmd builds the donation-records subcommand.
func newDonationRecordsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "donation-records [cosmicgame-addr]",
		Short: "Dump the ETH donation-with-info records from CosmicGame",
		Long: `Dump all EthDonationWithInfo records stored in the CosmicGame contract.

If no address is given, the default local Hardhat deployment address
` + defaultLocalGameAddr + ` is used.

` + readEnvHelp,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := defaultLocalGameAddr
			if len(args) == 1 {
				addr = args[0]
			} else {
				out := ethtx.Output{Verbose: true, W: cmd.OutOrStdout()}
				out.Section("DEFAULT ADDRESS")
				out.KeyValue("Using default", addr)
			}
			return runDonationRecords(cmd, addr)
		},
	}
}

func init() { register(newDonationRecordsCmd()) }

func runDonationRecords(cmd *cobra.Command, addrArg string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", addrArg)
	if err != nil {
		return err
	}

	net, out, err := connectNetwork(cmd)
	if err != nil {
		return err
	}
	out.ContractInfo("CosmicGame Address", gameAddr)

	game, err := cgcontracts.NewCosmicSignatureGame(gameAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicGame: %w", err)
	}

	copts := ethtx.CallOpts()

	numRecs, err := game.NumEthDonationWithInfoRecords(copts)
	if err != nil {
		return fmt.Errorf("getting NumEthDonationWithInfoRecords: %w", err)
	}

	out.Section("DONATION RECORDS")
	out.KeyValue("Total Records", numRecs.String())

	n := numRecs.Int64()
	if n == 0 {
		out.KeyValue("Note", "No ETH donation with info records found")
		return nil
	}

	for i := range n {
		rec, err := game.EthDonationWithInfoRecords(copts, big.NewInt(i))
		if err != nil {
			out.KeyValue("Error at record "+strconv.FormatInt(i, 10), err.Error())
			continue
		}
		out.Section("RECORD " + strconv.FormatInt(i, 10))
		out.KeyValue("Round Number", rec.RoundNum.String())
		out.KeyValue("Donor Address", rec.DonorAddress.String())
		out.KeyValueEth("Amount", rec.Amount)
		out.KeyValue("Data", rec.Data)
	}
	return nil
}
