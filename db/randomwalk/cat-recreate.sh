#!/usr/bin/env bash
# Output a single SQL script that: drops randomwalk tables, creates all tables,
# ranking tables, trigger functions, triggers, and indices (full schema reset).
# Does NOT execute anything; use: ./cat-recreate.sh > reset-randomwalk.sql && psql ... -f reset-randomwalk.sql
#
# Prerequisite: shared explorer tables (e.g. evt_log, transaction, address) must
# already exist in the target DB — same assumption as ETL in rwcg/etl/randomwalk.

set -e
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# Ensure the given content ends with ; (if non-empty and not already ending with ;)
ensure_semicolon() {
	local content trimmed
	content=$(cat)
	# trim trailing whitespace (Bash parameter expansion)
	trimmed="${content%"${content##*[![:space:]]}"}"
	if [ -n "$trimmed" ]; then
		printf '%s' "$content"
		[ "${trimmed: -1}" = ';' ] || printf ';'
		printf '\n'
	fi
}

# 1) DROP TABLE statements (order matches FKs / CASCADE)
cat drop-randomwalk.sql | ensure_semicolon
echo

# 2) CREATE TABLE — core RandomWalk / marketplace state
cat tables_randomwalk.sql | ensure_semicolon
echo

# 3) Ranking / Elo tables (depends on address for voter_aid FK)
cat token_ranking.sql | ensure_semicolon
echo

# 4) Wallet vote nonce migration (idempotent; apply after token_ranking.sql)
cat ranking_vote_wallet.sql | ensure_semicolon
echo

# 5) Trigger functions (before CREATE TRIGGER)
cat trigger-funcs.sql | ensure_semicolon
echo

# 6) Triggers
cat triggers.sql | ensure_semicolon
echo

# 7) Indices not defined inline elsewhere
cat indices_rwalk.sql | ensure_semicolon
