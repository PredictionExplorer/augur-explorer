package v2

import (
	"errors"
	"fmt"
	"math/big"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const participantCursorVersion = 1

var errInvalidParticipantCursor = errors.New("invalid participant cursor")

type participantCursor struct {
	Version   int                     `json:"v"`
	Kind      cgstore.ParticipantKind `json:"k"`
	SortValue string                  `json:"s"`
	AddressID int64                   `json:"a"`
}

func encodeParticipantCursor(cursor participantCursor) (string, error) {
	if !validParticipantCursor(cursor, cursor.Kind) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidParticipantCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidParticipantCursor, "participant cursor")
}

func decodeParticipantCursor(
	encoded string,
	expectedKind cgstore.ParticipantKind,
) (participantCursor, error) {
	cursor, err := decodeOpaqueCursor[participantCursor](encoded, errInvalidParticipantCursor)
	if err != nil {
		return participantCursor{}, err
	}
	if !validParticipantCursor(cursor, expectedKind) {
		return participantCursor{}, fmt.Errorf("%w: invalid fields", errInvalidParticipantCursor)
	}
	return cursor, nil
}

func validParticipantCursor(cursor participantCursor, expectedKind cgstore.ParticipantKind) bool {
	if cursor.Version != participantCursorVersion ||
		cursor.Kind != expectedKind ||
		!validParticipantKind(cursor.Kind) ||
		cursor.AddressID < 1 {
		return false
	}
	value, err := requiredAmount(cursor.SortValue)
	if err != nil || value != cursor.SortValue {
		return false
	}
	switch cursor.Kind {
	case cgstore.ParticipantBidders,
		cgstore.ParticipantWinners,
		cgstore.ParticipantRandomWalkStakers,
		cgstore.ParticipantDualStakers,
		cgstore.ParticipantCsTokenHolders:
		number, ok := new(big.Int).SetString(value, 10)
		return ok && number.IsInt64()
	default:
		return true
	}
}

func validParticipantKind(kind cgstore.ParticipantKind) bool {
	switch kind {
	case cgstore.ParticipantBidders,
		cgstore.ParticipantWinners,
		cgstore.ParticipantDonors,
		cgstore.ParticipantCSTStakers,
		cgstore.ParticipantRandomWalkStakers,
		cgstore.ParticipantDualStakers,
		cgstore.ParticipantCsTokenHolders,
		cgstore.ParticipantCosmicTokenHolders:
		return true
	default:
		return false
	}
}
