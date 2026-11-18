//go:build integration

// Gap tests for champion_durations.go: the legacy championDurations(uint256)
// fallback getter, the overflow guard, and the recovery pass's failure
// propagation.
package cosmicgame

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// registerLegacyChampionDurations stubs the game so that roundStats() is
// unknown (older v3.1 deployments) while the legacy championDurations(uint256)
// getter answers with the given pair.
func registerLegacyChampionDurations(t *testing.T, endurance, chrono *big.Int) {
	t.Helper()
	method := legacyChampionDurationsABI.Methods["championDurations"]
	testChain.RegisterCall(addr(fxGameAddr), func(input []byte) ([]byte, error) {
		if len(input) < 4 || string(input[:4]) != string(method.ID) {
			return nil, fmt.Errorf("pre-roundStats deployment: unknown selector %x", input)
		}
		return method.Outputs.Pack(endurance, chrono)
	})
}

func TestReadChampionDurationsLegacyGetter(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)

	registerLegacyChampionDurations(t, big.NewInt(31), big.NewInt(41))
	endurance, chrono, err := testHandlers.readChampionDurations(context.Background(), 5)
	if err != nil {
		t.Fatalf("legacy getter: %v", err)
	}
	if endurance != 31 || chrono != 41 {
		t.Fatalf("legacy durations = %d/%d, want 31/41", endurance, chrono)
	}
}

func TestReadChampionDurationsRejectsOverflow(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)

	huge := new(big.Int).Lsh(big.NewInt(1), 70) // > int64
	registerLegacyChampionDurations(t, huge, big.NewInt(1))
	if _, _, err := testHandlers.readChampionDurations(context.Background(), 5); err == nil ||
		!strings.Contains(err.Error(), "exceed int64") {
		t.Fatalf("expected overflow error, got %v", err)
	}
}

func TestReadChampionDurationsBothGettersMissing(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)

	// A V1-era contract: neither roundStats nor championDurations exists.
	testChain.RegisterCall(addr(fxGameAddr),
		testchain.MustContractStub(cgc.OwnableABI).Handler())
	if _, _, err := testHandlers.readChampionDurations(context.Background(), 5); err == nil {
		t.Fatal("expected error when both getters are missing")
	}
}

// seedRoundMissingDurations ingests a MainPrizeClaimed whose duration read
// fails, leaving the round in the recovery queue.
func seedRoundMissingDurations(t *testing.T, blockNum, roundNum int64) {
	t.Helper()
	testChain.Reorg(blockNum)
	testChain.RegisterCall(addr(fxGameAddr),
		testchain.MustContractStub(cgc.CosmicSignatureGameV3ABI).Handler())
	ingestTx(t, blockNum, addr(fxGameAddr), 0, []*types.Log{
		buildLog(t, gameV3ABI, "MainPrizeClaimed", addr(fxGameAddr),
			[]any{bigInt(roundNum), addr(fxAlice), bigInt(200)},
			[]any{eth(4), eth(12), bigInt(3), bigInt(600)}),
	})
}

func TestRecoverChampionDurationsPropagatesStoreFailure(t *testing.T) {
	resetDB(t)
	t.Cleanup(registerCallHandlers)
	const (
		blockNum = int64(5600)
		roundNum = int64(14)
	)
	seedRoundMissingDurations(t, blockNum, roundNum)

	registerCallHandlers()
	championDurationsMu.Lock()
	championDurations[roundNum] = [2]int64{703, 903}
	championDurationsMu.Unlock()

	// Hide the audit table: the recovery read succeeds, the write fails, and
	// the pass must report it (store failures abort, unlike read failures).
	ctx := context.Background()
	if _, err := testDB.Pool.Exec(ctx,
		"ALTER TABLE cg_live_state_updates RENAME TO cg_live_state_updates_hidden"); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, _ = testDB.Pool.Exec(context.Background(),
			"ALTER TABLE cg_live_state_updates_hidden RENAME TO cg_live_state_updates")
	})
	if err := testHandlers.RecoverChampionDurations(ctx, blockNum+1, 1767231000); err == nil {
		t.Fatal("expected store failure to propagate")
	}
}

func TestRecoverChampionDurationsQueryFailure(t *testing.T) {
	resetDB(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if err := testHandlers.RecoverChampionDurations(cancelled, 1, 1767231000); err == nil {
		t.Fatal("expected error from cancelled context")
	}
}

func TestBootstrapContractsQueryFailure(t *testing.T) {
	resetDB(t)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, _, err := BootstrapContracts(cancelled, cgRepo, dbStore); err == nil ||
		!strings.Contains(err.Error(), "reading contract addresses") {
		t.Fatalf("expected contract-address read failure, got %v", err)
	}
}
