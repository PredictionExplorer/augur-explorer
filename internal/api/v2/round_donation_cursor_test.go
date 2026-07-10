package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestRoundDonationCursorRoundTrip(t *testing.T) {
	t.Parallel()

	for _, resource := range []roundDonationResource{
		roundDonationResourceETH,
		roundDonationResourceERC20,
		roundDonationResourceNFT,
	} {
		t.Run(string(resource), func(t *testing.T) {
			t.Parallel()
			want := roundDonationCursor{
				Version:    roundDonationCursorVersion,
				Round:      4,
				Resource:   resource,
				EventLogID: 100,
			}
			encoded, err := encodeRoundDonationCursor(want)
			if err != nil {
				t.Fatal(err)
			}
			encodedAgain, err := encodeRoundDonationCursor(want)
			if err != nil {
				t.Fatal(err)
			}
			if encoded != encodedAgain {
				t.Fatalf("nondeterministic encoding: %q != %q", encoded, encodedAgain)
			}
			got, err := decodeRoundDonationCursor(encoded, want.Round, resource)
			if err != nil {
				t.Fatal(err)
			}
			if got != want {
				t.Fatalf("decoded cursor = %+v, want %+v", got, want)
			}
		})
	}
}

func TestDecodeRoundDonationCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	encodedJSON := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	tests := map[string]struct {
		value    string
		round    int64
		resource roundDonationResource
	}{
		"empty":            {"", 4, roundDonationResourceETH},
		"oversized":        {strings.Repeat("a", maxCursorLength+1), 4, roundDonationResourceETH},
		"invalid base64":   {"%%%", 4, roundDonationResourceETH},
		"unknown field":    {encodedJSON(`{"v":1,"r":4,"k":"eth","e":10,"x":1}`), 4, roundDonationResourceETH},
		"trailing payload": {encodedJSON(`{"v":1,"r":4,"k":"eth","e":10}{}`), 4, roundDonationResourceETH},
		"wrong version":    {encodedJSON(`{"v":2,"r":4,"k":"eth","e":10}`), 4, roundDonationResourceETH},
		"wrong round":      {encodedJSON(`{"v":1,"r":5,"k":"eth","e":10}`), 4, roundDonationResourceETH},
		"wrong resource":   {encodedJSON(`{"v":1,"r":4,"k":"nft","e":10}`), 4, roundDonationResourceETH},
		"unknown resource": {encodedJSON(`{"v":1,"r":4,"k":"other","e":10}`), 4, roundDonationResource("other")},
		"negative round":   {encodedJSON(`{"v":1,"r":-1,"k":"eth","e":10}`), -1, roundDonationResourceETH},
		"zero event":       {encodedJSON(`{"v":1,"r":4,"k":"eth","e":0}`), 4, roundDonationResourceETH},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := decodeRoundDonationCursor(tc.value, tc.round, tc.resource); !errors.Is(err, errInvalidRoundDonationCursor) {
				t.Fatalf("error = %v, want round donation cursor error", err)
			}
		})
	}
}

func TestEncodeRoundDonationCursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()

	for _, cursor := range []roundDonationCursor{
		{},
		{Version: 2, Round: 1, Resource: roundDonationResourceETH, EventLogID: 1},
		{Version: 1, Round: -1, Resource: roundDonationResourceETH, EventLogID: 1},
		{Version: 1, Round: 1, Resource: "other", EventLogID: 1},
		{Version: 1, Round: 1, Resource: roundDonationResourceETH, EventLogID: 0},
	} {
		if _, err := encodeRoundDonationCursor(cursor); !errors.Is(err, errInvalidRoundDonationCursor) {
			t.Errorf("encode cursor %+v: %v", cursor, err)
		}
	}
}

func FuzzDecodeRoundEthDonationCursor(f *testing.F) {
	fuzzRoundDonationCursor(f, roundDonationResourceETH)
}

func FuzzDecodeRoundERC20DonationCursor(f *testing.F) {
	fuzzRoundDonationCursor(f, roundDonationResourceERC20)
}

func FuzzDecodeRoundNFTDonationCursor(f *testing.F) {
	fuzzRoundDonationCursor(f, roundDonationResourceNFT)
}

func fuzzRoundDonationCursor(f *testing.F, resource roundDonationResource) {
	valid, err := encodeRoundDonationCursor(roundDonationCursor{
		Version:    roundDonationCursorVersion,
		Round:      7,
		Resource:   resource,
		EventLogID: 99,
	})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(valid, int64(7))
	f.Add("", int64(0))
	f.Add("%%%", int64(-1))

	f.Fuzz(func(t *testing.T, encoded string, round int64) {
		cursor, err := decodeRoundDonationCursor(encoded, round, resource)
		if err != nil {
			return
		}
		if !validRoundDonationCursor(cursor, round, resource) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}
