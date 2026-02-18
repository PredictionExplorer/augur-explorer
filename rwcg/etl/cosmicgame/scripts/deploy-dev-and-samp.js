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

async function main() {
    const [owner] = await hre.ethers.getSigners();
    const deployJsPath = path.join(process.cwd(), "scripts", "Deploy.js");
    const { basicDeployment } = require(deployJsPath);

    // Deploy with future activation (activationTime=1) so round stays inactive; we set dev params then activate
    console.log("Deploying CosmicSignatureGame stack (basicDeployment, activation in future)...");
    const result = await basicDeployment(owner, "", 1, "", false, true);
    const { cosmicGameProxy, prizesWallet } = result;
    const cosmicGameAddr = await cosmicGameProxy.getAddress();

    console.log("Setting dev-friendly time parameters (round still inactive)...");
    const g = { gasLimit: 1000000 };

    await (await cosmicGameProxy.connect(owner).setTimeoutDurationToClaimMainPrize(TIMEOUT_CLAIM_MAIN_PRIZE_SEC, g)).wait();
    await (await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementInMicroSeconds(MAIN_PRIZE_TIME_INCREMENT_MICRO, g)).wait();
    await (await cosmicGameProxy.connect(owner).setInitialDurationUntilMainPrizeDivisor(INITIAL_DURATION_UNTIL_MAIN_PRIZE_DIVISOR, g)).wait();
    await (await cosmicGameProxy.connect(owner).setCstDutchAuctionDurationDivisor(CST_DUTCH_AUCTION_DURATION_DIVISOR, g)).wait();
    await (await prizesWallet.connect(owner).setTimeoutDurationToWithdrawPrizes(TIMEOUT_WITHDRAW_PRIZES_SEC, g)).wait();

    console.log("Activating round (set round activation time to now)...");
    const block = await hre.ethers.provider.getBlock("latest");
    await (await cosmicGameProxy.connect(owner).setRoundActivationTime(block.timestamp - 1, g)).wait();

    console.log("Deploying Samp ERC20 contracts...");
    const Samp = await hre.ethers.getContractFactory("Samp");
    const samp1 = await Samp.connect(owner).deploy("ERC20 Token Sample1", "SAMP1");
    await samp1.waitForDeployment();
    const samp2 = await Samp.connect(owner).deploy("ERC20 Token Sample2", "SAMP2");
    await samp2.waitForDeployment();
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
