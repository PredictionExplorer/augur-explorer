package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

const (
	userCursorAlice = "0x2100000000000000000000000000000000000021"
	userCursorBob   = "0x2200000000000000000000000000000000000022"
)

func TestUserBidCursorRoundTrip(t *testing.T) {
	t.Parallel()
	want := userBidCursor{
		Version:    userBidCursorVersion,
		Address:    userCursorAlice,
		EventLogID: 5071,
	}
	encoded, err := encodeUserBidCursor(want)
	if err != nil {
		t.Fatalf("encodeUserBidCursor: %v", err)
	}
	got, err := decodeUserBidCursor(encoded, strings.ToUpper(userCursorAlice[:2])+userCursorAlice[2:])
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}
}

func TestUserBidCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	encoded := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cases := []string{
		"",
		"%%%",
		strings.Repeat("a", maxCursorLength+1),
		encoded(`{"v":2,"a":"` + userCursorAlice + `","e":1}`),
		encoded(`{"v":1,"a":"bad","e":1}`),
		encoded(`{"v":1,"a":"` + strings.ToUpper(userCursorAlice) + `","e":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","e":0}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","e":1,"x":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","e":1}{}`),
	}
	for _, value := range cases {
		if _, err := decodeUserBidCursor(value, userCursorAlice); !errors.Is(err, errInvalidUserBidCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}

	crossUser, err := encodeUserBidCursor(userBidCursor{
		Version: userBidCursorVersion, Address: userCursorAlice, EventLogID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeUserBidCursor(crossUser, userCursorBob); !errors.Is(err, errInvalidUserBidCursor) {
		t.Errorf("cross-user error = %v", err)
	}
}

func TestEncodeUserBidCursorRejectsNonCanonicalScope(t *testing.T) {
	t.Parallel()
	for _, cursor := range []userBidCursor{
		{Version: 2, Address: userCursorAlice, EventLogID: 1},
		{Version: 1, Address: userCursorAlice, EventLogID: 0},
		{Version: 1, Address: strings.ToUpper(userCursorAlice), EventLogID: 1},
		{Version: 1, Address: "bad", EventLogID: 1},
	} {
		if _, err := encodeUserBidCursor(cursor); !errors.Is(err, errInvalidUserBidCursor) {
			t.Errorf("encode %+v error = %v", cursor, err)
		}
	}
}

func FuzzDecodeUserBidCursor(f *testing.F) {
	valid, _ := encodeUserBidCursor(userBidCursor{
		Version: userBidCursorVersion, Address: userCursorAlice, EventLogID: 5071,
	})
	f.Add(valid, userCursorAlice)
	f.Add("", "")
	f.Add("%%%", userCursorBob)
	f.Fuzz(func(t *testing.T, value, address string) {
		cursor, err := decodeUserBidCursor(value, address)
		if err == nil && !validUserBidCursor(cursor, address) {
			t.Fatalf("accepted invalid cursor: %+v for %q", cursor, address)
		}
	})
}
