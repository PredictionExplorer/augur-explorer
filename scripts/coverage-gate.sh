#!/usr/bin/env bash
set -euo pipefail

repo_root=$(git rev-parse --show-toplevel)
cd "$repo_root"

policy="coverage/policy.json"
profile=""
diff_file=""
report=""
use_staged_diff=true
run_tests=true

usage() {
	echo "usage: $0 [--profile FILE] [--diff FILE|--no-diff] [--report FILE] [--policy FILE]" >&2
}

while (($# > 0)); do
	case "$1" in
		--profile)
			[[ $# -ge 2 ]] || { usage; exit 2; }
			profile=$2
			run_tests=false
			shift 2
			;;
		--diff)
			[[ $# -ge 2 ]] || { usage; exit 2; }
			diff_file=$2
			use_staged_diff=false
			shift 2
			;;
		--no-diff)
			use_staged_diff=false
			shift
			;;
		--report)
			[[ $# -ge 2 ]] || { usage; exit 2; }
			report=$2
			shift 2
			;;
		--policy)
			[[ $# -ge 2 ]] || { usage; exit 2; }
			policy=$2
			shift 2
			;;
		-h|--help)
			usage
			exit 0
			;;
		*)
			usage
			exit 2
			;;
	esac
done

command -v go >/dev/null || { echo "coverage gate: go is required" >&2; exit 2; }
command -v git >/dev/null || { echo "coverage gate: git is required" >&2; exit 2; }

cache_dir=$(git rev-parse --git-path coverage-gate)
mkdir -p "$cache_dir"

if [[ "$run_tests" == true ]]; then
	command -v docker >/dev/null || { echo "coverage gate: Docker is required" >&2; exit 2; }
	docker info >/dev/null 2>&1 || {
		echo "coverage gate: Docker daemon is unavailable" >&2
		exit 2
	}

	if ! git diff --quiet -- '*.go' go.mod go.sum 'db/migrations/*.sql' "$policy"; then
		echo "coverage gate: unstaged coverable changes would make staged coverage ambiguous" >&2
		echo "stage or stash those changes before committing" >&2
		exit 2
	fi
	if [[ -n $(git ls-files --others --exclude-standard -- '*.go') ]]; then
		echo "coverage gate: untracked Go files must be staged or removed" >&2
		exit 2
	fi

	cache_key=$(
		{
			git ls-files -s -- '*.go' go.mod go.sum 'db/migrations/*.sql' \
				'internal/testfixtures/**/*.sql' "$policy"
			go version
		} | git hash-object --stdin
	)
	profile="$cache_dir/$cache_key.out"
	if [[ ! -s "$profile" ]]; then
		tmp_profile="$profile.tmp.$$"
		trap 'rm -f "${tmp_profile:-}"' EXIT
		race_args=()
		if [[ "${COVERAGE_RACE:-0}" == "1" ]]; then
			race_args=(-race)
		fi
		echo "coverage gate: generating integration profile (cache $cache_key)"
		# The ${arr[@]+...} form keeps `set -u` happy on bash 3.2 (macOS
		# default) when the array is empty.
		go test ${race_args[@]+"${race_args[@]}"} -shuffle=on -tags=integration -timeout 20m \
			-covermode=atomic \
			-coverprofile="$tmp_profile" \
			-coverpkg=./cmd/...,./internal/... \
			./...
		mv "$tmp_profile" "$profile"
		trap - EXIT
	else
		echo "coverage gate: reusing cached integration profile $cache_key"
	fi
fi

[[ -n "$profile" && -s "$profile" ]] || {
	echo "coverage gate: coverage profile is missing or empty: $profile" >&2
	exit 2
}

if [[ "$use_staged_diff" == true ]]; then
	diff_file="$cache_dir/staged.diff"
	git diff --cached --unified=0 --no-color --diff-filter=ACMR -- '*.go' >"$diff_file"
fi

if [[ -z "$report" ]]; then
	report="$cache_dir/report.json"
fi

args=(-policy "$policy" -profile "$profile" -json "$report")
if [[ -n "$diff_file" ]]; then
	args+=(-diff "$diff_file")
fi

go run ./cmd/covergate "${args[@]}"
echo "coverage gate: report written to $report"
