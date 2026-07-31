#!/bin/bash
# api-baseline.sh — smoke test for cosmic1 (a1/rpc1) and cosmic2 (a2/rpc2) services
# usage: ./api-baseline.sh

PASS=0; FAIL=0

ok()  { printf '  \033[32mPASS\033[0m  %s\n' "$1"; PASS=$((PASS+1)); }
bad() { printf '  \033[31mFAIL\033[0m  %s  [%s]\n' "$1" "$2"; FAIL=$((FAIL+1)); }

# JSON-RPC check: must return a "result" field; prints block number
rpc() {  # rpc <url>
    local r blk
    r=$(curl -sk -m 10 -X POST -H 'Content-Type: application/json' \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' "$1")
    if echo "$r" | grep -q '"result"'; then
        blk=$(echo "$r" | grep -o '"result":"0x[0-9a-f]*"' | grep -o '0x[0-9a-f]*')
        ok "$1  block=$((16#${blk#0x}))"
    else
        bad "$1" "${r:-no response}"
    fi
}

# strict API check: must be HTTP 200 with a JSON object body
api() {  # api <url>
    local body code
    body=$(curl -sk -m 10 -w '\n%{http_code}' "$1")
    code=${body##*$'\n'}
    body=${body%$'\n'*}
    if [ "$code" = "200" ] && [ "${body:0:1}" = "{" ]; then
        ok "$1  200+json (${#body} bytes)"
    else
        bad "$1" "http=$code body=${body:0:80}"
    fi
}

echo "=== 1) RPC HTTP :38545 ==="
for h in rpc1.cosmicsignature.com rpc2.cosmicsignature.com; do
    rpc "http://$h:38545"
done

echo "=== 2) RPC HTTPS randomwalk :38547 ==="
for h in rpc1.cosmicsignature.com rpc2.cosmicsignature.com; do
    rpc "https://$h:38547"
done

echo "=== 3) RPC HTTPS cosmicgame :38546 ==="
for h in rpc1.cosmicsignature.com rpc2.cosmicsignature.com; do
    rpc "https://$h:38546"
done

echo "=== 4) RandomWalk HTTPS :1443 — market statistics (v1) ==="
for h in a1.cosmicsignature.com a2.cosmicsignature.com; do
    api "https://$h:1443/api/randomwalk/statistics/by_market"
done
echo "=== 5) RandomWalk HTTP :9393 — market statistics (v1) ==="
for h in a1.cosmicsignature.com a2.cosmicsignature.com; do
    api "http://$h:9393/api/randomwalk/statistics/by_market"
done

echo "=== 6) CosmicGame HTTP :2121 — statistics ==="
for h in a1.cosmicsignature.com a2.cosmicsignature.com; do
    api "http://$h:2121/api/v2/cosmicgame/statistics"
done

echo "=== 7) CosmicGame via nginx, port 80 — statistics ==="
for h in a1.cosmicsignature.com a2.cosmicsignature.com; do
    api "http://$h/api/v2/cosmicgame/statistics"
done

echo "=== 8) CosmicGame via nginx, port 443 — statistics ==="
for h in a1.cosmicsignature.com a2.cosmicsignature.com; do
    api "https://$h/api/v2/cosmicgame/statistics"
done

echo
echo "passed: $PASS   failed: $FAIL"
exit $((FAIL > 0))
