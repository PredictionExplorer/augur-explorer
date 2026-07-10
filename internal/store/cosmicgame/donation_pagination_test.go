package cosmicgame

import (
	"strings"
	"testing"
)

func TestDonationPageMethodsRejectInvalidArguments(t *testing.T) {
	t.Parallel()

	var repo Repo
	methods := map[string]func(int64, *DonationPageCursor, int) error{
		"eth": func(round int64, after *DonationPageCursor, limit int) error {
			_, _, err := repo.EthDonationsByRoundPage(t.Context(), round, after, limit)
			return err
		},
		"erc20": func(round int64, after *DonationPageCursor, limit int) error {
			_, _, err := repo.ERC20DonationsByRoundPage(t.Context(), round, after, limit)
			return err
		},
		"nft": func(round int64, after *DonationPageCursor, limit int) error {
			_, _, err := repo.NFTDonationsByRoundPage(t.Context(), round, after, limit)
			return err
		},
	}

	for name, method := range methods {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := method(-1, nil, 1); err == nil {
				t.Error("accepted a negative round")
			}
			for _, limit := range []int{-1, 0, maxDonationPageLimit + 1} {
				if err := method(1, nil, limit); err == nil {
					t.Errorf("accepted limit %d", limit)
				}
			}
			if err := method(1, &DonationPageCursor{}, 1); err == nil {
				t.Error("accepted a zero event-log cursor")
			}
			if err := method(1, &DonationPageCursor{EventLogID: -1}, 1); err == nil {
				t.Error("accepted a negative event-log cursor")
			}
		})
	}
}

func TestRoundEthDonationsPageSQLUsesCursorInBothBranches(t *testing.T) {
	t.Parallel()

	withoutCursor := roundEthDonationsPageSQL(false)
	withCursor := roundEthDonationsPageSQL(true)
	if withoutCursor == withCursor {
		t.Fatal("cursor and initial-page SQL are identical")
	}
	const cursorPredicate = "d.evtlog_id < $2"
	if count := strings.Count(withCursor, cursorPredicate); count != 2 {
		t.Fatalf("cursor predicate count = %d, want 2", count)
	}
	if count := strings.Count(withoutCursor, cursorPredicate); count != 0 {
		t.Fatalf("initial-page cursor predicate count = %d, want 0", count)
	}
}
