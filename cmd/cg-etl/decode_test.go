// Unit tests (no Docker) for the raw-data decode helpers of events whose
// signatures no generated ABI defines.
package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
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
			got, err := erc20_transfer_failed_amount(data)
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
		if _, err := erc20_transfer_failed_amount(make([]byte, n)); err == nil {
			t.Errorf("%d-byte data: want error, got nil", n)
		}
	}
}

func TestAdminUint256FromLogData(t *testing.T) {
	// Single-word body: the value is the last (and only) 32-byte word.
	word := make([]byte, 32)
	word[31] = 42
	got, err := admin_uint256_from_log_data(word)
	if err != nil {
		t.Fatalf("admin_uint256_from_log_data: %v", err)
	}
	if got.Int64() != 42 {
		t.Fatalf("value = %v, want 42", got)
	}
	if _, err := admin_uint256_from_log_data(make([]byte, 31)); err == nil {
		t.Error("short data: want error, got nil")
	}
}
