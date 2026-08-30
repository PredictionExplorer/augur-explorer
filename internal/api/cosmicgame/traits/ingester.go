// Package traits ingests the per-seed art trait packages published by the
// CS-Image-Generation pipeline into the database, so the ERC-721 metadata
// handler can serve them without ever touching the asset host inside a
// request.
//
// The loop is deliberately boring: scan for seeds whose next attempt is due,
// conditionally GET the trait file and its asset manifest, gate the result
// on the schema major and the seed, and upsert. Packages land minutes to
// hours after a mint, so the common outcome is a 404 that backs off rather
// than an error that alarms.
//
// Art facts are frozen at imprint by construction. When a re-fetched file
// disagrees with what is already stored, the ingester keeps the stored row
// and raises drift instead of adopting the new values — a determinism
// regression in the generator is an incident, not an update.
package traits

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/nftraits"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// Default tuning. The interval matches how promptly a new mint should flip
// from the fallback to the enriched response; the per-tick ceiling keeps a
// cold-start backlog from turning into a burst against the asset host.
const (
	defaultInterval      = 2 * time.Minute
	defaultBatchSize     = 100
	defaultMaxPerTick    = 500
	defaultRecheck       = 24 * time.Hour
	defaultRequestTimout = 15 * time.Second
	defaultNudgeTTL      = 5 * time.Minute
	defaultNudgeQueue    = 256

	minBackoff = 2 * time.Minute
	maxBackoff = 6 * time.Hour
)

// Store is the database surface the ingester needs. The concrete
// implementation is *cosmicgame.Repo; the interface keeps the loop testable
// without a container.
type Store interface {
	SeedsDueForTraitFetch(ctx context.Context, limit int) ([]cgstore.TraitFetchCandidate, error)
	TraitFetchCandidate(ctx context.Context, seed string) (cgstore.TraitFetchCandidate, error)
	UpsertTokenTraits(ctx context.Context, up cgstore.TokenTraitsUpsert) error
	RecordTraitFetchOutcome(ctx context.Context, seed string, status cgstore.TraitFetchStatus, cause error, nextAttempt time.Time) error
	RecordTraitDrift(ctx context.Context, seed, pipelineVersion string, nextAttempt time.Time) error
}

// Config carries the ingester's dependencies and tuning. Only Store and
// SourceBase are required.
type Config struct {
	// Store persists ingested packages and the fetch schedule.
	Store Store
	// SourceBase is the asset host root that publishes /traits/{seed}.json
	// and /asset-manifests/{seed}.json.
	SourceBase string
	// Interval is the scan period; zero selects two minutes.
	Interval time.Duration
	// RecheckInterval is how long a stored package waits before it is
	// re-fetched for drift; zero selects a day.
	RecheckInterval time.Duration
	// BatchSize is the database page size per scan; zero selects 100.
	BatchSize int
	// MaxPerTick bounds the seeds one tick may process; zero selects 500.
	MaxPerTick int
	// HTTPClient overrides the default client (tests inject httptest).
	HTTPClient *http.Client
	// Logger receives ingest diagnostics; nil discards.
	Logger *slog.Logger
	// Now overrides the clock, for tests.
	Now func() time.Time
}

// Ingester fills the trait tables from the asset host.
type Ingester struct {
	store           Store
	fetcher         *fetcher
	interval        time.Duration
	recheckInterval time.Duration
	batchSize       int
	maxPerTick      int
	logger          *slog.Logger
	now             func() time.Time

	nudges   chan string
	nudgeTTL time.Duration
	// recentNudges dedupes on-demand requests so a token that is trending
	// while its package is still rendering cannot turn page views into
	// asset-host traffic. Swept once per tick.
	recentNudges sync.Map
}

// New validates the configuration and returns a ready ingester. It returns
// an error when no asset host is configured, which the caller should treat
// as "trait ingestion is switched off", not as a startup failure: with no
// trait rows every token simply serves the fallback metadata.
func New(cfg Config) (*Ingester, error) {
	if cfg.Store == nil {
		return nil, errors.New("traits: store is required")
	}
	client := cfg.HTTPClient
	if client == nil {
		client = &http.Client{Timeout: defaultRequestTimout}
	}
	f, err := newFetcher(client, cfg.SourceBase)
	if err != nil {
		return nil, fmt.Errorf("traits: %w", err)
	}
	ing := &Ingester{
		store:           cfg.Store,
		fetcher:         f,
		interval:        orDuration(cfg.Interval, defaultInterval),
		recheckInterval: orDuration(cfg.RecheckInterval, defaultRecheck),
		batchSize:       orInt(cfg.BatchSize, defaultBatchSize),
		maxPerTick:      orInt(cfg.MaxPerTick, defaultMaxPerTick),
		logger:          cfg.Logger,
		now:             cfg.Now,
		nudges:          make(chan string, defaultNudgeQueue),
		nudgeTTL:        defaultNudgeTTL,
	}
	if ing.logger == nil {
		ing.logger = slog.New(slog.DiscardHandler)
	}
	if ing.now == nil {
		ing.now = time.Now
	}
	return ing, nil
}

