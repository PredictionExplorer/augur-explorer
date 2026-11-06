#!/usr/bin/env bash
#
# verify-rw-token-images.sh — verify RandomWalk token assets end to end,
# starting from the live blockchain.
#
# For each sampled token the script follows the same path an NFT marketplace
# does:
#
#   1. eth_call RandomWalkNFT.tokenURI(tokenId) on the RPC node
#   2. GET the metadata JSON that URI points to
#   3. HEAD every resource the token has — the two URLs in the metadata
#      (image, animation_url) plus the full 8-file variant set derived from
#      the image URL:
#        {white,black} x { .png, _thumb.jpg, _single.mp4, _triple.mp4 }
#      and the external_url detail page.
#
# Token ids are sampled at random from totalSupply() on the same RPC, so the
# script needs no database access.
#
# Usage:
#   ./scripts/verify-rw-token-images.sh [RPC_URL] [COUNT]
#
#   RPC_URL   JSON-RPC endpoint (default https://rpc1.cosmicsignature.com:38547)
#   COUNT     how many random tokens to sample (default 5; "all" = every token)
#
# Environment overrides:
#   RWALK_ADDR=0x...   RandomWalkNFT contract (default: Arbitrum production)
#   TOKEN_IDS="1 2 3"  check exactly these ids instead of sampling
#   JOBS=8             parallel tokens (default 4)
#   TIMEOUT=15         per-request timeout in seconds (default 15)
#   FAIL_ONLY=1        print only tokens with at least one failing request
#   RESOLVE_IP=1.2.3.4 pin the asset/metadata hosts to this server, so each
#                      production box can be checked individually (public DNS
#                      points at only one of them). Only hosts listed in
#                      RESOLVE_HOSTS are pinned; external website URLs keep
#                      normal DNS.
#   RESOLVE_HOSTS="..."  hosts affected by RESOLVE_IP (default:
#                      nfts.randomwalknft.com randomwalknft-api.com
#                      api.randomwalknft.com)
#
# Examples:
#   ./scripts/verify-rw-token-images.sh
#   ./scripts/verify-rw-token-images.sh https://rpc2.cosmicsignature.com:38547 25
#   TOKEN_IDS="0 215 4150" ./scripts/verify-rw-token-images.sh
#   RESOLVE_IP=161.129.67.58 ./scripts/verify-rw-token-images.sh   # cosmic2
#
# Requires: curl, python3. Exits 0 when every request returns HTTP 200.

set -uo pipefail

RPC_URL="${1:-https://rpc1.cosmicsignature.com:38547}"
COUNT="${2:-5}"
RWALK_ADDR="${RWALK_ADDR:-0x895a6F444BE4ba9d124F61DF736605792B35D66b}"
TIMEOUT="${TIMEOUT:-15}"
JOBS="${JOBS:-4}"
FAIL_ONLY="${FAIL_ONLY:-0}"
TOKEN_IDS="${TOKEN_IDS:-}"
RESOLVE_IP="${RESOLVE_IP:-}"
RESOLVE_HOSTS="${RESOLVE_HOSTS:-nfts.randomwalknft.com randomwalknft-api.com api.randomwalknft.com}"

