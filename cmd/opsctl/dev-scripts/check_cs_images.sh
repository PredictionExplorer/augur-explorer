#!/usr/bin/env bash
#
# check_cs_images.sh — verify Cosmic Signature token IMAGE assets return HTTP 200.
#
# Companion to `opsctl assets inventory` / `opsctl assets gen-thumbnails`:
# those check the assets on disk, this script checks the public URLs the
# frontend actually requests. For every minted token (seed) in the database it
# requests the three image URLs (videos are intentionally skipped):
#
#   thumbnail (micro) : /images/new/cosmicsignature/0x<seed>/thumb_micro.webp
#   thumbnail (card)  : /images/new/cosmicsignature/0x<seed>/thumb_card.webp
#   full image        : /images/new/cosmicsignature/0x<seed>.png
#
# Each request is graded success (HTTP 200, green) or failure (anything else,
# red). A per-token line is printed plus a final summary; a non-zero exit code
# is returned when any request failed.
#
# Usage:
#   source ~/configs/cg-prod.env   # (optional; script also sources --env)
#   ./check_cs_images.sh
#   ./check_cs_images.sh --limit 25            # only first 25 tokens
#   ./check_cs_images.sh --base-url https://nfts.cosmicsignature.com
#   ./check_cs_images.sh --env ~/configs/cg-prod.env --jobs 8 --fail-only
#
# Requires: psql, curl. DB connection comes from PGSQL_HOST (host[:port]),
# PGSQL_USERNAME, PGSQL_DATABASE, PGSQL_PASSWORD (read from --env or environment).

set -uo pipefail

# ---- defaults ---------------------------------------------------------------
ENV_FILE="${ENV_FILE:-$HOME/configs/cg-prod.env}"
BASE_URL="${BASE_URL:-https://nfts.cosmicsignature.com}"
SCHEMA="${SCHEMA:-public}"
LIMIT=0
TIMEOUT=15
JOBS=1
FAIL_ONLY=0

usage() {
	sed -n '2,32p' "$0" | sed 's/^# \{0,1\}//'
	exit "${1:-0}"
}

# ---- args -------------------------------------------------------------------
while [[ $# -gt 0 ]]; do
	case "$1" in
		--env)       ENV_FILE="$2"; shift 2 ;;
		--base-url)  BASE_URL="$2"; shift 2 ;;
		--schema)    SCHEMA="$2"; shift 2 ;;
		--limit)     LIMIT="$2"; shift 2 ;;
		--timeout)   TIMEOUT="$2"; shift 2 ;;
		--jobs)      JOBS="$2"; shift 2 ;;
		--fail-only) FAIL_ONLY=1; shift ;;
		-h|--help)   usage 0 ;;
		*) echo "unknown argument: $1" >&2; usage 1 ;;
	esac
done

# ---- colors (only when stdout is a terminal) --------------------------------
if [[ -t 1 ]]; then
	GREEN=$'\e[32m'; RED=$'\e[31m'; YELLOW=$'\e[33m'; CYAN=$'\e[36m'; BOLD=$'\e[1m'; DIM=$'\e[2m'; RESET=$'\e[0m'
else
	GREEN=""; RED=""; YELLOW=""; CYAN=""; BOLD=""; DIM=""; RESET=""
fi

die() { echo "${RED}error:${RESET} $*" >&2; exit 1; }

# ---- env / deps -------------------------------------------------------------
if [[ -f "$ENV_FILE" ]]; then
	# shellcheck disable=SC1090
	source "$ENV_FILE"
fi

command -v psql >/dev/null 2>&1 || die "psql not found on PATH"
command -v curl >/dev/null 2>&1 || die "curl not found on PATH"

: "${PGSQL_HOST:?PGSQL_HOST not set (source your env file or pass --env)}"
: "${PGSQL_USERNAME:?PGSQL_USERNAME not set}"
: "${PGSQL_DATABASE:?PGSQL_DATABASE not set}"
: "${PGSQL_PASSWORD:?PGSQL_PASSWORD not set}"

# PGSQL_HOST may be "host" or "host:port".
if [[ "$PGSQL_HOST" == *:* ]]; then
	PGHOST="${PGSQL_HOST%:*}"
	PGPORT="${PGSQL_HOST##*:}"
else
	PGHOST="$PGSQL_HOST"
	PGPORT="${PGSQL_PORT:-5432}"
fi

BASE_URL="${BASE_URL%/}"   # strip trailing slash

# ---- fetch seeds ------------------------------------------------------------
# DISTINCT ON (token_id) guards against duplicate mint-event rows so the token
# count matches the chain. Note token_id may be 0-based (ids 0..N-1).
QUERY="SELECT DISTINCT ON (token_id) token_id, lower(regexp_replace(seed, '^0x', '')) AS seed
       FROM ${SCHEMA}.cg_mint_event
       WHERE seed IS NOT NULL AND seed <> ''
       ORDER BY token_id"
if [[ "$LIMIT" =~ ^[0-9]+$ && "$LIMIT" -gt 0 ]]; then
	QUERY="$QUERY LIMIT $LIMIT"
fi

echo "${BOLD}Cosmic Signature image check${RESET}"
echo "  base url : ${CYAN}${BASE_URL}${RESET}"
echo "  database : ${PGSQL_USERNAME}@${PGHOST}:${PGPORT}/${PGSQL_DATABASE} (schema ${SCHEMA})"
echo "  parallel : ${JOBS} job(s), timeout ${TIMEOUT}s"
echo

