package toolutil

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

func TestLoadCollectorConfig(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	path := filepath.Join(dir, "collector.json")
	document := `{
		"rpc_url":"https://rpc.example",
		"output_dir":"backup",
		"start_block":42,
		"contract_addresses":["0x2100000000000000000000000000000000000021"]
	}`
	if err := os.WriteFile(path, []byte(document), 0o600); err != nil {
		t.Fatal(err)
	}
	config, err := LoadCollectorConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if config.RPCURL != "https://rpc.example" ||
		config.OutputDir != "backup" ||
		config.StartBlock != 42 {
		t.Fatalf("config = %+v", config)
	}

	for name, content := range map[string]string{
		"malformed": `{`,
		"no output": `{"rpc_url":"https://rpc.example"}`,
	} {
		testPath := filepath.Join(dir, name+".json")
		if err := os.WriteFile(testPath, []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
		if _, err := LoadCollectorConfig(testPath); err == nil {
			t.Errorf("%s config was accepted", name)
		}
	}
	if _, err := LoadCollectorConfig(filepath.Join(dir, "missing.json")); err == nil {
		t.Fatal("missing config was accepted")
	}
}

func TestResolveContractAddresses(t *testing.T) {
	t.Parallel()
	first := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	second := "0x2200000000000000000000000000000000000022"
	config := &CollectorConfig{
		ContractAddresses: []string{" " + first + " ", first},
		Contracts: &ContractAddrs{
			CosmicGameAddr:  first,
			CosmicTokenAddr: second,
		},
	}
	got, err := config.ResolveContractAddresses()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		ethcommon.HexToAddress(first).Hex(),
		ethcommon.HexToAddress(second).Hex(),
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("addresses = %v, want %v", got, want)
	}

	if _, err := (&CollectorConfig{
		ContractAddresses: []string{"not-an-address"},
	}).ResolveContractAddresses(); err == nil {
		t.Fatal("invalid address was accepted")
	}
	if _, err := (&CollectorConfig{}).ResolveContractAddresses(); err == nil {
		t.Fatal("empty address config was accepted")
	}
}
