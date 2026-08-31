package traits

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/nftraits"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

const (
	testSeed       = "0x100033"
	testPaddedSeed = "0x0000000000000000000000000000000000000000000000000000000000100033"
)

// fakeStore records what the ingester decided, so tests assert on outcomes
// rather than on database state.
type fakeStore struct {
	mu sync.Mutex

	due        []cgstore.TraitFetchCandidate
	bySeed     map[string]cgstore.TraitFetchCandidate
	dueErr     error
	lookupErr  error
	upsertErr  error
	upserts    []cgstore.TokenTraitsUpsert
	outcomes   []recordedOutcome
	drifts     []recordedDrift
	dueCalls   int
	scanCursor int

	// scanned and stored let real-time loop tests wait on progress instead
	// of sleeping. They are buffered and best-effort.
	scanned chan struct{}
	stored  chan struct{}
	// onOutcome fires after each recorded outcome, letting a test interrupt
	// the scan at a known point.
	onOutcome func()
}

type recordedOutcome struct {
	Seed        string
	Status      cgstore.TraitFetchStatus
	Cause       error
	NextAttempt time.Time
}

type recordedDrift struct {
	Seed            string
	PipelineVersion string
}

func newFakeStore(candidates ...cgstore.TraitFetchCandidate) *fakeStore {
	bySeed := make(map[string]cgstore.TraitFetchCandidate, len(candidates))
	for _, c := range candidates {
		bySeed[c.Seed] = c
	}
	return &fakeStore{
		due:     candidates,
		bySeed:  bySeed,
		scanned: make(chan struct{}, 64),
		stored:  make(chan struct{}, 64),
	}
}

// signal delivers a progress notification without ever blocking the loop
// under test.
func signal(ch chan struct{}) {
	select {
	case ch <- struct{}{}:
	default:
	}
}

// awaitSignal waits for one progress notification, failing the test if the
// loop stalls. The deadline is an outer guard around real network I/O, not a
// timing assertion.
func awaitSignal(tb testing.TB, ch chan struct{}, what string) {
	tb.Helper()
	select {
	case <-ch:
	case <-time.After(10 * time.Second):
		tb.Fatalf("timed out waiting for %s", what)
	}
}

// SeedsDueForTraitFetch serves the pending candidates one page at a time and
// then reports exhaustion, mirroring the real scan's drain semantics.
func (f *fakeStore) SeedsDueForTraitFetch(_ context.Context, limit int) ([]cgstore.TraitFetchCandidate, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.dueCalls++
	signal(f.scanned)
	if f.dueErr != nil {
		return nil, f.dueErr
	}
	if f.scanCursor >= len(f.due) {
		return nil, nil
	}
	end := min(f.scanCursor+limit, len(f.due))
	page := f.due[f.scanCursor:end]
	f.scanCursor = end
	return page, nil
}

func (f *fakeStore) TraitFetchCandidate(_ context.Context, seed string) (cgstore.TraitFetchCandidate, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.lookupErr != nil {
		return cgstore.TraitFetchCandidate{}, f.lookupErr
	}
	candidate, ok := f.bySeed[seed]
	if !ok {
		return cgstore.TraitFetchCandidate{}, store.ErrNotFound
	}
	return candidate, nil
}

func (f *fakeStore) UpsertTokenTraits(_ context.Context, up cgstore.TokenTraitsUpsert) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.upsertErr != nil {
		return f.upsertErr
	}
	f.upserts = append(f.upserts, up)
	signal(f.stored)
	return nil
}

func (f *fakeStore) RecordTraitFetchOutcome(
	_ context.Context,
	seed string,
	status cgstore.TraitFetchStatus,
	cause error,
	nextAttempt time.Time,
) error {
	f.mu.Lock()
	f.outcomes = append(f.outcomes, recordedOutcome{seed, status, cause, nextAttempt})
	hook := f.onOutcome
	f.mu.Unlock()
	if hook != nil {
		hook()
	}
	return nil
}

func (f *fakeStore) RecordTraitDrift(_ context.Context, seed, pipelineVersion string, _ time.Time) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.drifts = append(f.drifts, recordedDrift{seed, pipelineVersion})
	return nil
}

func (f *fakeStore) snapshot() ([]cgstore.TokenTraitsUpsert, []recordedOutcome, []recordedDrift) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]cgstore.TokenTraitsUpsert(nil), f.upserts...),
		append([]recordedOutcome(nil), f.outcomes...),
		append([]recordedDrift(nil), f.drifts...)
}