# Keep stderr OUT of the data: a psql WARNING/NOTICE on stderr must not become a
# fake row. Capture stdout to a var and stderr to a temp file.
DB_ERR="$(mktemp)"
DB_OUT="$(PGPASSWORD="$PGSQL_PASSWORD" psql -h "$PGHOST" -p "$PGPORT" \
	-U "$PGSQL_USERNAME" -d "$PGSQL_DATABASE" -tA -F'|' -c "$QUERY" 2>"$DB_ERR")"
DB_RC=$?
if [[ "$DB_RC" -ne 0 ]]; then
	msg="$(cat "$DB_ERR")"; rm -f "$DB_ERR"
	die "database query failed: ${msg}"
fi
if [[ -s "$DB_ERR" ]]; then
	echo "${YELLOW}psql:${RESET} $(cat "$DB_ERR")" >&2
fi
rm -f "$DB_ERR"

# Validate every row (token_id numeric, seed = hex). Anything else is dropped
# with a warning, so stray output can never inflate the token count.
ROWS=()
while IFS= read -r _line; do
	[[ -z "$_line" ]] && continue
	if [[ "$_line" =~ ^[0-9]+\|[0-9a-f]+$ ]]; then
		ROWS+=("$_line")
	else
		echo "${YELLOW}skipping unexpected query line:${RESET} ${_line}" >&2
	fi
done <<<"$DB_OUT"

if [[ "${#ROWS[@]}" -eq 0 ]]; then
	die "no minted tokens with a seed found"
fi

# Report the token-id range so 0-based numbering is obvious (rows are ordered).
first_id="${ROWS[0]%%|*}"
last_id="${ROWS[-1]%%|*}"
echo "Found ${BOLD}${#ROWS[@]}${RESET} distinct token(s) to check (3 image requests each)."
echo "  token_id range: ${BOLD}${first_id}${RESET}..${BOLD}${last_id}${RESET}"
echo

# ---- worker that checks one token -------------------------------------------
# Prints a human line (unless --fail-only and all passed) plus a machine line:
#   RESULT|<token_id>|<ok_count>|<fail_count>
check_token() {
	local row="$1"
	local token_id seed
	IFS='|' read -r token_id seed <<<"$row"
	[[ -z "$seed" ]] && { echo "RESULT|${token_id:-?}|0|0"; return; }

	local short="0x${seed:0:6}..${seed: -4}"
	local -a keys=(micro card full)
	local -A urls=(
		[micro]="${BASE_URL}/images/new/cosmicsignature/0x${seed}/thumb_micro.webp"
		[card]="${BASE_URL}/images/new/cosmicsignature/0x${seed}/thumb_card.webp"
		[full]="${BASE_URL}/images/new/cosmicsignature/0x${seed}.png"
	)

	local ok=0 fail=0 line cell key code
	line="$(printf '%s#%-6s%s %s' "$BOLD" "$token_id" "$RESET" "$short")"
	for key in "${keys[@]}"; do
		code="$(curl -s -o /dev/null -w '%{http_code}' --max-time "$TIMEOUT" "${urls[$key]}" 2>/dev/null)"
		if [[ "$code" == "200" ]]; then
			ok=$((ok+1))
			cell="${GREEN}${key}:OK${RESET}"
		else
			fail=$((fail+1))
			cell="${RED}${key}:FAIL(${code:-000})${RESET}"
		fi
		line+="  ${cell}"
	done

	if [[ "$FAIL_ONLY" -eq 0 || "$fail" -gt 0 ]]; then
		printf '%s\n' "$line"
	fi
	echo "RESULT|${token_id}|${ok}|${fail}"
}
export -f check_token
export BASE_URL TIMEOUT FAIL_ONLY BOLD RESET GREEN RED

# ---- run (capture once, then print + tally) ---------------------------------
TMP_OUT="$(mktemp)"
trap 'rm -f "$TMP_OUT"' EXIT

if [[ "$JOBS" =~ ^[0-9]+$ && "$JOBS" -gt 1 ]]; then
	printf '%s\n' "${ROWS[@]}" \
		| xargs -d '\n' -P "$JOBS" -I {} bash -c 'check_token "$@"' _ {} >"$TMP_OUT"
else
	for row in "${ROWS[@]}"; do
		check_token "$row"
	done >"$TMP_OUT"
fi

# Print the human-readable token lines (everything except RESULT records).
grep -v '^RESULT|' "$TMP_OUT" || true

# Tally from the RESULT records.
total_ok=0
total_fail=0
tokens_checked=0
tokens_failed=0
while IFS='|' read -r _ _tid oc fc; do
	tokens_checked=$((tokens_checked + 1))
	total_ok=$((total_ok + oc))
	total_fail=$((total_fail + fc))
	[[ "$fc" -gt 0 ]] && tokens_failed=$((tokens_failed + 1))
done < <(grep '^RESULT|' "$TMP_OUT")

total_req=$((total_ok + total_fail))

echo
echo "${BOLD}Summary${RESET}"
printf '  tokens checked      : %d\n' "$tokens_checked"
printf '  image requests      : %d\n' "$total_req"
printf '  %ssuccess (200)%s      : %d\n' "$GREEN" "$RESET" "$total_ok"
printf '  %sfailure (non-200)%s  : %d\n' "$RED" "$RESET" "$total_fail"
printf '  tokens with failures: %s%d%s\n' \
	"$([[ "$tokens_failed" -gt 0 ]] && echo "$RED" || echo "$GREEN")" "$tokens_failed" "$RESET"

if [[ "$total_fail" -gt 0 ]]; then
	echo
	echo "${RED}${BOLD}RESULT: FAILURE${RESET} — ${total_fail} request(s) did not return 200."
	exit 1
fi
echo
echo "${GREEN}${BOLD}RESULT: SUCCESS${RESET} — all ${total_req} image requests returned 200."
exit 0
