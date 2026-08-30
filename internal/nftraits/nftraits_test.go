package nftraits_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/nftraits"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// exampleSeed is the seed the vendored generator fixture describes.
const exampleSeed = "0x100033"

// paddedExampleSeed is the same seed in the 64-hex-character form the
// indexer stores, because NftMinted carries a uint256.
const paddedExampleSeed = "0x0000000000000000000000000000000000000000000000000000000000100033"

func mustParseExample(tb testing.TB) *nftraits.File {
	tb.Helper()
	file, err := nftraits.Parse(testfixtures.NftTraitsExample)
	if err != nil {
		tb.Fatalf("parsing the vendored fixture: %v", err)
	}
	return file
}

func TestParseExampleFixture(t *testing.T) {
	t.Parallel()
	file := mustParseExample(t)

	if file.SchemaVersion != "1.0.0" {
		t.Errorf("schema_version = %q, want 1.0.0", file.SchemaVersion)
	}
	if file.Seed != exampleSeed {
		t.Errorf("seed = %q, want %q", file.Seed, exampleSeed)
	}
	if file.PipelineVersion == "" {
		t.Error("pipeline_version is empty")
	}
	if len(file.Attributes) != 8 {
		t.Fatalf("len(attributes) = %d, want 8", len(file.Attributes))
	}
	if !strings.HasPrefix(file.DescriptionArt, "Three bodies trace") {
		t.Errorf("description_art = %q, want it to start with the art sentence", file.DescriptionArt)
	}
	if !bytes.Contains(file.Simulation, []byte(`"chaos_index"`)) {
		t.Error("simulation block does not carry chaos_index")
	}
	if !bytes.Contains(file.Generation, []byte(`"palette"`)) {
		t.Error("generation block does not carry palette")
	}
}

// TestAttributesRoundTripVerbatim is the core of the contract: an attribute
// that survives a decode/encode cycle unchanged is one the metadata handler
// can copy without altering the art facts.
func TestAttributesRoundTripVerbatim(t *testing.T) {
	t.Parallel()
	file := mustParseExample(t)

	encoded, err := json.Marshal(file.Attributes)
	if err != nil {
		t.Fatalf("re-encoding attributes: %v", err)
	}

	var original struct {
		Attributes []map[string]any `json:"attributes"`
	}
	if err := json.Unmarshal(testfixtures.NftTraitsExample, &original); err != nil {
		t.Fatalf("decoding the fixture: %v", err)
	}
	var roundTripped []map[string]any
	if err := json.Unmarshal(encoded, &roundTripped); err != nil {
		t.Fatalf("decoding the re-encoded attributes: %v", err)
	}

	wantJSON, err := json.Marshal(original.Attributes)
	if err != nil {
		t.Fatalf("re-encoding the fixture attributes: %v", err)
	}
	gotJSON, err := json.Marshal(roundTripped)
	if err != nil {
		t.Fatalf("re-encoding the round-tripped attributes: %v", err)
	}
	if !bytes.Equal(wantJSON, gotJSON) {
		t.Errorf("attributes changed across the round trip:\n got: %s\nwant: %s", gotJSON, wantJSON)
	}
}

// TestChaosAttributeKeepsNumericSpelling pins the specific hazard raw values
// exist for: a number that decodes to float64 and re-encodes differently.
func TestChaosAttributeKeepsNumericSpelling(t *testing.T) {
	t.Parallel()
	file := mustParseExample(t)

	for _, attr := range file.Attributes {
		if attr.TraitType != "Chaos" {
			continue
		}
		if got := string(attr.Value); got != "30" {
			t.Errorf("Chaos value = %q, want 30", got)
		}
		if got := string(attr.MaxValue); got != "100" {
			t.Errorf("Chaos max_value = %q, want 100", got)
		}
		if attr.DisplayType != "number" {
			t.Errorf("Chaos display_type = %q, want number", attr.DisplayType)
		}
		return
	}
	t.Fatal("the fixture has no Chaos attribute")
}

