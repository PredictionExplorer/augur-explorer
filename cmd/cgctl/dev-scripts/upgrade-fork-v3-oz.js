/**
 * upgrade-fork-v3-oz.js — the STANDARD OpenZeppelin upgrade procedure on the fork.
 * Counterpart to upgrade-fork-v3.js (which sends the raw upgradeToAndCall and
 * bypasses OZ). This one goes through hre.upgrades.upgradeProxy with NO unsafe
 * flags, exactly like the production upgrade task, so the validator gate is
 * exercised for real:
 *
 *   - the OZ manifest is seeded with the DEPLOYED V2 storage layout
 *     (test/data/deployed-storage-layouts/arbitrumOne-CosmicSignatureGameV2.json),
 *     which is what the preserved runners/.openzeppelin manifest provides on
 *     upgrade day;
 *   - on stock `main` sources this FAILS by design ("Renamed `bidderAddresses`
 *     to `roundStats`...") because the @custom:oz-renamed-from annotation is
 *     commented out — proving the gate works;
 *   - after flipping the annotations per
 *     tasks/docs/Cosmic-Signature-Game-Contract-Upgrade-And-Re-Registration.md
 *     (uncomment oz-renamed-from on roundStats; comment out the obsolete
 *     ToDo-202607169-1 annotations) and recompiling, it passes and performs
 *     the upgrade + reinitialize through the OZ plugin.
 *
 * The OZ plugin works through hre.network, so this script must run with
 * --network hardhat_on_localhost and the fork node must listen on port 8545:
 *
 *   HH_PORT=8545 ... fork-and-prepare.sh
 *   cd <Cosmic-Signature repo>
 *   npx hardhat run .../upgrade-fork-v3-oz.js --network hardhat_on_localhost
 *
 * Env: GAME_ADDR, ACTIVATION_DELAY_SEC, SKIP_PRIZES_WALLET — as in upgrade-fork-v3.js.
 */
const hre = global.hre ?? require("hardhat");
const fs = require("fs");
const path = require("path");

const GAME_ADDR = process.env.GAME_ADDR || "0x6a714Ae7B5b6eA520F6BCA23d2E609C4Fd5863F2";
const ACTIVATION_DELAY_SEC = BigInt(process.env.ACTIVATION_DELAY_SEC || "300");
const SKIP_PRIZES_WALLET = process.env.SKIP_PRIZES_WALLET === "1";

const ERC1967_IMPL_SLOT = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc";
const DEPLOYED_LAYOUT_JSON = "test/data/deployed-storage-layouts/arbitrumOne-CosmicSignatureGameV2.json";

async function implementationOf(proxyAddr) {
    const word = await hre.ethers.provider.getStorage(proxyAddr, ERC1967_IMPL_SLOT);
    return hre.ethers.getAddress("0x" + word.slice(-40));
}

