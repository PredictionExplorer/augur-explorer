package v2

import (
	"encoding/base64"
	"errors"
	"strings"
	"testing"
)

func TestROILeaderboardCursorRoundTrip(t *testing.T) {
	t.Parallel()
	for _, sort := range []RoiLeaderboardSort{NetProfit, Roi, WinRate, Spent, Nfts, Bids} {
		secondary := int64(0)
		if sort == WinRate {
			secondary = 3
		}
		want := roiLeaderboardCursor{
			Version: roiLeaderboardCursorVersion, Sort: sort, MinBids: 5,
			SortValue: "-0.5", Secondary: secondary, BidderAid: 21,
		}
		encoded, err := encodeROILeaderboardCursor(want)
		if err != nil {
			t.Fatalf("encode %s: %v", sort, err)
		}
		got, err := decodeROILeaderboardCursor(encoded, sort, 5)
		if err != nil || got != want {
			t.Fatalf("round trip %s = %+v, %v; want %+v", sort, got, err, want)
		}
	}
}

func TestStatisticsCursorsRejectInvalidInput(t *testing.T) {
	t.Parallel()
	encoded := func(payload string) string {
		return base64.RawURLEncoding.EncodeToString([]byte(payload))
	}
	roiCases := []string{
		"",
		"%%%",
		strings.Repeat("a", maxCursorLength+1),
		encoded(`{"v":2,"s":"roi","m":5,"k":"1","x":0,"a":1}`),
		encoded(`{"v":1,"s":"other","m":5,"k":"1","x":0,"a":1}`),
		encoded(`{"v":1,"s":"roi","m":5,"k":"bad","x":0,"a":1}`),
		encoded(`{"v":1,"s":"roi","m":5,"k":"1","x":0,"a":0}`),
		encoded(`{"v":1,"s":"roi","m":5,"k":"1","x":0,"a":1,"z":1}`),
	}
	for _, value := range roiCases {
		if _, err := decodeROILeaderboardCursor(value, Roi, 5); !errors.Is(err, errInvalidROILeaderboardCursor) {
			t.Errorf("ROI cursor %q error = %v", value, err)
		}
	}
	crossSort, _ := encodeROILeaderboardCursor(roiLeaderboardCursor{
		Version: 1, Sort: Roi, MinBids: 5, SortValue: "1", BidderAid: 1,
	})
	if _, err := decodeROILeaderboardCursor(crossSort, Spent, 5); !errors.Is(err, errInvalidROILeaderboardCursor) {
		t.Errorf("cross-sort error = %v", err)
	}
	if _, err := decodeROILeaderboardCursor(crossSort, Roi, 6); !errors.Is(err, errInvalidROILeaderboardCursor) {
		t.Errorf("cross-filter error = %v", err)
	}
	if _, err := decodeClaimSummaryCursor(encoded(`{"v":1,"r":-1,"e":1}`)); !errors.Is(err, errInvalidClaimSummaryCursor) {
		t.Errorf("claim summary error = %v", err)
	}
	if _, err := decodeClaimDetailCursor(
		encoded(`{"v":1,"r":2,"s":"attached","e":10,"g":0,"k":0}`),
		2,
		claimDetailTransactions,
	); !errors.Is(err, errInvalidClaimDetailCursor) {
		t.Errorf("cross-section error = %v", err)
	}
}

func TestClaimCursorsRoundTrip(t *testing.T) {
	t.Parallel()
	summary := claimSummaryCursor{Version: 1, Round: 2, EventLogID: 100}
	encodedSummary, err := encodeClaimSummaryCursor(summary)
	if err != nil {
		t.Fatal(err)
	}
	if got, err := decodeClaimSummaryCursor(encodedSummary); err != nil || got != summary {
		t.Fatalf("summary = %+v, %v", got, err)
	}
	for _, cursor := range []claimDetailCursor{
		{Version: 1, Round: 2, Section: claimDetailTransactions, EventLogID: 100},
		{Version: 1, Round: 2, Section: claimDetailAttached, EventLogID: 100},
		{Version: 1, Round: 2, Section: claimDetailUnclaimed, Segment: 1, Key: 100},
	} {
		encodedCursor, err := encodeClaimDetailCursor(cursor)
		if err != nil {
			t.Fatal(err)
		}
		got, err := decodeClaimDetailCursor(encodedCursor, cursor.Round, cursor.Section)
		if err != nil || got != cursor {
			t.Fatalf("detail = %+v, %v; want %+v", got, err, cursor)
		}
	}
}

func FuzzDecodeROILeaderboardCursor(f *testing.F) {
	valid, _ := encodeROILeaderboardCursor(roiLeaderboardCursor{
		Version: 1, Sort: Roi, MinBids: 5, SortValue: "0.5", BidderAid: 21,
	})
	f.Add(valid, string(Roi), 5)
	f.Add("", "", -1)
	f.Fuzz(func(t *testing.T, value, sort string, minBids int) {
		expected := RoiLeaderboardSort(sort)
		cursor, err := decodeROILeaderboardCursor(value, expected, minBids)
		if err == nil && !validROILeaderboardCursor(cursor, expected, minBids) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}

func FuzzDecodeClaimSummaryCursor(f *testing.F) {
	valid, _ := encodeClaimSummaryCursor(claimSummaryCursor{Version: 1, Round: 2, EventLogID: 100})
	f.Add(valid)
	f.Add("")
	f.Fuzz(func(t *testing.T, value string) {
		cursor, err := decodeClaimSummaryCursor(value)
		if err == nil && !validClaimSummaryCursor(cursor) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}

func FuzzDecodeClaimDetailCursor(f *testing.F) {
	valid, _ := encodeClaimDetailCursor(claimDetailCursor{
		Version: 1, Round: 2, Section: claimDetailTransactions, EventLogID: 100,
	})
	f.Add(valid, int64(2), string(claimDetailTransactions))
	f.Add("", int64(-1), "")
	f.Fuzz(func(t *testing.T, value string, round int64, section string) {
		expected := claimDetailSection(section)
		cursor, err := decodeClaimDetailCursor(value, round, expected)
		if err == nil && !validClaimDetailCursor(cursor, round, expected) {
			t.Fatalf("accepted invalid cursor: %+v", cursor)
		}
	})
}
