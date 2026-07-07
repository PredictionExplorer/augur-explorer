//go:build integration

package randomwalk

import (
	"testing"
)

func TestGetActiveOffers(t *testing.T) {
	sw := store(t)
	// Offer 1 was bought and offer 4 cancelled; 2 and 3 remain active.
	golden(t, "active_offers_default_order", func() any {
		return sw.Get_active_offers(aidRandomWalk, aidMarketplace, 0)
	})
	golden(t, "active_offers_price_desc", func() any {
		return sw.Get_active_offers(aidRandomWalk, aidMarketplace, 1)
	})
	golden(t, "active_offers_price_asc", func() any {
		return sw.Get_active_offers(aidRandomWalk, aidMarketplace, 2)
	})
}

func TestGetMintedTokensByPeriod(t *testing.T) {
	sw := store(t)
	golden(t, "minted_tokens_by_period", func() any {
		return sw.Get_minted_tokens_by_period(aidRandomWalk, 1767228600, 1767229000)
	})
	if got := sw.Get_minted_tokens_by_period(aidRandomWalk, 100, 200); len(got) != 0 {
		t.Errorf("expected no mints in 1970, got %d", len(got))
	}
}

func TestGetMintedTokensSequentially(t *testing.T) {
	sw := store(t)
	golden(t, "minted_tokens_sequentially", func() any {
		return sw.Get_minted_tokens_sequentially(aidRandomWalk, 0, 100)
	})
	golden(t, "minted_tokens_sequentially_paged", func() any {
		return sw.Get_minted_tokens_sequentially(aidRandomWalk, 1, 2)
	})
}

func TestGetTradingHistory(t *testing.T) {
	sw := store(t)
	golden(t, "trading_history", func() any {
		return sw.Get_trading_history(aidMarketplace, 0, 100)
	})
}

func TestGetRandomWalkStats(t *testing.T) {
	sw := store(t)
	golden(t, "random_walk_stats", func() any {
		return sw.Get_random_walk_stats(aidRandomWalk)
	})
}

func TestGetMarketStats(t *testing.T) {
	sw := store(t)
	golden(t, "market_stats", func() any {
		return sw.Get_market_stats(aidMarketplace)
	})
}

func TestGetTokenFullHistory(t *testing.T) {
	sw := store(t)
	// Token 10: mint, name, offer, sale, relist, cancel, withdrawal.
	golden(t, "token_full_history_10", func() any {
		return sw.Get_token_full_history(aidRandomWalk, 10, 0, 100)
	})
}

func TestGetMarketTradingVolumeByPeriod(t *testing.T) {
	sw := store(t)
	golden(t, "market_trading_volume_by_period", func() any {
		return sw.Get_market_trading_volume_by_period(aidMarketplace, 1767229100, 1767229400, 300)
	})
}

func TestGetNameChangesForToken(t *testing.T) {
	sw := store(t)
	golden(t, "name_changes_for_token_10", func() any {
		return sw.Get_name_changes_for_token(10)
	})
	if got := sw.Get_name_changes_for_token(11); len(got) != 0 {
		t.Errorf("expected no name changes for token 11, got %d", len(got))
	}
}

func TestGetRandomWalkTokensByUser(t *testing.T) {
	sw := store(t)
	// dave minted #11 and bought #10.
	golden(t, "random_walk_tokens_by_user_dave", func() any {
		return sw.Get_random_walk_tokens_by_user(aidDave)
	})
}

func TestGetFloorPrice(t *testing.T) {
	sw := store(t)
	noOffers, floorPrice, offerID, tokenID, err := sw.Get_floor_price(aidRandomWalk, aidMarketplace)
	if err != nil {
		t.Fatalf("Get_floor_price: %v", err)
	}
	if noOffers {
		t.Fatal("expected active offers in the fixture set")
	}
	// Cheapest active sell offer is dave's #11 at 2 ETH.
	if floorPrice != 2.0 || offerID != 2 || tokenID != 11 {
		t.Errorf("floor price: got (%v, offer %d, token %d), want (2, 2, 11)", floorPrice, offerID, tokenID)
	}
}

func TestGetTradingHistoryByUser(t *testing.T) {
	sw := store(t)
	golden(t, "trading_history_by_user_carol", func() any {
		return sw.Get_trading_history_by_user(aidCarol)
	})
}

func TestGetRwalkUserInfo(t *testing.T) {
	sw := store(t)
	golden(t, "rwalk_user_info_carol", func() any {
		info, err := sw.Get_rwalk_user_info(aidCarol, aidRandomWalk)
		if err != nil {
			t.Fatalf("Get_rwalk_user_info(carol): %v", err)
		}
		return info
	})
	// A user with no RandomWalk activity yields ErrNoRows, not a crash.
	if _, err := sw.Get_rwalk_user_info(aidAlice+979, aidRandomWalk); err == nil {
		t.Error("expected error for user without rw_user_stats row")
	}
}

func TestGetTop5TradedTokens(t *testing.T) {
	sw := store(t)
	golden(t, "top5_traded_tokens", func() any {
		return sw.Get_top5_traded_tokens()
	})
}

func TestGetRwalkTokenInfo(t *testing.T) {
	sw := store(t)
	golden(t, "rwalk_token_info_10", func() any {
		info, err := sw.Get_rwalk_token_info(aidRandomWalk, 10)
		if err != nil {
			t.Fatalf("Get_rwalk_token_info(10): %v", err)
		}
		return info
	})
	if _, err := sw.Get_rwalk_token_info(aidRandomWalk, 999); err == nil {
		t.Error("expected error for missing token 999")
	}
}

func TestCheckRwalkTokenExists(t *testing.T) {
	sw := store(t)
	exists, err := sw.Check_rwalk_token_exists(10)
	if err != nil {
		t.Fatalf("Check_rwalk_token_exists(10): %v", err)
	}
	if !exists {
		t.Error("expected token 10 to exist")
	}
	exists, err = sw.Check_rwalk_token_exists(999)
	if err != nil {
		t.Fatalf("Check_rwalk_token_exists(999): %v", err)
	}
	if exists {
		t.Error("expected token 999 to be missing")
	}
}

func TestGetRwalkMintIntervals(t *testing.T) {
	sw := store(t)
	golden(t, "rwalk_mint_intervals", func() any {
		return sw.Get_rwalk_mint_intervals(aidRandomWalk)
	})
}

func TestGetRwalkWithdrawalChart(t *testing.T) {
	sw := store(t)
	golden(t, "rwalk_withdrawal_chart", func() any {
		return sw.Get_rwalk_withdrawal_chart(aidRandomWalk)
	})
}

func TestGetSaleHistory(t *testing.T) {
	sw := store(t)
	golden(t, "sale_history", func() any {
		return sw.Get_sale_history(aidMarketplace, 0, 100)
	})
}

func TestGetRwalkFloorPriceForPeriods(t *testing.T) {
	sw := store(t)
	golden(t, "rwalk_floor_price_for_periods", func() any {
		return sw.Get_rwalk_floor_price_for_periods(aidRandomWalk, aidMarketplace, 1767229100, 1767229700, 300)
	})
}

func TestGetMintedTokensForCSV(t *testing.T) {
	sw := store(t)
	golden(t, "minted_tokens_for_csv", func() any {
		return sw.Get_minted_tokens_for_CSV(aidRandomWalk)
	})
}

func TestGetMintReport(t *testing.T) {
	sw := store(t)
	golden(t, "mint_report", func() any {
		return sw.Get_mint_report()
	})
}
