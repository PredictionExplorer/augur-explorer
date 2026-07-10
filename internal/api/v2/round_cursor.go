package v2

import (
	"errors"
	"fmt"
)

const roundCursorVersion = 1

var errInvalidRoundCursor = errors.New("invalid round cursor")

type roundCursor struct {
	Version    int   `json:"v"`
	RoundNum   int64 `json:"r"`
	EventLogID int64 `json:"e"`
}

func encodeRoundCursor(cursor roundCursor) (string, error) {
	if cursor.Version != roundCursorVersion || cursor.RoundNum < 0 || cursor.EventLogID < 1 {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRoundCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRoundCursor, "round cursor")
}

func decodeRoundCursor(encoded string) (roundCursor, error) {
	cursor, err := decodeOpaqueCursor[roundCursor](encoded, errInvalidRoundCursor)
	if err != nil {
		return roundCursor{}, err
	}
	if cursor.Version != roundCursorVersion || cursor.RoundNum < 0 || cursor.EventLogID < 1 {
		return roundCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRoundCursor)
	}
	return cursor, nil
}
