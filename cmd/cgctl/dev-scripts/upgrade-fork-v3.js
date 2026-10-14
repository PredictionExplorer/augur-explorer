/**
 * upgrade-fork-v3.js — rehearse the production V2->V3 upgrade on a local Hardhat FORK.
 *
 * Fork-only companion to upgrade-v3.js: instead of a signer key, it impersonates the
 * real owner via hardhat_impersonateAccount, so it CANNOT touch the real chain (the
 * impersonation RPC only exists on a Hardhat node; the script aborts anywhere else).
 *
 * It replays the production upgrade-day sequence with the same calldata the owner
 * will sign for real:
 *   1. defer round activation      (the upgrade authorizer requires an inactive round)
 *   2. deploy CosmicSignatureGameV3 and upgradeToAndCall(impl, reinitialize())
 *   3. deploy a new PrizesWallet(game) and setPrizesWallet(new)     [SKIP_PRIZES_WALLET=1]
 *   4. re-emit the 4 V3 config admin events (values unchanged) — reinitialize() only
 *      emits Initialized(3) (Comment-202610175), so the ETL needs these to populate
 *      the V3 config admin tables
 *   5. reactivate the round and verify every value reinitialize() wrote
 *
 * It bypasses hre's network config entirely (its own JsonRpcProvider from RPC_URL),
 * so the repo's hardhat.config.js needs no fork_local entry. Hardhat is only used
 * to compile and read artifacts. NOTE: this intentionally skips OpenZeppelin layout
 * validation (no .openzeppelin manifest exists for a fork); the committed
 * CosmicSignatureGameV3-DeployedV2StorageLayout.js test covers that separately.
 *
 * Typical use (after fork-and-prepare.sh + fork-v3-rehearsal.sh claim):
 *   cd /home/niko/eth/dev/b/cg-v3/Cosmic-Signature-v3.1-2026-08-19-review
 *   RPC_URL=http://127.0.0.1:10545 npx hardhat run \
 *     /home/niko/eth/dev/b/backend/cmd/cgctl/dev-scripts/upgrade-fork-v3.js
 *
 * Env:
 *   RPC_URL               fork node RPC            (default http://127.0.0.1:10545)
 *   GAME_ADDR             game proxy               (default: production proxy)
 *   ACTIVATION_DELAY_SEC  round reactivation delay (default 300; new-round jumps past it)
 *   SKIP_PRIZES_WALLET    set to 1 to keep the current PrizesWallet
 */
const hre = global.hre ?? require("hardhat");

const RPC_URL = process.env.RPC_URL || "http://127.0.0.1:10545";
const GAME_ADDR = process.env.GAME_ADDR || "0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2";
const ACTIVATION_DELAY_SEC = BigInt(process.env.ACTIVATION_DELAY_SEC || "300");
const SKIP_PRIZES_WALLET = process.env.SKIP_PRIZES_WALLET === "1";

const ERC1967_IMPL_SLOT = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc";

// Explicit gas limits: eth_estimateGas re-executes the tx repeatedly, and every
// cold state read goes to the upstream, whose historical-state window for the
// pinned block is only ~30 s on a non-archive nitro node. Fewer executions,
// fewer upstream round-trips. The fork's blockGasLimit is huge; generosity is free.
const gCall = { gasLimit: 2_000_000 };
const gDeploy = { gasLimit: 60_000_000 };

// Values written by reinitialize(3) — verified at the end. The review-branch
// contracts (claude/v3-contract-review-dc3c65) removed the CST price decline
// multiplier pair; the decline rate now derives from the bid CST reward rate,
// and cstDutchAuctionBeginningBidPriceMinLimit is frozen at its V2 value.
const V3_GETTERS = [
    "cstDutchAuctionBeginningBidPriceMinLimit",
    "getNextCstBidPriceBase",
    "roundLateBidDurationDivisor",
    "roundLateBidPricePremiumAmountBaseMultiplier",
    "roundLateBidPricePremiumAmountExponent",
    "mainPrizeNumCosmicSignatureNfts",
];

async function implementationOf(provider, proxyAddr) {
    const word = await provider.getStorage(proxyAddr, ERC1967_IMPL_SLOT);
    return hre.ethers.getAddress("0x" + word.slice(-40));
}

