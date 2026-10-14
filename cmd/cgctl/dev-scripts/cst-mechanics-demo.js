"use strict";
// cst-mechanics-demo.js — observational demo of the review-branch V3 CST bid pricing
// (decline = reward rate, additive restart with cap, premium floor, weight = fraction paid).
//
// Run from the Cosmic-Signature repo on the IN-PROCESS hardhat network (no node needed):
//   cd <Cosmic-Signature repo> && npx hardhat run <dev-scripts>/cst-mechanics-demo.js
//
// Deploys V1, completes round 0, upgrades V2 -> V3 with the repo's own test helpers,
// then prints: (A) the CST price curve of a fresh round, (B) a two-account
// "hold the deadline with CST" attack and its per-hour cost, (C) why waiting
// for a free bid loses the round. CST for the attackers is minted by
// impersonating the game proxy (test scaffolding; only the game may mint).

// `hardhat run` injects `hre` globally; fall back to resolving from the repo (cwd).
const hre = globalThis.hre ?? require(require.resolve("hardhat", { paths: [process.cwd()] }));

const ONE = 10n ** 18n;
const fmt = (wei_, dp_ = 2) => {
	const s = hre.ethers.formatEther(wei_);
	const i = s.indexOf(".");
	return (i < 0) ? s : s.slice(0, i + 1 + dp_);
};
const hours = (sec_) => (Number(sec_) / 3600).toFixed(2);

