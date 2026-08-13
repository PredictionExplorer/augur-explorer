#!/usr/bin/env bash
# validate-stats.sh -- run the statistical-counters validator
# (validate-stats.sql). Reads the same environment variables as cg-etl
# (DATABASE_URL, or PGSQL_HOST / PGSQL_USERNAME / PGSQL_PASSWORD /
# PGSQL_DATABASE); see validate-db.sh for details.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
exec "$SCRIPT_DIR/validate-db.sh" "$SCRIPT_DIR/validate-stats.sql"
