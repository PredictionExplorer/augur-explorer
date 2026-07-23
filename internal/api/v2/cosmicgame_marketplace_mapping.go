package v2

import (
	marketstore "github.com/PredictionExplorer/augur-explorer/internal/store/marketplace"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

func randomWalkStoreEventTx(record marketstore.EventTx) rwstore.EventTx {
	return rwstore.EventTx{
		EvtLogID:  record.EventLogID,
		BlockNum:  record.BlockNum,
		TxID:      record.TxID,
		TxHash:    record.TxHash,
		TimeStamp: record.TimeStamp,
		DateTime:  record.DateTime,
	}
}

func mapCosmicSignatureMarketplaceOffer(
	record marketstore.OfferRecord,
) (CosmicSignatureMarketplaceOffer, error) {
	return mapRandomWalkMarketplaceOffer(rwstore.OfferRecord{
		ListTx:    randomWalkStoreEventTx(record.ListTx),
		OfferID:   record.OfferID,
		OfferType: record.OfferType,
		TokenID:   record.TokenID,
		PriceWei:  record.PriceWei,
		MakerAid:  record.MakerAid,
		MakerAddr: record.MakerAddr,
	})
}

func mapCosmicSignatureMarketplaceOfferHistory(
	record marketstore.OfferHistoryRecord,
) (CosmicSignatureMarketplaceOfferHistoryEntry, error) {
	randomWalkRecord := rwstore.OfferHistoryRecord{
		ListTx:    randomWalkStoreEventTx(record.ListTx),
		OfferID:   record.OfferID,
		OfferType: record.OfferType,
		TokenID:   record.TokenID,
		PriceWei:  record.PriceWei,
		MakerAid:  record.MakerAid,
		MakerAddr: record.MakerAddr,
		Active:    record.Active,
		ProfitWei: record.ProfitWei,
	}
	if record.Purchase != nil {
		randomWalkRecord.Purchase = &rwstore.OfferOutcomePurchase{
			Tx:         randomWalkStoreEventTx(record.Purchase.Tx),
			BuyerAid:   record.Purchase.BuyerAid,
			BuyerAddr:  record.Purchase.BuyerAddr,
			SellerAid:  record.Purchase.SellerAid,
			SellerAddr: record.Purchase.SellerAddr,
		}
	}
	if record.Cancellation != nil {
		cancellation := randomWalkStoreEventTx(*record.Cancellation)
		randomWalkRecord.Cancellation = &cancellation
	}
	return mapRandomWalkOfferHistoryEntry(randomWalkRecord)
}

func mapCosmicSignatureMarketplaceTrade(
	record marketstore.TradeRecord,
) (CosmicSignatureMarketplaceTrade, error) {
	return mapRandomWalkTrade(rwstore.TradeRecord{
		Tx:         randomWalkStoreEventTx(record.Tx),
		OfferID:    record.OfferID,
		OfferType:  record.OfferType,
		TokenID:    record.TokenID,
		PriceWei:   record.PriceWei,
		BuyerAid:   record.BuyerAid,
		BuyerAddr:  record.BuyerAddr,
		SellerAid:  record.SellerAid,
		SellerAddr: record.SellerAddr,
		ProfitWei:  record.ProfitWei,
	})
}

func mapCosmicSignatureMarketplaceFloorPrice(
	record marketstore.FloorPriceRecord,
) (CosmicSignatureMarketplaceFloorPrice, error) {
	randomWalkRecord := rwstore.FloorPriceRecord{
		ActiveSellOfferCount: record.ActiveSellOfferCount,
	}
	if record.Floor != nil {
		randomWalkRecord.Floor = &rwstore.FloorListingRecord{
			OfferID:      record.Floor.OfferID,
			TokenID:      record.Floor.TokenID,
			PriceWei:     record.Floor.PriceWei,
			ListedAtTs:   record.Floor.ListedAtTs,
			ListedAtText: record.Floor.ListedAtText,
		}
	}
	return mapRandomWalkFloorPrice(randomWalkRecord)
}
