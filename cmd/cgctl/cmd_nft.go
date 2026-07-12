package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newNFTCmd builds the nft command group.
func newNFTCmd() *cobra.Command {
	group := &cobra.Command{
		Use:   "nft",
		Short: "ERC-721 / CosmicSignature NFT helpers",
	}
	group.AddCommand(newNFTApprovedCmd(), newNFTIsApprovedForAllCmd(), newNFTOwnerOfCmd(), newNFTSetNameCmd())
	return group
}

func init() { register(newNFTCmd()) }

func newNFTApprovedCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "approved <erc721-addr> <token-id>",
		Short: "Show the approved operator of a single ERC-721 token",
		Long: `Show the ERC-721 single-token approval status (getApproved) together with
the token owner.

` + readEnvHelp,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNFTApproved(cmd, args)
		},
	}
}

func newNFTIsApprovedForAllCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "is-approved-for-all <erc721-addr> <owner-addr> <operator-addr>",
		Short: "Show the ERC-721 operator-level approval status",
		Long: `Show the ERC-721 isApprovedForAll status (operator-level approval).

` + readEnvHelp,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNFTIsApprovedForAll(cmd, args)
		},
	}
}

func newNFTOwnerOfCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "owner-of <erc721-addr> <token-id>",
		Short: "Show the owner of a specific ERC-721 token",
		Long: `Show the owner of a specific ERC-721 token, plus the owner's NFT count
and ETH balance.

` + readEnvHelp,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNFTOwnerOf(cmd, args)
		},
	}
}

func newNFTSetNameCmd() *cobra.Command {
	var info bool
	c := &cobra.Command{
		Use:   "set-name <cosmicsignaturenft-addr> <token-id> [name]",
		Short: "Set the name of a CosmicSignatureNft token",
		Long: `Set the NFT name for a CosmicSignatureNft token. If the name is omitted,
the name is set to the empty string.

` + txEnvHelp,
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNFTSetName(cmd, info, args)
		},
	}
	addInfoFlag(c, &info)
	return c
}

func runNFTApproved(cmd *cobra.Command, args []string) error {
	nftAddr, err := parseAddress("erc721-addr", args[0])
	if err != nil {
		return err
	}
	tokenID, err := parseInt64("token-id", args[1])
	if err != nil {
		return err
	}

	net, out, err := connectNetwork(cmd)
	if err != nil {
		return err
	}
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

func runNFTIsApprovedForAll(cmd *cobra.Command, args []string) error {
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

	net, out, err := connectNetwork(cmd)
	if err != nil {
		return err
	}
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

func runNFTOwnerOf(cmd *cobra.Command, args []string) error {
	contractAddr, err := parseAddress("erc721-addr", args[0])
	if err != nil {
		return err
	}
	tokenID, err := parseInt64("token-id", args[1])
	if err != nil {
		return err
	}

	net, out, err := connectNetwork(cmd)
	if err != nil {
		return err
	}
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
	ownerEthBalance, err := net.Balance(cmd.Context(), owner)
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

func runNFTSetName(cmd *cobra.Command, verbose bool, args []string) error {
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

	s, err := newTxSession(cmd, verbose)
	if err != nil {
		return err
	}
	s.Out.ContractInfo("CosmicSignatureNft Address", nftAddr)
	nft, err := cgcontracts.NewCosmicSignatureNft(nftAddr, s.Net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate CosmicSignatureNft: %w", err)
	}

	s.Out.Section("NFT NAME CONFIG")
	s.Out.KeyValue("Token ID", tokenID.String())
	if nftName == "" {
		s.Out.KeyValue("New Name", "(empty)")
	} else {
		s.Out.KeyValue("New Name", nftName)
	}

	s.Out.TxSubmitting("SetNftName", nil, ethtx.GasLimitAdminCall, s.AdjustedGasPrice())
	tx, err := nft.SetNftName(s.TransactOpts(nil, ethtx.GasLimitAdminCall), tokenID, nftName)
	return s.FinishTx(cmd.Context(), tx, err)
}
