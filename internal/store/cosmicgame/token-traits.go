package cosmicgame

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// TokenTraitsRow is the stored art contract for one seed. The JSON columns
// are handed to the metadata handler as raw bytes: they were copied verbatim
// from the generator and are served the same way.
type TokenTraitsRow struct {
	Seed            string
	SchemaMajor     int
	PipelineVersion string
	Attributes      []byte
	DescriptionArt  string
	Simulation      []byte
	Generation      []byte
	// Assets is the assets.json payload, nil when no manifest was ingested.
	Assets      []byte
	ContentHash string
}

// TokenTraitsUpsert is one successful ingest: the art facts plus the
// validators that let the next fetch answer 304.
//
// The JSON fields are pre-encoded by the caller. Keeping the store in bytes
// rather than in the contract's Go types is deliberate — the trait file is
// an external contract, and the database layer should not have an opinion
// about its shape beyond "these blocks are JSON".
type TokenTraitsUpsert struct {
	Seed            string
	SchemaMajor     int
	PipelineVersion string
	Attributes      []byte
	DescriptionArt  string
	Simulation      []byte
	Generation      []byte
	Assets          []byte
	ContentHash     string
	SourceETag      string
	ManifestETag    string
	// NextAttemptAt schedules the next drift re-check for this seed.
	NextAttemptAt time.Time
}

// TraitFetchCandidate is one seed the ingest loop may try, carrying
// everything needed to make the request conditional, to detect drift and to
// size the backoff without a second query.
type TraitFetchCandidate struct {
	Seed         string
	AttemptCount int
	SourceETag   string
	ManifestETag string
	// HasTraits reports whether art facts are already stored, which turns
	// the fetch from a first ingest into a drift re-check.
	HasTraits bool
	// ContentHash is the stored art-fact digest, empty when HasTraits is
	// false. A re-fetch that hashes differently is drift.
	ContentHash string
	// DueAt is the scheduled next attempt, zero when the seed has never
	// been attempted. On-demand paths honour it so a persistently missing
	// package cannot be re-requested on every page view.
	DueAt time.Time
}

// TraitFetchStatus labels the outcome of one ingest attempt for operators
// reading cg_token_traits_fetch.last_status.
type TraitFetchStatus string

// The outcomes the ingest loop records. Missing is the common one: packages
// land minutes to hours after the mint, so a 404 is expected, not an error.
const (
	TraitFetchOK        TraitFetchStatus = "ok"
	TraitFetchMissing   TraitFetchStatus = "missing"
	TraitFetchNotMod    TraitFetchStatus = "not_modified"
	TraitFetchRejected  TraitFetchStatus = "rejected"
	TraitFetchDrifted   TraitFetchStatus = "drifted"
	TraitFetchTransient TraitFetchStatus = "error"
)

// TokenTraits returns the stored art contract for a canonical 0x-prefixed
// lowercase seed, or store.ErrNotFound when the package has not been
// ingested yet (the handler then serves the fallback metadata).
func (r *Repo) TokenTraits(ctx context.Context, seed string) (TokenTraitsRow, error) {
	const query = `SELECT
			seed,
			schema_major,
			pipeline_version,
			attributes,
			description_art,
			simulation,
			generation,
			assets,
			content_hash
		FROM cg_token_traits
		WHERE seed=$1`

	var row TokenTraitsRow
	var assets []byte
	err := r.q(ctx).QueryRow(ctx, query, seed).Scan(
		&row.Seed,
		&row.SchemaMajor,
		&row.PipelineVersion,
		&row.Attributes,
		&row.DescriptionArt,
		&row.Simulation,
		&row.Generation,
		&assets,
		&row.ContentHash,
	)
	if err != nil {
		return TokenTraitsRow{}, store.WrapError("token traits", err)
	}
	row.Assets = assets
	return row, nil
}

