package v2

import (
	"errors"
	"fmt"
)

type globalStakingEventResource string

const (
	globalStakingEventCstActions   globalStakingEventResource = "cstActions"
	globalStakingEventRwalkActions globalStakingEventResource = "randomWalkActions"
	globalStakingEventCstRaffle    globalStakingEventResource = "cstRaffle"
	globalStakingEventRwalkRaffle  globalStakingEventResource = "randomWalkRaffle"
)

const globalStakingEventCursorVersion = 1

var errInvalidGlobalStakingEventCursor = errors.New("invalid global staking event cursor")

type globalStakingEventCursor struct {
	Version    int                        `json:"v"`
	Resource   globalStakingEventResource `json:"k"`
	EventLogID int64                      `json:"e"`
}

func encodeGlobalStakingEventCursor(cursor globalStakingEventCursor) (string, error) {
	if !validGlobalStakingEventCursor(cursor, cursor.Resource) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidGlobalStakingEventCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidGlobalStakingEventCursor, "global staking event cursor")
}

func decodeGlobalStakingEventCursor(
	encoded string,
	expected globalStakingEventResource,
) (globalStakingEventCursor, error) {
	cursor, err := decodeOpaqueCursor[globalStakingEventCursor](encoded, errInvalidGlobalStakingEventCursor)
	if err != nil {
		return globalStakingEventCursor{}, err
	}
	if !validGlobalStakingEventCursor(cursor, expected) {
		return globalStakingEventCursor{}, fmt.Errorf("%w: invalid fields", errInvalidGlobalStakingEventCursor)
	}
	return cursor, nil
}

func validGlobalStakingEventCursor(
	cursor globalStakingEventCursor,
	expected globalStakingEventResource,
) bool {
	return cursor.Version == globalStakingEventCursorVersion &&
		cursor.EventLogID >= 1 &&
		cursor.Resource == expected &&
		validGlobalStakingEventResource(cursor.Resource)
}

func validGlobalStakingEventResource(resource globalStakingEventResource) bool {
	switch resource {
	case globalStakingEventCstActions,
		globalStakingEventRwalkActions,
		globalStakingEventCstRaffle,
		globalStakingEventRwalkRaffle:
		return true
	default:
		return false
	}
}

type globalStakedTokenResource string

const (
	globalStakedTokenCst   globalStakedTokenResource = "cst"
	globalStakedTokenRwalk globalStakedTokenResource = "randomWalk"
)

const globalStakedTokenCursorVersion = 1

var errInvalidGlobalStakedTokenCursor = errors.New("invalid global staked token cursor")

type globalStakedTokenCursor struct {
	Version  int                       `json:"v"`
	Resource globalStakedTokenResource `json:"k"`
	TokenID  int64                     `json:"t"`
}

func encodeGlobalStakedTokenCursor(cursor globalStakedTokenCursor) (string, error) {
	if !validGlobalStakedTokenCursor(cursor, cursor.Resource) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidGlobalStakedTokenCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidGlobalStakedTokenCursor, "global staked token cursor")
}

func decodeGlobalStakedTokenCursor(
	encoded string,
	expected globalStakedTokenResource,
) (globalStakedTokenCursor, error) {
	cursor, err := decodeOpaqueCursor[globalStakedTokenCursor](encoded, errInvalidGlobalStakedTokenCursor)
	if err != nil {
		return globalStakedTokenCursor{}, err
	}
	if !validGlobalStakedTokenCursor(cursor, expected) {
		return globalStakedTokenCursor{}, fmt.Errorf("%w: invalid fields", errInvalidGlobalStakedTokenCursor)
	}
	return cursor, nil
}

func validGlobalStakedTokenCursor(
	cursor globalStakedTokenCursor,
	expected globalStakedTokenResource,
) bool {
	return cursor.Version == globalStakedTokenCursorVersion &&
		cursor.TokenID >= 0 &&
		cursor.Resource == expected &&
		validGlobalStakedTokenResource(cursor.Resource)
}

func validGlobalStakedTokenResource(resource globalStakedTokenResource) bool {
	switch resource {
	case globalStakedTokenCst, globalStakedTokenRwalk:
		return true
	default:
		return false
	}
}

const globalStakingDepositCursorVersion = 1

var errInvalidGlobalStakingDepositCursor = errors.New("invalid global staking deposit cursor")

type globalStakingDepositCursor struct {
	Version   int   `json:"v"`
	DepositID int64 `json:"d"`
}

func encodeGlobalStakingDepositCursor(cursor globalStakingDepositCursor) (string, error) {
	if !validGlobalStakingDepositCursor(cursor) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidGlobalStakingDepositCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidGlobalStakingDepositCursor, "global staking deposit cursor")
}

func decodeGlobalStakingDepositCursor(encoded string) (globalStakingDepositCursor, error) {
	cursor, err := decodeOpaqueCursor[globalStakingDepositCursor](encoded, errInvalidGlobalStakingDepositCursor)
	if err != nil {
		return globalStakingDepositCursor{}, err
	}
	if !validGlobalStakingDepositCursor(cursor) {
		return globalStakingDepositCursor{}, fmt.Errorf("%w: invalid fields", errInvalidGlobalStakingDepositCursor)
	}
	return cursor, nil
}

func validGlobalStakingDepositCursor(cursor globalStakingDepositCursor) bool {
	return cursor.Version == globalStakingDepositCursorVersion && cursor.DepositID >= 0
}

const roundStakingRewardCursorVersion = 1

var errInvalidRoundStakingRewardCursor = errors.New("invalid round staking reward cursor")

type roundStakingRewardCursor struct {
	Version   int   `json:"v"`
	Round     int64 `json:"r"`
	DepositID int64 `json:"d"`
	StakerAid int64 `json:"a"`
}

func encodeRoundStakingRewardCursor(cursor roundStakingRewardCursor) (string, error) {
	if !validRoundStakingRewardCursor(cursor, cursor.Round) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRoundStakingRewardCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRoundStakingRewardCursor, "round staking reward cursor")
}

func decodeRoundStakingRewardCursor(
	encoded string,
	expectedRound int64,
) (roundStakingRewardCursor, error) {
	cursor, err := decodeOpaqueCursor[roundStakingRewardCursor](encoded, errInvalidRoundStakingRewardCursor)
	if err != nil {
		return roundStakingRewardCursor{}, err
	}
	if !validRoundStakingRewardCursor(cursor, expectedRound) {
		return roundStakingRewardCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRoundStakingRewardCursor)
	}
	return cursor, nil
}

func validRoundStakingRewardCursor(cursor roundStakingRewardCursor, expectedRound int64) bool {
	return cursor.Version == roundStakingRewardCursorVersion &&
		cursor.Round >= 0 &&
		cursor.Round == expectedRound &&
		cursor.DepositID >= 0 &&
		cursor.StakerAid >= 1
}
