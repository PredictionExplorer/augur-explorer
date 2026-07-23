#!/bin/bash
# fork-v3-rehearsal.sh — orchestrates the V2->V3 upgrade rehearsal on a mainnet fork.
# Pure bash + curl + sed (+ go for the claim/bid steps). No python/jq required.
#
# Sequence (run the subcommands in this order):
#   status       show round/prize/activation times and chain clock (steps 1-2)
#   set-delay    impersonated owner sets delayDurationBeforeRoundActivation = 20 min (step 4)
#   jump         advance chain time to mainPrizeTime + timeout + 1 s (step 3)
#   claim        claim the main prize with hardhat account #0 via claimprize.go (steps 5-6)
#   upgrade      run the V2->V3 upgrade while the round is inactive (step 7)
#   new-round    advance past the activation time and place the first ETH bid (step 8)
#
# The contract sets the next round's activation time AT CLAIM TIME:
#   roundActivationTime = claim_block.timestamp + delayDurationBeforeRoundActivation
# (MainPrize.sol), so "set-delay" must run BEFORE "claim" to get claim + 20 min.
#
# Environment (override as needed):
#   RPC_URL           fork node                (default http://161.129.67.42:10545)
#   GAME_ADDR         game proxy               (default: production proxy)
#   COSMIC_SIG_REPO   Cosmic-Signature repo    (for the "upgrade" step)
#   NEW_ROUND_PKEY    key for the first bid    (default: hardhat account #0)
set -e -o pipefail

# Optional local settings file (survives script updates): current dir first, then script dir.
SELF_DIR="$(cd "$(dirname "$0")" && pwd)"
if [ -f "./fork-env.sh" ]; then . "./fork-env.sh"
elif [ -f "$SELF_DIR/fork-env.sh" ]; then . "$SELF_DIR/fork-env.sh"; fi

# RPC_URL > FORK_RPC_URL > derived from HH_HOST/HH_PORT > public-IP default.
HH_PORT=${HH_PORT:-10545}
if [ -z "$RPC_URL" ]; then
	if [ -n "$FORK_RPC_URL" ]; then RPC_URL="$FORK_RPC_URL"
	elif [ -n "$HH_HOST" ] && [ "$HH_HOST" != "0.0.0.0" ]; then RPC_URL="http://$HH_HOST:$HH_PORT"
	else RPC_URL="http://161.129.67.42:$HH_PORT"; fi
fi
GAME_ADDR=${GAME_ADDR:-0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2}
COSMIC_SIG_REPO=${COSMIC_SIG_REPO:-${HH_REPO:-$HOME/Cosmic-Signature}}
# Hardhat account #0 (publicly known dev key)
ACCOUNT0_PKEY=${NEW_ROUND_PKEY:-ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}
DELAY_SECONDS=1200   # 20 minutes maintenance window after the claim

SCRIPTS_DIR="$(cd "$(dirname "$0")" && pwd)"
export RPC_URL

# ---------- JSON-RPC helpers (pure bash/sed) ----------
rpc() { # rpc <method> <params-json>  -> raw JSON response
	curl -s -X POST "$RPC_URL" -H 'Content-Type: application/json' \
		-d "{\"jsonrpc\":\"2.0\",\"method\":\"$1\",\"params\":$2,\"id\":1}"
}

rpc_result() { # rpc_result <method> <params-json>  -> "result" string; dies on RPC error
	local resp
	resp=$(rpc "$1" "$2")
	if [ -z "$resp" ]; then
		echo "no response from $RPC_URL" >&2; exit 1
	fi
	case "$resp" in *'"error"'*)
		echo "RPC error from $1: $resp" >&2; exit 1;;
	esac
	# extracts string results ("result":"0x...") and literal results ("result":null)
	printf '%s\n' "$resp" | sed -n 's/.*"result":"\{0,1\}\([^",}]*\)"\{0,1\}[,}].*/\1/p'
}

