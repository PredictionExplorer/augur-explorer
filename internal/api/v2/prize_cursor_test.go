package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestPrizeCursorRoundTrip(t *testing.T) {
	t.Parallel()

	want := prizeCursor{
		Version:     prizeCursorVersion,
		Round:       42,
		PrizeType:   10,
		WinnerIndex: 3,
	}
	encoded, err := encodePrizeCursor(want)
	if err != nil {
		t.Fatalf("encodePrizeCursor: %v", err)
	}
	encodedAgain, err := encodePrizeCursor(want)
	if err != nil {
		t.Fatalf("encodePrizeCursor again: %v", err)
	}
	if encoded != encodedAgain {
		t.Fatalf("cursor encoding is not deterministic: %q != %q", encoded, encodedAgain)
	}
	got, err := decodePrizeCursor(encoded, want.Round)
	if err != nil {
		t.Fatalf("decodePrizeCursor: %v", err)
	}
	if got != want {
		t.Fatalf("decoded cursor = %+v, want %+v", got, want)
	}
}

func TestDecodePrizeCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	valid, err := encodePrizeCursor(prizeCursor{
		Version:     prizeCursorVersion,
		Round:       4,
		PrizeType:   10,
		WinnerIndex: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	encodedJSON := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	tests := map[string]struct {
		value string
		round int64
	}{
		"empty":            {"", 4},
		"oversized":        {strings.Repeat("a", maxCursorLength+1), 4},
		"invalid base64":   {"%%%", 4},
		"invalid json":     {encodedJSON("{"), 4},
		"unknown field":    {encodedJSON(`{"v":1,"r":4,"t":10,"w":2,"x":1}`), 4},
		"trailing payload": {encodedJSON(`{"v":1,"r":4,"t":10,"w":2}{}`), 4},
		"wrong version":    {encodedJSON(`{"v":2,"r":4,"t":10,"w":2}`), 4},
		"wrong round":      {valid, 5},
		"negative round":   {encodedJSON(`{"v":1,"r":-1,"t":10,"w":2}`), -1},
		"negative type":    {encodedJSON(`{"v":1,"r":4,"t":-1,"w":2}`), 4},
		"unknown type":     {encodedJSON(`{"v":1,"r":4,"t":16,"w":2}`), 4},
		"negative winner":  {encodedJSON(`{"v":1,"r":4,"t":10,"w":-1}`), 4},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := decodePrizeCursor(tc.value, tc.round); !errors.Is(err, errInvalidPrizeCursor) {
				t.Fatalf("decodePrizeCursor() error = %v, want errInvalidPrizeCursor", err)
			}
		})
	}
}

func TestEncodePrizeCursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()

	tests := []prizeCursor{
		{},
		{Version: 2, Round: 1, PrizeType: 0, WinnerIndex: 0},
		{Version: 1, Round: -1, PrizeType: 0, WinnerIndex: 0},
		{Version: 1, Round: 1, PrizeType: -1, WinnerIndex: 0},
		{Version: 1, Round: 1, PrizeType: 16, WinnerIndex: 0},
		{Version: 1, Round: 1, PrizeType: 0, WinnerIndex: -1},
	}
	for _, tc := range tests {
		if _, err := encodePrizeCursor(tc); !errors.Is(err, errInvalidPrizeCursor) {
			t.Errorf("encodePrizeCursor(%+v) error = %v, want errInvalidPrizeCursor", tc, err)
		}
	}
}

func TestPrizeCursorFollows(t *testing.T) {
	t.Parallel()

	previous := cgstore.PrizePageCursor{PrizeType: 5, WinnerIndex: 2}
	tests := []struct {
		current cgstore.PrizePageCursor
		want    bool
	}{
		{cgstore.PrizePageCursor{PrizeType: 6, WinnerIndex: 0}, true},
		{cgstore.PrizePageCursor{PrizeType: 5, WinnerIndex: 3}, true},
		{cgstore.PrizePageCursor{PrizeType: 5, WinnerIndex: 2}, false},
		{cgstore.PrizePageCursor{PrizeType: 5, WinnerIndex: 1}, false},
		{cgstore.PrizePageCursor{PrizeType: 4, WinnerIndex: 99}, false},
	}
	for _, tc := range tests {
		if got := prizeCursorFollows(tc.current, previous); got != tc.want {
			t.Errorf("prizeCursorFollows(%+v,%+v) = %v, want %v",
				tc.current, previous, got, tc.want)
		}
	}
}

func FuzzDecodePrizeCursor(f *testing.F) {
	valid, err := encodePrizeCursor(prizeCursor{
		Version:     prizeCursorVersion,
		Round:       7,
		PrizeType:   10,
		WinnerIndex: 3,
	})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(valid, int64(7))
	f.Add("", int64(0))
	f.Add("%%%", int64(-1))
	f.Add(strings.Repeat("a", maxCursorLength+1), int64(7))

	f.Fuzz(func(t *testing.T, encoded string, round int64) {
		cursor, err := decodePrizeCursor(encoded, round)
		if err != nil {
			return
		}
		if !validPrizeCursor(cursor, round) {
			t.Fatalf("accepted invalid cursor: %+v for round %d", cursor, round)
		}
	})
}
