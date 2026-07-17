//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

// TestInsertMethodsPropagateAddressLookupFailure sweeps every
// Insert*(ctx, *Event) error method with a cancelled context and a
// zero-value event: the address resolution of the event's first address
// field must abort the insert and surface an error (either the empty-address
// validation or the cancelled query, depending on the method's first DB
// touch). The legacy layer called os.Exit(1) on these paths; the reflection
// sweep guarantees no newly added insert regresses to swallowing resolution
// failures.
func TestInsertMethodsPropagateAddressLookupFailure(t *testing.T) {
	r := repo(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	ctxType := reflect.TypeFor[context.Context]()
	errType := reflect.TypeFor[error]()

	rv := reflect.ValueOf(r)
	rt := rv.Type()
	swept := 0
	for i := range rt.NumMethod() {
		m := rt.Method(i)
		if !strings.HasPrefix(m.Name, "Insert") {
			continue
		}
		mt := m.Type
		// Sweep the uniform shape Insert*(ctx, *Event) error; the few
		// bespoke signatures (InsertDonationJSON, admin corrections) have
		// dedicated tests.
		if mt.NumIn() != 3 || mt.NumOut() != 1 ||
			mt.In(1) != ctxType || mt.In(2).Kind() != reflect.Pointer ||
			mt.Out(0) != errType {
			continue
		}
		evt := reflect.New(mt.In(2).Elem())
		out := rv.Method(i).Call([]reflect.Value{reflect.ValueOf(cancelled), evt})
		err, _ := out[0].Interface().(error)
		if err == nil {
			t.Errorf("%s with cancelled ctx and zero event succeeded, want address-resolution failure", m.Name)
		} else if !errors.Is(err, context.Canceled) && !strings.Contains(err.Error(), "empty address") {
			t.Errorf("%s failed outside the address-resolution guard: %v", m.Name, err)
		}
		swept++
	}
	// The write layer had 73 inserts when this sweep landed; a shrinking
	// count means methods changed shape without the sweep following.
	if swept < 70 {
		t.Errorf("swept only %d Insert methods; the reflection filter no longer matches the write layer", swept)
	}
}

// TestRoundStatisticsActivationTimeFallback pins the cg_adm_acttime fallback
// of CosmicGameRoundStatistics: when the round-stats row is missing (a round
// with no activity yet) or carries a NULL activation_time, the activation
// time is derived from the latest admin event applying to the next
// unclaimed round — and rounds beyond it come back with zero.
func TestRoundStatisticsActivationTimeFallback(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	pl := r.pool()

	// Fixture facts: rounds 0-2 are claimed, so the "next" round is 3, and
	// the latest cg_adm_acttime row (evtlog 5075) sets 1767228300.
	const wantActivation = int64(1767228300)

	// Null out round 3's trigger-computed activation time; restore after.
	var hadValue bool
	var oldEpoch int64
	err := pl.QueryRow(ctx,
		"SELECT activation_time IS NOT NULL, COALESCE(EXTRACT(EPOCH FROM activation_time)::BIGINT, 0) "+
			"FROM cg_round_stats WHERE round_num=3").Scan(&hadValue, &oldEpoch)
	if err != nil {
		t.Fatalf("reading round 3 activation time: %v", err)
	}
	if _, err := pl.Exec(ctx, "UPDATE cg_round_stats SET activation_time=NULL WHERE round_num=3"); err != nil {
		t.Fatalf("clearing activation time: %v", err)
	}
	t.Cleanup(func() {
		if hadValue {
			if _, err := pl.Exec(context.Background(),
				"UPDATE cg_round_stats SET activation_time=TO_TIMESTAMP($1) WHERE round_num=3", oldEpoch); err != nil {
				t.Errorf("restoring activation time: %v", err)
			}
		}
	})

	stats, err := r.CosmicGameRoundStatistics(ctx, 3)
	if err != nil {
		t.Fatalf("CosmicGameRoundStatistics(3): %v", err)
	}
	if stats.ActivationTime != wantActivation {
		t.Errorf("round 3 fallback ActivationTime = %d, want %d", stats.ActivationTime, wantActivation)
	}

	// Round 4 has no stats row and is not the next round: both fallback
	// layers miss and the zero-value shape is returned.
	stats, err = r.CosmicGameRoundStatistics(ctx, 4)
	if err != nil {
		t.Fatalf("CosmicGameRoundStatistics(4): %v", err)
	}
	if stats.RoundNum != 4 || stats.ActivationTime != 0 || stats.TotalBids != 0 {
		t.Errorf("round 4 shape = %+v, want zero row with RoundNum=4", stats)
	}
}

// TestUnclaimedERC20Items seeds an unclaimed ERC-20 donation (the fixture's
// only ERC-20 donation is claimed) and pins both consumers of the unclaimed
// state: the legacy ClaimsByRound inline items and the v2 keyset page's
// erc20 segment. The donation-stats trigger creates and later removes the
// aggregate row, so fixture state is restored by deleting the donation.
func TestUnclaimedERC20Items(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	pl := r.pool()

	// Round 0, token address id 27 (an address the fixtures registered but
	// never used as an ERC-20 donation token, so a fresh stats row appears
	// with claimed = FALSE).
	const donationBlock = 999901
	if _, err := pl.Exec(ctx,
		"INSERT INTO cg_erc20_donation(evtlog_id, block_num, tx_id, time_stamp, contract_aid, round_num, donor_aid, token_aid, amount, bid_id) "+
			"VALUES (NULL, $1, 1008, TO_TIMESTAMP(1767226301), 7, 0, 21, 27, 123450000000000000000, 2001)",
		donationBlock); err != nil {
		t.Fatalf("seeding unclaimed ERC-20 donation: %v", err)
	}
	t.Cleanup(func() {
		if _, err := pl.Exec(context.Background(),
			"DELETE FROM cg_erc20_donation WHERE block_num=$1", donationBlock); err != nil {
			t.Errorf("removing seeded donation: %v", err)
		}
	})

	summaries, err := r.ClaimsByRound(ctx)
	if err != nil {
		t.Fatalf("ClaimsByRound: %v", err)
	}
	var item *cgmodel.CGClaimUnclaimedItem
	for i := range summaries {
		if summaries[i].RoundNum != 0 {
			continue
		}
		for j := range summaries[i].UnclaimedItems {
			if summaries[i].UnclaimedItems[j].AssetType == "ERC20" {
				item = &summaries[i].UnclaimedItems[j]
			}
		}
	}
	if item == nil {
		t.Fatal("round 0 has no unclaimed ERC20 item after seeding one")
	}
	if item.AmountEth != 123.45 || item.TokenId != -1 || item.RecipientAddr == "" || item.TokenAddr == "" {
		t.Errorf("unclaimed ERC20 item = %+v", *item)
	}

	// The v2 unclaimed page must expose the same donation in its erc20
	// segment with exact base units and no ETH/token-id fields.
	records, hasMore, err := r.UnclaimedItemsPage(ctx, 0, nil, 50)
	if err != nil {
		t.Fatalf("UnclaimedItemsPage: %v", err)
	}
	if hasMore {
		t.Error("unexpected extra unclaimed pages")
	}
	var v2rec *UnclaimedItemRecord
	for i := range records {
		if records[i].AssetType == "erc20" {
			v2rec = &records[i]
		}
	}
	if v2rec == nil {
		t.Fatal("v2 unclaimed page has no erc20 record")
	}
	if v2rec.AmountBaseUnits == nil || *v2rec.AmountBaseUnits != "123450000000000000000" {
		t.Errorf("AmountBaseUnits = %v, want exact base units", v2rec.AmountBaseUnits)
	}
	if v2rec.TokenAddr == nil || *v2rec.TokenAddr == "" {
		t.Errorf("TokenAddr = %v, want the donated token address", v2rec.TokenAddr)
	}
	if v2rec.EthAmountWei != nil || v2rec.TokenID != nil {
		t.Errorf("erc20 record leaked eth/token fields: %+v", *v2rec)
	}
	if v2rec.RoundNum != 0 || v2rec.Segment != 2 {
		t.Errorf("erc20 record scope = round %d segment %d", v2rec.RoundNum, v2rec.Segment)
	}
}
