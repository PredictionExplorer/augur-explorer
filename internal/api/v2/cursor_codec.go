package v2

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

const maxCursorLength = 512

func encodeOpaqueCursor[T any](cursor T, invalid error, name string) (string, error) {
	payload, err := json.Marshal(cursor)
	if err != nil {
		return "", fmt.Errorf("marshal %s: %w", name, err)
	}
	encoded := base64.RawURLEncoding.EncodeToString(payload)
	if len(encoded) > maxCursorLength {
		return "", fmt.Errorf("%w: encoded value is too long", invalid)
	}
	return encoded, nil
}

func decodeOpaqueCursor[T any](encoded string, invalid error) (T, error) {
	var cursor T
	if encoded == "" || len(encoded) > maxCursorLength {
		return cursor, fmt.Errorf("%w: invalid length", invalid)
	}
	payload, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return cursor, fmt.Errorf("%w: invalid base64url", invalid)
	}

	dec := json.NewDecoder(bytes.NewReader(payload))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&cursor); err != nil {
		return cursor, fmt.Errorf("%w: invalid payload", invalid)
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return cursor, fmt.Errorf("%w: trailing payload", invalid)
	}
	return cursor, nil
}
