package cosmicgame

import (
	"testing"
)

func TestCanonicalParticipantSortValue(t *testing.T) {
	t.Parallel()
	for _, value := range []string{"0", "1", "999999999999999999999999999999999999"} {
		if !canonicalParticipantSortValue(value) {
			t.Errorf("canonicalParticipantSortValue(%q) = false", value)
		}
	}
	for _, value := range []string{"", "-1", "+1", "01", "1.0", "1e3", "not-a-number"} {
		if canonicalParticipantSortValue(value) {
			t.Errorf("canonicalParticipantSortValue(%q) = true", value)
		}
	}
}

func TestParticipantSortValueBounds(t *testing.T) {
	t.Parallel()
	const tooLargeForInt64 = "9223372036854775808"
	for _, kind := range []ParticipantKind{
		ParticipantBidders,
		ParticipantWinners,
		ParticipantRandomWalkStakers,
		ParticipantDualStakers,
	} {
		if validParticipantSortValue(kind, tooLargeForInt64) {
			t.Errorf("%s accepted an overflowing count", kind)
		}
	}
	for _, kind := range []ParticipantKind{ParticipantDonors, ParticipantCSTStakers} {
		if !validParticipantSortValue(kind, tooLargeForInt64) {
			t.Errorf("%s rejected a valid exact wei value", kind)
		}
	}
}

func TestParticipantPageMethodsRejectInvalidArguments(t *testing.T) {
	t.Parallel()
	var r Repo
	type pageCall func(*ParticipantPageCursor, int) error
	calls := map[ParticipantKind]pageCall{
		ParticipantBidders: func(after *ParticipantPageCursor, limit int) error {
			_, _, err := r.BidderParticipantsPage(t.Context(), after, limit)
			return err
		},
		ParticipantWinners: func(after *ParticipantPageCursor, limit int) error {
			_, _, err := r.WinnerParticipantsPage(t.Context(), after, limit)
			return err
		},
		ParticipantDonors: func(after *ParticipantPageCursor, limit int) error {
			_, _, err := r.DonorParticipantsPage(t.Context(), after, limit)
			return err
		},
		ParticipantCSTStakers: func(after *ParticipantPageCursor, limit int) error {
			_, _, err := r.CSTStakerParticipantsPage(t.Context(), after, limit)
			return err
		},
		ParticipantRandomWalkStakers: func(after *ParticipantPageCursor, limit int) error {
			_, _, err := r.RandomWalkStakerParticipantsPage(t.Context(), after, limit)
			return err
		},
		ParticipantDualStakers: func(after *ParticipantPageCursor, limit int) error {
			_, _, err := r.DualStakerParticipantsPage(t.Context(), after, limit)
			return err
		},
	}
	for kind, call := range calls {
		t.Run(string(kind), func(t *testing.T) {
			t.Parallel()
			for _, limit := range []int{-1, 0, maxStatisticsPageLimit + 1} {
				if err := call(nil, limit); err == nil {
					t.Errorf("accepted limit %d", limit)
				}
			}
			invalid := []*ParticipantPageCursor{
				{Kind: ParticipantKind("other"), SortValue: "1", AddressID: 1},
				{Kind: kind, SortValue: "", AddressID: 1},
				{Kind: kind, SortValue: "-1", AddressID: 1},
				{Kind: kind, SortValue: "01", AddressID: 1},
				{Kind: kind, SortValue: "1", AddressID: 0},
			}
			for _, cursor := range invalid {
				if err := call(cursor, 1); err == nil {
					t.Errorf("accepted cursor %+v", cursor)
				}
			}
		})
	}
}

func TestTrimParticipantPage(t *testing.T) {
	t.Parallel()
	rows := []int{1, 2, 3}
	page, more := trimParticipantPage(rows, 2)
	if !more || len(page) != 2 || page[0] != 1 || page[1] != 2 {
		t.Fatalf("trimParticipantPage = %v,%v", page, more)
	}
	page, more = trimParticipantPage(rows[:2], 2)
	if more || len(page) != 2 {
		t.Fatalf("exact page = %v,%v", page, more)
	}
}
