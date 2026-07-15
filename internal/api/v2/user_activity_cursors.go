package v2

import (
	"errors"
	"fmt"
)

const userOwnedTokenCursorVersion = 1

var errInvalidUserOwnedTokenCursor = errors.New("invalid user owned token cursor")

// userOwnedTokenCursor identifies the last token returned by the ascending
// owned Cosmic Signature token directory. The payload's key set is unique
// among cursor types, so cursors from other resources fail structurally.
type userOwnedTokenCursor struct {
	Version int    `json:"v"`
	Address string `json:"a"`
	TokenID int64  `json:"t"`
}

func encodeUserOwnedTokenCursor(cursor userOwnedTokenCursor) (string, error) {
	if !validUserOwnedTokenCursor(cursor, cursor.Address) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidUserOwnedTokenCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidUserOwnedTokenCursor, "user owned token cursor")
}

func decodeUserOwnedTokenCursor(encoded, expectedAddress string) (userOwnedTokenCursor, error) {
	cursor, err := decodeOpaqueCursor[userOwnedTokenCursor](encoded, errInvalidUserOwnedTokenCursor)
	if err != nil {
		return userOwnedTokenCursor{}, err
	}
	if !validUserOwnedTokenCursor(cursor, expectedAddress) {
		return userOwnedTokenCursor{}, fmt.Errorf("%w: invalid fields", errInvalidUserOwnedTokenCursor)
	}
	return cursor, nil
}

func validUserOwnedTokenCursor(cursor userOwnedTokenCursor, expectedAddress string) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == userOwnedTokenCursorVersion &&
		cursor.TokenID >= 0 &&
		cursorOK &&
		expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}
