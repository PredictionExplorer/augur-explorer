#!/bin/bash
# ssl-baseline.sh — SSL certificate expiry check for all servers/services
# usage: ./ssl-baseline.sh
#
# PASS (green)  = more than WARN_DAYS until expiry
# WARN (yellow) = expires within WARN_DAYS
# FAIL (red)    = expired, or no certificate could be fetched

WARN_DAYS=10   # matches srvmonitor SSLWarningDays

PASS=0; WARN=0; FAIL=0

ok()   { printf '  \033[32mPASS\033[0m  %s\n' "$1"; PASS=$((PASS+1)); }
warn() { printf '  \033[33mWARN\033[0m  %s\n' "$1"; WARN=$((WARN+1)); }
bad()  { printf '  \033[31mFAIL\033[0m  %s  [%s]\n' "$1" "$2"; FAIL=$((FAIL+1)); }

# cert <host:port> [servername]  — fetch cert, report subject + days left
cert() {
    local hp=$1 sni=${2:-${1%%:*}} pem enddate subject serial end_ts now_ts days
    pem=$(echo | timeout 10 openssl s_client -connect "$hp" -servername "$sni" 2>/dev/null \
          | openssl x509 2>/dev/null)
    if [ -z "$pem" ]; then
        bad "$hp" "no certificate (connection failed or not TLS)"
        return
    fi
    enddate=$(echo "$pem" | openssl x509 -noout -enddate | cut -d= -f2)
    subject=$(echo "$pem" | openssl x509 -noout -subject | sed 's/^subject=//; s/CN = //')
    serial=$(echo "$pem" | openssl x509 -noout -serial | cut -d= -f2)
    end_ts=$(date -d "$enddate" +%s)
    now_ts=$(date +%s)
    days=$(( (end_ts - now_ts) / 86400 ))
    local line
    line=$(printf '%-40s %-25s %4s days left  (until %s, serial %.8s...)' \
           "$hp" "[$subject]" "$days" "$(date -d "$enddate" '+%b %d %Y')" "$serial")
    if [ "$days" -lt 0 ]; then
        bad "$hp" "EXPIRED $((-days)) days ago ($enddate, CN=$subject)"
    elif [ "$days" -le "$WARN_DAYS" ]; then
        warn "$line"
    else
        ok "$line"
    fi
}

echo "=== 1) RPC HTTPS randomwalk :38547 ==="
for h in rpc1.cosmicsignature.com rpc2.cosmicsignature.com; do
    cert "$h:38547"
done

echo "=== 2) RPC HTTPS cosmicgame :38546 ==="
for h in rpc1.cosmicsignature.com rpc2.cosmicsignature.com; do
    cert "$h:38546"
done

echo "=== 3) CosmicGame via nginx :443 ==="
for h in a1.cosmicsignature.com a2.cosmicsignature.com nfts.cosmicsignature.com; do
    cert "$h:443"
done

echo "=== 4) RandomWalk HTTPS :1443 ==="
for h in a1.cosmicsignature.com a2.cosmicsignature.com api.randomwalknft.com; do
    cert "$h:1443"
done

echo "=== 5) RandomWalk web/API :443 ==="
for h in randomwalknft-api.com nfts.randomwalknft.com www.randomwalknft.com randomwalknft.com; do
    cert "$h:443"
done

echo "=== 6) Cosmic3 sepolia/dev (cosmicsignature) ==="
cert "devapi.cosmicsignature.com:443"
# TODO: add the sepolia/dev API+RPC host:port entries from the srvmonitor
# SSL_CERT*_HOST/_PORT env config, e.g.:
# cert "161.129.67.42:PORT" "some.cosmicsignature.com"   # 2nd arg = SNI servername

echo
echo "passed: $PASS   warned: $WARN   failed: $FAIL"
exit $((FAIL > 0))
