//go:build integration

package randomwalk

import (
	"context"
	"errors"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func TestActiveOffers(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	fetch := func(orderBy int) func() any {
		return func() any {
			offers, err := r.ActiveOffers(ctx, aidRandomWalk, aidMarketplace, orderBy)
			if err != nil {
				t.Fatalf("ActiveOffers(order %d): %v", orderBy, err)
			}
			return offers
		}
	}
	// Offer 1 was bought and offer 4 cancelled; 2 and 3 remain active.
	golden(t, "active_offers_default_order", fetch(0))
	golden(t, "active_offers_price_desc", fetch(1))
	golden(t, "active_offers_price_asc", fetch(2))
}

func TestMintedTokensByPeriod(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "minted_tokens_by_period", func() any {
		toks, err := r.MintedTokensByPeriod(ctx, aidRandomWalk, 1767228600, 1767229000)
		if err != nil {
			t.Fatalf("MintedTokensByPeriod: %v", err)
		}
		return toks
	})
	got, err := r.MintedTokensByPeriod(ctx, aidRandomWalk, 100, 200)
	if err != nil {
		t.Fatalf("MintedTokensByPeriod(1970): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no mints in 1970, got %d", len(got))
	}
}

func TestMintedTokensSequentially(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "minted_tokens_sequentially", func() any {
		toks, err := r.MintedTokensSequentially(ctx, aidRandomWalk, 0, 100)
		if err != nil {
			t.Fatalf("MintedTokensSequentially: %v", err)
		}
		return toks
	})
	golden(t, "minted_tokens_sequentially_paged", func() any {
		toks, err := r.MintedTokensSequentially(ctx, aidRandomWalk, 1, 2)
		if err != nil {
			t.Fatalf("MintedTokensSequentially(paged): %v", err)
		}
		return toks
	})
}

func TestTradingHistory(t *testing.T) {
	r := repo(t)
	golden(t, "trading_history", func() any {
		recs, err := r.TradingHistory(context.Background(), aidMarketplace, 0, 100)
		if err != nil {
			t.Fatalf("TradingHistory: %v", err)
		}
		return recs
	})
}

func TestRandomWalkStats(t *testing.T) {
	r := repo(t)
	golden(t, "random_walk_stats", func() any {
		stats, err := r.RandomWalkStats(context.Background(), aidRandomWalk)
		if err != nil {
			t.Fatalf("RandomWalkStats: %v", err)
		}
		return stats
	})
}

func TestMarketStats(t *testing.T) {
	r := repo(t)
	golden(t, "market_stats", func() any {
		stats, err := r.MarketStats(context.Background(), aidMarketplace)
		if err != nil {
			t.Fatalf("MarketStats: %v", err)
		}
		return stats
	})
}

func TestTokenFullHistory(t *testing.T) {
	r := repo(t)
	// Token 10: mint, name, offer, sale, relist, cancel, withdrawal.
	golden(t, "token_full_history_10", func() any {
		recs, err := r.TokenFullHistory(context.Background(), aidRandomWalk, 10, 0, 100)
		if err != nil {
			t.Fatalf("TokenFullHistory: %v", err)
		}
		return recs
	})
}

func TestMarketTradingVolumeByPeriod(t *testing.T) {
	r := repo(t)
	golden(t, "market_trading_volume_by_period", func() any {
		recs, err := r.MarketTradingVolumeByPeriod(context.Background(), aidMarketplace, 1767229100, 1767229400, 300)
		if err != nil {
			t.Fatalf("MarketTradingVolumeByPeriod: %v", err)
		}
		return recs
	})
}

func TestTokenNameChanges(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "name_changes_for_token_10", func() any {
		recs, err := r.TokenNameChanges(ctx, 10)
		if err != nil {
			t.Fatalf("TokenNameChanges: %v", err)
		}
		return recs
	})
	got, err := r.TokenNameChanges(ctx, 11)
	if err != nil {
		t.Fatalf("TokenNameChanges(11): %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected no name changes for token 11, got %d", len(got))
	}
}

func TestTokensByUser(t *testing.T) {
	r := repo(t)
	// dave minted #11 and bought #10.
	golden(t, "random_walk_tokens_by_user_dave", func() any {
		toks, err := r.TokensByUser(context.Background(), aidDave)
		if err != nil {
			t.Fatalf("TokensByUser: %v", err)
		}
		return toks
	})
}

