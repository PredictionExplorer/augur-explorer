#!/bin/bash
# List expected NFT images on a remote host, one `ls` per file.

SSH_HOST="${1:-}"
if [ -z "$SSH_HOST" ]; then
	echo "Usage: $0 <ssh-host>"
	exit 1
fi

DST_DIR="/home/frontend/nft-assets/new/cosmicsignature"
SSH="ssh -o BatchMode=yes -o ConnectTimeout=5 -l frontend $SSH_HOST"
HELPER="./img_upload"

if [ ! -x "$HELPER" ]; then
	echo "Error: helper not found or not executable at $HELPER"
	exit 1
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

	# Exactly one remote ls per file; on failure, print a clean MISSING line.
	$SSH "ls -l -- $REMOTE_PATH" || echo "MISSING	$REMOTE_PATH"

	COUNTER=$((COUNTER+1))
done
