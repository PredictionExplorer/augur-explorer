package buildjson

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// bindingPackages locates the two binding package directories relative to
// this source file.
func bindingPackages(t *testing.T) []string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot locate source file")
	}
	contractsDir := filepath.Join(filepath.Dir(thisFile), "..", "..")
	return []string{
		filepath.Join(contractsDir, "cosmicgame"),
		filepath.Join(contractsDir, "randomwalk"),
	}
}

// TestArtifactsMatchBindings pins the committed combined-JSON artifacts to
// the committed bindings in both directions: every binding's embedded
// ABI/Bin metadata must equal its artifact entry (modulo abigen's
// whitespace stripping), and every artifact entry must have a binding.
// Together with `make generate-check` (which re-runs the pinned abigen over
// the artifacts and diffs the output) this proves the bindings are
// reproducible from what the repository contains.
func TestArtifactsMatchBindings(t *testing.T) {
	for _, dir := range bindingPackages(t) {
		t.Run(filepath.Base(dir), func(t *testing.T) {
			fromBindings, err := Extract(dir)
			if err != nil {
				t.Fatalf("extracting bindings: %v", err)
			}
			if len(fromBindings) == 0 {
				t.Fatal("no bindings found")
			}

			artifactPath := filepath.Join(dir, "buildjson", "combined.json")
			data, err := os.ReadFile(artifactPath) //nolint:gosec // repo-relative test fixture path
			if err != nil {
				t.Fatalf("reading artifact: %v (run `go run ./contracts/internal/buildjson/extract %s`)", err, dir)
			}
			fromArtifact, err := Decode(data)
			if err != nil {
				t.Fatalf("decoding artifact: %v", err)
			}

			for _, typeName := range SortedKeys(fromBindings) {
				got, ok := fromArtifact[typeName]
				if !ok {
					t.Errorf("binding type %s missing from the artifact", typeName)
					continue
				}
				want := fromBindings[typeName]
				// The binding embeds the whitespace-stripped form of the
				// artifact ABI (abigen strips every space, including the
				// ones inside internalType values — the artifact carries
				// the restored solc form).
				if StripABIWhitespace(got.ABI) != want.ABI {
					t.Errorf("type %s: ABI differs between binding and artifact", typeName)
				}
				if got.Bin != want.Bin {
					t.Errorf("type %s: bytecode differs between binding and artifact", typeName)
				}
			}
			for _, typeName := range SortedKeys(fromArtifact) {
				if _, ok := fromBindings[typeName]; !ok {
					t.Errorf("artifact type %s has no binding", typeName)
				}
			}
		})
	}
}

// TestEncodeDecodeRoundTrip proves Encode output parses back to the same
// metadata (the property `make generate-check` depends on).
func TestEncodeDecodeRoundTrip(t *testing.T) {
	in := map[string]Meta{
		"Token":  {ABI: `[{"name":"Transfer","type":"event"}]`, Bin: "6080"},
		"Helper": {ABI: `[]`},
	}
	data, err := Encode(in)
	if err != nil {
		t.Fatal(err)
	}
	out, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != 2 {
		t.Fatalf("decoded %d types, want 2", len(out))
	}
	for name, want := range in {
		if out[name] != want {
			t.Errorf("%s: round trip = %+v, want %+v", name, out[name], want)
		}
	}
}

// TestDecodeCanonicalizesABI proves artifact ABIs may be pretty-printed or
// key-reordered by hand: Decode normalizes to the sorted-key compact form
// abigen embeds.
func TestDecodeCanonicalizesABI(t *testing.T) {
	data := []byte(`{
  "contracts": {
    "X.sol:X": {
      "abi": [ { "type": "event", "name": "E", "anonymous": false } ],
      "bin": ""
    }
  }
}`)
	out, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	want := `[{"anonymous":false,"name":"E","type":"event"}]`
	if out["X"].ABI != want {
		t.Errorf("canonical ABI = %s, want %s", out["X"].ABI, want)
	}
}

// TestRestoreInternalTypeSpaces pins the restoration of abigen's stripped
// internalType spaces for every composite marker (struct/enum/contract,
// nested tuple components included) while elementary types pass through,
// and that abigen's whitespace stripping is its exact inverse on canonical
// (sorted-key, stripped) input — the form every binding embeds.
func TestRestoreInternalTypeSpaces(t *testing.T) {
	in := `[{"inputs":[` +
		`{"components":[` +
		`{"internalType":"contractRandomWalkNFT","type":"address"},` +
		`{"internalType":"enumIGovernor.ProposalState","type":"uint8"},` +
		`{"internalType":"uint256","type":"uint256"},` +
		`{"internalType":"string","type":"string"}` +
		`],"internalType":"structICosmicSignatureToken.MintSpec[]","type":"tuple[]"}` +
		`],"name":"f","type":"function"}]`
	out, err := RestoreInternalTypeSpaces(in)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`"internalType":"struct ICosmicSignatureToken.MintSpec[]"`,
		`"internalType":"contract RandomWalkNFT"`,
		`"internalType":"enum IGovernor.ProposalState"`,
		`"internalType":"uint256"`,
		`"internalType":"string"`,
	} {
		if !strings.Contains(out, want) {
			t.Errorf("restored ABI missing %s:\n%s", want, out)
		}
	}
	// Stripping must be the exact inverse used by abigen.
	if got := StripABIWhitespace(out); got != in {
		t.Errorf("strip(restore(x)) != x:\n got %s\nwant %s", got, in)
	}
}

// TestRestoreOnePrefix pins the per-value edge cases: already-spaced values
// and bare markers pass through, lowercase continuations (elementary types
// like "string") never gain a space.
func TestRestoreOnePrefix(t *testing.T) {
	cases := map[string]string{
		"structFoo.Bar":      "struct Foo.Bar",
		"enumFoo.Bar":        "enum Foo.Bar",
		"contractFoo":        "contract Foo",
		"struct Already.Ok":  "struct Already.Ok",
		"contract Spaced":    "contract Spaced",
		"struct":             "struct",
		"string":             "string",
		"uint256":            "uint256",
		"structs_lowercase":  "structs_lowercase",
		"address":            "address",
		"structFoo.Bar[][3]": "struct Foo.Bar[][3]",
	}
	for in, want := range cases {
		if got := restoreOnePrefix(in); got != want {
			t.Errorf("restoreOnePrefix(%q) = %q, want %q", in, got, want)
		}
	}
}
