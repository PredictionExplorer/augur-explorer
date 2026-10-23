#!/usr/bin/env bash
# Chain-vs-database audit: thin wrapper around `opsctl validate-chain`
# (cmd/opsctl/cmd_validatechain.go), which replaced scripts/validate-chain.py.
#
# Environment: RPC_URL plus the same PGSQL_*/DATABASE_URL variables the other
# validate-*.sh scripts use. Optional modes and tuning are opsctl flags
# (--archive-dir/--archive-config, --full-log-parity, --state-file,
# --log-workers, --log-chunk); the old ARCHIVE_DIR, ARCHIVE_CONFIG,
# FULL_LOG_PARITY, VALIDATE_CHAIN_STATE, LOG_WORKERS and LOG_CHUNK environment
# variables still work as flag defaults. All arguments are passed through.
#
# Uses the already-built cmd/opsctl/opsctl binary when present, otherwise
# builds it first (requires the Go toolchain).
set -euo pipefail

REPO_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)"
BIN="$REPO_DIR/cmd/opsctl/opsctl"

if [[ ! -x "$BIN" ]]; then
    echo "building opsctl..." >&2
    (cd "$REPO_DIR" && go build -o cmd/opsctl/opsctl ./cmd/opsctl)
fi

exec "$BIN" validate-chain "$@"
