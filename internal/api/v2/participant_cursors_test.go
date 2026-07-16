package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func TestParticipantCursorRoundTrip(t *testing.T) {
	t.Parallel()
	for _, kind := range []cgstore.ParticipantKind{
		cgstore.ParticipantBidders,
		cgstore.ParticipantWinners,
		cgstore.ParticipantDonors,
		cgstore.ParticipantCSTStakers,
		cgstore.ParticipantRandomWalkStakers,
		cgstore.ParticipantDualStakers,
		cgstore.ParticipantCsTokenHolders,
		cgstore.ParticipantCosmicTokenHolders,
	} {
		sortValue := "9"
		if kind == cgstore.ParticipantDonors || kind == cgstore.ParticipantCSTStakers ||
			kind == cgstore.ParticipantCosmicTokenHolders {
			sortValue = "999999999999999999999999"
		}
		want := participantCursor{
			Version: participantCursorVersion, Kind: kind,
			SortValue: sortValue, AddressID: 21,
		}
		encoded, err := encodeParticipantCursor(want)
		if err != nil {
			t.Fatalf("encode %s: %v", kind, err)
		}
		got, err := decodeParticipantCursor(encoded, kind)
		if err != nil || got != want {
			t.Fatalf("round trip %s = %+v, %v; want %+v", kind, got, err, want)
		}
	}
}

func TestParticipantCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	encoded := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cases := []string{
		"",
		"%%%",
		strings.Repeat("a", maxCursorLength+1),
		encoded(`{"v":2,"k":"bidders","s":"1","a":1}`),
		encoded(`{"v":1,"k":"other","s":"1","a":1}`),
		encoded(`{"v":1,"k":"bidders","s":"-1","a":1}`),
		encoded(`{"v":1,"k":"bidders","s":"01","a":1}`),
		encoded(`{"v":1,"k":"bidders","s":"1.0","a":1}`),
		encoded(`{"v":1,"k":"bidders","s":"9223372036854775808","a":1}`),
		encoded(`{"v":1,"k":"bidders","s":"1","a":0}`),
		encoded(`{"v":1,"k":"bidders","s":"1","a":1,"x":1}`),
		encoded(`{"v":1,"k":"bidders","s":"1","a":1}{}`),
	}
	for _, value := range cases {
		if _, err := decodeParticipantCursor(value, cgstore.ParticipantBidders); !errors.Is(err, errInvalidParticipantCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	crossKind, err := encodeParticipantCursor(participantCursor{
		Version: 1, Kind: cgstore.ParticipantBidders, SortValue: "1", AddressID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeParticipantCursor(crossKind, cgstore.ParticipantWinners); !errors.Is(err, errInvalidParticipantCursor) {
		t.Errorf("cross-kind error = %v", err)
	}
	// The holder directories are participant kinds too; their cursors must
	// stay confined to their own directory.
	holders, err := encodeParticipantCursor(participantCursor{
		Version: 1, Kind: cgstore.ParticipantCsTokenHolders, SortValue: "3", AddressID: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeParticipantCursor(holders, cgstore.ParticipantCosmicTokenHolders); !errors.Is(err, errInvalidParticipantCursor) {
		t.Errorf("cross-holder-directory error = %v", err)
	}
	// Cosmic Signature holdings are bounded counts: a beyond-int64 sort
	// value is invalid there while the balance directory accepts it.
	if _, err := decodeParticipantCursor(encoded(
		`{"v":1,"k":"cosmicSignatureHolders","s":"999999999999999999999999","a":1}`,
	), cgstore.ParticipantCsTokenHolders); !errors.Is(err, errInvalidParticipantCursor) {
		t.Errorf("oversized holder count error = %v", err)
	}
}

func FuzzDecodeParticipantCursor(f *testing.F) {
	valid, _ := encodeParticipantCursor(participantCursor{
		Version: 1, Kind: cgstore.ParticipantBidders, SortValue: "5", AddressID: 21,
	})
	f.Add(valid, string(cgstore.ParticipantBidders))
	f.Add("", "")
	f.Fuzz(func(t *testing.T, value, kind string) {
		expected := cgstore.ParticipantKind(kind)
		cursor, err := decodeParticipantCursor(value, expected)
		if err == nil && !validParticipantCursor(cursor, expected) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}
