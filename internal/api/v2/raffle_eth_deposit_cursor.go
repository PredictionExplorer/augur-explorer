package v2

import (
	"errors"
	"fmt"
)

const raffleEthDepositCursorVersion = 1

var errInvalidRaffleEthDepositCursor = errors.New("invalid raffle ETH deposit cursor")

type raffleEthDepositCursor struct {
	Version     int   `json:"v"`
	Round       int64 `json:"r"`
	WinnerIndex int64 `json:"w"`
	EventLogID  int64 `json:"e"`
}

func encodeRaffleEthDepositCursor(cursor raffleEthDepositCursor) (string, error) {
	if !validRaffleEthDepositCursor(cursor, cursor.Round) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRaffleEthDepositCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRaffleEthDepositCursor, "raffle ETH deposit cursor")
}

func decodeRaffleEthDepositCursor(encoded string, expectedRound int64) (raffleEthDepositCursor, error) {
	cursor, err := decodeOpaqueCursor[raffleEthDepositCursor](encoded, errInvalidRaffleEthDepositCursor)
	if err != nil {
		return raffleEthDepositCursor{}, err
	}
	if !validRaffleEthDepositCursor(cursor, expectedRound) {
		return raffleEthDepositCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRaffleEthDepositCursor)
	}
	return cursor, nil
}

func validRaffleEthDepositCursor(cursor raffleEthDepositCursor, expectedRound int64) bool {
	return cursor.Version == raffleEthDepositCursorVersion &&
		cursor.Round == expectedRound &&
		cursor.Round >= 0 &&
		cursor.WinnerIndex >= 0 &&
		cursor.EventLogID >= 1
}