// Run scans for due seeds every Interval until ctx is cancelled, servicing
// on-demand nudges in between. It returns when ctx is done.
func (i *Ingester) Run(ctx context.Context) {
	i.logger.Info("trait ingester started",
		"source", i.fetcher.base,
		"interval", i.interval,
		"recheck_interval", i.recheckInterval,
	)
	ticker := time.NewTicker(i.interval)
	defer ticker.Stop()

	i.runOnce(ctx)
	for {
		select {
		case <-ctx.Done():
			i.logger.Info("trait ingester stopped")
			return
		case <-ticker.C:
			i.runOnce(ctx)
		case seed := <-i.nudges:
			i.ingestNudged(ctx, seed)
		}
	}
}

// Nudge asks for one seed to be ingested soon. It never blocks and never
// performs I/O, so it is safe to call from an HTTP handler: the metadata
// path stays a pure database read while a token whose package just landed
// still flips to the enriched response within seconds instead of waiting
// out the scan interval.
func (i *Ingester) Nudge(seed string) {
	canonical, err := nftraits.CanonicalSeed(seed)
	if err != nil {
		return
	}
	now := i.now()
	if last, ok := i.recentNudges.Load(canonical); ok {
		if at, ok := last.(time.Time); ok && now.Sub(at) < i.nudgeTTL {
			return
		}
	}
	i.recentNudges.Store(canonical, now)
	select {
	case i.nudges <- canonical:
	default:
		// The queue is full; the periodic scan covers this seed anyway.
	}
}

// runOnce processes one scan: drain up to MaxPerTick due seeds in database
// pages of BatchSize.
func (i *Ingester) runOnce(ctx context.Context) {
	i.sweepNudgeCache()

	started := i.now()
	var processed, stored, missing, failed int
	for processed < i.maxPerTick {
		if ctx.Err() != nil {
			return
		}
		page := min(i.batchSize, i.maxPerTick-processed)
		candidates, err := i.store.SeedsDueForTraitFetch(ctx, page)
		if err != nil {
			if ctx.Err() == nil {
				i.logger.Error("trait ingest scan failed", "error", err)
			}
			return
		}
		if len(candidates) == 0 {
			break
		}
		for _, candidate := range candidates {
			if ctx.Err() != nil {
				return
			}
			processed++
			switch i.ingestSeed(ctx, candidate) {
			case outcomeStored:
				stored++
			case outcomeMissing:
				missing++
			case outcomeFailed:
				failed++
			case outcomeUnchanged:
			}
		}
		if len(candidates) < page {
			break
		}
	}
	if processed > 0 {
		i.logger.Info("trait ingest scan complete",
			"processed", processed,
			"stored", stored,
			"missing", missing,
			"failed", failed,
			"duration", i.now().Sub(started),
		)
	}
}

// ingestNudged resolves the on-demand seed's current state and ingests it,
// honouring the persisted schedule so a nudge cannot bypass the backoff.
func (i *Ingester) ingestNudged(ctx context.Context, seed string) {
	candidate, err := i.store.TraitFetchCandidate(ctx, seed)
	if err != nil {
		if ctx.Err() == nil && !errors.Is(err, store.ErrNotFound) {
			i.logger.Error("trait nudge lookup failed", "seed", seed, "error", err)
		}
		return
	}
	if !candidate.DueAt.IsZero() && i.now().Before(candidate.DueAt) {
		return
	}
	i.ingestSeed(ctx, candidate)
}

// outcome classifies one seed's attempt for the scan summary.
type outcome int

const (
	outcomeStored outcome = iota
	outcomeUnchanged
	outcomeMissing
	outcomeFailed
)

