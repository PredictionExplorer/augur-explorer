package cosmicgame

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/nftraits"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

const (
	testAssetBase  = "https://nfts.cosmicsignature.com/images"
	testSourceBase = "https://nfts.cosmicsignature.com"
	// testSeed is the 64-hex-character form the indexer stores for the seed
	// the vendored generator fixture describes.
	testSeed = "0x0000000000000000000000000000000000000000000000000000000000100033"
	// mintedAt is 2026-08-28T09:36:15Z, the fixture's generation moment
	// (04:36:15-05:00) expressed in UTC.
	mintedAt = int64(1787909775)
)

// exampleTraits decodes the vendored generator fixture the way the ingester
// would have stored it, so assembly tests run against real generator output.
func exampleTraits(tb testing.TB, withManifest bool) *tokenTraits {
	tb.Helper()
	file, err := nftraits.Parse(testfixtures.NftTraitsExample)
	if err != nil {
		tb.Fatalf("parsing the vendored fixture: %v", err)
	}
	attributes, err := json.Marshal(file.Attributes)
	if err != nil {
		tb.Fatalf("encoding attributes: %v", err)
	}
	simulation, err := nftraits.CanonicalJSON(file.Simulation)
	if err != nil {
		tb.Fatalf("canonicalizing simulation: %v", err)
	}
	generation, err := nftraits.CanonicalJSON(file.Generation)
	if err != nil {
		tb.Fatalf("canonicalizing generation: %v", err)
	}
	row := cgstore.TokenTraitsRow{
		Seed:            "0x100033",
		SchemaMajor:     1,
		PipelineVersion: file.PipelineVersion,
		Attributes:      attributes,
		DescriptionArt:  file.DescriptionArt,
		Simulation:      simulation,
		Generation:      generation,
	}
	if withManifest {
		row.Assets = testfixtures.AssetManifestExample
	}
	traits, err := decodeTokenTraits(row)
	if err != nil {
		tb.Fatalf("decoding the stored row: %v", err)
	}
	return traits
}

func enrichedInput(tb testing.TB) tokenMetadataInput {
	tb.Helper()
	return tokenMetadataInput{
		TokenID:       47,
		RoundNum:      1,
		MintTimestamp: mintedAt,
		Seed:          testSeed,
		AssetBase:     testAssetBase,
		SourceBase:    testSourceBase,
		Allocation:    "Final Gesture",
		Traits:        exampleTraits(tb, true),
	}
}

// marshalMetadata renders the document the way the handler does, so
// assertions run against the bytes a marketplace actually receives.
func marshalMetadata(tb testing.TB, in tokenMetadataInput) map[string]any {
	tb.Helper()
	encoded, err := json.Marshal(buildTokenMetadata(in))
	if err != nil {
		tb.Fatalf("marshaling metadata: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		tb.Fatalf("decoding metadata: %v", err)
	}
	return decoded
}

// TestMetadataNeverCarriesOwner is the one guarantee that must hold on every
// path: ownership changes on transfer while marketplaces cache for days.
func TestMetadataNeverCarriesOwner(t *testing.T) {
	t.Parallel()
	cases := map[string]tokenMetadataInput{
		"enriched": enrichedInput(t),
		"fallback": {TokenID: 1, Seed: testSeed, AssetBase: testAssetBase},
	}
	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			doc := marshalMetadata(t, in)
			if _, ok := doc["owner"]; ok {
				t.Error("the document carries a top-level owner")
			}
			properties, _ := doc["properties"].(map[string]any)
			if _, ok := properties["owner"]; ok {
				t.Error("properties carries an owner")
			}
			encoded, err := json.Marshal(doc)
			if err != nil {
				t.Fatalf("re-encoding the document: %v", err)
			}
			if strings.Contains(string(encoded), `"owner"`) {
				t.Errorf("the serialized document mentions an owner: %s", encoded)
			}
		})
	}
}

