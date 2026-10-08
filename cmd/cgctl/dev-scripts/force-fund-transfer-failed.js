/**
 * force-fund-transfer-failed.js — make StakingWalletCosmicSignatureNft emit
 * CosmicSignatureEvents.FundTransferFailed on a local Hardhat node.
 *
 * The wallet emits that event from tryPerformMaintenance(charityAddress_)
 * when the ETH transfer of its balance to charityAddress_ fails. Nothing in a
 * normal game flow produces it, so the script manufactures the conditions:
 *
 *   1. tryPerformMaintenance requires numStakedNfts == 0 (it reverts with
 *      ThereAreStakedNfts otherwise); on a played stack the counter is
 *      zeroed for the one call with hardhat_setStorageAt and restored after;
 *   2. gives an empty wallet a balance with hardhat_setBalance (its deposit()
 *      cannot be called while nothing is staked, and only the game may call
 *      it), so the emitted amount is non-zero and visible in SQL;
 *   3. calls tryPerformMaintenance with a contract that has no receive() or
 *      fallback() — by default the CosmicSignatureNft — so the low-level
 *      call{value} reverts, the wallet swallows the failure and emits
 *      FundTransferFailed("ETH transfer to charity failed.", nft, amount).
 *
 * Run from the Cosmic-Signature repo (its hardhat.config.js / artifacts are used):
 *   cd /home/niko/eth/dev/b/cg-v3/Cosmic-Signature-v3.1-2026-08-19
 *   CADDR=0x... npx hardhat run \
 *     /home/niko/eth/dev/b/backend-v3.1/cmd/cgctl/dev-scripts/force-fund-transfer-failed.js --network localhost
 *
 * Env:
 *   CADDR          game proxy address (default: the deploy-dev-and-samp.js address)
 *   REJECT_ADDR    ETH-rejecting destination (default: the game's nft() contract)
 *   BALANCE_ETH    balance to plant in the wallet before the call (default 1)
 *   DRY_RUN        when set, only static-call tryPerformMaintenance and report
 *                  whether it would emit; no state is changed
 *
 * Prints FTF_TX= / FTF_BLOCK= / FTF_WALLET= / FTF_DEST= / FTF_AMOUNT= / FTF_ERRSTR=
 * lines for the shell wrapper (force-fund-transfer-failed.sh), which then
 * checks that cg-etl stored the row in cg_fund_transf_err.
 */
const hre = global.hre ?? require("hardhat");

const CADDR = process.env.CADDR || "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512";
const BALANCE_ETH = process.env.BALANCE_ETH || "1";
const FUND_TRANSFER_FAILED_TOPIC = hre.ethers.id("FundTransferFailed(string,address,uint256)");

// findStorageSlot returns the hex slot index (among the first few slots) whose
// value equals `value`. StakingWalletCosmicSignatureNft inherits Ownable
// (_owner, slot 0) and then StakingWalletNftBase (numStakedNfts, slot 1), so
// the scan resolves to slot 1 unless the layout changes; comparing values
// keeps the script honest about that.
async function findStorageSlot(address, value) {
    const want = hre.ethers.toBeHex(value, 32).toLowerCase();
    for (let slot = 0; slot < 8; slot++) {
        const slotHex = hre.ethers.toBeHex(slot, 32);
        const got = (await hre.ethers.provider.getStorage(address, slotHex)).toLowerCase();
        if (got === want) {
            return slotHex;
        }
    }
    throw new Error(`no storage slot in 0..7 of ${address} holds ${value}; the wallet's storage layout changed`);
}

