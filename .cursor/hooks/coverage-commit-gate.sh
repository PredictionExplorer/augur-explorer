#!/usr/bin/env bash
set -euo pipefail

input=$(</dev/stdin)
policy="${COVERAGE_POLICY:-coverage/policy.json}"
status=$(go run ./cmd/covergate -policy "$policy" -commit-status)
case "$status" in
	deferred\ *)
		echo '{ "permission": "allow" }'
		exit 0
		;;
	enabled\ *)
		;;
	*)
		echo '{
			"permission": "deny",
			"user_message": "Coverage commit policy is invalid.",
			"agent_message": "Fix the coverage policy before committing."
		}'
		exit 0
		;;
esac

if [[ "$input" == *"--no-verify"* ]]; then
	echo '{
		"permission": "deny",
		"user_message": "Coverage policy forbids bypassing the pre-commit gate.",
		"agent_message": "Do not use --no-verify. Run the coverage gate and fix failures."
	}'
	exit 0
fi

tracked_hook="${COVERAGE_TRACKED_HOOK:-.githooks/pre-commit}"
installed_hook="${COVERAGE_INSTALLED_HOOK:-$(git rev-parse --git-path hooks/pre-commit)}"
if [[ ! -x "$installed_hook" ]] || ! cmp -s "$tracked_hook" "$installed_hook"; then
	echo '{
		"permission": "deny",
		"user_message": "Install the repository coverage hook with `make hooks-install` before committing.",
		"agent_message": "The native pre-commit coverage hook is missing or stale. Run `make hooks-install`."
	}'
	exit 0
fi

echo '{ "permission": "allow" }'
