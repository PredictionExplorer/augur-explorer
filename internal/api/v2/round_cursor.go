package v2

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

const roundCursorVersion = 1

var errInvalidRoundCursor = errors.New("invalid round cursor")

type roundCursor struct {
	Version    int   `json:"v"`
	RoundNum   int64 `json:"r"`
	EventLogID int64 `json:"e"`
}

func encodeRoundCursor(cursor roundCursor) (string, error) {
	if cursor.Version != roundCursorVersion || cursor.RoundNum < 0 || cursor.EventLogID < 1 {
		return "", fmt.Errorf("%w: invalid fields", errInvalidRoundCursor)
	}
	payload, err := json.Marshal(cursor)
	if err != nil {
		return "", fmt.Errorf("marshal round cursor: %w", err)
	}
	encoded := base64.RawURLEncoding.EncodeToString(payload)
	if len(encoded) > maxCursorLength {
		return "", fmt.Errorf("%w: encoded value is too long", errInvalidRoundCursor)
	}
	return encoded, nil
}

func decodeRoundCursor(encoded string) (roundCursor, error) {
	if encoded == "" || len(encoded) > maxCursorLength {
		return roundCursor{}, fmt.Errorf("%w: invalid length", errInvalidRoundCursor)
	}
	payload, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return roundCursor{}, fmt.Errorf("%w: invalid base64url", errInvalidRoundCursor)
	}

	dec := json.NewDecoder(bytes.NewReader(payload))
	dec.DisallowUnknownFields()
	var cursor roundCursor
	if err := dec.Decode(&cursor); err != nil {
		return roundCursor{}, fmt.Errorf("%w: invalid payload", errInvalidRoundCursor)
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return roundCursor{}, fmt.Errorf("%w: trailing payload", errInvalidRoundCursor)
	}
	if cursor.Version != roundCursorVersion || cursor.RoundNum < 0 || cursor.EventLogID < 1 {
		return roundCursor{}, fmt.Errorf("%w: invalid fields", errInvalidRoundCursor)
	}
	return cursor, nil
}
