package v2

import (
	"encoding/base64"
	"errors"
	"testing"
)

func TestGlobalStakingEventCursorScopes(t *testing.T) {
	t.Parallel()
	resources := []globalStakingEventResource{
		globalStakingEventCstActions,
		globalStakingEventRwalkActions,
		globalStakingEventCstRaffle,
		globalStakingEventRwalkRaffle,
	}
	for _, resource := range resources {
		want := globalStakingEventCursor{
			Version:    globalStakingEventCursorVersion,
			Resource:   resource,
			EventLogID: 5055,
		}
		encoded, err := encodeGlobalStakingEventCursor(want)
		if err != nil {
			t.Fatalf("encode %q: %v", resource, err)
		}
		got, err := decodeGlobalStakingEventCursor(encoded, resource)
		if err != nil || got != want {
			t.Fatalf("round trip %q = %+v, %v; want %+v", resource, got, err, want)
		}
		for _, other := range resources {
			if other == resource {
				continue
			}
			if _, err := decodeGlobalStakingEventCursor(encoded, other); err == nil {
				t.Errorf("%q cursor decoded as %q", resource, other)
			}
		}
	}
}

func TestGlobalStakingEventCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	payload := func(body string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(body))
	}
	for _, value := range []string{
		"",
		"%%%",
		payload(`{"v":2,"k":"cstActions","e":1}`),
		payload(`{"v":1,"k":"unknown","e":1}`),
		payload(`{"v":1,"k":"cstActions","e":0}`),
		payload(`{"v":1,"k":"cstActions","e":1,"x":1}`),
		payload(`{"v":1,"k":"cstActions","e":1}{}`),
		payload(`{"v":1,"d":1}`),
	} {
		if _, err := decodeGlobalStakingEventCursor(
			value,
			globalStakingEventCstActions,
		); !errors.Is(err, errInvalidGlobalStakingEventCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	if _, err := encodeGlobalStakedTokenCursor(globalStakedTokenCursor{
		Version:  globalStakedTokenCursorVersion,
		Resource: "unknown",
		TokenID:  1,
	}); !errors.Is(err, errInvalidGlobalStakedTokenCursor) {
		t.Errorf("encode invalid staked-token cursor error = %v", err)
	}
	for _, cursor := range []globalStakingEventCursor{
		{Version: 2, Resource: globalStakingEventCstActions, EventLogID: 1},
		{Version: 1, Resource: "unknown", EventLogID: 1},
		{Version: 1, Resource: globalStakingEventCstActions, EventLogID: 0},
	} {
		if _, err := encodeGlobalStakingEventCursor(cursor); !errors.Is(
			err,
			errInvalidGlobalStakingEventCursor,
		) {
			t.Errorf("encode %+v error = %v", cursor, err)
		}
	}
}

func TestGlobalStakedTokenCursorScopes(t *testing.T) {
	t.Parallel()
	for _, resource := range []globalStakedTokenResource{
		globalStakedTokenCst,
		globalStakedTokenRwalk,
	} {
		want := globalStakedTokenCursor{
			Version:  globalStakedTokenCursorVersion,
			Resource: resource,
			TokenID:  42,
		}
		encoded, err := encodeGlobalStakedTokenCursor(want)
		if err != nil {
			t.Fatalf("encode %q: %v", resource, err)
		}
		got, err := decodeGlobalStakedTokenCursor(encoded, resource)
		if err != nil || got != want {
			t.Fatalf("round trip %q = %+v, %v; want %+v", resource, got, err, want)
		}
		other := globalStakedTokenCst
		if resource == other {
			other = globalStakedTokenRwalk
		}
		if _, err := decodeGlobalStakedTokenCursor(encoded, other); err == nil {
			t.Errorf("%q cursor decoded as %q", resource, other)
		}
	}

	payload := func(body string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(body))
	}
	for _, value := range []string{
		"",
		payload(`{"v":2,"k":"cst","t":1}`),
		payload(`{"v":1,"k":"unknown","t":1}`),
		payload(`{"v":1,"k":"cst","t":-1}`),
		payload(`{"v":1,"k":"cst","t":1,"x":1}`),
		payload(`{"v":1,"k":"cstActions","e":1}`),
	} {
		if _, err := decodeGlobalStakedTokenCursor(
			value,
			globalStakedTokenCst,
		); !errors.Is(err, errInvalidGlobalStakedTokenCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	if _, err := encodeRoundStakingRewardCursor(roundStakingRewardCursor{
		Version:   roundStakingRewardCursorVersion,
		Round:     1,
		DepositID: 1,
		StakerAid: 0,
	}); !errors.Is(err, errInvalidRoundStakingRewardCursor) {
		t.Errorf("encode invalid round-reward cursor error = %v", err)
	}
}

func TestGlobalStakingDepositCursor(t *testing.T) {
	t.Parallel()
	want := globalStakingDepositCursor{
		Version:   globalStakingDepositCursorVersion,
		DepositID: 501,
	}
	encoded, err := encodeGlobalStakingDepositCursor(want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := decodeGlobalStakingDepositCursor(encoded)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}
	payload := func(body string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(body))
	}
	for _, value := range []string{
		"",
		payload(`{"v":2,"d":1}`),
		payload(`{"v":1,"d":-1}`),
		payload(`{"v":1,"d":1,"x":1}`),
		payload(`{"v":1,"s":1}`),
	} {
		if _, err := decodeGlobalStakingDepositCursor(value); !errors.Is(
			err,
			errInvalidGlobalStakingDepositCursor,
		) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	if _, err := encodeGlobalStakingDepositCursor(globalStakingDepositCursor{
		Version:   globalStakingDepositCursorVersion,
		DepositID: -1,
	}); !errors.Is(err, errInvalidGlobalStakingDepositCursor) {
		t.Errorf("encode invalid deposit cursor error = %v", err)
	}
}

func TestRoundStakingRewardCursorScopesRound(t *testing.T) {
	t.Parallel()
	want := roundStakingRewardCursor{
		Version:   roundStakingRewardCursorVersion,
		Round:     7,
		DepositID: 9,
		StakerAid: 21,
	}
	encoded, err := encodeRoundStakingRewardCursor(want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := decodeRoundStakingRewardCursor(encoded, 7)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}
	if _, err := decodeRoundStakingRewardCursor(encoded, 8); err == nil {
		t.Error("round-scoped cursor decoded under another round")
	}
	payload := func(body string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(body))
	}
	for _, value := range []string{
		"",
		payload(`{"v":2,"r":7,"d":9,"a":21}`),
		payload(`{"v":1,"r":-1,"d":9,"a":21}`),
		payload(`{"v":1,"r":7,"d":-1,"a":21}`),
		payload(`{"v":1,"r":7,"d":9,"a":0}`),
		payload(`{"v":1,"r":7,"d":9,"a":21,"x":1}`),
		payload(`{"v":1,"d":9}`),
	} {
		if _, err := decodeRoundStakingRewardCursor(
			value,
			7,
		); !errors.Is(err, errInvalidRoundStakingRewardCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
}

func FuzzDecodeGlobalStakingEventCursor(f *testing.F) {
	valid, _ := encodeGlobalStakingEventCursor(globalStakingEventCursor{
		Version:    globalStakingEventCursorVersion,
		Resource:   globalStakingEventCstActions,
		EventLogID: 5055,
	})
	f.Add(valid, "cstActions")
	f.Add("", "unknown")
	f.Fuzz(func(t *testing.T, value, resource string) {
		expected := globalStakingEventResource(resource)
		cursor, err := decodeGlobalStakingEventCursor(value, expected)
		if err == nil && !validGlobalStakingEventCursor(cursor, expected) {
			t.Fatalf("accepted invalid cursor: %+v for %q", cursor, resource)
		}
	})
}

func FuzzDecodeGlobalStakedTokenCursor(f *testing.F) {
	valid, _ := encodeGlobalStakedTokenCursor(globalStakedTokenCursor{
		Version:  globalStakedTokenCursorVersion,
		Resource: globalStakedTokenCst,
		TokenID:  5,
	})
	f.Add(valid, "cst")
	f.Add("", "unknown")
	f.Fuzz(func(t *testing.T, value, resource string) {
		expected := globalStakedTokenResource(resource)
		cursor, err := decodeGlobalStakedTokenCursor(value, expected)
		if err == nil && !validGlobalStakedTokenCursor(cursor, expected) {
			t.Fatalf("accepted invalid cursor: %+v for %q", cursor, resource)
		}
	})
}

func FuzzDecodeGlobalStakingDepositCursor(f *testing.F) {
	valid, _ := encodeGlobalStakingDepositCursor(globalStakingDepositCursor{
		Version:   globalStakingDepositCursorVersion,
		DepositID: 1,
	})
	f.Add(valid)
	f.Add("")
	f.Fuzz(func(t *testing.T, value string) {
		cursor, err := decodeGlobalStakingDepositCursor(value)
		if err == nil && !validGlobalStakingDepositCursor(cursor) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}

func FuzzDecodeRoundStakingRewardCursor(f *testing.F) {
	valid, _ := encodeRoundStakingRewardCursor(roundStakingRewardCursor{
		Version:   roundStakingRewardCursorVersion,
		Round:     1,
		DepositID: 2,
		StakerAid: 3,
	})
	f.Add(valid, int64(1))
	f.Add("", int64(-1))
	f.Fuzz(func(t *testing.T, value string, round int64) {
		cursor, err := decodeRoundStakingRewardCursor(value, round)
		if err == nil && !validRoundStakingRewardCursor(cursor, round) {
			t.Fatalf("accepted invalid cursor: %+v for round %d", cursor, round)
		}
	})
}
