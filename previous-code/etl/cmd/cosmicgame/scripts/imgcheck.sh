#!/bin/bash
# List expected NFT artifacts (PNG + MP4) on a remote host, one `ls` per artifact.
# If any artifact is missing, generate locally and upload both.

SSH_HOST="${1:-}"
if [ -z "$SSH_HOST" ]; then
	echo "Usage: $0 <ssh-host>"
	exit 1
fi

DST_DIR="/home/frontend/nft-assets/new/cosmicsignature"
SSH="ssh -o BatchMode=yes -o ConnectTimeout=5 -l frontend $SSH_HOST"
HELPER="./img_upload"
EXEC_CMD="./three_body_problem"
LOCAL_IMG_DIR="./pics"	# expects $SEED.png and $SEED.mp4 here
LOCAL_VID_DIR="./vids"
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

	REMOTE_IMG="$DST_DIR/0x$SEED.png"
	REMOTE_VID="$DST_DIR/0x$SEED.mp4"

	MUST_REGEN=0

	# Exactly one remote `ls` per artifact
	if ! $SSH "ls -l -- \"$REMOTE_IMG\""; then
		echo "MISSING	$REMOTE_IMG"
		MUST_REGEN=1
	fi

	if ! $SSH "ls -l -- \"$REMOTE_VID\""; then
		echo "MISSING	$REMOTE_VID"
		MUST_REGEN=1
	fi

	if [ $MUST_REGEN -eq 1 ]; then
		if [ -x "$EXEC_CMD" ]; then
			echo "GENERATE	seed=$SEED"
			"$EXEC_CMD" --seed "$SEED" --file-name "0x$SEED" 1>/dev/null 2>&1
			gen_rc=$?
			if [ $gen_rc -ne 0 ]; then
				echo "Aborting: generator failed for seed=$SEED (rc=$gen_rc)"
				exit 1
			fi

			LOCAL_IMG="$LOCAL_IMG_DIR/0x$SEED.png"
			LOCAL_VID="$LOCAL_VID_DIR/0x$SEED.mp4"

			if [ ! -f "$LOCAL_IMG" ]; then
				echo "Aborting: generated image not found at $LOCAL_IMG"
				exit 1
			fi
			if [ ! -f "$LOCAL_VID" ]; then
				echo "Aborting: generated video not found at $LOCAL_VID"
				exit 1
			fi

			scp "$LOCAL_IMG" "frontend@$SSH_HOST:$REMOTE_IMG"
			scp_rc=$?
			if [ $scp_rc -ne 0 ]; then
				echo "Aborting: scp failed for $LOCAL_IMG -> $REMOTE_IMG (rc=$scp_rc)"
				exit 1
			fi
			echo "UPLOADED	$REMOTE_IMG"

			scp "$LOCAL_VID" "frontend@$SSH_HOST:$REMOTE_VID"
			scp_rc=$?
			if [ $scp_rc -ne 0 ]; then
				echo "Aborting: scp failed for $LOCAL_VID -> $REMOTE_VID (rc=$scp_rc)"
				exit 1
			fi
			echo "UPLOADED	$REMOTE_VID"
		fi
	else
		echo "OK	seed=$SEED	($REMOTE_IMG & $REMOTE_VID present)"
	fi

	COUNTER=$((COUNTER+1))
done

