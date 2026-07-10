package v2

import (
	"errors"
	"fmt"
)

const raffleNftWinnerCursorVersion = 1

var errInvalidRaffleNftWinnerCursor = errors.New("invalid raffle NFT winner cursor")

type raffleNftWinnerCursor struct {
	Version     int           `json:"v"`
	Round       int64         `json:"r"`
	Pool        RaffleNftPool `json:"p"`
	WinnerIndex int64         `json:"w"`
	EventLogID  int64         `json:"e"`
}

func encodeRaffleNftWinnerCursor(cursor raffleNftWinnerCursor) (string, error) {
	if !validRaffleNftWinnerCursor(cursor, cursor.Round, cursor.Pool) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRaffleNftWinnerCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRaffleNftWinnerCursor, "raffle NFT winner cursor")
}

func decodeRaffleNftWinnerCursor(
	encoded string,
	expectedRound int64,
	expectedPool RaffleNftPool,
) (raffleNftWinnerCursor, error) {
	cursor, err := decodeOpaqueCursor[raffleNftWinnerCursor](encoded, errInvalidRaffleNftWinnerCursor)
	if err != nil {
		return raffleNftWinnerCursor{}, err
	}
	if !validRaffleNftWinnerCursor(cursor, expectedRound, expectedPool) {
		return raffleNftWinnerCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRaffleNftWinnerCursor)
	}
	return cursor, nil
}

func validRaffleNftWinnerCursor(
	cursor raffleNftWinnerCursor,
	expectedRound int64,
	expectedPool RaffleNftPool,
) bool {
	return cursor.Version == raffleNftWinnerCursorVersion &&
		cursor.Round == expectedRound &&
		cursor.Round >= 0 &&
		cursor.Pool == expectedPool &&
		cursor.Pool.Valid() &&
		cursor.WinnerIndex >= 0 &&
		cursor.EventLogID >= 1
}
