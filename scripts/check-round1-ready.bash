#!/usr/bin/env bash

# Read-only sanity check that the CosmicSignatureGame proxy is running V3
# and that round 1 is (or is about to be) biddable.
#
# Usage:
#   ./scripts/check-round1-ready.bash [RPC_URL] [GAME_PROXY_ADDR]
#
# Only needs curl and python3; no Hardhat, no private keys, no writes.
# Exits 0 if all checks pass, 1 otherwise.

RPC_URL="${1:-https://rpc3.cosmicsignature.com:22946}"
GAME_ADDR="${2:-0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512}"

# ERC-1967 implementation slot.
IMPL_SLOT='0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'
# keccak256("Upgraded(address)")
UPGRADED_TOPIC='0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b'

# Function selectors (keccak-derived, verified against the deployed V3 branch).
SEL_ROUND_NUM='0x119b22b3'                 # roundNum()
SEL_ROUND_ACTIVATION_TIME='0x6e970834'     # roundActivationTime()
SEL_LAST_BIDDER='0xe5b3cd14'               # lastBidderAddress()
SEL_MAIN_PRIZE_TIME_INCREMENT='0xeb13430e' # mainPrizeTimeIncrementInMicroSeconds()
SEL_V3_PROBE='0x5fdf49cb'                  # mainPrizeNumCosmicSignatureNfts() - exists in V3 only
SEL_NEXT_ETH_BID_PRICE='0x4e452010'        # getNextEthBidPriceAdvanced(int256)
SEL_CLAIM_TIMEOUT='0x3b9d292e'             # timeoutDurationToClaimMainPrize()
SEL_MAIN_ETH_PRIZE='0x5b0a45d9'            # getMainEthPrizeAmount()

FAILURES=0

pass() { echo "  [PASS] $1"; }
fail() { echo "  [FAIL] $1"; FAILURES=$(( FAILURES + 1 )); }
warn() { echo "  [WARN] $1"; }
info() { echo "  [info] $1"; }

rpc() {
	# rpc <method> <params-json>
	curl -s --max-time 10 -X POST -H 'Content-Type: application/json' \
		--data '{"jsonrpc":"2.0","id":1,"method":"'"$1"'","params":'"$2"'}' \
		"${RPC_URL}"
}

# Extracts .result from a JSON-RPC response; prints nothing on error responses.
result_of() {
	python3 -c '
import sys, json
try:
	r = json.loads(sys.stdin.read())
except Exception:
	sys.exit(0)
if isinstance(r, dict) and "result" in r and r["result"] is not None:
	print(r["result"])
'
}

hex_to_dec() { python3 -c "print(int('$1', 16))"; }
wei_to_eth() { python3 -c "print(int('$1', 16) / 1e18)"; }

# eth_call helper: call <calldata>; prints hex result or nothing on revert.
call() {
	rpc eth_call '[{"to":"'"${GAME_ADDR}"'","data":"'"$1"'"},"latest"]' | result_of
}

echo "Checking game proxy ${GAME_ADDR}"
echo "via ${RPC_URL}"
echo

# ---------------------------------------------------------------- 1. RPC node.
echo "1. RPC node"
chain_id_hex="$(rpc eth_chainId '[]' | result_of)"
if [ -z "${chain_id_hex}" ]; then
	fail "RPC endpoint is not responding."
	echo
	echo "RESULT: NOT READY (${FAILURES} failed check(s))."
	exit 1
fi
chain_id="$(hex_to_dec "${chain_id_hex}")"
if [ "${chain_id}" = "31337" ]; then
	pass "Chain is reachable, chainId = 31337 (Hardhat)."
else
	fail "Unexpected chainId ${chain_id} (expected 31337)."
fi

latest_block_json="$(rpc eth_getBlockByNumber '["latest", false]')"
block_num_hex="$(echo "${latest_block_json}" | python3 -c 'import sys,json; print(json.load(sys.stdin)["result"]["number"])')"
block_ts_hex="$(echo "${latest_block_json}" | python3 -c 'import sys,json; print(json.load(sys.stdin)["result"]["timestamp"])')"
block_num="$(hex_to_dec "${block_num_hex}")"
block_ts="$(hex_to_dec "${block_ts_hex}")"
now="$(date +%s)"
info "Latest block: #${block_num}, timestamp $(date -d "@${block_ts}" '+%Y-%m-%d %H:%M:%S %Z' 2>/dev/null || date -r "${block_ts}"), $(( now - block_ts ))s behind system time."
if [ $(( now - block_ts )) -gt 86400 ]; then
	warn "Chain clock is more than a day behind system time. The node was likely started without HARDHAT_MODE_CODE=2, so initialDate \"2025-01-01\" applied. Time-based UI countdowns will look wrong."
fi
echo

# ------------------------------------------------- 2. Implementation version.
echo "2. Contract version"
impl_word="$(rpc eth_getStorageAt '["'"${GAME_ADDR}"'","'"${IMPL_SLOT}"'","latest"]' | result_of)"
impl_addr="0x${impl_word: -40}"
if [ "${impl_addr}" = "0x0000000000000000000000000000000000000000" ]; then
	fail "No implementation in the ERC-1967 slot; is ${GAME_ADDR} really the proxy?"
else
	info "Active implementation: ${impl_addr}."
fi

