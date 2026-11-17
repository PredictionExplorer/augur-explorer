//go:build integration

package chaoszero

import (
	"context"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	czmodel "github.com/PredictionExplorer/augur-explorer/internal/model/chaoszero"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Synthetic addresses for the write-error matrix (created on demand by the
// warm-up lookups; distinct from the fixture addresses so cache state is
// fully controlled by each subtest).
var errAddrs = []string{
	"0x00000000000000000000000000000000000000A1",
	"0x00000000000000000000000000000000000000A2",
	"0x00000000000000000000000000000000000000A3",
	"0x00000000000000000000000000000000000000A4",
}

// TestWriteErrorsSurfaceFromEveryAddressLookup drives every insert method
// through each of its address-resolution failure points: a fresh store's
// address cache is pre-warmed for the first N addresses (which then resolve
// without touching the database), and the cancelled context fails the first
// cold lookup. Every failure must surface as a wrapped error naming the
// target table.
func TestWriteErrorsSurfaceFromEveryAddressLookup(t *testing.T) {
	shared := repo(t) // harness guard
	_ = shared
	ctx := context.Background()
	cancelled, cancel := context.WithCancel(ctx)
	cancel()

	base := func() czmodel.EventBase {
		return czmodel.EventBase{
			EvtId: 999999, BlockNum: 1, TxId: 1,
			Contract: errAddrs[0], TimeStamp: 1767225600,
		}
	}
	amounts8 := []string{"1", "0", "0", "0", "0", "0", "0", "0"}

	cases := []struct {
		name string
		// addrCount is how many distinct addresses the method resolves;
		// one subtest per failure point (0..addrCount-1 warmed lookups).
		addrCount int
		wantOp    string
		call      func(r *Repo) error
	}{
		{"round_initialized", 1, "cz_round_initialized", func(r *Repo) error {
			return r.InsertRoundInitialized(cancelled, &czmodel.RoundInitialized{EventBase: base()})
		}},
		{"round_funded", 2, "cz_round_funded", func(r *Repo) error {
			return r.InsertRoundFunded(cancelled, &czmodel.RoundFunded{
				EventBase: base(), Caller: errAddrs[1],
				Donated: "1", TransfersUsed: "0", CarryUsed: "0", Allocations: amounts8,
			})
		}},
		{"trade", 2, "cz_trade", func(r *Repo) error {
			return r.InsertTrade(cancelled, &czmodel.Trade{
				EventBase: base(), User: errAddrs[1], Cash: "1", Units: "1",
			})
		}},
		{"position_split", 2, "cz_position_split", func(r *Repo) error {
			return r.InsertPositionSplit(cancelled, &czmodel.PositionSplit{
				EventBase: base(), User: errAddrs[1], Amount: "1",
			})
		}},
		{"positions_merged", 2, "cz_positions_merged", func(r *Repo) error {
			return r.InsertPositionsMerged(cancelled, &czmodel.PositionsMerged{
				EventBase: base(), User: errAddrs[1], Amounts: amounts8, CollateralOut: "1",
			})
		}},
		{"round_settled", 1, "cz_round_settled", func(r *Repo) error {
			return r.InsertRoundSettled(cancelled, &czmodel.RoundSettled{
				EventBase: base(), Answers: make([]bool, 8), Escrowed: amounts8, Carried: "0",
			})
		}},
		{"claim", 2, "cz_claim", func(r *Repo) error {
			return r.InsertClaim(cancelled, &czmodel.Claim{
				EventBase: base(), User: errAddrs[1], Amounts: amounts8, Total: "1",
			})
		}},
		{"answers_observed", 1, "cz_answers_observed", func(r *Repo) error {
			return r.InsertAnswersObserved(cancelled, &czmodel.AnswersObserved{EventBase: base()})
		}},
		{"cycle_escrowed", 1, "cz_cycle_escrowed", func(r *Repo) error {
			return r.InsertCycleEscrowed(cancelled, &czmodel.CycleEscrowed{
				EventBase: base(), Amounts: amounts8,
			})
		}},
		{"paid", 2, "cz_paid", func(r *Repo) error {
			return r.InsertPaid(cancelled, &czmodel.Paid{
				EventBase: base(), Recipient: errAddrs[1], Amount: "1",
			})
		}},
		{"position_transfer", 4, "cz_position_transfer", func(r *Repo) error {
			return r.InsertPositionTransfer(cancelled, &czmodel.PositionTransfer{
				EventBase: base(), Operator: errAddrs[1], From: errAddrs[2], To: errAddrs[3],
				Items: []czmodel.PositionTransferItem{{PositionId: "1", Amount: "1"}},
			})
		}},
		{"approval_for_all", 3, "cz_approval_for_all", func(r *Repo) error {
			return r.InsertApprovalForAll(cancelled, &czmodel.ApprovalForAll{
				EventBase: base(), Owner: errAddrs[1], Operator: errAddrs[2],
			})
		}},
	}

	for _, tc := range cases {
		for warm := 0; warm < tc.addrCount; warm++ {
			t.Run(tc.name, func(t *testing.T) {
				// A fresh store means a fresh address cache.
				st := store.NewFromPool(sharedStore.Pool())
				for i := 0; i < warm; i++ {
					if _, err := st.LookupOrCreateAddress(ctx, errAddrs[i], 1, 1); err != nil {
						t.Fatalf("warming address %d: %v", i, err)
					}
				}
				err := tc.call(NewRepo(st))
				if err == nil || !strings.Contains(err.Error(), tc.wantOp) {
					t.Fatalf("warm=%d err = %v, want substring %q", warm, err, tc.wantOp)
				}
			})
		}
	}
}

// TestProcessingStatusRecreateFailureSurfaces pins the lost-row recovery
// error path: when the singleton row is gone and the recreating INSERT
// cannot run (read-only transaction), the error names the recovery step.
func TestProcessingStatusRecreateFailureSurfaces(t *testing.T) {
	repo(t) // harness guard
	ctx := context.Background()

	cfg, err := pgxpool.ParseConfig(sharedConnString)
	if err != nil {
		t.Fatal(err)
	}
	cfg.ConnConfig.RuntimeParams["default_transaction_read_only"] = "on"
	roPool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(roPool.Close)

	if _, err := sharedStore.Pool().Exec(ctx, "DELETE FROM cz_proc_status"); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := sharedStore.Pool().Exec(context.Background(),
			"INSERT INTO cz_proc_status(last_evt_id, last_block) VALUES (0, 0)"); err != nil {
			t.Errorf("restoring cz_proc_status: %v", err)
		}
	})

	_, err = NewRepo(store.NewFromPool(roPool)).ProcessingStatus(ctx)
	if err == nil || !strings.Contains(err.Error(), "insert default row") {
		t.Fatalf("ProcessingStatus on read-only pool = %v, want the recreate failure", err)
	}
}

