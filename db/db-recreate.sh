#!/usr/bin/env bash
# Output a single SQL script that re-creates the RWCG schema (all tables,
# trigger functions, triggers, indices, seed rows) from the goose migration
# files, in order, using only their "Up" sections.
#
# Before creating anything it drops exactly the tables those files create
# (the drop list is derived from the files themselves, so it cannot go
# stale). Dropping a table CASCADE removes its indexes, triggers and serial
# sequences; the trigger functions are all CREATE OR REPLACE, so they need
# no explicit drops. Objects that are not part of this schema are left
# untouched.
#
# Does NOT execute anything; pipe the output into psql yourself:
#
#   ./cat-recreate.sh | psql -h localhost -U cosmicgame -d cosmicgame
#
# All schema data is lost, including the cg_contracts/rw_contracts registry
# rows; re-insert them before starting the ETLs (populate.js prints the
# INSERT).
#
# Note: a database initialized this way has no goose_db_version table, so
# `make migrate-up` / goose cannot manage it later. For a goose-managed
# database apply db/migrations with goose instead (see docs/operations.md).

set -euo pipefail
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MIGRATIONS_DIR="$SCRIPT_DIR/migrations"

# up_section prints the lines between "-- +goose Up" and "-- +goose Down".
up_section() {
	awk '/^-- \+goose Down/{exit} up{print} /^-- \+goose Up/{up=1}' "$1"
}

cat <<'SQL'
-- Quiet replay: later migration files re-create a few objects with
-- IF NOT EXISTS, which is expected here; hide those NOTICEs (and the
-- drop-cascade inventory) so real problems stand out.
SET client_min_messages = WARNING;

SQL

echo "-- ==== drop the tables this schema owns (generated from the files below) ===="
for f in "$MIGRATIONS_DIR"/*.sql; do
	up_section "$f"
done | grep -oE 'CREATE TABLE (IF NOT EXISTS )?[a-zA-Z0-9_]+' | awk '{print $NF}' | sort -u |
	while read -r t; do
		echo "DROP TABLE IF EXISTS $t CASCADE;"
	done
echo

for f in "$MIGRATIONS_DIR"/*.sql; do
	echo "-- ==== $(basename "$f") ===="
	up_section "$f"
	echo
done