// assetHost is a stand-in for the nginx locations the generator's packages
// are published behind.
type assetHost struct {
	traits    map[string]string
	manifests map[string]string
	etag      string
	traitsHits,
	manifestHits int
	mu sync.Mutex
}

func newAssetHost() *assetHost {
	return &assetHost{traits: map[string]string{}, manifests: map[string]string{}}
}

func (h *assetHost) serve(tb testing.TB) *httptest.Server {
	tb.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var table map[string]string
		switch {
		case strings.HasSuffix(r.URL.Path, "/metadata/nft_traits.json"):
			h.mu.Lock()
			h.traitsHits++
			h.mu.Unlock()
			table = h.traits
		case strings.HasSuffix(r.URL.Path, "/metadata/assets.json"):
			h.mu.Lock()
			h.manifestHits++
			h.mu.Unlock()
			table = h.manifests
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Both files live inside the seed's package directory, so the
		// seed is the leading path segment: /{seed}/metadata/{file}.json.
		name, _, _ := strings.Cut(strings.TrimPrefix(r.URL.Path, "/"), "/")
		body, ok := table[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if h.etag != "" {
			if r.Header.Get("If-None-Match") == h.etag {
				w.WriteHeader(http.StatusNotModified)
				return
			}
			w.Header().Set("ETag", h.etag)
		}
		_, _ = w.Write([]byte(body))
	}))
	tb.Cleanup(srv.Close)
	return srv
}

func (h *assetHost) hits() (traitsHits, manifestHits int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.traitsHits, h.manifestHits
}

// newTestIngester wires an ingester against a fake store and a stub asset
// host with a frozen clock.
func newTestIngester(tb testing.TB, st Store, base string) *Ingester {
	tb.Helper()
	frozen := time.Date(2026, 8, 29, 12, 0, 0, 0, time.UTC)
	ing, err := New(Config{
		Store:      st,
		SourceBase: base,
		Interval:   time.Minute,
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
		Now:        func() time.Time { return frozen },
	})
	if err != nil {
		tb.Fatalf("New: %v", err)
	}
	return ing
}

func TestNewValidatesConfig(t *testing.T) {
	t.Parallel()
	if _, err := New(Config{SourceBase: "https://example.com"}); err == nil {
		t.Error("New accepted a nil store")
	}
	if _, err := New(Config{Store: newFakeStore()}); err == nil {
		t.Error("New accepted an empty source base")
	}
	if _, err := New(Config{Store: newFakeStore(), SourceBase: "not a url"}); err == nil {
		t.Error("New accepted a malformed source base")
	}
	ing, err := New(Config{Store: newFakeStore(), SourceBase: "https://example.com"})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if ing.interval != defaultInterval || ing.batchSize != defaultBatchSize ||
		ing.maxPerTick != defaultMaxPerTick || ing.recheckInterval != defaultRecheck {
		t.Errorf("defaults not applied: %+v", ing)
	}
	if ing.logger == nil || ing.now == nil {
		t.Error("New left the logger or clock nil")
	}
}

// TestIngestStoresGatedPackage is the happy path: a published package lands
// in the store with its art facts, its manifest and its validators.
func TestIngestStoresGatedPackage(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	host.manifests[testSeed] = string(testfixtures.AssetManifestExample)
	host.etag = `"trait-v1"`
	srv := host.serve(t)

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	got := ing.ingestSeed(t.Context(), cgstore.TraitFetchCandidate{Seed: testSeed})
	if got != outcomeStored {
		t.Fatalf("ingestSeed = %v, want outcomeStored", got)
	}

	upserts, outcomes, drifts := st.snapshot()
	if len(upserts) != 1 {
		t.Fatalf("len(upserts) = %d, want 1", len(upserts))
	}
	if len(outcomes) != 0 || len(drifts) != 0 {
		t.Errorf("a stored package also recorded outcomes=%v drifts=%v", outcomes, drifts)
	}

	up := upserts[0]
	if up.Seed != testSeed {
		t.Errorf("seed = %q, want %q", up.Seed, testSeed)
	}
	if up.SchemaMajor != nftraits.SupportedSchemaMajor {
		t.Errorf("schema_major = %d, want %d", up.SchemaMajor, nftraits.SupportedSchemaMajor)
	}
	if up.PipelineVersion != "1.0.0" {
		t.Errorf("pipeline_version = %q, want 1.0.0", up.PipelineVersion)
	}
	if up.DescriptionArt == "" {
		t.Error("description_art is empty")
	}
	if up.ContentHash == "" {
		t.Error("content_hash is empty")
	}
	if up.SourceETag != `"trait-v1"` || up.ManifestETag != `"trait-v1"` {
		t.Errorf("validators = (%q, %q), want both %q", up.SourceETag, up.ManifestETag, `"trait-v1"`)
	}
	if len(up.Assets) == 0 {
		t.Error("the manifest was not stored")
	}

	var attributes []nftraits.Attribute
	if err := json.Unmarshal(up.Attributes, &attributes); err != nil {
		t.Fatalf("stored attributes are not decodable: %v", err)
	}
	if len(attributes) != 8 || attributes[0].TraitType != "Structure" {
		t.Errorf("stored attributes = %+v, want the fixture's 8 starting with Structure", attributes)
	}
	// The stored blocks must be canonical so re-ingesting identical content
	// produces byte-identical rows.
	if !json.Valid(up.Simulation) || !json.Valid(up.Generation) {
		t.Error("stored simulation/generation blocks are not valid JSON")
	}
	if strings.Contains(string(up.Simulation), "\n") {
		t.Error("stored simulation block was not canonicalized")
	}
}

