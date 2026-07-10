package v2

import (
	"errors"
	"fmt"
)

const roundDonationCursorVersion = 1

var errInvalidRoundDonationCursor = errors.New("invalid round donation cursor")

type roundDonationResource string

const (
	roundDonationResourceETH   roundDonationResource = "eth"
	roundDonationResourceERC20 roundDonationResource = "erc20"
	roundDonationResourceNFT   roundDonationResource = "nft"
)

type roundDonationCursor struct {
	Version    int                   `json:"v"`
	Round      int64                 `json:"r"`
	Resource   roundDonationResource `json:"k"`
	EventLogID int64                 `json:"e"`
}

func encodeRoundDonationCursor(cursor roundDonationCursor) (string, error) {
	if !validRoundDonationCursor(cursor, cursor.Round, cursor.Resource) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRoundDonationCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidRoundDonationCursor, "round donation cursor")
}

func decodeRoundDonationCursor(
	encoded string,
	expectedRound int64,
	expectedResource roundDonationResource,
) (roundDonationCursor, error) {
	cursor, err := decodeOpaqueCursor[roundDonationCursor](encoded, errInvalidRoundDonationCursor)
	if err != nil {
		return roundDonationCursor{}, err
	}
	if !validRoundDonationCursor(cursor, expectedRound, expectedResource) {
		return roundDonationCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRoundDonationCursor)
	}
	return cursor, nil
}

func validRoundDonationCursor(
	cursor roundDonationCursor,
	expectedRound int64,
	expectedResource roundDonationResource,
) bool {
	return cursor.Version == roundDonationCursorVersion &&
		cursor.Round == expectedRound &&
		cursor.Round >= 0 &&
		cursor.Resource == expectedResource &&
		validRoundDonationResource(cursor.Resource) &&
		cursor.EventLogID >= 1
}

func validRoundDonationResource(resource roundDonationResource) bool {
	switch resource {
	case roundDonationResourceETH, roundDonationResourceERC20, roundDonationResourceNFT:
		return true
	default:
		return false
	}
}
