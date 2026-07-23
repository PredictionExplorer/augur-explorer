package v2

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	cosmicMarketplaceCursorVersion = 1
	cosmicSignatureCollectionScope = "cosmicSignature"
)

type cosmicMarketplaceResource string

const (
	cosmicMarketplaceResourceOffers       cosmicMarketplaceResource = "cosmicSignatureMarketplaceOffers"
	cosmicMarketplaceResourceOfferHistory cosmicMarketplaceResource = "cosmicSignatureMarketplaceOfferHistory"
	cosmicMarketplaceResourceTrades       cosmicMarketplaceResource = "cosmicSignatureMarketplaceTrades"
)

var errInvalidCosmicMarketplaceCursor = errors.New("invalid cosmic marketplace cursor")

type cosmicMarketplaceOfferCursor struct {
	Version    int                       `json:"v"`
	Resource   cosmicMarketplaceResource `json:"k"`
	Collection string                    `json:"c"`
	Sort       RandomWalkOfferSort       `json:"s"`
	PriceWei   string                    `json:"p,omitempty"`
	EventLogID int64                     `json:"e"`
}

func encodeCosmicMarketplaceOfferCursor(
	cursor cosmicMarketplaceOfferCursor,
) (string, error) {
	if !validCosmicMarketplaceOfferCursor(cursor, cursor.Sort) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidCosmicMarketplaceCursor)
	}
	return encodeOpaqueCursor(
		cursor,
		errInvalidCosmicMarketplaceCursor,
		"cosmic marketplace offer cursor",
	)
}

func decodeCosmicMarketplaceOfferCursor(
	encoded string,
	expectedSort RandomWalkOfferSort,
) (cosmicMarketplaceOfferCursor, error) {
	cursor, err := decodeOpaqueCursor[cosmicMarketplaceOfferCursor](
		encoded,
		errInvalidCosmicMarketplaceCursor,
	)
	if err != nil {
		return cosmicMarketplaceOfferCursor{}, err
	}
	if !validCosmicMarketplaceOfferCursor(cursor, expectedSort) {
		return cosmicMarketplaceOfferCursor{},
			fmt.Errorf("%w: invalid fields", errInvalidCosmicMarketplaceCursor)
	}
	return cursor, nil
}

func validCosmicMarketplaceOfferCursor(
	cursor cosmicMarketplaceOfferCursor,
	expectedSort RandomWalkOfferSort,
) bool {
	if cursor.Version != cosmicMarketplaceCursorVersion ||
		cursor.Resource != cosmicMarketplaceResourceOffers ||
		cursor.Collection != cosmicSignatureCollectionScope ||
		cursor.EventLogID < 1 ||
		!cursor.Sort.Valid() ||
		cursor.Sort != expectedSort {
		return false
	}
	switch cursor.Sort {
	case PriceAsc, PriceDesc:
		return validWeiAmountString(cursor.PriceWei)
	case Newest, Oldest:
		return cursor.PriceWei == ""
	default:
		return false
	}
}

type cosmicMarketplaceLedgerCursor struct {
	Version    int                       `json:"v"`
	Resource   cosmicMarketplaceResource `json:"k"`
	Collection string                    `json:"c"`
	EventLogID int64                     `json:"e"`
}

func validCosmicMarketplaceLedgerResource(resource cosmicMarketplaceResource) bool {
	return resource == cosmicMarketplaceResourceOfferHistory ||
		resource == cosmicMarketplaceResourceTrades
}

func encodeCosmicMarketplaceLedgerCursor(
	cursor cosmicMarketplaceLedgerCursor,
) (string, error) {
	if !validCosmicMarketplaceLedgerCursor(cursor, cursor.Resource) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidCosmicMarketplaceCursor)
	}
	return encodeOpaqueCursor(
		cursor,
		errInvalidCosmicMarketplaceCursor,
		"cosmic marketplace ledger cursor",
	)
}

func decodeCosmicMarketplaceLedgerCursor(
	encoded string,
	expectedResource cosmicMarketplaceResource,
) (cosmicMarketplaceLedgerCursor, error) {
	cursor, err := decodeOpaqueCursor[cosmicMarketplaceLedgerCursor](
		encoded,
		errInvalidCosmicMarketplaceCursor,
	)
	if err != nil {
		return cosmicMarketplaceLedgerCursor{}, err
	}
	if !validCosmicMarketplaceLedgerCursor(cursor, expectedResource) {
		return cosmicMarketplaceLedgerCursor{},
			fmt.Errorf("%w: invalid fields", errInvalidCosmicMarketplaceCursor)
	}
	return cursor, nil
}

func validCosmicMarketplaceLedgerCursor(
	cursor cosmicMarketplaceLedgerCursor,
	expectedResource cosmicMarketplaceResource,
) bool {
	return cursor.Version == cosmicMarketplaceCursorVersion &&
		cursor.Collection == cosmicSignatureCollectionScope &&
		cursor.EventLogID >= 1 &&
		cursor.Resource == expectedResource &&
		validCosmicMarketplaceLedgerResource(cursor.Resource) &&
		validCosmicMarketplaceLedgerResource(expectedResource)
}

func invalidCosmicMarketplaceCursorProblem(instance string) Problem {
	return newProblem(
		http.StatusBadRequest,
		"invalid-cursor",
		"Invalid cursor",
		"The cursor is malformed or belongs to another marketplace resource, collection, or sort.",
		instance,
	)
}
