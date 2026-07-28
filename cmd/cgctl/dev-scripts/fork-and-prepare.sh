#!/bin/bash
# fork-and-prepare.sh — one-shot: fork Arbitrum One into a local Hardhat node (nohup,
# log in the current directory) and immediately run the transactions that make the
# contracts ready for the prize claim + V3 upgrade rehearsal:
#
#   1. launch  : npx hardhat node --fork <upstream> pinned at (head - 64), via nohup
#   2. wait    : poll until the fork answers JSON-RPC
#   3. prewarm : read all state the rehearsal needs, so Hardhat caches it locally
#                BEFORE the upstream full node prunes the pinned block's state
#   4. set-delay: impersonated owner sets delayDurationBeforeRoundActivation = 20 min
#   5. jump    : chain time -> mainPrizeTime + claim timeout + 1 s
#
# After this script succeeds, the round is claimable by anyone. Continue with:
#   ./fork-v3-rehearsal.sh claim      (hardhat account #0 claims; round goes inactive)
#   ./fork-v3-rehearsal.sh upgrade    (V2 -> V3 during the maintenance window)
#   ./fork-v3-rehearsal.sh new-round  (jump past activation, first gesture on V3)
#
# Environment (override as needed):
#   UPSTREAM_RPC   Arbitrum One node to fork from   (default http://69.10.55.2:38545)
#   HH_REPO        Cosmic-Signature repo            (default $HOME/Cosmic-Signature)
#   HH_HOST        interface for the fork node      (default 0.0.0.0)
#   HH_PORT        port for the fork node           (default 10545)
#   GAME_ADDR      game proxy                       (default: production proxy)
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
DELAY_SECONDS=1200   # 20-minute maintenance window after the claim

RUN_DIR="$(pwd)"
LOG_FILE="$RUN_DIR/hardhat-fork.log"
PID_FILE="$RUN_DIR/hardhat-fork.pid"
# FORK_RPC_URL overrides; otherwise derive from HH_HOST/HH_PORT.
if [ -n "$FORK_RPC_URL" ]; then FORK_RPC="$FORK_RPC_URL"
elif [ "$HH_HOST" = "0.0.0.0" ]; then FORK_RPC="http://127.0.0.1:$HH_PORT"
else FORK_RPC="http://$HH_HOST:$HH_PORT"; fi
echo "fork node: bind $HH_HOST:$HH_PORT | rpc $FORK_RPC | upstream $UPSTREAM_RPC"

