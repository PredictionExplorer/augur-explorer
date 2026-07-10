package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestRoundCursorRoundTrip(t *testing.T) {
	t.Parallel()

	want := roundCursor{
		Version:    roundCursorVersion,
		RoundNum:   42,
		EventLogID: 9001,
	}
	encoded, err := encodeRoundCursor(want)
	if err != nil {
		t.Fatalf("encodeRoundCursor: %v", err)
	}
	encodedAgain, err := encodeRoundCursor(want)
	if err != nil {
		t.Fatalf("encodeRoundCursor again: %v", err)
	}
	if encoded != encodedAgain {
		t.Fatalf("cursor encoding is not deterministic: %q != %q", encoded, encodedAgain)
	}

	got, err := decodeRoundCursor(encoded)
	if err != nil {
		t.Fatalf("decodeRoundCursor: %v", err)
	}
	if got != want {
		t.Fatalf("decoded cursor = %+v, want %+v", got, want)
	}
}

func TestDecodeRoundCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	encodedJSON := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	tests := map[string]string{
		"empty":            "",
		"oversized":        strings.Repeat("a", maxCursorLength+1),
		"invalid base64":   "%%%",
		"invalid json":     encodedJSON("{"),
		"unknown field":    encodedJSON(`{"v":1,"r":4,"e":10,"x":1}`),
		"trailing payload": encodedJSON(`{"v":1,"r":4,"e":10}{}`),
		"wrong version":    encodedJSON(`{"v":2,"r":4,"e":10}`),
		"negative round":   encodedJSON(`{"v":1,"r":-1,"e":10}`),
		"zero event log":   encodedJSON(`{"v":1,"r":4,"e":0}`),
	}
	for name, value := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := decodeRoundCursor(value); !errors.Is(err, errInvalidRoundCursor) {
				t.Fatalf("decodeRoundCursor() error = %v, want errInvalidRoundCursor", err)
			}
		})
	}
}

func TestEncodeRoundCursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()

	tests := []roundCursor{
		{},
		{Version: 2, RoundNum: 1, EventLogID: 1},
		{Version: 1, RoundNum: -1, EventLogID: 1},
		{Version: 1, RoundNum: 1, EventLogID: 0},
	}
	for _, tc := range tests {
		if _, err := encodeRoundCursor(tc); !errors.Is(err, errInvalidRoundCursor) {
			t.Errorf("encodeRoundCursor(%+v) error = %v, want errInvalidRoundCursor", tc, err)
		}
	}
}

func TestRoundCursorPrecedes(t *testing.T) {
	t.Parallel()

	previous := cgstore.RoundPageCursor{RoundNum: 5, EventLogID: 20}
	tests := []struct {
		current cgstore.RoundPageCursor
		want    bool
	}{
		{cgstore.RoundPageCursor{RoundNum: 4, EventLogID: 99}, true},
		{cgstore.RoundPageCursor{RoundNum: 5, EventLogID: 19}, true},
		{cgstore.RoundPageCursor{RoundNum: 5, EventLogID: 20}, false},
		{cgstore.RoundPageCursor{RoundNum: 5, EventLogID: 21}, false},
		{cgstore.RoundPageCursor{RoundNum: 6, EventLogID: 1}, false},
	}
	for _, tc := range tests {
		if got := roundCursorPrecedes(tc.current, previous); got != tc.want {
			t.Errorf("roundCursorPrecedes(%+v,%+v) = %v, want %v",
				tc.current, previous, got, tc.want)
		}
	}
}

func FuzzDecodeRoundCursor(f *testing.F) {
	valid, err := encodeRoundCursor(roundCursor{
		Version:    roundCursorVersion,
		RoundNum:   7,
		EventLogID: 99,
	})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(valid)
	f.Add("")
	f.Add("%%%")
	f.Add(strings.Repeat("a", maxCursorLength+1))

	f.Fuzz(func(t *testing.T, encoded string) {
		cursor, err := decodeRoundCursor(encoded)
		if err != nil {
			return
		}
		if cursor.Version != roundCursorVersion || cursor.RoundNum < 0 || cursor.EventLogID < 1 {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}
