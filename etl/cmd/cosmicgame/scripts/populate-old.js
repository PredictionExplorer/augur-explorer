const hre = require("hardhat");
const {
    basicDeployment
} = require("./Deploy.js");
const bidParamsEncoding = {
    type: "tuple(string,int256)",
    name: "bidparams",
    components: [{
            name: "msg",
            type: "string"
        },
        {
            name: "rwalk",
            type: "int256"
        },
    ],
};
async function main() {
    async function mint_rwalk(a) {
            tokenPrice = await randomWalkNFT.getMintPrice();
            let tx = await randomWalkNFT.connect(a).mint({
                value: tokenPrice
            });
            let receipt = await tx.wait();
            let topic_sig = randomWalkNFT.interface.getEvent("MintEvent").topicHash;
            let log = receipt.logs.find((x) => x.topics.indexOf(topic_sig) >= 0);
            let parsed_log = randomWalkNFT.interface.parseLog(log);
            let token_id = parsed_log.args[0];
            return token_id;
        }
	async function stake_available_nfts() {
		let tscst = await cosmicSignature.totalSupply();
		for (let i = 0; i < tscst; i++) {
			let ownr = await cosmicSignature.ownerOf(i);
			if (ownr == (await stakingWalletCosmicSignatureNft.getAddress())) {
				continue; // already staked
			}
			let owner_signer = await hre.ethers.getSigner(ownr);
			if (owner_signer === undefined) {
			} else {
			}
			try {
	            cosmicSignature.connect(owner_signer).setApprovalForAll(await stakingWalletCosmicSignatureNft.getAddress(), true);
			} catch (e) {
			}
			try {
				await stakingWalletCosmicSignatureNft.connect(owner_signer).stake(i);
			} catch (e) {
				//console.log("stake_available_nfts() stake() error: ",e);
			}
		}
	}
	async function unstake_all_nfts() {
		let num_a = await stakingWalletCosmicSignatureNft.actionCounter();
		let num_unstaked = 0;
		for (let i = 1; i <= num_a; i++) {
			let action_rec = (await stakingWalletCosmicSignatureNft.stakeActions(i)).toObject();
			let ownr = action_rec.nftOwnerAddress;
			let num_s = await stakingWalletCosmicSignatureNft.numStakedNfts();
			if (ownr == "0x0000000000000000000000000000000000000000") {
				continue; 
			}
			let owner_signer = await hre.ethers.getSigner(ownr);
			await stakingWalletCosmicSignatureNft.connect(owner_signer).unstake(i);
			num_unstaked=num_unstaked+1;
		}
	}
        [owner, addr1, addr2, addr3, addr4, addr5, ...addrs] =
        await ethers.getSigners();
    const {
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
    } = await basicDeployment(owner, "", 1, "", false, true);
    console.log("Addresses set");
    console.log(
        "INSERT INTO cg_contracts VALUES('" +
        (await cosmicGameProxy.getAddress()) +
        "','" +
        (await cosmicSignature.getAddress()) +
        "','" +
        (await cosmicToken.getAddress()) +
        "','" +
        (await cosmicDAO.getAddress()) +
        "','" +
        (await charityWallet.getAddress()) +
        "','" +
        (await prizesWallet.getAddress()) +
        "','" +
        (await randomWalkNFT.getAddress()) +
        "','" +
        (await stakingWalletCosmicSignatureNft.getAddress()) +
        "','" +
        (await stakingWalletRandomWalkNft.getAddress()) +
        "','" +
        (await marketingWallet.getAddress()) +
        "','" +
        implementationAddr +
        "')"
    );
	let tx;
	tx = await cosmicGameProxy.connect(owner).setTimeoutDurationToClaimMainPrize(120,{gasLimit:1000000});
	await tx.wait()
	tx = await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementInMicroSeconds(300000000,{gasLimit:1000000});
    await tx.wait()
	tx = await cosmicGameProxy.connect(owner).setInitialDurationUntilMainPrizeDivisor(1000000,{gasLimit:1000000});
	await tx.wait()
	let latestBlock = await hre.ethers.provider.getBlock("latest");
	let roundActivationTime = latestBlock.timestamp - 1;
	console.log("Setting activartion time to "+roundActivationTime)
	await (await cosmicGameProxy.setRoundActivationTime(roundActivationTime)).wait();
	let atime = await cosmicGameProxy.roundActivationTime();
	console.log("activation time after setting it:",atime);

	let taddr = await marketingWallet.treasurerAddress();
	console.log("treasurer:",taddr);
	console.log("owner:",owner.address);
	await marketingWallet.connect(owner).setTreasurerAddress(addr3.address);

	const Samp = await hre.ethers.getContractFactory("Samp");
	const samp1 = await Samp.deploy();
	await samp1.waitForDeployment();
	let samp1Addr = await samp1.getAddress();
	const samp2 = await Samp.deploy();
	await samp2.waitForDeployment();
	let samp2Addr = await samp2.getAddress();

    let donationAmount = hre.ethers.parseEther("100");
    await cosmicGameProxy.connect(addr5).donateEth({
        value: donationAmount
    });
    let donationData =
        '{"version":1,"title":"Hardhat donation","message":"Donation from HardHat","url":"http://hardhat.org"}';
    await cosmicGameProxy.connect(addr4).donateEthWithInfo(donationData, {
        value: hre.ethers.parseEther("60"),
    });
    donationData =
        '{"version":1,"title":"ArtBlocks donation","message":"ArtBlocks offers a platform for creators, buyers and sellers of digital assets and any non-digital products, services and/or benefits to be furnished by or on behalf of sellers in connection with such sales","url":"https://www.artblocks.io"}';

    await cosmicGameProxy
        .connect(addr2)
        .donateEthWithInfo(donationData, {
            value: hre.ethers.parseEther("90")
        });

	let numStakeActions = 5
    for (let i = 0; i < numStakeActions; i++) {
        let token_id = await mint_rwalk(addr1);
        await randomWalkNFT
            .connect(addr1)
            .setApprovalForAll(await stakingWalletRandomWalkNft.getAddress(), true);
        await stakingWalletRandomWalkNft.connect(addr1).stake(token_id);
    }
    for (let i = 0; i < 5; i++) {
        let token_id = await mint_rwalk(addr2);
        await randomWalkNFT
            .connect(addr2)
            .setApprovalForAll(await stakingWalletRandomWalkNft.getAddress(), true);
        await stakingWalletRandomWalkNft.connect(addr2).stake(token_id);
    }
    for (let i = 0; i < 50; i++) {
        let token_id = await mint_rwalk(addr3);
        await randomWalkNFT
            .connect(addr3)
            .setApprovalForAll(await stakingWalletRandomWalkNft.getAddress(), true);
        await stakingWalletRandomWalkNft.connect(addr3).stake(token_id);
    }

    let prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    console.log("Donation complete");

    const contractBalance = await ethers.provider.getBalance(
        await cosmicGameProxy.getAddress()
    );
    let bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 1", { value: bidPrice + 1000n }); // this works
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr2).bidWithEth(-1,"bid 1", { value: bidPrice + 1000n }); // this works