// ingestSeed runs the full attempt for one seed and records its outcome.
// Every branch writes the fetch schedule, so a seed can never be retried in
// a tight loop regardless of how the attempt ended.
func (i *Ingester) ingestSeed(ctx context.Context, candidate cgstore.TraitFetchCandidate) outcome {
	seed := candidate.Seed
	result, err := i.fetcher.get(ctx, i.fetcher.traitsURL(seed), candidate.SourceETag)
	if err != nil {
		if ctx.Err() != nil {
			return outcomeFailed
		}
		i.logger.Warn("trait fetch failed", "seed", seed, "error", err)
		i.record(ctx, seed, cgstore.TraitFetchTransient, err, i.backoff(candidate.AttemptCount))
		return outcomeFailed
	}

	switch result.Status {
	case fetchMissing:
		// Expected for recent mints: the generator has not published yet.
		i.record(ctx, seed, cgstore.TraitFetchMissing, nil, i.backoff(candidate.AttemptCount))
		return outcomeMissing
	case fetchNotModified:
		i.record(ctx, seed, cgstore.TraitFetchNotMod, nil, i.now().Add(i.recheckInterval))
		return outcomeUnchanged
	case fetchOK:
	}

	file, err := nftraits.Parse(result.Body)
	if err != nil {
		i.logger.Error("trait file malformed", "seed", seed, "error", err)
		i.record(ctx, seed, cgstore.TraitFetchRejected, err, i.backoff(candidate.AttemptCount))
		return outcomeFailed
	}
	if err := file.Gate(seed); err != nil {
		// A major schema bump or a seed mismatch means this build must not
		// serve the file. Keep the last good row and page a human.
		i.logger.Error("trait file refused by the schema gate",
			"seed", seed,
			"schema_version", file.SchemaVersion,
			"pipeline_version", file.PipelineVersion,
			"error", err,
		)
		i.record(ctx, seed, cgstore.TraitFetchRejected, err, i.backoff(candidate.AttemptCount))
		return outcomeFailed
	}

	hash, err := nftraits.ContentHash(file)
	if err != nil {
		i.logger.Error("hashing trait file failed", "seed", seed, "error", err)
		i.record(ctx, seed, cgstore.TraitFetchRejected, err, i.backoff(candidate.AttemptCount))
		return outcomeFailed
	}
	if candidate.HasTraits && candidate.ContentHash != hash {
		i.logger.Error("trait drift detected: stored art facts differ from the regenerated file",
			"seed", seed,
			"stored_content_hash", candidate.ContentHash,
			"fetched_content_hash", hash,
			"pipeline_version", file.PipelineVersion,
		)
		if err := i.store.RecordTraitDrift(ctx, seed, file.PipelineVersion, i.now().Add(i.recheckInterval)); err != nil {
			if ctx.Err() == nil {
				i.logger.Error("recording trait drift failed", "seed", seed, "error", err)
			}
		}
		return outcomeFailed
	}

	manifestBody, manifestETag := i.fetchManifest(ctx, seed, candidate.ManifestETag)

	upsert, err := buildUpsert(seed, file, hash, manifestBody, result.ETag, manifestETag, i.now().Add(i.recheckInterval))
	if err != nil {
		i.logger.Error("encoding trait file for storage failed", "seed", seed, "error", err)
		i.record(ctx, seed, cgstore.TraitFetchRejected, err, i.backoff(candidate.AttemptCount))
		return outcomeFailed
	}
	if err := i.store.UpsertTokenTraits(ctx, upsert); err != nil {
		if ctx.Err() == nil {
			i.logger.Error("storing trait file failed", "seed", seed, "error", err)
		}
		return outcomeFailed
	}
	i.logger.Debug("trait package ingested",
		"seed", seed,
		"pipeline_version", file.PipelineVersion,
		"attributes", len(file.Attributes),
		"manifest", manifestBody != nil,
	)
	return outcomeStored
}

// fetchManifest retrieves assets.json best-effort. It only enriches
// image_details and animation_details, so any failure degrades the served
// metadata by one optional object rather than blocking the ingest.
func (i *Ingester) fetchManifest(ctx context.Context, seed, etag string) (manifest *nftraits.Manifest, newETag string) {
	result, err := i.fetcher.get(ctx, i.fetcher.manifestURL(seed), etag)
	if err != nil {
		if ctx.Err() == nil {
			i.logger.Warn("asset manifest fetch failed", "seed", seed, "error", err)
		}
		return nil, ""
	}
	if result.Status != fetchOK {
		// 304 keeps the stored manifest through the upsert's COALESCE; 404
		// means the package predates schema 2 or is still uploading.
		return nil, ""
	}
	parsed, err := nftraits.ParseManifest(result.Body)
	if err != nil {
		i.logger.Warn("asset manifest malformed", "seed", seed, "error", err)
		return nil, ""
	}
	return parsed, result.ETag
}

