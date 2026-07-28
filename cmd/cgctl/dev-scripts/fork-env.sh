# fork-env.sh — local settings for the fork rehearsal scripts.
# Put this file next to the scripts (or in the directory you run them from).
# Every fork-*.sh script sources it automatically, so your host/port/repo
# settings survive script updates. Environment variables still win over
# values set here (each assignment respects an already-exported value).

# Interface the Hardhat fork node binds to, and where clients reach it.
HH_HOST=${HH_HOST:-161.129.67.42}
HH_PORT=${HH_PORT:-10545}

# Explicit RPC URL for reaching the fork node (overrides the derived default).
# Derived default: http://127.0.0.1:$HH_PORT when HH_HOST=0.0.0.0, else http://$HH_HOST:$HH_PORT
#FORK_RPC_URL=http://161.129.67.42:10545

# Arbitrum One node to fork from (archive endpoint recommended).
UPSTREAM_RPC=${UPSTREAM_RPC:-http://69.10.55.2:38545}
#UPSTREAM_RPC=https://arb-mainnet.g.alchemy.com/v2/<your-key>

# Cosmic-Signature repo (hardhat project with scripts/upgrade-fork-to-v3.js).
HH_REPO=${HH_REPO:-$HOME/Cosmic-Signature}

# Game proxy address.
GAME_ADDR=${GAME_ADDR:-0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2}
