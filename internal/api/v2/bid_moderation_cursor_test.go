package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestBidBanCursorRoundTrip(t *testing.T) {
	t.Parallel()
	want := bidBanCursor{Version: bidBanCursorVersion, ID: 42}
	encoded, err := encodeBidBanCursor(want)
	if err != nil {
		t.Fatalf("encodeBidBanCursor: %v", err)
	}
	got, err := decodeBidBanCursor(encoded)
	if err != nil {
		t.Fatalf("decodeBidBanCursor: %v", err)
	}
	if got != want {
		t.Fatalf("cursor = %+v, want %+v", got, want)
	}
}

func TestBidBanCursorRejectsInvalidValues(t *testing.T) {
	t.Parallel()
	if _, err := encodeBidBanCursor(bidBanCursor{}); !errors.Is(err, errInvalidBidBanCursor) {
		t.Fatalf("encode zero cursor = %v", err)
	}
	if _, err := encodeBidBanCursor(bidBanCursor{Version: 2, ID: 1}); !errors.Is(err, errInvalidBidBanCursor) {
		t.Fatalf("encode unsupported cursor = %v", err)
	}

	encodeRaw := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	for name, encoded := range map[string]string{
		"empty":       "",
		"bad base64":  "!",
		"zero id":     encodeRaw(`{"v":1,"i":0}`),
		"bad version": encodeRaw(`{"v":2,"i":1}`),
		"unknown":     encodeRaw(`{"v":1,"i":1,"x":2}`),
		"trailing":    encodeRaw(`{"v":1,"i":1}{}`),
		"too long":    strings.Repeat("a", maxCursorLength+1),
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := decodeBidBanCursor(encoded); !errors.Is(err, errInvalidBidBanCursor) {
				t.Fatalf("decodeBidBanCursor(%q) = %v", name, err)
			}
		})
	}
}

func FuzzDecodeBidBanCursor(f *testing.F) {
	valid, err := encodeBidBanCursor(bidBanCursor{Version: bidBanCursorVersion, ID: 42})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(valid)
	f.Add("")
	f.Add("not-base64")
	f.Fuzz(func(t *testing.T, encoded string) {
		cursor, err := decodeBidBanCursor(encoded)
		if err == nil && (cursor.Version != bidBanCursorVersion || cursor.ID < 1) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}
