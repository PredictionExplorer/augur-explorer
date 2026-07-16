package v2

import (
	"errors"
	"fmt"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"

	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// randomWalkTx is the mapped transaction identity shared by every
// RandomWalk event representation.
type randomWalkTx struct {
	EventLogID      int64
	BlockNumber     int64
	TransactionHash string
	OccurredAt      time.Time
}

func mapRandomWalkTx(tx rwstore.EventTx) (randomWalkTx, error) {
	if tx.EvtLogID < 1 || tx.BlockNum < 0 {
		return randomWalkTx{}, errors.New("invalid event identity")
	}
	if !isTransactionHash(tx.TxHash) {
		return randomWalkTx{}, errors.New("invalid transaction hash")
	}
	occurredAt, err := time.Parse(time.RFC3339Nano, tx.DateTime)
	if err != nil {
		return randomWalkTx{}, fmt.Errorf("parse event timestamp: %w", err)
	}
	return randomWalkTx{
		EventLogID:      tx.EvtLogID,
		BlockNumber:     tx.BlockNum,
		TransactionHash: strings.ToLower(tx.TxHash),
		OccurredAt:      occurredAt.UTC(),
	}, nil
}

// canonicalAnyAddress checksums an address that may legitimately be the
// zero address (burn transfers).
func canonicalAnyAddress(name, value string) (string, error) {
	if !ethcommon.IsHexAddress(value) {
		return "", fmt.Errorf("invalid %s address", name)
	}
	return ethcommon.HexToAddress(value).Hex(), nil
}

func mapRandomWalkToken(record rwstore.TokenRecord) (RandomWalkToken, error) {
	if record.TokenID < 0 {
		return RandomWalkToken{}, errors.New("invalid RandomWalk token identity")
	}
	transaction, err := mapRandomWalkTx(record.MintTx)
	if err != nil {
		return RandomWalkToken{}, fmt.Errorf("RandomWalk mint transaction: %w", err)
	}
	minter, err := canonicalNonZeroAddress("RandomWalk minter", record.MinterAddr)
	if err != nil {
		return RandomWalkToken{}, err
	}
	owner, err := canonicalNonZeroAddress("RandomWalk owner", record.CurOwnerAddr)
	if err != nil {
		return RandomWalkToken{}, err
	}
	if record.Seed == "" {
		return RandomWalkToken{}, errors.New("RandomWalk token misses its mint seed")
	}
	seedNumber, err := requiredAmount(record.SeedNum)
	if err != nil {
		return RandomWalkToken{}, fmt.Errorf("RandomWalk seed number: %w", err)
	}
	mintPrice, err := requiredAmount(record.MintPriceWei)
	if err != nil {
		return RandomWalkToken{}, fmt.Errorf("RandomWalk mint price: %w", err)
	}
	lastPrice, err := requiredAmount(record.LastPriceWei)
	if err != nil {
		return RandomWalkToken{}, fmt.Errorf("RandomWalk last price: %w", err)
	}
	volume, err := requiredAmount(record.TradingVolumeWei)
	if err != nil {
		return RandomWalkToken{}, fmt.Errorf("RandomWalk token volume: %w", err)
	}
	if record.TradeCount < 0 {
		return RandomWalkToken{}, errors.New("negative RandomWalk trade count")
	}
	result := RandomWalkToken{
		BlockNumber:         transaction.BlockNumber,
		CurrentOwnerAddress: owner,
		EventLogId:          transaction.EventLogID,
		LastPriceWei:        lastPrice,
		MintPriceWei:        mintPrice,
		MintedAt:            transaction.OccurredAt,
		MinterAddress:       minter,
		NftTokenId:          record.TokenID,
		Seed:                record.Seed,
		SeedNumber:          seedNumber,
		TradeCount:          record.TradeCount,
		TradingVolumeWei:    volume,
		TransactionHash:     transaction.TransactionHash,
	}
	if record.TokenName != "" {
		name := record.TokenName
		result.TokenName = &name
	}
	return result, nil
}

func mapRandomWalkTokenDetail(record rwstore.TokenDetailRecord) (RandomWalkTokenDetail, error) {
	token, err := mapRandomWalkToken(record.TokenRecord)
	if err != nil {
		return RandomWalkTokenDetail{}, err
	}
	detail := RandomWalkTokenDetail{
		BlockNumber:         token.BlockNumber,
		CurrentOwnerAddress: token.CurrentOwnerAddress,
		EventLogId:          token.EventLogId,
		LastPriceWei:        token.LastPriceWei,
		MintPriceWei:        token.MintPriceWei,
		MintedAt:            token.MintedAt,
		MinterAddress:       token.MinterAddress,
		NftTokenId:          token.NftTokenId,
		Seed:                token.Seed,
		SeedNumber:          token.SeedNumber,
		TokenName:           token.TokenName,
		TradeCount:          token.TradeCount,
		TradingVolumeWei:    token.TradingVolumeWei,
		TransactionHash:     token.TransactionHash,
	}
	if record.NameChangeText != "" {
		changedAt, err := time.Parse(time.RFC3339Nano, record.NameChangeText)
		if err != nil {
			return RandomWalkTokenDetail{}, fmt.Errorf("parse rename timestamp: %w", err)
		}
		utc := changedAt.UTC()
		detail.NameChangedAt = &utc
	}
	return detail, nil
}

func mapRandomWalkTokenNameChange(
	record rwstore.TokenNameChangeRecord,
) (RandomWalkTokenNameChange, error) {
	if record.TokenID < 0 {
		return RandomWalkTokenNameChange{}, errors.New("invalid rename token identity")
	}
	transaction, err := mapRandomWalkTx(record.Tx)
	if err != nil {
		return RandomWalkTokenNameChange{}, fmt.Errorf("RandomWalk rename transaction: %w", err)
	}
	owner, err := canonicalNonZeroAddress("RandomWalk renamer", record.Owner)
	if err != nil {
		return RandomWalkTokenNameChange{}, err
	}
	return RandomWalkTokenNameChange{
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogID,
		NewName:         record.NewName,
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		OwnerAddress:    owner,
		TransactionHash: transaction.TransactionHash,
	}, nil
}

func mapRandomWalkOfferSide(offerType int16) (RandomWalkOfferSide, error) {
	switch offerType {
	case 0:
		return Buy, nil
	case 1:
		return Sell, nil
	default:
		return "", fmt.Errorf("unknown offer side %d", offerType)
	}
}

func mapRandomWalkTokenEvent(record rwstore.TokenEventRecord) (RandomWalkTokenEvent, error) {
	if record.TokenID < 0 {
		return RandomWalkTokenEvent{}, errors.New("invalid token event identity")
	}
	transaction, err := mapRandomWalkTx(record.Tx)
	if err != nil {
		return RandomWalkTokenEvent{}, fmt.Errorf("token event transaction: %w", err)
	}
	result := RandomWalkTokenEvent{
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogID,
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		TransactionHash: transaction.TransactionHash,
	}
	switch record.Kind {
	case rwstore.TokenEventMint:
		result.EventType = RandomWalkTokenEventTypeMint
		minter, err := canonicalNonZeroAddress("mint event minter", record.MinterAddr)
		if err != nil {
			return RandomWalkTokenEvent{}, err
		}
		if record.Seed == "" {
			return RandomWalkTokenEvent{}, errors.New("mint event misses its seed")
		}
		seedNumber, err := requiredAmount(record.SeedNum)
		if err != nil {
			return RandomWalkTokenEvent{}, fmt.Errorf("mint event seed number: %w", err)
		}
		price, err := requiredAmount(record.PriceWei)
		if err != nil {
			return RandomWalkTokenEvent{}, fmt.Errorf("mint event price: %w", err)
		}
		seed := record.Seed
		result.MinterAddress = &minter
		result.Seed = &seed
		result.SeedNumber = &seedNumber
		result.PriceWei = &price
	case rwstore.TokenEventTransfer:
		result.EventType = RandomWalkTokenEventTypeTransfer
		from, err := canonicalAnyAddress("transfer sender", record.FromAddr)
		if err != nil {
			return RandomWalkTokenEvent{}, err
		}
		to, err := canonicalAnyAddress("transfer recipient", record.ToAddr)
		if err != nil {
			return RandomWalkTokenEvent{}, err
		}
		result.FromAddress = &from
		result.ToAddress = &to
	case rwstore.TokenEventNameChange:
		result.EventType = RandomWalkTokenEventTypeNameChange
		if !record.HasNewName {
			return RandomWalkTokenEvent{}, errors.New("rename event misses its name")
		}
		name := record.NewName
		result.NewName = &name
	case rwstore.TokenEventListed, rwstore.TokenEventOfferCanceled, rwstore.TokenEventPurchase:
		if !record.HasOffer {
			return RandomWalkTokenEvent{}, errors.New("marketplace event misses its offer")
		}
		if record.OfferID < 0 {
			return RandomWalkTokenEvent{}, errors.New("negative marketplace offer id")
		}
		side, err := mapRandomWalkOfferSide(record.OfferType)
		if err != nil {
			return RandomWalkTokenEvent{}, err
		}
		price, err := requiredAmount(record.PriceWei)
		if err != nil {
			return RandomWalkTokenEvent{}, fmt.Errorf("marketplace event price: %w", err)
		}
		offerID := record.OfferID
		result.OfferId = &offerID
		result.OfferSide = &side
		result.PriceWei = &price
		if record.Kind == rwstore.TokenEventPurchase {
			result.EventType = RandomWalkTokenEventTypePurchase
			buyer, err := canonicalNonZeroAddress("purchase buyer", record.BuyerAddr)
			if err != nil {
				return RandomWalkTokenEvent{}, err
			}
			seller, err := canonicalNonZeroAddress("purchase seller", record.SellerAddr)
			if err != nil {
				return RandomWalkTokenEvent{}, err
			}
			result.BuyerAddress = &buyer
			result.SellerAddress = &seller
		} else {
			result.EventType = RandomWalkTokenEventTypeListed
			if record.Kind == rwstore.TokenEventOfferCanceled {
				result.EventType = RandomWalkTokenEventTypeOfferCanceled
			}
			maker, err := canonicalNonZeroAddress("offer maker", record.MakerAddr)
			if err != nil {
				return RandomWalkTokenEvent{}, err
			}
			result.MakerAddress = &maker
		}
	default:
		return RandomWalkTokenEvent{}, fmt.Errorf("unknown token event kind %q", record.Kind)
	}
	return result, nil
}

func mapRandomWalkMarketplaceOffer(record rwstore.OfferRecord) (RandomWalkMarketplaceOffer, error) {
	if record.OfferID < 0 || record.TokenID < 0 {
		return RandomWalkMarketplaceOffer{}, errors.New("invalid marketplace offer identity")
	}
	transaction, err := mapRandomWalkTx(record.ListTx)
	if err != nil {
		return RandomWalkMarketplaceOffer{}, fmt.Errorf("marketplace offer transaction: %w", err)
	}
	side, err := mapRandomWalkOfferSide(record.OfferType)
	if err != nil {
		return RandomWalkMarketplaceOffer{}, err
	}
	price, err := requiredAmount(record.PriceWei)
	if err != nil {
		return RandomWalkMarketplaceOffer{}, fmt.Errorf("marketplace offer price: %w", err)
	}
	maker, err := canonicalNonZeroAddress("marketplace offer maker", record.MakerAddr)
	if err != nil {
		return RandomWalkMarketplaceOffer{}, err
	}
	return RandomWalkMarketplaceOffer{
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogID,
		ListedAt:        transaction.OccurredAt,
		MakerAddress:    maker,
		NftTokenId:      record.TokenID,
		OfferId:         record.OfferID,
		PriceWei:        price,
		Side:            side,
		TransactionHash: transaction.TransactionHash,
	}, nil
}

func mapRandomWalkOfferHistoryEntry(
	record rwstore.OfferHistoryRecord,
) (RandomWalkOfferHistoryEntry, error) {
	if record.OfferID < 0 || record.TokenID < 0 {
		return RandomWalkOfferHistoryEntry{}, errors.New("invalid offer history identity")
	}
	transaction, err := mapRandomWalkTx(record.ListTx)
	if err != nil {
		return RandomWalkOfferHistoryEntry{}, fmt.Errorf("offer history transaction: %w", err)
	}
	side, err := mapRandomWalkOfferSide(record.OfferType)
	if err != nil {
		return RandomWalkOfferHistoryEntry{}, err
	}
	price, err := requiredAmount(record.PriceWei)
	if err != nil {
		return RandomWalkOfferHistoryEntry{}, fmt.Errorf("offer history price: %w", err)
	}
	maker, err := canonicalNonZeroAddress("offer history maker", record.MakerAddr)
	if err != nil {
		return RandomWalkOfferHistoryEntry{}, err
	}
	result := RandomWalkOfferHistoryEntry{
		BlockNumber:     transaction.BlockNumber,
		EventLogId:      transaction.EventLogID,
		ListedAt:        transaction.OccurredAt,
		MakerAddress:    maker,
		NftTokenId:      record.TokenID,
		OfferId:         record.OfferID,
		PriceWei:        price,
		Side:            side,
		TransactionHash: transaction.TransactionHash,
	}
	switch {
	case record.Purchase != nil && record.Cancellation != nil:
		return RandomWalkOfferHistoryEntry{}, errors.New("offer is both bought and canceled")
	case record.Purchase != nil:
		if record.Active {
			return RandomWalkOfferHistoryEntry{}, errors.New("bought offer is still active")
		}
		result.Status = Bought
		purchaseTx, err := mapRandomWalkTx(record.Purchase.Tx)
		if err != nil {
			return RandomWalkOfferHistoryEntry{}, fmt.Errorf("offer purchase transaction: %w", err)
		}
		buyer, err := canonicalNonZeroAddress("offer buyer", record.Purchase.BuyerAddr)
		if err != nil {
			return RandomWalkOfferHistoryEntry{}, err
		}
		seller, err := canonicalNonZeroAddress("offer seller", record.Purchase.SellerAddr)
		if err != nil {
			return RandomWalkOfferHistoryEntry{}, err
		}
		result.Purchase = &RandomWalkOfferPurchase{
			BlockNumber:     purchaseTx.BlockNumber,
			BuyerAddress:    buyer,
			EventLogId:      purchaseTx.EventLogID,
			OccurredAt:      purchaseTx.OccurredAt,
			SellerAddress:   seller,
			TransactionHash: purchaseTx.TransactionHash,
		}
	case record.Cancellation != nil:
		if record.Active {
			return RandomWalkOfferHistoryEntry{}, errors.New("canceled offer is still active")
		}
		result.Status = Canceled
		cancelTx, err := mapRandomWalkTx(*record.Cancellation)
		if err != nil {
			return RandomWalkOfferHistoryEntry{}, fmt.Errorf("offer cancellation transaction: %w", err)
		}
		result.Cancellation = &RandomWalkOfferCancellation{
			BlockNumber:     cancelTx.BlockNumber,
			EventLogId:      cancelTx.EventLogID,
			OccurredAt:      cancelTx.OccurredAt,
			TransactionHash: cancelTx.TransactionHash,
		}
	default:
		if !record.Active {
			return RandomWalkOfferHistoryEntry{}, errors.New("closed offer has no closing event")
		}
		result.Status = Active
	}
	if record.ProfitWei != "" {
		profit, err := canonicalSignedInteger(record.ProfitWei)
		if err != nil {
			return RandomWalkOfferHistoryEntry{}, fmt.Errorf("offer seller profit: %w", err)
		}
		result.SellerProfitWei = &profit
	}
	return result, nil
}

func mapRandomWalkTrade(record rwstore.TradeRecord) (RandomWalkTrade, error) {
	if record.OfferID < 0 || record.TokenID < 0 {
		return RandomWalkTrade{}, errors.New("invalid trade identity")
	}
	transaction, err := mapRandomWalkTx(record.Tx)
	if err != nil {
		return RandomWalkTrade{}, fmt.Errorf("trade transaction: %w", err)
	}
	side, err := mapRandomWalkOfferSide(record.OfferType)
	if err != nil {
		return RandomWalkTrade{}, err
	}
	price, err := requiredAmount(record.PriceWei)
	if err != nil {
		return RandomWalkTrade{}, fmt.Errorf("trade price: %w", err)
	}
	buyer, err := canonicalNonZeroAddress("trade buyer", record.BuyerAddr)
	if err != nil {
		return RandomWalkTrade{}, err
	}
	seller, err := canonicalNonZeroAddress("trade seller", record.SellerAddr)
	if err != nil {
		return RandomWalkTrade{}, err
	}
	result := RandomWalkTrade{
		BlockNumber:     transaction.BlockNumber,
		BuyerAddress:    buyer,
		EventLogId:      transaction.EventLogID,
		NftTokenId:      record.TokenID,
		OccurredAt:      transaction.OccurredAt,
		OfferId:         record.OfferID,
		PriceWei:        price,
		SellerAddress:   seller,
		Side:            side,
		TransactionHash: transaction.TransactionHash,
	}
	if record.ProfitWei != "" {
		profit, err := canonicalSignedInteger(record.ProfitWei)
		if err != nil {
			return RandomWalkTrade{}, fmt.Errorf("trade seller profit: %w", err)
		}
		result.SellerProfitWei = &profit
	}
	return result, nil
}

func mapRandomWalkFloorPrice(record rwstore.FloorPriceRecord) (RandomWalkFloorPrice, error) {
	if record.ActiveSellOfferCount < 0 {
		return RandomWalkFloorPrice{}, errors.New("negative active sell-offer count")
	}
	result := RandomWalkFloorPrice{ActiveSellOfferCount: record.ActiveSellOfferCount}
	if record.Floor == nil {
		if record.ActiveSellOfferCount != 0 {
			return RandomWalkFloorPrice{}, errors.New("non-empty book without a floor listing")
		}
		return result, nil
	}
	if record.ActiveSellOfferCount == 0 {
		return RandomWalkFloorPrice{}, errors.New("floor listing on an empty book")
	}
	if record.Floor.OfferID < 0 || record.Floor.TokenID < 0 {
		return RandomWalkFloorPrice{}, errors.New("invalid floor listing identity")
	}
	price, err := requiredAmount(record.Floor.PriceWei)
	if err != nil {
		return RandomWalkFloorPrice{}, fmt.Errorf("floor price: %w", err)
	}
	listedAt, err := time.Parse(time.RFC3339Nano, record.Floor.ListedAtText)
	if err != nil {
		return RandomWalkFloorPrice{}, fmt.Errorf("parse floor listing timestamp: %w", err)
	}
	result.Floor = &RandomWalkFloorListing{
		ListedAt:   listedAt.UTC(),
		NftTokenId: record.Floor.TokenID,
		OfferId:    record.Floor.OfferID,
		PriceWei:   price,
	}
	return result, nil
}

// zeroRandomWalkUserProfile is the stable zero-activity shape returned for
// valid wallets the indexer has never seen.
func zeroRandomWalkUserProfile(address string) RandomWalkUserProfile {
	return RandomWalkUserProfile{
		Address:          address,
		ProfitWei:        "0",
		TradingVolumeWei: "0",
	}
}

func mapRandomWalkUserProfile(record rwstore.UserProfileRecord) (RandomWalkUserProfile, error) {
	address, err := canonicalNonZeroAddress("RandomWalk user", record.Address)
	if err != nil {
		return RandomWalkUserProfile{}, err
	}
	if record.MintedTokenCount < 0 || record.OwnedTokenCount < 0 ||
		record.TradeCount < 0 || record.WithdrawalCount < 0 {
		return RandomWalkUserProfile{}, errors.New("negative RandomWalk user counter")
	}
	volume, err := requiredAmount(record.TradingVolumeWei)
	if err != nil {
		return RandomWalkUserProfile{}, fmt.Errorf("RandomWalk user volume: %w", err)
	}
	profit, err := canonicalSignedInteger(record.ProfitWei)
	if err != nil {
		return RandomWalkUserProfile{}, fmt.Errorf("RandomWalk user profit: %w", err)
	}
	return RandomWalkUserProfile{
		Address:          address,
		MintedTokenCount: record.MintedTokenCount,
		OwnedTokenCount:  record.OwnedTokenCount,
		ProfitWei:        profit,
		TradeCount:       record.TradeCount,
		TradingVolumeWei: volume,
		WithdrawalCount:  record.WithdrawalCount,
	}, nil
}

func mapRandomWalkOwnedToken(record rwstore.OwnedTokenRecord) (RandomWalkOwnedToken, error) {
	if record.TokenID < 0 || record.TradeCount < 0 {
		return RandomWalkOwnedToken{}, errors.New("invalid owned token identity")
	}
	lastPrice, err := requiredAmount(record.LastPriceWei)
	if err != nil {
		return RandomWalkOwnedToken{}, fmt.Errorf("owned token last price: %w", err)
	}
	volume, err := requiredAmount(record.TradingVolumeWei)
	if err != nil {
		return RandomWalkOwnedToken{}, fmt.Errorf("owned token volume: %w", err)
	}
	result := RandomWalkOwnedToken{
		LastPriceWei:     lastPrice,
		NftTokenId:       record.TokenID,
		TradeCount:       record.TradeCount,
		TradingVolumeWei: volume,
	}
	if record.Seed != "" {
		seed := record.Seed
		seedNumber, err := requiredAmount(record.SeedNum)
		if err != nil {
			return RandomWalkOwnedToken{}, fmt.Errorf("owned token seed number: %w", err)
		}
		result.Seed = &seed
		result.SeedNumber = &seedNumber
	}
	if record.TokenName != "" {
		name := record.TokenName
		result.TokenName = &name
	}
	if record.HasMint {
		mintedAt, err := time.Parse(time.RFC3339Nano, record.MintText)
		if err != nil {
			return RandomWalkOwnedToken{}, fmt.Errorf("parse owned token mint timestamp: %w", err)
		}
		mintPrice, err := requiredAmount(record.MintPriceWei)
		if err != nil {
			return RandomWalkOwnedToken{}, fmt.Errorf("owned token mint price: %w", err)
		}
		utc := mintedAt.UTC()
		result.MintedAt = &utc
		result.MintPriceWei = &mintPrice
	}
	return result, nil
}

func mapRandomWalkStatistics(record rwstore.StatisticsRecord) (RandomWalkStatistics, error) {
	if record.MintedCount < 0 || record.UniqueOwnerCount < 0 ||
		record.TokenTradeCount < 0 || record.MarketTradeCount < 0 ||
		record.ActiveSellOfferCount < 0 || record.ActiveBuyOfferCount < 0 ||
		record.WithdrawalCount < 0 {
		return RandomWalkStatistics{}, errors.New("negative RandomWalk statistics counter")
	}
	tokenVolume, err := requiredAmount(record.TokenTradingVolumeWei)
	if err != nil {
		return RandomWalkStatistics{}, fmt.Errorf("collection volume: %w", err)
	}
	mintFunds, err := requiredAmount(record.MintFundsWei)
	if err != nil {
		return RandomWalkStatistics{}, fmt.Errorf("mint funds: %w", err)
	}
	marketVolume, err := requiredAmount(record.MarketTradingVolumeWei)
	if err != nil {
		return RandomWalkStatistics{}, fmt.Errorf("marketplace volume: %w", err)
	}
	result := RandomWalkStatistics{
		Tokens: RandomWalkTokenStatistics{
			MintFundsWei:     mintFunds,
			MintedCount:      record.MintedCount,
			TradeCount:       record.TokenTradeCount,
			TradingVolumeWei: tokenVolume,
			UniqueOwnerCount: record.UniqueOwnerCount,
		},
		Marketplace: RandomWalkMarketplaceStatistics{
			ActiveBuyOfferCount:  record.ActiveBuyOfferCount,
			ActiveSellOfferCount: record.ActiveSellOfferCount,
			TradeCount:           record.MarketTradeCount,
			TradingVolumeWei:     marketVolume,
		},
		Withdrawals: RandomWalkWithdrawalStatistics{
			Count: record.WithdrawalCount,
		},
	}
	if record.LastMint != nil {
		if record.LastMint.TokenID < 0 {
			return RandomWalkStatistics{}, errors.New("invalid last mint identity")
		}
		price, err := requiredAmount(record.LastMint.PriceWei)
		if err != nil {
			return RandomWalkStatistics{}, fmt.Errorf("last mint price: %w", err)
		}
		minter, err := canonicalNonZeroAddress("last mint minter", record.LastMint.MinterAddr)
		if err != nil {
			return RandomWalkStatistics{}, err
		}
		mintedAt, err := time.Parse(time.RFC3339Nano, record.LastMint.MintText)
		if err != nil {
			return RandomWalkStatistics{}, fmt.Errorf("parse last mint timestamp: %w", err)
		}
		result.Tokens.LastMint = &RandomWalkLastMint{
			MintedAt:      mintedAt.UTC(),
			MinterAddress: minter,
			NftTokenId:    record.LastMint.TokenID,
			PriceWei:      price,
		}
	}
	if record.LatestWithdrawal != nil {
		if record.LatestWithdrawal.TokenID < 0 {
			return RandomWalkStatistics{}, errors.New("invalid latest withdrawal identity")
		}
		amount, err := requiredAmount(record.LatestWithdrawal.AmountWei)
		if err != nil {
			return RandomWalkStatistics{}, fmt.Errorf("latest withdrawal amount: %w", err)
		}
		withdrawer, err := canonicalNonZeroAddress(
			"latest withdrawer", record.LatestWithdrawal.WithdrawerAddr)
		if err != nil {
			return RandomWalkStatistics{}, err
		}
		occurredAt, err := time.Parse(time.RFC3339Nano, record.LatestWithdrawal.OccurredText)
		if err != nil {
			return RandomWalkStatistics{}, fmt.Errorf("parse latest withdrawal timestamp: %w", err)
		}
		result.Withdrawals.Latest = &RandomWalkLatestWithdrawal{
			AmountWei:         amount,
			NftTokenId:        record.LatestWithdrawal.TokenID,
			OccurredAt:        occurredAt.UTC(),
			WithdrawerAddress: withdrawer,
		}
	}
	return result, nil
}

func mapRandomWalkWithdrawal(record rwstore.WithdrawalRecord) (RandomWalkWithdrawal, error) {
	if record.TokenID < 0 {
		return RandomWalkWithdrawal{}, errors.New("invalid withdrawal identity")
	}
	transaction, err := mapRandomWalkTx(record.Tx)
	if err != nil {
		return RandomWalkWithdrawal{}, fmt.Errorf("withdrawal transaction: %w", err)
	}
	withdrawer, err := canonicalNonZeroAddress("withdrawer", record.WithdrawerAddr)
	if err != nil {
		return RandomWalkWithdrawal{}, err
	}
	amount, err := requiredAmount(record.AmountWei)
	if err != nil {
		return RandomWalkWithdrawal{}, fmt.Errorf("withdrawal amount: %w", err)
	}
	return RandomWalkWithdrawal{
		AmountWei:         amount,
		BlockNumber:       transaction.BlockNumber,
		EventLogId:        transaction.EventLogID,
		NftTokenId:        record.TokenID,
		OccurredAt:        transaction.OccurredAt,
		TransactionHash:   transaction.TransactionHash,
		WithdrawerAddress: withdrawer,
	}, nil
}