// TestRoundDetailPropagatesSectionFailures breaks the per-section tables the
// summary query does not touch and asserts RoundDetail surfaces the error
// instead of answering a partial round.
func TestRoundDetailPropagatesSectionFailures(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	for _, table := range []string{"cz_answers_observed", "cz_paid"} {
		t.Run(table, func(t *testing.T) {
			if _, err := sharedStore.Pool().Exec(ctx,
				"ALTER TABLE "+table+" RENAME TO "+table+"_broken"); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if _, err := sharedStore.Pool().Exec(context.Background(),
					"ALTER TABLE "+table+"_broken RENAME TO "+table); err != nil {
					t.Errorf("restoring %s: %v", table, err)
				}
			})
			if _, err := r.RoundDetail(ctx, 0); err == nil {
				t.Fatalf("RoundDetail with %s broken = nil error", table)
			}
		})
	}
}

// TestQuestionOddsIgnoresOutOfRangeQuestions pins the defensive skip: a
// trade row whose question is outside 0..7 (impossible from the decoder,
// possible from manual writes) must not panic the odds array fill.
func TestQuestionOddsIgnoresOutOfRangeQuestions(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	if _, err := sharedStore.Pool().Exec(ctx, `INSERT INTO cz_trade(
			evtlog_id,block_num,tx_id,time_stamp,contract_aid,
			round_id,user_aid,question,yes,buy,cash,units
		) VALUES (NULL,166,1117,TO_TIMESTAMP(1767232200),31,0,21,9,true,true,1,1)`); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if _, err := sharedStore.Pool().Exec(context.Background(),
			"DELETE FROM cz_trade WHERE question = 9"); err != nil {
			t.Errorf("removing the out-of-range trade: %v", err)
		}
	})

	detail, err := r.RoundDetail(ctx, 0)
	if err != nil {
		t.Fatalf("RoundDetail: %v", err)
	}
	if len(detail.Odds) != 8 {
		t.Fatalf("odds rows = %d, want exactly 8", len(detail.Odds))
	}
	for q, odds := range detail.Odds {
		if int(odds.Question) != q {
			t.Errorf("odds[%d].Question = %d, want %d", q, odds.Question, q)
		}
	}
}
