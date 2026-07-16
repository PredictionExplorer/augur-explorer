package v2

import (
	"errors"
	"fmt"
)

// randomWalkResource scopes every RandomWalk cursor to exactly one
// collection so continuation cursors cannot cross resources, sorts, filters
// or wallets — and cannot be replayed against the CosmicGame collections,
// whose cursor payloads never carry these resource values.
type randomWalkResource string

const (
	randomWalkResourceTokens         randomWalkResource = "rwalkTokens"
	randomWalkResourceNameHistory    randomWalkResource = "rwalkNameHistory"
	randomWalkResourceTokenEvents    randomWalkResource = "rwalkTokenEvents"
	randomWalkResourceOfferBook      randomWalkResource = "rwalkOffers"
	randomWalkResourceOfferHistory   randomWalkResource = "rwalkOfferHistory"
	randomWalkResourceTrades         randomWalkResource = "rwalkTrades"
	randomWalkResourceWithdrawals    randomWalkResource = "rwalkWithdrawals"
	randomWalkResourceUserOffers     randomWalkResource = "rwalkUserOffers"
	randomWalkResourceUserTokens     randomWalkResource = "rwalkUserTokens"
	randomWalkResourceRankingRatings randomWalkResource = "rwalkRankingRatings"
)

const randomWalkCursorVersion = 1

var errInvalidRandomWalkCursor = errors.New("invalid random walk cursor")

// randomWalkTokenFilterScope reproduces the directory request's filters
// inside the cursor so a continuation cannot cross into another filtered
// view. At most one of Named and Name is set.
type randomWalkTokenFilterScope struct {
	Named       bool   `json:"n,omitempty"`
	Name        string `json:"q,omitempty"`
	MintedFrom  *int64 `json:"mf,omitempty"`
	MintedUntil *int64 `json:"mu,omitempty"`
}

func (f randomWalkTokenFilterScope) valid() bool {
	if f.Named && f.Name != "" {
		return false
	}
	if len(f.Name) > maxTokenNameSearchLength {
		return false
	}
	if f.MintedFrom != nil && *f.MintedFrom < 0 {
		return false
	}
	if f.MintedUntil != nil && *f.MintedUntil < 1 {
		return false
	}
	return true
}

func (f randomWalkTokenFilterScope) equal(other randomWalkTokenFilterScope) bool {
	return f.Named == other.Named &&
		f.Name == other.Name &&
		equalInt64Pointer(f.MintedFrom, other.MintedFrom) &&
		equalInt64Pointer(f.MintedUntil, other.MintedUntil)
}

func equalInt64Pointer(a, b *int64) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return *a == *b
}

// randomWalkTokenCursor identifies the last token returned by the
// RandomWalk directory under one filter and sort. TradeCount participates
// only in the mostTraded ranking.
type randomWalkTokenCursor struct {
	Version    int                        `json:"v"`
	Resource   randomWalkResource         `json:"k"`
	Filter     randomWalkTokenFilterScope `json:"f"`
	Sort       RandomWalkTokenSort        `json:"s"`
	TradeCount int64                      `json:"c,omitempty"`
	TokenID    int64                      `json:"t"`
}

func encodeRandomWalkTokenCursor(cursor randomWalkTokenCursor) (string, error) {
	if !validRandomWalkTokenCursor(cursor, cursor.Filter, cursor.Sort) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRandomWalkCursor, "random walk token cursor")
}

func decodeRandomWalkTokenCursor(
	encoded string,
	expectedFilter randomWalkTokenFilterScope,
	expectedSort RandomWalkTokenSort,
) (randomWalkTokenCursor, error) {
	cursor, err := decodeOpaqueCursor[randomWalkTokenCursor](encoded, errInvalidRandomWalkCursor)
	if err != nil {
		return randomWalkTokenCursor{}, err
	}
	if !validRandomWalkTokenCursor(cursor, expectedFilter, expectedSort) {
		return randomWalkTokenCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return cursor, nil
}

func validRandomWalkTokenCursor(
	cursor randomWalkTokenCursor,
	expectedFilter randomWalkTokenFilterScope,
	expectedSort RandomWalkTokenSort,
) bool {
	if cursor.Version != randomWalkCursorVersion ||
		cursor.Resource != randomWalkResourceTokens ||
		cursor.TokenID < 0 || cursor.TradeCount < 0 ||
		!cursor.Filter.valid() || !expectedFilter.valid() ||
		!cursor.Filter.equal(expectedFilter) ||
		!cursor.Sort.Valid() || cursor.Sort != expectedSort {
		return false
	}
	// The immutable tokenId order never carries a trade-count key.
	return cursor.Sort != TokenId || cursor.TradeCount == 0
}

// randomWalkTokenEventCursor identifies the last event returned by one
// token-scoped, newest-first RandomWalk event page.
type randomWalkTokenEventCursor struct {
	Version    int                `json:"v"`
	Resource   randomWalkResource `json:"k"`
	TokenID    int64              `json:"t"`
	EventLogID int64              `json:"e"`
}

func validRandomWalkTokenEventResource(resource randomWalkResource) bool {
	return resource == randomWalkResourceNameHistory ||
		resource == randomWalkResourceTokenEvents
}

func encodeRandomWalkTokenEventCursor(cursor randomWalkTokenEventCursor) (string, error) {
	if !validRandomWalkTokenEventCursor(cursor, cursor.TokenID, cursor.Resource) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRandomWalkCursor, "random walk token event cursor")
}

