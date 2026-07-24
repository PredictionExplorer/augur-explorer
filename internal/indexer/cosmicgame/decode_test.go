// Unit tests (no Docker) for the raw-data decode helpers of events whose
// signatures no generated ABI defines.
package cosmicgame

import (
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// packERC20TransferFailedData builds the canonical ABI encoding of the
// non-indexed (string errStr, uint256 amount) body.
func packERC20TransferFailedData(t *testing.T, errStr string, amount *big.Int) []byte {
	t.Helper()
	stringT, err := abi.NewType("string", "", nil)
	if err != nil {
		t.Fatalf("string abi type: %v", err)
	}
	uint256T, err := abi.NewType("uint256", "", nil)
	if err != nil {
		t.Fatalf("uint256 abi type: %v", err)
	}
	data, err := abi.Arguments{{Type: stringT}, {Type: uint256T}}.Pack(errStr, amount)
	if err != nil {
		t.Fatalf("packing: %v", err)
	}
	return data
}

func TestErc20TransferFailedAmount(t *testing.T) {
	cases := []struct {
		name   string
		errStr string
		amount *big.Int
	}{
		{"small amount", "CST transfer failed", big.NewInt(12345)},
		{"zero amount", "err", big.NewInt(0)},
		{"wei-scale amount", "", new(big.Int).Mul(big.NewInt(7), big.NewInt(1e18))},
		{"long message", string(make([]byte, 100)), big.NewInt(1)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := packERC20TransferFailedData(t, c.errStr, c.amount)
			got, err := erc20TransferFailedAmount(data)
			if err != nil {
				t.Fatalf("erc20_transfer_failed_amount: %v", err)
			}
			if got.Cmp(c.amount) != 0 {
				t.Fatalf("amount = %s, want %s", got, c.amount)
			}
		})
	}
}

func TestErc20TransferFailedAmountShortData(t *testing.T) {
	for _, n := range []int{0, 31, 32, 63} {
		if _, err := erc20TransferFailedAmount(make([]byte, n)); err == nil {
			t.Errorf("%d-byte data: want error, got nil", n)
		}
	}
}

// TestDecodeInitializedVersionBounds pins the totality guard on the
// OpenZeppelin Initialized(uint64 version) decode: the type(uint64).max
// sentinel emitted by _disableInitializers() maps to -1 (matching legacy
// data), any other version beyond int64 fails the batch loudly instead of
// wrapping negative in the database, and the maximum representable version
// decodes exactly.
func TestDecodeInitializedVersionBounds(t *testing.T) {
	h := newUnitHandlers(t)
	elog := &store.EthereumEventLog{EvtID: 1, BlockNum: 2, TxID: 3}

	// One ABI word carrying version = 2^63 — one past math.MaxInt64.
	word := make([]byte, 32)
	word[24] = 0x80
	lg := &types.Log{Data: word}
	if _, err := h.decodeInitialized(lg, elog); err == nil || !strings.Contains(err.Error(), "overflows int64") {
		t.Fatalf("decodeInitialized(2^63) error = %v, want version-overflow rejection", err)
	}

	// type(uint64).max is the _disableInitializers() sentinel, stored as -1.
	for i := 24; i < 32; i++ {
		word[i] = 0xff
	}
	evt, err := h.decodeInitialized(lg, elog)
	if err != nil {
		t.Fatalf("decodeInitialized(MaxUint64): %v", err)
	}
	if evt.Version != -1 {
		t.Fatalf("Version = %d, want -1 sentinel", evt.Version)
	}

	// math.MaxInt64 itself is the last exactly-representable value.
	word[24] = 0x7f
	evt, err = h.decodeInitialized(lg, elog)
	if err != nil {
		t.Fatalf("decodeInitialized(MaxInt64): %v", err)
	}
	if evt.Version != math.MaxInt64 {
		t.Fatalf("Version = %d, want math.MaxInt64", evt.Version)
	}
}

func TestAdminUint256FromLogData(t *testing.T) {
	// Single-word body: the value is the last (and only) 32-byte word.
	word := make([]byte, 32)
	word[31] = 42
	got, err := adminUint256FromLogData(word)
	if err != nil {
		t.Fatalf("admin_uint256_from_log_data: %v", err)
	}
	if got.Int64() != 42 {
		t.Fatalf("value = %v, want 42", got)
	}
	if _, err := adminUint256FromLogData(make([]byte, 31)); err == nil {
		t.Error("short data: want error, got nil")
	}
}