upgrade_count="$(rpc eth_getLogs '[{"address":"'"${GAME_ADDR}"'","fromBlock":"0x0","toBlock":"latest","topics":["'"${UPGRADED_TOPIC}"'"]}]' \
	| python3 -c 'import sys,json; print(len(json.load(sys.stdin).get("result",[])))')"
if [ "${upgrade_count}" = "3" ]; then
	pass "3 Upgraded events found (V1 birth -> V2 -> V3), as expected."
else
	fail "Expected 3 Upgraded events, found ${upgrade_count}."
fi

# mainPrizeNumCosmicSignatureNfts() exists only in V3; on V2 this call reverts.
v3_probe="$(call "${SEL_V3_PROBE}")"
if [ -n "${v3_probe}" ]; then
	pass "V3-only getter mainPrizeNumCosmicSignatureNfts() answers ($(hex_to_dec "${v3_probe}")) => this is V3, not V2."
else
	fail "V3-only getter reverted => the active implementation is NOT V3."
fi
echo

# ------------------------------------------------------------- 3. Round state.
echo "3. Round state"
round_num_hex="$(call "${SEL_ROUND_NUM}")"
if [ -z "${round_num_hex}" ]; then
	fail "roundNum() reverted."
	round_num=""
else
	round_num="$(hex_to_dec "${round_num_hex}")"
	if [ "${round_num}" = "1" ]; then
		pass "roundNum = 1 (round 0 was completed on V1, as designed)."
	else
		fail "roundNum = ${round_num}, expected 1."
	fi
fi

activation_hex="$(call "${SEL_ROUND_ACTIVATION_TIME}")"
if [ -z "${activation_hex}" ]; then
	fail "roundActivationTime() reverted."
else
	activation="$(hex_to_dec "${activation_hex}")"
	# The deploy config uses a far-future/sentinel value to keep a round parked.
	if [ "${activation}" -gt $(( now + 86400 )) ]; then
		fail "roundActivationTime = ${activation}, more than a day in the future => round is parked/blocked."
	elif [ "${activation}" -le "${now}" ]; then
		pass "roundActivationTime = ${activation} ($(date -d "@${activation}" '+%Y-%m-%d %H:%M:%S %Z' 2>/dev/null)) - already reached; bidding is open."
	else
		pass "roundActivationTime = ${activation} - activates in $(( activation - now ))s."
	fi
fi

last_bidder_word="$(call "${SEL_LAST_BIDDER}")"
if [ -n "${last_bidder_word}" ]; then
	last_bidder="0x${last_bidder_word: -40}"
	if [ "${last_bidder}" = "0x0000000000000000000000000000000000000000" ]; then
		info "lastBidderAddress = zero: no bids yet in this round (fresh round)."
	else
		info "lastBidderAddress = ${last_bidder}: the round already has bids."
	fi
fi
echo

# ----------------------------------------------------- 4. Timing parameters.
echo "4. Timing parameters"
incr_hex="$(call "${SEL_MAIN_PRIZE_TIME_INCREMENT}")"
if [ -z "${incr_hex}" ]; then
	fail "mainPrizeTimeIncrementInMicroSeconds() reverted."
else
	incr="$(hex_to_dec "${incr_hex}")"
	if [ "${incr}" = "0" ]; then
		fail "mainPrizeTimeIncrementInMicroSeconds = 0. The round-0 temporary hack was NOT restored; every bid would allow an instant claim!"
	else
		pass "mainPrizeTimeIncrementInMicroSeconds = ${incr} ($(python3 -c "print(${incr}/60e6)") minutes per bid)."
	fi
fi

timeout_hex="$(call "${SEL_CLAIM_TIMEOUT}")"
if [ -n "${timeout_hex}" ]; then
	timeout="$(hex_to_dec "${timeout_hex}")"
	info "timeoutDurationToClaimMainPrize = ${timeout}s ($(python3 -c "print(${timeout}/60)") minutes)."
fi
echo

# ------------------------------------------------------------ 5. Economics.
echo "5. Economics"
# getNextEthBidPriceAdvanced(0) - offset 0 means "as of the latest block".
price_hex="$(call "${SEL_NEXT_ETH_BID_PRICE}0000000000000000000000000000000000000000000000000000000000000000")"
if [ -z "${price_hex}" ]; then
	fail "getNextEthBidPriceAdvanced(0) reverted."
else
	price_dec="$(hex_to_dec "${price_hex}")"
	if [ "${price_dec}" -gt 0 ]; then
		pass "Next ETH bid price = $(wei_to_eth "${price_hex}") ETH."
	else
		fail "Next ETH bid price is 0."
	fi
fi

balance_hex="$(rpc eth_getBalance '["'"${GAME_ADDR}"'","latest"]' | result_of)"
if [ -n "${balance_hex}" ]; then
	info "Game contract balance = $(wei_to_eth "${balance_hex}") ETH (carried over from round 0 for the round-1 prize pool)."
fi
prize_hex="$(call "${SEL_MAIN_ETH_PRIZE}")"
if [ -n "${prize_hex}" ]; then
	info "Current main ETH prize amount = $(wei_to_eth "${prize_hex}") ETH."
fi
echo

# -------------------------------------------------------------------- Result.
if [ "${FAILURES}" = "0" ]; then
	echo "RESULT: READY. V3 is live and round ${round_num} is open (or opening) for bids."
	exit 0
else
	echo "RESULT: NOT READY (${FAILURES} failed check(s))."
	exit 1
fi