func TestParseRejectsMalformedInput(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		body string
	}{
		{"empty", ""},
		{"not json", "not json at all"},
		{"truncated", `{"schema_version":"1.0.0","seed":"0x1"`},
		{"array", `[]`},
		{"trailing content", `{"schema_version":"1.0.0","seed":"0x1","pipeline_version":"1","description_art":"a","attributes":[],"simulation":{},"generation":{}}{}`},
		{"missing schema_version", `{"seed":"0x1","pipeline_version":"1","description_art":"a","attributes":[],"simulation":{},"generation":{}}`},
		{"missing seed", `{"schema_version":"1.0.0","pipeline_version":"1","description_art":"a","attributes":[],"simulation":{},"generation":{}}`},
		{"missing pipeline_version", `{"schema_version":"1.0.0","seed":"0x1","description_art":"a","attributes":[],"simulation":{},"generation":{}}`},
		{"missing description_art", `{"schema_version":"1.0.0","seed":"0x1","pipeline_version":"1","attributes":[],"simulation":{},"generation":{}}`},
		{"missing attributes", `{"schema_version":"1.0.0","seed":"0x1","pipeline_version":"1","description_art":"a","simulation":{},"generation":{}}`},
		{"simulation not an object", `{"schema_version":"1.0.0","seed":"0x1","pipeline_version":"1","description_art":"a","attributes":[],"simulation":[],"generation":{}}`},
		{"generation not an object", `{"schema_version":"1.0.0","seed":"0x1","pipeline_version":"1","description_art":"a","attributes":[],"simulation":{},"generation":"x"}`},
		{"attribute without trait_type", `{"schema_version":"1.0.0","seed":"0x1","pipeline_version":"1","description_art":"a","attributes":[{"value":1}],"simulation":{},"generation":{}}`},
		{"attribute without value", `{"schema_version":"1.0.0","seed":"0x1","pipeline_version":"1","description_art":"a","attributes":[{"trait_type":"Structure"}],"simulation":{},"generation":{}}`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if _, err := nftraits.Parse([]byte(tc.body)); err == nil {
				t.Fatal("Parse accepted an invalid trait file")
			}
		})
	}
}

// TestParseToleratesUnknownFields protects the versioning policy: a minor
// bump adds fields, and refusing them would strand every token behind a
// deploy of this repository.
func TestParseToleratesUnknownFields(t *testing.T) {
	t.Parallel()
	body := `{"schema_version":"1.7.0","seed":"0x1","pipeline_version":"1.7.0",
		"description_art":"a","attributes":[{"trait_type":"Structure","value":"Triangle Web"}],
		"simulation":{"masses":[1,2,3]},"generation":{"palette":{}},
		"provenance":{"future":"field"},"generated_at":"2026-01-01T00:00:00Z"}`
	file, err := nftraits.Parse([]byte(body))
	if err != nil {
		t.Fatalf("Parse rejected a forward-compatible minor bump: %v", err)
	}
	if err := file.Gate("0x1"); err != nil {
		t.Errorf("Gate rejected a minor bump: %v", err)
	}
}

