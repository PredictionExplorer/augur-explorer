// Set the game's next ETH bid price to an arbitrary value on a dev chain by
// writing contract storage directly (hardhat_setStorageAt). No owner setter
// exists for prices, so this is dev-chain-only.
//
// Usage (from the Cosmic-Signature repo, so hardhat + artifacts resolve):
//   CADDR=0x... BID_PRICE_ETH=1 \
//     npx hardhat run <path>/set-bid-price.js --network hardhat_on_localhost
//
// Env vars:
//   CADDR          game proxy address (required)
//   BID_PRICE_ETH  target price of the NEXT ETH bid, in ETH (default 1)
//
// How it works. The effective price (getNextEthBidPrice) comes from one of
// two storage variables:
//   - nextEthBidPrice                   once the round has a bid — the getter
//                                       returns it verbatim, so we write the
//                                       target directly;
//   - ethDutchAuctionBeginningBidPrice  while the round has NO bids — the
//                                       auction declines from it over time,
//                                       so we calibrate: write the target,
//                                       read the effective price, scale
//                                       linearly, and write again.
// Storage slots are NOT hardcoded (the layout hides behind a 256-slot
// reserved gap and changes between versions). The script finds the slot by
// scanning the proxy storage for the variable's current value and
// confirming candidates with a write-marker probe through the getter.

"use strict";

const hre = global.hre ?? require("hardhat");

const CADDR = process.env.CADDR;
const MAX_SLOT = 2048; // scan this many slots; layout gap is 256, game vars follow
const SCAN_BATCH = 256;

const GAME_ABI = [
    "function roundNum() view returns (uint256)",
    "function roundActivationTime() view returns (uint256)",
    "function lastBidderAddress() view returns (address)",
    "function getNextEthBidPrice() view returns (uint256)",
    "function nextEthBidPrice() view returns (uint256)",
    "function ethDutchAuctionBeginningBidPrice() view returns (uint256)",
];

const word = (v) => hre.ethers.toBeHex(v, 32);

async function readSlot(provider, slot) {
    return BigInt(await provider.send("eth_getStorageAt", [CADDR, word(slot), "latest"]));
}

async function writeSlot(provider, slot, value) {
    await provider.send("hardhat_setStorageAt", [CADDR, word(slot), word(value)]);
}

// findSlot locates the storage slot backing a public variable: collect the
// slots holding the variable's current value, then write a marker into each
// candidate and keep the one the getter echoes back. Every probe restores
// the original word before moving on.
async function findSlot(provider, game, getterName) {
    const current = await game[getterName]();
    const candidates = [];
    for (let base = 0; base < MAX_SLOT; base += SCAN_BATCH) {
        const reads = [];
        for (let i = base; i < base + SCAN_BATCH; i++) {
            reads.push(readSlot(provider, i));
        }
        const values = await Promise.all(reads);
        values.forEach((v, j) => {
            if (v === current) {
                candidates.push(base + j);
            }
        });
    }
    if (candidates.length === 0) {
        throw new Error(`no slot in 0..${MAX_SLOT} holds the current value of ${getterName} (${current})`);
    }
    const marker = 0xC0FFEE0000000000000000000000000000000000000000000000000000n;
    for (const slot of candidates) {
        const original = await readSlot(provider, slot);
        await writeSlot(provider, slot, marker + BigInt(slot));
        const got = await game[getterName]();
        await writeSlot(provider, slot, original);
        if (got === marker + BigInt(slot)) {
            return slot;
        }
    }
    throw new Error(
        `${getterName}: none of the value-matching slots [${candidates}] responds to a write probe`,
    );
}

async function main() {
    if (!CADDR) {
        throw new Error("CADDR env var is required (game proxy address).");
    }
    const target = hre.ethers.parseEther(process.env.BID_PRICE_ETH || "1");
    const provider = hre.ethers.provider;
    const game = new hre.ethers.Contract(CADDR, GAME_ABI, provider);

    const roundNum = await game.roundNum();
    const hasBids = (await game.lastBidderAddress()) !== hre.ethers.ZeroAddress;
    const before = await game.getNextEthBidPrice();
    console.log(`round=${roundNum} hasBids=${hasBids}`);
    console.log(`current next ETH bid price = ${hre.ethers.formatEther(before)} ETH`);
    console.log(`target  next ETH bid price = ${hre.ethers.formatEther(target)} ETH`);

    if (hasBids) {
        // Mid-round: the getter returns nextEthBidPrice verbatim.
        const slot = await findSlot(provider, game, "nextEthBidPrice");
        console.log(`nextEthBidPrice storage slot = ${slot}`);
        await writeSlot(provider, slot, target);
    } else {
        // No bids yet: the price is a Dutch auction declining from
        // ethDutchAuctionBeginningBidPrice. The effective price is linear in
        // the beginning price (p = B * k for the current timestamp), so one
        // proportional correction lands on the target (within rounding).
        const slot = await findSlot(provider, game, "ethDutchAuctionBeginningBidPrice");
        console.log(`ethDutchAuctionBeginningBidPrice storage slot = ${slot}`);
        await writeSlot(provider, slot, target);
        const effective = await game.getNextEthBidPrice();
        if (effective !== target && effective > 0n) {
            const corrected = (target * target) / effective;
            await writeSlot(provider, slot, corrected);
            console.log(
                `auction is mid-decline: calibrated beginning price to ` +
                `${hre.ethers.formatEther(corrected)} ETH`,
            );
            console.log(
                "NOTE: the auction keeps declining, so the effective price drifts " +
                "below the target as chain time advances (exact at calibration time).",
            );
        }
    }

    const after = await game.getNextEthBidPrice();
    console.log(`new next ETH bid price = ${hre.ethers.formatEther(after)} ETH`);
}

main().then(() => process.exit(0)).catch((e) => { console.error(e); process.exit(1); });
