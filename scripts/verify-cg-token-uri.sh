#!/usr/bin/env bash
#
# verify-cg-token-uri.sh — verify Cosmic Signature token assets end to end,
# starting from the live blockchain (the marketplace view of a token).
#
# Complements verify-token-images.sh, which walks the API token list and
# checks the three website image renditions. This script instead follows the
# path OpenSea & co. take:
#
#   1. eth_call CosmicSignatureNft.tokenURI(tokenId) on the RPC node
#   2. GET the metadata JSON that URI points to
#   3. HEAD every http(s) URL found anywhere in the metadata: image,
#      animation_url, external_url and the full properties.media manifest
#      (previews, HQ videos, spectral sweeps, trait/asset JSONs, ...)
#   4. when the metadata declares image_details/animation_details.bytes,
#      compare them against the served Content-Length
#
# Directory-style URLs (trailing "/", e.g. generation_records and
# spectral_bins) are reported as warnings when they do not answer 200, since
# object stores often serve the files but no directory index.
#
# Token ids are sampled at random from totalSupply() on the same RPC, so the
# script needs no database access.
#
# Usage:
#   ./scripts/verify-cg-token-uri.sh [RPC_URL] [COUNT]
#
#   RPC_URL   JSON-RPC endpoint (default https://rpc1.cosmicsignature.com:38546)
#   COUNT     how many random tokens to sample (default 5; "all" = every token)
#
# Environment overrides:
#   CSNFT_ADDR=0x...   CosmicSignatureNft contract (default: Arbitrum production)
#   TOKEN_IDS="1 2 3"  check exactly these ids instead of sampling
#   JOBS=8             parallel tokens (default 4)
#   TIMEOUT=15         per-request timeout in seconds (default 15)
#   FAIL_ONLY=1        print only tokens with at least one failing request
#   RESOLVE_IP=1.2.3.4 pin the asset host to this server, so each production
#                      box can be checked individually (public DNS points at
#                      only one of them). Only hosts listed in RESOLVE_HOSTS
#                      are pinned; external website URLs keep normal DNS.
#   RESOLVE_HOSTS="..."  hosts affected by RESOLVE_IP
#                      (default: nfts.cosmicsignature.com)
#
# Examples:
#   ./scripts/verify-cg-token-uri.sh
#   ./scripts/verify-cg-token-uri.sh https://rpc2.cosmicsignature.com:38546 all
#   TOKEN_IDS="0 5 71" ./scripts/verify-cg-token-uri.sh
#   RESOLVE_IP=161.129.67.58 ./scripts/verify-cg-token-uri.sh   # cosmic2
#
# Requires: curl, python3. Exits 0 when every non-directory request returns
# HTTP 200 and every declared size matches.

set -uo pipefail

RPC_URL="${1:-https://rpc1.cosmicsignature.com:38546}"
COUNT="${2:-5}"
CSNFT_ADDR="${CSNFT_ADDR:-0xbb84Be3500A63581d3F2d5AC3bdF8685AAedad25}"
TIMEOUT="${TIMEOUT:-15}"
JOBS="${JOBS:-4}"
FAIL_ONLY="${FAIL_ONLY:-0}"
TOKEN_IDS="${TOKEN_IDS:-}"
RESOLVE_IP="${RESOLVE_IP:-}"
RESOLVE_HOSTS="${RESOLVE_HOSTS:-nfts.cosmicsignature.com}"

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

echo "${BOLD}Cosmic Signature tokenURI asset check (chain -> metadata -> assets)${RESET}"
echo "  rpc url  : ${CYAN}${RPC_URL}${RESET}"
echo "  contract : ${CYAN}${CSNFT_ADDR}${RESET}"
echo "  parallel : ${JOBS} job(s), timeout ${TIMEOUT}s"
[[ -n "$RESOLVE_IP" ]] && echo "  pinned   : ${CYAN}${RESOLVE_IP}${RESET} for ${RESOLVE_HOSTS}"
echo

