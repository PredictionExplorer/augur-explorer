package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/cmd/cgctl/internal/ethtx"
	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	group := &cobra.Command{
		Use:   "nft",
		Short: "ERC-721 / CosmicSignature NFT helpers",
	}

	group.AddCommand(&cobra.Command{
		Use:   "approved <erc721-addr> <token-id>",
		Short: "Show the approved operator of a single ERC-721 token",
		Long: `Show the ERC-721 single-token approval status (getApproved) together with
the token owner.

Environment:
  RPC_URL  Ethereum RPC endpoint (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNFTApproved(cmd.Context(), ethtx.NewPrinter(true), args)
		},
	})

	group.AddCommand(&cobra.Command{
		Use:   "is-approved-for-all <erc721-addr> <owner-addr> <operator-addr>",
		Short: "Show the ERC-721 operator-level approval status",
		Long: `Show the ERC-721 isApprovedForAll status (operator-level approval).

Environment:
  RPC_URL  Ethereum RPC endpoint (required)`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNFTIsApprovedForAll(cmd.Context(), ethtx.NewPrinter(true), args)
		},
	})

	group.AddCommand(&cobra.Command{
		Use:   "owner-of <erc721-addr> <token-id>",
		Short: "Show the owner of a specific ERC-721 token",
		Long: `Show the owner of a specific ERC-721 token, plus the owner's NFT count
and ETH balance.

Environment:
  RPC_URL  Ethereum RPC endpoint (required)`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNFTOwnerOf(cmd.Context(), ethtx.NewPrinter(true), args)
		},
	})

	var setNameInfo bool
	setName := &cobra.Command{
		Use:   "set-name <cosmicsignaturenft-addr> <token-id> [name]",
		Short: "Set the name of a CosmicSignatureNft token",
		Long: `Set the NFT name for a CosmicSignatureNft token. If the name is omitted,
the name is set to the empty string.

Environment:
  RPC_URL   Ethereum RPC endpoint (required)
  PKEY_HEX  64-char hex private key, no 0x prefix (required)`,
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNFTSetName(cmd.Context(), ethtx.NewPrinter(setNameInfo), args)
		},
	}
	setName.Flags().BoolVarP(&setNameInfo, "info", "i", false, "print detailed output")
	group.AddCommand(setName)

	register(group)
}

func runNFTApproved(ctx context.Context, out *ethtx.Printer, args []string) error {
	nftAddr, err := parseAddress("erc721-addr", args[0])
	if err != nil {
		return err
	}
	tokenID, err := parseInt64("token-id", args[1])
	if err != nil {
		return err
	}

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)

	erc721, err := cgcontracts.NewCosmicSignatureNft(nftAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC721 contract: %w", err)
	}

	copts := ethtx.CallOpts()

	operator, err := erc721.GetApproved(copts, big.NewInt(tokenID))
	if err != nil {
		return fmt.Errorf("GetApproved(): %w", err)
	}
	owner, err := erc721.OwnerOf(copts, big.NewInt(tokenID))
	if err != nil {
		return fmt.Errorf("OwnerOf(): %w", err)
	}

	out.Section("TOKEN INFO")
	out.KeyValue("Contract Address", nftAddr.String())
	out.KeyValue("Token ID", tokenID)

	out.Section("APPROVAL STATUS")
	out.KeyValue("Token Owner", owner.String())
	out.KeyValue("Approved Operator", operator.String())

	if operator == (common.Address{}) {
		out.KeyValue("Status", "NOT APPROVED - No operator set for this token")
	} else {
		out.KeyValue("Status", "APPROVED - Operator can transfer this token")
	}
	return nil
}

func runNFTIsApprovedForAll(ctx context.Context, out *ethtx.Printer, args []string) error {
	tokenAddr, err := parseAddress("erc721-addr", args[0])
	if err != nil {
		return err
	}
	ownerAddr, err := parseAddress("owner-addr", args[1])
	if err != nil {
		return err
	}
	operatorAddr, err := parseAddress("operator-addr", args[2])
	if err != nil {
		return err
	}

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)

	erc721, err := cgcontracts.NewCosmicSignatureNft(tokenAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC721 contract: %w", err)
	}

	copts := ethtx.CallOpts()

	isApproved, err := erc721.IsApprovedForAll(copts, ownerAddr, operatorAddr)
	if err != nil {
		return fmt.Errorf("calling IsApprovedForAll(): %w", err)
	}
	balance, err := erc721.BalanceOf(copts, ownerAddr)
	if err != nil {
		balance = nil
	}

	out.Section("TOKEN INFO")
	out.KeyValue("Token Contract", tokenAddr.String())

	out.Section("APPROVAL STATUS")
	out.KeyValue("Owner", ownerAddr.String())
	out.KeyValue("Operator", operatorAddr.String())
	if balance != nil {
		out.KeyValue("Owner's Token Balance", balance.String())
	}
	out.KeyValue("Is Approved For All", isApproved)

	if isApproved {
		out.KeyValue("Status", "APPROVED - Operator can transfer all owner's tokens")
	} else {
		out.KeyValue("Status", "NOT APPROVED - Operator cannot transfer owner's tokens")
	}
	return nil
}

func runNFTOwnerOf(ctx context.Context, out *ethtx.Printer, args []string) error {
	contractAddr, err := parseAddress("erc721-addr", args[0])
	if err != nil {
		return err
	}
	tokenID, err := parseInt64("token-id", args[1])
	if err != nil {
		return err
	}

	net, err := ethtx.Connect(ctx)
	if err != nil {
		return fmt.Errorf("network connection failed: %w", err)
	}
	out.NetworkInfo(net)

	erc721, err := cgcontracts.NewERC721(contractAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC721 contract: %w", err)
	}

	copts := ethtx.CallOpts()

	owner, err := erc721.OwnerOf(copts, big.NewInt(tokenID))
	if err != nil {
		return fmt.Errorf("calling OwnerOf(): %w (token may not exist)", err)
	}
	balance, err := erc721.BalanceOf(copts, owner)
	if err != nil {
		balance = nil
	}
	ownerEthBalance, err := net.Balance(ctx, owner)
	if err != nil {
		ownerEthBalance = nil
	}

	out.Section("TOKEN INFO")
	out.KeyValue("Contract Address", contractAddr.String())
	out.KeyValue("Token ID", tokenID)

	out.Section("OWNERSHIP INFO")
	out.KeyValue("Owner Address", owner.String())
	if balance != nil {
		out.KeyValue("Owner's Total NFTs", balance.String())
	}
	if ownerEthBalance != nil {
		out.KeyValueEth("Owner's ETH Balance", ownerEthBalance)
	}
	return nil
}

func runNFTSetName(ctx context.Context, out *ethtx.Printer, args []string) error {
	nftAddr, err := parseAddress("cosmicsignaturenft-addr", args[0])
	if err != nil {
		return err
	}
	tokenID, err := parseBigInt("token_id", args[1])
	if err != nil {
		return err
	}
	nftName := ""
	if len(args) == 3 {
		nftName = args[2]
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

	out.ContractInfo("CosmicSignatureNft Address", nftAddr)
	nft, err := cgcontracts.NewCosmicSignatureNft(nftAddr, net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicSignatureNft: %w", err)
	}

	out.Section("NFT NAME CONFIG")
	out.KeyValue("Token ID", tokenID.String())
	if nftName == "" {
		out.KeyValue("New Name", "(empty)")
	} else {
		out.KeyValue("New Name", nftName)
	}

	out.TxSubmitting("SetNftName", nil, ethtx.GasLimitAdminCall, net.GasPrice)
	txopts := net.TransactOpts(acc, nil, ethtx.GasLimitAdminCall)

	tx, err := nft.SetNftName(txopts, tokenID, nftName)
	if err != nil {
		return fmt.Errorf("setNftName: %w", err)
	}
	out.TxSubmitted(tx)
	return nil
}
