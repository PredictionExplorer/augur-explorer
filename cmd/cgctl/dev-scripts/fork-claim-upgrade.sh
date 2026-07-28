#!/bin/bash
# fork-claim-upgrade.sh — full V2->V3 rehearsal in ONE run, designed to finish all
# upstream-state-dependent work while the pinned block is still warm on a PRUNING
# upstream node:
#
#   0. compile the contracts FIRST (slowest step; needs no fork/upstream)
#   1. fork Arbitrum One (nohup, log in current dir), pinned at head - 8
#   2. wait for RPC; neutralize Arbitrum precompiles (0x64/0x6C report 0xfe upstream,
#      which burns all gas on a fork); materialize the zero address locally
#   3. prewarm ALL accounts the claim + upgrade flows touch:
#      proxy (code/balance/nonce/config reads), current ERC1967 implementation
#      (code/balance/nonce — the OZ upgrades plugin reads these), owner
#   4. impersonated owner sets delayDurationBeforeRoundActivation (maintenance window)
#   5. jump chain time to mainPrizeTime + claim timeout + 1 s
#   6. claim the main prize with hardhat account #0 (round ends, window opens)
#   7. run the V2->V3 upgrade (upgrade-fork-to-v3.js) immediately, state still warm
#
# Environment (override as needed):
#   UPSTREAM_RPC   node to fork from                (default http://69.10.55.2:38545)
#   HH_REPO        Cosmic-Signature repo            (default $HOME/Cosmic-Signature)
#   HH_HOST        fork node bind interface         (default 0.0.0.0)
#   HH_PORT        fork node port                   (default 10545)
#   GAME_ADDR      game proxy                       (default: production proxy)
#   DELAY_SECONDS  maintenance window (default 1200 = 20 min)
set -e -o pipefail

# Optional local settings file (survives script updates): current dir first, then script dir.
SELF_DIR="$(cd "$(dirname "$0")" && pwd)"
if [ -f "./fork-env.sh" ]; then . "./fork-env.sh"
elif [ -f "$SELF_DIR/fork-env.sh" ]; then . "$SELF_DIR/fork-env.sh"; fi

