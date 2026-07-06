package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	group := &cobra.Command{
		Use:   "erc20",
		Short: "ERC-20 token helpers (balance, allowance, approve, revoke)",
	}

	group.AddCommand(&cobra.Command{
		Use:   "balance <erc20-addr> <user-addr>",
		Short: "Show an ERC-20 token balance for a user address",
		Long: `Show an ERC-20 token balance for a user address, together with token
metadata and the user's ETH balance.

Environment:
  RPC_URL  Ethereum RPC endpoint (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runERC20Balance(cmd.Context(), ethtx.NewPrinter(true), args)
		},
	})

	group.AddCommand(&cobra.Command{
		Use:   "allowance <erc20-addr> <owner-addr> <spender-addr>",
		Short: "Show the ERC-20 allowance granted to a spender",
		Long: `Show the ERC-20 token allowance an owner has granted to a spender.

Environment:
  RPC_URL  Ethereum RPC endpoint (required)`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runERC20Allowance(cmd.Context(), ethtx.NewPrinter(true), args)
		},
	})

	var approveInfo bool
	approve := &cobra.Command{
		Use:   "approve <erc20-addr> <spender-addr>",
		Short: "Approve an unlimited (MAX_UINT256) ERC-20 allowance",
		Long: `Approve MAX_UINT256 allowance for a spender to spend your ERC-20 tokens.

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runERC20Approve(cmd.Context(), ethtx.NewPrinter(approveInfo), args)
		},
	}
	approve.Flags().BoolVarP(&approveInfo, "info", "i", false, "print detailed output")
	group.AddCommand(approve)

	var revokeInfo bool
	revoke := &cobra.Command{
		Use:   "revoke <erc20-addr> <spender-addr>",
		Short: "Revoke an ERC-20 allowance (set it to 0)",
		Long: `Revoke the allowance (set it to 0) for a spender of your ERC-20 tokens.

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runERC20Revoke(cmd.Context(), ethtx.NewPrinter(revokeInfo), args)
		},
	}
	revoke.Flags().BoolVarP(&revokeInfo, "info", "i", false, "print detailed output")
	group.AddCommand(revoke)

	register(group)
}

func runERC20Balance(ctx context.Context, out *ethtx.Printer, args []string) error {
	contractAddr, err := parseAddress("erc20-addr", args[0])
	if err != nil {
		return err
	}
	userAddr, err := parseAddress("user-addr", args[1])
	if err != nil {
		return err
	}

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)

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
	ethBalance, err := net.Balance(ctx, userAddr)
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

func runERC20Allowance(ctx context.Context, out *ethtx.Printer, args []string) error {
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

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)

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

func runERC20Approve(ctx context.Context, out *ethtx.Printer, args []string) error {
	tokenAddr, err := parseAddress("erc20-addr", args[0])
	if err != nil {
		return err
	}
	spenderAddr, err := parseAddress("spender-addr", args[1])
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

	erc20, err := cgcontracts.NewERC20(tokenAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC20 contract: %w", err)
	}

	copts := ethtx.CallOpts()

	symbol, err := erc20.Symbol(copts)
	if err != nil {
		symbol = "UNKNOWN"
	}
	currentAllowance, err := erc20.Allowance(copts, acc.Address, spenderAddr)
	if err != nil {
		return fmt.Errorf("getting current allowance: %w", err)
	}
	balance, err := erc20.BalanceOf(copts, acc.Address)
	if err != nil {
		return fmt.Errorf("getting token balance: %w", err)
	}

	out.Section("TOKEN INFO")
	out.KeyValue("Token Address", tokenAddr.String())
	out.KeyValue("Token Symbol", symbol)
	out.KeyValue("Your Balance", balance.String())
	out.KeyValue("Spender Address", spenderAddr.String())
	out.KeyValue("Current Allowance", currentAllowance.String())

	out.Section("APPROVAL INFO")
	out.KeyValue("New Allowance", "MAX_UINT256 (unlimited)")

	out.TxSubmitting("Approve (MAX_UINT256)", nil, ethtx.GasLimitERC20Approve, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitERC20Approve)

	tx, err := erc20.Approve(txopts, spenderAddr, ethtx.MaxUint256())
	if err != nil {
		return fmt.Errorf("approve: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}

func runERC20Revoke(ctx context.Context, out *ethtx.Printer, args []string) error {
	tokenAddr, err := parseAddress("erc20-addr", args[0])
	if err != nil {
		return err
	}
	spenderAddr, err := parseAddress("spender-addr", args[1])
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

	erc20, err := cgcontracts.NewERC20(tokenAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC20 contract: %w", err)
	}

	copts := ethtx.CallOpts()

	symbol, err := erc20.Symbol(copts)
	if err != nil {
		symbol = "UNKNOWN"
	}
	currentAllowance, err := erc20.Allowance(copts, acc.Address, spenderAddr)
	if err != nil {
		return fmt.Errorf("getting current allowance: %w", err)
	}

	out.Section("TOKEN INFO")
	out.KeyValue("Token Address", tokenAddr.String())
	out.KeyValue("Token Symbol", symbol)
	out.KeyValue("Spender Address", spenderAddr.String())
	out.KeyValue("Current Allowance", currentAllowance.String())

	out.Section("REVOKE INFO")
	out.KeyValue("New Allowance", "0 (revoking)")

	out.TxSubmitting("Approve (revoke to 0)", nil, ethtx.GasLimitERC20Approve, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitERC20Approve)

	tx, err := erc20.Approve(txopts, spenderAddr, big.NewInt(0))
	if err != nil {
		return fmt.Errorf("approve: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}
