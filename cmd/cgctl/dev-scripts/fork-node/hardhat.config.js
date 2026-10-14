// Minimal config for the rehearsal fork node (see package.json for why it is standalone).
// No contracts are compiled here; the Cosmic-Signature repo deploys over RPC.
module.exports = {
    solidity: "0.8.24",
    networks: {
        hardhat: {
            chainId: 31337,
            allowUnlimitedContractSize: true,
            // Matches the main repo's HARDHAT_MODE_CODE=2 node behavior:
            // no automine surprises for same-timestamp blocks during time jumps.
            allowBlocksWithSameTimestamp: true,
            blockGasLimit: 10_000 * 30_000_000,
        },
    },
};
