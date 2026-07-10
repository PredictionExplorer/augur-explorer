package v2

import (
	"errors"
	"fmt"
)

const prizeCursorVersion = 1

var errInvalidPrizeCursor = errors.New("invalid prize cursor")

type prizeCursor struct {
	Version     int   `json:"v"`
	Round       int64 `json:"r"`
	PrizeType   int64 `json:"t"`
	WinnerIndex int64 `json:"w"`
}

func encodePrizeCursor(cursor prizeCursor) (string, error) {
	if !validPrizeCursor(cursor, cursor.Round) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidPrizeCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidPrizeCursor, "prize cursor")
}

func decodePrizeCursor(encoded string, expectedRound int64) (prizeCursor, error) {
	cursor, err := decodeOpaqueCursor[prizeCursor](encoded, errInvalidPrizeCursor)
	if err != nil {
		return prizeCursor{}, err
	}
	if !validPrizeCursor(cursor, expectedRound) {
		return prizeCursor{}, fmt.Errorf("%w: invalid fields", errInvalidPrizeCursor)
	}
	return cursor, nil
}

func validPrizeCursor(cursor prizeCursor, expectedRound int64) bool {
	return cursor.Version == prizeCursorVersion &&
		cursor.Round == expectedRound &&
		cursor.Round >= 0 &&
		cursor.PrizeType >= 0 &&
		cursor.PrizeType <= 15 &&
		cursor.WinnerIndex >= 0
}
