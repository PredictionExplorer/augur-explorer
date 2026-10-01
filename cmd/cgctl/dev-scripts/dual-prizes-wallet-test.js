/**
 * dual-prizes-wallet-test.js — end-to-end scenario for DUAL PrizesWallet support.
 *
 * Reproduces the planned mainnet upgrade shape (Comment-202606264: the fixed
 * PrizesWallet is a fresh deployment, not an upgrade) and leaves unclaimed
 * prizes in BOTH wallets so the ETL and the frontend can be exercised:
 *
 *   1. Deploy the V1 stack, play bootstrap round 0 on V1, upgrade proxy to V2
 *      (same flow as deploy-v2.js; round 0 on V1 is required for V2 pricing).
 *   2. Donate 10 ETH, calibrate the next ETH bid price to ~1 ETH.
 *   3. ROUND A on V2: accounts #1/#2 bid (plus one donated RandomWalk NFT and
 *      one donated ERC-20), account #1 claims. Raffle/chrono ETH deposits and
 *      the donated NFT + ERC-20 now sit UNCLAIMED in PrizesWallet 1.
 *   4. Upgrade proxy to V3 (reinitialize + V3 config admin events), deploy
 *      PrizesWallet 2, game.setPrizesWallet(wallet2) → emits
 *      PrizesWalletAddressChanged, which is what the ETL keys on.
 *   5. ROUND B on V3: same dance. Deposits + donations sit UNCLAIMED in
 *      PrizesWallet 2. Nothing is withdrawn from either wallet.
 *   6. Verify and print per-wallet unclaimed state; leave the next round
 *      activated with a human-friendly countdown for manual play.
 *
 * Run from the Cosmic-Signature v3.1 repo (its hardhat.config.js / artifacts
 * are used), against a running `npx hardhat node`:
 *
 *   cd /home/niko/eth/dev/b/cg-v3/Cosmic-Signature-v3.1-2026-08-19
 *   NODE_PATH=$PWD/node_modules npx hardhat run \
 *     /home/niko/eth/dev/b/backend-v3.1/cmd/cgctl/dev-scripts/dual-prizes-wallet-test.js \
 *     --network localhost
 *
 * Afterwards:
 *   - insert the printed cg_contracts row (prizes_wallet_addr = WALLET 1 — the
 *     ETL discovers wallet 2 on its own from the PrizesWalletAddressChanged
 *     event and starts indexing it with no restart);
 *   - run the ETL + API + frontend;
 *   - connect hardhat account #1 (and #2) and collect: the My Allocations page
 *     must show rows from both wallets and "Retrieve All" must send one
 *     transaction per wallet with no errors.
 *
 * Tunables (env):
 *   BID_PRICE_ETH           target first-bid price per round          (default 1)
 *   DONATION_ETH            ETH donated into the game per round       (default 10)
 *   TIME_INCREMENT_SEC      seconds added to the prize timer per bid  (default 60)
 *   TIMEOUT_CLAIM_SEC       anyone-can-claim timeout after prize time (default 300)
 *   ACTIVATION_DELAY_SEC    delay before the next round activates     (default 5)
 *   INITIAL_DURATION_SEC    countdown the FIRST manual bid arms       (default 300)
 */
"use strict";

const hre = global.hre ?? require("hardhat");
const { basicDeployment } = require("./Deploy.js");

const BID_PRICE_ETH = process.env.BID_PRICE_ETH || "1";
const DONATION_ETH = process.env.DONATION_ETH || "10";
const TIME_INCREMENT_SEC = BigInt(process.env.TIME_INCREMENT_SEC || "60");
const TIMEOUT_CLAIM_SEC = BigInt(process.env.TIMEOUT_CLAIM_SEC || "300");
const ACTIVATION_DELAY_SEC = BigInt(process.env.ACTIVATION_DELAY_SEC || "5");
const INITIAL_DURATION_SEC = BigInt(process.env.INITIAL_DURATION_SEC || "300");

