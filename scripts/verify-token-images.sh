#!/usr/bin/env bash
#
# verify-token-images.sh — verify every Cosmic Signature token image is
# generated and served.
#
# Walks the public API's token directory (GET
# /api/v2/cosmicgame/cosmic-signature-tokens, cursor-paginated) and requests
# each minted token's three image assets:
#
#   thumbnail (micro) : <images>/new/cosmicsignature/0x<seed>/thumb_micro.webp
#   thumbnail (card)  : <images>/new/cosmicsignature/0x<seed>/thumb_card.webp
#   full image        : <images>/new/cosmicsignature/0x<seed>.png
#
# Unlike cmd/opsctl/dev-scripts/check_cs_images.sh (which reads seeds from
# the database via psql), this script needs only the public API: point it at
# an apiserver URL and it discovers the token list itself.
#
# Usage:
#   ./scripts/verify-token-images.sh API_URL [IMAGES_URL]
#
#   API_URL      apiserver base URL, e.g. https://a1.cosmicsignature.com
#   IMAGES_URL   image server base URL (optional). Defaults to API_URL —
#                the apiserver serves assets under /images on the same host
#                unless NFT_ASSETS_PUBLIC_BASE points elsewhere (CDN etc.).
#                "/images" is appended automatically when missing.
#
# Environment overrides:
#   JOBS=8        parallel image requests (default 4)
#   TIMEOUT=15    per-request timeout in seconds (default 15)
#   LIMIT=25      check only the first N tokens (default 0 = all)
#   FAIL_ONLY=1   print only tokens with at least one failing image
#
# Examples:
#   ./scripts/verify-token-images.sh https://a1.cosmicsignature.com
#   JOBS=8 FAIL_ONLY=1 ./scripts/verify-token-images.sh http://161.129.176.16:2121
#
# Requires: curl, python3. Exits 0 when every image answers HTTP 200.

set -uo pipefail

TIMEOUT="${TIMEOUT:-15}"
JOBS="${JOBS:-4}"
LIMIT="${LIMIT:-0}"
FAIL_ONLY="${FAIL_ONLY:-0}"

if [[ $# -lt 1 || $# -gt 2 ]]; then
	sed -n '2,37p' "$0" | sed 's/^# \{0,1\}//'
	exit 1
fi

API_URL="${1%/}"
IMAGES_URL="${2:-${IMAGES_URL:-$API_URL}}"
IMAGES_URL="${IMAGES_URL%/}"
[[ "$IMAGES_URL" == */images ]] || IMAGES_URL="${IMAGES_URL}/images"

# ---- colors (only when stdout is a terminal) --------------------------------
if [[ -t 1 ]]; then
	GREEN=$'\e[32m'; RED=$'\e[31m'; YELLOW=$'\e[33m'; CYAN=$'\e[36m'; BOLD=$'\e[1m'; RESET=$'\e[0m'
else
	GREEN=""; RED=""; YELLOW=""; CYAN=""; BOLD=""; RESET=""
fi

die() { echo "${RED}error:${RESET} $*" >&2; exit 1; }

command -v curl >/dev/null 2>&1 || die "curl not found on PATH"
command -v python3 >/dev/null 2>&1 || die "python3 not found on PATH"

echo "${BOLD}Cosmic Signature image check${RESET}"
echo "  api url    : ${CYAN}${API_URL}${RESET}"
echo "  images url : ${CYAN}${IMAGES_URL}${RESET}"
echo "  parallel   : ${JOBS} job(s), timeout ${TIMEOUT}s"
echo

# ---- fetch the token list from the API ---------------------------------------
# Pages through /api/v2/cosmicgame/cosmic-signature-tokens and prints one
# "token_id|seed" line per token (seed lowercased, 0x prefix stripped).
# On any API problem prints a single "FETCH_ERROR|<reason>" line.
fetch_tokens() {
	python3 - "$API_URL" "$TIMEOUT" <<'PY'
import json
import ssl
import sys
import urllib.parse
import urllib.request

api_url, timeout = sys.argv[1], float(sys.argv[2])
ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE  # mirrors the curl -k used for self-signed prod certs

cursor = None
while True:
    url = f"{api_url}/api/v2/cosmicgame/cosmic-signature-tokens?limit=200"
    if cursor:
        url += "&cursor=" + urllib.parse.quote(cursor, safe="")
    try:
        with urllib.request.urlopen(url, timeout=timeout, context=ctx) as resp:
            page = json.load(resp)
    except Exception as exc:
        print(f"FETCH_ERROR|{url}: {exc}")
        sys.exit(0)
    if not isinstance(page, dict) or not isinstance(page.get("data"), list):
        print(f"FETCH_ERROR|unexpected response from {url}: {str(page)[:120]}")
        sys.exit(0)
    for tok in page["data"]:
        seed = str(tok.get("seed", "")).lower().removeprefix("0x")
        print(f"{tok['nftTokenId']}|{seed}")
    cursor = page.get("meta", {}).get("nextCursor")
    if not cursor:
        break
PY
}

ROWS=()
while IFS= read -r line; do
	[[ -z "$line" ]] && continue
	if [[ "$line" == FETCH_ERROR\|* ]]; then
		die "${line#FETCH_ERROR|}"
	elif [[ "$line" =~ ^[0-9]+\|[0-9a-f]+$ ]]; then
		ROWS+=("$line")
	else
		echo "${YELLOW}skipping unexpected line:${RESET} ${line}" >&2
	fi
done < <(fetch_tokens)

[[ "${#ROWS[@]}" -gt 0 ]] || die "no minted tokens with a seed returned by the API"

if [[ "$LIMIT" =~ ^[0-9]+$ && "$LIMIT" -gt 0 && "${#ROWS[@]}" -gt "$LIMIT" ]]; then
	ROWS=("${ROWS[@]:0:$LIMIT}")
fi

first_id="${ROWS[0]%%|*}"
last_id="${ROWS[-1]%%|*}"
echo "Found ${BOLD}${#ROWS[@]}${RESET} token(s) to check (3 image requests each)."
echo "  token_id range: ${BOLD}${last_id}${RESET}..${BOLD}${first_id}${RESET} (API lists newest first)"
echo

# ---- worker that checks one token --------------------------------------------
# Prints a human line (unless FAIL_ONLY and all passed) plus a machine line:
#   RESULT|<token_id>|<ok_count>|<fail_count>
check_token() {
	local row="$1"
	local token_id seed
	IFS='|' read -r token_id seed <<<"$row"
	[[ -z "$seed" ]] && { echo "RESULT|${token_id:-?}|0|0"; return; }

	local short="0x${seed:0:6}..${seed: -4}"
	local -a keys=(micro card full)
	local -A urls=(
		[micro]="${IMAGES_URL}/new/cosmicsignature/0x${seed}/thumb_micro.webp"
		[card]="${IMAGES_URL}/new/cosmicsignature/0x${seed}/thumb_card.webp"
		[full]="${IMAGES_URL}/new/cosmicsignature/0x${seed}.png"
	)

	local ok=0 fail=0 line cell key code
	line="$(printf '%s#%-6s%s %s' "$BOLD" "$token_id" "$RESET" "$short")"
	for key in "${keys[@]}"; do
		code="$(curl -sk -o /dev/null -w '%{http_code}' --max-time "$TIMEOUT" "${urls[$key]}" 2>/dev/null)"
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
export IMAGES_URL TIMEOUT FAIL_ONLY BOLD RESET GREEN RED

# ---- run (capture once, then print + tally) -----------------------------------
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

grep -v '^RESULT|' "$TMP_OUT" || true

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
