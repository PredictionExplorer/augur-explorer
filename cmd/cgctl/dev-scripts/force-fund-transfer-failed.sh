#!/usr/bin/env bash
# Produce a FundTransferFailed event from StakingWalletCosmicSignatureNft on a
# local Hardhat node and verify that cg-etl stored it in cg_fund_transf_err
# (with err_str) and that cg_error_lookup resolves the event signature.
#
# The wallet only emits FundTransferFailed when tryPerformMaintenance cannot
# forward its balance to the charity address, which a normal game never
# triggers; force-fund-transfer-failed.js manufactures it (plants a balance
# with hardhat_setBalance, then targets a contract that rejects ETH).
#
# Usage:
#   ./force-fund-transfer-failed.sh [--deploy] [--network NETWORK] [COSMIC_SIGNATURE_DIR]
#
#   --deploy   deploy a fresh stack first (deploy-dev-and-samp.js) and use its
#              CADDR. Point cg-etl at the new addresses (cg_contracts) before
#              the verification window closes; the script polls for
#              ETL_WAIT_SEC seconds.
#
# Env:
#   COSMIC_SIGNATURE_DIR   Cosmic-Signature repo (default: the v3.1 checkout)
#   CADDR                  game proxy (default: deploy-dev-and-samp.js address)
#   REJECT_ADDR            ETH-rejecting destination (default: the game's nft())
#   BALANCE_ETH            balance planted in the wallet (default 1)
#   ETL_WAIT_SEC           how long to wait for cg-etl to index the tx (default 120)
#   DATABASE_URL or PGSQL_HOST/PGSQL_USERNAME/PGSQL_PASSWORD/PGSQL_DATABASE
#                          the cg-etl database (same variables cg-etl reads)
#
# Exit code 0 only when the row is present with the expected destination,
# amount and err_str.

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
NETWORK="${NETWORK:-localhost}"
COSMIC_SIGNATURE_DIR="${COSMIC_SIGNATURE_DIR:-/home/niko/eth/dev/b/cg-v3/Cosmic-Signature-v3.1-2026-08-19}"
ETL_WAIT_SEC="${ETL_WAIT_SEC:-120}"
DEPLOY=0

while [ $# -gt 0 ]; do
    case "$1" in
        --deploy) DEPLOY=1; shift ;;
        --network) NETWORK="$2"; shift 2 ;;
        -h|--help) sed -n '2,30p' "$0"; exit 0 ;;
        *) COSMIC_SIGNATURE_DIR="$1"; shift ;;
    esac
done

if [ ! -f "$COSMIC_SIGNATURE_DIR/hardhat.config.js" ]; then
    echo "Error: not a Cosmic-Signature hardhat repo: $COSMIC_SIGNATURE_DIR" >&2
    exit 1
fi

# psql connection: DATABASE_URL wins, else the legacy PGSQL_* settings
# (empty PGSQL_HOST = local Unix socket, as in .env.example; PGSQL_DATABASE
# defaults to the user name like psql does).
psql_conn() {
    if [ -n "${DATABASE_URL:-}" ]; then
        psql -X -q -A -t -v ON_ERROR_STOP=1 "$DATABASE_URL" "$@"
    else
        : "${PGSQL_USERNAME:?set DATABASE_URL or PGSQL_USERNAME (and PGSQL_HOST/PGSQL_PASSWORD/PGSQL_DATABASE as needed)}"
        local host_args=()
        if [ -n "${PGSQL_HOST:-}" ]; then
            host_args=(-h "${PGSQL_HOST%%:*}")
            case "$PGSQL_HOST" in *:*) host_args+=(-p "${PGSQL_HOST##*:}") ;; esac
        fi
        PGPASSWORD="${PGSQL_PASSWORD:-}" psql -X -q -A -t -v ON_ERROR_STOP=1 \
            "${host_args[@]}" -U "$PGSQL_USERNAME" -d "${PGSQL_DATABASE:-$PGSQL_USERNAME}" "$@"
    fi
}

# Fail early if the database is unreachable or migration 00034 is missing.
if ! psql_conn -c "SELECT 1 FROM cg_error_lookup LIMIT 1" >/dev/null; then
    echo "Error: cannot query cg_error_lookup; is the database reachable and migrated to 00034?" >&2
    exit 1
fi

cd "$COSMIC_SIGNATURE_DIR"
export NODE_PATH="$PWD/node_modules"

