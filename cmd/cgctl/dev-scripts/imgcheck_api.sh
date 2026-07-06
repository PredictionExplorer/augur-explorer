#!/bin/bash
# CosmicSignature NFT image/video checker and uploader (API-based).
# 1) Fetches all tokens from the CosmicGame public API
# 2) Determines which images/videos are missing on the destination host
# 3) Regenerates only missing artifacts via run.py (or GENERATOR)
# 4) Uploads the missing artifacts
#
# Usage: imgcheck_api.sh <ssh-host> [api_base_url]
# Example: imgcheck_api.sh my.server.com
# Example: imgcheck_api.sh my.server.com http://161.129.67.42:8353
#
# Generator: defaults to run.py in cwd. Override with GENERATOR:
#   GENERATOR=/path/to/run.py ./imgcheck_api.sh my.server.com

SSH_HOST="${1:-69.10.55.2}"
API_BASE_URL="${2:-http://161.129.67.42:8353}"

if [ -z "$SSH_HOST" ]; then
	echo "Usage: $0 <ssh-host> [api_base_url]"
	echo "  api_base_url  defaults to http://161.129.67.42:8353"
	echo "  Set GENERATOR=/path/to/run.py to use a different generator."
	exit 1
fi

DST_DIR="/home/frontend/nft-assets/new/cosmicsignature"
# Accept new host keys on first connect (avoid "Host key verification failed"); use SSH_OPTS_EXTRA to override
SSH_OPTS="-o BatchMode=yes -o ConnectTimeout=5 -o StrictHostKeyChecking=accept-new ${SSH_OPTS_EXTRA:-}"
SSH="ssh $SSH_OPTS -l frontend $SSH_HOST"
if [ -n "$GENERATOR" ]; then
	EXEC_CMD="$GENERATOR"
elif [ -x "./run.py" ]; then
	EXEC_CMD="./run.py"
elif [ -f "./run.py" ]; then
	EXEC_CMD="python3 ./run.py"
elif [ -x "./three_body_problem" ]; then
	EXEC_CMD="./three_body_problem"
elif [ -x "./target/release/three_body_problem" ]; then
	EXEC_CMD="./target/release/three_body_problem"
else
	EXEC_CMD="./run.py"
fi
LOCAL_IMG_DIR="./pics"
LOCAL_VID_DIR="./vids"

# Token list API: returns JSON with .CosmicSignatureTokenList[].Seed
CST_LIST_URL="${API_BASE_URL}/api/cosmicgame/cst/list/all/0/100000"

# Check generator exists: only the path (first word, or second word after python3) is checked
_EXEC_FIRST="${EXEC_CMD%% *}"
if [ "$_EXEC_FIRST" = "python3" ] || [ "$_EXEC_FIRST" = "python" ]; then
	_REST="${EXEC_CMD#* }"
	_EXEC_PATH="${_REST%% *}"
else
	_EXEC_PATH="$_EXEC_FIRST"
fi
GENERATOR_OK=0
if [ -x "$_EXEC_PATH" ] || [ -f "$_EXEC_PATH" ]; then
	GENERATOR_OK=1
elif [ "$_EXEC_FIRST" = "python3" ] || [ "$_EXEC_FIRST" = "python" ]; then
	command -v "$_EXEC_FIRST" >/dev/null 2>&1 && [ -f "$_EXEC_PATH" ] && GENERATOR_OK=1
fi
if [ "$GENERATOR_OK" -eq 0 ]; then
	echo "Warning: generator not found at $_EXEC_PATH"
	echo "Missing files will be reported but not generated."
	echo "Tried: GENERATOR, ./run.py, ./three_body_problem, ./target/release/three_body_problem"
fi
unset _EXEC_FIRST _EXEC_PATH _REST

# Fetch token list from API
echo "Fetching token list from $CST_LIST_URL ..."
RESP="$(curl -sS -f "$CST_LIST_URL")" || {
	echo "Error: failed to fetch token list (curl exit $?)"
	exit 1
}
API_STATUS="$(echo "$RESP" | jq -r '.status // 0')"
if [ "$API_STATUS" != "1" ]; then
	echo "Error: API returned status=$API_STATUS ($(echo "$RESP" | jq -r '.error // "unknown"'))"
	exit 1
fi

