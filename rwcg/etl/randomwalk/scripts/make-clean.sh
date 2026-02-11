#!/bin/bash
# Clean script binaries in rwcg/etl/randomwalk/scripts/

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "Cleaning randomwalk scripts binaries..."

rm -f mint
rm -f ownerof
rm -f price
rm -f status
# Binaries not in make.sh but have main():
rm -f accept_offer
rm -f approve
rm -f cancel_offer
rm -f new_offer
rm -f scan_rwmints
rm -f scan_transfers
rm -f setname
rm -f statusmkt
rm -f tokenuri
rm -f transfer
rm -f verify_erc20_transfers
rm -f verify_owner
rm -f withdrawal

echo "Done cleaning randomwalk scripts."
