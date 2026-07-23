package cosmicgame

import (
	"errors"
	"testing"
)

func TestV3StoreRejectsInvalidIdentitiesBeforeDatabaseAccess(t *testing.T) {
	t.Parallel()
	var repo Repo
	if err := repo.UpdateRoundChampionDurations(t.Context(), -1, 0, 0); err == nil {
		t.Fatal("negative champion round was accepted")
	}
	if _, err := repo.InsertLiveStateUpdateIfChanged(
		t.Context(), "", 0, -2, -1, -1, "bad",
	); err == nil {
		t.Fatal("invalid live-state identity was accepted")
	}
	if _, err := repo.BidRewardsByBidID(t.Context(), 0); err == nil {
		t.Fatal("invalid bid reward id was accepted")
	}
}

func TestEffectiveThisBidderReward(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		split, mint, total, want string
	}{
		{"10", "11", "12", "10"},
		{"", "11", "12", "11"},
		{"-1", "", "12", "12"},
	} {
		if got := effectiveThisBidderReward(test.split, test.mint, test.total); got != test.want {
			t.Errorf("effective reward = %q, want %q", got, test.want)
		}
	}
}

func TestConfiguredBidReward(t *testing.T) {
	t.Parallel()
	if got, err := configuredBidReward("12", nil); err != nil || got != "12" {
		t.Fatalf("configured reward = %q, %v", got, err)
	}
	for _, test := range []struct {
		value string
		err   error
	}{
		{err: errors.New("db down")},
		{value: ""},
		{value: "0"},
	} {
		if _, err := configuredBidReward(test.value, test.err); err == nil {
			t.Errorf("invalid configured reward %q/%v was accepted", test.value, test.err)
		}
	}
}
