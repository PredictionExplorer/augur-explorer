package marketplace

import (
	"context"
	"testing"
)

func TestLegacyOrderClauseWhitelist(t *testing.T) {
	t.Parallel()
	cases := map[int]string{
		-1: " ORDER BY o.id",
		0:  " ORDER BY o.id",
		1:  " ORDER BY o.price DESC",
		2:  " ORDER BY o.price ASC",
		3:  " ORDER BY o.id",
	}
	for input, want := range cases {
		if got := legacyOrderClause(input); got != want {
			t.Errorf("legacyOrderClause(%d) = %q, want %q", input, got, want)
		}
	}
}

func TestMarketplaceReadsRejectInvalidArgumentsBeforeQuery(t *testing.T) {
	t.Parallel()
	r := NewRepo(nil)
	ctx := context.Background()
	validScope := Scope{MarketplaceAid: 12, CollectionAid: 3}

	calls := map[string]func() error{
		"legacy offers scope": func() error {
			_, err := r.ActiveOffersLegacy(ctx, Scope{}, 0)
			return err
		},
		"legacy sales scope": func() error {
			_, err := r.SaleHistoryLegacy(ctx, Scope{}, 0, 1)
			return err
		},
		"legacy sales offset": func() error {
			_, err := r.SaleHistoryLegacy(ctx, validScope, -1, 1)
			return err
		},
		"legacy floor scope": func() error {
			_, _, _, _, err := r.FloorPriceETH(ctx, Scope{})
			return err
		},
		"offers zero limit": func() error {
			_, _, err := r.ActiveOffersPage(ctx, validScope, OfferSortNewest, nil, 0)
			return err
		},
		"offers bad sort": func() error {
			_, _, err := r.ActiveOffersPage(ctx, validScope, OfferSort("bad"), nil, 1)
			return err
		},
		"offers cursor missing price": func() error {
			_, _, err := r.ActiveOffersPage(
				ctx,
				validScope,
				OfferSortPriceAsc,
				&OfferPageCursor{EventLogID: 1},
				1,
			)
			return err
		},
		"history zero cursor": func() error {
			_, _, err := r.OfferHistoryPage(ctx, validScope, &EventPageCursor{}, 1)
			return err
		},
		"trades zero limit": func() error {
			_, _, err := r.TradesPage(ctx, validScope, nil, 0)
			return err
		},
		"exact floor scope": func() error {
			_, err := r.FloorPrice(ctx, Scope{})
			return err
		},
	}
	for name, call := range calls {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := call(); err == nil {
				t.Fatal("invalid arguments accepted")
			}
		})
	}
}
