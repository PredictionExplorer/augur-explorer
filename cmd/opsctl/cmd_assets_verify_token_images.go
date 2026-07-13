package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"

	opsassets "github.com/PredictionExplorer/augur-explorer/internal/ops/assets"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

const (
	// rwalkNFTAddr is the RandomWalk NFT contract whose minted-token count
	// bounds the scan.
	rwalkNFTAddr = "0x895a6F444BE4ba9d124F61DF736605792B35D66b"
	// rwalkImagesURL is the base URL of the RandomWalk image server.
	rwalkImagesURL = "https://api.randomwalknft.com:1443/images/randomwalk"
	// Token-image checks are independent; one stalled response must not block
	// the entire inventory indefinitely.
	rwalkImageRequestTimeout = 30 * time.Second
)

type assetsVerifyTokenImagesDeps struct {
	storeConfig func() store.Config
	openStore   func(context.Context, store.Config) (*store.Store, error)
	closeStore  func(*store.Store)
	newSource   func(*store.Store) opsassets.TokenCountSource
	client      opsassets.HTTPClient
	baseURL     string
	verify      func(context.Context, opsassets.VerifyTokenImagesOptions) (opsassets.ImageVerificationSummary, error)
}

func defaultAssetsVerifyTokenImagesDeps() assetsVerifyTokenImagesDeps {
	return assetsVerifyTokenImagesDeps{
		storeConfig: store.ConfigFromEnv,
		openStore:   store.New,
		closeStore: func(st *store.Store) {
			st.Close()
		},
		newSource: newRandomWalkTokenCountSource,
		client:    &http.Client{Timeout: rwalkImageRequestTimeout},
		baseURL:   rwalkImagesURL,
		verify:    opsassets.VerifyTokenImages,
	}
}

// newAssetsVerifyTokenImagesCmd builds `opsctl assets verify-token-images`,
// the replacement for the notibot verify-token-imgs script.
func newAssetsVerifyTokenImagesCmd() *cobra.Command {
	return newAssetsVerifyTokenImagesCmdWithDeps(defaultAssetsVerifyTokenImagesDeps())
}

func newAssetsVerifyTokenImagesCmdWithDeps(deps assetsVerifyTokenImagesDeps) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-token-images",
		Short: "Check every RandomWalk token image URL for HTTP 403 responses",
		Long: `Fetches the image of every minted RandomWalk token from the public image
server and reports tokens for which the server answers HTTP 403 (a known
RandomWalk webserver bug for early token ids).

The database connection is built from the PGSQL_* environment variables.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVerifyTokenImagesWithDeps(cmd, deps)
		},
	}
	return cmd
}

func runVerifyTokenImages(cmd *cobra.Command) error {
	return runVerifyTokenImagesWithDeps(cmd, defaultAssetsVerifyTokenImagesDeps())
}

func runVerifyTokenImagesWithDeps(cmd *cobra.Command, deps assetsVerifyTokenImagesDeps) error {
	ctx := cmd.Context()
	info := log.New(cmd.OutOrStdout(), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	st, err := deps.openStore(ctx, deps.storeConfig())
	if err != nil {
		return fmt.Errorf("failed to connect to storage: %w", err)
	}
	defer deps.closeStore(st)

	_, err = deps.verify(ctx, opsassets.VerifyTokenImagesOptions{
		Source:  deps.newSource(st),
		Client:  deps.client,
		BaseURL: deps.baseURL,
		Logger:  info,
	})
	return err
}

type randomWalkTokenCountSource struct {
	lookupAddressID func(context.Context, string) (int64, error)
	mintedTokens    func(context.Context, int64) (int64, error)
}

func newRandomWalkTokenCountSource(st *store.Store) opsassets.TokenCountSource {
	repo := rwstore.NewRepo(st)
	return randomWalkTokenCountSource{
		lookupAddressID: st.LookupAddressID,
		mintedTokens: func(ctx context.Context, addressID int64) (int64, error) {
			stats, err := repo.RandomWalkStats(ctx, addressID)
			if err != nil {
				return 0, err
			}
			return stats.TokensMinted, nil
		},
	}
}

func (s randomWalkTokenCountSource) MintedTokenCount(ctx context.Context) (int64, error) {
	rwalkAid, err := s.lookupAddressID(ctx, rwalkNFTAddr)
	if err != nil {
		return 0, fmt.Errorf("can't resolve RandomWalk contract address id: %w", err)
	}
	tokens, err := s.mintedTokens(ctx, rwalkAid)
	if err != nil {
		return 0, fmt.Errorf("can't read RandomWalk stats: %w", err)
	}
	return tokens, nil
}
