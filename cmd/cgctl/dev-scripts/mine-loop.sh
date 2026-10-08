#!/bin/bash
# Periodically mine blocks on a Hardhat/dev node so block.timestamp keeps
# tracking wall-clock time. Hardhat only stamps a new timestamp when a block
# is mined, so contract-side countdowns (round activation delay, main prize
# timer, Dutch auction prices) appear frozen between transactions. Run this
# in a spare terminal while testing the frontend against a local chain.
#
# Usage: ./mine-loop.sh [interval_seconds] [count]
#   interval_seconds  seconds between blocks (default 5)
#   count             number of blocks to mine (default: forever)
#
# Env: RPC_URL (default http://localhost:8545)
#      INCREASE_TIME  optional seconds of evm_increaseTime before each mine
#                     (use to fast-forward, e.g. INCREASE_TIME=60 for 12x speed)

RPC_URL="${RPC_URL:-http://localhost:8545}"
INTERVAL="${1:-5}"
COUNT="${2:-0}" # 0 = forever

rpc() {
	curl -s -X POST -H 'Content-Type: application/json' \
		--data "{\"jsonrpc\":\"2.0\",\"method\":\"$1\",\"params\":[$2],\"id\":1}" \
		"$RPC_URL"
}

mined=0
while :; do
	if [ -n "$INCREASE_TIME" ]; then
		rpc evm_increaseTime "$INCREASE_TIME" >/dev/null
	fi
	block=$(rpc evm_mine "" | sed 's/.*result[^0-9a-fx]*"\?\([0-9a-fx]*\)"\?.*/\1/')
	mined=$((mined + 1))
	echo "$(date '+%H:%M:%S') mined block ($mined)${block:+ result=$block}"
	[ "$COUNT" -gt 0 ] && [ "$mined" -ge "$COUNT" ] && break
	sleep "$INTERVAL"
done
