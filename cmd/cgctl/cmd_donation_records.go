package main

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	c := &cobra.Command{
		Use:   "donation-records [cosmicgame-addr]",
		Short: "Dump the ETH donation-with-info records from CosmicGame",
		Long: `Dump all EthDonationWithInfo records stored in the CosmicGame contract.

If no address is given, the default local Hardhat deployment address
` + defaultLocalGameAddr + ` is used.

Environment:
  RPC_URL  Ethereum RPC endpoint (required)`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			out := ethtx.NewPrinter(true)
			addr := defaultLocalGameAddr
			if len(args) == 1 {
				addr = args[0]
			} else {
				out.Section("DEFAULT ADDRESS")
				out.KeyValue("Using default", addr)
			}
			return runDonationRecords(cmd.Context(), out, addr)
		},
	}
	register(c)
}

func runDonationRecords(ctx context.Context, out *ethtx.Printer, addrArg string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", addrArg)
	if err != nil {
		return err
	}

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)
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

	for i := int64(0); i < n; i++ {
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
