//go:build integration

package cosmicgame

import (
	"context"
	"encoding/json"
	"errors"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/nftraits"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// Seeds used by the trait tests. They are hexadecimal because the trait
// tables constrain their key to hex, matching what the indexer writes
// (NftMinted carries a uint256, formatted as 64 hex characters).
const (
	traitSeedA = "0x0000000000000000000000000000000000000000000000000000000000100033"
	traitSeedB = "0x00000000000000000000000000000000000000000000000000000000001000ff"
)

// hexSeedFor temporarily gives one fixture token a hexadecimal seed and
// restores the original afterwards.
//
// The shared dataset deliberately uses non-hex placeholder seeds, and
// changing them would churn forty golden files for no behavioural reason.
// An UPDATE is safe here: on_mint_update is a no-op, so no aggregate moves.
func hexSeedFor(t *testing.T, tokenID int64, seed string) {
	t.Helper()
	ctx := context.Background()
	r := repo(t)

	var original string
	if err := r.q(ctx).QueryRow(ctx,
		"SELECT seed FROM cg_mint_event WHERE token_id=$1", tokenID,
	).Scan(&original); err != nil {
		t.Fatalf("reading the original seed of token %d: %v", tokenID, err)
	}
	bare := seed[2:] // the column stores the seed without the 0x prefix
	if _, err := r.q(ctx).Exec(ctx,
		"UPDATE cg_mint_event SET seed=$1 WHERE token_id=$2", bare, tokenID,
	); err != nil {
		t.Fatalf("setting a hex seed on token %d: %v", tokenID, err)
	}
	t.Cleanup(func() {
		if _, err := r.q(ctx).Exec(ctx,
			"UPDATE cg_mint_event SET seed=$1 WHERE token_id=$2", original, tokenID,
		); err != nil {
			t.Errorf("restoring the seed of token %d: %v", tokenID, err)
		}
	})
}

// cleanTraits removes any trait rows for the given seeds before and after a
// test, so cases stay independent of execution order.
func cleanTraits(t *testing.T, seeds ...string) {
	t.Helper()
	ctx := context.Background()
	r := repo(t)
	remove := func() {
		for _, seed := range seeds {
			if _, err := r.q(ctx).Exec(ctx, "DELETE FROM cg_token_traits WHERE seed=$1", seed); err != nil {
				t.Errorf("clearing cg_token_traits for %s: %v", seed, err)
			}
			if _, err := r.q(ctx).Exec(ctx, "DELETE FROM cg_token_traits_fetch WHERE seed=$1", seed); err != nil {
				t.Errorf("clearing cg_token_traits_fetch for %s: %v", seed, err)
			}
		}
	}
	remove()
	t.Cleanup(remove)
}

// exampleUpsert builds a storable payload from the vendored generator
// fixture, keyed on traitSeedA.
func exampleUpsert(t *testing.T) TokenTraitsUpsert {
	t.Helper()
	const seed = traitSeedA
	file, err := nftraits.Parse(testfixtures.NftTraitsExample)
	if err != nil {
		t.Fatalf("parsing the vendored fixture: %v", err)
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
	return TokenTraitsUpsert{
		Seed:            seed,
		SchemaMajor:     1,
		PipelineVersion: file.PipelineVersion,
		Attributes:      attributes,
		DescriptionArt:  file.DescriptionArt,
		Simulation:      simulation,
		Generation:      generation,
		Assets:          testfixtures.AssetManifestExample,
		ContentHash:     hash,
		SourceETag:      `"trait-v1"`,
		ManifestETag:    `"manifest-v1"`,
		NextAttemptAt:   time.Now().Add(24 * time.Hour),
	}
}

func TestTokenTraitsRoundTrip(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cleanTraits(t, traitSeedA)

	up := exampleUpsert(t)
	if err := r.UpsertTokenTraits(ctx, up); err != nil {
		t.Fatalf("UpsertTokenTraits: %v", err)
	}

	row, err := r.TokenTraits(ctx, traitSeedA)
	if err != nil {
		t.Fatalf("TokenTraits: %v", err)
	}
	if row.Seed != traitSeedA {
		t.Errorf("seed = %q, want %q", row.Seed, traitSeedA)
	}
	if row.SchemaMajor != 1 || row.PipelineVersion != up.PipelineVersion {
		t.Errorf("schema/pipeline = %d/%q, want 1/%q", row.SchemaMajor, row.PipelineVersion, up.PipelineVersion)
	}
	if row.DescriptionArt != up.DescriptionArt {
		t.Errorf("description_art = %q, want %q", row.DescriptionArt, up.DescriptionArt)
	}
	if row.ContentHash != up.ContentHash {
		t.Errorf("content_hash = %q, want %q", row.ContentHash, up.ContentHash)
	}

	// JSONB normalizes whitespace, so compare decoded values rather than bytes.
	var stored, want []nftraits.Attribute
	if err := json.Unmarshal(row.Attributes, &stored); err != nil {
		t.Fatalf("decoding stored attributes: %v", err)
	}
	if err := json.Unmarshal(up.Attributes, &want); err != nil {
		t.Fatalf("decoding source attributes: %v", err)
	}
	if len(stored) != len(want) {
		t.Fatalf("stored %d attributes, want %d", len(stored), len(want))
	}
	for i := range want {
		if stored[i].TraitType != want[i].TraitType || string(stored[i].Value) != string(want[i].Value) {
			t.Errorf("attribute %d = %+v, want %+v", i, stored[i], want[i])
		}
	}

	var simulation map[string]any
	if err := json.Unmarshal(row.Simulation, &simulation); err != nil {
		t.Fatalf("decoding stored simulation: %v", err)
	}
	if simulation["chaos_index"] != float64(30) {
		t.Errorf("simulation.chaos_index = %v, want 30", simulation["chaos_index"])
	}

	manifest, err := nftraits.ParseManifest(row.Assets)
	if err != nil {
		t.Fatalf("decoding stored assets: %v", err)
	}
	if manifest.EntryByPath("images/source/master.png") == nil {
		t.Error("the stored manifest lost its master image entry")
	}
}

func TestTokenTraitsNotFound(t *testing.T) {
	r := repo(t)
	if _, err := r.TokenTraits(context.Background(), "0xdeadbeef"); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("TokenTraits() error = %v, want store.ErrNotFound", err)
	}
}

// TestUpsertTokenTraitsPreservesManifest covers the independent publication
// of the two files: a trait re-ingest whose manifest request failed must not
// cost the token its image_details.
func TestUpsertTokenTraitsPreservesManifest(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cleanTraits(t, traitSeedA)
	hexSeedFor(t, 1, traitSeedA)

	if err := r.UpsertTokenTraits(ctx, exampleUpsert(t)); err != nil {
		t.Fatalf("first upsert: %v", err)
	}

	second := exampleUpsert(t)
	second.Assets = nil
	second.ManifestETag = ""
	second.PipelineVersion = "1.1.0"
	if err := r.UpsertTokenTraits(ctx, second); err != nil {
		t.Fatalf("second upsert: %v", err)
	}

	row, err := r.TokenTraits(ctx, traitSeedA)
	if err != nil {
		t.Fatalf("TokenTraits: %v", err)
	}
	if len(row.Assets) == 0 {
		t.Error("a manifest-less re-ingest dropped the stored manifest")
	}
	if row.PipelineVersion != "1.1.0" {
		t.Errorf("pipeline_version = %q, want the refreshed 1.1.0", row.PipelineVersion)
	}

	candidate, err := r.TraitFetchCandidate(ctx, traitSeedA)
	if err != nil {
		t.Fatalf("TraitFetchCandidate: %v", err)
	}
	if candidate.ManifestETag != `"manifest-v1"` {
		t.Errorf("manifest_etag = %q, want the preserved %q", candidate.ManifestETag, `"manifest-v1"`)
	}
}

func TestUpsertTokenTraitsRejectsEmptySeed(t *testing.T) {
	r := repo(t)
	if err := r.UpsertTokenTraits(context.Background(), TokenTraitsUpsert{}); err == nil {
		t.Fatal("UpsertTokenTraits accepted an empty seed")
	}
}

// TestSeedsDueForTraitFetchExcludesNonHexSeeds proves the SQL filter: the
// shared dataset's placeholder seeds cannot address a package and must never
// reach the trait tables, whose keys are constrained to hex.
func TestSeedsDueForTraitFetchExcludesNonHexSeeds(t *testing.T) {
	r := repo(t)
	due, err := r.SeedsDueForTraitFetch(context.Background(), 100)
	if err != nil {
		t.Fatalf("SeedsDueForTraitFetch: %v", err)
	}
	if len(due) != 0 {
		t.Fatalf("the placeholder-seed fixture produced %d candidates: %+v", len(due), due)
	}
}

func TestSeedsDueForTraitFetchRejectsBadLimit(t *testing.T) {
	r := repo(t)
	if _, err := r.SeedsDueForTraitFetch(context.Background(), 0); err == nil {
		t.Fatal("SeedsDueForTraitFetch accepted a zero limit")
	}
	if _, err := r.DriftedTraitSeeds(context.Background(), -1); err == nil {
		t.Fatal("DriftedTraitSeeds accepted a negative limit")
	}
}

func TestSeedsDueForTraitFetchSchedule(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cleanTraits(t, traitSeedA, traitSeedB)
	hexSeedFor(t, 1, traitSeedA)
	hexSeedFor(t, 2, traitSeedB)

	seedsOf := func(candidates []TraitFetchCandidate) []string {
		out := make([]string, len(candidates))
		for i, c := range candidates {
			out[i] = c.Seed
		}
		slices.Sort(out)
		return out
	}

	// Never-attempted seeds are due immediately.
	due, err := r.SeedsDueForTraitFetch(ctx, 100)
	if err != nil {
		t.Fatalf("SeedsDueForTraitFetch: %v", err)
	}
	if got := seedsOf(due); !slices.Equal(got, []string{traitSeedA, traitSeedB}) {
		t.Fatalf("due seeds = %v, want both hex seeds", got)
	}
	for _, c := range due {
		if c.AttemptCount != 0 || c.HasTraits || c.ContentHash != "" || !c.DueAt.IsZero() {
			t.Errorf("a never-attempted candidate carries state: %+v", c)
		}
	}

	// A scheduled backoff removes the seed from the scan.
	future := time.Now().Add(time.Hour)
	if err := r.RecordTraitFetchOutcome(ctx, traitSeedA, TraitFetchMissing, nil, future); err != nil {
		t.Fatalf("RecordTraitFetchOutcome: %v", err)
	}
	due, err = r.SeedsDueForTraitFetch(ctx, 100)
	if err != nil {
		t.Fatalf("SeedsDueForTraitFetch: %v", err)
	}
	if got := seedsOf(due); !slices.Equal(got, []string{traitSeedB}) {
		t.Fatalf("due seeds after a backoff = %v, want only %s", got, traitSeedB)
	}

	// A lapsed schedule brings it back, carrying the attempt count that
	// sizes the next backoff.
	past := time.Now().Add(-time.Hour)
	if err := r.RecordTraitFetchOutcome(ctx, traitSeedA, TraitFetchMissing, nil, past); err != nil {
		t.Fatalf("RecordTraitFetchOutcome: %v", err)
	}
	due, err = r.SeedsDueForTraitFetch(ctx, 100)
	if err != nil {
		t.Fatalf("SeedsDueForTraitFetch: %v", err)
	}
	// The lapsed seed sorts ahead of the never-attempted one only if its
	// schedule is older than the sentinel; never-attempted wins by design.
	if got := seedsOf(due); !slices.Equal(got, []string{traitSeedA, traitSeedB}) {
		t.Fatalf("due seeds after the schedule lapsed = %v, want both", got)
	}
	for _, c := range due {
		if c.Seed != traitSeedA {
			continue
		}
		if c.AttemptCount != 2 {
			t.Errorf("attempt_count = %d, want 2 after two recorded attempts", c.AttemptCount)
		}
		if c.DueAt.IsZero() {
			t.Error("a scheduled candidate reports no due time")
		}
	}

	// The limit bounds the page.
	page, err := r.SeedsDueForTraitFetch(ctx, 1)
	if err != nil {
		t.Fatalf("SeedsDueForTraitFetch: %v", err)
	}
	if len(page) != 1 {
		t.Errorf("len(page) = %d, want the requested 1", len(page))
	}
}

// TestSeedsDueForTraitFetchOrdersNeverAttemptedFirst keeps a cold start
// filling gaps before it re-checks packages it already has.
func TestSeedsDueForTraitFetchOrdersNeverAttemptedFirst(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cleanTraits(t, traitSeedA, traitSeedB)
	hexSeedFor(t, 1, traitSeedA)
	hexSeedFor(t, 2, traitSeedB)

	if err := r.RecordTraitFetchOutcome(ctx, traitSeedB, TraitFetchMissing, nil, time.Now().Add(-time.Minute)); err != nil {
		t.Fatalf("RecordTraitFetchOutcome: %v", err)
	}
	due, err := r.SeedsDueForTraitFetch(ctx, 100)
	if err != nil {
		t.Fatalf("SeedsDueForTraitFetch: %v", err)
	}
	if len(due) != 2 {
		t.Fatalf("len(due) = %d, want 2", len(due))
	}
	if due[0].Seed != traitSeedA {
		t.Errorf("first candidate = %s, want the never-attempted %s", due[0].Seed, traitSeedA)
	}
}

func TestTraitFetchCandidateReflectsStoredTraits(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cleanTraits(t, traitSeedA)
	hexSeedFor(t, 1, traitSeedA)

	up := exampleUpsert(t)
	if err := r.UpsertTokenTraits(ctx, up); err != nil {
		t.Fatalf("UpsertTokenTraits: %v", err)
	}

	candidate, err := r.TraitFetchCandidate(ctx, traitSeedA)
	if err != nil {
		t.Fatalf("TraitFetchCandidate: %v", err)
	}
	if !candidate.HasTraits {
		t.Error("HasTraits is false for a stored package")
	}
	if candidate.ContentHash != up.ContentHash {
		t.Errorf("content_hash = %q, want %q", candidate.ContentHash, up.ContentHash)
	}
	if candidate.SourceETag != up.SourceETag {
		t.Errorf("source_etag = %q, want %q", candidate.SourceETag, up.SourceETag)
	}
	if candidate.AttemptCount != 0 {
		t.Errorf("attempt_count = %d, want a successful upsert to reset it", candidate.AttemptCount)
	}
	if candidate.DueAt.IsZero() {
		t.Error("a stored package has no drift re-check scheduled")
	}
}

func TestTraitFetchCandidateUnknownSeed(t *testing.T) {
	r := repo(t)
	_, err := r.TraitFetchCandidate(context.Background(), "0xfeedface")
	if !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("TraitFetchCandidate() error = %v, want store.ErrNotFound", err)
	}
}

