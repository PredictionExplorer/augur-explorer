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
    
    // Leave some deposits unclaimed for testing
    // Skip last 3 rounds to keep some unclaimed rewards
    const currentRound = Number(await cosmicGameProxy.roundNum());
    const roundsToSkip = [
        currentRound - 1,  // Skip last round
        currentRound - 2,  // Skip second-to-last round
        currentRound - 3,  // Skip third-to-last round
    ];
    
    for (const event of events) {
        const winnerAddress = event.args.prizeWinnerAddress;
        const roundNum = Number(event.args.roundNum);
        const withdrawalKey = `${winnerAddress}_${roundNum}`;
        
        if (withdrawalDone[withdrawalKey]) continue;
        
        // Skip withdrawals for the last 3 rounds to leave unclaimed rewards
        if (roundsToSkip.includes(roundNum)) {
            console.log(`Skipping withdrawal for round ${roundNum} (leaving unclaimed for testing)`);
            continue;
        }
        
        try {
            const winnerSigner = await hre.ethers.getSigner(winnerAddress);
            await prizesWallet.connect(winnerSigner).withdrawEth(roundNum);
            withdrawalDone[withdrawalKey] = true;
            console.log(`Withdrawn round ${roundNum} for ${winnerAddress}`);
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
    
    // Distribute Samp tokens (owner has all from deployment)
    console.log("Distributing Samp ERC20 tokens...");
    const sampAmount = hre.ethers.parseEther("50000");
    await samp1.connect(owner).transfer(addr1.address, sampAmount);
    await samp1.connect(owner).transfer(addr2.address, sampAmount);
    await samp1.connect(owner).transfer(addr3.address, sampAmount);
    await samp2.connect(owner).transfer(addr1.address, sampAmount);
    await samp2.connect(owner).transfer(addr2.address, sampAmount);
    await samp2.connect(owner).transfer(addr3.address, sampAmount);
    console.log("✓ Distributed 50000 samp1 and samp2 to addr1, addr2, addr3");
    
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
    
    // === ENHANCEMENT: CST Bid + Donation Combos ===
    // This generates events for CST bid combined with token/NFT donations
    // Frontend needs this to display "CST bid + donation" transactions
    await ethers.provider.send("evm_increaseTime", [36000]);
    const cstPriceForCombo = await cosmicGameProxy.getNextCstBidPrice();
    
    // Approve ERC20 tokens for PrizesWallet (needed for donation)
    await samp1.connect(addr3).approve(await prizesWallet.getAddress(), hre.ethers.parseEther("9999999999999999"));
    await samp1.connect(addr3).approve(await cosmicGameProxy.getAddress(), hre.ethers.parseEther("9999999999999999"));
    
    console.log("Testing CST bid + ERC20 donation combo");
    
    // DEBUG: Check balances before bidding
    const addr3CstBal = await cosmicToken.balanceOf(addr3.address);
    const addr3Samp1Bal = await samp1.balanceOf(addr3.address);
    const cstBidPrice = await cosmicGameProxy.getNextCstBidPrice();
    console.log(`DEBUG addr3 CST balance: ${hre.ethers.formatEther(addr3CstBal)} CST`);
    console.log(`DEBUG CST bid price: ${hre.ethers.formatEther(cstBidPrice)} CST`);
    console.log(`DEBUG addr3 samp1 balance: ${hre.ethers.formatEther(addr3Samp1Bal)} samp1`);
    console.log(`DEBUG Can afford? CST=${addr3CstBal >= cstBidPrice}, samp1=${addr3Samp1Bal >= hre.ethers.parseEther("75")}`);
    
    await cosmicGameProxy.connect(addr3).bidWithCstAndDonateToken(
        cstPriceForCombo * 2n,  // priceMaxLimit
        "CST bid with ERC20 donation",
        await samp1.getAddress(),
        hre.ethers.parseEther("75")
    );
    
    await ethers.provider.send("evm_increaseTime", [36000]);
    const nftIdForCstDonation = await mintRwalk(addr3);
    await randomWalkNFT.connect(addr3).setApprovalForAll(await prizesWallet.getAddress(), true);
    
    console.log("Testing CST bid + NFT donation combo");
    await cosmicGameProxy.connect(addr3).bidWithCstAndDonateNft(
        cstPriceForCombo * 2n,  // priceMaxLimit
        "CST bid with NFT donation",
        await randomWalkNFT.getAddress(),
        nftIdForCstDonation
    );
    // === END: CST Bid + Donation Combos ===
    
    // Set delay and claim
    await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(1000);
    await advanceTimeAndClaim(addr3);
    
    // Claim donated tokens (only claim samp1, leave samp2 for timeout test)
    const roundNum = await cosmicGameProxy.roundNum();
    try {
        await prizesWallet.connect(addr3).claimDonatedToken(roundNum - 1n, await samp1.getAddress(), CONFIG.erc20Donations.round2);
        // NOTE: NOT claiming samp2 from Round 3 - leaving it for addr5 to claim after timeout
        console.log("Claimed samp1 from Round 3, left samp2 unclaimed for timeout test");
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
    
    // Claim donated NFTs (only claim 0 and 1, leave 2, 3, 4 for timeout test)
    await prizesWallet.connect(addr3).claimDonatedNft(0n);
    await prizesWallet.connect(addr3).claimDonatedNft(1n);
    // NOTE: NFT indexes 2, 3, 4 are intentionally LEFT UNCLAIMED
    // so addr5 can claim them after timeout (see timeout claiming section)
    
    // Withdraw unclaimed deposits
    await withdrawAllUnclaimedDeposits();
    
    // === ENHANCEMENT: Timeout-Based Claiming ===
    // This generates events where beneficiaryAddress ≠ prizeWinnerAddress
    // Frontend needs this to display "Claimed by someone else after timeout"
    console.log("\n=== Testing Timeout-Based Prize Claiming ===");
    
    const currentRoundNum = Number(await cosmicGameProxy.roundNum());
    const withdrawalTimeout = await prizesWallet.timeoutDurationToWithdrawPrizes();
    
    // Get the timeout time for round 1 (earliest unclaimed round)
    const round1TimeoutTime = await prizesWallet.roundTimeoutTimesToWithdrawPrizes(1);
    const currentBlockTime = (await ethers.provider.getBlock('latest')).timestamp;
    const timeToWait = Number(round1TimeoutTime) - currentBlockTime + 100;
    
    console.log(`Round 1 timeout expires at: ${round1TimeoutTime}`);
    console.log(`Current blockchain time: ${currentBlockTime}`);
    console.log(`Need to wait: ${timeToWait} seconds`);
    
    if (timeToWait > 0) {
        await ethers.provider.send("evm_increaseTime", [timeToWait]);
        await ethers.provider.send("evm_mine");
        console.log(`✓ Fast-forwarded past timeout`);
    }
    
    // Find unclaimed ETH deposits (withdrawAllUnclaimedDeposits skipped last 3 rounds)
    // Let addr5 claim prizes that belong to other winners (after timeout)
    const filter = prizesWallet.filters.EthReceived();
    const ethReceivedEvents = await prizesWallet.queryFilter(filter);
    
    let claimedByNonWinner = 0;
    for (const event of ethReceivedEvents) {
        const eventRound = Number(event.args.roundNum);
        const originalWinner = event.args.prizeWinnerAddress;
        
        // Check if this prize is still unclaimed
        const balance = await prizesWallet["getEthBalanceAmount(uint256,address)"](eventRound, originalWinner);
        
        if (balance > 0n && originalWinner !== addr5.address) {
            // addr5 (non-winner) claims the prize after timeout
            console.log(`addr5 claiming round ${eventRound} prize (original winner: ${originalWinner.substring(0, 10)}...)`);
            try {
                await prizesWallet.connect(addr5)["withdrawEth(uint256,address)"](eventRound, originalWinner);
                console.log(`  ✓ Generated EthWithdrawn event: beneficiary=${addr5.address.substring(0, 10)}... ≠ winner=${originalWinner.substring(0, 10)}...`);
                claimedByNonWinner++;
                if (claimedByNonWinner >= 2) break; // Test 2 examples
            } catch (e) {
                console.log(`  ✗ Could not withdraw: ${e.message}`);
            }
        }
    }
    
    // Test timeout claiming for donated ERC20 tokens
    // Round 3 has ERC20 donations, need to wait for Round 3's timeout
    const round3TimeoutTime = await prizesWallet.roundTimeoutTimesToWithdrawPrizes(3);
    const currentTime = (await ethers.provider.getBlock('latest')).timestamp;
    const timeToWaitR3 = Number(round3TimeoutTime) - currentTime + 100;
    
    console.log(`\nRound 3 timeout (for NFT/ERC20) expires at: ${round3TimeoutTime}`);
    console.log(`Current time: ${currentTime}`);
    console.log(`Additional time to wait for Round 3: ${timeToWaitR3} seconds`);
    
    if (timeToWaitR3 > 0) {
        await ethers.provider.send("evm_increaseTime", [timeToWaitR3]);
        await ethers.provider.send("evm_mine");
        console.log(`✓ Fast-forwarded past Round 3 timeout`);
    }
    
    // Now claim ERC20 donations from Round 3 (samp2 was left unclaimed)
    console.log("\nAttempting ERC20 timeout claims...");
    for (let r = 2; r <= 3; r++) {
        const mainWinner = await prizesWallet.mainPrizeBeneficiaryAddresses(r);
        console.log(`  Round ${r}: main winner = ${mainWinner.substring(0, 10)}...`);
        
        if (mainWinner !== ethers.ZeroAddress && mainWinner !== addr5.address) {
            try {
                const samp2Balance = await prizesWallet.getDonatedTokenBalanceAmount(r, await samp2.getAddress());
                console.log(`  Round ${r} samp2 balance: ${hre.ethers.formatEther(samp2Balance)} samp2`);
                
                if (samp2Balance > 0n) {
                    console.log(`addr5 claiming round ${r} donated samp2 (original winner: ${mainWinner.substring(0, 10)}...)`);
                    await prizesWallet.connect(addr5).claimDonatedToken(r, await samp2.getAddress(), 0n);
                    console.log(`  ✓ SUCCESS! DonatedTokenClaimed event: round=${r}, beneficiary=addr5 ≠ winner=${mainWinner.substring(0, 10)}...`);
                    break; // Only test once
                } else {
                    console.log(`  Round ${r} samp2 already claimed, checking next round...`);
                }
            } catch (e) {
                console.log(`  ✗ ERC20 claim failed: ${e.message.substring(0, 100)}`);
            }
        }
    }
    
    // Test timeout claiming for donated NFTs from Round 3
    const nextNftIndex = await prizesWallet.nextDonatedNftIndex();
    console.log(`Checking ${nextNftIndex} donated NFTs for timeout claims...`);
    
    for (let i = 2; i < nextNftIndex; i++) {
        try {
            const nft = await prizesWallet.donatedNfts(i);
            const nftObj = nft.toObject();
            
            if (nftObj.nftAddress !== ethers.ZeroAddress) {
                const nftRound = Number(nftObj.roundNum);
                const mainWinner = await prizesWallet.mainPrizeBeneficiaryAddresses(nftRound);
                
                if (mainWinner !== addr5.address) {
                    console.log(`addr5 claiming donated NFT index ${i} from ROUND ${nftRound} (original winner: ${mainWinner.substring(0, 10)}...)`);
                    try {
                        await prizesWallet.connect(addr5).claimDonatedNft(i);
                        console.log(`  ✓ SUCCESS! DonatedNftClaimed event: round=${nftRound}, beneficiary=addr5 ≠ winner=${mainWinner.substring(0, 10)}...`);
                        break; // Only test once (we got what we needed)
                    } catch (claimError) {
                        console.log(`  ✗ NFT claim failed: ${claimError.message.substring(0, 100)}`);
                    }
                }
            }
        } catch (e) {
            console.log(`  ✗ Error checking NFT ${i}: ${e.message.substring(0, 80)}`);
        }
    }
    
    console.log("=== Timeout-Based Claiming Complete ===\n");
    // === END: Timeout-Based Claiming ===
    
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
    
    // Unstake only 2 NFTs to test collection, leaving most staked with unclaimed rewards
    const numActions = await stakingWalletCosmicSignatureNft.actionCounter();
    const numToUnstake = 2; // Only unstake 2 NFTs, leaving the rest with unclaimed rewards
    
    for (let i = 1; i <= numToUnstake && i <= numActions; i++) {
        const actionRec = (await stakingWalletCosmicSignatureNft.stakeActions(i)).toObject();
        const owner = actionRec.nftOwnerAddress;
        
        if (owner === "0x0000000000000000000000000000000000000000") continue;
        
        try {
            const ownerSigner = await hre.ethers.getSigner(owner);
            await stakingWalletCosmicSignatureNft.connect(ownerSigner).unstake(i);
            console.log(`Unstaked action ${i} - rewards will be marked as collected`);
        } catch (e) {
            // Already unstaked or error
        }
    }
    console.log(`Unstaked only ${numToUnstake} out of ${numActions} staking actions`);
    
    await ethers.provider.send("evm_mine");
    await ethers.provider.send("evm_mine");
    
    // Skip maintenance call since we're leaving NFTs staked (tryPerformMaintenance requires all NFTs unstaked)
    // await stakingWalletCosmicSignatureNft.tryPerformMaintenance(owner.address);
    console.log("Skipped tryPerformMaintenance - NFTs remain staked with unclaimed rewards");
    
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
    
    // === ENHANCEMENT: Batch Operations ===
    // Generate events showing multiple operations in single transaction
    // Frontend needs this to display batch actions properly
    console.log("\n=== Testing Batch Operations ===");
    
    // Test batch ETH withdrawal (withdrawEthMany)
    try {
        const roundsWithPrizes = [];
        const finalRoundNum = Number(await cosmicGameProxy.roundNum());
        
        // Find rounds where addr1 has unclaimed ETH
        for (let r = 0; r < finalRoundNum; r++) {
            const balance = await prizesWallet["getEthBalanceAmount(uint256,address)"](r, addr1.address);
            if (balance > 0n) {
                roundsWithPrizes.push(r);
            }
        }
        
        if (roundsWithPrizes.length >= 2) {
            console.log(`addr1 batch-withdrawing ETH from ${roundsWithPrizes.length} rounds: ${roundsWithPrizes}`);
            await prizesWallet.connect(addr1).withdrawEthMany(roundsWithPrizes);
            console.log(`  ✓ Generated ${roundsWithPrizes.length} EthWithdrawn events in ONE transaction`);
        } else {
            console.log(`  ⚠️ addr1 has prizes in only ${roundsWithPrizes.length} round(s), skipping batch test`);
        }
    } catch (e) {
        console.log(`  ✗ Batch ETH withdrawal error: ${e.message}`);
    }
    
    // Test batch donated NFT claiming (claimManyDonatedNfts)
    // DISABLED: Conflicts with timeout claiming test
    // If we batch-claim with mainWinner, we can't test timeout claiming by non-winner
    console.log(`  ⚠️ Batch NFT claiming skipped (NFTs left for timeout test)`);
    
    console.log("=== Batch Operations Complete ===\n");
    // === END: Batch Operations ===
    
    console.log("=== Population Complete ===");
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });

