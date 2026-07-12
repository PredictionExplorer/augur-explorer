package ethtx

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestWeiToEthCompact(t *testing.T) {
	if got := WeiToEthCompact(nil); got != "0.0000" {
		t.Errorf("WeiToEthCompact(nil) = %q", got)
	}
	wei, _ := new(big.Int).SetString("1234567890000000000", 10)
	if got := WeiToEthCompact(wei); got != "1.234568" {
		t.Errorf("WeiToEthCompact = %q, want 1.234568", got)
	}
}

func TestFormatTokenAmount(t *testing.T) {
	cases := []struct {
		name     string
		amount   *big.Int
		decimals uint8
		symbol   string
		want     string
	}{
		{"nil amount", nil, 18, "CST", "0 CST"},
		{"zero decimals", big.NewInt(42), 0, "PTS", "42 PTS"},
		{"six decimals", big.NewInt(1_500_000), 6, "USDC", "1.500000 USDC"},
		{"eighteen decimals", big.NewInt(1_000_000_000_000_000_000), 18, "CST", "1.000000000000000000 CST"},
		{"sub-unit", big.NewInt(5), 6, "USDC", "0.000005 USDC"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FormatTokenAmount(tc.amount, tc.decimals, tc.symbol); got != tc.want {
				t.Errorf("FormatTokenAmount = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestFmtDuration(t *testing.T) {
	cases := []struct {
		secs int64
		want string
	}{
		{-3, "-3 sec (negative)"},
		{0, "0 sec"},
		{45, "45 sec"},
		{90, "1 min 30 sec"},
		{3600, "1h 0m 0s"},
		{3725, "1h 2m 5s"},
	}
	for _, tc := range cases {
		if got := FmtDuration(tc.secs); got != tc.want {
			t.Errorf("FmtDuration(%d) = %q, want %q", tc.secs, got, tc.want)
		}
	}
}

func TestConvertToPercentage(t *testing.T) {
	if got := ConvertToPercentage(nil); got != 0 {
		t.Errorf("ConvertToPercentage(nil) = %v", got)
	}
	if got := ConvertToPercentage(big.NewInt(0)); got != 0 {
		t.Errorf("ConvertToPercentage(0) = %v", got)
	}
	if got := ConvertToPercentage(big.NewInt(100)); got != 1 {
		t.Errorf("ConvertToPercentage(100) = %v, want 1", got)
	}
	if got := ConvertToPercentage(big.NewInt(10)); got != 10 {
		t.Errorf("ConvertToPercentage(10) = %v, want 10", got)
	}
}

func TestMaxUint256(t *testing.T) {
	want, _ := new(big.Int).SetString(strings.Repeat("f", 64), 16)
	if got := MaxUint256(); got.Cmp(want) != 0 {
		t.Errorf("MaxUint256 = %s, want 2^256-1", got)
	}
}

func TestAdjustGasPriceBy(t *testing.T) {
	if got := AdjustGasPriceBy(nil, 3); got.Sign() != 0 {
		t.Errorf("AdjustGasPriceBy(nil) = %s, want 0", got)
	}
	base := big.NewInt(1_000_000_000)
	if got := AdjustGasPriceBy(base, 1.0); got.Cmp(base) != 0 {
		t.Errorf("multiplier 1.0 = %s, want base unchanged", got)
	}
	if got := AdjustGasPriceBy(base, 1.5); got.Cmp(big.NewInt(1_500_000_000)) != 0 {
		t.Errorf("multiplier 1.5 = %s, want 1.5 gwei", got)
	}
	if got := AdjustGasPriceBy(base, 3); got.Cmp(big.NewInt(3_000_000_000)) != 0 {
		t.Errorf("multiplier 3 = %s, want 3 gwei", got)
	}
}

func TestOutputKeyValueDurationAndContractInfo(t *testing.T) {
	var out bytes.Buffer
	o := Output{Verbose: true, W: &out}
	o.KeyValueDuration("Waiting", 90)
	o.ContractInfo("CosmicGame Address", common.HexToAddress("0x1100000000000000000000000000000000000011"))
	text := out.String()
	for _, want := range []string{
		"Waiting                     = 90 (1 min 30 sec)",
		"CONTRACT",
		"CosmicGame Address  = 0x1100000000000000000000000000000000000011",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("output missing %q\noutput:\n%s", want, text)
		}
	}

	var quiet bytes.Buffer
	q := Output{Verbose: false, W: &quiet}
	q.KeyValueDuration("Waiting", 90)
	q.ContractInfo("X", common.Address{})
	if quiet.Len() != 0 {
		t.Errorf("quiet output printed %q", quiet.String())
	}
}
