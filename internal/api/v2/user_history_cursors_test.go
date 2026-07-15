package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestUserEventCursorRoundTrip(t *testing.T) {
	t.Parallel()
	for _, resource := range []userEventResource{
		userEventResourceRaffleEthDeposits,
		userEventResourceRaffleNftWins,
		userEventResourceEthDonations,
		userEventResourceErc20Donations,
		userEventResourceNftDonations,
		userEventResourceDonatedNfts,
	} {
		want := userEventCursor{
			Version:    userEventCursorVersion,
			Address:    userCursorAlice,
			Resource:   resource,
			EventLogID: 5040,
		}
		encoded, err := encodeUserEventCursor(want)
		if err != nil {
			t.Fatalf("encodeUserEventCursor(%s): %v", resource, err)
		}
		checksummed := strings.ToUpper(userCursorAlice[:2]) + userCursorAlice[2:]
		got, err := decodeUserEventCursor(encoded, checksummed, resource)
		if err != nil || got != want {
			t.Fatalf("round trip(%s) = %+v, %v; want %+v", resource, got, err, want)
		}
	}
}

func TestUserEventCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	encoded := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	deposits := string(userEventResourceRaffleEthDeposits)
	cases := []string{
		"",
		"%%%",
		strings.Repeat("a", maxCursorLength+1),
		encoded(`{"v":2,"a":"` + userCursorAlice + `","k":"` + deposits + `","e":1}`),
		encoded(`{"v":1,"a":"bad","k":"` + deposits + `","e":1}`),
		encoded(`{"v":1,"a":"` + strings.ToUpper(userCursorAlice) + `","k":"` + deposits + `","e":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","k":"unknown","e":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","k":"` + deposits + `","e":0}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","k":"` + deposits + `","e":1,"x":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","k":"` + deposits + `","e":1}{}`),
	}
	for _, value := range cases {
		if _, err := decodeUserEventCursor(
			value, userCursorAlice, userEventResourceRaffleEthDeposits,
		); !errors.Is(err, errInvalidUserEventCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}

	valid, err := encodeUserEventCursor(userEventCursor{
		Version:    userEventCursorVersion,
		Address:    userCursorAlice,
		Resource:   userEventResourceRaffleEthDeposits,
		EventLogID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeUserEventCursor(
		valid, userCursorBob, userEventResourceRaffleEthDeposits,
	); !errors.Is(err, errInvalidUserEventCursor) {
		t.Errorf("cross-user error = %v", err)
	}
	if _, err := decodeUserEventCursor(
		valid, userCursorAlice, userEventResourceNftDonations,
	); !errors.Is(err, errInvalidUserEventCursor) {
		t.Errorf("cross-resource error = %v", err)
	}
}

func TestEncodeUserEventCursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()
	for _, cursor := range []userEventCursor{
		{Version: 2, Address: userCursorAlice, Resource: userEventResourceRaffleNftWins, EventLogID: 1},
		{Version: 1, Address: userCursorAlice, Resource: userEventResourceRaffleNftWins, EventLogID: 0},
		{Version: 1, Address: strings.ToUpper(userCursorAlice), Resource: userEventResourceRaffleNftWins, EventLogID: 1},
		{Version: 1, Address: "bad", Resource: userEventResourceRaffleNftWins, EventLogID: 1},
		{Version: 1, Address: userCursorAlice, Resource: "unknown", EventLogID: 1},
	} {
		if _, err := encodeUserEventCursor(cursor); !errors.Is(err, errInvalidUserEventCursor) {
			t.Errorf("encode %+v error = %v", cursor, err)
		}
	}
}

func TestUserPrizeCursorRoundTrip(t *testing.T) {
	t.Parallel()
	want := userPrizeCursor{
		Version:     userPrizeCursorVersion,
		Address:     userCursorAlice,
		Round:       2,
		PrizeType:   9,
		WinnerIndex: 3,
	}
	encoded, err := encodeUserPrizeCursor(want)
	if err != nil {
		t.Fatalf("encodeUserPrizeCursor: %v", err)
	}
	checksummed := strings.ToUpper(userCursorAlice[:2]) + userCursorAlice[2:]
	got, err := decodeUserPrizeCursor(encoded, checksummed)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}
}

func TestEncodeUserPrizeCursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()
	for _, cursor := range []userPrizeCursor{
		{Version: 2, Address: userCursorAlice, Round: 0, PrizeType: 0, WinnerIndex: 0},
		{Version: 1, Address: "bad", Round: 0, PrizeType: 0, WinnerIndex: 0},
		{Version: 1, Address: userCursorAlice, Round: -1, PrizeType: 0, WinnerIndex: 0},
		{Version: 1, Address: userCursorAlice, Round: 0, PrizeType: 16, WinnerIndex: 0},
		{Version: 1, Address: userCursorAlice, Round: 0, PrizeType: 0, WinnerIndex: -1},
	} {
		if _, err := encodeUserPrizeCursor(cursor); !errors.Is(err, errInvalidUserPrizeCursor) {
			t.Errorf("encode %+v error = %v", cursor, err)
		}
	}
}

