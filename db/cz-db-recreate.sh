#!/usr/bin/env bash
# Output a single SQL script that re-creates ONLY the ChaosZero (cz_) schema
# from the goose migration files whose name contains "chaoszero", in order,
# using only their "Up" sections. The shared layer-1 tables (evt_log, block,
# transaction, address, ...) are left untouched; create them first with
# db-recreate.sh (or goose) on a fresh database.
#
# Before creating anything it drops exactly the tables those files create
# (the drop list is derived from the files themselves, so it cannot go
# stale). Dropping a table CASCADE removes its indexes and serial sequences.
#
# Does NOT execute anything; pipe the output into psql yourself:
#
#   ./cz-db-recreate.sh | psql -h localhost -U cosmicgame -d cosmicgame
#
# All ChaosZero data is lost, including the cz_contracts registry row and
# the cz_proc_status watermark; the migration re-seeds both with the
# production defaults. For a local hardhat/anvil deployment, apply the
# UPDATE statements that cmd/cz-etl/dev-scripts/deploy-local.js prints
# (local addresses + last_block = 0) before starting cz-etl.
#
# Note: replaying migrations this way bypasses goose bookkeeping. On a
# goose-managed database prefer `goose down/up` instead.

set -euo pipefail
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MIGRATIONS_DIR="$SCRIPT_DIR/migrations"

shopt -s nullglob
CZ_FILES=("$MIGRATIONS_DIR"/*chaoszero*.sql)
if [ ${#CZ_FILES[@]} -eq 0 ]; then
	echo "ERROR: no *chaoszero*.sql migration files in $MIGRATIONS_DIR" >&2
	exit 1
fi

# up_section prints the lines between "-- +goose Up" and "-- +goose Down".
up_section() {
	awk '/^-- \+goose Down/{exit} up{print} /^-- \+goose Up/{up=1}' "$1"
}

cat <<'SQL'
-- Quiet replay: hide the drop-cascade inventory so real problems stand out.
SET client_min_messages = WARNING;

SQL

echo "-- ==== drop the ChaosZero tables (generated from the files below) ===="
for f in "${CZ_FILES[@]}"; do
	up_section "$f"
done | grep -oE 'CREATE TABLE (IF NOT EXISTS )?[a-zA-Z0-9_]+' | awk '{print $NF}' | sort -u |
	while read -r t; do
		echo "DROP TABLE IF EXISTS $t CASCADE;"
	done
echo

for f in "${CZ_FILES[@]}"; do
	echo "-- ==== $(basename "$f") ===="
	up_section "$f"
	echo
done
