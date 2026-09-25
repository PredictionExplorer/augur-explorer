package cosmicgame

import (
	"context"
	"slices"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

func TestContractsAllIncludesEveryImplementationOnce(t *testing.T) {
	c := unitContracts()
	implA := ethcommon.Address{19: 0x10} // same as the registry implementation
	implB := ethcommon.Address{19: 0x20}
	c.Implementations = []ethcommon.Address{implA, implB, implB, {}}

	all := c.All()
	if n := countOf(all, implA); n != 1 {
		t.Errorf("registry implementation listed %d times, want 1", n)
	}
	if n := countOf(all, implB); n != 1 {
		t.Errorf("upgraded implementation listed %d times, want 1", n)
	}
	if slices.Contains(all, ethcommon.Address{}) {
		t.Error("zero address must not be watched")
	}
}

func TestNewImplementationJoinsInitializedSourcesAndFiresHook(t *testing.T) {
	h := newUnitHandlers(t)
	impl := ethcommon.Address{19: 0x77}
	if slices.Contains(h.initializedSources(), impl) {
		t.Fatal("unknown implementation admitted before Upgraded")
	}

	var got []ethcommon.Address
	var gotBlock int64
	h.SetOnNewImplementation(func(_ context.Context, a ethcommon.Address, block int64) {
		got = append(got, a)
		gotBlock = block
	})

	h.noteImplementation(context.Background(), impl, 4242)
	h.noteImplementation(context.Background(), impl, 4243) // repeat: no second callback
	h.noteImplementation(context.Background(), h.c.Implementation, 4244)

	if len(got) != 1 || got[0] != impl || gotBlock != 4242 {
		t.Fatalf("hook calls = %v (block %d), want one call for %s at 4242", got, gotBlock, impl.Hex())
	}
	if !slices.Contains(h.Implementations(), impl) {
		t.Error("implementation set did not grow")
	}
	// The registered Initialized handler re-evaluates its sources.
	for _, hd := range h.Registry().Handlers() {
		if hd.Name() == "Initialized" && !slices.Contains(hd.Sources(), impl) {
			t.Error("Initialized handler does not admit the new implementation")
		}
	}
}

func countOf(list []ethcommon.Address, a ethcommon.Address) int {
	n := 0
	for _, x := range list {
		if x == a {
			n++
		}
	}
	return n
}
