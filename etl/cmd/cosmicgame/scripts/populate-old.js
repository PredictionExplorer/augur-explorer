const hre = require("hardhat");
const { basicDeployment } = require("./Deploy.js");

// ============================================================================
// CONFIGURATION PARAMETERS
// ============================================================================
const CONFIG = {
    donations: {
        eth: [
            hre.ethers.parseEther("100"),  // Simple donation
        ],
        ethWithInfo: [
            {
                amount: hre.ethers.parseEther("60"),
                data: '{"version":1,"title":"Hardhat donation","message":"Donation from HardHat","url":"http://hardhat.org"}'
            },
            {
                amount: hre.ethers.parseEther("90"),
                data: '{"version":1,"title":"ArtBlocks donation","message":"ArtBlocks offers a platform for creators","url":"https://www.artblocks.io"}'
            },
            {
                amount: hre.ethers.parseEther("9"),
                data: '{"version":1,"title":"EF donation","message":"Ethereum Foundation","url":"http://ethereum.org/en"}'
            },
            {
                amount: hre.ethers.parseEther("8"),
                data: '{"version":1,"title":"EF donation","message":"Ethereum Foundation","url":"http://ethereum.org/en"}'
            },
        ],
        regular: [
            hre.ethers.parseEther("500"),  // Regular donation (round 6)
        ],
        charity: [
            hre.ethers.parseEther("4"),    // First charity donation
            hre.ethers.parseEther("4"),    // Second charity donation
        ],
    },
    bids: {
        round0: [
            { signer: 'addr1', message: 'bid 1', rwalkToken: -1, extraValue: 1000n },
            { signer: 'addr2', message: 'bid 1', rwalkToken: -1, extraValue: 1000n },
            { signer: 'addr1', message: 'bid 2', rwalkToken: -1 },
            { signer: 'addr1', message: 'bid 2', rwalkToken: -1 },
            { signer: 'owner', message: 'bidWithRWalk', rwalkToken: 'MINT_RWALK' },
            { signer: 'addr3', message: 'bid 3', rwalkToken: -1 },
            { signer: 'addr3', message: 'bid 3', rwalkToken: -1 },
            { signer: 'addr4', message: '', rwalkToken: -1 },
            { signer: 'addr4', message: '', rwalkToken: -1 },
            { signer: 'addr5', message: '', rwalkToken: -1 },
            { signer: 'addr5', message: '', rwalkToken: -1 },
        ],
        round1: [
            { signer: 'addr1', message: 'bid 4', rwalkToken: -1 },
        ],
        round2: [
            { signer: 'addr1', message: 'bid 5', rwalkToken: -1 },
        ],
        round3: [
            { signer: 'addr1', message: '', rwalkToken: -1 },
            { signer: 'addr1', message: '', rwalkToken: -1 },
        ],
        round4: [
            { signer: 'addr3', message: 'bid 3', rwalkToken: -1 },
            { signer: 'addr3', message: 'bid 3', rwalkToken: -1 },
        ],
        round5: [
            { signer: 'addr2', message: 'opening bid (addr2)', rwalkToken: -1 },
            { signer: 'addr1', message: 'CST bid (addr1)', useCst: true },
        ],
        round6: [
            { signer: 'addr3', message: 'bid eth', rwalkToken: -1 },
            { signer: 'addr3', message: 'CST bid addr1', useCst: true },
        ],
        round7: [
            { signer: 'addr3', message: 'bid eth', rwalkToken: -1 },
            { signer: 'addr3', message: 'CST bid addr1', useCst: true },
        ],
    },
    staking: {
        rwalkTokens: {
            addr1: 5,
            addr2: 5,
            addr3: 50,
        },
    },
    marketing: {
        rewards: [
            { addr: 'addr1', amount: hre.ethers.parseEther("7") },
            { addr: 'addr2', amount: hre.ethers.parseEther("7") },
            { addr: 'addr2', amount: hre.ethers.parseEther("2") },
            { addr: 'addr1', amount: hre.ethers.parseEther("6") },
            { addr: 'addr2', amount: hre.ethers.parseEther("5") },
            { addr: 'addr2', amount: hre.ethers.parseEther("5") },
            { addr: 'addr3', amount: hre.ethers.parseEther("5") },
            { addr: 'addr4', amount: hre.ethers.parseEther("1") },
            { addr: 'addr1', amount: hre.ethers.parseEther("11") },
        ],
    },
    erc20Donations: {
        round2: 10000000000000000000n,
        round4: 11000000000000000000n,
    },
    gameSettings: {
        timeoutClaimPrize: 120,
        prizeTimeIncrement: 900000000,
        initialDurationDivisor: 1000000,
        delayBeforeActivation: 1,
        numRaffleEthPrizes: 4,
        numRaffleNftsBidders: 6,
        numRaffleNftsStakers: 3,
        mainPrizePercentage: 30,
        charityPercentage: 5,
        rafflePercentage: 6,
        chronoPercentage: 8,
        stakingPercentage: 19,
        cstPrizeAmount: 99000000000000000000n,
        cstRewardForBidding: 130000000000000000000n,
        marketingWalletContribution: 120000000000000000000n,
        cstMinLimit: 150000000000000000000n,
        maxMessageLength: 199,
        cstDutchAuctionDuration: 13 * 3600,
    },
};

