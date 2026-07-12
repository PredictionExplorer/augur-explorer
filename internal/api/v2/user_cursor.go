package v2

import (
	"errors"
	"fmt"
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

const userBidCursorVersion = 1

var errInvalidUserBidCursor = errors.New("invalid user bid cursor")

type userBidCursor struct {
	Version    int    `json:"v"`
	Address    string `json:"a"`
	EventLogID int64  `json:"e"`
}

func encodeUserBidCursor(cursor userBidCursor) (string, error) {
	if !validUserBidCursor(cursor, cursor.Address) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidUserBidCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidUserBidCursor, "user bid cursor")
}

func decodeUserBidCursor(encoded, expectedAddress string) (userBidCursor, error) {
	cursor, err := decodeOpaqueCursor[userBidCursor](encoded, errInvalidUserBidCursor)
	if err != nil {
		return userBidCursor{}, err
	}
	if !validUserBidCursor(cursor, expectedAddress) {
		return userBidCursor{}, fmt.Errorf("%w: invalid fields", errInvalidUserBidCursor)
	}
	return cursor, nil
}

func validUserBidCursor(cursor userBidCursor, expectedAddress string) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == userBidCursorVersion &&
		cursor.EventLogID >= 1 &&
		cursorOK &&
		expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}

func normalizedUserAddress(address string) (string, bool) {
	if !ethcommon.IsHexAddress(address) {
		return "", false
	}
	return strings.ToLower(ethcommon.HexToAddress(address).Hex()), true
}
