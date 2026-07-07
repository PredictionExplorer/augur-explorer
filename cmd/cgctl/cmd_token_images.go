package main

import (
	"context"
	"fmt"
	"log"
	"os"

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
			sw, err := connectTokenStorage()
			if err != nil {
				return err
			}
			fmt.Printf("%v", sw.Get_erc721_token_total())
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
			sw, err := connectTokenStorage()
			if err != nil {
				return err
			}
			fmt.Printf("%v", sw.Get_erc721_token_seed(tokenID))
			return nil
		},
	})
}

// connectTokenStorage connects to PostgreSQL (PGSQL_* env vars) and wraps the
// connection in the cosmicgame storage helper. The pool lives for the
// remainder of the process (cgctl runs one command and exits).
func connectTokenStorage() (*cgstore.SQLStorageWrapper, error) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	st, err := store.New(context.Background(), store.ConfigFromEnv())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to storage: %w", err)
	}
	storage := store.NewSQLStorageFromDB(st.DB(), logger)
	storage.Db_set_schema_name("public")
	return &cgstore.SQLStorageWrapper{S: storage}, nil
}