const bidParamsEncoding = {
    type: "tuple(string,int256)",
    name: "bidparams",
    components: [
        { name: "msg", type: "string" },
        { name: "rwalk", type: "int256" },
    ],
};

// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

async function mintRwalk(signer) {
    const tokenPrice = await randomWalkNFT.getMintPrice();
    const tx = await randomWalkNFT.connect(signer).mint({ value: tokenPrice });
    const receipt = await tx.wait();
    const topicSig = randomWalkNFT.interface.getEvent("MintEvent").topicHash;
    const log = receipt.logs.find((x) => x.topics.indexOf(topicSig) >= 0);
    const parsedLog = randomWalkNFT.interface.parseLog(log);
    return parsedLog.args[0];
}

async function stakeAvailableNfts() {
    const totalSupply = await cosmicSignature.totalSupply();
    const stakingWalletAddr = await stakingWalletCosmicSignatureNft.getAddress();
    
    for (let i = 0; i < totalSupply; i++) {
        const owner = await cosmicSignature.ownerOf(i);
        if (owner === stakingWalletAddr) continue; // Already staked
        
        const ownerSigner = await hre.ethers.getSigner(owner);
        if (!ownerSigner) continue;
        
        try {
            await cosmicSignature.connect(ownerSigner).setApprovalForAll(stakingWalletAddr, true);
            await stakingWalletCosmicSignatureNft.connect(ownerSigner).stake(i);
        } catch (e) {
            // Ignore staking errors
        }
    }
}

async function unstakeAllNfts() {
    const numActions = await stakingWalletCosmicSignatureNft.actionCounter();
    let numUnstaked = 0;
    
    for (let i = 1; i <= numActions; i++) {
        const actionRec = (await stakingWalletCosmicSignatureNft.stakeActions(i)).toObject();
        const owner = actionRec.nftOwnerAddress;
        
        if (owner === "0x0000000000000000000000000000000000000000") continue;
        
        const ownerSigner = await hre.ethers.getSigner(owner);
        await stakingWalletCosmicSignatureNft.connect(ownerSigner).unstake(i);
        numUnstaked++;
    }
    return numUnstaked;
}

async function withdrawAllUnclaimedDeposits() {
    const filter = prizesWallet.filters.EthReceived();
    const events = await prizesWallet.queryFilter(filter);
    const withdrawalDone = {};
    
    for (const event of events) {
        const winnerAddress = event.args.prizeWinnerAddress;
        const roundNum = event.args.roundNum;
        const withdrawalKey = `${winnerAddress}_${roundNum}`;
        
        if (withdrawalDone[withdrawalKey]) continue;
        
        try {
            const winnerSigner = await hre.ethers.getSigner(winnerAddress);
            await prizesWallet.connect(winnerSigner).withdrawEth(roundNum);
            withdrawalDone[withdrawalKey] = true;
        } catch (e) {
            // Already withdrawn or no balance
        }
    }
}

async function placeBid(signer, rwalkTokenId, message, value) {
    const bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    const bidValue = value || bidPrice;
    await cosmicGameProxy.connect(signer).bidWithEth(rwalkTokenId, message, { value: bidValue });
}

async function placeCstBid(signer, message) {
    const cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    await cosmicGameProxy.connect(signer).bidWithCst(cstPrice, message);
}

