#!/bin/bash
# fork-to-round-end.sh — ONE script, whole procedure up to "round ended":
#
#   1. fork Arbitrum One into a local Hardhat node (nohup, log in current dir)
#   2. wait for the RPC and prewarm the production state (beats upstream pruning)
#   3. impersonated owner sets the maintenance window (delayDurationBeforeRoundActivation)
#   4. offset chain time to mainPrizeTime + claim timeout + 1 s
#   5. claim the main prize with hardhat account #0 (raw eth_sendTransaction — no Go needed)
#   6. verify: round ended, maintenance window open until the printed activation time
#
# After it finishes, run the V2->V3 upgrade during the maintenance window:
#   cd <Cosmic-Signature repo> && HARDHAT_MODE_CODE=2 npx hardhat run scripts/upgrade-fork-to-v3.js --network fork_local
#
# Environment (override as needed):
#   UPSTREAM_RPC   node to fork from; MUST serve state at the pinned block
#                  (archive endpoint recommended, e.g. Alchemy/Infura free tier)
#                  default: http://69.10.55.2:38545
#   HH_REPO        Cosmic-Signature repo (default $HOME/Cosmic-Signature)
#   HH_HOST        fork node bind interface (default 0.0.0.0)
#   HH_PORT        fork node port (default 10545)
#   GAME_ADDR      game proxy (default: production proxy)
#   DELAY_SECONDS  maintenance window length (default 1200 = 20 min)
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

# Hardhat dev account #0 — unlocked on the node, so eth_sendTransaction just works.
ACCOUNT0=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

RUN_DIR="$(pwd)"
LOG_FILE="$RUN_DIR/hardhat-fork.log"
PID_FILE="$RUN_DIR/hardhat-fork.pid"
# FORK_RPC_URL overrides; otherwise derive from HH_HOST/HH_PORT.
if [ -n "$FORK_RPC_URL" ]; then FORK_RPC="$FORK_RPC_URL"
elif [ "$HH_HOST" = "0.0.0.0" ]; then FORK_RPC="http://127.0.0.1:$HH_PORT"
else FORK_RPC="http://$HH_HOST:$HH_PORT"; fi
echo "fork node: bind $HH_HOST:$HH_PORT | rpc $FORK_RPC | upstream $UPSTREAM_RPC"

# ---------- JSON-RPC helpers (bash/sed only) ----------
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

