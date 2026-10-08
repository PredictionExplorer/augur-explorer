/**
 * upgrade-v3.js — upgrade a live CosmicSignatureGame proxy (on V2) to CosmicSignatureGameV3.
 * Companion to deploy-v2.js: deploy → play on V2 → run this → keep playing on V3.
 *
 * Flow: defer round activation (the upgrade authorizer requires an inactive round) →
 * upgrade proxy to CosmicSignatureGameV3 (reinitialize, reinitializer(3) sets V3 defaults) →
 * re-emit the 6 V3 config admin events (values unchanged; exercises indexer coverage) →
 * reactivate the round.
 *
 * IMPORTANT: the current round must be finished (main prize claimed) before running this.
 * If a round is active with bids, setRoundActivationTime reverts and this script aborts —
 * claim the prize first (frontend, or `cgctl claim-prize`), then run this during the
 * between-rounds window.
 *
 * Run from the Cosmic-Signature v3.1 repo:
 *   cd /home/niko/eth/dev/b/cg-v3/Cosmic-Signature-v3.1-2026-08-19
 *   CADDR=0x... NODE_PATH=$PWD/node_modules npx hardhat run \
 *     /home/niko/eth/dev/b/backend-v3.1/cmd/cgctl/dev-scripts/upgrade-v3.js --network localhost
 *
 * Env:
 *   CADDR                 game proxy address (required)
 *   ACTIVATION_DELAY_SEC  delay before the next round activates after the upgrade (default 5)
 */
const hre = global.hre ?? require("hardhat");

const ACTIVATION_DELAY_SEC = BigInt(process.env.ACTIVATION_DELAY_SEC || "5");
const g = { gasLimit: 1000000 };

async function main() {
    const proxyAddr = process.env.CADDR;
    if (!proxyAddr || !/^0x[0-9a-fA-F]{40}$/.test(proxyAddr)) {
        console.error("Error: set CADDR to the game proxy address (printed by deploy-v2.js).");
        process.exit(1);
    }

    const [owner] = await hre.ethers.getSigners();
    const V2 = await hre.ethers.getContractFactory("CosmicSignatureGameV2", owner);
    const proxy = V2.attach(proxyAddr);

    const roundNum = await proxy.roundNum();
    const lastBidder = await proxy.lastBidderAddress();
    console.log(`Proxy ${proxyAddr}: roundNum=${roundNum}, lastBidder=${lastBidder}`);
    if (lastBidder !== hre.ethers.ZeroAddress) {
        console.error("Error: the current round has bids. Claim the main prize first, then rerun.");
        process.exit(1);
    }

    // 1) Defer activation far into the future — _authorizeUpgrade requires an inactive round,
    //    and claimMainPrize schedules the next activation at ~now+delay.
    const block = await hre.ethers.provider.getBlock("latest");
    const farFuture = BigInt(block.timestamp) + 86400n;
    try {
        await (await proxy.connect(owner).setRoundActivationTime(farFuture, g)).wait();
        console.log(`Deferred round activation to ${farFuture} for the upgrade.`);
    } catch (err) {
        console.error("setRoundActivationTime failed — is the round active? Claim the prize first.");
        console.error(err.shortMessage || err.message || err);
        process.exit(1);
    }

    // 2) Upgrade to V3. reinitializer(3) sets the 5 new V3 params to their defaults.
    //    The real V2→V3 upgrade needs no unsafe flags (append-only layout); we keep them
    //    because deploy-v2.js forced unsafeSkipStorageCheck during the V1→V2 step,
    //    which leaves the OZ manifest without a validated V2 layout (same as populate-old-v3.js).
    console.log("Upgrading proxy to CosmicSignatureGameV3 (reinitialize)...");
    const V3 = await hre.ethers.getContractFactory("CosmicSignatureGameV3", owner);
    let proxyV3;
    try {
        proxyV3 = await hre.upgrades.upgradeProxy(proxy, V3, {
            kind: "uups",
            call: "reinitialize",
            unsafeAllowRenames: true,
            unsafeSkipStorageCheck: true,
        });
    } catch (err) {
        // If the OZ manifest (.openzeppelin/unknown-*.json) was lost (e.g. fresh checkout),
        // register the existing proxy and retry once.
        console.warn("upgradeProxy failed, trying forceImport + retry:", err.shortMessage || err.message);
        await hre.upgrades.forceImport(proxyAddr, V2, { kind: "uups" });
        proxyV3 = await hre.upgrades.upgradeProxy(proxy, V3, {
            kind: "uups",
            call: "reinitialize",
            unsafeAllowRenames: true,
            unsafeSkipStorageCheck: true,
        });
    }
    await proxyV3.waitForDeployment();
    const implAddr = await hre.upgrades.erc1967.getImplementationAddress(proxyAddr);
    console.log(`V3 upgrade complete, implementation=${implAddr}`);

    // 3) Re-emit the 6 V3 config admin events with their current (default) values so the
    //    backend indexer sees ISystemEventsV3 records (record types 40–45). No behavior change.
    console.log("Emitting V3 config admin events (values unchanged)...");
    const game = proxyV3.connect(owner);
    await (await game.setRoundLateBidDurationDivisor(await proxyV3.roundLateBidDurationDivisor(), g)).wait();
    await (await game.setRoundLateBidPricePremiumAmountBaseMultiplier(await proxyV3.roundLateBidPricePremiumAmountBaseMultiplier(), g)).wait();
    await (await game.setRoundLateBidPricePremiumAmountExponent(await proxyV3.roundLateBidPricePremiumAmountExponent(), g)).wait();
    await (await game.setCstBidPriceDeclineMultiplier(await proxyV3.cstBidPriceDeclineMultiplier(), g)).wait();
    await (await game.setCstBidPriceDeclineMultiplierChangeDivisor(await proxyV3.cstBidPriceDeclineMultiplierChangeDivisor(), g)).wait();
    await (await game.setMainPrizeNumCosmicSignatureNfts(await proxyV3.mainPrizeNumCosmicSignatureNfts(), g)).wait();

    // 4) Reactivate the round.
    const nowBlock = await hre.ethers.provider.getBlock("latest");
    const activationTs = BigInt(nowBlock.timestamp) + ACTIVATION_DELAY_SEC;
    await (await game.setRoundActivationTime(activationTs, g)).wait();

    console.log("");
    console.log(`Done. roundNum=${await proxyV3.roundNum()} now runs on V3 (implementation=${implAddr}).`);
    console.log(`Round activates at ${activationTs} (in ~${ACTIVATION_DELAY_SEC}s of chain time).`);
    console.log(`mainPrizeNumCosmicSignatureNfts=${await proxyV3.mainPrizeNumCosmicSignatureNfts()}`);

    // Keep the backend's contract registry in sync with the new implementation.
    console.log("");
    console.log(`UPDATE cg_contracts SET implementation_addr='${implAddr}' WHERE cosmic_game_addr='${proxyAddr}';`);
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err);
        process.exit(1);
    });
