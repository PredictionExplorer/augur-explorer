/**
 * fork-v3-smoke.js — post-upgrade operational test on the production fork.
 * Run AFTER upgrade-fork-v3.js. Proves the upgraded contracts are not bricked:
 *
 *   1. the round activates and accepts a first ETH bid
 *   2. the late-bid premium curve moves the quoted ETH price
 *   3. a second ETH bid pays the quote and mints the bid CST reward
 *      to the PREVIOUS bidder (the V3 reward flip)
 *   4. the CST Dutch auction declines to a biddable price and takes a CST bid
 *   5. the main prize claim pays ETH, mints mainPrizeNumCosmicSignatureNfts
 *      NFTs to the beneficiary, funds the NEW PrizesWallet, and schedules
 *      the next round
 *   6. the next round activates and accepts its first bid
 *
 * Uses the fork node's unlocked genesis accounts; aborts on a non-Hardhat RPC.
 * With a non-archive upstream (cosmic1) this must run right after the node
 * starts; with an archive upstream there is no time pressure.
 *
 * Typical use (from the Cosmic-Signature repo):
 *   RPC_URL=http://127.0.0.1:10545 npx hardhat run \
 *     /home/niko/eth/dev/b/backend/cmd/cgctl/dev-scripts/fork-v3-smoke.js
 *
 * Env: RPC_URL, GAME_ADDR, TOKEN_ADDR, NFT_ADDR (defaults: production).
 */
const hre = global.hre ?? require("hardhat");

const RPC_URL = process.env.RPC_URL || "http://127.0.0.1:10545";
const GAME_ADDR = process.env.GAME_ADDR || "0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2";
const TOKEN_ADDR = process.env.TOKEN_ADDR || "0xAD91843e6A58Ba560F577E676986AFb1dba6FBA0";
const NFT_ADDR = process.env.NFT_ADDR || "0xbb84Be3500A63581d3F2d5AC3bdF8685AAedad25";

// Hardhat genesis accounts #0 and #1 (publicly known dev accounts).
const ACCT0 = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266";
const ACCT1 = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8";

const gCall = { gasLimit: 3_000_000 };

let step = 0;
function pass(msg) {
    step += 1;
    console.log(`  ok ${step}: ${msg}`);
}
function fail(msg) {
    console.error(`  FAIL at step ${step + 1}: ${msg}`);
    process.exit(1);
}
function assert(cond, okMsg, failMsg) {
    if (cond) pass(okMsg); else fail(failMsg ?? okMsg);
}

