/**
 * deploy-v2.js — deploy the CosmicSignatureGame stack and upgrade the proxy to V2,
 * leaving the chain ready for manual play (bids, claims) on the V2 implementation.
 * Upgrade to V3 later with upgrade-v3.js.
 *
 * Flow: deploy V1 stack (Deploy.js, round inactive) → set dev-friendly time params →
 * upgrade proxy to CosmicSignatureGameV2 (reinitialize) → set V2-only CST auction params →
 * deploy 2 Samp ERC20s → activate round 0.
 *
 * Run from the Cosmic-Signature v3.1 repo (its hardhat.config.js / artifacts are used):
 *   cd /home/niko/eth/dev/b/cg-v3/Cosmic-Signature-v3.1-2026-08-19
 *   NODE_PATH=$PWD/node_modules npx hardhat run \
 *     /home/niko/eth/dev/b/backend-v3.1/cmd/cgctl/dev-scripts/deploy-v2.js --network localhost
 *
 * Tunables (env):
 *   TIME_INCREMENT_SEC      seconds added to the prize timer per bid   (default 60)
 *   TIMEOUT_CLAIM_SEC       anyone-can-claim timeout after prize time  (default 300)
 *   ACTIVATION_DELAY_SEC    delay before round 0 (and each next round) activates (default 5)
 *
 * Prints CADDR= / TSAMP1= / TSAMP2= lines (same convention as deploy-and-populate.sh).
 */
const hre = require("hardhat");
const { basicDeployment } = require("./Deploy.js");

const TIME_INCREMENT_SEC = BigInt(process.env.TIME_INCREMENT_SEC || "60");
const TIMEOUT_CLAIM_SEC = BigInt(process.env.TIMEOUT_CLAIM_SEC || "300");
const ACTIVATION_DELAY_SEC = BigInt(process.env.ACTIVATION_DELAY_SEC || "5");

// Same dev values the populate/deploy scripts use.
const INITIAL_DURATION_DIVISOR = 10000000; // near-immediate first prize time
const CST_DUTCH_AUCTION_DURATION_DIVISOR = 3600;
const CST_DUTCH_AUCTION_DURATION = 30 * 60; // V2-only setter
const CST_DUTCH_AUCTION_DURATION_CHANGE_DIVISOR = 250; // V2-only setter
const TIMEOUT_WITHDRAW_PRIZES_SEC = 3600;

const g = { gasLimit: 1000000 };