func TestFloorPrice(t *testing.T) {
	r := repo(t)
	noOffers, floorPrice, offerID, tokenID, err := r.FloorPrice(context.Background(), aidRandomWalk, aidMarketplace)
	if err != nil {
		t.Fatalf("FloorPrice: %v", err)
	}
	if noOffers {
		t.Fatal("expected active offers in the fixture set")
	}
	// Cheapest active sell offer is dave's #11 at 2 ETH.
	if floorPrice != 2.0 || offerID != 2 || tokenID != 11 {
		t.Errorf("floor price: got (%v, offer %d, token %d), want (2, 2, 11)", floorPrice, offerID, tokenID)
	}
}

func TestTradingHistoryByUser(t *testing.T) {
	r := repo(t)
	golden(t, "trading_history_by_user_carol", func() any {
		recs, err := r.TradingHistoryByUser(context.Background(), aidCarol)
		if err != nil {
			t.Fatalf("TradingHistoryByUser: %v", err)
		}
		return recs
	})
}

func TestUserInfo(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "rwalk_user_info_carol", func() any {
		info, err := r.UserInfo(ctx, aidCarol, aidRandomWalk)
		if err != nil {
			t.Fatalf("UserInfo(carol): %v", err)
		}
		return info
	})
	// A user with no RandomWalk activity yields ErrNotFound, not a crash.
	if _, err := r.UserInfo(ctx, aidAlice+979, aidRandomWalk); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("expected ErrNotFound for user without rw_user_stats row, got %v", err)
	}
}

func TestTop5TradedTokens(t *testing.T) {
	r := repo(t)
	golden(t, "top5_traded_tokens", func() any {
		toks, err := r.Top5TradedTokens(context.Background())
		if err != nil {
			t.Fatalf("Top5TradedTokens: %v", err)
		}
		return toks
	})
}

func TestTokenInfo(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "rwalk_token_info_10", func() any {
		info, err := r.TokenInfo(ctx, aidRandomWalk, 10)
		if err != nil {
			t.Fatalf("TokenInfo(10): %v", err)
		}
		return info
	})
	if _, err := r.TokenInfo(ctx, aidRandomWalk, 999); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("expected ErrNotFound for missing token 999, got %v", err)
	}
}

func TestTokenMinted(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	exists, err := r.TokenMinted(ctx, 10)
	if err != nil {
		t.Fatalf("TokenMinted(10): %v", err)
	}
	if !exists {
		t.Error("expected token 10 to exist")
	}
	exists, err = r.TokenMinted(ctx, 999)
	if err != nil {
		t.Fatalf("TokenMinted(999): %v", err)
	}
	if exists {
		t.Error("expected token 999 to be missing")
	}
}

func TestMintIntervals(t *testing.T) {
	r := repo(t)
	golden(t, "rwalk_mint_intervals", func() any {
		recs, err := r.MintIntervals(context.Background(), aidRandomWalk)
		if err != nil {
			t.Fatalf("MintIntervals: %v", err)
		}
		return recs
	})
}

func TestWithdrawalChart(t *testing.T) {
	r := repo(t)
	golden(t, "rwalk_withdrawal_chart", func() any {
		recs, err := r.WithdrawalChart(context.Background(), aidRandomWalk)
		if err != nil {
			t.Fatalf("WithdrawalChart: %v", err)
		}
		return recs
	})
}

func TestSaleHistory(t *testing.T) {
	r := repo(t)
	golden(t, "sale_history", func() any {
		recs, err := r.SaleHistory(context.Background(), aidMarketplace, 0, 100)
		if err != nil {
			t.Fatalf("SaleHistory: %v", err)
		}
		return recs
	})
}

func TestFloorPriceByPeriod(t *testing.T) {
	r := repo(t)
	golden(t, "rwalk_floor_price_for_periods", func() any {
		recs, err := r.FloorPriceByPeriod(context.Background(), aidRandomWalk, aidMarketplace, 1767229100, 1767229700, 300)
		if err != nil {
			t.Fatalf("FloorPriceByPeriod: %v", err)
		}
		return recs
	})
}

func TestMintedTokensCSV(t *testing.T) {
	r := repo(t)
	golden(t, "minted_tokens_for_csv", func() any {
		recs, err := r.MintedTokensCSV(context.Background(), aidRandomWalk)
		if err != nil {
			t.Fatalf("MintedTokensCSV: %v", err)
		}
		return recs
	})
}

func TestMintReport(t *testing.T) {
	r := repo(t)
	golden(t, "mint_report", func() any {
		recs, err := r.MintReport(context.Background())
		if err != nil {
			t.Fatalf("MintReport: %v", err)
		}
		return recs
	})
}