async function placeBidsForRound(roundName) {
    const bids = CONFIG.bids[roundName];
    if (!bids) {
        console.log(`No bids configured for ${roundName}`);
        return;
    }
    
    console.log(`Placing ${bids.length} bids for ${roundName}`);
    
    for (const bid of bids) {
        const signer = signerMap[bid.signer];
        
        if (bid.useCst) {
            // CST bid
            await placeCstBid(signer, bid.message);
        } else {
            // ETH bid
            let rwalkToken = bid.rwalkToken;
            
            // Handle special case: mint RandomWalk token for this bid
            if (rwalkToken === 'MINT_RWALK') {
                rwalkToken = Number(await mintRwalk(signer));
            }
            
            // Calculate bid value
            const bidPrice = await cosmicGameProxy.getNextEthBidPrice();
            const bidValue = bid.extraValue ? bidPrice + bid.extraValue : bidPrice;
            
            await cosmicGameProxy.connect(signer).bidWithEth(rwalkToken, bid.message, { value: bidValue });
        }
    }
}

async function advanceTimeAndClaim(claimer, gasLimit = 9000000) {
    const prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    
    const tx = await cosmicGameProxy.connect(claimer).claimMainPrize({ gasLimit });
    return await tx.wait();
}

async function mintAndStakeRwalkTokens(signer, count) {
    const tokens = [];
    const stakingWalletAddr = await stakingWalletRandomWalkNft.getAddress();
    
    for (let i = 0; i < count; i++) {
        const tokenId = await mintRwalk(signer);
        tokens.push(tokenId);
        await randomWalkNFT.connect(signer).setApprovalForAll(stakingWalletAddr, true);
        await stakingWalletRandomWalkNft.connect(signer).stake(tokenId);
    }
    return tokens;
}

async function donateNft(signer, nftContract, tokenId, bidMessage) {
    const bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    const nftAddr = await nftContract.getAddress();
    await nftContract.connect(signer).setApprovalForAll(await prizesWallet.getAddress(), true);
    await cosmicGameProxy.connect(signer).bidWithEthAndDonateNft(
        -1,
        bidMessage,
        nftAddr,
        tokenId,
        { value: bidPrice }
    );
}

async function donateErc20Token(tokenContract, amount, bidMessage) {
    const bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    const tokenAddr = await tokenContract.getAddress();
    await tokenContract.approve(await cosmicGameProxy.getAddress(), hre.ethers.parseEther("9999999999999999"));
    await tokenContract.approve(await prizesWallet.getAddress(), hre.ethers.parseEther("9999999999999999"));
    await cosmicGameProxy.bidWithEthAndDonateToken(-1, bidMessage, tokenAddr, amount, { value: bidPrice });
}

async function approveCosmicToken(signers) {
    const approvalAmount = hre.ethers.parseEther("10000000");
    const gameAddr = await cosmicGameProxy.getAddress();
    
    for (const signer of signers) {
        await cosmicToken.connect(signer).approve(gameAddr, approvalAmount);
    }
}

async function setActivationTimeToNow() {
    const latestBlock = await hre.ethers.provider.getBlock("latest");
    const roundActivationTime = latestBlock.timestamp - 1;
    await (await cosmicGameProxy.setRoundActivationTime(roundActivationTime)).wait();
    return roundActivationTime;
}

async function setupGameConfiguration() {
    const { gameSettings } = CONFIG;
    
    await cosmicGameProxy.connect(owner).setTimeoutDurationToClaimMainPrize(gameSettings.timeoutClaimPrize, { gasLimit: 1000000 });
    await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementInMicroSeconds(gameSettings.prizeTimeIncrement, { gasLimit: 1000000 });
    await cosmicGameProxy.connect(owner).setInitialDurationUntilMainPrizeDivisor(gameSettings.initialDurationDivisor, { gasLimit: 1000000 });
    await cosmicGameProxy.connect(owner).setNumRaffleEthPrizesForBidders(gameSettings.numRaffleEthPrizes);
    await cosmicGameProxy.connect(owner).setNumRaffleCosmicSignatureNftsForBidders(gameSettings.numRaffleNftsBidders);
    await cosmicGameProxy.connect(owner).setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(gameSettings.numRaffleNftsStakers);
    await cosmicGameProxy.connect(owner).setMainEthPrizeAmountPercentage(gameSettings.mainPrizePercentage);
    await cosmicGameProxy.connect(owner).setCharityEthDonationAmountPercentage(gameSettings.charityPercentage);
    await cosmicGameProxy.connect(owner).setRaffleTotalEthPrizeAmountForBiddersPercentage(gameSettings.rafflePercentage);
    await cosmicGameProxy.connect(owner).setChronoWarriorEthPrizeAmountPercentage(gameSettings.chronoPercentage);
    await cosmicGameProxy.connect(owner).setCosmicSignatureNftStakingTotalEthRewardAmountPercentage(gameSettings.stakingPercentage);
    await cosmicGameProxy.connect(owner).setCstPrizeAmount(gameSettings.cstPrizeAmount);
    await cosmicGameProxy.connect(owner).setCstRewardAmountForBidding(gameSettings.cstRewardForBidding);
    await cosmicGameProxy.connect(owner).setMarketingWalletCstContributionAmount(gameSettings.marketingWalletContribution);
    await cosmicGameProxy.connect(owner).setCstDutchAuctionBeginningBidPriceMinLimit(gameSettings.cstMinLimit);
    await cosmicGameProxy.connect(owner).setBidMessageLengthMaxLimit(gameSettings.maxMessageLength);
    await cosmicGameProxy.connect(owner).setCstDutchAuctionDurationDivisor(gameSettings.cstDutchAuctionDuration);
    
    console.log("Game configuration complete");
}

