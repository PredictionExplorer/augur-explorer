package v2

import (
	"errors"
	"fmt"
)

const globalTokenCursorVersion = 1

var errInvalidGlobalTokenCursor = errors.New("invalid global token cursor")

// maxTokenNameSearchLength bounds the name search term; the OpenAPI
// contract declares the same limit.
const maxTokenNameSearchLength = 64

// globalTokenFilterScope reproduces the request's filter inside the cursor
// so a continuation cannot cross into another filtered view of the
// directory. At most one of Named and Name is set.
type globalTokenFilterScope struct {
	Named bool   `json:"n,omitempty"`
	Name  string `json:"q,omitempty"`
}

func (f globalTokenFilterScope) valid() bool {
	if f.Named && f.Name != "" {
		return false
	}
	return len(f.Name) <= maxTokenNameSearchLength
}

// globalTokenCursor identifies the last token returned by the descending
// global Cosmic Signature directory under one filter.
type globalTokenCursor struct {
	Version int                    `json:"v"`
	Filter  globalTokenFilterScope `json:"f"`
	TokenID int64                  `json:"t"`
}

func encodeGlobalTokenCursor(cursor globalTokenCursor) (string, error) {
	if !validGlobalTokenCursor(cursor, cursor.Filter) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidGlobalTokenCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidGlobalTokenCursor, "global token cursor")
}

func decodeGlobalTokenCursor(
	encoded string,
	expectedFilter globalTokenFilterScope,
) (globalTokenCursor, error) {
	cursor, err := decodeOpaqueCursor[globalTokenCursor](encoded, errInvalidGlobalTokenCursor)
	if err != nil {
		return globalTokenCursor{}, err
	}
	if !validGlobalTokenCursor(cursor, expectedFilter) {
		return globalTokenCursor{}, fmt.Errorf("%w: invalid fields", errInvalidGlobalTokenCursor)
	}
	return cursor, nil
}

func validGlobalTokenCursor(cursor globalTokenCursor, expectedFilter globalTokenFilterScope) bool {
	return cursor.Version == globalTokenCursorVersion &&
		cursor.TokenID >= 0 &&
		cursor.Filter.valid() &&
		expectedFilter.valid() &&
		cursor.Filter == expectedFilter
}

// tokenEventResource scopes a tokenEventCursor to exactly one event-keyed
// per-token history so continuation cursors cannot cross resources or
// tokens.
type tokenEventResource string

const (
	tokenEventResourceNameHistory tokenEventResource = "nameHistory"
	tokenEventResourceTransfers   tokenEventResource = "transfers"
)

func validTokenEventResource(resource tokenEventResource) bool {
	switch resource {
	case tokenEventResourceNameHistory, tokenEventResourceTransfers:
		return true
	default:
		return false
	}
}

const tokenEventCursorVersion = 1

var errInvalidTokenEventCursor = errors.New("invalid token event cursor")

// tokenEventCursor identifies the last event returned by one token-scoped,
// newest-first event page.
type tokenEventCursor struct {
	Version    int                `json:"v"`
	Resource   tokenEventResource `json:"k"`
	TokenID    int64              `json:"t"`
	EventLogID int64              `json:"e"`
}

func encodeTokenEventCursor(cursor tokenEventCursor) (string, error) {
	if !validTokenEventCursor(cursor, cursor.TokenID, cursor.Resource) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidTokenEventCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidTokenEventCursor, "token event cursor")
}

func decodeTokenEventCursor(
	encoded string,
	expectedTokenID int64,
	expectedResource tokenEventResource,
) (tokenEventCursor, error) {
	cursor, err := decodeOpaqueCursor[tokenEventCursor](encoded, errInvalidTokenEventCursor)
	if err != nil {
		return tokenEventCursor{}, err
	}
	if !validTokenEventCursor(cursor, expectedTokenID, expectedResource) {
		return tokenEventCursor{}, fmt.Errorf("%w: invalid fields", errInvalidTokenEventCursor)
	}
	return cursor, nil
}

func validTokenEventCursor(
	cursor tokenEventCursor,
	expectedTokenID int64,
	expectedResource tokenEventResource,
) bool {
	return cursor.Version == tokenEventCursorVersion &&
		cursor.EventLogID >= 1 &&
		cursor.TokenID >= 0 &&
		cursor.TokenID == expectedTokenID &&
		cursor.Resource == expectedResource &&
		validTokenEventResource(cursor.Resource)
}

const supplyChangeCursorVersion = 1

var errInvalidSupplyChangeCursor = errors.New("invalid supply change cursor")

// supplyChangeCursor identifies the last bid returned by the ascending
// per-bid Cosmic Token supply ledger. The "s" key keeps the payload's key
// set unique among cursor types.
type supplyChangeCursor struct {
	Version    int   `json:"v"`
	EventLogID int64 `json:"s"`
}

func encodeSupplyChangeCursor(cursor supplyChangeCursor) (string, error) {
	if !validSupplyChangeCursor(cursor) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidSupplyChangeCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidSupplyChangeCursor, "supply change cursor")
}

func decodeSupplyChangeCursor(encoded string) (supplyChangeCursor, error) {
	cursor, err := decodeOpaqueCursor[supplyChangeCursor](encoded, errInvalidSupplyChangeCursor)
	if err != nil {
		return supplyChangeCursor{}, err
	}
	if !validSupplyChangeCursor(cursor) {
		return supplyChangeCursor{}, fmt.Errorf("%w: invalid fields", errInvalidSupplyChangeCursor)
	}
	return cursor, nil
}

func validSupplyChangeCursor(cursor supplyChangeCursor) bool {
	return cursor.Version == supplyChangeCursorVersion && cursor.EventLogID >= 1
}

const globalMarketingCursorVersion = 1

var errInvalidGlobalMarketingCursor = errors.New("invalid global marketing cursor")

// globalMarketingCursor identifies the last reward returned by the
// descending global marketing ledger. The "m" key keeps the payload's key
// set unique among cursor types.
type globalMarketingCursor struct {
	Version    int   `json:"v"`
	EventLogID int64 `json:"m"`
}

func encodeGlobalMarketingCursor(cursor globalMarketingCursor) (string, error) {
	if !validGlobalMarketingCursor(cursor) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidGlobalMarketingCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidGlobalMarketingCursor, "global marketing cursor")
}

func decodeGlobalMarketingCursor(encoded string) (globalMarketingCursor, error) {
	cursor, err := decodeOpaqueCursor[globalMarketingCursor](encoded, errInvalidGlobalMarketingCursor)
	if err != nil {
		return globalMarketingCursor{}, err
	}
	if !validGlobalMarketingCursor(cursor) {
		return globalMarketingCursor{}, fmt.Errorf("%w: invalid fields", errInvalidGlobalMarketingCursor)
	}
	return cursor, nil
}

func validGlobalMarketingCursor(cursor globalMarketingCursor) bool {
	return cursor.Version == globalMarketingCursorVersion && cursor.EventLogID >= 1
}
