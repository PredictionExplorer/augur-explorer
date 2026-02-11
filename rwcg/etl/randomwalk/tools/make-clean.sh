#!/bin/bash
# Clean tool binaries in rwcg/etl/randomwalk/tools/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning randomwalk tools binaries..."

rm -f rw_toprated
# Other mains in this directory:
rm -f discord_bot
rm -f discord_ch_name
rm -f discord_user_limit
rm -f ffmpeg-convert
rm -f notif_bot
rm -f twauthorize
rm -f tweet_mints
rm -f twitteroob
rm -f twsend
rm -f twsend_image
rm -f twsend_img_reply

echo "Done cleaning randomwalk tools."