send_tx() { # send_tx <from> <calldata> <gas-hex>  -> tx hash; force-mines and checks receipt
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

# ---------- 0. stop a previous fork node started by this script ----------
if [ -f "$PID_FILE" ] && kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
	echo "Stopping previous fork node (pid $(cat "$PID_FILE"))..."
	kill "$(cat "$PID_FILE")" || true
	sleep 2
fi

# ---------- 1. launch the fork ----------
echo "== 1/6 Launching Hardhat fork =="
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

# ---------- 2. wait + prewarm ----------
echo "== 2/6 Waiting for the fork RPC =="
for i in $(seq 1 60); do
	if rpc eth_blockNumber '[]' 2>/dev/null | grep -q '"result"'; then break; fi
	if ! kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
		echo "hardhat node died during startup; last log lines:" >&2
		tail -15 "$LOG_FILE" >&2; exit 1
	fi
	sleep 2
	[ "$i" = 60 ] && { echo "fork did not come up in 120 s; see $LOG_FILE" >&2; exit 1; }
done
# mine one local block so 'latest' is past the fork block (avoids hardfork-history edge)
rpc_result evm_mine "[]" >/dev/null
echo "fork is up at $FORK_RPC (block $(hex2dec "$(rpc_result eth_blockNumber '[]')"))"

# Neutralize Arbitrum precompiles: on the real chain ArbSys (0x64) and ArbGasInfo (0x6C)
# are precompiles, but eth_getCode reports placeholder bytecode 0xfe (INVALID opcode).
# The fork imports that placeholder, so any call to them burns ALL forwarded gas and the
# game's tryGetArb*() fallbacks are left with nothing (claim dies out-of-gas). Clearing
# the code makes calls succeed-with-empty-data cheaply, exactly like in plain Hardhat
# tests, and the game falls back to its non-Arbitrum randomness sources.
rpc_result hardhat_setCode '["0x0000000000000000000000000000000000000064","0x"]' >/dev/null
rpc_result hardhat_setCode '["0x000000000000000000000000000000000000006C","0x"]' >/dev/null
echo "Arbitrum precompiles (0x64, 0x6C) neutralized for fork execution"

# Materialize the zero address locally: tools default their eth_call sender to 0x0,
# and a later upstream lookup for its nonce/balance would fail once the pin is pruned.
ZERO=0x0000000000000000000000000000000000000000
rpc_result hardhat_setNonce "[\"$ZERO\",\"0x0\"]" >/dev/null
rpc_result hardhat_setBalance "[\"$ZERO\",\"0x0\"]" >/dev/null
rpc_result hardhat_setCode "[\"$ZERO\",\"0x\"]" >/dev/null

echo "== 3/6 Prewarming production state =="
OWNER=$(game_call_addr $SEL_OWNER)
ROUND_NUM=$(game_call $SEL_ROUND_NUM)
LAST_BIDDER=$(game_call_addr $SEL_LAST_BIDDER)
PRIZE_TIME=$(game_call $SEL_PRIZE_TIME)
TIMEOUT=$(game_call $SEL_TIMEOUT)
echo "owner         : $OWNER"
echo "roundNum      : $ROUND_NUM"
echo "lastBidder    : $LAST_BIDDER"
echo "chain time    : $(fmt_ts "$(chain_now)")"
echo "mainPrizeTime : $(fmt_ts "$PRIZE_TIME")"
echo "claim timeout : $TIMEOUT s"
if [ -z "$ROUND_NUM" ] || [ -z "$PRIZE_TIME" ] || [ "$PRIZE_TIME" = 0 ] || [ "$OWNER" = "0x" ]; then
	echo "PREWARM FAILED — upstream cannot serve state at pinned block $PIN." >&2
	echo "Use an archive upstream: UPSTREAM_RPC=https://arb-mainnet.g.alchemy.com/v2/<key> $0" >&2
	exit 1
fi

# ---------- 3. maintenance window duration (impersonated owner) ----------
echo "== 4/6 Setting maintenance window: delayDurationBeforeRoundActivation=$DELAY_SECONDS =="
rpc_result hardhat_impersonateAccount "[\"$OWNER\"]" >/dev/null
rpc_result hardhat_setBalance "[\"$OWNER\",\"0x21E19E0C9BAB2400000\"]" >/dev/null
TX=$(send_tx "$OWNER" "$SEL_SET_DELAY$(printf '%064x' "$DELAY_SECONDS")" "0x186a0")
echo "tx: $TX | delay now: $(game_call $SEL_GET_DELAY) s"

# ---------- 4. offset chain time past mainPrizeTime + timeout ----------
echo "== 5/6 Offsetting chain time to mainPrizeTime + timeout + 1 =="
TARGET=$((PRIZE_TIME + TIMEOUT + 1))
NOW=$(chain_now)
if [ "$NOW" -lt "$TARGET" ]; then
	rpc_result evm_setNextBlockTimestamp "[$TARGET]" >/dev/null
	rpc_result evm_mine "[]" >/dev/null
fi
echo "chain time now: $(fmt_ts "$(chain_now)") (prize claimable by anyone)"

# ---------- 5. claim the main prize with account #0 ----------
echo "== 6/6 Claiming the main prize with hardhat account #0 =="
# 6M gas: the real production round-0 claim used 4,025,850 gas (tx 0x83b8f2…, Jun 15 2026),
# so 6M gives ~50% headroom. NOTE: this only suffices because the Arbitrum precompiles
# (0x64/0x6C) are neutralized above — with their 0xfe placeholder code in place, the claim
# burns ANY limit to ~98% and fails (all-gas-consuming INVALID opcode, not a gas shortage).
TX=$(send_tx "$ACCOUNT0" "$SEL_CLAIM" "0x5B8D80")
echo "claim tx: $TX"

ROUND_AFTER=$(game_call $SEL_ROUND_NUM)
ACTIVATION=$(game_call $SEL_ACTIVATION)
echo
echo "ROUND ENDED. roundNum: $ROUND_NUM -> $ROUND_AFTER"
echo "Maintenance window open until: $(fmt_ts "$ACTIVATION")"
echo "(chain time only advances ~10 s per mined block, so no wall-clock rush)"
echo
echo "Next: run the V2->V3 upgrade during this window:"
echo "  cd $HH_REPO && HARDHAT_MODE_CODE=2 FORK_RPC_URL=$FORK_RPC npx hardhat run scripts/upgrade-fork-to-v3.js --network fork_local"
echo "Node log: $LOG_FILE | pid: $(cat "$PID_FILE") | pinned block: $PIN"