// TestMetadataHasNoSeedAttribute pins the trait-panel cleanup: a unique
// value can never filter anything, so it only crowds out traits that can.
func TestMetadataHasNoSeedAttribute(t *testing.T) {
	t.Parallel()
	cases := map[string]tokenMetadataInput{
		"enriched": enrichedInput(t),
		"fallback": {TokenID: 1, Seed: testSeed, AssetBase: testAssetBase},
	}
	for name, in := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, attr := range buildAttributes(in) {
				if strings.EqualFold(attr.TraitType, "seed") {
					t.Fatalf("the trait panel still carries a seed attribute: %+v", attr)
				}
			}
			properties, _ := marshalMetadata(t, in)["properties"].(map[string]any)
			if got := properties["seed"]; got != testSeed {
				t.Errorf("properties.seed = %v, want %q", got, testSeed)
			}
		})
	}
}

// TestArtAttributesAreVerbatimAndFirst is the merge contract: the generator
// owns every seed-derived fact and the handler only appends to it.
func TestArtAttributesAreVerbatimAndFirst(t *testing.T) {
	t.Parallel()
	in := enrichedInput(t)
	got := buildAttributes(in)

	art := in.Traits.Attributes
	if len(got) != len(art)+3 {
		t.Fatalf("len(attributes) = %d, want %d art + Round + Imprinted + Allocation", len(got), len(art))
	}
	for i, want := range art {
		if got[i].TraitType != want.TraitType {
			t.Errorf("attribute %d trait_type = %q, want %q", i, got[i].TraitType, want.TraitType)
		}
		if string(got[i].Value) != string(want.Value) {
			t.Errorf("attribute %q value = %s, want %s", want.TraitType, got[i].Value, want.Value)
		}
		if string(got[i].MaxValue) != string(want.MaxValue) {
			t.Errorf("attribute %q max_value = %s, want %s", want.TraitType, got[i].MaxValue, want.MaxValue)
		}
		if got[i].DisplayType != want.DisplayType {
			t.Errorf("attribute %q display_type = %q, want %q", want.TraitType, got[i].DisplayType, want.DisplayType)
		}
	}

	chain := got[len(art):]
	wantChain := []struct{ traitType, displayType, value string }{
		{"Round", "number", "1"},
		{"Imprinted", "date", "1787909775"},
		{"Allocation", "", `"Final Gesture"`},
	}
	for i, want := range wantChain {
		if chain[i].TraitType != want.traitType {
			t.Errorf("chain attribute %d = %q, want %q", i, chain[i].TraitType, want.traitType)
		}
		if chain[i].DisplayType != want.displayType {
			t.Errorf("%s display_type = %q, want %q", want.traitType, chain[i].DisplayType, want.displayType)
		}
		if string(chain[i].Value) != want.value {
			t.Errorf("%s value = %s, want %s", want.traitType, chain[i].Value, want.value)
		}
	}
}

// TestChaosAttributeSurvivesTheFullPipeline follows one numeric art trait
// from the generator's bytes to the served document.
func TestChaosAttributeSurvivesTheFullPipeline(t *testing.T) {
	t.Parallel()
	doc := marshalMetadata(t, enrichedInput(t))
	attributes, ok := doc["attributes"].([]any)
	if !ok {
		t.Fatal("attributes is not an array")
	}
	for _, raw := range attributes {
		attr, _ := raw.(map[string]any)
		if attr["trait_type"] != "Chaos" {
			continue
		}
		if got := attr["value"]; got != float64(30) {
			t.Errorf("Chaos value = %v, want 30", got)
		}
		if got := attr["max_value"]; got != float64(100) {
			t.Errorf("Chaos max_value = %v, want 100", got)
		}
		if got := attr["display_type"]; got != "number" {
			t.Errorf("Chaos display_type = %v, want number", got)
		}
		return
	}
	t.Fatal("the served document has no Chaos attribute")
}

