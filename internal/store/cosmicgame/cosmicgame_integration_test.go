//go:build integration

package cosmicgame

import (
	"context"
	"strconv"
	"testing"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

// itoa keeps golden names short in table-driven cases.
func itoa(n int64) string { return strconv.FormatInt(n, 10) }

func TestContractAddrs(t *testing.T) {
	r := repo(t)
	golden(t, "cosmic_game_contract_addrs", func() any {
		addrs, err := r.ContractAddrs(context.Background())
		if err != nil {
			t.Fatalf("ContractAddrs: %v", err)
		}
		return addrs
	})
}

// TestProcessingStatusRoundTrip covers ProcessingStatus (which lazily inserts
// the default row) and the update path, restoring the original watermark
// afterwards.
func TestProcessingStatusRoundTrip(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	initial, err := r.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("ProcessingStatus: %v", err)
	}
	t.Cleanup(func() {
		if err := r.UpdateProcessingStatus(ctx, &initial); err != nil {
			t.Errorf("restoring processing status: %v", err)
		}
	})

	want := cgmodel.CosmicGameProcStatus{LastEvtIdProcessed: 5098, LastBlockNum: 142}
	if err := r.UpdateProcessingStatus(ctx, &want); err != nil {
		t.Fatalf("UpdateProcessingStatus: %v", err)
	}

	got, err := r.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("ProcessingStatus after update: %v", err)
	}
	if got != want {
		t.Fatalf("processing status round trip: got %+v, want %+v", got, want)
	}
}