async function main() {
    const [owner] = await hre.ethers.getSigners();
    const game = await hre.ethers.getContractAt("CosmicSignatureGame", CADDR, owner);

    const walletAddr = await game.stakingWalletCosmicSignatureNft();
    const rejectAddr = process.env.REJECT_ADDR || (await game.nft());
    const wallet = await hre.ethers.getContractAt("StakingWalletCosmicSignatureNft", walletAddr, owner);

    console.log(`Game proxy:                      ${CADDR}`);
    console.log(`StakingWalletCosmicSignatureNft: ${walletAddr}`);
    console.log(`ETH-rejecting destination:       ${rejectAddr}`);

    const walletOwner = await wallet.owner();
    if (walletOwner.toLowerCase() !== owner.address.toLowerCase()) {
        throw new Error(`signer ${owner.address} is not the wallet owner ${walletOwner}; tryPerformMaintenance is onlyOwner`);
    }
    if ((await hre.ethers.provider.getCode(rejectAddr)) === "0x") {
        throw new Error(`${rejectAddr} has no code; a plain account accepts ETH and the transfer would succeed`);
    }

    // tryPerformMaintenance requires numStakedNfts == 0. On a played stack
    // NFTs are usually staked, so temporarily zero the counter's storage
    // slot (hardhat_setStorageAt) and restore it after the call. Only the
    // counter changes; stake actions and the wallet balance stay intact
    // because the transfer is rejected.
    const numStaked = await wallet.numStakedNfts();
    let stakedSlot = null;
    if (numStaked !== 0n) {
        stakedSlot = await findStorageSlot(walletAddr, numStaked);
        console.log(`numStakedNfts = ${numStaked} (storage slot ${stakedSlot}); will be zeroed for the call and restored`);
    }

    if (process.env.DRY_RUN) {
        // Prove the mechanics without touching chain state: the method
        // returns false when the transfer is rejected. Zeroing the counter
        // is part of the real run, so a staked wallet can only be reported.
        if (stakedSlot !== null) {
            console.log("DRY_RUN: staked wallet; the real run zeroes the counter first. Nothing changed.");
            return;
        }
        const wouldSucceed = await wallet.tryPerformMaintenance.staticCall(rejectAddr);
        console.log(`DRY_RUN: tryPerformMaintenance(${rejectAddr}) would ${wouldSucceed ? "SUCCEED (bad destination)" : "fail and emit FundTransferFailed"}`);
        if (wouldSucceed) {
            process.exit(2);
        }
        return;
    }

    if (stakedSlot !== null) {
        await hre.ethers.provider.send("hardhat_setStorageAt", [walletAddr, stakedSlot, hre.ethers.zeroPadValue("0x00", 32)]);
        if ((await wallet.numStakedNfts()) !== 0n) {
            throw new Error("zeroing the numStakedNfts slot did not take effect");
        }
    }
    const restoreStakedCounter = async () => {
        if (stakedSlot === null) {
            return;
        }
        await hre.ethers.provider.send("hardhat_setStorageAt", [walletAddr, stakedSlot, hre.ethers.toBeHex(numStaked, 32)]);
        const restored = await wallet.numStakedNfts();
        console.log(`numStakedNfts restored to ${restored}`);
        if (restored !== numStaked) {
            throw new Error(`numStakedNfts restore failed: ${restored} != ${numStaked}`);
        }
    };

    // Make sure the wallet has something to forward, so the emitted amount
    // is non-zero and visible in SQL. A fresh stack holds nothing (deposit()
    // is game-only and reverts while nothing is staked); a played stack
    // already holds staker rewards, which the rejected transfer leaves intact.
    let balance = await hre.ethers.provider.getBalance(walletAddr);
    if (balance === 0n) {
        const planted = hre.ethers.parseEther(BALANCE_ETH);
        await hre.ethers.provider.send("hardhat_setBalance", [walletAddr, hre.ethers.toBeHex(planted)]);
        balance = await hre.ethers.provider.getBalance(walletAddr);
        console.log(`Planted wallet balance:          ${hre.ethers.formatEther(balance)} ETH`);
    } else {
        console.log(`Wallet balance:                  ${hre.ethers.formatEther(balance)} ETH (left untouched)`);
    }

    let receipt;
    try {
        // Static call first: the method returns false on a failed transfer.
        const wouldSucceed = await wallet.tryPerformMaintenance.staticCall(rejectAddr);
        if (wouldSucceed) {
            throw new Error(`tryPerformMaintenance(${rejectAddr}) would succeed; pick a destination that rejects ETH (REJECT_ADDR)`);
        }
        const tx = await wallet.tryPerformMaintenance(rejectAddr, { gasLimit: 300000 });
        receipt = await tx.wait();
    } finally {
        await restoreStakedCounter();
    }
    console.log(`tryPerformMaintenance tx:        ${receipt.hash} (block ${receipt.blockNumber})`);

    const log = receipt.logs.find(
        (l) => l.address.toLowerCase() === walletAddr.toLowerCase() && l.topics[0] === FUND_TRANSFER_FAILED_TOPIC,
    );
    if (!log) {
        throw new Error("receipt carries no FundTransferFailed log from the wallet");
    }
    const iface = new hre.ethers.Interface([
        "event FundTransferFailed(string errStr, address indexed destinationAddress, uint256 amount)",
    ]);
    const parsed = iface.parseLog(log);
    console.log(`FundTransferFailed emitted:      errStr="${parsed.args.errStr}" destination=${parsed.args.destinationAddress} amount=${parsed.args.amount}`);

    const after = await hre.ethers.provider.getBalance(walletAddr);
    if (after !== balance) {
        throw new Error(`wallet balance changed from ${balance} to ${after}; the transfer was not rejected`);
    }

    console.log(`FTF_TX=${receipt.hash}`);
    console.log(`FTF_BLOCK=${receipt.blockNumber}`);
    console.log(`FTF_WALLET=${walletAddr}`);
    console.log(`FTF_DEST=${parsed.args.destinationAddress}`);
    console.log(`FTF_AMOUNT=${parsed.args.amount}`);
    console.log(`FTF_ERRSTR=${parsed.args.errStr}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