// Near-immediate first prize time while the script auto-plays rounds; replaced
// with a human countdown at the end (same convention as deploy-v2.js).
const INITIAL_DURATION_DIVISOR = 10000000;
const CST_DUTCH_AUCTION_DURATION_DIVISOR = 3600;
const CST_DUTCH_AUCTION_DURATION = 30 * 60; // V2-only setter
const CST_DUTCH_AUCTION_DURATION_CHANGE_DIVISOR = 250; // V2-only setter
const TIMEOUT_WITHDRAW_PRIZES_SEC = 7 * 24 * 3600; // long: nothing may expire while testing
const NUM_RAFFLE_ETH_PRIZES = 6; // 2 bidders × 6 draws → both accounts almost surely win ETH
const BID_SPACING_SEC = 30;
const ERC20_DONATION_ROUND_A = hre.ethers.parseEther("1000");
const ERC20_DONATION_ROUND_B = hre.ethers.parseEther("2000");

const g = { gasLimit: 1000000 };

// ---------------------------------------------------------------------------
// Bid-price calibration (condensed from set-bid-price.js).
// No owner setter exists for prices; on a dev chain we write the storage slot
// backing ethDutchAuctionBeginningBidPrice directly (hardhat_setStorageAt).
// The slot is found by value-scan + write-marker probe, never hardcoded.
// ---------------------------------------------------------------------------

async function findSlot(provider, gameAddr, getter) {
    const word = (v) => hre.ethers.toBeHex(v, 32);
    const readSlot = async (slot) =>
        BigInt(await provider.send("eth_getStorageAt", [gameAddr, word(slot), "latest"]));
    const writeSlot = (slot, value) =>
        provider.send("hardhat_setStorageAt", [gameAddr, word(slot), word(value)]);

    const current = await getter();
    const candidates = [];
    for (let base = 0; base < 2048; base += 256) {
        const reads = [];
        for (let i = base; i < base + 256; i++) reads.push(readSlot(i));
        (await Promise.all(reads)).forEach((v, j) => {
            if (v === current) candidates.push(base + j);
        });
    }
    const marker = 0xc0ffee0000000000000000000000000000000000000000000000000000n;
    for (const slot of candidates) {
        const original = await readSlot(slot);
        await writeSlot(slot, marker + BigInt(slot));
        const got = await getter();
        await writeSlot(slot, original);
        if (got === marker + BigInt(slot)) return { slot, writeSlot };
    }
    throw new Error(`no storage slot responds to a write probe (candidates: [${candidates}])`);
}

/** Set the next ETH bid price to ~target while the round has no bids yet. */
async function calibrateBidPrice(game, gameAddr, targetEth) {
    const provider = hre.ethers.provider;
    const target = hre.ethers.parseEther(targetEth);
    const { slot, writeSlot } = await findSlot(provider, gameAddr, () =>
        game.ethDutchAuctionBeginningBidPrice(),
    );
    await writeSlot(slot, target);
    // The Dutch auction declines from the beginning price; one proportional
    // correction lands the EFFECTIVE price on the target (within rounding).
    const effective = await game.getNextEthBidPrice();
    if (effective !== target && effective > 0n) {
        await writeSlot(slot, (target * target) / effective);
    }
    console.log(
        `  next ETH bid price calibrated to ${hre.ethers.formatEther(await game.getNextEthBidPrice())} ETH (slot ${slot})`,
    );
}

// ---------------------------------------------------------------------------
// Scenario helpers
// ---------------------------------------------------------------------------

async function increaseTime(sec) {
    await hre.ethers.provider.send("evm_increaseTime", [sec]);
    await hre.ethers.provider.send("evm_mine");
}

async function setActivationTimeToNow(game, owner) {
    const blk = await hre.ethers.provider.getBlock("latest");
    await (await game.connect(owner).setRoundActivationTime(BigInt(blk.timestamp) - 1n, g)).wait();
}

async function deferActivation(game, owner) {
    const blk = await hre.ethers.provider.getBlock("latest");
    await (await game.connect(owner).setRoundActivationTime(BigInt(blk.timestamp) + 86400n, g)).wait();
}

/**
 * V2/V3 share the bid signature: bidWithEth(rwalkNftId, msg, cstRewardMinLimit).
 * The value sent is 1.5× the quote: on V3 the late-bid price premium makes the
 * price grow every second, so the exact quoted amount can be insufficient by
 * the time the transaction mines. The game refunds the excess (Comment on
 * ethBidRefund), so overpaying is safe.
 */
function withBuffer(price) {
    return price + price / 2n;
}

