package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// defaultDevRPCURL is the RPC endpoint used when RPC_URL is not set: the
// standard local Hardhat node.
const defaultDevRPCURL = "http://127.0.0.1:8545"

// rwalkSamplesEnvHelp documents the environment of the rwalk_samples dev
// subcommand. Unlike the other transaction subcommands, RPC_URL is optional
// and defaults to the local Hardhat node.
const rwalkSamplesEnvHelp = `Environment:
  RPC_URL               Ethereum RPC endpoint (default ` + defaultDevRPCURL + `)
  PKEY_HEX              64-char hex private key, no 0x prefix (required)
  GAS_PRICE_MULTIPLIER  multiplier applied to the suggested gas price (default 2.0)`

// newRwalkSamplesCmd builds the rwalk_samples subcommand, a development
// helper that mints a batch of sample RandomWalk NFTs on a local chain so
// CosmicSignature features that consume RandomWalk tokens (RandomWalk bids,
// RandomWalk staking) have material to work with.
func newRwalkSamplesCmd() *cobra.Command {
	var verbose bool
	var count int
	c := &cobra.Command{
		Use:   "rwalk_samples [rwalk_addr]",
		Short: "Mint sample RandomWalk NFTs for development (default 10)",
		Long: "Mints a batch of sample RandomWalk NFTs to the signer's address, re-reading\n" +
			"the mint price before each mint (the price rises after every mint). This is a\n" +
			"development helper for populating a local CosmicSignature chain.\n\n" + rwalkSamplesEnvHelp,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRwalkSamples(cmd, verbose, count, args[0])
		},
	}
	c.Flags().IntVarP(&count, "count", "n", 10, "number of tokens to mint")
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newRwalkSamplesCmd()) }

// newDevTxSession is newTxSession with a development default for RPC_URL:
// when the variable is unset it connects to the local Hardhat node instead of
// failing. PKEY_HEX is still required.
func newDevTxSession(cmd *cobra.Command, verbose bool) (*ethtx.Session, error) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		rpcURL = defaultDevRPCURL
	}
	pkey, err := pkeyHexFromEnv()
	if err != nil {
		return nil, err
	}
	multiplier, err := gasMultiplierFromEnv()
	if err != nil {
		return nil, err
	}
	return ethtx.New(cmd.Context(), ethtx.Options{
		RPCURL:             rpcURL,
		PrivateKeyHex:      pkey,
		Verbose:            verbose,
		Out:                cmd.OutOrStdout(),
		GasPriceMultiplier: multiplier,
	})
}

func runRwalkSamples(cmd *cobra.Command, verbose bool, count int, addrArg string) error {
	if count < 1 {
		return fmt.Errorf("--count must be at least 1 (got %d)", count)
	}
	rwalkAddr, err := parseAddress("rwalk_addr", addrArg)
	if err != nil {
		return err
	}

	s, err := newDevTxSession(cmd, verbose)
	if err != nil {
		return err
	}
	rwalk, err := rwcontracts.NewRWalk(rwalkAddr, s.Net.Client)
	if err != nil {
		return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
	}
	s.Out.ContractInfo("RandomWalk Address", rwalkAddr)

	out := cmd.OutOrStdout()
	tokenIDs := make([]string, 0, count)
	for i := 1; i <= count; i++ {
		if i > 1 {
			// Pick up the incremented nonce and current gas price before
			// building the next transaction.
			if err = s.Refresh(cmd.Context()); err != nil {
				return err
			}
		}
		price, errPrice := rwalk.GetMintPrice(ethtx.CallOpts())
		if errPrice != nil {
			return fmt.Errorf("error at GetMintPrice(): %w", errPrice)
		}

		gasLimit := ethtx.GasLimitHighComplexity
		totalNeeded := new(big.Int).Mul(s.AdjustedGasPrice(), big.NewInt(int64(gasLimit)))
		totalNeeded.Add(totalNeeded, price)
		if s.Acc.Balance.Cmp(totalNeeded) < 0 {
			return fmt.Errorf("mint %d/%d: insufficient balance: need %s ETH (price + gas), have %s ETH",
				i, count, ethtx.WeiToEthText(totalNeeded), ethtx.WeiToEthText(s.Acc.Balance))
		}

		tx, errMint := rwalk.Mint(s.TransactOpts(price, gasLimit))
		if errMint != nil {
			return fmt.Errorf("mint %d/%d failed to send: %w", i, count, errMint)
		}
		receipt, errWait := s.WaitForReceipt(cmd.Context(), tx)
		if errWait != nil {
			return fmt.Errorf("mint %d/%d: receipt not received for tx %s: %w", i, count, tx.Hash(), errWait)
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("mint %d/%d: transaction %s reverted on-chain", i, count, tx.Hash())
		}

		tokenID := mintedTokenID(rwalk, rwalkAddr, receipt)
		tokenIDs = append(tokenIDs, tokenID)
		fmt.Fprintf(out, "Minted %d/%d: token %s (price %s wei, tx %s)\n",
			i, count, tokenID, price.String(), tx.Hash().String())
	}
	fmt.Fprintf(out, "Done. Minted %d RandomWalk tokens: %s\n", count, strings.Join(tokenIDs, ","))
	return nil
}

// mintedTokenID extracts the token ID from the MintEvent log of a successful
// mint receipt. A mint always emits the event; "?" is returned only if the
// logs cannot be parsed (e.g. an unexpected contract at the address).
func mintedTokenID(rwalk *rwcontracts.RWalk, rwalkAddr common.Address, receipt *types.Receipt) string {
	for _, lg := range receipt.Logs {
		if lg.Address != rwalkAddr {
			continue
		}
		ev, err := rwalk.ParseMintEvent(*lg)
		if err == nil {
			return ev.TokenId.String()
		}
	}
	return "?"
}
