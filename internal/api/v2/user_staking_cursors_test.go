package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestUserStakingTokenCursorRoundTrip(t *testing.T) {
	t.Parallel()
	for _, resource := range []userStakingTokenResource{
		userStakingTokenResourceCstStakedTokens,
		userStakingTokenResourceRwStakedTokens,
		userStakingTokenResourceCstTokenRewards,
	} {
		want := userStakingTokenCursor{
			Version:  userStakingTokenCursorVersion,
			Address:  userCursorAlice,
			Resource: resource,
			TokenID:  5,
		}
		encoded, err := encodeUserStakingTokenCursor(want)
		if err != nil {
			t.Fatalf("encodeUserStakingTokenCursor(%s): %v", resource, err)
		}
		checksummed := strings.ToUpper(userCursorAlice[:2]) + userCursorAlice[2:]
		got, err := decodeUserStakingTokenCursor(encoded, checksummed, resource)
		if err != nil || got != want {
			t.Fatalf("round trip(%s) = %+v, %v; want %+v", resource, got, err, want)
		}
	}
}

func TestUserStakingTokenCursorRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	encoded := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	staked := string(userStakingTokenResourceCstStakedTokens)
	cases := []string{
		"",
		"%%%",
		strings.Repeat("a", maxCursorLength+1),
		encoded(`{"v":2,"a":"` + userCursorAlice + `","k":"` + staked + `","t":1}`),
		encoded(`{"v":1,"a":"bad","k":"` + staked + `","t":1}`),
		encoded(`{"v":1,"a":"` + strings.ToUpper(userCursorAlice) + `","k":"` + staked + `","t":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","k":"unknown","t":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","k":"` + staked + `","t":-1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","k":"` + staked + `","t":1,"x":1}`),
		encoded(`{"v":1,"a":"` + userCursorAlice + `","k":"` + staked + `","t":1}{}`),
	}
	for _, value := range cases {
		if _, err := decodeUserStakingTokenCursor(
			value, userCursorAlice, userStakingTokenResourceCstStakedTokens,
		); !errors.Is(err, errInvalidUserStakingTokenCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}

	valid, err := encodeUserStakingTokenCursor(userStakingTokenCursor{
		Version:  userStakingTokenCursorVersion,
		Address:  userCursorAlice,
		Resource: userStakingTokenResourceCstStakedTokens,
		TokenID:  0,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := decodeUserStakingTokenCursor(
		valid, userCursorBob, userStakingTokenResourceCstStakedTokens,
	); !errors.Is(err, errInvalidUserStakingTokenCursor) {
		t.Errorf("cross-user error = %v", err)
	}
	if _, err := decodeUserStakingTokenCursor(
		valid, userCursorAlice, userStakingTokenResourceCstTokenRewards,
	); !errors.Is(err, errInvalidUserStakingTokenCursor) {
		t.Errorf("cross-resource error = %v", err)
	}
}

func TestEncodeUserStakingTokenCursorRejectsInvalidFields(t *testing.T) {
	t.Parallel()
	for _, cursor := range []userStakingTokenCursor{
		{Version: 2, Address: userCursorAlice, Resource: userStakingTokenResourceCstStakedTokens, TokenID: 1},
		{Version: 1, Address: "bad", Resource: userStakingTokenResourceCstStakedTokens, TokenID: 1},
		{Version: 1, Address: userCursorAlice, Resource: "unknown", TokenID: 1},
		{Version: 1, Address: userCursorAlice, Resource: userStakingTokenResourceCstStakedTokens, TokenID: -1},
	} {
		if _, err := encodeUserStakingTokenCursor(cursor); !errors.Is(err, errInvalidUserStakingTokenCursor) {
			t.Errorf("encode %+v error = %v", cursor, err)
		}
	}
}

func TestUserStakingDepositCursorRoundTripAndRejection(t *testing.T) {
	t.Parallel()
	want := userStakingDepositCursor{
		Version:   userStakingDepositCursorVersion,
		Address:   userCursorAlice,
		DepositID: 501,
	}
	encoded, err := encodeUserStakingDepositCursor(want)
	if err != nil {
		t.Fatalf("encodeUserStakingDepositCursor: %v", err)
	}
	got, err := decodeUserStakingDepositCursor(encoded, userCursorAlice)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}

	base64URL := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cases := []string{
		"",
		"%%%",
		base64URL(`{"v":2,"a":"` + userCursorAlice + `","d":1}`),
		base64URL(`{"v":1,"a":"bad","d":1}`),
		base64URL(`{"v":1,"a":"` + userCursorAlice + `","d":-1}`),
		base64URL(`{"v":1,"a":"` + userCursorAlice + `","d":1,"x":1}`),
	}
	for _, value := range cases {
		if _, err := decodeUserStakingDepositCursor(value, userCursorAlice); !errors.Is(err, errInvalidUserStakingDepositCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	if _, err := decodeUserStakingDepositCursor(encoded, userCursorBob); !errors.Is(err, errInvalidUserStakingDepositCursor) {
		t.Errorf("cross-user error = %v", err)
	}
	if _, err := encodeUserStakingDepositCursor(userStakingDepositCursor{
		Version: userStakingDepositCursorVersion, Address: userCursorAlice, DepositID: -1,
	}); !errors.Is(err, errInvalidUserStakingDepositCursor) {
		t.Error("encode accepted a negative deposit id")
	}
}

func TestUserStakingDepositRewardCursorRoundTripAndRejection(t *testing.T) {
	t.Parallel()
	want := userStakingDepositRewardCursor{
		Version:   userStakingDepositRewardCursorVersion,
		Address:   userCursorAlice,
		DepositID: 501,
		ActionID:  2,
	}
	encoded, err := encodeUserStakingDepositRewardCursor(want)
	if err != nil {
		t.Fatalf("encodeUserStakingDepositRewardCursor: %v", err)
	}
	got, err := decodeUserStakingDepositRewardCursor(encoded, userCursorAlice, 501)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}

	base64URL := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cases := []string{
		"",
		"%%%",
		base64URL(`{"v":2,"a":"` + userCursorAlice + `","d":501,"s":1}`),
		base64URL(`{"v":1,"a":"bad","d":501,"s":1}`),
		base64URL(`{"v":1,"a":"` + userCursorAlice + `","d":-1,"s":1}`),
		base64URL(`{"v":1,"a":"` + userCursorAlice + `","d":501,"s":-1}`),
	}
	for _, value := range cases {
		if _, err := decodeUserStakingDepositRewardCursor(value, userCursorAlice, 501); !errors.Is(err, errInvalidUserStakingDepositRewardCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	if _, err := decodeUserStakingDepositRewardCursor(encoded, userCursorBob, 501); !errors.Is(err, errInvalidUserStakingDepositRewardCursor) {
		t.Errorf("cross-user error = %v", err)
	}
	if _, err := decodeUserStakingDepositRewardCursor(encoded, userCursorAlice, 502); !errors.Is(err, errInvalidUserStakingDepositRewardCursor) {
		t.Errorf("cross-deposit error = %v", err)
	}
	if _, err := encodeUserStakingDepositRewardCursor(userStakingDepositRewardCursor{
		Version: userStakingDepositRewardCursorVersion, Address: userCursorAlice, DepositID: 501, ActionID: -1,
	}); !errors.Is(err, errInvalidUserStakingDepositRewardCursor) {
		t.Error("encode accepted a negative action id")
	}
}

func TestUserStakingTokenDepositCursorRoundTripAndRejection(t *testing.T) {
	t.Parallel()
	want := userStakingTokenDepositCursor{
		Version:   userStakingTokenDepositCursorVersion,
		Address:   userCursorAlice,
		TokenID:   1,
		DepositID: 501,
	}
	encoded, err := encodeUserStakingTokenDepositCursor(want)
	if err != nil {
		t.Fatalf("encodeUserStakingTokenDepositCursor: %v", err)
	}
	got, err := decodeUserStakingTokenDepositCursor(encoded, userCursorAlice, 1)
	if err != nil || got != want {
		t.Fatalf("round trip = %+v, %v; want %+v", got, err, want)
	}

	base64URL := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	cases := []string{
		"",
		"%%%",
		base64URL(`{"v":2,"a":"` + userCursorAlice + `","t":1,"d":501}`),
		base64URL(`{"v":1,"a":"bad","t":1,"d":501}`),
		base64URL(`{"v":1,"a":"` + userCursorAlice + `","t":-1,"d":501}`),
		base64URL(`{"v":1,"a":"` + userCursorAlice + `","t":1,"d":-1}`),
	}
	for _, value := range cases {
		if _, err := decodeUserStakingTokenDepositCursor(value, userCursorAlice, 1); !errors.Is(err, errInvalidUserStakingTokenDepositCursor) {
			t.Errorf("cursor %q error = %v", value, err)
		}
	}
	if _, err := decodeUserStakingTokenDepositCursor(encoded, userCursorBob, 1); !errors.Is(err, errInvalidUserStakingTokenDepositCursor) {
		t.Errorf("cross-user error = %v", err)
	}
	if _, err := decodeUserStakingTokenDepositCursor(encoded, userCursorAlice, 2); !errors.Is(err, errInvalidUserStakingTokenDepositCursor) {
		t.Errorf("cross-token error = %v", err)
	}
	if _, err := encodeUserStakingTokenDepositCursor(userStakingTokenDepositCursor{
		Version: userStakingTokenDepositCursorVersion, Address: userCursorAlice, TokenID: -1, DepositID: 1,
	}); !errors.Is(err, errInvalidUserStakingTokenDepositCursor) {
		t.Error("encode accepted a negative token id")
	}
}

func FuzzDecodeUserStakingTokenCursor(f *testing.F) {
	valid, _ := encodeUserStakingTokenCursor(userStakingTokenCursor{
		Version:  userStakingTokenCursorVersion,
		Address:  userCursorAlice,
		Resource: userStakingTokenResourceCstStakedTokens,
		TokenID:  5,
	})
	f.Add(valid, userCursorAlice, string(userStakingTokenResourceCstStakedTokens))
	f.Add("", "", "")
	f.Add("%%%", userCursorBob, string(userStakingTokenResourceCstTokenRewards))
	f.Fuzz(func(t *testing.T, value, address, resource string) {
		cursor, err := decodeUserStakingTokenCursor(value, address, userStakingTokenResource(resource))
		if err == nil && !validUserStakingTokenCursor(cursor, address, userStakingTokenResource(resource)) {
			t.Fatalf("accepted invalid cursor: %+v for %q/%q", cursor, address, resource)
		}
	})
}

func FuzzDecodeUserStakingDepositCursor(f *testing.F) {
	valid, _ := encodeUserStakingDepositCursor(userStakingDepositCursor{
		Version:   userStakingDepositCursorVersion,
		Address:   userCursorAlice,
		DepositID: 501,
	})
	f.Add(valid, userCursorAlice)
	f.Add("", "")
	f.Add("%%%", userCursorBob)
	f.Fuzz(func(t *testing.T, value, address string) {
		cursor, err := decodeUserStakingDepositCursor(value, address)
		if err == nil && !validUserStakingDepositCursor(cursor, address) {
			t.Fatalf("accepted invalid cursor: %+v for %q", cursor, address)
		}
	})
}

func FuzzDecodeUserStakingDepositRewardCursor(f *testing.F) {
	valid, _ := encodeUserStakingDepositRewardCursor(userStakingDepositRewardCursor{
		Version:   userStakingDepositRewardCursorVersion,
		Address:   userCursorAlice,
		DepositID: 501,
		ActionID:  2,
	})
	f.Add(valid, userCursorAlice, int64(501))
	f.Add("", "", int64(0))
	f.Add("%%%", userCursorBob, int64(-1))
	f.Fuzz(func(t *testing.T, value, address string, depositID int64) {
		cursor, err := decodeUserStakingDepositRewardCursor(value, address, depositID)
		if err == nil && !validUserStakingDepositRewardCursor(cursor, address, depositID) {
			t.Fatalf("accepted invalid cursor: %+v for %q/%d", cursor, address, depositID)
		}
	})
}

func FuzzDecodeUserStakingTokenDepositCursor(f *testing.F) {
	valid, _ := encodeUserStakingTokenDepositCursor(userStakingTokenDepositCursor{
		Version:   userStakingTokenDepositCursorVersion,
		Address:   userCursorAlice,
		TokenID:   1,
		DepositID: 501,
	})
	f.Add(valid, userCursorAlice, int64(1))
	f.Add("", "", int64(0))
	f.Add("%%%", userCursorBob, int64(-1))
	f.Fuzz(func(t *testing.T, value, address string, tokenID int64) {
		cursor, err := decodeUserStakingTokenDepositCursor(value, address, tokenID)
		if err == nil && !validUserStakingTokenDepositCursor(cursor, address, tokenID) {
			t.Fatalf("accepted invalid cursor: %+v for %q/%d", cursor, address, tokenID)
		}
	})
}