func TestRecordTraitFetchOutcomeAccumulates(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cleanTraits(t, traitSeedA)
	hexSeedFor(t, 1, traitSeedA)

	next := time.Now().Add(-time.Minute)
	for range 3 {
		if err := r.RecordTraitFetchOutcome(ctx, traitSeedA, TraitFetchTransient, errors.New("boom"), next); err != nil {
			t.Fatalf("RecordTraitFetchOutcome: %v", err)
		}
	}
	candidate, err := r.TraitFetchCandidate(ctx, traitSeedA)
	if err != nil {
		t.Fatalf("TraitFetchCandidate: %v", err)
	}
	if candidate.AttemptCount != 3 {
		t.Errorf("attempt_count = %d, want 3", candidate.AttemptCount)
	}

	var status, lastError string
	if err := r.q(ctx).QueryRow(ctx,
		"SELECT last_status, COALESCE(last_error,'') FROM cg_token_traits_fetch WHERE seed=$1", traitSeedA,
	).Scan(&status, &lastError); err != nil {
		t.Fatalf("reading the fetch row: %v", err)
	}
	if status != string(TraitFetchTransient) || lastError != "boom" {
		t.Errorf("status/error = %q/%q, want %q/boom", status, lastError, TraitFetchTransient)
	}

	// A success clears the accumulated failure state.
	if err := r.UpsertTokenTraits(ctx, exampleUpsert(t)); err != nil {
		t.Fatalf("UpsertTokenTraits: %v", err)
	}
	if err := r.q(ctx).QueryRow(ctx,
		"SELECT last_status, COALESCE(last_error,'') FROM cg_token_traits_fetch WHERE seed=$1", traitSeedA,
	).Scan(&status, &lastError); err != nil {
		t.Fatalf("reading the fetch row: %v", err)
	}
	if status != string(TraitFetchOK) || lastError != "" {
		t.Errorf("status/error after success = %q/%q, want ok/empty", status, lastError)
	}
}

