const { ethers } = require("hardhat");
require("./rpc-helpers.js");
async function customGetSigners() {

	const privateKeys = [
		"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		"0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
		"0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
		"0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
		"0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
		"0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba"
	];
	const signers = privateKeys.map(key => new ethers.Wallet(key, ethers.provider));
//	[owner, addr1, addr2, addr3, addr4, addr5, ...addrs] = await ethers.getSigners();
	return [signers[0],signers[1],signers[2],signers[3],signers[4],signers[5]];

}
async function getCosmicSignatureGameContract(cosmicSignatureGameContractName = "CosmicSignatureGame") {
    const cosmicSignatureGameAddr = process.env.CADDR;
    if (typeof cosmicSignatureGameAddr === "undefined" || cosmicSignatureGameAddr.length !== 42) {
        console.log("CADDR environment variable does not contain contract address of CosmicGame contract.");
        process.exit(1);
    }

    const cosmicSignatureGame = await hre.ethers.getContractAt(cosmicSignatureGameContractName, cosmicSignatureGameAddr);

    return cosmicSignatureGame;
}
async function getERC20SampleContracts(sampContractName = "Samp") {
	// gets address of dummy token contract (for donation testing)
    const sampContract1Addr = process.env.TSAMP1;
    const sampContract2Addr = process.env.TSAMP2;
    if (typeof sampContract1Addr === "undefined" || sampContract1Addr.length !== 42) {
        console.log("TSAMP1 environment variable does not contain contract address of token sample contract.");
        process.exit(1);
    }
    if (typeof sampContract2Addr === "undefined" || sampContract2Addr.length !== 42) {
        console.log("TSAMP2 environment variable does not contain contract address of token sample contract.");
        process.exit(1);
    }

    const sampContract1 = await hre.ethers.getContractAt(sampContractName, sampContract1Addr);
    const sampContract2 = await hre.ethers.getContractAt(sampContractName, sampContract2Addr);

    return [sampContract1,sampContract2];
}
async function getRandomWalkNft(game) {

	
}

async function main() {
    async function mint_rwalk(randomWalkNFT,a) {
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
			await stakingWalletCosmicSignatureNft.connect(owner_signer).unstake(i,10);
			num_unstaked=num_unstaked+1;
		}
	}
	[owner, addr1, addr2, addr3, addr4, addr5] = await customGetSigners();
    const cosmicGameProxy = await getCosmicSignatureGameContract()
	const [ samp1,samp2 ] = await getERC20SampleContracts();
	const rwalkAddr = await cosmicGameProxy.randomWalkNft();
	const randomWalkNFT = await ethers.getContractAt("RandomWalkNFT",rwalkAddr);
	const stakingWalletCstAddr = await cosmicGameProxy.stakingWalletCosmicSignatureNft();
	const stakingWalletCst = await ethers.getContractAt("StakingWalletCosmicSignatureNft",stakingWalletCstAddr);
	const stakingWalletRWalkAddr = await cosmicGameProxy.stakingWalletRandomWalkNft();
	const stakingWalletRandomWalkNft = await ethers.getContractAt("StakingWalletRandomWalkNft",stakingWalletRWalkAddr);

	let token_id;
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
    console.log("Donation complete");

	let numStakeActions = 5
    for (let i = 0; i < numStakeActions; i++) {
        let token_id = await mint_rwalk(randomWalkNFT,addr1);
        await randomWalkNFT
            .connect(addr1)
            .setApprovalForAll(await stakingWalletRandomWalkNft.getAddress(), true);
		let txdata = stakingWalletRandomWalkNft.interface.encodeFunctionData("stake",[token_id]);
		const tx = await addr1.sendTransaction({
			to: await stakingWalletRandomWalkNft.getAddress(),
			data: txdata,
			gasLimit: 500000
		});
		await tx.wait();
        //await stakingWalletRandomWalkNft.connect(addr1).stake(token_id);
    }
    for (let i = 0; i < 5; i++) {
        let token_id = await mint_rwalk(randomWalkNFT,addr2);
        await randomWalkNFT
            .connect(addr2)
            .setApprovalForAll(await stakingWalletRandomWalkNft.getAddress(), true);
		let txdata = stakingWalletRandomWalkNft.interface.encodeFunctionData("stake",[token_id]);
		const tx = await addr2.sendTransaction({
			to: await stakingWalletRandomWalkNft.getAddress(),
			data: txdata,
			gasLimit: 500000
		});
		await tx.wait();
        //await stakingWalletRandomWalkNft.connect(addr2).stake(token_id);
    }
    for (let i = 0; i < 15; i++) {
        let token_id = await mint_rwalk(randomWalkNFT,addr3);
        await randomWalkNFT
            .connect(addr3)
            .setApprovalForAll(await stakingWalletRandomWalkNft.getAddress(), true);
		let txdata = stakingWalletRandomWalkNft.interface.encodeFunctionData("stake",[token_id]);
		const tx = await addr3.sendTransaction({
			to: await stakingWalletRandomWalkNft.getAddress(),
			data: txdata,
			gasLimit: 500000
		});
		await tx.wait();
        //await stakingWalletRandomWalkNft.connect(addr3).stake(token_id);
    }
    let prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();

    const contractBalance = await ethers.provider.getBalance(
        await cosmicGameProxy.getAddress()
    );
    let bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	let txdata = cosmicGameProxy.interface.encodeFunctionData("bidWithEth",[-1,"bid 1"]);
	let tx = await addr1.sendTransaction({
		to: await cosmicGameProxy.getAddress(),
		data: txdata,
		gasLimit: 500000,
		value: bidPrice + 1000n,
	});
	await tx.wait();
    //await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 1", { value: bidPrice + 1000n }); // this works
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr2).bidWithEth(-1,"bid 1", { value: bidPrice + 1000n }); // this works