// TestIngestAcceptsPaddedIndexerSeed covers the real interop hazard: the
// indexer stores a uint256 zero-padded to 64 hex characters while the
// generator writes the seed it was given. The vendored fixture says
// "0x100033"; the indexer says the padded form.
//
// The row must be keyed on the indexer's spelling, because that is what the
// metadata handler derives from cg_mint_event and looks up. Keying on the
// file's spelling would store packages no token can ever find.
func TestIngestAcceptsPaddedIndexerSeed(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testPaddedSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	if got := ing.ingestSeed(t.Context(), cgstore.TraitFetchCandidate{Seed: testPaddedSeed}); got != outcomeStored {
		t.Fatalf("ingestSeed = %v, want outcomeStored", got)
	}
	upserts, _, _ := st.snapshot()
	if len(upserts) != 1 {
		t.Fatalf("len(upserts) = %d, want 1", len(upserts))
	}
	if upserts[0].Seed != testPaddedSeed {
		t.Errorf("stored seed = %q, want the indexer's spelling %q", upserts[0].Seed, testPaddedSeed)
	}
}

func TestIngestOutcomes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		body        string
		serve       bool
		etag        string
		candidate   cgstore.TraitFetchCandidate
		wantOutcome outcome
		wantStatus  cgstore.TraitFetchStatus
	}{
		{
			name:        "missing package backs off",
			serve:       false,
			candidate:   cgstore.TraitFetchCandidate{Seed: testSeed},
			wantOutcome: outcomeMissing,
			wantStatus:  cgstore.TraitFetchMissing,
		},
		{
			name:        "unchanged package is skipped",
			body:        string(testfixtures.NftTraitsExample),
			serve:       true,
			etag:        `"same"`,
			candidate:   cgstore.TraitFetchCandidate{Seed: testSeed, SourceETag: `"same"`},
			wantOutcome: outcomeUnchanged,
			wantStatus:  cgstore.TraitFetchNotMod,
		},
		{
			name:        "malformed file is rejected",
			body:        `{"schema_version":`,
			serve:       true,
			candidate:   cgstore.TraitFetchCandidate{Seed: testSeed},
			wantOutcome: outcomeFailed,
			wantStatus:  cgstore.TraitFetchRejected,
		},
		{
			name:        "major schema bump is refused",
			body:        majorBumpFixture(),
			serve:       true,
			candidate:   cgstore.TraitFetchCandidate{Seed: testSeed},
			wantOutcome: outcomeFailed,
			wantStatus:  cgstore.TraitFetchRejected,
		},
		{
			name:        "seed mismatch is refused",
			body:        string(testfixtures.NftTraitsExample),
			serve:       true,
			candidate:   cgstore.TraitFetchCandidate{Seed: "0xdead"},
			wantOutcome: outcomeFailed,
			wantStatus:  cgstore.TraitFetchRejected,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			host := newAssetHost()
			host.etag = tc.etag
			if tc.serve {
				host.traits[tc.candidate.Seed] = tc.body
			}
			srv := host.serve(t)

			st := newFakeStore()
			ing := newTestIngester(t, st, srv.URL)

			if got := ing.ingestSeed(t.Context(), tc.candidate); got != tc.wantOutcome {
				t.Fatalf("ingestSeed = %v, want %v", got, tc.wantOutcome)
			}
			upserts, outcomes, _ := st.snapshot()
			if len(upserts) != 0 {
				t.Errorf("a non-storing outcome still upserted %d rows", len(upserts))
			}
			if len(outcomes) != 1 {
				t.Fatalf("len(outcomes) = %d, want 1", len(outcomes))
			}
			if outcomes[0].Status != tc.wantStatus {
				t.Errorf("status = %q, want %q", outcomes[0].Status, tc.wantStatus)
			}
			if outcomes[0].NextAttempt.IsZero() {
				t.Error("no next attempt was scheduled; the seed would be retried in a tight loop")
			}
		})
	}
}