func TestRecordTraitFetchOutcomeTruncatesLongErrors(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cleanTraits(t, traitSeedA)

	long := make([]byte, 4096)
	for i := range long {
		long[i] = 'x'
	}
	if err := r.RecordTraitFetchOutcome(ctx, traitSeedA, TraitFetchTransient,
		errors.New(string(long)), time.Now()); err != nil {
		t.Fatalf("RecordTraitFetchOutcome: %v", err)
	}
	var stored string
	if err := r.q(ctx).QueryRow(ctx,
		"SELECT last_error FROM cg_token_traits_fetch WHERE seed=$1", traitSeedA,
	).Scan(&stored); err != nil {
		t.Fatalf("reading the fetch row: %v", err)
	}
	if len(stored) != maxTraitFetchErrorLen {
		t.Errorf("len(last_error) = %d, want it truncated to %d", len(stored), maxTraitFetchErrorLen)
	}
}

func TestRecordTraitFetchOutcomeRejectsEmptySeed(t *testing.T) {
	r := repo(t)
	if err := r.RecordTraitFetchOutcome(context.Background(), "", TraitFetchMissing, nil, time.Now()); err == nil {
		t.Fatal("RecordTraitFetchOutcome accepted an empty seed")
	}
	if err := r.RecordTraitDrift(context.Background(), "", "1.0.0", time.Now()); err == nil {
		t.Fatal("RecordTraitDrift accepted an empty seed")
	}
}