func TestGate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		schemaVersion string
		fileSeed      string
		wantSeed      string
		wantErr       error
	}{
		{"accepts the fixture", "1.0.0", exampleSeed, exampleSeed, nil},
		{"accepts a minor bump", "1.4.2", exampleSeed, exampleSeed, nil},
		{"accepts the padded indexer seed", "1.0.0", exampleSeed, paddedExampleSeed, nil},
		{"accepts an unpadded file seed", "1.0.0", paddedExampleSeed, exampleSeed, nil},
		{"accepts mixed case", "1.0.0", "0X10AB", "0x10ab", nil},
		{"refuses a major bump", "2.0.0", exampleSeed, exampleSeed, nftraits.ErrUnsupportedSchema},
		{"refuses a zero major", "0.9.0", exampleSeed, exampleSeed, nftraits.ErrUnsupportedSchema},
		{"refuses a different seed", "1.0.0", "0x100034", exampleSeed, nftraits.ErrSeedMismatch},
		{"refuses a non-hex seed", "1.0.0", "0xzz", exampleSeed, nftraits.ErrSeedMismatch},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			file := &nftraits.File{SchemaVersion: tc.schemaVersion, Seed: tc.fileSeed}
			err := file.Gate(tc.wantSeed)
			switch {
			case tc.wantErr == nil && err != nil:
				t.Fatalf("Gate() = %v, want nil", err)
			case tc.wantErr != nil && !errors.Is(err, tc.wantErr):
				t.Fatalf("Gate() = %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func TestGateRejectsMalformedSemver(t *testing.T) {
	t.Parallel()
	for _, version := range []string{"", "one", "1", "x.0.0", "-1.0.0"} {
		file := &nftraits.File{SchemaVersion: version, Seed: exampleSeed}
		if err := file.Gate(exampleSeed); err == nil {
			t.Errorf("Gate accepted schema_version %q", version)
		}
	}
}

func TestSemverMajor(t *testing.T) {
	t.Parallel()
	tests := []struct {
		version string
		want    int
		wantErr bool
	}{
		{"1.0.0", 1, false},
		{"2.13.4", 2, false},
		{" 1.0.0 ", 1, false},
		{"10.0.0", 10, false},
		{"0.1.0", 0, false},
		{"1", 0, true},
		{"", 0, true},
		{".1.0", 0, true},
		{"v1.0.0", 0, true},
		{"-3.0.0", 0, true},
	}
	for _, tc := range tests {
		t.Run(tc.version, func(t *testing.T) {
			t.Parallel()
			got, err := nftraits.SemverMajor(tc.version)
			if (err != nil) != tc.wantErr {
				t.Fatalf("SemverMajor(%q) error = %v, wantErr %v", tc.version, err, tc.wantErr)
			}
			if err == nil && got != tc.want {
				t.Errorf("SemverMajor(%q) = %d, want %d", tc.version, got, tc.want)
			}
		})
	}
}

func TestCanonicalSeed(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in      string
		want    string
		wantErr bool
	}{
		{"0x100033", "0x100033", false},
		{"100033", "0x100033", false},
		{"0X100ABC", "0x100abc", false},
		{"  0xAbC  ", "0xabc", false},
		{"0x0", "0x0", false},
		{paddedExampleSeed, paddedExampleSeed, false},
		{"", "", true},
		{"0x", "", true},
		{"seed0001", "", true},
		{"0xzz", "", true},
		{"0x12 34", "", true},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()
			got, err := nftraits.CanonicalSeed(tc.in)
			if (err != nil) != tc.wantErr {
				t.Fatalf("CanonicalSeed(%q) error = %v, wantErr %v", tc.in, err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("CanonicalSeed(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestSeedsEquivalent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		a, b string
		want bool
	}{
		{"identical", exampleSeed, exampleSeed, true},
		{"prefix optional", "100033", "0x100033", true},
		{"case insensitive", "0xABC", "0xabc", true},
		{"leading zeros ignored", paddedExampleSeed, exampleSeed, true},
		{"zero seeds", "0x0", "0x0000", true},
		{"different values", "0x100033", "0x100034", false},
		{"non-hex left", "seed01", "0x01", false},
		{"non-hex right", "0x01", "seed01", false},
		{"both empty", "", "", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := nftraits.SeedsEquivalent(tc.a, tc.b); got != tc.want {
				t.Errorf("SeedsEquivalent(%q, %q) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

// TestContentHashIsStableAcrossFormatting is what makes drift detection
// trustworthy: only a change in the art facts may change the digest, never a
// change in how the generator happened to format its JSON.
func TestContentHashIsStableAcrossFormatting(t *testing.T) {
	t.Parallel()
	file := mustParseExample(t)
	base, err := nftraits.ContentHash(file)
	if err != nil {
		t.Fatalf("ContentHash: %v", err)
	}

	var compact bytes.Buffer
	if err := json.Compact(&compact, testfixtures.NftTraitsExample); err != nil {
		t.Fatalf("compacting the fixture: %v", err)
	}
	reformatted, err := nftraits.Parse(compact.Bytes())
	if err != nil {
		t.Fatalf("parsing the compacted fixture: %v", err)
	}
	got, err := nftraits.ContentHash(reformatted)
	if err != nil {
		t.Fatalf("ContentHash: %v", err)
	}
	if got != base {
		t.Errorf("whitespace changed the content hash: %s != %s", got, base)
	}
}

// TestContentHashIgnoresNonArtFields pins the exclusions: a package
// re-uploaded by a newer pipeline with an identical render must not raise
// drift.
func TestContentHashIgnoresNonArtFields(t *testing.T) {
	t.Parallel()
	file := mustParseExample(t)
	base, err := nftraits.ContentHash(file)
	if err != nil {
		t.Fatalf("ContentHash: %v", err)
	}

	file.GeneratedAt = "2030-01-01T00:00:00Z"
	file.PipelineVersion = "9.9.9"
	file.SchemaVersion = "1.9.0"
	got, err := nftraits.ContentHash(file)
	if err != nil {
		t.Fatalf("ContentHash: %v", err)
	}
	if got != base {
		t.Error("a non-art field changed the content hash")
	}
}

func TestContentHashChangesWithArtFacts(t *testing.T) {
	t.Parallel()
	base, err := nftraits.ContentHash(mustParseExample(t))
	if err != nil {
		t.Fatalf("ContentHash: %v", err)
	}

	mutations := map[string]func(*nftraits.File){
		"attribute value": func(f *nftraits.File) {
			f.Attributes[0].Value = json.RawMessage(`"Tangent Caustics"`)
		},
		"attribute added": func(f *nftraits.File) {
			f.Attributes = append(f.Attributes, nftraits.Attribute{
				TraitType: "Wildcard", Value: json.RawMessage(`"Yes"`),
			})
		},
		"attribute order": func(f *nftraits.File) {
			f.Attributes[0], f.Attributes[1] = f.Attributes[1], f.Attributes[0]
		},
		"description": func(f *nftraits.File) { f.DescriptionArt += "." },
		"simulation":  func(f *nftraits.File) { f.Simulation = json.RawMessage(`{"chaos_index":31}`) },
		"generation":  func(f *nftraits.File) { f.Generation = json.RawMessage(`{"palette":{}}`) },
	}
	for name, mutate := range mutations {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			file := mustParseExample(t)
			mutate(file)
			got, err := nftraits.ContentHash(file)
			if err != nil {
				t.Fatalf("ContentHash: %v", err)
			}
			if got == base {
				t.Error("changing an art fact left the content hash unchanged")
			}
		})
	}
}

// TestContentHashPartsCannotCollide guards the length prefixing: moving a
// character across the boundary between two hashed parts must change the
// digest.
func TestContentHashPartsCannotCollide(t *testing.T) {
	t.Parallel()
	left := &nftraits.File{
		DescriptionArt: "ab",
		Simulation:     json.RawMessage(`{"x":1}`),
		Generation:     json.RawMessage(`{}`),
	}
	right := &nftraits.File{
		DescriptionArt: "a",
		Simulation:     json.RawMessage(`{"x":1}`),
		Generation:     json.RawMessage(`{}`),
	}
	lh, err := nftraits.ContentHash(left)
	if err != nil {
		t.Fatalf("ContentHash: %v", err)
	}
	rh, err := nftraits.ContentHash(right)
	if err != nil {
		t.Fatalf("ContentHash: %v", err)
	}
	if lh == rh {
		t.Error("different art facts hashed identically")
	}
}

func TestContentHashRejectsMalformedBlocks(t *testing.T) {
	t.Parallel()
	tests := map[string]*nftraits.File{
		"simulation": {Simulation: json.RawMessage(`{`), Generation: json.RawMessage(`{}`)},
		"generation": {Simulation: json.RawMessage(`{}`), Generation: json.RawMessage(`{`)},
		// A raw attribute value is copied through by the encoder, which
		// validates it: a corrupted value must fail loudly, not produce a
		// digest over invalid JSON.
		"attribute value": {
			Attributes: []nftraits.Attribute{{TraitType: "Structure", Value: json.RawMessage(`{`)}},
			Simulation: json.RawMessage(`{}`),
			Generation: json.RawMessage(`{}`),
		},
	}
	for name, file := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := nftraits.ContentHash(file); err == nil {
				t.Fatal("ContentHash accepted a malformed block")
			}
		})
	}
}

func TestCanonicalJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"sorts keys", `{"b":1,"a":2}`, `{"a":2,"b":1}`},
		{"strips whitespace", "{\n  \"a\" : 1\n}", `{"a":1}`},
		{"preserves float precision", `{"a":274.3025574565995}`, `{"a":274.3025574565995}`},
		{"preserves large integers", `{"a":9007199254740993}`, `{"a":9007199254740993}`},
		{"preserves negative exponents", `{"a":-4.68e-9}`, `{"a":-4.68e-9}`},
		{"handles arrays", `[3,1,2]`, `[3,1,2]`},
		{"handles nesting", `{"b":{"d":1,"c":2}}`, `{"b":{"c":2,"d":1}}`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := nftraits.CanonicalJSON(json.RawMessage(tc.in))
			if err != nil {
				t.Fatalf("CanonicalJSON: %v", err)
			}
			if string(got) != tc.want {
				t.Errorf("CanonicalJSON(%s) = %s, want %s", tc.in, got, tc.want)
			}
		})
	}
}

