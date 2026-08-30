//go:build integration

package apitest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/api/routes"
	"github.com/PredictionExplorer/augur-explorer/internal/nftraits"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// enrichedTokenID is the fixture token the trait tests lend a hexadecimal
// seed to. It is the round 0 main-prize mint, so its Allocation is the one
// the collection calls "Final Gesture".
const enrichedTokenID = 1

// enrichedSeed is the vendored fixture's seed 0x100033 in the 64-character
// zero-padded form the indexer writes, because NftMinted carries a uint256.
// Serving it proves the two spellings resolve to one package.
const enrichedSeed = "0x0000000000000000000000000000000000000000000000000000000000100033"

// lendHexSeed gives the fixture token a hexadecimal seed for the duration of
// one test and clears any trait rows written against it.
//
// The shared dataset keeps placeholder seeds on purpose — changing them
// would churn forty unrelated goldens — and on_mint_update is a no-op, so
// the borrowed seed leaves no aggregate behind. Tests using it must not run
// in parallel with the rest of the suite.
func lendHexSeed(t *testing.T, h *harness) {
	t.Helper()
	ctx := context.Background()

	var original string
	if err := h.db.QueryRowContext(ctx,
		"SELECT seed FROM cg_mint_event WHERE token_id=$1", enrichedTokenID,
	).Scan(&original); err != nil {
		t.Fatalf("reading the fixture seed: %v", err)
	}
	if _, err := h.db.ExecContext(ctx,
		"UPDATE cg_mint_event SET seed=$1 WHERE token_id=$2",
		strings.TrimPrefix(enrichedSeed, "0x"), enrichedTokenID,
	); err != nil {
		t.Fatalf("lending the token a hex seed: %v", err)
	}
	t.Cleanup(func() {
		if _, err := h.db.ExecContext(ctx,
			"UPDATE cg_mint_event SET seed=$1 WHERE token_id=$2", original, enrichedTokenID,
		); err != nil {
			t.Errorf("restoring the fixture seed: %v", err)
		}
		if _, err := h.db.ExecContext(ctx,
			"DELETE FROM cg_token_traits WHERE seed=$1", enrichedSeed,
		); err != nil {
			t.Errorf("clearing the trait row: %v", err)
		}
		if _, err := h.db.ExecContext(ctx,
			"DELETE FROM cg_token_traits_fetch WHERE seed=$1", enrichedSeed,
		); err != nil {
			t.Errorf("clearing the fetch row: %v", err)
		}
	})
}

// withIngestedTraits lends the fixture token a hexadecimal seed and stores
// the vendored generator package against it directly, bypassing the ingest
// loop so golden snapshots stay deterministic. TestTraitIngestEndToEnd
// covers the loop itself.
func withIngestedTraits(t *testing.T, h *harness, withManifest bool) {
	t.Helper()
	ctx := context.Background()
	repo := cgstore.NewRepo(h.store)
	lendHexSeed(t, h)

	file, err := nftraits.Parse(testfixtures.NftTraitsExample)
	if err != nil {
		t.Fatalf("parsing the vendored fixture: %v", err)
	}
	if err := file.Gate(enrichedSeed); err != nil {
		t.Fatalf("the vendored fixture does not gate against the padded seed: %v", err)
	}
	hash, err := nftraits.ContentHash(file)
	if err != nil {
		t.Fatalf("hashing the fixture: %v", err)
	}
	attributes, err := json.Marshal(file.Attributes)
	if err != nil {
		t.Fatalf("encoding attributes: %v", err)
	}
	simulation, err := nftraits.CanonicalJSON(file.Simulation)
	if err != nil {
		t.Fatalf("canonicalizing simulation: %v", err)
	}
	generation, err := nftraits.CanonicalJSON(file.Generation)
	if err != nil {
		t.Fatalf("canonicalizing generation: %v", err)
	}
	up := cgstore.TokenTraitsUpsert{
		Seed:            enrichedSeed,
		SchemaMajor:     1,
		PipelineVersion: file.PipelineVersion,
		Attributes:      attributes,
		DescriptionArt:  file.DescriptionArt,
		Simulation:      simulation,
		Generation:      generation,
		ContentHash:     hash,
		SourceETag:      `"trait-v1"`,
		NextAttemptAt:   time.Unix(1767230000, 0),
	}
	if withManifest {
		up.Assets = testfixtures.AssetManifestExample
		up.ManifestETag = `"manifest-v1"`
	}
	if err := repo.UpsertTokenTraits(ctx, up); err != nil {
		t.Fatalf("storing the trait package: %v", err)
	}
}

