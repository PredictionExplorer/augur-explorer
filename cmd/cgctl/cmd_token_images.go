package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// The total-tokens and token-seed subcommands replace the old img_upload
// helper binary used by imgcheck.sh: they read CosmicSignature token data
// from the database so image/video artifacts can be checked and regenerated.

func init() {
	register(&cobra.Command{
		Use:   "total-tokens",
		Short: "Print the total number of CosmicSignature tokens (from the DB)",
		Long: `Print the total number of CosmicSignature ERC-721 tokens recorded in the
database. The value is printed without a trailing newline so shell scripts can
capture it directly.

Environment:
  PGSQL_*  PostgreSQL connection (PGSQL_HOST, PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD)`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			repo, err := connectTokenRepo(cmd.Context())
			if err != nil {
				return err
			}
			total, err := repo.CosmicSignatureTokenCount(cmd.Context())
			if err != nil {
				return err
			}
			fmt.Printf("%v", total)
			return nil
		},
	})

	register(&cobra.Command{
		Use:   "token-seed <token-id>",
		Short: "Print the seed of a CosmicSignature token (from the DB)",
		Long: `Print the seed of the given CosmicSignature ERC-721 token as recorded in
the database. The value is printed without a trailing newline so shell scripts
can capture it directly.

Environment:
  PGSQL_*  PostgreSQL connection (PGSQL_HOST, PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD)`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			tokenID, err := parseInt64("token-id", args[0])
			if err != nil {
				return err
			}
			repo, err := connectTokenRepo(cmd.Context())
			if err != nil {
				return err
			}
			seed, err := repo.CosmicSignatureTokenSeed(cmd.Context(), tokenID)
			if err != nil {
				// The legacy helper printed an empty string for unknown
				// token ids; imgcheck.sh relies on that contract.
				if errors.Is(err, store.ErrNotFound) {
					fmt.Printf("%v", "")
					return nil
				}
				return err
			}
			fmt.Printf("%v", seed)
			return nil
		},
	})
}

// connectTokenRepo connects to PostgreSQL (PGSQL_* env vars) and returns the
// CosmicGame query repo. The pool lives for the remainder of the process
// (cgctl runs one command and exits).
func connectTokenRepo(ctx context.Context) (*cgstore.Repo, error) {
	st, err := store.New(ctx, store.ConfigFromEnv())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to storage: %w", err)
	}
	return cgstore.NewRepo(st), nil
}