async function bidEth(game, signer, message) {
    const price = await game.getNextEthBidPrice();
    await (await game.connect(signer).bidWithEth(-1n, message, 0n, { value: withBuffer(price), gasLimit: 3000000 })).wait();
    return price;
}

async function advanceToMainPrizeAndClaim(game, claimer) {
    // V3 returns int256 (can go negative past prize time); V2 clamps at 0.
    const duration = await game.getDurationUntilMainPrize();
    if (duration > 0n) await increaseTime(Number(duration) + 1);
    const tx = await game.connect(claimer).claimMainPrize({ gasLimit: 30000000 });
    return tx.wait();
}

async function mintRandomWalkNft(randomWalkNFT, signer) {
    const price = await randomWalkNFT.getMintPrice();
    const receipt = await (await randomWalkNFT.connect(signer).mint({ value: price })).wait();
    const topic = randomWalkNFT.interface.getEvent("MintEvent").topicHash;
    const log = receipt.logs.find((x) => x.topics.indexOf(topic) >= 0);
    return randomWalkNFT.interface.parseLog(log).args[0];
}

/**
 * Dev builds compiled with ENABLE_ASSERTS guard registerRoundEnd with
 * `assert(... mainPrizeBeneficiaryAddresses[roundNum_ - 1] != 0)`, which a
 * freshly deployed second wallet violates by design (round A was registered in
 * wallet 1). Production builds compile the asserts out; on a dev build the
 * contract exposes setBypassSomeAsserts(true) exactly for this.
 */
async function bypassAssertsIfDevBuild(wallet, owner, label) {
    try {
        await wallet.bypassSomeAsserts(); // getter only exists on assert-enabled builds
    } catch (_) {
        return; // production build — nothing to do
    }
    await (await wallet.connect(owner).setBypassSomeAsserts(true, g)).wait();
    console.log(`  ${label}: dev build with asserts detected — setBypassSomeAsserts(true)`);
}

/**
 * One full auto-played round: donate ETH into the pool, calibrate the bid
 * price, four spaced bids by accounts #1/#2 (one donating a RandomWalk NFT,
 * one donating ERC-20), then account #1 (the last bidder) claims. Nothing is
 * withdrawn afterwards — all secondary prizes stay in `wallet`.
 */
async function playRound(label, ctx, wallet) {
    const { game, gameAddr, owner, addr1, addr2, randomWalkNFT, samp } = ctx;
    console.log(`\n=== ${label} ===`);

    await (await game.connect(owner).donateEth({ value: hre.ethers.parseEther(DONATION_ETH), gasLimit: 1000000 })).wait();
    console.log(`  donated ${DONATION_ETH} ETH into the game pool`);

    await setActivationTimeToNow(game, owner);
    await calibrateBidPrice(game, gameAddr, BID_PRICE_ETH);

    const walletAddr = await wallet.getAddress();

    // Bid 1 (account #2) and bid 2 (account #1), plain.
    console.log(`  bid 1: account #2 (${hre.ethers.formatEther(await bidEth(game, addr2, `${label} bid 1 (acct2)`))} ETH)`);
    await increaseTime(BID_SPACING_SEC);
    console.log(`  bid 2: account #1 (${hre.ethers.formatEther(await bidEth(game, addr1, `${label} bid 2 (acct1)`))} ETH)`);

    // Bid 3 (account #2) donates a freshly minted RandomWalk NFT.
    await increaseTime(BID_SPACING_SEC);
    const nftId = await mintRandomWalkNft(randomWalkNFT, addr2);
    await (await randomWalkNFT.connect(addr2).setApprovalForAll(walletAddr, true)).wait();
    {
        const price = await game.getNextEthBidPrice();
        await (await game.connect(addr2).bidWithEthAndDonateNft(
            -1n, `${label} bid 3 + NFT donation`, 0n,
            await randomWalkNFT.getAddress(), nftId,
            { value: withBuffer(price), gasLimit: 3000000 },
        )).wait();
        console.log(`  bid 3: account #2 + donated RandomWalk NFT #${nftId}`);
    }

    // Bid 4 (account #1, last bidder → main prize beneficiary) donates ERC-20.
    await increaseTime(BID_SPACING_SEC);
    const erc20Amount = label.includes("V3") ? ERC20_DONATION_ROUND_B : ERC20_DONATION_ROUND_A;
    await (await samp.connect(addr1).approve(gameAddr, hre.ethers.MaxUint256)).wait();
    await (await samp.connect(addr1).approve(walletAddr, hre.ethers.MaxUint256)).wait();
    {
        const price = await game.getNextEthBidPrice();
        await (await game.connect(addr1).bidWithEthAndDonateToken(
            -1n, `${label} bid 4 + ERC20 donation`, 0n,
            await samp.getAddress(), erc20Amount,
            { value: withBuffer(price), gasLimit: 3000000 },
        )).wait();
        console.log(`  bid 4: account #1 + donated ${hre.ethers.formatEther(erc20Amount)} SAMP`);
    }

    const roundNum = await game.roundNum();
    await advanceToMainPrizeAndClaim(game, addr1);
    console.log(`  round ${roundNum} claimed by account #1 — secondary prizes deposited into ${walletAddr}`);
    return Number(roundNum);
}

