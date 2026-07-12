package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// gameSetterSpec describes a simple CosmicGame owner-only setter subcommand:
// read the current value, show current vs. new, and submit the setter call.
type gameSetterSpec struct {
	use       string // cobra Use line, e.g. "set-charity-percentage <cosmicgame-addr> <percentage>"
	short     string
	long      string
	section   string // section title for the current/new value block
	action    string // contract method name shown when submitting
	valueName string // name of the value argument for parse errors
	percent   bool   // append "%" when displaying values
	read      func(*cgcontracts.CosmicSignatureGame, *bind.CallOpts) (*big.Int, error)
	write     func(*cgcontracts.CosmicSignatureGame, *bind.TransactOpts, *big.Int) (*types.Transaction, error)
}

// newGameSetterCmd builds the cobra command for a gameSetterSpec.
func newGameSetterCmd(spec gameSetterSpec) *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   spec.use,
		Short: spec.short,
		Long:  spec.long + "\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGameSetter(cmd, info, spec, args[0], args[1])
		},
	}
	addInfoFlag(c, &info)
	return c
}

func runGameSetter(cmd *cobra.Command, verbose bool, spec gameSetterSpec, addrArg, valueArg string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", addrArg)
	if err != nil {
		return err
	}
	newValue, err := parseBigInt(spec.valueName, valueArg)
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

	currentValue, err := spec.read(game, ethtx.CallOpts())
	if err != nil {
		return fmt.Errorf("getting current value: %w", err)
	}

	suffix := ""
	if spec.percent {
		suffix = "%"
	}
	s.Out.Section(spec.section)
	s.Out.KeyValue("Current Value", currentValue.String()+suffix)
	s.Out.KeyValue("New Value", newValue.String()+suffix)

	s.Out.TxSubmitting(spec.action, nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	tx, err := spec.write(game, s.TransactOpts(nil, ethtx.GasLimitAdminCall), newValue)
	return s.FinishTx(cmd.Context(), tx, err)
}