if [ "$DEPLOY" = 1 ]; then
    echo "== Deploying a fresh stack (deploy-dev-and-samp.js) on $NETWORK"
    DEPLOY_OUT="$(npx hardhat run "$SCRIPT_DIR/deploy-dev-and-samp.js" --network "$NETWORK" 2>&1)"
    echo "$DEPLOY_OUT"
    CADDR="$(echo "$DEPLOY_OUT" | sed -n 's/^CADDR=\(0x[0-9a-fA-F]\{40\}\)$/\1/p' | tail -1)"
    if [ -z "$CADDR" ]; then
        echo "Error: could not parse CADDR from the deploy output" >&2
        exit 1
    fi
    echo "Deployed game proxy $CADDR. Register the new addresses in cg_contracts and (re)start cg-etl;"
    echo "the verification below polls for $ETL_WAIT_SEC seconds."
fi
export CADDR="${CADDR:-0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512}"

echo "== Forcing FundTransferFailed from the CST staking wallet (game $CADDR, network $NETWORK)"
JS_OUT="$(npx hardhat run "$SCRIPT_DIR/force-fund-transfer-failed.js" --network "$NETWORK" 2>&1)"
echo "$JS_OUT"

get_var() { echo "$JS_OUT" | sed -n "s/^$1=\(.*\)$/\1/p" | tail -1; }
FTF_TX="$(get_var FTF_TX)"
FTF_WALLET="$(get_var FTF_WALLET)"
FTF_DEST="$(get_var FTF_DEST)"
FTF_AMOUNT="$(get_var FTF_AMOUNT)"
FTF_ERRSTR="$(get_var FTF_ERRSTR)"
if [ -z "$FTF_TX" ] || [ -z "$FTF_DEST" ] || [ -z "$FTF_AMOUNT" ]; then
    echo "Error: the hardhat script did not report the emitted event" >&2
    exit 1
fi

echo "== Waiting up to ${ETL_WAIT_SEC}s for cg-etl to index tx $FTF_TX"
QUERY="SELECT f.block_num, c.addr, d.addr, f.amount::TEXT, COALESCE(f.err_str,'<NULL>'), e.topic0_sig
	FROM cg_fund_transf_err f
	JOIN transaction t ON t.id = f.tx_id
	JOIN address c ON c.address_id = f.contract_aid
	JOIN address d ON d.address_id = f.destination_aid
	JOIN evt_log e ON e.id = f.evtlog_id
	WHERE t.tx_hash = '$FTF_TX'"
ROW=""
deadline=$(( $(date +%s) + ETL_WAIT_SEC ))
while [ "$(date +%s)" -lt "$deadline" ]; do
    ROW="$(psql_conn -c "$QUERY" || true)"
    if [ -n "$ROW" ]; then
        break
    fi
    sleep 2
done
if [ -z "$ROW" ]; then
    echo "[FAIL] no cg_fund_transf_err row for tx $FTF_TX after ${ETL_WAIT_SEC}s (is cg-etl running against $CADDR's stack?)" >&2
    exit 1
fi

IFS='|' read -r DB_BLOCK DB_CONTRACT DB_DEST DB_AMOUNT DB_ERRSTR DB_TOPIC <<<"$ROW"
echo "cg_fund_transf_err: block=$DB_BLOCK contract=$DB_CONTRACT destination=$DB_DEST amount=$DB_AMOUNT err_str=\"$DB_ERRSTR\" topic0_sig=$DB_TOPIC"

FAIL=0
check() { # check <label> <got> <want>
    if [ "$(echo "$2" | tr 'A-F' 'a-f')" = "$(echo "$3" | tr 'A-F' 'a-f')" ]; then
        echo "  [PASS] $1: $2"
    else
        echo "  [FAIL] $1: got '$2', want '$3'"; FAIL=1
    fi
}
check "emitting contract is the staking wallet" "$DB_CONTRACT" "$FTF_WALLET"
check "destination" "$DB_DEST" "$FTF_DEST"
check "amount" "$DB_AMOUNT" "$FTF_AMOUNT"
check "err_str" "$DB_ERRSTR" "$FTF_ERRSTR"

LOOKUP="$(psql_conn -c "SELECT name || ': ' || COALESCE(err_str,'<NULL>') FROM cg_error_lookup WHERE kind='event' AND LEFT(signature,8) = '$DB_TOPIC'")"
if [ -n "$LOOKUP" ]; then
    echo "  [PASS] cg_error_lookup resolves topic0 $DB_TOPIC -> $LOOKUP"
else
    echo "  [FAIL] cg_error_lookup has no event row for topic0 $DB_TOPIC"; FAIL=1
fi

if [ "$FAIL" = 0 ]; then
    echo "== OK: FundTransferFailed from the staking wallet is indexed with its err_str."
else
    exit 1
fi