func TestEncodeUserDonatedErc20CursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()
	for _, cursor := range []userDonatedErc20Cursor{
		{Version: 2, Address: userCursorAlice, Round: 0, TokenID: 1},
		{Version: 1, Address: "bad", Round: 0, TokenID: 1},
		{Version: 1, Address: userCursorAlice, Round: -1, TokenID: 1},
		{Version: 1, Address: userCursorAlice, Round: 0, TokenID: 0},
	} {
		if _, err := encodeUserDonatedErc20Cursor(cursor); !errors.Is(err, errInvalidUserDonatedErc20Cursor) {
			t.Errorf("encode %+v error = %v", cursor, err)
		}
	}
}

func TestUserPrizeCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	encoded := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cases := []string{
		"",
		"%%%",
		strings.Repeat("a", maxCursorLength+1),
		encoded(`{"v":2,"a":"` + userCursorAlice + `","r":0,"t":0,"w":0}`),
		encoded(`{"v":1,"a":"bad","r":0,"t":0,"w":0}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","r":-1,"t":0,"w":0}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","r":0,"t":-1,"w":0}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","r":0,"t":16,"w":0}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","r":0,"t":0,"w":-1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","r":0,"t":0,"w":0,"x":1}`),
	}
	for _, value := range cases {
		if _, err := decodeUserPrizeCursor(value, userCursorAlice); !errors.Is(err, errInvalidUserPrizeCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}

	crossUser, err := encodeUserPrizeCursor(userPrizeCursor{
		Version: userPrizeCursorVersion, Address: userCursorAlice,
		Round: 0, PrizeType: 0, WinnerIndex: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeUserPrizeCursor(crossUser, userCursorBob); !errors.Is(err, errInvalidUserPrizeCursor) {
		t.Errorf("cross-user error = %v", err)
	}
}

func TestUserDonatedErc20CursorRoundTripAndRejection(t *testing.T) {
	t.Parallel()
	want := userDonatedErc20Cursor{
		Version: userDonatedErc20CursorVersion,
		Address: userCursorAlice,
		Round:   1,
		TokenID: 26,
	}
	encoded, err := encodeUserDonatedErc20Cursor(want)
	if err != nil {
		t.Fatalf("encodeUserDonatedErc20Cursor: %v", err)
	}
	got, err := decodeUserDonatedErc20Cursor(encoded, userCursorAlice)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}

	base64URL := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cases := []string{
		"",
		"%%%",
		base64URL(`{"v":2,"a":"` + userCursorAlice + `","r":0,"t":1}`),
		base64URL(`{"v":1,"a":"` + userCursorAlice + `","r":-1,"t":1}`),
		base64URL(`{"v":1,"a":"` + userCursorAlice + `","r":0,"t":0}`),
		base64URL(`{"v":1,"a":"bad","r":0,"t":1}`),
	}
	for _, value := range cases {
		if _, err := decodeUserDonatedErc20Cursor(value, userCursorAlice); !errors.Is(err, errInvalidUserDonatedErc20Cursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	if _, err := decodeUserDonatedErc20Cursor(encoded, userCursorBob); !errors.Is(err, errInvalidUserDonatedErc20Cursor) {
		t.Errorf("cross-user error = %v", err)
	}
}

func FuzzDecodeUserEventCursor(f *testing.F) {
	valid, _ := encodeUserEventCursor(userEventCursor{
		Version:    userEventCursorVersion,
		Address:    userCursorAlice,
		Resource:   userEventResourceRaffleEthDeposits,
		EventLogID: 5040,
	})
	f.Add(valid, userCursorAlice, string(userEventResourceRaffleEthDeposits))
	f.Add("", "", "")
	f.Add("%%%", userCursorBob, string(userEventResourceDonatedNfts))
	f.Fuzz(func(t *testing.T, value, address, resource string) {
		cursor, err := decodeUserEventCursor(value, address, userEventResource(resource))
		if err == nil && !validUserEventCursor(cursor, address, userEventResource(resource)) {
			t.Fatalf("accepted invalid cursor: %+v for %q/%q", cursor, address, resource)
		}
	})
}

func FuzzDecodeUserPrizeCursor(f *testing.F) {
	valid, _ := encodeUserPrizeCursor(userPrizeCursor{
		Version:     userPrizeCursorVersion,
		Address:     userCursorAlice,
		Round:       2,
		PrizeType:   9,
		WinnerIndex: 3,
	})
	f.Add(valid, userCursorAlice)
	f.Add("", "")
	f.Add("%%%", userCursorBob)
	f.Fuzz(func(t *testing.T, value, address string) {
		cursor, err := decodeUserPrizeCursor(value, address)
		if err == nil && !validUserPrizeCursor(cursor, address) {
			t.Fatalf("accepted invalid cursor: %+v for %q", cursor, address)
		}
	})
}

func FuzzDecodeUserDonatedErc20Cursor(f *testing.F) {
	valid, _ := encodeUserDonatedErc20Cursor(userDonatedErc20Cursor{
		Version: userDonatedErc20CursorVersion,
		Address: userCursorAlice,
		Round:   1,
		TokenID: 26,
	})
	f.Add(valid, userCursorAlice)
	f.Add("", "")
	f.Add("%%%", userCursorBob)
	f.Fuzz(func(t *testing.T, value, address string) {
		cursor, err := decodeUserDonatedErc20Cursor(value, address)
		if err == nil && !validUserDonatedErc20Cursor(cursor, address) {
			t.Fatalf("accepted invalid cursor: %+v for %q", cursor, address)
		}
	})
}
