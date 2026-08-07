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

while IFS=$'\t' read -r created message; do
	[ -n "$created" ] || continue
	epoch=$(date -d "$created" +%s 2>/dev/null) || continue
	age=$((now - epoch))
	if [ "$age" -le "$WINDOW_SECS" ]; then
		echo "CRASH $created $message"
		exit 2
	fi
done <<<"$candidates"

echo "OK"