// majorBumpFixture is the vendored example with its schema_version raised to
// a major this build must refuse.
func majorBumpFixture() string {
	var doc map[string]any
	if err := json.Unmarshal(testfixtures.NftTraitsExample, &doc); err != nil {
		panic(err)
	}
	doc["schema_version"] = "2.0.0"
	body, err := json.Marshal(doc)
	if err != nil {
		panic(err)
	}
	return string(body)
}

// TestIngestRaisesDriftWithoutOverwriting is the frozen-art guarantee: a
// regenerated package that disagrees with what collectors already saw must
// alarm, not silently replace the stored row.
func TestIngestRaisesDriftWithoutOverwriting(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	got := ing.ingestSeed(t.Context(), cgstore.TraitFetchCandidate{
		Seed:        testSeed,
		HasTraits:   true,
		ContentHash: "a-hash-from-a-previous-render",
	})
	if got != outcomeFailed {
		t.Fatalf("ingestSeed = %v, want outcomeFailed", got)
	}

	upserts, _, drifts := st.snapshot()
	if len(upserts) != 0 {
		t.Fatal("drift overwrote the stored art facts")
	}
	if len(drifts) != 1 {
		t.Fatalf("len(drifts) = %d, want 1", len(drifts))
	}
	if drifts[0].Seed != testSeed || drifts[0].PipelineVersion != "1.0.0" {
		t.Errorf("drift = %+v, want the seed and pipeline version", drifts[0])
	}
}

// TestIngestReingestsMatchingContent covers the benign re-check: identical
// art, possibly a newer pipeline, so the row is refreshed and no alarm fires.
func TestIngestReingestsMatchingContent(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	file, err := nftraits.Parse(testfixtures.NftTraitsExample)
	if err != nil {
		t.Fatalf("parsing the fixture: %v", err)
	}
	hash, err := nftraits.ContentHash(file)
	if err != nil {
		t.Fatalf("ContentHash: %v", err)
	}

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	got := ing.ingestSeed(t.Context(), cgstore.TraitFetchCandidate{
		Seed: testSeed, HasTraits: true, ContentHash: hash,
	})
	if got != outcomeStored {
		t.Fatalf("ingestSeed = %v, want outcomeStored", got)
	}
	upserts, _, drifts := st.snapshot()
	if len(drifts) != 0 {
		t.Errorf("identical content raised drift: %+v", drifts)
	}
	if len(upserts) != 1 {
		t.Fatalf("len(upserts) = %d, want 1", len(upserts))
	}
}

// TestIngestSurvivesMissingManifest keeps the manifest advisory: losing it
// costs image_details, never the whole ingest.
func TestIngestSurvivesMissingManifest(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	if got := ing.ingestSeed(t.Context(), cgstore.TraitFetchCandidate{Seed: testSeed}); got != outcomeStored {
		t.Fatalf("ingestSeed = %v, want outcomeStored", got)
	}
	upserts, _, _ := st.snapshot()
	if len(upserts[0].Assets) != 0 {
		t.Error("a missing manifest produced stored assets")
	}
	// The manifest is still requested: it is optional, not skipped.
	if _, manifestHits := host.hits(); manifestHits != 1 {
		t.Errorf("asset manifest requested %d times, want 1", manifestHits)
	}
}

func TestIngestSurvivesMalformedManifest(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	host.manifests[testSeed] = `{"schema_version":1,"assets":[]}`
	srv := host.serve(t)

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	if got := ing.ingestSeed(t.Context(), cgstore.TraitFetchCandidate{Seed: testSeed}); got != outcomeStored {
		t.Fatalf("ingestSeed = %v, want outcomeStored", got)
	}
	upserts, _, _ := st.snapshot()
	if len(upserts[0].Assets) != 0 {
		t.Error("an unsupported manifest schema was stored anyway")
	}
}

func TestIngestReportsTransportFailure(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer srv.Close()

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	if got := ing.ingestSeed(t.Context(), cgstore.TraitFetchCandidate{Seed: testSeed}); got != outcomeFailed {
		t.Fatalf("ingestSeed = %v, want outcomeFailed", got)
	}
	_, outcomes, _ := st.snapshot()
	if len(outcomes) != 1 || outcomes[0].Status != cgstore.TraitFetchTransient {
		t.Fatalf("outcomes = %+v, want one transient error", outcomes)
	}
	if outcomes[0].Cause == nil {
		t.Error("the transient failure recorded no cause")
	}
}