func TestCanonicalJSONEmptyAndInvalid(t *testing.T) {
	t.Parallel()
	got, err := nftraits.CanonicalJSON(nil)
	if err != nil || got != nil {
		t.Errorf("CanonicalJSON(nil) = (%s, %v), want (nil, nil)", got, err)
	}
	if _, err := nftraits.CanonicalJSON(json.RawMessage(`{oops`)); err == nil {
		t.Error("CanonicalJSON accepted invalid JSON")
	}
}

func TestParseManifest(t *testing.T) {
	t.Parallel()
	manifest, err := nftraits.ParseManifest(testfixtures.AssetManifestExample)
	if err != nil {
		t.Fatalf("ParseManifest: %v", err)
	}
	if manifest.SchemaVersion != nftraits.SupportedManifestSchema {
		t.Errorf("schema_version = %d, want %d", manifest.SchemaVersion, nftraits.SupportedManifestSchema)
	}

	master := manifest.EntryByPath("images/source/master.png")
	if master == nil {
		t.Fatal("the manifest has no master image entry")
	}
	if master.Width == nil || *master.Width != 3456 {
		t.Errorf("master width = %v, want 3456", master.Width)
	}
	if len(master.SHA256) != 64 {
		t.Errorf("master sha256 = %q, want 64 hex characters", master.SHA256)
	}

	video := manifest.EntryByPath("videos/web/main.mp4")
	if video == nil {
		t.Fatal("the manifest has no main video entry")
	}
	if video.Codec != "h264" {
		t.Errorf("video codec = %q, want h264", video.Codec)
	}
	if video.DurationSeconds == nil || *video.DurationSeconds != 30 {
		t.Errorf("video duration = %v, want 30", video.DurationSeconds)
	}

	bins := manifest.EntryByPath("spectral/")
	if bins == nil || bins.FileCount == nil || *bins.FileCount != 64 {
		t.Errorf("spectral bins entry = %+v, want file_count 64", bins)
	}
	if bins != nil && bins.SHA256 != "" {
		t.Error("the directory entry should carry no digest")
	}
}

