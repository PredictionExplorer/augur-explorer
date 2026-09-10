// Adjust the game's time parameters on a running dev chain (works on V2 and V3).
//
// Usage (from the Cosmic-Signature repo, so hardhat + artifacts resolve):
//   CADDR=0x... TIME_INCREMENT_SEC=300 ROUND_DELAY_SEC=600 \
//     npx hardhat run <path>/set-time-params.js --network hardhat_on_localhost
//
// Env vars (set only the ones you want to change):
//   CADDR               game proxy address (required)
//   TIME_INCREMENT_SEC  per-bid main prize time bump, seconds
//   ROUND_DELAY_SEC     delay before the next round activates after a round ends, seconds
//   TIMEOUT_CLAIM_SEC   anyone-can-claim timeout after main prize time, seconds
//   ACTIVATION_DELAY_SEC  when the round has to be deactivated to apply a change,
//                         re-activate it this many seconds later (default 5)
//
// Constraints (enforced by the contracts):
//   - ROUND_DELAY_SEC is onlyOwner and can be changed at any time.
//   - TIME_INCREMENT_SEC and TIMEOUT_CLAIM_SEC require an inactive round. If the
//     current round is active but has NO bids, this script temporarily deactivates
//     it, applies the change, and re-activates it. If a bid has already been placed
//     in the current round, the change is impossible until the round completes —
//     the script aborts with a message; claim the main prize first, then re-run
//     during the between-rounds window.

"use strict";

const hre = require("hardhat");

const CADDR = process.env.CADDR;
const ACTIVATION_DELAY_SEC = BigInt(process.env.ACTIVATION_DELAY_SEC || "5");

const GAME_ABI = [
    "function roundNum() view returns (uint256)",
    "function roundActivationTime() view returns (uint256)",
    "function lastBidderAddress() view returns (address)",
    "function mainPrizeTimeIncrementInMicroSeconds() view returns (uint256)",
    "function delayDurationBeforeRoundActivation() view returns (uint256)",
    "function timeoutDurationToClaimMainPrize() view returns (uint256)",
    "function setMainPrizeTimeIncrementInMicroSeconds(uint256)",
    "function setDelayDurationBeforeRoundActivation(uint256)",
    "function setTimeoutDurationToClaimMainPrize(uint256)",
    "function setRoundActivationTime(uint256)",
];

async function main() {
    if (!CADDR) {
        throw new Error("CADDR env var is required (game proxy address).");
    }
    const [owner] = await hre.ethers.getSigners();
    const game = new hre.ethers.Contract(CADDR, GAME_ABI, owner);
    const g = { gasLimit: 30_000_000 };

    const roundNum = await game.roundNum();
    const activationTime = await game.roundActivationTime();
    const lastBidder = await game.lastBidderAddress();
    // The chain may be idle (auto-mine, no interval mining), leaving the latest
    // block timestamp stale. The next transaction gets a wall-clock timestamp,
    // so judge round activity against wall-clock time.
    const latestBlock = await hre.ethers.provider.getBlock("latest");
    const now = BigInt(Math.max(Number(latestBlock.timestamp), Math.floor(Date.now() / 1000)));
    const roundIsActive = now >= activationTime;
    const hasBids = lastBidder !== hre.ethers.ZeroAddress;

    console.log(`roundNum=${roundNum} active=${roundIsActive} hasBids=${hasBids}`);
    console.log("Current values:");
    console.log(`  mainPrizeTimeIncrement          = ${(await game.mainPrizeTimeIncrementInMicroSeconds()) / 1_000_000n} s`);
    console.log(`  delayDurationBeforeRoundActivation = ${await game.delayDurationBeforeRoundActivation()} s`);
    console.log(`  timeoutDurationToClaimMainPrize = ${await game.timeoutDurationToClaimMainPrize()} s`);

    // Changes that are allowed at any time.
    if (process.env.ROUND_DELAY_SEC != null) {
        const v = BigInt(process.env.ROUND_DELAY_SEC);
        await (await game.setDelayDurationBeforeRoundActivation(v, g)).wait();
        console.log(`Set delayDurationBeforeRoundActivation = ${v} s`);
    }

    // Changes that require an inactive round.
    const inactiveOnlyChanges = [];
    if (process.env.TIME_INCREMENT_SEC != null) {
        const v = BigInt(process.env.TIME_INCREMENT_SEC);
        inactiveOnlyChanges.push({
            label: `mainPrizeTimeIncrement = ${v} s`,
            apply: () => game.setMainPrizeTimeIncrementInMicroSeconds(v * 1_000_000n, g),
        });
    }
    if (process.env.TIMEOUT_CLAIM_SEC != null) {
        const v = BigInt(process.env.TIMEOUT_CLAIM_SEC);
        inactiveOnlyChanges.push({
            label: `timeoutDurationToClaimMainPrize = ${v} s`,
            apply: () => game.setTimeoutDurationToClaimMainPrize(v, g),
        });
    }

    if (inactiveOnlyChanges.length > 0) {
        let mustReactivate = false;
        if (roundIsActive) {
            if (hasBids) {
                console.error(
                    "ERROR: a bid has already been placed in the current round, so the round " +
                    "cannot be deactivated and TIME_INCREMENT_SEC / TIMEOUT_CLAIM_SEC cannot be " +
                    "changed now. Finish the round (claim the main prize), then re-run this " +
                    "script during the between-rounds window."
                );
                process.exit(1);
            }
            // Round is active but bid-free: park activation in the future, change, re-arm.
            await (await game.setRoundActivationTime(now + 86400n, g)).wait();
            console.log("Temporarily deactivated the round.");
            mustReactivate = true;
        }
        for (const change of inactiveOnlyChanges) {
            await (await change.apply()).wait();
            console.log(`Set ${change.label}`);
        }
        if (mustReactivate) {
            const blk = await hre.ethers.provider.getBlock("latest");
            const activationTs = BigInt(blk.timestamp) + ACTIVATION_DELAY_SEC;
            await (await game.setRoundActivationTime(activationTs, g)).wait();
            console.log(`Round re-activates at ${activationTs} (in ~${ACTIVATION_DELAY_SEC}s of chain time).`);
        }
    }

    console.log("New values:");
    console.log(`  mainPrizeTimeIncrement          = ${(await game.mainPrizeTimeIncrementInMicroSeconds()) / 1_000_000n} s`);
    console.log(`  delayDurationBeforeRoundActivation = ${await game.delayDurationBeforeRoundActivation()} s`);
    console.log(`  timeoutDurationToClaimMainPrize = ${await game.timeoutDurationToClaimMainPrize()} s`);
}

main().then(() => process.exit(0)).catch((e) => { console.error(e); process.exit(1); });