/** Collect per-(round, winner) unclaimed ETH held by one wallet. */
async function unclaimedEthReport(wallet) {
    const events = await wallet.queryFilter(wallet.filters.EthReceived());
    const seen = new Set();
    const rows = [];
    for (const evt of events) {
        const round = Number(evt.args.roundNum);
        const winner = evt.args.prizeWinnerAddress;
        const key = `${round}|${winner}`;
        if (seen.has(key)) continue;
        seen.add(key);
        const amount = await wallet["getEthBalanceAmount(uint256,address)"](round, winner);
        if (amount > 0n) rows.push({ round, winner, amount });
    }
    return rows;
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

async function main() {
    const [owner, addr1, addr2] = await hre.ethers.getSigners();

    // ---- 1) V1 stack + bootstrap round 0 + upgrade to V2 (deploy-v2.js flow) ----
    console.log("Deploying CosmicSignatureGame stack (V1 proxy, round inactive)...");
    const deployed = await basicDeployment(owner, "", 1, "", false, true);
    const { cosmicGameProxy, prizesWallet: wallet1, randomWalkNFT } = deployed;
    const gameAddr = await cosmicGameProxy.getAddress();
    const wallet1Addr = await wallet1.getAddress();
    console.log(`Game proxy:     ${gameAddr}`);
    console.log(`PrizesWallet 1: ${wallet1Addr}`);

    console.log("Setting dev-friendly time parameters...");
    await (await cosmicGameProxy.connect(owner).setTimeoutDurationToClaimMainPrize(TIMEOUT_CLAIM_SEC, g)).wait();
    await (await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementInMicroSeconds(TIME_INCREMENT_SEC * 1_000_000n, g)).wait();
    await (await cosmicGameProxy.connect(owner).setInitialDurationUntilMainPrizeDivisor(INITIAL_DURATION_DIVISOR, g)).wait();
    await (await cosmicGameProxy.connect(owner).setCstDutchAuctionDurationDivisor(CST_DUTCH_AUCTION_DURATION_DIVISOR, g)).wait();
    await (await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(ACTIVATION_DELAY_SEC, g)).wait();
    await (await cosmicGameProxy.connect(owner).setNumRaffleEthPrizesForBidders(NUM_RAFFLE_ETH_PRIZES, g)).wait();
    await (await wallet1.connect(owner).setTimeoutDurationToWithdrawPrizes(TIMEOUT_WITHDRAW_PRIZES_SEC, g)).wait();
    // Make the wallet-1 address part of the on-chain admin-event history, the
    // same stream (PrizesWalletAddressChanged) the ETL derives the wallet
    // registry from. The wallet-2 switch below adds the second entry.
    await (await cosmicGameProxy.connect(owner).setPrizesWallet(wallet1Addr, g)).wait();
    await bypassAssertsIfDevBuild(wallet1, owner, "wallet 1");

    console.log("Playing round 0 on V1 (bootstrap bid + claim, required for V2 pricing)...");
    {
        const blk = await hre.ethers.provider.getBlock("latest");
        await (await cosmicGameProxy.connect(owner).setRoundActivationTime(BigInt(blk.timestamp), g)).wait();
        const price = await cosmicGameProxy.getNextEthBidPrice();
        await (await cosmicGameProxy.connect(owner).bidWithEth(-1n, "round 0 bootstrap (V1)", { value: price, gasLimit: 3000000 })).wait();
        const duration = await cosmicGameProxy.getDurationUntilMainPrize();
        if (duration > 0n) await increaseTime(Number(duration) + 1);
        await (await cosmicGameProxy.connect(owner).claimMainPrize({ gasLimit: 9000000 })).wait();
        console.log(`Round 0 claimed (roundNum=${await cosmicGameProxy.roundNum()}).`);
    }
    await deferActivation(cosmicGameProxy, owner);

    console.log("Upgrading proxy to CosmicSignatureGameV2 (reinitialize)...");
    const V2 = await hre.ethers.getContractFactory("CosmicSignatureGameV2", owner);
    let game = await hre.upgrades.upgradeProxy(cosmicGameProxy, V2, {
        kind: "uups",
        call: "reinitialize",
        unsafeAllowRenames: true,
        unsafeSkipStorageCheck: true,
    });
    await game.waitForDeployment();
    console.log(`V2 implementation: ${await hre.upgrades.erc1967.getImplementationAddress(gameAddr)}`);

    // V2-only CST setters (revert with NotImplemented on V3) + re-apply the dev
    // claim timeout that V2's reinitialize() reset to the production default.
    await (await game.connect(owner).setCstDutchAuctionDuration(CST_DUTCH_AUCTION_DURATION, g)).wait();
    await (await game.connect(owner).setCstDutchAuctionDurationChangeDivisor(CST_DUTCH_AUCTION_DURATION_CHANGE_DIVISOR, g)).wait();
    await (await game.connect(owner).setTimeoutDurationToClaimMainPrize(TIMEOUT_CLAIM_SEC, g)).wait();

    // ---- 2) Sample ERC-20 for donations ----
    let sampFactory;
    try {
        sampFactory = await hre.ethers.getContractFactory("Samp");
    } catch (_) {
        sampFactory = await hre.ethers.getContractFactory("FuzzTestMockErc20");
    }
    const sampArgs = sampFactory.interface.deploy.inputs.length === 2 ? ["Sample Token", "SAMP"] : [];
    const samp = await sampFactory.connect(owner).deploy(...sampArgs);
    await samp.waitForDeployment();
    if (typeof samp.mint === "function") {
        try { await (await samp.connect(owner).mint(owner.address, hre.ethers.parseEther("1000000"))).wait(); } catch (_) {}
    }
    await (await samp.connect(owner).transfer(addr1.address, hre.ethers.parseEther("100000"))).wait();
    console.log(`Sample ERC20: ${await samp.getAddress()} (100000 SAMP → account #1)`);

    const ctx = { game, gameAddr, owner, addr1, addr2, randomWalkNFT, samp };

    // ---- 3) ROUND A on V2 → prizes accumulate in PrizesWallet 1 ----
    const roundA = await playRound("ROUND A (V2, PrizesWallet 1)", ctx, wallet1);

    // ---- 4) Upgrade to V3 + deploy and switch to PrizesWallet 2 ----
    await deferActivation(game, owner);
    console.log("\nUpgrading proxy to CosmicSignatureGameV3 (reinitialize)...");
    const V3 = await hre.ethers.getContractFactory("CosmicSignatureGameV3", owner);
    game = await hre.upgrades.upgradeProxy(game, V3, {
        kind: "uups",
        call: "reinitialize",
        unsafeAllowRenames: true,
        unsafeSkipStorageCheck: true,
    });
    await game.waitForDeployment();
    const implV3 = await hre.upgrades.erc1967.getImplementationAddress(gameAddr);
    console.log(`V3 implementation: ${implV3}`);

    // V3 config admin events (record types 40–45) with current values, so the
    // ETL sees the ISystemEventsV3 stream (same as upgrade-v3.js).
    await (await game.connect(owner).setRoundLateBidDurationDivisor(await game.roundLateBidDurationDivisor(), g)).wait();
    await (await game.connect(owner).setRoundLateBidPricePremiumAmountBaseMultiplier(await game.roundLateBidPricePremiumAmountBaseMultiplier(), g)).wait();
    await (await game.connect(owner).setRoundLateBidPricePremiumAmountExponent(await game.roundLateBidPricePremiumAmountExponent(), g)).wait();
    await (await game.connect(owner).setCstBidPriceDeclineMultiplier(await game.cstBidPriceDeclineMultiplier(), g)).wait();
    await (await game.connect(owner).setCstBidPriceDeclineMultiplierChangeDivisor(await game.cstBidPriceDeclineMultiplierChangeDivisor(), g)).wait();
    await (await game.connect(owner).setMainPrizeNumCosmicSignatureNfts(await game.mainPrizeNumCosmicSignatureNfts(), g)).wait();

    console.log("Deploying PrizesWallet 2 and pointing the game at it (setPrizesWallet)...");
    const PW = await hre.ethers.getContractFactory("PrizesWallet", owner);
    const wallet2 = await PW.deploy(gameAddr);
    await wallet2.waitForDeployment();
    const wallet2Addr = await wallet2.getAddress();
    await (await wallet2.connect(owner).setTimeoutDurationToWithdrawPrizes(TIMEOUT_WITHDRAW_PRIZES_SEC, g)).wait();
    await bypassAssertsIfDevBuild(wallet2, owner, "wallet 2");
    // The round is inactive (deferred above), as setPrizesWallet requires.
    await (await game.connect(owner).setPrizesWallet(wallet2Addr, g)).wait();
    console.log(`PrizesWallet 2: ${wallet2Addr} (PrizesWalletAddressChanged emitted)`);

    ctx.game = game;

    // ---- 5) ROUND B on V3 → prizes accumulate in PrizesWallet 2 ----
    const roundB = await playRound("ROUND B (V3, PrizesWallet 2)", ctx, wallet2);

    // ---- 6) Verification ----
    console.log("\n=== VERIFICATION ===");
    const fmt = hre.ethers.formatEther;
    let failures = 0;
    const check = (cond, msg) => {
        console.log(`  ${cond ? "OK " : "FAIL"} ${msg}`);
        if (!cond) failures++;
    };

    check((await game.prizesWallet()) === wallet2Addr, "game.prizesWallet() is PrizesWallet 2");

    const [bal1, bal2] = [
        await hre.ethers.provider.getBalance(wallet1Addr),
        await hre.ethers.provider.getBalance(wallet2Addr),
    ];
    check(bal1 > 0n, `PrizesWallet 1 holds ETH: ${fmt(bal1)}`);
    check(bal2 > 0n, `PrizesWallet 2 holds ETH: ${fmt(bal2)}`);

    // Wallet 1 holds rounds ≤ A (the bootstrap round 0 also deposited a dust
    // prize for the owner there); wallet 2 must hold ONLY round B — any earlier
    // round appearing in wallet 2 would mean deposits leaked across the switch.
    const report1 = await unclaimedEthReport(wallet1);
    const report2 = await unclaimedEthReport(wallet2);
    check(report1.some((r) => r.round === roundA) && report1.every((r) => r.round <= roundA),
        `wallet 1 unclaimed ETH covers round ${roundA}, nothing later (${report1.length} rows)`);
    check(report2.length > 0 && report2.every((r) => r.round === roundB),
        `wallet 2 unclaimed ETH rows are all from round ${roundB} (${report2.length} rows)`);

    const who = (a) =>
        a === addr1.address ? "account #1" : a === addr2.address ? "account #2" : a;
    for (const [name, rows] of [["wallet 1", report1], ["wallet 2", report2]]) {
        for (const r of rows) {
            console.log(`       ${name} round ${r.round}: ${fmt(r.amount)} ETH → ${who(r.winner)} (${r.winner})`);
        }
    }
    for (const [idx, acct] of [[1, addr1], [2, addr2]]) {
        for (const [name, rows, round] of [["wallet 1", report1, roundA], ["wallet 2", report2, roundB]]) {
            if (!rows.some((r) => r.winner === acct.address)) {
                console.log(`  NOTE account #${idx} won no raffle ETH in ${name} round ${round} (random draw) — use the other account for that wallet`);
            }
        }
    }

    const sampAddr = await samp.getAddress();
    const [erc20W1, erc20W2] = [
        await wallet1.getDonatedTokenBalanceAmount(roundA, sampAddr),
        await wallet2.getDonatedTokenBalanceAmount(roundB, sampAddr),
    ];
    check(erc20W1 === ERC20_DONATION_ROUND_A, `wallet 1 holds donated ERC20: ${fmt(erc20W1)} SAMP (round ${roundA}, claimable by account #1)`);
    check(erc20W2 === ERC20_DONATION_ROUND_B, `wallet 2 holds donated ERC20: ${fmt(erc20W2)} SAMP (round ${roundB}, claimable by account #1)`);

    const [nftsW1, nftsW2] = [
        Number(await wallet1.nextDonatedNftIndex()),
        Number(await wallet2.nextDonatedNftIndex()),
    ];
    check(nftsW1 === 1, `wallet 1 holds 1 donated NFT (index 0, claimable by account #1)`);
    check(nftsW2 === 1, `wallet 2 holds 1 donated NFT (index 0, claimable by account #1)`);

    check(
        (await wallet1.mainPrizeBeneficiaryAddresses(roundA)) === addr1.address &&
        (await wallet2.mainPrizeBeneficiaryAddresses(roundB)) === addr1.address,
        "account #1 is the main prize beneficiary of both rounds (may claim both donations)",
    );

    // ---- 7) Leave the chain playable: next round with a human countdown ----
    const incrementMicroSec = await game.mainPrizeTimeIncrementInMicroSeconds();
    await (await game.connect(owner).setInitialDurationUntilMainPrizeDivisor(incrementMicroSec / INITIAL_DURATION_SEC, g)).wait();
    const blk = await hre.ethers.provider.getBlock("latest");
    await (await game.connect(owner).setRoundActivationTime(BigInt(blk.timestamp) + ACTIVATION_DELAY_SEC, g)).wait();
    console.log(`\nNext round (${await game.roundNum()}) activates in ~${ACTIVATION_DELAY_SEC}s; first bid arms a ~${INITIAL_DURATION_SEC}s countdown.`);

    // ---- 8) Backend bootstrap ----
    // prizes_wallet_addr is WALLET 1 on purpose: the ETL discovers wallet 2 by
    // itself from the PrizesWalletAddressChanged event and adds it to the watch
    // set at runtime — that is the code path this scenario exists to test.
    console.log("\nContract Addresses Deployed:");
    console.log(
        "INSERT INTO cg_contracts(" +
        "cosmic_game_addr,cosmic_signature_addr,cosmic_token_addr,cosmic_dao_addr," +
        "charity_wallet_addr,prizes_wallet_addr,random_walk_addr," +
        "staking_wallet_cst_addr,staking_wallet_rwalk_addr,marketing_wallet_addr," +
        "implementation_addr" +
        ") VALUES(" +
        `'${gameAddr}',` +
        `'${await deployed.cosmicSignature.getAddress()}',` +
        `'${await deployed.cosmicToken.getAddress()}',` +
        `'${await deployed.cosmicDAO.getAddress()}',` +
        `'${await deployed.charityWallet.getAddress()}',` +
        `'${wallet1Addr}',` +
        `'${await randomWalkNFT.getAddress()}',` +
        `'${await deployed.stakingWalletCosmicSignatureNft.getAddress()}',` +
        `'${await deployed.stakingWalletRandomWalkNft.getAddress()}',` +
        `'${await deployed.marketingWallet.getAddress()}',` +
        `'${implV3}')`,
    );
    console.log("-- optional, mirrors what you would do on mainnet after the switch:");
    console.log(`-- UPDATE cg_contracts SET prizes_wallet_addr='${wallet2Addr}' WHERE cosmic_game_addr='${gameAddr}';`);

    console.log("");
    console.log("CADDR=" + gameAddr);
    console.log("PWALLET1=" + wallet1Addr);
    console.log("PWALLET2=" + wallet2Addr);
    console.log("TSAMP1=" + sampAddr);

    console.log("\nWhat to verify next:");
    console.log("  ETL:      cg_adm_prizes_wallet_addr has 2 rows (wallet 1, wallet 2); deposits/donations");
    console.log(`            of round ${roundA} carry wallet 1's contract_aid, round ${roundB} wallet 2's.`);
    console.log("  API:      /api/cosmicgame/statistics/dashboard ContractAddrs.PrizesWalletAddrs lists both;");
    console.log("            unclaimed prize rows carry WalletAddr.");
    console.log("  Frontend: connect hardhat account #1 — My Allocations shows ETH + NFT + ERC20 rows from");
    console.log("            BOTH wallets; 'Retrieve All' sends one tx per wallet; all claims succeed.");

    if (failures > 0) throw new Error(`${failures} verification check(s) failed`);
    console.log("\nAll verification checks passed.");
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err);
        process.exit(1);
    });