// TestTraitIngestEndToEnd drives the whole path the way production does:
// a stub asset host publishes the package, the real ingest loop fetches and
// stores it, and the metadata route then serves the enriched document.
//
// It exists because the pieces can each be right while the seam between them
// is wrong. In particular the row must be keyed on the indexer's zero-padded
// spelling of the seed and not the generator's, which the vendored fixture
// writes unpadded — a mismatch every unit test on either side would miss.
func TestTraitIngestEndToEnd(t *testing.T) {
	h := server(t)
	ctx := context.Background()
	repo := cgstore.NewRepo(h.store)
	lendHexSeed(t, h)

	assetHost := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/traits/" + enrichedSeed + ".json":
			w.Header().Set("ETag", `"trait-v1"`)
			_, _ = w.Write(testfixtures.NftTraitsExample)
		case "/asset-manifests/" + enrichedSeed + ".json":
			w.Header().Set("ETag", `"manifest-v1"`)
			_, _ = w.Write(testfixtures.AssetManifestExample)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer assetHost.Close()

	// A module of its own, wired the way cmd/apiserver wires one, so the
	// running loop and the nudge hook it publishes cannot leak into the
	// shared harness that every other test snapshots.
	cgAPI, err := cosmicgame.New(ctx, cosmicgame.Config{
		Store:            h.store,
		EthClient:        h.ethClient,
		RPCClient:        h.rpcClient,
		AdminAPIKey:      adminKey,
		TraitsSourceBase: assetHost.URL,
	})
	if err != nil {
		t.Fatalf("building the module: %v", err)
	}
	runCtx, cancel := context.WithCancel(ctx)
	done := cgAPI.StartTraitIngestion(runCtx, cosmicgame.TraitIngestConfig{
		SourceBase: assetHost.URL,
		Interval:   50 * time.Millisecond,
	})
	t.Cleanup(func() {
		cancel()
		select {
		case <-done:
		case <-time.After(10 * time.Second):
			t.Error("the trait ingest loop did not stop on cancellation")
		}
	})
	router := routes.New(h.store, routes.Options{CosmicGame: cgAPI})

	// Poll rather than sleep: the loop does real HTTP and real database
	// work, so the deadline is only an outer guard.
	deadline := time.Now().Add(20 * time.Second)
	var row cgstore.TokenTraitsRow
	for {
		row, err = repo.TokenTraits(ctx, enrichedSeed)
		if err == nil {
			break
		}
		if !errors.Is(err, store.ErrNotFound) {
			t.Fatalf("reading the ingested row: %v", err)
		}
		if time.Now().After(deadline) {
			t.Fatal("the ingest loop never stored the published package")
		}
		time.Sleep(20 * time.Millisecond)
	}
	if row.Seed != enrichedSeed {
		t.Fatalf("stored seed = %q, want the indexer's spelling %q", row.Seed, enrichedSeed)
	}
	if len(row.Assets) == 0 {
		t.Error("the asset manifest was not ingested alongside the traits")
	}

	req := httptest.NewRequest(http.MethodGet, "/api/cosmicgame/cst/metadata/1", nil)
	req.RemoteAddr = "10.66.66.66:4242"
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200: %s", w.Code, w.Body.String())
	}
	var doc struct {
		Description  string                     `json:"description"`
		ImageDetails map[string]any             `json:"image_details"`
		Properties   map[string]json.RawMessage `json:"properties"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("decoding the response: %v", err)
	}
	if !strings.HasPrefix(doc.Description, "Three bodies trace orbit ribbons") {
		t.Errorf("description = %q, want the ingested art sentence", doc.Description)
	}
	if len(doc.ImageDetails) == 0 {
		t.Error("image_details is missing after a manifest ingest")
	}
	if _, ok := doc.Properties["simulation"]; !ok {
		t.Error("properties.simulation is missing after ingest")
	}
}

// TestMetadataFallsBackOnUndecodableTraitRow covers the storage-bug path: a
// row whose art blocks cannot be decoded must degrade to the fallback rather
// than serve half a document or fail the request.
func TestMetadataFallsBackOnUndecodableTraitRow(t *testing.T) {
	h := server(t)
	ctx := context.Background()
	lendHexSeed(t, h)

	// Valid JSON, but not the array of attributes the handler expects.
	if _, err := h.db.ExecContext(ctx,
		`INSERT INTO cg_token_traits(
			seed,schema_major,pipeline_version,attributes,description_art,
			simulation,generation,content_hash
		) VALUES($1,1,'1.0.0','"not-an-array"'::JSONB,'art','{}'::JSONB,'{}'::JSONB,'hash')`,
		enrichedSeed,
	); err != nil {
		t.Fatalf("inserting the corrupt row: %v", err)
	}

	w := h.get(t, "/api/cosmicgame/cst/metadata/1")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200: %s", w.Code, w.Body.String())
	}
	var doc map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("decoding the response: %v", err)
	}
	if !strings.HasPrefix(doc["description"].(string), "Discover the unique attributes") {
		t.Errorf("description = %v, want the fallback sentence", doc["description"])
	}
	properties, _ := doc["properties"].(map[string]any)
	if _, ok := properties["simulation"]; ok {
		t.Error("an undecodable row still produced an enriched document")
	}
}

// TestMetadataOmitsAllocationWhenProvenanceIsAmbiguous covers the data
// inconsistency path: a token matching several prize families has no single
// truthful provenance, so the attribute is dropped rather than guessed.
func TestMetadataOmitsAllocationWhenProvenanceIsAmbiguous(t *testing.T) {
	h := server(t)
	ctx := context.Background()

	// Widening round 0's main prize makes it claim tokens 1-3, so token 2
	// matches both the main prize and its bidder raffle. cg_prize_claim has
	// insert and delete triggers but none on update, so this leaves no
	// aggregate behind.
	if _, err := h.db.ExecContext(ctx,
		"UPDATE cg_prize_claim SET num_cs_nfts=3 WHERE round_num=0",
	); err != nil {
		t.Fatalf("widening the main prize: %v", err)
	}
	t.Cleanup(func() {
		if _, err := h.db.ExecContext(ctx,
			"UPDATE cg_prize_claim SET num_cs_nfts=1 WHERE round_num=0",
		); err != nil {
			t.Errorf("restoring the main prize width: %v", err)
		}
	})

	w := h.get(t, "/api/cosmicgame/cst/metadata/2")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200: %s", w.Code, w.Body.String())
	}
	var doc struct {
		Attributes []struct {
			TraitType string `json:"trait_type"`
		} `json:"attributes"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("decoding the response: %v", err)
	}
	for _, attr := range doc.Attributes {
		if attr.TraitType == "Allocation" {
			t.Fatal("an ambiguous provenance still produced an Allocation attribute")
		}
	}
}