// SeedsDueForTraitFetch returns at most limit minted seeds whose next ingest
// attempt is due, never-attempted seeds first and then oldest schedule
// first. Seeds are canonicalized to "0x" + lowercase here so the caller
// works in one form throughout.
//
// Non-hexadecimal seeds are filtered out in SQL: they cannot address a
// package on the asset host and must never reach the trait tables, whose
// primary keys are constrained to hex.
func (r *Repo) SeedsDueForTraitFetch(ctx context.Context, limit int) ([]TraitFetchCandidate, error) {
	const op = "seeds due for trait fetch"
	if limit <= 0 {
		return nil, fmt.Errorf("%s: limit must be positive", op)
	}
	query := traitFetchCandidateSelectSQL + `
		WHERE f.seed IS NULL OR f.next_attempt_at<=NOW()
		ORDER BY COALESCE(f.next_attempt_at, '-infinity'::TIMESTAMPTZ), s.seed
		LIMIT $1`

	return queryList(ctx, r, op, limit, query, scanTraitFetchCandidate, limit)
}

// TraitFetchCandidate returns the ingest state of one canonical seed, or
// store.ErrNotFound when no token was ever minted with it. The on-demand
// path uses it so a nudge from the metadata handler runs through exactly the
// same conditional-request and drift logic as the periodic scan.
func (r *Repo) TraitFetchCandidate(ctx context.Context, seed string) (TraitFetchCandidate, error) {
	const op = "trait fetch candidate"
	query := traitFetchCandidateSelectSQL + `
		WHERE s.seed=$1`

	var rec TraitFetchCandidate
	var dueAt sql.NullTime
	err := r.q(ctx).QueryRow(ctx, query, seed).Scan(
		&rec.Seed,
		&rec.AttemptCount,
		&rec.SourceETag,
		&rec.ManifestETag,
		&rec.HasTraits,
		&rec.ContentHash,
		&dueAt,
	)
	if err != nil {
		return TraitFetchCandidate{}, store.WrapError(op, err)
	}
	if dueAt.Valid {
		rec.DueAt = dueAt.Time
	}
	return rec, nil
}

// traitFetchCandidateSelectSQL projects one row per minted seed with its
// ingest state. Seeds are canonicalized to "0x" + lowercase to match the
// trait tables' keys, and non-hexadecimal seeds are excluded: they cannot
// address a package on the asset host and the tables constrain their key to
// hex.
const traitFetchCandidateSelectSQL = `WITH seeds AS (
			SELECT DISTINCT '0x' || lower(m.seed) AS seed
			FROM cg_mint_event m
			WHERE m.seed ~ '^[0-9a-fA-F]+$'
		)
		SELECT
			s.seed,
			COALESCE(f.attempt_count, 0),
			COALESCE(f.source_etag, ''),
			COALESCE(f.manifest_etag, ''),
			(t.seed IS NOT NULL),
			COALESCE(t.content_hash, ''),
			f.next_attempt_at
		FROM seeds s
			LEFT JOIN cg_token_traits_fetch f ON f.seed=s.seed
			LEFT JOIN cg_token_traits t ON t.seed=s.seed`

func scanTraitFetchCandidate(rows pgx.Rows, rec *TraitFetchCandidate) error {
	var dueAt sql.NullTime
	if err := rows.Scan(
		&rec.Seed,
		&rec.AttemptCount,
		&rec.SourceETag,
		&rec.ManifestETag,
		&rec.HasTraits,
		&rec.ContentHash,
		&dueAt,
	); err != nil {
		return err
	}
	if dueAt.Valid {
		rec.DueAt = dueAt.Time
	}
	return nil
}