async function main() {
    const provider = new hre.ethers.JsonRpcProvider(RPC_URL);
    try {
        await provider.send("hardhat_metadata", []);
    } catch {
        console.error(`Error: ${RPC_URL} is not a Hardhat node — refusing to run (production safety).`);
        process.exit(1);
    }

    // Idempotent: upgrade-fork-v3.js already does this, but smoke can run standalone.
    // See the comment there — forked Arbitrum precompiles are 0xfe (INVALID) and
    // burn all forwarded gas; empty code mimics a plain hardhat dev chain.
    await provider.send("hardhat_setCode", ["0x0000000000000000000000000000000000000064", "0x"]); // ArbSys
    await provider.send("hardhat_setCode", ["0x000000000000000000000000000000000000006C", "0x"]); // ArbGasInfo

    const chainNow = async () => BigInt((await provider.getBlock("latest")).timestamp);
    const jumpTo = async (ts) => {
        if ((await chainNow()) >= ts) return;
        await provider.send("evm_setNextBlockTimestamp", [Number(ts)]);
        await provider.send("evm_mine", []);
    };

    const gameArt = await hre.artifacts.readArtifact("CosmicSignatureGameV3");
    const tokenArt = await hre.artifacts.readArtifact("CosmicSignatureToken");
    const nftArt = await hre.artifacts.readArtifact("CosmicSignatureNft");
    const game = new hre.ethers.Contract(GAME_ADDR, gameArt.abi, provider);
    const token = new hre.ethers.Contract(TOKEN_ADDR, tokenArt.abi, provider);
    const nft = new hre.ethers.Contract(NFT_ADDR, nftArt.abi, provider);
    const acct0 = new hre.ethers.JsonRpcSigner(provider, ACCT0);
    const acct1 = new hre.ethers.JsonRpcSigner(provider, ACCT1);

    // ---- preconditions: V3 live, round has no bids yet ----
    const nftsPerPrize = await game.mainPrizeNumCosmicSignatureNfts().catch(() => null);
    if (nftsPerPrize === null) fail("mainPrizeNumCosmicSignatureNfts() reverted — proxy is not on V3; run upgrade-fork-v3.js first");
    pass(`proxy is on V3 (mainPrizeNumCosmicSignatureNfts=${nftsPerPrize})`);
    const roundNum = await game.roundNum();
    if ((await game.lastBidderAddress()) !== hre.ethers.ZeroAddress)
        fail(`round ${roundNum} already has bids — rerun on a fresh fork`);

    // ---- 1) activation and first ETH bid ----
    await jumpTo((await game.roundActivationTime()) + 1n);
    const quote1 = await game.getNextEthBidPrice();
    await (await game.connect(acct0).bidWithEth(-1n, "rehearsal: first ETH bid", 0n, { value: quote1, ...gCall })).wait();
    assert((await game.lastBidderAddress()) === ACCT0,
        `round ${roundNum} activated; first ETH bid accepted at ${hre.ethers.formatEther(quote1)} ETH`);
    const prizeTime1 = await game.mainPrizeTime();
    assert(prizeTime1 > (await chainNow()), `mainPrizeTime scheduled (${prizeTime1})`);

    // ---- 2) the quoted ETH price moves with time (late-bid premium curve) ----
    await jumpTo((await chainNow()) + 30n);
    const quote2 = await game.getNextEthBidPrice();
    assert(quote2 >= quote1,
        `ETH quote after 30s: ${hre.ethers.formatEther(quote2)} (was ${hre.ethers.formatEther(quote1)})`,
        `quote DROPPED: ${quote2} < ${quote1}`);

    // ---- 3) second ETH bid; CST reward goes to the PREVIOUS bidder (V3 flip) ----
    const acct0CstBefore = await token.balanceOf(ACCT0);
    await (await game.connect(acct1).bidWithEth(-1n, "rehearsal: second ETH bid", 0n, { value: quote2 * 2n, ...gCall })).wait();
    assert((await game.lastBidderAddress()) === ACCT1, "second ETH bid accepted (overpayment handled)");
    const acct0CstAfter = await token.balanceOf(ACCT0);
    assert(acct0CstAfter > acct0CstBefore,
        `bid CST reward minted to previous bidder: +${hre.ethers.formatEther(acct0CstAfter - acct0CstBefore)} CST`,
        "previous bidder received no CST reward");

    // ---- 4) CST Dutch auction declines; CST bid is accepted ----
    const [untilZero] = await game.getCstDutchAuctionDurations();
    await jumpTo((await chainNow()) + untilZero + 2n);
    const cstQuote = await game.getNextCstBidPrice();
    pass(`CST auction price after ${untilZero + 2n}s: ${hre.ethers.formatEther(cstQuote)} CST`);
    await (await game.connect(acct1).bidWithCst(cstQuote + 10n ** 18n, "rehearsal: CST bid", 0n, gCall)).wait();
    assert((await game.lastBidderAddress()) === ACCT1, "CST bid accepted");

    // ---- 5) main prize claim ----
    await jumpTo((await game.mainPrizeTime()) + 1n);
    const mainPrizeEth = await game.getMainEthPrizeAmount();
    const pwAddr = await game.prizesWallet();
    const pwEthBefore = await provider.getBalance(pwAddr);
    const acct1EthBefore = await provider.getBalance(ACCT1);
    const acct1NftsBefore = await nft.balanceOf(ACCT1);
    await (await game.connect(acct1).claimMainPrize({ gasLimit: 10_000_000 })).wait();
    assert((await game.roundNum()) === roundNum + 1n, `main prize claimed; roundNum advanced to ${roundNum + 1n}`);
    const acct1EthDelta = (await provider.getBalance(ACCT1)) - acct1EthBefore;
    assert(acct1EthDelta > (mainPrizeEth * 9n) / 10n,
        `beneficiary received the main ETH prize (+${hre.ethers.formatEther(acct1EthDelta)} ETH of ${hre.ethers.formatEther(mainPrizeEth)})`,
        `beneficiary ETH delta ${acct1EthDelta} vs prize ${mainPrizeEth}`);
    const nftsMinted = (await nft.balanceOf(ACCT1)) - acct1NftsBefore;
    assert(nftsMinted >= nftsPerPrize,
        `beneficiary received ${nftsMinted} Cosmic Signature NFTs (>= ${nftsPerPrize} main prize NFTs)`,
        `only ${nftsMinted} NFTs minted to beneficiary, expected >= ${nftsPerPrize}`);
    const pwEthDelta = (await provider.getBalance(pwAddr)) - pwEthBefore;
    assert(pwEthDelta > 0n,
        `NEW PrizesWallet ${pwAddr} received raffle/chrono ETH deposits (+${hre.ethers.formatEther(pwEthDelta)} ETH)`,
        "new PrizesWallet received nothing on claim");

    // ---- 6) the next round activates and accepts its first bid ----
    await jumpTo((await game.roundActivationTime()) + 1n);
    const quote3 = await game.getNextEthBidPrice();
    await (await game.connect(acct0).bidWithEth(-1n, "rehearsal: next round first bid", 0n, { value: quote3, ...gCall })).wait();
    assert((await game.lastBidderAddress()) === ACCT0,
        `round ${roundNum + 1n} activated and accepted its first ETH bid at ${hre.ethers.formatEther(quote3)} ETH`);

    console.log("");
    console.log(`SMOKE TEST PASSED: ${step} checks — the upgraded contracts operate across a full round lifecycle.`);
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err);
        process.exit(1);
    });