async function main() {
    const provider = new hre.ethers.JsonRpcProvider(RPC_URL);

    // ---- fork guard: hardhat_metadata only exists on a Hardhat node ----
    try {
        await provider.send("hardhat_metadata", []);
    } catch {
        console.error(`Error: ${RPC_URL} is not a Hardhat node — refusing to run (production safety).`);
        process.exit(1);
    }

    const v3Art = await hre.artifacts.readArtifact("CosmicSignatureGameV3");
    const pwArt = await hre.artifacts.readArtifact("PrizesWallet");

    // The V3 ABI is a superset covering everything we call pre- and post-upgrade
    // (owner/roundNum/lastBidderAddress/setRoundActivationTime/upgradeToAndCall all
    // exist on the live V2 too).
    const gameRO = new hre.ethers.Contract(GAME_ADDR, v3Art.abi, provider);
    const ownerAddr = await gameRO.owner();
    const roundNum = await gameRO.roundNum();
    const lastBidder = await gameRO.lastBidderAddress();
    const implBefore = await implementationOf(provider, GAME_ADDR);
    const pwBefore = await gameRO.prizesWallet();
    console.log(`Proxy ${GAME_ADDR} (fork of production)`);
    console.log(`  owner=${ownerAddr}  roundNum=${roundNum}`);
    console.log(`  implementation=${implBefore}  prizesWallet=${pwBefore}`);

    if (lastBidder !== hre.ethers.ZeroAddress) {
        console.error("Error: the current round has bids. Claim the main prize first:");
        console.error("  ./fork-v3-rehearsal.sh jump && ./fork-v3-rehearsal.sh claim");
        process.exit(1);
    }

    // Prewarm account metadata for every address this flow (and the follow-up
    // first bid) touches, while the upstream can still serve the pinned block
    // (see the gas-limit comment). Even hardhat_setCode needs the account
    // fetched first, so nothing can be seeded after the upstream prunes.
    const prewarmAddrs = [
        GAME_ADDR, ownerAddr, implBefore, pwBefore,
        hre.ethers.ZeroAddress,                         // go-ethereum binders probe it
        "0x0000000000000000000000000000000000000064",   // ArbSys (bids probe it, expect failure)
        "0x000000000000000000000000000000000000006C",   // ArbGasInfo (same)
        "0xbb84Be3500A63581d3F2d5AC3bdF8685AAedad25",   // CosmicSignature NFT
        "0xAD91843e6A58Ba560F577E676986AFb1dba6FBA0",   // CosmicSignatureToken
        "0xF3D52E1c681949be7E624778dB13DaD7F8c729db",   // DAO
        "0x96bB0ADB414d5350f435E52f94946B6C7A0760a9",   // CharityWallet
        "0x895a6F444BE4ba9d124F61DF736605792B35D66b",   // RandomWalk NFT
        "0x6308A405B4FF1eA890870Efe2a6D036750B81F7C",   // StakingWallet CST
        "0x5EB3396092841E6c5b0b51141699F6711E830529",   // StakingWallet RWalk
        "0xa3802c799f5e3D3D3562A9B513a41C6aAF92e25e",   // MarketingWallet
    ];
    for (const addr of prewarmAddrs) {
        await provider.send("eth_getTransactionCount", [addr, "latest"]);
        await provider.send("eth_getBalance", [addr, "latest"]);
        await provider.send("eth_getCode", [addr, "latest"]);
    }

    // Neutralize the Arbitrum precompiles. eth_getCode reports the placeholder
    // byte 0xfe (INVALID) for them, and a fork executes it literally: the game's
    // try/catch still catches the failure, but INVALID burns ALL forwarded gas,
    // which starves claimMainPrize. Empty code makes the calls return empty and
    // fail cheaply inside the try — identical to a plain hardhat dev chain
    // (ArbitrumError events + fallback randomness).
    await provider.send("hardhat_setCode", ["0x0000000000000000000000000000000000000064", "0x"]); // ArbSys
    await provider.send("hardhat_setCode", ["0x000000000000000000000000000000000000006C", "0x"]); // ArbGasInfo

    // ---- impersonate the real owner ----
    await provider.send("hardhat_impersonateAccount", [ownerAddr]);
    await provider.send("hardhat_setBalance", [ownerAddr, "0x21E19E0C9BAB2400000"]); // 10k ETH
    // Not provider.getSigner(ownerAddr): that validates against eth_accounts,
    // which does not list impersonated accounts.
    const owner = new hre.ethers.JsonRpcSigner(provider, ownerAddr);
    const game = gameRO.connect(owner);

    // Capture every completed round's bid count THROUGH THE DEPLOYED V2 ABI
    // (public mapping getter `bidderAddresses(round).numItems`). After the
    // upgrade the same storage slot is read back through the V3 ABI
    // (`roundStats(round).numBids`): equality proves the slot-259
    // reinterpretation (BidderAddresses -> RoundStats) on real deployed data —
    // the property the @custom:oz-renamed-from/oz-retyped-from annotations
    // assert for OpenZeppelin's off-chain validator, which this raw-transaction
    // rehearsal intentionally bypasses.
    const legacyAbi = ["function bidderAddresses(uint256 roundNum_) view returns (uint256 numItems)"];
    const legacyGame = new hre.ethers.Contract(GAME_ADDR, legacyAbi, provider);
    const legacyBidCounts = [];
    for (let r = 0n; r < roundNum; r++) {
        legacyBidCounts.push(await legacyGame.bidderAddresses(r));
    }
    console.log(`Pre-upgrade bid counts via deployed V2 ABI: [${legacyBidCounts.join(", ")}]`);

    // ---- 1) defer activation: _authorizeUpgrade and the setters need an inactive round ----
    const block = await provider.getBlock("latest");
    const farFuture = BigInt(block.timestamp) + 86400n;
    try {
        await (await game.setRoundActivationTime(farFuture, gCall)).wait();
        console.log(`Deferred round activation to ${farFuture} for the upgrade window.`);
    } catch (err) {
        console.error("setRoundActivationTime failed — is the round active? Claim the prize first.");
        console.error(err.shortMessage || err.message || err);
        process.exit(1);
    }

    // ---- 2) deploy the V3 implementation and upgrade (same calldata as production) ----
    console.log("Deploying CosmicSignatureGameV3 implementation...");
    const v3Factory = new hre.ethers.ContractFactory(v3Art.abi, v3Art.bytecode, owner);
    const v3Impl = await v3Factory.deploy(gDeploy);
    await v3Impl.waitForDeployment();
    console.log(`  implementation deployed at ${v3Impl.target}`);

    console.log("upgradeToAndCall(impl, reinitialize())...");
    const reinitData = game.interface.encodeFunctionData("reinitialize");
    await (await game.upgradeToAndCall(v3Impl.target, reinitData, gCall)).wait();

    const implAfter = await implementationOf(provider, GAME_ADDR);
    if (implAfter !== hre.ethers.getAddress(String(v3Impl.target))) {
        console.error(`Error: ERC1967 slot points at ${implAfter}, expected ${v3Impl.target}.`);
        process.exit(1);
    }
    console.log(`V3 upgrade complete, implementation=${implAfter}`);

    // ---- storage reinterpretation check: old rounds read back through V3 types ----
    for (let r = 0n; r < roundNum; r++) {
        const rs = await game.roundStats(r);
        const numBids = rs[0];
        const numCstBids = rs[1];
        if (numBids !== legacyBidCounts[Number(r)]) {
            console.error(`Error: round ${r}: V3 roundStats.numBids=${numBids} != deployed V2 numItems=${legacyBidCounts[Number(r)]}`);
            process.exit(1);
        }
        if (numCstBids !== 0n) {
            console.error(`Error: round ${r}: numCstBids=${numCstBids}, expected 0 for a V2-completed round`);
            process.exit(1);
        }
        const first = await game.getBidInfoAt(r, 0n);
        const last = await game.getBidInfoAt(r, numBids - 1n);
        if (first.bidderAddress === hre.ethers.ZeroAddress || last.bidderAddress === hre.ethers.ZeroAddress) {
            console.error(`Error: round ${r}: zero bidder address in reinterpreted bid info`);
            process.exit(1);
        }
        if (first.raffleCumulativeWeight !== 0n || last.raffleCumulativeWeight !== 0n) {
            console.error(`Error: round ${r}: nonzero raffleCumulativeWeight on a V2- bid (slot beyond the old word must be zero)`);
            process.exit(1);
        }
        console.log(`Round ${r}: ${numBids} bids read back identically through V3 types `
            + `(first=${first.bidderAddress}, last=${last.bidderAddress}, V3-only fields zero)`);
    }
    console.log("Storage reinterpretation verified: deployed V2 bidderAddresses data reads correctly as V3 roundStats.");

    // A second reinitialize must revert with InvalidInitialization.
    try {
        await game.reinitialize.staticCall();
        console.error("Error: a second reinitialize() did NOT revert — investigate before the real upgrade.");
        process.exit(1);
    } catch {
        console.log("Second reinitialize() reverts as expected (reinitializer(3) consumed).");
    }

    // ---- 3) replace the PrizesWallet (part of the production upgrade plan) ----
    let newPrizesWallet = null;
    if (SKIP_PRIZES_WALLET) {
        console.log("SKIP_PRIZES_WALLET=1 — keeping the current PrizesWallet.");
    } else {
        console.log("Deploying new PrizesWallet(game) and calling setPrizesWallet...");
        const pwFactory = new hre.ethers.ContractFactory(pwArt.abi, pwArt.bytecode, owner);
        const pw = await pwFactory.deploy(GAME_ADDR, gDeploy);
        await pw.waitForDeployment();
        await (await game.setPrizesWallet(pw.target, gCall)).wait();
        const pwOnGame = await game.prizesWallet();
        const gameOnPw = await new hre.ethers.Contract(pw.target, pwArt.abi, provider).game();
        if (pwOnGame !== hre.ethers.getAddress(String(pw.target)) || gameOnPw !== hre.ethers.getAddress(GAME_ADDR)) {
            console.error(`Error: cross-pointers wrong: game.prizesWallet=${pwOnGame}, wallet.game=${gameOnPw}`);
            process.exit(1);
        }
        newPrizesWallet = String(pw.target);
        console.log(`  new PrizesWallet=${newPrizesWallet} (game.prizesWallet and wallet.game verified)`);
    }

    // ---- 4) re-emit the 4 V3 config admin events (reinitialize emits only Initialized(3)) ----
    console.log("Emitting V3 config admin events (values unchanged)...");
    await (await game.setRoundLateBidDurationDivisor(await game.roundLateBidDurationDivisor(), gCall)).wait();
    await (await game.setRoundLateBidPricePremiumAmountBaseMultiplier(await game.roundLateBidPricePremiumAmountBaseMultiplier(), gCall)).wait();
    await (await game.setRoundLateBidPricePremiumAmountExponent(await game.roundLateBidPricePremiumAmountExponent(), gCall)).wait();
    await (await game.setMainPrizeNumCosmicSignatureNfts(await game.mainPrizeNumCosmicSignatureNfts(), gCall)).wait();

    // ---- 5) reactivate the round and verify what reinitialize() wrote ----
    const nowBlock = await provider.getBlock("latest");
    const activationTs = BigInt(nowBlock.timestamp) + ACTIVATION_DELAY_SEC;
    await (await game.setRoundActivationTime(activationTs, gCall)).wait();

    console.log("");
    console.log("Values written by reinitialize(3):");
    for (const getter of V3_GETTERS) {
        console.log(`  ${getter} = ${await game[getter]()}`);
    }
    console.log("");
    console.log(`Done. roundNum=${await game.roundNum()} now runs on V3.`);
    console.log(`Round activates at ${activationTs} (${ACTIVATION_DELAY_SEC}s of chain time; `
        + `'./fork-v3-rehearsal.sh new-round' jumps past it).`);

    // Keep the backend's contract registry in sync (run against the DB that tracks this fork).
    console.log("");
    console.log(`UPDATE cg_contracts SET implementation_addr='${implAfter}' WHERE cosmic_game_addr='${GAME_ADDR}';`);
    if (newPrizesWallet) {
        console.log(`UPDATE cg_contracts SET prizes_wallet_addr='${newPrizesWallet}' WHERE cosmic_game_addr='${GAME_ADDR}';`);
    }
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err);
        process.exit(1);
    });
