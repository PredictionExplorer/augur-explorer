#!/usr/bin/env bash
#
# idrac_check.sh -- READ-ONLY iDRAC crash detector for srvmonitor.
#
# Queries the Redfish System Event Log (SEL) of one iDRAC and reports
# whether the server recently logged a fatal hardware error (for example
# "A bus fatal error was detected on a component at slot 2"), which halts
# a PowerEdge R640 completely.
#
# This script NEVER changes server state: it performs HTTP GETs only.
# It deliberately contains no reset/power-off/power-on actions -- do not
# add any. Recovery stays a manual operator decision.
#
# Environment:
#   IDRAC                   iDRAC hostname or IP (required)
#   IDRAC_USER              iDRAC user (required)
#   IDRAC_PASS              iDRAC password (required)
#   IDRAC_CRIT_WINDOW_SECS  how far back a Critical SEL entry still counts
#                           as a current crash (default 3600)
#   IDRAC_CRIT_REGEX        case-insensitive message regex marking a
#                           fatal/halting error (default 'fatal error')
#
# Output / exit codes (consumed by srvmonitor's iDRAC monitor):
#   exit 0, line "OK"                          no recent fatal SEL entry
#   exit 2, line "CRASH <created> <message>"   fatal SEL entry in window
#   exit 1 (or other), error text              the check itself failed
set -euo pipefail

: "${IDRAC:?Set IDRAC}"
: "${IDRAC_USER:?Set IDRAC_USER}"
: "${IDRAC_PASS:?Set IDRAC_PASS}"

WINDOW_SECS="${IDRAC_CRIT_WINDOW_SECS:-3600}"
CRIT_REGEX="${IDRAC_CRIT_REGEX:-fatal error}"

SEL_URL="https://$IDRAC/redfish/v1/Managers/iDRAC.Embedded.1/LogServices/Sel/Entries"

# GET only. --insecure because iDRACs ship self-signed certificates.
entries=$(curl --fail --silent --show-error --insecure \
	--max-time 30 --request GET \
	--user "$IDRAC_USER:$IDRAC_PASS" \
	"$SEL_URL")

# Critical entries whose message marks a fatal/halting error, newest first.
candidates=$(jq -r --arg re "$CRIT_REGEX" '
	.Members
	| map(select(.Severity == "Critical")
		| select((.Message // "") | test($re; "i")))
	| sort_by(.Created)
	| reverse
	| .[]
	| [.Created, .Message]
	| @tsv
' <<<"$entries")

now=$(date -u +%s)

# date(1) parses timestamps incompatibly across implementations: GNU (Linux)
# takes -d, BSD (macOS) takes -j -f with an explicit format. Only GNU date
# understands --version, so it doubles as the probe.
if date --version >/dev/null 2>&1; then
	date_is_gnu=1
else
	date_is_gnu=0
fi

# parse_epoch prints the Unix epoch for one RFC 3339 SEL timestamp and
# returns non-zero when it cannot be parsed.
parse_epoch() {
	local ts="$1" normalized
	if [ "$date_is_gnu" -eq 1 ]; then
		date -d "$ts" +%s 2>/dev/null
		return
	fi
	# BSD date rejects the colon inside the numeric zone offset, so
	# 2026-06-18T01:27:24-05:00 has to become ...-0500 first.
	normalized=$(printf '%s' "$ts" | sed -E 's/([+-][0-9]{2}):([0-9]{2})$/\1\2/')
	date -j -f '%Y-%m-%dT%H:%M:%S%z' "$normalized" +%s 2>/dev/null ||
		date -j -u -f '%Y-%m-%dT%H:%M:%SZ' "$normalized" +%s 2>/dev/null
}

parsed=0
unparsed=0
while IFS=$'\t' read -r created message; do
	[ -n "$created" ] || continue
	if ! epoch=$(parse_epoch "$created"); then
		# Never skip quietly. Treating an unparseable timestamp as "no
		# crash" is how this check silently reported OK on every entry
		# when it only spoke GNU date.
		echo "cannot parse SEL timestamp: $created" >&2
		unparsed=$((unparsed + 1))
		continue
	fi
	parsed=$((parsed + 1))
	age=$((now - epoch))
	if [ "$age" -le "$WINDOW_SECS" ]; then
		echo "CRASH $created $message"
		exit 2
	fi
done <<<"$candidates"

# Fatal entries that all failed to parse mean the check could not answer.
# That is a check failure, not a healthy server.
if [ "$parsed" -eq 0 ] && [ "$unparsed" -gt 0 ]; then
	echo "no SEL timestamp could be parsed ($unparsed fatal entries)" >&2
	exit 1
fi

echo "OK"