func TestChainAttributesOmittedWhenUnavailable(t *testing.T) {
	t.Parallel()
	in := tokenMetadataInput{TokenID: 1, RoundNum: 4, Seed: testSeed, AssetBase: testAssetBase}
	got := buildAttributes(in)
	if len(got) != 1 || got[0].TraitType != "Round" {
		t.Fatalf("attributes = %+v, want only Round", got)
	}

	in.MintTimestamp = -1
	if got := buildAttributes(in); len(got) != 1 {
		t.Errorf("a negative mint timestamp produced an Imprinted attribute: %+v", got)
	}
}

func TestBuildDescription(t *testing.T) {
	t.Parallel()
	traits := exampleTraits(t, false)

	tests := []struct {
		name        string
		in          tokenMetadataInput
		wantPrefix  string
		wantContain []string
	}{
		{
			name:       "enriched leads with the art sentence",
			in:         tokenMetadataInput{TokenID: 47, RoundNum: 1, MintTimestamp: mintedAt, Traits: traits},
			wantPrefix: traits.DescriptionArt,
			wantContain: []string{
				"Imprinted in Round 1 on 28 Aug 2026.",
				"Same seed, same pixels — re-render it to verify.",
			},
		},
		{
			name:        "enriched without a mint time omits the date",
			in:          tokenMetadataInput{TokenID: 47, RoundNum: 2, Traits: traits},
			wantPrefix:  traits.DescriptionArt,
			wantContain: []string{"Imprinted in Round 2. Same seed"},
		},
		{
			name:        "fallback keeps the long-standing sentence",
			in:          tokenMetadataInput{TokenID: 9, RoundNum: 1},
			wantPrefix:  "Discover the unique attributes",
			wantContain: []string{"Cosmic Signature Token #9"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := buildDescription(tc.in)
			if !strings.HasPrefix(got, tc.wantPrefix) {
				t.Errorf("description = %q, want it to start with %q", got, tc.wantPrefix)
			}
			for _, want := range tc.wantContain {
				if !strings.Contains(got, want) {
					t.Errorf("description = %q, want it to contain %q", got, want)
				}
			}
		})
	}
}

// TestDescriptionDateIsUTC keeps the provenance sentence deterministic
// wherever the server happens to run.
func TestDescriptionDateIsUTC(t *testing.T) {
	t.Parallel()
	// 2026-01-01T00:30:00Z is still 2025-12-31 in the Americas.
	in := tokenMetadataInput{RoundNum: 3, MintTimestamp: 1767227400, Traits: exampleTraits(t, false)}
	if got := buildDescription(in); !strings.Contains(got, "on 01 Jan 2026") {
		t.Errorf("description = %q, want the UTC date 01 Jan 2026", got)
	}
}

func TestTokenDisplayName(t *testing.T) {
	t.Parallel()
	if got := tokenDisplayName("Genesis", 1); got != "Genesis" {
		t.Errorf("name = %q, want the custom on-chain name", got)
	}
	if got := tokenDisplayName("", 47); got != "Cosmic Signature #47" {
		t.Errorf("name = %q, want the numbered default", got)
	}
	if got := tokenDisplayName("", 0); got != "Cosmic Signature #0" {
		t.Errorf("name = %q, want token 0 to be numbered", got)
	}
}

func TestBuildTokenMetadataTopLevelFields(t *testing.T) {
	t.Parallel()
	doc := marshalMetadata(t, enrichedInput(t))

	want := map[string]string{
		"metadata_version": metadataVersion,
		"name":             "Cosmic Signature #47",
		"image":            testAssetBase + "/new/cosmicsignature/" + testSeed + ".png",
		"animation_url":    testAssetBase + "/new/cosmicsignature/" + testSeed + ".mp4",
		"background_color": "000000",
		"external_url":     "https://www.cosmicsignature.com/detail/47",
	}
	for key, expected := range want {
		if got := doc[key]; got != expected {
			t.Errorf("%s = %v, want %q", key, got, expected)
		}
	}
}

