/**
 * chaoszero-sim.js — optional ChaosZero market simulation for populate.js.
 *
 * Enabled by WITH_CHAOSZERO=1 (deploy-and-populate.sh --with-chaoszero).
 * Deploys the REAL production ChaosZero contracts (V3OutcomeSource +
 * IndependentOutcomeMarket, artifacts read from CHAOSZERO_DIR) on top of the
 * live local CosmicSignatureGame, so ChaosZero market rounds are the very
 * same rounds the cosmicgame ETL indexes — one game, one set of rounds,
 * exactly like the Arbitrum production deployment.
 *
 * populate.js calls the five phases between its own steps; each phase runs
 * while the game sits in a known round state:
 *
 *   phaseA  after S3 (round R1 open): deploy source+market with the real
 *           roundStats(R1-1) baseline; mint trader CST by impersonating the
 *           game; funding variants (donation, direct transfers absorbed by a
 *           buy AND a sell, sub-threshold held CST, zero-amount fundRound
 *           trigger); both-sides purchases on three questions (a purchased
 *           winner is guaranteed whatever the game decides); partial and
 *           full exits; splits; a genuine 4-id operator batch transfer;
 *           post-revocation transfer asserted to revert; self- and
 *           zero-value transfers; plural multi-question merge.
 *   phaseB  after S5 (R1 closed): post-closure pre-settlement merge;
 *           explicit settleRound(); post-settlement vault-paid merge;
 *           post-settlement transfer of the winning side to a fresh wallet
 *           that claims; claims asserted against computed entitlements
 *           (including a guaranteed zero-total claim); one winner left
 *           escrowed; round R2 opened by its first buy (carry rollover).
 *   phaseC  after S6 (R2 closed): post-closure merge on the still
 *           UNSETTLED round R2 (left unsettled on purpose).
 *   phaseD  after S9 (R4 open; R3 = S7's round got NO market activity):
 *           fundRound(R4) must settle R2 across the skipped round in the
 *           same receipt, with donation+transfers+carry all nonzero and
 *           allocation remainders within 1 wei; R2 claims; R4 both-sides
 *           purchases.
 *   phaseE  after S10 (R4 closed): claimKnown() triggers the settlement of
 *           R4 (third settlement path); final claims; one entitlement left
 *           escrowed.
 *   phaseF  after S10 + activation delay: open the NEXT game round with a
 *           bid and leave its market round rich and UNSETTLED — funding
 *           that absorbs the settled history's carry, both-sided buys, an
 *           early-exit sell, a direct-transfer absorb, a split, a partial
 *           merge and a gifted position. Every earlier round ends settled,
 *           so this is the only intermediate-state sample the populated
 *           database gets (open liquidity, OPEN badges, live odds). Ends
 *           with the event tally: per-type minimums and structural checks
 *           (throws, so populate.js reports POPULATE FAILED).
 *
 * Outcome values are never assumed: the game statistics are dictated by the
 * populate.js scenario, so every claim assertion derives the entitlement
 * from the actual settled values and live position balances.
 */
"use strict";

const fs = require("fs");
const path = require("path");

const ONE = 10n ** 18n;

function enabled() {
    return ["1", "true", "yes"].includes(String(process.env.WITH_CHAOSZERO || "").toLowerCase());
}

// --------------------------------------------------------------------------
// ChaosZero artifacts (hardhat v3 or foundry build output)
// --------------------------------------------------------------------------
function chaosZeroDir() {
    const dir = process.env.CHAOSZERO_DIR;
    if (!dir) throw new Error("WITH_CHAOSZERO is set but CHAOSZERO_DIR does not point at the ChaosZero checkout");
    if (!fs.existsSync(path.join(dir, "src", "IndependentOutcomeMarket.sol"))) {
        throw new Error(`CHAOSZERO_DIR=${dir} does not look like the ChaosZero repo`);
    }
    return dir;
}

function artifactRoots(root) {
    return [
        path.join(root, ".local", "hardhat-artifacts"),
        path.join(root, "out"),
        path.join(root, "artifacts"),
    ].filter((p) => fs.existsSync(p));
}

function findArtifact(root, name) {
    const target = `${name}.json`;
    const stack = artifactRoots(root);
    if (stack.length === 0) {
        throw new Error(`No ChaosZero artifacts under ${root}; build it first (pnpm build:deployment or forge build)`);
    }
    while (stack.length > 0) {
        const dir = stack.pop();
        for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
            const full = path.join(dir, entry.name);
            if (entry.isDirectory()) {
                if (entry.name === "build-info") continue;
                stack.push(full);
                continue;
            }
            if (entry.name !== target || full.endsWith(".dbg.json")) continue;
            const parsed = JSON.parse(fs.readFileSync(full, "utf8"));
            if (!parsed.abi) continue;
            const bytecode = typeof parsed.bytecode === "string" ? parsed.bytecode : parsed.bytecode && parsed.bytecode.object;
            return { abi: parsed.abi, bytecode: bytecode || "0x" };
        }
    }
    throw new Error(`ChaosZero artifact ${target} not found under ${artifactRoots(root).join(", ")}`);
}

