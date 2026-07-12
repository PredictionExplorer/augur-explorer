#!/usr/bin/env bash
set -euo pipefail

repo_root=$(git rev-parse --show-toplevel)
source_hook="$repo_root/.githooks/pre-commit"
target_hook=$(git rev-parse --git-path hooks/pre-commit)
marker="# augur-explorer managed coverage hook"

if [[ -e "$target_hook" ]] && ! grep -Fq "$marker" "$target_hook"; then
	echo "hooks install: refusing to overwrite unrelated hook at $target_hook" >&2
	exit 1
fi

mkdir -p "$(dirname "$target_hook")"
cp "$source_hook" "$target_hook"
chmod +x "$target_hook"
echo "hooks install: installed coverage gate at $target_hook"
