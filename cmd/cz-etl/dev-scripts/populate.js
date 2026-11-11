// ChaosZero sample-data generator for a local development chain (hardhat
// node or anvil, chain id 31337). Mirrors the spirit of
// cmd/cgctl/dev-scripts/populate.js: drive the market through several full
// cycles so every event cz-etl indexes appears in every variant the
// deployed contract can produce, assert the intended outcome of each step
// in flight (receipts + contract reads), and FAIL (exit 1) if any required
// event type or structural scenario is missing from the final log tally.
//
// The scenario spans four game cycles on top of the deploy-local.js seed
// (game rounds A..E, where D is deliberately skipped by the market):
//
//   Cycle A  donation funding; direct-transfer funding absorbed by a buy
//            AND by a sell; sub-threshold held CST ignored by trades, then
//            swept by the zero-amount fundRound trigger; gestures (ETH and
//            CST) driven until question 0 is mathematically known YES
//            against the real previous-cycle baseline; purchased winners
//            on the YES side; partial and full exits (sell YES + sell NO);
//            two splits; a genuine 4-id operator batch transfer between
//            users (operator, sender, recipient all distinct); transfer
//            after revocation asserted to revert; self- and zero-value
//            transfers; multi-question mergePositions before closure;
//            merge after closure but before settlement; explicit
//            settleRound(); post-settlement transfer of a winning
//            position to a fresh wallet that then claims; post-settlement
//            merge paid from the vault; claims asserted against computed
//            entitlements (including an intentional zero-total claim);
//            one winner left unclaimed so escrow persists.
//   Cycle B  combined funding: donation + held direct transfer + carry all
//            nonzero in ONE RoundFunded; two gestures so question 0
//            resolves NO and parity resolves NO; purchased winners on the
//            NO side; settlement triggered by claimKnown() alone.
//   Cycle C  opened by its first buy (carry-only rollover, no fundRound);
//            split-only user activity afterwards; left to settle later.
//   Cycle D  closed by the game but never funded by the market: the round
//            id gap proves skipped game rounds are harmless.
//   Cycle E  funded by fundRound(), which settles cycle C (the third
//            settlement trigger) across the skipped round.
//
// Usage (after deploy-local.js, against the same chain):
//   CHAOSZERO_DIR=/path/to/ChaosZero node populate.js
//
// Environment: RPC_URL (default http://127.0.0.1:8545),
// CZ_DEPLOYMENT_FILE (default ./cz-local-deployment.json).
"use strict";

const {
  ONE,
  chaosZeroDir,
  findArtifact,
  connect,
  readDeployment,
  quoteBySearch,
  rpc,
  ethers,
} = require("./lib.js");

// REQUIRED_EVENTS maps every event cz-etl indexes to its contract and the
// minimum count this scenario must produce.
const REQUIRED_EVENTS = {
  RoundInitialized: { src: "market", min: 4 }, // cycles A, B, C, E (D skipped)
  RoundFunded: { src: "market", min: 7 },
  Trade: { src: "market", min: 12 },
  AnswersObserved: { src: "market", min: 2 },
  PositionSplit: { src: "market", min: 3 },
  PositionsMerged: { src: "market", min: 3 },
  RoundSettled: { src: "market", min: 3 },
  Claimed: { src: "market", min: 4 },
  TransferSingle: { src: "market", min: 10 },
  TransferBatch: { src: "market", min: 4 },
  ApprovalForAll: { src: "market", min: 2 },
  CycleEscrowed: { src: "vault", min: 3 },
  Paid: { src: "vault", min: 4 },
};

function expect(cond, what) {
  if (!cond) throw new Error(`ASSERTION FAILED: ${what}`);
}

