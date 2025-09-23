#!/bin/bash
# List expected NFT images on a remote host, one `ls` per file.
# If an image is missing, generate it locally and upload it.

SSH_HOST="${1:-}"
if [ -z "$SSH_HOST" ]; then
	echo "Usage: $0 <ssh-host>"
	exit 1
fi

DST_DIR="/home/frontend/nft-assets/new/cosmicsignature"
SSH="ssh -o BatchMode=yes -o ConnectTimeout=5 -l frontend $SSH_HOST"
HELPER="./img_upload"
EXEC_CMD="./three_body_problem"
LOCAL_IMG_DIR="./pics"

if [ ! -x "$HELPER" ]; then
	echo "Error: helper not found or not executable at $HELPER"
	exit 1
fi

if [ ! -x "$EXEC_CMD" ]; then
	echo "Warning: generator not found or not executable at $EXEC_CMD"
	echo "Missing files will be reported but not generated."
fi

TOTAL_TOKENS="$($HELPER total_tokens)"
if ! [[ "$TOTAL_TOKENS" =~ ^[0-9]+$ ]]; then
	echo "Error: invalid total tokens: $TOTAL_TOKENS"
	exit 1
fi

COUNTER=0
while [ $COUNTER -lt "$TOTAL_TOKENS" ]; do
	SEED="$($HELPER token_seed "$COUNTER")"
	REMOTE_PATH="$DST_DIR/$SEED.png"

	# Exactly one remote ls per file.
	if ! $SSH "ls -l -- $REMOTE_PATH"; then
		echo "MISSING	$REMOTE_PATH"

		if [ -x "$EXEC_CMD" ]; then
			# Generate locally
			"$EXEC_CMD" --seed "$SEED" --file-name "$SEED"
			gen_rc=$?
			if [ $gen_rc -ne 0 ]; then
				echo "Aborting: generator failed for seed=$SEED (rc=$gen_rc)"
				exit 1
			fi

			LOCAL_IMG="$LOCAL_IMG_DIR/$SEED.png"
			if [ ! -f "$LOCAL_IMG" ]; then
				echo "Aborting: generated image not found at $LOCAL_IMG"
				exit 1
			fi

			# Upload
			scp "$LOCAL_IMG" "frontend@$SSH_HOST:$DST_DIR/$SEED.png"
			scp_rc=$?
			if [ $scp_rc -ne 0 ]; then
				echo "Aborting: scp failed for $LOCAL_IMG -> $REMOTE_PATH (rc=$scp_rc)"
				exit 1
			fi

			echo "UPLOADED	$REMOTE_PATH"
		fi
	fi

	COUNTER=$((COUNTER+1))
done