func TestBuildPropertiesEnriched(t *testing.T) {
	t.Parallel()
	doc := marshalMetadata(t, enrichedInput(t))
	properties, ok := doc["properties"].(map[string]any)
	if !ok {
		t.Fatal("properties is not an object")
	}
	if properties["seed"] != testSeed {
		t.Errorf("seed = %v, want %q", properties["seed"], testSeed)
	}
	if properties["token_id"] != float64(47) || properties["round_num"] != float64(1) {
		t.Errorf("token_id/round_num = %v/%v, want 47/1", properties["token_id"], properties["round_num"])
	}

	simulation, ok := properties["simulation"].(map[string]any)
	if !ok {
		t.Fatal("properties.simulation is not an object")
	}
	if simulation["chaos_index"] != float64(30) {
		t.Errorf("simulation.chaos_index = %v, want 30", simulation["chaos_index"])
	}
	generation, ok := properties["generation"].(map[string]any)
	if !ok {
		t.Fatal("properties.generation is not an object")
	}
	palette, _ := generation["palette"].(map[string]any)
	if palette["family"] != "Aurora Split" {
		t.Errorf("generation.palette.family = %v, want Aurora Split", palette["family"])
	}
}

func TestBuildPropertiesFallbackIsMinimal(t *testing.T) {
	t.Parallel()
	in := tokenMetadataInput{TokenID: 1, RoundNum: 0, Seed: testSeed, AssetBase: testAssetBase, SourceBase: testSourceBase}
	properties, ok := marshalMetadata(t, in)["properties"].(map[string]any)
	if !ok {
		t.Fatal("properties is not an object")
	}
	if len(properties) != 3 {
		t.Errorf("fallback properties = %v, want only seed, token_id and round_num", properties)
	}
	for _, key := range []string{"simulation", "generation", "media"} {
		if _, ok := properties[key]; ok {
			t.Errorf("the fallback carries properties.%s", key)
		}
	}
}

// TestBuildMediaAdvertisesEveryHostedAsset is the discoverability fix: the
// HQ video, the sweeps and the spectral bins are hosted but unreachable from
// the image and animation_url fields alone.
func TestBuildMediaAdvertisesEveryHostedAsset(t *testing.T) {
	t.Parallel()
	properties, _ := marshalMetadata(t, enrichedInput(t))["properties"].(map[string]any)
	media, ok := properties["media"].(map[string]any)
	if !ok {
		t.Fatal("properties.media is not an object")
	}

	pkg := testAssetBase + "/new/cosmicsignature/" + testSeed
	want := map[string]string{
		"hq_video":           pkg + "/videos/hq/main.mp4",
		"spectral_sweep":     pkg + "/videos/web/spectral_sweep.mp4",
		"spectral_sweep_hq":  pkg + "/videos/hq/spectral_sweep.mp4",
		"spectral_bins":      pkg + "/spectral/",
		"web_image":          pkg + "/images/web/full.webp",
		"source_image":       pkg + "/images/source/master.png",
		"preview_image":      pkg + "/images/web/preview.webp",
		"generation_records": pkg + "/metadata/",
		"asset_manifest":     testSourceBase + "/asset-manifests/" + testSeed + ".json",
		"trait_source":       testSourceBase + "/traits/" + testSeed + ".json",
	}
	for key, expected := range want {
		if got := media[key]; got != expected {
			t.Errorf("media.%s = %v, want %q", key, got, expected)
		}
	}
	if len(media) != len(want) {
		t.Errorf("media has %d entries, want %d: %v", len(media), len(want), media)
	}
}

// TestBuildMediaOmitsContractLinksWithoutAnAssetHost keeps the document
// honest: those two paths only exist where the asset host publishes them.
func TestBuildMediaOmitsContractLinksWithoutAnAssetHost(t *testing.T) {
	t.Parallel()
	in := enrichedInput(t)
	in.SourceBase = ""
	properties, _ := marshalMetadata(t, in)["properties"].(map[string]any)
	media, _ := properties["media"].(map[string]any)

	for _, key := range []string{"asset_manifest", "trait_source"} {
		if _, ok := media[key]; ok {
			t.Errorf("media.%s was emitted without a configured asset host", key)
		}
	}
	if _, ok := media["hq_video"]; !ok {
		t.Error("the package-relative media links were dropped too")
	}
}

