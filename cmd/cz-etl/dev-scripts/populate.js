// ChaosZero sample-data generator for a local development chain (hardhat
// node or anvil, chain id 31337). Mirrors the spirit of
// cmd/cgctl/dev-scripts/populate.js: drive one full market lifecycle so
// every event cz-etl indexes appears at least once, then tally the logs and
// FAIL (exit 1) if any event type is missing.
//
// Covered events (13):
//   market: RoundInitialized, RoundFunded, Trade, AnswersObserved,
//           PositionSplit, PositionsMerged, RoundSettled, Claimed,
//           TransferSingle, TransferBatch, ApprovalForAll
//   vault:  CycleEscrowed, Paid
//
// Scenario: fund the open round again; trade four questions with exact
// quotes (buy/sell amounts are found by probing the market's invariant
// check via staticCall); split a position set and move it around with
// direct, operator-batch, and approval flows; merge it back; make game
// gestures; close the cycle; fund the next round (which settles the old
// one into the vault); claim the winnings.
//
// Usage (after deploy-local.js, against the same chain):
//   CHAOSZERO_DIR=/path/to/ChaosZero \
//   NODE_PATH=$CHAOSZERO_DIR/node_modules \
//   node populate.js
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

// REQUIRED_EVENTS maps every event cz-etl indexes to the contract that
// emits it; the final tally fails if any is missing.
const REQUIRED_EVENTS = {
  RoundInitialized: "market",
  RoundFunded: "market",
  Trade: "market",
  AnswersObserved: "market",
  PositionSplit: "market",
  PositionsMerged: "market",
  RoundSettled: "market",
  Claimed: "market",
  TransferSingle: "market",
  TransferBatch: "market",
  ApprovalForAll: "market",
  CycleEscrowed: "vault",
  Paid: "vault",
};

