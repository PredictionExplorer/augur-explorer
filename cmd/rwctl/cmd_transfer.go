package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newTransferCmd builds the transfer subcommand (legacy transfer script).
func newTransferCmd() *cobra.Command {
	var verbose bool
	c := &cobra.Command{
		Use:   "transfer [rwalk_addr] [token_id] [new_owner_addr]",
		Short: "Transfer a RandomWalk token to a new owner",
		Long:  "Transfers RandomWalk token to a new owner.\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			rwalkAddr := common.HexToAddress(args[0])
			tokenID, err := parseInt64("token_id", args[1])
			if err != nil {
				return err
			}
			dstAddr := common.HexToAddress(args[2])

			s, err := newTxSession(cmd, verbose)
			if err != nil {
				return err
			}
			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, s.Net.Client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}

			s.Out.Section("TRANSFER INFO")
			s.Out.KeyValue("Token ID", tokenID)
			s.Out.KeyValue("From", s.Acc.Address.String())
			s.Out.KeyValue("To", dstAddr.String())

			gasLimit := ethtx.GasLimitContractCall
			txopts := s.TransactOpts(big.NewInt(0), gasLimit)
			s.Out.TxSubmitting("ERC721 TransferFrom", big.NewInt(0), gasLimit, ethtx.AdjustGasPrice(s.Net.GasPrice))

			tx, err := rwalk.TransferFrom(txopts, s.Acc.Address, dstAddr, big.NewInt(tokenID))
			return s.FinishTx(cmd.Context(), tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newTransferCmd()) }
