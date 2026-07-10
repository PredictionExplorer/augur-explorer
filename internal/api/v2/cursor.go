package v2

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

const (
	bidCursorVersion = 1
	maxCursorLength  = 512
)

var errInvalidCursor = errors.New("invalid cursor")

type bidCursor struct {
	Version     int   `json:"v"`
	Round       int64 `json:"r"`
	BidPosition int64 `json:"p"`
	EventLogID  int64 `json:"e"`
}

func encodeBidCursor(cursor bidCursor) (string, error) {
	if cursor.Version != bidCursorVersion || cursor.Round < 0 || cursor.BidPosition < 1 || cursor.EventLogID < 1 {
		return "", fmt.Errorf("%w: invalid fields", errInvalidCursor)
	}
	payload, err := json.Marshal(cursor)
	if err != nil {
		return "", fmt.Errorf("marshal cursor: %w", err)
	}
	encoded := base64.RawURLEncoding.EncodeToString(payload)
	if len(encoded) > maxCursorLength {
		return "", fmt.Errorf("%w: encoded value is too long", errInvalidCursor)
	}
	return encoded, nil
}

func decodeBidCursor(encoded string, expectedRound int64) (bidCursor, error) {
	if encoded == "" || len(encoded) > maxCursorLength {
		return bidCursor{}, fmt.Errorf("%w: invalid length", errInvalidCursor)
	}
	payload, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return bidCursor{}, fmt.Errorf("%w: invalid base64url", errInvalidCursor)
	}

	dec := json.NewDecoder(bytes.NewReader(payload))
	dec.DisallowUnknownFields()
	var cursor bidCursor
	if err := dec.Decode(&cursor); err != nil {
		return bidCursor{}, fmt.Errorf("%w: invalid payload", errInvalidCursor)
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return bidCursor{}, fmt.Errorf("%w: trailing payload", errInvalidCursor)
	}
	if cursor.Version != bidCursorVersion ||
		cursor.Round != expectedRound ||
		cursor.Round < 0 ||
		cursor.BidPosition < 1 ||
		cursor.EventLogID < 1 {
		return bidCursor{}, fmt.Errorf("%w: invalid fields", errInvalidCursor)
	}
	return cursor, nil
}
