#!/usr/bin/env bash
# Run every native Go fuzz target in the repository, one at a time
# (go test accepts only a single -fuzz pattern per invocation).
#
# Usage: scripts/fuzz-all.sh [fuzztime]
#   fuzztime  per-target fuzzing duration (go duration or Nx execs), default 10s
#
# Exit code is non-zero if any target found a failing input. Crashers are
# written by the Go toolchain to <pkg>/testdata/fuzz/<FuzzName>/ — commit the
# minimized input as a permanent regression seed after fixing the bug.
set -u

FUZZTIME="${1:-10s}"
fail=0
total=0

for pkg in $(go list ./...); do
    targets=$(go test -list '^Fuzz' "$pkg" 2>/dev/null | grep '^Fuzz' || true)
    [ -z "$targets" ] && continue
    for t in $targets; do
        total=$((total + 1))
        echo "=== fuzz ${pkg#*augur-explorer/} ${t} (${FUZZTIME})"
        if ! go test -run='^$' -fuzz="^${t}\$" -fuzztime="$FUZZTIME" "$pkg"; then
            echo "FUZZFAIL ${pkg} ${t}"
            fail=1
        fi
    done
done

echo "fuzz-all: ${total} targets, fuzztime ${FUZZTIME} each"
if [ "$fail" -ne 0 ]; then
    echo "fuzz-all: FAILURES FOUND (see FUZZFAIL lines above; crashers under testdata/fuzz/)"
fi
exit "$fail"
