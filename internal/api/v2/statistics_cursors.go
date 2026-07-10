package v2

import (
	"errors"
	"fmt"
)

const (
	roiLeaderboardCursorVersion = 1
	claimSummaryCursorVersion   = 1
	claimDetailCursorVersion    = 1
)

var (
	errInvalidROILeaderboardCursor = errors.New("invalid ROI leaderboard cursor")
	errInvalidClaimSummaryCursor   = errors.New("invalid claim summary cursor")
	errInvalidClaimDetailCursor    = errors.New("invalid claim detail cursor")
)

type roiLeaderboardCursor struct {
	Version   int                `json:"v"`
	Sort      RoiLeaderboardSort `json:"s"`
	MinBids   int                `json:"m"`
	SortValue string             `json:"k"`
	Secondary int64              `json:"x"`
	BidderAid int64              `json:"a"`
}

func encodeROILeaderboardCursor(cursor roiLeaderboardCursor) (string, error) {
	if !validROILeaderboardCursor(cursor, cursor.Sort, cursor.MinBids) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidROILeaderboardCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidROILeaderboardCursor, "ROI leaderboard cursor")
}

func decodeROILeaderboardCursor(
	encoded string,
	expectedSort RoiLeaderboardSort,
	expectedMinBids int,
) (roiLeaderboardCursor, error) {
	cursor, err := decodeOpaqueCursor[roiLeaderboardCursor](encoded, errInvalidROILeaderboardCursor)
	if err != nil {
		return roiLeaderboardCursor{}, err
	}
	if !validROILeaderboardCursor(cursor, expectedSort, expectedMinBids) {
		return roiLeaderboardCursor{}, fmt.Errorf("%w: invalid fields", errInvalidROILeaderboardCursor)
	}
	return cursor, nil
}

func validROILeaderboardCursor(
	cursor roiLeaderboardCursor,
	expectedSort RoiLeaderboardSort,
	expectedMinBids int,
) bool {
	if cursor.Version != roiLeaderboardCursorVersion ||
		cursor.Sort != expectedSort ||
		!cursor.Sort.Valid() ||
		cursor.MinBids != expectedMinBids ||
		cursor.MinBids < 0 ||
		cursor.BidderAid < 1 {
		return false
	}
	if _, err := canonicalDecimal(cursor.SortValue, true); err != nil {
		return false
	}
	if cursor.Sort == WinRate {
		return cursor.Secondary >= 0
	}
	return cursor.Secondary == 0
}

type claimSummaryCursor struct {
	Version    int   `json:"v"`
	Round      int64 `json:"r"`
	EventLogID int64 `json:"e"`
}

func encodeClaimSummaryCursor(cursor claimSummaryCursor) (string, error) {
	if !validClaimSummaryCursor(cursor) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidClaimSummaryCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidClaimSummaryCursor, "claim summary cursor")
}

func decodeClaimSummaryCursor(encoded string) (claimSummaryCursor, error) {
	cursor, err := decodeOpaqueCursor[claimSummaryCursor](encoded, errInvalidClaimSummaryCursor)
	if err != nil {
		return claimSummaryCursor{}, err
	}
	if !validClaimSummaryCursor(cursor) {
		return claimSummaryCursor{}, fmt.Errorf("%w: invalid fields", errInvalidClaimSummaryCursor)
	}
	return cursor, nil
}

func validClaimSummaryCursor(cursor claimSummaryCursor) bool {
	return cursor.Version == claimSummaryCursorVersion &&
		cursor.Round >= 0 &&
		cursor.EventLogID >= 1
}

type claimDetailSection string

const (
	claimDetailTransactions claimDetailSection = "transactions"
	claimDetailAttached     claimDetailSection = "attached"
	claimDetailUnclaimed    claimDetailSection = "unclaimed"
)

type claimDetailCursor struct {
	Version    int                `json:"v"`
	Round      int64              `json:"r"`
	Section    claimDetailSection `json:"s"`
	EventLogID int64              `json:"e"`
	Segment    int                `json:"g"`
	Key        int64              `json:"k"`
}

func encodeClaimDetailCursor(cursor claimDetailCursor) (string, error) {
	if !validClaimDetailCursor(cursor, cursor.Round, cursor.Section) {
		return "", fmt.Errorf("%w: invalid fields", errInvalidClaimDetailCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidClaimDetailCursor, "claim detail cursor")
}

func decodeClaimDetailCursor(
	encoded string,
	expectedRound int64,
	expectedSection claimDetailSection,
) (claimDetailCursor, error) {
	cursor, err := decodeOpaqueCursor[claimDetailCursor](encoded, errInvalidClaimDetailCursor)
	if err != nil {
		return claimDetailCursor{}, err
	}
	if !validClaimDetailCursor(cursor, expectedRound, expectedSection) {
		return claimDetailCursor{}, fmt.Errorf("%w: invalid fields", errInvalidClaimDetailCursor)
	}
	return cursor, nil
}

func validClaimDetailCursor(
	cursor claimDetailCursor,
	expectedRound int64,
	expectedSection claimDetailSection,
) bool {
	if cursor.Version != claimDetailCursorVersion ||
		cursor.Round != expectedRound ||
		cursor.Round < 0 ||
		cursor.Section != expectedSection {
		return false
	}
	switch cursor.Section {
	case claimDetailTransactions, claimDetailAttached:
		return cursor.EventLogID >= 1 && cursor.Segment == 0 && cursor.Key == 0
	case claimDetailUnclaimed:
		return cursor.EventLogID == 0 &&
			cursor.Segment >= 0 && cursor.Segment <= 2 &&
			cursor.Key >= 1
	default:
		return false
	}
}