// buildUpsert encodes the gated file into the store's byte-oriented payload.
//
// The row is keyed on requestedSeed — the indexer's spelling, zero-padded to
// 64 hex characters because NftMinted carries a uint256 — and never on the
// file's own, which the generator may write unpadded. Gate has already
// established that the two denote the same seed; keying on anything but the
// indexer's form would store rows the metadata handler cannot find.
//
// The blocks are canonicalized so an identical package always produces
// identical stored bytes, which keeps the served JSON and the golden tests
// stable across re-ingests.
func buildUpsert(
	requestedSeed string,
	file *nftraits.File,
	contentHash string,
	manifest *nftraits.Manifest,
	sourceETag, manifestETag string,
	nextAttempt time.Time,
) (cgstore.TokenTraitsUpsert, error) {
	seed, err := nftraits.CanonicalSeed(requestedSeed)
	if err != nil {
		return cgstore.TokenTraitsUpsert{}, err
	}
	major, err := nftraits.SemverMajor(file.SchemaVersion)
	if err != nil {
		return cgstore.TokenTraitsUpsert{}, err
	}
	attributes, err := json.Marshal(file.Attributes)
	if err != nil {
		return cgstore.TokenTraitsUpsert{}, fmt.Errorf("encoding attributes: %w", err)
	}
	simulation, err := nftraits.CanonicalJSON(file.Simulation)
	if err != nil {
		return cgstore.TokenTraitsUpsert{}, fmt.Errorf("encoding simulation: %w", err)
	}
	generation, err := nftraits.CanonicalJSON(file.Generation)
	if err != nil {
		return cgstore.TokenTraitsUpsert{}, fmt.Errorf("encoding generation: %w", err)
	}
	var assets []byte
	if manifest != nil {
		if assets, err = json.Marshal(manifest); err != nil {
			return cgstore.TokenTraitsUpsert{}, fmt.Errorf("encoding asset manifest: %w", err)
		}
	}
	return cgstore.TokenTraitsUpsert{
		Seed:            seed,
		SchemaMajor:     major,
		PipelineVersion: file.PipelineVersion,
		Attributes:      attributes,
		DescriptionArt:  file.DescriptionArt,
		Simulation:      simulation,
		Generation:      generation,
		Assets:          assets,
		ContentHash:     contentHash,
		SourceETag:      sourceETag,
		ManifestETag:    manifestETag,
		NextAttemptAt:   nextAttempt,
	}, nil
}

// record writes one non-storing outcome, logging rather than propagating a
// bookkeeping failure: the scan must keep going for the other seeds.
func (i *Ingester) record(
	ctx context.Context,
	seed string,
	status cgstore.TraitFetchStatus,
	cause error,
	nextAttempt time.Time,
) {
	if err := i.store.RecordTraitFetchOutcome(ctx, seed, status, cause, nextAttempt); err != nil {
		if ctx.Err() == nil {
			i.logger.Error("recording trait fetch outcome failed", "seed", seed, "status", status, "error", err)
		}
	}
}

// backoff schedules the next attempt after a fruitless one: exponential from
// the scan interval, capped, so thousands of not-yet-rendered packages
// settle into a slow trickle instead of a fixed-rate hammer.
func (i *Ingester) backoff(attemptCount int) time.Time {
	delay := max(i.interval, minBackoff)
	for range min(attemptCount, 12) {
		delay *= 2
		if delay >= maxBackoff {
			delay = maxBackoff
			break
		}
	}
	return i.now().Add(delay)
}

// sweepNudgeCache drops negative-cache entries that have aged out, keeping
// the map proportional to recent traffic rather than to total supply.
func (i *Ingester) sweepNudgeCache() {
	cutoff := i.now().Add(-i.nudgeTTL)
	i.recentNudges.Range(func(key, value any) bool {
		if at, ok := value.(time.Time); ok && at.Before(cutoff) {
			i.recentNudges.Delete(key)
		}
		return true
	})
}

func orDuration(v, fallback time.Duration) time.Duration {
	if v <= 0 {
		return fallback
	}
	return v
}

func orInt(v, fallback int) int {
	if v <= 0 {
		return fallback
	}
	return v
}