func decodeRandomWalkTokenEventCursor(
	encoded string,
	expectedTokenID int64,
	expectedResource randomWalkResource,
) (randomWalkTokenEventCursor, error) {
	cursor, err := decodeOpaqueCursor[randomWalkTokenEventCursor](encoded, errInvalidRandomWalkCursor)
	if err != nil {
		return randomWalkTokenEventCursor{}, err
	}
	if !validRandomWalkTokenEventCursor(cursor, expectedTokenID, expectedResource) {
		return randomWalkTokenEventCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return cursor, nil
}

func validRandomWalkTokenEventCursor(
	cursor randomWalkTokenEventCursor,
	expectedTokenID int64,
	expectedResource randomWalkResource,
) bool {
	return cursor.Version == randomWalkCursorVersion &&
		cursor.EventLogID >= 1 &&
		cursor.TokenID >= 0 &&
		cursor.TokenID == expectedTokenID &&
		cursor.Resource == expectedResource &&
		validRandomWalkTokenEventResource(cursor.Resource) &&
		validRandomWalkTokenEventResource(expectedResource)
}

// randomWalkOfferBookCursor identifies the last offer returned by the live
// order book under one sort. PriceWei participates only in the two price
// orders.
type randomWalkOfferBookCursor struct {
	Version    int                 `json:"v"`
	Resource   randomWalkResource  `json:"k"`
	Sort       RandomWalkOfferSort `json:"s"`
	PriceWei   string              `json:"p,omitempty"`
	EventLogID int64               `json:"e"`
}

func encodeRandomWalkOfferBookCursor(cursor randomWalkOfferBookCursor) (string, error) {
	if !validRandomWalkOfferBookCursor(cursor, cursor.Sort) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRandomWalkCursor, "random walk offer book cursor")
}