// TestRecordTraitDriftLeavesArtFactsAlone is the frozen-art guarantee at the
// storage layer: raising drift must never touch what collectors already saw.
func TestRecordTraitDriftLeavesArtFactsAlone(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	cleanTraits(t, traitSeedA)

	up := exampleUpsert(t)
	if err := r.UpsertTokenTraits(ctx, up); err != nil {
		t.Fatalf("UpsertTokenTraits: %v", err)
	}
	if err := r.RecordTraitDrift(ctx, traitSeedA, "9.9.9", time.Now().Add(time.Hour)); err != nil {
		t.Fatalf("RecordTraitDrift: %v", err)
	}

	row, err := r.TokenTraits(ctx, traitSeedA)
	if err != nil {
		t.Fatalf("TokenTraits: %v", err)
	}
	if row.ContentHash != up.ContentHash || row.PipelineVersion != up.PipelineVersion {
		t.Error("drift overwrote the stored art facts")
	}

	var driftAt time.Time
	var driftVersion string
	if err := r.q(ctx).QueryRow(ctx,
		"SELECT drift_detected_at, drift_pipeline_version FROM cg_token_traits_fetch WHERE seed=$1", traitSeedA,
	).Scan(&driftAt, &driftVersion); err != nil {
		t.Fatalf("reading the drift columns: %v", err)
	}
	if driftAt.IsZero() || driftVersion != "9.9.9" {
		t.Errorf("drift columns = %v/%q, want a timestamp and 9.9.9", driftAt, driftVersion)
	}

	drifted, err := r.DriftedTraitSeeds(ctx, 10)
	if err != nil {
		t.Fatalf("DriftedTraitSeeds: %v", err)
	}
	if !slices.Contains(drifted, traitSeedA) {
		t.Errorf("DriftedTraitSeeds = %v, want it to list %s", drifted, traitSeedA)
	}

	// The first observation is the one operators investigate, so a repeat
	// must not move the timestamp.
	if err := r.RecordTraitDrift(ctx, traitSeedA, "9.9.10", time.Now().Add(time.Hour)); err != nil {
		t.Fatalf("second RecordTraitDrift: %v", err)
	}
	var secondAt time.Time
	if err := r.q(ctx).QueryRow(ctx,
		"SELECT drift_detected_at FROM cg_token_traits_fetch WHERE seed=$1", traitSeedA,
	).Scan(&secondAt); err != nil {
		t.Fatalf("re-reading the drift columns: %v", err)
	}
	if !secondAt.Equal(driftAt) {
		t.Errorf("drift_detected_at moved from %v to %v", driftAt, secondAt)
	}
}

