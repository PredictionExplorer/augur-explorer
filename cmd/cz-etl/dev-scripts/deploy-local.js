// Deploys the ChaosZero local stack to a local development chain (hardhat
// node or anvil, chain id 31337), replicating the repo's own foundry script
// script/DeployIndependentLocal.s.sol:
//
//   1. LocalCosmicGame            (its constructor deploys LocalCst)
//   2. LocalOutcomeSource(game)
//   3. IndependentOutcomeMarket(cst, source)   (constructor deploys the vault)
//   4. cst.faucet(100_000 CST), approve, fundRound(game.roundNum(), 10_000 CST)
//
// The addresses are written to cz-local-deployment.json (consumed by
// populate.js) and the SQL that points a local backend database at this
// deployment is printed: cz_contracts must hold the local addresses and
// cz_proc_status.last_block must drop to 0, because the production default
// (517037316) would make cz-etl skip every local block.
//
// Usage (requires a built ChaosZero checkout and a running local chain):
//   CHAOSZERO_DIR=/path/to/ChaosZero \
//   NODE_PATH=$CHAOSZERO_DIR/node_modules \
//   node deploy-local.js
//
// Environment: RPC_URL (default http://127.0.0.1:8545),
// CZ_DEPLOYMENT_FILE (default ./cz-local-deployment.json next to this script).
"use strict";

const fs = require("fs");
const {
  ONE,
  chaosZeroDir,
  findArtifact,
  connect,
  deploymentFile,
  ethers,
} = require("./lib.js");

async function main() {
  const root = chaosZeroDir();
  const { provider, signers, rpcUrl } = await connect();
  const deployer = signers[0];

  const deployContract = async (name, args = []) => {
    const art = findArtifact(root, name);
    if (!art.bytecode || art.bytecode === "0x") {
      console.error(`Artifact ${art.path} has no bytecode (ABI-only build?)`);
      process.exit(1);
    }
    const factory = new ethers.ContractFactory(art.abi, art.bytecode, deployer);
    const contract = await factory.deploy(...args);
    await contract.waitForDeployment();
    const addr = await contract.getAddress();
    console.log(`${name.padEnd(28)} ${addr}`);
    return contract;
  };

  const startBlock = await provider.getBlockNumber();

  const game = await deployContract("LocalCosmicGame");
  const cstAddr = await game.token();
  console.log(`${"LocalCst (by game ctor)".padEnd(28)} ${cstAddr}`);
  const source = await deployContract("LocalOutcomeSource", [await game.getAddress()]);
  const market = await deployContract("IndependentOutcomeMarket", [
    cstAddr,
    await source.getAddress(),
  ]);
  const vaultAddr = await market.settlementVault();
  console.log(`${"OutcomeSettlementVault".padEnd(28)} ${vaultAddr}`);

  // Seed the first round exactly like DeployIndependentLocal.s.sol.
  const cst = new ethers.Contract(cstAddr, findArtifact(root, "LocalCst").abi, deployer);
  await (await cst.faucet(100000n * ONE)).wait();
  await (await cst.approve(await market.getAddress(), 100000n * ONE)).wait();
  const round = await game.roundNum();
  await (await market.fundRound(round, 10000n * ONE)).wait();
  console.log(`Funded round ${round} with 10000 CST (RoundInitialized + RoundFunded emitted)`);

  const deployment = {
    rpcUrl,
    chainId: 31337,
    startBlock,
    gameAddress: await game.getAddress(),
    tokenAddress: cstAddr,
    sourceAddress: await source.getAddress(),
    marketAddress: await market.getAddress(),
    vaultAddress: vaultAddr,
    deployedAt: new Date().toISOString(),
  };
  fs.writeFileSync(deploymentFile(), JSON.stringify(deployment, null, 2) + "\n");
  console.log(`\nDeployment written to ${deploymentFile()}`);

  console.log(`
-- Point a local backend database at this deployment before starting cz-etl:
UPDATE cz_contracts SET
	market_addr = '${deployment.marketAddress}',
	source_addr = '${deployment.sourceAddress}',
	vault_addr  = '${deployment.vaultAddress}';
UPDATE cz_proc_status SET last_evt_id = 0, last_block = 0;
`);
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
