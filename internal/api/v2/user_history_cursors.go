package v2

import (
	"errors"
	"fmt"
)

// userEventResource scopes a userEventCursor to exactly one event-keyed user
// history so continuation cursors cannot cross resources or wallets.
type userEventResource string

const (
	userEventResourceRaffleEthDeposits userEventResource = "raffleEthDeposits"
	userEventResourceRaffleNftWins     userEventResource = "raffleNftWins"
	userEventResourceEthDonations      userEventResource = "ethDonations"
	userEventResourceErc20Donations    userEventResource = "erc20Donations"
	userEventResourceNftDonations      userEventResource = "nftDonations"
	userEventResourceDonatedNfts       userEventResource = "donatedNfts"
)

const userEventCursorVersion = 1

var errInvalidUserEventCursor = errors.New("invalid user event cursor")

type userEventCursor struct {
	Version    int               `json:"v"`
	Address    string            `json:"a"`
	Resource   userEventResource `json:"k"`
	EventLogID int64             `json:"e"`
}

func encodeUserEventCursor(cursor userEventCursor) (string, error) {
	if !validUserEventCursor(cursor, cursor.Address, cursor.Resource) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidUserEventCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidUserEventCursor, "user event cursor")
}

func decodeUserEventCursor(
	encoded string,
	expectedAddress string,
	expectedResource userEventResource,
) (userEventCursor, error) {
	cursor, err := decodeOpaqueCursor[userEventCursor](encoded, errInvalidUserEventCursor)
	if err != nil {
		return userEventCursor{}, err
	}
	if !validUserEventCursor(cursor, expectedAddress, expectedResource) {
		return userEventCursor{}, fmt.Errorf("%w: invalid fields", errInvalidUserEventCursor)
	}
	return cursor, nil
}

func validUserEventCursor(
	cursor userEventCursor,
	expectedAddress string,
	expectedResource userEventResource,
) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == userEventCursorVersion &&
		cursor.EventLogID >= 1 &&
		cursor.Resource == expectedResource &&
		validUserEventResource(cursor.Resource) &&
		cursorOK &&
		expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}

func validUserEventResource(resource userEventResource) bool {
	switch resource {
	case userEventResourceRaffleEthDeposits,
		userEventResourceRaffleNftWins,
		userEventResourceEthDonations,
		userEventResourceErc20Donations,
		userEventResourceNftDonations,
		userEventResourceDonatedNfts:
		return true
	default:
		return false
	}
}

const userPrizeCursorVersion = 1

var errInvalidUserPrizeCursor = errors.New("invalid user prize cursor")

type userPrizeCursor struct {
	Version     int    `json:"v"`
	Address     string `json:"a"`
	Round       int64  `json:"r"`
	PrizeType   int64  `json:"t"`
	WinnerIndex int64  `json:"w"`
}

func encodeUserPrizeCursor(cursor userPrizeCursor) (string, error) {
	if !validUserPrizeCursor(cursor, cursor.Address) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidUserPrizeCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidUserPrizeCursor, "user prize cursor")
}

func decodeUserPrizeCursor(encoded, expectedAddress string) (userPrizeCursor, error) {
	cursor, err := decodeOpaqueCursor[userPrizeCursor](encoded, errInvalidUserPrizeCursor)
	if err != nil {
		return userPrizeCursor{}, err
	}
	if !validUserPrizeCursor(cursor, expectedAddress) {
		return userPrizeCursor{}, fmt.Errorf("%w: invalid fields", errInvalidUserPrizeCursor)
	}
	return cursor, nil
}

func validUserPrizeCursor(cursor userPrizeCursor, expectedAddress string) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == userPrizeCursorVersion &&
		cursor.Round >= 0 &&
		cursor.PrizeType >= 0 &&
		cursor.PrizeType <= 15 &&
		cursor.WinnerIndex >= 0 &&
		cursorOK &&
		expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}

const userDonatedErc20CursorVersion = 1

var errInvalidUserDonatedErc20Cursor = errors.New("invalid user donated erc20 cursor")

// userDonatedErc20Cursor keys on cg_erc20_donation_stats' (round, token)
// primary key. The token component is an internal address ID that only ever
// travels opaquely inside the cursor.
type userDonatedErc20Cursor struct {
	Version int    `json:"v"`
	Address string `json:"a"`
	Round   int64  `json:"r"`
	TokenID int64  `json:"t"`
}

func encodeUserDonatedErc20Cursor(cursor userDonatedErc20Cursor) (string, error) {
	if !validUserDonatedErc20Cursor(cursor, cursor.Address) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidUserDonatedErc20Cursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidUserDonatedErc20Cursor, "user donated erc20 cursor")
}

func decodeUserDonatedErc20Cursor(encoded, expectedAddress string) (userDonatedErc20Cursor, error) {
	cursor, err := decodeOpaqueCursor[userDonatedErc20Cursor](encoded, errInvalidUserDonatedErc20Cursor)
	if err != nil {
		return userDonatedErc20Cursor{}, err
	}
	if !validUserDonatedErc20Cursor(cursor, expectedAddress) {
		return userDonatedErc20Cursor{}, fmt.Errorf("%w: invalid fields", errInvalidUserDonatedErc20Cursor)
	}
	return cursor, nil
}

func validUserDonatedErc20Cursor(cursor userDonatedErc20Cursor, expectedAddress string) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == userDonatedErc20CursorVersion &&
		cursor.Round >= 0 &&
		cursor.TokenID >= 1 &&
		cursorOK &&
		expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}