// TestCstMetadataEnriched pins the served document for a token whose art
// package has been ingested. It is the contract test for the merge: art
// facts verbatim and first, chain facts appended, no owner, no seed
// attribute.
func TestCstMetadataEnriched(t *testing.T) {
	h := server(t)
	withIngestedTraits(t, h, true)

	w := h.get(t, "/api/cosmicgame/cst/metadata/1")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200: %s", w.Code, w.Body.String())
	}
	if got := w.Header().Get("Cache-Control"); got != "public, max-age=300" {
		t.Errorf("Cache-Control = %q, want %q", got, "public, max-age=300")
	}
	if w.Header().Get("ETag") == "" {
		t.Error("the response carries no ETag for marketplaces to revalidate against")
	}

	compareGolden(t, "cst_metadata_enriched", response{
		Status:      w.Code,
		ContentType: contentTypeOf(w),
		Body:        canonicalJSON(t, w.Body.Bytes()),
	})

	var doc struct {
		MetadataVersion string `json:"metadata_version"`
		Description     string `json:"description"`
		Attributes      []struct {
			TraitType string `json:"trait_type"`
		} `json:"attributes"`
		Properties map[string]json.RawMessage `json:"properties"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("decoding the response: %v", err)
	}
	if doc.MetadataVersion != "2.0.0" {
		t.Errorf("metadata_version = %q, want 2.0.0", doc.MetadataVersion)
	}
	if !strings.HasPrefix(doc.Description, "Three bodies trace orbit ribbons") {
		t.Errorf("description = %q, want it to lead with the art sentence", doc.Description)
	}
	if !strings.Contains(doc.Description, "Imprinted in Round 0") {
		t.Errorf("description = %q, want the chain provenance sentence", doc.Description)
	}

	wantOrder := []string{
		"Structure", "Underlay", "Palette", "Spectral Class", "Mass Balance",
		"Fate", "Chaos", "Syzygies", // art, verbatim and first
		"Round", "Imprinted", "Allocation", // chain, appended
	}
	if len(doc.Attributes) != len(wantOrder) {
		t.Fatalf("len(attributes) = %d, want %d", len(doc.Attributes), len(wantOrder))
	}
	for i, want := range wantOrder {
		if doc.Attributes[i].TraitType != want {
			t.Errorf("attribute %d = %q, want %q", i, doc.Attributes[i].TraitType, want)
		}
	}
	for _, key := range []string{"seed", "token_id", "round_num", "simulation", "generation", "media"} {
		if _, ok := doc.Properties[key]; !ok {
			t.Errorf("properties is missing %q", key)
		}
	}
	if _, ok := doc.Properties["owner"]; ok {
		t.Error("properties still carries an owner")
	}
	if strings.Contains(w.Body.String(), `"owner"`) {
		t.Error("the served document mentions an owner")
	}
}

// TestCstMetadataEnrichedWithoutManifest pins the degraded-but-correct
// document: no assets.json means no detail objects, everything else intact.
func TestCstMetadataEnrichedWithoutManifest(t *testing.T) {
	h := server(t)
	withIngestedTraits(t, h, false)

	w := h.get(t, "/api/cosmicgame/cst/metadata/1")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200: %s", w.Code, w.Body.String())
	}
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("decoding the response: %v", err)
	}
	for _, key := range []string{"image_details", "animation_details"} {
		if _, ok := doc[key]; ok {
			t.Errorf("%s was served without a manifest", key)
		}
	}
	if _, ok := doc["attributes"]; !ok {
		t.Error("the art attributes were dropped along with the manifest")
	}
}

// TestMetadataHostDispatchServesEnriched proves the bare tokenURI route the
// NFT contract points at gets the same enriched document.
func TestMetadataHostDispatchServesEnriched(t *testing.T) {
	h := server(t)
	withIngestedTraits(t, h, true)

	w := h.do(t, request{path: "/metadata/1", host: "nfts.cosmicsignature.com"})
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200: %s", w.Code, w.Body.String())
	}
	var doc struct {
		Properties struct {
			Seed  string          `json:"seed"`
			Media map[string]any  `json:"media"`
			Sim   json.RawMessage `json:"simulation"`
		} `json:"properties"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("decoding the response: %v", err)
	}
	if doc.Properties.Seed != enrichedSeed {
		t.Errorf("properties.seed = %q, want %q", doc.Properties.Seed, enrichedSeed)
	}
	wantTraitSource := traitsSourceBase + "/traits/" + enrichedSeed + ".json"
	if got := doc.Properties.Media["trait_source"]; got != wantTraitSource {
		t.Errorf("media.trait_source = %v, want %q", got, wantTraitSource)
	}
}

