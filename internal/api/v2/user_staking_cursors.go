package v2

import (
	"errors"
	"fmt"
)

// userStakingTokenResource scopes a userStakingTokenCursor to exactly one
// ascending token-keyed staking collection.
type userStakingTokenResource string

// #nosec G101 -- cursor resource scope labels, not credentials.
const (
	userStakingTokenResourceCstStakedTokens userStakingTokenResource = "cstStakedTokens"
	userStakingTokenResourceRwStakedTokens  userStakingTokenResource = "randomWalkStakedTokens"
	userStakingTokenResourceCstTokenRewards userStakingTokenResource = "cstTokenRewards"
)

const userStakingTokenCursorVersion = 1

var errInvalidUserStakingTokenCursor = errors.New("invalid user staking token cursor")

// userStakingTokenCursor continues an ascending token-keyed staking page.
type userStakingTokenCursor struct {
	Version  int                      `json:"v"`
	Address  string                   `json:"a"`
	Resource userStakingTokenResource `json:"k"`
	TokenID  int64                    `json:"t"`
}

func encodeUserStakingTokenCursor(cursor userStakingTokenCursor) (string, error) {
	if !validUserStakingTokenCursor(cursor, cursor.Address, cursor.Resource) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidUserStakingTokenCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidUserStakingTokenCursor, "user staking token cursor")
}

func decodeUserStakingTokenCursor(
	encoded string,
	expectedAddress string,
	expectedResource userStakingTokenResource,
) (userStakingTokenCursor, error) {
	cursor, err := decodeOpaqueCursor[userStakingTokenCursor](encoded, errInvalidUserStakingTokenCursor)
	if err != nil {
		return userStakingTokenCursor{}, err
	}
	if !validUserStakingTokenCursor(cursor, expectedAddress, expectedResource) {
		return userStakingTokenCursor{}, fmt.Errorf("%w: invalid fields", errInvalidUserStakingTokenCursor)
	}
	return cursor, nil
}

func validUserStakingTokenCursor(
	cursor userStakingTokenCursor,
	expectedAddress string,
	expectedResource userStakingTokenResource,
) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == userStakingTokenCursorVersion &&
		cursor.TokenID >= 0 &&
		cursor.Resource == expectedResource &&
		validUserStakingTokenResource(cursor.Resource) &&
		cursorOK &&
		expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}

func validUserStakingTokenResource(resource userStakingTokenResource) bool {
	switch resource {
	case userStakingTokenResourceCstStakedTokens,
		userStakingTokenResourceRwStakedTokens,
		userStakingTokenResourceCstTokenRewards:
		return true
	default:
		return false
	}
}

const userStakingDepositCursorVersion = 1

var errInvalidUserStakingDepositCursor = errors.New("invalid user staking deposit cursor")

// userStakingDepositCursor continues the newest-first staking deposit
// ledger of one wallet.
type userStakingDepositCursor struct {
	Version   int    `json:"v"`
	Address   string `json:"a"`
	DepositID int64  `json:"d"`
}

func encodeUserStakingDepositCursor(cursor userStakingDepositCursor) (string, error) {
	if !validUserStakingDepositCursor(cursor, cursor.Address) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidUserStakingDepositCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidUserStakingDepositCursor, "user staking deposit cursor")
}

func decodeUserStakingDepositCursor(encoded, expectedAddress string) (userStakingDepositCursor, error) {
	cursor, err := decodeOpaqueCursor[userStakingDepositCursor](encoded, errInvalidUserStakingDepositCursor)
	if err != nil {
		return userStakingDepositCursor{}, err
	}
	if !validUserStakingDepositCursor(cursor, expectedAddress) {
		return userStakingDepositCursor{}, fmt.Errorf("%w: invalid fields", errInvalidUserStakingDepositCursor)
	}
	return cursor, nil
}

func validUserStakingDepositCursor(cursor userStakingDepositCursor, expectedAddress string) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == userStakingDepositCursorVersion &&
		cursor.DepositID >= 0 &&
		cursorOK &&
		expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}

