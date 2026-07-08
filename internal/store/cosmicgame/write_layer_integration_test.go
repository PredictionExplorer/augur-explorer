//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// TestDeleteMethodsValidSQL sweeps every Delete* method on the Repo via
// reflection and executes it against the real schema. A DELETE with id -1
// removes nothing, but PostgreSQL still parses the statement and resolves
// the table, so a typo'd table name fails loudly here. (Two such typos
// shipped in the legacy layer and killed the ETL on every reorg of those
// events — cg_fund_transfer_err vs cg_fund_transf_err.)
func TestDeleteMethodsValidSQL(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	repoType := reflect.TypeOf(r)
	ctxType := reflect.TypeOf((*context.Context)(nil)).Elem()
	errType := reflect.TypeOf((*error)(nil)).Elem()

	swept := 0
	for i := 0; i < repoType.NumMethod(); i++ {
		m := repoType.Method(i)
		if !strings.HasPrefix(m.Name, "Delete") {
			continue
		}
		mt := m.Func.Type()
		// (receiver, ctx, int64) error — the by-evtlog-id delete shape
		// (DeleteBannedBid shares it; the sweep covers it too).
		if mt.NumIn() != 3 || mt.NumOut() != 1 ||
			mt.In(1) != ctxType || mt.In(2) != reflect.TypeOf(int64(0)) || mt.Out(0) != errType {
			continue
		}
		swept++
		out := m.Func.Call([]reflect.Value{
			reflect.ValueOf(r), reflect.ValueOf(ctx), reflect.ValueOf(int64(-1)),
		})
		if errVal := out[0].Interface(); errVal != nil {
			t.Errorf("%s(-1) = %v, want nil (invalid SQL or wrong table name)", m.Name, errVal)
		}
	}
	// The event-delete surface is 72 methods (+ DeleteBannedBid); a shrinking
	// count means methods were renamed out of the swept shape unnoticed.
	if swept < 73 {
		t.Errorf("swept %d Delete* methods, want at least 73 — reflection filter no longer matches", swept)
	}
}

// TestLookupOrCreateAddress exercises the Store-level address resolution
// (create, cached and uncached lookups, ErrNotFound, empty input) against
// the real database. The rows created here are removed again so the shared
// fixture dataset stays untouched for the golden tests.
func TestLookupOrCreateAddress(t *testing.T) {
	repo(t) // skip when the integration environment is unavailable
	ctx := context.Background()

	st, err := spareStore(ctx)
	if err != nil {
		t.Fatalf("connecting spare store: %v", err)
	}
	defer st.Close()

	const newAddr = "0xAaAa000000000000000000000000000000009999"
	t.Cleanup(func() {
		_, _ = st.Pool().Exec(context.Background(), "DELETE FROM address WHERE addr=$1", newAddr)
	})

	created, err := st.LookupOrCreateAddress(ctx, newAddr, 123, 45)
	if err != nil {
		t.Fatalf("creating address: %v", err)
	}
	if created <= 0 {
		t.Fatalf("created address id = %d, want > 0", created)
	}

	// Second call hits the per-Store cache.
	cached, err := st.LookupOrCreateAddress(ctx, newAddr, 999, 999)
	if err != nil {
		t.Fatalf("cached lookup: %v", err)
	}
	if cached != created {
		t.Fatalf("cached id = %d, want %d", cached, created)
	}

	// A fresh Store (empty cache) resolves the same id from the database,
	// and the recorded first-seen block/tx are the creator's values.
	fresh, err := spareStore(ctx)
	if err != nil {
		t.Fatalf("connecting fresh store: %v", err)
	}
	defer fresh.Close()
	uncached, err := fresh.LookupAddressID(ctx, newAddr)
	if err != nil {
		t.Fatalf("uncached LookupAddressID: %v", err)
	}
	if uncached != created {
		t.Fatalf("uncached id = %d, want %d", uncached, created)
	}
	var blockNum, txID int64
	if err := fresh.Pool().QueryRow(ctx, "SELECT block_num, tx_id FROM address WHERE addr=$1", newAddr).Scan(&blockNum, &txID); err != nil {
		t.Fatalf("reading created row: %v", err)
	}
	if blockNum != 123 || txID != 45 {
		t.Errorf("created row block/tx = %d/%d, want 123/45 (second call must not overwrite)", blockNum, txID)
	}

	if _, err := fresh.LookupAddressID(ctx, "0xBbBb000000000000000000000000000000008888"); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("missing address lookup = %v, want ErrNotFound in chain", err)
	}
	if _, err := fresh.LookupOrCreateAddress(ctx, "", 1, 2); err == nil {
		t.Error("empty address: want error, got nil")
	}

	// After a cache reset the same Store re-reads from the database.
	st.ResetAddressCache()
	again, err := st.LookupOrCreateAddress(ctx, newAddr, 0, 0)
	if err != nil {
		t.Fatalf("post-reset lookup: %v", err)
	}
	if again != created {
		t.Fatalf("post-reset id = %d, want %d", again, created)
	}
}