async function main() {
    if (hre.network.name === "hardhat") {
        console.error("Error: run with --network hardhat_on_localhost (fork node on port 8545).");
        process.exit(1);
    }
    try {
        await hre.network.provider.send("hardhat_metadata", []);
    } catch {
        console.error(`Error: network '${hre.network.name}' is not a Hardhat node — refusing to run (production safety).`);
        process.exit(1);
    }

    const gameArt = await hre.artifacts.readArtifact("CosmicSignatureGameV3");
    const gameRO = new hre.ethers.Contract(GAME_ADDR, gameArt.abi, hre.ethers.provider);
    const ownerAddr = await gameRO.owner();
    const roundNum = await gameRO.roundNum();
    const implBefore = await implementationOf(GAME_ADDR);
    console.log(`Proxy ${GAME_ADDR} (fork of production)`);
    console.log(`  owner=${ownerAddr}  roundNum=${roundNum}  implementation=${implBefore}`);
    if ((await gameRO.lastBidderAddress()) !== hre.ethers.ZeroAddress) {
        console.error("Error: the current round has bids. Claim the main prize first.");
        process.exit(1);
    }

    // Deployed V2 layout for the manifest seed.
    const deployed = JSON.parse(fs.readFileSync(path.join(hre.config.paths.root, DEPLOYED_LAYOUT_JSON), "utf8"));
    if (hre.ethers.getAddress(deployed.implementationAddress) !== implBefore) {
        console.error(`Error: on-chain implementation ${implBefore} != ${deployed.implementationAddress} in ${DEPLOYED_LAYOUT_JSON}.`);
        console.error("The deployed-layout snapshot does not describe the live implementation; refusing to seed the manifest with it.");
        process.exit(1);
    }

    // ---- impersonate the real owner ----
    await hre.network.provider.send("hardhat_impersonateAccount", [ownerAddr]);
    await hre.network.provider.send("hardhat_setBalance", [ownerAddr, "0x21E19E0C9BAB2400000"]);
    const owner = await hre.ethers.getSigner(ownerAddr);
    const game = gameRO.connect(owner);

    // Pre-upgrade bid counts through the deployed V2 ABI (storage reinterpretation check).
    const legacyGame = new hre.ethers.Contract(GAME_ADDR,
        ["function bidderAddresses(uint256 roundNum_) view returns (uint256 numItems)"], hre.ethers.provider);
    const legacyBidCounts = [];
    for (let r = 0n; r < roundNum; r++) legacyBidCounts.push(await legacyGame.bidderAddresses(r));
    console.log(`Pre-upgrade bid counts via deployed V2 ABI: [${legacyBidCounts.join(", ")}]`);

    // ---- defer activation (the upgrade authorizer requires an inactive round) ----
    const block = await hre.ethers.provider.getBlock("latest");
    await (await game.setRoundActivationTime(BigInt(block.timestamp) + 86400n)).wait();
    console.log("Deferred round activation for the upgrade window.");

    // ---- the OZ storage validation gate, against the DEPLOYED layout ----
    // Same API the committed CosmicSignatureGameV3-DeployedV2StorageLayout.js
    // test uses. The V3 layout comes from the compile-time validations, so the
    // @custom:oz-renamed-from / oz-retyped-from annotations in the SOURCES
    // decide the outcome — exactly like the upgrade-day task running against
    // the preserved manifest.
    // This script lives outside the hardhat project, so resolve the OZ modules
    // from the project's own node_modules.
    const fromRepo = (mod) => require(require.resolve(mod, { paths: [hre.config.paths.root] }));
    const { readValidations } = fromRepo("@openzeppelin/hardhat-upgrades/dist/utils/validations");
    const {
        getUnlinkedBytecode, getVersion, getStorageLayout,
        getStorageUpgradeReport, withValidationDefaults,
    } = fromRepo("@openzeppelin/upgrades-core");
    const V3 = await hre.ethers.getContractFactory("CosmicSignatureGameV3", owner);
    const validations = await readValidations(hre);
    const v3Version = getVersion(getUnlinkedBytecode(validations, V3.bytecode), V3.bytecode);
    const v3Layout = getStorageLayout(validations, v3Version);
    const report = getStorageUpgradeReport(deployed.layout, v3Layout, withValidationDefaults({}));
    if (!report.pass) {
        console.error("");
        console.error("OpenZeppelin validation REJECTED the upgrade (deployed V2 layout -> V3 sources):");
        console.error(report.explain());
        console.error("On stock `main` sources this is the expected outcome. To rehearse the real");
        console.error("deployment-branch state, follow the runbook (Cosmic-Signature-Game-Contract-");
        console.error("Upgrade-And-Re-Registration.md): in CosmicSignatureGameStorageV2Base.sol");
        console.error("uncomment `@custom:oz-renamed-from bidderAddresses` on roundStats and comment");
        console.error("out the obsolete annotations tagged ToDo-202607169-1, recompile, and rerun.");
        process.exit(2);
    }
    console.log("OZ storage validation PASSED against the DEPLOYED V2 layout (no unsafe flags).");

    // ---- the standard upgrade: hre.upgrades.upgradeProxy (deploy + upgradeToAndCall) ----
    // forceImport registers the proxy in the (ephemeral dev) manifest first; the
    // plugin then re-validates against the refactored V2 layout it recorded, which
    // tolerates identical annotations on both sides (see docs/v3-vs-v2-changes.md).
    const V2 = await hre.ethers.getContractFactory("CosmicSignatureGameV2", owner);
    await hre.upgrades.forceImport(GAME_ADDR, V2, { kind: "uups" });
    console.log("hre.upgrades.upgradeProxy(proxy, CosmicSignatureGameV3, { call: reinitialize })...");
    const proxyV3 = await hre.upgrades.upgradeProxy(GAME_ADDR, V3, { kind: "uups", call: "reinitialize" });
    await proxyV3.waitForDeployment();
    const implAfter = await implementationOf(GAME_ADDR);
    console.log(`OZ validation PASSED; V3 upgrade complete, implementation=${implAfter}`);

    // ---- storage reinterpretation check: old rounds read back through V3 types ----
    for (let r = 0n; r < roundNum; r++) {
        const rs = await game.roundStats(r);
        if (rs[0] !== legacyBidCounts[Number(r)]) {
            console.error(`Error: round ${r}: V3 roundStats.numBids=${rs[0]} != deployed V2 numItems=${legacyBidCounts[Number(r)]}`);
            process.exit(1);
        }
    }
    console.log(`Storage reinterpretation verified for rounds 0..${roundNum - 1n} (bid counts identical through V3 types).`);

    try {
        await game.reinitialize.staticCall();
        console.error("Error: a second reinitialize() did NOT revert — investigate before the real upgrade.");
        process.exit(1);
    } catch {
        console.log("Second reinitialize() reverts as expected (reinitializer(3) consumed).");
    }

    // ---- PrizesWallet replacement (not a proxy; plain deployment, as in production) ----
    if (SKIP_PRIZES_WALLET) {
        console.log("SKIP_PRIZES_WALLET=1 — keeping the current PrizesWallet.");
    } else {
        const PW = await hre.ethers.getContractFactory("PrizesWallet", owner);
        const pw = await PW.deploy(GAME_ADDR);
        await pw.waitForDeployment();
        await (await game.setPrizesWallet(pw.target)).wait();
        console.log(`New PrizesWallet=${pw.target} set on the game.`);
    }

    // ---- re-emit the 4 V3 config admin events; reactivate ----
    await (await game.setRoundLateBidDurationDivisor(await game.roundLateBidDurationDivisor())).wait();
    await (await game.setRoundLateBidPricePremiumAmountBaseMultiplier(await game.roundLateBidPricePremiumAmountBaseMultiplier())).wait();
    await (await game.setRoundLateBidPricePremiumAmountExponent(await game.roundLateBidPricePremiumAmountExponent())).wait();
    await (await game.setMainPrizeNumCosmicSignatureNfts(await game.mainPrizeNumCosmicSignatureNfts())).wait();
    console.log("V3 config admin events emitted (values unchanged).");

    const nowBlock = await hre.ethers.provider.getBlock("latest");
    const activationTs = BigInt(nowBlock.timestamp) + ACTIVATION_DELAY_SEC;
    await (await game.setRoundActivationTime(activationTs)).wait();

    console.log("");
    console.log(`Done via the STANDARD OZ procedure. roundNum=${await game.roundNum()} now runs on V3 (implementation=${implAfter}).`);
    console.log(`Round activates at ${activationTs}. Continue with fork-v3-rehearsal.sh smoke (RPC_URL on port 8545).`);
    console.log("");
    console.log(`UPDATE cg_contracts SET implementation_addr='${implAfter}' WHERE cosmic_game_addr='${GAME_ADDR}';`);
    if (!SKIP_PRIZES_WALLET) {
        console.log(`UPDATE cg_contracts SET prizes_wallet_addr='${await game.prizesWallet()}' WHERE cosmic_game_addr='${GAME_ADDR}';`);
    }
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err);
        process.exit(1);
    });
