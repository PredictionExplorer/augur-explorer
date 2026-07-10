package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestRaffleEthDepositCursorRoundTrip(t *testing.T) {
	t.Parallel()

	want := raffleEthDepositCursor{
		Version:     raffleEthDepositCursorVersion,
		Round:       4,
		WinnerIndex: 2,
		EventLogID:  100,
	}
	encoded, err := encodeRaffleEthDepositCursor(want)
	if err != nil {
		t.Fatal(err)
	}
	encodedAgain, err := encodeRaffleEthDepositCursor(want)
	if err != nil {
		t.Fatal(err)
	}
	if encoded != encodedAgain {
		t.Fatalf("nondeterministic cursor encoding: %q != %q", encoded, encodedAgain)
	}
	got, err := decodeRaffleEthDepositCursor(encoded, want.Round)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("decoded cursor = %+v, want %+v", got, want)
	}
}

func TestRaffleNftWinnerCursorRoundTrip(t *testing.T) {
	t.Parallel()

	want := raffleNftWinnerCursor{
		Version:     raffleNftWinnerCursorVersion,
		Round:       4,
		Pool:        RandomWalkStaker,
		WinnerIndex: 2,
		EventLogID:  100,
	}
	encoded, err := encodeRaffleNftWinnerCursor(want)
	if err != nil {
		t.Fatal(err)
	}
	encodedAgain, err := encodeRaffleNftWinnerCursor(want)
	if err != nil {
		t.Fatal(err)
	}
	if encoded != encodedAgain {
		t.Fatalf("nondeterministic cursor encoding: %q != %q", encoded, encodedAgain)
	}
	got, err := decodeRaffleNftWinnerCursor(encoded, want.Round, want.Pool)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("decoded cursor = %+v, want %+v", got, want)
	}
}

func TestDecodeRaffleCursorsRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	encodedJSON := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	ethTests := map[string]struct {
		value string
		round int64
	}{
		"empty":            {"", 4},
		"oversized":        {strings.Repeat("a", maxCursorLength+1), 4},
		"invalid base64":   {"%%%", 4},
		"unknown field":    {encodedJSON(`{"v":1,"r":4,"w":2,"e":10,"x":1}`), 4},
		"trailing payload": {encodedJSON(`{"v":1,"r":4,"w":2,"e":10}{}`), 4},
		"wrong version":    {encodedJSON(`{"v":2,"r":4,"w":2,"e":10}`), 4},
		"wrong round":      {encodedJSON(`{"v":1,"r":5,"w":2,"e":10}`), 4},
		"negative winner":  {encodedJSON(`{"v":1,"r":4,"w":-1,"e":10}`), 4},
		"zero event":       {encodedJSON(`{"v":1,"r":4,"w":2,"e":0}`), 4},
	}
	for name, tc := range ethTests {
		t.Run("eth/"+name, func(t *testing.T) {
			t.Parallel()
			if _, err := decodeRaffleEthDepositCursor(tc.value, tc.round); !errors.Is(err, errInvalidRaffleEthDepositCursor) {
				t.Fatalf("error = %v, want raffle ETH cursor error", err)
			}
		})
	}

	nftTests := map[string]struct {
		value string
		round int64
		pool  RaffleNftPool
	}{
		"empty":            {"", 4, Bidder},
		"oversized":        {strings.Repeat("a", maxCursorLength+1), 4, Bidder},
		"invalid base64":   {"%%%", 4, Bidder},
		"unknown field":    {encodedJSON(`{"v":1,"r":4,"p":"bidder","w":2,"e":10,"x":1}`), 4, Bidder},
		"trailing payload": {encodedJSON(`{"v":1,"r":4,"p":"bidder","w":2,"e":10}{}`), 4, Bidder},
		"wrong version":    {encodedJSON(`{"v":2,"r":4,"p":"bidder","w":2,"e":10}`), 4, Bidder},
		"wrong round":      {encodedJSON(`{"v":1,"r":5,"p":"bidder","w":2,"e":10}`), 4, Bidder},
		"wrong pool":       {encodedJSON(`{"v":1,"r":4,"p":"randomWalkStaker","w":2,"e":10}`), 4, Bidder},
		"unknown pool":     {encodedJSON(`{"v":1,"r":4,"p":"other","w":2,"e":10}`), 4, RaffleNftPool("other")},
		"negative winner":  {encodedJSON(`{"v":1,"r":4,"p":"bidder","w":-1,"e":10}`), 4, Bidder},
		"zero event":       {encodedJSON(`{"v":1,"r":4,"p":"bidder","w":2,"e":0}`), 4, Bidder},
	}
	for name, tc := range nftTests {
		t.Run("nft/"+name, func(t *testing.T) {
			t.Parallel()
			if _, err := decodeRaffleNftWinnerCursor(tc.value, tc.round, tc.pool); !errors.Is(err, errInvalidRaffleNftWinnerCursor) {
				t.Fatalf("error = %v, want raffle NFT cursor error", err)
			}
		})
	}
}

