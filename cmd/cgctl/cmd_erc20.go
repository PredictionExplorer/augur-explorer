package main

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newERC20Cmd builds the erc20 command group.
func newERC20Cmd() *cobra.Command {
	group := &cobra.Command{
		Use:   "erc20",
		Short: "ERC-20 token helpers (balance, allowance, approve, revoke)",
	}
	group.AddCommand(newERC20BalanceCmd(), newERC20AllowanceCmd(), newERC20ApproveCmd(), newERC20RevokeCmd())
	return group
}

func init() { register(newERC20Cmd()) }

func newERC20BalanceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "balance <erc20-addr> <user-addr>",
		Short: "Show an ERC-20 token balance for a user address",
		Long: `Show an ERC-20 token balance for a user address, together with token
metadata and the user's ETH balance.

` + readEnvHelp,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runERC20Balance(cmd, args)
		},
	}
}

func newERC20AllowanceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "allowance <erc20-addr> <owner-addr> <spender-addr>",
		Short: "Show the ERC-20 allowance granted to a spender",
		Long: `Show the ERC-20 token allowance an owner has granted to a spender.

` + readEnvHelp,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runERC20Allowance(cmd, args)
		},
	}
}

func newERC20ApproveCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "approve <erc20-addr> <spender-addr>",
		Short: "Approve an unlimited (MAX_UINT256) ERC-20 allowance",
		Long: `Approve MAX_UINT256 allowance for a spender to spend your ERC-20 tokens.

` + txEnvHelp,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runERC20Approve(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func newERC20RevokeCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "revoke <erc20-addr> <spender-addr>",
		Short: "Revoke an ERC-20 allowance (set it to 0)",
		Long: `Revoke the allowance (set it to 0) for a spender of your ERC-20 tokens.

` + txEnvHelp,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runERC20Revoke(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func runERC20Balance(cmd *cobra.Command, args []string) error {
	contractAddr, err := parseAddress("erc20-addr", args[0])
	if err != nil {
		return err
	}
	userAddr, err := parseAddress("user-addr", args[1])
	if err != nil {
		return err
	}

	net, out, err := connectNetwork(cmd)
	if err != nil {
		return err
	}
	erc20, err := cgcontracts.NewERC20(contractAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC20 contract: %w", err)
	}

	copts := ethtx.CallOpts()

	name, err := erc20.Name(copts)
	if err != nil {
		name = "UNKNOWN"
	}
	symbol, err := erc20.Symbol(copts)
	if err != nil {
		symbol = "UNKNOWN"
	}
	decimals, err := erc20.Decimals(copts)
	if err != nil {
		decimals = 18
	}
	totalSupply, err := erc20.TotalSupply(copts)
	if err != nil {
		return fmt.Errorf("getting total supply: %w", err)
	}
	balance, err := erc20.BalanceOf(copts, userAddr)
	if err != nil {
		return fmt.Errorf("getting balance: %w", err)
	}
	ethBalance, err := net.Balance(cmd.Context(), userAddr)
	if err != nil {
		ethBalance = nil
	}

	out.Section("TOKEN INFO")
	out.KeyValue("Contract Address", contractAddr.String())
	out.KeyValue("Name", name)
	out.KeyValue("Symbol", symbol)
	out.KeyValue("Decimals", decimals)
	out.KeyValue("Total Supply (raw)", totalSupply.String())
	out.KeyValue("Total Supply", ethtx.FormatTokenAmount(totalSupply, decimals, symbol))

	out.Section("USER BALANCE")
	out.KeyValue("User Address", userAddr.String())
	out.KeyValue("Balance (raw)", balance.String())
	out.KeyValue("Balance", ethtx.FormatTokenAmount(balance, decimals, symbol))
	if ethBalance != nil {
		out.KeyValueEth("ETH Balance", ethBalance)
	}
	return nil
}

func runERC20Allowance(cmd *cobra.Command, args []string) error {
	tokenAddr, err := parseAddress("erc20-addr", args[0])
	if err != nil {
		return err
	}
	ownerAddr, err := parseAddress("owner-addr", args[1])
	if err != nil {
		return err
	}
	spenderAddr, err := parseAddress("spender-addr", args[2])
	if err != nil {
		return err
	}

	net, out, err := connectNetwork(cmd)
	if err != nil {
		return err
	}
	erc20, err := cgcontracts.NewERC20(tokenAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC20 contract: %w", err)
	}

	copts := ethtx.CallOpts()

	symbol, err := erc20.Symbol(copts)
	if err != nil {
		symbol = "UNKNOWN"
	}
	decimals, err := erc20.Decimals(copts)
	if err != nil {
		decimals = 18
	}
	allowance, err := erc20.Allowance(copts, ownerAddr, spenderAddr)
	if err != nil {
		return fmt.Errorf("getting allowance: %w", err)
	}
	ownerBalance, err := erc20.BalanceOf(copts, ownerAddr)
	if err != nil {
		return fmt.Errorf("getting owner balance: %w", err)
	}

	out.Section("TOKEN INFO")
	out.KeyValue("Token Address", tokenAddr.String())
	out.KeyValue("Token Symbol", symbol)
	out.KeyValue("Decimals", decimals)

	out.Section("ALLOWANCE INFO")
	out.KeyValue("Owner", ownerAddr.String())
	out.KeyValue("Spender", spenderAddr.String())
	out.KeyValue("Owner Balance (raw)", ownerBalance.String())
	out.KeyValue("Owner Balance", ethtx.FormatTokenAmount(ownerBalance, decimals, symbol))
	out.KeyValue("Allowance (raw)", allowance.String())
	out.KeyValue("Allowance", ethtx.FormatTokenAmount(allowance, decimals, symbol))

	switch {
	case allowance.Cmp(ethtx.MaxUint256()) == 0:
		out.KeyValue("Status", "UNLIMITED (MAX_UINT256)")
	case allowance.Cmp(ownerBalance) >= 0:
		out.KeyValue("Status", "Sufficient for full balance")
	default:
		out.KeyValue("Status", "Limited allowance")
	}
	return nil
}

func runERC20Approve(cmd *cobra.Command, verbose bool, args []string) error {
	tokenAddr, err := parseAddress("erc20-addr", args[0])
	if err != nil {
		return err
	}
	spenderAddr, err := parseAddress("spender-addr", args[1])
	if err != nil {
		return err
	}

	s, err := newTxSession(cmd, verbose)
	if err != nil {
		return err
	}
	erc20, err := cgcontracts.NewERC20(tokenAddr, s.Net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC20 contract: %w", err)
	}

	copts := ethtx.CallOpts()

	symbol, err := erc20.Symbol(copts)
	if err != nil {
		symbol = "UNKNOWN"
	}
	currentAllowance, err := erc20.Allowance(copts, s.Acc.Address, spenderAddr)
	if err != nil {
		return fmt.Errorf("getting current allowance: %w", err)
	}
	balance, err := erc20.BalanceOf(copts, s.Acc.Address)
	if err != nil {
		return fmt.Errorf("getting token balance: %w", err)
	}

	s.Out.Section("TOKEN INFO")
	s.Out.KeyValue("Token Address", tokenAddr.String())
	s.Out.KeyValue("Token Symbol", symbol)
	s.Out.KeyValue("Your Balance", balance.String())
	s.Out.KeyValue("Spender Address", spenderAddr.String())
	s.Out.KeyValue("Current Allowance", currentAllowance.String())

	s.Out.Section("APPROVAL INFO")
	s.Out.KeyValue("New Allowance", "MAX_UINT256 (unlimited)")

	s.Out.TxSubmitting("Approve (MAX_UINT256)", nil, ethtx.GasLimitApprove, s.AdjustedGasPrice())
	tx, err := erc20.Approve(s.TransactOpts(nil, ethtx.GasLimitApprove), spenderAddr, ethtx.MaxUint256())
	return s.FinishTx(cmd.Context(), tx, err)
}

func runERC20Revoke(cmd *cobra.Command, verbose bool, args []string) error {
	tokenAddr, err := parseAddress("erc20-addr", args[0])
	if err != nil {
		return err
	}
	spenderAddr, err := parseAddress("spender-addr", args[1])
	if err != nil {
		return err
	}

	s, err := newTxSession(cmd, verbose)
	if err != nil {
		return err
	}
	erc20, err := cgcontracts.NewERC20(tokenAddr, s.Net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC20 contract: %w", err)
	}

	copts := ethtx.CallOpts()

	symbol, err := erc20.Symbol(copts)
	if err != nil {
		symbol = "UNKNOWN"
	}
	currentAllowance, err := erc20.Allowance(copts, s.Acc.Address, spenderAddr)
	if err != nil {
		return fmt.Errorf("getting current allowance: %w", err)
	}

	s.Out.Section("TOKEN INFO")
	s.Out.KeyValue("Token Address", tokenAddr.String())
	s.Out.KeyValue("Token Symbol", symbol)
	s.Out.KeyValue("Spender Address", spenderAddr.String())
	s.Out.KeyValue("Current Allowance", currentAllowance.String())

	s.Out.Section("REVOKE INFO")
	s.Out.KeyValue("New Allowance", "0 (revoking)")

	s.Out.TxSubmitting("Approve (revoke to 0)", nil, ethtx.GasLimitApprove, s.AdjustedGasPrice())
	tx, err := erc20.Approve(s.TransactOpts(nil, ethtx.GasLimitApprove), spenderAddr, big.NewInt(0))
	return s.FinishTx(cmd.Context(), tx, err)
}