// TestIngestSurvivesBookkeepingFailures pins the scan's resilience: a
// database that cannot record an outcome, a drift alarm or an upsert must
// not stop the loop from working through the rest of the seeds.
func TestIngestSurvivesBookkeepingFailures(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	tests := map[string]struct {
		candidate cgstore.TraitFetchCandidate
		missing   bool
		want      outcome
	}{
		"outcome write fails": {
			candidate: cgstore.TraitFetchCandidate{Seed: testSeed},
			missing:   true,
			want:      outcomeMissing,
		},
		"drift write fails": {
			candidate: cgstore.TraitFetchCandidate{
				Seed: testSeed, HasTraits: true, ContentHash: "a-previous-render",
			},
			want: outcomeFailed,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			base := srv.URL
			if tc.missing {
				// An asset host with nothing published drives the
				// missing-package branch into the failing recorder.
				base = newAssetHost().serve(t).URL
			}
			st := &brokenBookkeepingStore{fakeStore: newFakeStore()}
			ing := newTestIngester(t, st, base)
			if got := ing.ingestSeed(t.Context(), tc.candidate); got != tc.want {
				t.Fatalf("ingestSeed = %v, want %v", got, tc.want)
			}
		})
	}
}

// brokenBookkeepingStore accepts the scan but fails every write, standing in
// for a database that is up enough to read and not to write.
type brokenBookkeepingStore struct {
	*fakeStore
}

func (b *brokenBookkeepingStore) RecordTraitFetchOutcome(
	context.Context, string, cgstore.TraitFetchStatus, error, time.Time,
) error {
	return errors.New("database is down")
}

func (b *brokenBookkeepingStore) RecordTraitDrift(context.Context, string, string, time.Time) error {
	return errors.New("database is down")
}

// TestIngestReportsManifestTransportFailure keeps the manifest advisory when
// the asset host answers, but badly.
func TestIngestReportsManifestTransportFailure(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/metadata/assets.json") {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(testfixtures.NftTraitsExample)
	}))
	defer srv.Close()

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	if got := ing.ingestSeed(t.Context(), cgstore.TraitFetchCandidate{Seed: testSeed}); got != outcomeStored {
		t.Fatalf("ingestSeed = %v, want outcomeStored", got)
	}
	upserts, _, _ := st.snapshot()
	if len(upserts[0].Assets) != 0 {
		t.Error("a failed manifest request produced stored assets")
	}
}

// TestRunOnceStopsOnCancellation pins the scan's cancellation checks: a
// shutdown mid-backlog must not keep hammering the asset host.
func TestRunOnceStopsOnCancellation(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	srv := host.serve(t)

	candidates := make([]cgstore.TraitFetchCandidate, 6)
	for i := range candidates {
		candidates[i] = cgstore.TraitFetchCandidate{Seed: seedForIndex(i)}
	}
	st := newFakeStore(candidates...)
	ing := newTestIngester(t, st, srv.URL)
	ing.batchSize = 2

	ctx, cancel := context.WithCancel(t.Context())
	// Cancel as soon as the first seed has been attempted, so the loop is
	// interrupted between candidates rather than before it starts.
	st.onOutcome = func() { cancel() }

	ing.runOnce(ctx)

	_, outcomes, _ := st.snapshot()
	if len(outcomes) != 1 {
		t.Errorf("processed %d seeds after cancellation, want it to stop at 1", len(outcomes))
	}
}

func TestRunOnceStopsBeforeAnyWorkWhenAlreadyCancelled(t *testing.T) {
	t.Parallel()
	st := newFakeStore(cgstore.TraitFetchCandidate{Seed: testSeed})
	ing := newTestIngester(t, st, newAssetHost().serve(t).URL)

	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	ing.runOnce(ctx)

	st.mu.Lock()
	calls := st.dueCalls
	st.mu.Unlock()
	if calls != 0 {
		t.Errorf("a cancelled scan queried the store %d times, want 0", calls)
	}
}

func TestIngestReportsStoreFailure(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	st := newFakeStore()
	st.upsertErr = errors.New("database is down")
	ing := newTestIngester(t, st, srv.URL)

	if got := ing.ingestSeed(t.Context(), cgstore.TraitFetchCandidate{Seed: testSeed}); got != outcomeFailed {
		t.Fatalf("ingestSeed = %v, want outcomeFailed", got)
	}
}