// UpsertTokenTraits stores one successfully gated package and resets its
// fetch schedule, both in a single transaction so a crash can never leave
// art facts without the validators that describe them.
//
// A fetch that produced no manifest keeps the previously stored one: the
// trait file and assets.json are published independently, and losing
// image_details because a manifest request happened to fail would be a
// regression in the served metadata.
func (r *Repo) UpsertTokenTraits(ctx context.Context, up TokenTraitsUpsert) error {
	const op = "upsert token traits"
	if up.Seed == "" {
		return fmt.Errorf("%s: empty seed", op)
	}
	return r.store.InTx(ctx, func(ctx context.Context) error {
		const traitsQuery = `INSERT INTO cg_token_traits(
				seed,schema_major,pipeline_version,attributes,description_art,
				simulation,generation,assets,content_hash
			) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
			ON CONFLICT (seed) DO UPDATE SET
				schema_major=EXCLUDED.schema_major,
				pipeline_version=EXCLUDED.pipeline_version,
				attributes=EXCLUDED.attributes,
				description_art=EXCLUDED.description_art,
				simulation=EXCLUDED.simulation,
				generation=EXCLUDED.generation,
				assets=COALESCE(EXCLUDED.assets, cg_token_traits.assets),
				content_hash=EXCLUDED.content_hash,
				updated_at=NOW()`
		if _, err := r.q(ctx).Exec(ctx, traitsQuery,
			up.Seed,
			up.SchemaMajor,
			up.PipelineVersion,
			up.Attributes,
			up.DescriptionArt,
			up.Simulation,
			up.Generation,
			nullBytes(up.Assets),
			up.ContentHash,
		); err != nil {
			return store.WrapError(op, err)
		}

		const fetchQuery = `INSERT INTO cg_token_traits_fetch(
				seed,source_etag,manifest_etag,attempt_count,
				last_attempt_at,next_attempt_at,last_status,last_error
			) VALUES($1,$2,$3,0,NOW(),$4,$5,NULL)
			ON CONFLICT (seed) DO UPDATE SET
				source_etag=EXCLUDED.source_etag,
				manifest_etag=COALESCE(EXCLUDED.manifest_etag, cg_token_traits_fetch.manifest_etag),
				attempt_count=0,
				last_attempt_at=NOW(),
				next_attempt_at=EXCLUDED.next_attempt_at,
				last_status=EXCLUDED.last_status,
				last_error=NULL`
		if _, err := r.q(ctx).Exec(ctx, fetchQuery,
			up.Seed,
			nullString(up.SourceETag),
			nullString(up.ManifestETag),
			up.NextAttemptAt.UTC(),
			string(TraitFetchOK),
		); err != nil {
			return store.WrapError(op, err)
		}
		return nil
	})
}

// RecordTraitFetchOutcome schedules the next attempt for a seed that did not
// produce a new stored package: a 404 because the generator has not uploaded
// it yet, a 304 because the stored row is current, a rejected file, or a
// transport error. The art facts, if any, are untouched.
func (r *Repo) RecordTraitFetchOutcome(
	ctx context.Context,
	seed string,
	status TraitFetchStatus,
	cause error,
	nextAttempt time.Time,
) error {
	const op = "record trait fetch outcome"
	if seed == "" {
		return fmt.Errorf("%s: empty seed", op)
	}
	const query = `INSERT INTO cg_token_traits_fetch(
			seed,attempt_count,last_attempt_at,next_attempt_at,last_status,last_error
		) VALUES($1,1,NOW(),$2,$3,$4)
		ON CONFLICT (seed) DO UPDATE SET
			attempt_count=cg_token_traits_fetch.attempt_count+1,
			last_attempt_at=NOW(),
			next_attempt_at=EXCLUDED.next_attempt_at,
			last_status=EXCLUDED.last_status,
			last_error=EXCLUDED.last_error`
	var lastError any
	if cause != nil {
		lastError = truncateError(cause.Error())
	}
	if _, err := r.q(ctx).Exec(ctx, query, seed, nextAttempt.UTC(), string(status), lastError); err != nil {
		return store.WrapError(op, err)
	}
	return nil
}

// RecordTraitDrift flags a seed whose re-fetched trait file disagrees with
// the stored art facts and reschedules it. The stored row is deliberately
// left alone: art is frozen at imprint by construction, so the correct
// response is to keep serving what collectors already saw and to page a
// human, never to silently adopt the new values.
func (r *Repo) RecordTraitDrift(ctx context.Context, seed, pipelineVersion string, nextAttempt time.Time) error {
	const op = "record trait drift"
	if seed == "" {
		return fmt.Errorf("%s: empty seed", op)
	}
	const query = `INSERT INTO cg_token_traits_fetch(
			seed,attempt_count,last_attempt_at,next_attempt_at,last_status,
			drift_detected_at,drift_pipeline_version
		) VALUES($1,1,NOW(),$2,$3,NOW(),$4)
		ON CONFLICT (seed) DO UPDATE SET
			attempt_count=cg_token_traits_fetch.attempt_count+1,
			last_attempt_at=NOW(),
			next_attempt_at=EXCLUDED.next_attempt_at,
			last_status=EXCLUDED.last_status,
			drift_detected_at=COALESCE(cg_token_traits_fetch.drift_detected_at, NOW()),
			drift_pipeline_version=EXCLUDED.drift_pipeline_version`
	if _, err := r.q(ctx).Exec(ctx, query,
		seed,
		nextAttempt.UTC(),
		string(TraitFetchDrifted),
		nullString(pipelineVersion),
	); err != nil {
		return store.WrapError(op, err)
	}
	return nil
}