//    let nanoSecondsExtra = await cosmicGameProxy.nanoSecondsExtra();
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 2", { value: bidPrice });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr1).bidWithEth(-1, "bid 2", { value: bidPrice });
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    let token_id = await mint_rwalk(owner);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(owner).bidWithEth(Number(token_id),"bidWithRWlk", {value: bidPrice });

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {
        value: bidPrice
    });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {
        value: bidPrice
    });

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr4).bidWithEth(-1,"", { value: bidPrice });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr4).bidWithEth(-1,"", { value: bidPrice });

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr5).bidWithEth(-1,"", {value: bidPrice });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr5).bidWithEth(-1,"", {value: bidPrice });

    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime) - 100]);
    await ethers.provider.send("evm_mine");

    await ethers.provider.send("evm_increaseTime", [100]);
    await ethers.provider.send("evm_mine");

    let prizeAmount = await cosmicGameProxy.getMainEthPrizeAmount();
    let charityAmount = await cosmicGameProxy.getCharityEthDonationAmount();
    await cosmicGameProxy.connect(addr5).claimMainPrize({
        gasLimit: 30000000
    });
    let prizeAmount2 = await cosmicGameProxy.getMainEthPrizeAmount();
    let expectedprizeAmount = (prizeAmount - charityAmount) / 2n;

	await stake_available_nfts()
	latestBlock = await hre.ethers.provider.getBlock("latest");
	roundActivationTime = latestBlock.timestamp - 1;
	await (await cosmicGameProxy.setRoundActivationTime(roundActivationTime)).wait();
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 4", {
        value: bidPrice
    });

    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();

    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");

    prizeAmount = await cosmicGameProxy.getMainEthPrizeAmount();
    charityAmount = await cosmicGameProxy.getCharityEthDonationAmount();
    await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 3000000
    });

	latestBlock = await hre.ethers.provider.getBlock("latest");
	roundActivationTime = latestBlock.timestamp - 1;
	await (await cosmicGameProxy.setRoundActivationTime(roundActivationTime)).wait();
    prizeAmount2 = await cosmicGameProxy.getMainEthPrizeAmount();
	await stake_available_nfts();
    let ts = await cosmicSignature.totalSupply();
    let rn = await cosmicGameProxy.roundNum();
    let oldTotalSupply = ts;

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 5", { value: bidPrice });
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 5000000
    });
    receipt = await tx.wait();
    topic_sig = cosmicSignature.interface.getEvent("NftMinted").topicHash;
	let event_logs = receipt.logs.filter(log => log.topics[0] === topic_sig);
	let mint_found = false;
	for (let i=0; i< event_logs.length; i++) {
		let parsed_log = cosmicSignature.interface.parseLog(event_logs[i]);
	    let args = parsed_log.args.toObject();
		if (args.nftOwnerAddress === addr1.address) {
		  	token_id = args.nftId;
	    	await cosmicSignature.connect(addr1).setNftName(token_id, "name 0");
	    	await cosmicSignature.connect(addr1).setNftName(token_id, "name after 0");
			mint_found = true;
		}
		if (mint_found) {
			break;
		}
	}
	if (!mint_found) {
		console.log("No NFT was minted for the winner. Bug!")
		return;
	}
	await stake_available_nfts();

	latestBlock = await hre.ethers.provider.getBlock("latest");
	roundActivationTime = latestBlock.timestamp - 1;
	await (await cosmicGameProxy.setRoundActivationTime(roundActivationTime)).wait();

    await charityWallet.connect(addr1).send();

    tx = {
        to: await charityWallet.getAddress(),
        value: hre.ethers.parseEther("4"),
    };
    await addr2.sendTransaction(tx);
    await addr2.sendTransaction(tx);

    rn = await cosmicGameProxy.roundNum();
	await samp1.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await samp1.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp1.getAddress(),10000000000000000000n,{value:bidPrice});
	await samp2.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await samp2.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp2.getAddress(),10000000000000000000n,{value:bidPrice});

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 6", { value: bidPrice });
    await ethers.provider.send("evm_mine"); // mine empty block as spacing

    await randomWalkNFT
        .connect(addr1)
        .setApprovalForAll(await cosmicGameProxy.getAddress(), true);
    await randomWalkNFT
        .connect(addr2)
        .setApprovalForAll(await cosmicGameProxy.getAddress(), true);
    await randomWalkNFT
        .connect(addr3)
        .setApprovalForAll(await cosmicGameProxy.getAddress(), true);

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    token_id = await mint_rwalk(addr1);
	await randomWalkNFT.connect(addr1).setApprovalForAll(await prizesWallet.getAddress(), true);
    await cosmicGameProxy
        .connect(addr1)
        .bidWithEthAndDonateNft(-1,"donated token_id="+token_id,await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    token_id = await mint_rwalk(addr2);
	await randomWalkNFT.connect(addr2).setApprovalForAll(await prizesWallet.getAddress(), true);
    await cosmicGameProxy
        .connect(addr2)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    token_id = await mint_rwalk(addr3);
	await randomWalkNFT.connect(addr3).setApprovalForAll(await prizesWallet.getAddress(), true);
    await cosmicGameProxy
        .connect(addr3)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    token_id = await mint_rwalk(addr3);
	await randomWalkNFT.connect(addr3).setApprovalForAll(await prizesWallet.getAddress(), true);
    await cosmicGameProxy
        .connect(addr3)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
    await ethers.provider.send("evm_increaseTime", [36000]);
    await ethers.provider.send("evm_mine");
    token_id = await mint_rwalk(addr3);
	let cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    await cosmicGameProxy
        .connect(addr3)
        .bidWithCstAndDonateNft(cstPrice,"cst bid + donate1", await randomWalkNFT.getAddress(), token_id);
    await ethers.provider.send("evm_increaseTime", [36000]);
    await ethers.provider.send("evm_mine");
    token_id = await mint_rwalk(addr3);
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    await cosmicGameProxy
        .connect(addr3)
        .bidWithCstAndDonateNft(cstPrice,"cst bid + donate2", await randomWalkNFT.getAddress(), token_id);
    await ethers.provider.send("evm_increaseTime", [36000]);
    await ethers.provider.send("evm_mine");
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"bid using ERC20 token");
    await ethers.provider.send("evm_mine");
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"bid using ERC20 token");
    await ethers.provider.send("evm_mine");

	await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(1000);

    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
    receipt = await tx.wait();
	try {
		prizesWallet.connect(addr3).claimDonatedToken(rn,await samp1.getAddress(),10000000000000000000n);
	} catch (eee) {
		console.log(eee);
	}
	try {
		prizesWallet.connect(addr3).claimDonatedToken(rn,await samp2.getAddress(),10000000000000000000n);
	} catch (eee) {
		console.log(eee);
	}
	await stake_available_nfts();

    await cosmicGameProxy
        .connect(owner)
        .setCharityAddress(await charityWallet.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setRandomWalkNft(await randomWalkNFT.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setPrizesWallet(await prizesWallet.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletCosmicSignatureNft(await stakingWalletCosmicSignatureNft.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletRandomWalkNft(await stakingWalletRandomWalkNft.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setMarketingWallet(await marketingWallet.getAddress());
    await cosmicGameProxy.connect(owner).setNumRaffleEthPrizesForBidders(4);
    await cosmicGameProxy.connect(owner).setNumRaffleCosmicSignatureNftsForBidders(6);
    await cosmicGameProxy.connect(owner).setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(3);
    await cosmicGameProxy.connect(owner).setMainEthPrizeAmountPercentage(30);
    await cosmicGameProxy.connect(owner).setCharityEthDonationAmountPercentage(5);
    await cosmicGameProxy.connect(owner).setRaffleTotalEthPrizeAmountForBiddersPercentage(6);
    await cosmicGameProxy.connect(owner).setCharityAddress(addr3.address);
	await cosmicGameProxy.connect(owner).setChronoWarriorEthPrizeAmountPercentage(8);
    await cosmicGameProxy
        .connect(owner)
        .setCharityAddress(await charityWallet.getAddress());
    await cosmicGameProxy.connect(owner).setCosmicSignatureNftStakingTotalEthRewardAmountPercentage(19);
    await cosmicGameProxy
        .connect(owner)
        .setCosmicSignatureToken(await cosmicToken.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setCosmicSignatureNft(await cosmicSignature.getAddress());
    await cosmicGameProxy.connect(owner).setCstDutchAuctionDurationDivisor(13 * 3600);
	let tmpval = await cosmicGameProxy.ethDutchAuctionDurationDivisor();
	tmpval = tmpval + 11n;
	await cosmicGameProxy.connect(owner).setEthDutchAuctionDurationDivisor(tmpval);
	tmpval = await cosmicGameProxy.ethDutchAuctionEndingBidPriceDivisor();
	tmpval = tmpval + 40n;
	await cosmicGameProxy.connect(owner).setEthDutchAuctionEndingBidPriceDivisor(tmpval);
    await cosmicGameProxy.connect(owner).setCstPrizeAmountMultiplier(5);
    await cosmicGameProxy
        .connect(owner)
        .setCstDutchAuctionBeginningBidPriceMinLimit(150000000000000000000n);
    await cosmicGameProxy
        .connect(owner)
        .setMarketingWalletCstContributionAmount(120000000000000000000n);
    await cosmicGameProxy.connect(owner).setCstRewardAmountForBidding(130000000000000000000n);
    await cosmicGameProxy.connect(owner).setBidMessageLengthMaxLimit(199);
    await cosmicSignature.connect(owner).setNftGenerationScriptUri("ipfs://");
    await cosmicSignature.connect(owner).setNftBaseUri("nttp://");
    let iAddrBytes = await ethers.provider.getStorage(
        await cosmicGameProxy.getAddress(),
        "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
    );
    let iAddr = await ethers.AbiCoder.defaultAbiCoder()
        .decode(["address"], iAddrBytes)
        .toString();
    await cosmicGameProxy.connect(owner).upgradeToAndCall(iAddr,"0x");
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletCosmicSignatureNft(await stakingWalletCosmicSignatureNft.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletRandomWalkNft(await stakingWalletRandomWalkNft.getAddress());
    await cosmicGameProxy.connect(owner).setCstPrizeAmountMultiplier(999);
    let tmp = await cosmicGameProxy.mainPrizeTimeIncrementIncreaseDivisor();
    await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementIncreaseDivisor(tmp);
    tmp = await cosmicGameProxy.connect(owner).timeoutDurationToClaimMainPrize();
    await cosmicGameProxy.connect(owner).setTimeoutDurationToClaimMainPrize(tmp);
	let mainPrizeTimeout = Number(tmp);
    tmp = await cosmicGameProxy.ethBidPriceIncreaseDivisor();
    await cosmicGameProxy.connect(owner).setEthBidPriceIncreaseDivisor(tmp);
    tmp = await cosmicGameProxy.mainPrizeTimeIncrementInMicroSeconds();
    await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementInMicroSeconds(tmp);
    tmp = await cosmicGameProxy.initialDurationUntilMainPrizeDivisor();
    await cosmicGameProxy.connect(owner).setInitialDurationUntilMainPrizeDivisor(tmp);
    tmp = await cosmicGameProxy.delayDurationBeforeRoundActivation();
    await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(tmp);
	tmp = await prizesWallet.timeoutDurationToWithdrawPrizes();
	await prizesWallet.connect(owner).setTimeoutDurationToWithdrawPrizes(Number(tmp)/2);

	await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(1);
    await ethers.provider.send("evm_increaseTime", [1001]);
    await ethers.provider.send("evm_mine");

    await prizesWallet.connect(addr3).claimDonatedNft(0n);
    await prizesWallet.connect(addr3).claimDonatedNft(1n);
    topic_sig = prizesWallet.interface.getEvent("EthReceived").topicHash;
    deposit_logs = receipt.logs.filter((x) => x.topics.indexOf(topic_sig) >= 0);
    let withdrawal_done = [];
    for (let i = 0; i < deposit_logs.length; i++) {
        let wlog = prizesWallet.interface.parseLog(deposit_logs[i]);
        let winner_signer = await hre.ethers.getSigner(wlog.args.prizeWinnerAddress);
        if (typeof withdrawal_done[wlog.args.winner] === "undefined") {
            await prizesWallet.connect(winner_signer).withdrawEth();
            withdrawal_done[wlog.args.winner] = 1;
        } else {
            // skip
        }
    }

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr1).bidWithEth(-1,"", { value: bidPrice  });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr1).bidWithEth(-1,"", { value: bidPrice  });
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 30000000
    });

    await ethers.provider.send("evm_mine");
	await stake_available_nfts()

    donationData =
        '{"version":1,"title":"EF donation","message":"Ethereum Foundation is a non-profit and part of a community of organizations and people working to fund protocol development, grow the ecosystem, and advocate for Ethereum.","url":"http://ethereum.org/en"}';
    await cosmicGameProxy.donateEthWithInfo(donationData, {
        value: hre.ethers.parseEther("9"),
    });
    await cosmicGameProxy.donateEthWithInfo(donationData, {
        value: hre.ethers.parseEther("8"),
    });

    await marketingWallet.connect(addr3).payReward(addr1.address,hre.ethers.parseEther("7"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.connect(addr3).payReward(addr2.address,hre.ethers.parseEther("7"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.connect(addr3).payReward(addr2.address,hre.ethers.parseEther("2"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.connect(addr3).payReward(addr1.address,hre.ethers.parseEther("6"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.connect(addr3).payReward(addr2.address,hre.ethers.parseEther("5"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.connect(addr3).payReward(addr2.address,hre.ethers.parseEther("5"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.connect(addr3).payReward(addr3.address,hre.ethers.parseEther("5"));
    await marketingWallet.connect(addr3).payReward(addr4.address,hre.ethers.parseEther("1"));
    await marketingWallet.connect(addr3).payReward(addr1.address,hre.ethers.parseEther("11"));
	await unstake_all_nfts()
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
	await stakingWalletCosmicSignatureNft.tryPerformMaintenance(owner.address)
	await stake_available_nfts()

	await samp1.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await samp1.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp1.getAddress(),11000000000000000000n,{value:bidPrice});
	await samp2.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await samp2.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp2.getAddress(),11000000000000000000n,{value:bidPrice});

    // generate one deposit to charity and not to Staking Wallet
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3",{value: bidPrice});
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {value: bidPrice });
	rn = cosmicGameProxy.roundNum();
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });

	prizesWallet.connect(addr3).claimDonatedToken(rn,await samp1.getAddress(),11000000000000000000n);	// only claim one of the tokens (samp1, not samp2)
    await ethers.provider.send("evm_mine"); // mine empty block as spacing

    ts = await cosmicSignature.totalSupply();

    for (let i = Number(oldTotalSupply); i < Number(ts); i++) {
        let ownr = await cosmicSignature.ownerOf(i);
        if (ownr == (await stakingWalletCosmicSignatureNft.getAddress())) {
            continue;
        }
        let owner_signer = await hre.ethers.getSigner(ownr);
        try {
            await stakingWalletCosmicSignatureNft.connect(owner_signer).stake(i);
        } catch (e) {
            //console.log("ignoring stake() error for token " + i + ", owner " + ownr);
        }
    }

    await cosmicToken
        .connect(addr1)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
    await cosmicToken
        .connect(addr2)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
    await cosmicToken
        .connect(addr3)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
    await cosmicToken
        .connect(addr4)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );

    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr2).bidWithEth(-1,"opening bid (addr2)", {value: bidPrice });
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    await cosmicGameProxy.connect(addr1).bidWithCst(cstPrice,"CST bid (addr1)");
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
//    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
//    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid (addr3)", {value: bidPrice });
//    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [mainPrizeTimeout + 1]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });

    donationAmount = hre.ethers.parseEther("500");
    await cosmicGameProxy.connect(addr3).donateEth({
        value: donationAmount
    });
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice });
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"CST bid addr1");
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice });
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"CST bid addr1");
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await ethers.provider.send("evm_mine"); // mine empty block as spacing

    for (let i = 1; i <= 5; i++) {
        await stakingWalletRandomWalkNft.connect(addr1).unstake(i);
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