async function setupContractAddresses() {
    await cosmicGameProxy.connect(owner).setCharityAddress(await charityWallet.getAddress());
    await cosmicGameProxy.connect(owner).setRandomWalkNft(await randomWalkNFT.getAddress());
    await cosmicGameProxy.connect(owner).setPrizesWallet(await prizesWallet.getAddress());
    await cosmicGameProxy.connect(owner).setStakingWalletCosmicSignatureNft(await stakingWalletCosmicSignatureNft.getAddress());
    await cosmicGameProxy.connect(owner).setStakingWalletRandomWalkNft(await stakingWalletRandomWalkNft.getAddress());
    await cosmicGameProxy.connect(owner).setMarketingWallet(await marketingWallet.getAddress());
    await cosmicGameProxy.connect(owner).setCosmicSignatureToken(await cosmicToken.getAddress());
    await cosmicGameProxy.connect(owner).setCosmicSignatureNft(await cosmicSignature.getAddress());
    
    console.log("Contract addresses configured");
}

async function payMarketingRewards(treasurer) {
    for (const reward of CONFIG.marketing.rewards) {
        const recipientAddr = signerMap[reward.addr].address;
        await marketingWallet.connect(treasurer).payReward(recipientAddr, reward.amount);
        await ethers.provider.send("evm_mine");
    }
    console.log(`Paid ${CONFIG.marketing.rewards.length} marketing rewards`);
}

async function printContractAddresses() {
    console.log("Contract Addresses Deployed:");
    console.log(
        "INSERT INTO cg_contracts VALUES(" +
        `'${await cosmicGameProxy.getAddress()}',` +
        `'${await cosmicSignature.getAddress()}',` +
        `'${await cosmicToken.getAddress()}',` +
        `'${await cosmicDAO.getAddress()}',` +
        `'${await charityWallet.getAddress()}',` +
        `'${await prizesWallet.getAddress()}',` +
        `'${await randomWalkNFT.getAddress()}',` +
        `'${await stakingWalletCosmicSignatureNft.getAddress()}',` +
        `'${await stakingWalletRandomWalkNft.getAddress()}',` +
        `'${await marketingWallet.getAddress()}',` +
        `'${implementationAddr}')`
    );
}

// ============================================================================
// MAIN EXECUTION
// ============================================================================

let owner, addr1, addr2, addr3, addr4, addr5, addrs;
let cosmicGameProxy, cosmicToken, cosmicSignature, charityWallet, cosmicDAO;
let prizesWallet, randomWalkNFT, stakingWalletCosmicSignatureNft;
let stakingWalletRandomWalkNft, marketingWallet, implementationAddr;
let signerMap; // Map of signer names to signer objects

// Donation iterators
let ethDonationIndex = 0;
let ethWithInfoDonationIndex = 0;
let regularDonationIndex = 0;
let charityDonationIndex = 0;

function getNextEthDonation() {
    if (ethDonationIndex >= CONFIG.donations.eth.length) {
        throw new Error("No more ETH donations configured");
    }
    return CONFIG.donations.eth[ethDonationIndex++];
}

function getNextEthWithInfoDonation() {
    if (ethWithInfoDonationIndex >= CONFIG.donations.ethWithInfo.length) {
        throw new Error("No more ETH donations with info configured");
    }
    return CONFIG.donations.ethWithInfo[ethWithInfoDonationIndex++];
}