hex2dec() { # 0x-prefixed or bare hex -> decimal (values fit in 63 bits here)
	local h=${1#0x}
	h=$(printf '%s' "$h" | sed 's/^0*//')
	[ -z "$h" ] && h=0
	echo $((16#$h))
}

game_call() { # game_call <calldata> -> decimal
	hex2dec "$(rpc_result eth_call "[{\"to\":\"$GAME_ADDR\",\"data\":\"$1\"},\"latest\"]")"
}

game_call_addr() { # game_call_addr <calldata> -> 0x-address (last 20 bytes of the word)
	local word
	word=$(rpc_result eth_call "[{\"to\":\"$GAME_ADDR\",\"data\":\"$1\"},\"latest\"]")
	word=${word#0x}
	echo "0x${word: -40}"
}

chain_now() { # latest block timestamp, decimal
	local ts
	ts=$(rpc eth_getBlockByNumber '["latest",false]' \
		| sed -n 's/.*"timestamp":"0x\([0-9a-fA-F]*\)".*/\1/p')
	if [ -z "$ts" ]; then echo "could not read latest block from $RPC_URL" >&2; exit 1; fi
	echo $((16#$ts))
}

fmt_ts() { echo "$1  ($(date -u -d "@$1" '+%Y-%m-%d %H:%M:%S') UTC)"; }

# Selectors (keccak4): mainPrizeTime() timeoutDurationToClaimMainPrize()
# roundActivationTime() roundNum() owner() lastBidderAddress()
# setDelayDurationBeforeRoundActivation(uint256) delayDurationBeforeRoundActivation()
SEL_PRIZE_TIME=0x18305de2
SEL_TIMEOUT=0x3b9d292e
SEL_ACTIVATION=0x6e970834
SEL_ROUND_NUM=0x119b22b3
SEL_OWNER=0x8da5cb5b
SEL_LAST_BIDDER=0xe5b3cd14
SEL_SET_DELAY=0x09794bee
SEL_GET_DELAY=0xb9cf9ba5

case "${1:-}" in

status) # ---- steps 1-2: fetch the times ----
	echo "game proxy      : $GAME_ADDR"
	echo "owner           : $(game_call_addr $SEL_OWNER)"
	echo "roundNum        : $(game_call $SEL_ROUND_NUM)"
	echo "lastBidder      : $(game_call_addr $SEL_LAST_BIDDER)"
	echo "chain time      : $(fmt_ts "$(chain_now)")"
	echo "roundActivation : $(fmt_ts "$(game_call $SEL_ACTIVATION)")"
	echo "mainPrizeTime   : $(fmt_ts "$(game_call $SEL_PRIZE_TIME)")"
	echo "claim timeout   : $(game_call $SEL_TIMEOUT) s"
	echo "activation delay: $(game_call $SEL_GET_DELAY) s"
	;;

set-delay) # ---- step 4: activation = claim time + 20 min (must run BEFORE claim) ----
	OWNER=$(game_call_addr $SEL_OWNER)
	echo "Impersonating owner $OWNER and setting delayDurationBeforeRoundActivation=$DELAY_SECONDS..."
	rpc_result hardhat_impersonateAccount "[\"$OWNER\"]" >/dev/null || true
	rpc_result hardhat_setBalance "[\"$OWNER\",\"0x21E19E0C9BAB2400000\"]" >/dev/null || true
	ARG=$(printf '%064x' $DELAY_SECONDS)
	TX=$(rpc_result eth_sendTransaction "[{\"from\":\"$OWNER\",\"to\":\"$GAME_ADDR\",\"data\":\"$SEL_SET_DELAY$ARG\",\"gas\":\"0x186a0\"}]")
	echo "tx: $TX"
	echo "new delay: $(game_call $SEL_GET_DELAY) s"
	;;

jump) # ---- step 3: chain time = mainPrizeTime + timeout + 1 ----
	PRIZE_TIME=$(game_call $SEL_PRIZE_TIME)
	TIMEOUT=$(game_call $SEL_TIMEOUT)
	TARGET=$((PRIZE_TIME + TIMEOUT + 1))
	NOW=$(chain_now)
	if [ "$NOW" -ge "$TARGET" ]; then
		echo "chain time $NOW already >= target $TARGET, nothing to do"
		exit 0
	fi
	echo "Jumping chain time: $(fmt_ts "$NOW") -> $(fmt_ts "$TARGET")"
	rpc_result evm_setNextBlockTimestamp "[$TARGET]" >/dev/null
	rpc_result evm_mine "[]" >/dev/null
	echo "chain time now: $(fmt_ts "$(chain_now)")"
	echo "Anyone may claim the main prize (last-bidder exclusivity expired)."
	;;

claim) # ---- steps 5-6: claim with hardhat account #0; round goes inactive ----
	cd "$SCRIPTS_DIR"
	PKEY_HEX=$ACCOUNT0_PKEY go run claimprize.go "$GAME_ADDR"
	echo
	echo "roundNum now        : $(game_call $SEL_ROUND_NUM)"
	echo "next roundActivation: $(fmt_ts "$(game_call $SEL_ACTIVATION)")"
	echo "Maintenance window open until activation — run '$0 upgrade' now."
	;;

upgrade) # ---- step 7: V2 -> V3 while the round is inactive ----
	if [ ! -f "$COSMIC_SIG_REPO/scripts/upgrade-fork-to-v3.js" ]; then
		echo "upgrade-fork-to-v3.js not found under COSMIC_SIG_REPO=$COSMIC_SIG_REPO" >&2
		echo "Run manually:  cd <Cosmic-Signature repo> && HARDHAT_MODE_CODE=2 npx hardhat run scripts/upgrade-fork-to-v3.js --network fork_local" >&2
		exit 1
	fi
	cd "$COSMIC_SIG_REPO"
	HARDHAT_MODE_CODE=2 GAME_PROXY_ADDR=$GAME_ADDR npx hardhat run scripts/upgrade-fork-to-v3.js --network fork_local
	;;

new-round) # ---- step 8: pass the activation time and place the first gesture ----
	ACTIVATION=$(game_call $SEL_ACTIVATION)
	NOW=$(chain_now)
	if [ "$NOW" -lt "$((ACTIVATION + 1))" ]; then
		echo "Jumping chain time past activation: $(fmt_ts $((ACTIVATION + 1)))"
		rpc_result evm_setNextBlockTimestamp "[$((ACTIVATION + 1))]" >/dev/null
		rpc_result evm_mine "[]" >/dev/null
	fi
	cd "$SCRIPTS_DIR"
	PKEY_HEX=$ACCOUNT0_PKEY go run bid.go "$GAME_ADDR"
	echo
	echo "roundNum now: $(game_call $SEL_ROUND_NUM) — new round is live on V3."
	;;

*)
	echo "Usage: $0 {status|set-delay|jump|claim|upgrade|new-round}"
	echo "Run in order: status -> set-delay -> jump -> claim -> upgrade -> new-round"
	exit 1
	;;
esac
