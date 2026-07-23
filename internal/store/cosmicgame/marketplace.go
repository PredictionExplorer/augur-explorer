package cosmicgame

import (
	"context"

	marketstore "github.com/PredictionExplorer/augur-explorer/internal/store/marketplace"
)

// CosmicSignatureMarketplaceScope resolves the marketplace and Cosmic
// Signature collection addresses from the deployment registries into the
// database-local IDs used by the shared marketplace tables.
func (r *Repo) CosmicSignatureMarketplaceScope(
	ctx context.Context,
) (marketstore.Scope, error) {
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return marketstore.Scope{}, err
	}
	return r.marketplace.ResolveScope(
		ctx,
		addrs.MarketplaceAddr,
		addrs.CosmicSignatureAddr,
	)
}

// CosmicSignatureMarketplaceActiveOffers returns the v1 active-offer shape
// for Cosmic Signature NFTs.
func (r *Repo) CosmicSignatureMarketplaceActiveOffers(
	ctx context.Context,
	orderBy int,
) ([]marketstore.LegacyOffer, marketstore.Scope, error) {
	scope, err := r.CosmicSignatureMarketplaceScope(ctx)
	if err != nil {
		return nil, marketstore.Scope{}, err
	}
	records, err := r.marketplace.ActiveOffersLegacy(ctx, scope, orderBy)
	return records, scope, err
}

// CosmicSignatureMarketplaceFloorPriceETH returns the v1 floating-point ETH
// floor-price view for the Cosmic Signature collection.
func (r *Repo) CosmicSignatureMarketplaceFloorPriceETH(
	ctx context.Context,
) (noOffers bool, floorPrice float64, offerID int64, tokenID int64, scope marketstore.Scope, err error) {
	scope, err = r.CosmicSignatureMarketplaceScope(ctx)
	if err != nil {
		return false, 0, 0, 0, marketstore.Scope{}, err
	}
	noOffers, floorPrice, offerID, tokenID, err = r.marketplace.FloorPriceETH(ctx, scope)
	return noOffers, floorPrice, offerID, tokenID, scope, err
}

// CosmicSignatureMarketplaceSales returns the v1 completed-sale shape for
// Cosmic Signature NFTs.
func (r *Repo) CosmicSignatureMarketplaceSales(
	ctx context.Context,
	offset int,
	limit int,
) ([]marketstore.LegacyOffer, marketstore.Scope, error) {
	scope, err := r.CosmicSignatureMarketplaceScope(ctx)
	if err != nil {
		return nil, marketstore.Scope{}, err
	}
	records, err := r.marketplace.SaleHistoryLegacy(ctx, scope, offset, limit)
	return records, scope, err
}

// CosmicSignatureMarketplaceOffersPage returns one bounded live order-book
// page for the Cosmic Signature collection.
func (r *Repo) CosmicSignatureMarketplaceOffersPage(
	ctx context.Context,
	sort marketstore.OfferSort,
	after *marketstore.OfferPageCursor,
	limit int,
) ([]marketstore.OfferRecord, bool, error) {
	scope, err := r.CosmicSignatureMarketplaceScope(ctx)
	if err != nil {
		return nil, false, err
	}
	return r.marketplace.ActiveOffersPage(ctx, scope, sort, after, limit)
}

// CosmicSignatureMarketplaceOfferHistoryPage returns one bounded
// newest-first offer ledger page for the Cosmic Signature collection.
func (r *Repo) CosmicSignatureMarketplaceOfferHistoryPage(
	ctx context.Context,
	after *marketstore.EventPageCursor,
	limit int,
) ([]marketstore.OfferHistoryRecord, bool, error) {
	scope, err := r.CosmicSignatureMarketplaceScope(ctx)
	if err != nil {
		return nil, false, err
	}
	return r.marketplace.OfferHistoryPage(ctx, scope, after, limit)
}

// CosmicSignatureMarketplaceTradesPage returns one bounded newest-first
// purchase ledger page for the Cosmic Signature collection.
func (r *Repo) CosmicSignatureMarketplaceTradesPage(
	ctx context.Context,
	after *marketstore.EventPageCursor,
	limit int,
) ([]marketstore.TradeRecord, bool, error) {
	scope, err := r.CosmicSignatureMarketplaceScope(ctx)
	if err != nil {
		return nil, false, err
	}
	return r.marketplace.TradesPage(ctx, scope, after, limit)
}

// CosmicSignatureMarketplaceFloorPrice returns the exact v2 sell-side floor
// for the Cosmic Signature collection.
func (r *Repo) CosmicSignatureMarketplaceFloorPrice(
	ctx context.Context,
) (marketstore.FloorPriceRecord, error) {
	scope, err := r.CosmicSignatureMarketplaceScope(ctx)
	if err != nil {
		return marketstore.FloorPriceRecord{}, err
	}
	return r.marketplace.FloorPrice(ctx, scope)
}
