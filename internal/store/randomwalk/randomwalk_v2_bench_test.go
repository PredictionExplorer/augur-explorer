//go:build integration

package randomwalk

// Benchmarks for the heaviest v2 RandomWalk read paths (§4.5 of
// docs/MODERNIZATION.md), run against the seeded test container. Baselines
// live in docs/benchmarks.md; re-run with:
//
//	go test -tags=integration ./internal/store/randomwalk/ -bench BenchmarkRandomWalkV2Queries -benchmem -count=6
//
// Numbers include the container round trip; compare only against baselines
// captured the same way.

import (
	"context"
	"testing"
)

func benchRepo(b *testing.B) *Repo {
	b.Helper()
	if errSetupSkip != nil {
		b.Skipf("skipping: %v", errSetupSkip)
	}
	if sharedRepo == nil {
		b.Fatal("store harness not initialized (TestMain did not run?)")
	}
	return sharedRepo
}

func BenchmarkRandomWalkV2Queries(b *testing.B) {
	r := benchRepo(b)
	ctx := context.Background()

	// The two structurally heaviest reads: the six-branch per-token event
	// merge and the outcome-joined offer ledger; plus the price-ranked
	// order book that backs the marketplace UI.
	b.Run("token_events", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			records, _, err := r.TokenEventsPage(ctx, 10, nil, 50)
			if err != nil || len(records) == 0 {
				b.Fatalf("token events: %v (%d rows)", err, len(records))
			}
		}
	})

	b.Run("offer_history", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			records, _, err := r.OfferHistoryPage(ctx, nil, 50)
			if err != nil || len(records) == 0 {
				b.Fatalf("offer history: %v (%d rows)", err, len(records))
			}
		}
	})

	b.Run("offers_price_asc", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			records, _, err := r.ActiveOffersPage(ctx, OfferSortPriceAsc, nil, 50)
			if err != nil || len(records) == 0 {
				b.Fatalf("active offers: %v (%d rows)", err, len(records))
			}
		}
	})

	b.Run("statistics", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			record, err := r.StatisticsV2(ctx)
			if err != nil || record.MintedCount == 0 {
				b.Fatalf("statistics: %v (%+v)", err, record)
			}
		}
	})
}
