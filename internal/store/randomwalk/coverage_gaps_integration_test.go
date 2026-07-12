//go:build integration

package randomwalk

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
)

// TestMintReportMonthlyAggregation seeds mints inside MintReport's hardcoded
// 2021-11..2022-12 reporting window (the shared fixture mints all fall after
// it, so the scan body was never exercised) and pins the month naming, the
// exact wei sums and the cumulative redeem amount (half of everything
// deposited so far, in report order).
func TestMintReportMonthlyAggregation(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	pl := pool(t)

	// Two in-window mints for carol: June 2022 (2 ETH) and August 2022
	// (4 ETH). The mint trigger creates the rw_token rows (the fixture flow
	// creates them via rw_transfer first, so this also covers the trigger's
	// insert-if-missing branch); the cleanup deletes both tables to restore
	// fixture state (the delete trigger reverses the aggregates).
	if _, err := pl.Exec(ctx,
		`INSERT INTO rw_mint_evt(evtlog_id, block_num, tx_id, time_stamp, contract_aid, token_id, owner_aid, seed, seed_num, price) VALUES
			(NULL, 999801, 1030, '2022-06-15T12:00:00Z', 8, 90, 23, 'bb00000000000000000000000000000000000000000000000000000000000090', 90, 2000000000000000000),
			(NULL, 999802, 1030, '2022-08-10T12:00:00Z', 8, 91, 23, 'bb00000000000000000000000000000000000000000000000000000000000091', 91, 4000000000000000000)`); err != nil {
		t.Fatalf("seeding in-window mints: %v", err)
	}
	t.Cleanup(func() {
		cleanupCtx := context.Background()
		if _, err := pl.Exec(cleanupCtx, "DELETE FROM rw_mint_evt WHERE block_num IN (999801, 999802)"); err != nil {
			t.Errorf("removing seeded mints: %v", err)
		}
		if _, err := pl.Exec(cleanupCtx, "DELETE FROM rw_token WHERE rwalk_aid=8 AND token_id IN (90, 91)"); err != nil {
			t.Errorf("removing trigger-created tokens: %v", err)
		}
	})

	recs, err := r.MintReport(ctx)
	if err != nil {
		t.Fatalf("MintReport: %v", err)
	}
	if len(recs) != 2 {
		t.Fatalf("got %d report rows, want 2 (months with mints only): %+v", len(recs), recs)
	}

	june := recs[0]
	if june.Year != 2022 || june.Month != 6 || june.MonthStr != "June 2022" {
		t.Errorf("first row month = %d-%d %q, want June 2022", june.Year, june.Month, june.MonthStr)
	}
	if june.TotalMinted != 1 || june.TotalWei != "2000000000000000000" || june.TotalEth != 2 {
		t.Errorf("June row = %+v", june)
	}
	if june.RedeemAmount != 1 { // half of 2 ETH deposited so far
		t.Errorf("June RedeemAmount = %v, want 1", june.RedeemAmount)
	}

	august := recs[1]
	if august.MonthStr != "August 2022" {
		t.Errorf("second row = %q, want August 2022", august.MonthStr)
	}
	if august.TotalMinted != 1 || august.TotalWei != "4000000000000000000" {
		t.Errorf("August row = %+v", august)
	}
	if august.RedeemAmount != 3 { // half of the cumulative 6 ETH
		t.Errorf("August RedeemAmount = %v, want 3", august.RedeemAmount)
	}
}

// TestWriteMethodsPropagateCancellation sweeps every Insert*/Update*
// (ctx, *Event) error method with a cancelled context and a zero-value
// event: the first resolution step must abort with an error — either the
// empty-address validation or the cancelled query, depending on the method's
// first database touch. (The legacy layer treated several of these failures
// as "record does not exist" and silently dropped events — the data-loss bug
// fixed in Phase 1.)
func TestWriteMethodsPropagateCancellation(t *testing.T) {
	r := repo(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	ctxType := reflect.TypeFor[context.Context]()
	errType := reflect.TypeFor[error]()

	rv := reflect.ValueOf(r)
	rt := rv.Type()
	swept := 0
	for i := 0; i < rt.NumMethod(); i++ {
		m := rt.Method(i)
		if !strings.HasPrefix(m.Name, "Insert") && !strings.HasPrefix(m.Name, "Update") {
			continue
		}
		mt := m.Type
		if mt.NumIn() != 3 || mt.NumOut() != 1 ||
			mt.In(1) != ctxType || mt.In(2).Kind() != reflect.Pointer ||
			mt.Out(0) != errType {
			continue
		}
		evt := reflect.New(mt.In(2).Elem())
		out := rv.Method(i).Call([]reflect.Value{reflect.ValueOf(cancelled), evt})
		err, _ := out[0].Interface().(error)
		if err == nil {
			t.Errorf("%s with cancelled ctx and zero event succeeded, want a resolution failure", m.Name)
		} else if !errors.Is(err, context.Canceled) && !strings.Contains(err.Error(), "empty address") {
			t.Errorf("%s failed outside the first resolution guard: %v", m.Name, err)
		}
		swept++
	}
	if swept < 8 {
		t.Errorf("swept only %d write methods; the reflection filter no longer matches the write layer", swept)
	}
}