# ---- sample tokens and resolve their tokenURI via the RPC --------------------
resolve_tokens() {
	python3 - "$RPC_URL" "$CSNFT_ADDR" "$COUNT" "$TIMEOUT" "$TOKEN_IDS" <<'PY'
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

echo "totalSupply() = ${BOLD}${SUPPLY}${RESET}; checking ${BOLD}${#ROWS[@]}${RESET} token(s)."
echo

# ---- worker that checks one token --------------------------------------------
# Machine line: RESULT|<token_id>|<ok>|<fail>|<warn>
check_token() {
	local row="$1"
	local token_id meta_url
	IFS='|' read -r token_id meta_url <<<"$row"

	local ok=0 fail=0 warn=0 line cell
	line="$(printf '%s#%-4s%s' "$BOLD" "$token_id" "$RESET")"

	# 1) metadata JSON: HTTP 200 + parse; emits "label|dir|bytes|url" lines.
	local meta_file code urls rargs
	rargs=()
	mapfile -t rargs < <(resolve_args "$meta_url")
	meta_file="$(mktemp)"
	code="$(curl -sk "${rargs[@]}" -o "$meta_file" -w '%{http_code}' --max-time "$TIMEOUT" "$meta_url" 2>/dev/null)"
	if [[ "$code" != "200" ]]; then
		rm -f "$meta_file"
		line+="  ${RED}metadata:FAIL(${code:-000} ${meta_url})${RESET}"
		printf '%s\n' "$line"
		echo "RESULT|${token_id}|0|1|0"
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

expected = {
    str(meta.get("image", "")): int((meta.get("image_details") or {}).get("bytes") or 0),
    str(meta.get("animation_url", "")): int((meta.get("animation_details") or {}).get("bytes") or 0),
}

rows = []
seen = set()


def walk(node, label):
    if isinstance(node, dict):
        for k, v in node.items():
            walk(v, k)
    elif isinstance(node, list):
        for v in node:
            walk(v, label)
    elif isinstance(node, str) and node.startswith(("http://", "https://")):
        if node in seen:
            return
        seen.add(node)
        rows.append((label, node))


walk(meta, "url")
if not rows:
    sys.exit(2)
for label, url in rows:
    isdir = 1 if url.endswith("/") else 0
    print(f"{label}|{isdir}|{expected.get(url, 0)}|{url}")
PY
)"
	local status=$?
	rm -f "$meta_file"
	if [[ $status -ne 0 || -z "$urls" ]]; then
		line+="  ${RED}metadata:FAIL(parse ${meta_url})${RESET}"
		printf '%s\n' "$line"
		echo "RESULT|${token_id}|0|1|0"
		return
	fi
	ok=$((ok+1))
	line+="  ${GREEN}metadata:OK${RESET}"

	# 2) every URL: HEAD must return 200 (redirects followed); declared sizes
	# must match the final hop's Content-Length.
	local label isdir want_bytes url hdr_file http_code got_bytes
	hdr_file="$(mktemp)"
	while IFS='|' read -r label isdir want_bytes url; do
		[[ -z "$url" ]] && continue
		rargs=()
		mapfile -t rargs < <(resolve_args "$url")
		http_code="$(curl -skIL "${rargs[@]}" -o /dev/null -D "$hdr_file" -w '%{http_code}' \
			--max-time "$TIMEOUT" "$url" 2>/dev/null)"
		got_bytes="$(grep -i '^content-length:' "$hdr_file" | tail -1 | tr -dc '0-9')"
		if [[ "$http_code" == "200" ]]; then
			if [[ "$want_bytes" -gt 0 && -n "$got_bytes" && "$got_bytes" != "$want_bytes" ]]; then
				fail=$((fail+1))
				cell="${RED}${label}:SIZE(${got_bytes}/${want_bytes})${RESET}"
			else
				ok=$((ok+1))
				cell="${GREEN}${label}:OK${RESET}"
			fi
		elif [[ "$isdir" == "1" ]]; then
			warn=$((warn+1))
			cell="${YELLOW}${label}:WARN(dir ${http_code:-000})${RESET}"
		else
			fail=$((fail+1))
			cell="${RED}${label}:FAIL(${http_code:-000})${RESET}"
		fi
		line+="  ${cell}"
	done <<<"$urls"
	rm -f "$hdr_file"

	if [[ "$FAIL_ONLY" -eq 0 || "$fail" -gt 0 ]]; then
		printf '%s\n' "$line"
	fi
	echo "RESULT|${token_id}|${ok}|${fail}|${warn}"
}
export -f check_token resolve_args
export TIMEOUT FAIL_ONLY RESOLVE_IP RESOLVE_HOSTS BOLD RESET GREEN RED YELLOW

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
total_warn=0
tokens_checked=0
tokens_failed=0
while IFS='|' read -r _ _tid oc fc wc; do
	tokens_checked=$((tokens_checked + 1))
	total_ok=$((total_ok + oc))
	total_fail=$((total_fail + fc))
	total_warn=$((total_warn + wc))
	[[ "$fc" -gt 0 ]] && tokens_failed=$((tokens_failed + 1))
done < <(grep '^RESULT|' "$TMP_OUT")

total_req=$((total_ok + total_fail + total_warn))

echo
echo "${BOLD}Summary${RESET}"
printf '  tokens checked      : %d\n' "$tokens_checked"
printf '  requests made       : %d\n' "$total_req"
printf '  %ssuccess (200)%s      : %d\n' "$GREEN" "$RESET" "$total_ok"
printf '  %swarnings (dirs)%s    : %d\n' "$YELLOW" "$RESET" "$total_warn"
printf '  %sfailure%s            : %d\n' "$RED" "$RESET" "$total_fail"
printf '  tokens with failures: %s%d%s\n' \
	"$([[ "$tokens_failed" -gt 0 ]] && echo "$RED" || echo "$GREEN")" "$tokens_failed" "$RESET"

if [[ "$total_fail" -gt 0 ]]; then
	echo
	echo "${RED}${BOLD}RESULT: FAILURE${RESET} — ${total_fail} request(s) failed."
	exit 1
fi
echo
echo "${GREEN}${BOLD}RESULT: SUCCESS${RESET} — all ${total_req} checks passed" \
	"$([[ "$total_warn" -gt 0 ]] && echo "(${total_warn} directory-URL warning(s))")."
exit 0