# Extract seeds (normalize: strip 0x/0X prefix for consistent 0x$SEED filenames)
SEEDS="$(echo "$RESP" | jq -r '.CosmicSignatureTokenList[]? | .Seed // empty' | while read -r s; do
	[ -z "$s" ] && continue
	s="${s#0x}"; s="${s#0X}"
	printf '%s\n' "$s"
done)"

TOTAL="$(echo "$SEEDS" | grep -c . || true)"
if [ -z "$TOTAL" ] || [ "$TOTAL" -eq 0 ]; then
	echo "Error: no tokens returned from API or invalid response"
	echo "Response sample: $(echo "$RESP" | head -c 200)"
	exit 1
fi
echo "Found $TOTAL tokens from API."

COUNTER=0
while IFS= read -r SEED; do
	[ -z "$SEED" ] && continue
	# Filename uses 0x prefix (e.g. 0x1b2E85...)
	REMOTE_IMG="$DST_DIR/0x${SEED}.png"
	REMOTE_VID="$DST_DIR/0x${SEED}.mp4"
	MUST_REGEN=0

	if ! $SSH "ls -l -- \"$REMOTE_IMG\"" </dev/null 2>/dev/null; then
		echo "MISSING	$REMOTE_IMG"
		MUST_REGEN=1
	fi
	if ! $SSH "ls -l -- \"$REMOTE_VID\"" </dev/null 2>/dev/null; then
		echo "MISSING	$REMOTE_VID"
		MUST_REGEN=1
	fi

	if [ "$MUST_REGEN" -eq 1 ]; then
		if [ "$GENERATOR_OK" -eq 1 ]; then
			echo "GENERATE	seed=0x$SEED"
			$EXEC_CMD --seed "0x$SEED" --file-name "0x$SEED" </dev/null 1>/dev/null 2>&1
			gen_rc=$?
			if [ $gen_rc -ne 0 ]; then
				echo "Aborting: generator failed for seed=0x$SEED (rc=$gen_rc)"
				exit 1
			fi
			# Resolve local paths: generator may output 0xSEED.png or enhanced_0xSEED.png / classic_0xSEED.png (run.py)
			LOCAL_IMG=""
			for cand in "$LOCAL_IMG_DIR/0x$SEED.png" "$LOCAL_IMG_DIR/enhanced_0x$SEED.png" "$LOCAL_IMG_DIR/classic_0x$SEED.png"; do
				if [ -f "$cand" ]; then LOCAL_IMG="$cand"; break; fi
			done
			LOCAL_VID=""
			for cand in "$LOCAL_VID_DIR/0x$SEED.mp4" "$LOCAL_VID_DIR/enhanced_0x$SEED.mp4" "$LOCAL_VID_DIR/classic_0x$SEED.mp4"; do
				if [ -f "$cand" ]; then LOCAL_VID="$cand"; break; fi
			done
			if [ -z "$LOCAL_IMG" ] || [ ! -f "$LOCAL_IMG" ]; then
				echo "Aborting: generated image not found (tried 0x$SEED.png, enhanced_*, classic_* in $LOCAL_IMG_DIR)"
				exit 1
			fi
			if [ -z "$LOCAL_VID" ] || [ ! -f "$LOCAL_VID" ]; then
				echo "Aborting: generated video not found (tried 0x$SEED.mp4, enhanced_*, classic_* in $LOCAL_VID_DIR)"
				exit 1
			fi
			scp $SSH_OPTS "$LOCAL_IMG" "frontend@$SSH_HOST:$REMOTE_IMG" </dev/null || {
				echo "Aborting: scp failed for $LOCAL_IMG -> $REMOTE_IMG"
				exit 1
			}
			echo "UPLOADED	$REMOTE_IMG"
			scp $SSH_OPTS "$LOCAL_VID" "frontend@$SSH_HOST:$REMOTE_VID" </dev/null || {
				echo "Aborting: scp failed for $LOCAL_VID -> $REMOTE_VID"
				exit 1
			}
			echo "UPLOADED	$REMOTE_VID"
		fi
	else
		echo "OK	seed=0x$SEED	($REMOTE_IMG & $REMOTE_VID present)"
	fi
	COUNTER=$((COUNTER + 1))
done <<< "$SEEDS"

echo "Done. Processed $COUNTER tokens."