func decodeRandomWalkOfferBookCursor(
	encoded string,
	expectedSort RandomWalkOfferSort,
) (randomWalkOfferBookCursor, error) {
	cursor, err := decodeOpaqueCursor[randomWalkOfferBookCursor](encoded, errInvalidRandomWalkCursor)
	if err != nil {
		return randomWalkOfferBookCursor{}, err
	}
	if !validRandomWalkOfferBookCursor(cursor, expectedSort) {
		return randomWalkOfferBookCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return cursor, nil
}

func validRandomWalkOfferBookCursor(
	cursor randomWalkOfferBookCursor,
	expectedSort RandomWalkOfferSort,
) bool {
	if cursor.Version != randomWalkCursorVersion ||
		cursor.Resource != randomWalkResourceOfferBook ||
		cursor.EventLogID < 1 ||
		!cursor.Sort.Valid() || cursor.Sort != expectedSort {
		return false
	}
	switch cursor.Sort {
	case PriceAsc, PriceDesc:
		return validWeiAmountString(cursor.PriceWei)
	default:
		return cursor.PriceWei == ""
	}
}

// validWeiAmountString reports whether value is a plain non-negative
// decimal integer.
func validWeiAmountString(value string) bool {
	if value == "" {
		return false
	}
	for _, r := range value {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// randomWalkLedgerCursor identifies the last event returned by one of the
// newest-first RandomWalk ledgers (offer history, trades, withdrawals and
// the wallet-scoped offer view). Address is set exactly for the
// wallet-scoped resource.
type randomWalkLedgerCursor struct {
	Version    int                `json:"v"`
	Resource   randomWalkResource `json:"k"`
	Address    string             `json:"a,omitempty"`
	EventLogID int64              `json:"e"`
}

func validRandomWalkLedgerResource(resource randomWalkResource) bool {
	switch resource {
	case randomWalkResourceOfferHistory,
		randomWalkResourceTrades,
		randomWalkResourceWithdrawals,
		randomWalkResourceUserOffers:
		return true
	default:
		return false
	}
}

func encodeRandomWalkLedgerCursor(cursor randomWalkLedgerCursor) (string, error) {
	if !validRandomWalkLedgerCursor(cursor, cursor.Resource, cursor.Address) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRandomWalkCursor, "random walk ledger cursor")
}

func decodeRandomWalkLedgerCursor(
	encoded string,
	expectedResource randomWalkResource,
	expectedAddress string,
) (randomWalkLedgerCursor, error) {
	cursor, err := decodeOpaqueCursor[randomWalkLedgerCursor](encoded, errInvalidRandomWalkCursor)
	if err != nil {
		return randomWalkLedgerCursor{}, err
	}
	if !validRandomWalkLedgerCursor(cursor, expectedResource, expectedAddress) {
		return randomWalkLedgerCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return cursor, nil
}

func validRandomWalkLedgerCursor(
	cursor randomWalkLedgerCursor,
	expectedResource randomWalkResource,
	expectedAddress string,
) bool {
	if cursor.Version != randomWalkCursorVersion ||
		cursor.EventLogID < 1 ||
		cursor.Resource != expectedResource ||
		!validRandomWalkLedgerResource(cursor.Resource) {
		return false
	}
	if cursor.Resource != randomWalkResourceUserOffers {
		return cursor.Address == "" && expectedAddress == ""
	}
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursorOK && expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}

// randomWalkRankingRatingCursor identifies the last row returned by the
// ascending (rating, tokenId) beauty-contest rating directory. Rating is a
// float64 like the stored Elo value; Go's JSON encoding round-trips it
// exactly, so the store resumes from the precise keyset position.
type randomWalkRankingRatingCursor struct {
	Version  int                `json:"v"`
	Resource randomWalkResource `json:"k"`
	Rating   float64            `json:"r"`
	TokenID  int64              `json:"t"`
}

func encodeRandomWalkRankingRatingCursor(cursor randomWalkRankingRatingCursor) (string, error) {
	if !validRandomWalkRankingRatingCursor(cursor) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRandomWalkCursor, "random walk ranking rating cursor")
}

func decodeRandomWalkRankingRatingCursor(encoded string) (randomWalkRankingRatingCursor, error) {
	cursor, err := decodeOpaqueCursor[randomWalkRankingRatingCursor](encoded, errInvalidRandomWalkCursor)
	if err != nil {
		return randomWalkRankingRatingCursor{}, err
	}
	if !validRandomWalkRankingRatingCursor(cursor) {
		return randomWalkRankingRatingCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return cursor, nil
}

// validRandomWalkRankingRatingCursor accepts any finite rating (Elo can
// legitimately go negative) at a non-negative token position. Non-finite
// ratings cannot appear: encoding/json refuses NaN and infinities on both
// the encode and decode sides.
func validRandomWalkRankingRatingCursor(cursor randomWalkRankingRatingCursor) bool {
	return cursor.Version == randomWalkCursorVersion &&
		cursor.Resource == randomWalkResourceRankingRatings &&
		cursor.TokenID >= 0
}

// randomWalkUserTokenCursor identifies the last token returned by the
// ascending owned-token directory of one wallet.
type randomWalkUserTokenCursor struct {
	Version  int                `json:"v"`
	Resource randomWalkResource `json:"k"`
	Address  string             `json:"a"`
	TokenID  int64              `json:"t"`
}

func encodeRandomWalkUserTokenCursor(cursor randomWalkUserTokenCursor) (string, error) {
	if !validRandomWalkUserTokenCursor(cursor, cursor.Address) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRandomWalkCursor, "random walk user token cursor")
}

func decodeRandomWalkUserTokenCursor(
	encoded, expectedAddress string,
) (randomWalkUserTokenCursor, error) {
	cursor, err := decodeOpaqueCursor[randomWalkUserTokenCursor](encoded, errInvalidRandomWalkCursor)
	if err != nil {
		return randomWalkUserTokenCursor{}, err
	}
	if !validRandomWalkUserTokenCursor(cursor, expectedAddress) {
		return randomWalkUserTokenCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRandomWalkCursor)
	}
	return cursor, nil
}

func validRandomWalkUserTokenCursor(
	cursor randomWalkUserTokenCursor,
	expectedAddress string,
) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == randomWalkCursorVersion &&
		cursor.Resource == randomWalkResourceUserTokens &&
		cursor.TokenID >= 0 &&
		cursorOK && expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}
