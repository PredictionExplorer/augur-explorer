/**
 * Deploy CosmicSignatureGame stack with dev-friendly short time settings,
 * then deploy two Samp ERC20 contracts. Outputs CADDR, TSAMP1, TSAMP2 for use
 * by deploy-and-populate.sh.
 *
 * Must be run from Cosmic-Signature repo with:
 *   cd /path/to/Cosmic-Signature
 *   NODE_PATH=$PWD/node_modules npx hardhat run /path/to/deploy-dev-and-samp.js --network localhost
 *
 * Uses scripts/Deploy.js from current working directory (Cosmic-Signature).
 */
const path = require("path");
const hre = require("hardhat");

// Dev-friendly time constants (short for fast tests)
const TIMEOUT_CLAIM_MAIN_PRIZE_SEC = 60;           // 1 minute to claim main prize
const MAIN_PRIZE_TIME_INCREMENT_SEC = 60;          // 1 minute per bid
const MAIN_PRIZE_TIME_INCREMENT_MICRO = BigInt(MAIN_PRIZE_TIME_INCREMENT_SEC) * 1_000_000n;
const INITIAL_DURATION_UNTIL_MAIN_PRIZE_DIVISOR = 10000000;
const CST_DUTCH_AUCTION_DURATION_DIVISOR = 3600;   // short CST auction
const TIMEOUT_WITHDRAW_PRIZES_SEC = 3600;          // 1 hour for PrizesWallet
// Short delay before first round activates so populate.js can run admin config during this window (then sleep and proceed)
const ROUND_ACTIVATION_DELAY_SEC = 5;