# ---------- JSON-RPC helpers (bash/sed only) ----------
rpc_to() { # rpc_to <url> <method> <params-json>
	curl -s -X POST "$1" -H 'Content-Type: application/json' \
		-d "{\"jsonrpc\":\"2.0\",\"method\":\"$2\",\"params\":$3,\"id\":1}"
}
rpc() { rpc_to "$FORK_RPC" "$@"; }
rpc_result() { # dies on RPC error, prints "result"
	local resp
	resp=$(rpc "$1" "$2")
	if [ -z "$resp" ]; then echo "no response from $FORK_RPC" >&2; exit 1; fi
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
	local ts
	ts=$(rpc eth_getBlockByNumber '["latest",false]' | sed -n 's/.*"timestamp":"0x\([0-9a-fA-F]*\)".*/\1/p')
	[ -z "$ts" ] && { echo "could not read latest block" >&2; exit 1; }
	echo $((16#$ts))
}
fmt_ts() { echo "$1  ($(date -u -d "@$1" '+%Y-%m-%d %H:%M:%S') UTC)"; }

# Selectors: mainPrizeTime() timeoutDurationToClaimMainPrize() roundActivationTime()
# roundNum() owner() lastBidderAddress() setDelayDurationBeforeRoundActivation(uint256)
# delayDurationBeforeRoundActivation()
SEL_PRIZE_TIME=0x18305de2
SEL_TIMEOUT=0x3b9d292e
SEL_ACTIVATION=0x6e970834
SEL_ROUND_NUM=0x119b22b3
SEL_OWNER=0x8da5cb5b
SEL_LAST_BIDDER=0xe5b3cd14
SEL_SET_DELAY=0x09794bee
SEL_GET_DELAY=0xb9cf9ba5

# ---------- 0. stop a previous fork node started by this script ----------
if [ -f "$PID_FILE" ] && kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
	echo "Stopping previous fork node (pid $(cat "$PID_FILE"))..."
	kill "$(cat "$PID_FILE")" || true
	sleep 2
fi

# ---------- 1. launch the fork ----------
echo "== 1/5 Launching Hardhat fork =="
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

# ---------- 2. wait until the fork answers ----------
echo "== 2/5 Waiting for the fork RPC =="
for i in $(seq 1 60); do
	if rpc eth_blockNumber '[]' 2>/dev/null | grep -q '"result"'; then break; fi
	if ! kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
		echo "hardhat node died during startup; last log lines:" >&2
		tail -15 "$LOG_FILE" >&2
		exit 1
	fi
	sleep 2
	[ "$i" = 60 ] && { echo "fork did not come up in 120 s; see $LOG_FILE" >&2; exit 1; }
done
echo "fork is up at $FORK_RPC (block $(hex2dec "$(rpc_result eth_blockNumber '[]')"))"

# ---------- 3. prewarm: fetch state now, while the upstream can still serve it ----------
echo "== 3/5 Prewarming fork state (caches production state locally) =="
OWNER=$(game_call_addr $SEL_OWNER)
ROUND_NUM=$(game_call $SEL_ROUND_NUM)
LAST_BIDDER=$(game_call_addr $SEL_LAST_BIDDER)
PRIZE_TIME=$(game_call $SEL_PRIZE_TIME)
TIMEOUT=$(game_call $SEL_TIMEOUT)
ACTIVATION=$(game_call $SEL_ACTIVATION)
DELAY_NOW=$(game_call $SEL_GET_DELAY)
echo "owner           : $OWNER"
echo "roundNum        : $ROUND_NUM"
echo "lastBidder      : $LAST_BIDDER"
echo "chain time      : $(fmt_ts "$(chain_now)")"
echo "roundActivation : $(fmt_ts "$ACTIVATION")"
echo "mainPrizeTime   : $(fmt_ts "$PRIZE_TIME")"
echo "claim timeout   : $TIMEOUT s"
echo "activation delay: $DELAY_NOW s"
if [ -z "$ROUND_NUM" ] || [ -z "$PRIZE_TIME" ] || [ "$PRIZE_TIME" = 0 ] || [ "$OWNER" = "0x" ]; then
	echo "PREWARM FAILED — upstream cannot serve state at pinned block $PIN." >&2
	echo "The upstream node needs historical-state support (see notes); re-running" >&2
	echo "this script immediately can also win the race on a lucky run." >&2
	exit 1
fi

# ---------- 4. impersonated owner: activation delay = 20 min ----------
echo "== 4/5 Setting delayDurationBeforeRoundActivation=$DELAY_SECONDS (impersonated owner) =="
rpc_result hardhat_impersonateAccount "[\"$OWNER\"]" >/dev/null
rpc_result hardhat_setBalance "[\"$OWNER\",\"0x21E19E0C9BAB2400000\"]" >/dev/null
ARG=$(printf '%064x' $DELAY_SECONDS)
TX=$(rpc_result eth_sendTransaction "[{\"from\":\"$OWNER\",\"to\":\"$GAME_ADDR\",\"data\":\"$SEL_SET_DELAY$ARG\",\"gas\":\"0x186a0\"}]")
echo "tx: $TX"
echo "new delay: $(game_call $SEL_GET_DELAY) s"

# ---------- 5. jump chain time past mainPrizeTime + timeout ----------
echo "== 5/5 Jumping chain time to mainPrizeTime + timeout + 1 =="
TARGET=$((PRIZE_TIME + TIMEOUT + 1))
NOW=$(chain_now)
if [ "$NOW" -ge "$TARGET" ]; then
	echo "chain time $NOW already >= target $TARGET, nothing to do"
else
	rpc_result evm_setNextBlockTimestamp "[$TARGET]" >/dev/null
	rpc_result evm_mine "[]" >/dev/null
	echo "chain time now: $(fmt_ts "$(chain_now)")"
fi

echo
echo "READY: the main prize is claimable by anyone, and a 20-minute maintenance"
echo "window will follow the claim. Continue with:"
echo "  RPC_URL=http://<this-host>:$HH_PORT ./fork-v3-rehearsal.sh claim"
echo "  RPC_URL=http://<this-host>:$HH_PORT ./fork-v3-rehearsal.sh upgrade"
echo "  RPC_URL=http://<this-host>:$HH_PORT ./fork-v3-rehearsal.sh new-round"
echo "Node log: $LOG_FILE | pid: $(cat "$PID_FILE") | pinned block: $PIN"
