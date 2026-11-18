package main

import (
	"context"
	"errors"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestEnvUint(t *testing.T) {
	t.Setenv("VC_TEST_UINT", "")
	if got := envUint("VC_TEST_UINT", 7); got != 7 {
		t.Errorf("unset: got %d, want 7", got)
	}
	t.Setenv("VC_TEST_UINT", "12")
	if got := envUint("VC_TEST_UINT", 7); got != 12 {
		t.Errorf("set: got %d, want 12", got)
	}
	t.Setenv("VC_TEST_UINT", "not-a-number")
	if got := envUint("VC_TEST_UINT", 7); got != 7 {
		t.Errorf("garbage: got %d, want 7", got)
	}
	t.Setenv("VC_TEST_UINT", "0")
	if got := envUint("VC_TEST_UINT", 7); got != 7 {
		t.Errorf("zero: got %d, want 7", got)
	}
}

func TestDefaultChainStateFile(t *testing.T) {
	t.Setenv("VALIDATE_CHAIN_STATE", "/tmp/custom-state.json")
	if got := defaultChainStateFile(); got != "/tmp/custom-state.json" {
		t.Errorf("env override: got %q", got)
	}
	t.Setenv("VALIDATE_CHAIN_STATE", "")
	got := defaultChainStateFile()
	if !strings.HasSuffix(got, filepath.Join("cosmic-signature", "validate-chain-state.json")) &&
		got != "validate-chain-state.json" {
		t.Errorf("default: got %q", got)
	}
}

func TestTruncateList(t *testing.T) {
	short := []string{"a", "b"}
	if got := truncateList(short, 5); !reflect.DeepEqual(got, short) {
		t.Errorf("short: got %v", got)
	}
	long := []string{"a", "b", "c", "d"}
	want := []string{"a", "b", "..."}
	if got := truncateList(long, 2); !reflect.DeepEqual(got, want) {
		t.Errorf("long: got %v, want %v", got, want)
	}
}

func TestWrapOne(t *testing.T) {
	vals, err := wrapOne(big.NewInt(5), nil)
	if err != nil || len(vals) != 1 || vals[0].Int64() != 5 {
		t.Errorf("ok case: got %v, %v", vals, err)
	}
	boom := errors.New("boom")
	if _, err := wrapOne(nil, boom); !errors.Is(err, boom) {
		t.Errorf("error case: got %v", err)
	}
}

func TestBigIntSliceEqual(t *testing.T) {
	one, two := big.NewInt(1), big.NewInt(2)
	if !bigIntSliceEqual([]*big.Int{one, two}, []*big.Int{big.NewInt(1), big.NewInt(2)}) {
		t.Error("equal slices reported unequal")
	}
	if bigIntSliceEqual([]*big.Int{one}, []*big.Int{one, two}) {
		t.Error("length mismatch reported equal")
	}
	if bigIntSliceEqual([]*big.Int{one}, []*big.Int{two}) {
		t.Error("value mismatch reported equal")
	}
}

func TestVolatileMatchErrorPaths(t *testing.T) {
	boom := errors.New("boom")
	v := &chainValidator{ctx: context.Background()}
	okFn := func() ([]*big.Int, error) { return []*big.Int{big.NewInt(1)}, nil }

	if _, _, _, err := v.volatileMatch(
		func() ([]*big.Int, error) { return nil, boom }, okFn,
	); !errors.Is(err, boom) {
		t.Errorf("chain error: got %v", err)
	}
	if _, _, _, err := v.volatileMatch(
		okFn, func() ([]*big.Int, error) { return nil, boom },
	); !errors.Is(err, boom) {
		t.Errorf("db error: got %v", err)
	}

	// A persistent mismatch under a cancelled context aborts the retry wait.
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	v = &chainValidator{ctx: cancelled}
	if _, _, _, err := v.volatileMatch(
		okFn, func() ([]*big.Int, error) { return []*big.Int{big.NewInt(2)}, nil },
	); !errors.Is(err, context.Canceled) {
		t.Errorf("cancelled context: got %v", err)
	}
}

func TestLoadChainScanState(t *testing.T) {
	if s := loadChainScanState(filepath.Join(t.TempDir(), "absent.json")); len(s) != 0 {
		t.Errorf("missing file: got %v", s)
	}
	bad := filepath.Join(t.TempDir(), "bad.json")
	if err := os.WriteFile(bad, []byte("{broken"), 0o600); err != nil {
		t.Fatal(err)
	}
	if s := loadChainScanState(bad); len(s) != 0 {
		t.Errorf("bad JSON: got %v", s)
	}
}

func TestSaveChainScanStateErrors(t *testing.T) {
	// The parent "directory" is a regular file: MkdirAll must fail.
	parent := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(parent, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := saveChainScanState(filepath.Join(parent, "x", "state.json"),
		chainScanState{}); err == nil {
		t.Error("expected MkdirAll error")
	}
	// Round-trip through a valid path for the happy branches.
	good := filepath.Join(t.TempDir(), "sub", "state.json")
	state := chainScanState{"k": &chainScanEntry{From: 1, ScannedTo: 2,
		Counts: map[string]uint64{"0xab": 3}}}
	if err := saveChainScanState(good, state); err != nil {
		t.Fatalf("save: %v", err)
	}
	if got := loadChainScanState(good); got["k"] == nil || got["k"].ScannedTo != 2 {
		t.Errorf("round-trip: got %+v", got)
	}
}
