package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// newVerifyOwnerCmd builds the verify-owner subcommand (legacy verify_owner
// script).
func newVerifyOwnerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "verify-owner",
		Short: "Verify each token's on-chain owner matches the database",
		Long: "Verifies the owner of each token against the DB by querying directly to the RPC node.\n\n" +
			"Environment variables:\n  RPC_URL  Ethereum RPC endpoint (required)\n  PGSQL_*  PostgreSQL connection (required)",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			repo, baseStore, err := connectRWStorage(newInfoLogger())
			if err != nil {
				return err
			}
			caddrs, err := repo.ContractAddrs(ctx)
			if err != nil {
				return fmt.Errorf("resolving contract addresses: %w", err)
			}

			eclient, err := dialEthClient()
			if err != nil {
				return err
			}
			rwalkAddr := common.HexToAddress(caddrs.RandomWalk)
			rwalk, err := rwcontracts.NewRWalk(rwalkAddr, eclient)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWalk contract: %w", err)
			}
			copts := callOpts()
			numToksBig, err := rwalk.NextTokenId(copts)
			if err != nil {
				return fmt.Errorf("error getting num tokens: %w", err)
			}
			numToks := numToksBig.Int64()

			rwalkAid, err := baseStore.LookupAddressID(ctx, rwalkAddr.String())
			if err != nil {
				return fmt.Errorf("error looking up contract address id: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "num tokens: %v\n", numToks)

			stats, err := repo.RandomWalkStats(ctx, rwalkAid)
			if err != nil {
				return fmt.Errorf("error getting contract stats: %w", err)
			}
			if stats.TokensMinted != numToks {
				fmt.Fprintf(cmd.OutOrStdout(),
					"Error: num tokens doesn't match: real num tokens = %v, db num tokens = %v\n",
					numToks, stats.TokensMinted)

			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "Num tokens in database is set correctly (%v tokens)\n", numToks)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Starting verification process, will loop %v times\n", numToks)
			for i := int64(0); i < numToks; i++ {
				chainOwnerAddr, err := rwalk.OwnerOf(copts, big.NewInt(i))
				if err != nil {
					return fmt.Errorf("error during Owner() call: %w", err)
				}
				chainOwnerAid, err := baseStore.LookupAddressID(ctx, chainOwnerAddr.String())
				if err != nil {
					return fmt.Errorf("error during addr lookup: %w", err)
				}
				tokInfo, err := repo.TokenInfo(ctx, rwalkAid, i)
				if err != nil {
					return fmt.Errorf("error getting token info from db: %w", err)
				}
				if tokInfo.CurOwnerAid != chainOwnerAid {
					fmt.Fprintf(cmd.OutOrStdout(),
						"DB invalid: token_id=%v; owner mismatch, real owner %v, owner in db %v\n",
						i, chainOwnerAddr.String(), tokInfo.CurOwnerAddr)

				}
			}
			return nil
		},
	}
}

func init() { register(newVerifyOwnerCmd()) }
