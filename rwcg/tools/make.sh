#!/usr/bin/env bash
# Build all rwcg/tools binaries into this directory (same names as make-clean.sh).
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

REPO_ROOT="$SCRIPT_DIR"
while [[ "$REPO_ROOT" != "/" && ! -f "$REPO_ROOT/go.mod" ]]; do
	REPO_ROOT="$(dirname "$REPO_ROOT")"
done
if [[ ! -f "$REPO_ROOT/go.mod" ]]; then
	echo "error: could not find go.mod above $SCRIPT_DIR" >&2
	exit 1
fi

cd "$REPO_ROOT"

build() {
	local src="$1"
	local out="$2"
	echo "building $out <- rwcg/tools/$src"
	go build -o "$SCRIPT_DIR/$out" "./rwcg/tools/$src"
}

build archive_export.go archive_export
build arch_verify.go arch_verify
build arch_node_fill.go arch_node_fill
build db_verify.go db_verify
build evtlog_diff.go evtlog_diff
build transaction_collector.go transaction-collector
build transaction_collector_verify.go transaction-collector-verify

echo "done — binaries in $SCRIPT_DIR"