# Emits "--resolve host:port:RESOLVE_IP" (one token per line) when the URL's
# host is in RESOLVE_HOSTS; emits nothing otherwise.
resolve_args() {
	[[ -z "$RESOLVE_IP" ]] && return 0
	local hostport="${1#*://}" host port
	hostport="${hostport%%/*}"
	host="${hostport%%:*}"
	if [[ "$hostport" == *:* ]]; then
		port="${hostport##*:}"
	elif [[ "$1" == http://* ]]; then
		port=80
	else
		port=443
	fi
	local h
	for h in $RESOLVE_HOSTS; do
		[[ "$host" == "$h" ]] || continue
		printf -- '--resolve\n%s:%s:%s\n' "$host" "$port" "$RESOLVE_IP"
		return 0
	done
}

# ---- colors (only when stdout is a terminal) --------------------------------
if [[ -t 1 ]]; then
	GREEN=$'\e[32m'; RED=$'\e[31m'; YELLOW=$'\e[33m'; CYAN=$'\e[36m'; BOLD=$'\e[1m'; RESET=$'\e[0m'
else
	GREEN=""; RED=""; YELLOW=""; CYAN=""; BOLD=""; RESET=""
fi

die() { echo "${RED}error:${RESET} $*" >&2; exit 1; }

command -v curl >/dev/null 2>&1 || die "curl not found on PATH"
command -v python3 >/dev/null 2>&1 || die "python3 not found on PATH"

echo "${BOLD}RandomWalk token asset check (chain -> metadata -> assets)${RESET}"
echo "  rpc url  : ${CYAN}${RPC_URL}${RESET}"
echo "  contract : ${CYAN}${RWALK_ADDR}${RESET}"
echo "  parallel : ${JOBS} job(s), timeout ${TIMEOUT}s"
[[ -n "$RESOLVE_IP" ]] && echo "  pinned   : ${CYAN}${RESOLVE_IP}${RESET} for ${RESOLVE_HOSTS}"
echo

# ---- sample tokens and resolve their tokenURI via the RPC --------------------
# Prints one "token_id|metadata_url" line per token, or "FETCH_ERROR|reason".
resolve_tokens() {
	python3 - "$RPC_URL" "$RWALK_ADDR" "$COUNT" "$TIMEOUT" "$TOKEN_IDS" <<'PY'
import json
import random
import ssl
import sys
import urllib.request

rpc, addr, count, timeout, explicit = (
    sys.argv[1], sys.argv[2], sys.argv[3], float(sys.argv[4]), sys.argv[5])
ctx = ssl.create_default_context()
ctx.check_hostname = False
ctx.verify_mode = ssl.CERT_NONE  # mirrors the curl -k used for self-signed prod certs


def eth_call(data):
    body = json.dumps({"jsonrpc": "2.0", "id": 1, "method": "eth_call",
                       "params": [{"to": addr, "data": data}, "latest"]}).encode()
    req = urllib.request.Request(rpc, data=body,
                                 headers={"Content-Type": "application/json"})
    with urllib.request.urlopen(req, timeout=timeout, context=ctx) as resp:
        out = json.load(resp)
    if "result" not in out:
        raise RuntimeError(f"eth_call error: {out.get('error')}")
    return bytes.fromhex(out["result"][2:])


def decode_string(raw):
    strlen = int.from_bytes(raw[32:64], "big")
    return raw[64:64 + strlen].decode()


try:
    supply = int.from_bytes(eth_call("0x18160ddd"), "big")  # totalSupply()
except Exception as exc:
    print(f"FETCH_ERROR|totalSupply via {rpc}: {exc}")
    sys.exit(0)
print(f"SUPPLY|{supply}", file=sys.stderr)

if explicit.strip():
    ids = sorted({int(t) for t in explicit.split()})
    bad = [i for i in ids if not 0 <= i < supply]
    if bad:
        print(f"FETCH_ERROR|token ids {bad} out of range 0..{supply - 1}")
        sys.exit(0)
elif count == "all":
    ids = list(range(supply))
else:
    ids = sorted(random.sample(range(supply), min(int(count), supply)))

for tid in ids:
    data = "0xc87b56dd" + format(tid, "064x")  # tokenURI(uint256)
    try:
        uri = decode_string(eth_call(data))
    except Exception as exc:
        print(f"FETCH_ERROR|tokenURI({tid}) via {rpc}: {exc}")
        sys.exit(0)
    print(f"{tid}|{uri}")
PY
}

ROWS=()
SUPPLY="?"
while IFS= read -r line; do
	[[ -z "$line" ]] && continue
	if [[ "$line" == FETCH_ERROR\|* ]]; then
		die "${line#FETCH_ERROR|}"
	elif [[ "$line" == SUPPLY\|* ]]; then
		SUPPLY="${line#SUPPLY|}"
	elif [[ "$line" =~ ^[0-9]+\|https?:// ]]; then
		ROWS+=("$line")
	else
		echo "${YELLOW}skipping unexpected line:${RESET} ${line}" >&2
	fi
done < <(resolve_tokens 2>&1)

[[ "${#ROWS[@]}" -gt 0 ]] || die "no tokens resolved from the chain"

echo "totalSupply() = ${BOLD}${SUPPLY}${RESET}; checking ${BOLD}${#ROWS[@]}${RESET} token(s)" \
	"(1 metadata + 9 resource requests each)."
echo

# ---- worker that checks one token --------------------------------------------
# Prints a human line (unless FAIL_ONLY and all passed) plus a machine line:
#   RESULT|<token_id>|<ok_count>|<fail_count>
check_token() {
	local row="$1"
	local token_id meta_url
	IFS='|' read -r token_id meta_url <<<"$row"

	local ok=0 fail=0 line cell
	line="$(printf '%s#%-6s%s' "$BOLD" "$token_id" "$RESET")"

	# 1) metadata JSON: must be HTTP 200 and parse; emits "label|url" lines.
	local meta_file code urls rargs
	rargs=()
	mapfile -t rargs < <(resolve_args "$meta_url")
	meta_file="$(mktemp)"
	code="$(curl -sk "${rargs[@]}" -o "$meta_file" -w '%{http_code}' --max-time "$TIMEOUT" "$meta_url" 2>/dev/null)"
	if [[ "$code" != "200" ]]; then
		rm -f "$meta_file"
		line+="  ${RED}metadata:FAIL(${code:-000} ${meta_url})${RESET}"
		printf '%s\n' "$line"
		echo "RESULT|${token_id}|0|1"
		return
	fi
	urls="$(python3 - "$meta_file" <<'PY'
import json
import sys

try:
    with open(sys.argv[1]) as fh:
        meta = json.load(fh)
except Exception:
    sys.exit(1)

image = str(meta.get("image", ""))
if not image.endswith("_black.png"):
    sys.exit(2)  # unexpected naming; derived variants would be wrong
base = image.removesuffix("_black.png")

urls = {
    "meta_img": image,
    "meta_anim": str(meta.get("animation_url", "")),
    "ext_url": str(meta.get("external_url", "")),
}
for bg in ("white", "black"):
    urls[f"{bg[0]}png"] = f"{base}_{bg}.png"
    urls[f"{bg[0]}thumb"] = f"{base}_{bg}_thumb.jpg"
    urls[f"{bg[0]}single"] = f"{base}_{bg}_single.mp4"
    urls[f"{bg[0]}triple"] = f"{base}_{bg}_triple.mp4"

seen = set()
for label, url in urls.items():
    if not url.startswith("http") or url in seen:
        continue
    seen.add(url)
    print(f"{label}|{url}")
PY
)"
	local status=$?
	rm -f "$meta_file"
	if [[ $status -ne 0 || -z "$urls" ]]; then
		fail=1
		line+="  ${RED}metadata:FAIL(parse ${meta_url})${RESET}"
		printf '%s\n' "$line"
		echo "RESULT|${token_id}|0|1"
		return
	fi
	ok=$((ok+1))
	line+="  ${GREEN}metadata:OK${RESET}"

	# 2) every resource: HEAD must return 200.
	local label url code
	while IFS='|' read -r label url; do
		[[ -z "$url" ]] && continue
		rargs=()
		mapfile -t rargs < <(resolve_args "$url")
		code="$(curl -skI "${rargs[@]}" -o /dev/null -w '%{http_code}' --max-time "$TIMEOUT" "$url" 2>/dev/null)"
		if [[ "$code" == "200" ]]; then
			ok=$((ok+1))
			cell="${GREEN}${label}:OK${RESET}"
		else
			fail=$((fail+1))
			cell="${RED}${label}:FAIL(${code:-000})${RESET}"
		fi
		line+="  ${cell}"
	done <<<"$urls"

	if [[ "$FAIL_ONLY" -eq 0 || "$fail" -gt 0 ]]; then
		printf '%s\n' "$line"
	fi
	echo "RESULT|${token_id}|${ok}|${fail}"
}
export -f check_token resolve_args
export TIMEOUT FAIL_ONLY RESOLVE_IP RESOLVE_HOSTS BOLD RESET GREEN RED

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
printf '  requests made       : %d\n' "$total_req"
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
echo "${GREEN}${BOLD}RESULT: SUCCESS${RESET} — all ${total_req} requests returned 200."
exit 0
