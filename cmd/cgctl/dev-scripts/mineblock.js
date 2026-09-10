// Mine a single block on a Hardhat/dev node. On Hardhat, block.timestamp
// only moves when a block is mined, so contract countdowns (e.g. the
// delayDurationBeforeRoundActivation window after claimMainPrize) appear
// frozen until someone mines. Run this once — or periodically, see
// mine-loop.sh — to keep the chain clock ticking during manual testing.
//
// Usage (from the Solidity repo, node on localhost:8545):
//   npx hardhat run <backend>/cmd/cgctl/dev-scripts/mineblock.js --network localhost
//
// Uncomment the evm_increaseTime line to also jump the clock forward.

async function main() {
    // await ethers.provider.send("evm_increaseTime", [6000]);
    await ethers.provider.send("evm_mine");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
