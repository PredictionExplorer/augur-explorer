package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	var info bool
	c := &cobra.Command{
		Use:   "bid <cosmicgame-addr>",
		Short: "Make an ETH bid in the current CosmicGame round",
		Long: `Make an ETH bid in the current CosmicGame round at the next bid price.

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBid(cmd.Context(), ethtx.NewPrinter(info), args)
		},
	}
	c.Flags().BoolVarP(&info, "info", "i", false, "print detailed output (network, account, round info)")
	register(c)
}

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

func runBid(ctx context.Context, out *ethtx.Printer, args []string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", pickContractArg(args))
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

	out.Section("ROUND INFO")
	out.KeyValue("Round Number", roundNum.String())
	out.KeyValue("Total Bids This Round", totalBids.String())
	out.KeyValue("Last Bidder", lastBidder.String())
	out.KeyValueEth("Next Bid Price", bidPrice)

	if acc.Balance.Cmp(bidPrice) < 0 {
		return fmt.Errorf("insufficient balance: need %s ETH, have %s ETH",
			ethtx.WeiToEth(bidPrice), ethtx.WeiToEth(acc.Balance))
	}

	out.TxSubmitting("BidWithEth", bidPrice, ethtx.GasLimitBid, net.GasPrice)
	txopts := net.TransactOpts(acc, bidPrice, ethtx.GasLimitBid)

	tx, err := game.BidWithEth(txopts, big.NewInt(-1), "")
	if err != nil {
		return fmt.Errorf("bidWithEth: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}
