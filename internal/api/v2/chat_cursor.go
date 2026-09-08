package v2

import (
	"fmt"
	"time"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const chatCursorVersion = 1

type chatCursor struct {
	Version    int       `json:"v"`
	Round      int64     `json:"r"`
	Direction  string    `json:"d"`
	OccurredAt time.Time `json:"t"`
	EventLogID int64     `json:"e"`
	Revision   int64     `json:"g"`
}

func validChatCursor(cursor chatCursor) bool {
	return cursor.Version == chatCursorVersion && cursor.Round >= 0 && cursor.Revision > 0 &&
		(cursor.Direction == "older" || cursor.Direction == "after") && cursor.EventLogID >= 0 &&
		!cursor.OccurredAt.Before(time.Unix(0, 0)) &&
		(cursor.EventLogID > 0 || (cursor.Direction == "after" && cursor.OccurredAt.Equal(time.Unix(0, 0))))
}

func encodeChatCursor(round int64, direction string, position cgstore.ChatPosition, revision int64) (string, error) {
	cursor := chatCursor{
		Version: chatCursorVersion, Round: round, Direction: direction,
		OccurredAt: position.OccurredAt.UTC(), EventLogID: position.EventLogID, Revision: revision,
	}
	if !validChatCursor(cursor) {
		return "", fmt.Errorf("%w: invalid chat cursor fields", errInvalidCursor)
	}
	return encodeOpaqueCursor(cursor, errInvalidCursor, "chat cursor")
}

func decodeChatCursor(encoded string, round int64, direction string) (chatCursor, error) {
	cursor, err := decodeOpaqueCursor[chatCursor](encoded, errInvalidCursor)
	if err != nil {
		return chatCursor{}, err
	}
	if !validChatCursor(cursor) || cursor.Round != round || cursor.Direction != direction {
		return chatCursor{}, fmt.Errorf("%w: invalid chat cursor scope or fields", errInvalidCursor)
	}
	return cursor, nil
}

func compareChatPosition(a, b cgstore.ChatPosition) int {
	if c := a.OccurredAt.Compare(b.OccurredAt); c != 0 {
		return c
	}
	if a.EventLogID < b.EventLogID {
		return -1
	}
	if a.EventLogID > b.EventLogID {
		return 1
	}
	return 0
}