async function main() {
    const [owner] = await hre.ethers.getSigners();

    // 1) Deploy the V1 stack with activation in the future (activationTime=1 → now+3600),
    //    so all admin setters run while the round is inactive.
    console.log("Deploying CosmicSignatureGame stack (V1 proxy, round inactive)...");
    const result = await basicDeployment(owner, "", 1, "", false, true);
    const { cosmicGameProxy, prizesWallet } = result;
    const proxyAddr = await cosmicGameProxy.getAddress();
    console.log(`Proxy deployed at ${proxyAddr}`);

    // 2) Dev-friendly time parameters (settable on V1; carried through upgrades).
    console.log("Setting dev-friendly time parameters...");
    await (await cosmicGameProxy.connect(owner).setTimeoutDurationToClaimMainPrize(TIMEOUT_CLAIM_SEC, g)).wait();
    await (await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementInMicroSeconds(TIME_INCREMENT_SEC * 1_000_000n, g)).wait();
    await (await cosmicGameProxy.connect(owner).setInitialDurationUntilMainPrizeDivisor(INITIAL_DURATION_DIVISOR, g)).wait();
    await (await cosmicGameProxy.connect(owner).setCstDutchAuctionDurationDivisor(CST_DUTCH_AUCTION_DURATION_DIVISOR, g)).wait();
    await (await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(ACTIVATION_DELAY_SEC, g)).wait();
    await (await prizesWallet.connect(owner).setTimeoutDurationToWithdrawPrizes(TIMEOUT_WITHDRAW_PRIZES_SEC, g)).wait();

    // 3) Upgrade the proxy to V2 while the round is still inactive.
    //    Same options as populate-old-v3.js: the v3 branch renamed initializeV2() → reinitialize()
    //    (reinitializer(2)); V2 reuses V1 slots with renamed fields, hence the unsafe flags.
    console.log("Upgrading proxy to CosmicSignatureGameV2 (reinitialize)...");
    const V2 = await hre.ethers.getContractFactory("CosmicSignatureGameV2", owner);
    const proxyV2 = await hre.upgrades.upgradeProxy(cosmicGameProxy, V2, {
        kind: "uups",
        call: "reinitialize",
        unsafeAllowRenames: true,
        unsafeSkipStorageCheck: true,
    });
    await proxyV2.waitForDeployment();
    const implAddr = await hre.upgrades.erc1967.getImplementationAddress(proxyAddr);
    console.log(`V2 upgrade complete, implementation=${implAddr}`);

    // 4) V2-only CST auction params (these setters revert with NotImplemented on V3,
    //    so they must be emitted now — matches populate-old-v3.js ordering).
    await (await proxyV2.connect(owner).setCstDutchAuctionDuration(CST_DUTCH_AUCTION_DURATION, g)).wait();
    await (await proxyV2.connect(owner).setCstDutchAuctionDurationChangeDivisor(CST_DUTCH_AUCTION_DURATION_CHANGE_DIVISOR, g)).wait();

    // 5) Sample ERC20s (attachable to bids / donations from the frontend).
    //    Samp exists on older branches; the v3 branch ships FuzzTestMockErc20 instead
    //    (same fallback as populate-old-v3.js).
    console.log("Deploying sample ERC20 contracts...");
    let sampFactory, sampName;
    try {
        sampFactory = await hre.ethers.getContractFactory("Samp");
        sampName = "Samp";
    } catch (_) {
        sampFactory = await hre.ethers.getContractFactory("FuzzTestMockErc20");
        sampName = "FuzzTestMockErc20";
    }
    const sampArgs = sampName === "Samp" ? [["ERC20 Token Sample1", "SAMP1"], ["ERC20 Token Sample2", "SAMP2"]] : [[], []];
    const samp1 = await sampFactory.connect(owner).deploy(...sampArgs[0]);
    await samp1.waitForDeployment();
    const samp2 = await sampFactory.connect(owner).deploy(...sampArgs[1]);
    await samp2.waitForDeployment();
    if (sampName === "FuzzTestMockErc20") {
        const mintAmount = hre.ethers.parseEther("100000000000");
        await samp1.connect(owner).mint(owner.address, mintAmount);
        await samp2.connect(owner).mint(owner.address, mintAmount);
    }
    console.log(`Deployed sample ERC20 tokens (${sampName})`);

    // 6) Activate round 0 shortly.
    const block = await hre.ethers.provider.getBlock("latest");
    const activationTs = BigInt(block.timestamp) + ACTIVATION_DELAY_SEC;
    await (await proxyV2.connect(owner).setRoundActivationTime(activationTs, g)).wait();
    console.log(`Round 0 activates at ${activationTs} (in ~${ACTIVATION_DELAY_SEC}s of chain time).`);

    const roundNum = await proxyV2.roundNum();
    console.log("");
    console.log(`Ready to play on V2 (roundNum=${roundNum}).`);
    console.log(`Upgrade later with: CADDR=${proxyAddr} ... upgrade-v3.js (see that script's header).`);

    // Contract registry row for the backend (same format as populate-old-v3.js;
    // column order matches the cg_contracts table definition).
    console.log("");
    console.log("Contract Addresses Deployed:");
    console.log(
        "INSERT INTO cg_contracts VALUES(" +
        `'${proxyAddr}',` +
        `'${await result.cosmicSignature.getAddress()}',` +
        `'${await result.cosmicToken.getAddress()}',` +
        `'${await result.cosmicDAO.getAddress()}',` +
        `'${await result.charityWallet.getAddress()}',` +
        `'${await prizesWallet.getAddress()}',` +
        `'${await result.randomWalkNFT.getAddress()}',` +
        `'${await result.stakingWalletCosmicSignatureNft.getAddress()}',` +
        `'${await result.stakingWalletRandomWalkNft.getAddress()}',` +
        `'${await result.marketingWallet.getAddress()}',` +
        `'${implAddr}')`
    );

    console.log("");
    console.log("CADDR=" + proxyAddr);
    console.log("TSAMP1=" + (await samp1.getAddress()));
    console.log("TSAMP2=" + (await samp2.getAddress()));
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err);
        process.exit(1);
    });
