-- +goose Up
-- Per-seed art traits produced by the CS-Image-Generation pipeline and
-- ingested from the asset host, plus the bookkeeping the ingest loop needs.
--
-- The two concerns are deliberately separate tables: cg_token_traits holds
-- only facts that reach the served ERC-721 metadata, so a row existing means
-- "this seed can be served enriched". Retry state, ETags and drift alarms
-- live next door and never gate the read path.

CREATE TABLE cg_token_traits (
	seed				TEXT PRIMARY KEY CHECK (seed ~ '^0x[0-9a-f]+$'),
	schema_major		INT NOT NULL CHECK (schema_major >= 0),
	pipeline_version	TEXT NOT NULL,
	attributes			JSONB NOT NULL,
	description_art		TEXT NOT NULL,
	simulation			JSONB NOT NULL,
	generation			JSONB NOT NULL,
	assets				JSONB,
	content_hash		TEXT NOT NULL,
	first_seen_at		TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at			TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE cg_token_traits IS
	'Verbatim nft_traits.json payload per seed; the metadata handler copies these blocks without recomputing them.';
COMMENT ON COLUMN cg_token_traits.seed IS
	'Canonical 0x-prefixed lowercase hex seed, matching cg_mint_event.seed with the prefix added.';
COMMENT ON COLUMN cg_token_traits.content_hash IS
	'SHA-256 over the canonicalized art facts (attributes, description_art, simulation, generation). Art is frozen at imprint, so a changed hash is an incident, not an update.';
COMMENT ON COLUMN cg_token_traits.assets IS
	'assets.json schema_version 2 payload; NULL when the manifest was unavailable. Feeds image_details/animation_details.';

CREATE TABLE cg_token_traits_fetch (
	seed					TEXT PRIMARY KEY CHECK (seed ~ '^0x[0-9a-f]+$'),
	source_etag				TEXT,
	manifest_etag			TEXT,
	attempt_count			INT NOT NULL DEFAULT 0 CHECK (attempt_count >= 0),
	last_attempt_at			TIMESTAMPTZ,
	next_attempt_at			TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	last_status				TEXT,
	last_error				TEXT,
	drift_detected_at		TIMESTAMPTZ,
	drift_pipeline_version	TEXT
);
COMMENT ON TABLE cg_token_traits_fetch IS
	'Ingest bookkeeping per seed. Packages lag mints by minutes to hours, so 404s back off here instead of re-requesting every tick forever.';
COMMENT ON COLUMN cg_token_traits_fetch.next_attempt_at IS
	'Earliest next fetch. Exponential backoff after repeated 404s or errors; a short interval once traits are stored, for drift re-checks.';
COMMENT ON COLUMN cg_token_traits_fetch.drift_detected_at IS
	'Set when a re-fetched trait file disagreed with the stored art facts. The stored row is left untouched; operators must investigate the generator.';

-- The ingest scan is "seeds whose next attempt is due, oldest first".
CREATE INDEX idx_cg_token_traits_fetch_due
	ON cg_token_traits_fetch (next_attempt_at);

-- +goose Down
DROP TABLE cg_token_traits_fetch;
DROP TABLE cg_token_traits;