func TestParseManifestRejectsBadInput(t *testing.T) {
	t.Parallel()
	tests := map[string]string{
		"malformed":     `{`,
		"empty":         ``,
		"schema 1":      `{"schema_version":1,"assets":[]}`,
		"schema absent": `{"assets":[]}`,
	}
	for name, body := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := nftraits.ParseManifest([]byte(body)); err == nil {
				t.Fatal("ParseManifest accepted an invalid manifest")
			}
		})
	}
}

func TestEntryByPathMisses(t *testing.T) {
	t.Parallel()
	var nilManifest *nftraits.Manifest
	if got := nilManifest.EntryByPath("images/source/master.png"); got != nil {
		t.Error("a nil manifest returned an entry")
	}
	manifest := &nftraits.Manifest{Assets: []nftraits.ManifestEntry{{Path: "a"}}}
	if got := manifest.EntryByPath("b"); got != nil {
		t.Error("an unknown path returned an entry")
	}
}

// FuzzParse guards the one place this repository decodes bytes fetched from
// the asset host. Parse must classify every input as valid or invalid
// without panicking, and anything it accepts must survive gating and hashing.
func FuzzParse(f *testing.F) {
	f.Add(testfixtures.NftTraitsExample)
	f.Add([]byte(`{"schema_version":"1.0.0","seed":"0x1","pipeline_version":"1","description_art":"a","attributes":[],"simulation":{},"generation":{}}`))
	f.Add([]byte(`{"attributes":[{"trait_type":"a","value":1,"max_value":2,"display_type":"number"}]}`))
	f.Add([]byte(`{"simulation":{"masses":[1e309]}}`))
	f.Add([]byte(""))
	f.Add([]byte("{"))
	f.Add([]byte("null"))

	f.Fuzz(func(t *testing.T, body []byte) {
		file, err := nftraits.Parse(body)
		if err != nil {
			return
		}
		// Gating a parsed file must never panic, whatever it contains.
		_ = file.Gate(file.Seed)
		if _, err := nftraits.ContentHash(file); err != nil {
			// Hashing may fail only if a block is not valid JSON, which
			// Parse already rejected for the object check.
			t.Fatalf("ContentHash failed on a parsed file: %v", err)
		}
		if _, err := json.Marshal(file.Attributes); err != nil {
			t.Fatalf("re-encoding parsed attributes failed: %v", err)
		}
	})
}