func TestIngestStopsOnCancelledContext(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if got := ing.ingestSeed(ctx, cgstore.TraitFetchCandidate{Seed: testSeed}); got != outcomeFailed {
		t.Fatalf("ingestSeed = %v, want outcomeFailed", got)
	}
	upserts, outcomes, _ := st.snapshot()
	if len(upserts) != 0 {
		t.Error("a cancelled attempt still stored a row")
	}
	if len(outcomes) != 0 {
		t.Error("a cancelled attempt still rescheduled the seed")
	}
}

// TestRunOnceDrainsUpToTheTickCeiling pins the cold-start policy: a large
// backlog is worked through in pages but never in one unbounded burst.
func TestRunOnceDrainsUpToTheTickCeiling(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	srv := host.serve(t)

	candidates := make([]cgstore.TraitFetchCandidate, 25)
	for i := range candidates {
		candidates[i] = cgstore.TraitFetchCandidate{Seed: seedForIndex(i)}
	}
	st := newFakeStore(candidates...)
	ing := newTestIngester(t, st, srv.URL)
	ing.batchSize = 10
	ing.maxPerTick = 20

	ing.runOnce(t.Context())

	_, outcomes, _ := st.snapshot()
	if len(outcomes) != 20 {
		t.Fatalf("processed %d seeds, want the 20-seed ceiling", len(outcomes))
	}
	traitsHits, _ := host.hits()
	if traitsHits != 20 {
		t.Errorf("asset host saw %d requests, want 20", traitsHits)
	}
}

func TestRunOnceStopsWhenTheScanIsEmpty(t *testing.T) {
	t.Parallel()
	srv := newAssetHost().serve(t)
	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)

	ing.runOnce(t.Context())

	st.mu.Lock()
	calls := st.dueCalls
	st.mu.Unlock()
	if calls != 1 {
		t.Errorf("empty scan queried the store %d times, want 1", calls)
	}
}

func TestRunOnceSurvivesScanFailure(t *testing.T) {
	t.Parallel()
	srv := newAssetHost().serve(t)
	st := newFakeStore()
	st.dueErr = errors.New("database is down")
	ing := newTestIngester(t, st, srv.URL)

	ing.runOnce(t.Context())

	_, outcomes, _ := st.snapshot()
	if len(outcomes) != 0 {
		t.Error("a failed scan still processed seeds")
	}
}

func seedForIndex(i int) string {
	const hexDigits = "0123456789abcdef"
	return "0x1" + string(hexDigits[i/16]) + string(hexDigits[i%16])
}

// offlineIngester builds an ingester whose asset host is never contacted, so
// the clock-sensitive queue tests can run inside a synctest bubble (real
// network I/O is not durably blocked and would deadlock the bubble).
func offlineIngester(tb testing.TB) *Ingester {
	tb.Helper()
	ing, err := New(Config{Store: newFakeStore(), SourceBase: "https://assets.invalid"})
	if err != nil {
		tb.Fatalf("New: %v", err)
	}
	return ing
}

// TestNudgeDedupesWithinTheNegativeCacheWindow keeps a trending token whose
// package is still rendering from turning page views into asset-host traffic.
func TestNudgeDedupesWithinTheNegativeCacheWindow(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ing := offlineIngester(t)

		for range 5 {
			ing.Nudge(testSeed)
		}
		if got := len(ing.nudges); got != 1 {
			t.Fatalf("queued %d nudges, want 1", got)
		}
		<-ing.nudges

		// Still inside the window: the repeat is dropped.
		time.Sleep(defaultNudgeTTL / 2)
		ing.Nudge(testSeed)
		if got := len(ing.nudges); got != 0 {
			t.Fatalf("queued %d nudges inside the window, want 0", got)
		}

		// Past the window: the seed may be requested again.
		time.Sleep(defaultNudgeTTL)
		ing.Nudge(testSeed)
		if got := len(ing.nudges); got != 1 {
			t.Fatalf("queued %d nudges after the window, want 1", got)
		}
	})
}

func TestNudgeIgnoresUnusableSeeds(t *testing.T) {
	t.Parallel()
	srv := newAssetHost().serve(t)
	ing := newTestIngester(t, newFakeStore(), srv.URL)

	for _, seed := range []string{"", "0x", "not-hex", "seed0001"} {
		ing.Nudge(seed)
	}
	if got := len(ing.nudges); got != 0 {
		t.Fatalf("queued %d nudges for unusable seeds, want 0", got)
	}
}