func TestAssetDetailsFromTheManifest(t *testing.T) {
	t.Parallel()
	doc := marshalMetadata(t, enrichedInput(t))

	image, ok := doc["image_details"].(map[string]any)
	if !ok {
		t.Fatal("image_details is missing")
	}
	if image["width"] != float64(3456) || image["height"] != float64(2234) {
		t.Errorf("image dimensions = %v x %v, want 3456 x 2234", image["width"], image["height"])
	}
	if image["format"] != "png" {
		t.Errorf("image format = %v, want png", image["format"])
	}
	if sha, _ := image["sha256"].(string); len(sha) != 64 {
		t.Errorf("image sha256 = %v, want 64 hex characters", image["sha256"])
	}
	if _, ok := image["codec"]; ok {
		t.Error("image_details carries a codec")
	}

	animation, ok := doc["animation_details"].(map[string]any)
	if !ok {
		t.Fatal("animation_details is missing")
	}
	if animation["codec"] != "h264" {
		t.Errorf("animation codec = %v, want h264", animation["codec"])
	}
	if animation["duration_seconds"] != float64(30) {
		t.Errorf("animation duration = %v, want 30", animation["duration_seconds"])
	}
	if animation["bytes"] != float64(41982155) {
		t.Errorf("animation bytes = %v, want 41982155", animation["bytes"])
	}
}

func TestAssetDetailsOmittedWithoutAManifest(t *testing.T) {
	t.Parallel()
	in := enrichedInput(t)
	in.Traits = exampleTraits(t, false)
	doc := marshalMetadata(t, in)
	for _, key := range []string{"image_details", "animation_details"} {
		if _, ok := doc[key]; ok {
			t.Errorf("%s was emitted without a manifest", key)
		}
	}
}

func TestAssetDetailsOmittedForUndescribedPaths(t *testing.T) {
	t.Parallel()
	// An image-only generator run publishes no video entries.
	manifest := &nftraits.Manifest{
		SchemaVersion: 2,
		Assets: []nftraits.ManifestEntry{{
			Path: masterImagePath, Format: "png", SHA256: strings.Repeat("a", 64),
		}},
	}
	if imageDetails(manifest) == nil {
		t.Error("imageDetails dropped a described image")
	}
	if animationDetails(manifest) != nil {
		t.Error("animationDetails invented a video that the manifest does not describe")
	}
	if imageDetails(nil) != nil || animationDetails(nil) != nil {
		t.Error("a nil manifest produced details")
	}

	// A described entry with no usable fields is not worth an empty object.
	bare := &nftraits.Manifest{Assets: []nftraits.ManifestEntry{
		{Path: masterImagePath},
		{Path: mainVideoPath},
	}}
	if got := imageDetails(bare); got != nil {
		t.Errorf("imageDetails = %v, want nil for a field-less entry", got)
	}
	if got := animationDetails(bare); got != nil {
		t.Errorf("animationDetails = %v, want nil for a field-less entry", got)
	}
}

func TestAllocationLabel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		source cgstore.CosmicSignatureMintSource
		want   string
	}{
		{cgstore.MintSourceMainPrize, "Final Gesture"},
		{cgstore.MintSourceEnduranceChampion, "Endurance Champion"},
		{cgstore.MintSourceChronoWarriorPrize, "Chrono-Warrior"},
		{cgstore.MintSourceBidderRaffle, "Stellar Selection"},
		{cgstore.MintSourceCosmicSigStaker, "Anchored Selection"},
		{cgstore.MintSourceRandomWalkStaker, "Anchored Selection"},
		{cgstore.MintSourceLastCstBidder, "Last CST Gesture"},
		{cgstore.CosmicSignatureMintSource("somethingNew"), ""},
		{"", ""},
	}
	for _, tc := range tests {
		t.Run(string(tc.source), func(t *testing.T) {
			t.Parallel()
			if got := allocationLabel(tc.source); got != tc.want {
				t.Errorf("allocationLabel(%q) = %q, want %q", tc.source, got, tc.want)
			}
		})
	}
}