// TestCstMetadataFallbackShape pins the document every token serves for the
// minutes to hours between its mint and the generator publishing its package.
func TestCstMetadataFallbackShape(t *testing.T) {
	h := server(t)

	w := h.get(t, "/api/cosmicgame/cst/metadata/2")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200: %s", w.Code, w.Body.String())
	}
	var doc map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
		t.Fatalf("decoding the response: %v", err)
	}
	if doc["metadata_version"] != "2.0.0" {
		t.Errorf("metadata_version = %v, want 2.0.0", doc["metadata_version"])
	}
	for _, key := range []string{"image_details", "animation_details"} {
		if _, ok := doc[key]; ok {
			t.Errorf("the fallback carries %s", key)
		}
	}
	properties, _ := doc["properties"].(map[string]any)
	if _, ok := properties["owner"]; ok {
		t.Error("the fallback carries properties.owner")
	}
	if len(properties) != 3 {
		t.Errorf("fallback properties = %v, want only seed, token_id and round_num", properties)
	}

	attributes, _ := doc["attributes"].([]any)
	var traits []string
	for _, raw := range attributes {
		attr, _ := raw.(map[string]any)
		traits = append(traits, attr["trait_type"].(string))
	}
	// Token 2 is the round 0 bidder-raffle mint.
	want := []string{"Round", "Imprinted", "Allocation"}
	if len(traits) != len(want) {
		t.Fatalf("fallback attributes = %v, want %v", traits, want)
	}
	for i := range want {
		if traits[i] != want[i] {
			t.Errorf("attribute %d = %q, want %q", i, traits[i], want[i])
		}
	}
	for _, attr := range attributes {
		if m, _ := attr.(map[string]any); strings.EqualFold(m["trait_type"].(string), "seed") {
			t.Error("the fallback still carries a seed attribute")
		}
	}
}

