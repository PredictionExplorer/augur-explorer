// Shared helpers for the ChaosZero local dev scripts (deploy-local.js,
// populate.js). Plain CommonJS: `ethers` (v6) is resolved explicitly from
// the ChaosZero checkout's node_modules via createRequire, so neither a
// backend package.json nor NODE_PATH ordering (which prefers stray parent
// node_modules directories) can hijack the import.
//
// Contract artifacts (ABI + bytecode) are read from the ChaosZero checkout:
// a hardhat v3 build (.local/hardhat-artifacts) or a foundry build (out/),
// whichever exists.
"use strict";

const fs = require("fs");
const path = require("path");
const { createRequire } = require("module");

const LOCAL_CHAIN_ID = 31337n;
const ONE = 10n ** 18n;

// Well-known hardhat/anvil developer accounts (mnemonic "test test ... junk").
// NEVER use these keys anywhere but a disposable local chain.
const WELL_KNOWN_KEYS = [
  "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
  "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
  "0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
  "0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
  "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
  "0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba",
];

function chaosZeroDir() {
  const dir = process.env.CHAOSZERO_DIR;
  if (!dir) {
    console.error("CHAOSZERO_DIR must point at the ChaosZero checkout");
    process.exit(1);
  }
  if (!fs.existsSync(path.join(dir, "src", "IndependentOutcomeMarket.sol"))) {
    console.error(`CHAOSZERO_DIR=${dir} does not look like the ChaosZero repo`);
    process.exit(1);
  }
  return dir;
}

// ethers v6 from the ChaosZero checkout (an ethers v5 higher up the
// filesystem must not win the resolution).
const czRequire = createRequire(path.join(chaosZeroDir(), "package.json"));
const ethers = czRequire("ethers");
if (!ethers.JsonRpcProvider) {
  console.error(`ethers resolved from ${chaosZeroDir()} is not v6; run 'pnpm install' there`);
  process.exit(1);
}

// artifactRoots lists the artifact directories a ChaosZero build produces,
// in preference order.
function artifactRoots(root) {
  return [
    path.join(root, ".local", "hardhat-artifacts"),
    path.join(root, "out"),
    path.join(root, "artifacts"),
  ].filter((p) => fs.existsSync(p));
}

// findArtifact walks the artifact roots for <Name>.json (skipping debug and
// build-info files) and normalizes hardhat/foundry shapes into {abi,
// bytecode}.
function findArtifact(root, name) {
  const target = `${name}.json`;
  const stack = artifactRoots(root);
  if (stack.length === 0) {
    console.error(
      `No artifacts found under ${root}. Build ChaosZero first ` +
        "(pnpm install && pnpm build, or forge build).",
    );
    process.exit(1);
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
      const bytecode =
        typeof parsed.bytecode === "string"
          ? parsed.bytecode // hardhat artifact
          : parsed.bytecode && parsed.bytecode.object; // foundry artifact
      return { abi: parsed.abi, bytecode: bytecode || "0x", path: full };
    }
  }
  console.error(`Artifact ${target} not found under ${artifactRoots(root).join(", ")}`);
  process.exit(1);
}

async function connect() {
  const rpcUrl = process.env.RPC_URL || "http://127.0.0.1:8545";
  const provider = new ethers.JsonRpcProvider(rpcUrl, undefined, { polling: true, pollingInterval: 50 });
  const network = await provider.getNetwork();
  if (network.chainId !== LOCAL_CHAIN_ID) {
    console.error(
      `RPC ${rpcUrl} reports chain id ${network.chainId}; ` +
        `these scripts only ever run against a local ${LOCAL_CHAIN_ID} chain`,
    );
    process.exit(1);
  }
  // NonceManager serializes nonces locally; without it back-to-back
  // transactions race the node's pending-nonce view.
  const signers = WELL_KNOWN_KEYS.map((k) => new ethers.NonceManager(new ethers.Wallet(k, provider)));
  return { provider, signers, rpcUrl };
}

// deploymentFile is where deploy-local.js records the local addresses and
// populate.js reads them back.
function deploymentFile() {
  return process.env.CZ_DEPLOYMENT_FILE || path.join(__dirname, "cz-local-deployment.json");
}

function readDeployment() {
  const file = deploymentFile();
  if (!fs.existsSync(file)) {
    console.error(`Deployment file ${file} not found; run deploy-local.js first`);
    process.exit(1);
  }
  return JSON.parse(fs.readFileSync(file, "utf8"));
}

// quoteBySearch finds the largest output the market accepts for an exact
// trade using staticCall. buy()/sell() verify a weighted product invariant,
// so valid outputs form a contiguous [0..max] range: probe by doubling, then
// bisect the boundary. ~60 eth_calls per trade, negligible on a local node.
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
  // Invariant: ok(lo), !ok(hi).
  while (hi - lo > 1n) {
    const mid = (lo + hi) / 2n;
    if (await ok(mid)) lo = mid;
    else hi = mid;
  }
  if (lo === 0n) throw new Error("market did not accept any output for this trade");
  return lo;
}

// rpc sends a raw JSON-RPC request (evm_increaseTime etc.).
function rpc(provider, method, params = []) {
  return provider.send(method, params);
}

module.exports = {
  ONE,
  WELL_KNOWN_KEYS,
  chaosZeroDir,
  findArtifact,
  connect,
  deploymentFile,
  readDeployment,
  quoteBySearch,
  rpc,
  ethers,
};