//    let nanoSecondsExtra = await cosmicGameProxy.nanoSecondsExtra();
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 2", { value: bidPrice });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr1).bidWithEth(-1, "bid 2", { value: bidPrice });
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    token_id = await mint_rwalk(randomWalkNFT,owner);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(owner).bidWithEth(Number(token_id),"bidWithRWlk", {value: bidPrice });

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {
        value: bidPrice
    });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {
        value: bidPrice
    });

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr4).bidWithEth(-1,"", { value: bidPrice });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr4).bidWithEth(-1,"", { value: bidPrice });

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr5).bidWithEth(-1,"", {value: bidPrice });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
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
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
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
    prizeAmount2 = await cosmicGameProxy.getMainEthPrizeAmount();
	await stake_available_nfts();
    let ts = await cosmicSignature.totalSupply();
    let rn = await cosmicGameProxy.roundNum();
    let oldTotalSupply = ts;

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
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
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp1.getAddress(),10000000000000000000n,{value:bidPrice});
	await samp2.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await samp2.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp2.getAddress(),10000000000000000000n,{value:bidPrice});

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
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

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    token_id = await mint_rwalk(randomWalkNFT,addr1);
	await randomWalkNFT.connect(addr1).setApprovalForAll(await prizesWallet.getAddress(), true);
    await cosmicGameProxy
        .connect(addr1)
        .bidWithEthAndDonateNft(-1,"donated token_id="+token_id,await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    token_id = await mint_rwalk(randomWalkNFT,addr2);
	await randomWalkNFT.connect(addr2).setApprovalForAll(await prizesWallet.getAddress(), true);
    await cosmicGameProxy
        .connect(addr2)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	await randomWalkNFT.connect(addr3).setApprovalForAll(await prizesWallet.getAddress(), true);
    await cosmicGameProxy
        .connect(addr3)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	await randomWalkNFT.connect(addr3).setApprovalForAll(await prizesWallet.getAddress(), true);
    await cosmicGameProxy
        .connect(addr3)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
    await ethers.provider.send("evm_increaseTime", [36000]);
    await ethers.provider.send("evm_mine");
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	let cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    await cosmicGameProxy
        .connect(addr3)
        .bidWithCstAndDonateNft(cstPrice,"cst bid + donate1", await randomWalkNFT.getAddress(), token_id);
    await ethers.provider.send("evm_increaseTime", [36000]);
    await ethers.provider.send("evm_mine");
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    await cosmicGameProxy
        .connect(addr3)
        .bidWithCstAndDonateNft(cstPrice,"cst bid + donate2", await randomWalkNFT.getAddress(), token_id);
    await ethers.provider.send("evm_increaseTime", [36000]);
    await ethers.provider.send("evm_mine");
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"bid using ERC20 token");
    await ethers.provider.send("evm_mine");
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
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
	prizesWallet.connect(addr3).claimDonatedToken(rn,await samp1.getAddress());
	prizesWallet.connect(addr3).claimDonatedToken(rn,await samp2.getAddress());
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
    await cosmicGameProxy.connect(owner).setCstRewardAmountMultiplier(5);
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
    await cosmicGameProxy.connect(owner).upgradeTo(iAddr);
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletCosmicSignatureNft(await stakingWalletCosmicSignatureNft.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletRandomWalkNft(await stakingWalletRandomWalkNft.getAddress());
    await cosmicGameProxy.connect(owner).setCstRewardAmountMultiplier(999);
    let tmp = await cosmicGameProxy.mainPrizeTimeIncrementIncreaseDivisor();
    await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementIncreaseDivisor(tmp);
    tmp = await cosmicGameProxy.connect(owner).timeoutDurationToClaimMainPrize();
    await cosmicGameProxy.connect(owner).setTimeoutDurationToClaimMainPrize(tmp);
    tmp = await cosmicGameProxy.nextEthBidPriceIncreaseDivisor();
    await cosmicGameProxy.connect(owner).setNextEthBidPriceIncreaseDivisor(tmp);
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

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr1).bidWithEth(-1,"", { value: bidPrice  });
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
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

    await marketingWallet.payReward(addr1.address,hre.ethers.parseEther("7"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.payReward(addr3.address,hre.ethers.parseEther("7"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.payReward(addr2.address,hre.ethers.parseEther("2"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.payReward(addr1.address,hre.ethers.parseEther("6"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.payReward(addr2.address,hre.ethers.parseEther("5"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.payReward(addr2.address,hre.ethers.parseEther("5"));
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.payReward(addr3.address,hre.ethers.parseEther("5"));
    await marketingWallet.payReward(addr4.address,hre.ethers.parseEther("1"));
    await marketingWallet.payReward(addr1.address,hre.ethers.parseEther("11"));
	await unstake_all_nfts()
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
	await stakingWalletCosmicSignatureNft.tryPerformMaintenance(true,owner.address)
	await stake_available_nfts()

	await samp1.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await samp1.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp1.getAddress(),11000000000000000000n,{value:bidPrice});
	await samp2.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await samp2.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp2.getAddress(),11000000000000000000n,{value:bidPrice});

    // generate one deposit to charity and not to Staking Wallet
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3",{value: bidPrice});
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {value: bidPrice });
	rn = cosmicGameProxy.roundNum();
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });

	prizesWallet.connect(addr3).claimDonatedToken(rn,await samp1.getAddress());	// only claim one of the tokens (samp1, not samp2)
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
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice });
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    await cosmicGameProxy.connect(addr1).bidWithCst(cstPrice,"CST bid addr1");
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 3000000
    });

    donationAmount = hre.ethers.parseEther("500");
    await cosmicGameProxy.connect(addr3).donateEth({
        value: donationAmount
    });
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice });
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"CST bid addr1");
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice });
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
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
	numStakeActions = await stakingWalletCosmicSignatureNft.actionCounter();
    for (let i = 1; i <= numStakeActions; i++) {
        let action_rec = (await stakingWalletCosmicSignatureNft.stakeActions(i)).toObject();
        let ownr = action_rec.nftOwnerAddress;
		if (ownr == "0x0000000000000000000000000000000000000000") { continue; }
		if (ownr == addr1.address) {
			// make some address keeping its tokens
			continue;
		} else {
		}
        let owner_signer = await hre.ethers.getSigner(ownr);
		try {
	        await stakingWalletCosmicSignatureNft.connect(owner_signer).unstake(i,1);
	        await stakingWalletCosmicSignatureNft.connect(owner_signer).payReward(i,100);
		} catch (e) {
		//	console.log("unstake() error: ",e);
		}
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
