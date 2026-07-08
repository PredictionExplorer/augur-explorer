//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestCosmicSignatureTokens(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "cosmic_signature_nft_list", func() any {
		recs, err := r.CosmicSignatureTokens(ctx, 0, 100)
		if err != nil {
			t.Fatalf("CosmicSignatureTokens: %v", err)
		}
		return recs
	})
	golden(t, "cosmic_signature_nft_list_paged", func() any {
		recs, err := r.CosmicSignatureTokens(ctx, 3, 2)
		if err != nil {
			t.Fatalf("CosmicSignatureTokens(paged): %v", err)
		}
		return recs
	})
}

func TestCosmicSignatureTokenInfo(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// Token 1: alice's main-prize token, later named and CST-staked/unstaked.
	golden(t, "cosmic_signature_token_info_1", func() any {
		rec, err := r.CosmicSignatureTokenInfo(ctx, 1)
		if err != nil {
			t.Fatalf("CosmicSignatureTokenInfo(1): %v", err)
		}
		return rec
	})
	if _, err := r.CosmicSignatureTokenInfo(ctx, 999_999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("CosmicSignatureTokenInfo(999999) = %v, want ErrNotFound", err)
	}
}

func TestTokenNameHistory(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "cosmic_signature_token_name_history_1", func() any {
		recs, err := r.TokenNameHistory(ctx, 1)
		if err != nil {
			t.Fatalf("TokenNameHistory(1): %v", err)
		}
		return recs
	})
	got, err := r.TokenNameHistory(ctx, 2)
	if err != nil {
		t.Fatalf("TokenNameHistory(2): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no name history for token 2, got %d entries", len(got))
	}
}

func TestTokenOwnershipTransfers(t *testing.T) {
	r := repo(t)
	// Token 2: minted to dave, then transferred dave -> bob.
	golden(t, "cst_ownership_transfers_2", func() any {
		recs, err := r.TokenOwnershipTransfers(context.Background(), 2, 0, 100)
		if err != nil {
			t.Fatalf("TokenOwnershipTransfers(2): %v", err)
		}
		return recs
	})
}

func TestCosmicSignatureTokenDistribution(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_signature_token_distribution", func() any {
		recs, err := r.CosmicSignatureTokenDistribution(context.Background())
		if err != nil {
			t.Fatalf("CosmicSignatureTokenDistribution: %v", err)
		}
		return recs
	})
}

func TestNamedTokens(t *testing.T) {
	r := repo(t)
	golden(t, "named_tokens", func() any {
		recs, err := r.NamedTokens(context.Background())
		if err != nil {
			t.Fatalf("NamedTokens: %v", err)
		}
		return recs
	})
}

func TestSearchTokensByName(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// Token 1 is named "Genesis" by the fixture rename event; ILIKE makes
	// the lowercase substring match.
	golden(t, "search_tokens_by_name", func() any {
		recs, err := r.SearchTokensByName(ctx, "gene")
		if err != nil {
			t.Fatalf("SearchTokensByName(gene): %v", err)
		}
		return recs
	})
	got, err := r.SearchTokensByName(ctx, "no-token-has-this-name")
	if err != nil {
		t.Fatalf("SearchTokensByName(miss): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no search results, got %d", len(got))
	}
}

func TestCosmicSignatureTokenCount(t *testing.T) {
	r := repo(t)
	golden(t, "erc721_token_total", func() any {
		total, err := r.CosmicSignatureTokenCount(context.Background())
		if err != nil {
			t.Fatalf("CosmicSignatureTokenCount: %v", err)
		}
		return total
	})
}

func TestCosmicSignatureTokenSeed(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	const want = "seed0000000000000000000000000000000000000000000000000000000001"
	got, err := r.CosmicSignatureTokenSeed(ctx, 1)
	if err != nil {
		t.Fatalf("CosmicSignatureTokenSeed(1): %v", err)
	}
	if got != want {
		t.Errorf("token 1 seed: got %q, want %q", got, want)
	}
	if _, err := r.CosmicSignatureTokenSeed(ctx, 999_999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("CosmicSignatureTokenSeed(999999) = %v, want ErrNotFound", err)
	}
}