function getNextRegularDonation() {
    if (regularDonationIndex >= CONFIG.donations.regular.length) {
        throw new Error("No more regular donations configured");
    }
    return CONFIG.donations.regular[regularDonationIndex++];
}

function getNextCharityDonation() {
    if (charityDonationIndex >= CONFIG.donations.charity.length) {
        throw new Error("No more charity donations configured");
    }
    return CONFIG.donations.charity[charityDonationIndex++];
}

async function main() {
    // Deploy contracts
    [owner, addr1, addr2, addr3, addr4, addr5, ...addrs] = await ethers.getSigners();
    
    signerMap = { owner, addr1, addr2, addr3, addr4, addr5 };
    
    ({
        cosmicGameProxy,
        cosmicToken,
        cosmicSignature,
        charityWallet,
        cosmicDAO,
        prizesWallet,
        randomWalkNFT,
        stakingWalletCosmicSignatureNft,
        stakingWalletRandomWalkNft,
        marketingWallet,
        implementationAddr,
    } = await basicDeployment(owner, "", 1, "", false, true));
    
    printContractAddresses();
    
    // Initial setup
    await setupGameConfiguration();
    await setActivationTimeToNow();
    await marketingWallet.connect(owner).setTreasurerAddress(addr3.address);
    
    // Deploy sample ERC20 tokens for testing
    const Samp = await hre.ethers.getContractFactory("Samp");
    const samp1 = await Samp.deploy();
    await samp1.waitForDeployment();
    const samp2 = await Samp.deploy();
    await samp2.waitForDeployment();
    
    console.log("=== ROUND 0 ===");
    
    // Initial donations
    await cosmicGameProxy.connect(addr5).donateEth({ value: getNextEthDonation() });
    
    let donation = getNextEthWithInfoDonation();
    await cosmicGameProxy.connect(addr4).donateEthWithInfo(donation.data, { value: donation.amount });
    
    donation = getNextEthWithInfoDonation();
    await cosmicGameProxy.connect(addr2).donateEthWithInfo(donation.data, { value: donation.amount });
    
    // Stake RandomWalk NFTs
    await mintAndStakeRwalkTokens(addr1, CONFIG.staking.rwalkTokens.addr1);
    await mintAndStakeRwalkTokens(addr2, CONFIG.staking.rwalkTokens.addr2);
    await mintAndStakeRwalkTokens(addr3, CONFIG.staking.rwalkTokens.addr3);
    
    // Place bids for round 0
    await placeBidsForRound('round0');
    
    // Fast-forward to prize time
    let prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime) - 100]);
    await ethers.provider.send("evm_mine");
    await ethers.provider.send("evm_increaseTime", [100]);
    await ethers.provider.send("evm_mine");
    
    // Claim round 0
    await cosmicGameProxy.connect(addr5).claimMainPrize({ gasLimit: 90000000 });
    await stakeAvailableNfts();
    
    console.log("=== ROUND 1 ===");
    
    await setActivationTimeToNow();
    await placeBidsForRound('round1');
    await advanceTimeAndClaim(addr1);
    
    console.log("=== ROUND 2 ===");
    
    await setActivationTimeToNow();
    await stakeAvailableNfts();
    
    // Place bids
    await placeBidsForRound('round2');
    const receipt2 = await advanceTimeAndClaim(addr1, 5000000);
    
    // Find minted NFT and set name
    const nftMintedSig = cosmicSignature.interface.getEvent("NftMinted").topicHash;
    const eventLogs = receipt2.logs.filter(log => log.topics[0] === nftMintedSig);
    
    for (const log of eventLogs) {
        const parsedLog = cosmicSignature.interface.parseLog(log);
        const args = parsedLog.args.toObject();
        if (args.nftOwnerAddress === addr1.address) {
            const tokenId = args.nftId;
            await cosmicSignature.connect(addr1).setNftName(tokenId, "name 0");
            await cosmicSignature.connect(addr1).setNftName(tokenId, "name after 0");
            break;
        }
    }
    
    await stakeAvailableNfts();
    await setActivationTimeToNow();
    
    // Charity withdrawals
    await charityWallet.connect(addr1).send();
    await addr2.sendTransaction({ to: await charityWallet.getAddress(), value: getNextCharityDonation() });
    await addr2.sendTransaction({ to: await charityWallet.getAddress(), value: getNextCharityDonation() });
    
    // Donate ERC20 tokens
    await donateErc20Token(samp1, CONFIG.erc20Donations.round2, "bid&donateerc20");
    await donateErc20Token(samp2, CONFIG.erc20Donations.round2, "bid&donateerc20");
    
    // Donate NFTs
    for (let i = 0; i < 4; i++) {
        const tokenId = await mintRwalk(signerMap[`addr${(i % 3) + 1}`]);
        await donateNft(signerMap[`addr${(i % 3) + 1}`], randomWalkNFT, tokenId, `donated token_id=${tokenId}`);
    }
    
    // Bid with CST
    await ethers.provider.send("evm_increaseTime", [36000]);
    await placeCstBid(addr3, "cst bid + donate1");
    await ethers.provider.send("evm_increaseTime", [36000]);
    await placeCstBid(addr3, "cst bid + donate2");
    await ethers.provider.send("evm_increaseTime", [36000]);
    await placeCstBid(addr3, "bid using ERC20 token");
    await placeCstBid(addr3, "bid using ERC20 token");
    
    // Set delay and claim
    await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(1000);
    await advanceTimeAndClaim(addr3);
    
    // Claim donated tokens
    const roundNum = await cosmicGameProxy.roundNum();
    try {
        await prizesWallet.connect(addr3).claimDonatedToken(roundNum - 1n, await samp1.getAddress(), CONFIG.erc20Donations.round2);
        await prizesWallet.connect(addr3).claimDonatedToken(roundNum - 1n, await samp2.getAddress(), CONFIG.erc20Donations.round2);
    } catch (e) {
        console.log("Error claiming donated tokens:", e.message);
    }
    
    await stakeAvailableNfts();
    
    // Update contract addresses
    await setupContractAddresses();
    
    // Update various settings (testing admin events)
    const ethDutchAuctionDivisor = await cosmicGameProxy.ethDutchAuctionDurationDivisor();
    await cosmicGameProxy.connect(owner).setEthDutchAuctionDurationDivisor(ethDutchAuctionDivisor + 11n);
    
    const ethEndingPriceDivisor = await cosmicGameProxy.ethDutchAuctionEndingBidPriceDivisor();
    await cosmicGameProxy.connect(owner).setEthDutchAuctionEndingBidPriceDivisor(ethEndingPriceDivisor + 40n);
    
    // Update NFT metadata URIs
    await cosmicSignature.connect(owner).setNftGenerationScriptUri("ipfs://");
    await cosmicSignature.connect(owner).setNftBaseUri("nttp://");
    
    // Upgrade contract (testing)
    const iAddrBytes = await ethers.provider.getStorage(
        await cosmicGameProxy.getAddress(),
        "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
    );
    const iAddr = await ethers.AbiCoder.defaultAbiCoder().decode(["address"], iAddrBytes).toString();
    await cosmicGameProxy.connect(owner).upgradeToAndCall(iAddr, "0x");
    
    // Test setting same values (edge case testing)
    await cosmicGameProxy.connect(owner).setStakingWalletCosmicSignatureNft(await stakingWalletCosmicSignatureNft.getAddress());
    await cosmicGameProxy.connect(owner).setStakingWalletRandomWalkNft(await stakingWalletRandomWalkNft.getAddress());
    
    const mainPrizeIncreaseDivisor = await cosmicGameProxy.mainPrizeTimeIncrementIncreaseDivisor();
    await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementIncreaseDivisor(mainPrizeIncreaseDivisor);
    
    const timeoutDuration = await cosmicGameProxy.connect(owner).timeoutDurationToClaimMainPrize();
    await cosmicGameProxy.connect(owner).setTimeoutDurationToClaimMainPrize(timeoutDuration);
    const mainPrizeTimeout = Number(timeoutDuration);
    
    const bidPriceIncrease = await cosmicGameProxy.ethBidPriceIncreaseDivisor();
    await cosmicGameProxy.connect(owner).setEthBidPriceIncreaseDivisor(bidPriceIncrease);
    
    const prizeTimeIncrement = await cosmicGameProxy.mainPrizeTimeIncrementInMicroSeconds();
    await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementInMicroSeconds(prizeTimeIncrement);
    
    const initialDuration = await cosmicGameProxy.initialDurationUntilMainPrizeDivisor();
    await cosmicGameProxy.connect(owner).setInitialDurationUntilMainPrizeDivisor(initialDuration);
    
    const delayDuration = await cosmicGameProxy.delayDurationBeforeRoundActivation();
    await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(delayDuration);
    
    const withdrawTimeout = await prizesWallet.timeoutDurationToWithdrawPrizes();
    await prizesWallet.connect(owner).setTimeoutDurationToWithdrawPrizes(Number(withdrawTimeout) / 2);
    
    console.log("Admin settings updated");
    
    await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(CONFIG.gameSettings.delayBeforeActivation);
    await ethers.provider.send("evm_increaseTime", [1001]);
    await ethers.provider.send("evm_mine");
    
    // Claim donated NFTs
    await prizesWallet.connect(addr3).claimDonatedNft(0n);
    await prizesWallet.connect(addr3).claimDonatedNft(1n);
    
    // Withdraw unclaimed deposits
    await withdrawAllUnclaimedDeposits();
    
    console.log("=== ROUND 3 ===");
    
    await placeBidsForRound('round3');
    await advanceTimeAndClaim(addr1, 90000000);
    await ethers.provider.send("evm_mine");
    await stakeAvailableNfts();
    
    // More donations
    donation = getNextEthWithInfoDonation();
    await cosmicGameProxy.donateEthWithInfo(donation.data, { value: donation.amount });
    donation = getNextEthWithInfoDonation();
    await cosmicGameProxy.donateEthWithInfo(donation.data, { value: donation.amount });
    
    // Pay marketing rewards
    await payMarketingRewards(addr3);
    
    // Unstake all NFTs
    await unstakeAllNfts();
    await ethers.provider.send("evm_mine");
    await ethers.provider.send("evm_mine");
    
    // Perform maintenance
    await stakingWalletCosmicSignatureNft.tryPerformMaintenance(owner.address);
    await stakeAvailableNfts();
    
    console.log("=== ROUND 4 ===");
    
    // Donate more ERC20
    await donateErc20Token(samp1, CONFIG.erc20Donations.round4, "bid&donateerc20");
    await donateErc20Token(samp2, CONFIG.erc20Donations.round4, "bid&donateerc20");
    
    // Final bids
    await placeBidsForRound('round4');
    
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    const finalReceipt = await advanceTimeAndClaim(addr3);
    
    // Claim one donated token (test partial claiming)
    const finalRoundNum = await cosmicGameProxy.roundNum();
    await prizesWallet.connect(addr3).claimDonatedToken(finalRoundNum - 1n, await samp1.getAddress(), CONFIG.erc20Donations.round4);
    await ethers.provider.send("evm_mine");
    
    // Stake newly minted tokens
    const oldTotalSupply = await cosmicSignature.totalSupply();
    const newTotalSupply = await cosmicSignature.totalSupply();
    
    for (let i = Number(oldTotalSupply); i < Number(newTotalSupply); i++) {
        const tokenOwner = await cosmicSignature.ownerOf(i);
        if (tokenOwner === (await stakingWalletCosmicSignatureNft.getAddress())) continue;
        
        const ownerSigner = await hre.ethers.getSigner(tokenOwner);
        try {
            await stakingWalletCosmicSignatureNft.connect(ownerSigner).stake(i);
        } catch (e) {
            // Ignore
        }
    }
    
    // Approve CST tokens for all addresses
    await approveCosmicToken([addr1, addr2, addr3, addr4]);
    
    console.log("=== ROUND 5 ===");
    
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await placeBidsForRound('round5');
    
    // Timeout scenario
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    await ethers.provider.send("evm_increaseTime", [mainPrizeTimeout + 1]);
    await ethers.provider.send("evm_mine");
    
    await cosmicGameProxy.connect(addr3).claimMainPrize({ gasLimit: 9000000 });
    
    console.log("=== ROUND 6 ===");
    
    await cosmicGameProxy.connect(addr3).donateEth({ value: getNextRegularDonation() });
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await placeBidsForRound('round6');
    await advanceTimeAndClaim(addr3);
    
    console.log("=== ROUND 7 ===");
    
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await placeBidsForRound('round7');
    await advanceTimeAndClaim(addr3);
    
    await ethers.provider.send("evm_mine");
    await ethers.provider.send("evm_mine");
    
    // Unstake RandomWalk NFTs
    for (let i = 1; i <= 5; i++) {
        await stakingWalletRandomWalkNft.connect(addr1).unstake(i);
    }
    
    console.log("=== Population Complete ===");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

