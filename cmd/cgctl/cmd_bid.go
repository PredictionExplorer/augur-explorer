package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newBidCmd builds the bid subcommand.
func newBidCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "bid <cosmicgame-addr>",
		Short: "Make an ETH bid in the current CosmicGame round",
		Long:  "Make an ETH bid in the current CosmicGame round at the next bid price.\n\n" + txEnvHelp,
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBid(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func init() { register(newBidCmd()) }

// pickContractArg returns the argument that looks like an Ethereum address
// (0x + 40 hex chars). If one arg is given, it is returned. If two are given
// (e.g. tx hash + address), the address-shaped one wins.
func pickContractArg(args []string) string {
	if len(args) == 0 {
		return ""
	}
	if len(args) == 1 {
		return args[0]
	}
	for _, a := range args {
		trimmed := strings.TrimPrefix(a, "0x")
		if len(trimmed) == 40 {
			if strings.HasPrefix(a, "0x") {
				return a
			}
			return "0x" + trimmed
		}
	}
	return args[1]
}

func runBid(cmd *cobra.Command, verbose bool, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", pickContractArg(args))
	if err != nil {
		return err
	}

	s, err := newTxSession(cmd, verbose)
	if err != nil {
		return err
	}
	s.Out.ContractInfo("CosmicGame Address", gameAddr)
	game, err := cgcontracts.NewCosmicSignatureGame(gameAddr, s.Net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicGame: %w", err)
	}

	copts := ethtx.CallOpts()

	roundNum, err := game.RoundNum(copts)
	if err != nil {
		return fmt.Errorf("getting round number: %w", err)
	}
	bidPrice, err := game.GetNextEthBidPrice(copts)
	if err != nil {
		return fmt.Errorf("getting bid price: %w", err)
	}
	lastBidder, err := game.LastBidderAddress(copts)
	if err != nil {
		return fmt.Errorf("getting last bidder: %w", err)
	}
	totalBids, err := game.GetTotalNumBids(copts, roundNum)
	if err != nil {
		return fmt.Errorf("getting total bids: %w", err)
	}

	s.Out.Section("ROUND INFO")
	s.Out.KeyValue("Round Number", roundNum.String())
	s.Out.KeyValue("Total Bids This Round", totalBids.String())
	s.Out.KeyValue("Last Bidder", lastBidder.String())
	s.Out.KeyValueEth("Next Bid Price", bidPrice)

	if s.Acc.Balance.Cmp(bidPrice) < 0 {
		return fmt.Errorf("insufficient balance: need %s ETH, have %s ETH",
			ethtx.WeiToEthText(bidPrice), ethtx.WeiToEthText(s.Acc.Balance))
	}

	s.Out.TxSubmitting("BidWithEth", bidPrice, ethtx.GasLimitBid, s.AdjustedGasPrice())
	tx, err := game.BidWithEth(s.TransactOpts(bidPrice, ethtx.GasLimitBid), big.NewInt(-1), "")
	return s.FinishTx(cmd.Context(), tx, err)
}
