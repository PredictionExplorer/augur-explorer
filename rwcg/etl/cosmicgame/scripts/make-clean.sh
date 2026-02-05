#!/bin/bash
# Clean script binaries in rwcg/etl/cosmicgame/scripts/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning cosmicgame scripts binaries..."

rm -f approved
rm -f autobid
rm -f bid
rm -f cginfo
rm -f claimprize
rm -f claimprize-delay60
rm -f deploy-erc20
rm -f donate
rm -f dwi-records
rm -f erc20allowance
rm -f erc20approve
rm -f erc20bal
rm -f erc20revoke
rm -f isapproved4all
rm -f owner
rm -f set_charity_percentage
rm -f set_main_prize_percentage
rm -f set_num_nft_winners
rm -f set_num_raffle_winners
rm -f set_raffle_percentage
rm -f set_staking_percentage
rm -f set_token_name
rm -f set-initial-duration-divisor
rm -f set-time-increment
rm -f setactivation
rm -f setroundactivation
rm -f tokownerof

# Clean swap files
rm -f .*.swp

echo "Done cleaning cosmicgame scripts."
