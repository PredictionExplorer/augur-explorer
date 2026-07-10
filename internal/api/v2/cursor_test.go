package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestBidCursorRoundTrip(t *testing.T) {
	t.Parallel()

	want := bidCursor{
		Version:     bidCursorVersion,
		Round:       42,
		BidPosition: 17,
		EventLogID:  9001,
	}
	encoded, err := encodeBidCursor(want)
	if err != nil {
		t.Fatalf("encodeBidCursor: %v", err)
	}
	encodedAgain, err := encodeBidCursor(want)
	if err != nil {
		t.Fatalf("encodeBidCursor again: %v", err)
	}
	if encoded != encodedAgain {
		t.Fatalf("cursor encoding is not deterministic: %q != %q", encoded, encodedAgain)
	}

	got, err := decodeBidCursor(encoded, want.Round)
	if err != nil {
		t.Fatalf("decodeBidCursor: %v", err)
	}
	if got != want {
		t.Fatalf("decoded cursor = %+v, want %+v", got, want)
	}
}

func TestDecodeBidCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	valid, err := encodeBidCursor(bidCursor{
		Version:     bidCursorVersion,
		Round:       4,
		BidPosition: 2,
		EventLogID:  10,
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
		"unknown field":    {encodedJSON(`{"v":1,"r":4,"p":2,"e":10,"x":1}`), 4},
		"trailing payload": {encodedJSON(`{"v":1,"r":4,"p":2,"e":10}{}`), 4},
		"wrong version":    {encodedJSON(`{"v":2,"r":4,"p":2,"e":10}`), 4},
		"wrong round":      {valid, 5},
		"negative round":   {encodedJSON(`{"v":1,"r":-1,"p":2,"e":10}`), -1},
		"zero position":    {encodedJSON(`{"v":1,"r":4,"p":0,"e":10}`), 4},
		"zero event log":   {encodedJSON(`{"v":1,"r":4,"p":2,"e":0}`), 4},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := decodeBidCursor(tc.value, tc.round); !errors.Is(err, errInvalidCursor) {
				t.Fatalf("decodeBidCursor() error = %v, want errInvalidCursor", err)
			}
		})
	}
}

func TestEncodeBidCursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()

	tests := []bidCursor{
		{},
		{Version: 2, Round: 1, BidPosition: 1, EventLogID: 1},
		{Version: 1, Round: -1, BidPosition: 1, EventLogID: 1},
		{Version: 1, Round: 1, BidPosition: 0, EventLogID: 1},
		{Version: 1, Round: 1, BidPosition: 1, EventLogID: 0},
	}
	for _, tc := range tests {
		if _, err := encodeBidCursor(tc); !errors.Is(err, errInvalidCursor) {
			t.Errorf("encodeBidCursor(%+v) error = %v, want errInvalidCursor", tc, err)
		}
	}
}

func FuzzDecodeBidCursor(f *testing.F) {
	valid, err := encodeBidCursor(bidCursor{
		Version:     bidCursorVersion,
		Round:       7,
		BidPosition: 3,
		EventLogID:  99,
	})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(valid, int64(7))
	f.Add("", int64(0))
	f.Add("%%%", int64(-1))
	f.Add(strings.Repeat("a", maxCursorLength+1), int64(7))

	f.Fuzz(func(t *testing.T, encoded string, round int64) {
		cursor, err := decodeBidCursor(encoded, round)
		if err != nil {
			return
		}
		if cursor.Version != bidCursorVersion ||
			cursor.Round != round ||
			cursor.Round < 0 ||
			cursor.BidPosition < 1 ||
			cursor.EventLogID < 1 {
			t.Fatalf("accepted invalid cursor: %+v for round %d", cursor, round)
		}
	})
}
