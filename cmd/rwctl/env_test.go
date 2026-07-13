package main

import (
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPkeyHexFromEnv(t *testing.T) {
	t.Setenv("PKEY_HEX", "")
	if _, err := pkeyHexFromEnv(); err == nil || !strings.Contains(err.Error(), "required") {
		t.Errorf("missing PKEY_HEX error = %v", err)
	}

	t.Setenv("PKEY_HEX", "abcd")
	if _, err := pkeyHexFromEnv(); err == nil || !strings.Contains(err.Error(), "64 hex characters") {
		t.Errorf("short PKEY_HEX error = %v", err)
	}

	valid := strings.Repeat("ab", 32)
	t.Setenv("PKEY_HEX", valid)
	got, err := pkeyHexFromEnv()
	if err != nil || got != valid {
		t.Errorf("pkeyHexFromEnv = (%q, %v)", got, err)
	}
}

func TestParseInt64(t *testing.T) {
	if got, err := parseInt64("token_id", "42"); err != nil || got != 42 {
		t.Errorf("parseInt64 = (%d, %v)", got, err)
	}
	if _, err := parseInt64("token_id", "4.2"); err == nil || !strings.Contains(err.Error(), "token_id") {
		t.Errorf("bad int error = %v, want the argument name", err)
	}
}

func TestParseBigInt(t *testing.T) {
	want, _ := new(big.Int).SetString("123456789012345678901234567890", 10)
	got, err := parseBigInt("price_wei", "123456789012345678901234567890")
	if err != nil || got.Cmp(want) != 0 {
		t.Errorf("parseBigInt = (%s, %v)", got, err)
	}
	if _, err := parseBigInt("price_wei", "0x10"); err == nil || !strings.Contains(err.Error(), "price_wei") {
		t.Errorf("bad big int error = %v, want the argument name", err)
	}
}

func TestReadTwitterKeysRoundTrip(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("TWITTER_KEYS_FILE", "keys.json")
	if err := os.MkdirAll(filepath.Join(home, "configs"), 0o750); err != nil {
		t.Fatal(err)
	}

	if _, err := readTwitterKeys(); err == nil || !strings.Contains(err.Error(), "can't read") {
		t.Errorf("missing keys file error = %v", err)
	}

	path := filepath.Join(home, "configs", "keys.json")
	if err := os.WriteFile(path, []byte("{bad"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := readTwitterKeys(); err == nil || !strings.Contains(err.Error(), "can't parse") {
		t.Errorf("bad keys file error = %v", err)
	}

	if err := os.WriteFile(path, []byte(`{"ApiKey":"a","ApiSecret":"b","TokenKey":"c","TokenSecret":"d"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	keys, err := readTwitterKeys()
	if err != nil || keys.APIKey != "a" || keys.TokenSecret != "d" {
		t.Errorf("readTwitterKeys = (%+v, %v)", keys, err)
	}
}
