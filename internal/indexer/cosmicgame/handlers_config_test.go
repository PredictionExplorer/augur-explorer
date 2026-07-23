// Unit tests (no Docker) for the handler-set constructor's dependency
// validation and the FilterLogs contract enumeration cmd/cg-etl subscribes
// with.
package cosmicgame

import (
	"context"
	"math/big"
	"strings"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/ethcall"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// nopCaller satisfies bind.ContractCaller without a chain.
type nopCaller struct{}

func (nopCaller) CodeAt(ctx context.Context, contract ethcommon.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

func (nopCaller) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

func TestNewValidatesConfig(t *testing.T) {
	repo := cgstore.NewRepo(store.NewFromPool(nil))
	st := store.NewFromPool(nil)

	cases := []struct {
		name string
		cfg  Config
		want string
	}{
		{"missing repo", Config{Store: st, Caller: nopCaller{}}, "Config.Repo"},
		{"missing store", Config{Repo: repo, Caller: nopCaller{}}, "Config.Store"},
		{"missing caller", Config{Repo: repo, Store: st}, "Config.Caller"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := New(tc.cfg); err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Errorf("New(%s) = %v, want error mentioning %s", tc.name, err, tc.want)
			}
		})
	}

	// A complete config with a nil logger builds the full registry.
	h, err := New(Config{Repo: repo, Store: st, Caller: nopCaller{}})
	if err != nil {
		t.Fatalf("New with full config: %v", err)
	}
	if h.Registry() == nil {
		t.Fatal("New returned a handler set without a registry")
	}
}

// TestNewBoundsContractCaller pins the D22 wrap: the handlers' contract
// reads (donation info, donated-NFT tokenURI) run inside the per-block
// ingestion transaction, so New must bound every call — a raw unbounded
// caller can never reach h.caller.
func TestNewBoundsContractCaller(t *testing.T) {
	h, err := New(Config{
		Repo:   cgstore.NewRepo(store.NewFromPool(nil)),
		Store:  store.NewFromPool(nil),
		Caller: nopCaller{},
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if _, ok := h.caller.(ethcall.BoundedCaller); !ok {
		t.Fatalf("h.caller is %T, want ethcall.BoundedCaller (in-transaction reads must be time-bounded)", h.caller)
	}
}

func TestContractsAllListsEveryWatchedContract(t *testing.T) {
	c := Contracts{
		Game:            ethcommon.HexToAddress("0x01"),
		Signature:       ethcommon.HexToAddress("0x02"),
		Token:           ethcommon.HexToAddress("0x03"),
		Dao:             ethcommon.HexToAddress("0x04"),
		CharityWallet:   ethcommon.HexToAddress("0x05"),
		PrizesWallet:    ethcommon.HexToAddress("0x06"),
		StakingCST:      ethcommon.HexToAddress("0x07"),
		StakingRWalk:    ethcommon.HexToAddress("0x08"),
		MarketingWallet: ethcommon.HexToAddress("0x09"),
		Implementation:  ethcommon.HexToAddress("0x0a"),
	}
	all := c.All()
	if len(all) != 10 {
		t.Fatalf("All() returned %d contracts, want 10", len(all))
	}
	seen := make(map[ethcommon.Address]bool, len(all))
	for _, a := range all {
		if a == (ethcommon.Address{}) {
			t.Errorf("All() contains a zero address: a field was forgotten")
		}
		if seen[a] {
			t.Errorf("All() lists %s twice", a)
		}
		seen[a] = true
	}
}
