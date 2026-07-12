package main

// Tests for the environment seam: every place cgctl reads the process
// environment is validated here, including the fail-fast policy on
// malformed numeric configuration (the legacy code silently substituted
// defaults).

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRPCURLFromEnv(t *testing.T) {
	t.Setenv("RPC_URL", "")
	if _, err := rpcURLFromEnv(); err == nil {
		t.Error("empty RPC_URL accepted")
	}
	t.Setenv("RPC_URL", "http://example.invalid:8545")
	url, err := rpcURLFromEnv()
	if err != nil || url != "http://example.invalid:8545" {
		t.Errorf("rpcURLFromEnv = %q, %v", url, err)
	}
}

func TestPkeyHexFromEnv(t *testing.T) {
	t.Setenv("PKEY_HEX", "")
	if _, err := pkeyHexFromEnv(); err == nil {
		t.Error("empty PKEY_HEX accepted")
	}
	t.Setenv("PKEY_HEX", "abcd")
	if _, err := pkeyHexFromEnv(); err == nil || !strings.Contains(err.Error(), "64 hex characters") {
		t.Errorf("short PKEY_HEX error = %v", err)
	}
	t.Setenv("PKEY_HEX", strings.Repeat("a", 64))
	if _, err := pkeyHexFromEnv(); err != nil {
		t.Errorf("valid PKEY_HEX rejected: %v", err)
	}
}

func TestGasMultiplierFromEnv(t *testing.T) {
	cases := []struct {
		value   string
		want    float64
		wantErr bool
	}{
		{"", 0, false}, // unset: session default applies
		{"1.5", 1.5, false},
		{"3", 3, false},
		{"abc", 0, true},
		{"0", 0, true},
		{"-2", 0, true},
	}
	for _, tc := range cases {
		t.Run("value="+tc.value, func(t *testing.T) {
			t.Setenv("GAS_PRICE_MULTIPLIER", tc.value)
			got, err := gasMultiplierFromEnv()
			if tc.wantErr {
				if err == nil {
					t.Errorf("GAS_PRICE_MULTIPLIER=%q accepted, want error", tc.value)
				}
				return
			}
			if err != nil || got != tc.want {
				t.Errorf("gasMultiplierFromEnv() = %v, %v; want %v", got, err, tc.want)
			}
		})
	}
}

func TestNewTxSessionReportsEnvErrors(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.SetOut(&bytes.Buffer{})

	t.Setenv("RPC_URL", "")
	t.Setenv("PKEY_HEX", "")
	if _, err := newTxSession(cmd, false); err == nil || !strings.Contains(err.Error(), "RPC_URL") {
		t.Errorf("missing RPC_URL error = %v", err)
	}

	t.Setenv("RPC_URL", "http://127.0.0.1:1")
	if _, err := newTxSession(cmd, false); err == nil || !strings.Contains(err.Error(), "PKEY_HEX") {
		t.Errorf("missing PKEY_HEX error = %v", err)
	}

	t.Setenv("PKEY_HEX", strings.Repeat("a", 64))
	t.Setenv("GAS_PRICE_MULTIPLIER", "bogus")
	if _, err := newTxSession(cmd, false); err == nil || !strings.Contains(err.Error(), "GAS_PRICE_MULTIPLIER") {
		t.Errorf("malformed multiplier error = %v", err)
	}
}

func TestConnectNetworkReportsEnvError(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.SetOut(&bytes.Buffer{})
	t.Setenv("RPC_URL", "")
	if _, _, err := connectNetwork(cmd); err == nil || !strings.Contains(err.Error(), "RPC_URL") {
		t.Errorf("connectNetwork without RPC_URL = %v", err)
	}
}

func TestParseHelpers(t *testing.T) {
	if _, err := parseAddress("addr", "not-an-address"); err == nil {
		t.Error("invalid address accepted")
	}
	if _, err := parseAddress("addr", "0x1111111111111111111111111111111111111111"); err != nil {
		t.Errorf("valid address rejected: %v", err)
	}
	if _, err := parseBigInt("amount", "12x"); err == nil {
		t.Error("invalid big int accepted")
	}
	v, err := parseBigInt("amount", "123456789012345678901234567890")
	if err != nil || v.String() != "123456789012345678901234567890" {
		t.Errorf("parseBigInt = %v, %v", v, err)
	}
	if _, err := parseInt64("n", "abc"); err == nil {
		t.Error("invalid int64 accepted")
	}
	n, err := parseInt64("n", "-42")
	if err != nil || n != -42 {
		t.Errorf("parseInt64 = %v, %v", n, err)
	}
}

