package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestOpaqueCursorCodec(t *testing.T) {
	t.Parallel()

	type payload struct {
		Version int   `json:"v"`
		Value   int64 `json:"x"`
	}
	invalid := errors.New("test cursor")
	encoded, err := encodeOpaqueCursor(payload{Version: 1, Value: 42}, invalid, "test cursor")
	if err != nil {
		t.Fatalf("encodeOpaqueCursor: %v", err)
	}
	got, err := decodeOpaqueCursor[payload](encoded, invalid)
	if err != nil {
		t.Fatalf("decodeOpaqueCursor: %v", err)
	}
	if got != (payload{Version: 1, Value: 42}) {
		t.Fatalf("decoded payload = %+v", got)
	}

	tests := map[string]string{
		"empty":          "",
		"oversized":      strings.Repeat("a", maxCursorLength+1),
		"invalid base64": "!",
		"unknown field":  base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"x":42,"extra":true}`)),
		"trailing JSON":  base64.RawURLEncoding.EncodeToString([]byte(`{"v":1,"x":42}{}`)),
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := decodeOpaqueCursor[payload](input, invalid); !errors.Is(err, invalid) {
				t.Fatalf("decode error = %v, want wrapped sentinel", err)
			}
		})
	}
}

func TestOpaqueCursorCodecReportsMarshalFailure(t *testing.T) {
	t.Parallel()

	invalid := errors.New("test cursor")
	_, err := encodeOpaqueCursor(struct {
		Value chan int `json:"value"`
	}{Value: make(chan int)}, invalid, "test cursor")
	if err == nil || errors.Is(err, invalid) {
		t.Fatalf("marshal error = %v, want non-validation failure", err)
	}
}