async function main() {
    const [owner] = await hre.ethers.getSigners();
    // Prefer the repo's own scripts/Deploy.js (V2-era layout); the V3 repos
    // use a task-based deployment instead, so fall back to the V3-compatible
    // Deploy.js that lives next to this script.
    const fs = require("fs");
    let deployJsPath = path.join(process.cwd(), "scripts", "Deploy.js");
    if (!fs.existsSync(deployJsPath)) {
        deployJsPath = path.join(__dirname, "Deploy.js");
        console.log("Repo has no scripts/Deploy.js; using " + deployJsPath);
    }
    const { basicDeployment } = require(deployJsPath);

    // Deploy with future activation (activationTime=1) so round stays inactive; we set dev params then activate
    console.log("Deploying CosmicSignatureGame stack (basicDeployment, activation in future)...");
    const result = await basicDeployment(owner, "", 1, "", false, true);
    let { cosmicGameProxy, prizesWallet } = result;
    const cosmicGameAddr = await cosmicGameProxy.getAddress();

    // Mirror the mainnet history when the repo ships V2/V3 game contracts:
    // the V1->V2 upgrade reuses V1 storage slots with renamed fields and assumes
    // a completed round zero (the repo's own test helpers do
    // deployV1 -> completeRoundZero -> upgradeToV2 -> upgradeToV3).
    let gameVersion = 1;
    try {
        await hre.ethers.getContractFactory("CosmicSignatureGameV2", owner);
        await hre.ethers.getContractFactory("CosmicSignatureGameV3", owner);
        gameVersion = 3;
    } catch (e) { /* V1-only repo; leave the proxy as deployed */ }

    if (gameVersion === 3) {
        console.log("Completing round zero on V1 (required before the V2 upgrade)...");
        const gq = { gasLimit: 12000000 };
        const block0 = await hre.ethers.provider.getBlock("latest");
        await (await cosmicGameProxy.connect(owner).setRoundActivationTime(block0.timestamp + 1, { gasLimit: 1000000 })).wait();
        await hre.network.provider.send("evm_increaseTime", [2]);
        await hre.network.provider.send("evm_mine");
        await (await cosmicGameProxy.connect(owner).bidWithEth(-1, "round zero", { value: hre.ethers.parseEther("1"), ...gq })).wait();
        const mainPrizeTime = await cosmicGameProxy.mainPrizeTime();
        await hre.network.provider.send("evm_setNextBlockTimestamp", [Number(mainPrizeTime) + 1]);
        await hre.network.provider.send("evm_mine");
        await (await cosmicGameProxy.connect(owner).claimMainPrize(gq)).wait();
        console.log("Round zero claimed (roundNum=" + (await cosmicGameProxy.roundNum()) + ").");

        // _authorizeUpgrade is _onlyRoundIsInactive and the post-claim inactive
        // window can be very short on a real-time node: freeze the round
        // inactive before upgrading (reactivated with dev settings below).
        const blockC = await hre.ethers.provider.getBlock("latest");
        await (await cosmicGameProxy.connect(owner).setRoundActivationTime(blockC.timestamp + 1000000, { gasLimit: 1000000 })).wait();

        for (const upgradeName of ["CosmicSignatureGameV2", "CosmicSignatureGameV3"]) {
            const factory = await hre.ethers.getContractFactory(upgradeName, owner);
            cosmicGameProxy = await hre.upgrades.upgradeProxy(cosmicGameProxy, factory, {
                kind: "uups",
                call: "reinitialize",
            });
            await cosmicGameProxy.waitForDeployment();
            console.log("Proxy upgraded to " + upgradeName + ".");
        }
    }

    console.log("Setting dev-friendly time parameters (round still inactive)...");
    const g = { gasLimit: 1000000 };

    // Setter names differ per contract version (e.g. setCstDutchAuctionDurationDivisor
    // is V1-only; V3 reverts the V2 CST-auction setters with NotImplemented).
    async function trySet(contract, fnName, value) {
        const fn = contract.connect(owner)[fnName];
        if (typeof fn !== "function") {
            console.log(`  ${fnName}: not present on this contract version; skipped`);
            return;
        }
        try {
            await (await fn(value, g)).wait();
            console.log(`  ${fnName}(${value})`);
        } catch (e) {
            console.warn(`  ${fnName}: reverted; skipped (${String(e.message || e).slice(0, 100)})`);
        }
    }
    await trySet(cosmicGameProxy, "setTimeoutDurationToClaimMainPrize", TIMEOUT_CLAIM_MAIN_PRIZE_SEC);
    await trySet(cosmicGameProxy, "setMainPrizeTimeIncrementInMicroSeconds", MAIN_PRIZE_TIME_INCREMENT_MICRO);
    await trySet(cosmicGameProxy, "setInitialDurationUntilMainPrizeDivisor", INITIAL_DURATION_UNTIL_MAIN_PRIZE_DIVISOR);
    await trySet(cosmicGameProxy, "setCstDutchAuctionDurationDivisor", CST_DUTCH_AUCTION_DURATION_DIVISOR);
    await (await prizesWallet.connect(owner).setTimeoutDurationToWithdrawPrizes(TIMEOUT_WITHDRAW_PRIZES_SEC, g)).wait();

    // Delay before first round activates so populate can run admin setters during this window (see claim-and-configure.sh pattern)
    await (await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(ROUND_ACTIVATION_DELAY_SEC, g)).wait();
    const block = await hre.ethers.provider.getBlock("latest");
    try {
        console.log("Setting round activation in " + ROUND_ACTIVATION_DELAY_SEC + " seconds (config window for populate)...");
        const activationTs = BigInt(block.timestamp) + BigInt(ROUND_ACTIVATION_DELAY_SEC);
        await (await cosmicGameProxy.connect(owner).setRoundActivationTime(activationTs, g)).wait();
    } catch (err) {
        console.warn("setRoundActivationTime(now+" + ROUND_ACTIVATION_DELAY_SEC + ") failed (round may already be active):", err.message || err);
        console.warn("Continuing: Samp deploy and populate will run; populate will skip admin config if round is active.");
    }

    console.log("Deploying sample ERC20 contracts...");
    // The available test-token artifact differs per repo: "Samp" (V2-era,
    // sometimes with name/symbol args) or "FuzzTestMockErc20" (V3 repos).
    async function deploySample(label) {
        const attempts = [
            ["Samp", []],
            ["Samp", ["ERC20 Token " + label, label]],
            ["FuzzTestMockErc20", []],
        ];
        for (const [name, args] of attempts) {
            try {
                const factory = await hre.ethers.getContractFactory(name);
                const t = await factory.connect(owner).deploy(...args);
                await t.waitForDeployment();
                return t;
            } catch (e) { /* try the next variant */ }
        }
        throw new Error("no deployable sample ERC20 artifact (tried Samp, FuzzTestMockErc20)");
    }
    const samp1 = await deploySample("SAMP1");
    const samp2 = await deploySample("SAMP2");
    const tsamp1 = await samp1.getAddress();
    const tsamp2 = await samp2.getAddress();

    console.log("");
    console.log("CADDR=" + cosmicGameAddr);
    console.log("TSAMP1=" + tsamp1);
    console.log("TSAMP2=" + tsamp2);
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err);
        process.exit(1);
    });