async function main() {
  const root = chaosZeroDir();
  const dep = readDeployment();
  const { provider, signers } = await connect();
  const [deployer, alice, bob] = signers;

  const marketArt = findArtifact(root, "IndependentOutcomeMarket");
  const vaultArt = findArtifact(root, "OutcomeSettlementVault");
  const gameArt = findArtifact(root, "LocalCosmicGame");
  const cstArt = findArtifact(root, "LocalCst");

  const market = new ethers.Contract(dep.marketAddress, marketArt.abi, deployer);
  const game = new ethers.Contract(dep.gameAddress, gameArt.abi, deployer);
  const cst = new ethers.Contract(dep.tokenAddress, cstArt.abi, deployer);

  const send = async (label, promise) => {
    const receipt = await (await promise).wait();
    if (receipt.status !== 1) throw new Error(`${label} reverted`);
    console.log(`  ok ${label} (block ${receipt.blockNumber})`);
    return receipt;
  };
  const deadline = async () => BigInt((await provider.getBlock("latest")).timestamp) + 900n;

  // An exact-pair trade: probe the largest output the invariant check
  // accepts, then execute with that quote.
  const trade = async (signer, round, question, yes, isBuy, input) => {
    const m = market.connect(signer);
    const dl = await deadline();
    const fn = isBuy ? m.getFunction("buy") : m.getFunction("sell");
    const output = await quoteBySearch((out) =>
      fn.staticCall(round, question, yes, input, out, dl),
    );
    await send(
      `${isBuy ? "buy" : "sell"} q${question} ${yes ? "YES" : "NO"} ` +
        `${ethers.formatEther(input)} -> ${ethers.formatEther(output)}`,
      fn.send(round, question, yes, input, output, await deadline()),
    );
    return output;
  };

  const round = await game.roundNum();
  console.log(`Populating round ${round} on ${dep.rpcUrl}`);

  // --- Funding beyond the deployment seed (RoundFunded again) ------------
  console.log("Funding:");
  for (const s of [alice, bob]) {
    await send(`faucet ${await s.getAddress()}`, cst.connect(s).faucet(50000n * ONE));
    await send("approve market", cst.connect(s).approve(dep.marketAddress, 50000n * ONE));
  }
  await send("fundRound +500 CST", market.connect(alice).fundRound(round, 500n * ONE));

  // --- Game gestures: make question 0 mathematically known ---------------
  // (first cycle: any gesture count exceeds the nonexistent previous cycle)
  console.log("Gestures:");
  for (let i = 0; i < 5; i++) {
    const price = await game.getNextEthPrice();
    await send(`bidEth #${i + 1}`, game.connect(alice).bidEth({ value: price }));
  }

  // --- Trades (Trade; the first one also observes answers) ---------------
  console.log("Trades:");
  await trade(alice, round, 0, true, true, 100n * ONE);
  const q5Units = await trade(alice, round, 5, false, true, 50n * ONE);
  await trade(alice, round, 5, false, false, q5Units / 4n);
  await trade(bob, round, 7, true, true, 25n * ONE);
  await trade(bob, round, 3, false, true, 10n * ONE);

  // --- Position sets: split, transfer, operator batch, merge -------------
  console.log("Positions:");
  const splitAmount = 25n * ONE;
  await send("splitPosition q2", market.connect(bob).splitPosition(round, 2, splitAmount));
  const idYes = await market.positionId(round, 2, true);
  const idNo = await market.positionId(round, 2, false);
  await send(
    "safeTransferFrom q2 YES bob->alice",
    market
      .connect(bob)
      .safeTransferFrom(await bob.getAddress(), await alice.getAddress(), idYes, splitAmount, "0x"),
  );
  await send("setApprovalForAll bob->alice", market.connect(bob).setApprovalForAll(await alice.getAddress(), true));
  await send(
    "safeBatchTransferFrom q2 NO bob->alice (operator)",
    market
      .connect(alice)
      .safeBatchTransferFrom(await bob.getAddress(), await alice.getAddress(), [idNo], [splitAmount], "0x"),
  );
  await send("revoke approval", market.connect(bob).setApprovalForAll(await alice.getAddress(), false));
  await send("mergePosition q2", market.connect(alice).mergePosition(round, 2, splitAmount / 2n));

  // --- Close the cycle and settle into the next round --------------------
  console.log("Cycle close + settlement:");
  const mainPrizeTime = await game.mainPrizeTime();
  await rpc(provider, "evm_setNextBlockTimestamp", [Number(mainPrizeTime + 1n)]);
  await send("closeCycle", game.connect(alice).closeCycle());
  const nextRound = await game.roundNum();
  if (nextRound === round) throw new Error("closeCycle did not advance the round");
  await send(
    `fundRound(${nextRound}) settles round ${round}`,
    market.connect(deployer).fundRound(nextRound, 100n * ONE),
  );

  // --- Claims (Claimed + vault Paid) --------------------------------------
  console.log("Claims:");
  let claimed = 0;
  for (const [name, s] of [["alice", alice], ["bob", bob]]) {
    try {
      await send(`claimKnown(${round}) by ${name}`, market.connect(s).claimKnown(round));
      claimed++;
    } catch (err) {
      console.log(`  -- ${name} had nothing to claim (${err.shortMessage || err.message})`);
    }
  }
  if (claimed === 0) throw new Error("no account could claim winnings");

  // --- Event tally ---------------------------------------------------------
  console.log("\nEvent tally:");
  const ifaces = {
    market: new ethers.Interface(marketArt.abi),
    vault: new ethers.Interface(vaultArt.abi),
  };
  const counts = {};
  const logs = await provider.getLogs({
    address: [dep.marketAddress, dep.vaultAddress],
    fromBlock: dep.startBlock || 0,
    toBlock: "latest",
  });
  for (const log of logs) {
    const src = log.address.toLowerCase() === dep.marketAddress.toLowerCase() ? "market" : "vault";
    const parsed = ifaces[src].parseLog({ topics: [...log.topics], data: log.data });
    if (parsed) counts[parsed.name] = (counts[parsed.name] || 0) + 1;
  }
  let missing = 0;
  for (const [name, src] of Object.entries(REQUIRED_EVENTS)) {
    const n = counts[name] || 0;
    console.log(`  ${(n > 0 ? "OK  " : "MISS")} ${src.padEnd(6)} ${name.padEnd(18)} ${n}`);
    if (n === 0) missing++;
  }
  for (const [name, n] of Object.entries(counts)) {
    if (!(name in REQUIRED_EVENTS)) console.log(`  extra ${name}: ${n}`);
  }
  if (missing > 0) {
    console.error(`\nFAIL: ${missing} required event type(s) missing`);
    process.exit(1);
  }
  console.log(`\nAll ${Object.keys(REQUIRED_EVENTS).length} required event types emitted. ` +
    `Run cz-etl against this chain to index them.`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