UPSTREAM_RPC=${UPSTREAM_RPC:-http://69.10.55.2:38545}
HH_REPO=${HH_REPO:-$HOME/Cosmic-Signature}
HH_HOST=${HH_HOST:-0.0.0.0}
HH_PORT=${HH_PORT:-10545}
GAME_ADDR=${GAME_ADDR:-0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2}
DELAY_SECONDS=${DELAY_SECONDS:-1200}

ACCOUNT0=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
ZERO=0x0000000000000000000000000000000000000000
# ERC-1967 implementation slot
IMPL_SLOT=0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc

RUN_DIR="$(pwd)"
LOG_FILE="$RUN_DIR/hardhat-fork.log"
PID_FILE="$RUN_DIR/hardhat-fork.pid"
# FORK_RPC_URL overrides; otherwise derive from HH_HOST/HH_PORT.
if [ -n "$FORK_RPC_URL" ]; then FORK_RPC="$FORK_RPC_URL"
elif [ "$HH_HOST" = "0.0.0.0" ]; then FORK_RPC="http://127.0.0.1:$HH_PORT"
else FORK_RPC="http://$HH_HOST:$HH_PORT"; fi
echo "fork node: bind $HH_HOST:$HH_PORT | rpc $FORK_RPC | upstream $UPSTREAM_RPC"

# ---------- JSON-RPC helpers ----------
rpc_to() { curl -s -X POST "$1" -H 'Content-Type: application/json' \
	-d "{\"jsonrpc\":\"2.0\",\"method\":\"$2\",\"params\":$3,\"id\":1}"; }
rpc() { rpc_to "$FORK_RPC" "$@"; }
rpc_result() {
	local resp; resp=$(rpc "$1" "$2")
	[ -z "$resp" ] && { echo "no response from $FORK_RPC" >&2; exit 1; }
	case "$resp" in *'"error"'*) echo "RPC error from $1: $resp" >&2; exit 1;; esac
	printf '%s\n' "$resp" | sed -n 's/.*"result":"\{0,1\}\([^",}]*\)"\{0,1\}[,}].*/\1/p'
}
hex2dec() { local h=${1#0x}; h=$(printf '%s' "$h" | sed 's/^0*//'); [ -z "$h" ] && h=0; echo $((16#$h)); }
game_call() { hex2dec "$(rpc_result eth_call "[{\"to\":\"$GAME_ADDR\",\"data\":\"$1\"},\"latest\"]")"; }
game_call_addr() {
	local word; word=$(rpc_result eth_call "[{\"to\":\"$GAME_ADDR\",\"data\":\"$1\"},\"latest\"]")
	word=${word#0x}; echo "0x${word: -40}"
}
chain_now() {
	local ts; ts=$(rpc eth_getBlockByNumber '["latest",false]' \
		| sed -n 's/.*"timestamp":"0x\([0-9a-fA-F]*\)".*/\1/p')
	[ -z "$ts" ] && { echo "could not read latest block" >&2; exit 1; }
	echo $((16#$ts))
}
fmt_ts() { echo "$1  ($(date -u -d "@$1" '+%Y-%m-%d %H:%M:%S') UTC)"; }
warm_account() { # fetch code+balance+nonce so they are cached before the upstream prunes
	rpc_result eth_getCode "[\"$1\",\"latest\"]" >/dev/null
	rpc_result eth_getBalance "[\"$1\",\"latest\"]" >/dev/null
	rpc_result eth_getTransactionCount "[\"$1\",\"latest\"]" >/dev/null
}
send_tx() { # send_tx <from> <calldata> <gas-hex> -> tx hash; force-mines and checks receipt
	local txh
	txh=$(rpc_result eth_sendTransaction "[{\"from\":\"$1\",\"to\":\"$GAME_ADDR\",\"data\":\"$2\",\"gas\":\"$3\"}]")
	[ -z "$txh" ] && { echo "eth_sendTransaction returned no tx hash — aborting" >&2; exit 1; }
	rpc_result evm_mine "[]" >/dev/null
	local rcpt; rcpt=$(rpc eth_getTransactionReceipt "[\"$txh\"]")
	case "$rcpt" in
		*'"status":"0x1"'*) echo "$txh";;
		*) echo "transaction $txh FAILED: $rcpt" >&2; exit 1;;
	esac
}

# Selectors: mainPrizeTime() timeoutDurationToClaimMainPrize() roundActivationTime()
# roundNum() owner() lastBidderAddress() delayDurationBeforeRoundActivation()
# setDelayDurationBeforeRoundActivation(uint256) claimMainPrize()
SEL_PRIZE_TIME=0x18305de2
SEL_TIMEOUT=0x3b9d292e
SEL_ACTIVATION=0x6e970834
SEL_ROUND_NUM=0x119b22b3
SEL_OWNER=0x8da5cb5b
SEL_LAST_BIDDER=0xe5b3cd14
SEL_GET_DELAY=0xb9cf9ba5
SEL_SET_DELAY=0x09794bee
SEL_CLAIM=0x448c6eb1

# ---------- 0. compile BEFORE forking (needs no upstream; the slow part) ----------
echo "== 0/7 Compiling contracts (pre-fork, so the upgrade step starts instantly) =="
if [ ! -f "$HH_REPO/scripts/upgrade-fork-to-v3.js" ]; then
	echo "$HH_REPO/scripts/upgrade-fork-to-v3.js not found — copy it there first" >&2; exit 1
fi
( cd "$HH_REPO" && HARDHAT_MODE_CODE=2 npx hardhat compile ) || exit 1

# ---------- 0b. stop a previous fork node ----------
if [ -f "$PID_FILE" ] && kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
	echo "Stopping previous fork node (pid $(cat "$PID_FILE"))..."
	kill "$(cat "$PID_FILE")" || true
	sleep 2
fi

# ---------- 1. launch the fork ----------
echo "== 1/7 Launching Hardhat fork =="
LATEST_HEX=$(rpc_to "$UPSTREAM_RPC" eth_blockNumber '[]' | sed -n 's/.*"result":"0x\([0-9a-fA-F]*\)".*/\1/p')
[ -z "$LATEST_HEX" ] && { echo "failed to fetch latest block from $UPSTREAM_RPC" >&2; exit 1; }
PIN=$(( $((16#$LATEST_HEX)) - 8 ))
echo "upstream: $UPSTREAM_RPC | pin block: $PIN | log: $LOG_FILE"
echo "$PIN" > "$RUN_DIR/fork-block.txt"
cd "$HH_REPO"
HARDHAT_MODE_CODE=2 nohup npx hardhat node --fork "$UPSTREAM_RPC" --fork-block-number "$PIN" \
	--hostname "$HH_HOST" --port "$HH_PORT" > "$LOG_FILE" 2>&1 &
echo $! > "$PID_FILE"
cd "$RUN_DIR"
echo "node pid: $(cat "$PID_FILE")"

# ---------- 2. wait; neutralize precompiles; materialize zero address ----------
echo "== 2/7 Waiting for the fork RPC =="
for i in $(seq 1 60); do
	if rpc eth_blockNumber '[]' 2>/dev/null | grep -q '"result"'; then break; fi
	if ! kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
		echo "hardhat node died during startup; last log lines:" >&2
		tail -15 "$LOG_FILE" >&2; exit 1
	fi
	sleep 2
	[ "$i" = 60 ] && { echo "fork did not come up in 120 s; see $LOG_FILE" >&2; exit 1; }
done
rpc_result evm_mine "[]" >/dev/null
echo "fork is up at $FORK_RPC (block $(hex2dec "$(rpc_result eth_blockNumber '[]')"))"
rpc_result hardhat_setCode '["0x0000000000000000000000000000000000000064","0x"]' >/dev/null
rpc_result hardhat_setCode '["0x000000000000000000000000000000000000006C","0x"]' >/dev/null
rpc_result hardhat_setNonce "[\"$ZERO\",\"0x0\"]" >/dev/null
rpc_result hardhat_setBalance "[\"$ZERO\",\"0x0\"]" >/dev/null
rpc_result hardhat_setCode "[\"$ZERO\",\"0x\"]" >/dev/null
echo "precompiles neutralized; zero address materialized"

# ---------- 3. prewarm every account the claim + upgrade touch ----------
echo "== 3/7 Prewarming production state =="
warm_account "$GAME_ADDR"
IMPL_WORD=$(rpc_result eth_getStorageAt "[\"$GAME_ADDR\",\"$IMPL_SLOT\",\"latest\"]")
IMPL_ADDR="0x${IMPL_WORD#0x}"; IMPL_ADDR="0x${IMPL_ADDR: -40}"
warm_account "$IMPL_ADDR"
OWNER=$(game_call_addr $SEL_OWNER)
warm_account "$OWNER"
ROUND_NUM=$(game_call $SEL_ROUND_NUM)
LAST_BIDDER=$(game_call_addr $SEL_LAST_BIDDER)
PRIZE_TIME=$(game_call $SEL_PRIZE_TIME)
TIMEOUT=$(game_call $SEL_TIMEOUT)
echo "owner          : $OWNER"
echo "implementation : $IMPL_ADDR (V2, prewarmed for the upgrades plugin)"
echo "roundNum       : $ROUND_NUM"
echo "lastBidder     : $LAST_BIDDER"
echo "mainPrizeTime  : $(fmt_ts "$PRIZE_TIME")"
echo "claim timeout  : $TIMEOUT s"
if [ -z "$ROUND_NUM" ] || [ -z "$PRIZE_TIME" ] || [ "$PRIZE_TIME" = 0 ] || [ "$OWNER" = "0x" ]; then
	echo "PREWARM FAILED — upstream cannot serve state at pinned block $PIN." >&2
	echo "Use an archive upstream: UPSTREAM_RPC=https://arb-mainnet.g.alchemy.com/v2/<key> $0" >&2
	exit 1
fi

# ---------- 4. maintenance window (impersonated owner) ----------
echo "== 4/7 Setting delayDurationBeforeRoundActivation=$DELAY_SECONDS =="
rpc_result hardhat_impersonateAccount "[\"$OWNER\"]" >/dev/null
rpc_result hardhat_setBalance "[\"$OWNER\",\"0x21E19E0C9BAB2400000\"]" >/dev/null
TX=$(send_tx "$OWNER" "$SEL_SET_DELAY$(printf '%064x' "$DELAY_SECONDS")" "0x186a0")
echo "tx: $TX | delay now: $(game_call $SEL_GET_DELAY) s"

# ---------- 5. jump chain time ----------
echo "== 5/7 Jumping chain time to mainPrizeTime + timeout + 1 =="
TARGET=$((PRIZE_TIME + TIMEOUT + 1))
NOW=$(chain_now)
if [ "$NOW" -lt "$TARGET" ]; then
	rpc_result evm_setNextBlockTimestamp "[$TARGET]" >/dev/null
	rpc_result evm_mine "[]" >/dev/null
fi
echo "chain time now: $(fmt_ts "$(chain_now)") (prize claimable by anyone)"

# ---------- 6. claim (round ends, maintenance window opens) ----------
echo "== 6/7 Claiming the main prize with hardhat account #0 =="
# 6M gas: the real production round-0 claim used ~4.03M. Sufficient only because the
# precompiles were neutralized above (their 0xfe placeholder burns ANY gas limit).
TX=$(send_tx "$ACCOUNT0" "$SEL_CLAIM" "0x5B8D80")
ROUND_AFTER=$(game_call $SEL_ROUND_NUM)
ACTIVATION=$(game_call $SEL_ACTIVATION)
echo "claim tx: $TX"
echo "ROUND ENDED. roundNum: $ROUND_NUM -> $ROUND_AFTER"
echo "maintenance window until: $(fmt_ts "$ACTIVATION")"

# ---------- 7. upgrade to V3, state still warm ----------
echo "== 7/7 Running the V2 -> V3 upgrade =="
cd "$HH_REPO"
HARDHAT_MODE_CODE=2 FORK_RPC_URL="$FORK_RPC" GAME_PROXY_ADDR="$GAME_ADDR" \
	npx hardhat run scripts/upgrade-fork-to-v3.js --network fork_local
cd "$RUN_DIR"

echo
echo "ALL DONE: round claimed under V2, proxy upgraded to V3, maintenance window"
echo "open until $(fmt_ts "$ACTIVATION")."
echo "Start the first V3 round from your laptop:"
echo "  RPC_URL=http://<this-host>:$HH_PORT ./fork-v3-rehearsal.sh new-round"
echo "Node log: $LOG_FILE | pid: $(cat "$PID_FILE") | pinned block: $PIN"
