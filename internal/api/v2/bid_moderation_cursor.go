package v2

import (
	"errors"
	"fmt"
)

const bidBanCursorVersion = 1

var errInvalidBidBanCursor = errors.New("invalid bid-ban cursor")

type bidBanCursor struct {
	Version int   `json:"v"`
	ID      int64 `json:"i"`
}

func encodeBidBanCursor(cursor bidBanCursor) (string, error) {
	if cursor.Version != bidBanCursorVersion || cursor.ID < 1 {
		return "", fmt.Errorf("%w: invalid fields", errInvalidBidBanCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidBidBanCursor, "bid-ban cursor")
}

func decodeBidBanCursor(encoded string) (bidBanCursor, error) {
	cursor, err := decodeOpaqueCursor[bidBanCursor](encoded, errInvalidBidBanCursor)
	if err != nil {
		return bidBanCursor{}, err
	}
	if cursor.Version != bidBanCursorVersion || cursor.ID < 1 {
		return bidBanCursor{}, fmt.Errorf("%w: invalid fields", errInvalidBidBanCursor)
	}
	return cursor, nil
}