func TestPickContractArg(t *testing.T) {
	cases := []struct {
		name string
		args []string
		want string
	}{
		{"empty", nil, ""},
		{"single", []string{"0xabc"}, "0xabc"},
		{"prefixed address second", []string{"0xhash", "0x1111111111111111111111111111111111111111"}, "0x1111111111111111111111111111111111111111"},
		{"bare address first", []string{"1111111111111111111111111111111111111111", "other"}, "0x1111111111111111111111111111111111111111"},
		{"no address shaped", []string{"foo", "bar"}, "bar"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := pickContractArg(tc.args); got != tc.want {
				t.Errorf("pickContractArg(%v) = %q, want %q", tc.args, got, tc.want)
			}
		})
	}
}

func TestAutobidConfigValidation(t *testing.T) {
	valid := map[string]string{
		"RPC_URL":    "http://127.0.0.1:8545",
		"PKEY_HEX":   strings.Repeat("a", 64),
		"CGAME_ADDR": strings.Repeat("1", 40),
	}
	setAll := func(t *testing.T, overrides map[string]string) {
		t.Helper()
		for k, v := range valid {
			t.Setenv(k, v)
		}
		for _, k := range []string{
			"MAX_ETH_BID", "MAX_CST_BID", "RWALK_MIN_PRICE",
			"TIME_BEFORE_PRIZE", "CST_BID_ANYWAY", "AT_STARTUP_BID_UP_TO_PRICE_LEVEL",
			"GAS_PRICE_MULTIPLIER",
		} {
			t.Setenv(k, "")
		}
		for k, v := range overrides {
			t.Setenv(k, v)
		}
	}

	t.Run("defaults", func(t *testing.T) {
		setAll(t, nil)
		cfg, err := loadAutobidConfig()
		if err != nil {
			t.Fatalf("loadAutobidConfig: %v", err)
		}
		if cfg.Limits.MaxEthBid.String() != "5000000000000000000" {
			t.Errorf("MaxEthBid default = %s, want 5 ETH", cfg.Limits.MaxEthBid)
		}
		if cfg.Limits.TimeBeforePrize != 15 || !cfg.Limits.CstBidAnyway {
			t.Errorf("defaults = %+v", cfg.Limits)
		}
		if cfg.InitialBidPrice != nil {
			t.Errorf("InitialBidPrice default = %v, want nil", cfg.InitialBidPrice)
		}
	})

	t.Run("missing required env lists every problem", func(t *testing.T) {
		setAll(t, nil)
		t.Setenv("RPC_URL", "")
		t.Setenv("CGAME_ADDR", "abc")
		_, err := loadAutobidConfig()
		if err == nil {
			t.Fatal("invalid env accepted")
		}
		for _, want := range []string{"RPC_URL is required", "CGAME_ADDR must be 40 hex chars"} {
			if !strings.Contains(err.Error(), want) {
				t.Errorf("error %v missing %q", err, want)
			}
		}
	})

	t.Run("malformed numbers fail fast", func(t *testing.T) {
		for key, bad := range map[string]string{
			"MAX_ETH_BID":                      "lots",
			"MAX_CST_BID":                      "-1",
			"RWALK_MIN_PRICE":                  "x",
			"TIME_BEFORE_PRIZE":                "soon",
			"CST_BID_ANYWAY":                   "maybe",
			"AT_STARTUP_BID_UP_TO_PRICE_LEVEL": "high",
			"GAS_PRICE_MULTIPLIER":             "zero",
		} {
			setAll(t, map[string]string{key: bad})
			if _, err := loadAutobidConfig(); err == nil || !strings.Contains(err.Error(), key) {
				t.Errorf("%s=%q: err = %v, want mention of the variable", key, bad, err)
			}
		}
	})

	t.Run("custom values", func(t *testing.T) {
		setAll(t, map[string]string{
			"MAX_ETH_BID":                      "0.5",
			"TIME_BEFORE_PRIZE":                "30",
			"CST_BID_ANYWAY":                   "false",
			"AT_STARTUP_BID_UP_TO_PRICE_LEVEL": "0.01",
		})
		cfg, err := loadAutobidConfig()
		if err != nil {
			t.Fatalf("loadAutobidConfig: %v", err)
		}
		if cfg.Limits.MaxEthBid.String() != "500000000000000000" {
			t.Errorf("MaxEthBid = %s, want 0.5 ETH", cfg.Limits.MaxEthBid)
		}
		if cfg.Limits.TimeBeforePrize != 30 || cfg.Limits.CstBidAnyway {
			t.Errorf("limits = %+v", cfg.Limits)
		}
		if cfg.InitialBidPrice == nil || cfg.InitialBidPrice.String() != "10000000000000000" {
			t.Errorf("InitialBidPrice = %v, want 0.01 ETH", cfg.InitialBidPrice)
		}
	})

	t.Run("0x prefixed game address accepted", func(t *testing.T) {
		setAll(t, map[string]string{"CGAME_ADDR": "0x" + strings.Repeat("1", 40)})
		if _, err := loadAutobidConfig(); err != nil {
			t.Errorf("0x-prefixed CGAME_ADDR rejected: %v", err)
		}
	})
}