// TestEveryMintSourceHasAnAllocation fails when the indexer grows a prize
// family that the served provenance vocabulary has not caught up with.
func TestEveryMintSourceHasAnAllocation(t *testing.T) {
	t.Parallel()
	sources := []cgstore.CosmicSignatureMintSource{
		cgstore.MintSourceMainPrize,
		cgstore.MintSourceBidderRaffle,
		cgstore.MintSourceRandomWalkStaker,
		cgstore.MintSourceCosmicSigStaker,
		cgstore.MintSourceEnduranceChampion,
		cgstore.MintSourceLastCstBidder,
		cgstore.MintSourceChronoWarriorPrize,
	}
	for _, source := range sources {
		if allocationLabel(source) == "" {
			t.Errorf("mint source %q has no Allocation label", source)
		}
	}
	if len(allocationLabels) != len(sources) {
		t.Errorf("the allocation table has %d entries for %d mint sources", len(allocationLabels), len(sources))
	}
}

func TestMetadataSeed(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		in    string
		want  string
		exact bool
	}{
		{name: "bare hex gains the prefix", in: "100033", want: "0x100033"},
		{name: "prefixed hex is kept", in: "0x100033", want: "0x100033"},
		{name: "uppercase is lowered", in: "0X100ABC", want: "0x100abc"},
		{name: "padded indexer seed is kept", in: strings.TrimPrefix(testSeed, "0x"), want: testSeed},
		// Legacy and fixture rows that are not hexadecimal keep the
		// long-standing behaviour so their assets still resolve.
		{name: "non-hex keeps the legacy prefixing", in: "seed0001", want: "0xseed0001"},
		{name: "empty", in: "", want: "0x"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := metadataSeed(tc.in); got != tc.want {
				t.Errorf("metadataSeed(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestDecodeTokenTraitsRejectsCorruptRows(t *testing.T) {
	t.Parallel()
	tests := map[string]cgstore.TokenTraitsRow{
		"broken attributes": {Attributes: []byte(`{`)},
		"broken manifest":   {Attributes: []byte(`[]`), Assets: []byte(`{`)},
	}
	for name, row := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := decodeTokenTraits(row); err == nil {
				t.Fatal("decodeTokenTraits accepted a corrupt row")
			}
		})
	}

	traits, err := decodeTokenTraits(cgstore.TokenTraitsRow{Attributes: []byte(`[]`)})
	if err != nil {
		t.Fatalf("decodeTokenTraits: %v", err)
	}
	if traits.Manifest != nil {
		t.Error("a row without assets produced a manifest")
	}
}

func TestPutHelpers(t *testing.T) {
	t.Parallel()
	h := httpx.H{}
	putIfSet[int](h, "absent", nil)
	value := 7
	putIfSet(h, "present", &value)
	putIfNotEmpty(h, "empty", "")
	putIfNotEmpty(h, "filled", "x")

	if _, ok := h["absent"]; ok {
		t.Error("putIfSet stored a nil pointer")
	}
	if h["present"] != 7 {
		t.Errorf("present = %v, want 7", h["present"])
	}
	if _, ok := h["empty"]; ok {
		t.Error("putIfNotEmpty stored an empty string")
	}
	if h["filled"] != "x" {
		t.Errorf("filled = %v, want x", h["filled"])
	}
}

// TestMetadataIsDeterministic keeps the golden parity suite meaningful: the
// same inputs must always serialize to the same bytes.
func TestMetadataIsDeterministic(t *testing.T) {
	t.Parallel()
	in := enrichedInput(t)
	first, err := json.Marshal(buildTokenMetadata(in))
	if err != nil {
		t.Fatalf("marshaling: %v", err)
	}
	for range 20 {
		next, err := json.Marshal(buildTokenMetadata(in))
		if err != nil {
			t.Fatalf("marshaling: %v", err)
		}
		if string(next) != string(first) {
			t.Fatal("the served document is not byte-stable across renders")
		}
	}
}
