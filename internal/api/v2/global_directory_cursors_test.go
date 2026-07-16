package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestGlobalTokenCursorRoundTripPerFilter(t *testing.T) {
	t.Parallel()
	filters := []globalTokenFilterScope{
		{},
		{Named: true},
		{Name: "cosmic"},
	}
	for _, filter := range filters {
		want := globalTokenCursor{
			Version: globalTokenCursorVersion,
			Filter:  filter,
			TokenID: 42,
		}
		encoded, err := encodeGlobalTokenCursor(want)
		if err != nil {
			t.Fatalf("encodeGlobalTokenCursor(%+v): %v", filter, err)
		}
		got, err := decodeGlobalTokenCursor(encoded, filter)
		if err != nil || got != want {
			t.Fatalf("round trip(%+v) = %+v, %v; want %+v", filter, got, err, want)
		}
		// The same payload must fail against every other filter scope.
		for _, other := range filters {
			if other == filter {
				continue
			}
			if _, err := decodeGlobalTokenCursor(encoded, other); err == nil {
				t.Errorf("cursor for %+v decoded under %+v", filter, other)
			}
		}
		if _, err := decodeGlobalTokenCursor(encoded, globalTokenFilterScope{Name: "other"}); err == nil {
			t.Error("cursor decoded under a different search term")
		}
	}
}

func TestGlobalTokenCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	encoded := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cases := []string{
		"",
		"%%%",
		strings.Repeat("a", maxCursorLength+1),
		encoded(`{"v":2,"f":{},"t":1}`),
		encoded(`{"v":1,"f":{},"t":-1}`),
		encoded(`{"v":1,"f":{"n":true,"q":"x"},"t":1}`),
		encoded(`{"v":1,"f":{"q":"` + strings.Repeat("a", maxTokenNameSearchLength+1) + `"},"t":1}`),
		encoded(`{"v":1,"f":{},"t":1,"x":1}`),
		encoded(`{"v":1,"f":{},"t":1}{}`),
		// Payloads of sibling cursor types must fail structurally.
		encoded(`{"v":1,"a":"0x00000000000000000000000000000000000000aa","t":1}`),
		encoded(`{"v":1,"s":9}`),
	}
	for _, value := range cases {
		if _, err := decodeGlobalTokenCursor(value, globalTokenFilterScope{}); !errors.Is(err, errInvalidGlobalTokenCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	for _, cursor := range []globalTokenCursor{
		{Version: 2, TokenID: 1},
		{Version: 1, TokenID: -1},
		{Version: 1, Filter: globalTokenFilterScope{Named: true, Name: "x"}, TokenID: 1},
	} {
		if _, err := encodeGlobalTokenCursor(cursor); !errors.Is(err, errInvalidGlobalTokenCursor) {
			t.Errorf("encode %+v error = %v", cursor, err)
		}
	}
}

func TestTokenEventCursorScopes(t *testing.T) {
	t.Parallel()
	want := tokenEventCursor{
		Version:    tokenEventCursorVersion,
		Resource:   tokenEventResourceNameHistory,
		TokenID:    7,
		EventLogID: 5100,
	}
	encoded, err := encodeTokenEventCursor(want)
	if err != nil {
		t.Fatalf("encodeTokenEventCursor: %v", err)
	}
	got, err := decodeTokenEventCursor(encoded, 7, tokenEventResourceNameHistory)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}
	if _, err := decodeTokenEventCursor(encoded, 7, tokenEventResourceTransfers); err == nil {
		t.Error("name-history cursor decoded under transfers")
	}
	if _, err := decodeTokenEventCursor(encoded, 8, tokenEventResourceNameHistory); err == nil {
		t.Error("cursor decoded under another token")
	}

	payload := func(body string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(body))
	}
	for _, value := range []string{
		"",
		payload(`{"v":2,"k":"nameHistory","t":7,"e":1}`),
		payload(`{"v":1,"k":"unknown","t":7,"e":1}`),
		payload(`{"v":1,"k":"nameHistory","t":-1,"e":1}`),
		payload(`{"v":1,"k":"nameHistory","t":7,"e":0}`),
		payload(`{"v":1,"k":"nameHistory","t":7,"e":1,"x":1}`),
		// The user event cursor's key set must fail structurally.
		payload(`{"v":1,"a":"0x00000000000000000000000000000000000000aa","k":"nameHistory","e":1}`),
	} {
		if _, err := decodeTokenEventCursor(value, 7, tokenEventResourceNameHistory); !errors.Is(err, errInvalidTokenEventCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	if _, err := encodeTokenEventCursor(tokenEventCursor{
		Version:    tokenEventCursorVersion,
		Resource:   "unknown",
		TokenID:    7,
		EventLogID: 1,
	}); !errors.Is(err, errInvalidTokenEventCursor) {
		t.Errorf("encode unknown resource error = %v", err)
	}
}

func TestSupplyChangeAndGlobalMarketingCursors(t *testing.T) {
	t.Parallel()
	supply, err := encodeSupplyChangeCursor(supplyChangeCursor{
		Version:    supplyChangeCursorVersion,
		EventLogID: 5001,
	})
	if err != nil {
		t.Fatalf("encodeSupplyChangeCursor: %v", err)
	}
	if got, err := decodeSupplyChangeCursor(supply); err != nil || got.EventLogID != 5001 {
		t.Fatalf("supply round trip = %+v, %v", got, err)
	}
	marketing, err := encodeGlobalMarketingCursor(globalMarketingCursor{
		Version:    globalMarketingCursorVersion,
		EventLogID: 5002,
	})
	if err != nil {
		t.Fatalf("encodeGlobalMarketingCursor: %v", err)
	}
	if got, err := decodeGlobalMarketingCursor(marketing); err != nil || got.EventLogID != 5002 {
		t.Fatalf("marketing round trip = %+v, %v", got, err)
	}

	// The two single-key cursors must not decode as each other.
	if _, err := decodeSupplyChangeCursor(marketing); !errors.Is(err, errInvalidSupplyChangeCursor) {
		t.Errorf("marketing cursor decoded as supply cursor: %v", err)
	}
	if _, err := decodeGlobalMarketingCursor(supply); !errors.Is(err, errInvalidGlobalMarketingCursor) {
		t.Errorf("supply cursor decoded as marketing cursor: %v", err)
	}

	payload := func(body string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(body))
	}
	for _, value := range []string{"", payload(`{"v":2,"s":1}`), payload(`{"v":1,"s":0}`)} {
		if _, err := decodeSupplyChangeCursor(value); !errors.Is(err, errInvalidSupplyChangeCursor) {
			t.Errorf("supply cursor %q error = %v", value, err)
		}
	}
	for _, value := range []string{"", payload(`{"v":2,"m":1}`), payload(`{"v":1,"m":0}`)} {
		if _, err := decodeGlobalMarketingCursor(value); !errors.Is(err, errInvalidGlobalMarketingCursor) {
			t.Errorf("marketing cursor %q error = %v", value, err)
		}
	}
	if _, err := encodeSupplyChangeCursor(supplyChangeCursor{Version: 1}); !errors.Is(err, errInvalidSupplyChangeCursor) {
		t.Errorf("encode zero supply cursor error = %v", err)
	}
	if _, err := encodeGlobalMarketingCursor(globalMarketingCursor{Version: 1}); !errors.Is(err, errInvalidGlobalMarketingCursor) {
		t.Errorf("encode zero marketing cursor error = %v", err)
	}
}