async function main() {
  const root = chaosZeroDir();
  const dep = readDeployment();
  const { provider, signers } = await connect();
  const [deployer, alice, bob, carol, dave] = signers;
  const addr = {};
  for (const [name, s] of Object.entries({ deployer, alice, bob, carol, dave })) {
    addr[name] = await s.getAddress();
  }

  const marketArt = findArtifact(root, "IndependentOutcomeMarket");
  const vaultArt = findArtifact(root, "OutcomeSettlementVault");
  const gameArt = findArtifact(root, "LocalCosmicGame");
  const cstArt = findArtifact(root, "LocalCst");
  const sourceArt = findArtifact(root, "LocalOutcomeSource");

  const market = new ethers.Contract(dep.marketAddress, marketArt.abi, deployer);
  const game = new ethers.Contract(dep.gameAddress, gameArt.abi, deployer);
  const cst = new ethers.Contract(dep.tokenAddress, cstArt.abi, deployer);
  const source = new ethers.Contract(dep.sourceAddress, sourceArt.abi, deployer);
  const marketIface = new ethers.Interface(marketArt.abi);
  const vaultIface = new ethers.Interface(vaultArt.abi);

  const send = async (label, promise) => {
    const receipt = await (await promise).wait();
    expect(receipt.status === 1, `${label} must succeed`);
    console.log(`  ok ${label} (block ${receipt.blockNumber})`);
    return receipt;
  };
  const expectRevert = async (label, signer, promise) => {
    try {
      await (await promise).wait();
    } catch {
      // The NonceManager may have consumed a local nonce for the failed tx.
      signer.reset();
      console.log(`  ok ${label} reverted as expected`);
      return;
    }
    throw new Error(`ASSERTION FAILED: ${label} should have reverted`);
  };
  const deadline = async () => BigInt((await provider.getBlock("latest")).timestamp) + 900n;
  const events = (receipt, iface, name) =>
    receipt.logs
      .map((l) => {
        try {
          return iface.parseLog({ topics: [...l.topics], data: l.data });
        } catch {
          return null;
        }
      })
      .filter((p) => p && p.name === name);

  // An exact-pair trade: probe the largest output the invariant check
  // accepts via staticCall, then execute with that quote.
  const trade = async (signer, round, question, yes, isBuy, input) => {
    const m = market.connect(signer);
    const dl = await deadline();
    const fn = isBuy ? m.getFunction("buy") : m.getFunction("sell");
    const output = await quoteBySearch((out) =>
      fn.staticCall(round, question, yes, input, out, dl),
    );
    const receipt = await send(
      `${isBuy ? "buy" : "sell"} q${question} ${yes ? "YES" : "NO"} ` +
        `${ethers.formatEther(input)} -> ${ethers.formatEther(output)}`,
      fn.send(round, question, yes, input, output, await deadline()),
    );
    return { output, receipt };
  };

  const posId = (round, q, yes) => market.positionId(round, q, yes);
  const balance = (who, id) => market.balanceOf(who, id);
  // outcomeState returns (open, finalized, knownMask, valuesMask).
  const outcome = async (round) => {
    const [, finalized, known, values] = await source.outcomeState(round);
    return { finalized, known: Number(known), values: Number(values) };
  };
  // What claimKnown must pay a user: their balance on each winning side.
  const expectedClaim = async (user, round, values) => {
    let total = 0n;
    for (let q = 0; q < 8; q++) {
      total += await balance(user, await posId(round, q, (values & (1 << q)) !== 0));
    }
    return total;
  };
  const claimAndAssert = async (name, signer, round, values) => {
    const user = await signer.getAddress();
    const expected = await expectedClaim(user, round, values);
    const receipt = await send(
      `claimKnown(${round}) by ${name} (expect ${ethers.formatEther(expected)} CST)`,
      market.connect(signer).claimKnown(round),
    );
    const claimed = events(receipt, marketIface, "Claimed");
    expect(claimed.length === 1, `one Claimed event for ${name}`);
    expect(claimed[0].args.total === expected, `${name} Claimed.total matches entitlement`);
    const paidTotal = events(receipt, vaultIface, "Paid").reduce((s, p) => s + p.args.amount, 0n);
    expect(paidTotal === expected, `${name} vault Paid sum matches entitlement`);
    return receipt;
  };
  // Gas estimates for game calls are timestamp-sensitive (the mined block is
  // a second later than the estimation block, which can flip time-based
  // questions to "known" and add storage writes), so use a fixed limit.
  const GAME_GAS = { gasLimit: 2_000_000n };
  const gestureEth = async (signer) => {
    const price = await game.getNextEthPrice();
    await send(`bidEth ${ethers.formatEther(price)} ETH`, game.connect(signer).bidEth({ value: price, ...GAME_GAS }));
  };
  const gestureCst = async (signer) => {
    const price = await game.getNextCstPrice();
    await send(`bidCst <= ${ethers.formatEther(price)} CST`, game.connect(signer).bidCst(price, GAME_GAS));
  };
  const closeCycle = async (signer) => {
    const at = (await game.mainPrizeTime()) + 1n;
    await rpc(provider, "evm_setNextBlockTimestamp", [Number(at)]);
    await send("closeCycle", game.connect(signer).closeCycle(GAME_GAS));
  };
  const held = async () => {
    const [bal, active, carry] = await Promise.all([
      cst.balanceOf(dep.marketAddress),
      market.totalActiveCollateral(),
      market.pendingCarry(),
    ]);
    return bal > active + carry ? bal - active - carry : 0n;
  };

  const roundA = await game.roundNum();
  console.log(`Populating from round ${roundA} on ${dep.rpcUrl}`);

  // ========================================================================
  // CYCLE A
  // ========================================================================
  console.log("\n[cycle A] fund wallets:");
  for (const [name, s] of [["alice", alice], ["bob", bob], ["carol", carol]]) {
    await send(`faucet ${name}`, cst.connect(s).faucet(50000n * ONE));
    await send(`approve market (${name})`, cst.connect(s).approve(dep.marketAddress, 10n ** 27n));
  }

  console.log("[cycle A] funding variants:");
  // Explicit donation.
  await send("fundRound donation +500 CST", market.connect(alice).fundRound(roundA, 500n * ONE));
  // Direct transfer absorbed by a BUY: RoundFunded.transfersUsed must be 200.
  await send("direct transfer 200 CST", cst.connect(bob).transfer(dep.marketAddress, 200n * ONE));
  {
    const { receipt } = await trade(alice, roundA, 1, true, true, 30n * ONE + 7n);
    const funded = events(receipt, marketIface, "RoundFunded");
    expect(funded.length === 1, "buy absorbs held CST (RoundFunded in trade tx)");
    expect(funded[0].args.transfersUsed === 200n * ONE, "transfersUsed = 200 CST on absorbing buy");
  }
  // Direct transfer absorbed by a SELL.
  const q6Buy = await trade(bob, roundA, 6, false, true, 40n * ONE);
  await send("direct transfer 150 CST", cst.connect(carol).transfer(dep.marketAddress, 150n * ONE));
  {
    const { receipt } = await trade(bob, roundA, 6, false, false, q6Buy.output / 2n); // sell NO, partial exit
    const funded = events(receipt, marketIface, "RoundFunded");
    expect(funded.length === 1, "sell absorbs held CST");
    expect(funded[0].args.transfersUsed === 150n * ONE, "transfersUsed = 150 CST on absorbing sell");
  }
  // Below MIN_AUTO_TOP_UP (1 CST): trades leave it alone; a zero-amount
  // fundRound trigger needs the held total back at/over the threshold.
  await send("direct transfer 0.5 CST (sub-threshold)", cst.connect(alice).transfer(dep.marketAddress, ONE / 2n));
  expect(!(await market.fundingPreview(roundA, 0n)).canFund, "sub-threshold held CST must not be fundable");
  await trade(carol, roundA, 4, true, true, 5n * ONE);
  expect((await held()) === ONE / 2n, "trade must not absorb sub-threshold held CST");
  await send("direct transfer 0.6 CST (crosses threshold)", cst.connect(alice).transfer(dep.marketAddress, ONE / 2n + ONE / 10n));
  {
    const receipt = await send("fundRound zero-amount trigger", market.connect(carol).fundRound(roundA, 0n));
    const funded = events(receipt, marketIface, "RoundFunded");
    expect(funded[0].args.donated === 0n && funded[0].args.transfersUsed === ONE + ONE / 10n,
      "zero-amount trigger sweeps exactly the held 1.1 CST");
  }

  // Gestures until question 0 is mathematically known YES against the REAL
  // previous-cycle baseline (the local game seeds round 0 with 6 gestures;
  // on re-runs the baseline is whatever the last cycle produced).
  const prevStats = await game.roundStats(roundA - 1n);
  const prevG = prevStats.gestureCount;
  expect(prevG + 1n <= 40n, `baseline gesture count ${prevG} too large for a sane scenario`);
  console.log(`[cycle A] gestures: previous cycle had ${prevG}; making ${prevG + 1n} (ETH + CST):`);
  await gestureEth(alice); // first gesture of a cycle must be ETH
  for (let i = 1n; i < prevG; i++) await gestureEth(bob);
  await gestureCst(alice); // CST gesture; also the (prevG+1)-th and last
  {
    const { known, values } = await outcome(roundA);
    expect((known & 1) !== 0, "question 0 must be mathematically known after exceeding the baseline");
    expect((values & 1) !== 0, "question 0 must be known YES");
  }
  const parityA = Number((prevG + 1n) % 2n); // 1 = odd = q5 YES at closure

  console.log("[cycle A] trades with purchased winners (YES side):");
  const aliceQ0 = await trade(alice, roundA, 0, true, true, 100n * ONE); // winner, held
  await trade(bob, roundA, 0, false, true, 60n * ONE); // loser
  await trade(carol, roundA, 5, parityA === 1, true, 50n * ONE); // winner (parity known in advance)
  await trade(bob, roundA, 5, parityA !== 1, true, 20n * ONE); // loser
  await trade(alice, roundA, 0, true, false, aliceQ0.output / 3n); // partial exit: sell YES
  const q4Buy = await trade(bob, roundA, 4, true, true, 15n * ONE + 3n);
  await trade(bob, roundA, 4, true, false, q4Buy.output); // full exit: sell every unit

  console.log("[cycle A] splits, batch transfer, approvals:");
  await send("splitPosition q2 60 CST", market.connect(bob).splitPosition(roundA, 2, 60n * ONE));
  await send("splitPosition q6 25 CST", market.connect(bob).splitPosition(roundA, 6, 25n * ONE));
  const ids = {
    q2y: await posId(roundA, 2, true), q2n: await posId(roundA, 2, false),
    q6y: await posId(roundA, 6, true), q6n: await posId(roundA, 6, false),
  };
  await send("setApprovalForAll bob->dave", market.connect(bob).setApprovalForAll(addr.dave, true));
  {
    // Genuine multi-id batch between users: operator dave, sender bob,
    // recipient carol — all distinct; four ids across two questions.
    const receipt = await send(
      "safeBatchTransferFrom 4 ids bob->carol (operator dave)",
      market.connect(dave).safeBatchTransferFrom(
        addr.bob, addr.carol,
        [ids.q2y, ids.q2n, ids.q6y, ids.q6n],
        [20n * ONE, 20n * ONE, 10n * ONE, 10n * ONE], "0x"),
    );
    expect(events(receipt, marketIface, "TransferBatch").length === 1,
      "a 4-id transfer must emit TransferBatch (one id would emit TransferSingle)");
  }
  await send("revoke approval bob->dave", market.connect(bob).setApprovalForAll(addr.dave, false));
  await expectRevert("post-revocation transfer by dave", dave,
    market.connect(dave).safeTransferFrom(addr.bob, addr.carol, ids.q2y, ONE, "0x"));
  await send("self-transfer (alice->alice)",
    market.connect(alice).safeTransferFrom(addr.alice, addr.alice, await posId(roundA, 0, true), ONE, "0x"));
  await send("zero-value transfer (alice->bob)",
    market.connect(alice).safeTransferFrom(addr.alice, addr.bob, await posId(roundA, 0, true), 0n, "0x"));

  console.log("[cycle A] merges around closure and settlement:");
  // Multi-question plural merge before closure (carol merges what she received).
  await send("mergePositions {q2:20, q6:10} pre-closure",
    market.connect(carol).mergePositions(roundA, [0n, 0n, 20n * ONE, 0n, 0n, 0n, 10n * ONE, 0n]));
  await closeCycle(alice); // alice made the last gesture
  // Merge after closure, before settlement: paid from active collateral.
  await send("mergePosition q2 40 post-closure pre-settlement",
    market.connect(bob).mergePosition(roundA, 2, 40n * ONE));
  // Settlement trigger 1/3: explicit settleRound().
  const settleReceipt = await send("settleRound (explicit)", market.settleRound(roundA));
  const settledA = events(settleReceipt, marketIface, "RoundSettled");
  expect(settledA.length === 1, "explicit settleRound emits RoundSettled");
  // The field is named "values", which collides with Array.prototype.values
  // on the ethers Result object; getValue() bypasses the collision.
  const valuesA = Number(settledA[0].args.getValue("values"));
  expect((valuesA & 1) !== 0, "cycle A question 0 settles YES");
  expect(((valuesA >> 5) & 1) === parityA, "cycle A parity settles as planned");
  // Merge after settlement: the matched pair is paid from the VAULT.
  {
    const receipt = await send("mergePosition q6 15 post-settlement (vault-paid)",
      market.connect(bob).mergePosition(roundA, 6, 15n * ONE));
    expect(events(receipt, vaultIface, "Paid").length === 1, "post-settlement merge pays from the vault");
  }
  // Transfer a settled-round winning position to a fresh wallet; it claims.
  const q0yId = await posId(roundA, 0, true);
  await send("transfer winning q0 YES alice->dave (post-settlement)",
    market.connect(alice).safeTransferFrom(addr.alice, addr.dave, q0yId, 30n * ONE, "0x"));
  await claimAndAssert("dave (received winner)", dave, roundA, valuesA);
  await claimAndAssert("alice (purchased winner)", alice, roundA, valuesA);
  {
    // Bob holds only losing sides by now: his claim must total exactly zero.
    const expected = await expectedClaim(addr.bob, roundA, valuesA);
    expect(expected === 0n, "bob's cycle A entitlement is zero by construction");
    await claimAndAssert("bob (intentional zero claim)", bob, roundA, valuesA);
  }
  // Carol does NOT claim cycle A: her q5 winnings stay escrowed in the
  // vault through the following cycles.
  expect((await expectedClaim(addr.carol, roundA, valuesA)) > 0n, "carol leaves a claim escrowed");

  // ========================================================================
  // CYCLE B
  // ========================================================================
  const roundB = await game.roundNum();
  expect(roundB === roundA + 1n, "cycle B follows A");
  console.log(`\n[cycle B] round ${roundB}: combined funding (donation + transfers + carry):`);
  await send("direct transfer 25 CST", cst.connect(carol).transfer(dep.marketAddress, 25n * ONE));
  {
    const receipt = await send("fundRound donation 50 CST", market.fundRound(roundB, 50n * ONE));
    const f = events(receipt, marketIface, "RoundFunded")[0].args;
    expect(f.donated === 50n * ONE && f.transfersUsed === 25n * ONE && f.carryUsed > 0n,
      "one RoundFunded with donated, transfersUsed and carryUsed all nonzero");
    const total = f.donated + f.transfersUsed + f.carryUsed;
    const sum = f.allocations.reduce((s, a) => s + a, 0n);
    expect(sum === total, "allocations sum to total funding");
    const sorted = [...f.allocations].sort((a, b) => (a < b ? -1 : 1));
    expect(sorted[7] - sorted[0] <= 1n, "base-unit remainder spreads across pools by at most 1 wei");
  }
  console.log("[cycle B] two gestures => q0 NO (fewer than cycle A), parity NO (even):");
  await gestureEth(bob);
  await gestureEth(bob);
  console.log("[cycle B] purchased winners on the NO side:");
  await trade(carol, roundB, 0, false, true, 40n * ONE); // NO winner
  await trade(alice, roundB, 0, true, true, 30n * ONE); // loser
  await trade(bob, roundB, 5, false, true, 20n * ONE + 1n); // NO winner (even parity)
  await closeCycle(bob);
  {
    const { finalized, values } = await outcome(roundB);
    expect(finalized, "cycle B finalized after closeCycle");
    expect((values & 1) === 0, "cycle B question 0 resolves NO");
    expect(((values >> 5) & 1) === 0, "cycle B parity resolves NO");
    // Settlement trigger 2/3: claimKnown on the closed, unsettled round.
    const receipt = await claimAndAssert("carol (claim settles the round)", carol, roundB, values);
    expect(events(receipt, marketIface, "RoundSettled").length === 1, "claim triggered the settlement");
    expect(events(receipt, vaultIface, "CycleEscrowed").length === 1, "claim-settlement escrows the cycle");
    await claimAndAssert("bob (NO-side winner)", bob, roundB, values);
  }

  // ========================================================================
  // CYCLE C — opened by its first buy; split-only activity afterwards
  // ========================================================================
  const roundC = await game.roundNum();
  console.log(`\n[cycle C] round ${roundC}: opened by the first buy (carry rollover, no fundRound):`);
  {
    const { receipt } = await trade(alice, roundC, 3, true, true, 10n * ONE);
    const funded = events(receipt, marketIface, "RoundFunded");
    expect(funded.length === 1 && funded[0].args.carryUsed > 0n,
      "first buy initializes and funds the cycle from carry");
    expect(events(receipt, marketIface, "RoundInitialized").length === 1,
      "first buy emits RoundInitialized");
  }
  await send("splitPosition q1 10 CST (split-only user flow)",
    market.connect(bob).splitPosition(roundC, 1, 10n * ONE));
  await gestureEth(carol);
  await closeCycle(carol);

  // ========================================================================
  // CYCLE D — skipped by the market entirely
  // ========================================================================
  const roundD = await game.roundNum();
  console.log(`\n[cycle D] round ${roundD}: game-only cycle, never funded by the market:`);
  await gestureEth(dave);
  await closeCycle(dave);

  // ========================================================================
  // CYCLE E — fundRound settles cycle C across the gap
  // ========================================================================
  const roundE = await game.roundNum();
  expect(roundE === roundC + 2n, "cycle E follows the skipped cycle D");
  console.log(`\n[cycle E] round ${roundE}: fundRound settles cycle C across the skipped round:`);
  {
    // Settlement trigger 3/3: funding a later round settles the last funded one.
    const receipt = await send("fundRound 100 CST", market.fundRound(roundE, 100n * ONE));
    const settled = events(receipt, marketIface, "RoundSettled");
    expect(settled.length === 1 && settled[0].args.roundId === roundC,
      "funding cycle E settles cycle C (cycle D stays skipped)");
  }
  await trade(bob, roundE, 7, true, true, 12n * ONE + 345n); // leave E open with activity

  // ========================================================================
  // EVENT TALLY + STRUCTURAL CHECKS
  // ========================================================================
  console.log("\nEvent tally:");
  const logs = await provider.getLogs({
    address: [dep.marketAddress, dep.vaultAddress],
    fromBlock: dep.startBlock || 0,
    toBlock: "latest",
  });
  const counts = {};
  const parsed = [];
  for (const log of logs) {
    const src = log.address.toLowerCase() === dep.marketAddress.toLowerCase() ? "market" : "vault";
    const p = (src === "market" ? marketIface : vaultIface).parseLog({ topics: [...log.topics], data: log.data });
    if (!p) continue;
    counts[p.name] = (counts[p.name] || 0) + 1;
    parsed.push({ src, name: p.name, args: p.args });
  }
  let failures = 0;
  for (const [name, req] of Object.entries(REQUIRED_EVENTS)) {
    const n = counts[name] || 0;
    const ok = n >= req.min;
    console.log(`  ${ok ? "OK  " : "MISS"} ${req.src.padEnd(6)} ${name.padEnd(18)} ${n} (need >= ${req.min})`);
    if (!ok) failures++;
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
  ];
  for (const [what, ok] of structural) {
    console.log(`  ${ok ? "OK  " : "MISS"} struct ${what}`);
    if (!ok) failures++;
  }
  if (failures > 0) {
    console.error(`\nFAIL: ${failures} required event type(s) or scenario(s) missing`);
    process.exit(1);
  }
  console.log(`\nAll ${Object.keys(REQUIRED_EVENTS).length} event types and ${structural.length} ` +
    `structural scenarios covered across rounds ${roundA}..${roundE} (round ${roundD} skipped). ` +
    `Run cz-etl against this chain to index them.`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
