package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
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

			s, err := newTxSession(verbose)
			if err != nil {
				return err
			}
			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, s.net.client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}

			s.out.section("TRANSFER INFO")
			s.out.keyValue("Token ID", tokenID)
			s.out.keyValue("From", s.acc.address.String())
			s.out.keyValue("To", dstAddr.String())

			gasLimit := gasLimitContractCall
			txopts := transactOpts(s.net, s.acc, big.NewInt(0), gasLimit)
			s.out.txSubmitting("ERC721 TransferFrom", big.NewInt(0), gasLimit, adjustGasPrice(s.net.gasPrice))

			tx, err := rwalk.TransferFrom(txopts, s.acc.address, dstAddr, big.NewInt(tokenID))
			return s.finishTx(tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newTransferCmd()) }
