//go:build integration

package cosmicgame

import (
	"strconv"
	"testing"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

// itoa keeps golden names short in table-driven cases.
func itoa(n int64) string { return strconv.FormatInt(n, 10) }

func TestGetCosmicGameContractAddrs(t *testing.T) {
	sw := store(t)
	golden(t, "cosmic_game_contract_addrs", func() any {
		return sw.Get_cosmic_game_contract_addrs()
	})
}

// TestProcessingStatusRoundTrip covers Get_cosmic_game_processing_status
// (which lazily inserts the default row) and the update path, restoring the
// original watermark afterwards.
func TestProcessingStatusRoundTrip(t *testing.T) {
	sw := store(t)

	initial := sw.Get_cosmic_game_processing_status()
	t.Cleanup(func() { sw.Update_cosmic_game_process_status(&initial) })

	want := p.CosmicGameProcStatus{LastEvtIdProcessed: 5098, LastBlockNum: 142}
	sw.Update_cosmic_game_process_status(&want)

	got := sw.Get_cosmic_game_processing_status()
	if got != want {
		t.Fatalf("processing status round trip: got %+v, want %+v", got, want)
	}
}
