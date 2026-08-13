#!/usr/bin/env bash
# validate-db.sh -- run the game-mechanics database validator (validate-db.sql)
# using the same environment variables as cg-etl (internal/store.ConfigFromEnv
# / internal/toolutil.PostgresConnStringFromEnv):
#
#   DATABASE_URL    complete postgres:// URL; when set it wins over PGSQL_*.
#   PGSQL_HOST      "host" or "host:port"; empty selects the local Unix socket.
#   PGSQL_USERNAME  database user.
#   PGSQL_PASSWORD  database password (optional).
#   PGSQL_DATABASE  database name.
#
# An optional first argument selects a different validator SQL file (used by
# validate-stats.sh to run the statistical-counters validator).
#
# Like the ETL, connections built from PGSQL_* use sslmode=disable.
# Exits non-zero if any ERROR-severity check finds violations.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SQL_FILE="${1:-$SCRIPT_DIR/validate-db.sql}"

PSQL="${PSQL:-psql}"

if [[ -n "${DATABASE_URL:-}" ]]; then
	exec "$PSQL" "$DATABASE_URL" -v ON_ERROR_STOP=1 -f "$SQL_FILE"
fi

if [[ -z "${PGSQL_DATABASE:-}" ]]; then
	echo "error: DATABASE_URL, or PGSQL_HOST, PGSQL_USERNAME, and PGSQL_DATABASE must be set" >&2
	exit 2
fi

args=(-d "$PGSQL_DATABASE")

if [[ -n "${PGSQL_HOST:-}" ]]; then
	host="$PGSQL_HOST"
	port=""
	if [[ "$host" == *:* ]]; then
		port="${host##*:}"
		host="${host%%:*}"
	fi
	args+=(-h "$host")
	[[ -n "$port" ]] && args+=(-p "$port")
	export PGSSLMODE=disable
fi

[[ -n "${PGSQL_USERNAME:-}" ]] && args+=(-U "$PGSQL_USERNAME")
[[ -n "${PGSQL_PASSWORD:-}" ]] && export PGPASSWORD="$PGSQL_PASSWORD"

exec "$PSQL" "${args[@]}" -v ON_ERROR_STOP=1 -f "$SQL_FILE"