func TestDriftedTraitSeedsEmptyByDefault(t *testing.T) {
	r := repo(t)
	drifted, err := r.DriftedTraitSeeds(context.Background(), 10)
	if err != nil {
		t.Fatalf("DriftedTraitSeeds: %v", err)
	}
	if len(drifted) != 0 {
		t.Errorf("DriftedTraitSeeds = %v, want none in a healthy fixture", drifted)
	}
}

// TestCosmicSignatureMintSource pins the provenance of every prize family
// the fixture dataset mints a token through.
func TestCosmicSignatureMintSource(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	tests := []struct {
		tokenID int64
		want    CosmicSignatureMintSource
	}{
		{1, MintSourceMainPrize},
		{2, MintSourceBidderRaffle},
		{3, MintSourceRandomWalkStaker},
		{4, MintSourceLastCstBidder},
		{5, MintSourceEnduranceChampion},
		{6, MintSourceChronoWarriorPrize},
	}
	for _, tc := range tests {
		got, err := r.CosmicSignatureMintSource(ctx, tc.tokenID)
		if err != nil {
			t.Errorf("CosmicSignatureMintSource(%d): %v", tc.tokenID, err)
			continue
		}
		if got != tc.want {
			t.Errorf("CosmicSignatureMintSource(%d) = %q, want %q", tc.tokenID, got, tc.want)
		}
	}
}