async function main() {
	const root = hre.config.paths.root;
	const { loadFixtureDeployContractsForTesting } = require(`${root}/src/ContractTestingHelpers.js`);
	const { waitForTransactionReceipt } = require(`${root}/src/Helpers.js`);
	const { completeRoundZero, upgradeToV2 } = require(`${root}/test/src/V2UpgradeTestHelpers.js`);
	const { upgradeToV3 } = require(`${root}/test/src/V3UpgradeTestHelpers.js`);

	const now = async () => BigInt((await hre.ethers.provider.getBlock("latest")).timestamp);
	const sleepChain = async (sec_) => {
		await hre.network.provider.send("evm_increaseTime", [Number(sec_)]);
		await hre.network.provider.send("evm_mine");   // make view calls see the new time
	};

	// ---------- deploy V1, complete round 0, upgrade V2 -> V3 ----------
	console.log("Deploying V1 + peripherals, completing round 0, upgrading V2 -> V3 (repo helpers)...");
	// 2n = "activate the round 2 seconds from the current block" (relative-offset magic).
	const contracts = { ...(await loadFixtureDeployContractsForTesting(2n)) };
	await completeRoundZero(contracts);
	await upgradeToV2(contracts);
	await upgradeToV3(contracts);
	const game = contracts.cosmicSignatureGameV3Proxy ?? contracts.cosmicSignatureGameProxy;
	const token = contracts.cosmicSignatureToken;
	const owner = contracts.ownerSigner;
	const [a1, a2] = [contracts.signers[5], contracts.signers[6]];
	console.log(`roundNum=${await game.roundNum()} (V3 live on a local chain)`);

	// Derived scale facts (nothing stored):
	const increment = await game.getMainPrizeTimeIncrement();                  // seconds
	const rewardMul = await game.bidCstRewardAmountMultiplier();
	const incrementUs = await game.mainPrizeTimeIncrementInMicroSeconds();
	const ratePerSec = rewardMul / incrementUs;                                // CST wei / second
	const R = ratePerSec * increment;                                          // reward per increment
	console.log(`increment=${hours(increment)} h   decline/reward rate=${fmt(ratePerSec * 3600n)} CST/h   R=${fmt(R)} CST per increment`);

	// ---------- fund the attackers (test scaffolding: only the game may mint) ----------
	const gameAddr = await game.getAddress();
	await hre.network.provider.send("hardhat_setBalance", [gameAddr, "0x21E19E0C9BAB2400000"]);
	const gameSigner = await hre.ethers.getImpersonatedSigner(gameAddr);
	const WAR_CHEST = 1_000_000n * ONE;
	for (const a of [a1, a2]) await waitForTransactionReceipt(token.connect(gameSigner).mint(a.address, WAR_CHEST));
	console.log(`minted ${fmt(WAR_CHEST, 0)} CST to each of 2 attacker accounts (the "1 million CST" premise)\n`);

	// ---------- open round 1 and place its first (ETH) bid ----------
	await waitForTransactionReceipt(game.connect(owner).setRoundActivationTime(1n));
	await waitForTransactionReceipt(game.connect(a1).bidWithEth(-1n, "round-opening ETH bid", 0n, { value: ONE }));
	const t0 = await now();
	const S0 = await game.getNextCstBidPriceBaseAdvanced(0n);
	const deadline0 = await game.mainPrizeTime();
	console.log("DEMO A — CST asking price curve after the round's first ETH bid (view calls, no time travel)");
	console.log(`auction seeded at S = ${fmt(S0)} CST (= C, one initial-countdown of reward; derived, not stored)`);
	console.log("  t+hours   base price   full ask (incl. late-bid premium)");
	for (let h = 0; h <= 26; h += 2) {
		const base = await game.getNextCstBidPriceBaseAdvanced(BigInt(h * 3600));
		const ask = await game.getNextCstBidPriceAdvanced(BigInt(h * 3600));
		const note = (BigInt(h * 3600) > deadline0 - t0 && base === 0n) ? "   <- rule 3: sniping at zero base still costs ~4R" : "";
		console.log(`  ${String(h).padStart(7)}   ${fmt(base).padStart(10)}   ${fmt(ask).padStart(10)}${note}`);
	}
	console.log(`mainPrizeTime is t+${hours(deadline0 - t0)} h — the base hits 0 exactly when the initial countdown would end\n`);

	// ---------- DEMO B: hold-the-deadline attack, bids just under 1 increment apart ----------
	console.log("DEMO B — two accounts try to hold the round open with CST (spacing 3500 s < 1 h increment)");
	console.log("advancing chain time ~23 h so the ask is low (the cheapest possible starting point)...");
	await sleepChain(23n * 3600n - 60n);
	console.log("  bid  bidder  ask paid    restart S    reward to displaced holder");
	let steadyDrain = 0n;
	for (let i = 1; i <= 12; i++) {
		const bidder = (i % 2 === 1) ? a1 : a2;
		const bidderKey = (i % 2 === 1) ? "a1" : "a2";
		const ask = await game.getNextCstBidPrice();
		const reward = await game.getBidCstRewardAmount();
		await waitForTransactionReceipt(game.connect(bidder).bidWithCst(hre.ethers.MaxUint256, "", 0n));
		const S = await game.cstDutchAuctionBeginningBidPrice();
		const note = (i === 1) ? "  (reward went to the bidder itself — it also held the ETH-bid position)" : "";
		console.log(`  ${String(i).padStart(3)}  ${bidderKey}     ${fmt(ask).padStart(8)}     ${fmt(S).padStart(8)}     +${fmt(reward)}${note}`);
		steadyDrain = ask - reward;   // net CST destroyed per cycle once at the cap
		await sleepChain(3500n);
	}
	const drainPerIncrement = steadyDrain * increment / 3500n;
	const incrementsAffordable = (2n * WAR_CHEST) / drainPerIncrement;
	console.log(`steady state: each bid burns ${fmt(steadyDrain)} CST net (ask - reward recouped) per 3500 s`);
	console.log(`=> ~${fmt(drainPerIncrement)} CST per increment of delay; a 2,000,000 CST war chest lasts ~${incrementsAffordable} increments`);
	console.log(`   = ~${(Number(incrementsAffordable) * Number(increment) / 86400).toFixed(1)} days, strictly decreasing — vs FOREVER at ~zero cost in the pre-review draft`);
	console.log(`(income from holding the position is ${fmt(R)} CST/increment — ${Number(drainPerIncrement / R)}x less than the burn)\n`);

	// ---------- DEMO C: the free-bid strategy loses the round ----------
	console.log("DEMO C — stop paying and wait for a free (0 CST) bid instead:");
	const S = await game.cstDutchAuctionBeginningBidPrice();
	const tNow = await now();
	const tFree = S / ratePerSec;                       // seconds until ask reaches 0 (from last restart)
	const tBegin = await game.cstDutchAuctionBeginningTimeStamp();
	const freeAt = tBegin + tFree;
	const deadline = await game.mainPrizeTime();
	console.log(`  ask reaches 0 CST at   t+${hours(freeAt - tNow)} h  (S=${fmt(S)} declining ${fmt(ratePerSec * 3600n)}/h)`);
	console.log(`  mainPrizeTime is at    t+${hours(deadline - tNow)} h`);
	console.log(`  => the deadline arrives ${hours((freeAt - tNow) - (deadline - tNow))} h before the free bid exists; a bid extends the deadline only 1 increment,`);
	console.log(`     so free bids can never keep the round open. The last paying bidder now claims:`);
	await hre.network.provider.send("evm_setNextBlockTimestamp", [Number(deadline + 1n)]);
	const lastBidder = (await game.lastBidderAddress()) === a1.address ? a1 : a2;
	await waitForTransactionReceipt(game.connect(lastBidder).claimMainPrize());
	console.log(`  claimMainPrize OK — roundNum=${await game.roundNum()}. The "infinite 1-CST bidding" scenario is structurally impossible:`);
	console.log(`  there is no 1 CST minimum (restart is >= 2R = ${fmt(2n * R)} CST) and you never choose your price — you pay the auction's ask.`);
}

main().then(() => process.exit(0)).catch((e) => { console.error(e); process.exit(1); });
