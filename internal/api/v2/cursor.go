package v2

import (
	"errors"
	"fmt"
)

const bidCursorVersion = 1

var errInvalidCursor = errors.New("invalid cursor")

type bidCursor struct {
	Version     int   `json:"v"`
	Round       int64 `json:"r"`
	BidPosition int64 `json:"p"`
	EventLogID  int64 `json:"e"`
}

func encodeBidCursor(cursor bidCursor) (string, error) {
	if cursor.Version != bidCursorVersion || cursor.Round < 0 || cursor.BidPosition < 1 || cursor.EventLogID < 1 {
		return "", fmt.Errorf("%w: invalid fields", errInvalidCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidCursor, "cursor")
}

func decodeBidCursor(encoded string, expectedRound int64) (bidCursor, error) {
	cursor, err := decodeOpaqueCursor[bidCursor](encoded, errInvalidCursor)
	if err != nil {
		return bidCursor{}, err
	}
	if cursor.Version != bidCursorVersion ||
		cursor.Round != expectedRound ||
		cursor.Round < 0 ||
		cursor.BidPosition < 1 ||
		cursor.EventLogID < 1 {
		return bidCursor{}, fmt.Errorf("%w: invalid fields", errInvalidCursor)
	}
	return cursor, nil
}