// TestMetadataAllocationPerPrizeFamily walks every prize path the fixture
// dataset mints through and checks the provenance the collection publishes.
func TestMetadataAllocationPerPrizeFamily(t *testing.T) {
	h := server(t)
	tests := []struct {
		tokenID int
		want    string
	}{
		{1, "Final Gesture"},
		{2, "Stellar Selection"},
		{3, "Anchored Selection"}, // RandomWalk-staker raffle
		{4, "Last CST Gesture"},
		{5, "Endurance Champion"},
		{6, "Chrono-Warrior"},
	}
	for _, tc := range tests {
		w := h.get(t, "/api/cosmicgame/cst/metadata/"+strconv.Itoa(tc.tokenID))
		if w.Code != http.StatusOK {
			t.Errorf("token %d: status = %d, want 200", tc.tokenID, w.Code)
			continue
		}
		var doc struct {
			Attributes []struct {
				TraitType string `json:"trait_type"`
				Value     any    `json:"value"`
			} `json:"attributes"`
		}
		if err := json.Unmarshal(w.Body.Bytes(), &doc); err != nil {
			t.Errorf("token %d: decoding: %v", tc.tokenID, err)
			continue
		}
		var got string
		for _, attr := range doc.Attributes {
			if attr.TraitType == "Allocation" {
				got, _ = attr.Value.(string)
			}
		}
		if got != tc.want {
			t.Errorf("token %d Allocation = %q, want %q", tc.tokenID, got, tc.want)
		}
	}
}
