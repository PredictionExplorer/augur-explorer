// Unit tests (no Docker) for the RandomWalk handler-set constructor and the
// FilterLogs contract enumeration cmd/rw-etl subscribes with.
package randomwalk

import (
	"strings"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

func TestNewValidatesConfig(t *testing.T) {
	if _, err := New(Config{}); err == nil || !strings.Contains(err.Error(), "Config.Repo") {
		t.Errorf("New without repo = %v, want Config.Repo error", err)
	}

	h, err := New(Config{Repo: rwstore.NewRepo(store.NewFromPool(nil))})
	if err != nil {
		t.Fatalf("New with repo and nil logger: %v", err)
	}
	if h.Registry() == nil {
		t.Fatal("New returned a handler set without a registry")
	}
}

func TestContractsAllListsBothWatchedContracts(t *testing.T) {
	c := Contracts{
		Market:     ethcommon.HexToAddress("0x12"),
		RandomWalk: ethcommon.HexToAddress("0x08"),
	}
	all := c.All()
	if len(all) != 2 {
		t.Fatalf("All() returned %d contracts, want 2", len(all))
	}
	if all[0] != c.RandomWalk || all[1] != c.Market {
		t.Errorf("All() = %v, want [RandomWalk, Market]", all)
	}
}