func FuzzDecodeGlobalTokenCursor(f *testing.F) {
	all, _ := encodeGlobalTokenCursor(globalTokenCursor{
		Version: globalTokenCursorVersion,
		TokenID: 5,
	})
	named, _ := encodeGlobalTokenCursor(globalTokenCursor{
		Version: globalTokenCursorVersion,
		Filter:  globalTokenFilterScope{Named: true},
		TokenID: 5,
	})
	search, _ := encodeGlobalTokenCursor(globalTokenCursor{
		Version: globalTokenCursorVersion,
		Filter:  globalTokenFilterScope{Name: "cosmic"},
		TokenID: 5,
	})
	f.Add(all, false, "")
	f.Add(named, true, "")
	f.Add(search, false, "cosmic")
	f.Add("", false, "x")
	f.Fuzz(func(t *testing.T, value string, named bool, name string) {
		expected := globalTokenFilterScope{Named: named, Name: name}
		cursor, err := decodeGlobalTokenCursor(value, expected)
		if err == nil && !validGlobalTokenCursor(cursor, expected) {
			t.Fatalf("accepted invalid cursor: %+v for %+v", cursor, expected)
		}
	})
}

func FuzzDecodeTokenEventCursor(f *testing.F) {
	valid, _ := encodeTokenEventCursor(tokenEventCursor{
		Version:    tokenEventCursorVersion,
		Resource:   tokenEventResourceTransfers,
		TokenID:    3,
		EventLogID: 5100,
	})
	f.Add(valid, int64(3), "transfers")
	f.Add("", int64(0), "nameHistory")
	f.Add("%%%", int64(-1), "unknown")
	f.Fuzz(func(t *testing.T, value string, tokenID int64, resource string) {
		expected := tokenEventResource(resource)
		cursor, err := decodeTokenEventCursor(value, tokenID, expected)
		if err == nil && !validTokenEventCursor(cursor, tokenID, expected) {
			t.Fatalf("accepted invalid cursor: %+v for token %d resource %q",
				cursor, tokenID, resource)
		}
	})
}

func FuzzDecodeSupplyChangeCursor(f *testing.F) {
	valid, _ := encodeSupplyChangeCursor(supplyChangeCursor{
		Version:    supplyChangeCursorVersion,
		EventLogID: 5001,
	})
	f.Add(valid)
	f.Add("")
	f.Add("%%%")
	f.Fuzz(func(t *testing.T, value string) {
		cursor, err := decodeSupplyChangeCursor(value)
		if err == nil && !validSupplyChangeCursor(cursor) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}

func FuzzDecodeGlobalMarketingCursor(f *testing.F) {
	valid, _ := encodeGlobalMarketingCursor(globalMarketingCursor{
		Version:    globalMarketingCursorVersion,
		EventLogID: 5002,
	})
	f.Add(valid)
	f.Add("")
	f.Add("%%%")
	f.Fuzz(func(t *testing.T, value string) {
		cursor, err := decodeGlobalMarketingCursor(value)
		if err == nil && !validGlobalMarketingCursor(cursor) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}
