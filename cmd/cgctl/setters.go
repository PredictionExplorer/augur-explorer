package main

import (
	"context"
	"fmt"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
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
		Long: spec.long + `

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGameSetter(cmd.Context(), ethtx.NewPrinter(info), spec, args[0], args[1])
		},
	}
	c.Flags().BoolVarP(&info, "info", "i", false, "print detailed output")
	return c
}

func runGameSetter(ctx context.Context, out *ethtx.Printer, spec gameSetterSpec, addrArg, valueArg string) error {
	gameAddr, err := parseAddress("cosmicgame-addr", addrArg)
	if err != nil {
		return err
	}
	newValue, err := parseBigInt(spec.valueName, valueArg)
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

	currentValue, err := spec.read(game, ethtx.CallOpts())
	if err != nil {
		return fmt.Errorf("getting current value: %w", err)
	}

	suffix := ""
	if spec.percent {
		suffix = "%"
	}
	out.Section(spec.section)
	out.KeyValue("Current Value", currentValue.String()+suffix)
	out.KeyValue("New Value", newValue.String()+suffix)

	out.TxSubmitting(spec.action, nil, ethtx.GasLimitAdminCall, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall)

	tx, err := spec.write(game, txopts, newValue)
	if err != nil {
		return fmt.Errorf("%s: %w", spec.action, err)
	}
	out.TxSubmitted(tx)
	return nil
}