func TestNudgeCanonicalizesTheSeed(t *testing.T) {
	t.Parallel()
	srv := newAssetHost().serve(t)
	ing := newTestIngester(t, newFakeStore(), srv.URL)

	ing.Nudge("100033")
	select {
	case got := <-ing.nudges:
		if got != testSeed {
			t.Errorf("queued %q, want %q", got, testSeed)
		}
	default:
		t.Fatal("no nudge was queued")
	}
	// The canonical form is what the negative cache keys on, so the
	// prefixed spelling of the same seed is a duplicate.
	ing.Nudge(testSeed)
	if got := len(ing.nudges); got != 0 {
		t.Errorf("queued %d nudges for the same seed in two spellings, want 0", got)
	}
}

func TestNudgeNeverBlocksOnAFullQueue(t *testing.T) {
	t.Parallel()
	srv := newAssetHost().serve(t)
	ing := newTestIngester(t, newFakeStore(), srv.URL)

	for i := range defaultNudgeQueue + 50 {
		ing.Nudge(paddedSeed(i))
	}
	if got := len(ing.nudges); got != defaultNudgeQueue {
		t.Errorf("queue length = %d, want it capped at %d", got, defaultNudgeQueue)
	}
}

func paddedSeed(i int) string {
	const hexDigits = "0123456789abcdef"
	return "0x" + string(hexDigits[(i/256)%16]) + string(hexDigits[(i/16)%16]) + string(hexDigits[i%16]) + "1"
}

func TestSweepNudgeCacheDropsAgedEntries(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ing := offlineIngester(t)
		ing.Nudge(testSeed)
		<-ing.nudges

		ing.sweepNudgeCache()
		if _, ok := ing.recentNudges.Load(testSeed); !ok {
			t.Fatal("a fresh entry was swept")
		}

		time.Sleep(2 * defaultNudgeTTL)
		ing.sweepNudgeCache()
		if _, ok := ing.recentNudges.Load(testSeed); ok {
			t.Error("an aged entry survived the sweep")
		}
	})
}

func TestIngestNudgedHonoursTheSchedule(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	now := time.Date(2026, 8, 29, 12, 0, 0, 0, time.UTC)
	st := newFakeStore(cgstore.TraitFetchCandidate{
		Seed:  testSeed,
		DueAt: now.Add(time.Hour),
	})
	ing := newTestIngester(t, st, srv.URL)

	ing.ingestNudged(t.Context(), testSeed)

	if traitsHits, _ := host.hits(); traitsHits != 0 {
		t.Errorf("a backed-off seed was still fetched (%d requests)", traitsHits)
	}
}

func TestIngestNudgedRunsWhenDue(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	st := newFakeStore(cgstore.TraitFetchCandidate{Seed: testSeed})
	ing := newTestIngester(t, st, srv.URL)

	ing.ingestNudged(t.Context(), testSeed)

	upserts, _, _ := st.snapshot()
	if len(upserts) != 1 {
		t.Fatalf("len(upserts) = %d, want 1", len(upserts))
	}
}

func TestIngestNudgedIgnoresUnknownSeeds(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	srv := host.serve(t)

	st := newFakeStore()
	ing := newTestIngester(t, st, srv.URL)
	ing.ingestNudged(t.Context(), testSeed)

	if traitsHits, _ := host.hits(); traitsHits != 0 {
		t.Errorf("an unminted seed was fetched (%d requests)", traitsHits)
	}

	st.lookupErr = errors.New("database is down")
	ing.ingestNudged(t.Context(), testSeed)
	if traitsHits, _ := host.hits(); traitsHits != 0 {
		t.Errorf("a failed lookup still fetched (%d requests)", traitsHits)
	}
}

// TestBackoffGrowsAndSaturates pins the retry curve for the thousands of
// packages that are simply not rendered yet.
func TestBackoffGrowsAndSaturates(t *testing.T) {
	t.Parallel()
	now := time.Date(2026, 8, 29, 12, 0, 0, 0, time.UTC)
	ing := &Ingester{interval: 2 * time.Minute, now: func() time.Time { return now }}

	tests := []struct {
		attempt int
		want    time.Duration
	}{
		{0, 2 * time.Minute},
		{1, 4 * time.Minute},
		{2, 8 * time.Minute},
		{3, 16 * time.Minute},
		{6, 128 * time.Minute},
		{7, 256 * time.Minute},
		// 2m doubled 8 times is 8h32m, past the ceiling.
		{8, maxBackoff},
		{50, maxBackoff},
		{-1, 2 * time.Minute},
	}
	for _, tc := range tests {
		if got := ing.backoff(tc.attempt).Sub(now); got != tc.want {
			t.Errorf("backoff(%d) = %v, want %v", tc.attempt, got, tc.want)
		}
	}
}