// DriftedTraitSeeds returns the seeds currently flagged as drifted, newest
// first. Operators and the database validation scripts use it to turn the
// alarm into an actionable list.
func (r *Repo) DriftedTraitSeeds(ctx context.Context, limit int) ([]string, error) {
	const op = "drifted trait seeds"
	if limit <= 0 {
		return nil, fmt.Errorf("%s: limit must be positive", op)
	}
	const query = `SELECT seed
		FROM cg_token_traits_fetch
		WHERE drift_detected_at IS NOT NULL
		ORDER BY drift_detected_at DESC, seed
		LIMIT $1`
	return queryList(ctx, r, op, limit, query, func(rows pgx.Rows, seed *string) error {
		return rows.Scan(seed)
	}, limit)
}

// CosmicSignatureMintSource resolves which prize path minted one Cosmic
// Signature token, reusing the exactly-one-source rule the v2 ledger
// enforces: a token matching zero or several prize families is a data
// inconsistency, not a token with an ambiguous provenance.
func (r *Repo) CosmicSignatureMintSource(ctx context.Context, tokenID int64) (CosmicSignatureMintSource, error) {
	const op = "cosmic signature mint source"
	const query = `SELECT
			(pc.token_id IS NOT NULL),
			rnw.is_rwalk,
			rnw.is_staker,
			(endu.erc721_token_id IS NOT NULL),
			(stel.erc721_token_id IS NOT NULL),
			(cw.nft_id IS NOT NULL)
		FROM cg_mint_event m
			LEFT JOIN cg_prize_claim pc
				ON m.token_id>=pc.token_id AND m.token_id<pc.token_id+pc.num_cs_nfts
			LEFT JOIN cg_raffle_nft_prize rnw
				ON (rnw.token_id=m.token_id AND rnw.round_num=m.round_num)
			LEFT JOIN cg_endurance_prize endu
				ON (endu.erc721_token_id=m.token_id AND endu.round_num=m.round_num)
			LEFT JOIN cg_lastcst_prize stel
				ON (stel.erc721_token_id=m.token_id AND stel.round_num=m.round_num)
			LEFT JOIN cg_chrono_warrior_prize cw
				ON (cw.nft_id=m.token_id AND cw.round_num=m.round_num)
		WHERE m.token_id=$1`

	var (
		isMainPrize   bool
		raffleIsRWalk sql.NullBool
		raffleStaker  sql.NullBool
		isEndurance   bool
		isLastCst     bool
		isChrono      bool
	)
	err := r.q(ctx).QueryRow(ctx, query, tokenID).Scan(
		&isMainPrize, &raffleIsRWalk, &raffleStaker, &isEndurance, &isLastCst, &isChrono,
	)
	if err != nil {
		return "", store.WrapError(op, err)
	}
	source, err := deriveMintSource(tokenID, isMainPrize, raffleIsRWalk, raffleStaker, isEndurance, isLastCst, isChrono)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return source, nil
}

// maxTraitFetchErrorLen bounds what a remote failure can write into the
// bookkeeping table; the full text is already in the process log.
const maxTraitFetchErrorLen = 500

func truncateError(msg string) string {
	if len(msg) <= maxTraitFetchErrorLen {
		return msg
	}
	return msg[:maxTraitFetchErrorLen]
}

// nullString maps the empty string to a SQL NULL so "absent" and "empty
// validator" stay distinguishable in the fetch table.
func nullString(s string) any {
	if s == "" {
		return nil
	}
	return s
}

// nullBytes maps an empty payload to a SQL NULL, which is what lets the
// upsert's COALESCE preserve a previously stored manifest.
func nullBytes(b []byte) any {
	if len(b) == 0 {
		return nil
	}
	return b
}