const userStakingDepositRewardCursorVersion = 1

var errInvalidUserStakingDepositRewardCursor = errors.New("invalid user staking deposit reward cursor")

// userStakingDepositRewardCursor continues the ascending per-action reward
// page of one wallet inside one deposit; it is scoped to both.
type userStakingDepositRewardCursor struct {
	Version   int    `json:"v"`
	Address   string `json:"a"`
	DepositID int64  `json:"d"`
	ActionID  int64  `json:"s"`
}

func encodeUserStakingDepositRewardCursor(cursor userStakingDepositRewardCursor) (string, error) {
	if !validUserStakingDepositRewardCursor(cursor, cursor.Address, cursor.DepositID) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidUserStakingDepositRewardCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidUserStakingDepositRewardCursor, "user staking deposit reward cursor")
}

func decodeUserStakingDepositRewardCursor(
	encoded string,
	expectedAddress string,
	expectedDepositID int64,
) (userStakingDepositRewardCursor, error) {
	cursor, err := decodeOpaqueCursor[userStakingDepositRewardCursor](encoded, errInvalidUserStakingDepositRewardCursor)
	if err != nil {
		return userStakingDepositRewardCursor{}, err
	}
	if !validUserStakingDepositRewardCursor(cursor, expectedAddress, expectedDepositID) {
		return userStakingDepositRewardCursor{}, fmt.Errorf("%w: invalid fields", errInvalidUserStakingDepositRewardCursor)
	}
	return cursor, nil
}

func validUserStakingDepositRewardCursor(
	cursor userStakingDepositRewardCursor,
	expectedAddress string,
	expectedDepositID int64,
) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == userStakingDepositRewardCursorVersion &&
		cursor.DepositID >= 0 &&
		cursor.DepositID == expectedDepositID &&
		cursor.ActionID >= 0 &&
		cursorOK &&
		expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}

const userStakingTokenDepositCursorVersion = 1

var errInvalidUserStakingTokenDepositCursor = errors.New("invalid user staking token deposit cursor")

// userStakingTokenDepositCursor continues the ascending per-deposit reward
// page of one staked token; it is scoped to the wallet and the token.
type userStakingTokenDepositCursor struct {
	Version   int    `json:"v"`
	Address   string `json:"a"`
	TokenID   int64  `json:"t"`
	DepositID int64  `json:"d"`
}

func encodeUserStakingTokenDepositCursor(cursor userStakingTokenDepositCursor) (string, error) {
	if !validUserStakingTokenDepositCursor(cursor, cursor.Address, cursor.TokenID) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidUserStakingTokenDepositCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidUserStakingTokenDepositCursor, "user staking token deposit cursor")
}

func decodeUserStakingTokenDepositCursor(
	encoded string,
	expectedAddress string,
	expectedTokenID int64,
) (userStakingTokenDepositCursor, error) {
	cursor, err := decodeOpaqueCursor[userStakingTokenDepositCursor](encoded, errInvalidUserStakingTokenDepositCursor)
	if err != nil {
		return userStakingTokenDepositCursor{}, err
	}
	if !validUserStakingTokenDepositCursor(cursor, expectedAddress, expectedTokenID) {
		return userStakingTokenDepositCursor{}, fmt.Errorf("%w: invalid fields", errInvalidUserStakingTokenDepositCursor)
	}
	return cursor, nil
}

func validUserStakingTokenDepositCursor(
	cursor userStakingTokenDepositCursor,
	expectedAddress string,
	expectedTokenID int64,
) bool {
	cursorScope, cursorOK := normalizedUserAddress(cursor.Address)
	expectedScope, expectedOK := normalizedUserAddress(expectedAddress)
	return cursor.Version == userStakingTokenDepositCursorVersion &&
		cursor.TokenID >= 0 &&
		cursor.TokenID == expectedTokenID &&
		cursor.DepositID >= 0 &&
		cursorOK &&
		expectedOK &&
		cursor.Address == cursorScope &&
		cursorScope == expectedScope
}
