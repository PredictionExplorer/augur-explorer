package cosmicgame

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// metadataCacheControl lets marketplaces and CDNs hold a token document for
// five minutes. The ETag added by the conditional-request middleware makes
// revalidation cheap, and it changes naturally whenever a token is renamed
// or its traits are ingested.
const metadataCacheControl = "public, max-age=300"

// TokenMetadata is the exported entry point for the bare ERC-721 tokenURI
// route (GET /metadata/{tokenID}), dispatched by host in the shared router
// constructor. On the Cosmic Signature host it serves Cosmic Signature
// metadata.
func (a *API) TokenMetadata(c *httpx.Context) {
	a.handleCstMetadata(c)
}

// GET /api/cosmicgame/cst/metadata/:tokenID — OpenSea-compatible metadata JSON (image hosted under /images/...).
// Uses the same token row as /cst/info. Image base defaults to this API's origin + /images; optional NFT_ASSETS_PUBLIC_BASE overrides.
//
// The document has two shapes. When the seed's art package has been ingested
// the response is enriched: the generator's attributes verbatim, its
// description, the reproducibility blocks and the full media list. Until
// then it is the minimal fallback, which every new mint uses for the minutes
// to hours the pipeline takes to publish. Neither shape carries the owner:
// ownership changes on every transfer while marketplaces cache aggressively,
// so a baked-in owner is permanently stale and ownerOf() is the truth.
func (a *API) handleCstMetadata(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "CosmicGame module or database not available")
		return
	}
	base := common.NFTImagePublicBase(c, a.assetsPublicBase)
	if base == "" {
		c.JSON(http.StatusInternalServerError, httpx.H{
			"error": "cannot derive public /images base URL (set Host or NFT_ASSETS_PUBLIC_BASE)",
		})
		return
	}
	p := c.Param("token_id")
	var tokenID int64
	// Token IDs start at 0, so only negatives are invalid.
	n, err := fmt.Sscanf(p, "%d", &tokenID)
	if err != nil || n != 1 || tokenID < 0 {
		common.RespondErrorJSON(c, "invalid token_id")
		return
	}
	ctx := c.Request.Context()
	tokenInfo, err := a.repo.CosmicSignatureTokenInfo(ctx, tokenID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			c.JSON(http.StatusNotFound, httpx.H{"error": "record not found"})
			return
		}
		a.respondStoreError(c, err)
		return
	}

	in := tokenMetadataInput{
		TokenID:       tokenID,
		RoundNum:      tokenInfo.RoundNum,
		TokenName:     strings.TrimSpace(tokenInfo.TokenName),
		MintTimestamp: tokenInfo.Tx.TimeStamp,
		Seed:          metadataSeed(strings.TrimSpace(tokenInfo.Seed)),
		AssetBase:     base,
		Allocation:    a.tokenAllocation(ctx, tokenID),
	}
	in.Traits = a.tokenTraits(ctx, in.Seed)

	c.Writer.Header().Set("Cache-Control", metadataCacheControl)
	c.JSON(http.StatusOK, buildTokenMetadata(in))
}

// tokenTraits loads the stored art contract for a seed, or nil when the
// package has not been ingested yet.
//
// The asset host is never contacted from the request path. A miss instead
// nudges the background ingester, which dedupes and honours the persisted
// backoff, so the first view of a freshly published package pulls it in
// within seconds without the handler ever waiting on the network.
func (a *API) tokenTraits(ctx context.Context, seed string) *tokenTraits {
	row, err := a.repo.TokenTraits(ctx, seed)
	if err != nil {
		if !errors.Is(err, store.ErrNotFound) && ctx.Err() == nil {
			a.logger.Error("loading token traits failed", "seed", seed, "error", err)
		}
		if a.traitsIngester != nil {
			a.traitsIngester.Nudge(seed)
		}
		return nil
	}
	traits, err := decodeTokenTraits(row)
	if err != nil {
		// A row that cannot be decoded is a storage bug, not a token
		// problem: serve the fallback rather than a half-built document.
		a.logger.Error("decoding stored token traits failed", "seed", seed, "error", err)
		return nil
	}
	return traits
}

// tokenAllocation resolves the prize path that minted a token into its
// display name, or the empty string when provenance is unavailable. A token
// that matches zero or several prize families is a data inconsistency; the
// attribute is omitted rather than guessed, and the inconsistency is logged.
func (a *API) tokenAllocation(ctx context.Context, tokenID int64) string {
	source, err := a.repo.CosmicSignatureMintSource(ctx, tokenID)
	if err != nil {
		if !errors.Is(err, store.ErrNotFound) && ctx.Err() == nil {
			a.logger.Warn("resolving token mint source failed", "token_id", tokenID, "error", err)
		}
		return ""
	}
	label := allocationLabel(source)
	if label == "" {
		a.logger.Warn("unmapped token mint source", "token_id", tokenID, "source", string(source))
	}
	return label
}

// traitNudger is the slice of the trait ingester the handler depends on: a
// non-blocking request to ingest one seed soon. Narrowing it to this one
// method keeps the handler testable and makes it impossible for the request
// path to acquire a blocking dependency on the asset host by accident.
type traitNudger interface {
	Nudge(seed string)
}
