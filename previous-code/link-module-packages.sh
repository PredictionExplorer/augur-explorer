#!/usr/bin/env bash
# Create symlinks at the repository root so
# github.com/PredictionExplorer/augur-explorer/{package} resolves to
# previous-code/{package} (matches root go.mod module layout).
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

names=(amm contracts dbs etl layer1 primitives server statserv testing tweets uevm wanotif)
for name in "${names[@]}"; do
	ln -sfn "previous-code/${name}" "${name}"
	echo "linked ${name} -> previous-code/${name}"
done
