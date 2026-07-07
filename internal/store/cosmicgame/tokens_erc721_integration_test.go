//go:build integration

package cosmicgame

import "testing"

func TestGetCosmicSignatureNFTList(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_signature_nft_list", func() any {
		return sw.Get_cosmic_signature_nft_list(0, 100)
	})
	golden(t, "cosmic_signature_nft_list_paged", func() any {
		return sw.Get_cosmic_signature_nft_list(3, 2)
	})
}

func TestGetCosmicSignatureTokenInfo(t *testing.T) {
	sw := store(t)
	// Token 1: alice's main-prize token, later named and CST-staked/unstaked.
	golden(t, "cosmic_signature_token_info_1", func() any {
		found, rec := sw.Get_cosmic_signature_token_info(1)
		if !found {
			t.Fatal("expected token 1 to exist")
		}
		return rec
	})
	if found, _ := sw.Get_cosmic_signature_token_info(999_999); found {
		t.Error("expected token 999999 to be missing")
	}
}

func TestGetCosmicSignatureTokenNameHistory(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_signature_token_name_history_1", func() any {
		return sw.Get_cosmic_signature_token_name_history(1)
	})
	if got := sw.Get_cosmic_signature_token_name_history(2); len(got) != 0 {
		t.Errorf("expected no name history for token 2, got %d entries", len(got))
	}
}

func TestGetCstOwnershipTransfers(t *testing.T) {
	sw := store(t)
	// Token 2: minted to dave, then transferred dave -> bob.
	golden(t, "cst_ownership_transfers_2", func() any {
		return sw.Get_cst_ownership_transfers(2, 0, 100)
	})
}

func TestGetCosmicSignatureTokenDistribution(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_signature_token_distribution", func() any {
		return sw.Get_cosmic_signature_token_distribution()
	})
}

func TestGetNamedTokens(t *testing.T) {
	sw := store(t)
	golden(t, "named_tokens", func() any {
		return sw.Get_named_tokens()
	})
}

func TestGetERC721TokenTotal(t *testing.T) {
	sw := store(t)
	golden(t, "erc721_token_total", func() any {
		return sw.Get_erc721_token_total()
	})
}

func TestGetERC721TokenSeed(t *testing.T) {
	sw := store(t)
	const want = "seed0000000000000000000000000000000000000000000000000000000001"
	if got := sw.Get_erc721_token_seed(1); got != want {
		t.Errorf("token 1 seed: got %q, want %q", got, want)
	}
	if got := sw.Get_erc721_token_seed(999_999); got != "" {
		t.Errorf("missing token seed: got %q, want empty", got)
	}
}
