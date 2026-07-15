package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestUserOwnedTokenCursorRoundTrip(t *testing.T) {
	t.Parallel()
	want := userOwnedTokenCursor{
		Version: userOwnedTokenCursorVersion,
		Address: userCursorAlice,
		TokenID: 5,
	}
	encoded, err := encodeUserOwnedTokenCursor(want)
	if err != nil {
		t.Fatalf("encodeUserOwnedTokenCursor: %v", err)
	}
	checksummed := strings.ToUpper(userCursorAlice[:2]) + userCursorAlice[2:]
	got, err := decodeUserOwnedTokenCursor(encoded, checksummed)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}
}

func TestUserOwnedTokenCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	encoded := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cases := []string{
		"",
		"%%%",
		strings.Repeat("a", maxCursorLength+1),
		encoded(`{"v":2,"a":"` + userCursorAlice + `","t":1}`),
		encoded(`{"v":1,"a":"bad","t":1}`),
		encoded(`{"v":1,"a":"` + strings.ToUpper(userCursorAlice) + `","t":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","t":-1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","t":1,"x":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","t":1}{}`),
		// Payloads of sibling cursor types must fail structurally.
		encoded(`{"v":1,"a":"` + userCursorAlice + `","k":"cstStakedTokens","t":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","d":1}`),
	}
	for _, value := range cases {
		if _, err := decodeUserOwnedTokenCursor(value, userCursorAlice); !errors.Is(err, errInvalidUserOwnedTokenCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}

	valid, err := encodeUserOwnedTokenCursor(userOwnedTokenCursor{
		Version: userOwnedTokenCursorVersion,
		Address: userCursorAlice,
		TokenID: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeUserOwnedTokenCursor(valid, userCursorBob); !errors.Is(err, errInvalidUserOwnedTokenCursor) {
		t.Errorf("cross-user error = %v", err)
	}
}

func TestEncodeUserOwnedTokenCursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()
	for _, cursor := range []userOwnedTokenCursor{
		{Version: 2, Address: userCursorAlice, TokenID: 1},
		{Version: 1, Address: "bad", TokenID: 1},
		{Version: 1, Address: strings.ToUpper(userCursorAlice), TokenID: 1},
		{Version: 1, Address: userCursorAlice, TokenID: -1},
	} {
		if _, err := encodeUserOwnedTokenCursor(cursor); !errors.Is(err, errInvalidUserOwnedTokenCursor) {
			t.Errorf("encode %+v error = %v", cursor, err)
		}
	}
}

func TestUserEventCursorAcceptsActivityResources(t *testing.T) {
	t.Parallel()
	for _, resource := range []userEventResource{
		userEventResourceCsTransfers,
		userEventResourceCtTransfers,
		userEventResourceMarketingRewards,
	} {
		want := userEventCursor{
			Version:    userEventCursorVersion,
			Address:    userCursorAlice,
			Resource:   resource,
			EventLogID: 5049,
		}
		encoded, err := encodeUserEventCursor(want)
		if err != nil {
			t.Fatalf("encodeUserEventCursor(%s): %v", resource, err)
		}
		got, err := decodeUserEventCursor(encoded, userCursorAlice, resource)
		if err != nil || got != want {
			t.Fatalf("round trip(%s) = %+v, %v; want %+v", resource, got, err, want)
		}
		// The same payload must fail against every other activity scope.
		for _, other := range []userEventResource{
			userEventResourceCsTransfers,
			userEventResourceCtTransfers,
			userEventResourceMarketingRewards,
			userEventResourceCstStakingActions,
		} {
			if other == resource {
				continue
			}
			if _, err := decodeUserEventCursor(encoded, userCursorAlice, other); err == nil {
				t.Errorf("cursor for %s decoded under %s", resource, other)
			}
		}
	}
}

func FuzzDecodeUserOwnedTokenCursor(f *testing.F) {
	valid, _ := encodeUserOwnedTokenCursor(userOwnedTokenCursor{
		Version: userOwnedTokenCursorVersion,
		Address: userCursorAlice,
		TokenID: 5,
	})
	f.Add(valid, userCursorAlice)
	f.Add("", "")
	f.Add("%%%", userCursorBob)
	f.Fuzz(func(t *testing.T, value, address string) {
		cursor, err := decodeUserOwnedTokenCursor(value, address)
		if err == nil && !validUserOwnedTokenCursor(cursor, address) {
			t.Fatalf("accepted invalid cursor: %+v for %q", cursor, address)
		}
	})
}