func TestCosmicSignatureMintSourceUnknownToken(t *testing.T) {
	r := repo(t)
	_, err := r.CosmicSignatureMintSource(context.Background(), 999999)
	if !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("CosmicSignatureMintSource() error = %v, want store.ErrNotFound", err)
	}
}

// TestCosmicSignatureMintSourceAmbiguous covers the exactly-one-source rule.
// Widening round 0's main prize makes it claim tokens 1-3, so token 2 matches
// both the main prize and its bidder raffle — a data inconsistency the query
// must report rather than resolve arbitrarily. cg_prize_claim has insert and
// delete triggers but none on update, so this leaves no aggregate behind.
func TestCosmicSignatureMintSourceAmbiguous(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	if _, err := r.q(ctx).Exec(ctx, "UPDATE cg_prize_claim SET num_cs_nfts=3 WHERE round_num=0"); err != nil {
		t.Fatalf("widening the main prize: %v", err)
	}
	t.Cleanup(func() {
		if _, err := r.q(ctx).Exec(ctx, "UPDATE cg_prize_claim SET num_cs_nfts=1 WHERE round_num=0"); err != nil {
			t.Errorf("restoring the main prize width: %v", err)
		}
	})

	_, err := r.CosmicSignatureMintSource(ctx, 2)
	if err == nil {
		t.Fatal("an ambiguous token resolved to a single mint source")
	}
	if !strings.Contains(err.Error(), "2 mint sources") {
		t.Errorf("error = %v, want it to name the ambiguity", err)
	}
}

// TestTokenTraitsQueriesReportDatabaseFailures pins the error wrapping on
// every query in this file. A cancelled context is the cheapest faithful
// failure: pgx refuses to run and the method must surface that rather than
// return a zero value the caller would treat as real data.
func TestTokenTraitsQueriesReportDatabaseFailures(t *testing.T) {
	r := repo(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tests := map[string]func() error{
		"TokenTraits": func() error {
			_, err := r.TokenTraits(ctx, traitSeedA)
			return err
		},
		"SeedsDueForTraitFetch": func() error {
			_, err := r.SeedsDueForTraitFetch(ctx, 10)
			return err
		},
		"TraitFetchCandidate": func() error {
			_, err := r.TraitFetchCandidate(ctx, traitSeedA)
			return err
		},
		"UpsertTokenTraits": func() error {
			return r.UpsertTokenTraits(ctx, exampleUpsert(t))
		},
		"RecordTraitFetchOutcome": func() error {
			return r.RecordTraitFetchOutcome(ctx, traitSeedA, TraitFetchMissing, nil, time.Now())
		},
		"RecordTraitDrift": func() error {
			return r.RecordTraitDrift(ctx, traitSeedA, "1.0.0", time.Now())
		},
		"DriftedTraitSeeds": func() error {
			_, err := r.DriftedTraitSeeds(ctx, 10)
			return err
		},
		"CosmicSignatureMintSource": func() error {
			_, err := r.CosmicSignatureMintSource(ctx, 1)
			return err
		},
	}
	for name, call := range tests {
		t.Run(name, func(t *testing.T) {
			if err := call(); err == nil {
				t.Fatal("a cancelled context produced no error")
			} else if !errors.Is(err, context.Canceled) {
				t.Errorf("error = %v, want it to wrap context.Canceled", err)
			}
		})
	}
}