func TestEncodeRaffleCursorsRejectInvalidFields(t *testing.T) {
	t.Parallel()

	for _, cursor := range []raffleEthDepositCursor{
		{},
		{Version: 2, Round: 1, WinnerIndex: 0, EventLogID: 1},
		{Version: 1, Round: -1, WinnerIndex: 0, EventLogID: 1},
		{Version: 1, Round: 1, WinnerIndex: -1, EventLogID: 1},
		{Version: 1, Round: 1, WinnerIndex: 0, EventLogID: 0},
	} {
		if _, err := encodeRaffleEthDepositCursor(cursor); !errors.Is(err, errInvalidRaffleEthDepositCursor) {
			t.Errorf("encode ETH cursor %+v: %v", cursor, err)
		}
	}
	for _, cursor := range []raffleNftWinnerCursor{
		{},
		{Version: 2, Round: 1, Pool: Bidder, WinnerIndex: 0, EventLogID: 1},
		{Version: 1, Round: -1, Pool: Bidder, WinnerIndex: 0, EventLogID: 1},
		{Version: 1, Round: 1, Pool: RaffleNftPool("other"), WinnerIndex: 0, EventLogID: 1},
		{Version: 1, Round: 1, Pool: Bidder, WinnerIndex: -1, EventLogID: 1},
		{Version: 1, Round: 1, Pool: Bidder, WinnerIndex: 0, EventLogID: 0},
	} {
		if _, err := encodeRaffleNftWinnerCursor(cursor); !errors.Is(err, errInvalidRaffleNftWinnerCursor) {
			t.Errorf("encode NFT cursor %+v: %v", cursor, err)
		}
	}
}

func FuzzDecodeRaffleEthDepositCursor(f *testing.F) {
	valid, err := encodeRaffleEthDepositCursor(raffleEthDepositCursor{
		Version:     raffleEthDepositCursorVersion,
		Round:       7,
		WinnerIndex: 2,
		EventLogID:  99,
	})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(valid, int64(7))
	f.Add("", int64(0))
	f.Add("%%%", int64(-1))

	f.Fuzz(func(t *testing.T, encoded string, round int64) {
		cursor, err := decodeRaffleEthDepositCursor(encoded, round)
		if err != nil {
			return
		}
		if !validRaffleEthDepositCursor(cursor, round) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}

func FuzzDecodeRaffleNftWinnerCursor(f *testing.F) {
	valid, err := encodeRaffleNftWinnerCursor(raffleNftWinnerCursor{
		Version:     raffleNftWinnerCursorVersion,
		Round:       7,
		Pool:        Bidder,
		WinnerIndex: 2,
		EventLogID:  99,
	})
	if err != nil {
		f.Fatal(err)
	}
	f.Add(valid, int64(7), string(Bidder))
	f.Add("", int64(0), "")
	f.Add("%%%", int64(-1), "other")

	f.Fuzz(func(t *testing.T, encoded string, round int64, pool string) {
		expectedPool := RaffleNftPool(pool)
		cursor, err := decodeRaffleNftWinnerCursor(encoded, round, expectedPool)
		if err != nil {
			return
		}
		if !validRaffleNftWinnerCursor(cursor, round, expectedPool) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}