// TestBackoffRespectsAShortInterval keeps the first retry no sooner than the
// floor even when the scan interval is configured aggressively low.
func TestBackoffRespectsAShortInterval(t *testing.T) {
	t.Parallel()
	now := time.Now()
	ing := &Ingester{interval: time.Second, now: func() time.Time { return now }}
	if got := ing.backoff(0).Sub(now); got != minBackoff {
		t.Errorf("backoff(0) = %v, want the %v floor", got, minBackoff)
	}
}

// TestRunScansUntilCancelled exercises the loop's lifecycle: an immediate
// first pass, a pass per tick, and a prompt return on cancellation.
//
// It runs on the real clock because the loop performs real HTTP against the
// stub asset host, which a synctest bubble cannot contain. The barriers are
// channel signals from the fake store; the timeouts are only outer guards.
func TestRunScansUntilCancelled(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	st := newFakeStore(cgstore.TraitFetchCandidate{Seed: testSeed})
	ing, err := New(Config{
		Store:      st,
		SourceBase: srv.URL,
		Interval:   10 * time.Millisecond,
		HTTPClient: srv.Client(),
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	ctx, cancel := context.WithCancel(t.Context())
	done := make(chan struct{})
	go func() {
		defer close(done)
		ing.Run(ctx)
	}()

	awaitSignal(t, st.stored, "the first pass to store the package")
	// The candidate is consumed, so later ticks find nothing to do but must
	// keep scanning.
	for range 3 {
		awaitSignal(t, st.scanned, "a periodic scan")
	}

	cancel()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("Run did not return after cancellation")
	}

	upserts, _, _ := st.snapshot()
	if len(upserts) != 1 {
		t.Errorf("stored %d rows, want exactly 1", len(upserts))
	}
}

// TestRunServicesNudges proves the on-demand path reaches the asset host
// without waiting out the scan interval.
func TestRunServicesNudges(t *testing.T) {
	t.Parallel()
	host := newAssetHost()
	host.traits[testSeed] = string(testfixtures.NftTraitsExample)
	srv := host.serve(t)

	st := newFakeStore(cgstore.TraitFetchCandidate{Seed: testSeed})
	st.scanCursor = 1 // the periodic scan has nothing due
	ing, err := New(Config{
		Store:      st,
		SourceBase: srv.URL,
		Interval:   time.Hour,
		HTTPClient: srv.Client(),
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	ctx, cancel := context.WithCancel(t.Context())
	done := make(chan struct{})
	go func() {
		defer close(done)
		ing.Run(ctx)
	}()
	awaitSignal(t, st.scanned, "the first pass to find nothing due")

	ing.Nudge(testSeed)
	awaitSignal(t, st.stored, "the nudge to store the package")

	cancel()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("Run did not return after cancellation")
	}
}

func TestBuildUpsertRejectsUnusableFiles(t *testing.T) {
	t.Parallel()
	now := time.Now()
	tests := map[string]struct {
		seed string
		file *nftraits.File
	}{
		"non-hex requested seed": {"zz", &nftraits.File{Seed: testSeed, SchemaVersion: "1.0.0"}},
		"malformed semver":       {testSeed, &nftraits.File{Seed: testSeed, SchemaVersion: "one"}},
		"broken simulation":      {testSeed, &nftraits.File{Seed: testSeed, SchemaVersion: "1.0.0", Simulation: json.RawMessage(`{`)}},
		"broken generation":      {testSeed, &nftraits.File{Seed: testSeed, SchemaVersion: "1.0.0", Generation: json.RawMessage(`{`)}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, err := buildUpsert(tc.seed, tc.file, "hash", nil, "", "", now); err == nil {
				t.Fatal("buildUpsert accepted an unusable file")
			}
		})
	}
}

func TestOrHelpers(t *testing.T) {
	t.Parallel()
	if got := orDuration(0, time.Minute); got != time.Minute {
		t.Errorf("orDuration(0) = %v, want the fallback", got)
	}
	if got := orDuration(-time.Second, time.Minute); got != time.Minute {
		t.Errorf("orDuration(negative) = %v, want the fallback", got)
	}
	if got := orDuration(time.Hour, time.Minute); got != time.Hour {
		t.Errorf("orDuration(hour) = %v, want the value", got)
	}
	if got := orInt(0, 7); got != 7 {
		t.Errorf("orInt(0) = %d, want the fallback", got)
	}
	if got := orInt(-1, 7); got != 7 {
		t.Errorf("orInt(negative) = %d, want the fallback", got)
	}
	if got := orInt(3, 7); got != 3 {
		t.Errorf("orInt(3) = %d, want the value", got)
	}
}
