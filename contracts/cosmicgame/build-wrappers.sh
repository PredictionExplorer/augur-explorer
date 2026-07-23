#!/usr/bin/env bash
#
# Compile Solidity contracts and generate Go bindings (wrappers) for cosmicgame.
# Single pass: compile all from compiled-Cosmic-Signature, then generate Go.
# Output directory: /tmp/cosmicgame-wrappers
#
# Prerequisites: solc at /usr/local/bin/solc-0.8.30 (V1) and solc-0.8.34 (V2/V3), abigen at ~/bin/abigen-1.15.5
#

set -e

readonly OUT_DIR="/tmp/cosmicgame-wrappers"
readonly ABIGEN="${HOME}/bin/abigen-1.15.5"
readonly SOLC="/usr/local/bin/solc-0.8.30"
readonly SOLC_V2="/usr/local/bin/solc-0.8.34"
readonly OPENZEPPELIN_5_1="/home/niko/eth/dev/openzeppelin/5.1"
readonly OPENZEPPELIN_5_02="/home/niko/eth/dev/openzeppelin/5.02"
readonly ARBITRUM="/home/niko/eth/dev/compiled-Cosmic-Signature/@arbitrum"
readonly WORK_DIR="/home/niko/eth/dev/compiled-Cosmic-Signature"
readonly BASE_PATH="/home/niko/eth/dev/b/cursor-vref"
readonly PRODUCTION_SYMLINK="${BASE_PATH}/production"

readonly CONTRACTS=(
	CosmicSignatureGame
	CosmicSignatureGameV2
	CosmicSignatureGameV3
	CosmicSignatureNft
	CosmicSignatureToken
	CharityWallet
	PrizesWallet
	MarketingWallet
	CosmicSignatureDao
	StakingWalletCosmicSignatureNft
	StakingWalletRandomWalkNft
)

# Prefer WORK_DIR if it has the .sol files; else use resolved production symlink
SRC_DIR=""
if [[ -d "$WORK_DIR" ]] && [[ -f "$WORK_DIR/CosmicSignatureGame.sol" ]]; then
	SRC_DIR="$WORK_DIR"
else
	RESOLVED=$(readlink -f "$PRODUCTION_SYMLINK" 2>/dev/null || true)
	if [[ -d "$RESOLVED" ]] && [[ -f "$RESOLVED/CosmicSignatureGame.sol" ]]; then
		SRC_DIR="$RESOLVED"
	fi
fi
if [[ -z "$SRC_DIR" ]]; then
	echo "Error: no source dir with .sol files. Tried: $WORK_DIR and $PRODUCTION_SYMLINK" >&2
	exit 1
fi

# --- Step 1: Compile Solidity (per-contract) ---
echo "=== Compiling Solidity (from $SRC_DIR) ==="
mkdir -p "$OUT_DIR"
cd "$SRC_DIR"

for contract in "${CONTRACTS[@]}"; do
	entry="${contract}.sol"
	if [[ ! -f "$entry" ]] && { [[ "$contract" == "CosmicSignatureGameV2" ]] || [[ "$contract" == "CosmicSignatureGameV3" ]]; }; then
		# Upgrade implementations live alongside V1 in production; the compiled source tree may lag.
		upgrade_src="/home/niko/eth/dev/b/todays/Cosmic-Signature/contracts/production/${entry}"
		if [[ -f "$upgrade_src" ]]; then
			entry="$upgrade_src"
		fi
	fi
	if [[ ! -f "$entry" ]]; then
		echo "  Warning: skip $contract ($entry not found)" >&2
		continue
	fi
	comb_dir="$OUT_DIR/${contract}-combined"
	mkdir -p "$comb_dir"

	if [[ "$contract" == "CosmicSignatureGame" ]] || [[ "$contract" == "CosmicSignatureGameV2" ]] || [[ "$contract" == "CosmicSignatureGameV3" ]]; then
		remaps=(":@openzeppelin=$OPENZEPPELIN_5_1")
	elif [[ "$contract" == "PrizesWallet" ]]; then
		remaps=(":@openzeppelin=$OPENZEPPELIN_5_02")
	else
		remaps=(":@openzeppelin=$OPENZEPPELIN_5_02")
	fi
	# Arbitrum remap needed by CosmicSignatureGame, PrizesWallet, and any contract using ArbitrumHelpers (e.g. staking, CosmicSignatureNft)
	[[ -d "$ARBITRUM" ]] && remaps+=(":@arbitrum=$ARBITRUM")

	echo "  $contract"
	solc_bin="$SOLC"
	{ [[ "$contract" == "CosmicSignatureGameV2" ]] || [[ "$contract" == "CosmicSignatureGameV3" ]]; } && solc_bin="$SOLC_V2"
	if ! "$solc_bin" --overwrite --via-ir --combined-json bin,abi,userdoc,devdoc,metadata "${remaps[@]}" "$entry" -o "$comb_dir" 2>"$OUT_DIR/solc-errors-$contract.txt"; then
		echo "    Error: solc failed, see $OUT_DIR/solc-errors-$contract.txt" >&2
		continue
	fi
	if [[ ! -s "$comb_dir/combined.json" ]]; then
		echo "    Error: no combined.json produced" >&2
		continue
	fi
done

# --- Step 2: Generate Go wrappers ---
echo "=== Generating Go wrappers ==="
PKG="cosmicgame"

for contract in "${CONTRACTS[@]}"; do
	combined="$OUT_DIR/${contract}-combined/combined.json"
	outfile="$OUT_DIR/${contract}.go"
	if [[ ! -s "$combined" ]]; then
		echo "  Warning: skip $contract (no combined.json)" >&2
		continue
	fi
	abigen_args=(--combined-json "$combined" --pkg "$PKG" --type "$contract" --out "$outfile")
	# V2/V3 share mixin types already generated from V1; keep only each main implementation.
	if [[ "$contract" == "CosmicSignatureGameV2" ]] || [[ "$contract" == "CosmicSignatureGameV3" ]]; then
		upgrade_exc=$(python3 - "$combined" "$contract" <<'PY'
import json, sys
with open(sys.argv[1]) as f:
    names = sorted(set(k.split(":")[1] for k in json.load(f)["contracts"]))
print(",".join(f"*:{n}" for n in names if n != sys.argv[2]))
PY
)
		abigen_args+=(--exc="$upgrade_exc")
	fi
	if "$ABIGEN" "${abigen_args[@]}" 2>/dev/null; then
		echo "  $contract -> $outfile"
	else
		echo "  Warning: abigen failed for $contract" >&2
		rm -f "$outfile"
	fi
done

echo "=== Done. Wrappers in $OUT_DIR ==="
ls -la "$OUT_DIR"/*.go 2>/dev/null || true