// quoteBySearch finds the largest output buy()/sell() accept for an exact
// trade via staticCall: valid outputs form a contiguous [0..max] range under
// the weighted-product invariant, so double until rejected, then bisect.
async function quoteBySearch(tryOutput) {
    const ok = async (out) => {
        if (out <= 0n) return true;
        try {
            await tryOutput(out);
            return true;
        } catch {
            return false;
        }
    };
    let lo = 0n;
    let hi = 1n;
    while (await ok(hi)) {
        lo = hi;
        hi *= 2n;
        if (hi > 10n ** 30n) throw new Error("quote search diverged");
    }
    while (hi - lo > 1n) {
        const mid = (lo + hi) / 2n;
        if (await ok(mid)) lo = mid;
        else hi = mid;
    }
    if (lo === 0n) throw new Error("market did not accept any output for this trade");
    return lo;
}

// --------------------------------------------------------------------------
// Simulation factory
// --------------------------------------------------------------------------
function create(ctx) {
    const { ethers, game, token, TX } = ctx;
    const provider = ethers.provider;
    // Role mapping avoids signers whose CST balance other populate.js steps
    // depend on (addr3 makes CST bids in S4; owner/addr1 move CST in S8/S9
    // only after our claims replenish them).
    const [owner, addr1, addr2, /* addr3 */, addr4, addr5] = ctx.signers;
    const alice = addr4, bob = addr5, carol = addr2, dave = addr1;

    // ChaosZero market txs get a fixed gas limit: estimates are
    // timestamp-sensitive (the mined block is a second later than the
    // estimation block, and champion durations crossing a threshold flip
    // answer bits, adding storage writes and events beyond the estimate).
    const G = { gasLimit: 5_000_000 };

    const st = {
        market: null, source: null, marketAddr: null, sourceAddr: null, vaultAddr: null,
        marketIface: null, vaultIface: null, startBlock: 0,
        r1: 0n, r2: 0n, r4: 0n, rF: 0n,
        valuesR1: 0, aliceQ0Out: 0n, ids: null,
    };

    const expect = (cond, what) => {
        if (!cond) throw new Error("CZ ASSERTION FAILED: " + what);
    };
    const ok = (label, receipt) => console.log("  cz ok " + label + " (block " + receipt.blockNumber + ")");

    const events = (receipt, iface, emitter, name) => {
        const addr = emitter.toLowerCase();
        return receipt.logs
            .filter((l) => l.address.toLowerCase() === addr)
            .map((l) => {
                try { return iface.parseLog({ topics: [...l.topics], data: l.data }); }
                catch { return null; }
            })
            .filter((p) => p && p.name === name);
    };
    const mEvents = (receipt, name) => events(receipt, st.marketIface, st.marketAddr, name);
    const vEvents = (receipt, name) => events(receipt, st.vaultIface, st.vaultAddr, name);

    const deadline = async () => BigInt((await provider.getBlock("latest")).timestamp) + 900n;

    const trade = async (signer, who, round, question, yes, isBuy, input) => {
        const m = st.market.connect(signer);
        const fn = isBuy ? m.getFunction("buy") : m.getFunction("sell");
        const dl = await deadline();
        const output = await quoteBySearch((out) => fn.staticCall(round, question, yes, input, out, dl));
        const receipt = await TX(fn.send(round, question, yes, input, output, await deadline(), G));
        ok(`${who} ${isBuy ? "buy" : "sell"} q${question} ${yes ? "YES" : "NO"} ` +
            `${ethers.formatEther(input)} -> ${ethers.formatEther(output)}`, receipt);
        return { output, receipt };
    };

    const posId = (round, q, yes) => st.market.positionId(round, q, yes);
    const balance = (who, id) => st.market.balanceOf(who, id);
    // outcomeState returns (open, finalized, knownMask, valuesMask).
    const outcome = async (round) => {
        const [, finalized, known, values] = await st.source.outcomeState(round);
        return { finalized, known: Number(known), values: Number(values) };
    };
    const expectedClaim = async (user, round, values) => {
        let total = 0n;
        for (let q = 0; q < 8; q++) {
            total += await balance(user, await posId(round, q, (values & (1 << q)) !== 0));
        }
        return total;
    };
    const claimAndAssert = async (who, signer, round, values) => {
        const user = await signer.getAddress();
        const expected = await expectedClaim(user, round, values);
        const receipt = await TX(st.market.connect(signer).claimKnown(round, G));
        ok(`claimKnown(${round}) by ${who} (expect ${ethers.formatEther(expected)} CST)`, receipt);
        const claimed = mEvents(receipt, "Claimed");
        expect(claimed.length === 1, `one Claimed event for ${who}`);
        expect(claimed[0].args.total === expected, `${who} Claimed.total matches computed entitlement`);
        const paid = vEvents(receipt, "Paid").reduce((s, p) => s + p.args.amount, 0n);
        expect(paid === expected, `${who} vault Paid sum matches entitlement`);
        return receipt;
    };
    const held = async () => {
        const [bal, active, carry] = await Promise.all([
            token.balanceOf(st.marketAddr),
            st.market.totalActiveCollateral(),
            st.market.pendingCarry(),
        ]);
        return bal > active + carry ? bal - active - carry : 0n;
    };

    // The traders earn CST the real way: ETH bids mint the bidder's
    // time-based CST reward. (Impersonating the game to mint would be
    // faster, but hardhat serves impersonated transactions with fake
    // signature values that the backend ETL's TransactionByHash rejects.)
    // Waiting between bids grows the reward, keeping the loop short.
    const earnCst = async (recipients, target) => {
        for (let i = 0; i < 80; i++) {
            let needy = null;
            let needyName = "";
            for (const [name, r] of recipients) {
                if ((await token.balanceOf(await r.getAddress())) < target) {
                    needy = r;
                    needyName = name;
                    break;
                }
            }
            if (!needy) return;
            await ctx.increaseTime(600);
            const bidPrice = await game.getNextEthBidPrice();
            const receipt = await TX(game.connect(needy).bidWithEth(-1, "cz CST top-up bid", 0,
                { value: (bidPrice * 125n) / 100n, gasLimit: 1500000 }));
            ok(`CST top-up bid by ${needyName} (reward accrues)`, receipt);
        }
        throw new Error("could not accumulate enough CST through bids");
    };

    // ----------------------------------------------------------------------
    // phaseA — deploy on the live game; round R1 funding, trades, transfers
    // ----------------------------------------------------------------------
    async function phaseA() {
        const root = chaosZeroDir();
        const gameAddr = await game.getAddress();
        const tokenAddr = await token.getAddress();
        const r1 = await game.roundNum();
        st.r1 = r1;

        // The real previous-round statistics are the market baseline, same
        // as the production deployment procedure.
        const prev = await game.roundStats(r1 - 1n);
        expect(prev.numBids > 0n && prev.totalSpentEthAmount > 0n,
            `round ${r1 - 1n} must be a finished round with ETH bids to serve as baseline`);
        const baseline = {
            gestureCount: prev.numBids,
            ethSpent: prev.totalSpentEthAmount,
            cstBurned: prev.totalSpentCstAmount,
            highestCstPaid: prev.maxCstBidPrice,
            enduranceDuration: prev.enduranceChampionDuration,
            chronoDuration: prev.chronoWarriorDuration,
            sameGesture: (prev.flags & 4n) !== 0n,
        };

        const srcArt = findArtifact(root, "V3OutcomeSource");
        const mktArt = findArtifact(root, "IndependentOutcomeMarket");
        const vltArt = findArtifact(root, "OutcomeSettlementVault");

        const srcFactory = new ethers.ContractFactory(srcArt.abi, srcArt.bytecode, owner);
        const source = await srcFactory.deploy(gameAddr, r1, baseline);
        await source.waitForDeployment();
        const mktFactory = new ethers.ContractFactory(mktArt.abi, mktArt.bytecode, owner);
        const market = await mktFactory.deploy(tokenAddr, await source.getAddress());
        const mktReceipt = await market.deploymentTransaction().wait();

        st.source = source;
        st.market = market;
        st.sourceAddr = await source.getAddress();
        st.marketAddr = await market.getAddress();
        st.vaultAddr = await market.settlementVault();
        st.marketIface = new ethers.Interface(mktArt.abi);
        st.vaultIface = new ethers.Interface(vltArt.abi);
        st.startBlock = mktReceipt.blockNumber;
        console.log(`  cz deployed: source=${st.sourceAddr} market=${st.marketAddr} vault=${st.vaultAddr}` +
            ` (first market round ${r1}, baseline ${prev.numBids} bids)`);

        await earnCst([["alice", alice], ["bob", bob], ["carol", carol]], 150n * ONE);
        for (const [who, s] of [["alice", alice], ["bob", bob], ["carol", carol]]) {
            ok(`approve market (${who})`, await TX(token.connect(s).approve(st.marketAddr, 10n ** 27n)));
        }

        console.log("  cz -- funding variants on round " + r1 + ":");
        ok("fundRound donation 50 CST (alice)", await TX(st.market.connect(alice).fundRound(r1, 50n * ONE, G)));
        // Direct transfer absorbed by a BUY.
        ok("direct transfer 20 CST (bob)", await TX(token.connect(bob).transfer(st.marketAddr, 20n * ONE)));
        {
            const { receipt } = await trade(alice, "alice", r1, 1, true, true, 3n * ONE + 7n);
            const funded = mEvents(receipt, "RoundFunded");
            expect(funded.length === 1, "buy absorbs held CST (RoundFunded in trade tx)");
            expect(funded[0].args.transfersUsed === 20n * ONE, "transfersUsed = 20 CST on absorbing buy");
        }
        // Direct transfer absorbed by a SELL.
        const q6Buy = await trade(bob, "bob", r1, 6, false, true, 4n * ONE);
        ok("direct transfer 15 CST (carol)", await TX(token.connect(carol).transfer(st.marketAddr, 15n * ONE)));
        {
            const { receipt } = await trade(bob, "bob", r1, 6, false, false, q6Buy.output / 2n);
            const funded = mEvents(receipt, "RoundFunded");
            expect(funded.length === 1, "sell absorbs held CST");
            expect(funded[0].args.transfersUsed === 15n * ONE, "transfersUsed = 15 CST on absorbing sell");
        }
        // Sub-threshold held CST: ignored by trades, swept by the
        // zero-amount trigger only once at/over MIN_AUTO_TOP_UP (1 CST).
        ok("direct transfer 0.5 CST (sub-threshold)", await TX(token.connect(alice).transfer(st.marketAddr, ONE / 2n)));
        expect(!(await st.market.fundingPreview(r1, 0n)).canFund, "sub-threshold held CST must not be fundable");
        await trade(carol, "carol", r1, 4, true, true, ONE / 2n);
        expect((await held()) === ONE / 2n, "trade must not absorb sub-threshold held CST");
        ok("direct transfer 0.6 CST (crosses threshold)", await TX(token.connect(alice).transfer(st.marketAddr, ONE / 2n + ONE / 10n)));
        {
            const receipt = await TX(st.market.connect(carol).fundRound(r1, 0n, G));
            ok("fundRound zero-amount trigger (carol)", receipt);
            const funded = mEvents(receipt, "RoundFunded");
            expect(funded[0].args.donated === 0n && funded[0].args.transfersUsed === ONE + ONE / 10n,
                "zero-amount trigger sweeps exactly the held 1.1 CST");
        }

        console.log("  cz -- both-sides purchases (a winner is guaranteed either way):");
        const aliceQ0 = await trade(alice, "alice", r1, 0, true, true, 10n * ONE);
        st.aliceQ0Out = aliceQ0.output;
        await trade(bob, "bob", r1, 0, false, true, 6n * ONE);
        await trade(carol, "carol", r1, 5, true, true, 5n * ONE);
        await trade(bob, "bob", r1, 5, false, true, 2n * ONE);
        // Carol holds BOTH sides of q7: her escrow in phaseB is guaranteed.
        await trade(carol, "carol", r1, 7, true, true, 3n * ONE);
        await trade(carol, "carol", r1, 7, false, true, 3n * ONE);
        // Partial exit (sell a third) and full exit (sell every unit).
        await trade(alice, "alice", r1, 0, true, false, aliceQ0.output / 3n);
        const q4Buy = await trade(bob, "bob", r1, 4, true, true, ONE + ONE / 2n);
        await trade(bob, "bob", r1, 4, true, false, q4Buy.output);

        console.log("  cz -- splits, operator batch transfer, approvals:");
        ok("splitPosition q2 6 CST (bob)", await TX(st.market.connect(bob).splitPosition(r1, 2, 6n * ONE, G)));
        ok("splitPosition q6 2.5 CST (bob)", await TX(st.market.connect(bob).splitPosition(r1, 6, 2n * ONE + ONE / 2n, G)));
        const ids = {
            q2y: await posId(r1, 2, true), q2n: await posId(r1, 2, false),
            q6y: await posId(r1, 6, true), q6n: await posId(r1, 6, false),
        };
        st.ids = ids;
        const addrOf = { alice: await alice.getAddress(), bob: await bob.getAddress(), carol: await carol.getAddress(), dave: await dave.getAddress() };
        ok("setApprovalForAll bob->dave", await TX(st.market.connect(bob).setApprovalForAll(addrOf.dave, true)));
        {
            // Genuine multi-id batch: operator dave ≠ sender bob ≠ recipient
            // carol, four ids across two questions (one id would emit
            // TransferSingle, not TransferBatch).
            const receipt = await TX(st.market.connect(dave).safeBatchTransferFrom(
                addrOf.bob, addrOf.carol,
                [ids.q2y, ids.q2n, ids.q6y, ids.q6n],
                [2n * ONE, 2n * ONE, ONE, ONE], "0x", G));
            ok("safeBatchTransferFrom 4 ids bob->carol (operator dave)", receipt);
            expect(mEvents(receipt, "TransferBatch").length === 1, "a 4-id transfer must emit TransferBatch");
        }
        ok("revoke approval bob->dave", await TX(st.market.connect(bob).setApprovalForAll(addrOf.dave, false)));
        let reverted = false;
        try {
            await TX(st.market.connect(dave).safeTransferFrom(addrOf.bob, addrOf.carol, ids.q2y, ONE, "0x", G));
        } catch {
            reverted = true;
        }
        expect(reverted, "post-revocation transfer by dave must revert");
        console.log("  cz ok post-revocation transfer by dave reverted as expected");
        const q0yId = await posId(r1, 0, true);
        ok("self-transfer (alice->alice)", await TX(st.market.connect(alice).safeTransferFrom(addrOf.alice, addrOf.alice, q0yId, ONE, "0x", G)));
        ok("zero-value transfer (alice->bob)", await TX(st.market.connect(alice).safeTransferFrom(addrOf.alice, addrOf.bob, q0yId, 0n, "0x", G)));

        // Plural multi-question merge before closure (carol merges exactly
        // what the batch delivered).
        ok("mergePositions {q2:2, q6:1} pre-closure (carol)", await TX(
            st.market.connect(carol).mergePositions(r1, [0n, 0n, 2n * ONE, 0n, 0n, 0n, ONE, 0n], G)));
    }

    // ----------------------------------------------------------------------
    // phaseB — R1 closed: merge/settle/claim family; R2 opened by first buy
    // ----------------------------------------------------------------------
    async function phaseB() {
        const r1 = st.r1;
        const r2 = await game.roundNum();
        expect(r2 === r1 + 1n, `round ${r1} must be closed by now`);
        st.r2 = r2;

        // Post-closure, pre-settlement merge: paid from active collateral.
        {
            const receipt = await TX(st.market.connect(bob).mergePosition(r1, 2, 4n * ONE, G));
            ok("mergePosition q2 4 post-closure pre-settlement (bob)", receipt);
            expect(vEvents(receipt, "Paid").length === 0, "pre-settlement merge must not touch the vault");
        }
        // Settlement trigger 1/3: explicit settleRound().
        const settleReceipt = await TX(st.market.settleRound(r1, G));
        ok(`settleRound(${r1}) explicit`, settleReceipt);
        const settled = mEvents(settleReceipt, "RoundSettled");
        expect(settled.length === 1, "explicit settleRound emits RoundSettled");
        // The event field "values" collides with Array.prototype.values on
        // the ethers Result; getValue() bypasses the collision.
        st.valuesR1 = Number(settled[0].args.getValue("values"));
        {
            const o = await outcome(r1);
            expect(o.finalized && o.values === st.valuesR1, "settled values match the source outcome");
        }
        console.log(`  cz round ${r1} settled with values mask ${st.valuesR1} (game-decided)`);
        // Post-settlement merge: the matched pair is paid from the VAULT.
        {
            const receipt = await TX(st.market.connect(bob).mergePosition(r1, 6, ONE + ONE / 2n, G));
            ok("mergePosition q6 1.5 post-settlement (bob)", receipt);
            expect(vEvents(receipt, "Paid").length === 1, "post-settlement merge pays from the vault");
        }
        // Whoever bought the winning q0 side transfers part of it to dave,
        // who then claims: ownership moves AFTER settlement, payout follows.
        const q0Yes = (st.valuesR1 & 1) !== 0;
        const winner = q0Yes ? alice : bob;
        const winnerName = q0Yes ? "alice (q0 YES won)" : "bob (q0 NO won)";
        const winId = await posId(r1, 0, q0Yes);
        const daveAddr = await dave.getAddress();
        ok(`transfer winning q0 position ${winnerName} -> dave`, await TX(
            st.market.connect(winner).safeTransferFrom(await winner.getAddress(), daveAddr, winId, ONE, "0x", G)));
        {
            const expected = await expectedClaim(daveAddr, r1, st.valuesR1);
            expect(expected > 0n, "dave's received winning position must be worth something");
        }
        await claimAndAssert("dave (received winner)", dave, r1, st.valuesR1);
        await claimAndAssert("alice", alice, r1, st.valuesR1);
        await claimAndAssert("bob", bob, r1, st.valuesR1);
        {
            // owner never traded: a guaranteed zero-total claim, distinct
            // from the paying ones above.
            const expected = await expectedClaim(await owner.getAddress(), r1, st.valuesR1);
            expect(expected === 0n, "owner's entitlement is zero by construction");
            await claimAndAssert("owner (intentional zero claim)", owner, r1, st.valuesR1);
        }
        // Carol does NOT claim R1 (she holds both q7 sides, so she is owed
        // something whatever the outcome): escrow persists in the vault.
        expect((await expectedClaim(await carol.getAddress(), r1, st.valuesR1)) > 0n,
            "carol leaves a guaranteed claim escrowed");

        // Round R2: opened by its FIRST BUY, funded purely from carry.
        console.log("  cz -- round " + r2 + " opened by its first buy:");
        {
            const { receipt } = await trade(alice, "alice", r2, 3, true, true, ONE);
            const funded = mEvents(receipt, "RoundFunded");
            expect(funded.length === 1 && funded[0].args.carryUsed > 0n,
                "first buy initializes and funds the round from carry");
            expect(mEvents(receipt, "RoundInitialized").length === 1, "first buy emits RoundInitialized");
        }
        ok("splitPosition q1 1 CST (bob)", await TX(st.market.connect(bob).splitPosition(r2, 1, ONE, G)));
        // Both sides of q0 for round R2 (claimed in phaseD).
        await trade(carol, "carol", r2, 0, false, true, 4n * ONE);
        await trade(alice, "alice", r2, 0, true, true, 3n * ONE);
    }

    // ----------------------------------------------------------------------
    // phaseC — R2 closed but deliberately left UNSETTLED
    // ----------------------------------------------------------------------
    async function phaseC() {
        const r2 = st.r2;
        expect((await game.roundNum()) > r2, `round ${r2} must be closed by now`);
        expect((await outcome(r2)).finalized, `round ${r2} outcome must be finalized`);
        // Post-closure merge on an unsettled round (paid from collateral);
        // settlement itself is left for phaseD's fundRound to trigger.
        {
            const receipt = await TX(st.market.connect(bob).mergePosition(r2, 1, ONE / 2n, G));
            ok("mergePosition q1 0.5 post-closure on UNSETTLED round (bob)", receipt);
            expect(vEvents(receipt, "Paid").length === 0, "merge on the unsettled round must not touch the vault");
        }
    }

    // ----------------------------------------------------------------------
    // phaseD — R3 skipped; fundRound(R4) settles R2 across the gap
    // ----------------------------------------------------------------------
    async function phaseD() {
        const r2 = st.r2;
        const r4 = await game.roundNum();
        expect(r4 === r2 + 2n, `exactly one game round (${r2 + 1n}) must separate ${r2} and ${r4} with no market activity`);
        st.r4 = r4;

        // Combined funding: donation + held direct transfer + carry, all in
        // ONE RoundFunded — and the same tx must settle R2 (the last funded
        // round), proving the skipped round is harmless.
        ok("direct transfer 2.5 CST (carol)", await TX(token.connect(carol).transfer(st.marketAddr, 2n * ONE + ONE / 2n)));
        {
            const receipt = await TX(st.market.connect(carol).fundRound(r4, 5n * ONE, G));
            ok(`fundRound(${r4}, 5 CST) (carol)`, receipt);
            const settled = mEvents(receipt, "RoundSettled");
            expect(settled.length === 1 && settled[0].args.roundId === r2,
                `funding round ${r4} settles round ${r2} across the skipped round`);
            const f = mEvents(receipt, "RoundFunded")[0].args;
            expect(f.donated === 5n * ONE && f.transfersUsed === 2n * ONE + ONE / 2n && f.carryUsed > 0n,
                "one RoundFunded with donated, transfersUsed and carryUsed all nonzero");
            const total = f.donated + f.transfersUsed + f.carryUsed;
            const sum = f.allocations.reduce((s, a) => s + a, 0n);
            expect(sum === total, "allocations sum to the funding total");
            const sorted = [...f.allocations].sort((a, b) => (a < b ? -1 : 1));
            expect(sorted[7] - sorted[0] <= 1n, "base-unit remainder spreads across pools by at most 1 wei");
            expect(mEvents(receipt, "RoundInitialized").length === 1, `round ${r4} initialized by the funding`);
        }

        // R2 claims against the now-settled values.
        const values2 = (await outcome(r2)).values;
        console.log(`  cz round ${r2} settled with values mask ${values2} (game-decided)`);
        await claimAndAssert("alice", alice, r2, values2);
        await claimAndAssert("carol", carol, r2, values2);
        {
            // bob split q1 and merged back only half: he holds 0.5 CST of
            // BOTH q1 sides, so his claim is guaranteed to pay.
            const expected = await expectedClaim(await bob.getAddress(), r2, values2);
            expect(expected > 0n, "bob's split remainder guarantees a paying claim");
            await claimAndAssert("bob (split remainder)", bob, r2, values2);
        }

        // Round R4 positions for phaseE: both q0 sides, plus carol holding
        // both q6 sides (her escrow is guaranteed again).
        console.log("  cz -- round " + r4 + " purchases:");
        await trade(alice, "alice", r4, 0, true, true, 3n * ONE);
        await trade(bob, "bob", r4, 0, false, true, 2n * ONE);
        await trade(carol, "carol", r4, 6, true, true, ONE + ONE / 2n);
        await trade(carol, "carol", r4, 6, false, true, ONE + ONE / 2n);
    }

    // ----------------------------------------------------------------------
    // phaseE — claim-triggered settlement of R4; final tally
    // ----------------------------------------------------------------------
    async function phaseE() {
        const r4 = st.r4;
        expect((await game.roundNum()) > r4, `round ${r4} must be closed by now`);
        const o = await outcome(r4);
        expect(o.finalized, `round ${r4} outcome must be finalized`);
        console.log(`  cz round ${r4} finalized with values mask ${o.values} (game-decided)`);

        // Settlement trigger 2/3 of the remaining paths: claimKnown() on the
        // closed, unsettled round settles it in the same tx.
        const q0Yes = (o.values & 1) !== 0;
        const firstClaimer = q0Yes ? alice : bob;
        const receipt = await claimAndAssert(q0Yes ? "alice (q0 YES won, claim settles)" : "bob (q0 NO won, claim settles)",
            firstClaimer, r4, o.values);
        expect(mEvents(receipt, "RoundSettled").length === 1, "the claim triggered the settlement");
        expect(vEvents(receipt, "CycleEscrowed").length === 1, "the claim-settlement escrows the cycle");
        await claimAndAssert(q0Yes ? "bob (losing side)" : "alice (losing side)", q0Yes ? bob : alice, r4, o.values);
        // Carol holds both q6 sides and does not claim: escrow persists.
        expect((await expectedClaim(await carol.getAddress(), r4, o.values)) > 0n,
            "carol leaves a guaranteed claim escrowed on the final round");
    }

    // ----------------------------------------------------------------------
    // phaseF — open the next round and leave it rich and UNSETTLED
    // ----------------------------------------------------------------------
    async function phaseF() {
        // The caller advanced past the activation delay; the first ETH bid
        // starts the new game round (the market follows the game's rounds).
        const bidPrice = await game.getNextEthBidPrice();
        ok("first bid of the open round (alice)", await TX(game.connect(alice).bidWithEth(-1, "cz open-round bid", 0,
            { value: (bidPrice * 125n) / 100n, gasLimit: 1500000 })));
        const rF = await game.roundNum();
        expect(rF > st.r4, `game round ${rF} is newer than the last settled market round ${st.r4}`);
        st.rF = rF;
        await earnCst([["alice", alice], ["bob", bob], ["carol", carol]], 60n * ONE);

        // The initial funding of the open round absorbs the carry the
        // settled history rolled forward.
        {
            const receipt = await TX(st.market.connect(alice).fundRound(rF, 5n * ONE, G));
            const funded = mEvents(receipt, "RoundFunded");
            expect(funded.length === 1, "explicit fundRound initializes the open round");
            expect(funded[0].args.carryUsed > 0n, "the open round absorbed the settled history's carry");
            ok(`fundRound(${rF}) donation 5 CST + carry (alice)`, receipt);
        }

        console.log("  cz -- open-round activity (deliberately left unsettled):");
        // Both sides of q0: a two-sided trade history for the odds pages.
        await trade(alice, "alice", rF, 0, true, true, 6n * ONE);
        await trade(bob, "bob", rF, 0, false, true, 4n * ONE);
        await trade(carol, "carol", rF, 2, true, true, 5n * ONE);
        await trade(bob, "bob", rF, 3, false, true, 3n * ONE);
        // An early exit at market price while the round is open.
        const q5 = await trade(alice, "alice", rF, 5, true, true, 2n * ONE);
        await trade(alice, "alice", rF, 5, true, false, q5.output / 2n);
        // A direct transfer absorbed by the next buy (open-round top-up).
        ok("direct transfer 6 CST (bob)", await TX(token.connect(bob).transfer(st.marketAddr, 6n * ONE)));
        {
            const { receipt } = await trade(carol, "carol", rF, 4, false, true, 2n * ONE);
            const funded = mEvents(receipt, "RoundFunded");
            expect(funded.length === 1 && funded[0].args.transfersUsed === 6n * ONE,
                "the buy absorbed the held transfer into the open round");
        }
        // A split and a partial merge-back, both while the round is open.
        ok("splitPosition q7 2 CST (carol)", await TX(st.market.connect(carol).splitPosition(rF, 7, 2n * ONE, G)));
        ok("mergePositions {q7:1} while open (carol)", await TX(
            st.market.connect(carol).mergePositions(rF, [0n, 0n, 0n, 0n, 0n, 0n, 0n, ONE], G)));
        // A gifted position on the open round (zero-cost holder).
        ok("gift 1.0 q0 YES unit alice->dave", await TX(st.market.connect(alice).safeTransferFrom(
            await alice.getAddress(), await dave.getAddress(), await posId(rF, 0, true), ONE, "0x", G)));

        const o = await outcome(rF);
        expect(!o.finalized, `round ${rF} must still be open`);
        const active = await st.market.totalActiveCollateral();
        expect(active > 0n, "the open round holds live liquidity");
        console.log(`  cz round ${rF} left OPEN with ${ethers.formatEther(active)} CST active collateral`);

        await tally();
    }

    // ----------------------------------------------------------------------
    // Final event tally + structural checks + cz_contracts SQL
    // ----------------------------------------------------------------------
    const REQUIRED_EVENTS = {
        RoundInitialized: { src: "market", min: 4 }, // R1, R2, R4, RF (R3 skipped)
        RoundFunded: { src: "market", min: 8 },      // +2 on the open round (explicit, absorb)
        Trade: { src: "market", min: 21 },           // +6 on the open round
        AnswersObserved: { src: "market", min: 2 },
        PositionSplit: { src: "market", min: 4 },    // +1 on the open round
        PositionsMerged: { src: "market", min: 5 },  // +1 on the open round
        RoundSettled: { src: "market", min: 3 },     // RF stays unsettled (structural check)
        Claimed: { src: "market", min: 8 },
        TransferSingle: { src: "market", min: 15 },  // +7 on the open round (6 trades + gift)
        TransferBatch: { src: "market", min: 6 },    // +2 on the open round (split, merge)
        ApprovalForAll: { src: "market", min: 2 },
        CycleEscrowed: { src: "vault", min: 3 },
        Paid: { src: "vault", min: 6 },
    };

    async function tally() {
        console.log("\n  cz event tally:");
        const logs = await provider.getLogs({
            address: [st.marketAddr, st.vaultAddr],
            fromBlock: st.startBlock,
            toBlock: "latest",
        });
        const counts = {};
        const parsed = [];
        for (const log of logs) {
            const src = log.address.toLowerCase() === st.marketAddr.toLowerCase() ? "market" : "vault";
            const p = (src === "market" ? st.marketIface : st.vaultIface).parseLog({ topics: [...log.topics], data: log.data });
            if (!p) continue;
            counts[p.name] = (counts[p.name] || 0) + 1;
            parsed.push({ src, name: p.name, args: p.args });
        }
        let failures = 0;
        for (const [name, req] of Object.entries(REQUIRED_EVENTS)) {
            const n = counts[name] || 0;
            const good = n >= req.min;
            console.log(`  ${good ? "OK  " : "MISS"} ${req.src.padEnd(6)} ${name.padEnd(18)} ${n} (need >= ${req.min})`);
            if (!good) failures++;
        }
        const structural = [
            ["RoundFunded with direct transfers absorbed",
                parsed.some((p) => p.name === "RoundFunded" && p.args.transfersUsed > 0n)],
            ["RoundFunded with donation+transfers+carry all nonzero",
                parsed.some((p) => p.name === "RoundFunded" && p.args.donated > 0n && p.args.transfersUsed > 0n && p.args.carryUsed > 0n)],
            ["user-to-user TransferBatch with 4 ids",
                parsed.some((p) => p.name === "TransferBatch" && p.args.ids.length >= 4
                    && p.args.from !== ethers.ZeroAddress && p.args.to !== ethers.ZeroAddress)],
            ["paying Claimed on at least two rounds",
                new Set(parsed.filter((p) => p.name === "Claimed" && p.args.total > 0n).map((p) => p.args.roundId)).size >= 2],
            ["final round left open: initialized and funded, never settled",
                st.rF > 0n
                && parsed.some((p) => p.name === "RoundInitialized" && p.args.roundId === st.rF)
                && parsed.some((p) => p.name === "RoundFunded" && p.args.roundId === st.rF)
                && !parsed.some((p) => p.name === "RoundSettled" && p.args.roundId === st.rF)],
            ["open round holds live liquidity (totalActiveCollateral > 0)",
                (await st.market.totalActiveCollateral()) > 0n],
        ];
        for (const [what, good] of structural) {
            console.log(`  ${good ? "OK  " : "MISS"} struct ${what}`);
            if (!good) failures++;
        }

        console.log("\n-- ChaosZero contracts (same game, same rounds as cg_ tables):");
        console.log("-- SQL UPDATE for cz_contracts / cz_proc_status before starting cz-etl.");
        console.log("-- last_block is the block before the market deployment: a zero would");
        console.log("-- make cz-etl fall back to the shared block watermark, which skips");
        console.log("-- this history when the cosmicgame ETL has already indexed past it.");
        console.log("UPDATE cz_contracts SET");
        console.log(`\tmarket_addr = '${st.marketAddr}',`);
        console.log(`\tsource_addr = '${st.sourceAddr}',`);
        console.log(`\tvault_addr  = '${st.vaultAddr}';`);
        console.log(`UPDATE cz_proc_status SET last_evt_id = 0, last_block = ${st.startBlock - 1};`);

        if (failures > 0) {
            throw new Error(`${failures} required ChaosZero event type(s) or scenario(s) missing`);
        }
        console.log(`\n  cz OK: all ${Object.keys(REQUIRED_EVENTS).length} event types and ` +
            `${structural.length} structural scenarios covered on game rounds ${st.r1}, ${st.r2}, ${st.r4}` +
            ` (round ${st.r2 + 1n} skipped); round ${st.rF} left open with live data.`);
    }

    return { phaseA, phaseB, phaseC, phaseD, phaseE, phaseF };
}

module.exports = { enabled, create };
