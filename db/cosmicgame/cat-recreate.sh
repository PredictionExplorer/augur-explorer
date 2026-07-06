#!/usr/bin/env bash
# Output a single SQL script that: drops cosmicgame tables, creates all tables,
# creates trigger functions, triggers, and indices (full DB reset).
# Does NOT execute anything; use: ./cat-recreate.sh > reset.sql && psql ... -f reset.sql

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

# 1) DROP TABLE statements
cat drop-cosmicgame.sql | ensure_semicolon
echo

# 2) CREATE TABLE (mechanics, management, stats)
cat game-mechanics.sql game-management.sql game-stats.sql | ensure_semicolon
echo

# 3) Trigger functions
cat cosmicgame-funcs.sql | ensure_semicolon
echo

# 4) Triggers
cat triggers-cosmicgame.sql | ensure_semicolon
echo

# 5) Indices
cat indices_cosmicgame.sql | ensure_semicolon
